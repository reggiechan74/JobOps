# JobOps Dashboard & Application Tracker Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the deleted Go TUI with a cross-platform JobOps skill (`/jobops:dashboard`) that reconciles a YAML application tracker from the filesystem and drives an interactive "navigate → act" loop on Claude Code (read-only board on Codex).

**Architecture:** A single new skill, `plugins/jobops/skills/dashboard/SKILL.md`, is the only reader/writer of `tracker.yaml`. On each run it scans `Applications/` + `Company_Intelligence/`, merges filesystem-derived artifact flags with preserved human-status fields, recomputes each app's `next_action`, writes the YAML atomically, renders a markdown board, and (on Claude Code) loops `AskUserQuestion` to navigate and act. Setup gains one config key; existing configs self-heal.

**Tech Stack:** Markdown skill prompt (executed by Claude Code / Codex via Read/Glob/Bash/Write tools). YAML tracker file. Node-based contract validator (`npm test`). No compiled component.

**Reference spec:** `docs/superpowers/specs/2026-05-31-jobops-dashboard-tracker-design.md`

---

## Context an implementer needs (read before starting)

- **This is a prompt artifact, not executable code.** "Tests" here are: (a) the Codex contract validator `npm test`, and (b) a committed fixture workspace plus an `expected_tracker.yaml` that an agent reproduces by following the skill's reconcile algorithm, then verifies with `diff`. There is no unit-test runner to add.
- **Slug company/role is not parseable.** `{Company}_{Role}_{YYYYMMDD}` has an unrecoverable Company/Role boundary (both can contain `_`). Company is detected by prefix-matching `Company_Intelligence/` folder names; the rest is the role/title. This logic is specified verbatim in Task 3's skill body and mirrors the deleted `internal/scan/slug.go`.
- **Artifact→file mapping** (pinned against real skill outputs + `git show 6d9a6de:tools/dashboard/internal/scan/apps.go`):
  | Flag | True when |
  |---|---|
  | `assessment` | `assessment/assessment.md` exists |
  | `resume_draft` | `resume/step1_draft.md` **or** `resume/step2_provenance.md` exists |
  | `resume_final` | `resume/step3_final.md` exists |
  | `cover_letter` | `cover-letter/cover_letter.md` exists |
  | `osint` | `{company_intelligence}/{Company}/` folder exists |
  | `briefing` | any `interview/briefing*.md` exists |
  | `interview_prep` | any `interview/interview_prep*.md` exists |
- **Config validator rules** (`scripts/validate/validate-codex-plugin-compatibility.js`): every `skills/<dir>/SKILL.md` must have frontmatter `name == <dir>`, a non-empty `description`, and `disable-model-invocation: true`. New skill dirs are auto-discovered — no manifest edit needed.

---

## File Structure

- **Create** `plugins/jobops/skills/dashboard/SKILL.md` — the entire dashboard skill (config self-heal, reconcile, state machine, board render, interaction loop, Codex degradation).
- **Create** `test/fixtures/dashboard/workspace/` — committed fixture: `.jobops/config.json`, `Applications/<several apps at different stages>/…`, `Company_Intelligence/<one company>/…`, and a pre-existing `Applications/tracker.yaml` carrying human-status fields to prove the merge.
- **Create** `test/fixtures/dashboard/expected_tracker.yaml` — the exact YAML a correct reconcile must produce against the fixture.
- **Create** `test/fixtures/dashboard/README.md` — the verification scenario and `diff` command.
- **Modify** `plugins/jobops/skills/setup/SKILL.md` — add the `application_tracker` row (Step 2 table) and key (Step 6 JSON).
- **Modify** `docs/ARCHITECTURE.md` — note the tracker file in §4.
- **Modify** `plugins/jobops/README.md` and root `README.md` — document `/jobops:dashboard`.
- **Modify** `CHANGELOG.md`, `package.json`, `.claude-plugin/marketplace.json`, plugin.json files — version bump via the version-bump skill.

---

## Task 1: Add `application_tracker` to setup config

