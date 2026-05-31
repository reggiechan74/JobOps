package launch

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrefsRoundTripAndDefault(t *testing.T) {
	root := t.TempDir()
	_ = os.MkdirAll(filepath.Join(root, ".jobops"), 0o755)

	// Missing file -> default agent "claude".
	p := LoadPrefs(root)
	if p.Agent != "claude" {
		t.Errorf("default agent = %q, want claude", p.Agent)
	}

	p.Agent = "codex"
	if err := SavePrefs(root, p); err != nil {
		t.Fatal(err)
	}
	if LoadPrefs(root).Agent != "codex" {
		t.Errorf("saved agent not read back")
	}
}

func TestClipboardCommandDetection(t *testing.T) {
	// clipboardCommand picks the first available tool name from candidates.
	got := clipboardCommand([]string{"definitely-not-a-real-tool-xyz", "echo"})
	if got != "echo" {
		t.Errorf("clipboardCommand = %q, want echo (first found on PATH)", got)
	}
	if clipboardCommand([]string{"definitely-not-a-real-tool-xyz"}) != "" {
		t.Errorf("expected empty when no candidate exists")
	}
}
