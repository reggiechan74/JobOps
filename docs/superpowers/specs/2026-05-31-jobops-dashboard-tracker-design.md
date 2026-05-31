# JobOps Dashboard & Application Tracker — Design

**Date:** 2026-05-31
**Status:** Approved (brainstorming)
**Supersedes:** the deleted Go TUI in `tools/dashboard/` (staged for deletion on this branch)

## 1. Purpose

Replace the compiled Go TUI dashboard with a **skill** that runs inside Claude Code
*and* Codex. The skill renders an on-demand application statusboard and, on Claude
Code, drives an interactive "navigate → act" loop. Status is backed by a single
YAML tracker file under the user's workspace.

The Go binary required a separate toolchain and could not be invoked by the agent
mid-conversation. A skill removes the binary, works on both platforms, and can chain
directly into other JobOps skills.

## 2. Decisions (locked during brainstorming)

| Question | Decision |
|---|---|
| Source of truth | **Hybrid** — filesystem for artifact existence, YAML for human-only status |
| Interaction | **Two-level menu loop** — pick an application, then pick an action |
| Cross-platform | **Claude Code first** (full loop via `AskUserQuestion`); **Codex read-only** board |
| Scope | **Applications + linked company-intel** status (not career/crisis/contractor) |
| Sync | **Dashboard reconciles on each run** — the 35 existing skills are never modified |

## 3. Components

- **New skill:** `plugins/jobops/skills/dashboard/SKILL.md`. The *only* reader/writer of the tracker.
- **New config key:** `directories.application_tracker`, default `./Applications/tracker.yaml`.
- **Setup/migrate wiring:** `/jobops:setup` adds the key; `/jobops:migrate` backfills it idempotently.

The skill is plain skill markdown — the agent performs the scan via Read/Glob/Bash and
writes YAML via Write. There is no compiled component.

## 4. Tracker file

Location: `directories.application_tracker` (default `./Applications/tracker.yaml`),
co-located with the application folders it describes, gitignored with the rest of the
user's career data.

Schema (two clearly separated zones per application):

```yaml
version: 1
generated: 2026-05-31T14:02:00Z        # set every reconcile (UTC ISO-8601)
applications:
  - slug: Acme_SeniorPM_20260520        # = folder name, the stable key
    company: Acme
    role: Senior PM
    # ── human status: PRESERVED across reconciles, never auto-overwritten ──
    stage: interviewing                  # lead|applied|interviewing|offer|rejected|withdrawn|on_hold|archived
    applied_date: 2026-05-21
    next_deadline: 2026-06-02
    contact: "Jane Doe <jane@acme.com>"
    outcome: null
    notes: "Recruiter screen went well"
    # ── artifact flags: OVERWRITTEN every reconcile from the filesystem scan ──
    artifacts:
      assessment: true
      resume_draft: true                 # step1_draft.md OR step2_provenance.md present
      resume_final: true                 # step3_final.md present
      cover_letter: true
      osint: true                        # Company_Intelligence/{Company}/ folder exists
      briefing: false
      interview_prep: false
    next_action: briefing                # computed each reconcile
```

