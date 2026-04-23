---
description: Build a complete resume through all 3 steps (draft, provenance check, final)
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

Templates referenced by this skill: candidate_profile_schema, evidence_verification_framework

## Step 1: Creating Initial Resume Draft
✓ Initiating strike package assembly

I'll start by reading the job description file and creating the initial draft using the $2 cultural profile.

First, let me read the job description:

@$1

Now I'll run the step1-resume-draft agent with the specified cultural profile:

**✓ Deploying Step 1 Agent - Initial Draft Creation**

I'm launching the step1-resume-draft agent to create an initial tailored resume draft based on the job requirements. This agent will:
- Analyze the job requirements from $1
- Apply the $2 cultural profile preferences
- Use the HAM-Z methodology for strategic positioning
- **Review ALL education and credentials from master resume against job requirements**
- **Explicitly justify any credential exclusions in agent output**
- Create a targeted first draft optimized for the role

## Step 2: Provenance Analysis
✓ Executing credibility verification sweep

**✓ Deploying Step 2 Agent - Provenance Verification**

Now I'll launch the step2-provenance-check agent to analyze the Step 1 draft for:
- Credibility issues and evidence gaps
- Risk factors for each major claim
- Defensibility of quantified achievements
- Consistency across all experience sections
- **Missing credentials: Verify ALL education/designations from master resume are evaluated**
- **Credential-to-job mapping: Flag any excluded credentials that map to job requirements**
- Recommendations for strengthening credibility

## Step 3: Final Hardened Resume
✓ Producing deployment-ready final resume

**✓ Deploying Step 3 Agent - Hardened Resume Production**

Finally, I'll launch the step3-final-resume agent to:
- Incorporate all recommendations from Step 2 analysis
- Address credibility concerns while maintaining competitive positioning
- Produce a fully defensible final resume
- Ensure maximum interview potential

## Mission Summary

All three steps of the resume assembly process will be completed:
- Step 1: Initial targeted draft created
- Step 2: Comprehensive provenance verification performed
- Step 3: Final hardened resume produced

Your deployment-ready resume will be ready for mission execution with full credibility and competitive positioning.
