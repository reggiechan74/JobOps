package scan

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFrontmatter(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "assessment.md")
	body := "---\noutput_type: assessment\nstatus: draft\nnormalized_score: 84%\noverall_score: 168/200\n---\n# Body\n"
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	var m AppsMatter
	warn, err := ParseFrontmatter(path, &m)
	if err != nil {
		t.Fatal(err)
	}
	if warn != "" {
		t.Errorf("unexpected warn %q", warn)
	}
	if m.Status != "draft" || m.NormalizedScore != "84%" || m.OverallScore != "168/200" {
		t.Errorf("parsed = %+v", m)
	}
}

func TestParseFrontmatterMalformed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.md")
	if err := os.WriteFile(path, []byte("no frontmatter here\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	var m AppsMatter
	warn, err := ParseFrontmatter(path, &m)
	if err != nil {
		t.Fatalf("should not hard-error: %v", err)
	}
	if warn == "" {
		t.Errorf("expected a warn for missing frontmatter")
	}
}

func TestLeadingInt(t *testing.T) {
	cases := map[string]struct {
		n  int
		ok bool
	}{"84%": {84, true}, "168/200": {168, true}, "": {0, false}, "B+": {0, false}}
	for in, want := range cases {
		n, ok := leadingInt(in)
		if n != want.n || ok != want.ok {
			t.Errorf("leadingInt(%q) = (%d,%v), want (%d,%v)", in, n, ok, want.n, want.ok)
		}
	}
}
