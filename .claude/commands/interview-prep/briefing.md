---
description: Generate a comprehensive briefing note to address skill gaps or prepare for interviews
argument-hint: <assessment-report> <job-description> [gaps-only|prep-time]
---

You are creating a strategic briefing note to help a candidate prepare for their interview or address identified skill gaps. This command generates a detailed study guide based on the assessment report and job requirements, with time-aware study prioritization.

## Your Task

Create a comprehensive briefing note that either:
1. **Gaps-only mode**: Focuses exclusively on addressing identified weaknesses and skill gaps
2. **Full preparation mode**: Covers all job requirements for complete interview preparation

Both modes include **Study Priority tags** that guide candidates on how to allocate their limited preparation time.

## Arguments

- `{{ARG1}}`: Assessment report file (required) - The candidate assessment output from /assessjob
- `{{ARG2}}`: Job description file (required) - The original job posting
- `{{ARG3}}`: Mode/time flag (optional) - One of:
  - `gaps-only` - Focus only on skill gaps
  - `1d`, `2d`, `3d` - Days available for preparation
  - `1w`, `2w` - Weeks available for preparation
  - If omitted: defaults to full preparation mode with 1 week assumed

---

## WORKFLOW ARCHITECTURE

```
Phase 1 (Parallel batch):     Load assessment + job description (2 parallel reads)
Phase 2 (PARALLEL):           Extract gaps (main agent) ‖ Research learning resources (subagent)
Phase 3 (Sequential, visible): Generate study guides by priority tier
Phase 4 (Sequential):         Create practice schedule + interview strategy → Save
```

**Dependency Rules:**
- Phase 1 loads both files in parallel
- Phase 2: Gap extraction (main agent) feeds into research scope, BUT research subagent can start with job posting requirements immediately while main agent extracts specific gaps
- Phase 3 WAITS for both gap extraction AND research results
- Phase 4 WAITS for all study guides

---

## PROGRESS TRACKING (MANDATORY)

**Before starting any work**, create all tasks for user visibility:

| # | Task Subject | activeForm |
|---|-------------|------------|
| 1 | Load assessment and job description | Loading assessment and job description |
| 2 | Extract gaps and requirements | Extracting gaps and requirements from assessment |
| 3 | Research learning resources | Researching learning resources and best practices |
| 4 | Generate skill gap analysis | Generating skill gap analysis with priorities |
| 5 | Create priority-based study plan | Creating priority-based study plan |
| 6 | Write detailed study guides | Writing detailed study guides for each topic |
| 7 | Create interview strategy | Creating interview strategy and gap acknowledgment scripts |
| 8 | Save briefing note | Saving briefing note to Briefing_Notes |

**Task Update Rules:**
- Mark each task `in_progress` BEFORE starting work on it
- Mark each task `completed` AFTER finishing it

---

## Understanding Gap Severity vs Study Priority

| Attribute | Gap Severity | Study Priority |
|-----------|--------------|----------------|
| **Question answered** | "How important is this skill to the role?" | "What should I study first given my time?" |
| **Factors considered** | Role requirements, dealbreaker status | Severity + prep time + learning curve + interview likelihood |
| **Purpose** | Understand role fit | Guide study actions |

### Study Priority Tags

| Tag | Meaning | When to Apply |
|-----|---------|---------------|
| 🔴 **FOCUS NOW** | Study this first | High-impact, learnable in available time, likely interview topic |
| 🟡 **IF TIME PERMITS** | Secondary priority | Important but lower ROI given time constraints |
| 🟢 **SKIM ONLY** | Quick review, don't deep-dive | Low-impact OR requires too much time to master |

### Priority Assignment Logic

```
IF (Gap Severity = Critical) AND (Time to Bridge <= Available Prep Time):
    → 🔴 FOCUS NOW

IF (Gap Severity = Critical) AND (Time to Bridge > Available Prep Time):
    → 🟢 SKIM ONLY (with gap acknowledgment strategy)

IF (Gap Severity = High) AND (Time to Bridge <= Available Prep Time * 0.5):
    → 🔴 FOCUS NOW

IF (Gap Severity = High) AND (Time to Bridge > Available Prep Time * 0.5):
    → 🟡 IF TIME PERMITS

IF (Gap Severity = Medium) AND (High interview likelihood) AND (Quick win possible):
    → 🔴 FOCUS NOW

IF (Gap Severity = Medium):
    → 🟡 IF TIME PERMITS

IF (Gap Severity = Low):
    → 🟢 SKIM ONLY
```

Additional factors that elevate priority:
- Candidate has adjacent skills (faster learning curve)
- Topic is commonly asked in interviews for this role
- Quick wins available (certifications, portfolio pieces)
- Topic affects multiple job requirements

---

## YAML FRONT MATTER

