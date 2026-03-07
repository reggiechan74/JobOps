#!/bin/bash
set -euo pipefail

PLUGIN_ROOT_FILE="/tmp/.jobops-ic-plugin-root"
if [ ! -f "$PLUGIN_ROOT_FILE" ]; then
  echo "ERROR: Plugin root not found. Restart Claude Code to initialize." >&2
  exit 1
fi

PLUGIN_ROOT="$(cat "$PLUGIN_ROOT_FILE" | tr -d '[:space:]')"
SOURCE_DIR="$PLUGIN_ROOT/templates"
TARGET_DIR="${1:-.jobops/templates/default}"

if [ ! -d "$SOURCE_DIR" ]; then
  echo "ERROR: Template source not found at $SOURCE_DIR" >&2
  exit 1
fi

mkdir -p "$TARGET_DIR"
cp "$SOURCE_DIR"/* "$TARGET_DIR/"
COUNT=$(ls -1 "$TARGET_DIR" | wc -l)
echo "Copied $COUNT templates to $TARGET_DIR"
