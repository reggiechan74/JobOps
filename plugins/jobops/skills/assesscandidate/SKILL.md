---
name: assesscandidate
description: Assess candidate against job posting using pre-created scoring rubric
disable-model-invocation: true
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if this skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if this skill has jurisdiction-sensitive logic (crisis/legal skills accept `--jurisdiction=<ISO-3166-2>` to override).

## Templates

For each template used by this skill, resolve the full path as:

  {config.templates.base_dir}/{config.templates.active.<template_name>}/<filename>

Templates referenced by this skill: assessment_rubric_framework, assessment_report_structure, evidence_verification_framework

## Application Path Resolution

This skill writes to a per-application folder. Before writing any output:

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the job-posting filename, or honor `--app=<slug>` if supplied.
2. Compose the app folder: `{config.directories.applications_root}/{app_slug}/`.
3. Resolve this skill's sub-folder by category:
   - resume-development (buildresume, provenance-check) → `resume/`
   - cover-letter (coverletter) → `cover-letter/`
   - rubric / assessment (createrubric, assessjob, assesscandidate, auditjobposting) → `assessment/`
   - briefing / interview prep (briefing, interviewprep) → `interview/`
4. If the app folder does not exist, `mkdir -p` it, then copy
   `{config.directories.job_postings}/{filename}` → `{app_slug}/job_posting.md`
   so the pinned JD cannot silently change under completed work.
5. Exact-slug collisions (same Company+Role+Date) are not auto-suffixed. If the folder
   already contains the same output type, require the user to pass `--app=<distinct-slug>`.

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
Phase 2 (Sequential):        Read source files directly + Validate rubric alignment
Phase 3 (Sequential):        Optional domain research (if rubric is stale)
Phase 4 (Sequential, visible): Score Cat 1 → 2 → 3 → 4 → 5 → 6
Phase 5 (Sequential):        Generate report → Save files
```

**Dependency Rules:**
- Phase 2 starts after templates/rubric/posting loaded (Phase 1)
- Source reads and rubric validation can be done together in Phase 2
- Phase 4 WAITS for source reads AND rubric validation
- Phase 5 WAITS for all scoring (Phase 4)

---

## PROGRESS TRACKING (MANDATORY)

**Before starting any work**, create all tasks for user visibility:

| # | Task Subject | activeForm |
|---|-------------|------------|
| 1 | Load templates, rubric, and job posting | Loading templates, rubric, and job posting |
| 2 | Read candidate source files | Reading candidate source files directly |
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
- If source is a single file, mark task 2 `completed` immediately after reading it

---

## YAML FRONT MATTER

The generated assessment in `{applications_root}/{app_slug}/assessment/assessment.md` must start with:

```yaml
---
job_file: {config.directories.job_postings}/{{ARG2}}
rubric_file: {{ARG1}}
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
- `{config.templates.base_dir}/{config.templates.active[evidence_verification_framework]}/evidence_verification_framework.md` - Evidence-based scoring protocols
- `{config.templates.base_dir}/{config.templates.active[assessment_report_structure]}/assessment_report_structure.md` - Assessment report format
- `{{ARG1}}` — pre-created rubric file. Accept either an absolute path, a path relative to `{applications_root}/{app_slug}/assessment/rubric.md`, or a bare filename to resolve inside the current app folder.
- `{config.directories.job_postings}/{{ARG2}}` (add .md extension if needed)

If job posting doesn't exist in {config.directories.job_postings}/, check the root directory for legacy files.

> **Task:** Mark task 1 `completed`.

---

## PHASE 2: DATA ACQUISITION

> Mark tasks 2 and 3 as `in_progress`.

### Legacy Candidate Profile Artifact Check

Before reading candidate materials, set `source_path` to `{config.directories.resume_source}` and check for deprecated profile JSON artifacts.

