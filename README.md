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

## Directory Structure

```
resumeoptimizer/
├── .claude/                           # Claude Code configuration
│   ├── agents/                        # Specialized processing agents (7 agents)
│   │   ├── candidate-assessment.md           # HR assessment with domain expertise
│   │   ├── interview-briefing.md             # Gap analysis and study guides
│   │   ├── interview-question-generator.md   # Customized interview questions
│   │   ├── step1-resume-draft.md             # Initial targeted resume creation
│   │   ├── step2-provenance-check.md         # Enhanced credibility analysis
│   │   ├── step3-final-resume.md             # Final hardened resume production
│   │   └── step4-cover-letter.md             # Strategic cover letter generation
│   ├── commands/                      # Slash command definitions (9 commands)
│   │   ├── assessjob.md               # Candidate assessment with dynamic rubrics
│   │   ├── briefing.md                # Interview preparation briefing
│   │   ├── buildresume.md             # Complete 3-step resume build
│   │   ├── convert.md                 # Markdown to Word conversion
│   │   ├── coverletter.md             # Cover letter generation
│   │   ├── install-pandoc.md          # Pandoc installation utility
│   │   ├── interviewprep.md           # Customized interview questions
│   │   ├── prime.md                   # Load documentation into context
│   │   └── provenance.md              # Standalone provenance analysis
│   └── settings.local.json            # Local Claude Code settings
├── Briefing_Notes/                    # Interview preparation materials
│   └── .gitkeep                       # Directory placeholder
├── Job_Postings/                      # Target job descriptions
│   ├── .gitkeep                       # Directory placeholder
│   ├── AltoJobPost.md                 # Sample job posting
│   ├── City_of_Toronto_Director.md    # City of Toronto role
│   └── JLL_VP_Office_Leasing.md       # JLL VP position
├── OutputResumes/                     # Generated resumes and analysis
│   ├── .gitkeep                       # Directory placeholder
│   ├── Assessment_*.md                # Candidate assessment reports
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
├── README.md                          # This documentation
├── comprehensive_work_history_FAQ.md  # Master resume philosophy
└── LICENSE                            # Repository license
```

## Quick Start

### 1. Setup Your Master Resume

Place your comprehensive career inventory in `ResumeSourceFolder/`:
- Complete work history with all projects and achievements
- Technology capabilities and skills inventory
- Awards, publications, and certifications

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

### 5. Prepare for Interview

**Create study guide to address gaps:**
```
/briefing OutputResumes/Assessment_[Company]_[Role]_[Date].md AltoJobPost.md gaps-only
```

**Generate customized interview questions:**
```
/interviewprep OutputResumes/Step3_Final_Resume_[Role]_[Company]_[Date].md AltoJobPost.md
```

### 6. Finalize Application (Optional)

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

## Requirements

- Claude Code CLI installed and configured
- Master resume documents in markdown format
- Job descriptions in markdown format
- `pandoc` installed for Step 5 document conversion (optional)

## Why a Comprehensive Work History?

**The Hallucination Problem**: Traditional 2-page resumes create a fundamental issue when using AI for tailoring. When you try to customize a condensed resume, AI tools must fill gaps with plausible-sounding but fictional content, leading to claims you can't defend in interviews.

**The Solution**: A comprehensive 20+ page master resume serves as your "database of truth" - every achievement, metric, and skill is documented with full context. This allows AI to **select** relevant experiences rather than **invent** them.

### Key Benefits:
- **No Hallucinations**: Enhanced provenance system prevents AI from inventing capabilities
- **Interview Confidence**: Every bullet point traces to documented experience with line-by-line verification
- **Provenance Verification**: All claims verified against source material with mandatory quote-checking
- **Career Pattern Recognition**: Discover valuable themes in your experience
- **Evidence-Based Targeting**: Only documented skills and achievements used in tailoring

📖 **Read the complete explanation**: [`comprehensive_work_history_FAQ.md`](comprehensive_work_history_FAQ.md)

## Support

For issues or improvements, please check:
- `comprehensive_work_history_FAQ.md` for master resume philosophy
- `CLAUDE.md` for technical implementation details
- `SourceMaterial/` for methodology documentation
- `.claude/agents/` for agent-specific logic

## License

Private repository - All resume content is confidential
