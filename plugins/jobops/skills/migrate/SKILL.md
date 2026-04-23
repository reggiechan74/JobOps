---
description: Migrate legacy JobOps output files into the v2.0 application-centric layout - dry-run preview with user-editable mapping
disable-model-invocation: true
argument-hint: "[--dry-run]"
---

# JobOps Legacy Migration

This skill moves files from the legacy v1.x output layout
(`OutputResumes/`, `Briefing_Notes/`, `Scoring_Rubrics/`,
`Intelligence_Reports/`) into the v2.0 application-centric layout. It is
a standalone runnable version of Step 9 of `/jobops:setup`.

**Flags:**
- `--dry-run` — stop after the preview without moving any files.

---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.applications_root` and
`config.directories.company_intelligence` as destinations.

---

## Step 1: Scan legacy folders

Look for files in any of:
- `OutputResumes/`
- `Briefing_Notes/`
- `Scoring_Rubrics/`
- `Intelligence_Reports/`

If all four are absent or empty, print:

> Nothing to migrate. Legacy folders are empty or missing.

…and exit with success.

---

## Step 2: Dry-run parse

For each file, attempt to parse `{Company}_{Role}_{YYYYMMDD}` (or, for
Intelligence_Reports, `{Company}_{AgentType}_{YYYYMMDD}`) using the
patterns below.

**Pattern-to-folder mapping:**
- Patterns 1–3 apply to files under `OutputResumes/`, `Scoring_Rubrics/`,
  and `Briefing_Notes/`.
- Pattern 4 applies only to files under `Intelligence_Reports/`.

Patterns (tried in order within each folder's applicable set):

1. `^([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`
2. `^(?:Step\d+_)?(?:Draft_|Provenance_Analysis_|Final_Resume_|Cover_Letter_|Briefing_|Interview_Prep_)?([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`
3. `^Rubric_([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`
4. `^([A-Za-z0-9]+)_(Corporate|Legal|Leadership|Compensation|Culture|Market)_Intelligence_(\d{8})\.md$` — applied only to files inside `Intelligence_Reports/`.

Capture groups for patterns 1–3: `(Company)_(Role)_(Date)`.
Capture groups for pattern 4: `(Company)_(AgentType)_(Date)`.

Compose destinations:

- `OutputResumes/Step1_Draft_*` → `{applications_root}/{Company}_{Role}_{Date}/resume/step1_draft.md`
- `OutputResumes/Step2_Provenance_Analysis_*` → `.../resume/step2_provenance.md`
- `OutputResumes/Step3_Final_Resume_*` → `.../resume/step3_final.md`
- `OutputResumes/Cover_Letter_*` → `.../cover-letter/cover_letter.md`
- `Scoring_Rubrics/Rubric_*` → `.../assessment/rubric.md`
- `Briefing_Notes/*` (matches pattern 1 or 2) → `.../interview/briefing.md`
- `Intelligence_Reports/*` (matches pattern 4) → `{company_intelligence}/{Company}/{agent_lowercase}.md`
  - e.g. `AcmeCorp_Corporate_Intelligence_20260423.md` → `{company_intelligence}/AcmeCorp/corporate.md`
  - e.g. `AcmeCorp_Legal_Intelligence_20260423.md` → `{company_intelligence}/AcmeCorp/legal.md`
  - e.g. `AcmeCorp_Leadership_Intelligence_20260423.md` → `{company_intelligence}/AcmeCorp/leadership.md`
  - e.g. `AcmeCorp_Compensation_Intelligence_20260423.md` → `{company_intelligence}/AcmeCorp/compensation.md`
  - e.g. `AcmeCorp_Culture_Intelligence_20260423.md` → `{company_intelligence}/AcmeCorp/culture.md`
  - e.g. `AcmeCorp_Market_Intelligence_20260423.md` → `{company_intelligence}/AcmeCorp/market.md`

Files that do not match any pattern are added to an "unresolved" list.

---

## Step 3: Preview

Print planned moves grouped by destination app folder (or company folder),
then the unresolved list. Include total counts.

If `--dry-run` was passed, stop here and exit with success.

---

## Step 4: User-editable mapping

Offer:
1. **Execute as shown** — run the moves.
2. **Edit mapping** — drop into an inline editor where the user can
   reassign specific files to different app folders, different destination
   filenames, or mark as "skip".
3. **Cancel** — exit without moving anything.

---

## Step 5: Execute

For each confirmed move:
- `mkdir -p <destination parent>` if needed.
- If the file is tracked in git (`git ls-files --error-unmatch <file>` returns 0), use `git mv <src> <dst>`.
- Otherwise use plain `mv`.
- Count failures separately and report them; do not abort the whole batch on one failure.

---

## Step 6: Update config

Read `.jobops/config.json`. Set:
- `migration.completed = true`
- `migration.completed_at = <ISO-8601 now>`
- `migration.files_moved = <existing value> + <successful moves this run>`

Write back atomically.

---

## Step 7: Summary

Print:
- Files scanned, matched, unresolved.
- Files moved successfully vs failed.
- Path to the still-unresolved files (for manual handling).
- Reminder to review and commit the moves if the workspace is under version control.

Exit.
