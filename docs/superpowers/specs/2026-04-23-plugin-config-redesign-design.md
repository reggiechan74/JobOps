# JobOps Plugin Config & Output-Layout Redesign

**Date:** 2026-04-23
**Status:** Approved (brainstorm phase)
**Branch:** `feature/plugin-split`
**Supersedes:** Portions of `docs/plans/2026-03-07-plugin-split-design.md` related to configuration, path resolution, and output layout. The March design's marketplace structure, skill/agent inventory, HAM-Z methodology, and template concepts remain intact.
**Next step:** Implementation plan via `superpowers:writing-plans`.

---

## 1. Purpose

Two user needs drive this redesign:

1. **User-configurable target output folders with per-skill sub-folder conventions.** Users installing the plugin on their own workspace must be able to set where every top-level output folder lives, and the plugin must organize outputs application-centrically rather than artifact-type-centrically.
2. **A plugin-external user config file that the plugin populates through an interactive interview.** Config is user-owned, hand-editable, and not embedded in plugin internals.

A third driver is correcting design decisions in the March plan that were superseded by Claude Code platform changes between March and April 2026.

## 2. Platform changes since the March plan

Verified via Claude Code documentation research (April 2026). Relevant findings:

- `${CLAUDE_PLUGIN_ROOT}` now expands directly in skill markdown (not only in hooks/MCP configs). The `/tmp/.jobops-plugin-root` workaround from the March plan is obsolete.
- `${CLAUDE_PLUGIN_DATA}` provides persistent per-plugin storage that survives updates. Not used by this design (our state is per-workspace, not per-install).
- `plugin.json` gained a `userConfig` field for declarative per-plugin user configuration, prompted at enable time and stored in `settings.json`. **This design does not use `userConfig`** — user preference is for a dedicated JobOps-owned config file, which `/jobops:setup` populates via a guided interview.
- `plugin.json` gained a `dependencies` field for plugin-level requirements. The current schema accepts a bare array of plugin names (e.g., `["jobops"]`); version floors are **not** expressible in this field as of April 2026. `jobops-ic` uses `dependencies: ["jobops"]` instead of a handwritten per-skill prerequisite check.
- `disable-model-invocation: true` remains the correct frontmatter for manual-invoke-only skills.
- `claude plugin validate` and `--plugin-dir` remain the validation / local-testing commands.

## 3. Design decisions

| Decision | Choice |
|---|---|
| Output organization | Application-centric — one folder per job application |
| Top-level roots | Multiple, each independently configurable |
| Application identity | Derived from job-posting filename (`{Company}_{Role}_{YYYYMMDD}`); `--app=<slug>` override supported |
| Internal app-folder layout | Fully grouped, fixed by plugin convention (not user-configurable) |
| OSINT scope | Company-level — shared across applications to the same company |
| OSINT on existing data | Prompt: refresh / append / skip |
| Config file location | `.jobops/config.json` in workspace |
| Config scope | Single shared config; `jobops-ic` extends it rather than creating its own |
| Config gitignore | Gitignored by default (user can opt-in to commit) |
| Plugin-root path resolution | `${CLAUDE_PLUGIN_ROOT}` inline in skill markdown (no hooks, no `/tmp/` file) |
| Template swap | Kept — default/custom map in config supports jurisdictional variants |
| File archival | None — skills overwrite; git is the undo mechanism |
| Dependency enforcement | `dependencies` field in `jobops-ic/plugin.json` |
| Migration of legacy layouts | Optional step in `/jobops:setup` with dry-run preview and user-editable mapping; also available as standalone `/jobops:migrate` |
| Gitignore management | Setup writes a JobOps block to `.gitignore` after showing a preview |

## 4. Repository output layout

### 4.1 Top-level roots

Each path is independently configurable via `config.directories` in `.jobops/config.json`.

