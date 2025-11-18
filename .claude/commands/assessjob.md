---
description: Perform expert HR assessment of candidate against job posting with domain knowledge
argument-hint: <job-posting-file>
---

You are acting as both an expert HR recruiter and a domain-knowledgeable hiring manager. Perform a comprehensive assessment of the candidate's work history against the specified job posting.

## 🔥 CRITICAL OUTPUT CONSTRAINT 🔥

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

## Step-by-Step Process

### Load Required Templates

**CRITICAL: Read these framework templates before proceeding:**
- `.claude/templates/assessment_rubric_framework.md` - Master 100-point rubric structure
- `.claude/templates/evidence_verification_framework.md` - Evidence-based scoring protocols
- `.claude/templates/assessment_report_structure.md` - Assessment report format

These templates define the mandatory structure for rubrics and assessments.

### YAML front matter for generated files
Every markdown artifact you create (rubric and assessment report) must start with YAML metadata populated with real values.

- **Rubric file** (`Scoring_Rubrics/Rubric_*`):
  ```yaml
  ---
  job_file: Job_Postings/{{ARG1}}
  role: <role title>
  company: <company name>
  generated_by: /assessjob rubric
  generated_on: <ISO8601 timestamp>
  output_type: rubric
  status: final
  version: 1.0
  ---
  ```
- **Assessment file** (`OutputResumes/Assessment_*`):
  ```yaml
  ---
  job_file: Job_Postings/{{ARG1}}
  rubric_file: Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md
  role: <role title>
  company: <company name>
  candidate: <full candidate name>
  generated_by: /assessjob
  generated_on: <ISO8601 timestamp>
  output_type: assessment
  status: draft
  version: 1.0
  overall_score: <XX/100>
  ---
  ```

Insert the appropriate block before any headings, updating timestamps and scores, and bump `version` if you rerun the analysis.

### 1. Generate Candidate Profile (Context Optimization)
✓ Acquiring compressed candidate intelligence from source materials

**CRITICAL - Context Window Optimization**: Before loading full source materials, generate a structured candidate profile:

1. **Check for existing profile**: Look for `ResumeSourceFolder/.profile/candidate_profile.json`
   - If exists and recent (≤7 days old), use it directly
   - If exists but stale (>7 days old), regenerate
   - If doesn't exist, generate new profile

2. **Generate profile using resume-summarizer agent**:
   ```
   Use Task tool with subagent_type=general-purpose and prompt:
   "You are the resume-summarizer agent. Read all files in ResumeSourceFolder/ directory and create a structured JSON candidate profile following the schema in .claude/agents/resume-summarizer.md. Save output to ResumeSourceFolder/.profile/candidate_profile.json and ResumeSourceFolder/.profile/extraction_log.md"
   ```

3. **Expected outcome**:
   - JSON profile created: `ResumeSourceFolder/.profile/candidate_profile.json` (8K-10K tokens)
   - Extraction log created: `ResumeSourceFolder/.profile/extraction_log.md`
   - Token savings: 42K-72K tokens (85-90% reduction from loading 15 source files)

### 2. Load Required Documents
✓ Acquiring target coordinates from `Job_Postings/{{ARG1}}`
✓ Loading compressed candidate profile from structured JSON

- Read the job posting from `Job_Postings/{{ARG1}}` (add .md extension if needed)
- Read the candidate profile from `ResumeSourceFolder/.profile/candidate_profile.json`
- **Evidence Verification Protocol**: When citing specific achievements or skills in the assessment:
  - Use line references from JSON profile's evidence fields
  - Read specific sections from source files ONLY when verification needed
  - Quote exact text from source files for all scores ≥2 points
  - Maintain traceability: JSON profile → source file → line numbers

### 3. Create Dynamic Job-Specific Rubric
✓ Establishing targeting criteria and reconnaissance framework
Analyze the job posting to extract and categorize requirements, then create a customized scoring rubric.

