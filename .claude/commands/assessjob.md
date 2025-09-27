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
- Read the job posting from `Job_Postings/{{ARG1}}` (add .md extension if needed)
- Load ALL candidate work history files from `ResumeSourceFolder/` directory:
  - Use glob pattern to find all markdown files in ResumeSourceFolder/
  - Read each file to build complete candidate profile
  - Include all CV, resume, technology capability, and work history documents

### 2. Create Dynamic Job-Specific Rubric
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
Create a job-specific rubric following this template with MANDATORY detailed scoring breakdowns.

**CRITICAL REQUIREMENT**: You MUST include the complete detailed scoring breakdown (Expert/Proficient/Basic/None) for EVERY SINGLE skill and criteria. Do NOT create simplified versions. Do NOT skip the granular point allocation. The rubric MUST be as detailed as the template below - no shortcuts allowed.

**ENFORCEMENT CHECK**: After creating the rubric, verify that EVERY section includes:
- Individual skill breakdowns with 4-level scoring (Expert/Proficient/Basic/None)
- Specific measurable criteria for each level
- Overall assessment ranges
- Complete evaluation frameworks for scope & complexity
- Detailed scoring breakdowns for achievements & impact

**FAILURE TO INCLUDE DETAILED BREAKDOWNS VIOLATES THE COMMAND REQUIREMENTS**

