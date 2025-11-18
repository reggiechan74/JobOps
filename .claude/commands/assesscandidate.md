---
description: Assess candidate against job posting using pre-created scoring rubric
argument-hint: <rubric-file> <job-posting-file>
---

You are acting as both an expert HR recruiter and a domain-knowledgeable hiring manager. Perform a comprehensive assessment of the candidate's work history against the specified job posting using a pre-created scoring rubric.

## 🔥 CRITICAL OUTPUT CONSTRAINT 🔥

**MAXIMUM OUTPUT LIMIT: 20,000 TOKENS**

Your complete assessment report MUST NOT EXCEED 20,000 tokens. This is a hard limit.

**Strategies to stay within limit:**
- Be concise and focused in your analysis
- Use bullet points instead of lengthy paragraphs where appropriate
- Prioritize evidence-based scoring over verbose explanations
- Keep executive summary brief (2-3 sentences maximum)
- Focus on the most impactful evidence for each score
- Eliminate redundancy across sections

## Your Task

Use the existing scoring rubric from {{ARG1}} to evaluate the candidate against the {{ARG2}} job posting, providing a detailed assessment report with scores and evidence.

## Step-by-Step Process

### Load Required Templates

**CRITICAL: Read these framework templates before proceeding:**
- `.claude/templates/evidence_verification_framework.md` - Evidence-based scoring protocols
- `.claude/templates/assessment_report_structure.md` - Assessment report format

These templates define the mandatory verification and reporting standards.

### YAML front matter for assessment output
The generated assessment in `OutputResumes/Assessment_*` must start with:

```yaml
---
job_file: Job_Postings/{{ARG2}}
rubric_file: Scoring_Rubrics/{{ARG1}}
role: <role title>
company: <company name>
candidate: <full candidate name>
generated_by: /assesscandidate
generated_on: <ISO8601 timestamp>
output_type: assessment
status: draft
version: 1.0
overall_score: <XX/100>
---
```

Insert this block before any headings and update timestamps, scores, and versioning on reruns.

### 1. Generate Candidate Profile (Context Optimization)

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
- Read the scoring rubric from `Scoring_Rubrics/{{ARG1}}` (add .md extension if needed)
- Read the job posting from `Job_Postings/{{ARG2}}` (add .md extension if needed)
- If job posting doesn't exist in Job_Postings/, check the root directory for legacy files
- Read the candidate profile from `ResumeSourceFolder/.profile/candidate_profile.json`
- **Evidence Verification Protocol**: When citing specific achievements or skills in the assessment:
  - Use line references from JSON profile's evidence fields
  - Read specific sections from source files ONLY when verification needed
  - Quote exact text from source files for all scores ≥2 points
  - Maintain traceability: JSON profile → source file → line numbers

### 3. Validate Rubric-Job Alignment
- Verify that the rubric was created for the same job posting or compatible role
- Check that the rubric includes all detailed scoring breakdowns required
- Confirm rubric completeness against the framework requirements
- Note any misalignments between rubric and job posting

### 4. Acquire Additional Domain Knowledge (if needed)
If the rubric lacks current context, supplement with web research:
- Current market conditions and salary ranges
- Recent industry developments since rubric creation
- Updated technology trends or skill requirements
- Company developments or changes since rubric generation

### 5. Apply Evidence Verification Framework

Apply the complete evidence verification framework from `.claude/templates/evidence_verification_framework.md` to all scoring decisions. This framework is mandatory for all scores ≥2 points.

Key requirements include:
- Specific CV citations with line numbers for all non-basic scores
- Domain specificity verification (no assumption of equivalency)
- Experience type classification (Direct/Adjacent/Transferable/Assumed)
- Cross-reference protocol for all scored items
- Quality control checklist validation before finalization

### 6. Perform 100-Point Assessment Using Pre-Created Rubric

Apply the existing scoring rubric systematically to evaluate the candidate:

1. **Core Technical Skills (25 pts)** - Use the specific required/preferred skills from the rubric to map against candidate evidence
2. **Relevant Experience (25 pts)** - Evaluate against the years, industry, and domain requirements defined in the rubric
3. **Key Responsibilities (20 pts)** - Match candidate experience to the primary duties extracted in the rubric
4. **Achievements & Impact (15 pts)** - Verify metrics against the expected outcomes defined in the rubric
5. **Education & Certifications (10 pts)** - Check against the specific requirements listed in the rubric
6. **Cultural Fit (5 pts)** - Assess based on the company values and work environment from the rubric

### 7. Generate Comprehensive Assessment Report

Follow the report structure defined in `.claude/templates/assessment_report_structure.md` exactly.

**CRITICAL FORMAT REQUIREMENTS:**
- Use the detailed 3-level format with sub-bullets for all scored items:
  - **Rubric Criteria Applied**: [Specific criteria used from rubric]
  - **Candidate Evidence**: [Detailed evidence mapping with CV line numbers]
  - **Score Justification**: [Why this score level was assigned]
