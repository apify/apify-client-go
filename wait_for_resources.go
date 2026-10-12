package apify

import (
	"context"
	"fmt"
	"time"
)

// resourceLimitErrorTypes are the Apify API error types a run/task start is rejected with while
// the account has no free memory or concurrent-run slot for it. Both clear as other runs or
// builds of the account finish. Matches the reference client's RESOURCE_LIMIT_ERROR_TYPES.
var resourceLimitErrorTypes = map[string]bool{
	"actor-memory-limit-exceeded":    true,
	"concurrent-runs-limit-exceeded": true,
}

// waitForResourcesCooldown is the delay between two attempts to start a run rejected for lack of
// resources. Matches the reference client's WAIT_FOR_RESOURCES_COOLDOWN_MILLIS.
const waitForResourcesCooldown = 10 * time.Second

// WaitForResourcesIndefinitely, used as [ActorStartOptions.WaitForResources] or
// [TaskStartOptions.WaitForResources], retries a resource-limited start forever instead of
// bounding the retrying to a number of seconds. Matches the reference client's
// `waitForResources: true`.
const WaitForResourcesIndefinitely = int64(-1)

// startWaitingForResources calls start, retrying it every waitForResourcesCooldown while it
// fails with a resource-limit error (`actor-memory-limit-exceeded` or
// `concurrent-runs-limit-exceeded`), until it succeeds or the waitForResources budget elapses.
// Any other error is returned immediately, without retrying.
//
// waitForResources nil disables retrying, so the first error (of any kind) is returned
// immediately. [WaitForResourcesIndefinitely] retries with no time bound. Any other value is the
// number of seconds to keep retrying for (must be >= 0) before returning the last error.
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
