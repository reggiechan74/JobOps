# Resume Optimizer for Claude Code

A comprehensive job application system that transforms career inventories into targeted resumes and provides complete interview preparation using an 8-step methodology with provenance hardening.

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
- **Risk Assessment**: Categorizes claims by defensibility (High/Medium/Low risk)
- **Master Resume System**: Maintains comprehensive career inventory separate from targeted resumes
- **Dynamic Job Scoring**: Creates custom assessment rubrics for each job posting
- **Candidate Assessment**: Expert HR evaluation with domain-specific knowledge
- **Interview Preparation**: Gap analysis briefings and customized question generation
- **Strategic Coaching**: Comprehensive study guides with learning resources
- **Distributed OSINT Intelligence**: Professional-grade company research with parallel specialized agents

## Directory Structure

```
resumeoptimizer/
├── .claude/                           # Claude Code configuration
│   ├── agents/                        # Specialized processing agents (13 agents)
│   │   ├── candidate-assessment.md           # HR assessment with domain expertise
│   │   ├── interview-briefing.md             # Gap analysis and study guides
│   │   ├── interview-question-generator.md   # Customized interview questions
│   │   ├── osint-agent.md                    # Master OSINT orchestrator and unified analysis
│   │   ├── osint-compensation.md             # Salary benchmarking and total rewards analysis
│   │   ├── osint-corporate.md                # Corporate structure and financial intelligence
│   │   ├── osint-culture.md                  # Employee sentiment and workplace culture
│   │   ├── osint-leadership.md               # Executive backgrounds and leadership analysis
│   │   ├── osint-legal.md                    # Litigation history and regulatory compliance
│   │   ├── osint-market.md                   # Industry analysis and competitive positioning
│   │   ├── step1-resume-draft.md             # Initial targeted resume creation
│   │   ├── step2-provenance-check.md         # Enhanced credibility analysis
│   │   ├── step3-final-resume.md             # Final hardened resume production
│   │   └── step4-cover-letter.md             # Strategic cover letter generation
│   ├── commands/                      # Slash command definitions (11 commands)
│   │   ├── assessjob.md               # Candidate assessment with dynamic rubrics
│   │   ├── briefing.md                # Interview preparation briefing
│   │   ├── buildresume.md             # Complete 3-step resume build
│   │   ├── comparejobs.md             # Multi-assessment comparison and analysis
│   │   ├── convert.md                 # Markdown to Word conversion
│   │   ├── coverletter.md             # Cover letter generation
│   │   ├── install-pandoc.md          # Pandoc installation utility
│   │   ├── interviewprep.md           # Customized interview questions
│   │   ├── osint.md                   # Distributed OSINT intelligence orchestration
│   │   ├── prime.md                   # Load documentation into context
│   │   └── provenance.md              # Standalone provenance analysis
│   └── settings.local.json            # Local Claude Code settings
├── Briefing_Notes/                    # Interview preparation materials
│   └── .gitkeep                       # Directory placeholder
├── Intelligence_Reports/              # OSINT intelligence reports with citations
│   ├── [Company]_Corporate_Intelligence_[Date].md    # Corporate structure and financials
│   ├── [Company]_Legal_Intelligence_[Date].md        # Litigation history and compliance
│   ├── [Company]_Leadership_Intelligence_[Date].md   # Executive backgrounds and analysis
│   ├── [Company]_Compensation_Intelligence_[Date].md # Salary benchmarking and total rewards
│   ├── [Company]_Culture_Intelligence_[Date].md      # Employee sentiment and workplace culture
│   ├── [Company]_Market_Intelligence_[Date].md       # Industry analysis and competitive positioning
│   └── [Company]_Master_Intelligence_[Date].md       # Comprehensive integrated analysis
├── Job_Postings/                      # Target job descriptions
│   ├── .gitkeep                       # Directory placeholder
│   ├── AltoJobPost.md                 # Sample job posting
│   ├── City_of_Toronto_Director.md    # City of Toronto role
│   └── JLL_VP_Office_Leasing.md       # JLL VP position
├── OutputResumes/                     # Generated resumes and analysis
│   ├── .gitkeep                       # Directory placeholder
│   ├── Assessment_*.md                # Candidate assessment reports
│   ├── Comparison_*.md                # Multi-job comparison analyses
│   ├── Step1_Draft_*.md               # Initial targeted drafts
│   ├── Step2_Provenance_Analysis_*.md # Credibility analysis reports
│   ├── Step3_Final_Resume_*.md        # Interview-ready final resumes
│   ├── Step4_CoverLetter_*.md         # Strategic cover letters
│   └── *.pdf                          # Converted documents
├── ResumeSourceFolder/                # Master career inventory
│   ├── .gitkeep                       # Directory placeholder
│   ├── Comprehensive_CV_MASTER_COPY_v35.md       # Complete work history
│   └── Comprehensive_CV_Technology_Capability.md # Technical skills inventory
├── Sample_Output/                     # Example outputs for demonstration
│   ├── Assessment_Building_Ontario_Fund_*.md     # Sample assessment report
│   ├── Briefing_Building_Ontario_Fund_*.md       # Sample briefing note
│   ├── Building_Ontario_Fund.md                  # Sample job posting
│   └── Rubric_Building_Ontario_Fund_*.md         # Sample scoring rubric
├── Scoring_Rubrics/                   # Assessment rubrics and frameworks
│   └── Rubric_*_*_*.md               # Dynamic job-specific rubrics
├── SourceMaterial/                    # Reserved for methodology guides
│   └── .gitkeep                       # Directory placeholder
├── .gitignore                         # Git ignore rules
├── CLAUDE.md                          # System instructions for Claude Code
├── Master_Resume_Comprehensive_Setup_Guide_v1.md  # Ultra-detailed master resume creation guide
├── README.md                          # This documentation
├── comprehensive_work_history_FAQ.md  # Master resume philosophy
└── LICENSE                            # Repository license
```

