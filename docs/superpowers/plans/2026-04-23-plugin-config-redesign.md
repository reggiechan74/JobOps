# Plugin Config & Output-Layout Redesign — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Implement the plugin config and output-layout redesign: delete the `/tmp/` plugin-root workaround, add plugin dependencies, rewrite both setup skills, add a `/jobops:migrate` skill, standardize skill preambles, and switch all skills to application-centric / flat output paths.

**Architecture:** Single config file at `.jobops/config.json` populated by `/jobops:setup`. Plugin resolves its install path via `${CLAUDE_PLUGIN_ROOT}` directly in skill markdown (no hook, no tmp file). Outputs are organized application-centrically under `Applications/{Company}_{Role}_{YYYYMMDD}/` with fixed sub-folder convention; company-level OSINT is shared under `Company_Intelligence/{Company}/`; career/crisis/contractor emit flat timestamped files.

**Tech Stack:** Claude Code plugin system (v2.0 plugin.json with `dependencies`), JSON workspace config, markdown skills, bash for setup operations, git for history-preserving moves.

**Spec:** `docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md`

**Testing philosophy:** This plugin has no unit-test suite — the deliverables are markdown skill files plus JSON manifests. Verification uses: (a) `claude plugin validate`, (b) JSON/frontmatter lint, (c) grep checks that no skill references obsolete paths, (d) smoke-test invocation against the repo itself. Each task ends with an explicit verification command and an explicit commit.

---

## File Structure

### Files deleted
- `plugins/jobops/hooks/hooks.json` + parent `hooks/` dir
- `plugins/jobops/scripts/copy-templates.sh` + parent `scripts/` dir
- `plugins/jobops-ic/hooks/hooks.json` + parent `hooks/` dir
- `plugins/jobops-ic/scripts/copy-templates.sh` + parent `scripts/` dir

### Files created
- `plugins/jobops/skills/migrate/SKILL.md` — new standalone `/jobops:migrate`
- `docs/ARCHITECTURE.md` — contributor reference for config + path-resolution contract

### Files rewritten (full replace)
- `plugins/jobops/skills/setup/SKILL.md` — new 10-step flow
- `plugins/jobops-ic/skills/setup/SKILL.md` — extend-config flow
- `plugins/jobops-ic/.claude-plugin/plugin.json` — add `dependencies` field
- `plugins/jobops/.claude-plugin/plugin.json` — keyword list update (add `crisis-management`)

### Files edited (preamble + path references)
- 30 `jobops` skill files (all under `plugins/jobops/skills/*/SKILL.md` except `setup/` and `migrate/`)
- 10 `jobops-ic` skill files (all under `plugins/jobops-ic/skills/*/SKILL.md` except `setup/`)

### Unchanged
- Marketplace manifest, agents, templates, CLAUDE.md, README.md, package.json

---

## Preamble templates (reference — used in multiple tasks)

**JOBOPS_PREAMBLE** (used by Tasks 6, 8, 9, 10, 11):

```markdown
## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if this skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if this skill has jurisdiction-sensitive logic (crisis/legal skills accept `--jurisdiction=<ISO-3166-2>` to override).
```

**JOBOPS_IC_PREAMBLE** (used by Task 7):

```markdown
## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup (then /jobops-ic:setup) to initialize your workspace.

Use `config.directories.contractor_root` for output paths in this skill.
Use `config.preferences.default_currency` for pricing if applicable.
```

**TEMPLATES_BLOCK** (appended after preamble in template-consuming skills — Task 6/7):

```markdown
## Templates

For each template used by this skill, resolve the full path as:

  {config.templates.base_dir}/{config.templates.active.<template_name>}/<filename>

Templates referenced by this skill: <EXPLICIT-LIST — see each task for the list>
```

**APPLICATION_PATH_PROTOCOL** (used by application-scoped skills — Task 8):

```markdown
## Application Path Resolution

This skill writes to a per-application folder. Before writing any output:

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the job-posting filename, or honor `--app=<slug>` if supplied.
2. Compose the app folder: `{config.directories.applications_root}/{app_slug}/`.
3. Resolve this skill's sub-folder by category:
   - resume-development (buildresume, provenance-check) → `resume/`
   - cover-letter (coverletter) → `cover-letter/`
   - rubric / assessment (createrubric, assessjob, assesscandidate, auditjobposting) → `assessment/`
   - briefing / interview prep (briefing, interviewprep) → `interview/`
4. If the app folder does not exist, `mkdir -p` it, then copy
   `{config.directories.job_postings}/{filename}` → `{app_slug}/job_posting.md`
   so the pinned JD cannot silently change under completed work.
5. Exact-slug collisions (same Company+Role+Date) are not auto-suffixed. If the folder
   already contains the same output type, require the user to pass `--app=<distinct-slug>`.
```

---

## Skill categorization (used by later tasks)

### jobops — 32 skills total

| Category | Output root | Skills |
|---|---|---|
| setup (rewritten) | n/a | `setup` |
| migrate (new) | n/a | `migrate` |
| application (resume) | `applications_root/{app}/resume/` | `buildresume`, `provenance-check` |
| application (cover-letter) | `applications_root/{app}/cover-letter/` | `coverletter` |
| application (assessment) | `applications_root/{app}/assessment/` | `createrubric`, `assessjob`, `assesscandidate`, `auditjobposting` |
| application (interview) | `applications_root/{app}/interview/` | `briefing`, `interviewprep` |
| company intelligence | `company_intelligence/{Company}/` | `osint`, `layoff-intel` |
| career analysis | `career_analysis/` | `idealjob`, `change-one-thing`, `comparejobs`, `github-portfolio` |
| crisis management | `crisis_management/` | `severance-review`, `non-compete-analysis`, `constructive-dismissal`, `workplace-documentation`, `accommodation-request`, `discrimination-assessment`, `investigation-prep`, `unemployment-prep`, `reference-shield`, `code-red`, `assess-job-offer` |
| utility (no config paths) | n/a | `convert-to-pdf`, `convert-to-word`, `markdown-to-pdf`, `install-pandoc` |

Template-consuming skills (get `TEMPLATES_BLOCK`): `buildresume`, `provenance-check`, `createrubric`, `assessjob`, `assesscandidate`, `briefing`, `interviewprep`, `coverletter`.

### jobops-ic — 10 skills total

| Category | Output root | Skills |
|---|---|---|
| setup (rewritten) | n/a | `setup` |
| services | `contractor_root/services/` | `defineservices` |
| prospects | `contractor_root/prospects/` | `findclient` |
| proposals | `contractor_root/proposals/` | `proposaltemplate` |
| pitches | `contractor_root/pitches/` | `pitchdeck` |
| rate-cards | `contractor_root/rate-cards/` | `ratecard` |
| landing-pages | `contractor_root/landing-pages/` | `create-landing-page`, `copywrite`, `copywriting-spec`, `css-template` |

Template-consuming IC skill (gets `TEMPLATES_BLOCK` with `service_definition_schema`): `defineservices`.

---

## Task 1: Remove obsolete hooks and scripts

**Why:** `${CLAUDE_PLUGIN_ROOT}` now expands inside skill markdown. The SessionStart hook writing to `/tmp/.jobops-plugin-root` and the bash `copy-templates.sh` scripts are no longer needed. Both plugin folders simplify to `.claude-plugin/`, `templates/`, `skills/`, `agents/`.

**Files:**
- Delete: `plugins/jobops/hooks/hooks.json`
- Delete: `plugins/jobops/scripts/copy-templates.sh`
- Delete: `plugins/jobops/hooks/` (directory — empty after file deletion)
- Delete: `plugins/jobops/scripts/` (directory — empty after file deletion)
- Delete: `plugins/jobops-ic/hooks/hooks.json`
- Delete: `plugins/jobops-ic/scripts/copy-templates.sh`
- Delete: `plugins/jobops-ic/hooks/` (directory)
- Delete: `plugins/jobops-ic/scripts/` (directory)

- [ ] **Step 1: Confirm the current contents before deletion**

Run:
```bash
ls plugins/jobops/hooks plugins/jobops/scripts plugins/jobops-ic/hooks plugins/jobops-ic/scripts
```
Expected: each directory contains exactly the two files listed above. If anything else is present, STOP and report — do not delete.

- [ ] **Step 2: Delete the four files via `git rm`**

