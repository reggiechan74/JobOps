# jobops-dash

Terminal dashboard for tracking JobOps applications and launching `/jobops` and
`/jobops-ic` skills. Tabs: **Apps** (job applications), **Companies** (OSINT),
**Career**, **Crisis**, and **Contractor** (shown when jobops-ic is configured).

## Run

From your JobOps workspace (the directory containing `.jobops/`):

```bash
npm run dash               # via the repo
# or build a binary:
go build -o jobops-dash ./tools/dashboard && ./jobops-dash
```

Use `--workspace <path>` to point at a workspace explicitly.

## Keys

`↑↓`/`jk` move · `←→`/`tab` switch tabs · `↵` open skill palette · `c` copy
next-step command · `C` toggle claude/codex · `s` cycle lifecycle status (Apps) ·
`r` rescan · `q` quit.

Pressing `↵` and selecting a skill spawns an interactive `claude` (or `codex`)
session with the command pre-filled; on exit the dashboard rescans. This uses
your normal interactive session (no headless API cost).

## State the dashboard owns

- `Applications/{slug}/.tracker.json` — per-app lifecycle status.
- `.jobops/dashboard.json` — UI prefs (selected agent).

Neither is written by JobOps skills.
