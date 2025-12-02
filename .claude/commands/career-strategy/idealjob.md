---
description: Generate a synthetic job description for a near-perfect fit role based on career history and preferences
argument-hint: [output-filename]
---

I'll create a hypothetical job description that represents a near-perfect fit based on your comprehensive career history, stated preferences, and patterns from high-scoring assessments.

**Output File**: `$1` (defaults to `Job_Postings/IdealJob_Synthetic_[Date].md` if not specified)

---

## Phase 1: Comprehensive Career Intelligence Gathering

### 1.1 Load Career Inventory

Read the complete career inventory from `/workspaces/resumeoptimizer/ResumeSourceFolder/`:

**Experience Files:**
- Read all files in `Experience/` directory to understand:
  - Complete work history and role progression
  - Responsibilities that energized vs. drained the candidate
  - Achievements and quantified impact
  - Skills demonstrated and developed
  - Industries and company types worked in
  - Team sizes managed and organizational contexts

**Career Highlights:**
- Read all files in `CareerHighlights/` directory for:
  - Core competencies and strengths
  - Education and designations
  - Professional certifications
  - Publications and thought leadership
  - Professional activities and industry involvement

**Technology Capabilities:**
- Read files in `Technology/` directory for:
  - Technical skills and proficiency levels
  - Tools and platforms mastered
  - GitHub repositories and projects
  - Innovation and technology adoption patterns

**Job Preferences (Critical):**
- Read `Preferences/Vision.md` for:
  - Vision: What the candidate aspires to
  - Anti-Vision: What to explicitly avoid
  - Target role types
  - Compensation requirements (salary, hourly, entrepreneur)
  - Work arrangement preferences
  - Deal-breakers and must-haves

### 1.2 Check for Optimized Candidate Profile

If available, read `/workspaces/resumeoptimizer/ResumeSourceFolder/.profile/candidate_profile.json` for structured summary with evidence references.

---

## Phase 2: High-Scoring Assessment Pattern Analysis

### 2.1 Identify High-Scoring Assessments

Search for assessment reports with scores of 90 or higher:

1. **Check Sample_Output directory:**
   - Read files matching `Sample_Output/*Great_Fit*.md` or `Sample_Output/*_Fit_Assessment*.md`
   - Extract patterns from assessments scoring 90+

2. **Scan OutputResumes directory:**
   - Search assessment files for `overall_score: 9` (90-99) in YAML front matter
   - Or scan for "Overall Score: 9" patterns in content
   - Read top 3-5 highest-scoring assessments

### 2.2 Extract Success Patterns

From high-scoring assessments, identify:

**Role Characteristics:**
- Job titles that scored highest
- Industries and company types
- Seniority levels and scope
- Key responsibilities that aligned well

**Technical Skill Alignment:**
- Required skills where candidate scored Expert (3/3)
- Preferred skills where candidate scored Strong (2/2)
- Technical competencies with direct evidence

**Experience Alignment:**
- Years of experience requirements that matched
- Industry/domain requirements that aligned
- Role-specific experience that transferred well

**Cultural Fit Factors:**
- Work environment characteristics that aligned
- Company values that resonated
- Team structure preferences

**Gap Patterns:**
- Common gaps even in high-scoring assessments
- Areas where points were lost
- Cultural or soft skill misalignments

---

## Phase 3: Market Research & Validation

### 3.1 Search Current Job Market

Use web search to understand current market conditions:

**Search queries to execute:**
1. "[Primary skill domain] [target role level] job postings 2025"
2. "[Target industry] [role type] salary Canada 2025"
3. "[Key technology/methodology] career opportunities trends 2025"
4. "Emerging roles [candidate's intersection of skills] 2025"

**Research objectives:**
- Validate that the ideal role exists in the market
- Identify companies hiring for similar roles
- Confirm compensation ranges align with preferences
- Discover emerging role titles that match skill profile

### 3.2 Job Board Reconnaissance

Research job boards for comparable positions:
- LinkedIn job trends in candidate's domain
- Indeed/Glassdoor for role requirements and compensation
- Industry-specific job boards relevant to candidate's expertise

**Extract from market research:**
- Common job titles for the skill profile
- Typical required vs preferred qualifications
- Standard compensation ranges
- Growth industries and emerging opportunities
- Geographic hotspots for the role type

