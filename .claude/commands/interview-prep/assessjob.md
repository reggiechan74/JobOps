---
description: Perform expert HR assessment of candidate against job posting with domain knowledge
argument-hint: <job-posting-file> [resume-folder-or-file]
---

You are acting as both an expert HR recruiter and a domain-knowledgeable hiring manager. Perform a comprehensive assessment of the candidate's work history against the specified job posting.

## CRITICAL OUTPUT CONSTRAINT

**MAXIMUM ASSESSMENT REPORT LIMIT: 20,000 TOKENS**

Your assessment report output MUST NOT EXCEED 20,000 tokens. This is a hard limit.

**Important Clarifications:**
- **Rubric**: Must include ALL mandatory detailed scoring breakdowns (no token limit - completeness required)
- **Assessment Report**: MUST stay within 20,000 token limit (this is the constraint)

**Strategies to stay within limit for assessment report:**
- Be concise and focused in your analysis
- Use bullet points instead of lengthy paragraphs where appropriate
- Prioritize evidence-based scoring over verbose explanations
- Keep executive summary brief (2-3 sentences maximum)
- Focus on the most impactful evidence for each score
- Eliminate redundancy across sections

## Your Task

Create a job-specific scoring rubric from the {{ARG1}} job posting, then evaluate the candidate using this customized rubric to provide a detailed assessment report.

**Resume Source Location:**
- If {{ARG2}} is provided: Use {{ARG2}} as the resume source (file or folder)
- If {{ARG2}} is not provided: Default to `/workspaces/resumeoptimizer/ResumeSourceFolder`

---

## WORKFLOW ARCHITECTURE

**This workflow uses 5 phases with parallel dispatch where possible.**

```
Phase 1 (Sequential, fast):     Load templates + Validate inputs + Load job posting
Phase 2 (PARALLEL subagents):   Domain Research ‖ Candidate Profile Generation
Phase 3 (Sequential):           Create Rubric (synthesizes job posting + domain research)
Phase 4 (Sequential, visible):  Score Cat 1 → 2 → 3 → 4 → 5
Phase 5 (Sequential):           Generate Report → Save Files
```

**Dependency Rules:**
- Phase 2 starts after job posting is loaded (Phase 1)
- Phase 2 tasks are INDEPENDENT of each other - dispatch simultaneously
- Phase 3 WAITS for domain research results from Phase 2
- Phase 4 WAITS for rubric (Phase 3) AND candidate profile (Phase 2)
- Phase 5 WAITS for all scoring (Phase 4)

---

## PROGRESS TRACKING (MANDATORY)

**Before starting any work**, create all tasks for user visibility. Use TaskCreate for each task below, then use TaskUpdate to mark `in_progress` when starting and `completed` when done.

Create these tasks immediately at the start:

| # | Task Subject | activeForm | Phase |
|---|-------------|------------|-------|
| 1 | Load assessment framework templates | Loading assessment framework templates | 1 |
| 2 | Validate job posting and resume source | Validating job posting and resume source | 1 |
| 3 | Generate candidate profile | Generating candidate profile from source materials | 2 |
| 4 | Research domain and industry context | Researching domain and industry context | 2 |
| 5 | Create job-specific scoring rubric | Creating job-specific scoring rubric | 3 |
| 6 | Score Skills Inventory (Category 1) | Scoring Skills Inventory against rubric | 4 |
| 7 | Score Experience Relevance (Category 2) | Scoring Experience Relevance against rubric | 4 |
| 8 | Score Demonstrated Impact (Category 3) | Scoring Demonstrated Impact against rubric | 4 |
| 9 | Score Credentials (Category 4) | Scoring Credentials against rubric | 4 |
| 10 | Score Fit & Readiness (Category 5) | Scoring Fit & Readiness against rubric | 4 |
| 11 | Generate assessment report | Generating comprehensive assessment report | 5 |
| 12 | Save rubric and assessment files | Saving rubric and assessment files | 5 |

**Task Update Rules:**
- Mark each task `in_progress` BEFORE starting work on it
- Mark each task `completed` AFTER finishing it
- If a task is skipped (e.g., task 3 when source is a single file), mark it `completed` with a note