**Files:**
- Modify: `plugins/jobops/skills/setup/SKILL.md` (Step 2 table ~line 45-51; Step 6 JSON ~line 140-147)

- [ ] **Step 1: Add the directory row to the Step 2 interview table**

In `plugins/jobops/skills/setup/SKILL.md`, find the Step 2 table and add a row after the `crisis_management` row:

```markdown
| `application_tracker` | `./Applications/tracker.yaml` | Application status tracker maintained by `/jobops:dashboard` (one YAML file, not a folder) |
```

- [ ] **Step 2: Add the key to the Step 6 config JSON**

In the same file, in the Step 6 JSON block, change the `directories` object so it reads (note the added trailing comma on `crisis_management`):

```json
  "directories": {
    "resume_source": "<step-2 value>",
    "job_postings": "<step-2 value>",
    "applications_root": "<step-2 value>",
    "company_intelligence": "<step-2 value>",
    "career_analysis": "<step-2 value>",
    "crisis_management": "<step-2 value>",
    "application_tracker": "<step-2 value>"
  },
```

- [ ] **Step 3: Note the path-normalization exception**

The other directory values are folders; `application_tracker` is a file path. Directly under the Step 2 table, add a sentence so setup doesn't `mkdir` it:

```markdown
> `application_tracker` is a **file** path, not a directory — Step 3 must not `mkdir` it. Its parent directory (`applications_root`) is already created, so no extra action is needed.
```

- [ ] **Step 4: Verify the contract validator still passes**