Run:
```bash
git rm plugins/jobops/hooks/hooks.json \
       plugins/jobops/scripts/copy-templates.sh \
       plugins/jobops-ic/hooks/hooks.json \
       plugins/jobops-ic/scripts/copy-templates.sh
```

- [ ] **Step 3: Remove the now-empty directories**

Run:
```bash
rmdir plugins/jobops/hooks plugins/jobops/scripts plugins/jobops-ic/hooks plugins/jobops-ic/scripts
```
Expected: each `rmdir` succeeds silently (directory is empty).

- [ ] **Step 4: Verify the plugin trees are now four-subdir each**

Run:
```bash
ls plugins/jobops plugins/jobops-ic
```
Expected: each lists exactly `.claude-plugin  agents  skills  templates` (four entries).

- [ ] **Step 5: Confirm no skill references the obsolete `/tmp/` path**

Run:
```bash
grep -rn "/tmp/.jobops" plugins/ || echo "CLEAN"
grep -rn "copy-templates.sh" plugins/ || echo "CLEAN"
```
Expected: `CLEAN` for both (remaining references will be fixed in later tasks; this is a baseline).
Note any hits for later tasks.

- [ ] **Step 6: Commit**

```bash
git add -A plugins/jobops/hooks plugins/jobops/scripts plugins/jobops-ic/hooks plugins/jobops-ic/scripts
git commit -m "Remove obsolete hooks and copy-templates scripts

Plugin root resolution now uses \${CLAUDE_PLUGIN_ROOT} directly in skill
markdown, retiring the /tmp/.jobops-plugin-root workaround and the
bash copy-templates scripts. Both plugin directories simplify from six
subdirectories to four.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 2: Update plugin manifests

**Why:** `jobops-ic` needs a `dependencies: {jobops: "^2.0.0"}` field so Claude Code refuses to install it without `jobops`, replacing the handwritten prerequisite check each IC skill carried. `jobops` keywords gain `crisis-management` to match the v2.0 description.

**Files:**
- Modify: `plugins/jobops/.claude-plugin/plugin.json`
- Modify: `plugins/jobops-ic/.claude-plugin/plugin.json`

- [ ] **Step 1: Rewrite `plugins/jobops/.claude-plugin/plugin.json`**

Replace the entire file content with:
```json
{
  "name": "jobops",
  "version": "2.0.0",
  "description": "Intelligence-driven job application system - resume development, interview prep, OSINT intelligence, career strategy, and crisis management using HAM-Z methodology",
  "author": {
    "name": "Reggie Chan"
  },
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": ["resume", "job-search", "interview-prep", "osint", "career", "assessment", "crisis-management"]
}
```

- [ ] **Step 2: Rewrite `plugins/jobops-ic/.claude-plugin/plugin.json`**

Replace the entire file content with:
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

- [ ] **Step 3: Validate both JSON files parse**

Run:
```bash
python3 -m json.tool plugins/jobops/.claude-plugin/plugin.json > /dev/null && echo OK1
python3 -m json.tool plugins/jobops-ic/.claude-plugin/plugin.json > /dev/null && echo OK2
```
Expected: `OK1` and `OK2` (two lines).

- [ ] **Step 4: Confirm `dependencies` was added**

Run:
```bash
grep -A2 '"dependencies"' plugins/jobops-ic/.claude-plugin/plugin.json
```
Expected: shows `"jobops": "^2.0.0"`.

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/.claude-plugin/plugin.json plugins/jobops-ic/.claude-plugin/plugin.json
git commit -m "Add plugin dependencies and v2.0 keywords

jobops-ic now declares dependencies: {jobops: ^2.0.0}, letting Claude
Code enforce the requirement at install time instead of per-skill
runtime checks. jobops keywords gain crisis-management to match the
v2.0 description.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 3: Rewrite `/jobops:setup` SKILL

**Why:** The old setup creates a legacy directory set (OutputResumes, Briefing_Notes, etc.) and copies templates via the retired bash script. The redesign needs the new 10-step flow (Section 8 of the spec): directory interview, preferences, template install via `${CLAUDE_PLUGIN_ROOT}`, write config, gitignore with preview, optional career-history import, optional legacy migration with dry-run + user mapping, summary.

**Files:**
- Modify (full rewrite): `plugins/jobops/skills/setup/SKILL.md`

- [ ] **Step 1: Read the current file so you know what's being replaced**

Run:
```bash
wc -l plugins/jobops/skills/setup/SKILL.md
```
Note the line count for the commit message.

- [ ] **Step 2: Replace the full file content**

Write the following to `plugins/jobops/skills/setup/SKILL.md`:

```markdown
---
description: Initialize JobOps workspace - configure output directories, install templates, and optionally migrate legacy files
disable-model-invocation: true
argument-hint: [--reconfigure] [--skip-migration] [--skip-history]
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
```

- [ ] **Step 3: Verify the file parses as valid markdown with YAML frontmatter**

Run:
```bash
head -5 plugins/jobops/skills/setup/SKILL.md
```
Expected: starts with `---`, has `description:`, `disable-model-invocation: true`, `argument-hint:`, closing `---`.

- [ ] **Step 4: Confirm no reference to obsolete `/tmp/.jobops-plugin-root`**

Run:
```bash
grep -n "/tmp/.jobops" plugins/jobops/skills/setup/SKILL.md || echo "CLEAN"
```
Expected: `CLEAN`.

- [ ] **Step 5: Confirm `${CLAUDE_PLUGIN_ROOT}` is used for template copy**

Run:
```bash
grep -n 'CLAUDE_PLUGIN_ROOT' plugins/jobops/skills/setup/SKILL.md
```
Expected: at least one hit inside the Step 5 (template installation) section.

- [ ] **Step 6: Commit**

```bash
git add plugins/jobops/skills/setup/SKILL.md
git commit -m "Rewrite /jobops:setup for v2.0 config and output layout

10-step flow: welcome + existing-config check, directory interview
with per-path validation, directory creation, preferences interview
(cultural_profile + default_jurisdiction), template install via
\${CLAUDE_PLUGIN_ROOT}, atomic config write, gitignore update with
preview, optional career-history import, optional legacy migration
with dry-run/preview/user-editable mapping, summary.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 4: Rewrite `/jobops-ic:setup` SKILL

**Why:** The old IC setup creates `client_prospects` and tries to call a retired bash script. The new flow (Section 8.1 of the spec) extends `.jobops/config.json` with `directories.contractor_root`, `preferences.default_currency`, and the service_definition template entry; copies the IC template via `${CLAUDE_PLUGIN_ROOT}`; updates `.gitignore`. It must also short-circuit with a clear error if `.jobops/config.json` does not already exist (the plugin `dependencies` field prevents install without jobops, but the user could still have skipped `/jobops:setup`).

**Files:**
- Modify (full rewrite): `plugins/jobops-ic/skills/setup/SKILL.md`

- [ ] **Step 1: Replace the full file content**

Write the following to `plugins/jobops-ic/skills/setup/SKILL.md`:

```markdown
---
description: Extend JobOps workspace with independent-contractor directories, templates, and config - requires /jobops:setup to have run first
disable-model-invocation: true
---

# JobOps IC Setup

This skill extends an existing JobOps workspace with the directories,
templates, and config entries needed by the `jobops-ic` plugin. It does not
create `.jobops/config.json` — it adds to the one created by `/jobops:setup`.

---

## Step 1: Prerequisite check

Read `.jobops/config.json`.

If the file does not exist, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup first, then retry /jobops-ic:setup.

(The plugin `dependencies` field ensures `jobops` is installed alongside
`jobops-ic`; this check catches the separate case of the user skipping the
base setup.)

---

## Step 2: Welcome

Print:

> **JobOps IC Setup**
>
> This extends your existing JobOps workspace with:
> - A contractor output tree (services, prospects, proposals, pitches, rate cards, landing pages)
> - The IC service-definition template
> - Currency preference for pricing outputs

Continue.

---

## Step 3: Interview

Ask in order:

1. **Contractor root directory** — default `./Contractor`. Validate the parent is writable.
2. **Default currency** — ISO 4217 code, enum: `CAD` | `USD` | `EUR` | `GBP` | `AUD`. Default `CAD`. Used by rate cards and proposals.

---

## Step 4: Create directory tree

Run:
```bash
mkdir -p <contractor_root>/services \
         <contractor_root>/prospects \
         <contractor_root>/proposals \
         <contractor_root>/pitches \
         <contractor_root>/rate-cards \
         <contractor_root>/landing-pages
