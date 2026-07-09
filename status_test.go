package apify

import "testing"

func TestActorJobStatusIsTerminal(t *testing.T) {
	terminal := []ActorJobStatus{
		ActorJobStatusSucceeded,
		ActorJobStatusFailed,
		ActorJobStatusAborted,
		ActorJobStatusTimedOut,
	}
	for _, s := range terminal {
		if !s.IsTerminal() {
			t.Errorf("expected %q to be terminal", s)
		}
	}

	nonTerminal := []ActorJobStatus{
		ActorJobStatusReady,
		ActorJobStatusRunning,
		ActorJobStatusTimingOut,
		ActorJobStatusAborting,
	}
	for _, s := range nonTerminal {
		if s.IsTerminal() {
			t.Errorf("expected %q to be non-terminal", s)
		}
	}
}

// TestActorRunIsTerminalUsesStatus verifies the model helper delegates to the status type.
func TestActorRunIsTerminalUsesStatus(t *testing.T) {
	run := ActorRun{Status: ActorJobStatusRunning}
	if run.IsTerminal() {
		t.Fatal("RUNNING run must not be terminal")
	}
	run.Status = ActorJobStatusSucceeded
	if !run.IsTerminal() {
		t.Fatal("SUCCEEDED run must be terminal")
	}
}
