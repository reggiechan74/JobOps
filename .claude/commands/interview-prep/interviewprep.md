---
description: Generate customized interview questions based on job description and tailored resume
argument-hint: <resume-file> <job-description> [number-of-questions] [prep-time]
---

You are preparing a candidate for their interview by generating highly relevant interview questions based on the specific job description and their tailored resume that the employer has received. Questions are prioritized by likelihood of being asked to help candidates focus their limited practice time.

## Your Task

Generate a comprehensive set of interview questions that an employer would likely ask based on:
1. The specific requirements in the job description
2. The candidate's resume content that the employer is reviewing
3. Potential gaps or areas requiring clarification
4. Behavioral competencies relevant to the role

Each question receives a **Question Likelihood** tag to help candidates prioritize their preparation.

## Arguments

- `{{ARG1}}`: Tailored resume file (required) - The Step 3 final resume sent to employer
- `{{ARG2}}`: Job description file (required) - The original job posting
- `{{ARG3}}`: Number of questions (optional) - Defaults to 10 if not specified
- `{{ARG4}}`: Prep time (optional) - Practice time available: `1d`, `2d`, `3d`, `1w`, `2w`. Defaults to 3 days if not specified

---

## WORKFLOW ARCHITECTURE

```
Phase 1 (Parallel batch):     Load resume + job description (2 parallel reads)
Phase 2 (PARALLEL):           Analyze resume vs job ‖ Research interview patterns (subagent)
Phase 3 (Sequential, visible): Generate questions by category
Phase 4 (Sequential):         Create practice schedule → Save
```

**Dependency Rules:**
- Phase 1 loads both files in parallel
- Phase 2: Resume analysis (main agent) and interview pattern research (subagent) run concurrently
- Phase 3 WAITS for both Phase 2 tasks
- Phase 4 WAITS for all questions generated

---

## PROGRESS TRACKING (MANDATORY)

**Before starting any work**, create all tasks for user visibility:

| # | Task Subject | activeForm |
|---|-------------|------------|
| 1 | Load resume and job description | Loading resume and job description |
| 2 | Analyze resume vs job requirements | Analyzing resume vs job requirements |
| 3 | Research common interview patterns for role | Researching common interview patterns for role |
| 4 | Generate technical competency questions | Generating technical competency questions |
| 5 | Generate behavioral and situational questions | Generating behavioral and situational questions |
| 6 | Generate experience verification questions | Generating experience verification questions |
| 7 | Create practice schedule | Creating priority-based practice schedule |
| 8 | Save interview prep guide | Saving interview preparation guide |

**Task Update Rules:**
- Mark each task `in_progress` BEFORE starting work on it
- Mark each task `completed` AFTER finishing it

---

## Question Likelihood Tags

| Tag | Meaning | When to Apply |
|-----|---------|---------------|
| 🔴 **HIGH LIKELIHOOD** | Almost certain to be asked | Core job requirements, obvious resume claims to verify, standard behavioral questions for role level |
| 🟡 **MODERATE LIKELIHOOD** | Likely to come up | Secondary requirements, deeper probes on experience, role-specific scenarios |
| 🟢 **LOW LIKELIHOOD** | Possible but less common | Nice-to-have skills, edge cases, advanced scenarios |

### Likelihood Assignment Logic

```
IF (Question addresses must-have job requirement) AND (Directly verifies resume claim):
    → 🔴 HIGH LIKELIHOOD

IF (Question is standard behavioral for role level) AND (Common in industry):
    → 🔴 HIGH LIKELIHOOD

IF (Question probes gap between resume and requirements):
    → 🔴 HIGH LIKELIHOOD

IF (Question addresses preferred/secondary requirement):
    → 🟡 MODERATE LIKELIHOOD

IF (Question digs deeper into already-verified areas):
    → 🟡 MODERATE LIKELIHOOD

IF (Question addresses nice-to-have skills):
    → 🟢 LOW LIKELIHOOD

IF (Question covers edge cases or advanced scenarios):
    → 🟢 LOW LIKELIHOOD
```

---

## YAML FRONT MATTER

Save the question set to `Briefing_Notes/Interview_Prep_*` with this prefix:

```yaml
---
resume_file: {{ARG1}}
job_file: Job_Postings/{{ARG2}}
question_count: {{ARG3 or 10}}
prep_time: "<parsed prep time, e.g., '3 days' or '1 week'>"
role: <role title>
company: <company name>
candidate: <full candidate name>
generated_by: /interviewprep
generated_on: <ISO8601 timestamp>
output_type: interview_prep
status: draft
version: 1.1
---
```

---

## PHASE 1: LOAD INPUTS (Parallel batch)

> **Task:** Mark task 1 `in_progress`.

**Read both files in a single parallel batch:**
- Resume from `{{ARG1}}` (check OutputResumes/ directory)
- Job description from `Job_Postings/{{ARG2}}` (add .md extension if needed)

Set question count: Use {{ARG3}} if provided, otherwise default to 10.
Parse prep time: Use {{ARG4}} if provided, otherwise default to 3 days.

> **Task:** Mark task 1 `completed`.

---

## PHASE 2: PARALLEL ANALYSIS

> **CRITICAL: Dispatch both tasks simultaneously in a SINGLE message.**
> Mark tasks 2 and 3 as `in_progress` before dispatching.

### 2.1 Analyze Resume vs Job Requirements (Task 2 - Main agent)

#### From the Resume:
- Identify strongest qualifications and achievements
- Note any potential gaps or areas needing clarification
- Extract specific projects, metrics, and accomplishments
- Review technical skills and certifications claimed
- Identify leadership and team experiences

#### From the Job Description:
- Extract must-have technical requirements
- Identify key responsibilities and expectations
- Note soft skills and cultural fit requirements
- Understand reporting structure and stakeholder interactions
- Identify specific industry or domain requirements

> **Task:** Mark task 2 `completed`.

### 2.2 Research Common Interview Patterns (Task 3 - Subagent)

**Dispatch a research subagent** to run concurrently with resume analysis:

```
Use Task tool with subagent_type=general-purpose, model=sonnet, and prompt:
"Research common interview questions and patterns for [ROLE TITLE] at [COMPANY NAME] or similar roles:

1. Most commonly asked technical questions for this type of role
2. Standard behavioral questions at this seniority level
3. Industry-specific scenario questions
4. Company-specific interview format or culture (if publicly known)
5. Common follow-up probes interviewers use for this role type
6. Red flags that interviewers typically screen for at this level

Provide a structured summary organized by question category.
Focus on what's most likely to actually be asked in an interview."
```

> **Task:** Mark task 3 `completed` when subagent returns.

---

## PHASE 3: GENERATE QUESTIONS (Sequential, per-category visibility)

> **Prerequisites:** Both analysis (task 2) and research (task 3) must be `completed`.

Distribute questions across categories (adjust based on total number):
- **Technical Competency** (30-40% of questions)
- **Behavioral/Situational** (30-40% of questions)
- **Experience Verification** (20-30% of questions)
- **Cultural Fit** (10-20% of questions)
- **Role-Specific Scenarios** (10-20% of questions)

### 3.1 Generate Technical Competency Questions

> **Task:** Mark task 4 `in_progress`.

For each question, provide:

**A. The Question** - Clear, specific, with Likelihood tag (🔴/🟡/🟢)

**B. Question Context** - Why likely asked, what's being assessed, likelihood rationale

**C. Suggested Answer Approach** - Key points, STAR format guidance, specific resume examples, red flags to avoid

**D. Follow-up Questions** - 2-3 likely follow-up probes

**E. Practice Priority** - Minimum/recommended practice time and method

Question types:
- **Verify Claimed Expertise**: "You mention [specific skill/technology] on your resume. Can you describe a complex problem you solved using this?"
- **Applied Knowledge**: "How would you approach [specific technical challenge relevant to role]?"
- **Best Practices**: "What's your methodology for [key technical process]?"
- **Trade-offs**: "When would you choose [approach A] over [approach B]?"

> **Task:** Mark task 4 `completed`.

### 3.2 Generate Behavioral and Situational Questions

> **Task:** Mark task 5 `in_progress`.

- **Leadership**: "Tell me about a time when you led a team through a challenging project"
- **Problem-Solving**: "Describe a situation where you identified and solved a critical issue"
- **Stakeholder Management**: "Give an example of managing conflicting stakeholder priorities"
- **Innovation**: "Share an instance where you improved a process or introduced innovation"
- **Failure/Learning**: "Tell me about a project that didn't go as planned"