When saving the briefing to `Briefing_Notes/`, prepend this metadata:

```yaml
---
assessment_file: {{ARG1}}
job_file: Job_Postings/{{ARG2}}
mode: "<gaps-only | full>"
prep_time: "<parsed prep time, e.g., '3 days' or '1 week'>"
role: <role title>
company: <company name>
candidate: <full candidate name>
generated_by: /briefing
generated_on: <ISO8601 timestamp>
output_type: briefing
status: draft
version: 1.1
---
```

---

## PHASE 1: LOAD INPUTS (Parallel batch)

> **Task:** Mark task 1 `in_progress`.

**Read both files in a single parallel batch:**
- Assessment report from `{{ARG1}}` (check OutputResumes/ or Scoring_Rubrics/ directories)
- Job description from `Job_Postings/{{ARG2}}` (add .md extension if needed)

Parse {{ARG3}} to determine:
- Mode: gaps-only if {{ARG3}} == "gaps-only", otherwise full preparation
- Prep time: parse time value (1d=1 day, 3d=3 days, 1w=1 week, 2w=2 weeks), default 1 week

> **Task:** Mark task 1 `completed`.

---

## PHASE 2: PARALLEL ANALYSIS AND RESEARCH

> **CRITICAL: Dispatch both tasks simultaneously in a SINGLE message.**
> Mark tasks 2 and 3 as `in_progress` before dispatching.

### 2.1 Extract Gaps and Requirements (Task 2 - Main agent)

#### From Assessment Report:
- Identify all LOW scores (0-2 points or "None"/"Basic" ratings)
- Note MEDIUM scores (3-6 points or "Proficient" ratings) that need improvement
- Extract specific gaps, weaknesses, and areas of concern
- Review hiring manager recommendations and concerns

#### From Job Description:
- List all technical requirements and tools
- Extract key responsibilities and deliverables
- Identify industry-specific knowledge requirements
- Note soft skills and cultural fit expectations

> **Task:** Mark task 2 `completed`.

### 2.2 Research Learning Resources (Task 3 - Subagent)

**Dispatch a research subagent** to run concurrently with gap extraction:

```
Use Task tool with subagent_type=general-purpose and prompt:
"Research learning resources and preparation strategies for the role of [ROLE TITLE] at [COMPANY NAME].
The candidate needs to prepare for an interview with the following key requirements from the job posting:
[LIST TOP 8-10 REQUIREMENTS FROM JOB POSTING]

For each requirement area, find:
1. Best free/quick learning resources (tutorials, documentation, videos)
2. Realistic time-to-competency estimates for someone with adjacent skills
3. Common interview questions asked about this topic
4. Key terminology and concepts that demonstrate knowledge
5. Practical exercises or hands-on practice recommendations
6. Industry-specific best practices and current trends

Organize by topic area. Focus on resources that can be consumed in 1-14 days.
Prioritize practical, actionable resources over theoretical ones."
```

> **Task:** Mark task 3 `completed` when subagent returns.

---

## PHASE 3: GENERATE STUDY CONTENT (Sequential, with per-section visibility)

> **Prerequisites:** Both gap extraction (task 2) AND research (task 3) must be `completed`.

### 3.1 Generate Skill Gap Analysis

> **Task:** Mark task 4 `in_progress`.

For each gap/requirement, provide:
- **Current State**: Candidate's current level
- **Target State**: What the role requires
- **Gap Severity**: Critical/High/Medium/Low (role importance)
- **Time to Bridge**: Estimated hours/days needed
- **Study Priority**: 🔴 FOCUS NOW / 🟡 IF TIME PERMITS / 🟢 SKIM ONLY
- **Priority Rationale**: Brief explanation of why this priority was assigned
- **Minimum Viable Prep**: Absolute minimum hours for basic familiarity

> **Task:** Mark task 4 `completed`.

### 3.2 Create Priority-Based Study Plan

> **Task:** Mark task 5 `in_progress`.

Generate a day-by-day study plan based on available prep time:

**For 1-3 days available:**
- Day 1: FOCUS NOW topics (maximum impact)
- Day 2: FOCUS NOW completion + IF TIME PERMITS start
- Day 3: Review + SKIM ONLY + gap acknowledgment prep

**For 1-2 weeks available:**
- Days 1-2: FOCUS NOW topics (foundation building)
- Days 3-4: FOCUS NOW completion + IF TIME PERMITS start
- Days 5-6: IF TIME PERMITS + practice
- Day 7: Review + SKIM ONLY + final prep

Include executive summary with:
```
🔴 FOCUS NOW: X topics (estimated Y hours)
🟡 IF TIME PERMITS: X topics (estimated Y hours)
🟢 SKIM ONLY: X topics (estimated Y hours)
Total estimated prep time needed: Z hours
Available prep time: [from ARG3 or default]
Time sufficiency: [Sufficient / Tight / Insufficient - prioritization critical]
```