**Zone contract:** the human-status zone is owned by the user (edited only via the
dashboard's *Update status* action). The `artifacts` block, `next_action`, and
top-level `generated` are owned by the reconcile pass and overwritten every run.

## 5. Reconcile algorithm

Runs at the top of every dashboard invocation:

1. Load existing YAML (or start with an empty `applications: []`).
2. Scan `applications_root` for `{Company}_{Role}_{YYYYMMDD}/` folders. For each,
   probe the fixed sub-paths to set the `artifacts` block:
   | Flag | True when… |
   |---|---|
   | `assessment` | `assessment/assessment.md` exists |
   | `resume_draft` | `resume/step1_draft.md` **or** `resume/step2_provenance.md` exists |
   | `resume_final` | `resume/step3_final.md` exists |
   | `cover_letter` | `cover-letter/cover_letter.md` exists |
   | `osint` | `{company_intelligence}/{Company}/` folder exists |
   | `briefing` | any `interview/briefing*.md` exists |
   | `interview_prep` | any `interview/interview_prep*.md` exists |
   > Mapping pinned against each producing skill's output section and the deleted Go
   > scanner's lifecycle logic (`git show 6d9a6de:tools/dashboard/internal/scan/apps.go`).
   > `buildresume` produces all three resume steps internally (`step1-resume-draft` →
   > `step2-provenance-check` → `step3-final-resume`), so a single `buildresume` covers
   > the whole resume stage. `briefing`/`interview_prep` use prefix-glob to catch the
   > multi-part outputs (`interview_prep_part1.md`, …).
3. **Merge rule:**
   - slug already in YAML → keep the human-status zone verbatim; overwrite only `artifacts` + `next_action`.
   - new folder → append a skeleton entry (`stage: applied`, human fields null/empty, `company`/`role` parsed from slug).
   - slug whose folder no longer exists → keep the entry, set `stage: archived` (preserves notes/outcome); artifacts all false.
4. Recompute `next_action` (§6) and `generated`.
5. Write YAML back atomically (`.tmp` then `mv`), consistent with `/jobops:setup`.

## 6. Next-action state machine

First unmet step wins, ordered by the real pipeline (mirrors the Go scanner's proven
`pipelineOrder`, led by assessment):

```
assessment   == false   → assessjob
resume_final == false   → buildresume        (produces draft → provenance → final in one run)
cover_letter == false   → coverletter        (needs step3_final resume)
osint        == false   → osint
briefing     == false   → briefing
interview_prep == false → interviewprep
all artifacts true       → record-outcome     (prompt to set stage offer/rejected + outcome)
```

`archived` apps have `next_action: none`. `resume_draft` is not its own next-action — it
only distinguishes a half-dot (●◐○) in the board; `buildresume` is the action for any
incomplete resume.

## 7. Interaction loop

### 7.1 Claude Code (full loop)

1. Reconcile, then **always render the full text board** (§8) so every app is visible
   regardless of count.
2. **`AskUserQuestion` #1 — navigate.** Options = up-to-3 *most-actionable* apps
   (those with a pending `next_action`, sorted by `next_deadline` ascending), label
   `Company — Role`, description `stage · next: <skill>`. The tool's auto "Other"
   free-text lets the user type a slug or row number for anything outside the top 3.
   This honors the **2–4 option cap** of `AskUserQuestion` while scaling to any board size.
3. **`AskUserQuestion` #2 — act.** Bounded menu (≤4): `Run next: <skill>` ·
   `Update status` · `Open folder` · `Back to board`.
4. **Execute:**
   - *Run next* → invoke the skill via the Skill tool with the app folder as context; on return, reconcile + re-render.
   - *Update status* → follow-up `AskUserQuestion`(s) write human-zone field(s) (stage/date/contact/outcome) to YAML.
   - *Open folder* → print the path and `ls` it.
   - *Back* → redraw the board.
5. Loop until the user selects `Back to board` at the top level and declines to continue.

### 7.2 Codex (read-only)

Reconcile + render the board + a numbered "suggested next actions" list, then stop with
guidance to invoke the suggested skill manually (e.g. `/jobops:buildresume`). Same
reconcile/render path; only the interactive tail is omitted (the skill omits
`AskUserQuestion` when the tool is unavailable).

## 8. Board render

Markdown, renders on both platforms. Summary line, then a table sorted by stage then deadline:

```
JobOps Application Tracker · 5 active · reconciled 2026-05-31 14:02

Stage          Company / Role             Pipeline       Next             Deadline
─────────────  ─────────────────────────  ─────────────  ───────────────  ──────────
interviewing   Acme / Senior PM           ●●●●●○         interviewprep    Jun 2  ⚠
applied        Globex / Director Ops       ●●●◐○○         osint            Jun 5
applied        Initech / PM               ●◐○○○○         buildresume      —
offer          Umbrella / VP Product      ●●●●●●         record-outcome   —
archived       Soylent / Lead PM          ——             —                —

Pipeline dots: assess▸resume▸cover▸osint▸briefing▸prep
```

- Six dots, one per pipeline stage in order: `assess · resume · cover · osint · briefing · prep`.
- `●` = stage complete · `○` = missing · `◐` = resume only (draft present, no `step3_final`).
- `⚠` = `next_deadline` within 3 days of `generated`.

## 9. Setup, config self-heal & docs

- `/jobops:setup` Step 2 directory interview gains one row and Step 6 JSON gains one key:
  `application_tracker → ./Applications/tracker.yaml`.
- **Existing workspaces self-heal** (no `/jobops:migrate` change): if
  `config.directories.application_tracker` is absent, the dashboard defaults it to
  `{applications_root}/tracker.yaml` and persists the key back into `config.json`
  (atomic `.tmp`+`mv`) on first run. Idempotent thereafter.
- The dashboard skill follows the standard contract: frontmatter (`name`,
  `description`, `disable-model-invocation: true`), `## Configuration` preamble reading
  `.jobops/config.json`, and the `JOBOPS NOT CONFIGURED` exit when the config is absent.
- Docs to update: `docs/ARCHITECTURE.md` §4 (note the tracker file), both plugin READMEs,
  CHANGELOG, version bump via `/version-bump`.

## 10. Testing (scenario-based)

No Go test suite remains; verification is fixture-driven. A test workspace with several
`Applications/` folders at different pipeline stages plus a hand-written `tracker.yaml`
asserts:

1. Reconcile preserves human-zone fields (stage/dates/contact/notes/outcome) on an existing slug.
2. Artifact flags match the fixtures exactly.
3. `next_action` is correct for each pipeline stage (every branch of §6).
4. A slug whose folder was removed flips to `stage: archived`, `next_action: none`.
5. A brand-new folder appears as a skeleton entry with parsed company/role.
6. The board renders with correct dots, OSINT marks, and `⚠` deadline flag.
7. `npm test` (manifest/frontmatter contract) still passes with the new skill present.

## 11. Out of scope (YAGNI)

- Career / crisis / contractor tracking (filesystem browse stays manual).
- Multi-user or remote sync.
- Any compiled component or separate toolchain.
- Editing artifacts from the dashboard (it launches the owning skill instead).
