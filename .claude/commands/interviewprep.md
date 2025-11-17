---
description: Generate customized interview questions based on job description and tailored resume
argument-hint: <resume-file> <job-description> [number-of-questions]
---

You are preparing a candidate for their interview by generating highly relevant interview questions based on the specific job description and their tailored resume that the employer has received.

## Your Task

Generate a comprehensive set of interview questions that an employer would likely ask based on:
1. The specific requirements in the job description
2. The candidate's resume content that the employer is reviewing
3. Potential gaps or areas requiring clarification
4. Behavioral competencies relevant to the role

## Arguments

- `{{ARG1}}`: Tailored resume file (required) - The Step 3 final resume sent to employer
- `{{ARG2}}`: Job description file (required) - The original job posting
- `{{ARG3}}`: Number of questions (optional) - Defaults to 10 if not specified

## Step-by-Step Process

### YAML front matter for interview prep output
Save the question set to `OutputResumes/Interview_Prep_*` with this prefix:

```yaml
---
resume_file: {{ARG1}}
job_file: Job_Postings/{{ARG2}}
question_count: {{ARG3 or 10}}
role: <role title>
company: <company name>
candidate: <full candidate name>
generated_by: /interviewprep
generated_on: <ISO8601 timestamp>
output_type: interview_prep
status: draft
version: 1.0
---
```

Include the block before the first heading and update counts/timestamps on reruns.

### 1. Load Required Documents
- Read the tailored resume from `{{ARG1}}` (check OutputResumes/ directory)
- Read the job description from `Job_Postings/{{ARG2}}` (add .md extension if needed)
- Set question count: Use {{ARG3}} if provided, otherwise default to 10

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

#### B. Question Context
- Why this question is likely to be asked
- What the interviewer is trying to assess
- Connection to job requirements or resume claims

#### C. Suggested Answer Approach
- Key points to cover in response
- STAR format guidance where applicable
- Specific examples from resume to reference
- Red flags to avoid

#### D. Follow-up Questions
- 2-3 likely follow-up probes
- How to handle deeper drilling
- Areas where interviewer might challenge

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

### 7. Format and Save

Create a structured interview preparation guide:

```markdown
# Interview Preparation Guide
## [Position Title] at [Company Name]

### Executive Summary
- Total Questions: [Number]
- Estimated Interview Duration: [Time estimate]
- Key Focus Areas: [Top 3-4 areas]
- Critical Success Factors: [What will make candidate successful]

### Question Set

#### Question 1: [Category - e.g., Technical]
**Question:** [The actual question]

**Context:** [Why this will be asked]

**Recommended Approach:**
- [Key point 1]
- [Key point 2]
- [Example from resume to reference]

**Potential Follow-ups:**
1. [Follow-up question 1]
2. [Follow-up question 2]

---

[Repeat for all questions]

### Quick Reference Sheet
- Top 3 achievements to emphasize
- Key technical skills to demonstrate
- Behavioral examples to prepare (STAR format)
- Questions to ask the interviewer
```

Save as:
- Filename: `Interview_Prep_[Company]_[Role]_[Date].md`
- Location: `Briefing_Notes/` directory

### 8. Include Strategic Guidance

Add a section with:
- **Opening Statement**: 60-second pitch tailored to role
- **Closing Statement**: Summary of fit and interest
- **Questions for Interviewer**: 5 intelligent questions about role/company
- **Thank You Note Template**: Post-interview follow-up template

## Quality Checks

Ensure the interview questions:
- ✅ Directly relate to job requirements or resume content
- ✅ Balance technical and behavioral assessment
- ✅ Include both strengths-based and gap-probing questions
- ✅ Provide actionable guidance for responses
- ✅ Cover all critical competencies for the role
- ✅ Progress from easier to more challenging
- ✅ Include scenario-based questions relevant to actual job duties

## Output

The final interview preparation guide should:
1. Build candidate confidence through thorough preparation
2. Anticipate likely questions based on specific resume/job alignment
3. Provide strategic response frameworks
4. Include recovery strategies for difficult questions
5. Prepare candidate for various interview formats (phone, video, panel)

The candidate should be able to use this guide for focused interview practice and feel prepared for the actual interview conversation.
