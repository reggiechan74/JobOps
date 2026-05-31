# JobOps Dashboard TUI — Phase 2 Implementation Plan (Multi-tab shell + Companies tab)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Turn the single-tab dashboard into a tabbed shell and add a **Companies** tab over `Company_Intelligence/{Company}/`, so the engine's adapter pattern is proven with a second, artifact-only data source.

**Architecture:** Generalize `ui.Model` from one scanner to a slice of `tab`s (each owning its scanner, records, cursor, skills, and a `lifecycle` flag), selected by an `active` index with tab-switching keys. The render helpers become lifecycle-aware. A new `scan.CompaniesAdapter` returns one record per company with the 7 OSINT files as its "stages".

**Tech Stack:** Go 1.22+, Bubble Tea, Lip Gloss, existing `internal/{config,model,scan,launch,ui}` packages from Phase 1.

**Prerequisite:** Phase 1 is merged (the `tools/dashboard` module exists with `model`, `config`, `scan` (AppsAdapter), `launch`, and a single-tab `ui`). Work on a branch off `main`.

**Spec:** `docs/superpowers/specs/2026-05-31-jobops-dashboard-tui-design.md` (§3.2 Companies row, §6 tab bar).

**Deviations / decisions (made while planning):**
- **Lifecycle is per-tab, not per-record.** The table/detail show the Lifecycle column only for tabs whose `lifecycle` flag is true (Apps). Companies (and all Phase 3 tabs) render without it. This avoids the wrong "Interested" default leaking onto non-application records.
- **Companies "stages"** are the 7 canonical OSINT files (`corporate, legal, leadership, compensation, culture, market, summary`), each mapped to the `osint` skill; `NextSkill` is `osint` when any is missing.
- **`RenderDetail` only prints a pipeline line when the record has stages**, so flat tabs in Phase 3 don't show a misleading "Pipeline complete".
- **Rescan acts on the active tab only** (the tab the user just launched a skill from).

---

## File Structure

```
tools/dashboard/internal/
  scan/companies.go        # NEW: CompaniesAdapter (Scan/Skills/Name)
  scan/companies_test.go   # NEW
  ui/view.go               # MODIFY: RenderTable/RenderDetail gain showLifecycle bool
  ui/view_test.go          # MODIFY: pass the new bool arg
  ui/app.go                # MODIFY: tabs []tab + active int + TabSource + tab keys + tab bar
  ui/app_test.go           # MODIFY: New() takes []TabSource; tests index m.tabs[m.active]
main.go                    # MODIFY: build []ui.TabSource{Apps, Companies}
tools/dashboard/README.md  # MODIFY: document tabs + switching keys
```

---

## Task 1: Companies adapter

**Files:**
- Create: `tools/dashboard/internal/scan/companies.go`
- Test: `tools/dashboard/internal/scan/companies_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/scan/companies_test.go`:
```go
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
	// Acme: 2 of 7 OSINT files present.
	wr("Company_Intelligence/Acme/corporate.md", "# corp\n")
	wr("Company_Intelligence/Acme/summary.md", "# summary\n")
	// Beta Corp: all 7 present.
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
```

> Note: `findRec` and `stageState` already exist as helpers in `apps_test.go` (same package), so do not redefine them.

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/scan/ -run Companies`
Expected: FAIL — undefined `CompaniesAdapter`.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/scan/companies.go`:
```go
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
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/scan/`
Expected: PASS (all scan tests, including the new Companies ones).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/scan/companies.go tools/dashboard/internal/scan/companies_test.go
git commit -m "feat(dashboard): Companies adapter over Company_Intelligence"
```

---

## Task 2: Lifecycle-aware render helpers

**Files:**
- Modify: `tools/dashboard/internal/ui/view.go`
- Modify: `tools/dashboard/internal/ui/view_test.go`

- [ ] **Step 1: Update the tests first (red)**

Replace the THREE test functions in `tools/dashboard/internal/ui/view_test.go` (`TestRenderTableShowsRowsAndCursor`, `TestRenderDetailShowsStagesAndNext`, `TestRenderDetailBacklogNoScore`) with versions that pass the new `showLifecycle` argument, and add one test for the no-lifecycle path. Keep the existing `sampleRecords()` helper unchanged. The full replacement set:
```go
func TestRenderTableShowsRowsAndCursor(t *testing.T) {
	out := RenderTable(sampleRecords(), 0, 80, true)
	if !strings.Contains(out, "Acme Product Manager") || !strings.Contains(out, "Beta Capital Analyst") {
		t.Errorf("table missing rows:\n%s", out)
	}
	if !strings.Contains(out, "84%") {
		t.Errorf("table missing fit score:\n%s", out)
	}
	if !strings.Contains(out, "▶") {
		t.Errorf("table missing cursor marker:\n%s", out)
	}
	if !strings.Contains(out, "Lifecycle") {
		t.Errorf("lifecycle table should have a Lifecycle column:\n%s", out)
	}
}