```
<workspace>/
├── ResumeSourceFolder/          input:  master career data (user-maintained)
├── Job_Postings/                input:  target JDs
├── Applications/                output: per-application artifacts (app-centric)
├── Company_Intelligence/        output: OSINT at company level (shared)
├── Career_Analysis/             output: idealjob, change-one-thing, comparejobs
├── Crisis_Management/           output: severance, non-compete, discrimination, workplace docs, etc.
└── Contractor/                  output: jobops-ic artifacts (added if jobops-ic installed)
```

### 4.2 Internal application folder (fixed convention)

```
Applications/{Company}_{Role}_{YYYYMMDD}/
├── job_posting.md                        pinned copy of the JD for provenance
├── resume/
│   ├── step1_draft.md
│   ├── step2_provenance.md
│   └── step3_final.md
├── cover-letter/
│   └── cover_letter.md
├── assessment/
│   ├── rubric.md
│   └── assessment.md
└── interview/
    ├── briefing.md
    └── interview_prep.md
```

On first skill run against a given application, the skill copies `{config.directories.job_postings}/{filename}.md` to `{app_folder}/job_posting.md` so the pinned JD cannot silently change under completed work.

Exact-slug collisions (reapplying on the same date) are rare. The user supplies `--app=<distinct-slug>` to create a separate folder in that case; skills do not silently auto-suffix.

### 4.3 Internal company-intelligence folder (fixed convention)

```
Company_Intelligence/{Company}/
├── corporate.md
├── culture.md
├── leadership.md
├── legal.md
├── market.md
├── compensation.md
├── people/
│   └── {interviewer_name}.md             from osint-person (one per interviewer)
└── summary.md                             aggregated osint-agent synthesis
```

When `/jobops:osint` is run against a company that already has an intelligence folder, the skill prompts: **refresh** (overwrite all files), **append** (add a new timestamped summary alongside existing data), or **skip**.

### 4.4 Career-level, crisis, and contractor folders

Flat — one file per skill invocation, timestamped in the filename. Examples:

```
Career_Analysis/
├── idealjob_20260423.md
├── change_one_thing_20260418.md
└── comparison_20260420_google_vs_meta.md

Crisis_Management/
├── severance_review_20260423.md
├── non_compete_analysis_20260415.md
└── workplace_documentation_log.md         the single continuously-updated log

Contractor/
├── services/
├── prospects/
├── proposals/
├── pitches/
├── rate-cards/
└── landing-pages/
```

Within `Contractor/`, one subfolder per IC skill's output type — fixed by plugin convention.

## 5. Plugin manifests

### 5.1 `plugins/jobops/.claude-plugin/plugin.json`

```json
{
  "name": "jobops",
  "version": "2.0.0",
  "description": "Intelligence-driven job application system - resume development, interview prep, OSINT, career strategy, and crisis management using HAM-Z methodology",
  "author": {"name": "Reggie Chan"},
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": ["resume", "job-search", "interview-prep", "osint", "career", "assessment", "crisis-management"]
}
```

No `userConfig` block. All user state lives in `.jobops/config.json` (Section 6).

### 5.2 `plugins/jobops-ic/.claude-plugin/plugin.json`

```json
{
  "name": "jobops-ic",
  "version": "2.0.0",
  "description": "Independent contractor toolkit - services, client prospecting, pitch decks, proposals, rate cards, landing pages. Requires jobops.",
  "author": {"name": "Reggie Chan"},
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": ["independent-contractor", "consulting", "proposals", "rate-card", "pitch-deck", "landing-page"],
  "dependencies": ["jobops"]
}
```

The `dependencies` field makes Claude Code refuse to install `jobops-ic` without `jobops`. This replaces the handwritten "is resume-summarizer available" prerequisite check that every `jobops-ic` skill carried in the March plan.

### 5.3 `.claude-plugin/marketplace.json`

Unchanged from the March plan.

## 6. Workspace config file

### 6.1 Location and scope

