# JobOps Architecture

> **Audience:** contributors adding or modifying JobOps skills. End users should read the plugin READMEs and skill descriptions instead.

## 1. Two plugins, one config

- `jobops` owns resume, interview, OSINT, career, and crisis skills.
- `jobops-ic` adds an independent-contractor layer on top. It declares `dependencies: ["jobops"]` in its `plugin.json` so Claude Code refuses to install it without `jobops`. (The current plugin schema accepts a bare array of plugin names; version floors are not expressible here. Enforce version floors at runtime if ever needed.)
- Both plugins read the same `.jobops/config.json` in the user's workspace. `jobops-ic:setup` extends that file; it does not create a separate one.

## Platform compatibility

JobOps is distributed to both Claude Code and Codex from the same canonical plugin directories:

- Claude Code reads `.claude-plugin/marketplace.json` and each plugin's `.claude-plugin/plugin.json`.
- Codex reads `.agents/plugins/marketplace.json` and each plugin's `.codex-plugin/plugin.json`.
- Both platforms read the same `skills/` directories. Every skill frontmatter must include `name`, `description`, and `disable-model-invocation: true`.
- Keep skill bodies platform-neutral where practical. When a workflow needs platform-specific orchestration, describe the smallest mapping inline instead of creating a duplicate skill tree.
- Claude Code `Task tool` instructions map conceptually to Codex subagent spawning. Do not rewrite an entire skill solely to rename the orchestration primitive.
- `${CLAUDE_PLUGIN_ROOT}` remains a Claude-specific variable. Existing uses are tolerated for this compatibility pass; subsequent template/style changes should introduce a small cross-platform path-resolution pattern rather than duplicating skills.

Run `npm test` before committing any manifest, marketplace, or skill-frontmatter change.

## 2. Config file

Location: `.jobops/config.json` (workspace root, gitignored by default).

Schema: see `docs/superpowers/specs/2026-04-23-plugin-config-redesign-design.md` Section 6.2.

`config.candidate` holds the candidate's header contact fields (`name`, `credentials`, `location`, `phone`, `email`, `linkedin`, `github`), collected in `/jobops:setup` Step 4b. The `buildresume` and `coverletter` flows source the document header from this block so the resume and cover letter render identical contact lines. `phone` is a distinct field joined with ` | ` (or `•` on resumes); empty fields are omitted with their separator and never concatenated onto an adjacent field.

`config.preferences.cover_letter_mode` (`retrospective` | `forward`, default `retrospective`) selects how `/jobops:coverletter` writes the letter. A config written before this key existed has no `cover_letter_mode`; consumers treat the absence as `retrospective`. `/jobops:coverletter --mode=` overrides it per-invocation. `forward` mode runs a mandatory skill-level intake interview and dispatches `step4-cover-letter-forward` instead of `step4-cover-letter`; both agents share voice, provenance, and sub-agent-review rules.

`config.directories.tailored_cv` (default `./Tailored_CV`) is the base resume library for revise-first `/jobops:buildresume`: one user-curated `.md` base per role family, each carrying `output_type: resume_base` and a free-text `role_family` label in front matter (the taxonomy is entirely user-defined — no built-in category list exists). A config written before this key existed has no `tailored_cv`; `buildresume` treats the absent key (or an empty directory) as "no library" and falls back to the from-scratch pipeline with a setup hint — it never self-heals the config. The library is written only via buildresume's explicit promotion offer; the one-time `role_family` stamping of untagged files edits front matter only.

Creation: only by `/jobops:setup`. Extended by `/jobops-ic:setup`. Runtime skills never
write it, with one narrow exception: `/jobops:dashboard` adds the
`directories.application_tracker` key (preserving all other keys) if a pre-existing
workspace lacks it — a one-time self-heal.

Missing-file behavior: every runtime skill (except the two setup skills and `/jobops:migrate`) exits immediately with:

    JOBOPS NOT CONFIGURED
    Run /jobops:setup to initialize your workspace.

## 3. Plugin-root resolution

Skills that need the plugin's own templates or bundled files use `${CLAUDE_PLUGIN_ROOT}` directly inside shell commands in skill markdown. No hook directory and no `/tmp/` state file are involved.

## 4. Output layout

Four destination patterns.

**Application-centric** — per-application folders, fixed subfolder convention:

    {applications_root}/{Company}_{Role}_{YYYYMMDD}/
      ├── job_posting.md           (pinned copy of the JD)
      ├── resume/
      │   ├── step1_draft.md
      │   ├── step2_provenance.md
      │   └── step3_final.md
      ├── cover-letter/cover_letter.md
      ├── assessment/{domain_research,rubric,assessment}.md
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

