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

## Step-by-Step Process

### YAML front matter for interview prep output
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

Include the block before the first heading and update counts/timestamps on reruns.

### 1. Load Required Documents
- Read the tailored resume from `{{ARG1}}` (check OutputResumes/ directory)
- Read the job description from `Job_Postings/{{ARG2}}` (add .md extension if needed)
- Set question count: Use {{ARG3}} if provided, otherwise default to 10
- Parse prep time: Use {{ARG4}} if provided, otherwise default to 3 days

### 2. Analyze Resume vs. Job Requirements

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

### 3. Generate Question Categories

Distribute questions across these categories (adjust based on total number):
- **Technical Competency** (30-40% of questions)
- **Behavioral/Situational** (30-40% of questions)
- **Experience Verification** (20-30% of questions)
- **Cultural Fit** (10-20% of questions)
- **Role-Specific Scenarios** (10-20% of questions)

### 4. Question Generation Framework

For each question, provide:

#### A. The Question
- Clear, specific, and relevant to the role
- Appropriate difficulty level for the position
- Mix of open-ended and specific formats
- **Question Likelihood tag** (🔴/🟡/🟢)

#### B. Question Context
- Why this question is likely to be asked
- What the interviewer is trying to assess
- Connection to job requirements or resume claims
- **Likelihood rationale** - why this probability level

#### C. Suggested Answer Approach
- Key points to cover in response
- STAR format guidance where applicable
- Specific examples from resume to reference
- Red flags to avoid

#### D. Follow-up Questions
- 2-3 likely follow-up probes
- How to handle deeper drilling
- Areas where interviewer might challenge

#### E. Practice Priority
- **Minimum practice time**: X minutes
- **Recommended practice time**: Y minutes
- **Practice method**: Mock answer, written outline, or quick mental review

### 5. Question Types to Include

#### Technical Questions
- **Verify Claimed Expertise**: "You mention [specific skill/technology] on your resume. Can you describe a complex problem you solved using this?"
- **Applied Knowledge**: "How would you approach [specific technical challenge relevant to role]?"
- **Best Practices**: "What's your methodology for [key technical process]?"
- **Trade-offs**: "When would you choose [approach A] over [approach B]?"

#### Behavioral Questions (STAR Format)
- **Leadership**: "Tell me about a time when you led a team through a challenging project"
- **Problem-Solving**: "Describe a situation where you identified and solved a critical issue"
- **Stakeholder Management**: "Give an example of managing conflicting stakeholder priorities"
- **Innovation**: "Share an instance where you improved a process or introduced innovation"
- **Failure/Learning**: "Tell me about a project that didn't go as planned"

#### Experience Verification
- **Metrics Validation**: "You achieved [specific metric]. Walk me through how you measured and achieved this"
- **Role Clarification**: "Explain your specific role in [major project mentioned]"
- **Impact Assessment**: "What was the business impact of [specific achievement]?"

#### Cultural Fit
- **Work Style**: "How do you prefer to structure your workday and priorities?"
- **Team Dynamics**: "Describe your ideal team environment"
- **Values Alignment**: "Why are you interested in [company/organization]?"

#### Role-Specific Scenarios
- **Hypothetical Situations**: "If faced with [specific scenario related to role], how would you handle it?"
- **Strategic Thinking**: "How would you approach [strategic challenge specific to position]?"
- **First 90 Days**: "What would be your priorities in your first three months?"

### 6. Customize Based on Gaps

If the analysis reveals potential gaps between resume and requirements:
- Include questions that probe transferable skills
- Add questions about learning agility and adaptation
- Include scenario questions to test problem-solving without direct experience
- Mark gap-probing questions as 🔴 HIGH LIKELIHOOD

### 7. Generate Priority-Based Practice Schedule

Based on available prep time, create a practice schedule:

**For 1-2 days available:**
```markdown
## Practice Schedule (2 Days Available)

### Day 1: HIGH LIKELIHOOD Questions (4-5 hours)
- [ ] Q1: [Question summary] (30 min) - 🔴 HIGH LIKELIHOOD
- [ ] Q2: [Question summary] (30 min) - 🔴 HIGH LIKELIHOOD
- [ ] Q3: [Question summary] (30 min) - 🔴 HIGH LIKELIHOOD
- [ ] Mock interview practice with HIGH LIKELIHOOD questions (2 hrs)

### Day 2: MODERATE LIKELIHOOD + Review (3-4 hours)
- [ ] Q4: [Question summary] (20 min) - 🟡 MODERATE LIKELIHOOD
- [ ] Q5: [Question summary] (20 min) - 🟡 MODERATE LIKELIHOOD
- [ ] Review HIGH LIKELIHOOD answers (1 hr)
- [ ] LOW LIKELIHOOD quick review (30 min) - 🟢 Skim only
- [ ] Opening/closing statements practice (30 min)
```

**For 3-5 days available:**
```markdown
## Practice Schedule (3 Days Available)

### Day 1: HIGH LIKELIHOOD Deep Practice
- [ ] All HIGH LIKELIHOOD questions - full STAR answers written
- [ ] Record and review mock answers

### Day 2: MODERATE LIKELIHOOD + Technical Prep
- [ ] MODERATE LIKELIHOOD questions - outline key points
- [ ] Technical demonstration practice

### Day 3: Full Mock Interview + LOW LIKELIHOOD Skim
- [ ] Complete mock interview (all HIGH + select MODERATE)
- [ ] Quick review of LOW LIKELIHOOD questions
- [ ] Prepare questions for interviewer
```

