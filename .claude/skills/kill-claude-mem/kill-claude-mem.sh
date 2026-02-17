#!/bin/bash
# Kill stale claude-mem worker and MCP server processes
# Safe to run at any time — claude-mem's hooks will re-spawn fresh instances
#
# Usage:
#   ./kill-claude-mem.sh          # Kill all, report summary
#   ./kill-claude-mem.sh --quiet  # Kill all, no output (for hooks)

QUIET=false
[[ "$1" == "--quiet" ]] && QUIET=true

# Count before killing
WORKERS=$(pgrep -fc "worker-service.cjs" 2>/dev/null || echo 0)
MCP_SERVERS=$(pgrep -fc "mcp-server.cjs" 2>/dev/null || echo 0)
TOTAL=$((WORKERS + MCP_SERVERS))

if [ "$TOTAL" -eq 0 ]; then
  $QUIET || echo "No stale claude-mem processes found."
  exit 0
fi

# Capture memory before
MEM_BEFORE=$(free -m | awk '/^Mem:/{print $3}')

# Kill all instances — plugin hooks will re-spawn what's needed
pkill -f "worker-service.cjs" 2>/dev/null || true
pkill -f "mcp-server.cjs" 2>/dev/null || true

# Brief wait for process cleanup
sleep 1

# Capture memory after
MEM_AFTER=$(free -m | awk '/^Mem:/{print $3}')
MEM_FREED=$((MEM_BEFORE - MEM_AFTER))

if $QUIET; then
  exit 0
fi

echo "claude-mem cleanup complete"
echo "  Killed: ${WORKERS} workers, ${MCP_SERVERS} MCP servers"
echo "  Memory freed: ~${MEM_FREED} MB"
echo "  Status: Plugin hooks will re-spawn fresh instances on next interaction"
