# JobOps Dashboard TUI — Design

**Date:** 2026-05-31
**Status:** Approved (pending implementation plan)
**Author:** Reggie Chan (with Claude)

## 1. Purpose

A terminal dashboard for tracking job applications and launching JobOps skills.
Two problems it solves:

1. **Tracking** — JobOps produces a rich `Applications/{slug}/` tree per
   application, but there is no at-a-glance view of *which applications exist,
   how far each has progressed through the JobOps pipeline, their fit score, and
   their real-world lifecycle status* (applied / interviewing / offer / etc.).
2. **Launching** — JobOps skills are `/jobops:*` slash commands that only run
   inside an interactive Claude Code (or Codex) session. There is no way to go
   from "this application needs a cover letter next" to running that skill
   without manually typing the command and its arguments.

The dashboard reads the JobOps workspace to present the tracking view, and
launches skills by spawning an interactive `claude`/`codex` session with the
command pre-composed.

## 2. Key constraints and decisions

These were settled during brainstorming and are load-bearing for the design:

- **No headless execution.** `claude -p` (print/headless mode) is excluded — the
  user treats it as incurring pay-per-token API cost. The dashboard instead
  **spawns an interactive `claude` session**, which behaves exactly like opening
  Claude Code normally (same auth/billing as the user's existing usage) and lets
  skills ask follow-up questions.
- **Run action = spawn interactive Claude (primary) + copy-to-clipboard
  (fallback).** Enter spawns; `c` copies the command for when the user is already
  inside a Claude session; `C` toggles the spawn target between `claude` and
  `codex`. After a spawned session exits, the dashboard rescans.
- **Track both artifact progress and real-world lifecycle.** Artifact progress is
  derived for free from the folder tree + YAML frontmatter. Lifecycle status
  (interested/applied/interviewing/…) is **new state the dashboard owns**, stored
  per-application.
- **Runtime: Go + Bubble Tea** (Charm stack: `bubbletea` + `lipgloss` +
  `bubbles`). Single cross-compiled binary, no runtime dependency for end users,
  `tea.ExecProcess` is purpose-built for suspending the TUI to run an interactive
  child process, and `os/exec` + frontmatter parsing fit the task cleanly.
- **Layout: master-detail + action palette.** A scrollable application table on
  top, a detail pane for the selected record below, a key-hint status bar, and a
  modal skill-picker palette on Enter.
- **Scope: all five workflow domains as tabs**, but built on **one shared engine
  with five thin view adapters**, and **delivered Apps-tab-first** (see §9
  Phasing). The spec describes all five; the implementation plan sequences them.

## 3. Architecture

### 3.1 Shared core engine (built once)

- **Config loader** — reads `.jobops/config.json`. Resolves the workspace from
  `$PWD` upward (nearest ancestor containing `.jobops/`), overridable with
  `--workspace=<path>`. If no config is found, the app shows a "not configured"
  screen offering to spawn `/jobops:setup`.
- **Scanner** — given a data root, walks it and parses YAML frontmatter from
  markdown files (`adrg/frontmatter` with a `yaml.v3` unmarshaller) into typed
  records. Tolerant: a file with missing/malformed frontmatter still produces a
  record, flagged `⚠`, rather than aborting the scan.
- **Record model** — a generic record (see §4) carrying title, subtitle, optional
  status, optional score, an ordered artifact/stage checklist, last-modified
  time, and the source path(s). Tab adapters populate it.
- **Launcher** — composes a `/jobops:<skill> <args>` command from the selected
  record, then either spawns the configured agent CLI via `tea.ExecProcess`
  (Enter) or copies the command to the clipboard (`c`). See §5.
- **TUI shell** — the Bubble Tea root model: tab bar, master table, detail pane,
  status/help bar, and the modal launcher palette. Built from `bubbles`
  components (`table`, `list`, `help`, `textinput` for filter) styled with
  `lipgloss`.

### 3.2 Tab adapters (data root + record-shaping rule + skill menu)

| Tab | Data root(s) | Record granularity | Lifecycle? | Launch menu (skills) |
|-----|--------------|--------------------|------------|----------------------|
| **Apps** | `Job_Postings/*.md` ∪ `Applications/{slug}/` | one per slug | yes (`.tracker.json`) | auditjobposting, assessjob, buildresume, coverletter, osint, briefing, interviewprep, convert-to-pdf, convert-to-word |
| **Companies** | `Company_Intelligence/{Company}/` | one per company | no | osint |
| **Career** | `Career_Analysis/` | one per timestamped file | no | idealjob, comparejobs, change-one-thing |
| **Crisis** | `Crisis_Management/` | one per timestamped file | no | severance-review, non-compete-analysis, code-red, constructive-dismissal, accommodation-request, discrimination-assessment, reference-shield, unemployment-prep, workplace-documentation, layoff-intel, investigation-prep |
| **Contractor** | `Contractor/` (jobops-ic) | one per output file/folder | no | defineservices, findclient, pitchdeck, proposaltemplate, ratecard, create-landing-page |

Each adapter implements a small interface (see §4.3) so the shell treats all tabs
uniformly. Data-root keys come from `config.directories.*`, never hardcoded paths.

## 4. Data model

### 4.1 Generic record

```go
type Record struct {
    Title     string        // e.g. "Acme — Product Manager" or company / file name
    Subtitle  string        // e.g. applied date, role variant, timestamp
    Slug      string        // app slug or company key; "" for flat files
    Score     *int          // normalized fit %, when available (Apps)
    Lifecycle Lifecycle     // Apps only; zero value = Interested
    Stages    []Stage       // ordered pipeline/artifact checklist
    NextSkill string        // first incomplete stage's skill; "" if complete
    Updated   time.Time     // most-recent artifact mtime
    Paths     []string      // source files/dirs backing this record
    Warn      string        // non-empty if frontmatter was malformed
}

type Stage struct {
    Name   string
    Skill  string            // /jobops skill that produces it
    State  StageState        // Missing | Draft | Final
}
```

### 4.2 Apps adapter specifics

- **Discovery:** union by slug. `Job_Postings/*.md` filename stems and
  `Applications/*/` folder names are both slug sources. A slug present only in
  `Job_Postings/` is **backlog** (unstarted). A slug present only in
  `Applications/` (original posting moved away) still appears via the pinned
  `job_posting.md`.
- **Slug parse:** `{Company}_{Role}_{YYYYMMDD}` → Company, Role, Date for display.
  A `Job_Postings` file whose name does not parse is still listed (Company =
  filename, Role/Date empty) and flagged `⚠`.
- **Pipeline stages (ordered):** JD audit → Assess → Resume → Cover → OSINT →
  Briefing → Prep. Stage state:
  - **Final** (`✓`) — the stage's primary artifact exists with frontmatter
    `status: final` (or, for multi-file stages like resume, `step3_final.md`
    present).
  - **Draft** (`◐`) — artifact exists but `status: draft`, or only intermediate
    files exist (resume step1/2 without step3).
  - **Missing** (`○`) — no artifact.
- **Fit score:** `Applications/{slug}/assessment/assessment.md` frontmatter
  `normalized_score` (fallback: compute from `overall_score`). Drives the Fit %
  column and detail-pane grade.
- **OSINT stage** is cross-referenced: an app's OSINT stage reads
  `Company_Intelligence/{Company}/` (shared across applications to the same
  company), not a per-app folder.
- **Next step** = first stage not `Final`; its `Skill` preselects the launcher.

### 4.3 Adapter interface

```go
type TabAdapter interface {
    Name() string                       // tab label
    Scan(cfg Config) ([]Record, error)  // build records from this tab's root(s)
    Skills() []SkillSpec                // launch menu for this tab
    Compose(s SkillSpec, r Record) string // -> "/jobops:<skill> <args>"
}
```

The Apps adapter is the only one with lifecycle and a multi-stage pipeline; the
others return artifact-only `Stages` (or none) and an empty `Lifecycle`.

### 4.4 Lifecycle state (Apps only)

Stored at `Applications/{slug}/.tracker.json`, **owned and written by the
dashboard**. Co-located with the app so it survives folder rename/move; a dotfile
so it does not clutter listings. Written atomically (write `*.tmp`, then `rename`)
to match the config.json convention.

```json
{
  "lifecycle": "interviewing",
  "applied_on": "2026-05-20",
  "notes": "recruiter Jane; panel 06-03",
  "updated_at": "2026-05-29T14:02:00Z"
}
```

- **Enum:** `interested → applied → screening → interviewing → offer →
  accepted | rejected | withdrawn`. Default `interested` when the file is absent.
- The pinned `job_posting.md` is **never** mutated to hold lifecycle state (it is
  contractually immutable).
- Setting a status on a **backlog** posting (no `Applications/` folder yet)
  creates `Applications/{slug}/` and writes `.tracker.json` there. It does **not**
  pin the JD or run any skill — that only happens when a skill is launched.

## 5. Launcher

- **Compose:** each `SkillSpec` defines how to build its command from a record.
  The adapter fills inferable args — the JD filename (`{{ARG1}}`) and `--app=<slug>`
  for per-app skills. Args a skill still needs (e.g. a non-default resume source)
  are left for the user to answer **interactively in the spawned session**, so
  partial knowledge never blocks a launch.
- **Spawn (Enter):** `tea.ExecProcess(exec.Command(agent, composed), ...)` where
  `agent` is `claude` or `codex`. Bubble Tea releases the terminal, the agent runs
  interactively, and on exit the dashboard resumes and rescans the affected
  record's root.
- **Copy (`c`):** copy the composed command to the clipboard via the first
  available of `pbcopy` (macOS), `wl-copy` (Wayland), `xclip`/`xsel` (X11),
  `clip.exe` (WSL/Windows). If none is available, show the command in a modal the
  user can select/copy manually.
- **Agent toggle (`C`):** flips the spawn target between `claude` and `codex`;
  the choice persists in a dashboard preferences file (`.jobops/dashboard.json`,
  also dashboard-owned) so it survives restarts.

## 6. UX / navigation

**Layout (master-detail + palette):**

```
┌ JobOps · Apps | Companies | Career | Crisis | Contractor ───[12]┐
│ Company      Role          Fit  Lifecycle    Pipeline           │
│▶Acme         Product Mgr   84%  ● Interview   ███████░          │
│ Beta Capital Analyst       71%  ○ Applied     █████░░░          │
│ Gamma Fund   Assoc Dir     —    · Interested  ██░░░░░░          │
├ Acme — Product Manager ─────────────────────────────────────────┤
│ Fit 84% (B+)   Applied 2026-05-20                               │
│ JD✓ Assess✓ Resume✓ Cover✓ OSINT✓ Briefing✓ Prep○              │
│ Next: /jobops:interviewprep                                     │
│ Recent: assessment.md · 05-29   cover_letter.md · 05-28         │
├─────────────────────────────────────────────────────────────────┤
│ ↑↓ move  ↵ run  c copy  C agent  s status  o open  / filter  ?  │
└─────────────────────────────────────────────────────────────────┘
```

**Keys:** `↑↓`/`jk` move · `←→`/`tab`/`shift-tab` switch tabs · `↵` open launcher
palette · `c` copy next-step command · `C` toggle claude/codex · `s` set lifecycle
status (Apps) · `o` open selected artifact in `$PAGER`/`$EDITOR` · `/` filter list
· `r` rescan · `?` toggle help · `q` quit.

**Launcher palette (`↵`):** modal list of the current tab's skills, cursor
pre-positioned on the computed next step. Selecting runs (spawn) or `c` copies.

## 7. Project layout, build & distribution

Standalone Go module under `tools/dashboard/`, separate from the plugin trees
(which remain pure markdown skills):

```
JobOps/
  tools/dashboard/
    go.mod  go.sum
    main.go                       # load config, start Bubble Tea
    internal/
      config/   config.go         # read .jobops/config.json, workspace discovery
      scan/     scan.go           # walk roots, parse frontmatter -> records
                apps.go           # Apps adapter (slug union, pipeline, tracker)
                companies.go career.go crisis.go contractor.go
      model/    record.go         # Record, Stage, Lifecycle, SkillSpec
      launch/   launch.go         # compose, ExecProcess, clipboard, agent toggle
      ui/       app.go table.go detail.go palette.go statusbar.go help.go
    testdata/                     # fixture .jobops workspaces for tests
    README.md
```

- **Build:** `go build -o jobops-dash ./tools/dashboard`. A `.goreleaser.yaml`
  cross-compiles macOS (arm64/amd64), Linux (amd64/arm64), Windows.
- **Entry point:** add `"dash": "go run ./tools/dashboard"` to `package.json`
  scripts so existing users have one obvious command; document the prebuilt binary
  in the dashboard README.
- **Isolation:** the new `tools/` directory does not affect `npm test` or
  `claude plugin validate` (plugin trees untouched). CI gains a `go test ./...` +
  `go vet ./...` step.
- **Dashboard preferences:** `.jobops/dashboard.json` (dashboard-owned) stores the
  agent toggle and any future UI prefs; never written by JobOps skills.

## 8. Testing & error handling

**Testing (correctness lives in the scanner/adapters, so they get the most
coverage):**
- Table-driven Go tests for each adapter against `testdata/` fixture workspaces:
  backlog-only posting, partial pipeline, full pipeline, unparseable slug,
  missing/malformed frontmatter, multi-app same-company OSINT sharing, empty
  workspace.
- Launcher tests assert composed command strings per skill/record without
  executing; the spawn and clipboard operations sit behind interfaces and are
  faked in tests.
- TUI flows tested lightly with Bubble Tea's `teatest` (navigate, switch tab,
  open palette, set status); not pixel-exhaustive.