func TestRenderTableNoLifecycleColumn(t *testing.T) {
	out := RenderTable(sampleRecords(), 0, 80, false)
	if strings.Contains(out, "Lifecycle") {
		t.Errorf("non-lifecycle table should omit the Lifecycle column:\n%s", out)
	}
}

func TestRenderDetailShowsStagesAndNext(t *testing.T) {
	out := RenderDetail(sampleRecords()[0], true)
	if !strings.Contains(out, "Assess") || !strings.Contains(out, "✓") || !strings.Contains(out, "○") {
		t.Errorf("detail missing stage glyphs:\n%s", out)
	}
	if !strings.Contains(out, "/jobops:coverletter") {
		t.Errorf("detail missing next step:\n%s", out)
	}
	if !strings.Contains(out, "Interviewing") {
		t.Errorf("detail missing lifecycle:\n%s", out)
	}
}

func TestRenderDetailBacklogNoScore(t *testing.T) {
	out := RenderDetail(sampleRecords()[1], true)
	if !strings.Contains(out, "—") {
		t.Errorf("backlog detail should show — for missing score:\n%s", out)
	}
}

func TestRenderDetailNoLifecycle(t *testing.T) {
	out := RenderDetail(sampleRecords()[0], false)
	if strings.Contains(out, "Lifecycle") {
		t.Errorf("no-lifecycle detail should not print Lifecycle:\n%s", out)
	}
}
```

- [ ] **Step 2: Run tests to verify they fail (compile error)**

Run (from `tools/dashboard/`): `go test ./internal/ui/`
Expected: FAIL — too many arguments to `RenderTable`/`RenderDetail`.

- [ ] **Step 3: Update `view.go`**

Replace the `RenderTable` and `RenderDetail` functions in `tools/dashboard/internal/ui/view.go` with these (the helpers `scoreCell`, `lifecycleCell`, `pipelineBar`, `startedLabel` stay as-is):
```go
// RenderTable renders the record list. cursor is the selected row index; width
// caps title length; showLifecycle toggles the Lifecycle column (Apps only).
func RenderTable(recs []model.Record, cursor, width int, showLifecycle bool) string {
	titleW := 28
	if width > 0 && width < 80 {
		titleW = 18
	}
	var b strings.Builder
	if showLifecycle {
		b.WriteString(fmt.Sprintf("  %-*s %-5s %-13s %s\n", titleW, "Application", "Fit", "Lifecycle", "Pipeline"))
	} else {
		b.WriteString(fmt.Sprintf("  %-*s %-5s %s\n", titleW, "Name", "Fit", "Progress"))
	}
	for i, r := range recs {
		marker := " "
		if i == cursor {
			marker = "▶"
		}
		title := r.Title
		if len(title) > titleW {
			title = title[:titleW-1] + "…"
		}
		warn := ""
		if r.Warn != "" {
			warn = " ⚠"
		}
		if showLifecycle {
			b.WriteString(fmt.Sprintf("%s %-*s %-5s %-13s %s%s\n",
				marker, titleW, title, scoreCell(r), lifecycleCell(r), pipelineBar(r), warn))
		} else {
			b.WriteString(fmt.Sprintf("%s %-*s %-5s %s%s\n",
				marker, titleW, title, scoreCell(r), pipelineBar(r), warn))
		}
	}
	if len(recs) == 0 {
		b.WriteString("  (nothing here yet)\n")
	}
	return b.String()
}