```

Report created vs exists.

---

## Step 5: Template installation

Copy the IC service-definition schema into the workspace template tree:

```bash
cp ${CLAUDE_PLUGIN_ROOT}/templates/service_definition_schema.json .jobops/templates/default/
```

If the source file is missing, stop with a clear error. Do not silently skip.

---

## Step 6: Extend `.jobops/config.json`

Read the existing config, add the following keys, and write it back
atomically (write to `.jobops/config.json.tmp`, then `mv`):

- `directories.contractor_root = "<step-3 value>"`
- `preferences.default_currency = "<step-3 value>"`
- `templates.active.service_definition_schema = "default"`

Preserve all existing keys and values. Do not change `version`.

---

## Step 7: Gitignore update

If `.gitignore` has a `# JobOps workspace` block, append (if not already
present) the `contractor_root` path on a new line inside that block.
Default: `Contractor/`. If the block is missing, print a reminder to the
user that `/jobops:setup` controls the block and that they can re-run it
to refresh.

---

## Step 8: Summary

Print:

1. Contractor root path.
2. Default currency.
3. Template installed.
4. Recommended next steps: `/jobops-ic:defineservices`, `/jobops-ic:ratecard`.

Exit.
```

- [ ] **Step 2: Verify frontmatter**

Run:
```bash
head -4 plugins/jobops-ic/skills/setup/SKILL.md
```
Expected: `---`, description line, `disable-model-invocation: true`, `---`.

- [ ] **Step 3: Confirm no prerequisite-check boilerplate checking for `resume-summarizer` agent**

Run:
```bash
grep -n "resume-summarizer" plugins/jobops-ic/skills/setup/SKILL.md || echo "CLEAN"
```
Expected: `CLEAN`.

- [ ] **Step 4: Confirm `${CLAUDE_PLUGIN_ROOT}` is used for template copy**

Run:
```bash
grep -n 'CLAUDE_PLUGIN_ROOT' plugins/jobops-ic/skills/setup/SKILL.md
```
Expected: at least one hit.

- [ ] **Step 5: Confirm no reference to `/tmp/.jobops-ic-plugin-root`**

Run:
```bash
grep -n "/tmp/.jobops" plugins/jobops-ic/skills/setup/SKILL.md || echo "CLEAN"
```
Expected: `CLEAN`.

- [ ] **Step 6: Commit**

```bash
git add plugins/jobops-ic/skills/setup/SKILL.md
git commit -m "Rewrite /jobops-ic:setup for shared config extension

Prerequisite check reads .jobops/config.json; exits cleanly if the
base jobops setup has not run. Extends the existing config with
contractor_root, default_currency, and the service_definition_schema
template entry. Template copy uses \${CLAUDE_PLUGIN_ROOT} directly.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 5: Create `/jobops:migrate` SKILL

**Why:** Section 9 of the spec. Users who ran setup early and accumulated legacy files afterward, or who want to re-migrate a new batch, need a first-class skill that runs the same dry-run → preview → user-editable mapping → execute flow that lives inside `/jobops:setup` Step 9.

**Files:**
- Create: `plugins/jobops/skills/migrate/SKILL.md`

- [ ] **Step 1: Create the directory**

Run:
```bash
mkdir -p plugins/jobops/skills/migrate
```

- [ ] **Step 2: Write the skill file**

Write the following to `plugins/jobops/skills/migrate/SKILL.md`:

