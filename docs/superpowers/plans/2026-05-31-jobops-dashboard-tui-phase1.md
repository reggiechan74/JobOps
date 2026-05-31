# JobOps Dashboard TUI — Phase 1 Implementation Plan (Core Engine + Apps Tab)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a working Go + Bubble Tea terminal dashboard that lists job applications (backlog + in-progress), shows each one's JobOps pipeline progress, fit score, and lifecycle status, and launches `/jobops` skills by spawning an interactive `claude`/`codex` session.

**Architecture:** A shared core (config loader, frontmatter scanner, generic record model, launcher) plus an Apps tab adapter that unions `Job_Postings/` and `Applications/{slug}/`. The TUI is a master-detail layout with a modal skill-picker palette. Phase 1 ships the engine and the single Apps tab; later phases add more tabs as thin adapters.

**Tech Stack:** Go 1.22+, Bubble Tea (`charmbracelet/bubbletea`), Lip Gloss, Bubbles (`list`), `adrg/frontmatter` (+ `yaml.v3`) for YAML frontmatter. Tests use the standard `testing` package with `t.TempDir()` fixtures and `bubbletea/teatest` for one TUI smoke.

**Scope note:** This plan is Phase 1 from `docs/superpowers/specs/2026-05-31-jobops-dashboard-tui-design.md`. Phases 2–3 (Companies, Career, Crisis, Contractor tabs) are deferred to follow-up plans.

**Deviations from the spec (decided while planning, all intentional):**
- The `{Company}_{Role}_{YYYYMMDD}` slug has no delimiter between Company and Role and both may contain underscores, so it cannot be split reliably. The table therefore shows a single humanized **Application** column + **Date**, not separate Company/Role columns. The Company (needed for OSINT cross-reference and the `osint` launch arg) is detected by prefix-matching the slug against existing `Company_Intelligence/` folder names; if no match, Company is empty.
- Pipeline stages derived in Phase 1 are **Assess, Resume, Cover, OSINT, Briefing, Prep** (6). "JD audit" is offered as a launchable skill but is not a derived stage, because `auditjobposting`'s on-disk filename is not part of the fixed `Applications/` tree contract.
- A stage is **complete (`✓`) when its primary artifact exists**; **Draft (`◐`)** is used only for the resume intermediate case (step1/step2 present, step3 absent). Frontmatter `status` is not used for stage state, because several skills emit `status: draft` as their normal terminal state, which would make those stages never read as complete.
- `Compose` is a free function in the `launch` package (not an `Adapter` method as the spec sketched), so it is unit-testable without constructing an adapter.

---

## File Structure

All new code lives under `tools/dashboard/` as a standalone Go module, separate from the markdown plugin trees.

```
tools/dashboard/
  go.mod  go.sum
  main.go                       # workspace discovery, missing-config screen, start program
  internal/
    model/   model.go           # Record, Stage, StageState, Lifecycle, SkillSpec, ArgStyle, Adapter
    config/  config.go          # Config struct, Discover(), Load(), dir helpers
    scan/    frontmatter.go      # AppsMatter, ParseFrontmatter, leadingInt
             slug.go             # SplitSlug, DetectCompany
             tracker.go          # Tracker, ReadTracker, WriteTracker (atomic)
             apps.go             # AppsAdapter: Scan, Skills, deriveStages, readScore
    launch/  launch.go           # Compose, Copy (clipboard), Prefs (dashboard.json)
    ui/      view.go             # pure render helpers: renderTable, renderDetail, renderHelp
             app.go              # Bubble Tea Model: Init/Update/View, keys, palette, spawn
  README.md
```

Module path: `github.com/reggiechan74/jobops-dashboard`. Internal imports use that prefix (e.g. `github.com/reggiechan74/jobops-dashboard/internal/model`).

---

## Task 0: Install Go (prerequisite)

Go is not currently installed in this environment. Install Go 1.22+ before any other task.

- [ ] **Step 1: Install Go**

Run (Debian/Ubuntu-like environment):
```bash
sudo apt-get update && sudo apt-get install -y golang-go
```
If the packaged Go is older than 1.22, install from go.dev instead:
```bash
cd /tmp && curl -fsSL https://go.dev/dl/go1.22.5.linux-amd64.tar.gz -o go.tgz \
  && sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go.tgz
export PATH=$PATH:/usr/local/go/bin
```

- [ ] **Step 2: Verify**

Run: `go version`
Expected: `go version go1.22` (or newer).

---

## Task 1: Scaffold the Go module

**Files:**
- Create: `tools/dashboard/go.mod`
- Create: `tools/dashboard/main.go` (temporary stub, replaced in Task 12)

- [ ] **Step 1: Initialize the module**

Run:
```bash
mkdir -p tools/dashboard && cd tools/dashboard
go mod init github.com/reggiechan74/jobops-dashboard
```
Expected: creates `tools/dashboard/go.mod` with `module github.com/reggiechan74/jobops-dashboard` and a `go 1.22` line.

- [ ] **Step 2: Add dependencies**

Run (from `tools/dashboard/`):
```bash
go get github.com/charmbracelet/bubbletea@latest
go get github.com/charmbracelet/lipgloss@latest
go get github.com/charmbracelet/bubbles@latest
go get github.com/adrg/frontmatter@latest
go get gopkg.in/yaml.v3@latest
go get github.com/charmbracelet/x/exp/teatest@latest
```
Expected: `go.mod` lists these requires; `go.sum` is created.

- [ ] **Step 3: Write a temporary main stub**

Create `tools/dashboard/main.go`:
```go
package main

import "fmt"

func main() {
	fmt.Println("jobops-dash: not implemented yet")
}
```

- [ ] **Step 4: Verify it builds**

Run (from `tools/dashboard/`): `go build ./...`
Expected: no output, exit 0.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/go.mod tools/dashboard/go.sum tools/dashboard/main.go
git commit -m "feat(dashboard): scaffold Go module + bubbletea deps"
```

---

## Task 2: Model package

**Files:**
- Create: `tools/dashboard/internal/model/model.go`
- Test: `tools/dashboard/internal/model/model_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/model/model_test.go`:
```go
package model

import "testing"