## Quick Start

### 1. Setup Your Master Resume: The Anti-Hallucination Foundation

**The Core Problem**: Traditional 2-page resumes create a fundamental issue when using AI for tailoring. When you try to customize a condensed resume, AI tools must fill gaps with plausible-sounding but fictional content, leading to claims you can't defend in interviews.

**The Solution**: A comprehensive 20+ page master resume serves as your "database of truth" - every achievement, metric, and skill is documented with full context. This allows AI to **select** relevant experiences rather than **invent** them.

#### Why This Approach Works

**Key Benefits**:
- **No Hallucinations**: Enhanced provenance system prevents AI from inventing capabilities
- **Interview Confidence**: Every bullet point traces to documented experience with line-by-line verification
- **Provenance Verification**: All claims verified against source material with mandatory quote-checking
- **Career Pattern Recognition**: Discover valuable themes in your experience
- **Evidence-Based Targeting**: Only documented skills and achievements used in tailoring

#### Implementation Guide

**📖 Essential Reading**: Start with [`Master_Resume_Comprehensive_Setup_Guide_v1.md`](Master_Resume_Comprehensive_Setup_Guide_v1.md) - an ultra-detailed guide covering the complete methodology for building your career database.

Place your comprehensive career inventory in `ResumeSourceFolder/`:
- Complete work history with all projects and achievements using HAMZ-Extended framework
- Technology capabilities and skills inventory with proficiency assessments
- Awards, publications, and certifications with full context and validation
- Academic credentials including curriculum details and competency documentation

**Critical Standard**: Every claim must be defensible in interview settings with specific examples, metrics, and mechanisms documented.

#### Documentation Resources
📖 **Complete Setup Guide**: [`Master_Resume_Comprehensive_Setup_Guide_v1.md`](Master_Resume_Comprehensive_Setup_Guide_v1.md) - Ultra-detailed 75+ page methodology
📖 **Philosophy Overview**: [`comprehensive_work_history_FAQ.md`](comprehensive_work_history_FAQ.md) - Quick conceptual explanation

