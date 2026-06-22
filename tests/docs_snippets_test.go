package apify_test

// This file verifies the documentation requirement that "each in-documentation code snippet
// has to be a valid, runnable and properly formatted code." Go has no built-in doctest
// mechanism for Markdown (unlike Rust's `cargo test --doc`), so this test plays that role: it
// extracts every fenced ```go code block from the README and the docs/ pages, then for each
// snippet
//
//   - checks it is gofmt-formatted (proper formatting), and
//   - compiles it with `go build` (valid, runnable code).
//
// A snippet that begins with `package ` is treated as a complete program and built as-is;
// every other snippet is a fragment that assumes a configured `client` and a `ctx`, so it is
// wrapped in a function with those (and a few other doc-wide placeholders) predeclared before
// being compiled. The test is offline — it never executes the snippets against the API — so it
// runs without APIFY_TOKEN and is exercised by the standalone "Test examples" CI step.

import (
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// docFiles are the documentation files whose ```go snippets must compile and be formatted,
// relative to the module root (the parent of this tests/ directory).
var docFiles = []string{
	"README.md",
	"docs/README.md",
	"docs/actors.md",
	"docs/builds.md",
	"docs/runs.md",
	"docs/tasks.md",
	"docs/storages.md",
	"docs/schedules.md",
	"docs/webhooks.md",
	"docs/misc.md",
}

// goFenceRe matches a fenced ```go ... ``` block, capturing the snippet body.
var goFenceRe = regexp.MustCompile("(?s)```go\\n(.*?)\\n```")

// topLevelShortVarRe matches the left-hand side of a `:=` short variable declaration that
// sits at the outermost level of a fragment (column 0 — fragments are authored unindented, so
// nested declarations carry leading tabs). We emit a blank assignment for each such name to
// avoid "declared and not used" errors when a doc fragment introduces a variable only to
// illustrate a call. Nested declarations are intentionally skipped: their scope is the inner
// block, so a function-level discard would not even refer to them.
var topLevelShortVarRe = regexp.MustCompile(`(?m)^([a-zA-Z_][\w]*(?:\s*,\s*[a-zA-Z_][\w]*)*)\s*:=`)

// declaresName reports whether a fragment declares a given identifier itself (via `:=`, `=`,
// or `var`), in which case the wrapper must not also predeclare it.
func declaresName(body, name string) bool {
	// `name :=` / `name =` (also matches the identifier as part of a multi-name LHS).
	assign := regexp.MustCompile(`(?m)(^|[\s,(])` + regexp.QuoteMeta(name) + `\b\s*:?=`)
	if assign.MatchString(body) {
		return true
	}
	// `var name ...`
	return regexp.MustCompile(`(?m)\bvar\s+` + regexp.QuoteMeta(name) + `\b`).MatchString(body)
}

type snippet struct {
	file string
	body string
}

// collectSnippets extracts all ```go fenced blocks from the documentation files.
func collectSnippets(t *testing.T, root string) []snippet {
	t.Helper()
	var out []snippet
	for _, rel := range docFiles {
		data, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			t.Fatalf("read %s: %v", rel, err)
		}
		for _, m := range goFenceRe.FindAllStringSubmatch(string(data), -1) {
			out = append(out, snippet{file: rel, body: m[1]})
		}
	}
	if len(out) == 0 {
		t.Fatal("no ```go snippets found in documentation; extractor or docs changed")
	}
	return out
}

// assembleProgram turns a doc snippet into a self-contained, compilable Go source file.
//
// Full-program snippets (those that declare their own package) are used verbatim. Fragments
// are wrapped in a synthetic main package: the snippet runs inside an anonymous function with
// the doc-wide placeholders (client, ctx, and a few common IDs) predeclared, and every
// short-variable the fragment introduces is discarded with a trailing blank assignment so the
// fragment compiles without "declared and not used" noise.
func assembleProgram(s snippet) string {
	if strings.HasPrefix(strings.TrimSpace(s.body), "package ") {
		return s.body
	}

	declared := map[string]bool{}
	var discards strings.Builder
	for _, m := range topLevelShortVarRe.FindAllStringSubmatch(s.body, -1) {
		for _, name := range strings.Split(m[1], ",") {
			name = strings.TrimSpace(name)
			if name == "" || name == "_" || declared[name] {
				continue
			}
			declared[name] = true
			discards.WriteString("\t_ = " + name + "\n")
		}
	}

	var b strings.Builder
	b.WriteString("package main\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"io\"\n")
	b.WriteString("\t\"log\"\n")
	b.WriteString("\t\"net/http\"\n")
	b.WriteString("\t\"os\"\n")
	b.WriteString("\t\"time\"\n\n")
	b.WriteString("\tapify \"github.com/apify/apify-client-go\"\n")
	b.WriteString(")\n\n")
	// Reference the imports that a given fragment may not use, so the wrapper never fails on
	// unused imports regardless of which snippet it carries.
	b.WriteString("var (\n")
	b.WriteString("\t_ = fmt.Sprint\n")
	b.WriteString("\t_ = io.EOF\n")
	b.WriteString("\t_ = log.Print\n")
	b.WriteString("\t_ = http.MethodGet\n")
	b.WriteString("\t_ = os.Stdout\n")
	b.WriteString("\t_ = time.Second\n")
	b.WriteString(")\n\n")
	b.WriteString("func snippet() {\n")
	// Doc-wide placeholders documented in docs/README.md: snippets assume a configured client
	// and a context, plus a handful of resource IDs used by the per-resource pages. Each is
	// predeclared only when the fragment does not declare it itself, so a snippet that opens
	// with `client := apify.NewClient(...)` is not shadowed.
	placeholders := []struct{ name, decl string }{
		{"client", "var client *apify.ApifyClient"},
		{"ctx", "var ctx context.Context"},
		{"actorID", "var actorID string"},
		{"runID", "var runID string"},
		{"buildID", "var buildID string"},
	}
	for _, p := range placeholders {
		if declaresName(s.body, p.name) {
			continue
		}
		b.WriteString("\t" + p.decl + "\n")
		b.WriteString("\t_ = " + p.name + "\n")
	}
	b.WriteString("\n")
	b.WriteString(s.body)
	b.WriteString("\n")
	b.WriteString(discards.String())
	b.WriteString("}\n\n")
	b.WriteString("func main() { snippet() }\n")
	return b.String()
}

