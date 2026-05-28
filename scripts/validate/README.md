# Validation Scripts

This directory contains quality validation scripts for JobOps.

## Available Validators

### validate-job-postings.js

Validates job postings for YAML front matter and filename conventions.

**Usage:**
```bash
# Validate all job postings in Job_Postings/
node scripts/validate/validate-job-postings.js

# Validate specific directory
node scripts/validate/validate-job-postings.js Job_Postings/

# Validate single file
node scripts/validate/validate-job-postings.js Job_Postings/UofT_Exec_Director_AM.md
```

**Checks:**
- ✅ YAML front matter presence
- ✅ Required fields (company, role, location, posting_date, source_url, date_saved, status)
- ✅ Recommended fields (salary_range, employment_type, remote_policy, etc.)
- ✅ Filename convention (Company_Role_YYYY-MM-DD.md)
- ✅ Date format validation (YYYY-MM-DD)
- ✅ Status value validation (active, archived, applied, rejected, interview, offer)

**Exit codes:**
- 0: All files valid
- 1: One or more files have errors

### validate-legacy-profile-block.js

Validates that `/jobops:assessjob` and `/jobops:assesscandidate` block deprecated `candidate_profile.json` artifacts before reading candidate source materials.

**Usage:**
```bash
node scripts/validate/validate-legacy-profile-block.js
```

**Checks:**
- Legacy profile check is present before candidate source reads
- `{source_path}/.profile/candidate_profile.json` is blocked
- `{source_path}/candidate_profile.json` is blocked
- Single-file `candidate_profile.json` inputs are blocked
- The command instructs the candidate to delete the deprecated artifact and offers to delete it

**Exit codes:**
- 0: Both assessment skills include the required guard
- 1: One or more guard requirements are missing

## Adding New Validators

Future validators to implement:
- `validate-assessments.js` - Check assessment completeness and evidence citations
- `validate-docs.js` - Check for broken links in documentation
- `validate-rubrics.js` - Verify rubric structure and scoring breakdowns

See `/workspaces/resumeoptimizer/SourceMaterial/Recommendation_Plan_2025-11-16.md` for planned quality gates.
