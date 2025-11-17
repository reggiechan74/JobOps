---
description: Compare assessments selected by the user in the @OutputResumes/ folder
argument-hint: <assessment-file-1> <assessment-file-2> [assessment-file-3] [assessment-file-4]
---

You are an expert HR analyst specializing in comparative candidate assessment analysis. Compare the selected job assessments to provide strategic hiring insights and recommendations.

## Your Task

Compare 2-4 assessment files from the OutputResumes folder to analyze candidate performance across different roles, identify patterns, and provide strategic hiring recommendations.

## Step-by-Step Process

### YAML front matter for comparison report
Prepend the comparison output with:

```yaml
---
assessments:
  - {{ARG1}}
  - {{ARG2}}
  # Append {{ARG3}} and {{ARG4}} if provided
generated_by: /comparejobs
generated_on: <ISO8601 timestamp>
output_type: assessment_comparison
status: draft
version: 1.0
---
```

Replace the placeholder comment with the additional assessment filenames you compare. Insert before the report heading and update on reruns.

### 1. Load Assessment Files
- Read the specified assessment files from `OutputResumes/{{ARG1}}`, `OutputResumes/{{ARG2}}`, and optionally `OutputResumes/{{ARG3}}` and `OutputResumes/{{ARG4}}`
- If file extensions are missing, assume `.md` extension
- Extract key information from each assessment:
  - Overall scores and ratings
  - Technical skills scores
  - Experience ratings
  - Achievement metrics
  - Cultural fit assessments
  - Hiring recommendations
  - Company and role details

### 2. Analyze Assessment Data
Parse each assessment to extract:
- **Candidate Information**: Name, role applied for, company
- **Scoring Breakdown**: Detailed scores by category (Technical Skills, Experience, Responsibilities, Achievements, Education, Cultural Fit)
- **Strengths & Gaps**: Key strengths and areas of concern
- **Hiring Recommendations**: Final recommendations and rationale
- **Assessment Date**: When evaluation was conducted

### 3. Perform Comparative Analysis
Conduct multi-dimensional comparison across:

#### Score Comparison
- Overall scores and rankings
- Category-wise performance patterns
- Consistency across different role requirements
- Score distribution and variance analysis

#### Role Suitability Analysis
- Best-fit roles based on skill alignment
- Transferable skills across positions
- Role-specific strengths and limitations
- Growth trajectory implications

#### Competitive Positioning
- Relative candidate strengths
- Market positioning against role requirements
- Differentiation factors
- Risk/reward profiles

#### Pattern Recognition
- Consistent performance areas
- Recurring gaps or concerns
- Cultural fit trends
- Experience leverage opportunities

### 4. Generate Strategic Insights
Analyze deeper implications:
- **Career Trajectory**: How roles complement candidate's progression
- **Skill Development**: Areas for professional growth
- **Market Positioning**: How candidate compares across different markets/roles
- **Strategic Recommendations**: Optimal role selection and negotiation insights

### 5. Create Comprehensive Comparison Report

