---
description: Initialize JobOps workspace - configure directories, install templates, and optionally import career history
disable-model-invocation: true
---

# JobOps Workspace Setup

## Step 1: Welcome

Welcome to JobOps setup. This skill initializes your workspace with everything needed to run the JobOps intelligence-driven job application system.

Setup will walk you through:
- **Directory configuration** - establish the folder structure for resumes, job postings, assessments, and output
- **Template installation** - install scoring rubrics, evidence verification, and assessment report templates
- **Career history import** (optional) - parse existing resume files into the HAM-Z master inventory format
- **Profile selection** (optional) - choose a cultural profile that shapes resume tone and language

This is a one-time process. You can re-run it later to reconfigure.

---

## Step 2: Check for Existing Configuration

Check whether a configuration file already exists:

```bash
if [ -f ".jobops/config.json" ]; then
  echo "EXISTING_CONFIG_FOUND"
  cat .jobops/config.json
else
  echo "NO_CONFIG"
fi
```

**If `.jobops/config.json` exists**, present the user with two choices:

1. **Reconfigure** - Start fresh. The existing config will be overwritten and directories will be re-created as needed.
2. **Skip setup** - The workspace is already configured. Exit setup and show the current configuration summary.

If the user chooses to skip, jump directly to **Step 10: Summary** using the existing config values. Otherwise, continue with Step 3.

**If no config exists**, continue to Step 3.

---

## Step 3: Directory Configuration

Present the full directory layout to the user with default values. Ask the user to confirm or customize any paths before proceeding.

| Directory Key | Default Path | Purpose |
|---|---|---|
| `resume_source` | `./ResumeSourceFolder` | Master resume data (Experience/, CareerHighlights/, Technology/, Preferences/, .profile/) |
| `job_postings` | `./Job_Postings` | Target job descriptions as `.md` files |
| `output_resumes` | `./OutputResumes` | Generated drafts, analyses, and final resumes |
| `scoring_rubrics` | `./Scoring_Rubrics` | Assessment rubrics (`Rubric_[Company]_[Role]_[Date].md`) |
| `briefing_notes` | `./Briefing_Notes` | Interview prep and study guides |
| `intelligence_reports` | `./Intelligence_Reports` | OSINT and company intelligence reports |
| `sample_output` | `./Sample_Output` | Reference examples and sample documents |

Show all seven directories at once. The user may accept all defaults by confirming, or override individual paths. Store the confirmed values for use in later steps.

---

## Step 4: Create Directories

For each directory confirmed in Step 3, create it if it does not already exist:

```bash
mkdir -p "./ResumeSourceFolder"
mkdir -p "./Job_Postings"
mkdir -p "./OutputResumes"
mkdir -p "./Scoring_Rubrics"
mkdir -p "./Briefing_Notes"
mkdir -p "./Intelligence_Reports"
mkdir -p "./Sample_Output"
```

Use the actual paths the user confirmed. The `-p` flag ensures no errors if directories already exist and creates any intermediate parent directories.

Also create the JobOps internal config directory:

```bash
mkdir -p ".jobops"
```

Report which directories were newly created versus which already existed.

---

## Step 5: Template Installation

Install the default template set by running the template copy script:

```bash
bash "$(cat /tmp/.jobops-plugin-root)/scripts/copy-templates.sh" ".jobops/templates/default"
```

This copies the canonical templates into the workspace:
- `assessment_rubric_framework.md` - 200-point deduplicated scoring with role-based weight variants (Technical IC, People Manager, Executive)
- `evidence_verification_framework.md` - Citation requirements, domain verification, experience classification
- `assessment_report_structure.md` - Report format with 3-level evidence attribution
- `candidate_profile_schema.json` - JSON schema for optimized profiles (85-90% token reduction)

If the script fails or the plugin root file is missing, report the error and advise the user to check their plugin installation.

---

## Step 6: Create Custom Template Directory

Create the custom template directory for user overrides:

```bash
mkdir -p ".jobops/templates/custom"
```

Users can place modified templates here. When a custom version of a template exists, it takes precedence over the default. This is configured via the `templates.active` section of `config.json` (see Step 7).

---

## Step 7: Write config.json

Write the workspace configuration file at `.jobops/config.json` using the directory paths confirmed in Step 3 and default template settings.

