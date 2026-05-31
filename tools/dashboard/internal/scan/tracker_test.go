package scan

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func TestReadTrackerMissingDefaults(t *testing.T) {
	tr, err := ReadTracker(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if tr.Lifecycle != model.Interested {
		t.Errorf("missing tracker should default to Interested, got %q", tr.Lifecycle)
	}
}

func TestWriteThenReadTracker(t *testing.T) {
	dir := t.TempDir()
	in := Tracker{Lifecycle: model.Interviewing, AppliedOn: "2026-05-20", Notes: "panel 06-03"}
	if err := WriteTracker(dir, in); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".tracker.json")); err != nil {
		t.Fatalf(".tracker.json not written: %v", err)
	}
	out, err := ReadTracker(dir)
	if err != nil {
		t.Fatal(err)
	}
	if out.Lifecycle != model.Interviewing || out.AppliedOn != "2026-05-20" || out.Notes != "panel 06-03" {
		t.Errorf("round-trip mismatch: %+v", out)
	}
	if out.UpdatedAt == "" {
		t.Errorf("WriteTracker should stamp UpdatedAt")
	}
}