### 2. Add Job Description

Create a markdown file with the target job description in the root directory (e.g., `AltoJobPost.md`)

### 3. Assess Job Opportunity First

Before investing time in resume creation, assess your fit for the position:

```
/assessjob AltoJobPost.md
```

This generates a comprehensive assessment report showing your candidacy strength and identifies critical gaps. **Only proceed if the assessment shows reasonable fit** (typically 60%+ match).

### 4. Generate Optimized Resume (If Assessment is Positive)

Use the slash command to run the complete resume development process:

```
/buildresume AltoJobPost.md Canadian
```

This executes the core resume development process:
1. Creates initial targeted draft
2. Performs provenance analysis
3. Produces final hardened resume

### 5. Compare Multiple Opportunities (If Applicable)

**Compare assessments to prioritize opportunities:**
```
/comparejobs Assessment_Company1_Role1_Date.md Assessment_Company2_Role2_Date.md
```

### 6. Prepare for Interview

**Create study guide to address gaps:**
```
/briefing OutputResumes/Assessment_[Company]_[Role]_[Date].md AltoJobPost.md gaps-only
```

**Generate customized interview questions:**
```
/interviewprep OutputResumes/Step3_Final_Resume_[Role]_[Company]_[Date].md AltoJobPost.md
```

### 7. Finalize Application (Optional)

**Generate strategic cover letter:**
```
/coverletter OutputResumes/Step3_Final_Resume_[Role]_[Company]_[Date].md AltoJobPost.md
```

**Convert to professional Word format:**
```
/install-pandoc    # Install pandoc if needed
/convert all       # Convert all documents to Word format
```

## Available Commands

### Step 1: Opportunity Assessment
```
/assessjob <job-description-file>
```
- **START HERE**: Assess your fit before creating resume
- Creates dynamic job-specific scoring rubric from job posting
- Generates 100-point assessment with evidence
- Provides go/no-go recommendation based on candidacy strength
- Only proceed to resume creation if assessment shows reasonable fit (60%+)

### Steps 2-4: Resume Development (Use Only After Positive Assessment)
```
/buildresume <job-description-file> [cultural-profile]
```
- Runs the full 3-step optimization process
- Cultural profiles: Canadian (default), US, European, UK, Australian
- Output saved to `OutputResumes/` with timestamps

### Provenance Check Only
```
/provenance <draft-resume-file>
```
- Analyzes existing resume draft for credibility issues
- Cross-references against master resume documents
- Generates detailed risk assessment report

### Cover Letter Generation
```
/coverletter <step3-resume-file> <job-description-file> [hiring-manager-name]
```
- Creates strategic cover letter from Step 3 validated resume
- Includes requirements-matching table for visual impact
- Maps critical job requirements to proven experience
- Maintains evidence chain from master documents

### Document Conversion
```
/convert <file-path-or-pattern> [output-directory]
```
- Converts markdown documents to professional Word DOCX format
- Supports single files, patterns, or type-based conversion
- Uses pandoc for high-quality professional formatting
- Examples: `/convert resume`, `/convert "Step3_*.md"`, `/convert all`

### Pandoc Installation
```
/install-pandoc [force]
```
- Automatically installs pandoc for your operating system
- Detects platform and uses appropriate package manager
- Verifies installation and conversion capabilities
- Required for Step 5 document conversion functionality

### Interview Preparation Commands

#### Candidate Assessment
```
/assessjob <job-description-file>
```
- Creates dynamic job-specific scoring rubric from job posting
- Performs web research for domain expertise and industry standards
- Generates comprehensive 100-point assessment with evidence mapping
- Saves custom rubric to `Scoring_Rubrics/` for reuse and audit trail
- Provides hiring recommendations, interview strategies, and gap analysis
- Outputs both rubric and assessment report with full traceability

