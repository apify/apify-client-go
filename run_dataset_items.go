package apify

import (
	"context"
	"encoding/json"
	"errors"
)

const (
	// defaultRunIterateChunkSize is the number of dataset rows requested per API call when
	// [RunIterateDatasetItemsOptions.ChunkSize] is unset. Matches the reference client's
	// DEFAULT_ITERATE_CHUNK_SIZE.
	defaultRunIterateChunkSize = int64(1000)
	// defaultRunIteratePollIntervalSecs is how long the iterator waits for the run to finish
	// between polls when [RunIterateDatasetItemsOptions.PollIntervalSecs] is unset.
	defaultRunIteratePollIntervalSecs = int64(5)
)

// errRunIterateDesc is returned when the items options ask for newest-first order.
var errRunIterateDesc = errors.New("RunIterateDatasetItemsOptions: Items.Desc is not supported: the iterator reads rows by offset from the oldest, which a descending order would count from the newest")

// RunIterateDatasetItemsOptions configures [RunClient.IterateDatasetItems] and
// [IterateRunDatasetItems].
type RunIterateDatasetItemsOptions struct {
	// Items are the filters and projections applied to every page, the same as for
	// [DatasetClient.ListItems]. Items.Offset is the number of rows to skip from the beginning
	// (default 0), and Items.Limit the maximum number of dataset rows to scan (default no limit):
	// fewer items are yielded when filters drop some, more when Unwind splits a row into several.
	// Items.Desc is not supported, because the rows are read by offset from the oldest.
	Items DatasetListItemsOptions
	// ChunkSize is the maximum number of dataset rows requested per API call (default 1000).
	ChunkSize *int64
	// PollIntervalSecs is how long to wait for the run to finish between polls, in seconds
	// (default 5).
	PollIntervalSecs *int64
}

// runIteratePhase is the stage a [RunDatasetItemsIterator] is in.
type runIteratePhase int

const (
	// runIterateStart: nothing has been requested yet.
	runIterateStart runIteratePhase = iota
	// runIterateLive: the run may still be producing items. Each round reads the rows below the
	// dataset's itemCount, then waits for the run to finish.
	runIterateLive
	// runIterateTail: the run has finished. The rows past itemCount (which lags a few seconds
	// behind the pushed items) are read a page at a time until none are left.
	runIterateTail
	// runIterateDone: the iterator is exhausted.
	runIterateDone
)

// RunDatasetItemsIterator iterates over the items of a run's default dataset while the run is
// still producing them. Obtain one from [RunClient.IterateDatasetItems] or
// [IterateRunDatasetItems] and drain it by calling Next until it returns (nil, nil).
//
// While the run has not finished, each round yields the rows below the dataset's itemCount and
// then waits up to PollIntervalSecs for the run to finish, so the last rows are read as soon as it
// does. Each page is requested with a limit that ends at itemCount, so it covers exactly the
// rows it asks for, whatever the filters or Unwind do to the items. itemCount lags a few seconds
// behind the pushed items, so once the run reaches a terminal status the rows past it are read a
// page at a time until none are left. On a client from LastRun, the iterator sticks to the run
// its first request resolves to. Matches the reference client's iterateDatasetItems().
type RunDatasetItemsIterator[T any] struct {
	run     *RunClient
	options RunIterateDatasetItemsOptions

	pageSize int64
	pollSecs int64
	position int64
	end      int64
	hasEnd   bool

	phase      runIteratePhase
	dataset    *DatasetClient
	runClient  *RunClient
	currentRun ActorRun
	runPresent bool
	// roundFinished is whether the run was already finished when the current live round began.
	roundFinished bool
	// roundItemCount is the number of rows the current live round reads up to.
	roundItemCount int64
	roundStarted   bool

	buffer []T
	pos    int
}

// IterateRunDatasetItems returns an iterator over the items of the run's default dataset, each
// decoded into T, that follows the dataset while the run is still producing items. See
// [RunDatasetItemsIterator].
func IterateRunDatasetItems[T any](c *RunClient, options RunIterateDatasetItemsOptions) *RunDatasetItemsIterator[T] {
	it := &RunDatasetItemsIterator[T]{
		run:      c,
		options:  options,
		pageSize: defaultRunIterateChunkSize,
		pollSecs: defaultRunIteratePollIntervalSecs,
		position: offsetVal(options.Items.Offset),
	}
	if options.ChunkSize != nil && *options.ChunkSize > 0 {
		it.pageSize = *options.ChunkSize
	}
	if options.PollIntervalSecs != nil && *options.PollIntervalSecs >= 0 {
		it.pollSecs = *options.PollIntervalSecs
	}
	if options.Items.Limit != nil && *options.Items.Limit > 0 {
		it.hasEnd = true
		it.end = it.position + *options.Items.Limit
	}
	return it
}

// IterateDatasetItems is [IterateRunDatasetItems] with every item decoded into a
// json.RawMessage. See [RunDatasetItemsIterator].
func (c *RunClient) IterateDatasetItems(options RunIterateDatasetItemsOptions) *RunDatasetItemsIterator[json.RawMessage] {
	return IterateRunDatasetItems[json.RawMessage](c, options)
}

