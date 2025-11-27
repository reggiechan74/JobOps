#!/usr/bin/env node

/**
 * Stealth Browser MCP Server
 *
 * A custom MCP server that wraps playwright-extra with stealth plugin
 * to bypass bot detection when scraping job boards like hiring.cafe.
 *
 * Usage:
 *   node scripts/mcp/stealth-browser.js
 *
 * Add to Claude Code:
 *   claude mcp add stealth-browser node scripts/mcp/stealth-browser.js
 */

// Use subpath exports with wildcard pattern from package.json (./* -> ./dist/cjs/*)
const { Server } = require('@modelcontextprotocol/sdk/server/index.js');
const { StdioServerTransport } = require('@modelcontextprotocol/sdk/server/stdio.js');
const {
  CallToolRequestSchema,
  ListToolsRequestSchema,
} = require('@modelcontextprotocol/sdk/types.js');

let chromium, stealth, browser, page;

// Lazy load playwright-extra to handle missing dependencies gracefully
async function initBrowser() {
  if (!chromium) {
    const playwrightExtra = require('playwright-extra');
    chromium = playwrightExtra.chromium;
    stealth = require('puppeteer-extra-plugin-stealth')();
    chromium.use(stealth);
  }

  if (!browser) {
    browser = await chromium.launch({
      headless: true,
      args: [
        '--disable-blink-features=AutomationControlled',
        '--disable-dev-shm-usage',
        '--no-sandbox'
      ]
    });
    const context = await browser.newContext({
      userAgent: 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36',
      viewport: { width: 1920, height: 1080 },
      locale: 'en-US',
      timezoneId: 'America/New_York'
    });
    page = await context.newPage();
  }

  return { browser, page };
}

async function closeBrowser() {
  if (browser) {
    await browser.close();
    browser = null;
    page = null;
  }
}

// Create MCP server
const server = new Server(
  {
    name: 'stealth-browser',
    version: '1.0.0',
  },
  {
    capabilities: {
      tools: {},
    },
  }
);

// List available tools
server.setRequestHandler(ListToolsRequestSchema, async () => {
  return {
    tools: [
      {
        name: 'stealth_navigate',
        description: 'Navigate to a URL using stealth browser that bypasses bot detection',
        inputSchema: {
          type: 'object',
          properties: {
            url: {
              type: 'string',
              description: 'The URL to navigate to',
            },
            waitFor: {
              type: 'string',
              description: 'Wait condition: "load", "domcontentloaded", "networkidle"',
              default: 'networkidle'
            },
            timeout: {
              type: 'number',
              description: 'Timeout in milliseconds',
              default: 30000
            }
          },
          required: ['url'],
        },
      },
      {
        name: 'stealth_get_content',
        description: 'Get the text content of the current page',
        inputSchema: {
          type: 'object',
          properties: {
            selector: {
              type: 'string',
              description: 'CSS selector to get content from (optional, defaults to body)',
            },
          },
        },
      },
      {
        name: 'stealth_get_html',
        description: 'Get the HTML content of the current page',
        inputSchema: {
          type: 'object',
          properties: {
            selector: {
              type: 'string',
              description: 'CSS selector to get HTML from (optional, defaults to full page)',
            },
          },
        },
      },
      {
        name: 'stealth_evaluate',
        description: 'Execute JavaScript in the page context',
        inputSchema: {
          type: 'object',
          properties: {
            script: {
              type: 'string',
              description: 'JavaScript code to execute',
            },
          },
          required: ['script'],
        },
      },
      {
        name: 'stealth_fetch_api',
        description: 'Make a fetch request from within the browser context (inherits cookies/session)',
        inputSchema: {
          type: 'object',
          properties: {
            url: {
              type: 'string',
              description: 'API URL to fetch',
            },
            method: {
              type: 'string',
              description: 'HTTP method (GET, POST, etc.)',
              default: 'GET'
            },
            body: {
              type: 'string',
              description: 'Request body (JSON string for POST requests)',
            },
            headers: {
              type: 'object',
              description: 'Additional headers',
            }
          },
          required: ['url'],
        },
      },
      {
        name: 'stealth_wait',
        description: 'Wait for a specified time or element',
        inputSchema: {
          type: 'object',
          properties: {
            milliseconds: {
              type: 'number',
              description: 'Time to wait in milliseconds',
            },
            selector: {
              type: 'string',
              description: 'CSS selector to wait for',
            },
          },
        },
      },
      {
        name: 'stealth_screenshot',
        description: 'Take a screenshot of the current page',
        inputSchema: {
          type: 'object',
          properties: {
            path: {
              type: 'string',
              description: 'File path to save screenshot',
            },
            fullPage: {
              type: 'boolean',
              description: 'Capture full page',
              default: false
            }
          },
        },
      },
      {
        name: 'stealth_close',
        description: 'Close the browser',
        inputSchema: {
          type: 'object',
          properties: {},
        },
      },
    ],
  };
});

