package scan

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

// flatRecords turns a directory of timestamped .md outputs into browsable
// records with no pipeline, score, or lifecycle. Returns nil if dir is missing.
func flatRecords(dir string) []model.Record {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var recs []model.Record
	for _, e := range entries {
		if e.IsDir() || isHidden(e.Name()) || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		stem := strings.TrimSuffix(e.Name(), ".md")
		title, date := SplitSlug(stem)
		var mt time.Time
		if info, ierr := e.Info(); ierr == nil {
			mt = info.ModTime()
		}
		recs = append(recs, model.Record{
			Title:   title,
			Date:    date,
			Slug:    stem,
			Paths:   []string{filepath.Join(dir, e.Name())},
			Updated: mt,
		})
	}
	sortByUpdatedDesc(recs)
	return recs
}

// sortByUpdatedDesc sorts records newest-first, breaking ties by Slug.
func sortByUpdatedDesc(recs []model.Record) {
	sort.Slice(recs, func(i, j int) bool {
		if recs[i].Updated.Equal(recs[j].Updated) {
			return recs[i].Slug < recs[j].Slug
		}
		return recs[i].Updated.After(recs[j].Updated)
	})
}

// CareerAdapter lists Career_Analysis/ outputs and offers career-level skills.
type CareerAdapter struct {
	Cfg *config.Config
}

var _ model.Adapter = CareerAdapter{}

func (a CareerAdapter) Name() string { return "Career" }

func (a CareerAdapter) Skills() []model.SkillSpec {
	return []model.SkillSpec{
		{Name: "idealjob", Label: "Ideal job profile", Arg: model.ArgNone},
		{Name: "comparejobs", Label: "Compare jobs", Arg: model.ArgNone},
		{Name: "change-one-thing", Label: "Change one thing", Arg: model.ArgNone},
	}
}

func (a CareerAdapter) Scan() ([]model.Record, error) {
	return flatRecords(a.Cfg.CareerAnalysisDir()), nil
}

// CrisisAdapter lists Crisis_Management/ outputs and offers crisis skills.
type CrisisAdapter struct {
	Cfg *config.Config
}

var _ model.Adapter = CrisisAdapter{}

func (a CrisisAdapter) Name() string { return "Crisis" }

func (a CrisisAdapter) Skills() []model.SkillSpec {
	return []model.SkillSpec{
		{Name: "severance-review", Label: "Severance review", Arg: model.ArgNone},
		{Name: "non-compete-analysis", Label: "Non-compete analysis", Arg: model.ArgNone},
		{Name: "code-red", Label: "Code red", Arg: model.ArgNone},
		{Name: "constructive-dismissal", Label: "Constructive dismissal", Arg: model.ArgNone},
		{Name: "accommodation-request", Label: "Accommodation request", Arg: model.ArgNone},
		{Name: "discrimination-assessment", Label: "Discrimination assessment", Arg: model.ArgNone},
		{Name: "reference-shield", Label: "Reference shield", Arg: model.ArgNone},
		{Name: "unemployment-prep", Label: "Unemployment prep", Arg: model.ArgNone},
		{Name: "workplace-documentation", Label: "Workplace documentation", Arg: model.ArgNone},
		{Name: "layoff-intel", Label: "Layoff intel", Arg: model.ArgNone},
		{Name: "investigation-prep", Label: "Investigation prep", Arg: model.ArgNone},
	}
}

func (a CrisisAdapter) Scan() ([]model.Record, error) {
	return flatRecords(a.Cfg.CrisisManagementDir()), nil
}