// RenderDetail renders the detail pane for one record. showLifecycle toggles the
// lifecycle/started line; the pipeline line is shown only when the record has
// stages (flat tabs have none).
func RenderDetail(r model.Record, showLifecycle bool) string {
	var b strings.Builder
	header := r.Title
	if r.Date != "" {
		header += "  ·  " + r.Date
	}
	b.WriteString(header + "\n")
	if showLifecycle {
		b.WriteString(fmt.Sprintf("Fit %s   Lifecycle: %s   %s\n",
			scoreCell(r), lifecycleCell(r), startedLabel(r)))
	} else {
		b.WriteString(fmt.Sprintf("Fit %s\n", scoreCell(r)))
	}

	if len(r.Stages) > 0 {
		var stages []string
		for _, s := range r.Stages {
			stages = append(stages, fmt.Sprintf("%s %s", s.State.Glyph(), s.Name))
		}
		b.WriteString(strings.Join(stages, "   ") + "\n")
		if r.NextSkill != "" {
			b.WriteString("Next: /jobops:" + r.NextSkill + "\n")
		} else {
			b.WriteString("Pipeline complete\n")
		}
	}
	if r.Warn != "" {
		b.WriteString("⚠ " + r.Warn + "\n")
	}
	return b.String()
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run (from `tools/dashboard/`): `go test ./internal/ui/`
Expected: FAIL to COMPILE on `app.go` (it still calls the 3-arg `RenderTable`). That is expected — Task 3 updates `app.go`. To verify *this task* in isolation, temporarily check just the view helpers compile by running `go vet ./internal/ui/` will also fail on app.go. Therefore: **do Task 3 before running the full `ui` test suite.** For this task, confirm only that `view.go` and `view_test.go` are internally consistent by reading them; the green bar comes at the end of Task 3.

> This task and Task 3 are a coupled pair (they change a shared function signature). Commit this task's view changes, then immediately do Task 3; the `ui` package will not compile in between, which is expected.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/ui/view.go tools/dashboard/internal/ui/view_test.go
git commit -m "feat(dashboard): lifecycle-aware render helpers (showLifecycle)"
```

---

## Task 3: Multi-tab Model

**Files:**
- Modify: `tools/dashboard/internal/ui/app.go`
- Modify: `tools/dashboard/internal/ui/app_test.go`

- [ ] **Step 1: Update the tests first (red)**

Replace `newTestModel` and the `staticScanner`/`shrinkingScanner` usages in `tools/dashboard/internal/ui/app_test.go` so `New` takes `[]TabSource` and tests index the active tab. Replace the whole file body below the imports with:
```go
func newTestModel() Model {
	m := New("/tmp/ws", "claude", []TabSource{
		{Name: "Apps", Scanner: staticScanner{recs: sampleRecords()}, Lifecycle: true},
	})
	m.width, m.height = 100, 30
	return m
}

// staticScanner is a Scanner that returns canned records (no filesystem).
type staticScanner struct{ recs []model.Record }

func (s staticScanner) Scan() ([]model.Record, error) { return s.recs, nil }
func (s staticScanner) Skills() []model.SkillSpec {
	return []model.SkillSpec{{Name: "coverletter", Label: "Cover letter", Arg: model.ArgJD}}
}

func TestCursorMovesDown(t *testing.T) {
	m := newTestModel()
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	if updated.(Model).tabs[0].cursor != 1 {
		t.Errorf("cursor = %d, want 1", updated.(Model).tabs[0].cursor)
	}
}

func TestEnterOpensPalette(t *testing.T) {
	m := newTestModel()
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if updated.(Model).mode != modePalette {
		t.Errorf("mode = %v, want modePalette", updated.(Model).mode)
	}
}

func TestStatusKeyCyclesLifecycle(t *testing.T) {
	m := newTestModel()
	start := m.tabs[0].records[0].Lifecycle
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}})
	got := updated.(Model).tabs[0].records[0].Lifecycle
	if got == start {
		t.Errorf("status key should advance lifecycle from %q", start)
	}
}

