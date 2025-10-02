# Resume Optimizer for Claude Code

A comprehensive job application system that transforms career inventories into targeted resumes and provides complete interview preparation using an 8-step methodology with provenance hardening.

## Quick Start

1. Install dependencies: `npm run install-all`
2. Populate `ResumeSourceFolder/` with your master resume inventory and add target roles to `Job_Postings/`
3. Start the Playwright MCP server (see "Run the Playwright MCP server" in `SETUP.md`) so Claude Code can reach browser automation
4. Launch Claude Code in this repository and run slash commands such as `/assessjob Job_Postings/Example.md` followed by `/buildresume`

## Architecture Overview

- **Claude Code CLI** orchestrates the workflow: users issue slash commands that call specialized agents under `.claude/`
- **Playwright MCP server** provides headless browser automation for job sourcing and intelligence gathering
- **Content directories** (`ResumeSourceFolder/`, `Job_Postings/`, `OutputResumes/`, `Briefing_Notes/`) hold source materials, generated analyses, and deliverables
- **HAM-Z agents** execute each step: resume drafting, provenance hardening, and interview preparation
- **Scoring and OSINT resources** in `Scoring_Rubrics/` and `Intelligence_Reports/` feed assessments and research deliverables

## Overview

Resume Optimizer uses the HAM-Z™ methodology to create powerful, defensible resumes that are:
- **Strategically positioned** for specific roles
- **Culturally attuned** to regional preferences
- **Provenance-hardened** to withstand scrutiny
- **ATS-optimized** for applicant tracking systems

The system employs an 8-step process organized into three phases: resume development (Steps 1-3), interview preparation (Steps 4-6), and application finalization (Steps 7-8), ensuring every claim is backed by evidence and every achievement is quantified with mechanism, metrics, and timeframe.

### Theoretical Foundation: Assessment-First Hiring

This application was designed around the **Assessment-First Hiring Model**, a system dynamics approach that addresses the fundamental inefficiency in modern job markets. The theoretical framework, detailed in [`SourceMaterial/System_Dynamics_Analysis_Assessment_First_Hiring_v3.md`](SourceMaterial/System_Dynamics_Analysis_Assessment_First_Hiring_v3.md), demonstrates how AI-assisted resume embellishment has created a "Vicious Cycle of Embellishment" where:

- 73.4% of job seekers use AI tools to enhance resumes
- Signal-to-noise ratio has degraded to **-2.5 dB** (more noise than signal)
- Employers add screening layers, incentivizing further embellishment

The assessment-first model breaks this cycle by evaluating candidate fit *before* application creation, potentially achieving:
- **35-60% reduction** in misaligned applications
- **30-60% decrease** in time-to-hire
- **40-80% improvement** in overall market matching efficiency
- **+3.2 to +10.4 dB improvement** in signal quality

This theoretical foundation explains why the system starts with `/assessjob` rather than resume creation - preventing wasted effort on fundamentally poor matches while ensuring only high-probability applications receive full development.

## Key Features

- **HAM-Z™ Methodology**: Combines Hard Skills, Actions, Metrics, and Structure with XYZ narrative format
- **Provenance Hardening**: Comprehensive credibility analysis to identify and fix evidence gaps
- **Cultural Profiles**: Tailors resume style to regional expectations (Canadian, US, European, etc.)
- **Master Resume System**: Maintains comprehensive career inventory separate from targeted resumes
- **Dynamic Job Scoring**: Creates custom assessment rubrics for each job posting
- **Interview Preparation**: Gap analysis briefings and customized question generation
- **Distributed OSINT Intelligence**: Professional-grade company research with parallel specialized agents
- **Hybrid Job Search**: API search combined with Playwright scraping for complete job descriptions

## Directory Structure

