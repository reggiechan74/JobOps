# JobOps Plugin Config & Output-Layout Redesign — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use `superpowers:subagent-driven-development` (recommended) or `superpowers:executing-plans` to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Rework the `feature/plugin-split` branch so the JobOps plugin writes all artifacts to user-configurable, application-centric output folders driven by a dedicated workspace config file at `.jobops/config.json`, and eliminate obsolete plumbing (`/tmp/.jobops-plugin-root` hook, `copy-templates.sh`) now that `${CLAUDE_PLUGIN_ROOT}` expands inline in skill markdown.

**Architecture:** `plugin.json` stays lean (metadata + dependencies only). All user state lives in `.jobops/config.json`, populated by a rewritten `/jobops:setup` via interactive interview. Every skill reads the config, applies a fixed path-resolution protocol to pick an output folder, and writes there. A new `/jobops:migrate` skill relocates legacy flat-layout artifacts into the new application-centric layout.

**Tech Stack:** Claude Code plugin system, markdown skills with `disable-model-invocation: true`, JSON config, shell commands invoked from skill markdown, `${CLAUDE_PLUGIN_ROOT}` env var expansion.

**Spec:** `docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md`

**Prerequisite branch state:** `feature/plugin-split` (commit `443ee79` or later — the spec is committed on this branch). Assumes the March 2026 plugin scaffolding is in place.

---

## File Structure Map

### Files deleted

- `plugins/jobops/hooks/hooks.json`
- `plugins/jobops/scripts/copy-templates.sh`
- `plugins/jobops/hooks/` (empty directory after deletion)
- `plugins/jobops/scripts/` (empty directory after deletion)
- `plugins/jobops-ic/hooks/hooks.json`
- `plugins/jobops-ic/scripts/copy-templates.sh`
- `plugins/jobops-ic/hooks/` (empty directory)
- `plugins/jobops-ic/scripts/` (empty directory)

### Files created

- `plugins/jobops/skills/migrate/SKILL.md` — standalone migration skill
- `docs/ARCHITECTURE.md` — contributor-only config & path-resolution contract

### Files modified (full rewrites)

- `plugins/jobops/skills/setup/SKILL.md` — new 10-step flow
- `plugins/jobops-ic/skills/setup/SKILL.md` — light extend-config flow

### Files modified (targeted edits — preamble + path references)

**jobops, 29 skills (all except `setup`):**
- Resume-development: `buildresume`, `provenance-check`, `coverletter`
- Assessment: `assessjob`, `assesscandidate`, `createrubric`, `comparejobs`
- Interview-prep: `briefing`, `interviewprep`
- Job-search: `osint`, `auditjobposting`
- Career-strategy: `idealjob`, `change-one-thing`, `assess-job-offer`
- Career-crisis (11): `code-red`, `severance-review`, `workplace-documentation`, `non-compete-analysis`, `reference-shield`, `unemployment-prep`, `discrimination-assessment`, `investigation-prep`, `accommodation-request`, `layoff-intel`, `constructive-dismissal`
- Application-finalization: `convert-to-pdf`, `convert-to-word`, `markdown-to-pdf`
- System-setup: `install-pandoc`, `github-portfolio`

**jobops-ic, 9 skills (all except `setup`):**
`defineservices`, `findclient`, `pitchdeck`, `proposaltemplate`, `ratecard`, `copywrite`, `copywriting-spec`, `create-landing-page`, `css-template`

### Files modified (small)

- `plugins/jobops-ic/.claude-plugin/plugin.json` — add `dependencies` field
- `README.md` — refresh setup section (post-implementation)
- `CHANGELOG.md` — v2.0.0 entry covering the redesign

### Files unchanged

- All 17 jobops agents (`plugins/jobops/agents/*.md`)
- 1 jobops-ic agent (`plugins/jobops-ic/agents/landing-page-copywriter.md`)
- `plugins/jobops/templates/*.md` (bundled templates)
- `plugins/jobops-ic/templates/*` (bundled service_definition_schema.json)
- `.claude-plugin/marketplace.json`
- `plugins/jobops/.claude-plugin/plugin.json`
- `package.json` (already at 2.0.0)
- `CLAUDE.md` (already rewritten as contributor-only)

---

## Standard Skill Preamble (reference — used throughout Phase 4)

### Standard jobops preamble

```markdown
## Configuration

Read `.jobops/config.json`. If missing, stop with:

  JOBOPS NOT CONFIGURED
  Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if the skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if the skill has jurisdiction-sensitive logic.
```

### Standard jobops-ic preamble

```markdown
## Configuration

Read `.jobops/config.json`. If missing, stop with:

  JOBOPS NOT CONFIGURED
  Run /jobops:setup (then /jobops-ic:setup) to initialize your workspace.

Use `config.directories.contractor_root` for output paths in this skill.
Use `config.preferences.default_currency` for pricing if applicable.
```

### Templates block (for template-consuming skills only)

```markdown
## Templates

For each template needed, resolve path from:
  {config.templates.base_dir}/{config.templates.active.<name>}/<filename>

Templates referenced by this skill: <explicit list>
```

### Application-path resolution protocol (for application-writing skills only)

```markdown
## Application-Path Resolution

1. Derive the application slug from the job-posting filename: `{Company}_{Role}_{YYYYMMDD}`.
   If `--app=<slug>` argument is provided, use that instead.
2. Compose the app folder: `{config.directories.applications_root}/{app_slug}/`
3. This skill writes to sub-folder `<category>/` (one of: resume, cover-letter, assessment, interview).
4. If the app folder does not exist: `mkdir -p` the target path, then copy
   `{config.directories.job_postings}/{filename}.md` to `{app_folder}/job_posting.md`
   (provenance pin — only if `job_posting.md` does not already exist).
```

### Company-intelligence resolution (for OSINT skills only)

```markdown
## Company-Intelligence Resolution

1. Extract `{Company}` from the job-posting frontmatter or explicit argument.
2. Target: `{config.directories.company_intelligence}/{Company}/`
3. If the target folder exists, prompt the user: **refresh** (overwrite all files),
   **append** (write a new `summary_{YYYYMMDD}.md` alongside existing files), or **skip**.
4. Person-level reports from `osint-person` agent write to `{Company}/people/{interviewer_name}.md`.
```

---

## Phase 1: Foundation cleanup

### Task 1: Delete obsolete hooks and scripts

**Files:**
- Delete: `plugins/jobops/hooks/hooks.json`
- Delete: `plugins/jobops/scripts/copy-templates.sh`
- Delete: `plugins/jobops/hooks/` (empty directory)
- Delete: `plugins/jobops/scripts/` (empty directory)
- Delete: `plugins/jobops-ic/hooks/hooks.json`
- Delete: `plugins/jobops-ic/scripts/copy-templates.sh`
- Delete: `plugins/jobops-ic/hooks/` (empty directory)
- Delete: `plugins/jobops-ic/scripts/` (empty directory)

- [ ] **Step 1: Remove hook and script files**

```bash
rm plugins/jobops/hooks/hooks.json
rm plugins/jobops/scripts/copy-templates.sh
rm plugins/jobops-ic/hooks/hooks.json
rm plugins/jobops-ic/scripts/copy-templates.sh
```

- [ ] **Step 2: Remove now-empty directories**

```bash
rmdir plugins/jobops/hooks plugins/jobops/scripts
rmdir plugins/jobops-ic/hooks plugins/jobops-ic/scripts
```

- [ ] **Step 3: Verify plugin directory structure**

Run: `find plugins -maxdepth 2 -type d | sort`

Expected output:
```
plugins
plugins/jobops
plugins/jobops/.claude-plugin
plugins/jobops/agents
plugins/jobops/skills
plugins/jobops/templates
plugins/jobops-ic
plugins/jobops-ic/.claude-plugin
plugins/jobops-ic/agents
plugins/jobops-ic/skills
plugins/jobops-ic/templates
```

No `hooks/` or `scripts/` directories should appear.

- [ ] **Step 4: Commit**

```bash
git add -A
git commit -m "$(cat <<'EOF'
Remove obsolete hooks and copy-templates scripts

${CLAUDE_PLUGIN_ROOT} now expands directly in skill markdown, removing
the need for the SessionStart hook that wrote the plugin root to /tmp/
and the copy-templates.sh helper that read from it.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 2: Add `dependencies` field to `jobops-ic/plugin.json`

**Files:**
- Modify: `plugins/jobops-ic/.claude-plugin/plugin.json`

- [ ] **Step 1: Read current jobops-ic plugin.json**

Run: `cat plugins/jobops-ic/.claude-plugin/plugin.json`

Note the current structure. It should be pure metadata; no `dependencies` field yet.

- [ ] **Step 2: Rewrite the file with dependencies added**

Replace the entire contents of `plugins/jobops-ic/.claude-plugin/plugin.json` with:

```json
{
  "name": "jobops-ic",
  "version": "2.0.0",
  "description": "Independent contractor toolkit - services, client prospecting, pitch decks, proposals, rate cards, landing pages. Requires jobops.",
  "author": {
    "name": "Reggie Chan"
  },
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": ["independent-contractor", "consulting", "proposals", "rate-card", "pitch-deck", "landing-page"],
  "dependencies": {
    "jobops": "^2.0.0"
  }
}
```

- [ ] **Step 3: Verify JSON is valid**

Run: `python3 -c "import json; json.load(open('plugins/jobops-ic/.claude-plugin/plugin.json'))" && echo OK`

Expected: `OK`

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops-ic/.claude-plugin/plugin.json
git commit -m "$(cat <<'EOF'
Declare jobops-ic dependency on jobops via plugin.json

Replaces the handwritten 'verify resume-summarizer agent exists' prerequisite
check that every jobops-ic skill carried. Claude Code now enforces the
dependency at install time.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 3: Baseline plugin validation

**Files:** None (verification only)

- [ ] **Step 1: Validate jobops plugin structure**

Run: `claude plugin validate plugins/jobops`

Expected: No errors. If validation reports errors, stop and investigate — the scaffolding is broken and downstream tasks assume a valid baseline.

- [ ] **Step 2: Validate jobops-ic plugin structure**

Run: `claude plugin validate plugins/jobops-ic`

Expected: No errors.

- [ ] **Step 3: Validate marketplace**

Run: `claude plugin validate .`

Expected: Marketplace validated successfully. Both plugins listed.

- [ ] **Step 4: No commit (verification-only task)**

---

## Phase 2: Setup skills

### Task 4: Rewrite `plugins/jobops/skills/setup/SKILL.md` — Steps 1–7 (config-writing core)

**Files:**
- Modify: `plugins/jobops/skills/setup/SKILL.md` (full rewrite in two tasks; this one covers Steps 1–7)

- [ ] **Step 1: Delete the existing setup skill content**

Empty the file, retaining only the opening frontmatter block. Replace the current content entirely.

- [ ] **Step 2: Write the new setup skill, Steps 1–7**

Replace the entire contents of `plugins/jobops/skills/setup/SKILL.md` with:

```markdown
---
description: Initialize JobOps workspace - interview for directory paths, preferences, and templates; optionally import career history and migrate legacy outputs
disable-model-invocation: true
---

