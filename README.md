# JobOps - Intelligence-Driven Application Warfare System

<p align="center">
  <img src="Images/JobOps_logo.png" alt="JobOps Logo" width="400">
</p>

**Version 1.1.1** | [Changelog](CHANGELOG.md) | [Why I Built This](Why_I_Built_This.md)

**Wage war on unemployment.** A tactical job application platform that transforms career inventories into winning applications through systematic opportunity assessment, credible resume development, and strategic interview preparation using an 8-step intelligence-driven methodology.

> *"The system is broken. And everyone knows it. But nobody's fixing it."* — Read the full story in [Why_I_Built_This.md](Why_I_Built_This.md)

## Rapid Deployment Protocol

1. Install dependencies: `npm run install-all`
2. **New users**: Run `/create-career-history <your-resume.pdf>` to parse your existing resume and auto-populate your master career inventory with HAM-Z-enhanced content
3. Review and enhance the generated `ResumeSourceFolder/` files, then add target roles to `Job_Postings/`
4. Start the Playwright MCP server (see "Run the Playwright MCP server" in `SETUP.md`) so Claude Code can reach browser automation
5. Launch Claude Code in this repository and run slash commands such as `/assessjob Job_Postings/Example.md` followed by `/buildresume`

## Architecture Overview

- **Claude Code CLI** orchestrates the workflow: users issue slash commands that call specialized agents under `.claude/`
- **Playwright MCP server** provides headless browser automation for job sourcing and intelligence gathering
- **Content directories** (`ResumeSourceFolder/`, `Job_Postings/`, `OutputResumes/`, `Briefing_Notes/`) hold source materials, generated analyses, and deliverables
- **HAM-Z agents** execute each step: resume drafting, provenance hardening, and interview preparation
- **Scoring and OSINT resources** in `Scoring_Rubrics/` and `Intelligence_Reports/` feed assessments and research deliverables

## Tactical Capabilities

JobOps deploys the HAM-Z™ methodology to create powerful, defensible strike packages that are:
- **Strategically positioned** for specific roles
- **Culturally attuned** to regional preferences
- **Provenance-hardened** to withstand scrutiny
- **ATS-optimized** for applicant tracking systems

Deploy a tactical 8-step offensive organized into three phases: strike package assembly (Steps 1-3), interview preparation (Steps 4-6), and mission execution (Steps 7-8), ensuring every claim is verified and every achievement is weaponized with mechanism, metrics, and timeframe.

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

This theoretical foundation explains why JobOps begins with target reconnaissance (`/assessjob`) rather than immediate resume creation - preventing wasted effort on fundamentally poor targets while ensuring only high-probability opportunities receive full tactical deployment.

## Weapons Systems

- **HAM-Z™ Methodology**: Combines Hard Skills, Actions, Metrics, and Structure with XYZ narrative format to create precision-targeted achievements
- **Provenance Hardening**: Comprehensive credibility verification to identify and eliminate fabricated claims
- **Cultural Profiles**: Adapts resume style to regional theater expectations (Canadian, US, European, etc.)
- **Arsenal Inventory**: Maintains comprehensive career database separate from mission-specific deployments
- **Dynamic Target Scoring**: Creates custom reconnaissance rubrics for each opportunity
- **Interview Preparation**: Gap analysis intelligence briefings and tactical rehearsal question generation
- **Distributed OSINT Intelligence**: Professional-grade company reconnaissance with parallel specialized agents
- **Target Acquisition System**: Hybrid search combining API reconnaissance + Playwright deep-scan for complete intelligence