- **Path:** `.jobops/config.json` in the workspace root.
- **Ownership:** JobOps-owned. Hand-editable at any time.
- **Scope:** Single shared file; `/jobops-ic:setup` extends it rather than creating a second file.
- **Version control:** `.jobops/` is added to `.gitignore` by setup. Users who want to commit it (e.g., team-shared workspace) can opt in by editing `.gitignore`.
- **Creation:** By `/jobops:setup` via the interview flow (Section 8). Skills never create it.
- **Missing-file behavior:** Any skill that depends on the file stops with `JOBOPS NOT CONFIGURED — Run /jobops:setup`.

### 6.2 Schema

The example below shows the file after both `/jobops:setup` and `/jobops-ic:setup` have run. If only `/jobops:setup` has run, `directories.contractor_root`, `preferences.default_currency`, and `templates.active.service_definition_schema` are absent and added later by `/jobops-ic:setup`.

```json
{
  "version": "2.0",
  "directories": {
    "resume_source": "./ResumeSourceFolder",
    "job_postings": "./Job_Postings",
    "applications_root": "./Applications",
    "company_intelligence": "./Company_Intelligence",
    "career_analysis": "./Career_Analysis",
    "crisis_management": "./Crisis_Management",
    "contractor_root": "./Contractor"
  },
  "preferences": {
    "cultural_profile": "canadian",
    "default_jurisdiction": "CA-ON",
    "default_currency": "CAD"
  },
  "templates": {
    "base_dir": "./.jobops/templates",
    "active": {
      "assessment_rubric_framework": "default",
      "evidence_verification_framework": "default",
      "assessment_report_structure": "default",
      "candidate_profile_schema": "default",
      "service_definition_schema": "default"
    }
  },
  "migration": {
    "completed": false,
    "completed_at": null,
    "files_moved": 0
  }
}
```

### 6.3 Field semantics

| Field | Purpose |
|---|---|
| `version` | Schema version, bumped independently of plugin version |
| `directories.resume_source` | Master HAM-Z career data |
| `directories.job_postings` | User-maintained folder of target JDs |
| `directories.applications_root` | Per-application output tree (Section 4.2) |
| `directories.company_intelligence` | OSINT output tree (Section 4.3) |
| `directories.career_analysis` | Career-level analysis outputs |
| `directories.crisis_management` | Crisis-skill outputs |
| `directories.contractor_root` | `jobops-ic` outputs. Present only after `/jobops-ic:setup`. |
| `preferences.cultural_profile` | `"canadian"` \| `"american"` — resume voice |
| `preferences.default_jurisdiction` | ISO 3166-2 code (e.g., `CA-ON`). Crisis skills accept `--jurisdiction=X` to override per-invocation. |
| `preferences.default_currency` | ISO 4217 code. Used by `jobops-ic` rate cards and proposals. Present only after `/jobops-ic:setup`. |
| `templates.base_dir` | Where `default/` and `custom/` template folders live |
| `templates.active.<name>` | Per-template toggle: `"default"`, `"custom"`, or any named variant |
| `migration.completed` | Has `/jobops:setup` migrated legacy outputs? Prevents re-prompt. |
| `migration.completed_at` | ISO 8601 timestamp |
| `migration.files_moved` | Audit count |

### 6.4 Template directory structure

```
.jobops/
├── config.json
└── templates/
    ├── default/                           (copied from plugin on setup, read-only by convention)
    │   ├── assessment_rubric_framework.md
    │   ├── evidence_verification_framework.md
    │   ├── assessment_report_structure.md
    │   ├── candidate_profile_schema.json
    │   └── service_definition_schema.json (added on /jobops-ic:setup)
    └── custom/                            (user-created variants, toggled via templates.active)
```

To activate a custom variant, the user copies the default file into `custom/`, modifies it, and edits `.jobops/config.json` to change the `templates.active` entry for that template from `"default"` to `"custom"`.

