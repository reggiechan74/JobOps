package scan

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func TestFlatRecords(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Ideal_Job_Profile_20260101.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	_ = os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("x"), 0o644)
	_ = os.WriteFile(filepath.Join(dir, ".hidden.md"), []byte("x"), 0o644)

	recs := flatRecords(dir)
	if len(recs) != 1 {
		t.Fatalf("want 1 record, got %d", len(recs))
	}
	if recs[0].Title != "Ideal Job Profile" || recs[0].Date != "20260101" {
		t.Errorf("record = %q / %q", recs[0].Title, recs[0].Date)
	}
	if len(recs[0].Stages) != 0 || recs[0].Score != nil || recs[0].Lifecycle != "" {
		t.Errorf("flat record should have no stages/score/lifecycle: %+v", recs[0])
	}
}

func TestFlatRecordsMissingDir(t *testing.T) {
	if recs := flatRecords(filepath.Join(t.TempDir(), "nope")); recs != nil {
		t.Errorf("missing dir should yield nil, got %v", recs)
	}
}

func flatConfig(t *testing.T, key, rel string) *config.Config {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, ".jobops")
	_ = os.MkdirAll(dir, 0o755)
	_ = os.WriteFile(filepath.Join(dir, "config.json"),
		[]byte(`{"directories":{"`+key+`":"`+rel+`"}}`), 0o644)
	_ = os.MkdirAll(filepath.Join(root, rel), 0o755)
	_ = os.WriteFile(filepath.Join(root, rel, "Output_20260101.md"), []byte("x"), 0o644)
	cfg, err := config.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return cfg
}

func TestCareerAdapter(t *testing.T) {
	cfg := flatConfig(t, "career_analysis", "Career_Analysis")
	recs, err := CareerAdapter{Cfg: cfg}.Scan()
	if err != nil {
		t.Fatal(err)
	}
	if len(recs) != 1 {
		t.Fatalf("want 1 career record, got %d", len(recs))
	}
	if got := (CareerAdapter{}).Skills(); len(got) != 3 || got[0].Name != "idealjob" {
		t.Errorf("career skills = %+v", got)
	}
	var _ model.Adapter = CareerAdapter{}
}

func TestCrisisAdapter(t *testing.T) {
	cfg := flatConfig(t, "crisis_management", "Crisis_Management")
	recs, err := CrisisAdapter{Cfg: cfg}.Scan()
	if err != nil {
		t.Fatal(err)
	}
	if len(recs) != 1 {
		t.Fatalf("want 1 crisis record, got %d", len(recs))
	}
	skills := CrisisAdapter{}.Skills()
	if len(skills) != 11 {
		t.Errorf("want 11 crisis skills, got %d", len(skills))
	}
}
