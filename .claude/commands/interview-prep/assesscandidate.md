---
description: Assess candidate against job posting using pre-created scoring rubric
argument-hint: <rubric-file> <job-posting-file>
---

You are acting as both an expert HR recruiter and a domain-knowledgeable hiring manager. Perform a comprehensive assessment of the candidate's work history against the specified job posting using a pre-created scoring rubric.

## CRITICAL OUTPUT CONSTRAINT

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

---

## WORKFLOW ARCHITECTURE

```
Phase 1 (Parallel batch):    Load templates + rubric + job posting (4 parallel reads)
Phase 2 (PARALLEL):          Candidate profile gen (subagent) ‖ Validate rubric alignment
Phase 3 (Sequential):        Optional domain research (if rubric is stale)
Phase 4 (Sequential, visible): Score Cat 1 → 2 → 3 → 4 → 5 → 6
Phase 5 (Sequential):        Generate report → Save files
```

**Dependency Rules:**
- Phase 2 starts after templates/rubric/posting loaded (Phase 1)
- Profile gen and rubric validation are INDEPENDENT - run concurrently
- Phase 4 WAITS for both candidate profile AND rubric validation
- Phase 5 WAITS for all scoring (Phase 4)

---

## PROGRESS TRACKING (MANDATORY)

**Before starting any work**, create all tasks for user visibility:

| # | Task Subject | activeForm |
|---|-------------|------------|
| 1 | Load templates, rubric, and job posting | Loading templates, rubric, and job posting |
| 2 | Generate candidate profile | Generating candidate profile from source materials |
| 3 | Validate rubric-job alignment | Validating rubric-job posting alignment |
| 4 | Score Technical Skills & Competencies | Scoring Technical Skills & Competencies |
| 5 | Score Relevant Experience | Scoring Relevant Experience |
| 6 | Score Key Responsibilities | Scoring Key Responsibilities alignment |
| 7 | Score Achievements & Impact | Scoring Achievements & Impact |
| 8 | Score Education & Certifications | Scoring Education & Certifications |
| 9 | Score Cultural Fit | Scoring Cultural Fit |
| 10 | Generate assessment report | Generating comprehensive assessment report |
| 11 | Save assessment files | Saving assessment and rubric files |

**Task Update Rules:**
- Mark each task `in_progress` BEFORE starting work on it
- Mark each task `completed` AFTER finishing it
- If profile exists and is fresh, mark task 2 `completed` immediately

---

## YAML FRONT MATTER

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

---

## PHASE 1: LOAD INPUTS (Parallel batch)

> **Task:** Mark task 1 `in_progress`.

**Read all files in a single parallel batch:**
- `.claude/templates/evidence_verification_framework.md` - Evidence-based scoring protocols
- `.claude/templates/assessment_report_structure.md` - Assessment report format
- `Scoring_Rubrics/{{ARG1}}` (add .md extension if needed)
- `Job_Postings/{{ARG2}}` (add .md extension if needed)

If job posting doesn't exist in Job_Postings/, check the root directory for legacy files.

> **Task:** Mark task 1 `completed`.

---

## PHASE 2: PARALLEL DATA ACQUISITION

> **CRITICAL: Dispatch both tasks simultaneously in a SINGLE message.**
> Mark tasks 2 and 3 as `in_progress` before dispatching.

### 2.1 Generate Candidate Profile (Task 2 - Subagent)

**Check for existing profile**: Look for `ResumeSourceFolder/.profile/candidate_profile.json`
- If exists and recent (<=7 days old), use it directly (mark task 2 `completed` immediately)
- If exists but stale (>7 days old), regenerate
- If doesn't exist, generate new profile

**If regeneration needed**, dispatch subagent:
```
Use Task tool with subagent_type=resume-summarizer and prompt:
"Read all files in ResumeSourceFolder/ directory and create a structured JSON
candidate profile following the schema in .claude/agents/resume-summarizer.md.
Save output to ResumeSourceFolder/.profile/candidate_profile.json and
ResumeSourceFolder/.profile/extraction_log.md"
```

### 2.2 Validate Rubric-Job Alignment (Task 3 - Main agent)

