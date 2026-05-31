package scan

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

// Tracker is the dashboard-owned per-application lifecycle state, stored at
// Applications/{slug}/.tracker.json.
type Tracker struct {
	Lifecycle model.Lifecycle `json:"lifecycle"`
	AppliedOn string          `json:"applied_on,omitempty"`
	Notes     string          `json:"notes,omitempty"`
	UpdatedAt string          `json:"updated_at"`
}

const trackerFile = ".tracker.json"

// ReadTracker reads {appDir}/.tracker.json. A missing file yields a default
// Tracker (Interested) with no error; an unreadable/corrupt file also defaults,
// so a bad tracker never breaks a scan.
func ReadTracker(appDir string) (Tracker, error) {
	data, err := os.ReadFile(filepath.Join(appDir, trackerFile))
	if err != nil {
		// Missing or unreadable tracker -> default (Interested); never fatal.
		return Tracker{Lifecycle: model.DefaultLifecycle()}, nil
	}
	var t Tracker
	if json.Unmarshal(data, &t) != nil || !t.Lifecycle.Valid() {
		return Tracker{Lifecycle: model.DefaultLifecycle()}, nil
	}
	return t, nil
}

// WriteTracker writes {appDir}/.tracker.json atomically (temp file + rename),
// creating appDir if needed and stamping UpdatedAt.
func WriteTracker(appDir string, t Tracker) error {
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return err
	}
	t.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	tmp := filepath.Join(appDir, trackerFile+".tmp")
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, filepath.Join(appDir, trackerFile))
}