Exception: `idealjob` writes four flat siblings per invocation — `idealjob_{YYYYMMDD}.md` (summary, `output_type: ideal_job_summary`) plus `idealjob_{YYYYMMDD}_{anchor,stretch,pivot}.md` (`output_type: ideal_job_archetype`, with assessjob-compatible `overall_score`/`normalized_score` keys so `/comparejobs` can ingest them by direct path).

**Base library** — user-curated revise-mode inputs, written only via buildresume's explicit promotion offer:

    {tailored_cv}/<base>.md            (output_type: resume_base, role_family: <user label>)

**Application tracker** — a single YAML file (`config.directories.application_tracker`,
default `{applications_root}/tracker.yaml`) maintained exclusively by `/jobops:dashboard`.
It is reconciled from the filesystem on each run: filesystem presence drives the
`artifacts` flags and `next_action`; the human-status zone (`stage`, dates, `contact`,
`outcome`, `notes`) is preserved across reconciles. No other skill reads or writes it.

**Output-type contract** — every Markdown output carries a stable `output_type` key in its
YAML front matter. This is the **filename-independent** detection key: downstream tooling
(notably the dashboard reconcile) trusts it over the filename, falling back to canonical
paths and globs only as a safety net for drifted/legacy files. Filenames inside each
sub-folder are fixed (resolved by step 6 of `## Application Path Resolution`); only the
sub-folder is resolved dynamically. Canonical `output_type` values:

| `output_type` | Producing skill | Canonical file |
|---|---|---|
| `job_posting` | path-resolution JD pin | `job_posting.md` |
| `rubric` | createrubric / assessjob | `assessment/rubric.md` |
| `assessment` | assessjob | `assessment/assessment.md` |
| `resume_step1` | buildresume (step 1) | `resume/step1_draft.md` |
| `resume_provenance` | buildresume (step 2) / provenance-check | `resume/step2_provenance.md` |
| `resume_final` | buildresume (step 3) | `resume/step3_final.md` |
| `resume_manifest` | buildresume (revise-mode step 1) | `resume/step1_manifest.md` |
| `resume_base` | buildresume promotion offer | `{tailored_cv}/<base>.md` (user-defined filename) |
| `cover_letter` | coverletter | `cover-letter/cover_letter.md` |
| `osint_corporate` … `osint_market` | osint | `{company_intelligence}/{Company}/<area>.md` |
| `osint_summary` | osint | `{company_intelligence}/{Company}/summary.md` |
| `briefing` | briefing | `interview/briefing.md` (`_partN` if split) |
| `interview_prep` | interviewprep | `interview/interview_prep.md` (`_partN` if split) |

A PDF/TeX/DOCX derivative shares the **exact basename** of its source `.md` and lives in the
**same** sub-folder — never a separate `latex/` folder. `/jobops:normalize-apps` brings
drifted existing application folders into line with this contract (slug rename, file moves,
filename normalization, `latex/` fold, `output_type` backfill).

## 5. Skill-authoring contract

Every runtime skill:

1. Begins with YAML frontmatter that includes `name`, `description`, `disable-model-invocation: true`, and optional `argument-hint` where applicable.
2. Has a `## Configuration` block using either `JOBOPS_PREAMBLE` (for jobops skills) or `JOBOPS_IC_PREAMBLE` (for jobops-ic skills). See Sections 7.1 and 7.2 of the spec for the verbatim blocks.
3. If it consumes templates, has a `## Templates` block listing each template by name. Template path resolution is always `{config.templates.base_dir}/{config.templates.active.<name>}/<filename>`.
4. If it writes to an application folder, has an `## Application Path Resolution` block spelling out the resolution steps: canonical-slug parsing (leading PascalCase company, trailing 8-digit date; leading date/time prefixes rejected at creation), folder composition, sub-folder, JD pinning (with `output_type: job_posting` front matter on the pinned copy), collision handling, and the fixed output filename(s). Every Markdown output it writes carries its `output_type` (see Section 4).
5. If it writes to a company folder, has a `## Company-Intelligence Path Resolution` block including the refresh / append / skip prompt for the existing-folder case.
6. Never hardcodes a directory name — always reads `config.directories.<key>`.

## 6. Setup flow invariants

- `/jobops:setup` writes the config atomically (`.tmp` then `mv`) so a crash mid-write can't leave an invalid file.
- `/jobops:setup --reconfigure` is idempotent: running it twice in a row with the same answers produces the same config and filesystem state.
- Gitignore management writes a single block marked by `# JobOps workspace`; re-running setup replaces the block in place rather than appending duplicates.
- Legacy migration is opt-in, dry-run-first, and user-editable.

## 7. Template variants

Defaults ship in `plugins/<plugin>/templates/` and are copied to `.jobops/templates/default/` by setup. Users create variants under `.jobops/templates/custom/` and toggle the active one via `config.templates.active.<name>`. `default/` is treated as read-only by convention.
