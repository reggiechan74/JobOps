# Changelog

All notable changes to JobOps will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.5.2] - 2025-12-09

### Added
- **Landing Page Development Suite**: New set of slash commands for creating professional landing pages
  - `/landing-page:create` - Main landing page builder command
    - Integrates with `frontend-design:frontend-design` skill for production-grade HTML
    - Invokes `landing-page-copywriter` agent for strategic copy development
    - Three CSS template options: `tactical` (tech/B2B), `minimal` (creative), `corporate` (enterprise)
    - Proven landing page structure: Hero, Problem, Solution, How It Works, Social Proof, CTA
    - Mobile-first responsive design with accessibility compliance
    - Playwright preview for visual validation at desktop/mobile breakpoints
    - Output: `{output-dir}/{page-name}.html`

  - `/landing-page:css-template` - CSS design system management
    - **View mode**: List available templates and component classes
    - **Analyze mode**: Deep-dive analysis of design system variables, components, responsive behavior
    - **Create mode**: Generate new CSS templates following established patterns
    - Documents component usage patterns, extension guidelines, accessibility notes
    - Available templates: tactical (`docs/subpage-styles.css`), minimal, corporate

  - `/landing-page:copywrite` - Strategic landing page copywriting
    - Three copywriting frameworks: PAS (Problem-Agitate-Solution), AIDA (Attention-Interest-Desire-Action), StoryBrand
    - Three tone options: `professional` (default), `casual`, `bold`
    - Section-by-section copy: Meta tags, Hero, Problem, Solution, How It Works, Social Proof, CTA
    - Includes A/B test variations for headlines and CTAs
    - Persuasion psychology principles: social proof, authority, scarcity, urgency, reciprocity
    - Quality checklist: benefits over features, specificity, "you" language, active voice

  - `/landing-page:copywriting-spec` - Brand copywriting specification generator
    - 8-phase interactive interview process for brand discovery
    - Complete 20-section document template embedded (self-contained, no external dependencies)
    - Sections include: Brand Identity, Target Audience, Voice & Tone, Language Guidelines, Messaging Hierarchy, Proof Points, Competitor Differentiation, Channel Requirements
    - Generates value proposition variants (10/30/50/75/150 words)
    - Problem articulation library with approved/avoid language
    - Channel-specific guidelines for website, email, LinkedIn, proposals, case studies
    - Quality checklist with 20+ validation items
    - Quick-start mode for users with existing brand guidelines
    - Output: `Client_Prospects/[CompanyName]_Copywriting_Spec.md`

- **Landing Page Copywriter Agent**: New specialized agent at `.claude/agents/landing-page/landing-page-copywriter.md`
  - Expertise in direct response marketing (Ogilvy, Hopkins, Halbert, Schwartz principles)
  - Cialdini's 6 persuasion principles integration
  - Supports PAS, AIDA, and StoryBrand frameworks
  - Section-by-section copy generation with A/B test variations
  - Quality checklist for conversion optimization

## [1.5.1] - 2025-12-05

