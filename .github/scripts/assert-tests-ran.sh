#!/usr/bin/env bash
# Fail the CI step if a `go test -json` run did not actually execute any test, or if every
# test was skipped. Guards against a silently-green run: a -run/-skip filter that matches
# nothing, a package with no test functions, or an all-skip run all exit 0 otherwise.
#
# Usage: assert-tests-ran.sh <go-test-json-file> <label>
set -euo pipefail

file="${1:?usage: assert-tests-ran.sh <json-file> <label>}"
label="${2:-tests}"

if [ ! -s "$file" ]; then
  echo "::error::${label}: no go test -json output was produced."
  exit 1
fi

# Count per-test terminal events (Test field set). Package-level events have no Test field and
# are ignored here.
passed=$(jq -rs '[.[] | select(.Action == "pass"  and (.Test != null))] | length' "$file")
failed=$(jq -rs '[.[] | select(.Action == "fail"  and (.Test != null))] | length' "$file")
skipped=$(jq -rs '[.[] | select(.Action == "skip" and (.Test != null))] | length' "$file")
ran=$((passed + failed))

echo "${label}: ${ran} test(s) ran (${passed} passed, ${failed} failed), ${skipped} skipped."

if [ "$ran" -eq 0 ]; then
  echo "::error::${label}: no tests ran (all skipped or none matched the filter); failing to avoid a silent green run."
  exit 1
fi