---

## Phase 4: Synthetic Job Description Generation

### 4.1 Job Description Structure

Create a comprehensive, realistic job description following standard market formats:

```markdown
---
synthetic: true
generated_by: /idealjob
generated_on: [ISO8601 timestamp]
fit_score_target: 95+
based_on_assessments: [list of high-scoring assessment files analyzed]
market_validation: [summary of market research findings]
---

# [Ideal Job Title]

## Company Overview
[Create a realistic but fictional company profile that matches the candidate's preferred work environment, company size, culture, and industry. Include elements from the Vision preferences and avoid Anti-Vision characteristics.]

## About the Role
[Executive summary of the position emphasizing elements that align with candidate's strengths and preferences]

## Key Responsibilities
[4-6 primary responsibilities directly mapped to candidate's proven expertise and achievements]

## Required Qualifications
[List qualifications that match candidate's actual credentials - education, certifications, years of experience]

## Required Technical Skills
[Skills where candidate has documented Expert-level proficiency with evidence]

## Preferred Qualifications
[Nice-to-haves that candidate possesses or nearly possesses]

## What Success Looks Like
[Metrics and outcomes aligned with candidate's proven track record]

## Compensation & Benefits
[Based on candidate's stated requirements from Vision.md and market research]

## Work Environment
[Based on Vision preferences, avoiding Anti-Vision characteristics]

## Why This Role Exists
[Business context that would create demand for this specific skill combination]
```

### 4.2 Alignment Documentation

Include a section documenting the alignment rationale:

```markdown
## Alignment Analysis (Internal Reference)

### Strengths Leveraged
[List top 5-7 candidate strengths this role would utilize]

### Preferences Honored
[List Vision elements incorporated and Anti-Vision elements avoided]

### High-Scoring Assessment Patterns Applied
[Reference specific patterns from 90+ assessments that informed this description]

### Market Validation
[Evidence that similar roles exist and are being hired for]

### Predicted Fit Score
[Estimated score if candidate were assessed against this description]

### Search Strategy
[Suggested search queries to find real roles similar to this ideal]
```

---

## Phase 5: Reality Check & Actionable Output

### 5.1 Existence Validation

Research whether roles matching this description actually exist:
- Identify 2-3 real companies that might have similar roles
- Find actual job postings that share 70%+ of the characteristics
- Note any adjustments needed to make the role more realistic

### 5.2 Gap Analysis

Identify any gaps between the ideal role and candidate's current profile:
- Skills that could be strengthened to expand opportunities
- Certifications that would open more doors
- Experience gaps that limit certain aspects

### 5.3 Search Recommendations

Provide actionable job search guidance:
- Specific job titles to search for
- Companies likely to have similar roles
- Industries to target
- Keywords for job board searches
- Networking targets (roles/people who might know of such opportunities)

---

## Output Requirements

### Primary Output: Synthetic Job Description

Save to: `Job_Postings/IdealJob_Synthetic_[Date].md` (or specified filename)

The job description should:
- Be realistic enough to use as a search template
- Incorporate all Vision preferences
- Avoid all Anti-Vision characteristics
- Align with proven high-scoring assessment patterns
- Include market-validated compensation
- Specify a fictional but plausible company

### Secondary Output: Search Strategy Brief

Include at the end of the file:
- Real job titles to search
- Target companies to research
- Industry sectors to prioritize
- Networking recommendations
- Timeline considerations

---

## Validation Checklist

Before finalizing, verify:

- [ ] All Vision preferences incorporated
- [ ] All Anti-Vision elements avoided
- [ ] Compensation matches stated requirements
- [ ] Technical skills match documented expertise (Expert level)
- [ ] Experience requirements achievable with current background
- [ ] Role scope matches preference for challenge/autonomy
- [ ] Company culture description aligns with preferences
- [ ] Market research validates role viability
- [ ] Real comparable roles identified
- [ ] Actionable search strategy provided

---

## Example Usage

```bash
# Generate ideal job with default filename
/idealjob

# Generate with custom output filename
/idealjob Job_Postings/MyIdealRole_AI_RealEstate_2025.md
```

---

Now executing comprehensive career analysis to generate your ideal job description...
