# JobOps - Intelligence-Driven Career Management Plugins

<p align="center">
  <img src="Images/JobOps_logo.png" alt="JobOps Logo" width="400">
</p>

**Version 2.7.0** | [Changelog](CHANGELOG.md) | [Why I Built This](Why_I_Built_This.md)

Two Claude Code and Codex plugins for systematic, intelligence-driven career management — from resume development to independent consulting.

## Quick Start

### Claude Code

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

### Codex

```bash
# Add the marketplace
codex plugin marketplace add reggiechan74/JobOps

# Install the core plugin
codex plugin add jobops@jobops-marketplace

# Optional: install the independent contractor add-on
codex plugin add jobops-ic@jobops-marketplace
```

After installation, start a new Codex session and invoke `jobops:setup` through Codex's explicit skill/plugin invocation surface. If you installed `jobops-ic`, run `jobops-ic:setup` after the core setup completes.

## Plugins

| Plugin | Description | Skills | Prerequisite |
|--------|-------------|--------|--------------|
| **jobops** | Resume development, interview prep, OSINT intelligence, career strategy, crisis management, application finalization | 35 | None |
| **jobops-ic** | Independent contractor toolkit — service definitions, client prospecting, pitch decks, proposals, rate cards, landing pages | 10 | jobops |

---

## Building Your Career Inventory

Every `/jobops:*` skill that produces resumes, assessments, or interview prep reads from a structured career inventory in your `resume_source` directory (default: `ResumeSourceFolder/`). The inventory is a collection of markdown files with YAML frontmatter — one file per role, credential, or technology anchor — that serves as the single source of truth for everything the plugins generate. The deeper and more specific this inventory, the less the plugins have to guess; the less they guess, the less fabrication ends up in your output.

If you're building your inventory from scratch, start here:

| Doc | What it covers |
|-----|----------------|
| [comprehensive_work_history_FAQ.md](comprehensive_work_history_FAQ.md) | **Why** a comprehensive master inventory exists — the anti-hallucination rationale and how it differs from a 2-page resume |
| [Master_Resume_Comprehensive_Setup_Guide.md](Master_Resume_Comprehensive_Setup_Guide_v1.md) | **How** to structure the inventory — documentation frameworks for experience, credentials, skills, projects, and development |
| [Guides/Experience_Template.md](Guides/Experience_Template.md) | Frontmatter + section scaffold for an individual role file under `ResumeSourceFolder/Experience/` |
| [Guides/CareerHighlight_Template.md](Guides/CareerHighlight_Template.md) | Template for certifications, publications, and non-role highlights under `ResumeSourceFolder/CareerHighlights/` |
| [Guides/TechnologyMatrix_Template.md](Guides/TechnologyMatrix_Template.md) | Template for technology anchor files under `ResumeSourceFolder/Technology/` |
| [Guides/Metadata_Standards.md](Guides/Metadata_Standards.md) | YAML frontmatter conventions shared across all inventory files |
| [Guides/Maintenance_Checklist.md](Guides/Maintenance_Checklist.md) | What to update when you add a new role, certification, or tech anchor |

The folder names (`Experience/`, `CareerHighlights/`, `Technology/`) and the frontmatter shape are what every skill in the `jobops` plugin expects when it reads `resume_source`.

---

## jobops Plugin

### Setup

| Skill | Description |
|-------|-------------|
| `/jobops:setup` | Initialize workspace — configure directories, install templates, import career history |
| `/jobops:audit-source` | Audit `ResumeSourceFolder/` for structural completeness and layout conformance; interactively fill gaps with propose-then-confirm edits |

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
| `/jobops:steelman` | Adversarial application-package review using hiring-manager OSINT |

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
| `/jobops:latex-pdf` | Primary config-driven LaTeX PDF pipeline for resumes, cover letters, and documents |
| `/jobops:convert-to-pdf` | Deprecated PDF wrapper that delegates to `/jobops:latex-pdf` |
| `/jobops:convert-to-word` | Convert markdown to Word DOCX via pandoc |
| `/jobops:markdown-to-pdf` | Deprecated PDF wrapper that delegates to `/jobops:latex-pdf` |

### System Setup

| Skill | Description |
|-------|-------------|
| `/jobops:install-pandoc` | Install pandoc for document conversion |
| `/jobops:migrate` | Migrate legacy v1.x outputs into the v2 layout |
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

| Key | Default | Purpose |
|-----|---------|---------|
| `resume_source` | `./ResumeSourceFolder` | Master HAM-Z career data (input) |
| `job_postings` | `./Job_Postings` | Target job descriptions (input) |
| `applications_root` | `./Applications` | Per-application output tree — one folder per application, with fixed `resume/`, `cover-letter/`, `assessment/`, `interview/` subfolders |
| `company_intelligence` | `./Company_Intelligence` | OSINT output tree — one folder per company, shared across applications |
| `career_analysis` | `./Career_Analysis` | Career-level outputs (idealjob, comparejobs, change-one-thing) |
| `crisis_management` | `./Crisis_Management` | Crisis-skill outputs (severance, non-compete, etc.) |
| `contractor_root` | `./Contractor` | `jobops-ic` outputs; added by `/jobops-ic:setup` |

---

## For Contributors

### Repository Structure

```
JobOps/
  .claude-plugin/marketplace.json    # Claude Code marketplace manifest
  .agents/plugins/marketplace.json   # Codex marketplace manifest
  plugins/
    jobops/                          # Core plugin
      .claude-plugin/plugin.json
      .codex-plugin/plugin.json
      skills/                        # 35 skills (flat layout)
      agents/                        # 15 Claude Code agents
      styles/                        # CSS and rendering styles
      templates/                     # Bundled templates
    jobops-ic/                       # IC add-on plugin
      .claude-plugin/plugin.json
      .codex-plugin/plugin.json
      skills/                        # 10 skills
      agents/                        # 1 Claude Code agent
      styles/                        # CSS and rendering styles
      templates/                     # Bundled templates
```

### Development

```bash
# Test a Claude Code plugin locally
claude --plugin-dir plugins/jobops

# Validate Claude Code plugin structure
claude plugin validate plugins/jobops
claude plugin validate plugins/jobops-ic
claude plugin validate .

# Validate Codex plugin structure
npm test
codex plugin marketplace add ./
codex plugin list
```

### Key Conventions

- Skill frontmatter includes `name`, `description`, and `disable-model-invocation: true` (user-invoked only), plus optional `argument-hint` where applicable
- Every skill reads `.jobops/config.json` for paths and template resolution
- jobops-ic skills include prerequisite check for jobops plugin
- Templates bundled in plugin, copied to workspace during setup
- Agents are flat (no subdirectories)

---

## License

ISC