# JobOps Workspace Setup

**Arguments:**
- `--reconfigure` — skip welcome; go directly to editing existing config
- `--skip-migration` — skip Step 9 (legacy file migration)
- `--skip-history` — skip Step 8 (career history import)

---

## Step 1: Welcome & existing-config check

Check whether `.jobops/config.json` already exists:

```bash
if [ -f ".jobops/config.json" ]; then
  echo "EXISTING_CONFIG_FOUND"
  cat .jobops/config.json
else
  echo "NO_CONFIG"
fi
```

**If EXISTING_CONFIG_FOUND:** Display the current config and offer the user three choices:

1. **Reconfigure all sections** — walk through Steps 2–10 with current values shown as defaults
2. **Reconfigure specific sections** — user picks which sections to re-run (directories / preferences / templates / migration)
3. **Exit** — leave config unchanged

If `--reconfigure` was passed as an argument, skip the prompt and default to choice 1.

**If NO_CONFIG:** Display a welcome message and proceed to Step 2. Welcome message:

> Welcome to JobOps setup. This skill initializes your workspace with everything needed to run the JobOps intelligence-driven job application system. Setup walks you through directory configuration, preferences, template installation, optional career-history import, and optional migration of legacy outputs. Re-runnable at any time.

---

## Step 2: Directory interview

For each of the seven `config.directories` keys below, present the key, its purpose, and a default value. Ask the user to accept the default or supply a custom path. Normalize relative paths to be workspace-relative (e.g., `./Applications`). Validate that the parent directory is writable.

Present each in this format:

```
[1/7] resume_source
  Purpose: Master career inventory (HAM-Z Experience/, CareerHighlights/, Technology/, Preferences/)
  Default: ./ResumeSourceFolder
  Your path [Enter to accept default]:
```

The seven directories:

| Key | Default | Purpose |
|---|---|---|
| `resume_source` | `./ResumeSourceFolder` | Master career inventory |
| `job_postings` | `./Job_Postings` | Target job descriptions as `.md` files |
| `applications_root` | `./Applications` | Per-application output folders |
| `company_intelligence` | `./Company_Intelligence` | OSINT output at company level (shared across applications) |
| `career_analysis` | `./Career_Analysis` | `idealjob`, `change-one-thing`, `comparejobs` output |
| `crisis_management` | `./Crisis_Management` | Crisis-skill output (severance, non-compete, discrimination, etc.) |
| `contractor_root` | `./Contractor` | **Only asked if `jobops-ic` plugin is installed** |

To detect whether `jobops-ic` is installed, check `/plugin list` output for `jobops-ic` or check `${CLAUDE_PLUGIN_ROOT}/../jobops-ic/.claude-plugin/plugin.json` existence. If jobops-ic is not installed, skip the `contractor_root` prompt; it will be added later by `/jobops-ic:setup`.

After collecting all responses, show them back in one summary and ask for final confirmation before proceeding.

---

## Step 3: Create directories

For each confirmed directory path from Step 2, ensure it exists:

```bash
mkdir -p "<resume_source_path>"
mkdir -p "<job_postings_path>"
mkdir -p "<applications_root_path>"
mkdir -p "<company_intelligence_path>"
mkdir -p "<career_analysis_path>"
mkdir -p "<crisis_management_path>"
# Only if jobops-ic installed:
mkdir -p "<contractor_root_path>"
```

Also create the JobOps internal directory:

```bash
mkdir -p ".jobops"
```

Report which directories were newly created versus which already existed.

---

## Step 4: Preferences interview

Ask the user for each preference in turn:

**Cultural profile** (affects resume voice):
```
1. canadian (default) — collaborative, measured language; emphasizes team contribution
2. american — assertive, individual achievement focus; strong action verbs
```

**Default legal jurisdiction** (used by crisis skills; override per-invocation with `--jurisdiction=X`):
```
Enter an ISO 3166-2 code.
Examples: CA-ON (Ontario, default), CA-BC, US-CA, US-NY, GB-ENG
Default: CA-ON
```

`default_currency` is NOT asked here. It is owned by `/jobops-ic:setup` and added when that skill runs.

Store the responses for use in Step 6.

---

## Step 5: Template installation

Copy the plugin's bundled templates to the workspace, using `${CLAUDE_PLUGIN_ROOT}` to resolve the plugin path:

```bash
TEMPLATES_BASE="./.jobops/templates"
mkdir -p "${TEMPLATES_BASE}/default"
mkdir -p "${TEMPLATES_BASE}/custom"
cp "${CLAUDE_PLUGIN_ROOT}/templates/"*.md "${TEMPLATES_BASE}/default/"
cp "${CLAUDE_PLUGIN_ROOT}/templates/"*.json "${TEMPLATES_BASE}/default/"
```

After copying, list the installed templates:

```bash
ls -1 "${TEMPLATES_BASE}/default/"
```

Expected: four files —
- `assessment_rubric_framework.md`
- `evidence_verification_framework.md`
- `assessment_report_structure.md`
- `candidate_profile_schema.json`

If the copy fails (source directory missing, permissions, etc.), report the error and advise the user to check that the `jobops` plugin was installed correctly.

---

## Step 6: Write `.jobops/config.json`

Write the workspace configuration file at `.jobops/config.json`. Use the confirmed paths from Step 2 and preferences from Step 4.

```json
{
  "version": "2.0",
  "directories": {
    "resume_source": "<step2_value>",
    "job_postings": "<step2_value>",
    "applications_root": "<step2_value>",
    "company_intelligence": "<step2_value>",
    "career_analysis": "<step2_value>",
    "crisis_management": "<step2_value>"
  },
  "preferences": {
    "cultural_profile": "<step4_value>",
    "default_jurisdiction": "<step4_value>"
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

Do NOT include `directories.contractor_root`, `preferences.default_currency`, or `templates.active.service_definition_schema` in this file — those are added only when `/jobops-ic:setup` runs. If the user chose `--reconfigure` and the existing file already had those keys (meaning jobops-ic setup ran previously), preserve them verbatim.

---

## Step 7: Gitignore update

Propose appending a JobOps block to `.gitignore`. First read any existing `.gitignore`:

```bash
if [ -f .gitignore ]; then
  cat .gitignore
else
  echo "NO_GITIGNORE"
fi
```

Show the user the proposed block to append:

```
# JobOps workspace
.jobops/
<applications_root>
<company_intelligence>
<career_analysis>
<crisis_management>
```

(Substitute the user's actual configured paths.)

Ask: **Append this block? (Y / N / edit)**

- **Y:** Append (or create) `.gitignore` with the block. Skip any line that's already present in the file.
- **N:** Skip — user handles `.gitignore` manually.
- **edit:** Open an inline edit session; user supplies the final block text before writing.

---

## Step 8: Career history import (optional)

See Steps 8a–8h in the continuation section below.

## Step 9: Migration of legacy files (optional)

See continuation section below.

## Step 10: Summary

See continuation section below.
```

- [ ] **Step 3: Verify markdown syntax is valid and the frontmatter parses**

Run:
```bash
head -4 plugins/jobops/skills/setup/SKILL.md
```

Expected: frontmatter opens with `---`, has `description:` and `disable-model-invocation: true`, closes with `---`.

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills/setup/SKILL.md
git commit -m "$(cat <<'EOF'
Rewrite jobops:setup Steps 1-7 (config-writing core)

New flow reads ${CLAUDE_PLUGIN_ROOT} inline for template installation,
interviews the user for 7 directory paths and 2 preferences, writes a
validated .jobops/config.json, and proposes a .gitignore update. Step 8-10
added in follow-up task.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 5: Extend `plugins/jobops/skills/setup/SKILL.md` — Steps 8–10 (history, migration, summary)

**Files:**
- Modify: `plugins/jobops/skills/setup/SKILL.md`

- [ ] **Step 1: Replace Steps 8–10 stubs with the full flow**

Find the three stub lines at the bottom of `plugins/jobops/skills/setup/SKILL.md`:

```markdown
## Step 8: Career history import (optional)

See Steps 8a–8h in the continuation section below.

## Step 9: Migration of legacy files (optional)

See continuation section below.

## Step 10: Summary

See continuation section below.
```

Replace them with this full content:

```markdown
## Step 8: Career history import (optional)

Skip this section entirely if `--skip-history` was passed.

Otherwise, ask: **Do you have existing resume files to import?**

If the user declines, proceed to Step 9.

If they say yes, walk through Steps 8a–8h:

### 8a. Collect input files

Accept one or more resume files in these formats: PDF (`.pdf`), Word (`.docx`), plain text (`.txt`), markdown (`.md`). Ask the user to supply absolute paths.

### 8b. Parse and extract

Read each file and extract:
- Work experience — company names, titles, dates, responsibilities, accomplishments
- Education — degrees, institutions, dates, certifications
- Skills — technical skills, tools, platforms, methodologies
- Achievements — quantified results, awards, recognitions, project outcomes

### 8c. Transform to HAM-Z format

Convert extracted achievements into HAM-Z format:

> Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]

Each bullet structured as:
- **H** — Hard Skill
- **A** — Action
- **M** — Metrics
- **Z** — Structure

Flag items lacking metrics or hard skills for user review.

### 8d. Create ResumeSourceFolder structure

Using the user's confirmed `resume_source` path from Step 2:

```bash
RESUME_SOURCE="<user_confirmed_resume_source_path>"
mkdir -p "${RESUME_SOURCE}/Experience"
mkdir -p "${RESUME_SOURCE}/CareerHighlights"
mkdir -p "${RESUME_SOURCE}/Technology"
mkdir -p "${RESUME_SOURCE}/Preferences"
```

### 8e. Populate files

Write extracted data into the appropriate subdirectories. One file per role in `Experience/`. Top achievements in `CareerHighlights/`. Skills catalog in `Technology/` organized by category (languages, platforms, tools, methodologies). Preferences in `Preferences/`.

### 8f. Create Vision template

Create `${RESUME_SOURCE}/Preferences/Vision.md` if it does not already exist:

```markdown
# Career Vision

