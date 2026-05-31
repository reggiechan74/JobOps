// Package ui renders the dashboard. view.go holds pure string-producing helpers
// (no Bubble Tea state) so they can be unit-tested directly.
package ui

import (
	"fmt"
	"strings"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func scoreCell(r model.Record) string {
	if r.Score == nil {
		return "—"
	}
	return fmt.Sprintf("%d%%", *r.Score)
}

func lifecycleCell(r model.Record) string {
	l := r.Lifecycle
	if l == "" {
		l = model.DefaultLifecycle()
	}
	return l.Label()
}

// pipelineBar renders the stage states as a compact glyph run, e.g. "✓✓○✓○○".
func pipelineBar(r model.Record) string {
	var b strings.Builder
	for _, s := range r.Stages {
		b.WriteString(s.State.Glyph())
	}
	return b.String()
}

// RenderTable renders the record list. cursor is the selected row index; width
// caps title length; showLifecycle toggles the Lifecycle column (Apps only).
func RenderTable(recs []model.Record, cursor, width int, showLifecycle bool) string {
	titleW := 28
	if width > 0 && width < 80 {
		titleW = 18
	}
	var b strings.Builder
	if showLifecycle {
		b.WriteString(fmt.Sprintf("  %-*s %-5s %-13s %s\n", titleW, "Application", "Fit", "Lifecycle", "Pipeline"))
	} else {
		b.WriteString(fmt.Sprintf("  %-*s %-5s %s\n", titleW, "Name", "Fit", "Progress"))
	}
	for i, r := range recs {
		marker := " "
		if i == cursor {
			marker = "▶"
		}
		title := r.Title
		if len(title) > titleW {
			title = title[:titleW-1] + "…"
		}
		warn := ""
		if r.Warn != "" {
			warn = " ⚠"
		}
		if showLifecycle {
			b.WriteString(fmt.Sprintf("%s %-*s %-5s %-13s %s%s\n",
				marker, titleW, title, scoreCell(r), lifecycleCell(r), pipelineBar(r), warn))
		} else {
			b.WriteString(fmt.Sprintf("%s %-*s %-5s %s%s\n",
				marker, titleW, title, scoreCell(r), pipelineBar(r), warn))
		}
	}
	if len(recs) == 0 {
		b.WriteString("  (nothing here yet)\n")
	}
	return b.String()
}

// RenderDetail renders the detail pane for one record. showLifecycle toggles the
// lifecycle/started line; the pipeline line is shown only when the record has
// stages (flat tabs have none).
func RenderDetail(r model.Record, showLifecycle bool) string {
	var b strings.Builder
	header := r.Title
	if r.Date != "" {
		header += "  ·  " + r.Date
	}
	b.WriteString(header + "\n")
	if showLifecycle {
		b.WriteString(fmt.Sprintf("Fit %s   Lifecycle: %s   %s\n",
			scoreCell(r), lifecycleCell(r), startedLabel(r)))
	} else {
		b.WriteString(fmt.Sprintf("Fit %s\n", scoreCell(r)))
	}

	if len(r.Stages) > 0 {
		var stages []string
		for _, s := range r.Stages {
			stages = append(stages, fmt.Sprintf("%s %s", s.State.Glyph(), s.Name))
		}
		b.WriteString(strings.Join(stages, "   ") + "\n")
		if r.NextSkill != "" {
			b.WriteString("Next: /jobops:" + r.NextSkill + "\n")
		} else {
			b.WriteString("Pipeline complete\n")
		}
	}
	if r.Warn != "" {
		b.WriteString("⚠ " + r.Warn + "\n")
	}
	return b.String()
}

func startedLabel(r model.Record) string {
	if r.Started {
		return "(in progress)"
	}
	return "(backlog)"
}
