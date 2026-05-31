# JobOps Dashboard TUI — Phase 3 Implementation Plan (Career, Crisis, Contractor tabs)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add three flat, artifact-only tabs — **Career** (`Career_Analysis/`), **Crisis** (`Crisis_Management/`), and **Contractor** (`Contractor/`, jobops-ic) — each listing past outputs and launching its skills.

**Architecture:** Two new adapters (`CareerAdapter`, `CrisisAdapter`) share a `flatRecords` helper that turns a directory of timestamped `.md` files into records with no pipeline/lifecycle. `ContractorAdapter` walks the jobops-ic subfolders. Because jobops-ic skills use the `/jobops-ic:` prefix, `SkillSpec` gains a `Plugin` field. Because flat tabs can be empty yet still need to launch a skill, the palette is made openable without a selected record for `ArgNone` skills.

**Tech Stack:** Go 1.22+, existing `internal/{config,model,scan,launch,ui}` packages; depends on Phase 2's multi-tab `ui.Model` + `ui.TabSource`.

**Prerequisite:** Phases 1 and 2 are merged. The `ui.Model` already holds `tabs []tab`, `active int`, and `New(root, agent string, []TabSource)`; render helpers already take a `showLifecycle bool`; `scan` already has `isHidden`, `existState`, `latestMtime`, `SplitSlug`. Work on a branch off `main`.

**Spec:** `docs/superpowers/specs/2026-05-31-jobops-dashboard-tui-design.md` (§3.2 Career/Crisis/Contractor rows).

**Decisions (made while planning):**
- **Flat records carry no stages, score, or lifecycle.** They list past outputs (Title from the humanized filename, Date if the stem ends in an 8-digit token) for browsing; the skills generate *new* outputs and are therefore record-independent (`ArgNone`).
- **`SkillSpec.Plugin`** selects the command namespace (`jobops` default, `jobops-ic` for Contractor). `launch.Compose` uses it.
- **Record-less launching:** for tabs whose skills are `ArgNone`, the palette opens even when the tab has zero records, and `composeSelected` builds the command without a record. (Apps/Companies skills are `ArgJD`/`ArgCompany` and still require a selected record.)
- **Contractor tab is conditional:** shown only when `directories.contractor_root` is configured (jobops-ic may not be installed). Career and Crisis are always shown (their config keys exist from the core `/jobops:setup`).
- **Contractor outputs** include `.md` and `.html` (landing pages); records are titled `"<subfolder> / <stem>"`.

---

## File Structure

```
tools/dashboard/internal/
  model/model.go            # MODIFY: add Plugin to SkillSpec
  launch/launch.go          # MODIFY: Compose uses spec.Plugin
  launch/launch_test.go     # MODIFY: add a plugin test
  config/config.go          # MODIFY: add ContractorRoot + 3 dir helpers
  config/config_test.go     # MODIFY: add a test
  scan/flat.go              # NEW: flatRecords helper + CareerAdapter + CrisisAdapter
  scan/flat_test.go         # NEW
  scan/contractor.go        # NEW: ContractorAdapter
  scan/contractor_test.go   # NEW
  ui/app.go                 # MODIFY: record-less palette for ArgNone tabs
  ui/app_test.go            # MODIFY: add a record-less-palette test
main.go                     # MODIFY: append Career/Crisis/(Contractor) tabs
tools/dashboard/README.md   # MODIFY: document the tabs
```

---

## Task 1: `SkillSpec.Plugin` + namespace-aware Compose

**Files:**
- Modify: `tools/dashboard/internal/model/model.go`
- Modify: `tools/dashboard/internal/launch/launch.go`
- Modify: `tools/dashboard/internal/launch/launch_test.go`

- [ ] **Step 1: Add the failing test**

Append to `tools/dashboard/internal/launch/launch_test.go`:
```go
func TestComposePlugin(t *testing.T) {
	rec := model.Record{Slug: "X_Y_20260101"}
	got := Compose(model.SkillSpec{Name: "ratecard", Arg: model.ArgNone, Plugin: "jobops-ic"}, rec)
	if got != "/jobops-ic:ratecard" {
		t.Errorf("Compose with plugin = %q, want /jobops-ic:ratecard", got)
	}
	// Empty Plugin still defaults to jobops.
	got2 := Compose(model.SkillSpec{Name: "idealjob", Arg: model.ArgNone}, rec)
	if got2 != "/jobops:idealjob" {
		t.Errorf("Compose default plugin = %q, want /jobops:idealjob", got2)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/launch/ -run Plugin`
Expected: FAIL — `Plugin` is not a field of `model.SkillSpec`.

- [ ] **Step 3: Add the `Plugin` field**

