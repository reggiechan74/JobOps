package ui

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/reggiechan74/jobops-dashboard/internal/launch"
	"github.com/reggiechan74/jobops-dashboard/internal/model"
	"github.com/reggiechan74/jobops-dashboard/internal/scan"
)

// Scanner is the minimal adapter surface the UI needs; the real Apps adapter
// satisfies it, and tests use a static fake.
type Scanner interface {
	Scan() ([]model.Record, error)
	Skills() []model.SkillSpec
}

type uiMode int

const (
	modeNormal uiMode = iota
	modePalette
	modeNotice
)

// Model is the Bubble Tea root model for Phase 1 (single Apps tab).
type Model struct {
	root    string // workspace root
	agent   string // "claude" | "codex"
	scanner Scanner

	records []model.Record
	cursor  int

	mode      uiMode
	paletteAt int // cursor in the palette list
	skills    []model.SkillSpec
	notice    string

	width, height int
}

// New builds a Model. root and agent come from config/prefs; scanner is the
// Apps adapter (or a fake in tests).
func New(root, agent string, scanner Scanner) Model {
	recs, _ := scanner.Scan()
	return Model{
		root:    root,
		agent:   agent,
		scanner: scanner,
		records: recs,
		skills:  scanner.Skills(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

// rescanMsg is emitted after a spawned agent session returns.
type rescanMsg struct{}

// noticeMsg asks the model to display a transient notice.
type noticeMsg struct{ text string }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil
	case rescanMsg:
		m.records, _ = m.scanner.Scan()
		if m.cursor >= len(m.records) {
			m.cursor = max(0, len(m.records)-1)
		}
		if len(m.records) == 0 {
			m.mode = modeNormal
		}
		return m, nil
	case noticeMsg:
		m.notice = msg.text
		m.mode = modeNotice
		return m, nil
	case tea.KeyMsg:
		if m.mode == modePalette {
			return m.updatePalette(msg)
		}
		if m.mode == modeNotice {
			m.mode = modeNormal
			return m, nil
		}
		return m.updateNormal(msg)
	}
	return m, nil
}

func (m Model) updateNormal(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "down", "j":
		if m.cursor < len(m.records)-1 {
			m.cursor++
		}
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "enter":
		if len(m.records) > 0 {
			m.mode = modePalette
			m.paletteAt = m.nextSkillIndex()
		}
	case "c":
		if cmd := m.composeSelected(m.skillAt(m.nextSkillIndex())); cmd != "" {
			m = m.copyOrNotice(cmd)
		}
	case "C":
		if m.agent == "claude" {
			m.agent = "codex"
		} else {
			m.agent = "claude"
		}
		_ = launch.SavePrefs(m.root, launch.Prefs{Agent: m.agent})
	case "s":
		if len(m.records) > 0 {
			m = m.cycleLifecycle()
		}
	case "r":
		m.records, _ = m.scanner.Scan()
	}
	return m, nil
}

func (m Model) updatePalette(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "q":
		m.mode = modeNormal
	case "down", "j":
		if m.paletteAt < len(m.skills)-1 {
			m.paletteAt++
		}
	case "up", "k":
		if m.paletteAt > 0 {
			m.paletteAt--
		}
	case "c":
		cmd := m.composeSelected(m.skillAt(m.paletteAt))
		m.mode = modeNormal
		m = m.copyOrNotice(cmd)
	case "enter":
		cmd := m.composeSelected(m.skillAt(m.paletteAt))
		m.mode = modeNormal
		return m, m.spawn(cmd)
	}
	return m, nil
}

func (m Model) cycleLifecycle() Model {
	cur := m.records[m.cursor].Lifecycle
	if cur == "" {
		cur = model.DefaultLifecycle()
	}
	idx := 0
	for i, l := range model.LifecycleOrder {
		if l == cur {
			idx = i
			break
		}
	}
	next := model.LifecycleOrder[(idx+1)%len(model.LifecycleOrder)]
	m.records[m.cursor].Lifecycle = next
	if len(m.records[m.cursor].Paths) == 0 {
		return m
	}
	appDir := m.records[m.cursor].Paths[0]
	tr, _ := scan.ReadTracker(appDir)
	tr.Lifecycle = next
	_ = scan.WriteTracker(appDir, tr)
	m.records[m.cursor].Started = true
	return m
}

func (m Model) nextSkillIndex() int {
	if len(m.records) == 0 {
		return 0
	}
	next := m.records[m.cursor].NextSkill
	for i, s := range m.skills {
		if s.Name == next {
			return i
		}
	}
	return 0
}

func (m Model) skillAt(i int) model.SkillSpec {
	if i < 0 || i >= len(m.skills) {
		return model.SkillSpec{}
	}
	return m.skills[i]
}

func (m Model) composeSelected(spec model.SkillSpec) string {
	if spec.Name == "" || len(m.records) == 0 {
		return ""
	}
	return launch.Compose(spec, m.records[m.cursor])
}

func (m Model) copyOrNotice(cmd string) Model {
	if cmd == "" {
		return m
	}
	if launch.Copy(cmd) {
		m.notice = "Copied: " + cmd
	} else {
		m.notice = "No clipboard tool. Copy manually:\n  " + cmd
	}
	m.mode = modeNotice
	return m
}

// spawn suspends the TUI to run the interactive agent with cmd as its prompt,
// then triggers a rescan when it returns. If the agent is not on PATH it copies
// the command to the clipboard and reports that via a noticeMsg instead.
func (m Model) spawn(cmd string) tea.Cmd {
	if cmd == "" {
		return nil
	}
	if _, err := exec.LookPath(m.agent); err != nil {
		copied := launch.Copy(cmd)
		text := "Agent '" + m.agent + "' not found on PATH. "
		if copied {
			text += "Copied command to clipboard:\n  " + cmd
		} else {
			text += "Copy it manually:\n  " + cmd
		}
		return func() tea.Msg { return noticeMsg{text} }
	}
	c := exec.Command(m.agent, cmd)
	return tea.ExecProcess(c, func(error) tea.Msg { return rescanMsg{} })
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("┌ JobOps · Apps  (agent: %s) ─[%d]\n", m.agent, len(m.records)))
	b.WriteString(RenderTable(m.records, m.cursor, m.width))
	b.WriteString("├─────────────\n")
	if len(m.records) > 0 {
		b.WriteString(RenderDetail(m.records[m.cursor]))
	}
	b.WriteString("├─────────────\n")
	switch m.mode {
	case modePalette:
		b.WriteString(m.renderPalette())
	case modeNotice:
		b.WriteString(m.notice + "\n[any key] dismiss\n")
	default:
		b.WriteString("↑↓ move  ↵ run  c copy  C agent  s status  r rescan  q quit\n")
	}
	return b.String()
}

func (m Model) renderPalette() string {
	var b strings.Builder
	b.WriteString("Run on " + m.records[m.cursor].Title + ":\n")
	for i, s := range m.skills {
		cursor := "  "
		if i == m.paletteAt {
			cursor = "› "
		}
		b.WriteString(fmt.Sprintf("%s%s\n", cursor, s.Label))
	}
	b.WriteString("↵ spawn " + m.agent + "  c copy  esc cancel\n")
	return b.String()
}
