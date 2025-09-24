---
description: Perform expert HR assessment of candidate against job posting with domain knowledge
signature: "<job-posting-file>"
project: true
---

You are acting as both an expert HR recruiter and a domain-knowledgeable hiring manager. Perform a comprehensive assessment of the candidate's work history against the specified job posting.

## Your Task

Evaluate the candidate from {{ARG1}} job posting using the scoring rubric and provide a detailed assessment report.

## Step-by-Step Process

### 1. Load Required Documents
- Load the scoring rubric from `SourceMaterial/jobscoringrubric.md`
- Read the job posting from `Job_Postings/{{ARG1}}` (add .md extension if needed)
- Load ALL candidate work history files from `ResumeSourceFolder/` directory:
  - Use glob pattern to find all markdown files in ResumeSourceFolder/
  - Read each file to build complete candidate profile
  - Include all CV, resume, technology capability, and work history documents

### 2. Acquire Domain Knowledge
Use web research to understand:
- Industry standards and role expectations
- Required vs nice-to-have skills for this specific role
- Typical responsibilities and seniority indicators
- Company context, culture, and technology stack
- Current market conditions and competitive landscape

### 3. Perform 100-Point Assessment

Apply the scoring rubric systematically:
1. **Core Technical Skills (25 pts)** - Map each skill to specific evidence
2. **Relevant Experience (25 pts)** - Calculate years, industry, role alignment
3. **Key Responsibilities (20 pts)** - Match duties and scope
4. **Achievements & Impact (15 pts)** - Verify metrics and innovation
5. **Education & Certifications (10 pts)** - Check requirements
6. **Cultural Fit (5 pts)** - Assess communication and values

### 4. Generate Comprehensive Report

Create assessment report with:
```markdown
# Candidate Assessment: [Role] at [Company]

## Executive Summary
[2-3 sentence overview of fit and recommendation]

## Overall Score: [XX/100]

### Detailed Scoring Breakdown

#### 1. Technical Skills & Competencies: [XX/25]
**Required Skills (15 points total):**
- [Skill 1]: [X/3] - [Evidence and reasoning]
- [Skill 2]: [X/3] - [Evidence and reasoning]
- [Skill 3]: [X/3] - [Evidence and reasoning]
- [Skill 4]: [X/3] - [Evidence and reasoning]
- [Skill 5]: [X/3] - [Evidence and reasoning]

**Preferred Skills (10 points total):**
- [Preferred Skill 1]: [X/2] - [Evidence and reasoning]
- [Preferred Skill 2]: [X/2] - [Evidence and reasoning]
- [Preferred Skill 3]: [X/2] - [Evidence and reasoning]
- [Preferred Skill 4]: [X/2] - [Evidence and reasoning]
- [Preferred Skill 5]: [X/2] - [Evidence and reasoning]

#### 2. Relevant Experience: [XX/25]
**Years of Experience (10 points):**
- Total Years: [X/5] - [Actual years vs requirement]
- Relevant Years: [X/5] - [Directly applicable experience]

**Industry/Domain Experience (10 points):**
- Industry Match: [X/5] - [Specific industry alignment]
- Domain Knowledge: [X/5] - [Technical domain expertise]

**Role-Specific Experience (5 points):**
- Similar Roles: [X/5] - [Comparable position experience]

#### 3. Key Responsibilities: [XX/20]
**Primary Duties Match (12 points):**
- [Responsibility 1]: [X/3] - [Evidence of capability]
- [Responsibility 2]: [X/3] - [Evidence of capability]
- [Responsibility 3]: [X/3] - [Evidence of capability]
- [Responsibility 4]: [X/3] - [Evidence of capability]

**Scope & Complexity (8 points):**
- Team Size/Budget: [X/4] - [Scale of responsibility]
- Project Complexity: [X/4] - [Technical/business complexity]

#### 4. Achievements & Impact: [XX/15]
**Quantifiable Results (10 points):**
- Measurable Outcomes: [X/5] - [Specific metrics achieved]
- Business Impact: [X/5] - [Revenue, cost, efficiency gains]

**Innovation & Leadership (5 points):**
- Innovation: [X/2.5] - [New approaches, technologies]
- Leadership: [X/2.5] - [Team leadership, mentoring]

#### 5. Education & Certifications: [XX/10]
**Education Requirements (6 points):**
- Degree Level: [X/3] - [Bachelor's/Master's/PhD alignment]
- Field Relevance: [X/3] - [Subject matter alignment]

**Certifications (4 points):**
- Required Certs: [X/2] - [Must-have certifications]
- Preferred Certs: [X/2] - [Nice-to-have certifications]

#### 6. Cultural Fit: [XX/5]
**Communication Skills (3 points):**
- Written Communication: [X/1.5] - [Evidence from documents]
- Verbal/Presentation: [X/1.5] - [Inferred from experience]

**Values Alignment (2 points):**
- Company Values: [X/2] - [Alignment with stated values]

## Detailed Analysis

### Strengths
[Bullet points with specific evidence from work history]

### Gaps & Concerns
[Identified gaps with risk levels: High/Medium/Low]

### Domain-Specific Insights
[Industry knowledge and market positioning from research]

## Evidence Mapping

### Technical Skills Match
| Required Skill | Candidate Evidence | Experience | Score |
|---|---|---|---|
[Detailed skill mapping table]

### Key Achievements Relevant to Role
[Top 5 achievements with metrics and context]

## Interview Strategy

### Must-Probe Areas
[Technical and behavioral areas requiring validation]

### Recommended Assessments
[Specific technical challenges or case studies]

## Competitive Analysis
[How candidate compares to typical applicants based on market research]

## Hiring Recommendation
[ ] Strongly Recommend - Exceptional fit
[ ] Recommend - Strong candidate with minor gaps
[ ] Consider - Meets requirements with development needs
[ ] Not Recommended - Significant gaps

### Justification
[Evidence-based reasoning]

## Next Steps
[Specific recommendations for hiring team]
```

### 5. Save Assessment
Save report to: `OutputResumes/Assessment_[Company]_[Role]_[Date].md`

## Important Notes
- Be objective and evidence-based in scoring
- Use web research to calibrate expectations
- Flag any inconsistencies or credibility concerns
- Provide actionable insights for hiring decisions
- Note confidence levels where evidence is limited