In `tools/dashboard/internal/model/model.go`, replace the `SkillSpec` struct with:
```go
// SkillSpec describes one launchable skill in a tab's palette.
type SkillSpec struct {
	Name   string // skill name without the prefix, e.g. "assessjob"
	Label  string // display label in the palette
	Arg    ArgStyle
	Plugin string // command namespace; "" means "jobops" (use "jobops-ic" for IC skills)
}
```

- [ ] **Step 4: Use the plugin in Compose**

In `tools/dashboard/internal/launch/launch.go`, replace the `Compose` function with:
```go
// Compose builds the slash command for running spec against rec. spec.Plugin
// selects the namespace ("" defaults to "jobops").
func Compose(spec model.SkillSpec, rec model.Record) string {
	plugin := spec.Plugin
	if plugin == "" {
		plugin = "jobops"
	}
	base := "/" + plugin + ":" + spec.Name
	switch spec.Arg {
	case model.ArgJD:
		return fmt.Sprintf("%s %s.md --app=%s", base, rec.Slug, rec.Slug)
	case model.ArgCompany:
		if rec.Company != "" {
			return base + " " + rec.Company
		}
		return base
	default:
		return base
	}
}
```

- [ ] **Step 5: Run tests to verify they pass**

Run (from `tools/dashboard/`): `go test ./internal/launch/ ./internal/model/`
Expected: PASS (the existing `TestCompose`/`TestComposeCompanyUnknown` still pass because empty `Plugin` defaults to `jobops`).