func TestViewRendersWithoutPanic(t *testing.T) {
	m := newTestModel()
	if m.View() == "" {
		t.Errorf("View() returned empty string")
	}
}

func TestSpawnMissingAgentFallback(t *testing.T) {
	m := newTestModel()
	m.agent = "definitely-not-a-real-agent-xyz"
	m.mode = modePalette
	m.paletteAt = 0
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected a command from the spawn fallback")
	}
	msg := cmd()
	nm, ok := msg.(noticeMsg)
	if !ok {
		t.Fatalf("expected noticeMsg, got %T", msg)
	}
	final, _ := updated.(Model).Update(nm)
	fm := final.(Model)
	if fm.mode != modeNotice {
		t.Errorf("mode = %v, want modeNotice", fm.mode)
	}
	if !strings.Contains(fm.notice, "not found") {
		t.Errorf("notice = %q, want it to mention 'not found'", fm.notice)
	}
}

type shrinkingScanner struct{ recs []model.Record }

func (s *shrinkingScanner) Scan() ([]model.Record, error) { return s.recs, nil }
func (s *shrinkingScanner) Skills() []model.SkillSpec      { return nil }

func TestRescanClampsCursor(t *testing.T) {
	three := []model.Record{
		{Title: "A", Slug: "A_20260101"},
		{Title: "B", Slug: "B_20260101"},
		{Title: "C", Slug: "C_20260101"},
	}
	sc := &shrinkingScanner{recs: three}
	m := New("/tmp/ws", "claude", []TabSource{{Name: "Apps", Scanner: sc, Lifecycle: true}})
	m.width, m.height = 100, 30
	m.tabs[0].cursor = 2

	sc.recs = three[:2]

	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'r'}})
	if cmd == nil {
		t.Fatal("expected r to return a rescan command")
	}
	updated, _ = updated.(Model).Update(cmd())
	fm := updated.(Model)
	if fm.tabs[0].cursor != 1 {
		t.Errorf("cursor = %d, want clamped to 1 after shrink-rescan", fm.tabs[0].cursor)
	}
	_ = fm.View()
}

func TestTabSwitchWraps(t *testing.T) {
	m := New("/tmp/ws", "claude", []TabSource{
		{Name: "Apps", Scanner: staticScanner{recs: sampleRecords()}, Lifecycle: true},
		{Name: "Companies", Scanner: staticScanner{recs: nil}},
	})
	m.width, m.height = 100, 30
	// right moves to tab 1
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRight})
	if updated.(Model).active != 1 {
		t.Errorf("active = %d, want 1 after right", updated.(Model).active)
	}
	// right again wraps to 0
	updated, _ = updated.(Model).Update(tea.KeyMsg{Type: tea.KeyRight})
	if updated.(Model).active != 0 {
		t.Errorf("active = %d, want 0 after wrap", updated.(Model).active)
	}
}
```

- [ ] **Step 2: Run tests to verify they fail**

Run (from `tools/dashboard/`): `go test ./internal/ui/`
Expected: FAIL — `New` signature mismatch, `TabSource`/`m.tabs`/`m.active` undefined.

- [ ] **Step 3: Rewrite `app.go`**

Replace the entire contents of `tools/dashboard/internal/ui/app.go` with:
```go
package ui

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/reggiechan74/jobops-dashboard/internal/launch"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
	"github.com/reggiechan74/jobops-dashboard/internal/scan"
)

// Scanner is the minimal adapter surface the UI needs; the real adapters satisfy
// it, and tests use static fakes.
type Scanner interface {
	Scan() ([]model.Record, error)
	Skills() []model.SkillSpec
}

// TabSource describes one tab the dashboard shows.
type TabSource struct {
	Name      string
	Scanner   Scanner
	Lifecycle bool // records carry a real-world lifecycle (Apps only)
}

type uiMode int