#### Multi-Job Comparison Analysis
```
/comparejobs <assessment-file-1> <assessment-file-2> [assessment-file-3] [assessment-file-4]
```
- Compare 2-4 assessment files from `OutputResumes/` folder for strategic insights
- Analyzes performance patterns across different roles and companies
- Provides role prioritization and optimal career strategy recommendations
- Identifies transferable skills and consistent strengths/gaps across opportunities
- Generates comprehensive comparison report with negotiation positioning advice
- Supports strategic career planning and multi-offer decision making

#### Gap Analysis Briefing
```
/briefing <assessment-report> <job-description> [gaps-only]
```
- Analyzes assessment results to identify critical skill gaps
- Researches current best practices and learning resources
- Creates detailed study guide with actionable timelines
- Two modes: gaps-only (focus on weaknesses) or comprehensive preparation
- Includes hands-on exercises, quick references, and interview strategies
- Output saved to `Briefing_Notes/Briefing_[Company]_[Role]_[Mode]_[Date].md`

#### Interview Questions Generator
```
/interviewprep <resume-file> <job-description> [number-of-questions]
```
- Generates customized interview questions based on resume-job alignment
- Defaults to 10 questions if number not specified
- Balances technical and behavioral questions using STAR format
- Provides strategic answer guidance and follow-up question preparation
- Includes red flags to avoid and strength positioning strategies
- Output saved to `Briefing_Notes/Interview_Prep_[Company]_[Role]_[Date].md`

### Distributed OSINT Intelligence System

#### Comprehensive Company Intelligence
```
/osint <company-name>
```
- **Professional-grade intelligence gathering** using distributed system of 6 specialized agents
- **Parallel processing** for maximum efficiency and comprehensive coverage
- **Master Intelligence Report** with integrated findings and strategic recommendations

#### OSINT Agent Specializations:
- **Corporate Intelligence** (osint-corporate): Company structure, financials, strategic positioning
- **Legal Intelligence** (osint-legal): Litigation history, regulatory compliance, legal risks
- **Leadership Intelligence** (osint-leadership): Executive backgrounds, leadership analysis
- **Compensation Intelligence** (osint-compensation): Salary benchmarking, benefits, total rewards
- **Culture Intelligence** (osint-culture): Employee sentiment, workplace culture analysis
- **Market Intelligence** (osint-market): Industry analysis, competitive positioning

#### Key Features:
- **Multi-Jurisdictional Legal Research**: CanLII, PACER, BAILII, and international legal databases
- **Comprehensive Compensation Analysis**: Glassdoor, Levels.fyi, PayScale, Blind, and equity analysis
- **Employee Sentiment Tracking**: Anonymous forums, review platforms, and culture assessment
- **Executive Background Verification**: Professional networks, board memberships, reputation analysis
- **Industry Competitive Intelligence**: Market positioning, disruption risks, growth opportunities
- **Corporate Financial Health**: Revenue trends, debt analysis, strategic partnerships
- **Rigorous Source Attribution**: Every claim includes citations with URLs and reliability ratings
- **Intelligence Classification**: All information marked as [VERIFIED], [INFERRED], [CALCULATED], or [ASSUMPTION]
- **Confidence Indicators**: HIGH/MEDIUM/LOW confidence levels for all major findings
- **Professional Citation Standards**: Numbered footnotes with detailed source information and access dates

#### Use Cases:
- **Pre-Interview Intelligence**: Comprehensive company research before job interviews
- **Job Offer Evaluation**: Due diligence before accepting employment offers
- **Negotiation Preparation**: Compensation benchmarking and leverage analysis
- **Business Partnership Assessment**: Risk evaluation for strategic partnerships
- **Investment Due Diligence**: Comprehensive company assessment for investment decisions
- **Competitive Intelligence**: Market positioning and competitor analysis

#### Intelligence Reporting Standards:
All OSINT reports follow rigorous intelligence community standards with complete source attribution:

