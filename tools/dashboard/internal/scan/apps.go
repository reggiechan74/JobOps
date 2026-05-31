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

// AppsAdapter implements model.Adapter for the Apps tab: the union of
// Job_Postings backlog and in-progress Applications.
type AppsAdapter struct {
	Cfg *config.Config
}

func (a AppsAdapter) Name() string { return "Apps" }

// Skills returns the launch palette for the Apps tab, in pipeline order.
func (a AppsAdapter) Skills() []model.SkillSpec {
	return []model.SkillSpec{
		{Name: "auditjobposting", Label: "Audit JD", Arg: model.ArgJD},
		{Name: "assessjob", Label: "Assess fit", Arg: model.ArgJD},
		{Name: "buildresume", Label: "Build resume", Arg: model.ArgJD},
		{Name: "coverletter", Label: "Cover letter", Arg: model.ArgJD},
		{Name: "osint", Label: "Company OSINT", Arg: model.ArgCompany},
		{Name: "briefing", Label: "Interview briefing", Arg: model.ArgJD},
		{Name: "interviewprep", Label: "Interview prep", Arg: model.ArgJD},
		{Name: "convert-to-pdf", Label: "Convert to PDF", Arg: model.ArgNone},
		{Name: "convert-to-word", Label: "Convert to Word", Arg: model.ArgNone},
	}
}

// pipelineOrder is the ordered list of derived stages and the skill that
// produces each. "JD audit" is intentionally not a derived stage.
var pipelineOrder = []struct{ name, skill string }{
	{"Assess", "assessjob"},
	{"Resume", "buildresume"},
	{"Cover", "coverletter"},
	{"OSINT", "osint"},
	{"Briefing", "briefing"},
	{"Prep", "interviewprep"},
}

// Scan unions Job_Postings stems and Applications folders into records.
func (a AppsAdapter) Scan() ([]model.Record, error) {
	slugs := map[string]bool{}

	if entries, err := os.ReadDir(a.Cfg.JobPostingsDir()); err == nil {
		for _, e := range entries {
			name := e.Name()
			if e.IsDir() || !strings.HasSuffix(name, ".md") || strings.HasPrefix(name, ".") {
				continue
			}
			slugs[strings.TrimSuffix(name, ".md")] = true
		}
	}
	if entries, err := os.ReadDir(a.Cfg.ApplicationsDir()); err == nil {
		for _, e := range entries {
			if e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
				slugs[e.Name()] = true
			}
		}
	}

	var recs []model.Record
	for slug := range slugs {
		recs = append(recs, a.buildRecord(slug))
	}
	sort.Slice(recs, func(i, j int) bool {
		if recs[i].Updated.Equal(recs[j].Updated) {
			return recs[i].Slug < recs[j].Slug
		}
		return recs[i].Updated.After(recs[j].Updated)
	})
	return recs, nil
}

func (a AppsAdapter) buildRecord(slug string) model.Record {
	appDir := filepath.Join(a.Cfg.ApplicationsDir(), slug)
	started := dirExists(appDir)
	title, date := SplitSlug(slug)
	company := DetectCompany(slug, a.Cfg.CompanyIntelDir())

	rec := model.Record{
		Title:   title,
		Date:    date,
		Slug:    slug,
		Company: company,
		Started: started,
	}
	if title == "" {
		rec.Title = slug
		rec.Warn = "unparseable slug"
	}

	tr, _ := ReadTracker(appDir)
	rec.Lifecycle = tr.Lifecycle

	rec.Stages = a.deriveStages(appDir, company)
	for _, s := range rec.Stages {
		if s.State != model.Final {
			rec.NextSkill = s.Skill
			break
		}
	}
	rec.Score = readScore(appDir)
	rec.Updated = latestMtime(appDir)
	rec.Paths = []string{appDir}
	return rec
}

func (a AppsAdapter) deriveStages(appDir, company string) []model.Stage {
	stages := make([]model.Stage, 0, len(pipelineOrder))
	for _, p := range pipelineOrder {
		stages = append(stages, model.Stage{Name: p.name, Skill: p.skill, State: a.stageState(appDir, company, p.name)})
	}
	return stages
}

func (a AppsAdapter) stageState(appDir, company, stage string) model.StageState {
	switch stage {
	case "Assess":
		return existState(filepath.Join(appDir, "assessment", "assessment.md"))
	case "Resume":
		if fileExists(filepath.Join(appDir, "resume", "step3_final.md")) {
			return model.Final
		}
		if fileExists(filepath.Join(appDir, "resume", "step1_draft.md")) ||
			fileExists(filepath.Join(appDir, "resume", "step2_provenance.md")) {
			return model.Draft
		}
		return model.Missing
	case "Cover":
		return existState(filepath.Join(appDir, "cover-letter", "cover_letter.md"))
	case "OSINT":
		if company != "" && dirExists(filepath.Join(a.Cfg.CompanyIntelDir(), company)) {
			return model.Final
		}
		return model.Missing
	case "Briefing":
		return globState(filepath.Join(appDir, "interview"), "briefing")
	case "Prep":
		return globState(filepath.Join(appDir, "interview"), "interview_prep")
	}
	return model.Missing
}

// readScore reads the normalized fit % from assessment.md, falling back to
// computing it from overall_score "N/200".
func readScore(appDir string) *int {
	path := filepath.Join(appDir, "assessment", "assessment.md")
	if !fileExists(path) {
		return nil
	}
	var m AppsMatter
	if _, err := ParseFrontmatter(path, &m); err != nil {
		return nil
	}
	if n, ok := leadingInt(m.NormalizedScore); ok {
		return &n
	}
	if num, ok := leadingInt(m.OverallScore); ok {
		if i := strings.Index(m.OverallScore, "/"); i >= 0 {
			if den, ok2 := leadingInt(m.OverallScore[i+1:]); ok2 && den > 0 {
				pct := int(float64(num)/float64(den)*100 + 0.5)
				return &pct
			}
		}
	}
	return nil
}

func fileExists(p string) bool {
	info, err := os.Stat(p)
	return err == nil && !info.IsDir()
}

func dirExists(p string) bool {
	info, err := os.Stat(p)
	return err == nil && info.IsDir()
}

func existState(p string) model.StageState {
	if fileExists(p) {
		return model.Final
	}
	return model.Missing
}

// globState returns Final if any file in dir starts with prefix.
func globState(dir, prefix string) model.StageState {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return model.Missing
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasPrefix(e.Name(), prefix) {
			return model.Final
		}
	}
	return model.Missing
}

// latestMtime returns the most recent mtime under dir, or zero time if empty.
func latestMtime(dir string) time.Time {
	var latest time.Time
	_ = filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
		if err == nil && info.Mode().IsRegular() && info.ModTime().After(latest) {
			latest = info.ModTime()
		}
		return nil
	})
	return latest
}
