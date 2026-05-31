package scan

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func contractorWorkspace(t *testing.T) *config.Config {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, ".jobops")
	_ = os.MkdirAll(dir, 0o755)
	_ = os.WriteFile(filepath.Join(dir, "config.json"),
		[]byte(`{"directories":{"contractor_root":"./Contractor"}}`), 0o644)
	wr := func(p string) {
		full := filepath.Join(root, "Contractor", p)
		_ = os.MkdirAll(filepath.Dir(full), 0o755)
		_ = os.WriteFile(full, []byte("x"), 0o644)
	}
	wr("services/svc.md")
	wr("rate-cards/rates.md")
	wr("landing-pages/page.html")
	wr("services/ignore.txt")
	cfg, err := config.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return cfg
}

func TestContractorScan(t *testing.T) {
	cfg := contractorWorkspace(t)
	recs, err := ContractorAdapter{Cfg: cfg}.Scan()
	if err != nil {
		t.Fatal(err)
	}
	if len(recs) != 3 {
		t.Fatalf("want 3 contractor records, got %d: %+v", len(recs), recs)
	}
	titles := map[string]bool{}
	for _, r := range recs {
		titles[r.Title] = true
	}
	for _, want := range []string{"services / svc", "rate-cards / rates", "landing-pages / page"} {
		if !titles[want] {
			t.Errorf("missing record titled %q (got %v)", want, titles)
		}
	}
}

func TestContractorSkills(t *testing.T) {
	skills := ContractorAdapter{}.Skills()
	if len(skills) != 6 {
		t.Fatalf("want 6 contractor skills, got %d", len(skills))
	}
	for _, s := range skills {
		if s.Plugin != "jobops-ic" {
			t.Errorf("skill %q has Plugin %q, want jobops-ic", s.Name, s.Plugin)
		}
	}
	var _ model.Adapter = ContractorAdapter{}
}