```markdown
# Job Assessment Comparison Report
**Analysis Date:** [Date]
**Assessments Compared:** [Number] assessments across [Number] roles/companies

## Executive Summary
[2-3 sentence overview highlighting key findings and strategic recommendations]

## Assessment Overview

| Assessment | Role | Company | Overall Score | Recommendation | Date |
|------------|------|---------|---------------|-----------------|------|
| {{ARG1}} | [Role] | [Company] | [Score]/100 | [Recommendation] | [Date] |
| {{ARG2}} | [Role] | [Company] | [Score]/100 | [Recommendation] | [Date] |
| {{ARG3}} | [Role] | [Company] | [Score]/100 | [Recommendation] | [Date] |
| {{ARG4}} | [Role] | [Company] | [Score]/100 | [Recommendation] | [Date] |

## Detailed Score Comparison

### Overall Performance Rankings
1. **[Highest Scoring Role]**: [Score]/100 - [Company] - [Brief rationale]
2. **[Second Highest]**: [Score]/100 - [Company] - [Brief rationale]
3. **[Third]**: [Score]/100 - [Company] - [Brief rationale]
4. **[Lowest]**: [Score]/100 - [Company] - [Brief rationale]

### Category Performance Analysis

#### Technical Skills & Competencies
| Role/Company | Required Skills | Preferred Skills | Total | Analysis |
|--------------|-----------------|-------------------|-------|----------|
| [Role 1] | [X]/15 | [X]/10 | [X]/25 | [Strength/Gap summary] |
| [Role 2] | [X]/15 | [X]/10 | [X]/25 | [Strength/Gap summary] |
| [Role 3] | [X]/15 | [X]/10 | [X]/25 | [Strength/Gap summary] |
| [Role 4] | [X]/15 | [X]/10 | [X]/25 | [Strength/Gap summary] |

**Key Insights:**
- [Consistent technical strengths across roles]
- [Skills gaps that appear across multiple assessments]
- [Role-specific technical advantages]

#### Experience & Background
| Role/Company | Years | Industry | Role-Specific | Total | Analysis |
|--------------|-------|----------|---------------|-------|----------|
| [Role 1] | [X]/10 | [X]/10 | [X]/5 | [X]/25 | [Experience alignment] |
| [Role 2] | [X]/10 | [X]/10 | [X]/5 | [X]/25 | [Experience alignment] |
| [Role 3] | [X]/10 | [X]/10 | [X]/5 | [X]/25 | [Experience alignment] |
| [Role 4] | [X]/10 | [X]/10 | [X]/5 | [X]/25 | [Experience alignment] |

**Key Insights:**
- [Best experience matches and why]
- [Experience transfer opportunities]
- [Seniority level alignment across roles]

#### Key Responsibilities Alignment
| Role/Company | Primary Duties | Scope & Complexity | Total | Analysis |
|--------------|----------------|--------------------|-------|----------|
| [Role 1] | [X]/12 | [X]/8 | [X]/20 | [Responsibility match] |
| [Role 2] | [X]/12 | [X]/8 | [X]/20 | [Responsibility match] |
| [Role 3] | [X]/12 | [X]/8 | [X]/20 | [Responsibility match] |
| [Role 4] | [X]/12 | [X]/8 | [X]/20 | [Responsibility match] |

**Key Insights:**
- [Strongest capability alignments]
- [Scope considerations across roles]
- [Leadership/management fit analysis]

#### Achievements & Impact Performance
| Role/Company | Quantifiable Results | Innovation/Leadership | Total | Analysis |
|--------------|---------------------|----------------------|-------|----------|
| [Role 1] | [X]/10 | [X]/5 | [X]/15 | [Impact potential] |
| [Role 2] | [X]/10 | [X]/5 | [X]/15 | [Impact potential] |
| [Role 3] | [X]/10 | [X]/5 | [X]/15 | [Impact potential] |
| [Role 4] | [X]/10 | [X]/5 | [X]/15 | [Impact potential] |

**Key Insights:**
- [Consistent achievement patterns]
- [Role-specific impact opportunities]
- [Leadership potential across different contexts]

#### Cultural Fit Analysis
| Role/Company | Communication | Values Alignment | Total | Analysis |
|--------------|---------------|-------------------|-------|----------|
| [Role 1] | [X]/3 | [X]/2 | [X]/5 | [Cultural fit assessment] |
| [Role 2] | [X]/3 | [X]/2 | [X]/5 | [Cultural fit assessment] |
| [Role 3] | [X]/3 | [X]/2 | [X]/5 | [Cultural fit assessment] |
| [Role 4] | [X]/3 | [X]/2 | [X]/5 | [Cultural fit assessment] |

**Key Insights:**
- [Best cultural matches and reasons]
- [Cultural adaptation requirements]
- [Communication style fit analysis]

## Strategic Analysis

### Candidate Strengths Across Roles
[Bullet points of consistent strengths that translate across multiple roles]
- [Strength 1]: Evidence from [X] assessments showing [specific examples]
- [Strength 2]: Consistent high scoring in [category] across [all/most] roles
- [Strength 3]: [Achievement/skill] that provides competitive advantage

### Recurring Gaps & Development Areas
[Areas that need attention across multiple role applications]
- [Gap 1]: [Specific area] identified in [X] assessments, requiring [development approach]
- [Gap 2]: [Cultural/skill gap] that could limit success in [specific contexts]
- [Gap 3]: [Experience/knowledge gap] that impacts [specific role types]

### Role-Specific Advantages
#### Best Fit: [Role Name] at [Company]
- **Why it's the strongest match**: [Detailed reasoning]
- **Score advantage**: [X] points higher than average
- **Key differentiators**: [Specific factors that make this role ideal]
- **Success probability**: [High/Medium/Low] with rationale

#### Alternative Strong Fit: [Role Name] at [Company]
- **Positioning rationale**: [Why this is second choice]
- **Advantages**: [Specific benefits of this role]
- **Trade-offs**: [What candidate gives up vs. top choice]

#### Development Opportunity: [Role Name] at [Company]
- **Growth potential**: [How this role stretches candidate]
- **Learning curve**: [Skills/experience to develop]
- **Risk/reward**: [Assessment of challenges vs. opportunities]

### Market Positioning Insights
- **Salary negotiation position**: [Strong/Moderate/Weak] based on score differentials
- **Competitive advantage**: [Unique value propositions across roles]
- **Market demand alignment**: [How candidate meets current market needs]

## Interview Strategy Comparison

### Common Interview Focus Areas
[Areas that should be explored across all opportunities]
- [Focus Area 1]: [Why important across roles]
- [Focus Area 2]: [Validation needed across opportunities]
- [Focus Area 3]: [Potential concerns to address]

### Role-Specific Interview Priorities
#### [Role 1] - [Company 1]
- **Must validate**: [Specific areas from assessment]
- **Opportunity to highlight**: [Candidate's strongest areas for this role]
- **Risk mitigation**: [Address specific concerns]

#### [Role 2] - [Company 2]
- **Must validate**: [Specific areas from assessment]
- **Opportunity to highlight**: [Candidate's strongest areas for this role]
- **Risk mitigation**: [Address specific concerns]

[Continue for additional roles]

## Strategic Recommendations

### Optimal Strategy
1. **Primary Target**: [Role/Company] - [Reasoning]
2. **Secondary Option**: [Role/Company] - [Reasoning]
3. **Development Play**: [Role/Company] - [Reasoning]

### Negotiation Positioning
- **Strongest negotiating position**: [Role] due to [specific advantages]
- **Compensation expectations**: [Insights based on scores and market position]
- **Non-financial considerations**: [Benefits, growth, culture factors]

### Risk Management
- **Highest risk scenario**: [Role/situation] due to [specific gaps]
- **Mitigation strategies**: [Specific approaches to address concerns]
- **Backup planning**: [Alternative scenarios and preparation]

## Implementation Roadmap

### Immediate Actions (Next 2 weeks)
1. [Specific action based on analysis]
2. [Interview preparation focus]
3. [Research or networking activities]

### Medium-term Development (Next 3-6 months)
1. [Skill development priorities]
2. [Experience building activities]
3. [Network expansion strategies]

### Long-term Career Planning (6-18 months)
1. [Strategic role progression]
2. [Professional development investments]
3. [Market positioning improvements]

## Appendix

### Assessment Scoring Summary
[Detailed table showing all scores across all categories for reference]

### Key Evidence Mapping
[Cross-reference of major claims/achievements across assessments]

### Market Context Notes
[Relevant industry/market information that influenced the analysis]
```

