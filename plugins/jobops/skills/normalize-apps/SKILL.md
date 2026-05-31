---
name: normalize-apps
description: Normalize existing application folders to the canonical output naming standard - rename leading-date slugs, move flat files into sub-folders, fix drifted filenames, fold stray latex outputs, and backfill output_type front matter. Dry-run preview with user-editable mapping.
disable-model-invocation: true
argument-hint: "[--dry-run]"
---

# JobOps Application Folder Normalization

This skill brings **already-migrated** v2.0 application folders into line with the
canonical output naming standard (`{Company}_{Role}_{YYYYMMDD}` slugs; fixed sub-folder
filenames; `output_type` front matter on every Markdown output). It is distinct from
`/jobops:migrate`, which moves legacy v1.x top-level folders (`OutputResumes/`,
`Briefing_Notes/`, …) into the v2.0 layout — run that first if those exist.

Use this skill when application folders already live under `applications_root` but drifted
in shape: leading-date slug names, improvised filenames (`resume_step3_final.md`,
`Step3_Final_Resume_*.md`), flat root files that were never moved into sub-folders, a stray
`latex/` folder, or Markdown outputs missing their `output_type` key. After it runs,
`/jobops:dashboard` should show **zero false `○`** for work that actually exists.

**Flags:**
- `--dry-run` — stop after the preview without renaming or moving any files.