---

## YAML FRONT MATTER FOR GENERATED FILES

Every markdown artifact you create (rubric and assessment report) must start with YAML metadata populated with real values.

- **Rubric file** (`Scoring_Rubrics/Rubric_*`):
  ```yaml
  ---
  job_file: Job_Postings/{{ARG1}}
  role: <role title>
  company: <company name>
  role_variant: <Technical IC | People Manager | Executive>
  total_points: 200
  generated_by: /assessjob rubric
  generated_on: <ISO8601 timestamp>
  output_type: rubric
  status: final
  version: 2.0
  ---
  ```
- **Assessment file** (`OutputResumes/Assessment_*`):
  ```yaml
  ---
  job_file: Job_Postings/{{ARG1}}
  resume_source: {{ARG2}} or /workspaces/resumeoptimizer/ResumeSourceFolder
  rubric_file: Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md
  role: <role title>
  company: <company name>
  role_variant: <Technical IC | People Manager | Executive>
  candidate: <full candidate name>
  generated_by: /assessjob
  generated_on: <ISO8601 timestamp>
  output_type: assessment
  status: draft
  version: 2.0
  overall_score: <XX/200>
  normalized_score: <XX%>
  ---
  ```

Insert the appropriate block before any headings, updating timestamps and scores, and bump `version` if you rerun the analysis.

---

## PHASE 1: INITIALIZATION (Sequential)

> **Tasks:** Mark task 1 `in_progress`, then task 2 `in_progress` after templates load.

### 1.1 Load Required Templates

**CRITICAL: Read all three framework templates in a single parallel batch:**
- `.claude/templates/assessment_rubric_framework.md` - Master 200-point rubric structure with role variants
- `.claude/templates/evidence_verification_framework.md` - Evidence-based scoring protocols
- `.claude/templates/assessment_report_structure.md` - Assessment report format

Use three parallel Read tool calls. These templates define the mandatory structure for rubrics and assessments.

> **Task:** Mark task 1 `completed`.

### 1.2 Validate Inputs and Load Job Posting

> **Task:** Mark task 2 `in_progress`.

**Resume Source Path Resolution:**

1. **Determine source path**:
   - If {{ARG2}} is provided: Use the specified path
   - If {{ARG2}} is not provided: Default to `/workspaces/resumeoptimizer/ResumeSourceFolder`

2. **Validate source path**:
   - If path is a file: Read the single resume file directly (skip profile generation in Phase 2)
   - If path is a directory: Will trigger profile generation in Phase 2
   - If path doesn't exist: Report error and halt

3. **Load job posting**:
   - Read the job posting from `Job_Postings/{{ARG1}}` (add .md extension if needed)
   - If not found, report error and suggest available alternatives

> **Task:** Mark task 2 `completed`.

---

## PHASE 2: PARALLEL DATA ACQUISITION

> **CRITICAL: Dispatch both tasks simultaneously using parallel Task tool calls in a SINGLE message.**
> Mark tasks 3 and 4 as `in_progress` before dispatching.

Phase 2 launches two independent subagents that run concurrently. Neither depends on the other. Both need only the job posting content (loaded in Phase 1).

### 2.1 Generate Candidate Profile (Task 3)

**NOTE**: Skip this step (mark task 3 `completed` immediately) if resume source is a single file.

**For folder-based sources**, dispatch a subagent to generate/load the candidate profile:

1. **Check for existing profile**: Look for `[resume-source-folder]/.profile/candidate_profile.json`
   - If exists and recent (<=7 days old), use it directly (mark task 3 `completed`)
   - If exists but stale (>7 days old), regenerate
   - If doesn't exist, generate new profile

2. **Dispatch profile generation subagent**:
   ```
   Use Task tool with subagent_type=resume-summarizer, model=sonnet, and prompt:
   "Read all files in [resume-source-folder]/ directory and create a structured JSON
   candidate profile following the schema in .claude/agents/resume-summarizer.md.
   Save output to [resume-source-folder]/.profile/candidate_profile.json and
   [resume-source-folder]/.profile/extraction_log.md"
   ```