```markdown
# Job-Specific Scoring Rubric: [Role Title] at [Company]
Generated: [Date]
Job Posting: {{ARG1}}

## Total Score: 100 Points

### 1. Core Technical Skills & Competencies (25 points)

#### Required Technical Skills (15 points) - Customized for [Role]
**MANDATORY: Include detailed scoring breakdown for each skill (3 points each)**

- **[Skill 1]** (3 points): [Specific requirement from job]
  - Expert (3 points): [Role-specific expert criteria - 5+ years, mentoring capability, etc.]
  - Proficient (2 points): [Role-specific proficient criteria - 2-4 years, independent execution, etc.]
  - Basic (1 point): [Role-specific basic criteria - <2 years or supported execution]
  - None (0 points): No demonstrated experience

- **[Skill 2]** (3 points): [Specific requirement from job]
  - Expert (3 points): [Role-specific expert criteria]
  - Proficient (2 points): [Role-specific proficient criteria]
  - Basic (1 point): [Role-specific basic criteria]
  - None (0 points): No demonstrated experience

- **[Skill 3]** (3 points): [Specific requirement from job]
  - Expert (3 points): [Role-specific expert criteria]
  - Proficient (2 points): [Role-specific proficient criteria]
  - Basic (1 point): [Role-specific basic criteria]
  - None (0 points): No demonstrated experience

- **[Skill 4]** (3 points): [Specific requirement from job]
  - Expert (3 points): [Role-specific expert criteria]
  - Proficient (2 points): [Role-specific proficient criteria]
  - Basic (1 point): [Role-specific basic criteria]
  - None (0 points): No demonstrated experience

- **[Skill 5]** (3 points): [Specific requirement from job]
  - Expert (3 points): [Role-specific expert criteria]
  - Proficient (2 points): [Role-specific proficient criteria]
  - Basic (1 point): [Role-specific basic criteria]
  - None (0 points): No demonstrated experience

**Overall Technical Skills Assessment:**
- **Expert Match (13-15 points)**: Demonstrates all required skills with deep expertise (3+ years recent experience in each)
- **Strong Match (10-12 points)**: Possesses 80%+ of required skills with solid experience (2+ years each)
- **Good Match (7-9 points)**: Has 60-80% of required skills with adequate experience (1+ years each)
- **Partial Match (4-6 points)**: Shows 40-60% of required skills or transferable experience
- **Minimal Match (0-3 points)**: Less than 40% skill alignment with limited depth

#### Preferred Skills (10 points) - Nice-to-Have for [Role]
**MANDATORY: Include detailed scoring breakdown for each preferred skill (2 points each)**

- **[Preferred 1]** (2 points): [From job posting]
  - Strong (2 points): [Role-specific strong criteria - extensive experience, competitive advantage]
  - Basic (1 point): [Role-specific basic criteria - some experience or exposure]
  - None (0 points): No demonstrated experience

- **[Preferred 2]** (2 points): [From job posting]
  - Strong (2 points): [Role-specific strong criteria]
  - Basic (1 point): [Role-specific basic criteria]
  - None (0 points): No demonstrated experience

- **[Preferred 3]** (2 points): [From job posting]
  - Strong (2 points): [Role-specific strong criteria]
  - Basic (1 point): [Role-specific basic criteria]
  - None (0 points): No demonstrated experience

- **[Preferred 4]** (2 points): [From job posting]
  - Strong (2 points): [Role-specific strong criteria]
  - Basic (1 point): [Role-specific basic criteria]
  - None (0 points): No demonstrated experience

- **[Preferred 5]** (2 points): [From job posting]
  - Strong (2 points): [Role-specific strong criteria]
  - Basic (1 point): [Role-specific basic criteria]
  - None (0 points): No demonstrated experience

**Overall Preferred Skills Assessment:**
- **Exceptional (9-10 points)**: Exceeds all preferred qualifications with depth (4-5 skills at strong level)
- **Strong (7-8 points)**: Meets 75%+ of preferred qualifications (3-4 skills demonstrated)
- **Good (5-6 points)**: Meets 50-75% of preferred qualifications (2-3 skills shown)
- **Basic (3-4 points)**: Meets 25-50% of preferred qualifications (1-2 skills evident)
- **Limited (0-2 points)**: Meets less than 25% of preferred qualifications (minimal evidence)

### 2. Relevant Experience (25 points) - Tailored to [Company] Requirements

#### Years of Experience (10 points)
**MANDATORY: Specify exact requirements and detailed scoring criteria**
- Required: [X years from job posting]
- Industry-specific: [Industry from job posting]
- Domain expertise: [Specific domains mentioned]

**Scoring Breakdown:**
- **Exceeds (9-10 points)**: Exceeds required experience by 3+ years
- **Meets Plus (7-8 points)**: Exceeds required experience by 1-2 years
- **Meets (5-6 points)**: Meets exact experience requirements
- **Near Match (3-4 points)**: Within 1-2 years of requirement
- **Below (0-2 points)**: More than 2 years below requirement

#### Industry/Domain Experience (10 points)
**MANDATORY: Define specific industry requirements and scoring levels**
- Primary industry: [From job posting]
- Related industries: [Adjacent acceptable industries]
- Technical domains: [Specific technical areas]

**Scoring Breakdown:**
- **Direct Match (9-10 points)**: Same industry with deep domain expertise
- **Adjacent Industry (7-8 points)**: Related industry with transferable knowledge
- **Partial Overlap (5-6 points)**: Some industry exposure or similar domains
- **Transferable (3-4 points)**: Different industry but applicable experience
- **No Alignment (0-2 points)**: No relevant industry experience

#### Role-Specific Experience (5 points)
**MANDATORY: Define comparable roles and scope requirements**
- Previous roles: [Similar titles from job posting]
- Scope indicators: [Team size, budget, etc. from job posting]

**Scoring Breakdown:**
- **Exact Match (5 points)**: Performed identical role previously
- **Very Similar (4 points)**: 80%+ overlap with target role
- **Similar (3 points)**: 60-80% overlap with responsibilities
- **Related (2 points)**: 40-60% overlap
- **Different (0-1 points)**: Less than 40% overlap

### 3. Key Responsibilities (20 points) - Based on [Role] Requirements

#### Primary Duties (12 points)
**MANDATORY: Extract 4 main responsibilities with detailed scoring criteria (3 points each)**

- **[Responsibility 1]** (3 points): [Specific duty from job posting]
- **[Responsibility 2]** (3 points): [Specific duty from job posting]
- **[Responsibility 3]** (3 points): [Specific duty from job posting]
- **[Responsibility 4]** (3 points): [Specific duty from job posting]

**Overall Primary Responsibilities Assessment:**
- **Complete Alignment (11-12 points)**: Demonstrated success in all primary responsibilities
- **Strong Alignment (8-10 points)**: Proven experience in 75%+ of primary duties
- **Good Alignment (5-7 points)**: Experience in 50-75% of primary duties
- **Partial Alignment (3-4 points)**: Experience in 25-50% of primary duties
- **Limited Alignment (0-2 points)**: Less than 25% responsibility match

#### Scope & Complexity (8 points)
**MANDATORY: Define specific scope requirements and detailed evaluation framework**
- Team/Project scale: [From job posting - specific numbers, scope indicators]
- Technical complexity: [Level indicated in posting]
- Business impact: [Strategic importance mentioned]

**Evaluation Framework:** Budget size, team size, geographic scope, stakeholder complexity, project duration

**Scoring Breakdown:**
- **Greater Scope (7-8 points)**: Managed larger scope/more complex projects
  - Budget: >2x target role budget requirements
  - Team: Led teams 50% larger than role requirement
  - Geography: Multi-regional/international vs. local requirement
  - Stakeholders: Complex matrix organizations, external partners
  - Duration: Multi-year programs vs. annual cycles

- **Equal Scope (5-6 points)**: Similar scope and complexity
  - Budget: Within 25% of target role budget size
  - Team: Comparable team sizes (±25%)
  - Geography: Same market coverage scope
  - Stakeholders: Similar complexity levels
  - Duration: Comparable project/program lengths

- **Slightly Less (3-4 points)**: Somewhat smaller scope but relevant
  - Budget: 25-50% below target role requirements
  - Team: Smaller teams but relevant leadership experience
  - Geography: Regional vs. national, or local vs. regional
  - Stakeholders: Simpler stakeholder landscape
  - Duration: Shorter cycles but relevant experience

- **Limited Scope (1-2 points)**: Significantly smaller scope
  - Budget: >50% below requirements
  - Team: Individual contributor or very small teams
  - Geography: Local only vs. broader requirement
  - Stakeholders: Internal only or limited external
  - Duration: Short-term projects only

- **No Comparison (0 points)**: No comparable scope experience
  - Cannot demonstrate relevant scope indicators
  - No leadership or project management evidence

### 4. Achievements & Impact (15 points) - Aligned with [Company] Goals

#### Expected Outcomes (10 points)
**MANDATORY: Define specific success metrics and detailed evaluation criteria**
[Based on job posting's success metrics]
- Performance indicators: [KPIs mentioned]
- Business objectives: [Goals from posting]

**Evaluation Criteria:** Revenue impact, cost savings, efficiency gains, scale metrics, timeline performance

**Scoring Breakdown:**
- **Exceptional Impact (9-10 points)**: Multiple high-impact achievements exceeding role level
  - Revenue/savings: >$10M impact or >50% improvement
  - Scale: Large team/project leadership (100+ people or $100M+ budgets)
  - Efficiency: >30% process improvements with quantified results
  - Timeline: Consistently delivered ahead of schedule

- **Strong Impact (7-8 points)**: Clear measurable achievements at appropriate level
  - Revenue/savings: $1M-$10M impact or 20-50% improvement
  - Scale: Mid-level leadership (25-100 people or $10M-$100M budgets)
  - Efficiency: 15-30% improvements documented
  - Timeline: Met aggressive deadlines consistently

- **Good Impact (5-6 points)**: Some quantified achievements relevant to role
  - Revenue/savings: $100K-$1M impact or 10-20% improvement
  - Scale: Team lead role (5-25 people or $1M-$10M budgets)
  - Efficiency: 5-15% improvements shown
  - Timeline: Generally met project timelines

- **Basic Impact (3-4 points)**: Few measurable outcomes documented
  - Revenue/savings: <$100K impact or <10% improvement
  - Scale: Individual contributor with some leadership
  - Efficiency: Minor improvements mentioned
  - Timeline: Mixed performance record

- **No Metrics (0-2 points)**: No quantifiable achievements presented
  - No financial impact demonstrated
  - No scale indicators provided
  - No efficiency metrics available
  - Cannot demonstrate timeline performance

#### Innovation & Leadership (5 points)
**MANDATORY: Define specific leadership expectations and innovation requirements**
- Leadership level: [From job posting]
- Innovation expectations: [Creative requirements]

**Scoring Breakdown:**
- **Transformational (5 points)**: Led major initiatives or innovations
- **Significant (4 points)**: Drove important improvements or changes
- **Contributor (3 points)**: Active participant in improvements
- **Supporter (1-2 points)**: Supported others' initiatives
- **None Shown (0 points)**: No innovation or leadership evidence

### 5. Education & Certifications (10 points) - Per [Company] Requirements

#### Education (6 points)
**MANDATORY: Define specific educational requirements and scoring levels**
- Required degree: [From job posting]
- Preferred field: [Specific major/field]
- Advanced degrees: [If mentioned as preferred]

**Scoring Breakdown:**
- **Exceeds (6 points)**: Advanced degree beyond requirement
- **Meets Plus (5 points)**: Required degree plus relevant additional education
- **Meets (3-4 points)**: Meets exact educational requirements
- **Equivalent (2 points)**: Equivalent experience or alternative credential
- **Below (0-1 points)**: Does not meet educational requirements

#### Certifications (4 points)
**MANDATORY: List specific certifications and define scoring criteria**
- Required certifications: [List from job posting]
- Preferred certifications: [Nice-to-have certs]

**Scoring Breakdown:**
- **Multiple Relevant (4 points)**: All required certifications plus 2+ additional industry-relevant designations
- **Complete (3 points)**: All required certifications current and in good standing, or 1 premium certification exceeding requirements
- **Partial (2 points)**: 50-75% of required certifications, or working toward premium designation with related certifications
- **In Progress (1 point)**: 25-50% of required certifications, or enrolled in required certification program with completion timeline
- **None (0 points)**: No relevant certifications and no active pursuit of required credentials

### 6. Cultural Fit (5 points) - [Company] Values & Environment

#### Communication (3 points)
**MANDATORY: Define specific communication requirements and evaluation criteria**
- Style requirements: [From job posting]
- Collaboration needs: [Team interaction level]

**Evaluation Criteria:** Written communication, presentation skills, cross-functional collaboration, conflict resolution, stakeholder management

**Scoring Breakdown:**
- **Exceptional (3 points)**: Clear evidence of superior communication/collaboration
  - Advanced presentation/public speaking experience
  - Cross-functional leadership across departments/divisions
  - Documented conflict resolution and negotiation success
  - Stakeholder management at C-level or external clients
  - Written communication (reports, proposals) at executive level

- **Strong (2 points)**: Good examples of team collaboration
  - Regular presentation duties to senior management
  - Cross-functional project leadership
  - Some conflict resolution experience
  - Internal stakeholder management
  - Professional writing capabilities demonstrated

- **Basic (1 point)**: Some evidence of teamwork
  - Team meeting facilitation
  - Intra-departmental collaboration
  - Basic conflict management
  - Peer-level stakeholder interaction
  - Standard business communication skills

- **Limited (0 points)**: No clear communication/collaboration evidence
  - Individual contributor only
  - No presentation or facilitation experience
  - No stakeholder management shown
  - Limited communication examples

#### Values Alignment (2 points)
**MANDATORY: Extract specific company values and define alignment criteria**
- Company values: [Extract from posting]
- Work environment: [Remote/hybrid/office preferences]

**Evaluation Criteria:** Company values demonstration, work environment preferences, adaptability, growth mindset

**Scoring Breakdown:**
- **Strong Match (2 points)**: Clear alignment with company values/culture
  - Concrete examples demonstrating core company values
  - Work style preferences align with company environment
  - Evidence of adaptability and change management
  - Growth mindset and continuous learning demonstrated
  - Cultural add potential (brings complementary strengths)

- **Good Match (1 point)**: Some cultural alignment indicators
  - General value alignment shown
  - Acceptable work style fit
  - Some adaptability evidence
  - Basic learning orientation
  - No obvious cultural conflicts

- **Unknown (0 points)**: No evidence of cultural fit
  - Cannot assess value alignment
  - Work style preferences unclear
  - No adaptability indicators
  - No learning/growth evidence shown

## Scoring Guidelines
**MANDATORY: Include complete scoring interpretation from framework rubric**

### Overall Score Ranges:
- **90-100**: Exceptional candidate - exceeds requirements significantly
- **80-89**: Excellent candidate - strong match, minor gaps if any
- **70-79**: Good candidate - solid match, some development areas
- **60-69**: Potential candidate - meets core requirements, notable gaps
- **50-59**: Borderline candidate - significant gaps, consider if high potential
- **Below 50**: Not recommended - major gaps in critical areas

## Usage Guidelines

### When Scoring:
1. Review job description thoroughly before scoring
2. Map each requirement to candidate's experience
3. Look for evidence, not assumptions
4. Consider recency of experience (weight recent more heavily)
5. Account for career progression and growth trajectory

### Critical Factors:
- Any "must-have" requirement scored below 3 may disqualify
- Consider minimum acceptable scores for each category
- Weight sections differently based on role priorities

### Documentation:
- Note specific examples supporting each score
- Document concerns or areas needing clarification
- Identify questions for interview follow-up
- Record any special circumstances or context

## Role-Specific Considerations
**MANDATORY: Add specific considerations based on job posting analysis**
[Add any unique aspects of this particular role, market conditions, competitive landscape, etc.]
```

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

[ ] **90-100**: Exceptional candidate - exceeds requirements significantly
[ ] **80-89**: Excellent candidate - strong match, minor gaps if any
[ ] **70-79**: Good candidate - solid match, some development areas
[ ] **60-69**: Potential candidate - meets core requirements, notable gaps
[ ] **50-59**: Borderline candidate - significant gaps, consider if high potential
[ ] **Below 50**: Not recommended - major gaps in critical areas

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
- **MANDATORY DETAILED SCORING**: Every rubric MUST include the granular scoring breakdowns from the framework rubric. This is not optional - it is required.
- **Dynamic Rubric Creation**: Each job posting generates its own tailored rubric with job-specific criteria but framework-level detail
- **Complete Scoring Levels**: All sections must include the detailed point breakdowns (Expert/Proficient/Basic, Exceptional/Strong/Good, etc.)
- **Role-Specific Adaptation**: Customize the content and thresholds for the specific role while maintaining scoring structure depth
- **Consistency**: Use the embedded framework as the foundation but customize content to job requirements
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