// Handle tool calls
server.setRequestHandler(CallToolRequestSchema, async (request) => {
  const { name, arguments: args } = request.params;

  try {
    switch (name) {
      case 'stealth_navigate': {
        const { browser, page } = await initBrowser();
        await page.goto(args.url, {
          waitUntil: args.waitFor || 'networkidle',
          timeout: args.timeout || 30000,
        });
        const title = await page.title();
        const url = page.url();
        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify({ success: true, title, url }, null, 2),
            },
          ],
        };
      }

      case 'stealth_get_content': {
        if (!page) throw new Error('No page open. Call stealth_navigate first.');
        const content = args.selector
          ? await page.$eval(args.selector, el => el.innerText)
          : await page.evaluate(() => document.body.innerText);
        return {
          content: [{ type: 'text', text: content }],
        };
      }

      case 'stealth_get_html': {
        if (!page) throw new Error('No page open. Call stealth_navigate first.');
        const html = args.selector
          ? await page.$eval(args.selector, el => el.outerHTML)
          : await page.content();
        return {
          content: [{ type: 'text', text: html }],
        };
      }

      case 'stealth_evaluate': {
        if (!page) throw new Error('No page open. Call stealth_navigate first.');
        const result = await page.evaluate(args.script);
        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify(result, null, 2),
            },
          ],
        };
      }

      case 'stealth_fetch_api': {
        if (!page) throw new Error('No page open. Call stealth_navigate first.');
        const result = await page.evaluate(async (params) => {
          const options = {
            method: params.method || 'GET',
            headers: {
              'Content-Type': 'application/json',
              ...params.headers
            }
          };
          if (params.body) {
            options.body = params.body;
          }
          const response = await fetch(params.url, options);
          const text = await response.text();
          try {
            return { status: response.status, data: JSON.parse(text) };
          } catch {
            return { status: response.status, data: text };
          }
        }, args);
        return {
          content: [
            {
              type: 'text',
              text: JSON.stringify(result, null, 2),
            },
          ],
        };
      }

      case 'stealth_wait': {
        if (!page) throw new Error('No page open. Call stealth_navigate first.');
        if (args.milliseconds) {
          await page.waitForTimeout(args.milliseconds);
        }
        if (args.selector) {
          await page.waitForSelector(args.selector);
        }
        return {
          content: [{ type: 'text', text: 'Wait completed' }],
        };
      }

      case 'stealth_screenshot': {
        if (!page) throw new Error('No page open. Call stealth_navigate first.');
        const path = args.path || `/tmp/screenshot-${Date.now()}.png`;
        await page.screenshot({ path, fullPage: args.fullPage || false });
        return {
          content: [{ type: 'text', text: `Screenshot saved to ${path}` }],
        };
      }

      case 'stealth_close': {
        await closeBrowser();
        return {
          content: [{ type: 'text', text: 'Browser closed' }],
        };
      }

      default:
        throw new Error(`Unknown tool: ${name}`);
    }
  } catch (error) {
    return {
      content: [
        {
          type: 'text',
          text: JSON.stringify({ error: error.message }, null, 2),
        },
      ],
      isError: true,
    };
  }
});

// Start server
async function main() {
  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error('Stealth Browser MCP server started');
}

main().catch(console.error);

// Cleanup on exit
process.on('SIGINT', async () => {
  await closeBrowser();
  process.exit(0);
});

process.on('SIGTERM', async () => {
  await closeBrowser();
  process.exit(0);
});
