package scan

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func companyWorkspace(t *testing.T) *config.Config {
	t.Helper()
	root := t.TempDir()
	wr := func(p, body string) {
		_ = os.MkdirAll(filepath.Dir(filepath.Join(root, p)), 0o755)
		_ = os.WriteFile(filepath.Join(root, p), []byte(body), 0o644)
	}
	wr(".jobops/config.json", `{"directories":{"company_intelligence":"./Company_Intelligence"}}`)
	wr("Company_Intelligence/Acme/corporate.md", "# corp\n")
	wr("Company_Intelligence/Acme/summary.md", "# summary\n")
	for _, f := range []string{"corporate", "legal", "leadership", "compensation", "culture", "market", "summary"} {
		wr("Company_Intelligence/Beta Corp/"+f+".md", "# x\n")
	}
	cfg, err := config.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return cfg
}

func TestCompaniesScan(t *testing.T) {
	cfg := companyWorkspace(t)
	recs, err := CompaniesAdapter{Cfg: cfg}.Scan()
	if err != nil {
		t.Fatal(err)
	}
	if len(recs) != 2 {
		t.Fatalf("want 2 company records, got %d", len(recs))
	}
	acme := findRec(recs, "Acme")
	if acme == nil {
		t.Fatal("Acme missing")
	}
	if acme.Company != "Acme" || acme.Title != "Acme" {
		t.Errorf("Acme title/company = %q/%q", acme.Title, acme.Company)
	}
	if len(acme.Stages) != 7 {
		t.Errorf("want 7 OSINT stages, got %d", len(acme.Stages))
	}
	if stageState(acme, "corporate") != model.Final || stageState(acme, "legal") != model.Missing {
		t.Errorf("Acme stage states wrong: %+v", acme.Stages)
	}
	if acme.NextSkill != "osint" {
		t.Errorf("Acme NextSkill = %q, want osint (files missing)", acme.NextSkill)
	}
	beta := findRec(recs, "Beta Corp")
	if beta == nil {
		t.Fatal("Beta Corp missing")
	}
	if beta.NextSkill != "" {
		t.Errorf("Beta NextSkill = %q, want empty (all 7 present)", beta.NextSkill)
	}
}

func TestCompaniesSkills(t *testing.T) {
	skills := CompaniesAdapter{}.Skills()
	if len(skills) != 1 || skills[0].Name != "osint" || skills[0].Arg != model.ArgCompany {
		t.Errorf("Companies skills = %+v, want single osint/ArgCompany", skills)
	}
}
