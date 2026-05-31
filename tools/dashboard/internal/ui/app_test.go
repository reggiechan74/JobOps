package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func newTestModel() Model {
	m := New("/tmp/ws", "claude", []TabSource{
		{Name: "Apps", Scanner: staticScanner{recs: sampleRecords()}, Lifecycle: true},
	})
	m.width, m.height = 100, 30
	return m
}

// staticScanner is a Scanner that returns canned records (no filesystem).
type staticScanner struct{ recs []model.Record }

func (s staticScanner) Scan() ([]model.Record, error) { return s.recs, nil }
func (s staticScanner) Skills() []model.SkillSpec {
	return []model.SkillSpec{{Name: "coverletter", Label: "Cover letter", Arg: model.ArgJD}}
}

func TestCursorMovesDown(t *testing.T) {
	m := newTestModel()
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	if updated.(Model).tabs[0].cursor != 1 {
		t.Errorf("cursor = %d, want 1", updated.(Model).tabs[0].cursor)
	}
}

func TestEnterOpensPalette(t *testing.T) {
	m := newTestModel()
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if updated.(Model).mode != modePalette {
		t.Errorf("mode = %v, want modePalette", updated.(Model).mode)
	}
}

func TestStatusKeyCyclesLifecycle(t *testing.T) {
	m := newTestModel()
	start := m.tabs[0].records[0].Lifecycle
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}})
	got := updated.(Model).tabs[0].records[0].Lifecycle
	if got == start {
		t.Errorf("status key should advance lifecycle from %q", start)
	}
}

func TestViewRendersWithoutPanic(t *testing.T) {
	m := newTestModel()
	if m.View() == "" {
		t.Errorf("View() returned empty string")
	}
}

func TestSpawnMissingAgentFallback(t *testing.T) {
	m := newTestModel()
	m.agent = "definitely-not-a-real-agent-xyz"
	m.mode = modePalette
	m.paletteAt = 0
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected a command from the spawn fallback")
	}
	msg := cmd()
	nm, ok := msg.(noticeMsg)
	if !ok {
		t.Fatalf("expected noticeMsg, got %T", msg)
	}
	final, _ := updated.(Model).Update(nm)
	fm := final.(Model)
	if fm.mode != modeNotice {
		t.Errorf("mode = %v, want modeNotice", fm.mode)
	}
	if !strings.Contains(fm.notice, "not found") {
		t.Errorf("notice = %q, want it to mention 'not found'", fm.notice)
	}
}

type shrinkingScanner struct{ recs []model.Record }

func (s *shrinkingScanner) Scan() ([]model.Record, error) { return s.recs, nil }
func (s *shrinkingScanner) Skills() []model.SkillSpec      { return nil }

func TestRescanClampsCursor(t *testing.T) {
	three := []model.Record{
		{Title: "A", Slug: "A_20260101"},
		{Title: "B", Slug: "B_20260101"},
		{Title: "C", Slug: "C_20260101"},
	}
	sc := &shrinkingScanner{recs: three}
	m := New("/tmp/ws", "claude", []TabSource{{Name: "Apps", Scanner: sc, Lifecycle: true}})
	m.width, m.height = 100, 30
	m.tabs[0].cursor = 2

	sc.recs = three[:2]

	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'r'}})
	if cmd == nil {
		t.Fatal("expected r to return a rescan command")
	}
	updated, _ = updated.(Model).Update(cmd())
	fm := updated.(Model)
	if fm.tabs[0].cursor != 1 {
		t.Errorf("cursor = %d, want clamped to 1 after shrink-rescan", fm.tabs[0].cursor)
	}
	_ = fm.View()
}

func TestTabSwitchWraps(t *testing.T) {
	m := New("/tmp/ws", "claude", []TabSource{
		{Name: "Apps", Scanner: staticScanner{recs: sampleRecords()}, Lifecycle: true},
		{Name: "Companies", Scanner: staticScanner{recs: nil}},
	})
	m.width, m.height = 100, 30
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRight})
	if updated.(Model).active != 1 {
		t.Errorf("active = %d, want 1 after right", updated.(Model).active)
	}
	updated, _ = updated.(Model).Update(tea.KeyMsg{Type: tea.KeyRight})
	if updated.(Model).active != 0 {
		t.Errorf("active = %d, want 0 after wrap", updated.(Model).active)
	}
}

func TestPerTabCursorIndependent(t *testing.T) {
	m := New("/tmp/ws", "claude", []TabSource{
		{Name: "Apps", Scanner: staticScanner{recs: sampleRecords()}, Lifecycle: true},
		{Name: "Companies", Scanner: staticScanner{recs: sampleRecords()}},
	})
	m.width, m.height = 100, 30
	// Move cursor down on tab 0.
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	// Switch to tab 1.
	updated, _ = updated.(Model).Update(tea.KeyMsg{Type: tea.KeyRight})
	fm := updated.(Model)
	if fm.tabs[0].cursor != 1 {
		t.Errorf("tab 0 cursor = %d, want 1 (preserved)", fm.tabs[0].cursor)
	}
	if fm.tabs[1].cursor != 0 {
		t.Errorf("tab 1 cursor = %d, want 0 (independent)", fm.tabs[1].cursor)
	}
}

func TestStatusNoOpOnNonLifecycleTab(t *testing.T) {
	m := New("/tmp/ws", "claude", []TabSource{
		{Name: "Companies", Scanner: staticScanner{recs: sampleRecords()}}, // Lifecycle: false
	})
	m.width, m.height = 100, 30
	start := m.tabs[0].records[0].Lifecycle
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}})
	if updated.(Model).tabs[0].records[0].Lifecycle != start {
		t.Errorf("s on a non-lifecycle tab should be a no-op; lifecycle changed from %q", start)
	}
}

func TestTabSwitchLeftWraps(t *testing.T) {
	m := New("/tmp/ws", "claude", []TabSource{
		{Name: "Apps", Scanner: staticScanner{recs: sampleRecords()}, Lifecycle: true},
		{Name: "Companies", Scanner: staticScanner{recs: nil}},
	})
	m.width, m.height = 100, 30
	// From tab 0, left should wrap to the last tab.
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyLeft})
	if updated.(Model).active != 1 {
		t.Errorf("active = %d, want 1 after left-wrap from 0", updated.(Model).active)
	}
}
