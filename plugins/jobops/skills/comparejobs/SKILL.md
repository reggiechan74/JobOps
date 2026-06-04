---
name: comparejobs
description: Compare assessment files across application folders under applications_root
disable-model-invocation: true
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if this skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if this skill has jurisdiction-sensitive logic (crisis/legal skills accept `--jurisdiction=<ISO-3166-2>` to override).

## Your Task

Compare 2-4 assessment files produced by `/assessjob` (one per application folder under `{config.directories.applications_root}`) — or ideal-role archetype files produced by `/idealjob` — to analyze candidate performance across different roles, identify patterns, and provide strategic hiring recommendations.

Each `{{ARGn}}` is either:

- an **application slug** (`{Company}_{Role}_{Date}`) — the skill loads `{applications_root}/{{ARGn}}/assessment/assessment.md`, or
- a **direct file path** — any argument containing `/` or ending in `.md` is read as-is (e.g., `{career_analysis}/idealjob_{YYYYMMDD}_anchor.md`). `/idealjob` archetype files carry assessjob-compatible score keys (`overall_score: <XX/200>`, `normalized_score: <XX%>`) and compare like any assessment. Their `synthetic: true` front matter MUST be called out in the report (label the column/row "synthetic benchmark") so an ideal-role benchmark is never mistaken for a real application.

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

**Read ALL assessment files in a single parallel batch using multiple Read tool calls.** Resolve each `{{ARGn}}` first:

- If `{{ARGn}}` contains `/` or ends in `.md`, read it directly as a file path (e.g., an `/idealjob` archetype file).
- Otherwise read `{config.directories.applications_root}/{{ARGn}}/assessment/assessment.md`.

Load `{{ARG1}}` and `{{ARG2}}` always; `{{ARG3}}` and `{{ARG4}}` if provided.

**CRITICAL**: Use parallel Read tool calls for all files in a single message. Do NOT read them sequentially.

> **Task:** Mark task 1 `completed`.

---

## PHASE 2: EXTRACT DATA

> **Task:** Mark task 2 `in_progress`.

Extract key information from each assessment:
- **Candidate Information**: Name, role applied for, company
- **Front Matter Keys**: `overall_score` (XX/200), `normalized_score` (XX%), `role_variant` (Technical IC | People Manager | Executive), and `synthetic` if present
- **Scoring Breakdown**: Per-category raw points, variant-adjusted maximum, and percentage for the five rubric categories (Skills Inventory, Experience Relevance, Demonstrated Impact, Credentials, Fit & Readiness)
- **Strengths & Gaps**: Key strengths and areas of concern
- **Hiring Recommendations**: Final recommendations and rationale
- **Assessment Date**: When evaluation was conducted

**Legacy inputs:** if an assessment carries a bare 0–100 `overall_score` (no `/200` suffix), treat that value as its normalized percentage and compare it at the overall level only — its category cells in the report read `N/A (legacy scale)`.

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

| Assessment | Role | Company | Variant | Overall Score | Recommendation | Date |
|------------|------|---------|---------|---------------|-----------------|------|
| {{ARG1}} | [Role] | [Company] | [Variant] | [XXX]/200 ([XX%]) | [Recommendation] | [Date] |
| {{ARG2}} | [Role] | [Company] | [Variant] | [XXX]/200 ([XX%]) | [Recommendation] | [Date] |
| {{ARG3}} | [Role] | [Company] | [Variant] | [XXX]/200 ([XX%]) | [Recommendation] | [Date] |
| {{ARG4}} | [Role] | [Company] | [Variant] | [XXX]/200 ([XX%]) | [Recommendation] | [Date] |

For any input whose front matter has `synthetic: true`, append ` (synthetic benchmark)` to its Company cell in this table and prefix its name with `Synthetic:` everywhere it appears in the report.

## Detailed Score Comparison

### Overall Performance Rankings

Rank by `normalized_score` — the percentage is the cross-assessment comparison basis; raw points are context.

1. **[Highest Scoring Role]**: [XXX]/200 ([XX%]) - [Company] - [Brief rationale]
2. **[Second Highest]**: [XXX]/200 ([XX%]) - [Company] - [Brief rationale]
3. **[Third]**: [XXX]/200 ([XX%]) - [Company] - [Brief rationale]
4. **[Lowest]**: [XXX]/200 ([XX%]) - [Company] - [Brief rationale]

### Category Performance Analysis

Category maxima vary by role variant (e.g., Skills Inventory is /60 for a Technical IC assessment but /40 for People Manager), so **compare on percentages, not raw points**. `[Max]` is each assessment's variant-adjusted category maximum from its rubric. For legacy 0–100 inputs, every category cell reads `N/A (legacy scale)`.

#### 1. Skills Inventory
| Role/Company | Points | Variant Max | % | Analysis |
|--------------|--------|-------------|---|----------|
| [Role 1] | [X] | [Max] | [XX%] | [Strength/Gap summary] |
| [Role 2] | [X] | [Max] | [XX%] | [Strength/Gap summary] |

**Key Insights:**
- [Consistent technical strengths across roles]
- [Skills gaps that appear across multiple assessments]
- [Role-specific technical advantages]

#### 2. Experience Relevance
| Role/Company | Points | Variant Max | % | Analysis |
|--------------|--------|-------------|---|----------|
| [Role 1] | [X] | [Max] | [XX%] | [Experience alignment] |
| [Role 2] | [X] | [Max] | [XX%] | [Experience alignment] |

**Key Insights:**
- [Best experience matches and why]
- [Experience transfer opportunities]

#### 3. Demonstrated Impact
| Role/Company | Points | Variant Max | % | Analysis |
|--------------|--------|-------------|---|----------|
| [Role 1] | [X] | [Max] | [XX%] | [Impact evidence] |
| [Role 2] | [X] | [Max] | [XX%] | [Impact evidence] |

#### 4. Credentials
| Role/Company | Points | Variant Max | % | Analysis |
|--------------|--------|-------------|---|----------|
| [Role 1] | [X] | [Max] | [XX%] | [Credential match] |
| [Role 2] | [X] | [Max] | [XX%] | [Credential match] |

#### 5. Fit & Readiness
| Role/Company | Points | Variant Max | % | Analysis |
|--------------|--------|-------------|---|----------|
| [Role 1] | [X] | [Max] | [XX%] | [Fit assessment] |
| [Role 2] | [X] | [Max] | [XX%] | [Fit assessment] |

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
- **Score advantage**: [X] percentage points above the average normalized score
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

Save the comparative analysis to: `{config.directories.career_analysis}/comparison_{YYYYMMDD}_{slug}.md`, where `{slug}` is built from the companies being compared, lowercased and joined with `_vs_` (e.g., `google_vs_meta`, or `google_vs_meta_vs_amazon` for 3+). For an input with no company (e.g., an `/idealjob` archetype), use its input filename stem instead (e.g., `google_vs_idealjob_20260604_anchor`).

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
claude /comparejobs JLL_VP_Office_Leasing_2025-09-25 Canerector_Vice_President_Real_Estate_2025-09-26

# Benchmark a real application against an /idealjob archetype (direct file path)
claude /comparejobs JLL_VP_Office_Leasing_2025-09-25 Career_Analysis/idealjob_20260604_anchor.md
```
(Each argument is either an application slug under `{config.directories.applications_root}` or a direct assessment-file path — see Your Task.)
