# jobops-dash

Terminal dashboard for tracking JobOps applications and launching `/jobops`
skills. Phase 1 ships the Apps tab.

## Run

From your JobOps workspace (the directory containing `.jobops/`):

```bash
npm run dash               # via the repo
# or build a binary:
go build -o jobops-dash ./tools/dashboard && ./jobops-dash
```

Use `--workspace <path>` to point at a workspace explicitly.

## Keys

`↑↓`/`jk` move · `↵` open skill palette · `c` copy next-step command ·
`C` toggle claude/codex · `s` cycle lifecycle status · `r` rescan · `q` quit.

Pressing `↵` and selecting a skill spawns an interactive `claude` (or `codex`)
session with the command pre-filled; on exit the dashboard rescans. This uses
your normal interactive session (no headless API cost).

## State the dashboard owns

- `Applications/{slug}/.tracker.json` — per-app lifecycle status.
- `.jobops/dashboard.json` — UI prefs (selected agent).

Neither is written by JobOps skills.
