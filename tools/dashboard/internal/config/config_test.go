package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeConfig(t *testing.T, root string) {
	t.Helper()
	dir := filepath.Join(root, ".jobops")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := `{
  "directories": {
    "resume_source": "./ResumeSourceFolder",
    "job_postings": "./Job_Postings",
    "applications_root": "./Applications",
    "company_intelligence": "./Company_Intelligence",
    "career_analysis": "./Career_Analysis",
    "crisis_management": "./Crisis_Management"
  }
}`
	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestDiscoverFindsAncestor(t *testing.T) {
	root := t.TempDir()
	writeConfig(t, root)
	nested := filepath.Join(root, "a", "b")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	got, err := Discover(nested)
	if err != nil {
		t.Fatal(err)
	}
	if got != root {
		t.Errorf("Discover = %q, want %q", got, root)
	}
}

func TestDiscoverMissing(t *testing.T) {
	if _, err := Discover(t.TempDir()); err == nil {
		t.Errorf("expected error when no .jobops ancestor exists")
	}
}

func TestLoadResolvesDirs(t *testing.T) {
	root := t.TempDir()
	writeConfig(t, root)
	cfg, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(root, "Job_Postings")
	if cfg.JobPostingsDir() != want {
		t.Errorf("JobPostingsDir = %q, want %q", cfg.JobPostingsDir(), want)
	}
	if cfg.ApplicationsDir() != filepath.Join(root, "Applications") {
		t.Errorf("ApplicationsDir = %q", cfg.ApplicationsDir())
	}
}