- **Citation Requirements**: Every factual claim includes inline citations¹ with source URLs where available
- **Information Classification System**:
  - `[VERIFIED]` - Confirmed from primary sources with direct URLs and verification
  - `[INFERRED]` - Logical conclusions drawn from available data with reasoning explained
  - `[CALCULATED]` - Mathematical derivations from source data with methodology disclosed
  - `[ASSUMPTION]` - Reasonable assumptions with stated reasoning and limitations
- **Confidence Indicators**: All findings labeled as HIGH CONFIDENCE, MEDIUM CONFIDENCE, or LOW CONFIDENCE
- **Source Quality Ratings**: Sources rated as (HIGH RELIABILITY), (MEDIUM RELIABILITY), or (LOW RELIABILITY)
- **Professional Footnotes**: Numbered footnote system with complete source information, publication dates, and access dates
- **Methodology Transparency**: Calculation methods, sampling limitations, and analytical frameworks disclosed
- **Intelligence Gaps**: Explicit identification of information unavailable or requiring additional research

#### Sample Intelligence Citation Format:
```
Company revenue of $2.5 billion¹ [VERIFIED - HIGH CONFIDENCE]
Market position suggests strong competitive advantage² [INFERRED - MEDIUM CONFIDENCE]
Employee productivity calculated at $425,000 per employee³ [CALCULATED - HIGH CONFIDENCE]
Estimated growth rate of 15% based on industry standards⁴ [ASSUMPTION - LOW CONFIDENCE]

¹ Example Corp, "2024 Annual Report", March 15, 2025, https://investor.example.com/annual-report-2024, (HIGH RELIABILITY) - Accessed September 26, 2025
² TechAnalyst Weekly, "Market Analysis", Sept 10, 2025, https://techanalyst.com/analysis, (MEDIUM RELIABILITY) - Accessed September 26, 2025
³ Calculation methodology: Revenue ($2.5B) ÷ Employees (5,882) = $425,000 per employee
⁴ Industry benchmark source with stated assumptions and limitations
```

## The Complete 8-Step Application Process

### Phase 1: Opportunity Assessment (Step 1)

#### Step 1: Job Assessment & Candidate Evaluation
- **Dynamic Rubric Creation**: Generates job-specific 100-point scoring framework from job posting
- **Domain Knowledge Integration**: Performs web research for industry standards and role expectations
- **Comprehensive Evaluation**: Maps candidate evidence to each job requirement systematically
- **Risk Analysis**: Identifies gaps and provides mitigation strategies for hiring decisions
- **Go/No-Go Decision**: Determines if opportunity is worth pursuing based on fit score and gap analysis
- **Outputs**:
  - `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md` (custom scoring criteria)
  - `OutputResumes/Assessment_[Company]_[Role]_[Date].md` (detailed evaluation report)

### Phase 2: Resume Development (Steps 2-4)

#### Step 2: Initial Draft Creation
- **Job Analysis**: Deep analysis of requirements, keywords, and cultural context
- **Cultural Profiling**: Applies region-specific resume conventions (Canadian, US, European, etc.)
- **HAM-Z Application**: Transforms experience using Hard Skills + Action + Metrics + Structure formula
- **ATS Optimization**: Creates targeted first draft optimized for applicant tracking systems
- **Output**: `Step1_Draft_[Role]_[Company]_[Date].md`

#### Step 3: Provenance Analysis (Enhanced Anti-Hallucination System)
- **Mandatory Evidence Verification**: Every claim cross-referenced against master documents with exact quotes
- **Line Number Validation**: Verifies all quoted references actually exist in source materials
- **Critical Risk Detection**: Flags fabricated capabilities, skills, or experience not documented
- **Hallucination Prevention**: Zero tolerance for assumed language skills, geographic experience, or industry knowledge
- **Risk Categorization**: Classifies issues as Critical/High/Medium/Low with specific remediation strategies
- **Output**: `Step2_Provenance_Analysis_[Role]_[Company]_[Date].md`

