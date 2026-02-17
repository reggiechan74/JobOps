---
name: kill-claude-mem
description: Kill stale claude-mem worker and MCP server processes that leak memory across sessions. Use when Codespace feels slow or memory is low.
allowed-tools: Bash
---

# Kill Stale claude-mem Processes

**Purpose:** claude-mem spawns worker-service and mcp-server processes per session that aren't always cleaned up when sessions end. Over multiple sessions, these accumulate and consume gigabytes of RAM. This command kills all instances; the plugin's hooks will re-spawn fresh ones on the next interaction.

**Upstream bug:** <https://github.com/thedotmack/claude-mem/issues/1089>

## Portable Installation

This skill folder is self-contained. To add it to another repo:

1. Copy the folder: `cp -r .claude/skills/kill-claude-mem /path/to/other/repo/.claude/skills/`
2. Run the installer: `cd /path/to/other/repo && bash .claude/skills/kill-claude-mem/install-hook.sh`

The installer adds a `SessionStart` hook to that project's `.claude/settings.json` so stale processes are cleaned automatically on every session startup.

## Files

| File | Purpose |
|------|---------|
| `SKILL.md` | This file — slash command definition |
| `kill-claude-mem.sh` | Cleanup script (used by hook and slash command) |
| `install-hook.sh` | One-time installer for the SessionStart hook |

## Execution Instructions

When this skill is invoked, run the cleanup script:

```bash
bash "$CLAUDE_PROJECT_DIR/.claude/skills/kill-claude-mem/kill-claude-mem.sh"
```

Then verify the system state:

```bash
free -h | head -2
```

Report the results to the user in this format:

```
claude-mem cleanup
  Killed: X workers, Y MCP servers
  Memory freed: ~Z MB
  RAM: X.XGi used / 15Gi total
  Status: Fresh instances will spawn on next tool use
```
