#!/bin/bash
# Install the kill-claude-mem SessionStart hook into the current project
#
# Usage:
#   cd /path/to/your/repo
#   bash .claude/skills/kill-claude-mem/install-hook.sh
#
# What it does:
#   1. Makes kill-claude-mem.sh executable
#   2. Adds a SessionStart hook to .claude/settings.json
#   3. Hook runs on startup|resume to kill stale claude-mem processes
#
# Safe to run multiple times — skips if hook is already installed.
# Requires: jq

set -euo pipefail

# Resolve paths
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
SKILL_DIR_REL=".claude/skills/kill-claude-mem"
SETTINGS=".claude/settings.json"

# Must be run from project root (where .claude/ lives)
if [ ! -d ".claude" ]; then
  echo "Error: Run this from your project root (the directory containing .claude/)"
  echo "  cd /path/to/your/repo && bash ${SKILL_DIR_REL}/install-hook.sh"
  exit 1
fi

# Check jq is available
if ! command -v jq &>/dev/null; then
  echo "Error: jq is required but not installed."
  echo "  Install: sudo apt install jq  (Linux) or brew install jq  (macOS)"
  exit 1
fi

# Make cleanup script executable
chmod +x "${SKILL_DIR_REL}/kill-claude-mem.sh"

# Define the hook entry
HOOK_COMMAND="bash \"\$CLAUDE_PROJECT_DIR/${SKILL_DIR_REL}/kill-claude-mem.sh\" --quiet"

# Create settings.json if it doesn't exist
if [ ! -f "$SETTINGS" ]; then
  echo '{}' > "$SETTINGS"
  echo "Created ${SETTINGS}"
fi

# Check if hook is already installed
if jq -e '.hooks.SessionStart[]?.hooks[]? | select(.command | contains("kill-claude-mem"))' "$SETTINGS" &>/dev/null; then
  echo "Hook already installed in ${SETTINGS} — nothing to do."
  exit 0
fi

# Add the SessionStart hook using jq
# Handle three cases: no hooks key, hooks but no SessionStart, or existing SessionStart array
jq --arg cmd "$HOOK_COMMAND" '
  .hooks //= {} |
  .hooks.SessionStart //= [] |
  .hooks.SessionStart += [{
    "matcher": "startup|resume",
    "hooks": [{
      "type": "command",
      "command": $cmd,
      "timeout": 10,
      "statusMessage": "Cleaning up stale claude-mem processes"
    }]
  }]
' "$SETTINGS" > "${SETTINGS}.tmp" && mv "${SETTINGS}.tmp" "$SETTINGS"

echo "Installed kill-claude-mem hook"
echo "  Settings: ${SETTINGS}"
echo "  Trigger:  SessionStart (startup|resume)"
echo "  Script:   ${SKILL_DIR_REL}/kill-claude-mem.sh"
echo ""
echo "Takes effect on next Claude Code session."
