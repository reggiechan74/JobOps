---
description: Initialize JobOps workspace - configure output directories, install templates, and optionally migrate legacy files
disable-model-invocation: true
argument-hint: "[--reconfigure] [--skip-migration] [--skip-history]"
---

# JobOps Workspace Setup

This skill initializes or reconfigures a JobOps workspace. It writes a single
user-owned config file at `.jobops/config.json`, creates the output directory
tree, installs templates from the plugin into the workspace, updates
`.gitignore`, and offers optional career-history import and legacy-file migration.

**Flags:**
- `--reconfigure` — re-run setup against an existing workspace; prompts before overwriting anything.
- `--skip-migration` — skip the legacy-file migration step (Step 9) unconditionally.
- `--skip-history` — skip the optional career-history import (Step 8).

---

## Step 1: Welcome and existing-config check

If `.jobops/config.json` already exists, read it and present three choices:

1. **Reconfigure all** — walk the full interview; overwrite at the end.
2. **Reconfigure specific sections** — directories only / preferences only / templates only / migration only.
3. **Exit** — leave everything as-is.

If `--reconfigure` was passed, skip the prompt and go straight to option 1.

If no config exists, proceed to Step 2.

---

## Step 2: Directory interview

Walk the user through every path in `config.directories`. Present the default,
a one-sentence purpose, and allow inline editing. Paths are normalized
(strip trailing slash, resolve `~`, keep as relative-to-workspace form that
starts with `./`).

| Key | Default | Purpose |
|---|---|---|
| `resume_source` | `./ResumeSourceFolder` | Master HAM-Z career data (user-maintained input) |
| `job_postings` | `./Job_Postings` | Target job descriptions (user-maintained input) |
| `applications_root` | `./Applications` | Per-application output tree (one folder per application) |
| `company_intelligence` | `./Company_Intelligence` | OSINT output tree (shared across applications to same company) |
| `career_analysis` | `./Career_Analysis` | Career-level outputs (idealjob, change-one-thing, comparejobs) |
| `crisis_management` | `./Crisis_Management` | Crisis-skill outputs (severance, non-compete, etc.) |

For each path: validate the parent is writable. If not, surface the specific
error and allow the user to correct it before continuing.

---

## Step 3: Create directories

For each confirmed path, run `mkdir -p <path>`. Report each as
**created** (newly made) or **exists** (already present). Do not fail if a
path already exists.

Also `mkdir -p .jobops .jobops/templates/default .jobops/templates/custom`.

---

## Step 4: Preferences interview

Ask in order:

1. **Cultural profile** — enum `canadian` | `american`. Default `canadian`.
   Controls resume voice and spelling conventions.
2. **Default jurisdiction** — ISO 3166-2 code (e.g., `CA-ON`, `US-CA`).
   Default `CA-ON`. Used by crisis skills; they accept
   `--jurisdiction=<code>` to override per-invocation.

Do **not** ask for `default_currency` here — that is owned by `/jobops-ic:setup`
(see Step 4 of that flow).

---

## Step 5: Template installation

Copy the plugin's bundled templates into the workspace:

```bash
cp ${CLAUDE_PLUGIN_ROOT}/templates/* .jobops/templates/default/
```

Report the count of files copied. Expected count after this task runs:
4 files (`assessment_rubric_framework.md`, `evidence_verification_framework.md`,
`assessment_report_structure.md`, `candidate_profile_schema.json`).

If the glob fails in the current execution context (no files matched), fall
back to:
```bash
find ${CLAUDE_PLUGIN_ROOT}/templates -maxdepth 1 -type f -exec cp {} .jobops/templates/default/ \;
```

Do not touch `.jobops/templates/custom/` — it is user-owned.

---

## Step 6: Write `.jobops/config.json`

Emit the full schema below with the values gathered in Steps 2 and 4.
`migration.completed` starts as `false` so Step 9 runs on first setup.

```json
{
  "version": "2.0",
  "directories": {
    "resume_source": "<step-2 value>",
    "job_postings": "<step-2 value>",
    "applications_root": "<step-2 value>",
    "company_intelligence": "<step-2 value>",
    "career_analysis": "<step-2 value>",
    "crisis_management": "<step-2 value>"
  },
  "preferences": {
    "cultural_profile": "<step-4 value>",
    "default_jurisdiction": "<step-4 value>"
  },
  "templates": {
    "base_dir": "./.jobops/templates",
    "active": {
      "assessment_rubric_framework": "default",
      "evidence_verification_framework": "default",
      "assessment_report_structure": "default",
      "candidate_profile_schema": "default"
    }
  },
  "migration": {
    "completed": false,
    "completed_at": null,
    "files_moved": 0
  }
}
```