## Target Roles
<!-- List the roles you are pursuing -->

## Industry Preferences
<!-- Industries or sectors of interest -->

## Work Style
<!-- Remote, hybrid, on-site preferences -->

## Geographic Preferences
<!-- Locations, willingness to relocate -->

## Compensation Expectations
<!-- Salary range, equity preferences, benefits priorities -->

## Non-Negotiables
<!-- Hard requirements for your next role -->

## Growth Goals
<!-- Skills to develop, career trajectory -->
```

### 8g. Gap analysis

After import, report:
- Count of experience entries, highlights, skills imported
- Bullets that could not be converted to HAM-Z (missing metrics or hard skills)
- Recommendations for strengthening weak entries

### 8h. Profile generation suggestion

If the import produced substantial data, suggest:

> To generate a compressed candidate profile for faster assessments, run the `resume-summarizer` agent. It produces an optimized JSON profile with 85–90% token reduction while preserving all key career data.

---

## Step 9: Migration of legacy files (optional)

Skip this section entirely if `--skip-migration` was passed OR if `migration.completed == true` in the existing config.

Legacy folder names to check:

| Legacy folder | New destination |
|---|---|
| `OutputResumes/` | `{applications_root}/{Company}_{Role}_{YYYYMMDD}/resume/` |
| `Briefing_Notes/` | `{applications_root}/{Company}_{Role}_{YYYYMMDD}/interview/` |
| `Scoring_Rubrics/` | `{applications_root}/{Company}_{Role}_{YYYYMMDD}/assessment/` |
| `Intelligence_Reports/` | `{company_intelligence}/{Company}/` |

### 9a. Scan

```bash
for dir in OutputResumes Briefing_Notes Scoring_Rubrics Intelligence_Reports; do
  if [ -d "$dir" ] && [ "$(ls -A "$dir" 2>/dev/null)" ]; then
    count=$(find "$dir" -type f | wc -l)
    echo "$dir: $count files"
  fi
done
```

If no legacy folders contain files, skip to Step 10.

Otherwise, offer: **Migrate legacy files into the new layout? (Y / N / custom)**

- **Y:** proceed to 9b
- **N:** skip to Step 10; set `migration.completed = true` so this is not re-asked
- **custom:** let the user pick specific legacy folders to migrate; others skipped

### 9b. Dry-run pass — parse filenames

Walk each legacy folder's files. For each filename, attempt to extract `{Company}`, `{Role}`, and `{YYYYMMDD}` using these ordered patterns (first match wins):

1. `Rubric_(?P<company>[^_]+)_(?P<role>[^_]+)_(?P<date>\d{8})\.md` — Scoring_Rubrics style
2. `Rubric_(?P<company>[^_]+)_(?P<role>[^_]+)_(?P<date>\d{4}-\d{2}-\d{2})\.md` — dash-date variant
3. `Briefing_(?P<company>[^_]+)_(?P<role>[^_]+)(?:_\w+)*_(?P<date>\d{8})\.md` — Briefing_Notes style
4. `InterviewPrep_(?P<company>[^_]+)_(?P<role>[^_]+)_(?P<date>\d{8})\.md` — interview prep variant
5. `(?P<company>[^_]+)_(?P<type>\w+)_Intelligence_(?P<date>\d{4}-\d{2}-\d{2})\.md` — Intelligence_Reports style (company-level, no role)
6. `(?P<date>\d{4}-\d{2}-\d{2})_\d+_(?P<company>[^_]+)_(?P<role>[^/]+)` — OutputResumes subfolder style (date first)
7. `Step\d+_(?:Draft|Provenance|Final)(?:_Resume)?_(?P<company>[^_]+)_(?P<role>[^_]+)_(?P<date>\d{8})\.md` — OutputResumes flat-file style

Normalize date formats to `YYYYMMDD` (drop dashes). Build a mapping:

```
LEGACY_FILE                                              → TARGET_PATH
OutputResumes/Step3_Final_Google_SeniorPM_20260423.md    → Applications/Google_SeniorPM_20260423/resume/step3_final.md
Briefing_Notes/Briefing_Google_SeniorPM_20260423.md      → Applications/Google_SeniorPM_20260423/interview/briefing.md
Scoring_Rubrics/Rubric_Google_SeniorPM_20260423.md       → Applications/Google_SeniorPM_20260423/assessment/rubric.md
Intelligence_Reports/Google_Corporate_Intelligence_2026-01-15.md → Company_Intelligence/Google/corporate.md
```

Files that do not match any pattern are listed as **unresolved** and remain in their original locations.

### 9c. Preview

Display the mapping in a structured format, grouped by target app folder for Applications-bound files and by company for intelligence files. List unresolved files separately.

Ask the user:
- **accept** — execute all planned moves
- **reassign** — the user specifies which files should map to a different target (inline edit per file)
- **cancel** — abort migration, no files moved

### 9d. Execute

For each accepted move:

```bash
TARGET_DIR=$(dirname "$TARGET_PATH")
mkdir -p "$TARGET_DIR"

if git ls-files --error-unmatch "$LEGACY_FILE" > /dev/null 2>&1; then
  git mv "$LEGACY_FILE" "$TARGET_PATH"
else
  mv "$LEGACY_FILE" "$TARGET_PATH"
fi
```

Track counts: `MOVED_COUNT`, `UNRESOLVED_COUNT`.

### 9e. Update config

Update `.jobops/config.json` to reflect migration status:

```json
{
  "migration": {
    "completed": true,
    "completed_at": "<ISO 8601 timestamp>",
    "files_moved": <MOVED_COUNT>
  }
}
```

Preserve all other config fields exactly as they were.

### 9f. Report

Summarize:
- Total files moved
- Legacy folders now empty (OutputResumes, Briefing_Notes, Scoring_Rubrics, Intelligence_Reports — list which are safe to delete)
- Unresolved files still in original locations

---

## Step 10: Summary

Display a final summary:

### Workspace configured

- `.jobops/config.json` written
- Directories created/confirmed: list each path with `(created)` or `(existed)`
- Templates installed: list files in `.jobops/templates/default/`
- Custom templates directory: `.jobops/templates/custom/` (empty, for user variants)
- Gitignore: `(updated)` / `(skipped by user)`

### Preferences

- Cultural profile: `<canadian | american>`
- Default jurisdiction: `<ISO 3166-2 code>`

### Career history

- `(imported N experiences, M highlights, K skills)` / `(skipped)`
- Items needing review: list any flagged HAM-Z gaps

### Migration

- `(moved N files, P unresolved)` / `(skipped)` / `(already completed)`

### Next steps

1. Drop a job posting into `{job_postings}/` — filename convention `Company_Role_YYYYMMDD.md`
2. Try `/jobops:buildresume <path-to-jd>` for a tailored resume
3. Try `/jobops:assessjob <path-to-jd>` for a 200-point assessment
4. If you use the contractor toolkit, install `jobops-ic` and run `/jobops-ic:setup`
```

- [ ] **Step 2: Verify file is well-formed**

Run:
```bash
wc -l plugins/jobops/skills/setup/SKILL.md
grep -c "^## Step" plugins/jobops/skills/setup/SKILL.md
```

Expected: ~350–450 lines. Ten `## Step` headings.

- [ ] **Step 3: Commit**

```bash
git add plugins/jobops/skills/setup/SKILL.md
git commit -m "$(cat <<'EOF'
Complete jobops:setup with history import and legacy migration

Adds Step 8 (career-history import with HAM-Z transformation),
Step 9 (legacy file migration with dry-run preview and user-editable
mapping), and Step 10 (final summary).

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 6: Rewrite `plugins/jobops-ic/skills/setup/SKILL.md`

**Files:**
- Modify: `plugins/jobops-ic/skills/setup/SKILL.md`

- [ ] **Step 1: Replace file contents**

Replace the entire contents of `plugins/jobops-ic/skills/setup/SKILL.md` with:

```markdown
---
description: Extend JobOps config for the independent contractor toolkit - contractor root directory, default currency, and service-definition template
disable-model-invocation: true
---

# JobOps-IC Workspace Setup

Lightweight extension of the base JobOps config. Adds a Contractor output root and currency preference, plus the `service_definition_schema` template.

## Step 1: Prerequisite check

Check for the base config:

```bash
if [ ! -f ".jobops/config.json" ]; then
  echo "JOBOPS NOT CONFIGURED"
  echo "Run /jobops:setup first to initialize the base JobOps workspace,"
  echo "then re-run /jobops-ic:setup."
  exit 1
fi
```

If the file is missing, stop with the message above. Do not proceed.

## Step 2: Interview

**Contractor output root:**
```
contractor_root
  Purpose: Output for jobops-ic artifacts (services, prospects, proposals, pitches, rate cards, landing pages)
  Default: ./Contractor
  Your path [Enter to accept default]:
```

**Default currency** (for rate cards and proposals):
```
Enter an ISO 4217 code.
Options: CAD (default), USD, EUR, GBP, AUD
Default: CAD
```

## Step 3: Create directory

```bash
mkdir -p "<contractor_root_path>"
mkdir -p "<contractor_root_path>/services"
mkdir -p "<contractor_root_path>/prospects"
mkdir -p "<contractor_root_path>/proposals"
mkdir -p "<contractor_root_path>/pitches"
mkdir -p "<contractor_root_path>/rate-cards"
mkdir -p "<contractor_root_path>/landing-pages"
```

## Step 4: Template installation

Copy the `service_definition_schema.json` template from the plugin into the workspace's active template directory:

```bash
# Read templates.base_dir from existing .jobops/config.json
TEMPLATES_BASE=$(python3 -c "import json; print(json.load(open('.jobops/config.json'))['templates']['base_dir'])")
cp "${CLAUDE_PLUGIN_ROOT}/templates/service_definition_schema.json" "${TEMPLATES_BASE}/default/"
```

Verify:

```bash
ls -1 "${TEMPLATES_BASE}/default/service_definition_schema.json"
```

## Step 5: Extend config

Read `.jobops/config.json`, add the following keys while preserving all existing content:

- `directories.contractor_root` = confirmed path from Step 2
- `preferences.default_currency` = confirmed value from Step 2
- `templates.active.service_definition_schema` = `"default"`

