# Resume Optimizer for Claude Code

A sophisticated resume optimization system that transforms comprehensive career inventories into targeted, credible resumes using a three-step methodology with provenance hardening.

## Overview

Resume Optimizer uses the HAM-Z™ methodology to create powerful, defensible resumes that are:
- **Strategically positioned** for specific roles
- **Culturally attuned** to regional preferences
- **Provenance-hardened** to withstand scrutiny
- **ATS-optimized** for applicant tracking systems

The system employs a three-step process that ensures every claim is backed by evidence and every achievement is quantified with mechanism, metrics, and timeframe.

## Key Features

- **HAM-Z™ Methodology**: Combines Hard Skills, Actions, Metrics, and Structure with XYZ narrative format
- **Provenance Hardening**: Comprehensive credibility analysis to identify and fix evidence gaps
- **Cultural Profiles**: Tailors resume style to regional expectations (Canadian, US, European, etc.)
- **Risk Assessment**: Categorizes claims by defensibility (High/Medium/Low risk)
- **Master Resume System**: Maintains comprehensive career inventory separate from targeted resumes
- **Dynamic Job Scoring**: Creates custom assessment rubrics for each job posting
- **Candidate Assessment**: Expert HR evaluation with domain-specific knowledge

## Directory Structure

```
resumeoptimizer/
├── .claude/                           # Claude Code configuration
│   ├── agents/                        # Specialized processing agents (5 agents)
│   │   ├── candidate-assessment.md    # HR assessment with domain expertise
│   │   ├── step1-resume-draft.md      # Initial targeted resume creation
│   │   ├── step2-provenance-check.md  # Enhanced credibility analysis
│   │   ├── step3-final-resume.md      # Final hardened resume production
│   │   └── step4-cover-letter.md      # Strategic cover letter generation
│   ├── commands/                      # Slash command definitions (7 commands)
│   │   ├── assessjob.md               # Candidate assessment with dynamic rubrics
│   │   ├── buildresume.md             # Complete 3-step resume build
│   │   ├── convert.md                 # Markdown to Word conversion
│   │   ├── coverletter.md             # Cover letter generation
│   │   ├── install-pandoc.md          # Pandoc installation utility
│   │   ├── prime.md                   # Load documentation into context
│   │   └── provenance.md              # Standalone provenance analysis
│   └── settings.local.json            # Local Claude Code settings
├── Job_Postings/                      # Target job descriptions
│   └── AltoJobPost.md                 # Sample job posting
├── OutputResumes/                     # Generated resumes and analysis
│   ├── Assessment_*.md                # Candidate assessment reports
│   ├── Step1_Draft_*.md               # Initial targeted drafts
│   ├── Step2_Provenance_Analysis_*.md # Credibility analysis reports
│   ├── Step3_Final_Resume_*.md        # Interview-ready final resumes
│   ├── Step4_CoverLetter_*.md         # Strategic cover letters
│   └── *.docx                         # Word format conversions
├── ResumeSourceFolder/                # Master career inventory (2 files)
│   ├── Comprehensive_CV_MASTER_COPY_v35.md       # Complete work history
│   └── Comprehensive_CV_Technology_Capability.md # Technical skills inventory
├── Scoring_Rubrics/                   # Assessment rubrics and frameworks
│   └── Rubric_*_*_*.md               # Dynamic job-specific rubrics
├── SourceMaterial/                    # Reserved for methodology guides
│   └── .gitkeep                       # Placeholder (guides stored elsewhere)
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

### 3. Generate Optimized Resume

Use the slash command to run the complete optimization process:

```
/buildresume AltoJobPost.md Canadian
```

This executes all three steps automatically:
1. Creates initial targeted draft
2. Performs provenance analysis
3. Produces final hardened resume

**Optional Step 4**: Generate a strategic cover letter with requirements table:
```
/coverletter OutputResumes/Step3_Final_Resume_[Role]_[Company]_[Date].md AltoJobPost.md
```

**Optional Step 5**: Convert to professional Word format for submission:
```
/install-pandoc    # Install pandoc if needed
/convert all       # Convert all documents to Word format
```

## Available Commands

### Complete Resume Build
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

### Candidate Assessment
```
/assessjob <job-description-file>
```
- Creates dynamic job-specific scoring rubric from job posting
- Performs web research for domain expertise and industry standards
- Generates comprehensive 100-point assessment with evidence mapping
- Saves custom rubric to `Scoring_Rubrics/` for reuse and audit trail
- Provides hiring recommendations, interview strategies, and gap analysis
- Outputs both rubric and assessment report with full traceability

## The Six-Step Process (Five Core + Assessment)

### Step 1: Initial Draft Creation
- **Job Analysis**: Deep analysis of requirements, keywords, and cultural context
- **Cultural Profiling**: Applies region-specific resume conventions (Canadian, US, European, etc.)
- **HAM-Z Application**: Transforms experience using Hard Skills + Action + Metrics + Structure formula
- **ATS Optimization**: Creates targeted first draft optimized for applicant tracking systems
- **Output**: `Step1_Draft_[Role]_[Company]_[Date].md`

### Step 2: Provenance Analysis (Enhanced Anti-Hallucination System)
- **Mandatory Evidence Verification**: Every claim cross-referenced against master documents with exact quotes
- **Line Number Validation**: Verifies all quoted references actually exist in source materials
- **Critical Risk Detection**: Flags fabricated capabilities, skills, or experience not documented
- **Hallucination Prevention**: Zero tolerance for assumed language skills, geographic experience, or industry knowledge
- **Risk Categorization**: Classifies issues as Critical/High/Medium/Low with specific remediation strategies
- **Output**: `Step2_Provenance_Analysis_[Role]_[Company]_[Date].md`

### Step 3: Final Resume Production
- **Systematic Hardening**: Incorporates all provenance recommendations while maintaining competitive positioning
- **Evidence-Based Claims**: Ensures every bullet point can withstand interview-level scrutiny
- **Defensive Positioning**: Balances aggressive positioning with complete defensibility
- **Interview Readiness**: Produces final version ready for submission and interview preparation
- **Output**: `Step3_Final_Resume_[Role]_[Company]_[Date].md`

### Step 4: Cover Letter Generation (Optional)
- **Strategic Narrative**: Creates compelling cover letter from validated Step 3 resume
- **Requirements Mapping**: Features visual requirements-matching table for hiring manager impact
- **Evidence Chain**: Maps specific job requirements to proven achievements from master documents
- **Hiring Manager Focus**: Tailors messaging to decision-maker priorities and pain points
- **Output**: `Step4_CoverLetter_[Role]_[Company]_[Date].md`

### Step 5: Document Conversion (Optional)
- **Professional Formatting**: Converts markdown documents to submission-ready Word DOCX format
- **Pandoc Integration**: Uses pandoc for high-quality conversion preserving tables and styling
- **Batch Processing**: Supports single files, patterns, or bulk conversion (`resume`, `coverletter`, `all`)
- **Enterprise Ready**: Creates polished documents meeting corporate submission standards
- **Output**: `[OriginalName].docx` for any converted markdown document

### Step 6: Candidate Assessment (Optional - From Hiring Manager Perspective)
- **Dynamic Rubric Creation**: Generates job-specific 100-point scoring framework from job posting
- **Domain Knowledge Integration**: Performs web research for industry standards and role expectations
- **Comprehensive Evaluation**: Maps candidate evidence to each job requirement systematically
- **Risk Analysis**: Identifies gaps and provides mitigation strategies for hiring decisions
- **Interview Strategy**: Recommends specific areas to probe and assessment approaches
- **Outputs**:
  - `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md` (custom scoring criteria)
  - `OutputResumes/Assessment_[Company]_[Role]_[Date].md` (detailed evaluation report)

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
- Step 1: `Step1_Draft_[Role]_[Company]_[Date].md`
- Step 2: `Step2_Provenance_Analysis_[Role]_[Company]_[Date].md`
- Step 3: `Step3_Final_Resume_[Role]_[Company]_[Date].md`
- Step 4: `Step4_CoverLetter_[Role]_[Company]_[Date].md`
- Step 5: `[OriginalName].docx` (Word versions of any markdown document)

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
