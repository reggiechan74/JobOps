# Hiring.Cafe Job Search Guide

This guide explains how to use the integrated hiring.cafe job search functionality in the ResumeOptimizer system.

## Overview

The ResumeOptimizer system now includes a specialized agent (`hiringcafe-search`) that searches the hiring.cafe job board using their API. This provides structured access to job postings with advanced filtering capabilities.

## Quick Start

### Basic Search
```bash
/searchjobs "software engineer"
```

### Location-Specific Search
```bash
/searchjobs "data analyst" "Toronto, Ontario"
```

### Company-Specific Search
```bash
/searchjobs "real estate analyst" --company=Purolator
```

### Save Results to File
```bash
/searchjobs "python developer" "Mississauga" --save
```

## Command Syntax

```bash
/searchjobs "<search-query>" [location] [--company=name] [--save] [--limit=N]
```

### Parameters

- **search-query** (required): Keywords, job title, skills, or company name
- **location** (optional): City, state/province, or country
- **--company=NAME** (optional): Filter by specific company
- **--save** (optional): Save results to Job_Postings/ folder
- **--limit=N** (optional): Limit results to N jobs (default: 20)

## Search Examples

### 1. Keyword Search
Find all positions matching specific keywords:
```bash
/searchjobs "machine learning engineer"
/searchjobs "product manager"
/searchjobs "senior developer remote"
```

### 2. Location-Based Search
Search within specific geographic areas:
```bash
/searchjobs "software developer" "Vancouver, BC"
/searchjobs "analyst" "Toronto"
/searchjobs "engineer" "Ontario, Canada"
```

### 3. Company-Focused Search
Find opportunities at specific companies:
```bash
/searchjobs "Shopify"
/searchjobs "RBC" "Toronto"
/searchjobs "" --company=Scotiabank
```

### 4. Multi-Criteria Search
Combine multiple search parameters:
```bash
/searchjobs "data scientist" "Montreal" --save
/searchjobs "DevOps engineer" "remote" --limit=50
/searchjobs "financial analyst" "Toronto" --company=CIBC --save
```

### 5. Broad Exploration
Cast a wide net to see what's available:
```bash
/searchjobs "remote" --limit=100
/searchjobs "Toronto" --limit=50
/searchjobs "entry level" "Ontario"
```

## Understanding Results

### Standard Output Format
```
Job Title: [Position Title]
Company: [Company Name]
Location: [City, State/Province, Country]
Category: [Job Category]
Seniority: [Entry Level / Mid Level / Senior Level]
Commitment: [Full Time / Part Time / Contract / Internship]
Apply URL: [Direct Application Link]
---
```

### Result Summary
Each search provides:
- **Total jobs found**: Complete count of matching positions
- **Top companies**: Most frequent employers in results
- **Geographic distribution**: Where jobs are located
- **Category breakdown**: Types of roles available

## Saved Results

When using `--save`, results are saved to:
```
Job_Postings/SearchResults_[Query]_[Date].md
```

Saved files include:
- Complete search criteria
- All matching job listings
- Full job descriptions (first 500 chars)
- Apply URLs for easy reference

## API Technical Details

### How It Works

The agent uses the hiring.cafe REST API:
- **Endpoint**: `https://hiring.cafe/api/search-jobs`
- **Method**: POST
- **Format**: JSON

### Search Strategy

1. **Broad keyword matching**: Searches job titles, descriptions, and company names
2. **Location filtering**: Checks multiple location fields for matches
3. **Post-processing**: Uses jq to refine and filter results
4. **Structured output**: Returns organized, easy-to-read listings

### Search Tips

**DO:**
- ✅ Use specific job titles when possible
- ✅ Include location in the search query
- ✅ Try company names for company-specific searches
- ✅ Start broad, then refine

**DON'T:**
- ❌ Use overly complex boolean queries
- ❌ Rely solely on location parameter (include in keywords too)
- ❌ Assume exact string matching (search is fuzzy)

## Integration with Resume Workflow

### Typical Workflow

1. **Search for Jobs**
   ```bash
   /searchjobs "software engineer" "Toronto" --save
   ```

2. **Review Saved Results**
   Check `Job_Postings/SearchResults_*.md`