Example before (jobops-only):
```json
{
  "directories": { "resume_source": "./ResumeSourceFolder", ... },
  "preferences": { "cultural_profile": "canadian", "default_jurisdiction": "CA-ON" },
  "templates": { "active": { "assessment_rubric_framework": "default", ... } }
}
```

Example after (extended):
```json
{
  "directories": { "resume_source": "./ResumeSourceFolder", ..., "contractor_root": "./Contractor" },
  "preferences": { "cultural_profile": "canadian", "default_jurisdiction": "CA-ON", "default_currency": "CAD" },
  "templates": { "active": { "assessment_rubric_framework": "default", ..., "service_definition_schema": "default" } }
}
```

Write the updated config back to `.jobops/config.json`.

## Step 6: Gitignore update

Propose appending `<contractor_root>` to `.gitignore` if it is not already present:

```bash
if ! grep -Fxq "<contractor_root>" .gitignore 2>/dev/null; then
  echo "Proposed .gitignore addition: <contractor_root>"
  # Ask user: Append? (Y/N)
fi
```

## Step 7: Summary

Report:
- Contractor root: `(created)` / `(existed)` — with subfolders list
- Template installed: `service_definition_schema.json`
- Config extended with: `directories.contractor_root`, `preferences.default_currency`, `templates.active.service_definition_schema`
- Gitignore: `(updated)` / `(skipped)`

### Next steps

1. `/jobops-ic:defineservices` — define your consulting services and pricing
2. `/jobops-ic:ratecard` — generate a rate card document
3. `/jobops-ic:findclient <job-file>` — identify prospects matching your services
```

- [ ] **Step 2: Verify frontmatter**

```bash
head -4 plugins/jobops-ic/skills/setup/SKILL.md
```

Expected: frontmatter with `description:` and `disable-model-invocation: true`.

- [ ] **Step 3: Commit**

```bash
git add plugins/jobops-ic/skills/setup/SKILL.md
git commit -m "$(cat <<'EOF'
Rewrite jobops-ic:setup as light config-extension flow

Replaces the March 2026 setup with a focused 7-step extension of the
existing .jobops/config.json: contractor_root, default_currency, and
service_definition_schema template. Prerequisite check fails fast if
jobops:setup has not been run first.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

## Phase 3: Migration skill

### Task 7: Create standalone `/jobops:migrate` skill

**Files:**
- Create: `plugins/jobops/skills/migrate/SKILL.md`

- [ ] **Step 1: Create the skill directory**

```bash
mkdir -p plugins/jobops/skills/migrate
```

- [ ] **Step 2: Write the skill file**

Create `plugins/jobops/skills/migrate/SKILL.md`:

```markdown
---
description: Migrate legacy flat-layout outputs (OutputResumes, Briefing_Notes, Scoring_Rubrics, Intelligence_Reports) into the new application-centric layout
disable-model-invocation: true
---

# JobOps Legacy Migration

Relocates pre-v2.0 artifact files into the new `Applications/` and `Company_Intelligence/` layouts. Same logic as `/jobops:setup` Step 9, extracted so it can be run standalone after setup or on subsequent batches of legacy files.

## Configuration

Read `.jobops/config.json`. If missing, stop with:

  JOBOPS NOT CONFIGURED
  Run /jobops:setup to initialize your workspace.

Use `config.directories.applications_root` and `config.directories.company_intelligence` for target paths.

## Step 1: Scan legacy folders

```bash
for dir in OutputResumes Briefing_Notes Scoring_Rubrics Intelligence_Reports; do
  if [ -d "$dir" ] && [ "$(ls -A "$dir" 2>/dev/null)" ]; then
    count=$(find "$dir" -type f | wc -l)
    echo "$dir: $count files"
  fi
done
```

If no legacy folders contain files, report "No legacy files found." and exit.

## Step 2: Dry-run parse

For each file in each legacy folder, attempt to extract `{Company}`, `{Role}`, and `{YYYYMMDD}` using these ordered patterns (first match wins):

1. `Rubric_(?P<company>[^_]+)_(?P<role>[^_]+)_(?P<date>\d{8})\.md` — Scoring_Rubrics style (compact date)
2. `Rubric_(?P<company>[^_]+)_(?P<role>[^_]+)_(?P<date>\d{4}-\d{2}-\d{2})\.md` — dash-date variant
3. `Briefing_(?P<company>[^_]+)_(?P<role>[^_]+)(?:_\w+)*_(?P<date>\d{8})\.md` — Briefing_Notes style
4. `InterviewPrep_(?P<company>[^_]+)_(?P<role>[^_]+)_(?P<date>\d{8})\.md` — interview prep variant
5. `(?P<company>[^_]+)_(?P<type>\w+)_Intelligence_(?P<date>\d{4}-\d{2}-\d{2})\.md` — Intelligence_Reports style (company-level, role omitted)
6. `(?P<date>\d{4}-\d{2}-\d{2})_\d+_(?P<company>[^_]+)_(?P<role>[^/]+)` — OutputResumes directory-per-run style (date first)
7. `Step\d+_(?:Draft|Provenance|Final)(?:_Resume)?_(?P<company>[^_]+)_(?P<role>[^_]+)_(?P<date>\d{8})\.md` — OutputResumes flat-file style

Normalize dates to `YYYYMMDD` (drop dashes).

Map each file to its destination:

| Source folder | Pattern type | Destination |
|---|---|---|
| `Scoring_Rubrics/` | 1 or 2 | `{applications_root}/{Company}_{Role}_{YYYYMMDD}/assessment/rubric.md` |
| `Briefing_Notes/` | 3 | `{applications_root}/{Company}_{Role}_{YYYYMMDD}/interview/briefing.md` |
| `Briefing_Notes/` | 4 | `{applications_root}/{Company}_{Role}_{YYYYMMDD}/interview/interview_prep.md` |
| `Intelligence_Reports/` | 5 | `{company_intelligence}/{Company}/{type_lowercased}.md` |
| `OutputResumes/` (file) | 7 | `{applications_root}/{Company}_{Role}_{YYYYMMDD}/resume/{step_lowercased}.md` |
| `OutputResumes/` (subdir) | 6 | `{applications_root}/{Company}_{Role}_{YYYYMMDD}/resume/` (copy directory contents) |

Files that match no pattern are listed as **unresolved** and skipped.

Special cases for pattern 7 step-name mapping:
- `Step1_Draft_Resume_...` → `step1_draft.md`
- `Step2_Provenance_...` → `step2_provenance.md` (or `step2_provenance_analysis.md` if the filename has "Analysis")
- `Step3_Final_Resume_...` → `step3_final.md`

Special cases for pattern 5 type mapping:
- `Corporate_Intelligence` → `corporate.md`
- `Culture_Intelligence` → `culture.md`
- `Leadership_Intelligence` → `leadership.md`
- `Legal_Intelligence` → `legal.md`
- `Market_Intelligence` → `market.md`
- `Compensation_Intelligence` → `compensation.md`
- `Master_Intelligence` → `summary.md`
- Files prefixed `EXAMPLE_` → unresolved (don't migrate examples)

## Step 3: Preview

Display the planned moves, grouped by target app folder (for Applications-bound files) and by company (for intelligence files). List unresolved files separately.

Ask the user:
- **accept** — execute all planned moves
- **reassign** — user specifies a different target for specific files; update the mapping then preview again
- **cancel** — no files moved, exit skill

## Step 4: Execute

For each accepted move, create the target directory and move the file, preferring `git mv` to preserve history:

```bash
TARGET_DIR=$(dirname "$TARGET_PATH")
mkdir -p "$TARGET_DIR"

if git ls-files --error-unmatch "$LEGACY_FILE" > /dev/null 2>&1; then
  git mv "$LEGACY_FILE" "$TARGET_PATH"
else
  mv "$LEGACY_FILE" "$TARGET_PATH"
fi
```

## Step 5: Update config

Read `.jobops/config.json`, update `migration.completed`, `migration.completed_at` (ISO 8601), and `migration.files_moved` (add the new count to any existing count).

Preserve all other config fields exactly.

## Step 6: Report

Summarize:
- Total files moved this run: N
- Total files moved to date (cumulative): M
- Legacy folders now empty (safe to delete): list
- Unresolved files still in place: list with reason each was unresolvable
```

- [ ] **Step 3: Verify file is well-formed**

```bash
head -4 plugins/jobops/skills/migrate/SKILL.md
wc -l plugins/jobops/skills/migrate/SKILL.md
```

Expected: frontmatter present, ~130 lines.

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills/migrate/SKILL.md
git commit -m "$(cat <<'EOF'
Add /jobops:migrate skill for standalone legacy file migration

Extracts the migration logic from /jobops:setup Step 9 into a first-class
skill. Seven regex patterns cover Scoring_Rubrics, Briefing_Notes,
Intelligence_Reports, and OutputResumes naming conventions. Dry-run
preview with user-editable mapping before execution.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

## Phase 4: Skill preamble and path updates

Each task in this phase edits one category of skill files. For every skill file:

1. Replace the existing `## Configuration` preamble (if any) with the appropriate standard preamble from the "Standard Skill Preamble" reference section at the top of this plan.
2. Add the `## Templates` block if the skill reads templates.
3. Add the `## Application-Path Resolution` block if the skill writes to an application folder.
4. Add the `## Company-Intelligence Resolution` block if the skill is an OSINT skill.
5. Replace all hardcoded output paths (e.g., `OutputResumes/Step3_Final_...`) with `{config.directories.<key>}` references following the new layout.
6. Remove any `/tmp/.jobops-plugin-root` references and replace with `${CLAUDE_PLUGIN_ROOT}` where plugin-relative paths are needed.
7. Remove any handwritten "verify jobops plugin is installed" prerequisite blocks (jobops-ic skills only).

### Task 8: Update resume-development skills (3 files)

**Files:**
- Modify: `plugins/jobops/skills/buildresume/SKILL.md`
- Modify: `plugins/jobops/skills/provenance-check/SKILL.md`
- Modify: `plugins/jobops/skills/coverletter/SKILL.md`

- [ ] **Step 1: Read current buildresume SKILL.md**

```bash
cat plugins/jobops/skills/buildresume/SKILL.md
```

Note the current preamble (typically starts with "## Configuration" referring to `.jobops/config.json`) and any hardcoded paths.

- [ ] **Step 2: Replace buildresume preamble**

In `plugins/jobops/skills/buildresume/SKILL.md`, locate the existing `## Configuration` block (lines ~6–15 in the current file) and replace it with the Standard jobops preamble + Templates block + Application-Path Resolution block.

