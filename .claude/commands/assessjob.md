---
description: Perform expert HR assessment of candidate against job posting with domain knowledge
signature: "<job-posting-file>"
project: true
---

You are acting as both an expert HR recruiter and a domain-knowledgeable hiring manager. Perform a comprehensive assessment of the candidate's work history against the specified job posting.

## Your Task

Create a job-specific scoring rubric from the {{ARG1}} job posting, then evaluate the candidate using this customized rubric to provide a detailed assessment report.

## Step-by-Step Process

### 1. Load Required Documents
- Load the framework rubric from `Scoring_Rubrics/jobscoringrubric.md` as a template
- Read the job posting from `Job_Postings/{{ARG1}}` (add .md extension if needed)
- Load ALL candidate work history files from `ResumeSourceFolder/` directory:
  - Use glob pattern to find all markdown files in ResumeSourceFolder/
  - Read each file to build complete candidate profile
  - Include all CV, resume, technology capability, and work history documents

### 2. Create Dynamic Job-Specific Rubric
Analyze the job posting to extract and categorize requirements, then create a customized scoring rubric:

#### Parse Job Posting to Identify:
- **Required Technical Skills**: List all must-have technical competencies
- **Preferred Skills**: List nice-to-have skills and technologies
- **Experience Requirements**: Years, industry, domain specifics
- **Key Responsibilities**: Primary duties and scope expectations
- **Education/Certifications**: Required and preferred credentials
- **Soft Skills/Cultural Fit**: Communication, collaboration, values

#### Generate Custom Rubric Structure:
Create a job-specific rubric following this template:

```markdown
# Job-Specific Scoring Rubric: [Role Title] at [Company]
Generated: [Date]
Job Posting: {{ARG1}}

## Total Score: 100 Points

### 1. Core Technical Skills & Competencies (25 points)

#### Required Technical Skills (15 points) - Customized for [Role]
[Dynamically list the specific required skills from job posting]
- **[Skill 1]** (3 points): [Specific requirement from job]
- **[Skill 2]** (3 points): [Specific requirement from job]
- **[Skill 3]** (3 points): [Specific requirement from job]
- **[Skill 4]** (3 points): [Specific requirement from job]
- **[Skill 5]** (3 points): [Specific requirement from job]

#### Preferred Skills (10 points) - Nice-to-Have for [Role]
[Dynamically list preferred skills, each worth 2 points]
- **[Preferred 1]** (2 points): [From job posting]
- **[Preferred 2]** (2 points): [From job posting]
- **[Preferred 3]** (2 points): [From job posting]
- **[Preferred 4]** (2 points): [From job posting]
- **[Preferred 5]** (2 points): [From job posting]

### 2. Relevant Experience (25 points) - Tailored to [Company] Requirements

#### Years of Experience (10 points)
- Required: [X years from job posting]
- Industry-specific: [Industry from job posting]
- Domain expertise: [Specific domains mentioned]

#### Industry/Domain Experience (10 points)
- Primary industry: [From job posting]
- Related industries: [Adjacent acceptable industries]
- Technical domains: [Specific technical areas]

#### Role-Specific Experience (5 points)
- Previous roles: [Similar titles from job posting]
- Scope indicators: [Team size, budget, etc. from job posting]

### 3. Key Responsibilities (20 points) - Based on [Role] Requirements

#### Primary Duties (12 points)
[Extract 4 main responsibilities from job posting, 3 points each]
- **[Responsibility 1]** (3 points): [Specific duty]
- **[Responsibility 2]** (3 points): [Specific duty]
- **[Responsibility 3]** (3 points): [Specific duty]
- **[Responsibility 4]** (3 points): [Specific duty]

#### Scope & Complexity (8 points)
- Team/Project scale: [From job posting]
- Technical complexity: [Level indicated in posting]
- Business impact: [Strategic importance mentioned]

### 4. Achievements & Impact (15 points) - Aligned with [Company] Goals

#### Expected Outcomes (10 points)
[Based on job posting's success metrics]
- Performance indicators: [KPIs mentioned]
- Business objectives: [Goals from posting]

#### Innovation & Leadership (5 points)
- Leadership level: [From job posting]
- Innovation expectations: [Creative requirements]

### 5. Education & Certifications (10 points) - Per [Company] Requirements

#### Education (6 points)
- Required degree: [From job posting]
- Preferred field: [Specific major/field]
- Advanced degrees: [If mentioned as preferred]

#### Certifications (4 points)
- Required certifications: [List from job posting]
- Preferred certifications: [Nice-to-have certs]

### 6. Cultural Fit (5 points) - [Company] Values & Environment

#### Communication (3 points)
- Style requirements: [From job posting]
- Collaboration needs: [Team interaction level]

#### Values Alignment (2 points)
- Company values: [Extract from posting]
- Work environment: [Remote/hybrid/office preferences]

## Scoring Guidelines
[Keep standard scoring interpretation from framework]

## Role-Specific Considerations
[Add any unique aspects of this particular role]
```