- If `source_path` is a single-file source named `candidate_profile.json`, stop before assessment. Tell the candidate this legacy JSON profile is deprecated and assessments now read source markdown directly. Instruct the candidate to delete the deprecated file and provide a source markdown file or folder instead. Offer to delete the legacy file for them.
- If `source_path` is a directory, check for these legacy artifacts before reading candidate materials:
  - `{source_path}/.profile/candidate_profile.json`
  - `{source_path}/candidate_profile.json`
- If either legacy artifact exists, stop before assessment. List the exact file path(s), tell the candidate each deprecated candidate profile JSON file must be deleted because JobOps now reads source markdown directly, and offer to delete the legacy artifact(s) for them.
- If the candidate approves deletion, delete only the listed legacy artifact file(s), re-run this check, then continue only if no legacy profile artifact remains.
- Never load, summarize, or score from `candidate_profile.json`.

### 2.1 Read Candidate Source Files (Task 2)

Determine source structure:
- **Single-file source** (resume.md or similar passed by user): read that file directly. Skip everything below; you have the candidate data in-hand.
- **Folder source** (`{config.directories.resume_source}/`): read these files directly for the rubric scoring you are about to do:
  - `Identity/Name.md`, `Identity/CurrentRole.md` (candidate identity)
  - `Technology/TechStack.md` (skill inventory)
  - `Technology/Certifications.md` (active credentials)
  - `WorkHistory/*.md` (roles, achievements, scope)
  - `Projects/*.md` if present (case studies)
  - `Education/*.md` if present
  - `Preferences/Vision.md` (cultural-fit signals)

If a required file is missing, prompt the user to run `/jobops:audit-source` and stop. Do NOT attempt to generate or load `candidate_profile.json` — that artifact is removed in v2.2.0.

Token budget: most rubric scoring needs ~30K of source markdown loaded at once for a thorough assessment. Read what you need; do not pre-summarize.

> **Task:** Mark task 2 `completed` when source reads are done.

### 2.2 Validate Rubric-Job Alignment (Task 3 - Main agent)

Validate the rubric:
- Verify that the rubric was created for the same job posting or compatible role
- Check that the rubric includes all detailed scoring breakdowns required
- Confirm rubric completeness against the framework requirements
- Note any misalignments between rubric and job posting
- If rubric lacks current context, note areas needing supplemental domain research

> **Task:** Mark task 3 `completed` when validation is done.

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

> **Prerequisites:** Source file reads (task 2) AND rubric validation (task 3) must both be `completed`.

### 4.0 Score Against Source Files

Score each rubric category against the source files you read in Phase 2. For each score, cite the specific source `{filepath}:{line_number}` you anchored on. Do not invent enums (proficiency_level, company_size, impact_category) — judge from the source prose with citation.

**CRITICAL**: Apply the evidence verification framework from the evidence verification framework template to all scoring decisions.

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

Follow the report structure defined in the assessment report structure template exactly.

**CRITICAL FORMAT REQUIREMENTS:**
- Use the detailed 3-level format with sub-bullets for all scored items:
  - **Rubric Criteria Applied**: [Specific criteria used from rubric]
  - **Candidate Evidence**: [Detailed evidence mapping with source file citations {filepath}:{line_number}]
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

**File Save Location (app-centric layout):**

Resolve `{app_slug}` per the Application Path Resolution protocol at the top of this skill, and ensure `{applications_root}/{app_slug}/assessment/` exists (`mkdir -p`).

1. **Pin the rubric inside the app folder**: if `{{ARG1}}` points to a rubric outside the app folder, copy it to `{applications_root}/{app_slug}/assessment/rubric.md` (if one is not already pinned there). The pinned rubric is the authoritative copy used for audit.

2. **Save Assessment Report**: `{applications_root}/{app_slug}/assessment/assessment.md`
   - Include reference to using pre-created rubric (use the pinned path)
   - Document all scores with evidence mapping
   - Provide clear traceability between rubric criteria and candidate evaluation

The app folder itself is the self-contained audit container — no timestamped audit sub-folder is needed.

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
