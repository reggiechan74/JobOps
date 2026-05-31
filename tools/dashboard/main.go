// Command jobops-dash is a terminal dashboard for tracking job applications and
// launching /jobops skills.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/reggiechan74/jobops-dashboard/internal/config"
	"github.com/reggiechan74/jobops-dashboard/internal/launch"
	"github.com/reggiechan74/jobops-dashboard/internal/scan"
	"github.com/reggiechan74/jobops-dashboard/internal/ui"
)

func main() {
	wsFlag := flag.String("workspace", "", "path to the JobOps workspace (default: discover from CWD)")
	flag.Parse()

	start := *wsFlag
	if start == "" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot determine working directory:", err)
			os.Exit(1)
		}
		start = cwd
	}

	root, err := config.Discover(start)
	if err != nil {
		// Not configured: offer to run setup.
		runUnconfigured(start)
		return
	}

	cfg, err := config.Load(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read .jobops/config.json:", err)
		os.Exit(1)
	}

	prefs := launch.LoadPrefs(root)
	sources := []ui.TabSource{
		{Name: "Apps", Scanner: scan.AppsAdapter{Cfg: cfg}, Lifecycle: true},
		{Name: "Companies", Scanner: scan.CompaniesAdapter{Cfg: cfg}},
		{Name: "Career", Scanner: scan.CareerAdapter{Cfg: cfg}},
		{Name: "Crisis", Scanner: scan.CrisisAdapter{Cfg: cfg}},
	}
	// Contractor (jobops-ic) only when its root is configured.
	if cfg.Directories.ContractorRoot != "" {
		sources = append(sources, ui.TabSource{Name: "Contractor", Scanner: scan.ContractorAdapter{Cfg: cfg}})
	}
	model := ui.New(root, prefs.Agent, sources)

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "dashboard error:", err)
		os.Exit(1)
	}
}

// runUnconfigured prints guidance and, if claude is available, offers to run
// /jobops:setup directly.
func runUnconfigured(start string) {
	fmt.Println("No .jobops workspace found at or above:", start)
	fmt.Println()
	if _, err := exec.LookPath("claude"); err == nil {
		fmt.Print("Run /jobops:setup now with claude? [y/N] ")
		var ans string
		fmt.Scanln(&ans)
		if ans == "y" || ans == "Y" {
			c := exec.Command("claude", "/jobops:setup")
			c.Stdin, c.Stdout, c.Stderr = os.Stdin, os.Stdout, os.Stderr
			_ = c.Run()
			return
		}
	}
	fmt.Println("Run `/jobops:setup` in Claude Code to initialize the workspace, then re-run jobops-dash.")
}
