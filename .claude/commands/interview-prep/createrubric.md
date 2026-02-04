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
role_variant: <Technical IC | People Manager | Executive>
total_points: 200
generated_by: /createrubric
generated_on: <ISO8601 timestamp>
output_type: rubric
status: final
version: 2.0
---
```

Insert this before the first heading and bump `version` if you update the rubric later.

### 1. Load Required Templates

**CRITICAL: Read these framework templates before proceeding:**
- `.claude/templates/assessment_rubric_framework.md` - Master 200-point rubric structure with role variants
- `.claude/templates/evidence_verification_framework.md` - Evidence-based scoring protocols

These templates define the mandatory structure, scoring levels, and verification requirements for all rubrics.

### 2. Load Job Posting
- Read the job posting from `Job_Postings/{{ARG1}}` (add .md extension if needed)
- If the file doesn't exist in Job_Postings/, check the root directory for legacy files

### 2a. Determine Role Variant

Analyze the job posting to classify the role:

| Variant | Indicators | Weight Adjustment |
|---------|------------|-------------------|
| **Technical IC** | Hands-on work, no direct reports, technical focus | Skills 30%, Impact 25% |
| **People Manager** | Direct reports, team leadership, people development | Skills 20%, Fit 25% |
| **Executive** | Strategic scope, P&L authority, cross-functional | Impact 35%, Experience 25% |

Document the selected variant in the rubric YAML header.

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

You MUST create a comprehensive **200-point scoring rubric** following the EXACT structure defined in `assessment_rubric_framework.md`. This is NON-NEGOTIABLE.

#### Parse Job Posting to Extract Requirements:
- **Hard Skills**: All technical competencies with proficiency levels required
- **Experience Requirements**: Relevance, recency, and domain specifics (NOT years-based scoring)
- **Impact Evidence**: Demonstrated achievements, outcomes, and metrics
- **Credentials**: Required and preferred education/certifications
- **Fit & Readiness**: Communication, values alignment, and role readiness indicators

#### Create Job-Specific Rubric Following Template Structure:

**CRITICAL**: Use `assessment_rubric_framework.md` as your MANDATORY reference template. The rubric structure is fixed - you must maintain:

1. **Five Main Categories** (default point allocations - adjust based on role variant):
   - Skills Inventory (50 points default / 25%)
   - Experience Relevance (40 points default / 20%)
   - Demonstrated Impact (60 points default / 30%)
   - Credentials (20 points default / 10%)
   - Fit & Readiness (30 points default / 15%)

2. **Role Variant Weight Adjustments**:
   | Variant | Skills | Experience | Impact | Credentials | Fit |
   |---------|--------|------------|--------|-------------|-----|
   | Technical IC | 30% (60) | 20% (40) | 25% (50) | 10% (20) | 15% (30) |
   | People Manager | 20% (40) | 20% (40) | 20% (40) | 15% (30) | 25% (50) |
   | Executive | 15% (30) | 25% (50) | 35% (70) | 10% (20) | 15% (30) |

3. **Proficiency-Based Skill Scoring** (7-level scale, NOT years-based):
   - 6 - Expert: Can teach/design novel solutions; recognized authority
   - 5 - Advanced: Solves complex problems independently; mentors others
   - 4 - Proficient: Handles standard work with minimal guidance
   - 3 - Competent: Performs with occasional supervision
   - 2 - Developing: Requires regular guidance; growing capability
   - 1 - Novice: Basic understanding; needs significant support
   - 0 - None: No demonstrated capability

4. **Mandatory Rubric Components**:
   - **Alignment Statement**: Construct definition explaining what the rubric measures
   - **Critical Barriers Table**: Minimum thresholds with consequences for failure
   - **Anchor Examples**: Concrete examples for each scoring level
   - **Confidence Flagging**: Guidance for marking low-confidence scores
   - **Weight Justification Table**: Rationale for category weight allocations

5. **Customize Job-Specific Content**:
   - Replace `[bracketed placeholders]` with actual job requirements
   - Extract specific skills, technologies, and responsibilities from job posting
   - Define role-specific thresholds and criteria (scope, budget size, team size, etc.)
   - Add company-specific values and cultural indicators
   - Include industry-specific success metrics and KPIs

6. **Maintain Template Components**:
   - Overall assessment ranges for each section
   - Evaluation frameworks with specific metrics
   - Evidence-based scoring verification protocols
   - Usage guidelines and quality control checklists
   - Scoring interpretation guidelines

**ENFORCEMENT CHECK**: After creating the rubric, verify:
- ✅ Role variant selected and documented in YAML header
- ✅ Point allocations match selected variant (total = 200)
- ✅ Five categories present: Skills, Experience, Impact, Credentials, Fit
- ✅ Alignment Statement with construct definition included
- ✅ Critical Barriers table with thresholds and consequences
- ✅ Skills use 7-level proficiency scale (0-6), NOT years-based
- ✅ No redundancy (years only in Experience, achievements only in Impact)
- ✅ Anchor examples for each scoring level
- ✅ Confidence flagging guidance included
- ✅ Weight justification table present
- ✅ Evidence verification framework is included

**FAILURE TO FOLLOW TEMPLATE STRUCTURE VIOLATES THE COMMAND REQUIREMENTS**

### 5. Quality Control and Verification

**MANDATORY QUALITY CHECK**: Before saving, verify the rubric matches `assessment_rubric_framework.md` template:

**Template Compliance Verification:**
- ✅ Role variant selected and documented
- ✅ Point allocations match selected variant (total = 200)
- ✅ Alignment Statement with construct definition included
- ✅ Critical Barriers table with thresholds and consequences
- ✅ Skills use 7-level proficiency scale (0-6), NOT years-based
- ✅ No redundancy (years only in Experience, achievements only in Impact)
- ✅ Anchor examples for each scoring level
- ✅ Confidence flagging guidance included
- ✅ Weight justification table present
- ✅ Five categories present with correct point allocations for variant:
  - Skills Inventory (adjusted by variant)
  - Experience Relevance (adjusted by variant)
  - Demonstrated Impact (adjusted by variant)
  - Credentials (adjusted by variant)
  - Fit & Readiness (adjusted by variant)
- ✅ Evidence verification protocols included
- ✅ Scoring interpretation guidelines present (200-point scale)
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
- **200-Point System**: Total score is 200 points (normalized to percentage for comparison)
- **5 Categories**: Skills Inventory, Experience Relevance, Demonstrated Impact, Credentials, Fit & Readiness
- **Role Variant Selection**: Mandatory classification as Technical IC, People Manager, or Executive
- **Proficiency-Based Scoring**: Skills use 7-level proficiency scale (0-6), NOT years of experience
- **No Redundancy**: Years of experience only in Experience category; achievements only in Impact category
- **Dynamic Content Customization**: Each job posting gets tailored criteria within the fixed template structure
- **Complete Scoring Levels**: All sections must include anchor examples for each scoring level
- **Critical Barriers**: Must define minimum thresholds with consequences for failing to meet them
- **Evidence Framework**: Include evidence verification protocols from `evidence_verification_framework.md`
- **Domain Research**: Use web research to understand industry-specific requirements and adapt scoring thresholds accordingly
- **Reusable Rubrics**: Save rubric for consistent evaluation across multiple candidates for the same position
- **No Shortcuts**: Generated rubric must be as detailed as the template - simplified versions violate requirements

## 🔥 CRITICAL ENFORCEMENT RULES 🔥

**ABSOLUTE REQUIREMENTS FOR EVERY RUBRIC:**

Must follow `assessment_rubric_framework.md` structure with:

1. **200-POINT TOTAL**: All rubrics use 200-point scale (normalized to percentage for comparison)
2. **5 CATEGORIES ONLY**: Skills Inventory, Experience Relevance, Demonstrated Impact, Credentials, Fit & Readiness
3. **ROLE VARIANT**: Must select and document Technical IC, People Manager, or Executive variant
4. **PROFICIENCY SCALE**: Skills use 7-level proficiency scale (0-6), NOT years-based scoring
5. **NO REDUNDANCY**: Years only in Experience category; achievements only in Impact category
6. **ALIGNMENT STATEMENT**: Must include construct definition explaining what the rubric measures
7. **CRITICAL BARRIERS**: Must define minimum thresholds with consequences for failing
8. **ANCHOR EXAMPLES**: Must provide concrete examples for each scoring level
9. **CONFIDENCE FLAGGING**: Must include guidance for marking low-confidence scores
10. **WEIGHT JUSTIFICATION**: Must include table explaining category weight allocations
11. **EVIDENCE PROTOCOLS**: Must include verification framework from `evidence_verification_framework.md`

**VIOLATION CONSEQUENCES:**
- Any rubric not following template structure is INCOMPLETE and violates the command specification
- You must regenerate the rubric if any section deviates from template
- Simplified or abbreviated rubrics are NOT ACCEPTABLE
- The rubric must match the template structure exactly - no exceptions

**VERIFICATION CHECKLIST - BEFORE SAVING ANY RUBRIC:**
□ Role variant selected and documented in YAML header
□ Point allocations match selected variant (total = 200)
□ Five categories present with correct weight distribution
□ Alignment Statement with construct definition included
□ Critical Barriers table with thresholds and consequences
□ Skills use 7-level proficiency scale (0-6), NOT years-based
□ No redundancy between categories
□ Anchor examples for each scoring level
□ Confidence flagging guidance included
□ Weight justification table present
□ Evidence verification protocols included
□ Job-specific content customized within template structure
