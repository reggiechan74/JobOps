---
description: Assess candidate against job posting using pre-created scoring rubric
signature: "<rubric-file> <job-posting-file>"
project: true
---

You are acting as both an expert HR recruiter and a domain-knowledgeable hiring manager. Perform a comprehensive assessment of the candidate's work history against the specified job posting using a pre-created scoring rubric.

## Your Task

Use the existing scoring rubric from {{ARG1}} to evaluate the candidate against the {{ARG2}} job posting, providing a detailed assessment report with scores and evidence.

## Step-by-Step Process

### 1. Load Required Documents
- Read the scoring rubric from `Scoring_Rubrics/{{ARG1}}` (add .md extension if needed)
- Read the job posting from `Job_Postings/{{ARG2}}` (add .md extension if needed)
- If job posting doesn't exist in Job_Postings/, check the root directory for legacy files
- Load ALL candidate work history files from `ResumeSourceFolder/` directory:
  - Use glob pattern to find all markdown files in ResumeSourceFolder/
  - Read each file to build complete candidate profile
  - Include all CV, resume, technology capability, and work history documents

### 2. Validate Rubric-Job Alignment
- Verify that the rubric was created for the same job posting or compatible role
- Check that the rubric includes all detailed scoring breakdowns required
- Confirm rubric completeness against the framework requirements
- Note any misalignments between rubric and job posting

### 3. Acquire Additional Domain Knowledge (if needed)
If the rubric lacks current context, supplement with web research:
- Current market conditions and salary ranges
- Recent industry developments since rubric creation
- Updated technology trends or skill requirements
- Company developments or changes since rubric generation

### 4. Perform 100-Point Assessment Using Pre-Created Rubric

Apply the existing scoring rubric systematically to evaluate the candidate:

1. **Core Technical Skills (25 pts)** - Use the specific required/preferred skills from the rubric to map against candidate evidence
2. **Relevant Experience (25 pts)** - Evaluate against the years, industry, and domain requirements defined in the rubric
3. **Key Responsibilities (20 pts)** - Match candidate experience to the primary duties extracted in the rubric
4. **Achievements & Impact (15 pts)** - Verify metrics against the expected outcomes defined in the rubric
5. **Education & Certifications (10 pts)** - Check against the specific requirements listed in the rubric
6. **Cultural Fit (5 pts)** - Assess based on the company values and work environment from the rubric

### 5. Generate Comprehensive Assessment Report

Create assessment report using the pre-created rubric scores:

```markdown
# Candidate Assessment: [Role] at [Company]
**Assessment Date:** [Date]
**Rubric Used:** {{ARG1}}
**Job Posting:** {{ARG2}}

## Executive Summary
[2-3 sentence overview of fit and recommendation based on rubric criteria]

## Overall Score: [XX/100]
*Scored against pre-created rubric: {{ARG1}}*

### Pre-Created Rubric Applied
**Rubric Creation Date:** [Date from rubric file]
**Job-Specific Criteria:** [Brief summary of rubric focus areas]

### Detailed Scoring Breakdown

#### 1. Technical Skills & Competencies: [XX/25]
**Required Skills (15 points total) - From Pre-Created Rubric:**
- [Skill 1 from rubric]: [X/3] - [Evidence from candidate's history]
  - Rubric Criteria Applied: [Specific criteria used from rubric]
  - Candidate Evidence: [Detailed evidence mapping]
  - Score Justification: [Why this score level was assigned]

- [Skill 2 from rubric]: [X/3] - [Evidence from candidate's history]
  - Rubric Criteria Applied: [Specific criteria used from rubric]
  - Candidate Evidence: [Detailed evidence mapping]
  - Score Justification: [Why this score level was assigned]

- [Skill 3 from rubric]: [X/3] - [Evidence from candidate's history]
  - Rubric Criteria Applied: [Specific criteria used from rubric]
  - Candidate Evidence: [Detailed evidence mapping]
  - Score Justification: [Why this score level was assigned]

- [Skill 4 from rubric]: [X/3] - [Evidence from candidate's history]
  - Rubric Criteria Applied: [Specific criteria used from rubric]
  - Candidate Evidence: [Detailed evidence mapping]
  - Score Justification: [Why this score level was assigned]

- [Skill 5 from rubric]: [X/3] - [Evidence from candidate's history]
  - Rubric Criteria Applied: [Specific criteria used from rubric]
  - Candidate Evidence: [Detailed evidence mapping]
  - Score Justification: [Why this score level was assigned]

**Preferred Skills (10 points total) - From Pre-Created Rubric:**
- [Preferred 1 from rubric]: [X/2] - [Evidence and reasoning]
  - Rubric Criteria Applied: [Specific criteria used]
  - Score Justification: [Reasoning for score]

- [Preferred 2 from rubric]: [X/2] - [Evidence and reasoning]
  - Rubric Criteria Applied: [Specific criteria used]
  - Score Justification: [Reasoning for score]

- [Preferred 3 from rubric]: [X/2] - [Evidence and reasoning]
  - Rubric Criteria Applied: [Specific criteria used]
  - Score Justification: [Reasoning for score]

- [Preferred 4 from rubric]: [X/2] - [Evidence and reasoning]
  - Rubric Criteria Applied: [Specific criteria used]
  - Score Justification: [Reasoning for score]

- [Preferred 5 from rubric]: [X/2] - [Evidence and reasoning]
  - Rubric Criteria Applied: [Specific criteria used]
  - Score Justification: [Reasoning for score]

#### 2. Relevant Experience: [XX/25]
**Years of Experience (10 points) - Per Rubric Requirements:**
- Total Years: [X/5] - [Actual years vs rubric requirement]
  - Rubric Requirement: [Years specified in rubric]
  - Candidate Experience: [Actual experience]
  - Score Level: [Which rubric level achieved]

- Relevant Years: [X/5] - [Directly applicable experience]
  - Rubric Criteria: [Relevance criteria from rubric]
  - Assessment: [How candidate meets criteria]

**Industry/Domain Experience (10 points) - Per Rubric Definitions:**
- Industry Match: [X/5] - [Specific industry alignment vs rubric]
  - Required Industry: [From rubric]
  - Candidate Industry: [Actual background]
  - Alignment Level: [Which scoring tier achieved]

- Domain Knowledge: [X/5] - [Technical domain expertise vs rubric]
  - Required Domains: [From rubric]
  - Candidate Domains: [Demonstrated expertise]
  - Score Justification: [Evidence for score level]

**Role-Specific Experience (5 points) - Per Rubric Standards:**
- Similar Roles: [X/5] - [Comparable position experience vs rubric]
  - Rubric Requirements: [Role criteria from rubric]
  - Candidate Roles: [Previous positions]
  - Overlap Assessment: [Percentage match analysis]

#### 3. Key Responsibilities: [XX/20]
**Primary Duties Match (12 points) - From Pre-Created Rubric:**
- [Responsibility 1 from rubric]: [X/3] - [Evidence of capability]
  - Rubric Expectation: [Specific duty requirement]
  - Candidate Evidence: [Demonstrated experience]
  - Score Level: [Expert/Proficient/Basic/None explanation]

- [Responsibility 2 from rubric]: [X/3] - [Evidence of capability]
  - Rubric Expectation: [Specific duty requirement]
  - Candidate Evidence: [Demonstrated experience]
  - Score Level: [Expert/Proficient/Basic/None explanation]

- [Responsibility 3 from rubric]: [X/3] - [Evidence of capability]
  - Rubric Expectation: [Specific duty requirement]
  - Candidate Evidence: [Demonstrated experience]
  - Score Level: [Expert/Proficient/Basic/None explanation]

- [Responsibility 4 from rubric]: [X/3] - [Evidence of capability]
  - Rubric Expectation: [Specific duty requirement]
  - Candidate Evidence: [Demonstrated experience]
  - Score Level: [Expert/Proficient/Basic/None explanation]

**Scope & Complexity (8 points) - Per Rubric Framework:**
- Team Size/Budget: [X/4] - [Scale of responsibility vs rubric metrics]
  - Rubric Thresholds: [Specific metrics from rubric]
  - Candidate Scale: [Actual scope managed]
  - Score Tier: [Which level achieved with evidence]

- Project Complexity: [X/4] - [Technical/business complexity vs rubric]
  - Rubric Standards: [Complexity criteria from rubric]
  - Candidate Projects: [Demonstrated complexity]
  - Assessment: [How candidate meets criteria]

#### 4. Achievements & Impact: [XX/15]
**Quantifiable Results (10 points) - Per Rubric Metrics:**
- Measurable Outcomes: [X/5] - [Specific metrics vs rubric thresholds]
  - Rubric Thresholds: [Quantitative criteria from rubric]
  - Candidate Achievements: [Actual metrics delivered]
  - Score Level: [Which tier achieved]

- Business Impact: [X/5] - [Revenue, cost, efficiency gains vs rubric]
  - Rubric Standards: [Impact criteria from rubric]
  - Candidate Impact: [Demonstrated business results]
  - Score Justification: [Evidence for score level]

**Innovation & Leadership (5 points) - Per Rubric Expectations:**
- Innovation: [X/2.5] - [New approaches, technologies vs rubric]
  - Rubric Criteria: [Innovation expectations from rubric]
  - Candidate Innovation: [Demonstrated innovation]
  - Score Assessment: [Level achieved]

- Leadership: [X/2.5] - [Team leadership, mentoring vs rubric]
  - Rubric Standards: [Leadership criteria from rubric]
  - Candidate Leadership: [Leadership evidence]
  - Score Rationale: [Justification for score]

#### 5. Education & Certifications: [XX/10]
**Education Requirements (6 points) - Per Rubric Standards:**
- Degree Level: [X/3] - [Educational achievement vs rubric requirement]
  - Rubric Requirement: [Degree requirement from rubric]
  - Candidate Education: [Actual educational background]
  - Score Level: [Which tier achieved]

- Field Relevance: [X/3] - [Subject matter alignment vs rubric]
  - Required Field: [From rubric]
  - Candidate Field: [Actual major/specialization]
  - Relevance Assessment: [Alignment evaluation]

**Certifications (4 points) - Per Rubric Specifications:**
- Required Certs: [X/2] - [Must-have certifications vs rubric]
  - Rubric Requirements: [Required certifications from rubric]
  - Candidate Certifications: [Current certifications held]
  - Score Assessment: [Compliance level]

- Preferred Certs: [X/2] - [Nice-to-have certifications vs rubric]
  - Rubric Preferences: [Preferred certifications from rubric]
  - Additional Certifications: [Extra certifications held]
  - Value Assessment: [Additional value provided]

#### 6. Cultural Fit: [XX/5]
**Communication Skills (3 points) - Per Rubric Criteria:**
- Written Communication: [X/1.5] - [Evidence vs rubric standards]
  - Rubric Expectations: [Communication criteria from rubric]
  - Candidate Evidence: [Demonstrated communication skills]
  - Score Level: [Assessment against criteria]

- Verbal/Presentation: [X/1.5] - [Inferred from experience vs rubric]
  - Rubric Standards: [Presentation criteria from rubric]
  - Candidate Indicators: [Experience suggesting capability]
  - Assessment: [Evaluation against standards]

**Values Alignment (2 points) - Per Rubric Framework:**
- Company Values: [X/2] - [Alignment vs rubric values]
  - Company Values: [From rubric]
  - Candidate Alignment: [Evidence of value alignment]
  - Score Justification: [Assessment reasoning]

## Detailed Analysis

### Strengths
[Bullet points with specific evidence from work history mapped to rubric criteria]

### Gaps & Concerns
[Identified gaps with risk levels: High/Medium/Low based on rubric requirements]

### Rubric-Specific Insights
[Analysis specific to the pre-created rubric criteria and thresholds]

## Evidence Mapping

### Technical Skills Match
| Required Skill (from Rubric) | Candidate Evidence | Experience Years | Rubric Score | Assigned Score |
|---|---|---|---|---|
[Detailed skill mapping table using rubric criteria]

### Key Achievements Relevant to Role
[Top 5 achievements with metrics mapped to rubric expectations]

## Interview Strategy

### Must-Probe Areas (Based on Rubric)
[Technical and behavioral areas requiring validation per rubric criteria]

### Recommended Assessments
[Specific technical challenges or case studies based on rubric requirements]

## Rubric Application Analysis

### Rubric Effectiveness
- **Rubric Completeness**: [Assessment of how well rubric covered evaluation needs]
- **Criteria Clarity**: [How clear and measurable the rubric criteria were]
- **Score Differentiation**: [How well rubric criteria distinguished performance levels]

### Rubric vs Candidate Alignment
- **Strong Matches**: [Areas where candidate clearly met rubric criteria]
- **Unclear Areas**: [Where rubric criteria were difficult to apply]
- **Missing Criteria**: [Candidate qualities not captured by rubric]

## Competitive Analysis
[How candidate compares to typical applicants based on rubric standards and market research]

## Hiring Recommendation

[ ] **90-100**: Exceptional candidate - exceeds requirements significantly
[ ] **80-89**: Excellent candidate - strong match, minor gaps if any
[ ] **70-79**: Good candidate - solid match, some development areas
[ ] **60-69**: Potential candidate - meets core requirements, notable gaps
[ ] **50-59**: Borderline candidate - significant gaps, consider if high potential
[ ] **Below 50**: Not recommended - major gaps in critical areas

### Justification
[Evidence-based reasoning using rubric criteria and scoring]

### Rubric-Based Decision Factors
[Key factors from rubric that influenced recommendation]

## Next Steps
[Specific recommendations for hiring team based on rubric assessment]

## Audit Trail
- **Rubric File**: {{ARG1}}
- **Job Posting**: {{ARG2}}
- **Assessment Method**: Pre-created rubric application
- **Evaluator Notes**: [Any deviations from rubric or special considerations]
```

### 6. Save Assessment Report

Save the assessment report to: `OutputResumes/Assessment_[Company]_[Role]_[Date].md`

Include in the filename reference to using pre-created rubric for audit trail.

Provide a summary of:
- Overall score and recommendation
- Key strengths and gaps identified
- Rubric application effectiveness
- Recommended next steps

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