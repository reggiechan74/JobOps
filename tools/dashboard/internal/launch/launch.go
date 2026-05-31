// Package launch composes /jobops skill commands and provides clipboard +
// preference helpers. Spawning the interactive agent (claude/codex) is wired in
// the ui package via tea.ExecProcess; this package stays bubbletea-free so it
// is unit-testable.
package launch

import (
	"fmt"

	"github.com/reggiechan74/jobops-dashboard/internal/model"
)

// Compose builds the slash command for running spec against rec. spec.Plugin
// selects the namespace ("" defaults to "jobops").
func Compose(spec model.SkillSpec, rec model.Record) string {
	plugin := spec.Plugin
	if plugin == "" {
		plugin = "jobops"
	}
	base := "/" + plugin + ":" + spec.Name
	switch spec.Arg {
	case model.ArgJD:
		return fmt.Sprintf("%s %s.md --app=%s", base, rec.Slug, rec.Slug)
	case model.ArgCompany:
		if rec.Company != "" {
			return base + " " + rec.Company
		}
		return base
	default:
		return base
	}
}
