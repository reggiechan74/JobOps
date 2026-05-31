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

// Scanner is the minimal adapter surface the UI needs; the real adapters satisfy
// it, and tests use static fakes.
type Scanner interface {
	Scan() ([]model.Record, error)
	Skills() []model.SkillSpec
}

// TabSource describes one tab the dashboard shows.
type TabSource struct {
	Name      string
	Scanner   Scanner
	Lifecycle bool // records carry a real-world lifecycle (Apps only)
}

type uiMode int

const (
	modeNormal uiMode = iota
	modePalette
	modeNotice
)

// tab holds the per-tab state.
type tab struct {
	name      string
	scanner   Scanner
	skills    []model.SkillSpec
	lifecycle bool
	records   []model.Record
	cursor    int
}

// Model is the Bubble Tea root model: a set of tabs plus shared UI state.
type Model struct {
	root  string
	agent string

	tabs   []tab
	active int

	mode      uiMode
	paletteAt int
	notice    string

	width, height int
}

// New builds a Model from one TabSource per tab, scanning each upfront.
func New(root, agent string, sources []TabSource) Model {
	tabs := make([]tab, 0, len(sources))
	for _, s := range sources {
		recs, _ := s.Scanner.Scan()
		tabs = append(tabs, tab{
			name:      s.Name,
			scanner:   s.Scanner,
			skills:    s.Scanner.Skills(),
			lifecycle: s.Lifecycle,
			records:   recs,
		})
	}
	return Model{root: root, agent: agent, tabs: tabs}
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
		t := &m.tabs[m.active]
		t.records, _ = t.scanner.Scan()
		if t.cursor >= len(t.records) {
			t.cursor = max(0, len(t.records)-1)
		}
		if len(t.records) == 0 {
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
	t := &m.tabs[m.active]
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "down", "j":
		if t.cursor < len(t.records)-1 {
			t.cursor++
		}
	case "up", "k":
		if t.cursor > 0 {
			t.cursor--
		}
	case "right", "tab":
		m.active = (m.active + 1) % len(m.tabs)
	case "left", "shift+tab":
		m.active = (m.active - 1 + len(m.tabs)) % len(m.tabs)
	case "enter":
		if len(t.records) > 0 || len(t.skills) > 0 {
			m.mode = modePalette
			m.paletteAt = m.nextSkillIndex()
		}
	case "c":
		if len(t.records) > 0 || len(t.skills) > 0 {
			if cmd := m.composeSelected(m.skillAt(m.nextSkillIndex())); cmd != "" {
				m = m.copyOrNotice(cmd)
			}
		}
	case "C":
		if m.agent == "claude" {
			m.agent = "codex"
		} else {
			m.agent = "claude"
		}
		_ = launch.SavePrefs(m.root, launch.Prefs{Agent: m.agent})
	case "s":
		if t.lifecycle && len(t.records) > 0 {
			m = m.cycleLifecycle()
		}
	case "r":
		return m, func() tea.Msg { return rescanMsg{} }
	}
	return m, nil
}

func (m Model) updatePalette(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	t := &m.tabs[m.active]
	switch msg.String() {
	case "esc", "q":
		m.mode = modeNormal
	case "down", "j":
		if m.paletteAt < len(t.skills)-1 {
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
	t := &m.tabs[m.active]
	rec := &t.records[t.cursor]
	cur := rec.Lifecycle
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
	rec.Lifecycle = next
	if len(rec.Paths) == 0 {
		return m
	}
	appDir := rec.Paths[0]
	tr, _ := scan.ReadTracker(appDir)
	tr.Lifecycle = next
	_ = scan.WriteTracker(appDir, tr)
	rec.Started = true
	return m
}

func (m Model) nextSkillIndex() int {
	t := m.tabs[m.active]
	if len(t.records) == 0 {
		return 0
	}
	next := t.records[t.cursor].NextSkill
	for i, s := range t.skills {
		if s.Name == next {
			return i
		}
	}
	return 0
}

func (m Model) skillAt(i int) model.SkillSpec {
	t := m.tabs[m.active]
	if i < 0 || i >= len(t.skills) {
		return model.SkillSpec{}
	}
	return t.skills[i]
}

func (m Model) composeSelected(spec model.SkillSpec) string {
	if spec.Name == "" {
		return ""
	}
	if spec.Arg == model.ArgNone {
		return launch.Compose(spec, model.Record{})
	}
	t := m.tabs[m.active]
	if len(t.records) == 0 {
		return ""
	}
	return launch.Compose(spec, t.records[t.cursor])
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
	if len(m.tabs) == 0 {
		return "No tabs configured.\n"
	}
	t := m.tabs[m.active]
	var b strings.Builder
	b.WriteString(m.renderTabBar())
	b.WriteString(RenderTable(t.records, t.cursor, m.width, t.lifecycle))
	b.WriteString("├─────────────\n")
	if len(t.records) > 0 {
		b.WriteString(RenderDetail(t.records[t.cursor], t.lifecycle))
	}
	b.WriteString("├─────────────\n")
	switch m.mode {
	case modePalette:
		b.WriteString(m.renderPalette())
	case modeNotice:
		b.WriteString(m.notice + "\n[any key] dismiss\n")
	default:
		b.WriteString("↑↓ move  ←→ tab  ↵ run  c copy  C agent  s status  r rescan  q quit\n")
	}
	return b.String()
}

func (m Model) renderTabBar() string {
	var names []string
	for i, tb := range m.tabs {
		if i == m.active {
			names = append(names, "["+tb.name+"]")
		} else {
			names = append(names, tb.name)
		}
	}
	return fmt.Sprintf("┌ JobOps · %s  (agent: %s) ─[%d]\n",
		strings.Join(names, " "), m.agent, len(m.tabs[m.active].records))
}

func (m Model) renderPalette() string {
	t := m.tabs[m.active]
	target := t.name
	if len(t.records) > 0 {
		target = t.records[t.cursor].Title
	}
	var b strings.Builder
	b.WriteString("Run on " + target + ":\n")
	for i, s := range t.skills {
		cursor := "  "
		if i == m.paletteAt {
			cursor = "› "
		}
		b.WriteString(fmt.Sprintf("%s%s\n", cursor, s.Label))
	}
	b.WriteString("↵ spawn " + m.agent + "  c copy  esc cancel\n")
	return b.String()
}