#### Step 4: Final Resume Production
- **Systematic Hardening**: Incorporates all provenance recommendations while maintaining competitive positioning
- **Evidence-Based Claims**: Ensures every bullet point can withstand interview-level scrutiny
- **Defensive Positioning**: Balances aggressive positioning with complete defensibility
- **Interview Readiness**: Produces final version ready for submission and interview preparation
- **Output**: `Step3_Final_Resume_[Role]_[Company]_[Date].md`

### Phase 3: Interview Preparation (Steps 5-6)

#### Step 5: Gap Analysis and Study Guide Creation
- **Critical Gap Identification**: Analyzes Step 1 assessment to pinpoint skill gaps and weaknesses
- **Learning Resource Research**: Web research for current best practices and training materials
- **Structured Study Plan**: Creates detailed study guide with actionable timelines
- **Two Modes**: gaps-only (focus on critical weaknesses) or comprehensive preparation
- **Practical Elements**: Includes hands-on exercises, quick references, and interview strategies
- **Output**: `Briefing_Notes/Briefing_[Company]_[Role]_[Mode]_[Date].md`

#### Step 6: Customized Interview Question Generation
- **Resume-Job Alignment**: Generates questions based on Step 4 final resume content vs job requirements
- **Question Balance**: Mix of technical, behavioral, and experience verification questions
- **Strategic Guidance**: STAR format coaching and answer frameworks for each question
- **Follow-up Preparation**: Anticipated follow-up questions and response strategies
- **Risk Mitigation**: Red flags to avoid and strength positioning techniques based on Step 1 gaps
- **Output**: `Briefing_Notes/Interview_Prep_[Company]_[Role]_[Date].md`

### Phase 4: Application Finalization (Steps 7-8)

#### Step 7: Cover Letter Generation (Optional)
- **Strategic Narrative**: Creates compelling cover letter from validated Step 4 final resume
- **Requirements Mapping**: Features visual requirements-matching table for hiring manager impact
- **Evidence Chain**: Maps specific job requirements to proven achievements from master documents
- **Hiring Manager Focus**: Tailors messaging to decision-maker priorities and pain points
- **Output**: `Step4_CoverLetter_[Role]_[Company]_[Date].md`

#### Step 8: Document Conversion (Optional)
- **Professional Formatting**: Converts markdown documents to submission-ready Word DOCX format
- **Pandoc Integration**: Uses pandoc for high-quality conversion preserving tables and styling
- **Batch Processing**: Supports single files, patterns, or bulk conversion (`resume`, `coverletter`, `all`)
- **Enterprise Ready**: Creates polished documents meeting corporate submission standards
- **Output**: `[OriginalName].docx` for any converted markdown document


## HAM-Z™ Methodology

Every achievement follows this formula:

**Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]**

Example:
> Achieved 50% reduction in lease processing time (60→30 days) by leveraging VTS platform expertise to implement automated workflow system across 150-property portfolio

## Provenance Risk Categories

### Critical Risk (Cannot Proceed)
- **Fabricated Claims**: Any capability, skill, or experience not documented in master materials
- **Language Skills**: Claims about bilingual abilities without explicit documentation
- **Industry Experience**: Assumed expertise in sectors not covered in work history
- **Geographic Experience**: Claims about regional knowledge without supporting evidence

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

## Best Practices

1. **Maintain Master Resume**: Keep your career inventory comprehensive and updated
2. **One Job, One Resume**: Always tailor to specific role requirements
3. **Evidence First**: Every claim should be defensible with documentation
4. **Quantify Everything**: Use numbers, percentages, timeframes consistently
5. **Cultural Awareness**: Match resume style to target market expectations

## File Naming Convention

Output files follow this pattern:
**Opportunity Assessment (Step 1):**
- Step 1: `Assessment_[Company]_[Role]_[Date].md` and `Rubric_[Company]_[Role]_[Date].md`

