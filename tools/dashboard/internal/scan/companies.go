package scan

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

// CompaniesAdapter implements the dashboard scanner for the Companies tab: one
// record per Company_Intelligence/{Company}/ folder, artifact-only (no lifecycle).
type CompaniesAdapter struct {
	Cfg *config.Config
}

var _ model.Adapter = CompaniesAdapter{}

func (a CompaniesAdapter) Name() string { return "Companies" }

// Skills returns the launch palette for the Companies tab.
func (a CompaniesAdapter) Skills() []model.SkillSpec {
	return []model.SkillSpec{
		{Name: "osint", Label: "Company OSINT", Arg: model.ArgCompany},
	}
}

// osintFiles is the canonical 7-file OSINT output set, in display order.
var osintFiles = []string{"corporate", "legal", "leadership", "compensation", "culture", "market", "summary"}

// Scan returns one record per subdirectory of Company_Intelligence/.
func (a CompaniesAdapter) Scan() ([]model.Record, error) {
	dir := a.Cfg.CompanyIntelDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, nil // missing root = empty tab, not an error
	}

	var recs []model.Record
	for _, e := range entries {
		if !e.IsDir() || isHidden(e.Name()) {
			continue
		}
		company := e.Name()
		companyDir := filepath.Join(dir, company)

		stages := make([]model.Stage, 0, len(osintFiles))
		next := ""
		for _, f := range osintFiles {
			st := existState(filepath.Join(companyDir, f+".md"))
			stages = append(stages, model.Stage{Name: f, Skill: "osint", State: st})
			if st != model.Final && next == "" {
				next = "osint"
			}
		}

		recs = append(recs, model.Record{
			Title:     company,
			Company:   company,
			Slug:      company,
			Stages:    stages,
			NextSkill: next,
			Updated:   latestMtime(companyDir),
			Paths:     []string{companyDir},
		})
	}
	sort.Slice(recs, func(i, j int) bool {
		if recs[i].Updated.Equal(recs[j].Updated) {
			return recs[i].Slug < recs[j].Slug
		}
		return recs[i].Updated.After(recs[j].Updated)
	})
	return recs, nil
}

// isHidden reports whether a directory entry name is a dotfile.
func isHidden(name string) bool {
	return len(name) > 0 && name[0] == '.'
}