func TestStageStateGlyph(t *testing.T) {
	cases := map[StageState]string{Missing: "○", Draft: "◐", Final: "✓"}
	for st, want := range cases {
		if got := st.Glyph(); got != want {
			t.Errorf("StageState(%d).Glyph() = %q, want %q", st, got, want)
		}
	}
}

func TestLifecycleValidAndDefault(t *testing.T) {
	if !Interviewing.Valid() {
		t.Errorf("Interviewing should be valid")
	}
	if Lifecycle("bogus").Valid() {
		t.Errorf("bogus should be invalid")
	}
	if DefaultLifecycle() != Interested {
		t.Errorf("default = %q, want %q", DefaultLifecycle(), Interested)
	}
}

func TestLifecycleLabel(t *testing.T) {
	if Interviewing.Label() != "Interviewing" {
		t.Errorf("Label() = %q, want Interviewing", Interviewing.Label())
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/model/`
Expected: FAIL — undefined: StageState, Missing, Draft, Final, Lifecycle, etc.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/model/model.go`:
```go
// Package model defines the shared types the dashboard engine and its tab
// adapters exchange: a generic Record, its pipeline Stages, and the SkillSpec
// describing a launchable /jobops skill.
package model

import (
	"strings"
	"time"
)

// StageState is the completion state of a single pipeline stage.
type StageState int

const (
	Missing StageState = iota // no artifact on disk
	Draft                     // intermediate artifact only (e.g. resume step1/2)
	Final                     // primary artifact present
)

// Glyph returns the single-rune indicator used in the table/detail views.
func (s StageState) Glyph() string {
	switch s {
	case Final:
		return "✓"
	case Draft:
		return "◐"
	default:
		return "○"
	}
}

// Stage is one step of the JobOps pipeline for a record.
type Stage struct {
	Name  string // e.g. "Resume"
	Skill string // /jobops skill that produces it, e.g. "buildresume"
	State StageState
}

// Lifecycle is the real-world application status the dashboard owns.
type Lifecycle string

const (
	Interested   Lifecycle = "interested"
	Applied      Lifecycle = "applied"
	Screening    Lifecycle = "screening"
	Interviewing Lifecycle = "interviewing"
	Offer        Lifecycle = "offer"
	Accepted     Lifecycle = "accepted"
	Rejected     Lifecycle = "rejected"
	Withdrawn    Lifecycle = "withdrawn"
)

// LifecycleOrder is the cycle order used by the status setter.
var LifecycleOrder = []Lifecycle{
	Interested, Applied, Screening, Interviewing, Offer, Accepted, Rejected, Withdrawn,
}

// Valid reports whether l is a known lifecycle value.
func (l Lifecycle) Valid() bool {
	for _, v := range LifecycleOrder {
		if v == l {
			return true
		}
	}
	return false
}

// Label returns a human-readable title-cased label.
func (l Lifecycle) Label() string {
	if l == "" {
		return ""
	}
	s := string(l)
	return strings.ToUpper(s[:1]) + s[1:]
}

// DefaultLifecycle is the status assumed when no tracker file exists.
func DefaultLifecycle() Lifecycle { return Interested }

// ArgStyle describes how a skill's command line is composed from a record.
type ArgStyle int

const (
	ArgJD      ArgStyle = iota // "<slug>.md --app=<slug>"
	ArgCompany                 // "<company>"
	ArgNone                    // no positional args
)

// SkillSpec describes one launchable /jobops skill in a tab's palette.
type SkillSpec struct {
	Name  string // skill name without the /jobops: prefix, e.g. "assessjob"
	Label string // display label in the palette
	Arg   ArgStyle
}

// Record is the generic unit every tab renders.
type Record struct {
	Title     string    // humanized application title (slug minus date)
	Date      string    // trailing date token from the slug, "" if none
	Slug      string    // application slug / folder name
	Company   string    // detected company (for OSINT), "" if unknown
	Score     *int      // normalized fit %, nil if not assessed
	Lifecycle Lifecycle // zero value treated as Interested
	Stages    []Stage   // ordered pipeline
	NextSkill string    // first incomplete stage's skill, "" if complete
	Started   bool      // true if an Applications/{slug}/ folder exists
	Updated   time.Time // most recent artifact mtime (zero if backlog)
	Paths     []string  // backing files/dirs
	Warn      string    // non-empty if slug/frontmatter was malformed
}

// Adapter is implemented by each tab. Phase 1 ships only the Apps adapter.
type Adapter interface {
	Name() string
	Scan() ([]Record, error)
	Skills() []SkillSpec
}
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/model/`
Expected: PASS (`ok ... internal/model`).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/model/
git commit -m "feat(dashboard): model types (Record, Stage, Lifecycle, SkillSpec)"
```

---

## Task 3: Config package

**Files:**
- Create: `tools/dashboard/internal/config/config.go`
- Test: `tools/dashboard/internal/config/config_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/config/config_test.go`:
```go
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
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/config/`
Expected: FAIL — undefined: Discover, Load.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/config/config.go`:
```go
// Package config loads the JobOps workspace configuration (.jobops/config.json)
// and resolves the configured directories to absolute paths.
package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// Config mirrors the parts of .jobops/config.json the dashboard reads.
type Config struct {
	Directories struct {
		ResumeSource        string `json:"resume_source"`
		JobPostings         string `json:"job_postings"`
		ApplicationsRoot    string `json:"applications_root"`
		CompanyIntelligence string `json:"company_intelligence"`
		CareerAnalysis      string `json:"career_analysis"`
		CrisisManagement    string `json:"crisis_management"`
	} `json:"directories"`

	Root string `json:"-"` // workspace root (the dir that contains .jobops/)
}

// Discover walks up from start looking for the nearest ancestor that contains a
// .jobops directory, and returns that ancestor path.
func Discover(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}
	for {
		info, err := os.Stat(filepath.Join(dir, ".jobops"))
		if err == nil && info.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("no .jobops directory found in any parent of " + start)
		}
		dir = parent
	}
}

// Load reads .jobops/config.json from the given workspace root.
func Load(root string) (*Config, error) {
	data, err := os.ReadFile(filepath.Join(root, ".jobops", "config.json"))
	if err != nil {
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	c.Root = root
	return &c, nil
}

// resolve joins a configured (possibly relative, possibly "./"-prefixed) path
// with the workspace root, unless it is already absolute.
func (c *Config) resolve(p string) string {
	if p == "" {
		return ""
	}
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(c.Root, filepath.Clean(p))
}

func (c *Config) JobPostingsDir() string  { return c.resolve(c.Directories.JobPostings) }
func (c *Config) ApplicationsDir() string { return c.resolve(c.Directories.ApplicationsRoot) }
func (c *Config) CompanyIntelDir() string { return c.resolve(c.Directories.CompanyIntelligence) }
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/config/`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/config/
git commit -m "feat(dashboard): config loader + workspace discovery"
```

---

## Task 4: Frontmatter parsing

**Files:**
- Create: `tools/dashboard/internal/scan/frontmatter.go`
- Test: `tools/dashboard/internal/scan/frontmatter_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/scan/frontmatter_test.go`:
```go
package scan

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFrontmatter(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "assessment.md")
	body := "---\noutput_type: assessment\nstatus: draft\nnormalized_score: 84%\noverall_score: 168/200\n---\n# Body\n"
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	var m AppsMatter
	warn, err := ParseFrontmatter(path, &m)
	if err != nil {
		t.Fatal(err)
	}
	if warn != "" {
		t.Errorf("unexpected warn %q", warn)
	}
	if m.Status != "draft" || m.NormalizedScore != "84%" || m.OverallScore != "168/200" {
		t.Errorf("parsed = %+v", m)
	}
}

func TestParseFrontmatterMalformed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.md")
	if err := os.WriteFile(path, []byte("no frontmatter here\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	var m AppsMatter
	warn, err := ParseFrontmatter(path, &m)
	if err != nil {
		t.Fatalf("should not hard-error: %v", err)
	}
	if warn == "" {
		t.Errorf("expected a warn for missing frontmatter")
	}
}

func TestLeadingInt(t *testing.T) {
	cases := map[string]struct {
		n  int
		ok bool
	}{"84%": {84, true}, "168/200": {168, true}, "": {0, false}, "B+": {0, false}}
	for in, want := range cases {
		n, ok := leadingInt(in)
		if n != want.n || ok != want.ok {
			t.Errorf("leadingInt(%q) = (%d,%v), want (%d,%v)", in, n, ok, want.n, want.ok)
		}
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/scan/`
Expected: FAIL — undefined: AppsMatter, ParseFrontmatter, leadingInt.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/scan/frontmatter.go`:
```go
// Package scan walks JobOps data roots and builds dashboard records from the
// files and YAML frontmatter on disk.
package scan

import (
	"os"
	"strconv"

	"github.com/adrg/frontmatter"
)

// AppsMatter holds the frontmatter fields the Apps adapter cares about.
type AppsMatter struct {
	OutputType      string `yaml:"output_type"`
	Status          string `yaml:"status"`
	NormalizedScore string `yaml:"normalized_score"`
	OverallScore    string `yaml:"overall_score"`
}

// ParseFrontmatter parses the YAML frontmatter of a markdown file into out.
// A missing or malformed frontmatter block is returned as a non-empty warn
// string rather than a hard error, so a scan never aborts on one bad file.
func ParseFrontmatter(path string, out interface{}) (warn string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, perr := frontmatter.Parse(f, out); perr != nil {
		return "could not parse frontmatter: " + perr.Error(), nil
	}
	return "", nil
}

// leadingInt extracts the leading run of digits from s (e.g. "84%" -> 84,
// "168/200" -> 168). Returns ok=false when s has no leading digit.
func leadingInt(s string) (int, bool) {
	end := 0
	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		end++
	}
	if end == 0 {
		return 0, false
	}
	n, err := strconv.Atoi(s[:end])
	if err != nil {
		return 0, false
	}
	return n, true
}
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/scan/`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/scan/frontmatter.go tools/dashboard/internal/scan/frontmatter_test.go
git commit -m "feat(dashboard): tolerant frontmatter parsing + leadingInt"
```

---

## Task 5: Slug parsing & company detection

**Files:**
- Create: `tools/dashboard/internal/scan/slug.go`
- Test: `tools/dashboard/internal/scan/slug_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/scan/slug_test.go`:
```go
package scan

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplitSlug(t *testing.T) {
	cases := []struct {
		in          string
		title, date string
	}{
		{"Acme_Product_Manager_20260531", "Acme Product Manager", "20260531"},
		{"Beta_Capital_Analyst", "Beta Capital Analyst", ""},
		{"Gamma_20251231", "Gamma", "20251231"},
	}
	for _, c := range cases {
		title, date := SplitSlug(c.in)
		if title != c.title || date != c.date {
			t.Errorf("SplitSlug(%q) = (%q,%q), want (%q,%q)", c.in, title, date, c.title, c.date)
		}
	}
}

func TestDetectCompany(t *testing.T) {
	root := t.TempDir()
	intel := filepath.Join(root, "Company_Intelligence", "Acme")
	if err := os.MkdirAll(intel, 0o755); err != nil {
		t.Fatal(err)
	}
	got := DetectCompany("Acme_Product_Manager_20260531", filepath.Join(root, "Company_Intelligence"))
	if got != "Acme" {
		t.Errorf("DetectCompany = %q, want Acme", got)
	}
	if DetectCompany("Zzz_Role_20260101", filepath.Join(root, "Company_Intelligence")) != "" {
		t.Errorf("expected empty company for non-matching slug")
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/scan/ -run 'Slug|Company'`
Expected: FAIL — undefined: SplitSlug, DetectCompany.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/scan/slug.go`:
```go
package scan

import (
	"os"
	"path/filepath"
	"strings"
)

// SplitSlug separates a "{Company}_{Role}_{YYYYMMDD}" stem into a humanized
// title (underscores -> spaces, trailing 8-digit date removed) and the date
// token. The Company/Role boundary is not recoverable, so the whole non-date
// portion is returned as a single title.
func SplitSlug(stem string) (title, date string) {
	parts := strings.Split(stem, "_")
	if n := len(parts); n > 1 && isDate8(parts[n-1]) {
		date = parts[n-1]
		parts = parts[:n-1]
	}
	return strings.Join(parts, " "), date
}

func isDate8(s string) bool {
	if len(s) != 8 {
		return false
	}
	for i := 0; i < 8; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// DetectCompany returns the name of the Company_Intelligence subfolder whose
// slugified name is a prefix of slug (case-insensitive), or "" if none matches.
func DetectCompany(slug, companyIntelDir string) string {
	entries, err := os.ReadDir(companyIntelDir)
	if err != nil {
		return ""
	}
	lowSlug := strings.ToLower(slug)
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		cslug := strings.ToLower(slugify(e.Name()))
		if strings.HasPrefix(lowSlug, cslug+"_") || lowSlug == cslug {
			return e.Name()
		}
	}
	return ""
}

// slugify converts a company folder name to the underscore form used in slugs.
func slugify(name string) string {
	r := strings.NewReplacer(" ", "_", "&", "_", "-", "_", ",", "", ".", "")
	out := r.Replace(name)
	for strings.Contains(out, "__") {
		out = strings.ReplaceAll(out, "__", "_")
	}
	return strings.Trim(out, "_")
}

var _ = filepath.Join // keep filepath imported for future helpers
```

> Note: remove the `var _ = filepath.Join` line if `filepath` ends up used elsewhere in this file; it exists only to avoid an unused-import error in this isolated task. (It is used in the test, not the impl.) If `go build` complains the import is unused, delete the `filepath` import and the `var _` line together.

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/scan/ -run 'Slug|Company'`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/scan/slug.go tools/dashboard/internal/scan/slug_test.go
git commit -m "feat(dashboard): slug split + company detection"
```

---

## Task 6: Lifecycle tracker (read/write)

**Files:**
- Create: `tools/dashboard/internal/scan/tracker.go`
- Test: `tools/dashboard/internal/scan/tracker_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/scan/tracker_test.go`:
```go
package scan

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func TestReadTrackerMissingDefaults(t *testing.T) {
	tr, err := ReadTracker(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if tr.Lifecycle != model.Interested {
		t.Errorf("missing tracker should default to Interested, got %q", tr.Lifecycle)
	}
}

func TestWriteThenReadTracker(t *testing.T) {
	dir := t.TempDir()
	in := Tracker{Lifecycle: model.Interviewing, AppliedOn: "2026-05-20", Notes: "panel 06-03"}
	if err := WriteTracker(dir, in); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".tracker.json")); err != nil {
		t.Fatalf(".tracker.json not written: %v", err)
	}
	out, err := ReadTracker(dir)
	if err != nil {
		t.Fatal(err)
	}
	if out.Lifecycle != model.Interviewing || out.AppliedOn != "2026-05-20" || out.Notes != "panel 06-03" {
		t.Errorf("round-trip mismatch: %+v", out)
	}
	if out.UpdatedAt == "" {
		t.Errorf("WriteTracker should stamp UpdatedAt")
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/scan/ -run Tracker`
Expected: FAIL — undefined: Tracker, ReadTracker, WriteTracker.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/scan/tracker.go`:
```go
package scan

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

// Tracker is the dashboard-owned per-application lifecycle state, stored at
// Applications/{slug}/.tracker.json.
type Tracker struct {
	Lifecycle model.Lifecycle `json:"lifecycle"`
	AppliedOn string          `json:"applied_on,omitempty"`
	Notes     string          `json:"notes,omitempty"`
	UpdatedAt string          `json:"updated_at"`
}

const trackerFile = ".tracker.json"

// ReadTracker reads {appDir}/.tracker.json. A missing file yields a default
// Tracker (Interested) with no error; an unreadable/corrupt file also defaults,
// so a bad tracker never breaks a scan.
func ReadTracker(appDir string) (Tracker, error) {
	data, err := os.ReadFile(filepath.Join(appDir, trackerFile))
	if err != nil {
		if os.IsNotExist(err) {
			return Tracker{Lifecycle: model.DefaultLifecycle()}, nil
		}
		return Tracker{Lifecycle: model.DefaultLifecycle()}, nil
	}
	var t Tracker
	if json.Unmarshal(data, &t) != nil || !t.Lifecycle.Valid() {
		return Tracker{Lifecycle: model.DefaultLifecycle()}, nil
	}
	return t, nil
}

// WriteTracker writes {appDir}/.tracker.json atomically (temp file + rename),
// creating appDir if needed and stamping UpdatedAt.
func WriteTracker(appDir string, t Tracker) error {
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return err
	}
	t.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	tmp := filepath.Join(appDir, trackerFile+".tmp")
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, filepath.Join(appDir, trackerFile))
}
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/scan/ -run Tracker`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/scan/tracker.go tools/dashboard/internal/scan/tracker_test.go
git commit -m "feat(dashboard): atomic per-app lifecycle tracker.json"
```

---

## Task 7: Apps adapter (Scan, stages, score)

**Files:**
- Create: `tools/dashboard/internal/scan/apps.go`
- Test: `tools/dashboard/internal/scan/apps_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/scan/apps_test.go`:
```go
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
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/scan/ -run 'AppsScan|ResumeDraft'`
Expected: FAIL — undefined: AppsAdapter.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/scan/apps.go`:
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
// produces each. "JD audit" is intentionally not a derived stage (see plan).
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
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/scan/`
Expected: PASS (all scan tests).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/scan/apps.go tools/dashboard/internal/scan/apps_test.go
git commit -m "feat(dashboard): Apps adapter — union scan, pipeline stages, fit score"
```

---

## Task 8: Launcher — command composition

**Files:**
- Create: `tools/dashboard/internal/launch/launch.go`
- Test: `tools/dashboard/internal/launch/launch_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/launch/launch_test.go`:
```go
package launch

import (
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func TestCompose(t *testing.T) {
	rec := model.Record{Slug: "Acme_Product_Manager_20260531", Company: "Acme"}
	cases := []struct {
		spec model.SkillSpec
		want string
	}{
		{model.SkillSpec{Name: "assessjob", Arg: model.ArgJD},
			"/jobops:assessjob Acme_Product_Manager_20260531.md --app=Acme_Product_Manager_20260531"},
		{model.SkillSpec{Name: "osint", Arg: model.ArgCompany}, "/jobops:osint Acme"},
		{model.SkillSpec{Name: "convert-to-pdf", Arg: model.ArgNone}, "/jobops:convert-to-pdf"},
	}
	for _, c := range cases {
		if got := Compose(c.spec, rec); got != c.want {
			t.Errorf("Compose(%s) = %q, want %q", c.spec.Name, got, c.want)
		}
	}
}

func TestComposeCompanyUnknown(t *testing.T) {
	rec := model.Record{Slug: "X_Y_20260101"} // no Company
	got := Compose(model.SkillSpec{Name: "osint", Arg: model.ArgCompany}, rec)
	if got != "/jobops:osint" {
		t.Errorf("Compose osint w/o company = %q, want /jobops:osint", got)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/launch/`
Expected: FAIL — undefined: Compose.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/launch/launch.go`:
```go
// Package launch composes /jobops skill commands and provides clipboard +
// preference helpers. Spawning the interactive agent (claude/codex) is wired in
// the ui package via tea.ExecProcess; this package stays bubbletea-free so it
// is unit-testable.
package launch

import (
	"fmt"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

// Compose builds the slash command for running spec against rec.
func Compose(spec model.SkillSpec, rec model.Record) string {
	base := "/jobops:" + spec.Name
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

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/launch/`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/launch/launch.go tools/dashboard/internal/launch/launch_test.go
git commit -m "feat(dashboard): launcher command composition"
```

---

## Task 9: Launcher — clipboard & agent preferences

**Files:**
- Modify: `tools/dashboard/internal/launch/launch.go`
- Create: `tools/dashboard/internal/launch/prefs.go`
- Test: `tools/dashboard/internal/launch/prefs_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/launch/prefs_test.go`:
```go
package launch

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrefsRoundTripAndDefault(t *testing.T) {
	root := t.TempDir()
	_ = os.MkdirAll(filepath.Join(root, ".jobops"), 0o755)

	// Missing file -> default agent "claude".
	p := LoadPrefs(root)
	if p.Agent != "claude" {
		t.Errorf("default agent = %q, want claude", p.Agent)
	}

	p.Agent = "codex"
	if err := SavePrefs(root, p); err != nil {
		t.Fatal(err)
	}
	if LoadPrefs(root).Agent != "codex" {
		t.Errorf("saved agent not read back")
	}
}

func TestClipboardCommandDetection(t *testing.T) {
	// clipboardCommand picks the first available tool name from candidates.
	got := clipboardCommand([]string{"definitely-not-a-real-tool-xyz", "echo"})
	if got != "echo" {
		t.Errorf("clipboardCommand = %q, want echo (first found on PATH)", got)
	}
	if clipboardCommand([]string{"definitely-not-a-real-tool-xyz"}) != "" {
		t.Errorf("expected empty when no candidate exists")
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/launch/ -run 'Prefs|Clipboard'`
Expected: FAIL — undefined: LoadPrefs, SavePrefs, clipboardCommand.

- [ ] **Step 3: Write the prefs + clipboard implementation**

Create `tools/dashboard/internal/launch/prefs.go`:
```go
package launch

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
)

// Prefs is dashboard-owned UI state, stored at .jobops/dashboard.json. It is
// never written by JobOps skills.
type Prefs struct {
	Agent string `json:"agent"` // "claude" or "codex"
}

const prefsFile = "dashboard.json"

// LoadPrefs reads .jobops/dashboard.json from root, defaulting Agent to
// "claude" when the file is missing or unreadable.
func LoadPrefs(root string) Prefs {
	def := Prefs{Agent: "claude"}
	data, err := os.ReadFile(filepath.Join(root, ".jobops", prefsFile))
	if err != nil {
		return def
	}
	var p Prefs
	if json.Unmarshal(data, &p) != nil || (p.Agent != "claude" && p.Agent != "codex") {
		return def
	}
	return p
}

// SavePrefs writes .jobops/dashboard.json atomically.
func SavePrefs(root string, p Prefs) error {
	dir := filepath.Join(root, ".jobops")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	tmp := filepath.Join(dir, prefsFile+".tmp")
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, filepath.Join(dir, prefsFile))
}

// clipboardCommand returns the first candidate command name found on PATH, or
// "" if none is available.
func clipboardCommand(candidates []string) string {
	for _, c := range candidates {
		if _, err := exec.LookPath(c); err == nil {
			return c
		}
	}
	return ""
}

// Copy writes text to the system clipboard using the first available tool.
// Returns false if no clipboard tool is available (caller should then show the
// command in a modal for manual copying).
func Copy(text string) bool {
	cmd := clipboardCommand([]string{"pbcopy", "wl-copy", "xclip", "xsel", "clip.exe"})
	if cmd == "" {
		return false
	}
	args := []string{}
	if cmd == "xclip" {
		args = []string{"-selection", "clipboard"}
	} else if cmd == "xsel" {
		args = []string{"--clipboard", "--input"}
	}
	c := exec.Command(cmd, args...)
	stdin, err := c.StdinPipe()
	if err != nil {
		return false
	}
	if err := c.Start(); err != nil {
		return false
	}
	_, _ = stdin.Write([]byte(text))
	_ = stdin.Close()
	return c.Wait() == nil
}
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/launch/`
Expected: PASS (all launch tests).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/launch/prefs.go tools/dashboard/internal/launch/prefs_test.go
git commit -m "feat(dashboard): clipboard copy + agent preference persistence"
```

---

## Task 10: UI render helpers (pure functions)

**Files:**
- Create: `tools/dashboard/internal/ui/view.go`
- Test: `tools/dashboard/internal/ui/view_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/ui/view_test.go`:
```go
package ui

import (
	"strings"
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func sampleRecords() []model.Record {
	score := 84
	return []model.Record{
		{
			Title: "Acme Product Manager", Date: "20260531", Slug: "Acme_Product_Manager_20260531",
			Company: "Acme", Score: &score, Lifecycle: model.Interviewing, Started: true,
			NextSkill: "coverletter",
			Stages: []model.Stage{
				{Name: "Assess", State: model.Final}, {Name: "Resume", State: model.Final},
				{Name: "Cover", State: model.Missing}, {Name: "OSINT", State: model.Final},
				{Name: "Briefing", State: model.Missing}, {Name: "Prep", State: model.Missing},
			},
		},
		{Title: "Beta Capital Analyst", Date: "20260101", Slug: "Beta_Capital_Analyst_20260101", Lifecycle: model.Interested},
	}
}

func TestRenderTableShowsRowsAndCursor(t *testing.T) {
	out := RenderTable(sampleRecords(), 0, 80)
	if !strings.Contains(out, "Acme Product Manager") || !strings.Contains(out, "Beta Capital Analyst") {
		t.Errorf("table missing rows:\n%s", out)
	}
	if !strings.Contains(out, "84%") {
		t.Errorf("table missing fit score:\n%s", out)
	}
	if !strings.Contains(out, "▶") {
		t.Errorf("table missing cursor marker:\n%s", out)
	}
}

func TestRenderDetailShowsStagesAndNext(t *testing.T) {
	out := RenderDetail(sampleRecords()[0])
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
	out := RenderDetail(sampleRecords()[1])
	if !strings.Contains(out, "—") {
		t.Errorf("backlog detail should show — for missing score:\n%s", out)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/ui/`
Expected: FAIL — undefined: RenderTable, RenderDetail.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/ui/view.go`:
```go
// Package ui renders the dashboard. view.go holds pure string-producing helpers
// (no Bubble Tea state) so they can be unit-tested directly.
package ui

import (
	"fmt"
	"strings"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func scoreCell(r model.Record) string {
	if r.Score == nil {
		return "—"
	}
	return fmt.Sprintf("%d%%", *r.Score)
}

func lifecycleCell(r model.Record) string {
	l := r.Lifecycle
	if l == "" {
		l = model.DefaultLifecycle()
	}
	return l.Label()
}

// pipelineBar renders the stage states as a compact glyph run, e.g. "✓✓○✓○○".
func pipelineBar(r model.Record) string {
	var b strings.Builder
	for _, s := range r.Stages {
		b.WriteString(s.State.Glyph())
	}
	return b.String()
}

// RenderTable renders the application list. cursor is the selected row index;
// width is the available terminal width (currently used only to cap titles).
func RenderTable(recs []model.Record, cursor, width int) string {
	titleW := 28
	if width > 0 && width < 80 {
		titleW = 18
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("  %-*s %-5s %-13s %s\n", titleW, "Application", "Fit", "Lifecycle", "Pipeline"))
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
		b.WriteString(fmt.Sprintf("%s %-*s %-5s %-13s %s%s\n",
			marker, titleW, title, scoreCell(r), lifecycleCell(r), pipelineBar(r), warn))
	}
	if len(recs) == 0 {
		b.WriteString("  (no applications found — add a JD to Job_Postings/ or run /jobops:assessjob)\n")
	}
	return b.String()
}

// RenderDetail renders the detail pane for one record.
func RenderDetail(r model.Record) string {
	var b strings.Builder
	header := r.Title
	if r.Date != "" {
		header += "  ·  " + r.Date
	}
	b.WriteString(header + "\n")
	b.WriteString(fmt.Sprintf("Fit %s   Lifecycle: %s   %s\n",
		scoreCell(r), lifecycleCell(r), startedLabel(r)))

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
	if r.Warn != "" {
		b.WriteString("⚠ " + r.Warn + "\n")
	}
	return b.String()
}

func startedLabel(r model.Record) string {
	if r.Started {
		return "(in progress)"
	}
	return "(backlog)"
}
```

- [ ] **Step 4: Run test to verify it passes**

Run (from `tools/dashboard/`): `go test ./internal/ui/`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/ui/view.go tools/dashboard/internal/ui/view_test.go
git commit -m "feat(dashboard): pure table + detail render helpers"
```

---

## Task 11: UI model (Bubble Tea) — keys, palette, spawn, status

**Files:**
- Create: `tools/dashboard/internal/ui/app.go`
- Test: `tools/dashboard/internal/ui/app_test.go`

- [ ] **Step 1: Write the failing test**

Create `tools/dashboard/internal/ui/app_test.go`:
```go
package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func newTestModel() Model {
	m := New("/tmp/ws", "claude", staticScanner{recs: sampleRecords()})
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
	if updated.(Model).cursor != 1 {
		t.Errorf("cursor = %d, want 1", updated.(Model).cursor)
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
	m.cursor = 0
	start := m.records[0].Lifecycle
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}})
	got := updated.(Model).records[0].Lifecycle
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
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `tools/dashboard/`): `go test ./internal/ui/ -run 'Cursor|Palette|Status|View'`
Expected: FAIL — undefined: New, Model, modePalette, etc.

- [ ] **Step 3: Write the implementation**

Create `tools/dashboard/internal/ui/app.go`:
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

// Scanner is the minimal adapter surface the UI needs; the real Apps adapter
// satisfies it, and tests use a static fake.
type Scanner interface {
	Scan() ([]model.Record, error)
	Skills() []model.SkillSpec
}

type uiMode int

const (
	modeNormal uiMode = iota
	modePalette
	modeNotice
)

// Model is the Bubble Tea root model for Phase 1 (single Apps tab).
type Model struct {
	root    string // workspace root
	agent   string // "claude" | "codex"
	scanner Scanner

	records []model.Record
	cursor  int

	mode      uiMode
	paletteAt int // cursor in the palette list
	skills    []model.SkillSpec
	notice    string

	width, height int
}

// New builds a Model. root and agent come from config/prefs; scanner is the
// Apps adapter (or a fake in tests).
func New(root, agent string, scanner Scanner) Model {
	recs, _ := scanner.Scan()
	return Model{
		root:    root,
		agent:   agent,
		scanner: scanner,
		records: recs,
		skills:  scanner.Skills(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

// rescanMsg is emitted after a spawned agent session returns.
type rescanMsg struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil
	case rescanMsg:
		m.records, _ = m.scanner.Scan()
		if m.cursor >= len(m.records) {
			m.cursor = max(0, len(m.records)-1)
		}
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
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "down", "j":
		if m.cursor < len(m.records)-1 {
			m.cursor++
		}
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "enter":
		if len(m.records) > 0 {
			m.mode = modePalette
			m.paletteAt = m.nextSkillIndex()
		}
	case "c":
		if cmd := m.composeSelected(m.skillAt(m.nextSkillIndex())); cmd != "" {
			m.copyOrNotice(cmd)
		}
	case "C":
		if m.agent == "claude" {
			m.agent = "codex"
		} else {
			m.agent = "claude"
		}
		_ = launch.SavePrefs(m.root, launch.Prefs{Agent: m.agent})
	case "s":
		if len(m.records) > 0 {
			m.cycleLifecycle()
		}
	case "r":
		m.records, _ = m.scanner.Scan()
	}
	return m, nil
}

func (m Model) updatePalette(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "q":
		m.mode = modeNormal
	case "down", "j":
		if m.paletteAt < len(m.skills)-1 {
			m.paletteAt++
		}
	case "up", "k":
		if m.paletteAt > 0 {
			m.paletteAt--
		}
	case "c":
		cmd := m.composeSelected(m.skillAt(m.paletteAt))
		m.mode = modeNormal
		m.copyOrNotice(cmd)
	case "enter":
		cmd := m.composeSelected(m.skillAt(m.paletteAt))
		m.mode = modeNormal
		return m, m.spawn(cmd)
	}
	return m, nil
}

func (m *Model) cycleLifecycle() {
	cur := m.records[m.cursor].Lifecycle
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
	m.records[m.cursor].Lifecycle = next
	appDir := m.records[m.cursor].Paths[0]
	tr, _ := scan.ReadTracker(appDir)
	tr.Lifecycle = next
	_ = scan.WriteTracker(appDir, tr)
	m.records[m.cursor].Started = true
}

func (m Model) nextSkillIndex() int {
	if len(m.records) == 0 {
		return 0
	}
	next := m.records[m.cursor].NextSkill
	for i, s := range m.skills {
		if s.Name == next {
			return i
		}
	}
	return 0
}

func (m Model) skillAt(i int) model.SkillSpec {
	if i < 0 || i >= len(m.skills) {
		return model.SkillSpec{}
	}
	return m.skills[i]
}

func (m Model) composeSelected(spec model.SkillSpec) string {
	if spec.Name == "" || len(m.records) == 0 {
		return ""
	}
	return launch.Compose(spec, m.records[m.cursor])
}

func (m *Model) copyOrNotice(cmd string) {
	if cmd == "" {
		return
	}
	if launch.Copy(cmd) {
		m.notice = "Copied: " + cmd
	} else {
		m.notice = "No clipboard tool. Copy manually:\n  " + cmd
	}
	m.mode = modeNotice
}

// spawn suspends the TUI to run the interactive agent with cmd as its prompt,
// then triggers a rescan when it returns.
func (m Model) spawn(cmd string) tea.Cmd {
	if cmd == "" {
		return nil
	}
	if _, err := exec.LookPath(m.agent); err != nil {
		// Agent not installed: fall back to clipboard.
		mm := m
		mm.copyOrNotice(cmd)
		return func() tea.Msg { return rescanMsg{} }
	}
	c := exec.Command(m.agent, cmd)
	return tea.ExecProcess(c, func(error) tea.Msg { return rescanMsg{} })
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("┌ JobOps · Apps  (agent: %s) ─[%d]\n", m.agent, len(m.records)))
	b.WriteString(RenderTable(m.records, m.cursor, m.width))
	b.WriteString("├─────────────\n")
	if len(m.records) > 0 {
		b.WriteString(RenderDetail(m.records[m.cursor]))
	}
	b.WriteString("├─────────────\n")
	switch m.mode {
	case modePalette:
		b.WriteString(m.renderPalette())
	case modeNotice:
		b.WriteString(m.notice + "\n[any key] dismiss\n")
	default:
		b.WriteString("↑↓ move  ↵ run  c copy  C agent  s status  r rescan  q quit\n")
	}
	return b.String()
}

func (m Model) renderPalette() string {
	var b strings.Builder
	b.WriteString("Run on " + m.records[m.cursor].Title + ":\n")
	for i, s := range m.skills {
		cursor := "  "
		if i == m.paletteAt {
			cursor = "› "
		}
		b.WriteString(fmt.Sprintf("%s%s\n", cursor, s.Label))
	}
	b.WriteString("↵ spawn " + m.agent + "  c copy  esc cancel\n")
	return b.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run (from `tools/dashboard/`): `go test ./internal/ui/`
Expected: PASS (view + app tests).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/internal/ui/app.go tools/dashboard/internal/ui/app_test.go
git commit -m "feat(dashboard): Bubble Tea model — keys, palette, spawn, status cycle"
```

---

## Task 12: main.go — workspace discovery, config screen, wiring

**Files:**
- Modify: `tools/dashboard/main.go` (replace the Task 1 stub)

- [ ] **Step 1: Replace the stub**

Overwrite `tools/dashboard/main.go`:
```go
// Command jobops-dash is a terminal dashboard for tracking job applications and
// launching /jobops skills.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/launch"
	"github.com/reggiechan74/jobops-dashboard/internal/scan"
	"github.com/reggiechan74/jobops-dashboard/internal/ui"
)

func main() {
	wsFlag := flag.String("workspace", "", "path to the JobOps workspace (default: discover from CWD)")
	flag.Parse()

	start := *wsFlag
	if start == "" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot determine working directory:", err)
			os.Exit(1)
		}
		start = cwd
	}

	root, err := config.Discover(start)
	if err != nil {
		// Not configured: offer to run setup.
		runUnconfigured(start)
		return
	}

	cfg, err := config.Load(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read .jobops/config.json:", err)
		os.Exit(1)
	}

	prefs := launch.LoadPrefs(root)
	model := ui.New(root, prefs.Agent, scan.AppsAdapter{Cfg: cfg})

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "dashboard error:", err)
		os.Exit(1)
	}
}

// runUnconfigured prints guidance and, if claude is available, offers to run
// /jobops:setup directly.
func runUnconfigured(start string) {
	fmt.Println("No .jobops workspace found at or above:", start)
	fmt.Println()
	if _, err := exec.LookPath("claude"); err == nil {
		fmt.Print("Run /jobops:setup now with claude? [y/N] ")
		var ans string
		fmt.Scanln(&ans)
		if ans == "y" || ans == "Y" {
			c := exec.Command("claude", "/jobops:setup")
			c.Stdin, c.Stdout, c.Stderr = os.Stdin, os.Stdout, os.Stderr
			_ = c.Run()
			return
		}
	}
	fmt.Println("Run `/jobops:setup` in Claude Code to initialize the workspace, then re-run jobops-dash.")
}
```

- [ ] **Step 2: Verify the whole module builds and all tests pass**

Run (from `tools/dashboard/`):
```bash
go build ./...
go vet ./...
go test ./...
```
Expected: build OK, vet clean, all tests PASS.

- [ ] **Step 3: Manual smoke test against the repo's own workspace**

Run (from `/home/reggiechan/JobOps`):
```bash
go run ./tools/dashboard --workspace "$PWD"
```
Expected: the dashboard renders. If `.jobops/` doesn't exist in this repo, it prints the unconfigured guidance instead — that is correct behavior. Press `q` to quit. (Create a throwaway workspace to see records: `mkdir -p /tmp/ws/.jobops && echo '{"directories":{"job_postings":"./Job_Postings","applications_root":"./Applications","company_intelligence":"./Company_Intelligence"}}' > /tmp/ws/.jobops/config.json && mkdir -p /tmp/ws/Job_Postings && touch /tmp/ws/Job_Postings/Acme_PM_20260531.md && go run ./tools/dashboard --workspace /tmp/ws`.)

- [ ] **Step 4: Commit**

```bash
cd /home/reggiechan/JobOps
git add tools/dashboard/main.go
git commit -m "feat(dashboard): main entry — discovery, config screen, program wiring"
```

---

## Task 13: Integration — npm script, CI, README

**Files:**
- Modify: `package.json:5-10` (scripts block)
- Create: `tools/dashboard/README.md`
- Modify: CI workflow if one exists (search first)

- [ ] **Step 1: Add the `dash` npm script**

In `package.json`, add a `"dash"` entry to the `scripts` object (keep existing entries):
```json
"scripts": {
  "install-browsers": "npx playwright install chrome",
  "install-all": "npm install && npm run install-browsers",
  "validate:codex-plugin": "node scripts/validate/validate-codex-plugin-compatibility.js",
  "test": "npm run validate:codex-plugin",
  "dash": "go run ./tools/dashboard"
}
```

- [ ] **Step 2: Write the dashboard README**

Create `tools/dashboard/README.md`:
```markdown
# jobops-dash

Terminal dashboard for tracking JobOps applications and launching `/jobops`
skills. Phase 1 ships the Apps tab.

## Run

From your JobOps workspace (the directory containing `.jobops/`):

```bash
npm run dash               # via the repo
# or build a binary:
go build -o jobops-dash ./tools/dashboard && ./jobops-dash
```

Use `--workspace <path>` to point at a workspace explicitly.

## Keys

`↑↓`/`jk` move · `↵` open skill palette · `c` copy next-step command ·
`C` toggle claude/codex · `s` cycle lifecycle status · `r` rescan · `q` quit.

Pressing `↵` and selecting a skill spawns an interactive `claude` (or `codex`)
session with the command pre-filled; on exit the dashboard rescans. This uses
your normal interactive session (no headless API cost).

## State the dashboard owns

- `Applications/{slug}/.tracker.json` — per-app lifecycle status.
- `.jobops/dashboard.json` — UI prefs (selected agent).

Neither is written by JobOps skills.
```

- [ ] **Step 3: Add Go CI if a workflow exists**

Run: `ls .github/workflows/ 2>/dev/null`
- If a workflow file exists, add these steps to it (after checkout), matching its existing YAML style:
```yaml
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - name: Go vet & test (dashboard)
        working-directory: tools/dashboard
        run: |
          go vet ./...
          go test ./...
```
- If `.github/workflows/` does not exist, skip this step (no CI to extend) and note it in the commit message.

- [ ] **Step 4: Verify nothing else broke**

Run (from `/home/reggiechan/JobOps`):
```bash
npm test
claude plugin validate plugins/jobops
(cd tools/dashboard && go test ./...)
```
Expected: Codex compat passes, plugin validates, all Go tests PASS.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add package.json tools/dashboard/README.md .github/workflows/ 2>/dev/null
git commit -m "feat(dashboard): npm dash script, README, Go CI step"
```

---

## Self-Review

**1. Spec coverage** (against `2026-05-31-jobops-dashboard-tui-design.md`):
- §2 no-headless / spawn interactive → Task 11 `spawn` via `tea.ExecProcess`; copy fallback via Task 9 `Copy` + Task 11 `copyOrNotice`; agent toggle `C` + persistence → Task 9 `Prefs`, Task 11. ✓
- §3.1 config loader/discovery → Task 3; scanner/frontmatter → Task 4; record model → Task 2; launcher → Tasks 8–9; TUI shell → Tasks 10–11. ✓
- §4.2 Apps discovery union, slug parse, pipeline stages, fit score, OSINT cross-ref → Tasks 5, 7. ✓
- §4.4 lifecycle `.tracker.json` atomic, default Interested, backlog status-change creates folder → Task 6 + Task 11 `cycleLifecycle`. ✓
- §5 command composition, arg inference, spawn, clipboard tools, agent toggle/prefs → Tasks 8, 9, 11. ✓
- §6 master-detail layout, palette, keys → Tasks 10, 11. ✓
- §7 project layout, `dash` script, CI, prefs file → Tasks 1, 13. ✓
- §8 scanner/adapter table tests, launcher tests, TUI teatest-style flow tests, error handling (missing config, malformed frontmatter `⚠`, missing agent → clipboard, missing clipboard → modal) → Tasks 3–12. ✓
- §9 Apps-first phasing → this plan is Phase 1 only. ✓
- §10 out-of-scope items not implemented. ✓

**2. Placeholder scan:** No TBD/TODO; every code step contains complete code; every test step has real assertions. The one prose note (Task 5 `var _ = filepath.Join`) gives an explicit remediation. ✓

**3. Type consistency:** `model.Record`/`Stage`/`Lifecycle`/`SkillSpec`/`ArgStyle` defined in Task 2 are used unchanged in Tasks 5–11. `AppsAdapter{Cfg: ...}` field name is consistent (Tasks 7, 12). `scan.ReadTracker`/`WriteTracker` signatures match between Task 6 and Task 11. `launch.Compose`/`Copy`/`LoadPrefs`/`SavePrefs`/`Prefs{Agent}` consistent across Tasks 8, 9, 11, 12. UI `Scanner` interface (Task 11) is satisfied by `scan.AppsAdapter` (Tasks 7, 12). ✓

Note for the implementer: when teatest is needed for richer flows in a later phase, the dependency is already added (Task 1); Phase 1 tests the model directly via `Update`/`View`, which is sufficient.
