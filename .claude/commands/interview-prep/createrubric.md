---
description: Create a detailed scoring rubric from job posting without performing assessment
argument-hint: <job-posting-file>
---

You are acting as an expert HR recruiter and domain-knowledgeable hiring manager. Create a comprehensive, job-specific scoring rubric from the specified job posting that can be used for consistent candidate evaluation.

## Your Task

Analyze the {{ARG1}} job posting and generate a detailed, reusable scoring rubric that extracts all requirements and creates standardized evaluation criteria.

## Step-by-Step Process

### YAML front matter for rubric output
Write the rubric to `Scoring_Rubrics/Rubric_*` and begin the file with:

```yaml
---
job_file: Job_Postings/{{ARG1}}
role: <role title>
company: <company name>
generated_by: /createrubric
generated_on: <ISO8601 timestamp>
output_type: rubric
status: final
version: 1.0
---
```

Insert this before the first heading and bump `version` if you update the rubric later.

### 1. Load Required Templates

**CRITICAL: Read these framework templates before proceeding:**
- `.claude/templates/assessment_rubric_framework.md` - Master 100-point rubric structure
- `.claude/templates/evidence_verification_framework.md` - Evidence-based scoring protocols

These templates define the mandatory structure, scoring levels, and verification requirements for all rubrics.

### 2. Load Job Posting
- Read the job posting from `Job_Postings/{{ARG1}}` (add .md extension if needed)
- If the file doesn't exist in Job_Postings/, check the root directory for legacy files

### 3. Acquire Domain Knowledge
Use web research to understand:
- Industry standards and role expectations for the specific position
- Required vs nice-to-have skills based on current market standards
- Typical responsibilities and seniority indicators in this domain
- Company context, culture, technology stack, and recent developments
- Current market conditions, salary ranges, and competitive landscape
- Industry-specific terminology and best practices

### 4. Create Job-Specific Rubric

**🚨 MANDATORY DETAILED SCORING REQUIREMENT 🚨**

You MUST create a comprehensive 100-point scoring rubric following the EXACT structure defined in `assessment_rubric_framework.md`. This is NON-NEGOTIABLE.

#### Parse Job Posting to Extract Requirements:
- **Required Technical Skills**: All must-have technical competencies
- **Preferred Skills**: Nice-to-have skills and technologies
- **Experience Requirements**: Years, industry, domain specifics
- **Key Responsibilities**: Primary duties and scope expectations
- **Education/Certifications**: Required and preferred credentials
- **Soft Skills/Cultural Fit**: Communication, collaboration, values

#### Create Job-Specific Rubric Following Template Structure:

**CRITICAL**: Use `assessment_rubric_framework.md` as your MANDATORY reference template. The rubric structure is fixed - you must maintain:

1. **Six Main Sections** (point allocations cannot change):
   - Core Technical Skills & Competencies (25 points)
   - Relevant Experience (25 points)
   - Key Responsibilities (20 points)
   - Achievements & Impact (15 points)
   - Education & Certifications (10 points)
   - Cultural Fit (5 points)

2. **Detailed Scoring Levels** (MUST include for every criterion):
   - Required skills: Expert/Proficient/Basic/None (4-level)
   - Preferred skills: Strong/Basic/None (3-level)
   - Experience sections: 5-level detailed breakdowns
   - All other sections: Full scoring frameworks as shown in template

3. **Customize Job-Specific Content**:
   - Replace `[bracketed placeholders]` with actual job requirements
   - Extract specific skills, technologies, and responsibilities from job posting
   - Define role-specific thresholds and criteria (years, budget size, team size, etc.)
   - Add company-specific values and cultural indicators
   - Include industry-specific success metrics and KPIs

4. **Maintain Template Components**:
   - Overall assessment ranges for each section
   - Evaluation frameworks with specific metrics
   - Evidence-based scoring verification protocols
   - Usage guidelines and quality control checklists
   - Scoring interpretation guidelines

**ENFORCEMENT CHECK**: After creating the rubric, verify:
- ✅ Every required skill has Expert/Proficient/Basic/None scoring with specific criteria
- ✅ Every preferred skill has Strong/Basic/None scoring with role-specific thresholds
- ✅ Experience sections have 5-level detailed scoring breakdowns
- ✅ Primary duties have complete evaluation frameworks
- ✅ Scope & complexity has detailed metrics (budget, team, geography, stakeholders, duration)
- ✅ Achievements have 5-level scoring with quantitative thresholds
- ✅ Education & certifications have multi-level detailed scoring
- ✅ Cultural fit has comprehensive evaluation criteria
- ✅ Evidence verification framework is included
- ✅ All sections match template structure exactly