## 7. Skill contract

### 7.1 Standard preamble (jobops)

```markdown
## Configuration

Read `.jobops/config.json`. If missing, stop with:

  JOBOPS NOT CONFIGURED
  Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if the skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if the skill has jurisdiction-sensitive logic.
```

### 7.2 Standard preamble (jobops-ic)

```markdown
## Configuration

Read `.jobops/config.json`. If missing, stop with:

  JOBOPS NOT CONFIGURED
  Run /jobops:setup (then /jobops-ic:setup) to initialize your workspace.

Use `config.directories.contractor_root` for output paths in this skill.
Use `config.preferences.default_currency` for pricing if applicable.
```

No handwritten prerequisite check — the `dependencies` field handles plugin-level enforcement.

### 7.3 Template-consuming skills

The following skills also carry a `## Templates` block:

- `buildresume`, `provenance-check`, `createrubric`, `assessjob`, `assesscandidate`, `briefing`, `interviewprep`, `coverletter`
- `defineservices` (via jobops-ic, uses `service_definition_schema.json`)

```markdown
## Templates

For each template needed, resolve path from:
  {config.templates.base_dir}/{config.templates.active.<name>}/<filename>

Templates referenced by this skill: <explicit list>
```

### 7.4 Application-path resolution protocol

Any skill that writes to an application folder performs these four steps, spelled out explicitly in the skill's markdown:

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the job-posting filename, or honor `--app=<slug>` override.
2. Compose the app folder: `{config.directories.applications_root}/{app_slug}/`
3. Resolve the skill's sub-folder by category:
   - resume-development → `resume/`
   - cover-letter → `cover-letter/`
   - rubric / assessment → `assessment/`
   - briefing / interview prep → `interview/`
4. If the app folder does not exist, create it and copy `{config.directories.job_postings}/{filename}` → `{app_folder}/job_posting.md`.

### 7.5 Company-intelligence resolution

1. Extract `{Company}` from the job-posting frontmatter or explicit argument.
2. Target: `{config.directories.company_intelligence}/{Company}/`.
3. If the folder exists, prompt refresh / append / skip.
4. Person-level reports from `osint-person` land in `{Company}/people/{interviewer_name}.md`.

### 7.6 Flat-output skills

Career-level, crisis, and non-application-specific contractor skills emit a single timestamped file under their respective root:

- `{config.directories.career_analysis}/<skill>_{date}.md`
- `{config.directories.crisis_management}/<skill>_{date}.md`
- `{config.directories.contractor_root}/<subfolder>/<descriptive_name>.md`

## 8. `/jobops:setup` flow

Signature: `/jobops:setup [--reconfigure] [--skip-migration] [--skip-history]`

1. **Welcome & existing-config check.** If `.jobops/config.json` exists, present the user: (a) Reconfigure all, (b) Reconfigure specific sections, (c) Exit. Otherwise proceed.
2. **Directory interview.** Walk through each path in `config.directories` with default, purpose, and inline edit. Paths are normalized and validated (writable parent).
3. **Create directories.** `mkdir -p` each confirmed path. Report newly-created vs already-existed.
4. **Preferences interview.** `cultural_profile` (enum), `default_jurisdiction` (string, ISO 3166-2). `default_currency` is not asked here — it is owned by `/jobops-ic:setup` (Section 8.1).
5. **Template installation.** `cp ${CLAUDE_PLUGIN_ROOT}/templates/* {templates.base_dir}/default/`. Create `{templates.base_dir}/custom/`.
6. **Write `.jobops/config.json`.** Emit the full schema with confirmed values. `migration.completed = false`.
7. **Gitignore update.** Show a preview of the JobOps block to be appended to `.gitignore`. User accepts, edits inline, or skips. Default block:
   ```
   # JobOps workspace
   .jobops/
   Applications/
   Company_Intelligence/
   Career_Analysis/
   Crisis_Management/
   Contractor/
   ```
