package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func newTestModel() Model {
	m := New("/tmp/ws", "claude", staticScanner{recs: sampleRecords()})
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
	if updated.(Model).cursor != 1 {
		t.Errorf("cursor = %d, want 1", updated.(Model).cursor)
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
	m.cursor = 0
	start := m.records[0].Lifecycle
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}})
	got := updated.(Model).records[0].Lifecycle
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