```markdown
---
description: Migrate legacy JobOps output files into the v2.0 application-centric layout - dry-run preview with user-editable mapping
disable-model-invocation: true
argument-hint: [--dry-run]
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

For each file, attempt to parse `{Company}_{Role}_{YYYYMMDD}` using the
patterns (tried in order):

1. `^([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`
2. `^(?:Step\d+_)?(?:Draft_|Provenance_Analysis_|Final_Resume_|Cover_Letter_)?([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`
3. `^Rubric_([A-Za-z0-9]+)_([A-Za-z0-9]+)_(\d{8})\.md$`

Compose destinations:

- `OutputResumes/Step1_Draft_*` → `{applications_root}/{Company}_{Role}_{Date}/resume/step1_draft.md`
- `OutputResumes/Step2_Provenance_Analysis_*` → `.../resume/step2_provenance.md`
- `OutputResumes/Step3_Final_Resume_*` → `.../resume/step3_final.md`
- `OutputResumes/Cover_Letter_*` → `.../cover-letter/cover_letter.md`
- `Scoring_Rubrics/Rubric_*` → `.../assessment/rubric.md`
- `Briefing_Notes/*` (matches pattern 1 or 2) → `.../interview/briefing.md`
- `Intelligence_Reports/<Company>_Corporate_*` → `{company_intelligence}/<Company>/corporate.md`
- `Intelligence_Reports/<Company>_Legal_*` → `{company_intelligence}/<Company>/legal.md`
- `Intelligence_Reports/<Company>_Leadership_*` → `{company_intelligence}/<Company>/leadership.md`
- `Intelligence_Reports/<Company>_Compensation_*` → `{company_intelligence}/<Company>/compensation.md`
- `Intelligence_Reports/<Company>_Culture_*` → `{company_intelligence}/<Company>/culture.md`
- `Intelligence_Reports/<Company>_Market_*` → `{company_intelligence}/<Company>/market.md`

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
```

- [ ] **Step 3: Verify frontmatter and `argument-hint`**

Run:
```bash
head -5 plugins/jobops/skills/migrate/SKILL.md
```
Expected: YAML frontmatter with description, `disable-model-invocation: true`, `argument-hint: [--dry-run]`.

- [ ] **Step 4: Confirm standard `.jobops/config.json` preamble**

Run:
```bash
grep -n "JOBOPS NOT CONFIGURED" plugins/jobops/skills/migrate/SKILL.md
```
Expected: at least one hit.

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/skills/migrate/SKILL.md
git commit -m "Add /jobops:migrate standalone skill

Extracts the Step 9 migration flow from /jobops:setup into a first-class
skill so users who accumulate legacy files after initial setup can
re-run migration on new batches. Same dry-run → preview →
user-editable mapping → execute flow as the setup variant, plus a
--dry-run flag for preview-only runs.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 6: Standardize preamble in non-setup jobops skills

**Why:** All 30 non-setup jobops skills need the standard `JOBOPS_PREAMBLE` (Section 7.1). Template-consuming skills additionally need `TEMPLATES_BLOCK` listing the specific templates they read. Current preambles vary (some already have Section 7.1-ish text, some don't); this task normalizes them. Path references inside skill bodies are handled in Tasks 8–11 by category.

**Files (30 total):**

All of:
```
plugins/jobops/skills/accommodation-request/SKILL.md
plugins/jobops/skills/assess-job-offer/SKILL.md
plugins/jobops/skills/assesscandidate/SKILL.md
plugins/jobops/skills/assessjob/SKILL.md
plugins/jobops/skills/auditjobposting/SKILL.md
plugins/jobops/skills/briefing/SKILL.md
plugins/jobops/skills/buildresume/SKILL.md
plugins/jobops/skills/change-one-thing/SKILL.md
plugins/jobops/skills/code-red/SKILL.md
plugins/jobops/skills/comparejobs/SKILL.md
plugins/jobops/skills/constructive-dismissal/SKILL.md
plugins/jobops/skills/convert-to-pdf/SKILL.md
plugins/jobops/skills/convert-to-word/SKILL.md
plugins/jobops/skills/coverletter/SKILL.md
plugins/jobops/skills/createrubric/SKILL.md
plugins/jobops/skills/discrimination-assessment/SKILL.md
plugins/jobops/skills/github-portfolio/SKILL.md
plugins/jobops/skills/idealjob/SKILL.md
plugins/jobops/skills/install-pandoc/SKILL.md
plugins/jobops/skills/interviewprep/SKILL.md
plugins/jobops/skills/investigation-prep/SKILL.md
plugins/jobops/skills/layoff-intel/SKILL.md
plugins/jobops/skills/markdown-to-pdf/SKILL.md
plugins/jobops/skills/non-compete-analysis/SKILL.md
plugins/jobops/skills/osint/SKILL.md
plugins/jobops/skills/provenance-check/SKILL.md
plugins/jobops/skills/reference-shield/SKILL.md
plugins/jobops/skills/severance-review/SKILL.md
plugins/jobops/skills/unemployment-prep/SKILL.md
plugins/jobops/skills/workplace-documentation/SKILL.md
```

**Utility skills that should NOT get a config preamble** (they do not touch workspace paths):
`convert-to-pdf`, `convert-to-word`, `markdown-to-pdf`, `install-pandoc` — these stay as-is for the preamble (skip them in Step 2 below). They still appear in the task because their frontmatter must continue to carry `disable-model-invocation: true`.

- [ ] **Step 1: Baseline — list every non-setup skill and record whether each currently has a `## Configuration` heading**

Run:
```bash
for f in plugins/jobops/skills/*/SKILL.md; do
  case "$f" in
    *setup/SKILL.md|*migrate/SKILL.md) continue ;;
  esac
  has=$(grep -c '^## Configuration' "$f" || true)
  echo "$has  $f"
done
```
Note the files that currently have `0` — those are missing the preamble entirely.

- [ ] **Step 2: For each of the 26 non-utility, non-setup, non-migrate skills, replace (or insert) the `## Configuration` block with the JOBOPS_PREAMBLE text**

The 26 target files are all of the above minus: `convert-to-pdf`, `convert-to-word`, `markdown-to-pdf`, `install-pandoc`.

For each target file, the replacement rule is:
- If `## Configuration` exists: replace from that heading up to (but not including) the next `## ` heading at the same depth.
- If it does not exist: insert the block immediately after the YAML frontmatter's closing `---` line.

Use this exact block (verbatim):

```markdown
## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if this skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if this skill has jurisdiction-sensitive logic (crisis/legal skills accept `--jurisdiction=<ISO-3166-2>` to override).

```

(Blank line at the end is required to separate from the next heading.)

- [ ] **Step 3: Append `## Templates` block to the 8 template-consuming skills**

The 8 files:
- `plugins/jobops/skills/buildresume/SKILL.md` — templates: `candidate_profile_schema`, `evidence_verification_framework`
- `plugins/jobops/skills/provenance-check/SKILL.md` — templates: `evidence_verification_framework`
- `plugins/jobops/skills/createrubric/SKILL.md` — templates: `assessment_rubric_framework`
- `plugins/jobops/skills/assessjob/SKILL.md` — templates: `assessment_rubric_framework`, `assessment_report_structure`
- `plugins/jobops/skills/assesscandidate/SKILL.md` — templates: `assessment_rubric_framework`, `assessment_report_structure`, `evidence_verification_framework`
- `plugins/jobops/skills/briefing/SKILL.md` — templates: `assessment_report_structure`
- `plugins/jobops/skills/interviewprep/SKILL.md` — templates: `assessment_report_structure`
- `plugins/jobops/skills/coverletter/SKILL.md` — templates: `candidate_profile_schema`

For each, insert (immediately after the `## Configuration` block, before the next section) this verbatim block with `<EXPLICIT-LIST>` replaced by the template names from the list above, comma-separated:

```markdown
## Templates

For each template used by this skill, resolve the full path as:

  {config.templates.base_dir}/{config.templates.active.<template_name>}/<filename>

Templates referenced by this skill: <EXPLICIT-LIST>

```

(Blank line at the end is required.)

If a `## Templates` block already exists, replace it with the standardized form.

- [ ] **Step 4: Verify preamble is present (and identical) in all 26 target files**

Run:
```bash
for f in plugins/jobops/skills/*/SKILL.md; do
  case "$f" in
    *setup/SKILL.md|*migrate/SKILL.md) continue ;;
    *convert-to-pdf/*|*convert-to-word/*|*markdown-to-pdf/*|*install-pandoc/*) continue ;;
  esac
  grep -q 'JOBOPS NOT CONFIGURED' "$f" || echo "MISSING: $f"
  grep -q 'config.directories.<key>' "$f" || echo "WRONG: $f"
done
```
Expected: no output (no MISSING and no WRONG lines).

- [ ] **Step 5: Verify the 8 template-consumer skills have the `## Templates` block**

Run:
```bash
for f in buildresume provenance-check createrubric assessjob assesscandidate briefing interviewprep coverletter; do
  grep -q 'Templates referenced by this skill' plugins/jobops/skills/$f/SKILL.md \
    || echo "MISSING TEMPLATES BLOCK: $f"
done
```
Expected: no output.

- [ ] **Step 6: Verify the 4 utility skills were NOT modified**

Run:
```bash
git diff --name-only plugins/jobops/skills/convert-to-pdf plugins/jobops/skills/convert-to-word plugins/jobops/skills/markdown-to-pdf plugins/jobops/skills/install-pandoc
```
Expected: empty (no files modified).

- [ ] **Step 7: Commit**

```bash
git add plugins/jobops/skills
git commit -m "Standardize config and templates preamble across jobops skills

All 26 non-utility jobops skills now carry the same Configuration
block (reads .jobops/config.json, exits cleanly if missing, uses
config.directories / cultural_profile / default_jurisdiction). The
8 template-consuming skills additionally carry a Templates block
listing the specific templates they read by name.

Utility skills (convert-to-pdf, convert-to-word, markdown-to-pdf,
install-pandoc) stay preamble-free because they do not read workspace
paths.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 7: Standardize preamble in non-setup jobops-ic skills

**Why:** Same rationale as Task 6 for the IC plugin. The handwritten `resume-summarizer`-check prerequisite block (currently in most IC skills) is removed — the plugin `dependencies` field enforces that at install time (Task 2).

**Files (9 total, excluding setup):**
```
plugins/jobops-ic/skills/copywrite/SKILL.md
plugins/jobops-ic/skills/copywriting-spec/SKILL.md
plugins/jobops-ic/skills/create-landing-page/SKILL.md
plugins/jobops-ic/skills/css-template/SKILL.md
plugins/jobops-ic/skills/defineservices/SKILL.md
plugins/jobops-ic/skills/findclient/SKILL.md
plugins/jobops-ic/skills/pitchdeck/SKILL.md
plugins/jobops-ic/skills/proposaltemplate/SKILL.md
plugins/jobops-ic/skills/ratecard/SKILL.md
```

- [ ] **Step 1: Baseline — check current prerequisite blocks**

Run:
```bash
grep -l "resume-summarizer" plugins/jobops-ic/skills/*/SKILL.md
```
Note the list of files that currently carry the handwritten check — they each need their `## Prerequisite Check` block removed.

- [ ] **Step 2: For each of the 9 files, apply two edits**

**Edit A — remove the handwritten prerequisite block.** If the file contains a `## Prerequisite Check` section that mentions `resume-summarizer`, delete the entire block (from the `## Prerequisite Check` line up to but not including the next `## ` heading).

**Edit B — replace (or insert) the `## Configuration` block with the IC preamble.** Rule is the same as Task 6 Step 2: if the heading exists, replace up to next `## `; if not, insert after frontmatter's closing `---`. The block is:

```markdown
## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup (then /jobops-ic:setup) to initialize your workspace.

Use `config.directories.contractor_root` for output paths in this skill.
Use `config.preferences.default_currency` for pricing if applicable.

```

- [ ] **Step 3: Append `## Templates` block to defineservices only**

`plugins/jobops-ic/skills/defineservices/SKILL.md` is the only IC template-consumer. Insert immediately after the `## Configuration` block:

```markdown
## Templates

For each template used by this skill, resolve the full path as:

  {config.templates.base_dir}/{config.templates.active.<template_name>}/<filename>

Templates referenced by this skill: service_definition_schema

```

- [ ] **Step 4: Verify no IC skill still carries the obsolete prerequisite text**

Run:
```bash
grep -l "resume-summarizer" plugins/jobops-ic/skills/*/SKILL.md || echo "CLEAN"
```
Expected: `CLEAN`.

- [ ] **Step 5: Verify preamble is present and identical in all 9 files**

Run:
```bash
for f in plugins/jobops-ic/skills/*/SKILL.md; do
  [[ "$f" == *setup/SKILL.md ]] && continue
  grep -q 'JOBOPS NOT CONFIGURED' "$f" || echo "MISSING: $f"
  grep -q 'config.directories.contractor_root' "$f" || echo "WRONG: $f"
done
```
Expected: no output.

- [ ] **Step 6: Verify defineservices template block**

Run:
```bash
grep -q "Templates referenced by this skill: service_definition_schema" \
  plugins/jobops-ic/skills/defineservices/SKILL.md && echo OK
```
Expected: `OK`.

- [ ] **Step 7: Commit**

```bash
git add plugins/jobops-ic/skills
git commit -m "Standardize config preamble across jobops-ic skills

All 9 non-setup jobops-ic skills drop the handwritten
resume-summarizer prerequisite check (the dependencies field in
plugin.json now enforces it at install time) and adopt the
standardized IC Configuration block that reads contractor_root and
default_currency from .jobops/config.json. defineservices gains the
standardized Templates block referencing service_definition_schema.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 8: Update output paths in application-scoped jobops skills

**Why:** The resume-development, cover-letter, assessment, and interview skills need to stop writing to flat roots like `OutputResumes/` and `Scoring_Rubrics/`, and start writing to the fixed `Applications/{app}/<subfolder>/` layout. They also need the `APPLICATION_PATH_PROTOCOL` block so the resolution logic is spelled out in-skill.

**Files (9 total):**
```
plugins/jobops/skills/buildresume/SKILL.md
plugins/jobops/skills/provenance-check/SKILL.md
plugins/jobops/skills/coverletter/SKILL.md
plugins/jobops/skills/createrubric/SKILL.md
plugins/jobops/skills/assessjob/SKILL.md
plugins/jobops/skills/assesscandidate/SKILL.md
plugins/jobops/skills/auditjobposting/SKILL.md
plugins/jobops/skills/briefing/SKILL.md
plugins/jobops/skills/interviewprep/SKILL.md
```

- [ ] **Step 1: Insert `APPLICATION_PATH_PROTOCOL` block into each file**

For each of the 9 files, insert this block immediately after the `## Templates` block (if present) or after `## Configuration` (if no templates block):

```markdown
## Application Path Resolution

This skill writes to a per-application folder. Before writing any output:

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the job-posting filename, or honor `--app=<slug>` if supplied.
2. Compose the app folder: `{config.directories.applications_root}/{app_slug}/`.
3. Resolve this skill's sub-folder by category:
   - resume-development (buildresume, provenance-check) → `resume/`
   - cover-letter (coverletter) → `cover-letter/`
   - rubric / assessment (createrubric, assessjob, assesscandidate, auditjobposting) → `assessment/`
   - briefing / interview prep (briefing, interviewprep) → `interview/`
4. If the app folder does not exist, `mkdir -p` it, then copy
   `{config.directories.job_postings}/{filename}` → `{app_slug}/job_posting.md`
   so the pinned JD cannot silently change under completed work.
5. Exact-slug collisions (same Company+Role+Date) are not auto-suffixed. If the folder
   already contains the same output type, require the user to pass `--app=<distinct-slug>`.

```

- [ ] **Step 2: Replace obsolete-root references inside each skill's body**

For each file, perform these body-text replacements (targeting prose like "save to `{config.directories.output_resumes}/...`" and path examples like `{config.directories.scoring_rubrics}/Rubric_...`):

| Old reference | New reference |
|---|---|
| `{config.directories.output_resumes}/Step1_Draft_*` | `{applications_root}/{app_slug}/resume/step1_draft.md` |
| `{config.directories.output_resumes}/Step2_Provenance_Analysis_*` | `{applications_root}/{app_slug}/resume/step2_provenance.md` |
| `{config.directories.output_resumes}/Step3_Final_Resume_*` | `{applications_root}/{app_slug}/resume/step3_final.md` |
| `{config.directories.output_resumes}/Cover_Letter_*` | `{applications_root}/{app_slug}/cover-letter/cover_letter.md` |
| `{config.directories.scoring_rubrics}/Rubric_*` | `{applications_root}/{app_slug}/assessment/rubric.md` |
| `{config.directories.scoring_rubrics}/Assessment_*` | `{applications_root}/{app_slug}/assessment/assessment.md` |
| `{config.directories.briefing_notes}/Briefing_*` | `{applications_root}/{app_slug}/interview/briefing.md` |
| `{config.directories.briefing_notes}/Interview_Prep_*` | `{applications_root}/{app_slug}/interview/interview_prep.md` |

Apply with careful per-file grep; do NOT use a blanket `sed -i` across all files at once — review each match in context to avoid false positives (e.g., documentation tables that are describing *old* paths intentionally should be reworded, not path-swapped).

- [ ] **Step 3: Remove references to `config.directories.output_resumes`, `scoring_rubrics`, `briefing_notes`**

These config keys no longer exist. After Step 2 replacements, verify:

```bash
grep -n "directories\.\(output_resumes\|scoring_rubrics\|briefing_notes\)" \
  plugins/jobops/skills/{buildresume,provenance-check,coverletter,createrubric,assessjob,assesscandidate,auditjobposting,briefing,interviewprep}/SKILL.md \
  || echo "CLEAN"
```
Expected: `CLEAN`.

If hits remain, re-read each and convert to the appropriate `applications_root/{app_slug}/<subfolder>/...` form.

- [ ] **Step 4: Verify all 9 files now have the application-path protocol block**

Run:
```bash
for f in buildresume provenance-check coverletter createrubric assessjob assesscandidate auditjobposting briefing interviewprep; do
  grep -q "Application Path Resolution" plugins/jobops/skills/$f/SKILL.md \
    || echo "MISSING: $f"
done
```
Expected: no output.

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/skills
git commit -m "Move application-scoped outputs to app-centric layout

buildresume, provenance-check, coverletter, createrubric, assessjob,
assesscandidate, auditjobposting, briefing, and interviewprep now
write to Applications/{Company}_{Role}_{Date}/<subfolder>/ instead of
the former flat OutputResumes / Scoring_Rubrics / Briefing_Notes
roots. Each carries the Application Path Resolution block spelling out
slug parsing, folder creation, and JD pinning.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 9: Update output paths in company-intelligence jobops skills

**Why:** OSINT outputs move from the legacy flat `Intelligence_Reports/` root (where filenames like `Google_Corporate_Intelligence_20260423.md` encoded both company and agent) to the per-company layout `Company_Intelligence/{Company}/<agent>.md`. The `/jobops:osint` skill also gains the refresh / append / skip prompt when the target company folder already exists.

**Files (2):**
```
plugins/jobops/skills/osint/SKILL.md
plugins/jobops/skills/layoff-intel/SKILL.md
```

Also touched: the OSINT sub-agent files under `plugins/jobops/agents/osint-*.md` reference the legacy filename convention in their system prompts. Those are updated too (7 files).

- [ ] **Step 1: Add the Company-Intelligence path block to `osint/SKILL.md`**

Insert immediately after the `## Configuration` block:

```markdown
## Company-Intelligence Path Resolution

This skill writes to a per-company folder under
`{config.directories.company_intelligence}/{Company}/`. Before dispatching
any sub-agent:

1. Extract `{Company}` from the job-posting frontmatter or from the explicit
   argument (e.g., `/jobops:osint Google`).
2. Compose the target folder:
   `{config.directories.company_intelligence}/{Company}/`.
3. If the folder already exists, prompt the user:
   - **refresh** — overwrite all per-agent files.
   - **append** — keep existing files, write a new `summary_{YYYYMMDD}.md`
     alongside them (do not modify per-agent files in place).
   - **skip** — exit without running.
4. Sub-agents write to fixed filenames:
   `corporate.md`, `legal.md`, `leadership.md`, `compensation.md`,
   `culture.md`, `market.md`.
5. Person-level reports from `osint-person` land under
   `{Company}/people/{interviewer_name}.md`.
6. The aggregated synthesis lands at `{Company}/summary.md` on refresh
   or `{Company}/summary_{YYYYMMDD}.md` on append.

```

- [ ] **Step 2: Replace legacy path references inside `osint/SKILL.md`**

In the body of `osint/SKILL.md`:
- Replace any reference to `{config.directories.intelligence_reports}` with `{config.directories.company_intelligence}/{Company}`.
- Replace filename patterns like `[CompanyName]_Corporate_Intelligence_[Date].md` with plain `corporate.md`. Apply the analogous rename for the six legal/leadership/compensation/culture/market/corporate filenames.
- Remove any instruction that tells agents to encode Company or Date in filenames — the folder path encodes the company and refresh/append handles dating.

- [ ] **Step 3: Update `layoff-intel/SKILL.md`**

layoff-intel writes a single output. Replace references to
`{config.directories.intelligence_reports}/Layoff_Intel_*` with
`{config.directories.company_intelligence}/{Company}/layoff_intel.md`
(the target is company-scoped, same as other OSINT outputs).

- [ ] **Step 4: Update the seven OSINT agent system prompts**

Files:
```
plugins/jobops/agents/osint-agent.md
plugins/jobops/agents/osint-corporate.md
plugins/jobops/agents/osint-legal.md
plugins/jobops/agents/osint-leadership.md
plugins/jobops/agents/osint-compensation.md
plugins/jobops/agents/osint-culture.md
plugins/jobops/agents/osint-market.md
plugins/jobops/agents/osint-person.md
```

In each, replace instructions to write to
`{CompanyName}_<Domain>_Intelligence_<Date>.md` under
`Intelligence_Reports/` with instructions to write to the fixed filename
(`corporate.md`, `legal.md`, etc.) under
`{company_intelligence}/{Company}/`. For `osint-person`, the destination
is `{company_intelligence}/{Company}/people/{interviewer_name}.md`. For
`osint-agent` (the synthesis commander), the destination is
`{company_intelligence}/{Company}/summary.md` (or
`summary_{YYYYMMDD}.md` on append mode, which the calling skill
communicates via argument).

- [ ] **Step 5: Verify no remaining `intelligence_reports` references**

Run:
```bash
grep -rn "directories\.intelligence_reports" plugins/jobops \
  || echo "CLEAN"
grep -rn "_Intelligence_\[Date\]" plugins/jobops || echo "CLEAN"
grep -rn "Intelligence_Reports/" plugins/jobops || echo "CLEAN"
```
Expected: `CLEAN` for all three.

- [ ] **Step 6: Verify refresh/append/skip prompt is present in osint SKILL**

Run:
```bash
grep -n "refresh.*append.*skip\|append mode" plugins/jobops/skills/osint/SKILL.md
```
Expected: at least one hit.

- [ ] **Step 7: Commit**

```bash
git add plugins/jobops/skills plugins/jobops/agents
git commit -m "Move OSINT outputs to per-company shared folders

/jobops:osint and /jobops:layoff-intel now write to
Company_Intelligence/{Company}/ with fixed per-agent filenames
(corporate.md, legal.md, leadership.md, compensation.md, culture.md,
market.md, summary.md). Sub-agents (osint-corporate .. osint-market,
osint-person, osint-agent) updated to emit the new path convention.
/jobops:osint gains the refresh / append / skip prompt when the
target company folder already contains intelligence.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 10: Update output paths in career-analysis jobops skills

**Why:** Career-level skills emit flat timestamped files into `{career_analysis}/`. Any references to legacy roots or non-timestamped filenames need to go.

**Files (4):**
```
plugins/jobops/skills/idealjob/SKILL.md
plugins/jobops/skills/change-one-thing/SKILL.md
plugins/jobops/skills/comparejobs/SKILL.md
plugins/jobops/skills/github-portfolio/SKILL.md
```

- [ ] **Step 1: Apply per-file path updates**

| Skill | Output filename |
|---|---|
| `idealjob` | `{career_analysis}/idealjob_{YYYYMMDD}.md` |
| `change-one-thing` | `{career_analysis}/change_one_thing_{YYYYMMDD}.md` |
| `comparejobs` | `{career_analysis}/comparison_{YYYYMMDD}_{slug}.md` where `{slug}` comes from the companies being compared (e.g., `google_vs_meta`) |
| `github-portfolio` | `{career_analysis}/github_portfolio_{YYYYMMDD}.md` |

In each skill's body, replace any legacy root references
(`{config.directories.output_resumes}`, `{config.directories.intelligence_reports}`,
`{config.directories.briefing_notes}`) with
`{config.directories.career_analysis}` plus the timestamped filename above.

- [ ] **Step 2: Verify no legacy-root references remain**

Run:
```bash
for f in idealjob change-one-thing comparejobs github-portfolio; do
  grep -n "directories\.\(output_resumes\|scoring_rubrics\|briefing_notes\|intelligence_reports\)" \
    plugins/jobops/skills/$f/SKILL.md \
    && echo "HITS IN: $f" || true
done
```
Expected: no `HITS IN:` lines.

- [ ] **Step 3: Verify each file references `career_analysis`**

Run:
```bash
for f in idealjob change-one-thing comparejobs github-portfolio; do
  grep -q "config.directories.career_analysis" plugins/jobops/skills/$f/SKILL.md \
    || echo "MISSING: $f"
done
```
Expected: no output.

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills
git commit -m "Move career-analysis outputs to career_analysis root

idealjob, change-one-thing, comparejobs, and github-portfolio now emit
single timestamped files under {config.directories.career_analysis}/
(idealjob_{YYYYMMDD}.md, change_one_thing_{YYYYMMDD}.md, etc.) instead
of the legacy OutputResumes / Intelligence_Reports roots.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 11: Update output paths in crisis-management jobops skills

**Why:** Crisis skills emit flat timestamped files into `{crisis_management}/`, with one exception — `workplace-documentation` maintains a single continuously-updated log file (not timestamped).

**Files (11):**
```
plugins/jobops/skills/accommodation-request/SKILL.md
plugins/jobops/skills/assess-job-offer/SKILL.md
plugins/jobops/skills/code-red/SKILL.md
plugins/jobops/skills/constructive-dismissal/SKILL.md
plugins/jobops/skills/discrimination-assessment/SKILL.md
plugins/jobops/skills/investigation-prep/SKILL.md
plugins/jobops/skills/non-compete-analysis/SKILL.md
plugins/jobops/skills/reference-shield/SKILL.md
plugins/jobops/skills/severance-review/SKILL.md
plugins/jobops/skills/unemployment-prep/SKILL.md
plugins/jobops/skills/workplace-documentation/SKILL.md
```

- [ ] **Step 1: Apply per-file path updates**

| Skill | Output filename |
|---|---|
| `accommodation-request` | `{crisis_management}/accommodation_request_{YYYYMMDD}.md` |
| `assess-job-offer` | `{crisis_management}/offer_assessment_{YYYYMMDD}_{Company}.md` |
| `code-red` | `{crisis_management}/code_red_{YYYYMMDD}.md` |
| `constructive-dismissal` | `{crisis_management}/constructive_dismissal_{YYYYMMDD}.md` |
| `discrimination-assessment` | `{crisis_management}/discrimination_assessment_{YYYYMMDD}.md` |
| `investigation-prep` | `{crisis_management}/investigation_prep_{YYYYMMDD}.md` |
| `non-compete-analysis` | `{crisis_management}/non_compete_analysis_{YYYYMMDD}.md` |
| `reference-shield` | `{crisis_management}/reference_shield_{YYYYMMDD}.md` |
| `severance-review` | `{crisis_management}/severance_review_{YYYYMMDD}.md` |
| `unemployment-prep` | `{crisis_management}/unemployment_prep_{YYYYMMDD}.md` |
| `workplace-documentation` | `{crisis_management}/workplace_documentation_log.md` (continuously-appended, no date) |

For each file, replace any legacy-root references with the above path, and
ensure `--jurisdiction=<ISO-3166-2>` is accepted as an override on the
jurisdiction-sensitive skills (all of the above except `assess-job-offer`
and `reference-shield` — those don't have jurisdictional logic).

- [ ] **Step 2: Spot-check `workplace-documentation` uses APPEND semantics**

`workplace-documentation` should write to the single log file using append
semantics (new entries added with timestamp header; existing content
preserved). Confirm the skill body describes this and does not call for
overwrite.

Run:
```bash
grep -n "append\|continuous\|log file" plugins/jobops/skills/workplace-documentation/SKILL.md
```
Expected: at least one hit describing append / continuous / log semantics.

- [ ] **Step 3: Verify no legacy-root references remain**

Run:
```bash
for f in accommodation-request assess-job-offer code-red constructive-dismissal discrimination-assessment investigation-prep non-compete-analysis reference-shield severance-review unemployment-prep workplace-documentation; do
  grep -n "directories\.\(output_resumes\|scoring_rubrics\|briefing_notes\|intelligence_reports\)" \
    plugins/jobops/skills/$f/SKILL.md \
    && echo "HITS IN: $f" || true
done
```
Expected: no `HITS IN:` lines.

- [ ] **Step 4: Verify each crisis skill references `crisis_management`**

Run:
```bash
for f in accommodation-request assess-job-offer code-red constructive-dismissal discrimination-assessment investigation-prep non-compete-analysis reference-shield severance-review unemployment-prep workplace-documentation; do
  grep -q "config.directories.crisis_management" plugins/jobops/skills/$f/SKILL.md \
    || echo "MISSING: $f"
done
```
Expected: no output.

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/skills
git commit -m "Move crisis-management outputs to crisis_management root

All 11 crisis skills now emit flat timestamped files under
{config.directories.crisis_management}/, with workplace-documentation
using append semantics against a single continuously-updated log
(workplace_documentation_log.md). Jurisdiction-sensitive skills accept
--jurisdiction=<ISO-3166-2> to override the workspace default.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 12: Update output paths in jobops-ic skills

**Why:** IC skills move from `client_prospects/...` to the new `contractor_root/<subfolder>/...` layout.

**Files (9, excluding setup):**

| Skill | Destination |
|---|---|
| `defineservices` | `{contractor_root}/services/service_definition_{YYYYMMDD}.md` (+ `.json` alongside for `--from-profile`) |
| `findclient` | `{contractor_root}/prospects/{ProspectCompany}_{YYYYMMDD}.md` |
| `proposaltemplate` | `{contractor_root}/proposals/{ClientCompany}_{YYYYMMDD}.md` |
| `pitchdeck` | `{contractor_root}/pitches/{ClientCompany}_{YYYYMMDD}.md` |
| `ratecard` | `{contractor_root}/rate-cards/rate_card_{YYYYMMDD}.md` |
| `create-landing-page` | `{contractor_root}/landing-pages/{slug}/index.html` (+ assets) |
| `copywrite` | `{contractor_root}/landing-pages/{slug}/copy.md` |
| `copywriting-spec` | `{contractor_root}/landing-pages/{slug}/spec.md` |
| `css-template` | `{contractor_root}/landing-pages/{slug}/styles.css` |

- [ ] **Step 1: Apply per-file path updates**

For each skill above, replace references to `{config.directories.client_prospects}`
and any other legacy root with the new destination from the table.

- [ ] **Step 2: Verify no references to obsolete `client_prospects` config key remain**

Run:
```bash
grep -rn "client_prospects" plugins/jobops-ic || echo "CLEAN"
```
Expected: `CLEAN`.

- [ ] **Step 3: Verify each IC skill references `contractor_root`**

Run:
```bash
for f in defineservices findclient proposaltemplate pitchdeck ratecard create-landing-page copywrite copywriting-spec css-template; do
  grep -q "config.directories.contractor_root" plugins/jobops-ic/skills/$f/SKILL.md \
    || echo "MISSING: $f"
done
```
Expected: no output.

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops-ic/skills
git commit -m "Move jobops-ic outputs to contractor_root subfolders

All 9 non-setup IC skills now write under
{config.directories.contractor_root}/<subfolder>/ — services/,
prospects/, proposals/, pitches/, rate-cards/, and landing-pages/ —
replacing the legacy client_prospects root. Landing-page skills
(create-landing-page, copywrite, copywriting-spec, css-template)
share a per-slug folder under landing-pages/{slug}/.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 13: Add `docs/ARCHITECTURE.md`

**Why:** Section 10.3 of the spec calls for a contributor-only reference documenting the config + path-resolution contract so future skill authors produce files that follow the convention.

**Files:**
- Create: `docs/ARCHITECTURE.md`

- [ ] **Step 1: Write the file**

Write the following to `docs/ARCHITECTURE.md`:

```markdown
# JobOps Architecture

> **Audience:** contributors adding or modifying JobOps skills. End users should read the plugin READMEs and skill descriptions instead.

## 1. Two plugins, one config

- `jobops` owns resume, interview, OSINT, career, and crisis skills.
- `jobops-ic` adds an independent-contractor layer on top. It declares `dependencies: {jobops: "^2.0.0"}` in its `plugin.json` so Claude Code refuses to install it without `jobops`.
- Both plugins read the same `.jobops/config.json` in the user's workspace. `jobops-ic:setup` extends that file; it does not create a separate one.

## 2. Config file

Location: `.jobops/config.json` (workspace root, gitignored by default).

Schema: see `docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md` Section 6.2.

Creation: only by `/jobops:setup`. Extended by `/jobops-ic:setup`. Never written by runtime skills.

Missing-file behavior: every runtime skill (except the two setup skills and `/jobops:migrate`) exits immediately with:

    JOBOPS NOT CONFIGURED
    Run /jobops:setup to initialize your workspace.

## 3. Plugin-root resolution

Skills that need the plugin's own templates or bundled files use `${CLAUDE_PLUGIN_ROOT}` directly inside shell commands in skill markdown. No `hooks/` directory and no `/tmp/` state file are involved.

## 4. Output layout

Three destination patterns.

**Application-centric** — per-application folders, fixed subfolder convention:

    {applications_root}/{Company}_{Role}_{YYYYMMDD}/
      ├── job_posting.md           (pinned copy of the JD)
      ├── resume/
      │   ├── step1_draft.md
      │   ├── step2_provenance.md
      │   └── step3_final.md
      ├── cover-letter/cover_letter.md
      ├── assessment/{rubric,assessment}.md
      └── interview/{briefing,interview_prep}.md

**Company-centric** — per-company OSINT, shared across applications to the same company:

    {company_intelligence}/{Company}/
      ├── {corporate,legal,leadership,compensation,culture,market}.md
      ├── people/{interviewer_name}.md
      └── summary.md                (or summary_{YYYYMMDD}.md on append)

**Flat** — one timestamped file per invocation:

    {career_analysis}/<skill>_{YYYYMMDD}[.optional_slug].md
    {crisis_management}/<skill>_{YYYYMMDD}.md
    {contractor_root}/<subfolder>/<filename>

Exception: `workplace-documentation` appends to a single continuously-updated log (`workplace_documentation_log.md`), not a timestamped file.

## 5. Skill-authoring contract

Every runtime skill:

1. Begins with YAML frontmatter carrying `description`, `disable-model-invocation: true`, and (optional) `argument-hint`.
2. Has a `## Configuration` block using either `JOBOPS_PREAMBLE` (for jobops skills) or `JOBOPS_IC_PREAMBLE` (for jobops-ic skills). See Sections 7.1 and 7.2 of the spec for the verbatim blocks.
3. If it consumes templates, has a `## Templates` block listing each template by name. Template path resolution is always `{config.templates.base_dir}/{config.templates.active.<name>}/<filename>`.
4. If it writes to an application folder, has an `## Application Path Resolution` block spelling out the four resolution steps (slug parsing, folder composition, sub-folder, JD pinning).
5. If it writes to a company folder, has a `## Company-Intelligence Path Resolution` block including the refresh / append / skip prompt for the existing-folder case.
6. Never hardcodes a directory name — always reads `config.directories.<key>`.

## 6. Setup flow invariants

- `/jobops:setup` writes the config atomically (`.tmp` then `mv`) so a crash mid-write can't leave an invalid file.
- `/jobops:setup --reconfigure` is idempotent: running it twice in a row with the same answers produces the same config and filesystem state.
- Gitignore management writes a single block marked by `# JobOps workspace`; re-running setup replaces the block in place rather than appending duplicates.
- Legacy migration is opt-in, dry-run-first, and user-editable.

## 7. Template variants

Defaults ship in `plugins/<plugin>/templates/` and are copied to `.jobops/templates/default/` by setup. Users create variants under `.jobops/templates/custom/` and toggle the active one via `config.templates.active.<name>`. `default/` is treated as read-only by convention.
```

- [ ] **Step 2: Verify the file is markdown-valid**

Run:
```bash
head -3 docs/ARCHITECTURE.md
wc -l docs/ARCHITECTURE.md
```
Expected: first line is `# JobOps Architecture`; line count around 80–100.

- [ ] **Step 3: Commit**

```bash
git add docs/ARCHITECTURE.md
git commit -m "Add docs/ARCHITECTURE.md contributor reference

Documents the config + path-resolution contract so future skill
authors produce files that follow the conventions: two plugins with
one config, CLAUDE_PLUGIN_ROOT for plugin-local resources, three
output layouts (application-centric, company-centric, flat), and the
six skill-authoring rules. Sources its authority from the approved
spec (docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md)
but is denormalized so contributors don't need to read the full spec.

Part of the plugin config redesign (spec: 2026-04-23)."
```

---

## Task 14: Structural validation

**Why:** Catch any broken JSON, broken frontmatter, or overlooked legacy reference before handing off to smoke testing. The `claude plugin validate` command is the plugin system's built-in linter.

- [ ] **Step 1: Validate both plugins and the marketplace**

Run:
```bash
claude plugin validate plugins/jobops
claude plugin validate plugins/jobops-ic
claude plugin validate .
```
Expected: each run reports valid / no errors. If any command reports an error, stop and fix before continuing.

If the `claude` CLI is not available in the execution environment, fall back to:
```bash
python3 -c "import json,sys,glob; [json.load(open(f)) or print(f,'OK') for f in glob.glob('plugins/**/*.json', recursive=True) + glob.glob('.claude-plugin/*.json')]"
```
and a frontmatter-syntax check:
```bash
for f in plugins/*/skills/*/SKILL.md; do
  head -1 "$f" | grep -q '^---$' || echo "FRONTMATTER_MISSING: $f"
done
```
Expected: no FRONTMATTER_MISSING lines.

- [ ] **Step 2: Full sweep for any remaining legacy references**

Run:
```bash
{
  grep -rn "directories\.output_resumes" plugins/
  grep -rn "directories\.scoring_rubrics" plugins/
  grep -rn "directories\.briefing_notes" plugins/
  grep -rn "directories\.intelligence_reports" plugins/
  grep -rn "directories\.client_prospects" plugins/
  grep -rn "directories\.sample_output" plugins/
  grep -rn "/tmp/.jobops" plugins/
  grep -rn "resume-summarizer" plugins/jobops-ic/
  grep -rn "copy-templates.sh" plugins/
} || echo "ALL CLEAN"
```
Expected: `ALL CLEAN`. Any hit must be resolved before continuing.

- [ ] **Step 3: Every non-utility non-setup skill has a config preamble**

Run:
```bash
fail=0
for f in plugins/jobops/skills/*/SKILL.md plugins/jobops-ic/skills/*/SKILL.md; do
  case "$f" in
    *jobops/skills/setup/*|*jobops/skills/migrate/*|*jobops-ic/skills/setup/*) continue ;;
    *convert-to-pdf/*|*convert-to-word/*|*markdown-to-pdf/*|*install-pandoc/*) continue ;;
  esac
  grep -q "JOBOPS NOT CONFIGURED" "$f" || { echo "MISSING PREAMBLE: $f"; fail=1; }
done
[ $fail -eq 0 ] && echo "OK"
```
Expected: `OK`.

- [ ] **Step 4: Every template-consumer has the Templates block**

Run:
```bash
for f in plugins/jobops/skills/buildresume plugins/jobops/skills/provenance-check plugins/jobops/skills/createrubric plugins/jobops/skills/assessjob plugins/jobops/skills/assesscandidate plugins/jobops/skills/briefing plugins/jobops/skills/interviewprep plugins/jobops/skills/coverletter plugins/jobops-ic/skills/defineservices; do
  grep -q "Templates referenced by this skill" "$f/SKILL.md" || echo "MISSING: $f"
done
```
Expected: no output.

- [ ] **Step 5: Application-scoped skills have the Application Path Resolution block**

Run:
```bash
for f in buildresume provenance-check coverletter createrubric assessjob assesscandidate auditjobposting briefing interviewprep; do
  grep -q "Application Path Resolution" plugins/jobops/skills/$f/SKILL.md \
    || echo "MISSING: $f"
done
```
Expected: no output.

- [ ] **Step 6: Commit nothing if all clean; otherwise fix and commit per-task**

No commit for validation-only. If any step produced a fix, commit it separately with a message beginning `fix: validation — `.

---

## Task 15: Smoke tests against the repo itself

**Why:** Structural validation catches syntax issues, but the best integration test is running the new setup end-to-end on a real workspace. The JobOps repo itself contains example job postings and is the most realistic target.

These tests are **manual** in nature — they involve running Claude Code commands. The implementing engineer/subagent should report the observed behavior rather than auto-proceed if any test fails.

- [ ] **Step 1: Fresh-install walkthrough**

1. In a scratch directory (not this repo), create an empty workspace.
2. Install both plugins via `--plugin-dir` pointed at this repo.
3. Run `/jobops:setup` and answer the interview with all defaults.
4. Observe: six root directories created, `.jobops/config.json` populated, templates installed under `.jobops/templates/default/`, gitignore block written.
5. Verify `.jobops/config.json` matches the Section 6.2 schema (version, six directories, both preferences, templates.active with four entries, migration block).

- [ ] **Step 2: Dependency enforcement**

1. In a scratch directory, attempt to install `jobops-ic` via `--plugin-dir` WITHOUT `jobops` installed.
2. Observe: install is rejected with a dependency error.

- [ ] **Step 3: Setup reconfigure idempotence**

1. In the Step 1 workspace, run `/jobops:setup --reconfigure` and accept all existing values.
2. Observe: config file content unchanged (diff is empty); no duplicate gitignore block; no errors.

- [ ] **Step 4: IC setup extends rather than overwrites**

1. In the Step 1 workspace, run `/jobops-ic:setup` with defaults.
2. Observe: `.jobops/config.json` gains `directories.contractor_root`, `preferences.default_currency`, `templates.active.service_definition_schema`, nothing removed. Contractor/ directory tree created.

- [ ] **Step 5: Spot-check a skill invocation**

1. In the Step 1 workspace, drop a sample job posting into `Job_Postings/` named `Google_SeniorPM_20260423.md`.
2. Run `/jobops:createrubric Job_Postings/Google_SeniorPM_20260423.md`.
3. Observe: `Applications/Google_SeniorPM_20260423/assessment/rubric.md` created; `Applications/Google_SeniorPM_20260423/job_posting.md` is a copy of the source JD.

- [ ] **Step 6: Migration dry-run against the JobOps repo**

1. In THIS repo (which has accumulated legacy test files), run `/jobops:migrate --dry-run`.
2. Observe: preview lists planned moves; nothing is actually moved; config is not updated; unresolved files are surfaced.
3. Do NOT execute the migration against this repo unless the user explicitly asks.

- [ ] **Step 7: Record results**

Compile a short report of which steps passed / failed, with the observed error messages for any failures. The report is the output of this task.

---

## Task 16: Finalize and prepare for review

- [ ] **Step 1: Review the full commit series**

Run:
```bash
git log --oneline main..HEAD
```
Expected: roughly 11–13 commits matching the task commits above.

- [ ] **Step 2: Double-check nothing unexpected is still pending**

Run:
```bash
git status
```
Expected: clean working tree.

- [ ] **Step 3: Summarize file change count**

Run:
```bash
git diff --stat main..HEAD | tail -1
```
Report the line. Expected scale: ~45 files changed, roughly 2000+ lines added / 1000+ lines removed (the bulk is the setup rewrite + 40 skill preamble/path edits + the new migrate skill).

- [ ] **Step 4: Hand off**

Stop. Report the summary from Step 3 plus the smoke-test report from Task 15 Step 7. Do not open a PR, push to remote, or make further changes without explicit user instruction.

---

## Post-plan notes

**Order matters** only where noted: Task 1 (deletions) can be safely interleaved with Task 2 (manifest updates). Tasks 3–5 (setup rewrites + new migrate skill) are independent of each other and of Tasks 6–12. Tasks 6 and 7 (preamble standardization) must complete before Tasks 8–12 (path updates inside skill bodies) because the path-update tasks assume the `## Configuration` block already exists at the expected position. Task 13 (ARCHITECTURE.md) is independent. Tasks 14–16 are final.

**If executed by subagents in parallel**, the safest split is:
- Agent A: Tasks 1, 2, 3, 4, 5, 13.
- Agent B: Tasks 6, 8.
- Agent C: Tasks 7, 12.
- Agent D: Tasks 9, 10, 11.
- Merge → Agent E: Tasks 14, 15, 16.

**If executed sequentially by a single subagent**, proceed in the listed order 1 → 16.