**While profile generates**, validate the rubric:
- Verify that the rubric was created for the same job posting or compatible role
- Check that the rubric includes all detailed scoring breakdowns required
- Confirm rubric completeness against the framework requirements
- Note any misalignments between rubric and job posting
- If rubric lacks current context, note areas needing supplemental domain research

> **Task:** Mark task 3 `completed` when validation is done.
> **Task:** Mark task 2 `completed` when profile subagent returns (or immediately if fresh profile exists).

---

## PHASE 3: SUPPLEMENTAL RESEARCH (if needed)

If rubric validation identified stale context or missing information:
- Current market conditions and salary ranges
- Recent industry developments since rubric creation
- Updated technology trends or skill requirements
- Company developments or changes since rubric generation

**Skip this phase if rubric is current and complete.**

---

## PHASE 4: SCORE CANDIDATE (Sequential - needs rubric + candidate data)

> **Prerequisites:** Candidate profile (task 2) AND rubric validation (task 3) must both be `completed`.

### 4.0 Load Candidate Materials

Read the candidate profile from `ResumeSourceFolder/.profile/candidate_profile.json`

**Evidence Verification Protocol**: When citing specific achievements or skills:
- Use line references from JSON profile's evidence fields
- Read specific sections from source files ONLY when verification needed
- Quote exact text from source files for all scores >=2 points
- Maintain traceability: JSON profile -> source file -> line numbers

**CRITICAL**: Apply the evidence verification framework from `.claude/templates/evidence_verification_framework.md` to all scoring decisions.

### 4.1 Score Technical Skills & Competencies (25 pts)

> **Task:** Mark task 4 `in_progress`.

Use the specific required/preferred skills from the rubric to map against candidate evidence.

> **Task:** Mark task 4 `completed`.

### 4.2 Score Relevant Experience (25 pts)

> **Task:** Mark task 5 `in_progress`.

Evaluate against the years, industry, and domain requirements defined in the rubric.

> **Task:** Mark task 5 `completed`.

### 4.3 Score Key Responsibilities (20 pts)

> **Task:** Mark task 6 `in_progress`.

Match candidate experience to the primary duties extracted in the rubric.

> **Task:** Mark task 6 `completed`.

### 4.4 Score Achievements & Impact (15 pts)

> **Task:** Mark task 7 `in_progress`.

Verify metrics against the expected outcomes defined in the rubric.

> **Task:** Mark task 7 `completed`.

### 4.5 Score Education & Certifications (10 pts)

> **Task:** Mark task 8 `in_progress`.

Check against the specific requirements listed in the rubric.

> **Task:** Mark task 8 `completed`.

### 4.6 Score Cultural Fit (5 pts)

> **Task:** Mark task 9 `in_progress`.

Assess based on the company values and work environment from the rubric.

> **Task:** Mark task 9 `completed`.

---

## PHASE 5: REPORT AND SAVE

### 5.1 Generate Comprehensive Assessment Report

> **Task:** Mark task 10 `in_progress`.

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

> **Task:** Mark task 10 `completed`.

### 5.2 Save Assessment Report

> **Task:** Mark task 11 `in_progress`.

**CRITICAL: Folder Structure and Timestamps**

Before saving any files:
1. **Get Current Eastern Time**: `TZ='America/New_York' date '+%Y-%m-%d_%H%M%S'`
2. **Create Timestamped Sub-Folder**: Format as `YYYY-MM-DD_HHMMSS_[Company]_[Role]`
   - Remove spaces; use CamelCase or underscores
   - Create in `/workspaces/resumeoptimizer/OutputResumes/`

**File Save Locations:**

1. **Copy Rubric to Assessment Folder**: `OutputResumes/[timestamp]_[Company]_[Role]/Rubric_[Company]_[Role]_[Date].md`
   - Copy the pre-created rubric from Scoring_Rubrics/ for audit trail

2. **Save Assessment Report**: `OutputResumes/[timestamp]_[Company]_[Role]/Assessment_[Company]_[Role]_[Date].md`
   - Include reference to using pre-created rubric
   - Document all scores with evidence mapping
   - Provide clear traceability between rubric criteria and candidate evaluation

**Provide a summary of:**
- Overall score and recommendation
- Key strengths and gaps identified
- Rubric application effectiveness
- Recommended next steps
- File locations for both rubric copy and assessment report

> **Task:** Mark task 11 `completed`.

---

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
