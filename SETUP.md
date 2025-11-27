# JobOps Setup Instructions

## Overview

**JobOps is a Claude Code repository**. This means it runs entirely through Claude Code's slash commands and agents. The vast majority of functionality requires **only Claude Code** to be installed.

## Required Setup

### Install Claude Code CLI

This is the **only** required dependency for JobOps core functionality:

```bash
npm install -g @anthropic-ai/claude-code
```

More information at [claude.ai/code](https://claude.ai/code)

### Verify Installation

```bash
claude --version
```

### Start Using JobOps

Once Claude Code is installed, navigate to the repository and launch Claude:

```bash
cd /path/to/resumeoptimizer
claude
```

All slash commands are now available:
- `/create-career-history` - Parse existing resumes
- `/assessjob` - Evaluate job fit
- `/buildresume` - Create tailored resumes
- `/briefing` - Generate interview prep materials
- `/osint` - Company intelligence gathering
- And many more (see README.md)

## Optional Dependencies

### Pandoc (Document Conversion)

**Required for**: `/convert` command (markdown to Word DOCX)

**Installation**: Use the built-in installer:
```bash
/install-pandoc
```

Or install manually:
- **Ubuntu/Debian**: `sudo apt-get install pandoc`
- **macOS**: `brew install pandoc`
- **Windows**: Download from [pandoc.org](https://pandoc.org/installing.html)

### Stealth Browser MCP Server (Job Search)

**Required for**: `/searchjobs` command

**Status**: Works with residential IPs, blocked from cloud environments

The `/searchjobs` command uses a custom stealth browser MCP to bypass hiring.cafe's bot detection.

#### How It Works

JobOps includes a custom **stealth-browser MCP server** that:
- Uses `playwright-extra` with `puppeteer-extra-plugin-stealth`
- Bypasses Vercel's bot detection (browser fingerprinting)
- Patches 10+ detection vectors (webdriver property, HeadlessChrome, WebGL, etc.)

**Important Limitation**: Even with stealth, hiring.cafe still blocks cloud/datacenter IPs. The `/searchjobs` command works best from:
- Local machines with residential IP addresses
- Non-cloud/non-VPN networks

#### Installation (Already Included)

The stealth browser is pre-configured in this repository:

1. **Dependencies are in package.json:**
   ```bash
   npm install
   ```

2. **MCP is configured in .mcp.json:**
   ```json
   {
     "mcpServers": {
       "stealth-browser": {
         "type": "stdio",
         "command": "node",
         "args": ["scripts/mcp/stealth-browser.js"]
       }
     }
   }
   ```

3. **Restart Claude Code** to load the MCP server

4. **Verify connection:**
   ```bash
   claude mcp list
   ```
   You should see `stealth-browser` listed.

#### Available Stealth Browser Tools

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

#### Troubleshooting

**"Too many requests" or "Please disable VPN/proxy" error:**
- This is IP-based blocking, not bot detection
- Run from a local machine with residential IP
- Use alternative job APIs instead (The Muse, Arbeitnow, etc.)

**Stealth browser not connecting:**
```bash
# Verify dependencies
npm install

# Test the MCP server directly
node scripts/mcp/stealth-browser.js

# Restart Claude Code
claude mcp list
```

**Browser not found:**
```bash
npx playwright install chromium
```

### Alternative: Standard Playwright MCP

For other scraping tasks (not hiring.cafe), you can use the standard Playwright MCP:

```bash
claude mcp add playwright -s local -- npx @playwright/mcp@latest
```

## Architecture Summary

```
JobOps
├── Claude Code CLI (REQUIRED)
│   └── All core functionality
│       - Resume development (/buildresume)
│       - Assessment (/assessjob)
│       - Interview prep (/briefing, /interviewprep)
│       - OSINT (/osint)
│       - Career analysis (/change-one-thing)
│
├── Pandoc (OPTIONAL)
│   └── Document conversion (/convert)
│
└── Stealth Browser MCP (OPTIONAL)
    └── Job search automation (/searchjobs)
        - Works from residential IPs
        - Blocked from cloud environments
```

## Free Job Search APIs (Alternative to hiring.cafe)

If `/searchjobs` is blocked due to IP restrictions, use these free public APIs:

| API | Endpoint | Focus |
|-----|----------|-------|
| **The Muse** | `https://www.themuse.com/api/public/jobs` | US corporate jobs (477K+) |
| **Arbeitnow** | `https://arbeitnow.com/api/job-board-api` | EU/Global tech jobs |
| **Remotive** | `https://remotive.com/api/remote-jobs` | Remote jobs only |
| **Jobicy** | `https://jobicy.com/api/v2/remote-jobs` | Remote jobs only |

Example:
```bash
curl -s 'https://www.themuse.com/api/public/jobs?page=1&per_page=10' | jq '.results[] | {title: .name, company: .company.name}'
```

## What About the package.json?

The `package.json` file contains dependencies for optional features:

- **@modelcontextprotocol/sdk** - For the stealth browser MCP server
- **playwright-extra** - Stealth browser automation
- **puppeteer-extra-plugin-stealth** - Bot detection bypass
- **@playwright/test** - Browser installation

These packages are **not required** for core JobOps functionality.

## Quick Start Checklist

- [x] Install Claude Code CLI: `npm install -g @anthropic-ai/claude-code`
- [x] Clone/download this repository
- [x] Launch Claude Code in the repository: `claude`
- [ ] (Optional) Install Pandoc for document conversion: `/install-pandoc`
- [ ] (Optional) Install npm dependencies for job search: `npm install`
- [ ] (Optional) Test job search from local machine: `/searchjobs "software engineer"`

## Getting Help

For tactical assistance:
- **README.md** - Complete command reference and workflow guide
- **CLAUDE.md** - Technical implementation details for Claude Code
- **comprehensive_work_history_FAQ.md** - Master resume philosophy
- **SourceMaterial/** - Methodology documentation
