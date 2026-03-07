# CLAUDE.md

> **This file is for contributors only.** Operational instructions are embedded in each plugin's skills and agents.

## Repository Purpose

**JobOps v2.0.0** — Intelligence-driven career management plugins for Claude Code, distributed via a self-hosted marketplace.

Two plugins:
- **jobops** — Resume development, interview prep, OSINT intelligence, career strategy, career crisis management, application finalization
- **jobops-ic** — Independent contractor toolkit: service definitions, client prospecting, pitch decks, proposals, rate cards, landing pages (requires jobops)

## Repository Structure

```
JobOps/
  .claude-plugin/marketplace.json     # Marketplace manifest
  plugins/
    jobops/                           # Core plugin
      .claude-plugin/plugin.json
      hooks/hooks.json
      scripts/copy-templates.sh
      templates/                      # Bundled templates (copied to workspace on setup)
      skills/                         # 31 skills (flat, no subdirectories)
      agents/                         # 16 agents
    jobops-ic/                        # Independent contractor add-on
      .claude-plugin/plugin.json
      hooks/hooks.json
      scripts/copy-templates.sh
      templates/
      skills/                         # 10 skills
      agents/                         # 1 agent
  .claude/styles/                     # CSS for PDF conversion (shared)
```

## User Data Directories

Configured via `.jobops/config.json` (created by `/jobops:setup`):

| Directory | Purpose |
|-----------|---------|
| `ResumeSourceFolder/` | Master career inventory |
| `Job_Postings/` | Target job descriptions |
| `OutputResumes/` | Generated outputs |
| `Scoring_Rubrics/` | Assessment rubrics |
| `Briefing_Notes/` | Interview prep guides |
| `Intelligence_Reports/` | OSINT reports |
| `Client_Prospects/` | IC service definitions and prospects |

## Development & Testing

```bash
# Test a plugin locally
claude --plugin-dir plugins/jobops

# Validate plugin structure
/plugin validate

# Run Playwright tests
npx playwright test
```

## File Naming Conventions

- Job Postings: `Job_Postings/CompanyName_Role_Date.md`
- Rubrics: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`
- Outputs: Auto-named with step, role, company, date

## Coding Style

- Skills use `disable-model-invocation: true` frontmatter
- Every skill (except setup) starts with config preamble reading `.jobops/config.json`
- jobops-ic skills include prerequisite check for jobops plugin
- Templates bundled in plugin, copied to `.jobops/templates/default/` via setup

## Version Management

Semantic versioning. Update: `package.json`, plugin.json files, `README.md`, `CHANGELOG.md`. Tag releases: `git tag -a vX.Y.Z`.
