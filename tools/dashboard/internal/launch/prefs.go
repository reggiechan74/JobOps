package launch

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
)

// Prefs is dashboard-owned UI state, stored at .jobops/dashboard.json. It is
// never written by JobOps skills.
type Prefs struct {
	Agent string `json:"agent"` // "claude" or "codex"
}

const prefsFile = "dashboard.json"

// LoadPrefs reads .jobops/dashboard.json from root, defaulting Agent to
// "claude" when the file is missing or unreadable.
func LoadPrefs(root string) Prefs {
	def := Prefs{Agent: "claude"}
	data, err := os.ReadFile(filepath.Join(root, ".jobops", prefsFile))
	if err != nil {
		return def
	}
	var p Prefs
	if json.Unmarshal(data, &p) != nil || (p.Agent != "claude" && p.Agent != "codex") {
		return def
	}
	return p
}

// SavePrefs writes .jobops/dashboard.json atomically.
func SavePrefs(root string, p Prefs) error {
	dir := filepath.Join(root, ".jobops")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	tmp := filepath.Join(dir, prefsFile+".tmp")
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, filepath.Join(dir, prefsFile))
}

// clipboardCommand returns the first candidate command name found on PATH, or
// "" if none is available.
func clipboardCommand(candidates []string) string {
	for _, c := range candidates {
		if _, err := exec.LookPath(c); err == nil {
			return c
		}
	}
	return ""
}

// Copy writes text to the system clipboard using the first available tool.
// Returns false if no clipboard tool is available (caller should then show the
// command in a modal for manual copying).
func Copy(text string) bool {
	cmd := clipboardCommand([]string{"pbcopy", "wl-copy", "xclip", "xsel", "clip.exe"})
	if cmd == "" {
		return false
	}
	args := []string{}
	if cmd == "xclip" {
		args = []string{"-selection", "clipboard"}
	} else if cmd == "xsel" {
		args = []string{"--clipboard", "--input"}
	}
	c := exec.Command(cmd, args...)
	stdin, err := c.StdinPipe()
	if err != nil {
		return false
	}
	if err := c.Start(); err != nil {
		return false
	}
	_, _ = stdin.Write([]byte(text))
	_ = stdin.Close()
	return c.Wait() == nil
}