**Resume Development (Steps 2-4):**
- Step 2: `Step1_Draft_[Role]_[Company]_[Date].md`
- Step 3: `Step2_Provenance_Analysis_[Role]_[Company]_[Date].md`
- Step 4: `Step3_Final_Resume_[Role]_[Company]_[Date].md`

**Interview Preparation (Steps 5-6):**
- Step 5: `Briefing_[Company]_[Role]_[Mode]_[Date].md`
- Step 6: `Interview_Prep_[Company]_[Role]_[Date].md`

**Application Finalization (Steps 7-8):**
- Step 7: `Step4_CoverLetter_[Role]_[Company]_[Date].md`
- Step 8: `[OriginalName].docx` (Word versions of any markdown document)

## Example Workflows

### Multi-Job Comparison Workflow

When evaluating multiple opportunities simultaneously, use the comparison command to make strategic decisions:

```bash
# First, assess each opportunity individually
/assessjob JLL_VP_Office_Leasing.md
/assessjob Canerector_Vice_President_Real_Estate.md
/assessjob CityOfToronto_SeniorDirector.md

# Then compare assessments to identify best fit
/comparejobs Assessment_JLL_VP_Office_Leasing_2025-09-25.md Assessment_Canerector_Vice_President_Real_Estate_2025-09-26.md Assessment_CityOfToronto_SeniorDirectorAssetManagement_2025-09-25.md
```

**Sample Output**:
- **Strategic Rankings**: Which roles offer the best career advancement
- **Score Comparison**: Detailed breakdown across all evaluation categories
- **Negotiation Positioning**: Where you have strongest leverage for salary/terms
- **Risk Analysis**: Which opportunities have highest probability of success
- **Interview Strategy**: Tailored approaches for each opportunity

### Comprehensive Company Intelligence Workflow

For thorough company research before interviews or major decisions:

```bash
# Comprehensive OSINT intelligence gathering
/osint Canerector

# Follow up with job-specific assessment
/assessjob Canerector_Vice_President_Real_Estate.md

# Create targeted resume if assessment is positive
/buildresume Canerector_Vice_President_Real_Estate.md Canadian

# Prepare for interview with study guide
/briefing Assessment_Canerector_Vice_President_Real_Estate_2025-09-26.md Canerector_Vice_President_Real_Estate.md gaps-only
```

**OSINT Sample Output**:
- **Master Intelligence Report**: Comprehensive analysis across all intelligence domains with integrated findings
- **6 Specialized Intelligence Reports**: Individual detailed reports saved to `Intelligence_Reports/` folder:
  - `[Company]_Corporate_Intelligence_[Date].md` - Financial health, strategic positioning with source citations
  - `[Company]_Legal_Intelligence_[Date].md` - Litigation history with case citations and court database URLs
  - `[Company]_Leadership_Intelligence_[Date].md` - Executive backgrounds with LinkedIn and interview sources
  - `[Company]_Compensation_Intelligence_[Date].md` - Salary data with Glassdoor/Levels.fyi citations and sample sizes
  - `[Company]_Culture_Intelligence_[Date].md` - Employee reviews with platform URLs and confidence ratings
  - `[Company]_Market_Intelligence_[Date].md` - Industry analysis with market research report citations
- **Enhanced Attribution**: Every claim includes citations, confidence levels, and source reliability ratings
- **Professional Standards**: Intelligence community-level reporting with footnotes and methodology disclosure
- **Strategic Recommendations**: Employment/investment/partnership guidance with complete evidence chain

## Requirements

- Claude Code CLI installed and configured
- Master resume documents in markdown format
- Job descriptions in markdown format
- `pandoc` installed for Step 5 document conversion (optional)

## Support

For issues or improvements, please check:
- `comprehensive_work_history_FAQ.md` for master resume philosophy
- `CLAUDE.md` for technical implementation details
- `SourceMaterial/` for methodology documentation
- `.claude/agents/` for agent-specific logic

## License

Private repository - All resume content is confidential
