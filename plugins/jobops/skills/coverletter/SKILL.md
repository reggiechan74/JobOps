---
description: Generate a strategic cover letter with requirements-matching table from Step 3 resume
disable-model-invocation: true
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if this skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if this skill has jurisdiction-sensitive logic (crisis/legal skills accept `--jurisdiction=<ISO-3166-2>` to override).

## Templates

For each template used by this skill, resolve the full path as:

  {config.templates.base_dir}/{config.templates.active.<template_name>}/<filename>

Templates referenced by this skill: candidate_profile_schema

## Application Path Resolution

This skill writes to a per-application folder. Before writing any output:

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the job-posting filename, or honor `--app=<slug>` if supplied.
2. Compose the app folder: `{config.directories.applications_root}/{app_slug}/`.
3. Resolve this skill's sub-folder by category:
   - resume-development (buildresume, provenance-check) → `resume/`
   - cover-letter (coverletter) → `cover-letter/`
   - rubric / assessment (createrubric, assessjob, assesscandidate, auditjobposting) → `assessment/`
   - briefing / interview prep (briefing, interviewprep) → `interview/`
4. If the app folder does not exist, `mkdir -p` it, then copy
   `{config.directories.job_postings}/{filename}` → `{app_slug}/job_posting.md`
   so the pinned JD cannot silently change under completed work.
5. Exact-slug collisions (same Company+Role+Date) are not auto-suffixed. If the folder
   already contains the same output type, require the user to pass `--app=<distinct-slug>`.

## Arguments

- `$1`: Step 3 final resume file path (required)
- `$2`: Job description file path (required)
- `$3`: Hiring manager name (optional, defaults to "Hiring Manager")

Produces a compelling cover letter based on the validated Step 3 resume, featuring a strategic requirements-matching table that directly demonstrates fit for the role. Process:

1. Load and analyze the validated Step 3 resume.
2. Extract critical requirements from the job description.
3. Create a requirements-matching table with evidence.
4. Generate a compelling narrative with cultural alignment.
5. Produce an interview-ready cover letter.

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
`{applications_root}/{app_slug}/cover-letter/cover_letter.md`

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
- Demonstrate immediate value to the employer
- Provide visual proof of requirements match
- Maintain credibility chain from master documents
- Create compelling case for interview invitation
- Stand out from generic applications

## Next Steps

After generation, you'll have:
1. A compelling, evidence-based cover letter
2. A requirements-matching table for quick scanning
3. Consistent messaging across all application materials
4. Strong positioning for interview success

The cover letter and resume work together as a cohesive application package, with every claim defensible and every achievement traceable to your comprehensive work history.