8. **Career history import (optional).** Only offered if `ResumeSourceFolder/` is empty or near-empty. Accepts PDF / DOCX / TXT / MD inputs; extracts and transforms to HAM-Z; populates `Experience/`, `CareerHighlights/`, `Technology/`, `Preferences/`; creates `Vision.md` stub. Flags entries lacking metrics or hard skills.
9. **Migration of legacy files (optional).** Runs only if `migration.completed == false` and any legacy folders (`OutputResumes/`, `Briefing_Notes/`, `Scoring_Rubrics/`, `Intelligence_Reports/`) contain files. Flow:
   - **Dry-run pass.** Walk each legacy folder; parse `{Company}_{Role}_{YYYYMMDD}` from each filename.
   - **Preview.** Show planned moves grouped by target app folder; list unresolved files separately.
   - **User-editable mapping.** User can reassign specific files to different application folders before executing.
   - **Execute.** `git mv` for tracked files (preserves history); plain `mv` for untracked.
   - **Update config.** Set `migration.completed = true`, `completed_at`, `files_moved`.
   - Unresolved files stay put; user handles manually.
10. **Summary & next steps.** Recap what was configured; recommend first skills to try.

### 8.1 `/jobops-ic:setup` flow

1. **Prerequisite check.** `.jobops/config.json` must exist. If not, stop with a pointer to `/jobops:setup`.
2. **Interview.** `contractor_root` (default `./Contractor`), `default_currency` (enum).
3. **Create directory.** `mkdir -p {contractor_root}`.
4. **Template installation.** Copy `service_definition_schema.json` from the `jobops-ic` plugin to `{templates.base_dir}/default/`.
5. **Extend config.** Add `directories.contractor_root`, `preferences.default_currency`, and `templates.active.service_definition_schema` to the existing `.jobops/config.json`.
6. **Gitignore update.** Append `Contractor/` if not already present.
7. **Summary.**

## 9. `/jobops:migrate` skill (new, standalone)

Extracts the migration logic (step 9 of `/jobops:setup`) into a first-class skill for users who already completed setup and accumulated legacy files afterward, or who want to re-run migration on new batches.

- Same dry-run → preview → user-editable mapping → execute flow as setup step 9.
- Updates `migration.files_moved` on each run (cumulative).
- No-op if all legacy folders are empty.

## 10. Scope of changes to existing `feature/plugin-split` work

### 10.1 Deletions

| File | Reason |
|---|---|
| `plugins/jobops/hooks/hooks.json` | `${CLAUDE_PLUGIN_ROOT}` expansion in skill markdown removes the need for the `/tmp/` hack |
| `plugins/jobops-ic/hooks/hooks.json` | Same |
| `plugins/jobops/scripts/copy-templates.sh` | Setup skill runs `cp ${CLAUDE_PLUGIN_ROOT}/templates/*` inline |
| `plugins/jobops-ic/scripts/copy-templates.sh` | Same |
| `plugins/jobops/hooks/`, `plugins/jobops/scripts/` (directories) | Empty after file deletions |
| `plugins/jobops-ic/hooks/`, `plugins/jobops-ic/scripts/` (directories) | Same |

Both plugin directories simplify from six subdirectories to four: `.claude-plugin/`, `templates/`, `skills/`, `agents/`.

### 10.2 Rewrites

- `plugins/jobops/skills/setup/SKILL.md` — full rewrite of the 10-step flow (Section 8). ~400 lines.
- `plugins/jobops-ic/skills/setup/SKILL.md` — rewrite of the light extend-config flow. ~80 lines.
- `plugins/jobops-ic/.claude-plugin/plugin.json` — add `dependencies` field.
- All 30 `jobops` skills (except `setup`) — replace preamble with Section 7 pattern; update all output path references to the new app-centric / flat-output layouts.
- All 10 `jobops-ic` skills (except `setup`) — drop handwritten prerequisite block (covered by `dependencies`); update all output paths to `contractor_root` subfolders.

