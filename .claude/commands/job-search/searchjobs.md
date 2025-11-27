---
description: Search hiring.cafe for job postings based on keywords, location, company, and other criteria with complete verbatim job descriptions
argument-hint: <search-query> [location] [--company=name] [--save] [--limit=N]
---

You are the Job Search Operations Coordinator orchestrating a hybrid search workflow that combines stealth browser automation to deliver complete, verbatim job descriptions from hiring.cafe.

## Important: Stealth Browser Required

This command uses the **stealth-browser MCP** which includes anti-bot detection measures. The stealth browser bypasses Vercel's bot protection but **may still be blocked by IP-based restrictions** when running from cloud environments (AWS, Azure, GCP, etc.).

**For best results:** Run this command from a local machine with a residential IP address.

**If blocked:** You'll see messages like "Too many requests" or "Please disable VPN/proxy". In that case:
1. Try running from a different network
2. Use the free public APIs instead (The Muse, Arbeitnow, Remotive, Jobicy)
3. Manually browse hiring.cafe and save job postings to `Job_Postings/`

## Command Overview

This command searches hiring.cafe for job postings using a stealth browser approach:
1. **Phase 1**: Navigate to hiring.cafe using stealth browser (bypasses bot detection)
2. **Phase 2**: Execute API search from within browser context (inherits session)
3. **Phase 3**: Extract and format job listings with complete descriptions

Users can search by:
- Keywords (job title, skills, technologies)
- Location (city, state/province, country)
- Company name
- Multiple criteria combined

## Arguments

- **{{ARG1}}** (required): Primary search query - keywords, job title, or company name
- **{{ARG2}}** (optional): Location filter - city, region, or country
- **{{ARG3}}** (optional): Additional flags:
  - `--company=NAME`: Filter by specific company
  - `--save`: Save results to file in Job_Postings/ folder
  - `--limit=N`: Limit results to N jobs (default: 20, max recommended: 50)

## Execution Protocol

### Phase 1: Parse Search Criteria

Extract and clarify the user's search requirements:
- Primary keywords from {{ARG1}}
- Location from {{ARG2}} if provided
- Company filter if `--company` flag present
- Output preferences if `--save` flag present
- Result limit from `--limit` flag (default: 20)

### Phase 2: Initialize Stealth Browser and Navigate

Use the stealth-browser MCP tools to bypass bot detection:

```
1. Call mcp__stealth-browser__stealth_navigate with:
   - url: "https://hiring.cafe"
   - waitFor: "networkidle"
   - timeout: 30000

2. Call mcp__stealth-browser__stealth_wait with:
   - milliseconds: 2000
   (Allow page to fully load and establish session)
```

If the navigation returns a "Security Checkpoint" or error, inform the user that bot detection is active and suggest alternatives.

### Phase 3: Execute API Search via Browser Context

Use the stealth browser to make the API call (inherits cookies/session):

```
Call mcp__stealth-browser__stealth_fetch_api with:
- url: "https://hiring.cafe/api/search-jobs"
- method: "POST"
- body: JSON.stringify({
    size: [limit from --limit or 20],
    page: 0,
    searchState: {
      searchQuery: "{{ARG1}} {{ARG2}}",
      sortBy: "date"
    }
  })
```

Parse the response to extract job listings.

### Phase 4: Handle Response

**If successful (status 200):**
- Extract job data: title, company_name, description, apply_url, location, category, etc.
- Proceed to Phase 5

**If blocked (status 403 or error message):**
```
⚠️ hiring.cafe is blocking this request due to IP-based restrictions.

This typically happens when running from cloud environments (codespaces, CI/CD, etc.)

**Alternatives:**
1. Run this command from your local machine with a residential IP
2. Search hiring.cafe manually in your browser
3. Use free public job APIs instead:
   - The Muse: https://www.themuse.com/api/public/jobs
   - Arbeitnow: https://arbeitnow.com/api/job-board-api
   - Remotive: https://remotive.com/api/remote-jobs
   - Jobicy: https://jobicy.com/api/v2/remote-jobs
```

### Phase 5: Format and Present Results

For each job in the results:

```markdown
## Job Search Results for "{{ARG1}}"

**Search Criteria:**
- Keywords: {{ARG1}}
- Location: {{ARG2}}
- Total Results Found: [count]

---

### 1. [Job Title]
**Company:** [Company Name]
**Location:** [City, State/Province, Country]
**Category:** [Job Category]
**Seniority:** [Experience Level]
**Commitment:** [Full Time/Part Time/Contract]
**Salary:** [If available]
**Apply URL:** [URL]

**Job Description:**
[Complete job description from API response]

---

### 2. [Repeat for each result...]
```

### Phase 6: Fetch Complete Descriptions (if needed)

If the API response contains truncated descriptions, navigate to individual job pages:

```
For each job needing full description:
1. Call mcp__stealth-browser__stealth_navigate with job apply_url
2. Call mcp__stealth-browser__stealth_wait with milliseconds: 2000
3. Call mcp__stealth-browser__stealth_get_content to extract job text
```

### Phase 7: Save Results (if --save flag present)

When the `--save` flag is provided, save the complete search results to:
`Job_Postings/SearchResults_[SanitizedQuery]_[Date].md`

Prepend the file with YAML frontmatter:

```yaml
---
query: "<search query from ARG1>"
location: "<location from ARG2, or leave blank if none>"
company_filter: "<company value from --company flag, or leave blank>"
result_count: <number of jobs saved>
generated_by: /searchjobs
generated_on: <ISO8601 timestamp>
output_type: job_search_results
status: final
version: 1.0
---
```

### Phase 8: Cleanup

Always close the stealth browser when done:

```
Call mcp__stealth-browser__stealth_close
```

## Usage Examples

### Example 1: Basic Keyword Search
```
/searchjobs "software engineer"
```
Searches for software engineer positions globally.

### Example 2: Location-Specific Search
```
/searchjobs "data analyst" "Toronto, Ontario"
```
Searches for data analyst jobs in Toronto area.

### Example 3: Company-Specific Search
```
/searchjobs "real estate analyst" --company=Purolator
```
Searches for real estate analyst positions at Purolator specifically.

### Example 4: Multi-Criteria with Save
```
/searchjobs "python developer" "Mississauga" --save --limit=30
```
Searches for Python developer jobs in Mississauga, saves to file.

## Stealth Browser MCP Tools Reference

The following tools are available from the stealth-browser MCP:

| Tool | Description |
|------|-------------|
| `stealth_navigate` | Navigate to URL with bot detection bypass |
| `stealth_get_content` | Get text content of current page |
| `stealth_get_html` | Get HTML content of current page |
| `stealth_evaluate` | Execute JavaScript in page context |
| `stealth_fetch_api` | Make fetch request with browser cookies/session |
| `stealth_wait` | Wait for time or element |
| `stealth_screenshot` | Take screenshot of current page |
| `stealth_close` | Close the browser |

## Error Handling

### Bot Detection Active
```
The stealth browser bypassed Vercel security, but hiring.cafe detected
a cloud/datacenter IP and is blocking the request.

Solutions:
1. Run from a local machine with residential IP
2. Use alternative job APIs (see above)
3. Manually browse and save job postings
```

### No Results Found
```
No jobs found matching "[query]" [in location].

Suggestions:
- Try broader keywords
- Remove location filter
- Check spelling of company names
```

### API Errors
```
Error connecting to hiring.cafe API.

Troubleshooting:
- Check internet connection
- Try again in a few moments
- Use alternative search methods
```

## Technical Notes

**Why Stealth Browser?**
- hiring.cafe uses Vercel's bot protection
- Standard Playwright/Puppeteer is detected and blocked
- The stealth plugin patches 10+ detection vectors:
  - navigator.webdriver property
  - HeadlessChrome user agent
  - WebGL fingerprinting
  - Canvas fingerprinting
  - Chrome runtime properties

**Remaining Limitation:**
- IP-based blocking cannot be bypassed by stealth plugins
- Cloud provider IPs (Azure, AWS, GCP) are often blocked
- Residential IPs typically work without issues

**Dependencies:**
- playwright-extra
- puppeteer-extra-plugin-stealth
- @modelcontextprotocol/sdk

Execute this job search operation and deliver results that help the user find their next opportunity.
