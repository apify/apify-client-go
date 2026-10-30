package apify

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	// emptyLogStreamRetry is how long a [StreamedLog] waits before it reopens a log stream that
	// delivered nothing: the API serves the log of a run that has not logged anything yet as an
	// empty stream that ends at once. The pause doubles up to emptyLogStreamMaxRetry, which bounds
	// the request rate while a run waits long to start. Matches the reference client's
	// EMPTY_LOG_STREAM_RETRY_MILLIS and EMPTY_LOG_STREAM_MAX_RETRY_MILLIS.
	emptyLogStreamRetry    = 500 * time.Millisecond
	emptyLogStreamMaxRetry = 5 * time.Second
	// stopGrace is how long [StreamedLog.Stop] waits for the messages already in flight before it
	// aborts the log stream. Matches the reference client's STOP_GRACE_MILLIS.
	stopGrace = time.Second
	// logRedirectPrefixSeparator ends the default message prefix, which names the Actor and the run.
	logRedirectPrefixSeparator = " -> "
	// logReadChunkBytes is the size of the buffer the log stream is read through.
	logReadChunkBytes = 32 * 1024
)

// logMessageMarker matches the ISO timestamp that starts every message of an Actor run log. The
// log is split into messages at these markers, so a message that spans several lines (a stack
// trace, say) is redirected as one.
var logMessageMarker = regexp.MustCompile(`(?:\n|^)(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z)`)

// StreamedLogOptions configures [RunClient.NewStreamedLog] and, through
// [ActorStartOptions.Log], the log redirection of [ActorClient.Call].
type StreamedLogOptions struct {
	// Disabled turns the log redirection off. It only has an effect in [ActorStartOptions.Log].
	Disabled bool
	// Output receives every redirected message, one per line, prefixed with Prefix. Defaults to
	// os.Stderr. Writes are serialized, so the writer need not be safe for concurrent use.
	Output io.Writer
	// Prefix is written in front of every message. Defaults to "{Actor name} runId:{run ID} -> ",
	// which costs two requests (the run and its Actor); a non-nil Prefix (an empty one included)
	// skips them.
	Prefix *string
	// FromStart controls whether messages logged before the redirection was created are
	// redirected too (default true). With false, they are dropped.
	FromStart *bool
}

// StreamedLog redirects the live log of an Actor run to a writer, message by message. Obtain one
// from [RunClient.NewStreamedLog], begin the redirection with [StreamedLog.Start] and end it with
// [StreamedLog.Stop]. Matches the reference client's StreamedLog.
//
// Each message of the log, which starts with the timestamp the platform puts in front of it, is
// written on its own line, prefixed. Messages that were already logged when the redirection was
// created are skipped if FromStart is false. The log of a run that has not logged anything yet is
// reopened until it delivers its first bytes.
type StreamedLog struct {
	log       *LogClient
	output    io.Writer
	prefix    string
	startedAt time.Time
	fromStart bool

	mu           sync.Mutex // serializes writes to output and guards the fields below
	running      bool
	stopping     bool
	startCtx     context.Context // the context of the Start call, which ends the redirection
	stopCtx      context.Context // the context of the Stop call, which bounds the final read
	cancelStream context.CancelFunc
	wake         chan struct{}
	done         chan error
	buffer       bytes.Buffer
}

// NewStreamedLog creates the redirection of this run's log, which does nothing until
// [StreamedLog.Start]. With the default prefix it fetches the run and its Actor to name them, and
// the bool reports whether the run exists: it is false (with a nil [StreamedLog]) for a run that
// does not. A non-nil [StreamedLogOptions.Prefix] skips the lookups, so the bool is then always
// true.
func (c *RunClient) NewStreamedLog(ctx context.Context, options StreamedLogOptions) (*StreamedLog, bool, error) {
	prefix := ""
	if options.Prefix != nil {
		prefix = *options.Prefix
	} else {
		run, present, err := getResourceWithTimeout[ActorRun](ctx, c.ctx, "", NewQueryParams(), c.ctx.longTimeout())
		if err != nil || !present {
			return nil, false, err
		}
		actorName := ""
		if run.ActID != "" {
			actorClient := c.root.Actor(run.ActID)
			actor, found, err := getResourceWithTimeout[Actor](ctx, actorClient.ctx, "", NewQueryParams(), actorClient.ctx.longTimeout())
			if err != nil {
				return nil, false, err
			}
			if found {
				actorName = actor.Name
			}
		}
		parts := []string{"runId:" + run.ID}
		if actorName != "" {
			parts = append([]string{actorName}, parts...)
		}
		prefix = strings.Join(parts, " ") + logRedirectPrefixSeparator
	}

	output := options.Output
	if output == nil {
		output = os.Stderr
	}
	return &StreamedLog{
		log:       c.Log(),
		output:    output,
		prefix:    prefix,
		startedAt: time.Now(),
		fromStart: options.FromStart == nil || *options.FromStart,
	}, true, nil
}