## Arsenal Organization

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
│   ├── commands/                      # 16 slash command definitions
│   │   ├── assesscandidate.md        # /assesscandidate - Use pre-created rubric
│   │   ├── assessjob.md              # /assessjob - Complete assessment
│   │   ├── briefing.md               # /briefing - Interview prep guide
│   │   ├── buildresume.md            # /buildresume - 3-step resume process
│   │   ├── change-one-thing.md       # /change-one-thing - Career retrospective
│   │   ├── comparejobs.md            # /comparejobs - Multi-job analysis
│   │   ├── convert.md                # /convert - Markdown to DOCX
│   │   ├── coverletter.md            # /coverletter - Generate letter
│   │   ├── create-career-history.md  # /create-career-history - Setup wizard
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
├── Guides/                            # Templates and metadata standards
│   ├── CareerHighlight_Template.md    # CareerHighlights file template
│   ├── Experience_Template.md         # Experience file template
│   ├── TechnologyMatrix_Template.md   # Technology capability template
│   ├── Metadata_Standards.md          # YAML frontmatter standards
│   ├── Maintenance_Checklist.md       # Master resume maintenance guide
│   └── Migration_Tracker.md           # Migration status tracker
├── Intelligence_Reports/              # OSINT company intelligence reports
│   ├── [Company]_*_Intelligence_*.md  # Specialized intelligence reports
│   └── [Company]_Master_Intelligence_*.md # Consolidated reports
├── Job_Postings/                      # Job descriptions and search results
│   ├── .template.md                   # Standard job posting template with YAML metadata
│   ├── [Company]_[Role]_YYYY-MM-DD.md # Manual job postings (new naming convention)
│   └── SearchResults_*_*_*.md        # Automated search results with full descriptions
├── OutputResumes/                     # Generated resumes and assessments (v1.1.1+ uses timestamped folders)
│   ├── YYYY-MM-DD_HHMMSS_[Company]_[Role]/  # Timestamped assessment folder (new in v1.1.1)
│   │   ├── Assessment_[Company]_[Role]_[Date].md # Job fit assessment with scores
│   │   └── Rubric_[Company]_[Role]_[Date].md     # Scoring rubric (audit trail copy)
│   ├── Change_One_Thing_Analysis_Part*_*.md # Career retrospective analysis
│   ├── Change_One_Thing_EXECUTIVE_SUMMARY_*.md # Analysis executive summary
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

## Tactical Operations Workflow

### 1. Build Your Arsenal Inventory

**The Core Problem**: Traditional 2-page resumes create a critical shortage when deploying AI for tailoring. When you attempt to customize condensed resumes, AI tools must fabricate plausible-sounding but fictional claims, leading to failures under interview scrutiny.

**The Solution**: A comprehensive 20+ page arsenal inventory serves as your "full career database" - every achievement, metric, and skill is documented with complete context. This allows AI to **deploy** verified content rather than **fabricate** it.

**📖 Essential Reading**: Start with [`Master_Resume_Comprehensive_Setup_Guide_v1.md`](Master_Resume_Comprehensive_Setup_Guide_v1.md) for complete methodology and [`comprehensive_work_history_FAQ.md`](comprehensive_work_history_FAQ.md) for quick conceptual explanation.

**🚀 Quick Start for New Users**: If your `ResumeSourceFolder/` is empty, run the intelligent parser with your existing resume:
```bash
/create-career-history path/to/your/resume.pdf
# Or with multiple files:
/create-career-history resume.docx linkedin_export.pdf certifications.txt
```
This parses your resume(s), extracts structured career data, transforms content to HAM-Z format, and creates pre-populated master career inventory files using standardized templates from `Guides/` directory. Provides gap analysis and enhancement recommendations.

Place your comprehensive career inventory in `ResumeSourceFolder/`:
- Complete work history with all projects and achievements using HAMZ-Extended framework
- Technology capabilities and skills inventory with proficiency assessments
- Awards, publications, and certifications with full context and validation

**Critical Standard**: Every claim must withstand interview interrogation with verified examples, metrics, and mechanisms documented.

### 2. Acquire Target Coordinates

**Option A: Deploy target acquisition sweep**
```bash
/searchjobs "your search query" "location" --save
```
Conducts reconnaissance sweep of hiring.cafe and automatically extracts complete verbatim target intelligence to `Job_Postings/` folder.

**Option B: Manual target upload**
Create a markdown file in `Job_Postings/` directory (e.g., `Job_Postings/CompanyName_Role_Date.md`)

### 3. Conduct Target Reconnaissance

Before committing strike resources, conduct reconnaissance on target viability:

```bash
/assessjob Job_Postings/AltoJobPost.md
```

Generates comprehensive reconnaissance report showing engagement probability and identifies critical capability gaps. **Only deploy full strike package if reconnaissance shows viable engagement** (typically 60%+ hit probability).

### 4. Assemble Strike Package (If Reconnaissance is Positive)

```bash
/buildresume Job_Postings/AltoJobPost.md Canadian
```

Executes the core 3-step resume assembly: initial draft creation, provenance verification, final hardened resume.

### 5. Prioritize Multiple Targets (If Applicable)

```bash
/comparejobs Assessment_Company1_Role1_Date.md Assessment_Company2_Role2_Date.md
```

