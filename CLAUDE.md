# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

This is a resume optimization system that uses a three-step methodology to create tailored, credible resumes from master career inventory documents. The system employs the HAM-Z methodology (Hard Skill, Action, Metrics, Structure combined with XYZ narrative structure) and includes provenance hardening to ensure all claims are defensible.

## Key Directories

- **ResumeSourceFolder/**: Contains master resume documents (comprehensive career inventory)
  - `Comprehensive_CV_MASTER_COPY_v35.md`: Primary master resume with complete work history
  - `Comprehensive_CV_Technology_Capability.md`: Technical skills and capabilities inventory

- **SourceMaterial/**: Contains methodology documentation
  - `CV_Master_Resume_Guide_v3.md`: Complete HAM-Z methodology guide
  - `CV_Provenance_Hardening_Pass_Check_Prompt.md`: Provenance analysis framework

- **OutputResumes/**: Generated resume drafts and analysis outputs

## Three-Step Resume Process

1. **Step 1 - Initial Draft**: Creates tailored resume using HAM-Z methodology and cultural profile preferences
2. **Step 2 - Provenance Check**: Analyzes credibility, evidence gaps, and risk factors
3. **Step 3 - Final Resume**: Produces hardened version addressing all credibility concerns

## Available Slash Commands

- `/buildresume <job-description-file> [cultural-profile]`: Runs complete 3-step process
  - Default cultural profile: Canadian
  - Creates draft, performs provenance check, produces final resume

- `/provenance <draft-resume-file>`: Performs standalone provenance analysis
  - Checks draft against master resume documents
  - Identifies credibility issues and evidence gaps

## Custom Agents

The repository includes specialized agents for each step:
- `step1-resume-draft`: Initial targeted resume creation
- `step2-provenance-check`: Comprehensive credibility analysis
- `step3-final-resume`: Final hardened version production
- `resume-tailoring-specialist`: Deprecated orchestrator (use 3-step process instead)

## Working with Resume Files

When processing resumes:
1. Always read from `ResumeSourceFolder/` for master data
2. Job descriptions are typically stored as `.md` files in the root
3. Output resumes should be saved to `OutputResumes/` with descriptive timestamps
4. Use the HAM-Z formula: **Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]**

## Provenance Analysis Categories

High Risk issues to flag:
- Unbounded metrics without timeframes
- Unsupported superlatives ("market-leading", "transformational")
- Cross-document inconsistencies
- Benchmark claims without defined peer sets

Medium Risk issues:
- Mechanism-free outcomes (results without "how")
- Duplicate achievements across roles
- Tool names without measurable outcomes

## Cultural Profiles

The system supports different cultural profiles for resume tailoring. When not specified, defaults to Canadian style which emphasizes:
- Quantified achievements
- Clear role boundaries
- Specific technologies and methodologies
- Conservative language without superlatives