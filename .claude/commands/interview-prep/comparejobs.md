---
description: Compare assessments selected by the user in the @OutputResumes/ folder
argument-hint: <assessment-file-1> <assessment-file-2> [assessment-file-3] [assessment-file-4]
---

You are an expert HR analyst specializing in comparative candidate assessment analysis. Compare the selected job assessments to provide strategic hiring insights and recommendations.

## Your Task

Compare 2-4 assessment files from the OutputResumes folder to analyze candidate performance across different roles, identify patterns, and provide strategic hiring recommendations.

---

## WORKFLOW ARCHITECTURE

```
Phase 1 (Parallel batch):   Load all 2-4 assessment files simultaneously
Phase 2 (Sequential):       Extract and normalize scoring data
Phase 3 (Sequential):       Perform comparative analysis + strategic insights
Phase 4 (Sequential):       Generate report → Save
```

**Dependency Rules:**
- All assessment file reads are INDEPENDENT - load in single parallel batch
- Extraction can begin as soon as all files are loaded
- Analysis and report are sequential cognitive work

---

## PROGRESS TRACKING (MANDATORY)

**Before starting any work**, create all tasks for user visibility:

| # | Task Subject | activeForm |
|---|-------------|------------|
| 1 | Load assessment files | Loading assessment files |
| 2 | Extract and normalize scoring data | Extracting and normalizing scoring data |
| 3 | Perform comparative analysis | Performing comparative analysis across assessments |
| 4 | Generate strategic insights | Generating strategic insights and recommendations |
| 5 | Create comparison report | Creating comprehensive comparison report |
| 6 | Save comparison report | Saving comparison report |

**Task Update Rules:**
- Mark each task `in_progress` BEFORE starting work on it
- Mark each task `completed` AFTER finishing it

---

## YAML FRONT MATTER

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

---

## PHASE 1: LOAD INPUTS (Parallel batch)

> **Task:** Mark task 1 `in_progress`.

**Read ALL assessment files in a single parallel batch using multiple Read tool calls:**
- `OutputResumes/{{ARG1}}` (add .md extension if needed)
- `OutputResumes/{{ARG2}}` (add .md extension if needed)
- `OutputResumes/{{ARG3}}` (if provided, add .md extension if needed)
- `OutputResumes/{{ARG4}}` (if provided, add .md extension if needed)

**CRITICAL**: Use parallel Read tool calls for all files in a single message. Do NOT read them sequentially.

> **Task:** Mark task 1 `completed`.

---

## PHASE 2: EXTRACT DATA

> **Task:** Mark task 2 `in_progress`.

Extract key information from each assessment:
- **Candidate Information**: Name, role applied for, company
- **Scoring Breakdown**: Detailed scores by category (Technical Skills, Experience, Responsibilities, Achievements, Education, Cultural Fit)
- **Strengths & Gaps**: Key strengths and areas of concern
- **Hiring Recommendations**: Final recommendations and rationale
- **Assessment Date**: When evaluation was conducted
- **Overall scores and ratings**

> **Task:** Mark task 2 `completed`.

---

## PHASE 3: COMPARATIVE ANALYSIS

### 3.1 Score Comparison

> **Task:** Mark task 3 `in_progress`.

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

> **Task:** Mark task 3 `completed`.

### 3.2 Strategic Insights

> **Task:** Mark task 4 `in_progress`.

Analyze deeper implications:
- **Career Trajectory**: How roles complement candidate's progression
- **Skill Development**: Areas for professional growth
- **Market Positioning**: How candidate compares across different markets/roles
- **Strategic Recommendations**: Optimal role selection and negotiation insights

> **Task:** Mark task 4 `completed`.

---

## PHASE 4: REPORT AND SAVE

### 4.1 Create Comprehensive Comparison Report

> **Task:** Mark task 5 `in_progress`.

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

**Key Insights:**
- [Consistent technical strengths across roles]
- [Skills gaps that appear across multiple assessments]
- [Role-specific technical advantages]