```
resumeoptimizer/
├── .claude/                           # Claude Code configuration
│   ├── agents/                        # 15 specialized processing agents
│   │   ├── candidate-assessment.md    # HR-level job fit assessment
│   │   ├── hiringcafe-search.md      # Job search via hiring.cafe API
│   │   ├── interview-briefing.md     # Study guide and gap analysis
│   │   ├── interview-question-generator.md # Custom interview questions
│   │   ├── osint-agent.md            # Master OSINT orchestrator
│   │   ├── osint-compensation.md     # Salary and benefits intelligence
│   │   ├── osint-corporate.md        # Company structure and financials
│   │   ├── osint-culture.md          # Workplace culture analysis
│   │   ├── osint-leadership.md       # Executive background checks
│   │   ├── osint-legal.md            # Litigation and compliance
│   │   ├── osint-market.md           # Industry and competition
│   │   ├── step1-resume-draft.md     # Initial tailored resume
│   │   ├── step2-provenance-check.md # Credibility verification
│   │   ├── step3-final-resume.md     # Hardened final resume
│   │   └── step4-cover-letter.md     # Strategic cover letter
│   ├── commands/                      # 14 slash command definitions
│   │   ├── assesscandidate.md        # /assesscandidate - Use pre-created rubric
│   │   ├── assessjob.md              # /assessjob - Complete assessment
│   │   ├── briefing.md               # /briefing - Interview prep guide
│   │   ├── buildresume.md            # /buildresume - 3-step resume process
│   │   ├── comparejobs.md            # /comparejobs - Multi-job analysis
│   │   ├── convert.md                # /convert - Markdown to DOCX
│   │   ├── coverletter.md            # /coverletter - Generate letter
│   │   ├── createrubric.md           # /createrubric - Rubric only
│   │   ├── install-pandoc.md         # /install-pandoc - Setup converter
│   │   ├── interviewprep.md          # /interviewprep - Question generator
│   │   ├── osint.md                  # /osint - Company intelligence
│   │   ├── prime.md                  # /prime - System primer
│   │   ├── provenance.md             # /provenance - Standalone check
│   │   └── searchjobs.md             # /searchjobs - hiring.cafe search
│   └── settings.local.json            # Local Claude Code settings
├── .playwright-mcp/                   # Playwright browser automation cache
├── Briefing_Notes/                    # Interview preparation materials
│   └── Briefing_*_*_*_*.md           # Gap analysis and study guides
├── Intelligence_Reports/              # OSINT company intelligence reports
│   ├── [Company]_*_Intelligence_*.md  # Specialized intelligence reports
│   └── [Company]_Master_Intelligence_*.md # Consolidated reports
├── Job_Postings/                      # Job descriptions and search results
│   ├── [Company]_[Role].md           # Manual job postings
│   └── SearchResults_*_*_*.md        # Automated search results with full descriptions
├── OutputResumes/                     # Generated resumes and assessments
│   ├── Assessment_*_*_*.md           # Job fit assessments with scores
│   ├── Comparison_*_*_*.md           # Multi-job comparison reports
│   ├── Interview_Prep_*_*_*.md       # Customized interview questions
│   ├── Step1_Draft_*_*_*.md          # Initial tailored resumes
│   ├── Step2_Provenance_Analysis_*.md # Credibility analysis reports
│   ├── Step3_Final_Resume_*_*_*.md   # Final hardened resumes
│   └── Step4_CoverLetter_*_*_*.md    # Strategic cover letters
├── ResumeSourceFolder/                # Master career inventory (source of truth)
│   ├── Comprehensive_CV_MASTER_COPY_v35.md    # Complete work history
│   └── Comprehensive_CV_Technology_Capability.md # Technical skills inventory
├── Sample_Output/                     # Example outputs for reference
│   ├── *_Fit_Assessment_*.md         # Sample assessments (Poor/Moderate/Great)
│   ├── Briefing_*_*_*.md             # Sample briefing notes
│   ├── Interview_Prep_*_*_*.md       # Sample interview preparations
│   └── Rubric_*_*_*.md               # Sample scoring rubrics
├── Scoring_Rubrics/                   # Reusable assessment rubrics
│   └── Rubric_[Company]_[Role]_[Date].md # Job-specific scoring frameworks
├── SourceMaterial/                    # Methodology and analysis documents
│   ├── System_Dynamics_Analysis_Assessment_First_Hiring_v3.md # Theoretical foundation
│   ├── Critical_Analysis_Model_Assumptions.md # System analysis
│   └── [various development files]    # Internal development documentation
├── CLAUDE.md                          # System instructions for Claude Code
├── JOB_SEARCH_GUIDE.md               # Job search workflow documentation
├── LICENSE                            # License file
├── Master_Resume_Comprehensive_Setup_Guide_v1.md # Master resume methodology
├── README.md                          # This documentation
├── SETUP.md                          # Setup and installation instructions
├── comprehensive_work_history_FAQ.md  # Master resume philosophy FAQ
└── package.json                       # Node.js dependencies (for pandoc)
```

## Detailed Workflow

### 1. Setup Your Master Resume

**The Core Problem**: Traditional 2-page resumes create a fundamental issue when using AI for tailoring. When you try to customize a condensed resume, AI tools must fill gaps with plausible-sounding but fictional content, leading to claims you can't defend in interviews.

**The Solution**: A comprehensive 20+ page master resume serves as your "database of truth" - every achievement, metric, and skill is documented with full context. This allows AI to **select** relevant experiences rather than **invent** them.

**📖 Essential Reading**: Start with [`Master_Resume_Comprehensive_Setup_Guide_v1.md`](Master_Resume_Comprehensive_Setup_Guide_v1.md) for complete methodology and [`comprehensive_work_history_FAQ.md`](comprehensive_work_history_FAQ.md) for quick conceptual explanation.

