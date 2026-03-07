# JobOps - Intelligence-Driven Career Management Plugins

<p align="center">
  <img src="Images/JobOps_logo.png" alt="JobOps Logo" width="400">
</p>

**Version 2.0.0** | [Changelog](CHANGELOG.md) | [Why I Built This](Why_I_Built_This.md)

Two Claude Code plugins for systematic, intelligence-driven career management — from resume development to independent consulting.

## Quick Start

```bash
# Add the marketplace
/plugin marketplace add reggiechan74/JobOps

# Install the core plugin
/plugin install jobops@jobops-marketplace

# Initialize your workspace
/jobops:setup

# Optional: install the independent contractor add-on
/plugin install jobops-ic@jobops-marketplace
/jobops-ic:setup
```

## Plugins

| Plugin | Description | Skills | Prerequisite |
|--------|-------------|--------|--------------|
| **jobops** | Resume development, interview prep, OSINT intelligence, career strategy, crisis management, application finalization | 31 | None |
| **jobops-ic** | Independent contractor toolkit — service definitions, client prospecting, pitch decks, proposals, rate cards, landing pages | 10 | jobops |

---

## jobops Plugin

### Setup

| Skill | Description |
|-------|-------------|
| `/jobops:setup` | Initialize workspace — configure directories, install templates, import career history |

### Resume Development

| Skill | Description |
|-------|-------------|
| `/jobops:buildresume` | Complete 3-step resume process (draft, provenance check, final) |
| `/jobops:provenance-check` | Standalone credibility analysis on a draft resume |
| `/jobops:coverletter` | Strategic cover letter with requirements-matching table |

### Interview Prep

| Skill | Description |
|-------|-------------|
| `/jobops:assessjob` | Dynamic rubric generation + full candidate assessment |
| `/jobops:assesscandidate` | Assess candidate using pre-created scoring rubric |
| `/jobops:createrubric` | Create reusable 200-point scoring rubric |
| `/jobops:comparejobs` | Compare 2-4 job assessments side-by-side |
| `/jobops:briefing` | Study guide with priority-tagged skill gap analysis |
| `/jobops:interviewprep` | Interview questions with likelihood and difficulty tags |

### Job Search & Intelligence

| Skill | Description |
|-------|-------------|
| `/jobops:osint` | 7-agent parallel company intelligence gathering |
| `/jobops:auditjobposting` | 100-point job posting quality and realism audit |

### Career Strategy

| Skill | Description |
|-------|-------------|
| `/jobops:idealjob` | Generate synthetic ideal job description from career history |
| `/jobops:change-one-thing` | Career retrospective with counterfactual pivot analysis |
| `/jobops:assess-job-offer` | Comprehensive offer analysis (compensation, legal, alignment) |

### Career Crisis Management

| Skill | Description |
|-------|-------------|
| `/jobops:code-red` | Employment crisis intervention (PIPs, HR conflicts, termination risk) |
| `/jobops:severance-review` | Severance package analysis and negotiation strategy |
| `/jobops:workplace-documentation` | Structured incident logging and timeline building |
| `/jobops:non-compete-analysis` | Restrictive covenant enforceability analysis |
| `/jobops:reference-shield` | Reference risk management and strategy |
| `/jobops:unemployment-prep` | Unemployment/EI claim preparation and defense |
| `/jobops:discrimination-assessment` | Protected class treatment pattern assessment |
| `/jobops:investigation-prep` | Workplace investigation preparation |
| `/jobops:accommodation-request` | Accommodation request builder (disability, religious, medical) |
| `/jobops:layoff-intel` | Layoff risk assessment and proactive preparation |
| `/jobops:constructive-dismissal` | Constructive dismissal conditions assessment |

### Application Finalization

| Skill | Description |
|-------|-------------|
| `/jobops:convert-to-pdf` | Convert markdown resume to PDF via Playwright |
| `/jobops:convert-to-word` | Convert markdown to Word DOCX via pandoc |
| `/jobops:markdown-to-pdf` | Convert markdown documents to PDF with Obsidian styling |

### System Setup

| Skill | Description |
|-------|-------------|
| `/jobops:install-pandoc` | Install pandoc for document conversion |
| `/jobops:github-portfolio` | Create or update GitHub portfolio documentation |

---

## jobops-ic Plugin

> **Requires:** jobops plugin must be installed first.

### Setup

| Skill | Description |
|-------|-------------|
| `/jobops-ic:setup` | Extend workspace with IC directories and templates |

### Independent Contractor

| Skill | Description |
|-------|-------------|
| `/jobops-ic:defineservices` | Define service catalog with pricing and positioning |
| `/jobops-ic:findclient` | B2B prospect discovery with 10-point fit scoring |
| `/jobops-ic:pitchdeck` | Generate provenance-hardened B2B pitch deck |
| `/jobops-ic:proposaltemplate` | McKinsey/BCG-style consulting proposal with pricing |
| `/jobops-ic:ratecard` | Professional rate card with pricing validation |

### Landing Pages

| Skill | Description |
|-------|-------------|
| `/jobops-ic:copywrite` | Strategic copy using PAS/AIDA/StoryBrand frameworks |
| `/jobops-ic:copywriting-spec` | Comprehensive brand copywriting specification |
| `/jobops-ic:create-landing-page` | Professional landing page with CSS design system |
| `/jobops-ic:css-template` | Manage CSS design system templates |

---

## Configuration

### Workspace Setup

Running `/jobops:setup` creates a `.jobops/` directory in your workspace:

```
.jobops/
  config.json                    # Directory paths and template settings
  templates/
    default/                     # Plugin-provided templates (read-only by convention)
    custom/                      # Your custom template overrides
```

### Template Customization

1. Copy a default template to `custom/`: `cp .jobops/templates/default/assessment_rubric_framework.md .jobops/templates/custom/`
2. Edit the custom version
3. Update `.jobops/config.json` to point to `"custom"` for that template

### Directory Configuration

All output directories are configurable via `.jobops/config.json`. Default paths:

| Key | Default |
|-----|---------|
| `resume_source` | `./ResumeSourceFolder` |
| `job_postings` | `./Job_Postings` |
| `output_resumes` | `./OutputResumes` |
| `scoring_rubrics` | `./Scoring_Rubrics` |
| `briefing_notes` | `./Briefing_Notes` |
| `intelligence_reports` | `./Intelligence_Reports` |
| `client_prospects` | `./Client_Prospects` |

---

## For Contributors

### Repository Structure

```
JobOps/
  .claude-plugin/marketplace.json    # Marketplace manifest
  plugins/
    jobops/                          # Core plugin
      .claude-plugin/plugin.json
      skills/                        # 31 skills (flat layout)
      agents/                        # 16 agents
      templates/                     # Bundled templates
      hooks/                         # SessionStart hook
      scripts/                       # Template copy script
    jobops-ic/                       # IC add-on plugin
      .claude-plugin/plugin.json
      skills/                        # 10 skills
      agents/                        # 1 agent
      templates/                     # Bundled templates
      hooks/                         # SessionStart hook
      scripts/                       # Template copy script
```

### Development

```bash
# Test a plugin locally
claude --plugin-dir plugins/jobops

# Validate plugin structure
/plugin validate
```

### Key Conventions

- Skills use `disable-model-invocation: true` (user-invoked only)
- Every skill reads `.jobops/config.json` for paths and template resolution
- jobops-ic skills include prerequisite check for jobops plugin
- Templates bundled in plugin, copied to workspace via setup script
- Agents are flat (no subdirectories)

---

## License

ISC