### 6. Save Comparison Report
Save the comparative analysis to: `OutputResumes/Comparison_[Role1]_[Role2]_[etc]_[Date].md`

## Usage Guidelines

### When to Use This Command
- Candidate is applying to multiple roles simultaneously
- Need to prioritize multiple opportunities
- Strategic career planning and role selection
- Comparative analysis for negotiation positioning
- Understanding transferable skills across roles
- Identifying consistent patterns in assessments

### Best Practices
1. **Minimum 2 assessments**: Ensure meaningful comparison
2. **Recent assessments**: Use assessments from similar time periods
3. **Diverse roles**: Compare different role types/companies for insights
4. **Complete analysis**: Don't skip sections - full analysis provides best insights
5. **Evidence-based**: Reference specific scores and evidence from assessments
6. **Strategic focus**: Emphasize actionable insights and recommendations

### Limitations
- Comparisons only as good as underlying assessment quality
- Cultural fit may vary significantly by company even for similar roles
- Market conditions may change between assessment dates
- Candidate preferences and priorities not captured in assessments

## Example Usage
```bash
claude /comparejobs Assessment_JLL_VP_Office_Leasing_2025-09-25.md Assessment_Canerector_Vice_President_Real_Estate_2025-09-26.md
```

This will compare the JLL VP Office Leasing assessment with the Canerector VP Real Estate assessment, providing strategic insights for role selection and career planning.