- [ ] **Step 6: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/model/model.go tools/dashboard/internal/launch/launch.go tools/dashboard/internal/launch/launch_test.go
git commit -m "feat(dashboard): SkillSpec.Plugin for /jobops-ic namespace"
```

---

## Task 2: Config dir helpers for career/crisis/contractor

**Files:**
- Modify: `tools/dashboard/internal/config/config.go`
- Modify: `tools/dashboard/internal/config/config_test.go`

- [ ] **Step 1: Add the failing test**

Append to `tools/dashboard/internal/config/config_test.go`:
```go
func TestFlatAndContractorDirs(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, ".jobops")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := `{
  "directories": {
    "career_analysis": "./Career_Analysis",
    "crisis_management": "./Crisis_Management",
    "contractor_root": "./Contractor"
  }
}`
	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.CareerAnalysisDir() != filepath.Join(root, "Career_Analysis") {
		t.Errorf("CareerAnalysisDir = %q", cfg.CareerAnalysisDir())
	}
	if cfg.CrisisManagementDir() != filepath.Join(root, "Crisis_Management") {
		t.Errorf("CrisisManagementDir = %q", cfg.CrisisManagementDir())
	}
	if cfg.ContractorDir() != filepath.Join(root, "Contractor") {
		t.Errorf("ContractorDir = %q", cfg.ContractorDir())
	}
	if cfg.Directories.ContractorRoot != "./Contractor" {
		t.Errorf("ContractorRoot field = %q", cfg.Directories.ContractorRoot)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/config/ -run FlatAndContractor`
Expected: FAIL — `ContractorRoot` field and the three methods are undefined.

- [ ] **Step 3: Add the field and helpers**

In `tools/dashboard/internal/config/config.go`, add `ContractorRoot` to the `Directories` struct (after `CrisisManagement`):
```go
		CrisisManagement    string `json:"crisis_management"`
		ContractorRoot      string `json:"contractor_root"`
```
And add these three helper methods next to the existing `JobPostingsDir`/`ApplicationsDir`/`CompanyIntelDir`:
```go
func (c *Config) CareerAnalysisDir() string   { return c.resolve(c.Directories.CareerAnalysis) }
func (c *Config) CrisisManagementDir() string { return c.resolve(c.Directories.CrisisManagement) }
func (c *Config) ContractorDir() string       { return c.resolve(c.Directories.ContractorRoot) }
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/config/`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/config/config.go tools/dashboard/internal/config/config_test.go
git commit -m "feat(dashboard): config helpers for career/crisis/contractor dirs"
```

---

## Task 3: Flat records helper + Career & Crisis adapters

**Files:**
- Create: `tools/dashboard/internal/scan/flat.go`
- Test: `tools/dashboard/internal/scan/flat_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/scan/flat_test.go`:
```go
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
	// Non-.md and hidden files are ignored.
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
	if got := CareerAdapter{}.Skills(); len(got) != 3 || got[0].Name != "idealjob" {
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
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/scan/ -run 'Flat|Career|Crisis'`
Expected: FAIL — undefined `flatRecords`, `CareerAdapter`, `CrisisAdapter`.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/scan/flat.go`:
```go
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
```

> Note: `isHidden` and `SplitSlug` already exist in the `scan` package (from Phase 2's `companies.go` and Phase 1's `slug.go`). Do not redefine them. If `go vet` reports `sortByUpdatedDesc` conflicts with an existing identifier, it means a prior task already added it — in that case remove the duplicate here and reuse the existing one; report it.

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/scan/`
Expected: PASS (all scan tests).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/scan/flat.go tools/dashboard/internal/scan/flat_test.go
git commit -m "feat(dashboard): flatRecords helper + Career & Crisis adapters"
```

---

## Task 4: Contractor adapter

**Files:**
- Create: `tools/dashboard/internal/scan/contractor.go`
- Test: `tools/dashboard/internal/scan/contractor_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/scan/contractor_test.go`:
```go
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
	wr("services/ignore.txt") // ignored (not .md/.html)
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
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/scan/ -run Contractor`
Expected: FAIL — undefined `ContractorAdapter`.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/scan/contractor.go`:
```go
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
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/scan/`
Expected: PASS (all scan tests).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/scan/contractor.go tools/dashboard/internal/scan/contractor_test.go
git commit -m "feat(dashboard): Contractor adapter (jobops-ic outputs)"
```

---

## Task 5: Record-less palette launching for ArgNone tabs

A flat tab with no past outputs has zero records, but the user must still be able to launch its skills. This task lets the palette open without a selected record and compose `ArgNone` commands record-free.

**Files:**
- Modify: `tools/dashboard/internal/ui/app.go`
- Modify: `tools/dashboard/internal/ui/app_test.go`

- [ ] **Step 1: Add the failing test**

Append to `tools/dashboard/internal/ui/app_test.go`:
```go
// emptyScanner returns no records but advertises ArgNone skills.
type emptyScanner struct{ skills []model.SkillSpec }

func (s emptyScanner) Scan() ([]model.Record, error) { return nil, nil }
func (s emptyScanner) Skills() []model.SkillSpec      { return s.skills }

func TestPaletteOpensWithNoRecords(t *testing.T) {
	src := emptyScanner{skills: []model.SkillSpec{{Name: "idealjob", Label: "Ideal", Arg: model.ArgNone}}}
	m := New("/tmp/ws", "claude", []TabSource{{Name: "Career", Scanner: src}})
	m.width, m.height = 100, 30

	// Enter opens the palette even though there are no records.
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	um := updated.(Model)
	if um.mode != modePalette {
		t.Fatalf("palette should open with no records when ArgNone skills exist; mode=%v", um.mode)
	}
	_ = um.View() // must not panic with zero records

	// Selecting the skill composes a record-free command; fallback returns a notice.
	um.agent = "definitely-not-a-real-agent-xyz"
	_, cmd := um.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected a command from ArgNone spawn")
	}
	nm, ok := cmd().(noticeMsg)
	if !ok {
		t.Fatalf("expected noticeMsg, got %T", cmd())
	}
	if !strings.Contains(nm.text, "/jobops:idealjob") {
		t.Errorf("ArgNone command should compose without a record: %q", nm.text)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/ui/ -run PaletteOpensWithNoRecords`
Expected: FAIL — the palette does not open (current `enter` guard requires `len(records) > 0`), or `composeSelected` returns "" with no record.

- [ ] **Step 3: Update three pieces of `app.go`**

(a) In `updateNormal`, replace the `"enter"` case so the palette opens when the tab has records OR has skills:
```go
	case "enter":
		if len(t.records) > 0 || len(t.skills) > 0 {
			m.mode = modePalette
			m.paletteAt = m.nextSkillIndex()
		}
```

(b) Replace `composeSelected` so `ArgNone` skills compose without a record, while `ArgJD`/`ArgCompany` still require one:
```go
func (m Model) composeSelected(spec model.SkillSpec) string {
	if spec.Name == "" {
		return ""
	}
	if spec.Arg == model.ArgNone {
		return launch.Compose(spec, model.Record{})
	}
	t := m.tabs[m.active]
	if len(t.records) == 0 {
		return ""
	}
	return launch.Compose(spec, t.records[t.cursor])
}
```

(c) Replace `renderPalette` so its header doesn't index an empty record slice:
```go
func (m Model) renderPalette() string {
	t := m.tabs[m.active]
	target := t.name
	if len(t.records) > 0 {
		target = t.records[t.cursor].Title
	}
	var b strings.Builder
	b.WriteString("Run on " + target + ":\n")
	for i, s := range t.skills {
		cursor := "  "
		if i == m.paletteAt {
			cursor = "› "
		}
		b.WriteString(fmt.Sprintf("%s%s\n", cursor, s.Label))
	}
	b.WriteString("↵ spawn " + m.agent + "  c copy  esc cancel\n")
	return b.String()
}
```

- [ ] **Step 4: Run the ui suite**

Run (from `tools/dashboard/`): `go test ./internal/ui/ && go vet ./internal/ui/`
Expected: all `ui` tests PASS (the new one plus all existing), vet clean.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/ui/app.go tools/dashboard/internal/ui/app_test.go
git commit -m "feat(dashboard): launch ArgNone skills from empty flat tabs"
```

---

## Task 6: Wire the new tabs into main

**Files:**
- Modify: `tools/dashboard/main.go`

- [ ] **Step 1: Update the sources slice**

In `tools/dashboard/main.go`, replace the `sources := []ui.TabSource{...}` block (added in Phase 2) with:
```go
	prefs := launch.LoadPrefs(root)
	sources := []ui.TabSource{
		{Name: "Apps", Scanner: scan.AppsAdapter{Cfg: cfg}, Lifecycle: true},
		{Name: "Companies", Scanner: scan.CompaniesAdapter{Cfg: cfg}},
		{Name: "Career", Scanner: scan.CareerAdapter{Cfg: cfg}},
		{Name: "Crisis", Scanner: scan.CrisisAdapter{Cfg: cfg}},
	}
	// Contractor (jobops-ic) only when its root is configured.
	if cfg.Directories.ContractorRoot != "" {
		sources = append(sources, ui.TabSource{Name: "Contractor", Scanner: scan.ContractorAdapter{Cfg: cfg}})
	}
	model := ui.New(root, prefs.Agent, sources)
```

- [ ] **Step 2: Build, vet, and test the whole module**

Run (from `tools/dashboard/`):
```bash
go build ./...
go vet ./...
go test ./... -count=1
```
Expected: build OK, vet clean, all packages PASS.

- [ ] **Step 3: Verify the unconfigured path (no TTY)**

Run (from `tools/dashboard/`):
```bash
tmp=$(mktemp -d)
echo "n" | go run . --workspace "$tmp"
```
Expected: prints the unconfigured guidance, exits 0, no TUI.

- [ ] **Step 4: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/main.go
git commit -m "feat(dashboard): show Career, Crisis, and (conditional) Contractor tabs"
```

---

## Task 7: README + final verification

**Files:**
- Modify: `tools/dashboard/README.md`

- [ ] **Step 1: Update the README**

In `tools/dashboard/README.md`, replace the first paragraph with:
```markdown
Terminal dashboard for tracking JobOps applications and launching `/jobops` and
`/jobops-ic` skills. Tabs: **Apps** (job applications), **Companies** (OSINT),
**Career**, **Crisis**, and **Contractor** (shown when jobops-ic is configured).
```

- [ ] **Step 2: Full green check**

Run (from `/home/reggiechan/JobOps`):
```bash
(cd tools/dashboard && go build ./... && go vet ./... && go test ./... -count=1)
npm test
claude plugin validate plugins/jobops
```
Expected: Go all green; `npm test` (Codex compat) passes; plugin validates.

- [ ] **Step 3: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/README.md
git commit -m "docs(dashboard): document Career/Crisis/Contractor tabs"
```

---

## Self-Review

**1. Spec coverage (§3.2 Career/Crisis/Contractor rows):**
- Career tab over `Career_Analysis/`, skills idealjob/comparejobs/change-one-thing → Task 3. ✓
- Crisis tab over `Crisis_Management/`, 11 crisis skills → Task 3. ✓
- Contractor tab over `Contractor/` with jobops-ic skills → Tasks 1 (Plugin), 4. ✓
- `/jobops-ic:` namespace → Task 1 (`SkillSpec.Plugin` + Compose). ✓
- Config plumbing for the three roots → Task 2. ✓
- All four new tabs surfaced (Contractor conditional) → Task 6. ✓
- Flat tabs launchable when empty → Task 5. ✓

**2. Placeholder scan:** No TBD/TODO; every code step has complete code; tests have real assertions. The `sortByUpdatedDesc` duplicate note in Task 3 includes an explicit remediation. ✓

**3. Type consistency:** `SkillSpec.Plugin` (Task 1) is read by `launch.Compose` (Task 1) and set by `ContractorAdapter.Skills` (Task 4). `flatRecords`/`sortByUpdatedDesc` (Task 3) reused by `ContractorAdapter` (Task 4). `CareerAdapter`/`CrisisAdapter`/`ContractorAdapter` all use the `{Cfg *config.Config}` shape and satisfy `model.Adapter`. `config.ContractorRoot` + `CareerAnalysisDir`/`CrisisManagementDir`/`ContractorDir` (Task 2) are consumed in Tasks 3, 4, 6. `ui.TabSource` (no `Lifecycle` for the flat tabs) and `composeSelected`/`renderPalette`/`enter` changes (Task 5) are consistent with Phase 2's `Model`. ✓

**Potential follow-up (not blocking):** `sortByUpdatedDesc` is introduced in `flat.go` (Task 3); `apps.go` and `companies.go` still have their own inline `sort.Slice`. A later cleanup could route both through `sortByUpdatedDesc`, but that refactor is out of scope here to keep the diff focused.