**For 1-2 weeks available:**
```markdown
## Practice Schedule (1 Week Available)

### Days 1-2: HIGH LIKELIHOOD Mastery
[Detailed schedule]

### Days 3-4: MODERATE LIKELIHOOD Development
[Detailed schedule]

### Days 5-6: Mock Interviews + Refinement
[Detailed schedule]

### Day 7: Final Review + LOW LIKELIHOOD Skim
[Detailed schedule]
```

### 8. Format and Save

Create a structured interview preparation guide:

```markdown
# Interview Preparation Guide
## [Position Title] at [Company Name]

### Executive Summary
- Total Questions: [Number]
- Estimated Interview Duration: [Time estimate]
- **Likelihood Breakdown**:
  ```
  🔴 HIGH LIKELIHOOD: X questions (practice time: Y hours)
  🟡 MODERATE LIKELIHOOD: X questions (practice time: Y hours)
  🟢 LOW LIKELIHOOD: X questions (practice time: Y minutes)
  Total recommended practice time: Z hours
  Available prep time: [from ARG4 or default]
  ```
- Key Focus Areas: [Top 3-4 areas]
- Critical Success Factors: [What will make candidate successful]

### Practice Schedule
[Generated based on prep time]

### Question Set

#### Question 1: [Category - e.g., Technical] 🔴 HIGH LIKELIHOOD
**Question:** [The actual question]

**Likelihood Rationale:** [Why this is likely to be asked]

**Context:** [What interviewer is assessing]

**Recommended Approach:**
- [Key point 1]
- [Key point 2]
- [Example from resume to reference]

**Practice Priority:**
- Minimum: X minutes
- Recommended: Y minutes
- Method: [Mock answer / Written outline / Mental review]

**Potential Follow-ups:**
1. [Follow-up question 1]
2. [Follow-up question 2]

**Red Flags to Avoid:**
- [What not to say/do]

---

[Repeat for all questions, organized by likelihood: HIGH first, then MODERATE, then LOW]

### Quick Reference Sheet
- Top 3 achievements to emphasize
- Key technical skills to demonstrate
- Behavioral examples to prepare (STAR format)
- Questions to ask the interviewer
```

**Multi-Part Output Structure:**
- If the complete guide exceeds 25,000 tokens, split it into multiple parts
- Each part must not exceed 25,000 tokens
- Save parts as:
  - `Interview_Prep_[Company]_[Role]_Part1_[Date].md`
  - `Interview_Prep_[Company]_[Role]_Part2_[Date].md`
  - `Interview_Prep_[Company]_[Role]_Part3_[Date].md`
  - etc.
- Location: `Briefing_Notes/` directory

**Logical Splitting Guidelines:**
- Part 1: Executive Summary + Practice Schedule + HIGH LIKELIHOOD questions
- Part 2: MODERATE LIKELIHOOD questions
- Part 3: LOW LIKELIHOOD questions + Strategic Guidance + Quick Reference Sheet
- Split at natural question boundaries, never mid-question
- Each part should include a header referencing total parts (e.g., "Part 1 of 3")
- Include navigation references at start/end of each part

### 9. Include Strategic Guidance

Add a section with:
- **Opening Statement**: 60-second pitch tailored to role
- **Closing Statement**: Summary of fit and interest
- **Questions for Interviewer**: 5 intelligent questions about role/company
- **Thank You Note Template**: Post-interview follow-up template

## Quality Checks

Ensure the interview questions:
- ✅ Directly relate to job requirements or resume content
- ✅ Each question has a Likelihood tag (🔴/🟡/🟢)
- ✅ Likelihood Breakdown included in Executive Summary
- ✅ Practice Schedule generated based on prep time
- ✅ Balance technical and behavioral assessment
- ✅ Include both strengths-based and gap-probing questions
- ✅ Provide actionable guidance for responses
- ✅ Cover all critical competencies for the role
- ✅ Questions organized by likelihood (HIGH first)
- ✅ Include scenario-based questions relevant to actual job duties
- ✅ **Each individual part does not exceed 25,000 tokens**
- ✅ **Multi-part reports split logically at question boundaries**

## Output

The final interview preparation guide should:
1. Build candidate confidence through thorough preparation
2. **Prioritize practice time on most likely questions**
3. Anticipate likely questions based on specific resume/job alignment
4. Provide strategic response frameworks
5. Include recovery strategies for difficult questions
6. Prepare candidate for various interview formats (phone, video, panel)

**MULTI-PART OUTPUT REQUIREMENTS**:
- Create as many parts as needed to cover all questions comprehensively
- **Each individual part must not exceed 25,000 tokens**
- Split content logically at question boundaries (never mid-question)
- Each part should be self-contained enough to be useful independently
- Include clear navigation between parts (e.g., "Continued in Part 2", "Continued from Part 1")
- Part 1 should always include Executive Summary, Practice Schedule, and HIGH LIKELIHOOD questions
- Final part should always include Strategic Guidance and Quick Reference Sheet

The candidate should be able to use this guide for focused interview practice, knowing exactly which questions deserve the most preparation time.