Write atomically: write to `.jobops/config.json.tmp`, then `mv` over the final
name.

---

## Step 7: Gitignore update

Show the user the JobOps block that will be appended to `.gitignore`:

```
# JobOps workspace
.jobops/
Applications/
Company_Intelligence/
Career_Analysis/
Crisis_Management/
```

(If the user customized any of the directory names in Step 2, substitute those
paths in the block.)

Offer three choices:
1. **Append as shown** — add the block to the end of `.gitignore`.
2. **Edit inline** — let the user modify the block before appending.
3. **Skip** — do not modify `.gitignore`.

If `.gitignore` does not exist, create it. If the block (detected by the
`# JobOps workspace` marker line) already exists, replace it in place rather
than appending a second copy.

---

## Step 8: Optional career-history import

Only offer this step if `--skip-history` was not passed and the
`resume_source` directory is empty or contains fewer than 3 files.

Prompt: "Do you have an existing resume to import as the basis for your HAM-Z
career inventory? (yes/no)"

If yes, accept a path to a `.pdf`, `.docx`, `.txt`, or `.md` file. Extract
text, then populate:
- `{resume_source}/Experience/<Company>_<Role>.md` — one file per role
- `{resume_source}/CareerHighlights/highlights.md`
- `{resume_source}/Technology/stack.md`
- `{resume_source}/Preferences/preferences.md`
- `{resume_source}/Vision.md` — stub, to be filled in manually

After population, flag any entries that lack metrics or hard skills so the
user knows where to focus subsequent manual enrichment.

---

## Step 9: Optional legacy migration

Skip this step if `--skip-migration` was passed OR `migration.completed` is
already `true` in the config being reconfigured.

Scan for legacy folders at the workspace root and subfolders named:
- `OutputResumes/`
- `Briefing_Notes/`
- `Scoring_Rubrics/`
- `Intelligence_Reports/`

If all four are absent or empty, skip this step silently (and still set
`migration.completed = true`).

Otherwise run:

**9a. Dry-run parse.** For each file in the legacy folders, attempt to parse
`{Company}_{Role}_{YYYYMMDD}` from the filename using the following patterns
(tried in order):
1. `^([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`
2. `^(?:Step\d+_)?(?:Draft_|Provenance_Analysis_|Final_Resume_)?([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`
3. `^Rubric_([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`

For each match, compose a planned destination:
- `OutputResumes/Step1_*` and `Step2_*` and `Step3_*` → `{applications_root}/{Company}_{Role}_{Date}/resume/<stepN>.md`
- `OutputResumes/Cover_Letter_*` → `{applications_root}/{Company}_{Role}_{Date}/cover-letter/cover_letter.md`
- `Scoring_Rubrics/Rubric_*` → `{applications_root}/{Company}_{Role}_{Date}/assessment/rubric.md`
- `Briefing_Notes/*` → `{applications_root}/{Company}_{Role}_{Date}/interview/briefing.md`
- `Intelligence_Reports/<Company>_*` → `{company_intelligence}/{Company}/<agent>.md` (parse the agent name from the filename segment between Company and Date)

**9b. Preview.** Print the planned moves grouped by target app folder. List
unresolved files (no regex match) in a separate "manual" group.

**9c. User-editable mapping.** Offer three options:
1. **Execute as shown** — proceed with the preview.
2. **Edit mapping** — drop into an inline editor where the user can reassign
   specific files to different application folders or mark them as "skip".
3. **Cancel** — leave everything in place.

**9d. Execute.** For each confirmed move:
- If the file is tracked in git, use `git mv <src> <dst>` (preserves history).
- Otherwise use plain `mv`.
- Create any missing destination parent folders first.

**9e. Update config.** Set:
- `migration.completed = true`
- `migration.completed_at = <ISO-8601 now>`
- `migration.files_moved = <count of successful moves>`

Unresolved files stay in place; the user handles them manually.

---

## Step 10: Summary and next steps

Print:

1. What was configured (each path + confirm/created/exists).
2. What was installed (template count).
3. Whether career history was imported (and how many files).
4. Whether migration ran (and how many files moved).
5. Recommended next steps:
   - Drop a job posting into `{job_postings}/` named `{Company}_{Role}_{YYYYMMDD}.md`.
   - Run `/jobops:buildresume <job_posting.md>` to produce a tailored resume.
   - Run `/jobops:osint <Company>` to gather company intelligence.
   - If you are an independent contractor, run `/jobops-ic:setup` to enable IC features.

Exit.
