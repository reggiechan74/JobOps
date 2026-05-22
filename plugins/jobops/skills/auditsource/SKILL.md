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

<!-- TASK-MARKER: Step 1 inserted by Task 4 -->
<!-- TASK-MARKER: Step 2 inserted by Task 5 -->
<!-- TASK-MARKER: Step 3 inserted by Task 6 -->
<!-- TASK-MARKER: Step 4 inserted by Task 7 -->
<!-- TASK-MARKER: Step 5 inserted by Task 8 -->
<!-- TASK-MARKER: MUST NOT inserted by Task 9 -->