- Include all sections from the template (Executive Summary, Detailed Scoring, Analysis, Evidence Mapping, Interview Strategy, etc.)
- Maintain command-specific content:
  - **Rubric Used**: {{ARG1}}
  - **Pre-Created Rubric Applied** section with creation date and criteria summary
  - **Rubric Application Analysis** section evaluating rubric effectiveness
  - **Audit Trail** with rubric file reference and assessment method

This is the gold standard assessment format with full rubric criteria attribution.

### 8. Save Assessment Report

**CRITICAL: Folder Structure and Timestamps**

Before saving any files:
1. **Get Current Eastern Time**: Retrieve current date/time in America/New_York timezone
2. **Create Timestamped Sub-Folder**: Format as `YYYY-MM-DD_HHMMSS_[Company]_[Role]`
   - Example: `2025-11-17_143022_UniversityOfToronto_ExecutiveDirectorAssetManagement`
   - Remove spaces from company and role names; use CamelCase or underscores
   - Ensure folder is created in `/workspaces/resumeoptimizer/OutputResumes/`

**File Save Locations:**

1. **Copy Rubric to Assessment Folder**: `OutputResumes/[YYYY-MM-DD_HHMMSS]_[Company]_[Role]/Rubric_[Company]_[Role]_[Date].md`
   - Copy the pre-created rubric from Scoring_Rubrics/ to the timestamped folder
   - Ensures assessment folder is self-contained with all artifacts for audit trail

2. **Save Assessment Report**: `OutputResumes/[YYYY-MM-DD_HHMMSS]_[Company]_[Role]/Assessment_[Company]_[Role]_[Date].md`
   - Include reference to using pre-created rubric for audit trail
   - Document all scores with evidence mapping
   - Provide clear traceability between rubric criteria and candidate evaluation

**Example Full Paths:**
```
/workspaces/resumeoptimizer/OutputResumes/2025-11-17_143022_UofT_ExecutiveDirectorAssetManagement/Rubric_UofT_ExecutiveDirectorAssetManagement_20251117.md
/workspaces/resumeoptimizer/OutputResumes/2025-11-17_143022_UofT_ExecutiveDirectorAssetManagement/Assessment_UofT_ExecutiveDirectorAssetManagement_20251117.md
```

**Folder Creation Steps:**
1. Use Bash tool to get Eastern time: `TZ='America/New_York' date '+%Y-%m-%d_%H%M%S'`
2. Extract company name and role from job posting or rubric
3. Create folder path: `OutputResumes/[timestamp]_[Company]_[Role]/`
4. Create folder using Bash: `mkdir -p "OutputResumes/[timestamp]_[Company]_[Role]"`
5. Copy rubric from Scoring_Rubrics/ to timestamped folder
6. Save assessment to timestamped folder

**Provide a summary of:**
- Overall score and recommendation
- Key strengths and gaps identified
- Rubric application effectiveness
- Recommended next steps
- **File locations** for both rubric copy and assessment report

## Important Notes

- **Rubric Fidelity**: Strictly apply the pre-created rubric criteria without modification
- **Evidence-Based Scoring**: All scores must be supported by specific evidence from candidate materials
- **Rubric Validation**: Verify rubric completeness and alignment with job posting
- **Scoring Consistency**: Use exact scoring criteria and thresholds from the rubric
- **Documentation**: Maintain clear audit trail linking rubric criteria to assigned scores
- **Gap Analysis**: Identify areas where rubric may not fully capture candidate qualities
- **Objectivity**: Apply rubric criteria consistently without bias or assumptions
- **Traceability**: Document specific rubric sections used for each evaluation area
- **Quality Control**: Ensure all scoring follows the detailed breakdowns in the pre-created rubric
- **Context Consideration**: Note any changes in market/industry since rubric creation

## Validation Requirements

Before finalizing assessment:
1. **Rubric Completeness Check**: Verify the rubric contains all required detailed scoring breakdowns
2. **Alignment Verification**: Confirm rubric matches the job posting requirements
3. **Evidence Mapping**: Ensure every score has corresponding evidence from candidate materials
4. **Consistency Review**: Check that scoring follows rubric criteria without deviation
5. **Documentation Audit**: Verify clear traceability from rubric criteria to assigned scores

## Error Handling

If issues are encountered:
- **Missing Rubric**: Report if rubric file cannot be found and suggest available alternatives
- **Incomplete Rubric**: Note any missing detailed scoring breakdowns and suggest rubric regeneration
- **Misaligned Rubric**: Flag if rubric doesn't match job posting and recommend creating new rubric
- **Insufficient Candidate Data**: Document where candidate evidence is limited for fair assessment