// Start begins the redirection in the background. The redirection also ends when ctx is done.
// It returns an error if the redirection is already running.
func (s *StreamedLog) Start(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.running {
		return errors.New("StreamedLog: already started")
	}
	streamCtx, cancel := context.WithCancel(ctx)
	s.running = true
	s.stopping = false
	s.startCtx = ctx
	s.stopCtx = nil
	s.cancelStream = cancel
	s.wake = make(chan struct{})
	s.done = make(chan error, 1)
	s.buffer.Reset()
	go func() {
		s.done <- s.streamLog(ctx, streamCtx)
	}()
	return nil
}

// Stop ends the redirection. It gives the messages already in flight up to one second, then
// aborts the log stream; if the stream delivered nothing yet, it reads the whole log in one
// request instead, so a short run that finished before the stream connected is not lost. It
// returns the error that made the redirection stop early, if any. ctx bounds the wait for it,
// the final read of the whole log included: when ctx is done first, Stop gives up on the rest of
// the log and returns the error of ctx. It returns an error if the redirection is not running or
// is already being stopped.
func (s *StreamedLog) Stop(ctx context.Context) error {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return errors.New("StreamedLog: not started")
	}
	if s.stopping {
		s.mu.Unlock()
		return errors.New("StreamedLog: already stopping")
	}
	s.stopping = true
	s.stopCtx = ctx
	close(s.wake)
	cancel, done, startCtx := s.cancelStream, s.done, s.startCtx
	s.mu.Unlock()

	timer := time.AfterFunc(stopGrace, cancel)
	defer timer.Stop()
	var err error
	select {
	case err = <-done:
	case <-ctx.Done():
		select {
		case err = <-done: // it finished as the context ended
		default:
			// The redirection is still running: cutting it off loses the rest of the log, which
			// the caller must be able to tell from a complete one.
			cancel()
			if err = <-done; err == nil && startCtx.Err() == nil {
				err = ctx.Err()
			}
		}
	}
	cancel()
	// A read cut short by the context of this call reports a plain cancellation: name the cause,
	// unless the context of Start, which ends the redirection on purpose, ended as well.
	if ctxErr := ctx.Err(); ctxErr != nil && errors.Is(err, context.Canceled) && startCtx.Err() == nil {
		err = ctxErr
	}

	s.mu.Lock()
	s.running = false
	s.mu.Unlock()
	return err
}

// finalReadContext returns the context of the final read of the whole log: it ends with the
// context of the redirection and with the one of the Stop call.
func (s *StreamedLog) finalReadContext(parent context.Context) (context.Context, context.CancelFunc) {
	s.mu.Lock()
	stopCtx := s.stopCtx
	s.mu.Unlock()
	readCtx, cancel := context.WithCancel(parent)
	stopWatching := context.AfterFunc(stopCtx, cancel)
	return readCtx, func() {
		stopWatching()
		cancel()
	}
}

// logReadError wraps the failure of the final read of the whole log, which leaves the log
// unredirected. Unlike a stream that is cut on purpose, it is never swallowed.
type logReadError struct{ err error }

func (e *logReadError) Error() string { return e.err.Error() }
func (e *logReadError) Unwrap() error { return e.err }

// isContextError reports whether err is the cancellation or the deadline of a context.
func isContextError(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

func (s *StreamedLog) isStopping() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stopping
}

// streamLog redirects the log until the stream ends for good. parent is the context of the whole
// redirection, and streamCtx the one of the log stream, which Stop cancels after its grace period.
// It returns nil when the redirection was ended on purpose.
func (s *StreamedLog) streamLog(parent, streamCtx context.Context) error {
	err := s.redirect(parent, streamCtx)
	var lost *logReadError
	if errors.As(err, &lost) && (!isContextError(lost.err) || parent.Err() == nil) {
		// The log was not redirected: the caller must hear of it, unless the cause is the end of
		// the caller's own context, which ends the redirection on purpose.
		return lost.err
	}
	if err != nil && streamCtx.Err() != nil && isContextError(err) {
		return nil // ended on purpose: by Stop, or because the caller's context is done
	}
	return err
}

