package apify

import (
	"context"
	"fmt"
	"time"
)

// resourceLimitErrorTypes are the error types of a start rejected for lack of memory or concurrent
// run slots.
var resourceLimitErrorTypes = map[string]bool{
	"actor-memory-limit-exceeded":    true,
	"concurrent-runs-limit-exceeded": true,
}

// waitForResourcesCooldown is the delay between attempts to start a run.
const waitForResourcesCooldown = 10 * time.Second

// WaitForResourcesIndefinitely, used as [ActorStartOptions.WaitForResources] or
// [TaskStartOptions.WaitForResources], retries a resource-limited start forever instead of bounding
// the retrying to a number of seconds.
const WaitForResourcesIndefinitely = int64(-1)

// startWaitingForResources calls start again every waitForResourcesCooldown while it fails with a
// resource-limit error, until it succeeds or waitForResources seconds pass. Any other error is
// returned at once. A nil waitForResources does not retry, and [WaitForResourcesIndefinitely]
// retries without a time bound.
func startWaitingForResources[T any](ctx context.Context, waitForResources *int64, start func() (T, error)) (T, error) {
	if waitForResources == nil {
		return start()
	}
	indefinite := *waitForResources == WaitForResourcesIndefinitely
	if !indefinite && *waitForResources < 0 {
		var zero T
		return zero, fmt.Errorf("WaitForResources must be >= 0 or WaitForResourcesIndefinitely, got %d", *waitForResources)
	}
	var deadline time.Time
	if !indefinite {
		deadline = time.Now().Add(time.Duration(*waitForResources) * time.Second)
	}

	for {
		result, err := start()
		if err == nil {
			return result, nil
		}
		apiErr, ok := AsAPIError(err)
		if !ok || !resourceLimitErrorTypes[apiErr.Type] {
			return result, err
		}

		delay := waitForResourcesCooldown
		if !indefinite {
			remaining := time.Until(deadline)
			if remaining <= 0 {
				return result, err
			}
			if remaining < delay {
				delay = remaining
			}
		}
		if !sleepWithContext(ctx, delay) {
			return result, ctx.Err()
		}
	}
}