---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.applications_root` as the scan root and
`config.directories.company_intelligence` to validate company tokens. Use
`config.directories.<key>` for all other paths.

---

## Step 1: Scan application folders

List the immediate sub-directories of `config.directories.applications_root` whose name is
not dot-prefixed. Each is one application folder. If there are none, print:

> Nothing to normalize. No application folders found.

…and exit with success.

For each application folder, plan four classes of change (Steps 2–5). Collect every planned
action without executing it yet.

---

## Step 2: Slug / folder rename (leading-date → canonical)

For each folder name, decide the canonical slug `{Company}_{Role}_{YYYYMMDD}`:

- **Already canonical** — leading PascalCase company token, trailing 8-digit date, no
  leading date/time prefix → no rename.
- **Leading date+time prefix** `^(\d{4})-(\d{2})-(\d{2})_(\d{6})_(.+)$` (e.g.
  `2026-04-15_214414_GTAA_DirectorAirportCommercialRealEstate`) → move the date to the tail
  as a compact 8-digit date and drop the time:
  `GTAA_DirectorAirportCommercialRealEstate_20260415`.
- **Leading date only** `^(\d{4})-(\d{2})-(\d{2})_(.+)$` → same, dropping no time component.
- Validate the resulting **leading** token against `company_intelligence` sub-folders
  (slugified, case-insensitive). If it does not match any company folder, keep the rename
  but flag it in the preview as `company-token-unverified` so the user can correct it via the
  editable mapping.

A folder whose target slug collides with an existing folder is added to the **unresolved**
list (never silently merged).

---

## Step 3: Move flat root files into sub-folders + rename to canonical

Within each application folder, map root-level (and otherwise mis-placed) files to their
canonical destination by filename pattern (case-insensitive; tried in order):

| Source pattern (anywhere in the app folder) | Canonical destination |
|---|---|
| `Assessment_*.md`, `assessment_report*.md` | `assessment/assessment.md` |
| `Rubric_*.md`, `*_rubric.md` | `assessment/rubric.md` |
| `*domain_research*.md` | `assessment/domain_research.md` |
| `Step1_Draft_*.md`, `resume_step1*.md`, `*step1*draft*.md`, `*_draft.md` | `resume/step1_draft.md` |
| `Step2_Provenance*.md`, `resume_step2*.md`, `resume_provenance.md`, `*provenance*.md` | `resume/step2_provenance.md` |
| `Step3_Final_Resume_*.{md,pdf,tex,docx}`, `resume_step3_final.{md,…}`, `resume_final.{md,…}`, `*step3*final*.{md,…}`, `*_final.{md,…}` (excluding `*provenance*`) | `resume/step3_final.{md,pdf,tex,docx}` |
| `Step4_CoverLetter_*.md`, `CoverLetter_*.md`, `cover_letter_*.md`, `*cover*letter*.{md,pdf,tex,docx}` | `cover-letter/cover_letter.{md,…}` |
| `Briefing_*_Part{N}_*.md`, `briefing_part{N}*.md` | `interview/briefing_part{N}.md` |
| `Briefing_*.md` (no part marker) | `interview/briefing.md` |
| `InterviewPrep_*.md`, `interview_prep_*.md` | `interview/interview_prep.md` |

Notes:
- A PDF/TeX/DOCX derivative shares the **exact basename** of its canonical `.md` source and
  lands in the **same** sub-folder (`resume/step3_final.pdf`, `cover-letter/cover_letter.docx`).
- `job_posting.md` at the app root stays put — only its front matter is touched (Step 6).
- If two source files would map to the same canonical destination (e.g. two `*final*.md`),
  add both to the **unresolved** list rather than overwriting.

---

## Step 4: Rename drifted structured filenames

For files already inside the correct sub-folder but under a non-canonical name, rename to the
canonical basename (`resume/resume_step3_final.md` → `resume/step3_final.md`,
`resume/resume_final.md` → `resume/step3_final.md`, `cover-letter/cover.md` →
`cover-letter/cover_letter.md`, etc.). The Step 3 patterns also catch these; treat Step 4 as
the in-sub-folder pass so nothing is missed.

---

## Step 5: Fold stray `latex/` outputs

If an app folder contains a `latex/` sub-folder, move each output into the matching content
sub-folder, basename-matched: a resume PDF/TeX → `resume/`, a cover-letter PDF/TeX →
`cover-letter/`, named to match the canonical `.md` basename in that sub-folder
(`latex/resume.pdf` → `resume/step3_final.pdf`). Remove the now-empty `latex/` folder. If the
target sub-folder has no canonical `.md` to match against, add the file to the **unresolved**
list.

---

## Step 6: Backfill `output_type` front matter

For each Markdown file now at a canonical location, ensure its YAML front matter carries the
correct `output_type` (the durable, filename-independent detection key). If front matter is
absent, wrap a new block; if present but missing the key, add it. Infer from location:

| Canonical location | `output_type` |
|---|---|
| `job_posting.md` | `job_posting` |
| `assessment/rubric.md` | `rubric` |
| `assessment/assessment.md` | `assessment` |
| `resume/step1_draft.md` | `resume_step1` |
| `resume/step2_provenance.md` | `resume_provenance` |
| `resume/step3_final.md` | `resume_final` |
| `cover-letter/cover_letter.md` | `cover_letter` |
| `interview/briefing*.md` | `briefing` |
| `interview/interview_prep*.md` | `interview_prep` |

Never overwrite an existing `output_type` value — only add it where missing.

---

## Step 7: Preview

Print the full plan grouped per application folder, in this order: slug rename (Step 2),
file moves/renames (Steps 3–5), and `output_type` backfills (Step 6). Then print the
**unresolved** list (collisions, unverified company tokens, ambiguous matches). Include total
counts of each action class.

If `--dry-run` was passed, stop here and exit with success.

---

## Step 8: User-editable mapping

Offer:
1. **Execute as shown** — run every planned action.
2. **Edit mapping** — drop into an inline editor where the user can reassign a specific file
   to a different destination, correct a company token, or mark any action as **skip**.
3. **Cancel** — exit without changing anything.

---

## Step 9: Execute (reversible, fail-soft)

Process slug renames (Step 2) before in-folder moves so destination paths are stable. For
each confirmed action:
- `mkdir -p <destination parent>` if needed.
- If the path is tracked in git (`git ls-files --error-unmatch <path>` returns 0), use
  `git mv <src> <dst>`; otherwise plain `mv`.
- For Step 6 backfills, edit the file in place (front-matter insert/merge only — never touch
  the body).
- Count failures separately and report them; do not abort the whole batch on one failure.

Reversibility: because every action is a `git mv` / in-place edit on a tracked tree (or a
plain `mv` otherwise), the user can revert with `git checkout`/`git restore`. State this in
the summary.

---

## Step 10: Summary

Print:
- Folders scanned; slugs renamed; files moved/renamed; `latex/` folders folded;
  `output_type` keys backfilled.
- Actions succeeded vs failed.
- The still-unresolved list (for manual handling).
- Reminder to run `/jobops:dashboard` to confirm zero false `○`, and to review and commit the
  changes if the workspace is under version control.

Exit.