**🚨 MANDATORY DETAILED SCORING REQUIREMENT 🚨**
The dynamic rubric you create MUST include granular point allocation for every single criterion. This is NON-NEGOTIABLE. You cannot create simplified rubrics or skip detailed breakdowns. Every skill, every experience level, every responsibility MUST have the complete 4-level or 5-level scoring framework with specific, measurable criteria.

#### Parse Job Posting to Identify:
- **Required Technical Skills**: List all must-have technical competencies
- **Preferred Skills**: List nice-to-have skills and technologies
- **Experience Requirements**: Years, industry, domain specifics
- **Key Responsibilities**: Primary duties and scope expectations
- **Education/Certifications**: Required and preferred credentials
- **Soft Skills/Cultural Fit**: Communication, collaboration, values

#### Generate Custom Rubric Structure:

**CRITICAL REQUIREMENT**: Follow the structure defined in `.claude/templates/assessment_rubric_framework.md` exactly. Customize all `[bracketed]` content with job-specific requirements from the job posting.

**ENFORCEMENT CHECK**: After creating the rubric, verify that EVERY section includes:
- Individual skill breakdowns with 4-level scoring (Expert/Proficient/Basic/None)
- Specific measurable criteria for each level
- Overall assessment ranges
- Complete evaluation frameworks for scope & complexity
- Detailed scoring breakdowns for achievements & impact

**FAILURE TO INCLUDE DETAILED BREAKDOWNS VIOLATES THE COMMAND REQUIREMENTS**

#### Save Dynamic Rubric:
**MANDATORY QUALITY CHECK**: Before saving, verify the rubric includes ALL required detailed breakdowns:
- ✅ Every required skill has Expert/Proficient/Basic/None scoring with specific criteria
- ✅ Every preferred skill has Strong/Basic/None scoring with role-specific thresholds
- ✅ Experience sections have 5-level detailed scoring breakdowns
- ✅ Primary duties have Expert/Proficient/Basic/None frameworks
- ✅ Scope & complexity has detailed evaluation framework with specific metrics
- ✅ Achievements & impact has 5-level scoring with quantitative thresholds
- ✅ Education & certifications have detailed multi-level scoring
- ✅ Cultural fit has comprehensive evaluation criteria and scoring levels

**IF ANY SECTION LACKS DETAILED BREAKDOWNS, THE RUBRIC IS INCOMPLETE AND MUST BE REGENERATED**

