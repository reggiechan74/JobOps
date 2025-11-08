---
description: Generate a strategic cover letter with requirements-matching table from Step 3 resume
argument-hint: <step3-resume-file> <job-description-file> [hiring-manager-name]
---

I'll create a compelling cover letter based on your final Step 3 resume, featuring a strategic requirements-matching table that directly demonstrates your fit for the role.

**Arguments:**
- `$1`: Step 3 final resume file path (required)
- `$2`: Job description file path (required)
- `$3`: Hiring manager name (optional, defaults to "Hiring Manager")

**Process:**
1. Load and analyze the validated Step 3 resume
2. Extract critical requirements from job description
3. Create requirements-matching table with evidence
4. Generate compelling narrative with cultural alignment
5. Produce interview-ready cover letter

## Step 1: Loading Final Resume

First, let me read your Step 3 final resume:

@$1

## Step 2: Analyzing Job Requirements

Now loading the job description to extract critical requirements:

@$2

## Step 3: Generating Cover Letter

**Running Step 4 Agent...**

I'm launching the step4-cover-letter agent to create your cover letter. This agent will:

### Cover Letter Components:
- **Opening Hook**: Compelling introduction tailored to the specific role
- **Requirements Alignment Table**: Visual proof of your perfect fit
- **Value Proposition**: Narrative expanding on your strongest qualifications
- **Cultural Fit**: Demonstration of alignment with company mission
- **Strong Close**: Clear call-to-action for next steps

### Requirements-Matching Table:
The cover letter will feature a two-column table that:
- Maps their top 5-6 requirements to your specific achievements
- Uses quantified evidence from your validated resume
- Provides scannable proof of fit for busy hiring managers
- Creates natural interview talking points

### Quality Assurance:
- All claims traced to Step 3 validated content
- No hallucinations or unsupported assertions
- Company and role names verified for accuracy
- Cultural tone matches resume style
- Optimized for ATS and human review

## Output

Your strategic cover letter will be saved as:
`OutputResumes/Step4_CoverLetter_[Role]_[Company]_[Date].md`

Before the letter content, write this YAML metadata block with actual values:

```yaml
---
job_file: $2
role: <role title>
company: <company name>
candidate: <full candidate name>
generated_by: /coverletter
generated_on: <ISO8601 timestamp>
output_type: cover_letter
status: final
version: 1.0
---
```

Update timestamps and increment `version` on subsequent iterations.

The cover letter will:
- ✅ Demonstrate immediate value to the employer
- ✅ Provide visual proof of requirements match
- ✅ Maintain credibility chain from master documents
- ✅ Create compelling case for interview invitation
- ✅ Stand out from generic applications

## Next Steps

After generation, you'll have:
1. A compelling, evidence-based cover letter
2. A requirements-matching table for quick scanning
3. Consistent messaging across all application materials
4. Strong positioning for interview success

The cover letter and resume work together as a cohesive application package, with every claim defensible and every achievement traceable to your comprehensive work history.