Also generate Cultural Fit and Role-Specific Scenario questions in this step.

> **Task:** Mark task 5 `completed`.

### 3.3 Generate Experience Verification Questions

> **Task:** Mark task 6 `in_progress`.

- **Metrics Validation**: "You achieved [specific metric]. Walk me through how you measured and achieved this"
- **Role Clarification**: "Explain your specific role in [major project mentioned]"
- **Impact Assessment**: "What was the business impact of [specific achievement]?"

Customize based on gaps between resume and requirements. Mark gap-probing questions as 🔴 HIGH LIKELIHOOD.

> **Task:** Mark task 6 `completed`.

---

## PHASE 4: PRACTICE SCHEDULE AND SAVE

### 4.1 Create Practice Schedule

> **Task:** Mark task 7 `in_progress`.

Generate a day-by-day practice schedule based on available prep time:

**For 1-2 days available:**
- Day 1: HIGH LIKELIHOOD questions (4-5 hours) + mock practice
- Day 2: MODERATE LIKELIHOOD + review HIGH + opening/closing statements

**For 3-5 days available:**
- Day 1: All HIGH LIKELIHOOD - full STAR answers written
- Day 2: MODERATE LIKELIHOOD - outline key points + technical prep
- Day 3: Full mock interview + LOW LIKELIHOOD skim + questions for interviewer

**For 1-2 weeks available:**
- Days 1-2: HIGH LIKELIHOOD mastery
- Days 3-4: MODERATE LIKELIHOOD development
- Days 5-6: Mock interviews + refinement
- Day 7: Final review + LOW LIKELIHOOD skim

Include strategic guidance:
- **Opening Statement**: 60-second pitch tailored to role
- **Closing Statement**: Summary of fit and interest
- **Questions for Interviewer**: 5 intelligent questions about role/company
- **Thank You Note Template**: Post-interview follow-up template

> **Task:** Mark task 7 `completed`.

### 4.2 Format and Save

> **Task:** Mark task 8 `in_progress`.

Create a structured interview preparation guide with:
- Executive Summary with likelihood breakdown
- Practice Schedule based on prep time
- All questions organized by likelihood (HIGH first, then MODERATE, then LOW)
- Quick Reference Sheet

**Multi-Part Output Structure:**
- If the complete guide exceeds 25,000 tokens, split into multiple parts
- Each part must not exceed 25,000 tokens
- Save parts as:
  - `Interview_Prep_[Company]_[Role]_Part1_[Date].md`
  - `Interview_Prep_[Company]_[Role]_Part2_[Date].md`
  - etc.
- Location: `Briefing_Notes/` directory

**Logical Splitting Guidelines:**
- Part 1: Executive Summary + Practice Schedule + HIGH LIKELIHOOD questions
- Part 2: MODERATE LIKELIHOOD questions
- Part 3: LOW LIKELIHOOD questions + Strategic Guidance + Quick Reference Sheet
- Split at natural question boundaries, never mid-question
- Each part includes header referencing total parts and navigation references

> **Task:** Mark task 8 `completed`.

---

## Quality Checks

Ensure the interview questions:
- Directly relate to job requirements or resume content
- Each question has a Likelihood tag (🔴/🟡/🟢)
- Likelihood Breakdown included in Executive Summary
- Practice Schedule generated based on prep time
- Balance technical and behavioral assessment
- Include both strengths-based and gap-probing questions
- Provide actionable guidance for responses
- Cover all critical competencies for the role
- Questions organized by likelihood (HIGH first)
- Include scenario-based questions relevant to actual job duties
- Each individual part does not exceed 25,000 tokens
- Multi-part reports split logically at question boundaries

## Output

The final interview preparation guide should:
1. Build candidate confidence through thorough preparation
2. **Prioritize practice time on most likely questions**
3. Anticipate likely questions based on specific resume/job alignment
4. Provide strategic response frameworks
5. Include recovery strategies for difficult questions
6. Prepare candidate for various interview formats (phone, video, panel)

The candidate should be able to use this guide for focused interview practice, knowing exactly which questions deserve the most preparation time.