// Next returns the next item, or (nil, nil) once every item has been read (or the Limit of
// scanned rows is reached). It calls the API only when the current in-memory page is used up.
func (it *RunDatasetItemsIterator[T]) Next(ctx context.Context) (*T, error) {
	for it.pos >= len(it.buffer) {
		if it.phase == runIterateDone {
			return nil, nil
		}
		if err := it.advance(ctx); err != nil {
			return nil, err
		}
	}
	item := it.buffer[it.pos]
	it.pos++
	return &item, nil
}

// advance performs the next step of the iteration, which loads the next page into the buffer, or
// moves to the next phase.
func (it *RunDatasetItemsIterator[T]) advance(ctx context.Context) error {
	it.buffer, it.pos = nil, 0
	switch it.phase {
	case runIterateStart:
		return it.start(ctx)
	case runIterateLive:
		return it.liveStep(ctx)
	case runIterateTail:
		return it.tailStep(ctx)
	default:
		return nil
	}
}

// start resolves the run and the client of its default dataset.
func (it *RunDatasetItemsIterator[T]) start(ctx context.Context) error {
	if it.options.Items.Desc != nil && *it.options.Items.Desc {
		return errRunIterateDesc
	}
	it.runClient = it.run
	run, present, err := getResourceWithTimeout[ActorRun](ctx, it.run.ctx, "", NewQueryParams(), it.run.ctx.longTimeout())
	if err != nil {
		return err
	}
	it.currentRun, it.runPresent = run, present
	// A LastRun client resolves runs/last on each request, so a newer run would swap the dataset
	// mid-iteration: pin the iteration to the run the first request resolved to.
	if present && run.ID != it.run.id {
		it.runClient = it.run.root.Run(run.ID)
	}
	it.dataset = it.runClient.Dataset()
	it.phase = runIterateLive
	return nil
}

// liveStep reads one page of the current live round, or ends the round: it waits for the run to
// finish and starts another one, or moves on to the tail.
func (it *RunDatasetItemsIterator[T]) liveStep(ctx context.Context) error {
	if !it.roundStarted {
		it.roundFinished = !it.runPresent || it.currentRun.IsTerminal()
		dataset, present, err := getResourceWithTimeout[Dataset](ctx, it.dataset.ctx, "", NewQueryParams(), it.dataset.ctx.longTimeout())
		if err != nil {
			return err
		}
		it.roundItemCount = 0
		if present {
			it.roundItemCount = dataset.ItemCount
		}
		if it.hasEnd {
			it.roundItemCount = minInt64(it.roundItemCount, it.end)
		}
		it.roundStarted = true
	}

	if it.position < it.roundItemCount {
		pageLimit := minInt64(it.pageSize, it.roundItemCount-it.position)
		items, _, err := it.listPage(ctx, it.position, pageLimit)
		if err != nil {
			return err
		}
		it.buffer = items
		it.position += pageLimit
		return nil
	}

	it.roundStarted = false
	if it.hasEnd && it.position >= it.end {
		it.phase = runIterateDone
		return nil
	}
	if it.roundFinished {
		it.phase = runIterateTail
		return nil
	}
	run, err := it.runClient.WaitForFinish(ctx, &it.pollSecs)
	if err != nil {
		return err
	}
	it.currentRun, it.runPresent = run, true
	return nil
}

// tailStep reads one page past itemCount after the run has finished. Only an empty page ends the
// iteration, since filters can shorten a full one.
func (it *RunDatasetItemsIterator[T]) tailStep(ctx context.Context) error {
	pageLimit := it.pageSize
	if it.hasEnd {
		pageLimit = minInt64(it.pageSize, it.end-it.position)
	}
	items, scanned, err := it.listPage(ctx, it.position, pageLimit)
	if err != nil {
		return err
	}
	it.buffer = items

	// A page that Clean, SkipEmpty or Unwind emptied past a lagging itemCount reports no scanned
	// rows either, so a plain read confirms the dataset really has no more rows.
	if len(items) == 0 && scanned == 0 {
		confirmed, err := it.confirmNoMoreRows(ctx)
		if err != nil {
			return err
		}
		if confirmed {
			it.phase = runIterateDone
			return nil
		}
	}
	it.position += pageLimit
	if it.hasEnd && it.position >= it.end {
		it.phase = runIterateDone
	}
	return nil
}

// confirmNoMoreRows reports whether the dataset has no row at the current position. Without a
// filter that can empty a page, an empty page already means that.
func (it *RunDatasetItemsIterator[T]) confirmNoMoreRows(ctx context.Context) (bool, error) {
	items := it.options.Items
	filtered := (items.Clean != nil && *items.Clean) || (items.SkipEmpty != nil && *items.SkipEmpty) || len(items.Unwind) > 0
	if !filtered {
		return true, nil
	}
	plain, err := ListDatasetItems[json.RawMessage](ctx, it.dataset, DatasetListItemsOptions{Offset: &it.position, Limit: Ptr(int64(1))})
	if err != nil {
		return false, err
	}
	return len(plain.Items) == 0, nil
}

// listPage reads the page of the dataset that starts at offset and holds at most limit rows,
// returning its items and the number of rows the API scanned for it.
func (it *RunDatasetItemsIterator[T]) listPage(ctx context.Context, offset, limit int64) ([]T, int64, error) {
	options := it.options.Items
	options.Offset = &offset
	options.Limit = &limit
	page, err := ListDatasetItems[T](ctx, it.dataset, options)
	if err != nil {
		return nil, 0, err
	}
	return page.Items, page.scanned, nil
}