### 6. Conduct Interview Preparation

**Generate intelligence brief to address capability gaps:**
```bash
/briefing OutputResumes/Assessment_[Company]_[Role]_[Date].md Job_Postings/AltoJobPost.md gaps-only
```

**Deploy interview prep questions:**
```bash
/interviewprep OutputResumes/Step3_Final_Resume_[Role]_[Company]_[Date].md Job_Postings/AltoJobPost.md
```

### 7. Execute Final Mission Sequence (Optional)

**Deploy strategic cover letter:**
```bash
/coverletter OutputResumes/Step3_Final_Resume_[Role]_[Company]_[Date].md Job_Postings/AltoJobPost.md
```

**Convert to deployment-ready format:**
```bash
/install-pandoc    # Install pandoc if needed
/convert all       # Convert all documents to Word format
```

## Command Arsenal

### Target Reconnaissance
```bash
/assessjob <job-description-file>
```
**START HERE** - Deploys dynamic reconnaissance rubric and 100-point viability assessment with engage/disengage recommendation.

### Resume Assembly
```bash
/buildresume <job-description-file> [cultural-profile]
```
Executes full 3-step assembly protocol (draft creation, provenance verification, hardened final resume). Theater profiles: Canadian (default), US, European, UK, Australian.

### Modular Reconnaissance Workflow
```bash
/createrubric <job-posting-file>                              # Establish reusable targeting criteria
/assesscandidate <rubric-file> <job-posting-file>            # Execute standardized reconnaissance
/comparejobs <assessment-1> <assessment-2> [assessment-3]    # Prioritize multiple targets
```

### Interview Preparation
```bash
/briefing <assessment-report> <job-description> [gaps-only]  # Generate intelligence brief
/interviewprep <resume-file> <job-description> [N]           # Deploy interview prep questions
```

### Target Acquisition System
```bash
/searchjobs <query> [location] [--company=name] [--save] [--limit=N]
```
Two-phase hybrid reconnaissance: API sweep + Playwright deep-scan for complete verbatim target intelligence. Default 20 targets, max 50 recommended.

**Examples:**
```bash
/searchjobs "software engineer" "Toronto"
/searchjobs "data analyst" --company=Deloitte --save
/searchjobs "python developer" "Mississauga" --save --limit=30
```

### Intelligence Operations
```bash
/osint <company-name>
```
Deploys 6 specialized intelligence agents in parallel for comprehensive target analysis: Corporate, Legal, Leadership, Compensation, Culture, Market.

### Mission Support
```bash
/provenance <draft-resume-file>                              # Provenance verification sweep
/coverletter <step3-resume> <job-description> [manager]      # Deploy cover letter
/convert <file-path-or-pattern> [output-directory]          # Convert to deployment format
/install-pandoc [force]                                      # Install conversion utility
```

### Career Strategy & Analysis
```bash
/change-one-thing <resume-folder>
```
Conducts comprehensive 16-lens career retrospective analysis to identify the single highest-leverage, non-obvious change you could have made at a specific inflection point. Generates 3-4 detailed analysis parts plus executive summary with:
- External patterns: Skills, network, technical decisions, visibility, cross-pollination, geographic positioning
- Internal constraints: Family, financial capacity, health/energy analysis
- Temporal mechanics: Hindsight bias correction, butterfly effects, irreversibility
- Probabilistic outcomes: Luck vs. skill, values evolution, market timing
- Multi-candidate recommendations optimized for wealth/optionality/impact/balance
- Quantified opportunity costs and risk-adjusted returns
- Actionable "what you can do now" recovery plans with immediate next steps

## The 8-Step Protocol

### Phase 1: Target Reconnaissance (Step 1)
**Target Assessment & Engagement Viability**
- Dynamic targeting criteria with 100-point probability framework
- Domain intelligence integration via web research
- Comprehensive capability mapping against target requirements
- Engage/disengage decision based on hit probability and capability gaps

### Phase 2: Resume Development (Steps 2-4)
**Step 2: Initial Draft Creation**
- Target analysis with theater-specific profiling
- HAM-Z methodology deployment
- ATS penetration optimization

**Step 3: Provenance Verification**
- Mandatory credibility verification against arsenal inventory
- Line-by-line validation with exact evidence citations
- Critical risk detection for fabricated capabilities
- Risk categorization (Critical/High/Medium/Low)

