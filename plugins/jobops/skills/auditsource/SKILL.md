---
description: Audit ResumeSourceFolder for structural completeness and layout conformance; interactively fill gaps and write edits back to source markdown
disable-model-invocation: true
argument-hint: "[--deep] [--migrate-layout] [--section=<path>] [--dry-run]"
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.resume_source` for the source folder root.
The canonical layout contract lives at `{plugin_root}/skills/auditsource/source_layout.md`. **Read it before running checks** — the contract evolves; do not rely on memory.

## Flags

- `--deep` — also run semantic checks (achievements without metrics, vague responsibilities, cross-file skill inconsistencies). Slower, more questions. Default off.
- `--migrate-layout` — interactively reorganize an existing folder into the canonical layout. Run once at adoption time. Refuses to run if layout already matches.
- `--section=<path>` — audit only the named subfolder (e.g. `WorkHistory/`). Default: whole tree.
- `--dry-run` — produce the gap report without prompting for fills or writing edits.

## Your Task

Audit the source folder so downstream skills can rely on it as authoritative — no inference, no hallucination, no synthesized fields. Every gap is either filled with user-authored content or explicitly marked as known-incomplete.

The remaining sections (Step 1 through Step 5, and "What this skill MUST NOT do") follow.

---

## Step 1: Layout check

Read `{plugin_root}/skills/auditsource/source_layout.md` for the canonical structure. Then walk `config.directories.resume_source` and compare.

For each REQUIRED file in the canonical layout, check existence. For each file in the source folder, check whether its path matches the canonical layout.

Produce a layout diff:

| Status | Canonical path | Found at |
|---|---|---|
| OK | Identity/Name.md | Identity/Name.md |
| MISSING | Identity/CurrentRole.md | (not found) |
| MISLOCATED | Technology/TechStack.md | Skills.md (root) |

Decision tree:

- All canonical-required files present AND no mislocated files, WITHOUT `--migrate-layout` → proceed to Step 2.
- All canonical-required files present AND no mislocated files, WITH `--migrate-layout` → print "Layout already matches canonical contract; no reorganization needed." and stop.
- Missing files OR mislocated files present, WITHOUT `--migrate-layout` → print the diff and stop. Tell the user: "Layout does not match canonical contract. Re-run with `--migrate-layout` to interactively reorganize, or fix manually and re-run."
- Missing files OR mislocated files present, WITH `--migrate-layout` → enter Step 1a (migration).

### Step 1a: Migration (only with --migrate-layout)

Refuse to run if `git status` shows uncommitted changes in the source folder. Tell the user to commit or stash first.

For each MISLOCATED file:
1. Display the current path, the proposed target path, and the first 5 lines of the file as preview.
2. Use `AskUserQuestion` with options: `Move to proposed path` / `Specify different target` / `Leave in place (mark as orphan)` / `Skip this file`.
3. On confirm: use `Bash` to `git mv` the file (preserves history).
4. Append a line to `<resume_source>/migration_log.md` recording (timestamp, old path, new path, decision).

For each MISSING required file:
1. Note in the migration log: file required but not present; user must create it (audit-source will block on this in Step 2 if not addressed).

After all moves are confirmed and logged, proceed to Step 2.

---
## Step 2: Structural audit (deterministic — no LLM judgment)

Every check in this step is a literal pattern match against file content. If a check requires LLM interpretation, it belongs in Step 5 (semantic audit, behind `--deep`).

For each file in the canonical layout, apply the checks below. Each match (or miss) produces a gap record: `{ file, line, category, description, severity, suggestion }`.

### Identity/CurrentRole.md

| Check | Severity | Description if failed |
|---|---|---|
| Contains a level-1 or level-2 heading | blocking | No job title heading found |
| Contains a line matching `(?i)company:` or a heading naming a company | blocking | No company identified |
| Contains a line matching `(?i)start[:\s]+\d{4}-\d{2}` OR an explicit `Status: Unemployed since YYYY-MM` line | blocking | No start_date in YYYY-MM format |

### WorkHistory/NN_company_role.md (for each file)

| Check | Severity | Description if failed |
|---|---|---|
| Filename matches `^\d{2}_.+\.md$` | advisory | Filename does not follow NN_ ordering prefix |
| Contains a heading with company name and title | blocking | No role title heading |
| Contains a line `(?i)start[:\s]+\d{4}-\d{2}` | blocking | No start_date |
| Contains a line `(?i)end[:\s]+(\d{4}-\d{2}|present)` | blocking | No end_date (or 'Present') |
| Contains a `## Responsibilities` or `## Achievements` heading | blocking | No content sections |
| For each line containing `%`, `$`, or `\d+x`, check if any of the 5 lines preceding contain a non-numeric context word (>3 chars) | advisory | Numeric claim without nearby context |

### Technology/TechStack.md

| Check | Severity | Description if failed |
|---|---|---|
| Contains at least one bullet, table row, or sub-heading | blocking | File has no enumerated skills |
| (cross-file, --deep only) Every skill word that appears in WorkHistory file headings or tech-listing lines also appears here | advisory | Skill in WorkHistory missing from TechStack |

### Technology/Certifications.md

For each bullet or section in the file:

| Check | Severity | Description if failed |
|---|---|---|
| Line contains a cert name AND issuer (or `None` as file content) | blocking | Cert lacks issuer |
| Line contains a date in `YYYY-MM` format OR explicit `In-Progress` | blocking | Cert lacks date |
| Line contains status `Active`/`Expired`/`In-Progress` | blocking | Cert lacks status |

### Preferences/Vision.md and AntiVision.md

| Check | Severity | Description if failed |
|---|---|---|
| File exists | blocking | Required preference file missing |
| File is non-empty OR contains the single word `None` | blocking | Preference file empty without explicit `None` |

### Cross-file timeline checks

Sort all WorkHistory files by `start_date`. For each adjacent pair (file_a ends, file_b starts):

| Check | Severity | Description if failed |
|---|---|---|
| `file_b.start_date >= file_a.end_date` (date arithmetic, deterministic) | advisory | Timeline conflict: roles overlap |
| `(file_b.start_date - file_a.end_date).days < 30` OR a gap is documented in `Preferences/` or `Identity/` | advisory | Unexplained gap > 30 days |

Also check `Identity/CurrentRole.md` vs the most recent WorkHistory entry: if CurrentRole says "currently" or "present" but the most recent WorkHistory file has an explicit end_date in the past, flag as advisory: "CurrentRole says 'currently' but most recent role has end_date < today".

### Output

Build `gaps[]` as a list of `{ file, line, category, description, severity, suggestion }`. Hand to Step 3.

---
<!-- TASK-MARKER: Step 3 inserted by Task 6 -->
<!-- TASK-MARKER: Step 4 inserted by Task 7 -->
<!-- TASK-MARKER: Step 5 inserted by Task 8 -->
<!-- TASK-MARKER: MUST NOT inserted by Task 9 -->