const (
	modeNormal uiMode = iota
	modePalette
	modeNotice
)

// tab holds the per-tab state.
type tab struct {
	name      string
	scanner   Scanner
	skills    []model.SkillSpec
	lifecycle bool
	records   []model.Record
	cursor    int
}

// Model is the Bubble Tea root model: a set of tabs plus shared UI state.
type Model struct {
	root  string
	agent string

	tabs   []tab
	active int

	mode      uiMode
	paletteAt int
	notice    string

	width, height int
}

// New builds a Model from one TabSource per tab, scanning each upfront.
func New(root, agent string, sources []TabSource) Model {
	tabs := make([]tab, 0, len(sources))
	for _, s := range sources {
		recs, _ := s.Scanner.Scan()
		tabs = append(tabs, tab{
			name:      s.Name,
			scanner:   s.Scanner,
			skills:    s.Scanner.Skills(),
			lifecycle: s.Lifecycle,
			records:   recs,
		})
	}
	return Model{root: root, agent: agent, tabs: tabs}
}

func (m Model) Init() tea.Cmd { return nil }

// rescanMsg is emitted after a spawned agent session returns.
type rescanMsg struct{}

// noticeMsg asks the model to display a transient notice.
type noticeMsg struct{ text string }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil
	case rescanMsg:
		t := &m.tabs[m.active]
		t.records, _ = t.scanner.Scan()
		if t.cursor >= len(t.records) {
			t.cursor = max(0, len(t.records)-1)
		}
		if len(t.records) == 0 {
			m.mode = modeNormal
		}
		return m, nil
	case noticeMsg:
		m.notice = msg.text
		m.mode = modeNotice
		return m, nil
	case tea.KeyMsg:
		if m.mode == modePalette {
			return m.updatePalette(msg)
		}
		if m.mode == modeNotice {
			m.mode = modeNormal
			return m, nil
		}
		return m.updateNormal(msg)
	}
	return m, nil
}

func (m Model) updateNormal(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	t := &m.tabs[m.active]
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "down", "j":
		if t.cursor < len(t.records)-1 {
			t.cursor++
		}
	case "up", "k":
		if t.cursor > 0 {
			t.cursor--
		}
	case "right", "tab":
		m.active = (m.active + 1) % len(m.tabs)
	case "left", "shift+tab":
		m.active = (m.active - 1 + len(m.tabs)) % len(m.tabs)
	case "enter":
		if len(t.records) > 0 {
			m.mode = modePalette
			m.paletteAt = m.nextSkillIndex()
		}
	case "c":
		if cmd := m.composeSelected(m.skillAt(m.nextSkillIndex())); cmd != "" {
			m = m.copyOrNotice(cmd)
		}
	case "C":
		if m.agent == "claude" {
			m.agent = "codex"
		} else {
			m.agent = "claude"
		}
		_ = launch.SavePrefs(m.root, launch.Prefs{Agent: m.agent})
	case "s":
		if t.lifecycle && len(t.records) > 0 {
			m = m.cycleLifecycle()
		}
	case "r":
		return m, func() tea.Msg { return rescanMsg{} }
	}
	return m, nil
}

func (m Model) updatePalette(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	t := &m.tabs[m.active]
	switch msg.String() {
	case "esc", "q":
		m.mode = modeNormal
	case "down", "j":
		if m.paletteAt < len(t.skills)-1 {
			m.paletteAt++
		}
	case "up", "k":
		if m.paletteAt > 0 {
			m.paletteAt--
		}
	case "c":
		cmd := m.composeSelected(m.skillAt(m.paletteAt))
		m.mode = modeNormal
		m = m.copyOrNotice(cmd)
	case "enter":
		cmd := m.composeSelected(m.skillAt(m.paletteAt))
		m.mode = modeNormal
		return m, m.spawn(cmd)
	}
	return m, nil
}

