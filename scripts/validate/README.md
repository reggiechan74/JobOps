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

## Adding New Validators

Future validators to implement:
- `validate-assessments.js` - Check assessment completeness and evidence citations
- `validate-docs.js` - Check for broken links in documentation
- `validate-rubrics.js` - Verify rubric structure and scoring breakdowns

See `/workspaces/resumeoptimizer/SourceMaterial/Recommendation_Plan_2025-11-16.md` for planned quality gates.