#### Save Dynamic Rubric:
Save the generated rubric to: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`
Also save a reference copy with the assessment for audit trail

### 3. Acquire Domain Knowledge
Use web research to understand:
- Industry standards and role expectations for the specific position
- Required vs nice-to-have skills based on current market standards
- Typical responsibilities and seniority indicators in this domain
- Company context, culture, technology stack, and recent developments
- Current market conditions, salary ranges, and competitive landscape
- Industry-specific terminology and best practices

### 4. Perform 100-Point Assessment Using Dynamic Rubric

Apply the job-specific scoring rubric you created systematically:
1. **Core Technical Skills (25 pts)** - Map each required/preferred skill from the custom rubric to specific evidence in candidate's history
2. **Relevant Experience (25 pts)** - Evaluate against the specific years, industry, and domain requirements identified in the rubric
3. **Key Responsibilities (20 pts)** - Match candidate's experience to the 4 primary duties extracted from the job posting
4. **Achievements & Impact (15 pts)** - Verify metrics against the expected outcomes defined in the rubric
5. **Education & Certifications (10 pts)** - Check against the specific requirements listed in the custom rubric
6. **Cultural Fit (5 pts)** - Assess based on the company values and work environment extracted from the posting

### 5. Generate Comprehensive Report

Create assessment report using the dynamic rubric scores:
```markdown
# Candidate Assessment: [Role] at [Company]
**Assessment Date:** [Date]
**Rubric Used:** Rubric_[Company]_[Role]_[Date].md

## Executive Summary
[2-3 sentence overview of fit and recommendation based on job-specific criteria]

## Overall Score: [XX/100]
*Scored against job-specific rubric tailored for this position*

### Job-Specific Rubric Applied
[Brief note about the custom rubric created from the job posting]

### Detailed Scoring Breakdown

#### 1. Technical Skills & Competencies: [XX/25]
**Required Skills (15 points total) - As Defined in Job Posting:**
- [Actual Skill 1 from rubric]: [X/3] - [Evidence from candidate's history]
- [Actual Skill 2 from rubric]: [X/3] - [Evidence from candidate's history]
- [Actual Skill 3 from rubric]: [X/3] - [Evidence from candidate's history]
- [Actual Skill 4 from rubric]: [X/3] - [Evidence from candidate's history]
- [Actual Skill 5 from rubric]: [X/3] - [Evidence from candidate's history]

**Preferred Skills (10 points total) - Nice-to-Haves from Posting:**
- [Actual Preferred 1 from rubric]: [X/2] - [Evidence and reasoning]
- [Actual Preferred 2 from rubric]: [X/2] - [Evidence and reasoning]
- [Actual Preferred 3 from rubric]: [X/2] - [Evidence and reasoning]
- [Actual Preferred 4 from rubric]: [X/2] - [Evidence and reasoning]
- [Actual Preferred 5 from rubric]: [X/2] - [Evidence and reasoning]

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

### 6. Save Assessment and Rubric
1. **Save Dynamic Rubric**: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`
   - This becomes the official scoring criteria for this specific position
   - Include all extracted requirements from the job posting
   - Document the scoring breakdown with job-specific details

2. **Save Assessment Report**: `OutputResumes/Assessment_[Company]_[Role]_[Date].md`
   - Include reference to the specific rubric used
   - Document all scores with evidence mapping
   - Provide clear traceability between rubric criteria and candidate evaluation

## Important Notes
- **Dynamic Rubric Creation**: Each job posting generates its own tailored rubric
- **Consistency**: Use the framework from `jobscoringrubric.md` but customize content
- **Traceability**: Always reference which rubric was used for which assessment
- **Objectivity**: Be evidence-based in scoring using the job-specific criteria
- **Domain Research**: Use web research to understand industry-specific requirements
- **Documentation**: Save both rubric and assessment for complete audit trail
- **Confidence Levels**: Note where evidence is limited or assumptions are made