3. **Expected outcome**:
   - JSON profile: `[resume-source-folder]/.profile/candidate_profile.json` (8K-10K tokens)
   - Extraction log: `[resume-source-folder]/.profile/extraction_log.md`
   - Token savings: 42K-72K tokens (85-90% reduction)

### 2.2 Research Domain and Industry Context (Task 4)

**Dispatch a domain research subagent** to run concurrently with profile generation:

```
Use Task tool with subagent_type=general-purpose, model=sonnet, and prompt:
"Research the following for the role of [ROLE TITLE] at [COMPANY NAME]:

1. Industry standards and typical role expectations for this specific position
2. Required vs nice-to-have skills based on current market standards
3. Typical responsibilities and seniority indicators for this role level
4. Company context: culture, values, technology stack, recent developments, size, reputation
5. Current market conditions: salary ranges, demand, competitive landscape
6. Industry-specific terminology, certifications, and best practices
7. What differentiates strong vs average candidates for this type of role

Provide a structured research summary organized by these 7 areas.
Focus on actionable intelligence that would help calibrate a scoring rubric.
Be specific - cite sources and data points where possible."
```

**IMPORTANT**: Both subagents (2.1 and 2.2) MUST be dispatched in the SAME message using parallel Task tool calls. Do NOT wait for one to complete before starting the other.

> **Task:** When each subagent completes, mark its respective task (3 or 4) as `completed`.

---

## PHASE 3: CREATE RUBRIC (Sequential - needs Phase 2 domain research)

> **Task:** Mark task 5 `in_progress`.
> **Prerequisite:** Domain research (task 4) must be `completed`. Candidate profile (task 3) is NOT needed yet.

### 3.1 Synthesize Job Posting + Domain Research

Combine the job posting requirements with domain research findings to create an informed rubric. The domain research helps you:
- Set appropriate proficiency thresholds calibrated to industry standards
- Weight skills correctly based on market reality (not just posting emphasis)
- Identify implicit requirements the posting may not state explicitly
- Calibrate experience expectations to the seniority level
- Set realistic achievement thresholds for the industry/company size

### 3.2 Determine Role Variant

Analyze the job posting to select the appropriate weight variant:

**Technical Individual Contributor** - Select if:
- Role is hands-on technical work (engineer, analyst, specialist)
- No direct reports mentioned
- Technical skills are primary focus

**People Manager** - Select if:
- Direct reports or team leadership mentioned
- Performance management, hiring, or development responsibilities
- "Manager", "Director", "Lead" with team oversight

**Executive** - Select if:
- Strategic/organizational scope
- P&L or budget authority
- Cross-functional leadership
- "VP", "C-level", "Executive Director", "Head of"

Document selected variant in rubric header.

### 3.3 Parse Job Posting Requirements

Extract and categorize:
- **Required Technical Skills**: All must-have technical competencies
- **Preferred Skills**: Nice-to-have skills and technologies
- **Experience Requirements**: Years, industry, domain specifics
- **Key Responsibilities**: Primary duties and scope expectations
- **Education/Certifications**: Required and preferred credentials
- **Soft Skills/Cultural Fit**: Communication, collaboration, values

### 3.4 Generate Custom Rubric Structure

**MANDATORY DETAILED SCORING REQUIREMENT**
The dynamic rubric you create MUST include granular point allocation for every single criterion. This is NON-NEGOTIABLE. You cannot create simplified rubrics or skip detailed breakdowns. Every skill, every experience level, every responsibility MUST have the complete scoring framework with specific, measurable criteria.

**CRITICAL REQUIREMENT**: Follow the structure defined in `.claude/templates/assessment_rubric_framework.md` exactly. Customize all `[bracketed]` content with job-specific requirements from the job posting, informed by domain research findings.

