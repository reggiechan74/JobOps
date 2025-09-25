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

- **Scoring_Rubrics/**: Contains assessment rubrics and scoring frameworks
  - `jobscoringrubric.md`: Master template for job assessment scoring
  - Dynamic rubrics: Job-specific rubrics created by `/assessjob` command
  - Format: `Rubric_[Company]_[Role]_[Date].md` for each assessment

- **OutputResumes/**: Generated resume drafts and analysis outputs
  - Step 1 drafts: Initial targeted resumes
  - Step 2 analyses: Provenance check reports
  - Step 3 finals: Hardened, interview-ready resumes

- **Job_Postings/**: Target job descriptions in markdown format
  - Store job postings as `.md` files for processing
  - Use descriptive filenames (e.g., `CompanyName_Role_Date.md`)

## Five-Step Application Process

1. **Step 1 - Initial Draft**: Creates tailored resume using HAM-Z methodology and cultural profile preferences
2. **Step 2 - Provenance Check**: Analyzes credibility, evidence gaps, and risk factors
3. **Step 3 - Final Resume**: Produces hardened version addressing all credibility concerns
4. **Step 4 - Cover Letter** (Optional): Generates strategic cover letter with requirements-matching table
5. **Step 5 - Document Conversion** (Optional): Converts markdown to professional Word DOCX format

## Available Slash Commands

- `/buildresume <job-description-file> [cultural-profile]`: Runs complete 3-step process
  - Default cultural profile: Canadian
  - Creates draft, performs provenance check, produces final resume

- `/provenance <draft-resume-file>`: Performs standalone provenance analysis
  - Checks draft against master resume documents
  - Identifies credibility issues and evidence gaps

- `/coverletter <step3-resume-file> <job-description-file> [hiring-manager-name]`: Creates cover letter
  - Generates from validated Step 3 resume
  - Includes strategic requirements-matching table
  - Maintains provenance chain

- `/convert <file-path-or-pattern> [output-directory]`: Converts to Word DOCX format
  - Uses pandoc for professional conversion
  - Supports patterns: `resume`, `coverletter`, `all`, file paths
  - Preserves formatting and creates submission-ready documents

- `/install-pandoc [force]`: Installs pandoc for document conversion
  - Auto-detects operating system and package manager
  - Verifies installation and conversion capabilities
  - Required dependency for Step 5 functionality

- `/assessjob <job-posting-file>`: Expert HR assessment of candidate vs job
  - Creates dynamic job-specific scoring rubric from job posting
  - Saves custom rubric to Scoring_Rubrics/ folder for reuse
  - Performs web research for domain expertise
  - Generates 100-point assessment with evidence
  - Provides hiring recommendations and interview strategy

## Custom Agents

The repository includes specialized agents for each step:
- `step1-resume-draft`: Initial targeted resume creation
- `step2-provenance-check`: Comprehensive credibility analysis
- `step3-final-resume`: Final hardened version production
- `step4-cover-letter`: Strategic cover letter with requirements table
- `step5-document-converter`: Markdown to Word DOCX conversion using pandoc
- `candidate-assessment`: Expert HR/domain assessment against job descriptions
- `resume-tailoring-specialist`: Deprecated orchestrator (use 5-step process instead)

## Working with Resume Files

When processing resumes:
1. Always read from `ResumeSourceFolder/` for master data
2. Job descriptions should be stored in `Job_Postings/` directory as `.md` files
3. Legacy job descriptions may exist in root directory (e.g., `AltoJobPost.md`)
4. Output resumes should be saved to `OutputResumes/` with descriptive timestamps
5. Use the HAM-Z formula: **Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]**

## File Organization Best Practices

- **Job Postings**: Use format `Job_Postings/CompanyName_Role_Date.md`
- **Output Files**: Automatically named with step, role, company, and date
- **Master Resume**: Keep only verified, defensible content in source documents
- **Evidence Trail**: Maintain clear provenance from master to final resume

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

## Important Commands and Tools

### Development Commands
- No build/compile commands needed (markdown-based system)
- No test framework (validation through provenance analysis)
- Git for version control: `git status`, `git add`, `git commit`
- Pandoc for document conversion: `pandoc --version` to verify installation

### Key Validation Steps
When creating or modifying resumes:
1. Run provenance check after any draft creation
2. Verify all metrics have timeframes and mechanisms
3. Ensure cultural profile alignment
4. Check for duplicate achievements across roles

## System Architecture

The resume optimization system follows a pipeline architecture:
1. **Input Layer**: Job postings + Master resume data
2. **Processing Layer**: Three-step transformation pipeline with HAM-Z methodology
3. **Validation Layer**: Provenance hardening and risk assessment
4. **Output Layer**: Credible, targeted resumes ready for submission

Each step in the pipeline has dedicated agents that maintain consistency and quality throughout the transformation process.