Run: `npm test`
Expected: PASS — output ends with the existing success summary; no new failures. (Setup edits don't touch frontmatter, so this confirms nothing regressed.)

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/skills/setup/SKILL.md
git commit -m "feat(setup): add application_tracker config key for dashboard"
```

---

## Task 2: Build the fixture workspace and expected output (the test)

This task creates the regression fixture **before** the skill exists, so Task 3 has a concrete target. Build it with a script so it is reproducible.

**Files:**
- Create: `test/fixtures/dashboard/workspace/.jobops/config.json`
- Create: `test/fixtures/dashboard/workspace/Applications/...` (four app folders)
- Create: `test/fixtures/dashboard/workspace/Company_Intelligence/Acme/summary.md`
- Create: `test/fixtures/dashboard/workspace/Applications/tracker.yaml` (pre-existing, with human-status fields)
- Create: `test/fixtures/dashboard/expected_tracker.yaml`
- Create: `test/fixtures/dashboard/README.md`

- [ ] **Step 1: Write a fixture-builder script and run it**

Create `test/fixtures/dashboard/build-fixture.sh` with exactly:

```bash
#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")" && pwd)/workspace"
rm -rf "$ROOT"
mkdir -p "$ROOT/.jobops"

cat > "$ROOT/.jobops/config.json" <<'JSON'
{
  "version": "2.0",
  "directories": {
    "resume_source": "./ResumeSourceFolder",
    "job_postings": "./Job_Postings",
    "applications_root": "./Applications",
    "company_intelligence": "./Company_Intelligence",
    "career_analysis": "./Career_Analysis",
    "crisis_management": "./Crisis_Management"
  }
}
JSON

# App A — Acme: assessment+resume(final)+cover+osint done; briefing/prep missing -> next: briefing
A="$ROOT/Applications/Acme_SeniorPM_20260520"
mkdir -p "$A/assessment" "$A/resume" "$A/cover-letter" "$A/interview"
echo x > "$A/assessment/assessment.md"
echo x > "$A/resume/step1_draft.md"
echo x > "$A/resume/step2_provenance.md"
echo x > "$A/resume/step3_final.md"
echo x > "$A/cover-letter/cover_letter.md"

# App B — Globex: assessment+resume(draft only); no cover -> resume_final false -> next: buildresume
B="$ROOT/Applications/Globex_DirectorOps_20260524"
mkdir -p "$B/assessment" "$B/resume"
echo x > "$B/assessment/assessment.md"
echo x > "$B/resume/step1_draft.md"

# App C — Initech: empty pipeline -> next: assessjob
C="$ROOT/Applications/Initech_PM_20260528"
mkdir -p "$C"

# App D — Umbrella: everything done -> next: record-outcome
D="$ROOT/Applications/Umbrella_VPProduct_20260510"
mkdir -p "$D/assessment" "$D/resume" "$D/cover-letter" "$D/interview"
echo x > "$D/assessment/assessment.md"
echo x > "$D/resume/step3_final.md"
echo x > "$D/cover-letter/cover_letter.md"
echo x > "$D/interview/briefing.md"
echo x > "$D/interview/interview_prep_part1.md"

# Company intel for Acme + Umbrella only (Globex/Initech have none)
mkdir -p "$ROOT/Company_Intelligence/Acme" "$ROOT/Company_Intelligence/Umbrella"
echo x > "$ROOT/Company_Intelligence/Acme/summary.md"
echo x > "$ROOT/Company_Intelligence/Umbrella/summary.md"

# Pre-existing tracker with human-status fields to prove the merge preserves them,
# plus a vanished slug (Soylent) that must flip to archived.
cat > "$ROOT/Applications/tracker.yaml" <<'YAML'
version: 1
generated: 2026-05-30T09:00:00Z
applications:
  - slug: Acme_SeniorPM_20260520
    company: Acme
    role: Senior PM
    stage: interviewing
    applied_date: 2026-05-21
    next_deadline: 2026-06-02
    contact: "Jane Doe <jane@acme.com>"
    outcome: null
    notes: "Recruiter screen went well"
    artifacts: {assessment: false, resume_draft: false, resume_final: false, cover_letter: false, osint: false, briefing: false, interview_prep: false}
    next_action: assessjob
  - slug: Soylent_LeadPM_20260415
    company: Soylent
    role: Lead PM
    stage: applied
    applied_date: 2026-04-16
    next_deadline: null
    contact: null
    outcome: null
    notes: "Folder archived offline"
    artifacts: {assessment: true, resume_draft: true, resume_final: true, cover_letter: true, osint: false, briefing: false, interview_prep: false}
    next_action: osint
YAML

echo "fixture built at $ROOT"
```

Run:
```bash
chmod +x test/fixtures/dashboard/build-fixture.sh
./test/fixtures/dashboard/build-fixture.sh
```
Expected: prints `fixture built at …/workspace`.

- [ ] **Step 2: Write the expected reconciled output**

Create `test/fixtures/dashboard/expected_tracker.yaml` with exactly the YAML a correct reconcile must produce. (Entries sorted by `slug` ascending for deterministic comparison; `generated` is normalized to the literal `<RECONCILE_TS>` placeholder so the diff ignores the timestamp.)

```yaml
version: 1
generated: <RECONCILE_TS>
applications:
  - slug: Acme_SeniorPM_20260520
    company: Acme
    role: Senior PM
    stage: interviewing
    applied_date: 2026-05-21
    next_deadline: 2026-06-02
    contact: "Jane Doe <jane@acme.com>"
    outcome: null
    notes: "Recruiter screen went well"
    artifacts:
      assessment: true
      resume_draft: true
      resume_final: true
      cover_letter: true
      osint: true
      briefing: false
      interview_prep: false
    next_action: briefing
  - slug: Globex_DirectorOps_20260524
    company: ""
    role: Globex DirectorOps
    stage: applied
    applied_date: null
    next_deadline: null
    contact: null
    outcome: null
    notes: null
    artifacts:
      assessment: true
      resume_draft: true
      resume_final: false
      cover_letter: false
      osint: false
      briefing: false
      interview_prep: false
    next_action: buildresume
  - slug: Initech_PM_20260528
    company: ""
    role: Initech PM
    stage: applied
    applied_date: null
    next_deadline: null
    contact: null
    outcome: null
    notes: null
    artifacts:
      assessment: false
      resume_draft: false
      resume_final: false
      cover_letter: false
      osint: false
      briefing: false
      interview_prep: false
    next_action: assessjob
  - slug: Soylent_LeadPM_20260415
    company: Soylent
    role: Lead PM
    stage: archived
    applied_date: 2026-04-16
    next_deadline: null
    contact: null
    outcome: null
    notes: "Folder archived offline"
    artifacts:
      assessment: false
      resume_draft: false
      resume_final: false
      cover_letter: false
      osint: false
      briefing: false
      interview_prep: false
    next_action: none
  - slug: Umbrella_VPProduct_20260510
    company: Umbrella
    role: VPProduct
    stage: applied
    applied_date: null
    next_deadline: null
    contact: null
    outcome: null
    notes: null
    artifacts:
      assessment: true
      resume_draft: false
      resume_final: true
      cover_letter: true
      osint: true
      briefing: true
      interview_prep: true
    next_action: record-outcome
```

> Note the deliberate teaching cases: Globex/Initech have no `Company_Intelligence` folder so `company: ""` and the humanized slug becomes the role; Umbrella has `step3_final.md` but no `step1`/`step2`, so `resume_draft: false` while `resume_final: true` (the ◐ half-dot case); Soylent's folder is gone so it flips to `archived`/`next_action: none` while keeping its notes.

- [ ] **Step 3: Write the verification README**

Create `test/fixtures/dashboard/README.md`:

```markdown
# Dashboard reconcile fixture

`build-fixture.sh` generates `workspace/` — a JobOps workspace with four live
application folders (Acme, Globex, Initech, Umbrella) at different pipeline
stages, company intel for Acme/Umbrella, and a pre-existing `tracker.yaml`
that carries human-status fields (Acme) and a vanished slug (Soylent).

## Verify a reconcile

1. `./build-fixture.sh`
2. Following the algorithm in `plugins/jobops/skills/dashboard/SKILL.md` §Reconcile,
   produce `workspace/Applications/tracker.yaml`.
3. Normalize the volatile timestamp and compare:
   ```bash
   sed 's/^generated:.*/generated: <RECONCILE_TS>/' \
     workspace/Applications/tracker.yaml > /tmp/actual_tracker.yaml
   diff -u expected_tracker.yaml /tmp/actual_tracker.yaml
   ```
   Expected: no diff.

This asserts: human-zone preservation (Acme), artifact flags, every next_action
branch, archived flip with notes retained (Soylent), new-folder skeletons
(Globex/Initech/Umbrella), and the resume draft/final distinction (Umbrella).
```

- [ ] **Step 4: Commit (the test, currently unreproducible — no skill yet)**

```bash
git add test/fixtures/dashboard/
git commit -m "test(dashboard): reconcile fixture workspace + expected output"
```

---

## Task 3: Write the dashboard skill

**Files:**
- Create: `plugins/jobops/skills/dashboard/SKILL.md`

- [ ] **Step 1: Create the skill file with full content**

Create `plugins/jobops/skills/dashboard/SKILL.md` with exactly:

````markdown
---
name: dashboard
description: Render an application statusboard from a reconciled YAML tracker and (on Claude Code) drive an interactive navigate-and-act loop over the application pipeline
disable-model-invocation: true
argument-hint: "[--board-only]"
---

# JobOps Application Dashboard

Reconciles `tracker.yaml` from the filesystem, renders a statusboard, and — on
Claude Code — drives a two-level menu loop (pick an application, then pick an
action). On Codex it renders a read-only board with suggested next actions.

**Flags:**
- `--board-only` — reconcile and render the board, then stop (no menu loop) on any platform.

---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.

**Tracker path resolution (self-heal):** read `config.directories.application_tracker`.
If that key is absent, set it to `{config.directories.applications_root}/tracker.yaml`,
write the updated config back atomically (`.jobops/config.json.tmp` then `mv`), and use
that path. This lets workspaces created before the dashboard existed work with no
migration step.

---

## Reconcile

Run this every invocation, before rendering.

1. **Load** the existing tracker YAML at the resolved path, or start with
   `{version: 1, applications: []}` if the file does not exist.
2. **Scan** `config.directories.applications_root` for immediate sub-directories whose
   name is not dot-prefixed. Each such folder is an application keyed by its name (`slug`).
   For each slug, set the `artifacts` map by testing these paths relative to the app folder:
   | Flag | True when |
   |---|---|
   | `assessment` | `assessment/assessment.md` exists |
   | `resume_draft` | `resume/step1_draft.md` **or** `resume/step2_provenance.md` exists |
   | `resume_final` | `resume/step3_final.md` exists |
   | `cover_letter` | `cover-letter/cover_letter.md` exists |
   | `osint` | the detected company folder (below) exists under `company_intelligence` |
   | `briefing` | any file matching `interview/briefing*.md` exists |
   | `interview_prep` | any file matching `interview/interview_prep*.md` exists |
3. **Detect company / role** for each slug:
   - Drop a trailing 8-digit date token (`_YYYYMMDD`) from the slug. The remainder,
     with `_` → space, is the **humanized title**.
   - List sub-folders of `company_intelligence`. Slugify each folder name (spaces, `&`,
     `-`, `,` → `_`; collapse repeats; trim `_`) and lowercase. If a slugified folder name
     `C` satisfies `lower(slug) == C` or `lower(slug)` starts with `C + "_"`, the company is
     that folder's name and `osint` is true. Otherwise `company` is `""`.
   - `role`: if a company was detected, the humanized title with the company's leading
     tokens removed; otherwise the full humanized title.
4. **Merge** each scanned slug into the tracker:
   - **Existing slug:** keep every human-editable field verbatim — `company`, `role`,
     `stage`, `applied_date`, `next_deadline`, `contact`, `outcome`, `notes`. (Role
     humanized from a slug is lossy, e.g. `SeniorPM`, so the user's refined `company`/`role`
     are preserved once set.) Overwrite only `artifacts` and `next_action` from the scan.
   - **New slug:** append `{slug, company, role, stage: applied, applied_date: null,
     next_deadline: null, contact: null, outcome: null, notes: null, artifacts, next_action}`
     using the derived `company`/`role`.
5. **Archive** tracker entries whose slug no longer has a folder: set `stage: archived`,
   all `artifacts` false, `next_action: none`. Keep all human-status fields.
6. **Compute `next_action`** (first false wins; skip for archived → `none`):
   ```
   assessment    == false → assessjob
   resume_final  == false → buildresume
   cover_letter  == false → coverletter
   osint         == false → osint
   briefing      == false → briefing
   interview_prep== false → interviewprep
   else                   → record-outcome
   ```
7. **Sort** `applications` by `slug` ascending. Set top-level `generated` to the current
   UTC time in ISO-8601 (`YYYY-MM-DDTHH:MM:SSZ`).
8. **Write** the YAML back atomically (`tracker.yaml.tmp` then `mv`). Emit the two zones in
   this field order per entry: `slug, company, role, stage, applied_date, next_deadline,
   contact, outcome, notes, artifacts, next_action`; the `artifacts` map in this order:
   `assessment, resume_draft, resume_final, cover_letter, osint, briefing, interview_prep`.

---

## Board render

After reconcile, print a summary line then a table sorted by stage rank
(`interviewing, offer, applied, lead, on_hold, withdrawn, rejected, archived`),
then by `next_deadline` ascending (nulls last). Count of non-archived apps in the header.

```
JobOps Application Tracker · {N} active · reconciled {generated}

Stage          Company / Role             Pipeline       Next             Deadline
─────────────  ─────────────────────────  ─────────────  ───────────────  ──────────
{stage}        {company / role}           {dots}         {next_action}    {deadline}{warn}

Pipeline dots: assess▸resume▸cover▸osint▸briefing▸prep
```

- **Company / Role cell:** `"{company} / {role}"` when company is non-empty, else the
  humanized title alone.
- **Dots** — six glyphs in order `assess · resume · cover · osint · briefing · prep`:
  `●` when that flag is true; for resume use `◐` when `resume_draft && !resume_final`;
  `○` otherwise. Archived rows show `——`.
- **Deadline:** `next_deadline` as `Mon D`, or `—` if null. Append ` ⚠` when the deadline
  is within 3 days of `generated`.

---

## Interaction

If `--board-only` was passed, stop after the board.

**Claude Code (AskUserQuestion available):** loop until the user backs out.
1. Render the board.
2. **Navigate** — ask one `AskUserQuestion` whose options are the up-to-3 apps with a
   pending `next_action` (i.e. not `record-outcome`/`none`), sorted by `next_deadline`
   ascending. Label `"{company-or-title} — {role}"`; description `"{stage} · next: {next_action}"`.
   The tool's auto "Other" lets the user type a slug or table row number for any other app.
   Resolve the selection to one application.
3. **Act** — ask a second `AskUserQuestion` with ≤4 options:
   - `Run next: {next_action}` — invoke that skill (`/jobops:{next_action}`, or for
     `record-outcome` go straight to the outcome update below), passing the app folder
     `{applications_root}/{slug}` as context. On return, re-run **Reconcile** and loop.
   - `Update status` — ask follow-up `AskUserQuestion`(s) for the field to change
     (`stage`, `next_deadline`, `contact`, `outcome`, `notes`), write the new value into
     that app's human-status zone, re-write the YAML atomically, and loop.
   - `Open folder` — print `{applications_root}/{slug}` and list its contents.
   - `Back to board` — redraw the board; if chosen again at the top level, exit.

**Codex (no AskUserQuestion):** render the board, then a numbered list of the apps with a
pending `next_action` and the suggested skill for each, and stop with:

> Invoke the suggested skill manually, e.g. `/jobops:buildresume`.

---

## Notes

- This skill is the only reader/writer of `tracker.yaml`. Other skills never touch it;
  their outputs are picked up on the next reconcile.
- Never delete tracker entries — archive them, preserving human-status fields.
````

- [ ] **Step 2: Verify the contract validator accepts the new skill**

Run: `npm test`
Expected: PASS. The validator discovers `skills/dashboard/SKILL.md`, confirms `name: dashboard`, a non-empty description, and `disable-model-invocation: true`. No failures.

- [ ] **Step 3: Verify the reconcile against the fixture**

Rebuild the fixture, then — following the skill's **Reconcile** section by hand against
`test/fixtures/dashboard/workspace/` — produce `workspace/Applications/tracker.yaml`, then:

```bash
./test/fixtures/dashboard/build-fixture.sh
# (produce workspace/Applications/tracker.yaml per the skill's Reconcile section)
sed 's/^generated:.*/generated: <RECONCILE_TS>/' \
  test/fixtures/dashboard/workspace/Applications/tracker.yaml > /tmp/actual_tracker.yaml
diff -u test/fixtures/dashboard/expected_tracker.yaml /tmp/actual_tracker.yaml
```
Expected: no diff. If the diff is non-empty, fix the skill's Reconcile wording (not the
expected file) until the output matches, then re-run.

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills/dashboard/SKILL.md
git commit -m "feat(dashboard): reconcile + statusboard + menu-loop skill"
```

---

## Task 4: Documentation

**Files:**
- Modify: `docs/ARCHITECTURE.md` (§4 Output layout)
- Modify: `plugins/jobops/README.md`
- Modify: root `README.md`

- [ ] **Step 1: Note the tracker file in ARCHITECTURE.md §4**

In `docs/ARCHITECTURE.md`, in the "Flat" / output-layout area of §4, add:

```markdown
**Application tracker** — a single YAML file (`config.directories.application_tracker`,
default `{applications_root}/tracker.yaml`) maintained exclusively by `/jobops:dashboard`.
It is reconciled from the filesystem on each run: filesystem presence drives the
`artifacts` flags and `next_action`; the human-status zone (`stage`, dates, `contact`,
`outcome`, `notes`) is preserved across reconciles. No other skill reads or writes it.
```

- [ ] **Step 2: Document the skill in `plugins/jobops/README.md`**

Add `/jobops:dashboard` to the skill listing in `plugins/jobops/README.md`, following the
existing entry format used by neighboring skills:

```markdown
- **`/jobops:dashboard`** — Reconcile the application tracker from your `Applications/`
  folders and render a statusboard. On Claude Code, navigate to an application and launch
  its next pipeline step interactively; on Codex, prints a read-only board with suggested
  next steps. Pass `--board-only` to skip the interactive loop.
```

- [ ] **Step 3: Mention the dashboard in the root `README.md`**

In the root `README.md`, in the jobops feature/skill summary, add a one-line mention of
the dashboard consistent with the surrounding prose (e.g. under the resume/application
workflow bullets):

```markdown
- **Application dashboard** (`/jobops:dashboard`) — a reconciled statusboard across all
  your applications with an interactive "do the next step" loop on Claude Code.
```

- [ ] **Step 4: Verify nothing regressed**

Run: `npm test`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add docs/ARCHITECTURE.md plugins/jobops/README.md README.md
git commit -m "docs(dashboard): document /jobops:dashboard skill and tracker file"
```

---

## Task 5: Version bump and finalize deletion of the Go tool

**Files:**
- Modify: `package.json`, `.claude-plugin/marketplace.json`, `.agents/plugins/marketplace.json`, `plugins/jobops/.claude-plugin/plugin.json`, `plugins/jobops/.codex-plugin/plugin.json`, `CHANGELOG.md`
- Delete: staged `tools/dashboard/**` (already staged for deletion on this branch)

- [ ] **Step 1: Confirm the Go tool deletion is part of this change**

Run: `git status --short tools/dashboard | head`
Expected: the `tools/dashboard/**` files show as staged deletions (`D `). They are replaced
by this skill; keep them staged so the commit in Step 4 removes them.

- [ ] **Step 2: Run the version-bump skill**

Invoke the `version-bump` skill (it updates `package.json`, both `marketplace.json` files,
both `plugin.json` files, `README.md`, and `CHANGELOG.md` consistently). Choose a **minor**
bump (new feature, backward compatible): `2.6.1 → 2.7.0`.

- [ ] **Step 3: Write the CHANGELOG entry**

Ensure the new `CHANGELOG.md` section for `2.7.0` includes:

```markdown
### Added
- `/jobops:dashboard` — application statusboard reconciled from the filesystem into a
  YAML tracker, with an interactive navigate-and-act loop on Claude Code and a read-only
  board on Codex. New config key `directories.application_tracker` (self-heals on existing
  workspaces).

### Removed
- The standalone Go TUI under `tools/dashboard/` (replaced by `/jobops:dashboard`).
```

- [ ] **Step 4: Verify and commit**

Run: `npm test`
Expected: PASS.

```bash
git add -A
git commit -m "chore(release): v2.7.0 — dashboard skill replaces Go TUI"
```

- [ ] **Step 5: Final verification of the whole feature**

Run:
```bash
npm test
claude plugin validate plugins/jobops
./test/fixtures/dashboard/build-fixture.sh && echo "fixture OK"
```
Expected: `npm test` PASS; plugin validation PASS; fixture builds. Confirm
`git status` is clean (all work committed).

---

## Self-Review (completed during planning)

- **Spec coverage:** every spec section maps to a task — §4 schema → Task 3 Reconcile + Task 2 expected output; §5 reconcile → Task 3; §6 state machine → Task 3 step 6 + Task 2 next_action cases; §7 interaction → Task 3 Interaction; §8 board → Task 3 Board render; §9 setup/self-heal → Task 1 + Task 3 Configuration; §10 testing → Task 2 + Task 3 step 3; docs/version → Tasks 4–5.
- **Placeholder scan:** no TBD/TODO; all skill, JSON, YAML, and shell content is given in full. `<RECONCILE_TS>` is an intentional normalization token, documented where used.
- **Type/name consistency:** artifact keys (`assessment, resume_draft, resume_final, cover_letter, osint, briefing, interview_prep`), `next_action` values (`assessjob, buildresume, coverletter, osint, briefing, interviewprep, record-outcome, none`), and field order are identical across the skill body, the fixture, and `expected_tracker.yaml`.