**200-Point Scoring Categories (Default Weights):**
1. **Skills Inventory (50 pts / 25%)** - Technical and professional competencies scored on 0-6 proficiency scale
2. **Experience Relevance (40 pts / 20%)** - Years, industry alignment, role progression (5-level scoring)
3. **Demonstrated Impact (60 pts / 30%)** - Achievements with quantified outcomes and metrics
4. **Credentials (20 pts / 10%)** - Education, certifications, and formal qualifications
5. **Fit & Readiness (30 pts / 15%)** - Cultural alignment, availability, and role-specific readiness

**ENFORCEMENT CHECK**: After creating the rubric, verify:
- Role variant selected and documented (Technical IC / Manager / Executive)
- Point allocations match selected variant
- Alignment Statement with construct definition
- Critical Barriers table with explicit thresholds
- Confidence Flagging protocol included
- Skills scored on proficiency (0-6), NOT years
- No redundancy between categories (years only in Experience, achievements only in Impact)
- Each skill has 7-level scoring with anchor examples
- Experience section has 5-level scoring without years-in-skill overlap
- Demonstrated Impact has quantified thresholds

**FAILURE TO INCLUDE DETAILED BREAKDOWNS VIOLATES THE COMMAND REQUIREMENTS**

### 3.5 Validate and Save Dynamic Rubric

**MANDATORY QUALITY CHECK**: Before saving, verify the rubric includes ALL required detailed breakdowns:
- Every required skill has 7-level proficiency scoring (0-6) with specific anchor examples
- Every preferred skill has proficiency-based scoring with role-specific thresholds
- Experience sections have 5-level detailed scoring breakdowns without years-in-skill overlap
- Demonstrated Impact has 5-level scoring with quantitative thresholds
- Credentials have detailed multi-level scoring
- Fit & Readiness has comprehensive evaluation criteria and scoring levels
- Critical Barriers table documents knockout criteria
- Role variant weights are correctly applied

**IF ANY SECTION LACKS DETAILED BREAKDOWNS, THE RUBRIC IS INCOMPLETE AND MUST BE REGENERATED**