**FAILURE TO FOLLOW TEMPLATE STRUCTURE VIOLATES THE COMMAND REQUIREMENTS**

### 5. Quality Control and Verification

**MANDATORY QUALITY CHECK**: Before saving, verify the rubric matches `assessment_rubric_framework.md` template:

**Template Compliance Verification:**
- ✅ All six main sections present with correct point allocations (25/25/20/15/10/5)
- ✅ Every required skill has Expert/Proficient/Basic/None scoring with specific criteria
- ✅ Every preferred skill has Strong/Basic/None scoring with role-specific thresholds
- ✅ Experience sections have 5-level detailed scoring breakdowns (Exceeds/Meets Plus/Meets/Near/Below)
- ✅ Primary duties have complete evaluation frameworks
- ✅ Scope & complexity has detailed evaluation framework with 5 specific metrics
- ✅ Achievements & impact has 5-level scoring with quantitative thresholds
- ✅ Education & certifications have detailed multi-level scoring
- ✅ Cultural fit has comprehensive evaluation criteria and scoring levels
- ✅ Evidence verification protocols included
- ✅ Scoring interpretation guidelines present
- ✅ Usage guidelines and quality control checklists included

**IF ANY SECTION LACKS REQUIRED COMPONENTS, THE RUBRIC IS INCOMPLETE AND MUST BE REGENERATED**

### 6. Save the Rubric

Save the generated rubric to: `Scoring_Rubrics/Rubric_[Company]_[Role]_[Date].md`

Provide a summary of:
- Template compliance confirmation
- Total skills/requirements identified
- Key differentiators for this role
- Notable industry-specific requirements
- Job-specific customizations made to template
- Recommended usage guidelines

## Important Notes

- **Template-Based Approach**: Every rubric MUST follow `assessment_rubric_framework.md` structure exactly
- **MANDATORY DETAILED SCORING**: All granular scoring breakdowns from the template are required - this is not optional
- **Dynamic Content Customization**: Each job posting gets tailored criteria within the fixed template structure
- **Complete Scoring Levels**: All sections must include detailed point breakdowns as shown in template
- **Role-Specific Adaptation**: Customize content and thresholds for the specific role while maintaining template structure
- **Evidence Framework**: Include evidence verification protocols from `evidence_verification_framework.md`
- **Domain Research**: Use web research to understand industry-specific requirements and adapt scoring thresholds accordingly
- **Reusable Rubrics**: Save rubric for consistent evaluation across multiple candidates for the same position
- **No Shortcuts**: Generated rubric must be as detailed as the template - simplified versions violate requirements

## 🔥 CRITICAL ENFORCEMENT RULES 🔥

**ABSOLUTE REQUIREMENTS FOR EVERY RUBRIC:**

Must follow `assessment_rubric_framework.md` structure with:

1. **EVERY REQUIRED SKILL** must have Expert (3) / Proficient (2) / Basic (1) / None (0) with specific criteria
2. **EVERY PREFERRED SKILL** must have Strong (2) / Basic (1) / None (0) with role-specific thresholds
3. **EXPERIENCE SECTIONS** must have 5-level detailed breakdowns (Exceeds/Meets Plus/Meets/Near/Below)
4. **PRIMARY DUTIES** must have Expert/Proficient/Basic/None with quantitative thresholds
5. **SCOPE & COMPLEXITY** must include detailed evaluation framework with 5 specific metrics
6. **ACHIEVEMENTS** must have 5-level scoring with quantitative thresholds
7. **EDUCATION & CERTIFICATIONS** must have multi-level detailed scoring
8. **CULTURAL FIT** must have comprehensive evaluation criteria
9. **EVIDENCE PROTOCOLS** must include verification framework from `evidence_verification_framework.md`

**VIOLATION CONSEQUENCES:**
- Any rubric not following template structure is INCOMPLETE and violates the command specification
- You must regenerate the rubric if any section deviates from template
- Simplified or abbreviated rubrics are NOT ACCEPTABLE
- The rubric must match the template structure exactly - no exceptions

**VERIFICATION CHECKLIST - BEFORE SAVING ANY RUBRIC:**
□ Template structure followed exactly (all sections, point allocations, scoring levels)
□ Every skill has multi-level scoring with specific, measurable criteria
□ Every section includes detailed point allocation breakdowns as shown in template
□ All evaluation frameworks include quantitative thresholds
□ Overall assessment ranges provided for each section
□ Evidence verification protocols included
□ Job-specific content customized within template structure
