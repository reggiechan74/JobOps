---
description: Bump JobOps version across plugin/marketplace JSON, README, and CHANGELOG
argument-hint: <patch|minor|major> [- short description of change]
allowed-tools: Read, Edit, Bash, Grep, Glob
---

# /version-bump

Bump the JobOps version consistently across every file that carries it, and add a new CHANGELOG entry categorized by release type.

## Inputs

`$ARGUMENTS` should contain:
1. **Release type** (required, first token): `patch` | `minor` | `major`
   - `patch` → bug fix (X.Y.Z → X.Y.Z+1) → CHANGELOG section: **Fixed**
   - `minor` → new feature, backwards-compatible (X.Y.Z → X.Y+1.0) → CHANGELOG section: **Added** / **Changed**
   - `major` → breaking change (X.Y.Z → X+1.0.0) → CHANGELOG section: **Changed — BREAKING**
2. **Description** (optional, everything after the type): short summary of the change. If omitted, derive one from the current git diff and recent conversation context.

If `$ARGUMENTS` is empty or the release type is missing/invalid, ask the user via `AskUserQuestion` which type to apply before proceeding.

## Files that carry the version

These eight files must all be bumped to the same version in a single pass. They split into Claude Code manifests, Codex manifests (the `.codex-plugin/` + `.agents/` pair added for Codex compatibility), and the package/README. If any of these paths no longer exists or new version-bearing files have been added, surface the discrepancy and ask before proceeding.

| File | Field |
|------|-------|
| `package.json` | `"version": "X.Y.Z"` |
| `.claude-plugin/marketplace.json` | `metadata.version` (Claude Code marketplace) |
| `.agents/plugins/marketplace.json` | `metadata.version` (Codex marketplace) |
| `plugins/jobops/.claude-plugin/plugin.json` | `"version": "X.Y.Z"` (Claude Code manifest) |
| `plugins/jobops/.codex-plugin/plugin.json` | `"version": "X.Y.Z"` (Codex manifest) |
| `plugins/jobops-ic/.claude-plugin/plugin.json` | `"version": "X.Y.Z"` (Claude Code manifest) |
| `plugins/jobops-ic/.codex-plugin/plugin.json` | `"version": "X.Y.Z"` (Codex manifest) |
| `README.md` | `**Version X.Y.Z**` line near the top |

`npm test` validates that the Codex marketplace and both Codex manifests match — run it after bumping to catch any Codex file you missed.

Also update:
- `CHANGELOG.md` — insert a new `## [X.Y.Z] - YYYY-MM-DD` entry directly above the current top entry. Use the **current date** from the system reminder (or `date +%Y-%m-%d` via Bash if the reminder is unavailable).

## Process

1. **Read** `package.json` to determine the current version. Treat this as the source of truth — every other file should already match it; if any don't, surface the drift before bumping.
2. **Parse `$ARGUMENTS`**. If release type is missing or invalid, ask the user. If description is missing, summarize the pending git diff (`git status` + `git diff`) and the recent conversation in one or two sentences.
3. **Compute** the next version per semver based on release type.
4. **Edit each file** in the table above to the new version. Use exact-string `Edit` calls; do not use `replace_all` on a bare `2.2.0` because that string can appear in other contexts. Match enough surrounding context to make each replacement unique.
5. **Edit `CHANGELOG.md`** by inserting a new entry above the current top version heading. Format:

   ```markdown
   ## [X.Y.Z] - YYYY-MM-DD

   ### {Fixed | Added | Changed | Changed — BREAKING}

   - **{file path or skill/agent name}** — {one-paragraph description of the change, why it exists, and what failure mode it addresses}.
   - {additional bullets for related sub-changes, one per logical unit}

   ## [previous version] - previous date
   ```

   For `patch`, default heading is `### Fixed`. For `minor`, use `### Added` for new capabilities and `### Changed` for modifications — include both subsections if both apply. For `major`, use `### Changed — BREAKING`. Match the existing CHANGELOG voice: lead each bullet with the file path or component name in bold, then explain the why and the failure mode addressed — not just the what.
6. **Verify** that no version-bearing file was missed: run `npm test` (it fails if the Codex marketplace or either Codex manifest still has the old version) and `grep -rn "<old version>"` across the eight files. Both must come back clean before reporting success.
7. **Report back** to the user: list the eight files bumped, the new version, and quote the CHANGELOG entry you wrote so they can confirm it reflects the change.

## Guardrails

- **Do not commit, tag, or push.** Version bump only edits files. The user runs `git commit` and `git tag -a vX.Y.Z` themselves.
- **Do not bump `jobops-ic` independently** of `jobops` — the marketplace ships them as a coordinated pair. If a release only touches one plugin, still bump both for marketplace alignment, and note in the report that one plugin had no functional changes so the user can decide whether to revert.
- **Do not invent CHANGELOG content** beyond what the diff and conversation support. If you cannot describe what changed and why, ask the user for the description rather than guessing.
- **Verify uniqueness** of each `Edit` `old_string`. The string `"version": "2.2.0",` appears in four plugin.json files (Claude + Codex manifests for both plugins) with identical surrounding lines — include the preceding `"name"` line to disambiguate, and note that the jobops Claude and Codex manifests share the same `"name": "jobops"` line, so `Read` each file first and edit one at a time rather than relying on context alone.
- **Do not skip files.** If any of the eight version-bearing files cannot be edited (missing, malformed JSON, etc.), stop and surface the problem — do not produce a partial bump.

## Example invocations

- `/version-bump patch - added Step 6a sub-agent sentence review to step4-cover-letter agent`
- `/version-bump minor - new /jobops:proposalreview skill`
- `/version-bump major - removed deprecated candidate_profile.json artifact`
- `/version-bump patch` (description will be inferred from git diff + conversation)
- `/version-bump` (will prompt for release type)
