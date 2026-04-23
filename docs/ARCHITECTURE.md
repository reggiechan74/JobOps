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
