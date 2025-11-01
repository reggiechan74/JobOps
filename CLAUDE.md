# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

**JobOps - Intelligence-Driven Application Warfare System**

This is a tactical job application platform that uses an 8-step methodology to create tailored, credible resumes from master career inventory documents. The system employs the HAM-Z methodology (Hard Skill, Action, Metrics, Structure combined with XYZ narrative structure) and includes provenance hardening to ensure all claims are defensible.

**Current Version:** 1.0.0 (see CHANGELOG.md for version history)

## Key Directories

- **ResumeSourceFolder/**: Contains master resume documents (comprehensive career inventory)
 

- **SourceMaterial/**: Contains methodology documentation
  - `CV_Master_Resume_Guide_v3.md`: Complete HAM-Z methodology guide
  - `CV_Provenance_Hardening_Pass_Check_Prompt.md`: Provenance analysis framework

- **Scoring_Rubrics/**: Contains assessment rubrics and scoring frameworks
  - Dynamic rubrics: Job-specific rubrics created by `/assessjob` or `/createrubric` commands with embedded detailed scoring framework
  - Reusable rubrics: Created by `/createrubric` for consistent evaluation across multiple candidates
  - Format: `Rubric_[Company]_[Role]_[Date].md` for each assessment

- **OutputResumes/**: Generated resume drafts and analysis outputs
  - Step 1 drafts: Initial targeted resumes
  - Step 2 analyses: Provenance check reports
  - Step 3 finals: Hardened, interview-ready resumes

- **Job_Postings/**: Target job descriptions in markdown format
  - Store job postings as `.md` files for processing
  - Use descriptive filenames (e.g., `CompanyName_Role_Date.md`)
  - Search results automatically saved as `SearchResults_[Query]_[Date].md` with complete job descriptions

- **Briefing_Notes/**: Interview preparation and skill gap study guides
  - Gap analysis briefings: Focus on addressing identified weaknesses
  - Full preparation briefings: Comprehensive interview readiness guides
  - Interview question sets: Customized questions with answer strategies
  - Format: `Briefing_[Company]_[Role]_[Mode]_[Date].md` and `Interview_Prep_[Company]_[Role]_[Date].md`

## Complete Application Process

### Core Resume Development (Steps 1-3)
1. **Step 1 - Initial Draft**: Creates tailored resume using HAM-Z methodology and cultural profile preferences
2. **Step 2 - Provenance Check**: Analyzes credibility, evidence gaps, and risk factors
3. **Step 3 - Final Resume**: Produces hardened version addressing all credibility concerns

### Interview Preparation (Steps 4-6)
4. **Step 4 - Assessment & Gap Analysis**: Evaluates candidate against job requirements with scoring rubric
5. **Step 5 - Study Guide Creation**: Generates briefing notes to address skill gaps or comprehensive prep
6. **Step 6 - Interview Question Prep**: Creates customized questions with strategic answer guidance

### Application Finalization (Steps 7-8)
7. **Step 7 - Cover Letter** (Optional): Generates strategic cover letter with requirements-matching table
8. **Step 8 - Document Conversion** (Optional): Converts markdown to professional Word DOCX format

## Available Slash Commands

### Core Resume Development
- `/buildresume <job-description-file> [cultural-profile]`: Runs complete 3-step resume process
  - Default cultural profile: Canadian
  - Creates draft, performs provenance check, produces final resume

- `/provenance <draft-resume-file>`: Performs standalone provenance analysis
  - Checks draft against master resume documents
  - Identifies credibility issues and evidence gaps

### Interview Preparation Workflow

#### Modular Assessment Commands
- `/createrubric <job-posting-file>`: Create reusable scoring rubric only
  - Generates detailed job-specific 100-point scoring framework
  - Performs web research for domain expertise and industry standards
  - Saves to Scoring_Rubrics/ folder for consistent evaluation
  - Enables standardized assessment across multiple candidates

- `/assesscandidate <rubric-file> <job-posting-file>`: Assess using pre-created rubric
  - Uses existing scoring rubric for consistent candidate evaluation
  - Applies rubric criteria without modification for fairness
  - Maps candidate evidence to specific rubric requirements
  - Maintains complete audit trail and traceability

- `/assessjob <job-posting-file>`: Complete assessment in one step (Step 4)
  - Creates dynamic job-specific scoring rubric from job posting
  - Saves custom rubric to Scoring_Rubrics/ folder for reuse
  - Performs web research for domain expertise
  - Generates 100-point assessment with evidence
  - Provides hiring recommendations and interview strategy

- `/comparejobs <assessment-file-1> <assessment-file-2> [assessment-file-3] [assessment-file-4]`: Compare multiple job assessments
  - Analyzes 2-4 assessments from OutputResumes/ folder
  - Provides strategic hiring insights and role prioritization
  - Identifies consistent patterns and transferable skills
  - Generates comprehensive comparison report with recommendations
  - Supports career planning and negotiation positioning

- `/briefing <assessment-report> <job-description> [gaps-only]`: Creates study guide briefing (Step 5)
  - Analyzes assessment to identify skill gaps and weaknesses
  - Researches current best practices and learning resources
  - Generates detailed study guide with timelines
  - Two modes: gaps-only focus or comprehensive preparation
  - Includes interview questions, hands-on exercises, and quick references

- `/interviewprep <resume-file> <job-description> [number-of-questions]`: Generate interview questions (Step 6)
  - Creates likely interview questions based on resume-job alignment
  - Defaults to 10 questions if number not specified
  - Balances technical and behavioral questions
  - Provides answer strategies and STAR format guidance
  - Includes follow-up questions and red flags to avoid

### Application Finalization
- `/coverletter <step3-resume-file> <job-description-file> [hiring-manager-name]`: Creates cover letter (Step 7)
  - Generates from validated Step 3 resume
  - Includes strategic requirements-matching table
  - Maintains provenance chain

- `/convert <file-path-or-pattern> [output-directory]`: Converts to Word DOCX format (Step 8)
  - Uses pandoc for professional conversion
  - Supports patterns: `resume`, `coverletter`, `all`, file paths
  - Preserves formatting and creates submission-ready documents

### Job Search & Intelligence
- `/searchjobs <search-query> [location] [--company=name] [--save] [--limit=N]`: Search hiring.cafe for job postings
  - Two-phase hybrid search: Fast API search + Playwright scraping for complete job descriptions
  - Search by keywords, location, company name
  - Optional flags: --company=name, --save, --limit=N (default: 20, max recommended: 50)
  - Returns structured job listings with complete verbatim job descriptions
  - Automatically saves results to Job_Postings/ folder when --save flag provided
  - Format: `SearchResults_[Query]_[Date].md` with full market analysis

- `/osint <company-name>`: Comprehensive company intelligence gathering
  - Deploys 6 specialized OSINT agents in parallel
  - Analyzes corporate structure, legal, leadership, compensation, culture, market
  - Generates Master Intelligence Report for strategic decision-making

### System Setup
- `/create-career-history <resume-file-1> [resume-file-2] ...`: Intelligent career history parser
  - Accepts one or more resume files (PDF, DOCX, TXT, MD) as input
  - Parses existing resumes and extracts structured career information
  - Creates complete ResumeSourceFolder structure with pre-populated files
  - Transforms content to HAM-Z format (Hard Skill + Action + Metrics)
  - Identifies content gaps and provides enhancement recommendations
  - Marks estimated metrics and flags provenance concerns
  - Essential first step for new users before building tailored resumes

- `/install-pandoc [force]`: Installs pandoc for document conversion
  - Auto-detects operating system and package manager
  - Verifies installation and conversion capabilities
  - Required dependency for document conversion functionality

### Portfolio Management
- `/github-portfolio [-create|-update] [repository-url-or-analysis]`: Create or update GitHub portfolio documentation
  - **Default mode (-create)**: Creates initial GitHub_Repositories.md with comprehensive structure
  - **Update mode (-update <repo-url>)**: Adds new repository to existing portfolio
  - **Refresh mode (-update)**: Full portfolio reorganization and content refresh
  - Provides repository analysis script for user to run in target repository's codespace
  - Creates comprehensive repository entries with technical architecture, skills demonstrated, business impact
  - Maintains 11 capability matrices tracking proficiency across all technical areas
  - Updates Table of Contents, conclusion, and version tracking automatically
  - Output: `ResumeSourceFolder/Technology/GitHub_Repositories.md`

## Custom Agents

The repository includes specialized agents for each step:

### Resume Development Agents
- `step1-resume-draft`: Initial targeted resume creation
- `step2-provenance-check`: Comprehensive credibility analysis
- `step3-final-resume`: Final hardened version production

### Interview Preparation Agents
- `candidate-assessment`: Expert HR/domain assessment against job descriptions (Step 4)
- `interview-briefing`: Creates comprehensive study guides for skill gaps and interview prep (Step 5)
- `interview-question-generator`: Generates customized interview questions with answer strategies (Step 6)

### Modular Assessment Agents
- `general-purpose`: Used by `/createrubric` for rubric-only generation
- `candidate-assessment`: Used by `/assesscandidate` for rubric-based evaluation

### Application Support Agents
- `step4-cover-letter`: Strategic cover letter with requirements table (Step 7)
- `step5-document-converter`: Markdown to Word DOCX conversion using pandoc (Step 8)

### Job Search & Intelligence Agents
- `hiringcafe-search`: Phase 1 API search agent - fast job discovery with structured metadata
  - Searches hiring.cafe API using keyword, location, and company filters
  - Returns job metadata: title, company, location, category, seniority, apply URLs
  - Provides summary statistics and market insights
  - Phase 2 (main session) uses Playwright MCP to scrape complete verbatim job descriptions
- `osint-agent`: Orchestrates comprehensive company intelligence gathering operations
- `osint-corporate`: Corporate structure, financials, and strategic positioning analysis
- `osint-legal`: Litigation history, regulatory compliance, and legal risk assessment
- `osint-leadership`: Executive backgrounds and leadership analysis
- `osint-compensation`: Salary benchmarking and total rewards intelligence
- `osint-culture`: Employee sentiment and workplace culture analysis
- `osint-market`: Industry analysis and competitive landscape intelligence

### Legacy Agents
- `resume-tailoring-specialist`: Deprecated orchestrator (use 8-step process instead)

## Working with Resume Files

**For new users**: If your `ResumeSourceFolder/` is empty or doesn't exist, run `/create-career-history <your-resume-file.pdf>` first to parse your existing resume and set up the complete folder structure with pre-populated, HAM-Z-enhanced content.

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

### Assessment Workflow Options

#### Option 1: Complete Assessment (Traditional)
```bash
/assessjob JobPosting.md  # Creates rubric + performs assessment
```

#### Option 2: Modular Assessment (Recommended for Multiple Candidates)
```bash
/createrubric JobPosting.md  # Create reusable rubric once
/assesscandidate Rubric_Company_Role_Date.md JobPosting.md  # Assess each candidate
```

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

The resume optimization system follows a comprehensive 8-step pipeline architecture:

### Phase 1: Resume Development (Steps 1-3)
1. **Input Layer**: Job postings + Master resume data
2. **Processing Layer**: HAM-Z methodology transformation with cultural profile adaptation
3. **Validation Layer**: Provenance hardening and credibility assessment
4. **Output Layer**: Final, defensible resume ready for submission

### Phase 2: Interview Preparation (Steps 4-6)
5. **Assessment Layer**: HR-level candidate evaluation with scoring rubrics
6. **Learning Layer**: Gap analysis and skill development planning
7. **Practice Layer**: Customized interview question generation with strategic guidance

### Phase 3: Application Finalization (Steps 7-8)
8. **Communication Layer**: Strategic cover letter with requirements matching
9. **Presentation Layer**: Professional document conversion for submission

Each step in the pipeline has dedicated agents that maintain consistency and quality throughout the complete application process, from initial resume creation to final interview preparation.

## Version Tracking and Updates

JobOps follows [Semantic Versioning](https://semver.org/) (MAJOR.MINOR.PATCH):

### When Making Changes

**PATCH version (1.0.X)** - Increment for:
- Bug fixes and minor improvements
- Documentation updates
- Small enhancements to existing features
- Non-breaking changes

**MINOR version (1.X.0)** - Increment for:
- New features and capabilities
- New slash commands or agents
- Backward-compatible functionality additions
- Significant enhancements

**MAJOR version (X.0.0)** - Increment for:
- Breaking changes or incompatible API changes
- Major system rewrites or architecture changes
- Removal of deprecated features
- Changes requiring user intervention

### Update Process

When making significant changes:

1. **Update version in package.json**:
   - Change the `version` field to reflect the new version number

2. **Update version in README.md**:
   - Update the version badge near the top of the file

3. **Update CHANGELOG.md**:
   - Add new version section following Keep a Changelog format
   - Document all changes under appropriate categories:
     - Added (new features)
     - Changed (changes to existing functionality)
     - Deprecated (soon-to-be-removed features)
     - Removed (removed features)
     - Fixed (bug fixes)
     - Security (security improvements)

4. **Commit with version tag**:
   ```bash
   git add package.json README.md CHANGELOG.md
   git commit -m "Release version X.Y.Z"
   git tag -a vX.Y.Z -m "Version X.Y.Z"
   git push && git push --tags
   ```

### Version Information Location

- **package.json**: Primary version source (line 3)
- **README.md**: Version badge with changelog link (line 7)
- **CHANGELOG.md**: Complete version history and release notes
- **CLAUDE.md**: Current version reference (line 11)

Always keep version numbers synchronized across all files when releasing new versions.