**Error handling:**
- No `.jobops/` found → config screen offering to spawn `/jobops:setup`.
- Malformed frontmatter → record still listed with a `⚠` flag; never crashes the
  scan.
- `claude`/`codex` not on `PATH` → Enter falls back to copy-to-clipboard with a
  notice.
- No clipboard utility → show the command in a selectable modal.
- Empty/missing data root for a tab → tab shows an empty-state hint, not an error.

## 9. Phasing (delivery order for the implementation plan)

The spec covers all five tabs; the implementation plan builds them in this order
so value lands early and the engine is validated before replication:

- **Phase 1 — Core engine + Apps tab.** Config loader, scanner, generic record,
  launcher (spawn + copy + agent toggle), TUI shell, and the full Apps adapter
  (backlog ∪ applications, pipeline derivation, fit score, `.tracker.json`
  lifecycle). Proves the entire pattern end-to-end.
- **Phase 2 — Companies tab.** Artifact-only adapter over
  `Company_Intelligence/{Company}/`; simplest second proof of the adapter
  interface.
- **Phase 3 — Career, Crisis, Contractor tabs.** Flat timestamped roots with
  near-identical artifact-only adapters and per-tab skill menus.

## 10. Out of scope (v1)

- Editing artifact contents in the TUI (use `o` to open in `$EDITOR`).
- Headless skill execution / capturing skill output in the TUI.
- Creating or editing job-posting files (the dashboard launches skills that do;
  it does not author JDs itself).
- Analytics/charts beyond the per-app fit score and pipeline progress.
- Multi-workspace switching within a session (one `--workspace` per run).
