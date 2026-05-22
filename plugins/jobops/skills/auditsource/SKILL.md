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
<!-- TASK-MARKER: Step 2 inserted by Task 5 -->
<!-- TASK-MARKER: Step 3 inserted by Task 6 -->
<!-- TASK-MARKER: Step 4 inserted by Task 7 -->
<!-- TASK-MARKER: Step 5 inserted by Task 8 -->
<!-- TASK-MARKER: MUST NOT inserted by Task 9 -->