func (m Model) cycleLifecycle() Model {
	t := &m.tabs[m.active]
	rec := &t.records[t.cursor]
	cur := rec.Lifecycle
	if cur == "" {
		cur = model.DefaultLifecycle()
	}
	idx := 0
	for i, l := range model.LifecycleOrder {
		if l == cur {
			idx = i
			break
		}
	}
	next := model.LifecycleOrder[(idx+1)%len(model.LifecycleOrder)]
	rec.Lifecycle = next
	if len(rec.Paths) == 0 {
		return m
	}
	appDir := rec.Paths[0]
	tr, _ := scan.ReadTracker(appDir)
	tr.Lifecycle = next
	_ = scan.WriteTracker(appDir, tr)
	rec.Started = true
	return m
}

func (m Model) nextSkillIndex() int {
	t := m.tabs[m.active]
	if len(t.records) == 0 {
		return 0
	}
	next := t.records[t.cursor].NextSkill
	for i, s := range t.skills {
		if s.Name == next {
			return i
		}
	}
	return 0
}

func (m Model) skillAt(i int) model.SkillSpec {
	t := m.tabs[m.active]
	if i < 0 || i >= len(t.skills) {
		return model.SkillSpec{}
	}
	return t.skills[i]
}

func (m Model) composeSelected(spec model.SkillSpec) string {
	t := m.tabs[m.active]
	if spec.Name == "" || len(t.records) == 0 {
		return ""
	}
	return launch.Compose(spec, t.records[t.cursor])
}

func (m Model) copyOrNotice(cmd string) Model {
	if cmd == "" {
		return m
	}
	if launch.Copy(cmd) {
		m.notice = "Copied: " + cmd
	} else {
		m.notice = "No clipboard tool. Copy manually:\n  " + cmd
	}
	m.mode = modeNotice
	return m
}

// spawn suspends the TUI to run the interactive agent with cmd as its prompt,
// then triggers a rescan when it returns. If the agent is not on PATH it copies
// the command to the clipboard and reports that via a noticeMsg instead.
func (m Model) spawn(cmd string) tea.Cmd {
	if cmd == "" {
		return nil
	}
	if _, err := exec.LookPath(m.agent); err != nil {
		copied := launch.Copy(cmd)
		text := "Agent '" + m.agent + "' not found on PATH. "
		if copied {
			text += "Copied command to clipboard:\n  " + cmd
		} else {
			text += "Copy it manually:\n  " + cmd
		}
		return func() tea.Msg { return noticeMsg{text} }
	}
	c := exec.Command(m.agent, cmd)
	return tea.ExecProcess(c, func(error) tea.Msg { return rescanMsg{} })
}

func (m Model) View() string {
	t := m.tabs[m.active]
	var b strings.Builder
	b.WriteString(m.renderTabBar())
	b.WriteString(RenderTable(t.records, t.cursor, m.width, t.lifecycle))
	b.WriteString("├─────────────\n")
	if len(t.records) > 0 {
		b.WriteString(RenderDetail(t.records[t.cursor], t.lifecycle))
	}
	b.WriteString("├─────────────\n")
	switch m.mode {
	case modePalette:
		b.WriteString(m.renderPalette())
	case modeNotice:
		b.WriteString(m.notice + "\n[any key] dismiss\n")
	default:
		b.WriteString("↑↓ move  ←→ tab  ↵ run  c copy  C agent  s status  r rescan  q quit\n")
	}
	return b.String()
}

func (m Model) renderTabBar() string {
	var names []string
	for i, tb := range m.tabs {
		if i == m.active {
			names = append(names, "["+tb.name+"]")
		} else {
			names = append(names, tb.name)
		}
	}
	return fmt.Sprintf("┌ JobOps · %s  (agent: %s) ─[%d]\n",
		strings.Join(names, " "), m.agent, len(m.tabs[m.active].records))
}