Save the generated rubric to: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`
Also save a reference copy with the assessment for audit trail

### 4. Acquire Domain Knowledge
✓ Conducting intelligence gathering on target domain and industry
Use web research to understand:
- Industry standards and role expectations for the specific position
- Required vs nice-to-have skills based on current market standards
- Typical responsibilities and seniority indicators in this domain
- Company context, culture, technology stack, and recent developments
- Current market conditions, salary ranges, and competitive landscape
- Industry-specific terminology and best practices

### 4a. Apply Evidence-Based Scoring Verification

**CRITICAL**: Apply the evidence verification framework from `.claude/templates/evidence_verification_framework.md` to all scoring decisions.

### 5. Perform 100-Point Assessment Using Dynamic Rubric
✓ Executing reconnaissance assessment against targeting criteria

Apply the job-specific scoring rubric you created systematically:
1. **Core Technical Skills (25 pts)** - Map each required/preferred skill from the custom rubric to specific evidence in candidate's history
2. **Relevant Experience (25 pts)** - Evaluate against the specific years, industry, and domain requirements identified in the rubric
3. **Key Responsibilities (20 pts)** - Match candidate's experience to the 4 primary duties extracted from the job posting
4. **Achievements & Impact (15 pts)** - Verify metrics against the expected outcomes defined in the rubric
5. **Education & Certifications (10 pts)** - Check against the specific requirements listed in the custom rubric
6. **Cultural Fit (5 pts)** - Assess based on the company values and work environment extracted from the posting

### 6. Generate Comprehensive Report
✓ Compiling reconnaissance report with engagement recommendation

**CRITICAL**: Follow the report structure defined in `.claude/templates/assessment_report_structure.md` exactly. Use the detailed 3-level format with rubric criteria attribution for all scores.

### 7. Save Assessment and Rubric

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
   - Document all scores with evidence mapping
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

## Validation Requirements

Before finalizing assessment:
1. **Rubric Completeness Check**: Verify the rubric contains all required detailed scoring breakdowns
2. **Alignment Verification**: Confirm rubric matches the job posting requirements
3. **Evidence Mapping**: Ensure every score has corresponding evidence from candidate materials
4. **Consistency Review**: Check that scoring follows rubric criteria without deviation
5. **Documentation Audit**: Verify clear traceability from rubric criteria to assigned scores

## Error Handling

If issues are encountered:
- **Missing Job Posting**: Report if job posting file cannot be found and suggest available alternatives
- **Incomplete Job Posting**: Note any missing critical information and document assumptions made
- **Insufficient Candidate Data**: Document where candidate evidence is limited for fair assessment
- **Domain Knowledge Gaps**: Flag areas where additional research or expertise would improve assessment accuracy

## Important Notes
- **MANDATORY DETAILED SCORING**: Every rubric MUST include the granular scoring breakdowns from the framework rubric. This is not optional - it is required.
- **Dynamic Rubric Creation**: Each job posting generates its own tailored rubric with job-specific criteria but framework-level detail
- **Complete Scoring Levels**: All sections must include the detailed point breakdowns (Expert/Proficient/Basic, Exceptional/Strong/Good, etc.)
- **Role-Specific Adaptation**: Customize the content and thresholds for the specific role while maintaining scoring structure depth
- **Consistency**: Use the template framework as the foundation but customize content to job requirements
- **Traceability**: Always reference which rubric was used for which assessment
- **Objectivity**: Be evidence-based in scoring using the job-specific criteria with detailed justification
- **Domain Research**: Use web research to understand industry-specific requirements and adapt scoring thresholds accordingly
- **Documentation**: Save both rubric and assessment for complete audit trail
- **Confidence Levels**: Note where evidence is limited or assumptions are made
- **Quality Control**: The generated rubric must be as detailed as the framework rubric - no shortcuts or simplified versions allowed

## 🔥 CRITICAL ENFORCEMENT RULES 🔥

**ABSOLUTE REQUIREMENTS FOR EVERY RUBRIC:**
1. **EVERY REQUIRED SKILL** must have Expert (3) / Proficient (2) / Basic (1) / None (0) with specific criteria
2. **EVERY PREFERRED SKILL** must have Strong (2) / Basic (1) / None (0) with role-specific thresholds
3. **EXPERIENCE SECTIONS** must have 5-level detailed breakdowns (Exceeds/Meets Plus/Meets/Near/Below)
4. **PRIMARY DUTIES** must have Expert/Proficient/Basic/None with quantitative thresholds
5. **SCOPE & COMPLEXITY** must include detailed evaluation framework with specific metrics
6. **ACHIEVEMENTS** must have 5-level scoring with quantitative thresholds (revenue, portfolio size, etc.)
7. **EDUCATION & CERTIFICATIONS** must have multi-level detailed scoring
8. **CULTURAL FIT** must have comprehensive evaluation criteria

**VIOLATION CONSEQUENCES:**
- Any rubric missing detailed breakdowns is INCOMPLETE and violates the command specification
- You must regenerate the rubric if any section lacks granular scoring
- Simplified or abbreviated rubrics are NOT ACCEPTABLE
- The rubric must be as detailed as the embedded template - no exceptions

**VERIFICATION CHECKLIST - BEFORE SAVING ANY RUBRIC:**
□ Every skill has 4-level scoring with specific, measurable criteria
□ Every section includes detailed point allocation breakdowns
□ All evaluation frameworks include quantitative thresholds
□ Overall assessment ranges are provided for each section
□ Role-specific criteria are adapted from the template structure