The Templates block for `buildresume` lists:
```
Templates referenced by this skill:
  - assessment_rubric_framework
  - evidence_verification_framework
  - assessment_report_structure
  - candidate_profile_schema
```

The Application-Path Resolution block uses category `resume`.

Replace all hardcoded path references throughout the skill:
- `OutputResumes/Step1_Draft_*` → `{app_folder}/resume/step1_draft.md`
- `OutputResumes/Step2_Provenance_Analysis_*` → `{app_folder}/resume/step2_provenance.md`
- `OutputResumes/Step3_Final_Resume_*` → `{app_folder}/resume/step3_final.md`

Where `{app_folder}` = `{config.directories.applications_root}/{app_slug}/` resolved per the protocol.

- [ ] **Step 3: Replace provenance-check preamble**

In `plugins/jobops/skills/provenance-check/SKILL.md`, apply the same pattern:
- Standard jobops preamble
- Templates block listing `evidence_verification_framework` and `candidate_profile_schema`
- Application-Path Resolution block, category `resume`
- Output path: `{app_folder}/resume/step2_provenance.md`

- [ ] **Step 4: Replace coverletter preamble**

In `plugins/jobops/skills/coverletter/SKILL.md`, apply the same pattern:
- Standard jobops preamble (no Templates block — coverletter does not consume templates)
- Application-Path Resolution block, category `cover-letter`
- Output path: `{app_folder}/cover-letter/cover_letter.md`

- [ ] **Step 5: Verify all three files have valid frontmatter**

```bash
for f in plugins/jobops/skills/{buildresume,provenance-check,coverletter}/SKILL.md; do
  echo "=== $f ==="
  head -4 "$f"
done
```

Expected: each shows `---`, `description:`, `disable-model-invocation: true`, `---`.

- [ ] **Step 6: Commit**

```bash
git add plugins/jobops/skills/{buildresume,provenance-check,coverletter}/SKILL.md
git commit -m "$(cat <<'EOF'
Update resume-development skills for v2.0 config and path layout

buildresume, provenance-check, and coverletter now read .jobops/config.json,
resolve output paths via the application-centric protocol, and write to
{applications_root}/{app_slug}/resume/ or cover-letter/.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 9: Update assessment skills (4 files)

**Files:**
- Modify: `plugins/jobops/skills/assessjob/SKILL.md`
- Modify: `plugins/jobops/skills/assesscandidate/SKILL.md`
- Modify: `plugins/jobops/skills/createrubric/SKILL.md`
- Modify: `plugins/jobops/skills/comparejobs/SKILL.md`

- [ ] **Step 1: Update assessjob**

Replace preamble with:
- Standard jobops preamble
- Templates block listing `assessment_rubric_framework`, `evidence_verification_framework`, `assessment_report_structure`, `candidate_profile_schema`
- Application-Path Resolution block, category `assessment`
- Output paths: `{app_folder}/assessment/rubric.md` (if skill creates rubric) and `{app_folder}/assessment/assessment.md`

Replace all hardcoded path references (`Scoring_Rubrics/Rubric_*`, `OutputResumes/Assessment_*`) with the new resolved paths.

- [ ] **Step 2: Update assesscandidate**

Replace preamble with:
- Standard jobops preamble
- Templates block listing `assessment_rubric_framework`, `evidence_verification_framework`, `assessment_report_structure`, `candidate_profile_schema`
- Application-Path Resolution block, category `assessment`
- Output: `{app_folder}/assessment/assessment.md`

Takes an existing rubric as input — reads it from `{app_folder}/assessment/rubric.md` rather than from the legacy `Scoring_Rubrics/` folder.

- [ ] **Step 3: Update createrubric**

Replace preamble with:
- Standard jobops preamble
- Templates block listing `assessment_rubric_framework`, `evidence_verification_framework`
- Application-Path Resolution block, category `assessment`
- Output: `{app_folder}/assessment/rubric.md`

- [ ] **Step 4: Update comparejobs**

Replace preamble with:
- Standard jobops preamble (no Templates block)
- Cross-application skill — reads from multiple app folders and writes to `{config.directories.career_analysis}/comparison_{YYYYMMDD}_{companyA}_vs_{companyB}.md`

Input args: two or more app slugs or app folder paths.

- [ ] **Step 5: Verify frontmatter on all four files**

```bash
for f in plugins/jobops/skills/{assessjob,assesscandidate,createrubric,comparejobs}/SKILL.md; do
  echo "=== $f ==="
  head -4 "$f"
done
```

- [ ] **Step 6: Commit**

```bash
git add plugins/jobops/skills/{assessjob,assesscandidate,createrubric,comparejobs}/SKILL.md
git commit -m "$(cat <<'EOF'
Update assessment skills for v2.0 config and path layout

assessjob, assesscandidate, createrubric write to {app_folder}/assessment/;
comparejobs writes cross-application comparisons to career_analysis/.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 10: Update interview-prep skills (2 files)

**Files:**
- Modify: `plugins/jobops/skills/briefing/SKILL.md`
- Modify: `plugins/jobops/skills/interviewprep/SKILL.md`

- [ ] **Step 1: Update briefing**

Replace preamble with:
- Standard jobops preamble
- Templates block listing `assessment_report_structure`
- Application-Path Resolution block, category `interview`
- Output: `{app_folder}/interview/briefing.md`

Replace all hardcoded path references (`Briefing_Notes/Briefing_*`).

- [ ] **Step 2: Update interviewprep**

Replace preamble with:
- Standard jobops preamble (no Templates block)
- Application-Path Resolution block, category `interview`
- Output: `{app_folder}/interview/interview_prep.md`

Replace all hardcoded path references (`Briefing_Notes/InterviewPrep_*`).

- [ ] **Step 3: Verify frontmatter**

```bash
for f in plugins/jobops/skills/{briefing,interviewprep}/SKILL.md; do
  echo "=== $f ==="
  head -4 "$f"
done
```

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills/{briefing,interviewprep}/SKILL.md
git commit -m "$(cat <<'EOF'
Update interview-prep skills for v2.0 config and path layout

briefing and interviewprep write to {app_folder}/interview/.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 11: Update job-search skills (2 files)

**Files:**
- Modify: `plugins/jobops/skills/osint/SKILL.md`
- Modify: `plugins/jobops/skills/auditjobposting/SKILL.md`

- [ ] **Step 1: Update osint**