Save the generated rubric to: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`
(Audit copy saved later in Phase 5 with the assessment)

> **Task:** Mark task 5 `completed`.

---

## PHASE 4: SCORE CANDIDATE (Sequential - needs rubric + candidate data)

> **Prerequisites:** Rubric (task 5) AND candidate profile (task 3) must both be `completed`.

### 4.1 Load Candidate Materials

**Load candidate materials (depends on source type):**
- **If single file**: Read the resume file directly
- **If folder**: Read the candidate profile from `[resume-source-folder]/.profile/candidate_profile.json`

**Evidence Verification Protocol** (from `.claude/templates/evidence_verification_framework.md`):
- **For folder-based profiles**: When citing specific achievements or skills:
  - Use line references from JSON profile's evidence fields
  - Read specific sections from source files ONLY when verification needed
  - Quote exact text from source files for all scores >=2 points
  - Maintain traceability: JSON profile -> source file -> line numbers

- **For single file resumes**: When citing specific achievements or skills:
  - Reference the resume file directly with line numbers
  - Quote exact text from the resume for all scores >=2 points
  - Maintain traceability: resume file -> line numbers

### 4.2 Score Skills Inventory (Category 1)

> **Task:** Mark task 6 `in_progress`.

Map each required/preferred skill from the custom rubric to specific evidence in candidate's history using 0-6 proficiency scale.

**Score each criterion with confidence level:**
| Criterion | Score | Confidence | Evidence Citation |
|-----------|-------|------------|-------------------|
| [Skill 1] | X/6 | HIGH/MED/LOW | "exact quote" or inference note |

> **Task:** Mark task 6 `completed`.

### 4.3 Score Experience Relevance (Category 2)

> **Task:** Mark task 7 `in_progress`.

Evaluate against the specific years, industry, and domain requirements identified in the rubric (5-level scoring).

> **Task:** Mark task 7 `completed`.

### 4.4 Score Demonstrated Impact (Category 3)

> **Task:** Mark task 8 `in_progress`.

Verify quantified achievements against the expected outcomes defined in the rubric. This is the highest-weighted category - be thorough.

> **Task:** Mark task 8 `completed`.

### 4.5 Score Credentials (Category 4)

> **Task:** Mark task 9 `in_progress`.

Check against the specific education and certification requirements listed in the custom rubric.

> **Task:** Mark task 9 `completed`.

### 4.6 Score Fit & Readiness (Category 5)

> **Task:** Mark task 10 `in_progress`.

Assess based on the company values, work environment, and role-specific readiness factors.

**Normalized Score Calculation:**
- Raw Score: Sum of all category scores (max 200)
- Normalized Score: (Raw Score / 200) x 100%

> **Task:** Mark task 10 `completed`.

---

## PHASE 5: REPORT AND SAVE

### 5.1 Generate Comprehensive Report

> **Task:** Mark task 11 `in_progress`.

**CRITICAL**: Follow the report structure defined in `.claude/templates/assessment_report_structure.md` exactly. Use the detailed 3-level format with rubric criteria attribution for all scores.

**Required Report Elements:**
- Executive Summary with normalized score and role variant
- Category-by-category scoring with confidence flags
- Evidence citations for all scores
- Critical Barriers assessment (pass/fail for each barrier)
- Gap analysis with severity ratings
- Recommendation with confidence level

> **Task:** Mark task 11 `completed`.

### 5.2 Save Assessment and Rubric

> **Task:** Mark task 12 `in_progress`.

**CRITICAL: Folder Structure and Timestamps**

Before saving any files:
1. **Get Current Eastern Time**: Retrieve current date/time in America/New_York timezone
2. **Create Timestamped Sub-Folder**: Format as `YYYY-MM-DD_HHMMSS_[Company]_[Role]`
   - Example: `2025-11-17_143022_UniversityOfToronto_ExecutiveDirectorAssetManagement`
   - Remove spaces from company and role names; use CamelCase or underscores
   - Ensure folder is created in `/workspaces/resumeoptimizer/OutputResumes/`

**File Save Locations:**

1. **Save Dynamic Rubric to BOTH locations**:
   - **Primary**: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`
     - This becomes the official scoring criteria repository
     - Include all extracted requirements from the job posting
     - Document the scoring breakdown with job-specific details

   - **Audit Copy**: `OutputResumes/[YYYY-MM-DD_HHMMSS]_[Company]_[Role]/Rubric_[Company]_[Role]_[Date].md`
     - Exact copy of rubric for complete audit trail
     - Ensures assessment folder is self-contained with all artifacts

2. **Save Assessment Report**: `OutputResumes/[YYYY-MM-DD_HHMMSS]_[Company]_[Role]/Assessment_[Company]_[Role]_[Date].md`
   - Include reference to the specific rubric used
   - Document all scores with evidence mapping and confidence flags
   - Provide clear traceability between rubric criteria and candidate evaluation

**Example Full Paths:**
```
/workspaces/resumeoptimizer/Scoring_Rubrics/Rubric_UofT_ExecutiveDirectorAssetManagement_20251117.md
/workspaces/resumeoptimizer/OutputResumes/2025-11-17_143022_UofT_ExecutiveDirectorAssetManagement/Rubric_UofT_ExecutiveDirectorAssetManagement_20251117.md
/workspaces/resumeoptimizer/OutputResumes/2025-11-17_143022_UofT_ExecutiveDirectorAssetManagement/Assessment_UofT_ExecutiveDirectorAssetManagement_20251117.md
```

**Folder Creation Steps:**
1. Use Bash tool to get Eastern time: `TZ='America/New_York' date '+%Y-%m-%d_%H%M%S'`
2. Extract company name and role from job posting
3. Create folder path: `OutputResumes/[timestamp]_[Company]_[Role]/`
4. Create folder using Bash: `mkdir -p "OutputResumes/[timestamp]_[Company]_[Role]"`
5. Save rubric to Scoring_Rubrics/ (primary location)
6. Copy rubric to timestamped folder (audit trail)
7. Save assessment to timestamped folder

> **Task:** Mark task 12 `completed`.

---

## Validation Requirements