func (s *StreamedLog) redirect(parent, streamCtx context.Context) error {
	raw := true
	options := LogOptions{Raw: &raw}
	retry := emptyLogStreamRetry
	gotBytes := false

	for !gotBytes {
		if s.isStopping() {
			// The stream's context ends with the grace period, which must not cut this read
			// short: it ends with the context of the caller's Stop instead.
			readCtx, cancelRead := s.finalReadContext(parent)
			body, _, err := s.log.GetBufferWithOptions(readCtx, options)
			cancelRead()
			if err != nil {
				return &logReadError{err: err}
			}
			s.consume(body)
			break
		}
		stream, err := s.log.StreamWithOptions(streamCtx, options)
		if err != nil {
			if s.isStopping() && parent.Err() == nil {
				continue // the grace period ended while connecting: read the log in one request
			}
			return err
		}
		if stream == nil {
			return nil // the log does not exist
		}
		if s.isStopping() {
			_ = stream.Close()
			continue
		}
		var readErr error
		gotBytes, readErr = s.consumeStream(stream)
		_ = stream.Close()
		if readErr != nil && streamCtx.Err() == nil {
			return readErr
		}
		if !gotBytes {
			select {
			case <-time.After(retry):
			case <-s.wake:
			case <-streamCtx.Done():
			}
			retry = min(retry*2, emptyLogStreamMaxRetry)
		}
	}
	s.flush()
	return nil
}

// consumeStream redirects the complete messages of stream and reports whether it delivered any
// byte. It stops at the end of the stream, an error, or when the redirection is being stopped.
func (s *StreamedLog) consumeStream(stream io.Reader) (bool, error) {
	chunk := make([]byte, logReadChunkBytes)
	gotBytes := false
	for {
		n, err := stream.Read(chunk)
		if n > 0 {
			gotBytes = true
			s.consume(chunk[:n])
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				return gotBytes, nil
			}
			return gotBytes, err
		}
		if s.isStopping() {
			return gotBytes, nil
		}
	}
}

// consume adds data to the buffer and redirects every complete message in it: everything before
// the last end of line. An incomplete last line stays in the buffer for the next chunk.
func (s *StreamedLog) consume(data []byte) {
	s.buffer.Write(data)
	buffered := s.buffer.Bytes()
	lastEOL := bytes.LastIndexByte(buffered, '\n')
	if lastEOL < 0 {
		return
	}
	complete := string(buffered[:lastEOL])
	rest := append([]byte(nil), buffered[lastEOL:]...)
	s.buffer.Reset()
	s.buffer.Write(rest)
	s.redirectMessages(complete)
}

// flush redirects whatever is left in the buffer when the log ends: an incomplete last message,
// or a last line without an end of line.
func (s *StreamedLog) flush() {
	last := strings.TrimSpace(s.buffer.String())
	s.buffer.Reset()
	if last != "" {
		s.write(last)
	}
}

// redirectMessages splits the complete part of the log into messages at their timestamps and
// writes them.
func (s *StreamedLog) redirectMessages(text string) {
	markers := logMessageMarker.FindAllStringSubmatchIndex(text, -1)
	for i, marker := range markers {
		contentEnd := len(text)
		if i+1 < len(markers) {
			contentEnd = markers[i+1][0]
		}
		timestamp := text[marker[2]:marker[3]]
		if !s.fromStart {
			if logged, err := time.Parse(time.RFC3339Nano, timestamp); err == nil && logged.Before(s.startedAt) {
				continue
			}
		}
		s.write(strings.TrimSpace(timestamp + text[marker[3]:contentEnd]))
	}
}

// write writes one prefixed message to the output.
func (s *StreamedLog) write(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, _ = fmt.Fprintf(s.output, "%s%s\n", s.prefix, message)
}

// startLogRedirection starts redirecting the log of run, as [ActorClient.Call] does by default.
// A failure never fails the call: it is reported to the redirection's output and the call goes on
// without redirected logs. It returns nil when the redirection is disabled or could not start.
func startLogRedirection(ctx context.Context, run *RunClient, options *StreamedLogOptions) *StreamedLog {
	settings := StreamedLogOptions{}
	if options != nil {
		settings = *options
	}
	if settings.Disabled {
		return nil
	}
	streamed, found, err := run.NewStreamedLog(ctx, settings)
	if err == nil && found {
		err = streamed.Start(ctx)
	}
	if err != nil {
		reportLogRedirectionError(settings.Output, err)
		return nil
	}
	if !found {
		return nil // the run is gone: there is no log to redirect
	}
	return streamed
}

// stopLogRedirection ends the redirection started by [startLogRedirection], if there is one.
func stopLogRedirection(ctx context.Context, streamed *StreamedLog, options *StreamedLogOptions) {
	if streamed == nil {
		return
	}
	if err := streamed.Stop(ctx); err != nil {
		var output io.Writer
		if options != nil {
			output = options.Output
		}
		reportLogRedirectionError(output, err)
	}
}

// reportLogRedirectionError tells the user that the log redirection ended early.
func reportLogRedirectionError(output io.Writer, err error) {
	if output == nil {
		output = os.Stderr
	}
	_, _ = fmt.Fprintf(output, "Log redirection stopped due to error: %v\n", err)
}
