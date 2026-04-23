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

Configured via `.jobops/config.json` (created by `/jobops:setup`, extended by `/jobops-ic:setup`). Config keys (not hardcoded paths) are authoritative — see `docs/ARCHITECTURE.md` for the full contract.

| Config key | Default | Purpose |
|------------|---------|---------|
| `resume_source` | `ResumeSourceFolder/` | Master HAM-Z career inventory (input) |
| `job_postings` | `Job_Postings/` | Target job descriptions (input) |
| `applications_root` | `Applications/` | Per-application output tree with fixed `resume/`, `cover-letter/`, `assessment/`, `interview/` subfolders + pinned `job_posting.md` |
| `company_intelligence` | `Company_Intelligence/` | OSINT output tree, one folder per company |
| `career_analysis` | `Career_Analysis/` | Flat timestamped career-level outputs |
| `crisis_management` | `Crisis_Management/` | Flat timestamped crisis-skill outputs |
| `contractor_root` | `Contractor/` | `jobops-ic` outputs (services, prospects, proposals, pitches, rate-cards, landing-pages); added by `/jobops-ic:setup` |

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

- Job Postings: `Job_Postings/{Company}_{Role}_{YYYYMMDD}.md`
- Application artifacts: `Applications/{Company}_{Role}_{YYYYMMDD}/<subfolder>/<fixed-filename>.md` (slug parsed from the JD filename; sub-folder and filename are plugin convention — see `docs/ARCHITECTURE.md` Section 4)
- OSINT: `Company_Intelligence/{Company}/{corporate,legal,leadership,compensation,culture,market,summary}.md`
- Career / crisis / contractor: single timestamped file per invocation under the appropriate root

## Coding Style

- Skills use `disable-model-invocation: true` frontmatter
- Every skill (except setup) starts with config preamble reading `.jobops/config.json`
- jobops-ic skills include prerequisite check for jobops plugin
- Templates bundled in plugin, copied to `.jobops/templates/default/` via setup

## Version Management

Semantic versioning. Update: `package.json`, plugin.json files, `README.md`, `CHANGELOG.md`. Tag releases: `git tag -a vX.Y.Z`.