3. **Select Target Job**
   Copy job posting to new file: `Job_Postings/CompanyName_Role_Date.md`

4. **Build Resume**
   ```bash
   /buildresume Job_Postings/CompanyName_Role_Date.md
   ```

5. **Research Company** (Optional)
   ```bash
   /osint CompanyName
   ```

6. **Assess Fit**
   ```bash
   /assessjob Job_Postings/CompanyName_Role_Date.md
   ```

7. **Prepare for Interview**
   ```bash
   /briefing OutputResumes/Assessment_* Job_Postings/CompanyName_Role_Date.md
   /interviewprep OutputResumes/Step3_* Job_Postings/CompanyName_Role_Date.md
   ```

## Advanced Usage

### Filtering by Multiple Criteria

Search broadly, then use grep or jq to filter:
```bash
# Save all Toronto tech jobs
/searchjobs "software developer data engineer DevOps" "Toronto" --save --limit=100

# Manually filter saved results file for specific criteria
```

### Regular Job Monitoring

Set up regular searches for your target roles:
```bash
# Weekly search for your ideal role
/searchjobs "senior data scientist" "Toronto" --save --limit=50
```

### Market Research

Use broad searches to understand the job market:
```bash
# See what's available in your city
/searchjobs "Toronto" --limit=100 --save

# Understand salary landscape
/searchjobs "software engineer salary 150k" "Toronto"
```

## Troubleshooting

### No Results Found
**Problem**: Search returns 0 results

**Solutions**:
- Use broader keywords (e.g., "developer" instead of "senior staff software engineer III")
- Remove location filter or try nearby cities
- Check spelling of company names
- Try related job titles or skills

### Too Many Results
**Problem**: Overwhelming number of results

**Solutions**:
- Add location filter: `/searchjobs "query" "city"`
- Use company filter: `/searchjobs "query" --company=name`
- Be more specific: `/searchjobs "senior python developer"`
- Increase limit to see more: `/searchjobs "query" --limit=100`

### Results Don't Match Criteria
**Problem**: Getting irrelevant results

**Solutions**:
- Be more specific in search query
- Include multiple related keywords
- Use company filter for company-specific searches
- Check saved results and manually filter

## API Limitations

### Known Limitations
- Location filtering via API parameter is inconsistent
- Maximum 100 results per page
- No salary range filtering in API
- Remote/hybrid/onsite filters not directly supported

### Workarounds
- Include location in search keywords
- Use pagination for large result sets
- Filter results with jq post-retrieval
- Include "remote" or "hybrid" in search query

## Support & Resources

### Getting Help
- Check CLAUDE.md for complete command documentation
- Review example searches in this guide
- Ask Claude for specific search strategies

### Related Commands
- `/osint <company>` - Research companies from search results
- `/buildresume <job-file>` - Create targeted resume for job
- `/assessjob <job-file>` - Evaluate your fit for the role

## Why hiring.cafe?

**Advantages**:
- Aggregates jobs from 30,000+ company career pages
- AI-enhanced job descriptions with extracted metadata
- Refreshed 3x daily for current listings
- No ATS (Applicant Tracking System) clutter
- Direct apply links to company career pages

**Compared to Indeed/LinkedIn**:
- More focused on company career pages
- Better data extraction and categorization
- Fewer duplicate listings
- Cleaner, more structured data

## Examples from Real Searches

### Example 1: Finding the Purolator Job
```bash
# Search specifically for real estate analyst roles
/searchjobs "Analyst Real Estate" "Mississauga"

# Result found:
Job Title: Analyst Real Estate (Mississauga, ON, CA, L5N 0E1)
Company: Purolator
Apply URL: https://careers.purolator.com/job/Mississauga-Analyst-Real-Estate-ON-L5N-0E1/1328182700/
```

### Example 2: Toronto Financial Services
```bash
/searchjobs "financial analyst" "Toronto" --limit=30

# Returns positions from RBC, Scotiabank, TD, BMO, CIBC, etc.
```

### Example 3: Remote Tech Jobs
```bash
/searchjobs "remote software engineer" --limit=50 --save

# Saves 50 remote software engineering positions to file for review
```

---

**Last Updated**: 2025-09-29
**Version**: 1.0