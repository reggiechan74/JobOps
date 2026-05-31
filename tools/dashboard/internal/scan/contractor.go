package scan

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

// ContractorAdapter lists jobops-ic outputs under contractor_root and offers the
// IC skills (which live in the /jobops-ic namespace).
type ContractorAdapter struct {
	Cfg *config.Config
}

var _ model.Adapter = ContractorAdapter{}

// contractorSubdirs are the fixed output folders created by /jobops-ic:setup.
var contractorSubdirs = []string{"services", "prospects", "proposals", "pitches", "rate-cards", "landing-pages"}

func (a ContractorAdapter) Name() string { return "Contractor" }

func (a ContractorAdapter) Skills() []model.SkillSpec {
	const p = "jobops-ic"
	return []model.SkillSpec{
		{Name: "defineservices", Label: "Define services", Arg: model.ArgNone, Plugin: p},
		{Name: "findclient", Label: "Find client", Arg: model.ArgNone, Plugin: p},
		{Name: "pitchdeck", Label: "Pitch deck", Arg: model.ArgNone, Plugin: p},
		{Name: "proposaltemplate", Label: "Proposal", Arg: model.ArgNone, Plugin: p},
		{Name: "ratecard", Label: "Rate card", Arg: model.ArgNone, Plugin: p},
		{Name: "create-landing-page", Label: "Landing page", Arg: model.ArgNone, Plugin: p},
	}
}

// Scan lists .md/.html outputs across the known contractor subfolders.
func (a ContractorAdapter) Scan() ([]model.Record, error) {
	root := a.Cfg.ContractorDir()
	var recs []model.Record
	for _, sub := range contractorSubdirs {
		d := filepath.Join(root, sub)
		entries, err := os.ReadDir(d)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() || isHidden(e.Name()) {
				continue
			}
			ext := filepath.Ext(e.Name())
			if ext != ".md" && ext != ".html" {
				continue
			}
			stem := strings.TrimSuffix(e.Name(), ext)
			var mt time.Time
			if info, ierr := e.Info(); ierr == nil {
				mt = info.ModTime()
			}
			recs = append(recs, model.Record{
				Title:   sub + " / " + stem,
				Slug:    sub + "/" + stem,
				Paths:   []string{filepath.Join(d, e.Name())},
				Updated: mt,
			})
		}
	}
	sortByUpdatedDesc(recs)
	return recs, nil
}
