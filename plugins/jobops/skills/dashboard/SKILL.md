---
name: dashboard
description: Render an application statusboard from a reconciled YAML tracker and (on Claude Code) drive an interactive navigate-and-act loop over the application pipeline
disable-model-invocation: true
argument-hint: "[--board-only] [--fast]"
---

# JobOps Application Dashboard

Reconciles `tracker.yaml` from the filesystem, renders a statusboard, and — on
Claude Code — drives a two-level menu loop (pick an application, then pick an
action). On Codex it renders a read-only board with suggested next actions.

**Flags:**
- `--board-only` — reconcile and render the board, then stop (no menu loop) on any platform.
- `--fast` — skip the reconcile entirely: render the board from the tracker **exactly as
  stored**, then stop (no scan, no tracker write, no menu loop). Use for a quick read-only
  glance; artifact flags and `next_action` reflect the last full reconcile, not the current
  filesystem. Takes precedence over `--board-only` (both stop after the board). If `--fast`
  and `--board-only` are passed together, `--fast` wins.

---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.

**Tracker path resolution (self-heal):** read `config.directories.application_tracker`.
If that key is absent, set it to `{config.directories.applications_root}/tracker.yaml`,
load the full existing config object, add only `directories.application_tracker` (leaving
every other key — `directories.*`, `preferences`, `candidate`, `templates`, `migration` —
untouched), and write the whole object back atomically to `.jobops/config.json.tmp` in the
same directory, then `mv` over `.jobops/config.json`, and use that path. This lets
workspaces created before the dashboard existed work with no migration step.

Under `--fast`, resolve the path **read-only**: if the key is absent, use the default
`{config.directories.applications_root}/tracker.yaml` without writing config back (a fast
glance never mutates state). If no tracker file exists at the resolved path, print
`No tracker yet — run /jobops:dashboard (without --fast) to build it.` and exit.

---

## Reconcile

Run this every invocation, before rendering — **unless `--fast` was passed**, in which case
skip this entire section and render the board directly from the stored tracker (no scan, no
recompute, no write).

1. **Load** the existing tracker YAML at the resolved path, or start with
   `{version: 1, applications: []}` if the file does not exist.
2. **Scan** `config.directories.applications_root` for immediate sub-directories whose
   name is not dot-prefixed. Each such folder is an application keyed by its name (`slug`).
   For each slug, set the `artifacts` map using the resolution order below — **first match
   wins**. The canonical path is checked first; the `output_type` YAML front-matter key is the
   **durable** fallback (trust it over the filename); a filename glob is the last-resort safety
   net for files written before producers were standardized. Scan the **whole app folder**
   (recursively), not just the expected sub-folder, so flat-layout legacy apps still register.
   When testing `output_type`, read the first front-matter block (the file's first ~30 lines).
   | Flag | True when (first match wins) |
   |---|---|
   | `assessment` | `assessment/assessment.md` exists, **OR** any `.md` in the app has `output_type: assessment` |
   | `resume_draft` | `resume/step1_draft.md` or `resume/step2_provenance.md` exists, **OR** any `.md` has `output_type ∈ {resume_step1, resume_provenance}`, **OR** glob `resume/*step1*.md`, `resume/*draft*.md`, `resume/*provenance*.md` matches |
   | `resume_final` | `resume/step3_final.md` exists, **OR** any `.md` has `output_type: resume_final`, **OR** glob `resume/*step3*final*.md` / `resume/*final*.md` matches (excluding `*provenance*`) |
   | `cover_letter` | `cover-letter/cover_letter.md` exists, **OR** any `.md` has `output_type: cover_letter`, **OR** glob `cover-letter/*cover*letter*.md` matches |
   | `osint` | the detected company folder (below) exists under `company_intelligence` |
   | `briefing` | glob `interview/briefing*.md` matches, **OR** any `.md` has `output_type: briefing` |
   | `interview_prep` | glob `interview/interview_prep*.md` matches, **OR** any `.md` has `output_type: interview_prep` |

   If `applications_root` does not exist or contains no application folders, treat the scan
   as yielding zero apps (do not error) and continue — existing tracker entries are still
   archived per step 5 and an empty board is rendered.
3. **Detect company / role** for each slug:
   - Drop a trailing 8-digit date token (`_YYYYMMDD`) from the slug. The remainder,
     with `_` → space, is the **humanized title**. A folder whose name has no trailing
     8-digit date is still a valid app; its full humanized name becomes the title.
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
   - **New slug:** append `{slug, company, role, stage: lead, applied_date: null,
     next_deadline: null, contact: null, outcome: null, notes: null, artifacts, next_action}`
     using the derived `company`/`role`. `lead` is the default stage in the absence of any
     other information — the user advances it via **Update status** as the application
     progresses.
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
8. **Write** the YAML back atomically: write `tracker.yaml.tmp` in the same directory as
   `tracker.yaml`, then `mv` it over `tracker.yaml` so the rename is atomic. Emit the two zones in
   this field order per entry: `slug, company, role, stage, applied_date, next_deadline,
   contact, outcome, notes, artifacts, next_action`; the `artifacts` map in this order:
   `assessment, resume_draft, resume_final, cover_letter, osint, briefing, interview_prep`.

---

## Board render

After reconcile (or, under `--fast`, directly from the stored tracker — the header's
`reconciled {generated}` then reflects the *last* full reconcile, signalling staleness),
print a summary line then a table sorted by stage rank
(`interviewing, offer, applied, preparing, lead, on_hold, withdrawn, rejected, archived`),
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

If `--board-only` or `--fast` was passed, stop after the board (no menu loop).

**Claude Code (AskUserQuestion available):** loop until the user backs out.
1. Render the board.
2. **Navigate** — ask one `AskUserQuestion` whose options are the up-to-3 apps with a
   pending `next_action` (i.e. not `record-outcome`/`none`), sorted by `next_deadline`
   ascending. Label `"{company-or-title} — {role}"`; description `"{stage} · next: {next_action}"`.
   The tool's auto "Other" lets the user type a slug or table row number for any other app.
   Resolve the selection to one application. Apps already at `record-outcome` are reachable
   only via the auto "Other" option (the navigate picker excludes them); selecting one jumps
   straight to the outcome fields.
3. **Act** — ask a second `AskUserQuestion` with ≤4 options:
   - `Run next: {next_action}` — invoke that skill (`/jobops:{next_action}`, or for
     `record-outcome` go straight to the outcome update below), passing the app folder
     `{applications_root}/{slug}` as context. On return, re-run **Reconcile** and loop.
   - `Update status` — ask follow-up `AskUserQuestion`(s) for the field to change
     (`stage`, `next_deadline`, `contact`, `outcome`, `notes`), write the new value into
     that app's human-status zone, re-write the YAML atomically, and loop. For `stage`, the
     choices are the pipeline stages in lifecycle order — `lead → preparing → applied →
     interviewing → offer` — plus the side states `on_hold`, `withdrawn`, and `rejected`
     (`archived` is set automatically when the app folder disappears, not chosen here).
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