func (m Model) renderPalette() string {
	t := m.tabs[m.active]
	var b strings.Builder
	b.WriteString("Run on " + t.records[t.cursor].Title + ":\n")
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

- [ ] **Step 4: Run the full ui suite**

Run (from `tools/dashboard/`): `go test ./internal/ui/ && go vet ./internal/ui/ && go build ./...`
Expected: all `ui` tests PASS, vet clean. (`go build ./...` will still fail on `main.go` until Task 4 — that's expected; if so, confirm `go test ./internal/ui/` alone is green.)

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/ui/app.go tools/dashboard/internal/ui/app_test.go
git commit -m "feat(dashboard): multi-tab Model with tab switching"
```

---

## Task 4: Wire the Companies tab into main

**Files:**
- Modify: `tools/dashboard/main.go`

- [ ] **Step 1: Update the program construction**

In `tools/dashboard/main.go`, the current code builds a single-scanner model:
```go
	prefs := launch.LoadPrefs(root)
	model := ui.New(root, prefs.Agent, scan.AppsAdapter{Cfg: cfg})
```
Replace those two lines with a `[]ui.TabSource` slice:
```go
	prefs := launch.LoadPrefs(root)
	sources := []ui.TabSource{
		{Name: "Apps", Scanner: scan.AppsAdapter{Cfg: cfg}, Lifecycle: true},
		{Name: "Companies", Scanner: scan.CompaniesAdapter{Cfg: cfg}},
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

- [ ] **Step 3: Verify the unconfigured path still works (no TTY)**

Run (from `tools/dashboard/`):
```bash
tmp=$(mktemp -d)
echo "n" | go run . --workspace "$tmp"
```
Expected: prints "No .jobops workspace found at or above: <tmp>" and the setup guidance; exits 0; no TUI.

- [ ] **Step 4: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/main.go
git commit -m "feat(dashboard): show Apps + Companies tabs"
```

---

## Task 5: Update the README

**Files:**
- Modify: `tools/dashboard/README.md`

- [ ] **Step 1: Update the README**

In `tools/dashboard/README.md`, replace the first paragraph (`Terminal dashboard ... Phase 1 ships the Apps tab.`) with:
```markdown
Terminal dashboard for tracking JobOps applications and launching `/jobops`
skills. Tabs: **Apps** (job applications) and **Companies** (OSINT intelligence).
```
And in the `## Keys` section, replace the keys line with one that includes tab switching:
```markdown
`↑↓`/`jk` move · `←→`/`tab` switch tabs · `↵` open skill palette · `c` copy
next-step command · `C` toggle claude/codex · `s` cycle lifecycle status (Apps) ·
`r` rescan · `q` quit.
```

- [ ] **Step 2: Verify it's still valid + everything green**

Run (from `/home/reggiechan/JobOps`):
```bash
(cd tools/dashboard && go build ./... && go vet ./... && go test ./... -count=1)
npm test
```
Expected: Go all green; `npm test` (Codex compat) passes.

- [ ] **Step 3: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/README.md
git commit -m "docs(dashboard): document tabs + switching keys"
```

---

## Self-Review

**1. Spec coverage (Companies row of §3.2, tab bar §6):**
- Companies adapter over `Company_Intelligence/{Company}/`, 7 OSINT stages, `osint` launch → Task 1. ✓
- Tab bar + `←→`/`tab` switching → Task 3 (`renderTabBar`, tab keys, `TestTabSwitchWraps`). ✓
- Per-tab lifecycle (Companies has none) → Tasks 2, 3 (`showLifecycle`, `tab.lifecycle`). ✓
- Companies wired into the app → Task 4. ✓

**2. Placeholder scan:** No TBD/TODO; complete code in every code step; tests have real assertions. The Task 2↔3 "package won't compile in between" note is an explicit, intentional coupling, not a placeholder. ✓

**3. Type consistency:** `TabSource{Name, Scanner, Lifecycle}`, `tab{name, scanner, skills, lifecycle, records, cursor}`, `New(root, agent string, []TabSource)`, `RenderTable(recs, cursor, width int, showLifecycle bool)`, `RenderDetail(r, showLifecycle bool)` are used identically across Tasks 1–4. `CompaniesAdapter{Cfg}` matches `AppsAdapter{Cfg}`. `existState`/`latestMtime`/`isHidden` reuse existing scan helpers (`isHidden` is new in companies.go; confirm `apps.go` doesn't already define it — it does not; it uses inline `strings.HasPrefix(name, ".")`). ✓
