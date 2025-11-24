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

### Playwright MCP Server (Job Search)

**Required for**: `/searchjobs` command only

**Status**: Currently blocked/non-functional

The `/searchjobs` command uses Playwright browser automation to scrape hiring.cafe. This requires additional setup:

#### Prerequisites

- Node.js (v18 or higher)
- npm (comes with Node.js)

#### Installation Steps

1. **Install npm dependencies:**
   ```bash
   npm install
   ```

2. **Install Playwright Chrome browser:**
   ```bash
   npm run install-browsers
   ```

   Or directly:
   ```bash
   npx playwright install chrome
   ```

#### Configure Playwright MCP Server

**Note**: The package name in the old configuration was incorrect. The correct package is `@playwright/mcp`.

1. **Create MCP configuration** (if needed):

   Create `mcp/playwright.config.json`:
   ```json
   {
     "browser": "chromium",
     "launchOptions": {
       "headless": true
     },
     "pages": [
       {
         "name": "hiring-cafe",
         "url": "https://hiring.cafe"
       }
     ]
   }
   ```

2. **Register with Claude Code**:

   Add to your Claude Code configuration (usually `~/.claude/config.json`):
   ```json
   {
     "mcpServers": {
       "playwright": {
         "command": "npx",
         "args": ["@playwright/mcp"],
         "cwd": "/path/to/resumeoptimizer"
       }
     }
   }
   ```

3. **Restart Claude Code** to pick up the new MCP server.

4. **Verify connection**:
   ```bash
   claude mcp status
   ```

   You should see `playwright` listed as `connected`.

#### Troubleshooting Playwright

**Browser Not Found Error:**
```bash
npx playwright install --force chrome
```

**Verify Installation:**
```bash
npx playwright --version
google-chrome --version  # Or chromium-browser --version
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
└── Playwright MCP (OPTIONAL - Currently blocked)
    └── Job search automation (/searchjobs)
```

## What About the package.json?

The `package.json` file exists solely for the optional `/searchjobs` functionality. The dependencies listed there are:

- **No runtime dependencies** (removed incorrect package)
- **@playwright/test** (dev dependency) - for browser installation only

These packages are **not required** for core JobOps functionality.

## Quick Start Checklist

- [x] Install Claude Code CLI: `npm install -g @anthropic-ai/claude-code`
- [x] Clone/download this repository
- [x] Launch Claude Code in the repository: `claude`
- [ ] (Optional) Install Pandoc for document conversion
- [ ] (Optional) Configure Playwright MCP for job search (currently blocked)

## Getting Help

For tactical assistance:
- **README.md** - Complete command reference and workflow guide
- **CLAUDE.md** - Technical implementation details for Claude Code
- **comprehensive_work_history_FAQ.md** - Master resume philosophy
- **SourceMaterial/** - Methodology documentation
