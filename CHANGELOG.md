# Changelog

All notable changes to JobOps will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2025-11-09

### Added
- **Career Retrospective Analysis Tool**: New `/change-one-thing` command for rigorous career analysis
  - 16-lens comprehensive pattern recognition framework across 4 analytical phases
  - External patterns: Skill trajectory, network building, technical decisions, learning gaps, visibility, cross-pollination, geographic positioning
  - Internal constraints: Family dynamics, financial capacity, health/energy analysis
  - Temporal mechanics: Information asymmetry correction, butterfly effect mapping, irreversibility analysis
  - Probabilistic outcomes: Luck surface area, values evolution, skill vs. timing attribution
  - Multi-candidate recommendation system optimized for different objectives (wealth, optionality, impact, balance)
  - Counterfactual rigor with quantified opportunity costs and risk assessment
  - Automatic file output: 3-4 comprehensive analysis parts + executive summary (≤6,000 tokens)
  - Values-aligned recommendations based on user interview about priorities and constraints
  - "What you can still do now" actionable recovery plans with immediate next steps
  - Outputs saved to `OutputResumes/Change_One_Thing_Analysis_Part[1-3]_YYYYMMDD.md` and `Change_One_Thing_EXECUTIVE_SUMMARY_YYYYMMDD.md`

## [1.0.1] - 2025-11-07

### Fixed
- **Credential Completeness Verification**: Enhanced `/buildresume` command to prevent omission of relevant education and credentials
  - Step 1 agent now reviews ALL credentials from master resume against job requirements
  - Step 1 agent must explicitly justify any credential exclusions in output
  - Step 2 provenance check now verifies all credentials are evaluated
  - Step 2 flags any excluded credentials that map to job requirements
  - Fixes issue where relevant qualifications (e.g., BCIT Financial Management diploma for financial analysis roles) were excluded without justification

## [1.0.0] - 2025-10-02

### Added - Initial Release

#### Core Resume Development System
- **HAM-Z™ Methodology**: 3-step resume assembly (draft → provenance verification → final hardened resume)
- **Provenance Hardening**: Comprehensive credibility analysis with risk categorization (Critical/High/Medium/Low)
- **Cultural Profiles**: Theater-specific resume adaptation (Canadian, US, European, UK, Australian)
- **Master Resume System**: Arsenal inventory approach for comprehensive career documentation

#### Assessment & Intelligence
- **Dynamic Job Scoring**: Custom 100-point assessment rubrics generated from job postings
- **Assessment-First Hiring Model**: Evidence-based candidate evaluation preventing misaligned applications
- **Distributed OSINT System**: 6 parallel specialized intelligence agents
  - Corporate intelligence (structure, financials, strategy)
  - Legal intelligence (litigation, regulatory, compliance)
  - Leadership intelligence (executives, personnel)
  - Compensation intelligence (salaries, benefits, total rewards)
  - Culture intelligence (employee sentiment, workplace environment)
  - Market intelligence (industry analysis, competitive landscape)

#### Interview Preparation
- **Gap Analysis**: Identifies skill gaps with severity ratings and training timelines
- **Intelligence Briefings**: Study guides with web-researched learning resources
- **Interview Question Generator**: Customized technical and behavioral questions with STAR format coaching
- **Modular Assessment Workflow**: Reusable rubrics for consistent candidate evaluation

#### Job Search & Targeting
- **Hybrid Job Search**: Two-phase system combining API search + Playwright scraping
- **Complete Job Descriptions**: Verbatim job posting extraction with full details
- **Target Acquisition**: Smart filtering by keywords, location, company, with batch processing
- **Automatic Saving**: Structured markdown output with metadata for downstream processing

#### Commands & Agents
**Core Commands:**
- `/assessjob` - Target reconnaissance with dynamic rubric creation
- `/buildresume` - Complete 3-step resume assembly
- `/briefing` - Intelligence brief generation for interview prep
- `/interviewprep` - Interview question generation
- `/searchjobs` - Hybrid job search with complete descriptions
- `/osint` - Comprehensive company intelligence gathering
- `/coverletter` - Strategic cover letter with requirements-matching table
- `/comparejobs` - Multi-opportunity prioritization analysis
- `/createrubric` - Standalone rubric generation for reuse
- `/assesscandidate` - Assessment using pre-created rubrics
- `/provenance` - Standalone credibility verification
- `/convert` - Markdown to Word DOCX conversion
- `/install-pandoc` - Pandoc installation utility

**Specialized Agents:**
- 15 specialized agents for resume development, assessment, intelligence, and job search
- Parallel agent deployment for maximum efficiency
- Standardized YAML front matter for all outputs

#### Documentation & Branding
- **Tactical Warfare Theme**: Intelligence-driven application warfare positioning
- **JobOps Logo**: Custom tactical branding with social media preview
- **Comprehensive Documentation**: README, setup guides, FAQs, methodology documentation
- **Tactical CLI Messaging**: Warfare-themed status messages (✓ checkmarks, tactical terminology)
- **Optional Statusline**: Repository-specific tactical statusline configuration

#### File Organization
- **Structured Directories**: Arsenal inventory, target coordinates, intelligence reports, briefing notes
- **YAML Front Matter**: Standardized metadata for all generated documents
- **Naming Conventions**: Consistent file naming across all outputs
- **Version Tracking**: Built-in versioning for all generated artifacts

#### Theoretical Foundation
- **Assessment-First Hiring Model**: System dynamics approach documented in SourceMaterial/
- **Signal-to-Noise Analysis**: Addresses AI-assisted resume embellishment crisis
- **Vicious Cycle of Embellishment**: Breaks the cycle with evidence-based assessment

### Technical Stack
- Claude Code CLI with custom agents and slash commands
- Playwright MCP server for browser automation
- Pandoc for document conversion
- Node.js for dependencies
- Markdown-based workflow with YAML metadata

### Documentation
- README.md - Complete system documentation with tactical branding
- SETUP.md - Installation and configuration guide
- CLAUDE.md - System instructions and agent documentation
- Master_Resume_Comprehensive_Setup_Guide_v1.md - Arsenal inventory methodology
- comprehensive_work_history_FAQ.md - Master resume philosophy
- JOB_SEARCH_GUIDE.md - Job search workflow documentation
- SourceMaterial/ - Theoretical foundation and analysis documents

---

## Version History

### Versioning Scheme

JobOps follows [Semantic Versioning](https://semver.org/):
- **MAJOR.MINOR.PATCH** (e.g., 1.0.0)
- **MAJOR**: Breaking changes, incompatible API changes, major rewrites
- **MINOR**: New features, enhancements, backward-compatible functionality
- **PATCH**: Bug fixes, documentation updates, minor improvements

### Planned Releases

#### v1.1.0 (Planned)
- Enhanced OSINT capabilities with additional intelligence sources
- Advanced interview simulation and practice mode
- Resume comparison and version tracking
- Expanded cultural profile options

#### v1.2.0 (Planned)
- AI-powered interview coaching with real-time feedback
- Salary negotiation intelligence and strategy
- Career trajectory analysis and planning
- Industry-specific resume templates

#### v2.0.0 (Planned)
- Multi-language support for international job markets
- Advanced analytics dashboard for application tracking
- Integration with major job boards and ATS systems
- Team collaboration features for career coaches

---

[1.1.0]: https://github.com/reggiechan74/JobOps/releases/tag/v1.1.0
[1.0.1]: https://github.com/reggiechan74/JobOps/releases/tag/v1.0.1
[1.0.0]: https://github.com/reggiechan74/JobOps/releases/tag/v1.0.0
