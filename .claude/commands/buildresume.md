---
description: Build a complete resume through all 3 steps (draft, provenance check, final)
argument-hint: <job-description-file> [cultural-profile]
---

I'll create a complete resume by running all three steps of the resume optimization process sequentially.

**Arguments:**
- `$1`: Job description file path (required)
- `$2`: Cultural profile (optional, defaults to "Canadian")

**Process:**
1. **Step 1**: Create initial tailored resume draft using HAM-Z methodology
2. **Step 2**: Perform comprehensive provenance analysis for credibility
3. **Step 3**: Create final hardened resume addressing all issues

### YAML front matter requirement

Every markdown file generated during this command must begin with a YAML front matter block so downstream tooling can parse metadata without heuristics. Populate the fields with real values before writing any body content.

- **Step 1 draft** (`OutputResumes/Step1_Draft_*`):
  ```yaml
  ---
  job_file: {{ARG1}}
  role: <target role title>
  company: <hiring company>
  candidate: <full candidate name>
  generated_by: /buildresume step1-resume-draft
  generated_on: <ISO8601 timestamp>
  output_type: resume_step1
  status: draft
  version: 1.0
  ---
  ```
- **Step 2 provenance analysis** (`OutputResumes/Step2_Provenance_Analysis_*`): use the same fields but set `generated_by` to `/buildresume step2-provenance-check`, `output_type` to `resume_provenance`, and `status` to `analysis`.
- **Step 3 final resume** (`OutputResumes/Step3_Final_Resume_*`): set `generated_by` to `/buildresume step3-final-resume`, `output_type` to `resume_final`, and `status` to `final`. Increment `version` if you regenerate after revisions.

Always include the front matter before any markdown headings or narrative.

Let me run the complete 3-step resume building process:

## Step 1: Creating Initial Resume Draft

I'll start by reading the job description file and creating the initial draft using the $2 cultural profile.

First, let me read the job description:

@$1

Now I'll run the step1-resume-draft agent with the specified cultural profile:

**Running Step 1 Agent...**

I'm launching the step1-resume-draft agent to create an initial tailored resume draft based on the job requirements. This agent will:
- Analyze the job requirements from $1
- Apply the $2 cultural profile preferences
- Use the HAM-Z methodology for strategic positioning
- Create a targeted first draft optimized for the role

## Step 2: Provenance Analysis

**Running Step 2 Agent...**

Now I'll launch the step2-provenance-check agent to analyze the Step 1 draft for:
- Credibility issues and evidence gaps
- Risk factors for each major claim
- Defensibility of quantified achievements
- Consistency across all experience sections
- Recommendations for strengthening credibility

## Step 3: Final Hardened Resume

**Running Step 3 Agent...**

Finally, I'll launch the step3-final-resume agent to:
- Incorporate all recommendations from Step 2 analysis
- Address credibility concerns while maintaining competitive positioning
- Produce a fully defensible final resume
- Ensure maximum interview potential

## Summary

All three steps of the resume optimization process will be completed:
- ✅ Step 1: Initial targeted draft created
- ✅ Step 2: Comprehensive provenance analysis performed
- ✅ Step 3: Final hardened resume produced

Your complete, optimized resume will be ready for application with full credibility and competitive positioning.
