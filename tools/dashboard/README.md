# jobops-dash

Terminal dashboard for tracking JobOps applications and launching `/jobops` and
`/jobops-ic` skills. Tabs: **Apps** (job applications), **Companies** (OSINT),
**Career**, **Crisis**, and **Contractor** (shown when jobops-ic is configured).

## Run

The dashboard is a self-contained Go module in `tools/dashboard/` (its own
`go.mod`). Requires Go 1.22+ and a real terminal. Run it from your JobOps
workspace — the directory that contains `.jobops/`, created by `/jobops:setup`.

Build a binary once (from inside the module), then run it:

```bash
cd tools/dashboard && go build -o jobops-dash .   # -> tools/dashboard/jobops-dash
# put it on your PATH, or run it from your workspace:
cd /path/to/your/workspace && /abs/path/to/JobOps/tools/dashboard/jobops-dash
```

Or run it without building, pointing at a workspace explicitly:

```bash
cd tools/dashboard && go run . --workspace /path/to/your/workspace
```

From the repo root, `npm run dash` is a shortcut (it `cd`s into the module);
pass a workspace with `npm run dash -- --workspace /path/to/your/workspace`.

If launched without a configured `.jobops/` workspace, it prints setup guidance
instead of starting the TUI.

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
