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

## Understanding Gap Severity vs Study Priority

These are two distinct concepts that work together:

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

Assign Study Priority based on this decision matrix:

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

## Step-by-Step Process

### YAML front matter for briefing note
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

Insert it before the first heading and refresh fields each time you regenerate.

### 1. Load Required Documents
✓ Loading reconnaissance assessment and target specifications
- Read the assessment report from `{{ARG1}}` (check OutputResumes/ or Scoring_Rubrics/ directories)
- Read the job description from `Job_Postings/{{ARG2}}` (add .md extension if needed)
- Parse {{ARG3}} to determine:
  - Mode: gaps-only if {{ARG3}} == "gaps-only", otherwise full preparation
  - Prep time: parse time value (1d=1 day, 3d=3 days, 1w=1 week, 2w=2 weeks), default 1 week

### 2. Extract Key Information

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

### 3. Conduct Research for Each Gap/Requirement
✓ Gathering intelligence on skill gaps and requirements

Use WebSearch and WebFetch to research:
- Current best practices and methodologies
- Industry standards and trends
- Technical concepts and implementations
- Certification requirements and preparation materials
- Real-world application examples
- **Realistic time-to-competency estimates for each skill**

### 4. Generate Briefing Note Structure
✓ Compiling intelligence brief with training protocol

Create a comprehensive briefing document with the following sections:

#### A. Executive Summary
- Overall preparedness assessment
- **Priority Breakdown**:
  ```
  🔴 FOCUS NOW: X topics (estimated Y hours)
  🟡 IF TIME PERMITS: X topics (estimated Y hours)
  🟢 SKIM ONLY: X topics (estimated Y hours)
  Total estimated prep time needed: Z hours
  Available prep time: [from ARG3 or default]
  Time sufficiency: [Sufficient / Tight / Insufficient - prioritization critical]
  ```
- Critical gaps requiring immediate attention
- Recommended preparation timeline
- Key focus areas for interview success

#### B. Skill Gap Analysis (or Full Requirements Overview)
For each gap/requirement, provide:
- **Current State**: Candidate's current level
- **Target State**: What the role requires
- **Gap Severity**: Critical/High/Medium/Low (role importance)
- **Time to Bridge**: Estimated hours/days needed
- **Study Priority**: 🔴 FOCUS NOW / 🟡 IF TIME PERMITS / 🟢 SKIM ONLY
- **Priority Rationale**: Brief explanation of why this priority was assigned

Example format:
```markdown
##### Kubernetes Orchestration
- **Current State**: No direct experience; familiar with Docker
- **Target State**: Production cluster management
- **Gap Severity**: Critical (core job requirement)
- **Time to Bridge**: 40+ hours for competency
- **Study Priority**: 🟢 SKIM ONLY
- **Priority Rationale**: Despite critical severity, requires more time than available; focus on terminology and concepts, prepare gap acknowledgment strategy
- **Minimum Viable Prep**: 2 hours (understand key concepts, prepare to discuss learning plan)
```

```markdown
##### Python Scripting
- **Current State**: Basic scripting experience
- **Target State**: Automation and tooling development
- **Gap Severity**: Medium (used for automation tasks)
- **Time to Bridge**: 4-6 hours with existing programming background
- **Study Priority**: 🔴 FOCUS NOW
- **Priority Rationale**: Adjacent skills accelerate learning; high ROI for time invested; likely technical screen topic
- **Minimum Viable Prep**: 3 hours (complete practice exercises, prepare code samples)
```

#### C. Priority-Based Study Plan

Generate a day-by-day study plan based on available prep time:

**For 1-3 days available:**
```markdown
## Priority-Based Study Plan (3 Days Available)

### Day 1: FOCUS NOW Topics (Priority: Maximum Impact)
- [ ] [Topic 1] (X hrs) - 🔴 FOCUS NOW
- [ ] [Topic 2] (X hrs) - 🔴 FOCUS NOW
Daily total: X hours

### Day 2: FOCUS NOW Completion + IF TIME PERMITS
- [ ] [Topic 3] (X hrs) - 🔴 FOCUS NOW
- [ ] [Topic 4] (X hrs) - 🟡 IF TIME PERMITS
Daily total: X hours

### Day 3: Review + SKIM ONLY
- [ ] Review FOCUS NOW topics (2 hrs)
- [ ] [Topic 5] terminology review (30 min) - 🟢 SKIM ONLY
- [ ] [Topic 6] quick overview (30 min) - 🟢 SKIM ONLY
- [ ] Prepare gap acknowledgment talking points (30 min)
Daily total: X hours
```

**For 1-2 weeks available:**
```markdown
## Priority-Based Study Plan (1 Week Available)

### Days 1-2: FOCUS NOW Topics (Foundation Building)
[List topics with time allocations]

### Days 3-4: FOCUS NOW Completion + IF TIME PERMITS
[List topics with time allocations]

### Days 5-6: IF TIME PERMITS + Practice
[List topics with time allocations]

### Day 7: Review + SKIM ONLY + Final Prep
[List review activities and SKIM ONLY topics]
```

#### D. Detailed Study Guide
For each topic area (organized by Study Priority - FOCUS NOW topics first):

##### 1. Core Concepts
- Fundamental principles and theory
- Key terminology and definitions
- Industry context and applications

