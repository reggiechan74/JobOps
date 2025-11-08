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

## Run the Playwright MCP server

### 1. Create a local configuration

Create `mcp/playwright.config.json` in the repository root (run `mkdir -p mcp` first if the directory does not exist) and tailor it to the sites you want the server to automate. The example below keeps Chrome headless and opens the hiring.cafe search experience used by `/searchjobs`:

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

Store the file in version control if you want a shared default, or add it to `.gitignore` and customize locally.

### 2. Launch the server

From the project root run:

```bash
npx @modelcontextprotocol/server-playwright --config mcp/playwright.config.json
```

The server listens on STDIN/STDOUT. Leave it running while you use Claude Code.

### 3. Register the server with Claude Code

Add the server to your local Claude Code configuration (usually `~/.claude/config.json`):

```json
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": ["@modelcontextprotocol/server-playwright", "--config", "mcp/playwright.config.json"],
      "cwd": "/path/to/resumeoptimizer"
    }
  }
}
```

Restart Claude Code so it picks up the new MCP endpoint.

### 4. Smoke-test the connection

With the server running, open Claude Code in this repository and run:

```bash
claude mcp status
```

You should see `playwright` listed as `connected`. If not, verify the config path, working directory, and that the server process is still running.

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
