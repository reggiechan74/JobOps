package ui

import (
	"strings"
	"testing"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

func sampleRecords() []model.Record {
	score := 84
	return []model.Record{
		{
			Title: "Acme Product Manager", Date: "20260531", Slug: "Acme_Product_Manager_20260531",
			Company: "Acme", Score: &score, Lifecycle: model.Interviewing, Started: true,
			NextSkill: "coverletter",
			Stages: []model.Stage{
				{Name: "Assess", State: model.Final}, {Name: "Resume", State: model.Final},
				{Name: "Cover", State: model.Missing}, {Name: "OSINT", State: model.Final},
				{Name: "Briefing", State: model.Missing}, {Name: "Prep", State: model.Missing},
			},
		},
		{Title: "Beta Capital Analyst", Date: "20260101", Slug: "Beta_Capital_Analyst_20260101", Lifecycle: model.Interested},
	}
}

func TestRenderTableShowsRowsAndCursor(t *testing.T) {
	out := RenderTable(sampleRecords(), 0, 80)
	if !strings.Contains(out, "Acme Product Manager") || !strings.Contains(out, "Beta Capital Analyst") {
		t.Errorf("table missing rows:\n%s", out)
	}
	if !strings.Contains(out, "84%") {
		t.Errorf("table missing fit score:\n%s", out)
	}
	if !strings.Contains(out, "▶") {
		t.Errorf("table missing cursor marker:\n%s", out)
	}
}

func TestRenderDetailShowsStagesAndNext(t *testing.T) {
	out := RenderDetail(sampleRecords()[0])
	if !strings.Contains(out, "Assess") || !strings.Contains(out, "✓") || !strings.Contains(out, "○") {
		t.Errorf("detail missing stage glyphs:\n%s", out)
	}
	if !strings.Contains(out, "/jobops:coverletter") {
		t.Errorf("detail missing next step:\n%s", out)
	}
	if !strings.Contains(out, "Interviewing") {
		t.Errorf("detail missing lifecycle:\n%s", out)
	}
}

func TestRenderDetailBacklogNoScore(t *testing.T) {
	out := RenderDetail(sampleRecords()[1])
	if !strings.Contains(out, "—") {
		t.Errorf("backlog detail should show — for missing score:\n%s", out)
	}
}
