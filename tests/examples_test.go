package apify_test

import (
	"os"
	"os/exec"
	"testing"
)

// runExample runs `go run ./examples/<name>` as a subprocess and fails the test if it does
// not exit successfully. Each documentation example is exercised this way, mirroring the
// reference clients' example smoke tests, so the docs stay runnable.
func runExample(t *testing.T, name string) {
	t.Helper()
	if os.Getenv("APIFY_TOKEN") == "" {
		t.Skip("skipping: APIFY_TOKEN is not set")
	}
	cmd := exec.Command("go", "run", "./examples/"+name)
	cmd.Dir = ".." // run from the module root, where examples/ lives
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("example %q failed: %v\n%s", name, err, out)
	}
}

func TestExampleGetAccount(t *testing.T)            { runExample(t, "get_account") }
func TestExampleStorages(t *testing.T)              { runExample(t, "storages") }
func TestExampleRunStoreActor(t *testing.T)         { runExample(t, "run_store_actor") }
func TestExampleRunAndLastRunStorages(t *testing.T) { runExample(t, "run_and_last_run_storages") }
func TestExampleIterateStore(t *testing.T)          { runExample(t, "iterate_store") }
func TestExampleLogRedirection(t *testing.T)        { runExample(t, "log_redirection") }
func TestExampleCreateBuildRunActor(t *testing.T)   { runExample(t, "create_build_run_actor") }
