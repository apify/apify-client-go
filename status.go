package apify

// ActorJobStatus is the lifecycle status of an Actor run or build.
//
// It is a named string type so the fixed set of statuses the API can report is discoverable
// and comparable against the exported constants below, while still round-tripping to/from the
// raw JSON string. The values mirror the Apify platform's ACTOR_JOB_STATUSES.
type ActorJobStatus string

// The Actor job statuses reported by the Apify API for runs and builds.
const (
	// ActorJobStatusReady means the job is created but not yet started.
	ActorJobStatusReady ActorJobStatus = "READY"
	// ActorJobStatusRunning means the job is currently running.
	ActorJobStatusRunning ActorJobStatus = "RUNNING"
	// ActorJobStatusSucceeded means the job finished successfully (terminal).
	ActorJobStatusSucceeded ActorJobStatus = "SUCCEEDED"
	// ActorJobStatusFailed means the job failed (terminal).
	ActorJobStatusFailed ActorJobStatus = "FAILED"
	// ActorJobStatusTimingOut means the job is in the process of timing out.
	ActorJobStatusTimingOut ActorJobStatus = "TIMING-OUT"
	// ActorJobStatusTimedOut means the job timed out (terminal).
	ActorJobStatusTimedOut ActorJobStatus = "TIMED-OUT"
	// ActorJobStatusAborting means the job is in the process of being aborted.
	ActorJobStatusAborting ActorJobStatus = "ABORTING"
	// ActorJobStatusAborted means the job was aborted (terminal).
	ActorJobStatusAborted ActorJobStatus = "ABORTED"
)

// terminalActorJobStatuses is the set of statuses in which a run or build is finished and its
// status will not change further.
var terminalActorJobStatuses = map[ActorJobStatus]struct{}{
	ActorJobStatusSucceeded: {},
	ActorJobStatusFailed:    {},
	ActorJobStatusAborted:   {},
	ActorJobStatusTimedOut:  {},
}

// IsTerminal reports whether the status is a terminal (finished) status: the run or build has
// stopped and its status will not change further.
func (s ActorJobStatus) IsTerminal() bool {
	_, ok := terminalActorJobStatuses[s]
	return ok
}
