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

## Directory Structure

```
resumeoptimizer/
├── ResumeSourceFolder/       # Master resume documents (your career inventory)
├── SourceMaterial/           # Methodology guides and frameworks
├── OutputResumes/            # Generated resumes and analysis reports
├── .claude/                  # Claude Code agents and commands
└── [job-description].md      # Job postings for targeting
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

## The Four-Step Process

### Step 1: Initial Draft Creation
- Analyzes job requirements and keywords
- Applies cultural profile preferences
- Uses HAM-Z methodology for bullet points
- Creates targeted first draft optimized for ATS

### Step 2: Provenance Analysis
- Verifies every claim against master documents
- Identifies evidence gaps and credibility risks
- Categorizes issues by severity (High/Medium/Low)
- Provides specific recommendations for fixes

### Step 3: Final Resume Production
- Incorporates all provenance recommendations
- Maintains competitive positioning
- Ensures complete defensibility
- Produces interview-ready final version

### Step 4: Cover Letter Generation (Optional)
- Creates compelling cover letter from validated resume
- Features requirements-matching table for visual impact
- Maps job requirements to specific achievements
- Maintains provenance chain from master documents

## HAM-Z™ Methodology

Every achievement follows this formula:

**Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]**

Example:
> Achieved 50% reduction in lease processing time (60→30 days) by leveraging VTS platform expertise to implement automated workflow system across 150-property portfolio

## Provenance Risk Categories

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

## Requirements

- Claude Code CLI installed and configured
- Master resume documents in markdown format
- Job descriptions in markdown format

## Support

For issues or improvements, please check:
- `CLAUDE.md` for technical implementation details
- `SourceMaterial/` for methodology documentation
- `.claude/agents/` for agent-specific logic

## License

Private repository - All resume content is confidential