Place your comprehensive career inventory in `ResumeSourceFolder/`:
- Complete work history with all projects and achievements using HAMZ-Extended framework
- Technology capabilities and skills inventory with proficiency assessments
- Awards, publications, and certifications with full context and validation

**Critical Standard**: Every claim must be defensible in interview settings with specific examples, metrics, and mechanisms documented.

### 2. Find and Add Job Descriptions

**Option A: Search hiring.cafe**
```bash
/searchjobs "your search query" "location" --save
```
Searches hiring.cafe and automatically saves complete verbatim job descriptions to `Job_Postings/` folder.

**Option B: Manual job posting**
Create a markdown file in `Job_Postings/` directory (e.g., `Job_Postings/CompanyName_Role_Date.md`)

### 3. Assess Job Opportunity First

Before investing time in resume creation, assess your fit:

```bash
/assessjob Job_Postings/AltoJobPost.md
```

Generates comprehensive assessment report showing candidacy strength and identifies critical gaps. **Only proceed if assessment shows reasonable fit** (typically 60%+ match).

### 4. Generate Optimized Resume (If Assessment is Positive)

```bash
/buildresume Job_Postings/AltoJobPost.md Canadian
```

Executes the core 3-step resume development process: draft creation, provenance analysis, final hardened resume.

### 5. Compare Multiple Opportunities (If Applicable)

```bash
/comparejobs Assessment_Company1_Role1_Date.md Assessment_Company2_Role2_Date.md
```

### 6. Prepare for Interview

**Create study guide to address gaps:**
```bash
/briefing OutputResumes/Assessment_[Company]_[Role]_[Date].md Job_Postings/AltoJobPost.md gaps-only
```

**Generate customized interview questions:**
```bash
/interviewprep OutputResumes/Step3_Final_Resume_[Role]_[Company]_[Date].md Job_Postings/AltoJobPost.md
```

### 7. Finalize Application (Optional)

**Generate strategic cover letter:**
```bash
/coverletter OutputResumes/Step3_Final_Resume_[Role]_[Company]_[Date].md Job_Postings/AltoJobPost.md
```

**Convert to professional Word format:**
```bash
/install-pandoc    # Install pandoc if needed
/convert all       # Convert all documents to Word format
```

## Core Commands Reference

### Opportunity Assessment
```bash
/assessjob <job-description-file>
```
**START HERE** - Creates dynamic scoring rubric and 100-point assessment with go/no-go recommendation.

### Resume Development
```bash
/buildresume <job-description-file> [cultural-profile]
```
Runs full 3-step optimization (draft, provenance check, final resume). Cultural profiles: Canadian (default), US, European, UK, Australian.

### Modular Assessment Workflow
```bash
/createrubric <job-posting-file>                              # Create reusable scoring rubric only
/assesscandidate <rubric-file> <job-posting-file>            # Assess using pre-created rubric
/comparejobs <assessment-1> <assessment-2> [assessment-3]    # Compare multiple assessments
```

### Interview Preparation
```bash
/briefing <assessment-report> <job-description> [gaps-only]  # Create study guide
/interviewprep <resume-file> <job-description> [N]           # Generate interview questions
```

### Job Search System
```bash
/searchjobs <query> [location] [--company=name] [--save] [--limit=N]
```
Two-phase hybrid search: API search + Playwright scraping for complete verbatim job descriptions. Default 20 jobs, max 50 recommended.

**Examples:**
```bash
/searchjobs "software engineer" "Toronto"
/searchjobs "data analyst" --company=Deloitte --save
/searchjobs "python developer" "Mississauga" --save --limit=30
```

### Company Intelligence
```bash
/osint <company-name>
```
Deploys 6 specialized OSINT agents in parallel for comprehensive intelligence: Corporate, Legal, Leadership, Compensation, Culture, Market.

### Document Management
```bash
/provenance <draft-resume-file>                              # Standalone credibility check
/coverletter <step3-resume> <job-description> [manager]      # Strategic cover letter
/convert <file-path-or-pattern> [output-directory]          # Markdown to Word DOCX
/install-pandoc [force]                                      # Install pandoc utility
```

## The 8-Step Application Process

### Phase 1: Opportunity Assessment (Step 1)
**Job Assessment & Candidate Evaluation**
- Dynamic rubric creation with 100-point scoring framework
- Domain knowledge integration via web research
- Comprehensive evaluation mapping candidate evidence to requirements
- Go/no-go decision based on fit score and gap analysis

### Phase 2: Resume Development (Steps 2-4)
**Step 2: Initial Draft Creation**
- Job analysis with cultural profiling
- HAM-Z methodology application
- ATS optimization