### Fixed
- **Mobile Responsiveness for Landing Pages**: Comprehensive mobile CSS fixes across all JobOps documentation pages (GitHub Issue #39)
  - **`docs/subpage-styles.css`**: Added tablet (768px) and mobile portrait (480px) breakpoints
    - Score ranges and priority tags stack vertically on mobile
    - Method cards and related features single-column on mobile
    - Command display with overflow handling
    - Scoring grid single-column on small screens
    - Terminal and code blocks with horizontal scroll
    - Pattern lists responsive layout
    - Reduced font sizes and padding for mobile
  - **`docs/index.html`**: Enhanced mobile responsiveness
    - Nav CTA button compact sizing
    - Hero actions full-width on mobile
    - Stats bar 2-column grid with reduced gaps
    - Phase headers and step cards compact layout
    - Install commands with word-break handling
    - Terminal body with overflow protection
    - Quote and CTA sections responsive sizing
  - **`docs/wargames.html`**: WarGames theme mobile optimization
    - DEFCON bar compact layout
    - WOPR header and typing effect responsive
    - Nav grid single-column on mobile
    - Protocol phases and steps compact sizing
    - Arsenal cards responsive padding
    - Demo terminal with overflow handling
    - Launch section compact layout
  - Tested viewports: Mobile Portrait (375px), Mobile Landscape (667px), Tablet Portrait (768px), Tablet Landscape (1024px)
  - All pages now properly handle: navigation collapse, grid layouts, text overflow, code blocks, spacing, and tap targets

## [1.5.0] - 2025-12-04

### Added
- **Independent Contractor Business Development**: New suite of 5 slash commands for B2B consulting business development
  - `/defineservices` - Foundation command for defining B2B service offerings with pricing and positioning
    - Three modes: `--guided` (interactive), `--from-profile` (auto-generate), `--update` (refresh existing)
    - Pricing validation with consistency checks (daily = hourly × 8, ascending tiers)
    - JSON schema compliance for structured data
    - Output: `Client_Prospects/Service_Definition_[Date].md`

  - `/findclient` - B2B client discovery with prospect scoring
    - 10-point B2B fit scoring with 5 weighted factors (Contractor History 25%, Procurement 20%, Domain 25%, Budget 15%, Geographic 15%)
    - Priority classification: HIGH (8-10), MEDIUM (5-7.9), LOW (1-4.9)
    - Entry point identification (vendor portals, warm intros, intermediaries)
    - Filters: `--industry`, `--size`, `--location`, `--limit`
    - Output: `Client_Prospects/Prospects_[Domain]_[Date].md`

  - `/ratecard` - Professional rate card document generation
    - Pricing validation (daily = hourly × 8, ascending tiers, premium multipliers 1.3x-2.0x)
    - Multiple formats: `--format=md|html|pdf` (PDF via Playwright)
    - Currency support: `--currency=CAD|USD`
    - Optional `--include-justification` for credentials section
    - Output: `Client_Prospects/Rate_Card_[Date].[format]`

  - `/pitchdeck` - B2B service pitch deck generation
    - 12-slide structure following consulting best practices
    - Provenance hardening with >90% validation target for all claims
    - Three targeting modes: `--prospect`, `--industry`, `--service`
    - Research protocol with 5-10 minute time budget
    - Output: `Client_Prospects/Pitch_[Target]_[Date].md` (+ optional .pptx)

  - `/proposaltemplate` - Professional consulting proposal generation
    - Four engagement types: `--type=project|retainer|staff-aug|workshop`
    - Pricing algorithms with volume discounts and transparency
    - 10-section McKinsey/BCG-style structure
    - Pricing calculation metadata in YAML frontmatter
    - Output: `Client_Prospects/Proposal_[Client]_[Service]_[Date].md`

- **Service Definition Schema**: New JSON schema for validating service definitions
  - Located at `.claude/templates/service_definition_schema.json`
  - Comprehensive structure for consultant identity, services, engagement models, rate card, differentiation, and target markets

- **Client Prospects Directory**: New `Client_Prospects/` directory for independent contractor outputs
  - Service definitions, prospect lists, rate cards, pitch decks, and proposals
  - `.gitignore` configured to exclude generated files while preserving directory structure

## [1.4.0] - 2025-12-02

### Added
- **Professional PDF Resume Formatting**: New `/formatresume` command for converting markdown resumes to professionally designed PDFs
  - Uses Playwright browser automation for pixel-perfect PDF generation
  - Three theme options: `modern` (blue accents, contemporary), `classic` (serif, formal), `minimal` (clean, whitespace-focused)
  - Configurable page count: `1` (entry-level), `2` (mid-career), `3` (executive), or `auto` (content-based)
  - Iterative refinement with automatic spacing adjustments to match target page count
  - ATS-friendly output with parseable text layer
  - Visual review and quality validation with screenshots
  - Self-contained HTML with embedded CSS (no external dependencies)
  - Output: `{original_name}_formatted.pdf`
  - CSS theme files in `.claude/styles/resume-themes/`

- **Ideal Job Generator**: New `/idealjob` command for creating synthetic job descriptions
  - Analyzes complete career inventory from ResumeSourceFolder
  - Reviews Vision/Anti-Vision preferences for role alignment
  - Studies patterns from high-scoring (90+) assessment reports
  - Conducts market research to validate role viability
  - Generates realistic job description optimized for candidate fit
  - Includes search strategy with real job titles, target companies, and keywords
  - Output: `Job_Postings/IdealJob_Synthetic_[Date].md`

## [1.3.0] - 2025-11-29

### Added
- **Job Posting Quality Audit**: New `/auditjobposting` command for comprehensive job posting evaluation
  - **100-point scoring rubric** evaluating posting quality and realism before applying
  - **Internal Consistency (25 pts)**: Title-responsibility alignment, experience coherence, skills-duties match
  - **Market Realism (25 pts)**: Technology timeline validity (detects impossible experience requirements like "5 years React Native" in 2018), skill breadth feasibility, candidate pool reality assessment
  - **Compensation Assessment (20 pts)**: Salary-role/experience/location alignment with real-time market research
  - **Role Scope & Definition (20 pts)**: Single role focus (unicorn detection for unrealistic multi-role expectations), responsibility clarity, success metrics definition
  - **Red Flag Assessment (10 pts)**: Language quality analysis, organizational health signals (Glassdoor sentiment), work-life balance indicators
  - Conducts web research for technology release dates, salary benchmarks, and company reviews
  - Generates actionable interview questions to address identified concerns
  - Grade interpretation: A (90-100) pursue aggressively, B (80-89) worth applying, C (70-79) proceed with caution, D (60-69) poor quality, F (<60) avoid
  - Output: `OutputResumes/JobAudit_[Company]_[Role]_[Date].md`

## [1.2.0] - 2025-11-27

### Added
- **Job Preferences System**: New `Preferences/` folder in ResumeSourceFolder for job fit assessment
  - `Vision.md` template captures career vision and anti-vision framework
  - Comprehensive job preferences schema (v1.2.0) with 9 subsections:
    - `target_roles`: Ideal role, acceptable alternatives, target/avoid industries
    - `employment_type`: Preferred type (Employee/Consultant/Entrepreneur), acceptable types
    - `compensation`: Salary (min/target/ideal), hourly rates, entrepreneur income, equity preferences
    - `work_arrangement`: Remote/Hybrid/On-site preference, hours per week, geographic preferences
    - `travel`: Tolerance percentage and level, international travel willingness
    - `benefits`: Vacation weeks, required benefits (medical/dental/vision)
    - `work_environment`: Preferred/avoid characteristics, company size, autonomy level
    - `intellectual_property`: IP rights retention, side gig requirements
    - `deal_breakers`: Categorized absolute deal-breakers with evidence references
  - Assessment commands can now evaluate job fit against candidate preferences
  - `/create-career-history` automatically creates Preferences folder with Vision.md template

### Changed
- **Candidate Profile Schema**: Updated to v1.2.0 with `job_preferences` section
  - Schema now includes complete job preferences extraction
  - Token budget increased from 11K to 12K to accommodate preferences data
  - All preferences include evidence traceability (file + line numbers)
- **Resume Summarizer Agent**: Enhanced to extract preferences from Preferences/ folder
  - New Section 11 in extraction guidelines for job preferences
  - Quality Control Checklist updated with 3 new preference-related items
  - Success Metrics updated to verify preferences extraction

## [1.1.2] - 2025-11-17

### Changed
- **Assessment Performance Optimization**: Enhanced `/assessjob` and `/assesscandidate` workflow efficiency
  - Validated candidate profile optimization delivering 85-90% context window reduction
  - Assessment reports now consistently under 20,000 token limit through focused analysis
  - Evidence traceability maintained through structured JSON profile references
  - Improved assessment generation speed through optimized candidate data loading
  - All evidence citations include source file paths and line numbers from profile

## [1.1.1] - 2025-11-17

### Changed
- **Assessment File Organization**: Enhanced `/assessjob` and `/assesscandidate` commands with timestamped sub-folder structure
  - Assessment reports now saved to position-specific sub-folders in `OutputResumes/`
  - Sub-folder naming convention: `YYYY-MM-DD_HHMMSS_[Company]_[Role]` using Eastern time
  - Example: `OutputResumes/2025-11-17_143022_UofT_ExecutiveDirectorAssetManagement/`
  - Each assessment folder contains complete audit trail:
    - Assessment report: `Assessment_[Company]_[Role]_[Date].md`
    - Rubric copy: `Rubric_[Company]_[Role]_[Date].md` (self-contained for audit)
  - Primary rubric repository unchanged: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`
  - Benefits: Chronological organization, complete audit trail per position, easy comparison across time periods

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

#### v1.5.0 (Planned)
- Enhanced OSINT capabilities with additional intelligence sources
- Advanced interview simulation and practice mode
- Resume comparison and version tracking
- Expanded cultural profile options

#### v1.6.0 (Planned)
- AI-powered interview coaching with real-time feedback
- Salary negotiation intelligence and strategy
- Industry-specific resume templates
- Networking strategy recommendations

#### v2.0.0 (Planned)
- Multi-language support for international job markets
- Advanced analytics dashboard for application tracking
- Integration with major job boards and ATS systems
- Team collaboration features for career coaches

---

[1.5.2]: https://github.com/reggiechan74/JobOps/releases/tag/v1.5.2
[1.5.1]: https://github.com/reggiechan74/JobOps/releases/tag/v1.5.1
[1.5.0]: https://github.com/reggiechan74/JobOps/releases/tag/v1.5.0
[1.4.0]: https://github.com/reggiechan74/JobOps/releases/tag/v1.4.0
[1.3.0]: https://github.com/reggiechan74/JobOps/releases/tag/v1.3.0
[1.2.0]: https://github.com/reggiechan74/JobOps/releases/tag/v1.2.0
[1.1.2]: https://github.com/reggiechan74/JobOps/releases/tag/v1.1.2
[1.1.1]: https://github.com/reggiechan74/JobOps/releases/tag/v1.1.1
[1.1.0]: https://github.com/reggiechan74/JobOps/releases/tag/v1.1.0
[1.0.1]: https://github.com/reggiechan74/JobOps/releases/tag/v1.0.1
[1.0.0]: https://github.com/reggiechan74/JobOps/releases/tag/v1.0.0