Before finalizing assessment:
1. **Rubric Completeness Check**: Verify the rubric contains all required detailed scoring breakdowns
2. **Role Variant Verification**: Confirm role variant is documented and weights match
3. **Alignment Verification**: Confirm rubric matches the job posting requirements
4. **Evidence Mapping**: Ensure every score has corresponding evidence from candidate materials
5. **Confidence Flagging**: Verify all scores include HIGH/MED/LOW confidence indicators
6. **Consistency Review**: Check that scoring follows rubric criteria without deviation
7. **Documentation Audit**: Verify clear traceability from rubric criteria to assigned scores

## Error Handling

If issues are encountered:
- **Missing Job Posting**: Report if job posting file cannot be found and suggest available alternatives
- **Incomplete Job Posting**: Note any missing critical information and document assumptions made
- **Insufficient Candidate Data**: Document where candidate evidence is limited for fair assessment
- **Domain Knowledge Gaps**: Flag areas where additional research or expertise would improve assessment accuracy

## Important Notes
- **MANDATORY DETAILED SCORING**: Every rubric MUST include the granular scoring breakdowns from the framework rubric. This is not optional - it is required.
- **200-POINT SYSTEM**: Total score is 200 points; normalize to percentage using (Raw/200) x 100%
- **ROLE VARIANT SELECTION**: Every assessment must identify and document the appropriate role variant
- **PROFICIENCY-BASED SKILLS**: Skills use 0-6 proficiency scale, NOT years-based scoring
- **NO CATEGORY REDUNDANCY**: Years only in Experience Relevance, achievements only in Demonstrated Impact
- **Dynamic Rubric Creation**: Each job posting generates its own tailored rubric with job-specific criteria but framework-level detail
- **Complete Scoring Levels**: All sections must include the detailed point breakdowns with specific anchor examples
- **Role-Specific Adaptation**: Customize the content and thresholds for the specific role while maintaining scoring structure depth
- **Consistency**: Use the template framework as the foundation but customize content to job requirements
- **Traceability**: Always reference which rubric was used for which assessment
- **Objectivity**: Be evidence-based in scoring using the job-specific criteria with detailed justification
- **Domain Research**: Domain research informs rubric calibration - thresholds and weights should reflect industry reality
- **Documentation**: Save both rubric and assessment for complete audit trail
- **Confidence Levels**: All scores must include HIGH/MED/LOW confidence flags
- **Quality Control**: The generated rubric must be as detailed as the framework rubric - no shortcuts or simplified versions allowed

## CRITICAL ENFORCEMENT RULES

**ABSOLUTE REQUIREMENTS FOR EVERY RUBRIC:**
1. **ROLE VARIANT** must be selected and weights adjusted accordingly
2. **SKILLS** use 0-6 proficiency scale (7 levels), NOT years-based
3. **NO REDUNDANCY**: Years only in Experience Relevance, achievements only in Demonstrated Impact
4. **DEMONSTRATED IMPACT** is highest-weighted category (25-35% depending on variant)
5. **CONFIDENCE FLAGS** required for all scores in assessment
6. **CRITICAL BARRIERS** table with explicit thresholds and consequences
7. **NORMALIZED SCORE** = Raw/200 x 100%

**VIOLATION CONSEQUENCES:**
- Any rubric missing detailed breakdowns is INCOMPLETE and violates the command specification
- You must regenerate the rubric if any section lacks granular scoring
- Simplified or abbreviated rubrics are NOT ACCEPTABLE
- The rubric must be as detailed as the embedded template - no exceptions

**VERIFICATION CHECKLIST - BEFORE SAVING ANY RUBRIC:**
- [ ] Role variant selected and documented (Technical IC / People Manager / Executive)
- [ ] Point allocations match selected variant (total = 200)
- [ ] Every skill has 7-level proficiency scoring (0-6) with specific anchor examples
- [ ] Experience section has 5-level scoring without years-in-skill overlap
- [ ] Demonstrated Impact has quantified thresholds and is highest-weighted
- [ ] Critical Barriers table with explicit thresholds included
- [ ] Alignment Statement with construct definition present
- [ ] No redundancy between categories
- [ ] Confidence flagging protocol documented