**Step 4: Hardened Resume Production**
- Systematic hardening incorporating all verification recommendations
- Evidence-based claims withstanding interview interrogation
- Deployment-ready final resume

### Phase 3: Interview Preparation (Steps 5-6)
**Step 5: Gap Analysis and Intelligence Brief Creation**
- Critical capability gap identification from reconnaissance
- Learning resource research with web intelligence
- Structured training protocol with actionable timelines
- Two modes: gaps-only or comprehensive interview prep

**Step 6: Interview Question Generation**
- Resume-job alignment analysis
- Technical, behavioral, and verification question deployment
- STAR format tactical coaching and response frameworks
- Follow-up preparation and defensive risk mitigation

### Phase 4: Mission Execution (Steps 7-8)
**Step 7: Cover Letter Deployment (Optional)**
- Strategic communication from verified resume
- Requirements-matching table for visual impact
- Complete evidence chain to arsenal inventory

**Step 8: Format Conversion (Optional)**
- Markdown to Word DOCX conversion
- Pandoc integration for professional formatting
- Batch deployment support

## HAM-Z™ Methodology

Every achievement follows this formula:

**Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]**

Example:
> Achieved 50% reduction in lease processing time (60→30 days) by leveraging VTS platform expertise to implement automated workflow system across 150-property portfolio

## Provenance Risk Categories

### Critical Risk (Mission Abort)
- Fabricated capabilities, skills, or experience not documented in arsenal inventory
- Language skills without explicit documentation
- Industry or geographic experience without supporting evidence

### High Risk (Must Fix)
- Unbounded metrics without timeframes
- Unsupported superlatives ("market-leading", "transformational")
- Cross-document inconsistencies
- Benchmark claims without defined peer sets

### Medium Risk (Should Fix)
- Results without mechanism ("how")
- Duplicate achievements across roles
- Tools listed without outcomes

### Low Risk (Consider Fixing)
- Ambiguous scope or geography
- Confidentiality concerns
- Minor formatting inconsistencies

## Tactical Operations Examples

### Multi-Target Prioritization Workflow
```bash
# Option 1: Complete reconnaissance workflow
/assessjob JLL_VP_Office_Leasing.md
/assessjob Canerector_Vice_President_Real_Estate.md
/assessjob CityOfToronto_SeniorDirector.md
/comparejobs Assessment_JLL*.md Assessment_Canerector*.md Assessment_CityOfToronto*.md

# Option 2: Modular reconnaissance workflow
/createrubric JLL_VP_Office_Leasing.md
/createrubric Canerector_Vice_President_Real_Estate.md
/assesscandidate Rubric_JLL*.md JLL_VP_Office_Leasing.md
/assesscandidate Rubric_Canerector*.md Canerector_Vice_President_Real_Estate.md
/comparejobs Assessment_JLL*.md Assessment_Canerector*.md
```

### Full Deployment Workflow
```bash
# Complete operations from target acquisition to mission execution
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
- **Career Analysis**: `Change_One_Thing_Analysis_Part[1-3]_YYYYMMDD.md`, `Change_One_Thing_EXECUTIVE_SUMMARY_YYYYMMDD.md`

## Equipment & Readiness

- Claude Code CLI installed and configured
- Arsenal inventory documents in markdown format in `ResumeSourceFolder/`
- Target coordinates in markdown format in `Job_Postings/`
- `pandoc` installed for document conversion (optional, use `/install-pandoc`)

### Optional: Tactical Statusline Configuration

Add JobOps branding to your Claude Code statusline (local repository only):

Edit `.claude/settings.local.json` and add:

```json
{
  "statusLine": {
    "type": "command",
    "command": "bash",
    "args": ["-c", "jq -r '\"⚔️ JobOps | \" + (.model.display_name // \"unknown\") + \" | 📁 \" + (.workspace.current_dir | split(\"/\") | last)'"],
    "padding": 1
  }
}
```

This displays: **`⚔️ JobOps | [Model] | 📁 [Directory]`** at the bottom of your terminal during Claude Code sessions.

## Mission Support

For tactical assistance:
- `comprehensive_work_history_FAQ.md` - Arsenal inventory philosophy
- `CLAUDE.md` - Technical implementation details
- `SourceMaterial/` - Methodology documentation
- `.claude/agents/` - Agent-specific operational protocols

## License

Private repository - All tactical operations are confidential
