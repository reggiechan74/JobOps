---
description: Generate a comprehensive briefing note to address skill gaps or prepare for interviews
signature: "<assessment-report> <job-description> [gaps-only]"
project: true
---

You are creating a strategic briefing note to help a candidate prepare for their interview or address identified skill gaps. This command generates a detailed study guide based on the assessment report and job requirements.

## Your Task

Create a comprehensive briefing note that either:
1. **Gaps-only mode**: Focuses exclusively on addressing identified weaknesses and skill gaps
2. **Full preparation mode**: Covers all job requirements for complete interview preparation

## Arguments

- `{{ARG1}}`: Assessment report file (required) - The candidate assessment output from /assessjob
- `{{ARG2}}`: Job description file (required) - The original job posting
- `{{ARG3}}`: Mode flag (optional) - "gaps-only" for gap-focused briefing, omit for full preparation

## Step-by-Step Process

### 1. Load Required Documents
- Read the assessment report from `{{ARG1}}` (check OutputResumes/ or Scoring_Rubrics/ directories)
- Read the job description from `Job_Postings/{{ARG2}}` (add .md extension if needed)
- Determine mode: gaps-only if {{ARG3}} == "gaps-only", otherwise full preparation

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

Use WebSearch and WebFetch to research:
- Current best practices and methodologies
- Industry standards and trends
- Technical concepts and implementations
- Certification requirements and preparation materials
- Real-world application examples

### 4. Generate Briefing Note Structure

Create a comprehensive briefing document with the following sections:

#### A. Executive Summary
- Overall preparedness assessment
- Critical gaps requiring immediate attention
- Recommended preparation timeline
- Key focus areas for interview success

#### B. Skill Gap Analysis (or Full Requirements Overview)
For each gap/requirement, provide:
- **Current State**: Candidate's current level
- **Target State**: What the role requires
- **Gap Severity**: Critical/High/Medium/Low
- **Time to Bridge**: Estimated hours/days needed

#### C. Detailed Study Guide
For each topic area:

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

#### D. Action Plan
- **24 Hours Before Interview**:
  - Critical topics to review
  - Key talking points to prepare
  - Questions to ask the interviewer

- **3 Days Before Interview**:
  - Hands-on practice priorities
  - Mock interview topics
  - Portfolio pieces to prepare

- **1 Week Before Interview**:
  - Deeper learning objectives
  - Projects to complete
  - Skills to demonstrate

#### E. Interview Strategy
- How to acknowledge gaps professionally
- Demonstrating learning ability and growth mindset
- Pivoting to related strengths
- Using transferable skills effectively

#### F. Quick Reference Guide
- One-page summary of key concepts
- Formula sheet or command reference
- Common acronyms and terminology
- Industry-specific metrics and KPIs

### 5. Customize Based on Mode

**If gaps-only mode:**
- Focus exclusively on addressing weaknesses
- Prioritize critical gaps that could be deal-breakers
- Provide accelerated learning paths
- Include mitigation strategies for unchangeable gaps

**If full preparation mode:**
- Cover all job requirements comprehensively
- Balance between strengthening weaknesses and showcasing strengths
- Include competitive differentiation strategies
- Prepare for technical deep-dives in strength areas

### 6. Format and Save

Save the briefing note as:
- Filename: `Briefing_[Company]_[Role]_[Mode]_[Date].md`
- Location: `Briefing_Notes/` directory
- Format: Markdown with clear headings and sections
- Include table of contents for easy navigation

## Quality Checks

Ensure the briefing note:
- ✅ Addresses every identified gap or requirement
- ✅ Provides actionable, specific learning resources
- ✅ Includes realistic timelines for skill development
- ✅ Offers practical exercises and hands-on practice
- ✅ Prepares for likely interview scenarios
- ✅ Balances depth with time constraints
- ✅ Uses current, relevant industry information

## Output

The final briefing note should be a comprehensive, actionable study guide that:
1. Clearly identifies what needs to be learned
2. Provides structured learning paths
3. Includes practical exercises and examples
4. Prepares for specific interview scenarios
5. Offers quick-reference materials for last-minute review

The candidate should be able to use this briefing note as their primary preparation resource for the interview.