> **Task:** Mark task 5 `completed`.

### 3.3 Write Detailed Study Guides

> **Task:** Mark task 6 `in_progress`.

For each topic area (organized by Study Priority - FOCUS NOW topics first):

**1. Core Concepts** - Fundamental principles, key terminology, industry context

**2. Technical Deep Dive** - Implementation details, best practices, common pitfalls, tools

**3. Practical Exercises** - Hands-on tutorials, sample problems, real-world scenarios

**4. Interview Preparation** - Common questions, how to demonstrate knowledge, stories/examples

**5. Learning Resources** organized by timeframe:
- **Immediate** (1-2 days): Quick tutorials, video overviews, cheat sheets
- **Short-term** (1 week): Online courses, practical projects, community forums
- **Long-term** (2-4 weeks): Comprehensive courses, certifications, books

> **Task:** Mark task 6 `completed`.

---

## PHASE 4: INTERVIEW STRATEGY AND SAVE

### 4.1 Create Interview Strategy

> **Task:** Mark task 7 `in_progress`.

#### Acknowledging Gaps Professionally
For each 🟢 SKIM ONLY topic with Critical/High severity, provide a gap acknowledgment script:

```markdown
**Gap: [Topic Name]**
**Acknowledgment Script:**
"While I haven't had direct production experience with [Topic], I understand its importance
for this role. I've familiarized myself with [key concepts/terminology]. In my experience
with [related technology], I [relevant transferable experience]. I'm committed to ramping
up quickly—I've already [specific action taken] and have a learning plan that includes
[specific resources/timeline]."

**Key Points to Emphasize:**
- Awareness of the skill's importance
- Related/transferable experience
- Concrete learning plan
- Quick ramp-up evidence from past experience
```

#### Additional Strategy Sections:
- **Action Plan**: 24 hours / 3 days / 1 week before interview priorities
- **Quick Reference Guide**: One-page summary of key concepts for FOCUS NOW topics
- **Demonstrating learning ability and growth mindset**
- **Pivoting to related strengths**

> **Task:** Mark task 7 `completed`.

### 4.2 Customize and Save

> **Task:** Mark task 8 `in_progress`.

**If gaps-only mode:**
- Focus exclusively on addressing weaknesses
- Prioritize critical gaps that could be deal-breakers
- Provide accelerated learning paths
- Include mitigation strategies for unchangeable gaps
- Emphasize gap acknowledgment strategies for SKIM ONLY items

**If full preparation mode:**
- Cover all job requirements comprehensively
- Balance between strengthening weaknesses and showcasing strengths
- Include competitive differentiation strategies
- Prepare for technical deep-dives in strength areas

**Multi-Part Output Structure:**
- If the complete briefing exceeds 25,000 tokens, split into multiple parts
- Each part must not exceed 25,000 tokens
- Save parts as:
  - `Briefing_[Company]_[Role]_[Mode]_Part1_[Date].md`
  - `Briefing_[Company]_[Role]_[Mode]_Part2_[Date].md`
  - etc.
- Location: `Briefing_Notes/` directory

**Logical Splitting Guidelines:**
- Part 1: Executive Summary + Priority Breakdown + Skill Gap Analysis + Study Plan
- Part 2: Detailed Study Guides for FOCUS NOW and IF TIME PERMITS topics
- Part 3: SKIM ONLY study guides + Action Plan + Interview Strategy + Quick Reference
- Split at natural section boundaries, never mid-topic
- Each part includes header referencing total parts and navigation references

> **Task:** Mark task 8 `completed`.

---

## Quality Checks

Ensure the briefing note:
- Addresses every identified gap or requirement
- Assigns Study Priority tag to every topic
- Includes Priority Breakdown in Executive Summary
- Provides Priority-Based Study Plan with daily schedule
- Provides actionable, specific learning resources
- Includes realistic timelines for skill development
- Offers practical exercises and hands-on practice
- Prepares for likely interview scenarios
- Balances depth with time constraints
- Uses current, relevant industry information
- Includes gap acknowledgment scripts for SKIM ONLY critical/high items
- Each individual part does not exceed 25,000 tokens
- Multi-part reports split logically at section boundaries

## Output

The final briefing note should be a comprehensive, actionable study guide that:
1. Clearly identifies what needs to be learned with priority tags
2. Provides time-aware structured learning paths
3. Includes practical exercises and examples
4. Prepares for specific interview scenarios
5. Offers quick-reference materials for last-minute review
6. Provides professional gap acknowledgment strategies

The candidate should be able to use this briefing note as their primary preparation resource for the interview, with clear guidance on where to focus their limited preparation time.