Total skill-file edits: **38**. Each is shallow (preamble + path references); volume is the primary cost.

### 10.3 Additions

| File | Purpose |
|---|---|
| `plugins/jobops/skills/migrate/SKILL.md` | New standalone `/jobops:migrate` (Section 9) |
| `docs/ARCHITECTURE.md` | Contributor-only reference for the config + path-resolution contract |

### 10.4 Unchanged from March plan

- Marketplace layout (`.claude-plugin/marketplace.json`)
- Skill inventory (30 + 10) and agent inventory (17 + 1)
- Template bundling inside the plugin
- `disable-model-invocation: true` on every skill
- `CLAUDE.md` contributor-only (already committed)
- `README.md` marketplace-focused (already committed; may need a refresh of the setup section)
- `package.json` v2.0.0 (already committed)

## 11. Test strategy

1. **Plugin structural validation.** `claude plugin validate plugins/jobops`, `claude plugin validate plugins/jobops-ic`, `claude plugin validate .` (marketplace).
2. **Dependency enforcement.** Attempt to install `jobops-ic` via `--plugin-dir` without `jobops` present; confirm rejection.
3. **Setup walk-through.** Run `/jobops:setup` on the current JobOps repo itself — the repo's accumulated legacy files are the most realistic migration test.
4. **Smoke tests.** After setup: `/jobops:buildresume` (exercises config read, template resolution, app-path creation, JD pinning); `/jobops:osint` (exercises company-intelligence path + prompt-on-existing); `/jobops:severance-review` (exercises crisis-path with jurisdiction preference).
5. **Reconfiguration.** Re-run `/jobops:setup --reconfigure` and confirm idempotence.
6. **Migration standalone.** Run `/jobops:migrate` on a subset of legacy files after adding new ones.
7. **`jobops-ic` install.** `/jobops-ic:setup`, then `/jobops-ic:defineservices` and `/jobops-ic:ratecard` to confirm shared config extension.

## 12. Risks

- **Migration regex robustness.** Legacy filenames (e.g., `Step3_Final_Resume_Google_SeniorPM_20260423.md`) may require multiple parse patterns. Mitigation: dry-run preview with user-editable mapping.
- **Shell path expansion in skill execution context.** The `cp ${CLAUDE_PLUGIN_ROOT}/templates/*` pattern must work reliably inside Claude's skill runner. Mitigation: early smoke test during implementation; fall back to an explicit `find`-based copy if the glob fails.
- **OSINT append mode semantics.** "Append" hasn't been implemented before. Decision: append writes a new timestamped `summary_{YYYYMMDD}.md` alongside existing files; does not modify existing per-agent files in place. Refresh is the overwrite mode.
- **Config drift.** If the user hand-edits `.jobops/config.json` and introduces a schema-invalid value, skills will fail at read time. Mitigation: `/jobops:setup --reconfigure` validates and re-writes the file; skills produce clear error messages identifying the offending field.

## 13. Explicit non-goals

- No changes to agents, HAM-Z methodology, rubric content, evidence protocol, or scoring logic.
- No new skills beyond `/jobops:migrate`.
- No removal of existing skills.
- No use of `userConfig` in `plugin.json`.
- No file-archival machinery (`.archive/` directory).
- No `/jobops:reconfigure` alias — the `--reconfigure` flag on `/jobops:setup` covers the need.

## 14. Open questions deferred to the implementation plan

- Exact regex set for the migration parser.
- Error-message copy for each failure mode.
- Format of the user-editable migration mapping (inline edit of a generated TOML/JSON? numbered reassignment prompts?).
- Whether `docs/ARCHITECTURE.md` lives in the plugin-split branch or a follow-up.

These are implementation details that the writing-plans step will resolve.