##### 2. Technical Deep Dive
- Implementation details
- Best practices and patterns
- Common pitfalls to avoid
- Tools and technologies overview

##### 3. Practical Exercises
- Hands-on tutorials and labs
- Sample problems with solutions
- Real-world scenarios to practice
- Code examples and implementations

##### 4. Interview Preparation
- Common interview questions for this topic
- How to demonstrate knowledge effectively
- Stories/examples to showcase understanding
- Whiteboard problem approaches

##### 5. Learning Resources
- **Immediate** (1-2 days):
  - Quick tutorials and documentation
  - Video overviews and crash courses
  - Cheat sheets and reference guides

- **Short-term** (1 week):
  - Online courses and workshops
  - Practical projects to build
  - Community forums and discussions

- **Long-term** (2-4 weeks):
  - Comprehensive courses or certifications
  - Books and in-depth resources
  - Mentorship opportunities

#### E. Action Plan
- **24 Hours Before Interview**:
  - Critical topics to review (FOCUS NOW items only)
  - Key talking points to prepare
  - Questions to ask the interviewer
  - Gap acknowledgment scripts ready

- **3 Days Before Interview**:
  - Hands-on practice priorities
  - Mock interview topics
  - Portfolio pieces to prepare

- **1 Week Before Interview**:
  - Deeper learning objectives
  - Projects to complete
  - Skills to demonstrate

#### F. Interview Strategy

##### Acknowledging Gaps Professionally
For each 🟢 SKIM ONLY topic with Critical/High severity, provide a gap acknowledgment script:

```markdown
**Gap: [Topic Name]**
**Acknowledgment Script:**
"While I haven't had direct production experience with [Topic], I understand its importance for this role. I've familiarized myself with [key concepts/terminology]. In my experience with [related technology], I [relevant transferable experience]. I'm committed to ramping up quickly—I've already [specific action taken] and have a learning plan that includes [specific resources/timeline]."

**Key Points to Emphasize:**
- Awareness of the skill's importance
- Related/transferable experience
- Concrete learning plan
- Quick ramp-up evidence from past experience
```

##### Additional Interview Strategies
- Demonstrating learning ability and growth mindset
- Pivoting to related strengths
- Using transferable skills effectively
- Turning gaps into growth narrative

#### G. Quick Reference Guide
- One-page summary of key concepts (FOCUS NOW topics)
- Formula sheet or command reference
- Common acronyms and terminology
- Industry-specific metrics and KPIs

### 5. Customize Based on Mode

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

### 6. Format and Save

**Multi-Part Output Structure:**
- If the complete briefing exceeds 25,000 tokens, split it into multiple parts
- Each part must not exceed 25,000 tokens
- Save parts as:
  - `Briefing_[Company]_[Role]_[Mode]_Part1_[Date].md`
  - `Briefing_[Company]_[Role]_[Mode]_Part2_[Date].md`
  - `Briefing_[Company]_[Role]_[Mode]_Part3_[Date].md`
  - etc.
- Location: `Briefing_Notes/` directory
- Format: Markdown with clear headings and sections

**Logical Splitting Guidelines:**
- Part 1: Executive Summary + Priority Breakdown + Skill Gap Analysis + Priority-Based Study Plan
- Part 2: Detailed Study Guides for FOCUS NOW and IF TIME PERMITS topics
- Part 3: SKIM ONLY study guides + Action Plan + Interview Strategy + Quick Reference
- Split at natural section boundaries, never mid-topic
- Each part should include a header referencing total parts (e.g., "Part 1 of 3")
- Include navigation references at start/end of each part

## Quality Checks

Ensure the briefing note:
- ✅ Addresses every identified gap or requirement
- ✅ Assigns Study Priority tag to every topic
- ✅ Includes Priority Breakdown in Executive Summary
- ✅ Provides Priority-Based Study Plan with daily schedule
- ✅ Provides actionable, specific learning resources
- ✅ Includes realistic timelines for skill development
- ✅ Offers practical exercises and hands-on practice
- ✅ Prepares for likely interview scenarios
- ✅ Balances depth with time constraints
- ✅ Uses current, relevant industry information
- ✅ Includes gap acknowledgment scripts for SKIM ONLY critical/high items
- ✅ **Each individual part does not exceed 25,000 tokens**
- ✅ **Multi-part reports split logically at section boundaries**

## Output

The final briefing note should be a comprehensive, actionable study guide that:
1. Clearly identifies what needs to be learned with priority tags
2. Provides time-aware structured learning paths
3. Includes practical exercises and examples
4. Prepares for specific interview scenarios
5. Offers quick-reference materials for last-minute review
6. Provides professional gap acknowledgment strategies

**MULTI-PART OUTPUT REQUIREMENTS**:
- Create as many parts as needed to cover all content comprehensively
- **Each individual part must not exceed 25,000 tokens**
- Split content logically at section boundaries (never mid-topic)
- Each part should be self-contained enough to be useful independently
- Include clear navigation between parts (e.g., "Continued in Part 2", "Continued from Part 1")
- Part 1 should always include the Executive Summary, Priority Breakdown, and Study Plan
- Final part should always include the Interview Strategy and Quick Reference Guide

The candidate should be able to use this briefing note as their primary preparation resource for the interview, with clear guidance on where to focus their limited preparation time.