The structure must follow this format (substitute the user's confirmed paths):

```json
{
  "version": "1.0",
  "directories": {
    "resume_source": "./ResumeSourceFolder",
    "job_postings": "./Job_Postings",
    "output_resumes": "./OutputResumes",
    "scoring_rubrics": "./Scoring_Rubrics",
    "briefing_notes": "./Briefing_Notes",
    "intelligence_reports": "./Intelligence_Reports",
    "sample_output": "./Sample_Output"
  },
  "templates": {
    "base_dir": "./.jobops/templates",
    "active": {
      "assessment_rubric_framework": "default",
      "evidence_verification_framework": "default",
      "assessment_report_structure": "default",
      "candidate_profile_schema": "default"
    }
  }
}
```

To switch a template to a custom version later, the user changes the value from `"default"` to `"custom"` for that template key.

---

## Step 8: Career History Import (Optional)

Ask the user: **Do you have existing resume files to import?**

If the user declines, skip to Step 9.

If the user has files, proceed with the career history import process:

### 8a. Collect Input Files

Accept one or more resume files in the following formats:
- PDF (`.pdf`)
- Word (`.docx`)
- Plain text (`.txt`)
- Markdown (`.md`)

Ask the user to provide the file paths.

### 8b. Parse and Extract

Read each provided file and extract structured career data:
- **Work experience** - company names, titles, dates, responsibilities, accomplishments
- **Education** - degrees, institutions, dates, certifications
- **Skills** - technical skills, tools, platforms, methodologies
- **Achievements** - quantified results, awards, recognitions, project outcomes

### 8c. Transform to HAM-Z Format

Convert extracted achievements and experience into HAM-Z format:

> **Achieved [Metric-Driven Result] by leveraging [Hard Skill] to [perform specific action/process]**

Each bullet point should follow this structure:
- **H** - Hard Skill: the specific technical or professional skill applied
- **A** - Action: the concrete action or process performed
- **M** - Metrics: quantified outcomes (percentages, dollar amounts, time savings, scale)
- **Z** - Structure: the organizational framing that ties it together

Flag items that lack metrics or hard skills for the user to review and enhance.

### 8d. Create ResumeSourceFolder Structure

Build out the master resume data directory:

```bash
mkdir -p "./ResumeSourceFolder/Experience"
mkdir -p "./ResumeSourceFolder/CareerHighlights"
mkdir -p "./ResumeSourceFolder/Technology"
mkdir -p "./ResumeSourceFolder/Preferences"
```

Use the user's confirmed `resume_source` path from Step 3 instead of the default if it was customized.

### 8e. Populate Files

Write the extracted and transformed data into the appropriate subdirectories:

- **Experience/** - One file per role or employer, with HAM-Z formatted bullets
- **CareerHighlights/** - Top achievements extracted and ranked by impact
- **Technology/** - Skills inventory organized by category (languages, platforms, tools, methodologies)
- **Preferences/** - Job search preferences and constraints

### 8f. Create Vision Template

Create a vision template if one does not already exist:

```bash
cat > "./ResumeSourceFolder/Preferences/Vision.md" << 'EOF'
# Career Vision

## Target Roles
<!-- List the roles you are pursuing -->

## Industry Preferences
<!-- Industries or sectors of interest -->

## Work Style
<!-- Remote, hybrid, on-site preferences -->

## Geographic Preferences
<!-- Locations, willingness to relocate -->

## Compensation Expectations
<!-- Salary range, equity preferences, benefits priorities -->

## Non-Negotiables
<!-- Hard requirements for your next role -->

## Growth Goals
<!-- Skills to develop, career trajectory -->
EOF
```

### 8g. Gap Analysis

After import, provide:
- Count of experience entries, highlights, and skills imported
- Bullets that could not be converted to HAM-Z (missing metrics or hard skills)
- Recommendations for strengthening weak entries
- Suggested next steps for completing the career inventory

### 8h. Profile Generation

If the import produced a substantial career history, suggest running the **resume-summarizer** agent to generate an optimized candidate profile:

> To generate a compressed candidate profile for faster assessments, use the `resume-summarizer` agent. This produces a JSON profile with 85-90% token reduction while preserving all key career data.

---

## Step 9: Profile Selection (Optional)

Ask the user for their preferred cultural profile:

### Canadian Profile (Default)
- Emphasizes collaborative language and team contribution
- Uses modest, measured tone ("contributed to" rather than "single-handedly drove")
- Highlights cross-functional collaboration and stakeholder engagement
- Frames achievements within team and organizational context

### American Profile
- Emphasizes individual achievement and assertive language
- Uses strong action verbs and direct impact statements ("drove", "spearheaded", "delivered")
- Highlights personal ownership and leadership initiative
- Frames achievements around individual contribution and competitive results

Store the selection in `.jobops/config.json` by adding a `profile` key:

```json
{
  "profile": {
    "cultural": "canadian"
  }
}
```

Valid values: `"canadian"` or `"american"`.

If the user skips this step, default to `"canadian"`.

---

## Step 10: Summary

Display a complete summary of what was configured:

### Workspace Configuration
- List each directory and whether it was newly created or already existed
- Show the path to `.jobops/config.json`

### Templates Installed
- List each template file installed in `.jobops/templates/default/`
- Note the custom template directory at `.jobops/templates/custom/`

### Career History (if imported)
- Number of experience entries created
- Number of career highlights extracted
- Number of skills catalogued
- Files that need review (incomplete HAM-Z conversions)

### Profile
- Cultural profile selected (Canadian or American)

### Next Steps

Recommend the following actions:

1. **Review and customize career history files** - Open `ResumeSourceFolder/` and refine the imported data, especially any entries flagged during gap analysis.
2. **Try `/jobops:buildresume`** - Point it at a job posting to create a targeted, tailored resume using the 3-step process (Draft, Provenance Check, Final).
3. **Try `/jobops:assessjob`** - Evaluate a job posting against your profile with a full 200-point scored assessment.
