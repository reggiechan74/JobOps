---
name: hiringcafe-search
description: "DEPRECATED: This agent is deprecated due to hiring.cafe's bot detection. Use the /searchjobs slash command instead, which uses the stealth-browser MCP to bypass Vercel's security. The /searchjobs command will only work reliably from a local machine with a residential IP address."
model: sonnet
tools: Bash, WebSearch
---

# DEPRECATED AGENT

**This agent has been deprecated** because hiring.cafe has implemented aggressive bot detection that blocks:

1. **Vercel Security Checkpoint** - Detects automated browsers
2. **IP-based blocking** - Blocks cloud/datacenter IPs (Azure, AWS, GCP)

## Use Instead

Use the **/searchjobs** slash command, which:
- Uses the **stealth-browser MCP** with playwright-extra and puppeteer-extra-plugin-stealth
- Bypasses Vercel's bot detection (browser fingerprinting)
- Provides clear error messages when IP-blocked

**Note:** Even with stealth browser, job searches will only work reliably from:
- Local machines with residential IP addresses
- Non-cloud environments

## Alternative Job Sources

If running from a cloud environment, use these free public APIs instead:

| API | Endpoint | Focus |
|-----|----------|-------|
| The Muse | `https://www.themuse.com/api/public/jobs` | US corporate jobs (477K+) |
| Arbeitnow | `https://arbeitnow.com/api/job-board-api` | EU/Global tech jobs |
| Remotive | `https://remotive.com/api/remote-jobs` | Remote jobs only |
| Jobicy | `https://jobicy.com/api/v2/remote-jobs` | Remote jobs only |

## Example API Calls

### The Muse (Largest Free Dataset)
```bash
curl -s 'https://www.themuse.com/api/public/jobs?page=1&per_page=20' | jq '.results[] | {title: .name, company: .company.name, location: .locations[0].name}'
```

### Arbeitnow
```bash
curl -s 'https://arbeitnow.com/api/job-board-api' | jq '.data[:10] | .[] | {title, company_name, location}'
```

### Remotive
```bash
curl -s 'https://remotive.com/api/remote-jobs?limit=10' | jq '.jobs[] | {title, company_name, candidate_required_location}'
```

### Jobicy
```bash
curl -s 'https://jobicy.com/api/v2/remote-jobs?count=10' | jq '.jobs[] | {title: .jobTitle, company: .companyName, location: .jobGeo}'
```

## Migration Path

To migrate existing workflows:

1. Replace `hiringcafe-search` agent calls with `/searchjobs` command
2. If running in cloud, switch to The Muse or Arbeitnow APIs
3. For hiring.cafe specifically, run locally with residential IP

---

*This agent documentation is preserved for historical reference.*