func TestDocSnippetsFormatted(t *testing.T) {
	root := ".."
	for _, s := range collectSnippets(t, root) {
		// gofmt is whitespace-sensitive; the snippet body in the doc must already be the
		// canonical gofmt output. We compare against gofmt of the body itself (parsed as a
		// fragment via a minimal wrapper) to keep the check local to the snippet.
		formatted, err := gofmtFragment(s.body)
		if err != nil {
			t.Errorf("%s: snippet does not parse: %v\n%s", s.file, err, s.body)
			continue
		}
		if formatted != s.body {
			t.Errorf("%s: snippet is not gofmt-formatted.\n--- have ---\n%s\n--- want ---\n%s",
				s.file, s.body, formatted)
		}
	}
}

func TestDocSnippetsCompile(t *testing.T) {
	root := ".."
	dir := t.TempDir()

	// A throwaway module that depends on the local client via a replace directive, so the
	// snippets compile against the exact source in this checkout.
	abs, err := filepath.Abs(root)
	if err != nil {
		t.Fatalf("abs root: %v", err)
	}
	gomod := "module docsnippets\n\ngo 1.23\n\nrequire github.com/apify/apify-client-go v0.0.0\n\nreplace github.com/apify/apify-client-go => " + abs + "\n"
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(gomod), 0o644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}

	snips := collectSnippets(t, root)
	for i, s := range snips {
		prog := assembleProgram(s)
		sub := filepath.Join(dir, fmt.Sprintf("snip%03d", i))
		if err := os.MkdirAll(sub, 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(filepath.Join(sub, "main.go"), []byte(prog), 0o644); err != nil {
			t.Fatalf("write snippet program: %v", err)
		}
	}

	// Build every assembled snippet program in one pass.
	cmd := exec.Command("go", "build", "./...")
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOFLAGS=-mod=mod")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("doc snippets failed to compile:\n%s", out)
	}
}

// gofmtFragment formats a snippet body using the same rules gofmt applies to a file. Fragments
// (which are not valid standalone files) are wrapped in a function, formatted, then unwrapped,
// so the comparison reflects exactly what an author would get from `gofmt` on the snippet.
func gofmtFragment(body string) (string, error) {
	trimmed := strings.TrimSpace(body)
	if strings.HasPrefix(trimmed, "package ") {
		out, err := gofmtSource(body)
		return strings.TrimRight(out, "\n"), err
	}

	wrapped := "package p\n\nfunc _f() {\n" + indent(body) + "\n}\n"
	out, err := gofmtSource(wrapped)
	if err != nil {
		return "", err
	}
	return dedentFuncBody(out), nil
}

// gofmtSource formats a complete source file with the same engine gofmt uses
// (go/format.Source), in-process — no subprocess and no separate parse step (Source reports a
// parse error itself).
func gofmtSource(src string) (string, error) {
	out, err := format.Source([]byte(src))
	return string(out), err
}

// indent adds one tab to every non-empty line so a fragment sits correctly inside a function
// body before formatting.
func indent(s string) string {
	lines := strings.Split(s, "\n")
	for i, ln := range lines {
		if strings.TrimSpace(ln) != "" {
			lines[i] = "\t" + ln
		}
	}
	return strings.Join(lines, "\n")
}

// dedentFuncBody extracts the body of the synthetic `_f` function produced by gofmtFragment and
// removes the single leading tab gofmt added, recovering the canonical formatting of the
// original fragment.
func dedentFuncBody(formatted string) string {
	lines := strings.Split(formatted, "\n")
	var body []string
	in := false
	for _, ln := range lines {
		if !in {
			if strings.HasPrefix(ln, "func _f() {") {
				in = true
			}
			continue
		}
		if ln == "}" {
			break
		}
		body = append(body, strings.TrimPrefix(ln, "\t"))
	}
	// Drop leading/trailing blank lines introduced by the wrapper.
	for len(body) > 0 && strings.TrimSpace(body[0]) == "" {
		body = body[1:]
	}
	for len(body) > 0 && strings.TrimSpace(body[len(body)-1]) == "" {
		body = body[:len(body)-1]
	}
	return strings.Join(body, "\n")
}
