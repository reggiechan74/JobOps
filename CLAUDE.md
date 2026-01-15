# CLAUDE.md

> **MAINTAINER NOTE:** This file is loaded into context with every prompt. Keep it under 150 lines. Detailed command/agent docs belong in `.claude/commands/` and `.claude/agents/`. Only add quick-reference information here.

## Repository Purpose

**JobOps v1.5.2** - Intelligence-driven job application system using 8-step methodology to create tailored, credible resumes from master career inventory. Uses HAM-Z methodology (Hard Skill, Action, Metrics, Structure) with provenance hardening.

## Key Directories

| Directory | Purpose |
|-----------|---------|
| `ResumeSourceFolder/` | Master resume data (Experience/, CareerHighlights/, Technology/, Preferences/, .profile/) |
| `Job_Postings/` | Target job descriptions as `.md` files |
| `OutputResumes/` | Generated drafts, analyses, finals |
| `Scoring_Rubrics/` | Assessment rubrics (`Rubric_[Company]_[Role]_[Date].md`) |
| `Briefing_Notes/` | Interview prep and study guides |
| `Client_Prospects/` | Independent contractor service definitions and prospects |
| `.claude/templates/` | Assessment framework templates (rubric, evidence, report schemas) |
| `.claude/commands/` | Slash command definitions |
| `.claude/agents/` | Specialized agent definitions |

## 8-Step Application Process

**Resume Development (Steps 1-3):** Draft → Provenance Check → Final Resume
**Interview Prep (Steps 4-6):** Assessment → Study Guide → Interview Questions
**Finalization (Steps 7-8):** Cover Letter → Document Conversion

## Slash Commands Quick Reference

### Core Resume (`/buildresume`, `/provenance`)
- `/buildresume <job-file> [profile]` - Complete 3-step resume process (default: Canadian profile)
- `/provenance <draft-file>` - Standalone credibility analysis

### Assessment (`/createrubric`, `/assessjob`, `/assesscandidate`, `/comparejobs`)
- `/createrubric <job-file>` - Create reusable 100-point scoring rubric
- `/assessjob <job-file>` - Dynamic rubric + full assessment
- `/assesscandidate <rubric-file> <job-file>` - Assess using pre-created rubric
- `/comparejobs <assessment-1> <assessment-2> [3] [4]` - Compare 2-4 assessments

### Interview Prep (`/briefing`, `/interviewprep`)
- `/briefing <assessment> <job-file> [gaps-only|1d|2d|1w]` - Study guide with priority tags
- `/interviewprep <resume> <job-file> [count] [prep-time]` - Interview questions with likelihood tags

### Finalization (`/coverletter`, `/formatresume`, `/convert`)
- `/coverletter <step3-resume> <job-file> [manager-name]` - Strategic cover letter
- `/formatresume <md-file> [modern|classic|minimal] [1|2|3|auto]` - PDF via Playwright
- `/convert <file-or-pattern> [output-dir]` - Word DOCX via pandoc

### Job Search (`/searchjobs`, `/osint`, `/auditjobposting`)
- `/searchjobs <query> [location] [--company=X] [--save] [--limit=N]` - hiring.cafe search
- `/osint <company>` - 6-agent parallel company intelligence
- `/auditjobposting <job-file>` - 100-point job posting quality audit

### Career Strategy (`/idealjob`, `/change-one-thing`, `/assess-job-offer`)
- `/idealjob [filename]` - Generate synthetic ideal job description
- `/change-one-thing <folder>` - Career retrospective with counterfactual analysis
- `/assess-job-offer <offer-file> [job-posting] [--counter-offer]` - Comprehensive offer analysis (compensation, legal, alignment)

### Career Crisis Management
- `/code-red [docs] [--mode=assess|respond|plan|exit]` - Employment crisis intervention
- `/severance-review <agreement> [--benchmark] [--counter-offer]` - Severance analysis
- `/workplace-documentation [--new-incident|--review|--timeline]` - Incident logging
- `/non-compete-analysis <agreement> [--state=XX]` - Restrictive covenant analysis
- `/reference-shield [--assess|--build|--rescue]` - Reference risk management
- `/unemployment-prep [--state=XX] [--appeal]` - UI claim preparation
- `/discrimination-assessment [log] [--protected-class=X]` - Pattern assessment
- `/investigation-prep [--accused|--complainant|--witness]` - Investigation prep
- `/accommodation-request [--disability|--religious|--medical]` - Accommodation builder
- `/layoff-intel [--company=X] [--assess|--prepare]` - Layoff risk assessment
- `/constructive-dismissal [log] [--jurisdiction=X]` - Constructive dismissal analysis

### Independent Contractor
- `/defineservices [--guided|--from-profile|--update]` - Service catalog creation
- `/ratecard [--format=md|pdf] [--currency=CAD|USD]` - Professional rate card
- `/findclient [job-file] [--industry=X] [--size=X]` - B2B prospect discovery
- `/pitchdeck [--prospect=X] [--industry=X]` - B2B pitch deck generation
- `/proposaltemplate [--client=X] [--type=project|retainer]` - Consulting proposals

### Landing Pages
- `/landing-page:create <name> [--template=tactical|minimal|corporate]`
- `/landing-page:css-template [--view|--analyze|--create]`
- `/landing-page:copywrite <purpose> [--tone=X] [--framework=PAS|AIDA]`

### System Setup
- `/create-career-history <resume-files...>` - Parse existing resumes into ResumeSourceFolder
- `/github-portfolio [-create|-update]` - GitHub portfolio documentation
- `/install-pandoc` - Install pandoc for document conversion

## HAM-Z Formula

**Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]**

## Provenance Risk Categories

**High Risk:** Unbounded metrics, unsupported superlatives, cross-document inconsistencies, benchmark claims without peer sets
**Medium Risk:** Mechanism-free outcomes, duplicate achievements, tool names without outcomes

## File Naming Conventions

- Job Postings: `Job_Postings/CompanyName_Role_Date.md`
- Rubrics: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`
- Outputs: Auto-named with step, role, company, date

## Template Architecture

Templates in `.claude/templates/` define canonical structures:
- `assessment_rubric_framework.md` - 100-point scoring (Tech 25, Experience 25, Responsibilities 20, Achievements 15, Education 10, Cultural 5)
- `evidence_verification_framework.md` - Citation requirements, domain verification, experience classification
- `assessment_report_structure.md` - Report format with 3-level evidence attribution
- `candidate_profile_schema.json` - JSON schema for optimized profiles (85-90% token reduction)

## Key Agents

| Agent | Purpose |
|-------|---------|
| `step1-resume-draft` | Initial tailored resume |
| `step2-provenance-check` | Credibility analysis |
| `step3-final-resume` | Hardened final version |
| `candidate-assessment` | HR-level evaluation |
| `interview-briefing` | Study guide creation |
| `interview-question-generator` | Question generation |
| `resume-summarizer` | Profile optimization (85-90% context reduction) |
| `osint-*` | 6 specialized intelligence agents |
| `landing-page-copywriter` | Landing page copy |

## Version Management

Semantic versioning (MAJOR.MINOR.PATCH). Update: `package.json`, `README.md` badge, `CHANGELOG.md`. Tag releases: `git tag -a vX.Y.Z`.