Replace preamble with:
- Standard jobops preamble (no Templates block — OSINT agents don't use templates)
- Company-Intelligence Resolution block (not Application-Path Resolution — OSINT is company-level)
- Output: `{config.directories.company_intelligence}/{Company}/`, one file per OSINT sub-agent:
  - `corporate.md`, `culture.md`, `leadership.md`, `legal.md`, `market.md`, `compensation.md`, `summary.md`
  - `people/{interviewer_name}.md` for `osint-person` agent outputs

When the target folder exists, implement the refresh/append/skip prompt:
- **refresh:** overwrite all files in the target folder
- **append:** write `summary_{YYYYMMDD}.md` alongside existing files (leave others intact)
- **skip:** no action, exit skill

- [ ] **Step 2: Update auditjobposting**

Replace preamble with:
- Standard jobops preamble (no Templates block)
- Application-Path Resolution block, category `assessment` (audit produces an assessment-adjacent artifact)
- Output: `{app_folder}/assessment/jd_audit.md`

- [ ] **Step 3: Verify frontmatter**

```bash
for f in plugins/jobops/skills/{osint,auditjobposting}/SKILL.md; do
  echo "=== $f ==="
  head -4 "$f"
done
```

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills/{osint,auditjobposting}/SKILL.md
git commit -m "$(cat <<'EOF'
Update job-search skills for v2.0 config and path layout

osint writes to company-level {company_intelligence}/{Company}/ with
refresh/append/skip handling on existing folders. auditjobposting writes
to {app_folder}/assessment/jd_audit.md.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 12: Update career-strategy skills (3 files)

**Files:**
- Modify: `plugins/jobops/skills/idealjob/SKILL.md`
- Modify: `plugins/jobops/skills/change-one-thing/SKILL.md`
- Modify: `plugins/jobops/skills/assess-job-offer/SKILL.md`

- [ ] **Step 1: Update idealjob**

Replace preamble with:
- Standard jobops preamble (no Templates block)
- No Application-Path Resolution (flat output under `career_analysis`)
- Output: `{config.directories.career_analysis}/idealjob_{YYYYMMDD}.md`

Uses `config.preferences.cultural_profile` to shape the voice.

- [ ] **Step 2: Update change-one-thing**

Replace preamble with:
- Standard jobops preamble (no Templates block)
- No Application-Path Resolution (flat output)
- Output: `{config.directories.career_analysis}/change_one_thing_{YYYYMMDD}.md`

- [ ] **Step 3: Update assess-job-offer**

Replace preamble with:
- Standard jobops preamble (no Templates block)
- No Application-Path Resolution (the offer analysis is semi-application-scoped; but since the user may not have a full application folder for exploratory offers, place outputs under `career_analysis` with a descriptive filename)
- Output: `{config.directories.career_analysis}/offer_assessment_{Company}_{YYYYMMDD}.md`

Uses `config.preferences.default_jurisdiction` for any legal/tax context.

- [ ] **Step 4: Verify frontmatter**

```bash
for f in plugins/jobops/skills/{idealjob,change-one-thing,assess-job-offer}/SKILL.md; do
  echo "=== $f ==="
  head -4 "$f"
done
```

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/skills/{idealjob,change-one-thing,assess-job-offer}/SKILL.md
git commit -m "$(cat <<'EOF'
Update career-strategy skills for v2.0 config and path layout

idealjob, change-one-thing, and assess-job-offer write flat-output files
to {config.directories.career_analysis}/ with timestamped filenames.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 13: Update career-crisis skills (11 files)

**Files:**
- Modify: `plugins/jobops/skills/code-red/SKILL.md`
- Modify: `plugins/jobops/skills/severance-review/SKILL.md`
- Modify: `plugins/jobops/skills/workplace-documentation/SKILL.md`
- Modify: `plugins/jobops/skills/non-compete-analysis/SKILL.md`
- Modify: `plugins/jobops/skills/reference-shield/SKILL.md`
- Modify: `plugins/jobops/skills/unemployment-prep/SKILL.md`
- Modify: `plugins/jobops/skills/discrimination-assessment/SKILL.md`
- Modify: `plugins/jobops/skills/investigation-prep/SKILL.md`
- Modify: `plugins/jobops/skills/accommodation-request/SKILL.md`
- Modify: `plugins/jobops/skills/layoff-intel/SKILL.md`
- Modify: `plugins/jobops/skills/constructive-dismissal/SKILL.md`

All 11 skills share the same pattern: Standard jobops preamble, no Templates block, flat output under `crisis_management` with a descriptive timestamped filename, jurisdiction-aware via `config.preferences.default_jurisdiction` (overridable per-invocation with `--jurisdiction=X`).

- [ ] **Step 1: Update each file's preamble**

For each of the 11 files, replace the existing preamble with the Standard jobops preamble. Remove any `default_jurisdiction` hardcoded to `CA-ON` from skill content; read it from config instead.

- [ ] **Step 2: Update each file's output path**

Replace the current output path with a `crisis_management`-rooted path:

| Skill | Output path |
|---|---|
| `code-red` | `{crisis_management}/code_red_plan_{YYYYMMDD}.md` |
| `severance-review` | `{crisis_management}/severance_review_{YYYYMMDD}.md` |
| `workplace-documentation` | `{crisis_management}/workplace_documentation_log.md` *(single continuously-updated file)* |
| `non-compete-analysis` | `{crisis_management}/non_compete_analysis_{YYYYMMDD}.md` |
| `reference-shield` | `{crisis_management}/reference_shield_{YYYYMMDD}.md` |
| `unemployment-prep` | `{crisis_management}/unemployment_prep_{YYYYMMDD}.md` |
| `discrimination-assessment` | `{crisis_management}/discrimination_assessment_{YYYYMMDD}.md` |
| `investigation-prep` | `{crisis_management}/investigation_prep_{YYYYMMDD}.md` |
| `accommodation-request` | `{crisis_management}/accommodation_request_{YYYYMMDD}.md` |
| `layoff-intel` | `{crisis_management}/layoff_intel_{Company}_{YYYYMMDD}.md` |
| `constructive-dismissal` | `{crisis_management}/constructive_dismissal_{YYYYMMDD}.md` |

- [ ] **Step 3: Verify frontmatter on all 11 files**

```bash
for skill in code-red severance-review workplace-documentation non-compete-analysis reference-shield unemployment-prep discrimination-assessment investigation-prep accommodation-request layoff-intel constructive-dismissal; do
  f="plugins/jobops/skills/$skill/SKILL.md"
  echo "=== $f ==="
  head -4 "$f"
done
```

Expected: 11 files all showing valid frontmatter.

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills/{code-red,severance-review,workplace-documentation,non-compete-analysis,reference-shield,unemployment-prep,discrimination-assessment,investigation-prep,accommodation-request,layoff-intel,constructive-dismissal}/SKILL.md
git commit -m "$(cat <<'EOF'
Update career-crisis skills for v2.0 config and path layout

All 11 crisis skills now read .jobops/config.json, use
preferences.default_jurisdiction (removing hardcoded CA-ON defaults),
and write flat-output files to config.directories.crisis_management
with timestamped filenames. workplace-documentation remains a single
continuously-updated log.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 14: Update application-finalization skills (3 files)

**Files:**
- Modify: `plugins/jobops/skills/convert-to-pdf/SKILL.md`
- Modify: `plugins/jobops/skills/convert-to-word/SKILL.md`
- Modify: `plugins/jobops/skills/markdown-to-pdf/SKILL.md`

These skills operate on existing files (convert `.md` to PDF/DOCX). They don't produce new application-specific outputs — they transform inputs. The preamble mainly validates config exists and references `config.directories.applications_root` for default input locations.

- [ ] **Step 1: Update convert-to-pdf**

Replace preamble with:
- Standard jobops preamble
- No Templates block
- No Application-Path Resolution block (not an output-creating skill in the app-centric sense)

The skill takes a markdown file path as input and writes the PDF alongside it. Default input search location is `{config.directories.applications_root}/` (recursive) when the user doesn't specify a full path.

- [ ] **Step 2: Update convert-to-word**

Same pattern as Step 1. Output: DOCX alongside the input file.

- [ ] **Step 3: Update markdown-to-pdf**

Same pattern as Step 1. Same output behavior.

- [ ] **Step 4: Verify frontmatter**

```bash
for f in plugins/jobops/skills/{convert-to-pdf,convert-to-word,markdown-to-pdf}/SKILL.md; do
  echo "=== $f ==="
  head -4 "$f"
done
```

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/skills/{convert-to-pdf,convert-to-word,markdown-to-pdf}/SKILL.md
git commit -m "$(cat <<'EOF'
Update finalization skills for v2.0 config

convert-to-pdf, convert-to-word, and markdown-to-pdf now read
.jobops/config.json for default input search locations and validate
that the workspace is configured before running.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 15: Update system-setup skills (2 files)

**Files:**
- Modify: `plugins/jobops/skills/install-pandoc/SKILL.md`
- Modify: `plugins/jobops/skills/github-portfolio/SKILL.md`

These are utility skills with minimal path dependencies.

- [ ] **Step 1: Update install-pandoc**

Replace preamble with the Standard jobops preamble (config-existence check is valuable as a consistency signal; the skill doesn't itself use directory paths).

Alternative: if the install-pandoc skill is purely environmental (installs pandoc system-wide, has no workspace interaction), skip the Configuration block entirely and rely on the skill's existing guidance. Decision: include the Standard jobops preamble for consistency across all skills.

- [ ] **Step 2: Update github-portfolio**

Replace preamble with the Standard jobops preamble. This skill reads the workspace to generate portfolio documentation. It may reference `config.directories.resume_source` as input.

- [ ] **Step 3: Verify frontmatter**

```bash
for f in plugins/jobops/skills/{install-pandoc,github-portfolio}/SKILL.md; do
  echo "=== $f ==="
  head -4 "$f"
done
```

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills/{install-pandoc,github-portfolio}/SKILL.md
git commit -m "$(cat <<'EOF'
Update system-setup skills for v2.0 config consistency

install-pandoc and github-portfolio now carry the Standard jobops
preamble for consistency with other skills.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 16: Update jobops-ic skills (9 files)

**Files:**
- Modify: `plugins/jobops-ic/skills/defineservices/SKILL.md`
- Modify: `plugins/jobops-ic/skills/findclient/SKILL.md`
- Modify: `plugins/jobops-ic/skills/pitchdeck/SKILL.md`
- Modify: `plugins/jobops-ic/skills/proposaltemplate/SKILL.md`
- Modify: `plugins/jobops-ic/skills/ratecard/SKILL.md`
- Modify: `plugins/jobops-ic/skills/copywrite/SKILL.md`
- Modify: `plugins/jobops-ic/skills/copywriting-spec/SKILL.md`
- Modify: `plugins/jobops-ic/skills/create-landing-page/SKILL.md`
- Modify: `plugins/jobops-ic/skills/css-template/SKILL.md`

All 9 skills share the same pattern: Standard jobops-ic preamble (simpler — no handwritten prerequisite check since `dependencies` handles it), flat output under `contractor_root/<subfolder>/`.

- [ ] **Step 1: Update each file's preamble**

For each of the 9 files:
- Remove any existing `## Prerequisite Check` block (handwritten "verify jobops plugin is installed" check)
- Replace the `## Configuration` block with the Standard jobops-ic preamble
- Add a Templates block only for `defineservices` (which uses `service_definition_schema`)

- [ ] **Step 2: Update each file's output path**

| Skill | Output path |
|---|---|
| `defineservices` | `{contractor_root}/services/{service_name}.md` |
| `findclient` | `{contractor_root}/prospects/{company}_{YYYYMMDD}.md` |
| `pitchdeck` | `{contractor_root}/pitches/{prospect_or_industry}_{YYYYMMDD}.md` |
| `proposaltemplate` | `{contractor_root}/proposals/{client}_{type}_{YYYYMMDD}.md` |
| `ratecard` | `{contractor_root}/rate-cards/ratecard_{YYYYMMDD}.md` |
| `copywrite` | `{contractor_root}/landing-pages/{page_name}_copy.md` |
| `copywriting-spec` | `{contractor_root}/landing-pages/{page_name}_spec.md` |
| `create-landing-page` | `{contractor_root}/landing-pages/{page_name}/` (directory with HTML/CSS files) |
| `css-template` | `{contractor_root}/landing-pages/_templates/{template_name}.css` |

- [ ] **Step 3: Verify frontmatter on all 9 files**

```bash
for skill in defineservices findclient pitchdeck proposaltemplate ratecard copywrite copywriting-spec create-landing-page css-template; do
  f="plugins/jobops-ic/skills/$skill/SKILL.md"
  echo "=== $f ==="
  head -4 "$f"
done
```

- [ ] **Step 4: Confirm no handwritten prerequisite checks remain**

```bash
grep -l "resume-summarizer agent is available" plugins/jobops-ic/skills/*/SKILL.md
grep -l "PREREQUISITE MISSING" plugins/jobops-ic/skills/*/SKILL.md
```

Expected: no output from either grep (all handwritten prerequisite checks removed).

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops-ic/skills/{defineservices,findclient,pitchdeck,proposaltemplate,ratecard,copywrite,copywriting-spec,create-landing-page,css-template}/SKILL.md
git commit -m "$(cat <<'EOF'
Update jobops-ic skills for v2.0 config and dependencies field

All 9 non-setup jobops-ic skills drop the handwritten jobops
prerequisite check (now enforced by plugin.json dependencies field),
adopt the Standard jobops-ic preamble, and write to
config.directories.contractor_root subfolders.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

## Phase 5: Documentation

### Task 17: Write `docs/ARCHITECTURE.md`

**Files:**
- Create: `docs/ARCHITECTURE.md`

- [ ] **Step 1: Create the file**

Write `docs/ARCHITECTURE.md` with:

```markdown
# JobOps Plugin Architecture

Contributor-facing reference for the JobOps plugin family (`jobops` + `jobops-ic`). User-facing documentation lives in the README.

## Plugin family

- **`jobops`** — Resume development, interview prep, OSINT, career strategy, career crisis, application finalization. 30 skills, 17 agents.
- **`jobops-ic`** — Independent contractor toolkit. 10 skills, 1 agent. Depends on `jobops` via the `dependencies` field in `plugin.json`.

Both plugins ship from the `jobops-marketplace` in the JobOps repository.

## Configuration contract

User state lives in `.jobops/config.json` in the workspace root. Plugin metadata files (`plugin.json`) contain no user state and do not use the `userConfig` mechanism.

### Schema

See the authoritative schema in `docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md` Section 6.2.

The config has four top-level objects:
- `directories` — paths for input and output roots
- `preferences` — user-level preferences (profile, jurisdiction, currency)
- `templates` — base directory and per-template active variant
- `migration` — legacy-migration state

### Ownership

- `/jobops:setup` owns creation and primary mutation of `.jobops/config.json`.
- `/jobops-ic:setup` extends the file (adds `contractor_root`, `default_currency`, `service_definition_schema` active entry) — does not replace it.
- `/jobops:migrate` updates only the `migration` section.
- Skills treat the config as read-only.

### Missing-file behavior

Every skill except the two setup skills and `migrate` begins by checking `.jobops/config.json` exists. If missing, the skill stops with:

```
JOBOPS NOT CONFIGURED
Run /jobops:setup to initialize your workspace.
```

## Path-resolution protocols

### Application-writing skills

Skills that produce per-application artifacts (resume, cover-letter, assessment, interview) use this protocol:

1. Derive `{Company}_{Role}_{YYYYMMDD}` slug from the job-posting filename, or honor `--app=<slug>` override.
2. Compose app folder: `{config.directories.applications_root}/{app_slug}/`
3. Resolve fixed sub-folder by skill category: `resume`, `cover-letter`, `assessment`, `interview`.
4. If the app folder doesn't exist, `mkdir -p` and copy the job-posting to `{app_folder}/job_posting.md` (provenance pin).

### OSINT skills

Company-scoped output to `{config.directories.company_intelligence}/{Company}/`. One file per OSINT sub-agent. `osint-person` reports go to `{Company}/people/{name}.md`. Existing-folder handling: prompt refresh / append / skip.

### Flat-output skills

Career-level, crisis, and non-application-specific contractor skills emit timestamped files directly under their respective root:

- `{config.directories.career_analysis}/<skill>_{YYYYMMDD}.md`
- `{config.directories.crisis_management}/<skill>_{YYYYMMDD}.md`
- `{config.directories.contractor_root}/<subfolder>/<name>_{YYYYMMDD}.md`

## Skill contract

Every skill must:

1. Declare `disable-model-invocation: true` in frontmatter — skills run only via explicit `/plugin:skill` invocation.
2. Include a `## Configuration` preamble that reads `.jobops/config.json` and stops fast if missing.
3. Include a `## Templates` block only if the skill resolves template paths.
4. Include an `## Application-Path Resolution`, `## Company-Intelligence Resolution`, or flat-output description matching its output category.
5. Never reference `/tmp/.jobops-plugin-root` — that was removed in v2.0. Use `${CLAUDE_PLUGIN_ROOT}` directly when plugin-bundled files are needed.

## Template system

Templates are bundled inside the plugin at `plugins/<plugin>/templates/`. On `/jobops:setup`, they are copied to `.jobops/templates/default/` in the workspace.

To customize a template:
1. Copy `.jobops/templates/default/<template>.md` to `.jobops/templates/custom/<template>.md`.
2. Edit the custom copy.
3. Set `.jobops/config.json` → `templates.active.<template_name>` from `"default"` to `"custom"`.

Skills resolve `{templates.base_dir}/{templates.active.<name>}/<filename>` at runtime.

## Adding a new skill

1. Create `plugins/<plugin>/skills/<skill-name>/SKILL.md` with `disable-model-invocation: true`.
2. Add the Standard preamble (jobops or jobops-ic variant).
3. Add Templates and Path-Resolution blocks matching the skill's output category.
4. Reference `config.directories.<key>` for all file paths.
5. Reference `config.preferences.<key>` for user preferences.
6. Never hardcode a path like `OutputResumes/` or `.jobops/templates/default/`.
7. Update `plugins/<plugin>/templates/` if the skill needs a new template; document it in the Templates block.

## Validation

- `claude plugin validate plugins/jobops` — structural validation of the jobops plugin.
- `claude plugin validate plugins/jobops-ic` — structural validation of jobops-ic.
- `claude plugin validate .` — marketplace validation.

Local development uses `claude --plugin-dir ./plugins/jobops [--plugin-dir ./plugins/jobops-ic]`.

## References

- Design spec: `docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md`
- Implementation plan: `docs/plans/2026-04-23-plugin-config-redesign-implementation.md`
- March 2026 plan (superseded in part): `docs/plans/2026-03-07-plugin-split-design.md`
```

- [ ] **Step 2: Verify the file**

```bash
wc -l docs/ARCHITECTURE.md
head -5 docs/ARCHITECTURE.md
```

Expected: ~150 lines, starts with `# JobOps Plugin Architecture`.

- [ ] **Step 3: Commit**

```bash
git add docs/ARCHITECTURE.md
git commit -m "$(cat <<'EOF'
Add docs/ARCHITECTURE.md contributor reference

Documents the config contract, path-resolution protocols, skill contract,
template system, and how to add new skills. References the design spec
and implementation plan for deeper detail.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 18: Update README setup section

**Files:**
- Modify: `README.md`

The README was rewritten for marketplace distribution in commit `85cbc9b`. The setup section may reference the old `/tmp/` flow or old directory layout. Update only what's stale.

- [ ] **Step 1: Read the README**

Run: `cat README.md | head -120`

Identify any references to:
- `/tmp/.jobops-plugin-root`
- `OutputResumes/`, `Briefing_Notes/`, `Scoring_Rubrics/`, `Intelligence_Reports/` as expected folder names
- Old config schema with `sample_output` key

- [ ] **Step 2: Update the Quick Start section**

If the Quick Start section doesn't already match, ensure it reads:

```markdown
## Quick Start

Install:

```bash
/plugin marketplace add reggiechan74/JobOps
/plugin install jobops@jobops-marketplace
/jobops:setup
```

Optional — add the independent contractor toolkit:

```bash
/plugin install jobops-ic@jobops-marketplace
/jobops-ic:setup
```

`/jobops:setup` walks you through:
- Directory configuration (where to store resumes, job postings, applications, OSINT reports, etc.)
- Cultural profile and default legal jurisdiction
- Template installation
- Optional career-history import from existing resume files
- Optional migration of legacy outputs
```

- [ ] **Step 3: Update the Output Layout section (if present)**

If the README has an "Output Layout" or similar section, ensure it describes the new application-centric layout:

```markdown
## Output Layout

JobOps organizes artifacts by job application, not by artifact type:

```
<workspace>/
├── Applications/
│   └── {Company}_{Role}_{YYYYMMDD}/
│       ├── job_posting.md        (pinned copy)
│       ├── resume/
│       ├── cover-letter/
│       ├── assessment/
│       └── interview/
├── Company_Intelligence/         OSINT, shared across applications to the same company
├── Career_Analysis/              idealjob, change-one-thing, comparisons
├── Crisis_Management/            severance, non-compete, discrimination, etc.
└── Contractor/                   jobops-ic artifacts (if installed)
```

All top-level paths are configurable via `.jobops/config.json`.
```

- [ ] **Step 4: Commit**

```bash
git add README.md
git commit -m "$(cat <<'EOF'
Refresh README setup and output-layout sections for v2.0 redesign

Quick Start now references /jobops:setup (not direct config editing);
Output Layout shows the application-centric structure.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

### Task 19: Update CHANGELOG.md

**Files:**
- Modify: `CHANGELOG.md`

- [ ] **Step 1: Read current CHANGELOG**

```bash
head -50 CHANGELOG.md
```

Note the existing 2.0.0 entry (if any). The March plan added a 2.0.0 entry covering the plugin-split. Extend it with the redesign content from this plan.

- [ ] **Step 2: Add or extend the 2.0.0 entry**

Ensure the top of `CHANGELOG.md` contains:

```markdown
## [2.0.0] - 2026-04-23

### Breaking Changes

- **Output layout is now application-centric.** Top-level folders `OutputResumes/`, `Briefing_Notes/`, `Scoring_Rubrics/`, and `Intelligence_Reports/` are replaced by `Applications/{Company}_{Role}_{YYYYMMDD}/` (with `resume/`, `cover-letter/`, `assessment/`, `interview/` subfolders) and `Company_Intelligence/{Company}/`.
- **Configuration moved to `.jobops/config.json`.** A new per-workspace file populated by `/jobops:setup` replaces the previous ad-hoc config references in skills.
- **Plugin marketplace distribution.** JobOps is now two plugins (`jobops` + `jobops-ic`) installed via Claude Code's plugin system.
- **v1.x skills removed.** `.claude/commands/` and `.claude/agents/` are replaced by plugin `skills/` and `agents/`.

### Added

- `/jobops:setup` — interactive workspace setup (10-step flow)
- `/jobops:migrate` — standalone migration of v1.x legacy outputs into v2.0 layout
- `/jobops-ic:setup` — extends the shared config for the contractor toolkit
- `docs/ARCHITECTURE.md` — contributor reference for the config and path-resolution contract
- `dependencies` field in `jobops-ic/plugin.json` — enforces the jobops prerequisite at install time

### Changed

- All skills now read `.jobops/config.json` and use `${CLAUDE_PLUGIN_ROOT}` for plugin-bundled files
- `default_jurisdiction` is now a user preference (was hardcoded to `CA-ON` for all crisis skills)
- `cultural_profile` is now a user preference (was per-invocation only)

### Removed

- `/tmp/.jobops-plugin-root` SessionStart hook — obsolete now that `${CLAUDE_PLUGIN_ROOT}` expands in skill markdown
- `scripts/copy-templates.sh` in both plugins — replaced by inline `cp ${CLAUDE_PLUGIN_ROOT}/templates/*`
- Handwritten "verify jobops plugin installed" prerequisite blocks in all 9 non-setup jobops-ic skills

### Migration

Users upgrading from v1.x should:
1. Install the new plugins via `/plugin install jobops@jobops-marketplace`
2. Run `/jobops:setup` — when prompted, accept migration of legacy files
3. Verify the generated `.jobops/config.json` reflects their preferred paths
```

- [ ] **Step 3: Commit**

```bash
git add CHANGELOG.md
git commit -m "$(cat <<'EOF'
Document v2.0.0 breaking changes and migration path in CHANGELOG

Covers the output layout flip, .jobops/config.json introduction,
plugin marketplace distribution, and removed plumbing (/tmp/ hook,
copy-templates.sh, handwritten prerequisite checks).

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

---

## Phase 6: Validation and smoke testing

### Task 20: Plugin structural validation

**Files:** None (verification only)

- [ ] **Step 1: Validate jobops**

Run: `claude plugin validate plugins/jobops`

Expected: No errors. If any skill file has invalid frontmatter, missing required fields, or malformed markdown, the validator reports the specific file. Fix the issue (likely a missed edit from Phase 4) and re-run.

- [ ] **Step 2: Validate jobops-ic**

Run: `claude plugin validate plugins/jobops-ic`

Expected: No errors.

- [ ] **Step 3: Validate marketplace**

Run: `claude plugin validate .`

Expected: "Marketplace validated successfully" or equivalent. Both `jobops` and `jobops-ic` listed with correct versions.

- [ ] **Step 4: No commit (verification only)**

---

### Task 21: Dependency enforcement test

**Files:** None (verification only)

- [ ] **Step 1: Attempt to install jobops-ic alone (should fail)**

In a new Claude Code session:

```bash
claude --plugin-dir ./plugins/jobops-ic
```

Inside the session, invoke `/jobops-ic:setup`. Expected: install-time rejection with a message pointing to the missing `jobops` dependency. If `--plugin-dir` does not enforce dependencies at session-load time, instead run `/plugin validate` inside the session and confirm it surfaces the missing dependency.

- [ ] **Step 2: Install both plugins together (should succeed)**

```bash
claude --plugin-dir ./plugins/jobops --plugin-dir ./plugins/jobops-ic
```

Inside the session, run `/plugin list`. Expected: both `jobops` and `jobops-ic` listed, no errors.

- [ ] **Step 3: No commit**

---

### Task 22: End-to-end smoke on the user's workspace

**Files:** None (runs on workspace data; any resulting file moves become part of the eventual test artifacts)

- [ ] **Step 1: Start Claude Code with both plugins**

```bash
claude --plugin-dir ./plugins/jobops --plugin-dir ./plugins/jobops-ic
```

- [ ] **Step 2: Run `/jobops:setup` on the JobOps repo itself**

Walk through the full 10-step flow. Accept defaults for directories. Accept the gitignore update. Skip career-history import (the repo already has `ResumeSourceFolder/` populated). Proceed to migration.

Expected: `.jobops/config.json` created with the schema from Task 4/5. Templates copied. Migration preview displayed for:
- `OutputResumes/` — 3 subdirectories (currently named `YYYY-MM-DD_HHMMSS_Company_Role`)
- `Briefing_Notes/` — 7 files (3 Alto + 3 GTAA + 1 InterviewPrep)
- `Scoring_Rubrics/` — ~30 files (various rubrics)
- `Intelligence_Reports/` — 8 files (7 Canerector + 1 EXAMPLE)

The `EXAMPLE_Corporate_Intelligence_2025-09-26.md` file should be flagged as unresolved (filename prefix skip rule).

- [ ] **Step 3: Execute migration**

Accept the preview. Confirm files move to their new locations. Verify:

```bash
find Applications/ -type f | head -20
find Company_Intelligence/ -type f
```

Expected: Application folders contain the migrated files, organized by app slug. `Company_Intelligence/Canerector/` has 7 intelligence files (`corporate.md`, `culture.md`, etc.).

- [ ] **Step 4: Smoke test a buildresume invocation**

Pick an existing job posting, e.g., `Job_Postings/Alto_Director_Real_Estate.md`.

Run `/jobops:buildresume Job_Postings/Alto_Director_Real_Estate.md`.

Expected:
- Skill reads `.jobops/config.json`
- Application slug derived (e.g., `Alto_DirectorRealEstate_20260423`)
- App folder created at `Applications/<slug>/`
- `job_posting.md` pinned
- `resume/step1_draft.md`, `step2_provenance.md`, `step3_final.md` generated

- [ ] **Step 5: Smoke test an OSINT invocation**

Run `/jobops:osint Google` (or another company name).

Expected:
- Company-intelligence folder created: `Company_Intelligence/Google/`
- Sub-agent outputs: `corporate.md`, `culture.md`, `leadership.md`, `legal.md`, `market.md`, `compensation.md`, `summary.md`

If `Company_Intelligence/Canerector/` exists (from migration), also test the refresh/append/skip prompt by running `/jobops:osint Canerector`.

- [ ] **Step 6: Smoke test a crisis skill with jurisdiction override**

Run `/jobops:severance-review <fake-agreement.md> --jurisdiction=US-CA`.

Expected:
- Output at `Crisis_Management/severance_review_20260423.md`
- Report uses California jurisdiction (not the default `CA-ON` from config)

- [ ] **Step 7: Smoke test a jobops-ic skill**

Run `/jobops-ic:setup` to extend the config.
Then run `/jobops-ic:ratecard`.

Expected:
- `.jobops/config.json` now has `directories.contractor_root`, `preferences.default_currency`, `templates.active.service_definition_schema`
- `Contractor/` directory created with expected subfolders
- Rate card output at `Contractor/rate-cards/ratecard_20260423.md`

- [ ] **Step 8: Review all generated artifacts**

Manually review a representative artifact from each category. Look for: correct frontmatter, HAM-Z formatting, no placeholder text, proper template rendering, jurisdiction/profile preferences honored.

- [ ] **Step 9: Decide on committing smoke-test artifacts**

Smoke-test runs produce real artifacts in the repo. Two options:
- **Keep:** commit as `samples/` or in the application folders as reference material
- **Discard:** revert the smoke-test changes before final commit

Recommended: keep the migrated files (that's the real v1-to-v2 payload), but discard new smoke-test artifacts from Step 4–7. This task's "commit" is the migration result, not the new skill outputs.

- [ ] **Step 10: Commit the migration result**

```bash
git add Applications/ Company_Intelligence/ .jobops/config.json
git add -u OutputResumes/ Briefing_Notes/ Scoring_Rubrics/ Intelligence_Reports/  # record deletions from migration
git commit -m "$(cat <<'EOF'
Migrate v1 repo data to v2 application-centric layout

Run /jobops:setup on the JobOps repo itself as the primary integration
test of the migration logic. Moves all legacy Briefing_Notes,
Scoring_Rubrics, Intelligence_Reports, and OutputResumes files into the
new Applications/ and Company_Intelligence/ layouts.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>
EOF
)"
```

- [ ] **Step 11: Discard smoke-test artifacts**

```bash
# Remove the files generated during Steps 4-7 (buildresume, osint, severance, ratecard)
# Keep only the migration result.
git status  # inspect what's untracked
# Selectively delete smoke-test output files
```

---

## Phase 7: Final commit and PR

### Task 23: Final review and push

**Files:** None (branch operations only)

- [ ] **Step 1: Review commit log**

```bash
git log --oneline main..feature/plugin-split
```

Expected: The 13 pre-existing commits plus ~23 new commits from this plan, in logical order.

- [ ] **Step 2: Verify clean working tree**

```bash
git status
```

Expected: clean.

- [ ] **Step 3: Run all validations one final time**

```bash
claude plugin validate plugins/jobops
claude plugin validate plugins/jobops-ic
claude plugin validate .
```

All three should pass.

- [ ] **Step 4: Push to remote**

```bash
git push origin feature/plugin-split
```

- [ ] **Step 5: Open PR to main**

```bash
gh pr create --title "v2.0.0: Plugin marketplace distribution with config redesign" --body "$(cat <<'EOF'
## Summary

- Split the monolithic `.claude/` JobOps project into two Claude Code plugins (`jobops` + `jobops-ic`) distributed via a self-hosted marketplace
- Introduce per-workspace config at `.jobops/config.json` populated by `/jobops:setup` (interactive 10-step flow)
- Flip output layout from type-centric (`OutputResumes/`, `Briefing_Notes/`, ...) to application-centric (`Applications/{Company}_{Role}_{YYYYMMDD}/resume/`, `assessment/`, `interview/`, ...)
- Eliminate obsolete plumbing: `/tmp/.jobops-plugin-root` hook, `copy-templates.sh`, handwritten jobops-ic prerequisite checks
- Add `/jobops:migrate` for standalone legacy-layout conversion
- Add `docs/ARCHITECTURE.md` contributor reference

Supersedes the relevant portions of `docs/plans/2026-03-07-plugin-split-design.md` per `docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md`.

## Breaking changes

- v1.x output folders are migrated into the new layout via `/jobops:setup` Step 9 or `/jobops:migrate`
- `.claude/commands/` and `.claude/agents/` are removed in favor of plugin `skills/` and `agents/`
- All skills now require `.jobops/config.json`; pre-existing workflows must run `/jobops:setup` before re-use
- Jurisdiction and cultural profile defaults are now user preferences, not hardcoded

## Test plan

- [x] `claude plugin validate plugins/jobops` — passes
- [x] `claude plugin validate plugins/jobops-ic` — passes
- [x] `claude plugin validate .` — marketplace passes
- [x] Dependency enforcement: installing jobops-ic alone is rejected
- [x] End-to-end smoke on this repo's own data (migration of 45+ legacy files completed successfully)
- [x] `buildresume`, `osint`, `severance-review`, `ratecard` smoke-tested in `--plugin-dir` mode

🤖 Generated with [Claude Code](https://claude.com/claude-code)
EOF
)"
```

- [ ] **Step 6: Record the PR URL**

Capture the output of `gh pr create` for reference.

---

## Appendix: Self-review checklist

Run this checklist after the plan is complete before handing off to execution:

1. **Spec coverage.** Every section of `docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md` must have a corresponding task:
   - §4 Directory layout → covered by Tasks 4, 5, 7 (setup + migrate), and the output paths in every Phase 4 task
   - §5 Plugin manifests → Task 2
   - §6 Workspace config file → Tasks 4, 5, 6
   - §7 Skill contract → Tasks 8–16
   - §8 `/jobops:setup` flow → Tasks 4, 5
   - §8.1 `/jobops-ic:setup` flow → Task 6
   - §9 `/jobops:migrate` → Task 7
   - §10 Scope of changes → Tasks 1, 8–16
   - §11 Test strategy → Tasks 20–22
   - §12 Risks → mitigated inline in Task 7 (parser patterns) and Task 22 (smoke)
   - §13 Non-goals → not implemented; plan carries no tasks for them
2. **Placeholder scan.** No "TBD", "TODO", "fill in details", or "similar to Task N" without repeating the code.
3. **Type consistency.** Property names (`applications_root`, `default_jurisdiction`, etc.) match across all tasks and the spec.
4. **File coverage.** Every file listed in "File Structure Map" appears in at least one task.