**Step 3: Provenance Analysis**
- Mandatory evidence verification against master documents
- Line number validation with exact quotes
- Critical risk detection for fabricated capabilities
- Risk categorization (Critical/High/Medium/Low)

**Step 4: Final Resume Production**
- Systematic hardening incorporating all provenance recommendations
- Evidence-based claims withstanding interview scrutiny
- Interview-ready final version

### Phase 3: Interview Preparation (Steps 5-6)
**Step 5: Gap Analysis and Study Guide Creation**
- Critical gap identification from assessment
- Learning resource research with web search
- Structured study plan with actionable timelines
- Two modes: gaps-only or comprehensive preparation

**Step 6: Customized Interview Question Generation**
- Resume-job alignment analysis
- Technical, behavioral, and verification questions
- STAR format coaching and answer frameworks
- Follow-up preparation and risk mitigation

### Phase 4: Application Finalization (Steps 7-8)
**Step 7: Cover Letter Generation (Optional)**
- Strategic narrative from validated resume
- Requirements-matching table for visual impact
- Evidence chain to master documents

**Step 8: Document Conversion (Optional)**
- Markdown to Word DOCX conversion
- Pandoc integration for professional formatting
- Batch processing support

## HAM-Z™ Methodology

Every achievement follows this formula:

**Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]**

Example:
> Achieved 50% reduction in lease processing time (60→30 days) by leveraging VTS platform expertise to implement automated workflow system across 150-property portfolio

## Provenance Risk Categories

### Critical Risk (Cannot Proceed)
- Fabricated capabilities, skills, or experience not documented in master materials
- Language skills without explicit documentation
- Industry or geographic experience without supporting evidence

### High Risk (Must Fix)
- Unbounded metrics without timeframes
- Unsupported superlatives ("market-leading", "transformational")
- Cross-document inconsistencies
- Benchmark claims without peer sets

### Medium Risk (Should Fix)
- Results without mechanism ("how")
- Duplicate achievements across roles
- Tools listed without outcomes

### Low Risk (Consider Fixing)
- Ambiguous scope or geography
- Confidentiality concerns
- Minor formatting inconsistencies

## Example Workflows

### Multi-Job Comparison Workflow
```bash
# Option 1: Complete assessment workflow
/assessjob JLL_VP_Office_Leasing.md
/assessjob Canerector_Vice_President_Real_Estate.md
/assessjob CityOfToronto_SeniorDirector.md
/comparejobs Assessment_JLL*.md Assessment_Canerector*.md Assessment_CityOfToronto*.md

# Option 2: Modular assessment workflow
/createrubric JLL_VP_Office_Leasing.md
/createrubric Canerector_Vice_President_Real_Estate.md
/assesscandidate Rubric_JLL*.md JLL_VP_Office_Leasing.md
/assesscandidate Rubric_Canerector*.md Canerector_Vice_President_Real_Estate.md
/comparejobs Assessment_JLL*.md Assessment_Canerector*.md
```

### Job Search and Intelligence Workflow
```bash
# Complete workflow from job discovery to application
/searchjobs "real estate director" "Toronto" --save --limit=20
/osint Deloitte
/assessjob Job_Postings/SearchResults_Leasing_Toronto_2025-09-29.md
/buildresume Job_Postings/SearchResults_Leasing_Toronto_2025-09-29.md Canadian
/briefing Assessment_Deloitte*.md Job_Postings/SearchResults*.md gaps-only
/interviewprep OutputResumes/Step3_Final_Resume*.md Job_Postings/SearchResults*.md
```

## File Naming Conventions

- **Assessment**: `Assessment_[Company]_[Role]_[Date].md` and `Rubric_[Company]_[Role]_[Date].md`
- **Resume Development**: `Step1_Draft_[Role]_[Company]_[Date].md`, `Step2_Provenance_Analysis_[Role]_[Company]_[Date].md`, `Step3_Final_Resume_[Role]_[Company]_[Date].md`
- **Interview Prep**: `Briefing_[Company]_[Role]_[Mode]_[Date].md`, `Interview_Prep_[Company]_[Role]_[Date].md`
- **Application**: `Step4_CoverLetter_[Role]_[Company]_[Date].md`, `[OriginalName].docx`

## Requirements

- Claude Code CLI installed and configured
- Master resume documents in markdown format in `ResumeSourceFolder/`
- Job descriptions in markdown format in `Job_Postings/`
- `pandoc` installed for document conversion (optional, use `/install-pandoc`)

## Support

For issues or improvements:
- `comprehensive_work_history_FAQ.md` - Master resume philosophy
- `CLAUDE.md` - Technical implementation details
- `SourceMaterial/` - Methodology documentation
- `.claude/agents/` - Agent-specific logic

## License

Private repository - All resume content is confidential