#### Experience & Background
| Role/Company | Years | Industry | Role-Specific | Total | Analysis |
|--------------|-------|----------|---------------|-------|----------|
| [Role 1] | [X]/10 | [X]/10 | [X]/5 | [X]/25 | [Experience alignment] |
| [Role 2] | [X]/10 | [X]/10 | [X]/5 | [X]/25 | [Experience alignment] |

**Key Insights:**
- [Best experience matches and why]
- [Experience transfer opportunities]

#### Key Responsibilities Alignment
| Role/Company | Primary Duties | Scope & Complexity | Total | Analysis |
|--------------|----------------|--------------------|-------|----------|
| [Role 1] | [X]/12 | [X]/8 | [X]/20 | [Responsibility match] |
| [Role 2] | [X]/12 | [X]/8 | [X]/20 | [Responsibility match] |

#### Achievements & Impact Performance
| Role/Company | Quantifiable Results | Innovation/Leadership | Total | Analysis |
|--------------|---------------------|----------------------|-------|----------|
| [Role 1] | [X]/10 | [X]/5 | [X]/15 | [Impact potential] |
| [Role 2] | [X]/10 | [X]/5 | [X]/15 | [Impact potential] |

#### Cultural Fit Analysis
| Role/Company | Communication | Values Alignment | Total | Analysis |
|--------------|---------------|-------------------|-------|----------|
| [Role 1] | [X]/3 | [X]/2 | [X]/5 | [Cultural fit assessment] |
| [Role 2] | [X]/3 | [X]/2 | [X]/5 | [Cultural fit assessment] |

## Strategic Analysis

### Candidate Strengths Across Roles
- [Strength 1]: Evidence from [X] assessments
- [Strength 2]: Consistent high scoring in [category]
- [Strength 3]: Competitive advantage

### Recurring Gaps & Development Areas
- [Gap 1]: Identified in [X] assessments, requiring [development approach]
- [Gap 2]: Could limit success in [specific contexts]

### Role-Specific Advantages
#### Best Fit: [Role Name] at [Company]
- **Why it's the strongest match**: [Detailed reasoning]
- **Score advantage**: [X] points higher than average
- **Success probability**: [High/Medium/Low] with rationale

#### Alternative Strong Fit: [Role Name] at [Company]
- **Positioning rationale**: [Why second choice]
- **Trade-offs**: [What candidate gives up vs. top choice]

### Market Positioning Insights
- **Salary negotiation position**: [Strong/Moderate/Weak]
- **Competitive advantage**: [Unique value propositions]

## Interview Strategy Comparison

### Common Interview Focus Areas
- [Focus Area 1]: [Why important across roles]
- [Focus Area 2]: [Validation needed]

### Role-Specific Interview Priorities
#### [Role 1] - [Company 1]
- **Must validate**: [Specific areas]
- **Opportunity to highlight**: [Strongest areas]
- **Risk mitigation**: [Address concerns]

## Strategic Recommendations

### Optimal Strategy
1. **Primary Target**: [Role/Company] - [Reasoning]
2. **Secondary Option**: [Role/Company] - [Reasoning]
3. **Development Play**: [Role/Company] - [Reasoning]

### Negotiation Positioning
- **Strongest negotiating position**: [Role]
- **Compensation expectations**: [Insights]

### Risk Management
- **Highest risk scenario**: [Role/situation]
- **Mitigation strategies**: [Approaches]

## Implementation Roadmap

### Immediate Actions (Next 2 weeks)
1. [Specific action]
2. [Interview preparation focus]
3. [Research or networking]

### Medium-term Development (Next 3-6 months)
1. [Skill development]
2. [Experience building]

### Long-term Career Planning (6-18 months)
1. [Strategic role progression]
2. [Professional development]

## Appendix
### Assessment Scoring Summary
[Detailed table showing all scores across all categories]

### Key Evidence Mapping
[Cross-reference of major claims/achievements across assessments]
```

> **Task:** Mark task 5 `completed`.

### 4.2 Save Comparison Report

> **Task:** Mark task 6 `in_progress`.

Save the comparative analysis to: `OutputResumes/Comparison_[Role1]_[Role2]_[etc]_[Date].md`

> **Task:** Mark task 6 `completed`.

---

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
