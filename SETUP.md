# ResumeOptimizer Setup Instructions

## Prerequisites

- **Claude Code CLI**: Install via npm:
  ```bash
  npm install -g @anthropic-ai/claude-code
  ```
  More information at [claude.ai/code](https://claude.ai/code)
- Node.js (v18 or higher)
- npm (comes with Node.js)

## Installation

### Install All Dependencies

Run the following command to install both npm packages and Playwright browser:

```bash
npm run install-all
```

This will:
1. Install the MCP Playwright server package
2. Install Playwright test framework
3. Download and install the Chrome browser for Playwright

### Manual Installation Steps

If you prefer to install components separately:

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

## Verification

After installation, verify that all components are working:

```bash
# Verify Claude Code
claude --version

# Verify Playwright
npx playwright --version

# Verify Chrome browser
google-chrome --version
```

## Troubleshooting

### Browser Not Found Error

If you encounter an error like "Chromium distribution 'chrome' is not found", run:

```bash
npx playwright install --force chrome
```

### Permission Issues

If you encounter permission issues during installation, you may need to run with elevated privileges (though this is rarely needed in development environments).

## Dependencies

- **@modelcontextprotocol/server-playwright**: MCP server for Playwright automation
- **@playwright/test**: Playwright testing framework and browser automation library

The Playwright browsers are installed separately from the npm packages to keep package sizes manageable.