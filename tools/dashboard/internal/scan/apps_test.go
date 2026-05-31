package scan

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

// buildWorkspace creates a .jobops workspace with the given postings and
// application sub-paths, and returns a loaded Config.
func buildWorkspace(t *testing.T) *config.Config {
	t.Helper()
	root := t.TempDir()
	mk := func(p string) { _ = os.MkdirAll(filepath.Join(root, p), 0o755) }
	wr := func(p, body string) {
		_ = os.MkdirAll(filepath.Dir(filepath.Join(root, p)), 0o755)
		_ = os.WriteFile(filepath.Join(root, p), []byte(body), 0o644)
	}
	mk(".jobops")
	wr(".jobops/config.json", `{"directories":{"job_postings":"./Job_Postings","applications_root":"./Applications","company_intelligence":"./Company_Intelligence"}}`)

	// Backlog posting (no Applications folder).
	wr("Job_Postings/Beta_Capital_Analyst_20260101.md", "# JD\n")
	// In-progress application: assessed (84%) + resume final, no cover yet.
	wr("Job_Postings/Acme_Product_Manager_20260531.md", "# JD\n")
	wr("Applications/Acme_Product_Manager_20260531/job_posting.md", "# pinned\n")
	wr("Applications/Acme_Product_Manager_20260531/assessment/assessment.md",
		"---\noutput_type: assessment\nstatus: final\nnormalized_score: 84%\n---\n")
	wr("Applications/Acme_Product_Manager_20260531/resume/step3_final.md", "# resume\n")
	// Company intelligence for Acme (drives OSINT stage + Company detection).
	mk("Company_Intelligence/Acme")

	cfg, err := config.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return cfg
}

func findRec(recs []model.Record, slug string) *model.Record {
	for i := range recs {
		if recs[i].Slug == slug {
			return &recs[i]
		}
	}
	return nil
}

func stageState(r *model.Record, name string) model.StageState {
	for _, s := range r.Stages {
		if s.Name == name {
			return s.State
		}
	}
	return model.Missing
}

func TestAppsScanUnionAndStages(t *testing.T) {
	cfg := buildWorkspace(t)
	recs, err := AppsAdapter{Cfg: cfg}.Scan()
	if err != nil {
		t.Fatal(err)
	}
	if len(recs) != 2 {
		t.Fatalf("want 2 records (backlog + started), got %d", len(recs))
	}

	beta := findRec(recs, "Beta_Capital_Analyst_20260101")
	if beta == nil || beta.Started {
		t.Fatalf("Beta should be a backlog (unstarted) record: %+v", beta)
	}

	acme := findRec(recs, "Acme_Product_Manager_20260531")
	if acme == nil || !acme.Started {
		t.Fatalf("Acme should be started")
	}
	if acme.Score == nil || *acme.Score != 84 {
		t.Errorf("Acme score = %v, want 84", acme.Score)
	}
	if acme.Company != "Acme" {
		t.Errorf("Acme company = %q, want Acme", acme.Company)
	}
	if stageState(acme, "Assess") != model.Final {
		t.Errorf("Assess should be Final")
	}
	if stageState(acme, "Resume") != model.Final {
		t.Errorf("Resume should be Final")
	}
	if stageState(acme, "OSINT") != model.Final {
		t.Errorf("OSINT should be Final (Company_Intelligence/Acme exists)")
	}
	if stageState(acme, "Cover") != model.Missing {
		t.Errorf("Cover should be Missing")
	}
	if acme.NextSkill != "coverletter" {
		t.Errorf("NextSkill = %q, want coverletter", acme.NextSkill)
	}
	if acme.Title != "Acme Product Manager" || acme.Date != "20260531" {
		t.Errorf("title/date = %q / %q", acme.Title, acme.Date)
	}
}

func TestResumeDraftState(t *testing.T) {
	cfg := buildWorkspace(t)
	// Add an application with only step1 (draft, no step3).
	app := filepath.Join(cfg.ApplicationsDir(), "Delta_Corp_Director_20260201")
	_ = os.MkdirAll(filepath.Join(app, "resume"), 0o755)
	_ = os.WriteFile(filepath.Join(app, "resume", "step1_draft.md"), []byte("# draft\n"), 0o644)

	recs, _ := AppsAdapter{Cfg: cfg}.Scan()
	d := findRec(recs, "Delta_Corp_Director_20260201")
	if d == nil {
		t.Fatal("Delta record missing")
	}
	if stageState(d, "Resume") != model.Draft {
		t.Errorf("Resume should be Draft with only step1, got %v", stageState(d, "Resume"))
	}
	if d.NextSkill != "assessjob" {
		t.Errorf("NextSkill = %q, want assessjob (first incomplete)", d.NextSkill)
	}
}

func TestReadScoreOverallFallback(t *testing.T) {
	dir := t.TempDir()
	adir := filepath.Join(dir, "assessment")
	if err := os.MkdirAll(adir, 0o755); err != nil {
		t.Fatal(err)
	}
	// overall_score only, no normalized_score: 150/200 -> 75%.
	body := "---\noutput_type: assessment\noverall_score: 150/200\n---\n"
	if err := os.WriteFile(filepath.Join(adir, "assessment.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	got := readScore(dir)
	if got == nil || *got != 75 {
		t.Errorf("readScore overall fallback = %v, want 75", got)
	}
}
