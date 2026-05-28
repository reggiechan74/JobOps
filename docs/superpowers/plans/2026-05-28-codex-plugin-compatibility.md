# Codex Plugin Compatibility Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make the existing `jobops` and `jobops-ic` Claude Code plugins installable, discoverable, and skill-scan clean in Codex without duplicating the plugin tree.

**Architecture:** Add Codex manifests and a Codex repo marketplace alongside the existing Claude manifests. Keep one canonical `skills/` tree per plugin, add Codex-required `name` frontmatter to every skill, and add a static validator that guards marketplace shape, manifest drift, and skill metadata.

**Tech Stack:** Node.js validation script, JSON plugin manifests, Markdown skill frontmatter, existing JobOps docs, Codex CLI smoke checks, optional Claude Code validator.

---

## Context Checked

- Approved design spec: `docs/superpowers/specs/2026-05-28-codex-plugin-compatibility-design.md`
- Codex plugin docs: `https://developers.openai.com/codex/plugins/build`
- Codex skills docs: `https://developers.openai.com/codex/skills`
- Codex subagents docs: `https://developers.openai.com/codex/subagents`
- Existing Claude manifests:
  - `.claude-plugin/marketplace.json`
  - `plugins/jobops/.claude-plugin/plugin.json`
  - `plugins/jobops-ic/.claude-plugin/plugin.json`
- Existing validator style:
  - `scripts/validate/validate-legacy-profile-block.js`
  - `scripts/validate/validate-job-postings.js`

## File Structure

**Create:**
- `.agents/plugins/marketplace.json` - Codex repo marketplace exposing `jobops` and `jobops-ic`.
- `plugins/jobops/.codex-plugin/plugin.json` - Codex manifest for the core plugin.
- `plugins/jobops-ic/.codex-plugin/plugin.json` - Codex manifest for the IC add-on.
- `scripts/validate/validate-codex-plugin-compatibility.js` - static package contract validator.

**Modify:**
- `package.json` - add `test` and `validate:codex-plugin` scripts.
- `scripts/validate/README.md` - document the new validator.
- All `plugins/jobops/skills/*/SKILL.md` files - add a `name` value matching each skill directory.
- All `plugins/jobops-ic/skills/*/SKILL.md` files - add a `name` value matching each skill directory.
- `README.md` - describe Claude Code and Codex install paths.
- `SETUP.md` - add Codex setup flow and troubleshooting.
- `CLAUDE.md` - update contributor structure and validation notes.
- `docs/ARCHITECTURE.md` - add the dual-platform compatibility contract.
- `CHANGELOG.md` - record Codex install/discovery compatibility.

**Do not create:**
- A mirrored Codex-only plugin tree.
- Per-skill `agents/openai.yaml` metadata files in this pass.
- Codex custom-agent TOML files for the existing Claude agents.

---

### Task 1: Add the Static Codex Compatibility Validator

**Files:**
- Create: `scripts/validate/validate-codex-plugin-compatibility.js`
- Modify: `package.json`
- Modify: `scripts/validate/README.md`

This task creates the failing contract first. It should fail before the Codex manifests and skill names exist.

- [ ] **Step 1: Create the validator**

Create `scripts/validate/validate-codex-plugin-compatibility.js`:

```javascript
#!/usr/bin/env node
/**
 * Validate JobOps Codex plugin compatibility.
 *
 * Usage: node scripts/validate/validate-codex-plugin-compatibility.js
 */

const fs = require('fs');
const path = require('path');

const REPO_ROOT = path.resolve(__dirname, '..', '..');
const PLUGINS = [
  {
    name: 'jobops',
    root: 'plugins/jobops',
    displayName: 'JobOps',
    category: 'Productivity'
  },
  {
    name: 'jobops-ic',
    root: 'plugins/jobops-ic',
    displayName: 'JobOps IC',
    category: 'Productivity'
  }
];

let failures = 0;

function fail(message) {
  failures += 1;
  console.error(`FAIL: ${message}`);
}

function check(condition, message) {
  if (!condition) {
    fail(message);
  }
}

function repoPath(relativePath) {
  return path.join(REPO_ROOT, relativePath);
}

function exists(relativePath) {
  return fs.existsSync(repoPath(relativePath));
}

function readText(relativePath) {
  const absolutePath = repoPath(relativePath);
  if (!fs.existsSync(absolutePath)) {
    fail(`Missing file: ${relativePath}`);
    return null;
  }
  return fs.readFileSync(absolutePath, 'utf8');
}

function readJson(relativePath) {
  const text = readText(relativePath);
  if (text === null) {
    return null;
  }
  try {
    return JSON.parse(text);
  } catch (error) {
    fail(`${relativePath} is not valid JSON: ${error.message}`);
    return null;
  }
}

function extractFrontmatter(relativePath) {
  const text = readText(relativePath);
  if (text === null) {
    return null;
  }
  const match = text.match(/^---\n([\s\S]*?)\n---/);
  if (!match) {
    fail(`${relativePath} must start with YAML frontmatter`);
    return null;
  }

  const data = {};
  for (const rawLine of match[1].split('\n')) {
    const line = rawLine.trim();
    if (line === '' || line.startsWith('#')) {
      continue;
    }
    const colonIndex = line.indexOf(':');
    if (colonIndex === -1) {
      continue;
    }
    const key = line.slice(0, colonIndex).trim();
    const value = line.slice(colonIndex + 1).trim().replace(/^["']|["']$/g, '');
    data[key] = value;
  }
  return data;
}

function assertVersion(label, actual, expected) {
  check(actual === expected, `${label} version must be ${expected}, got ${actual || '(missing)'}`);
}

function assertRelativePath(label, actual, expected) {
  check(actual === expected, `${label} must be ${expected}, got ${actual || '(missing)'}`);
}

function validatePackageScripts(packageJson) {
  check(packageJson.scripts, 'package.json must define scripts');
  if (!packageJson.scripts) {
    return;
  }
  check(
    packageJson.scripts['validate:codex-plugin'] === 'node scripts/validate/validate-codex-plugin-compatibility.js',
    'package.json scripts.validate:codex-plugin must run the Codex compatibility validator'
  );
  check(
    packageJson.scripts.test === 'npm run validate:codex-plugin',
    'package.json scripts.test must run validate:codex-plugin'
  );
}

function validateClaudeMarketplace(expectedVersion) {
  const marketplace = readJson('.claude-plugin/marketplace.json');
  if (!marketplace) {
    return;
  }

  check(marketplace.name === 'jobops-marketplace', '.claude-plugin marketplace name must be jobops-marketplace');
  check(marketplace.metadata, '.claude-plugin marketplace must include metadata');
  if (marketplace.metadata) {
    assertVersion('.claude-plugin marketplace metadata', marketplace.metadata.version, expectedVersion);
  }
  check(Array.isArray(marketplace.plugins), '.claude-plugin marketplace plugins must be an array');
}

function validateCodexMarketplace(expectedVersion) {
  const marketplace = readJson('.agents/plugins/marketplace.json');
  if (!marketplace) {
    return;
  }

  check(marketplace.name === 'jobops-marketplace', '.agents/plugins marketplace name must be jobops-marketplace');
  check(marketplace.interface, '.agents/plugins marketplace must include interface metadata');
  if (marketplace.interface) {
    check(
      marketplace.interface.displayName === 'JobOps Marketplace',
      '.agents/plugins marketplace interface.displayName must be JobOps Marketplace'
    );
  }
  check(marketplace.metadata, '.agents/plugins marketplace must include metadata');
  if (marketplace.metadata) {
    assertVersion('.agents/plugins marketplace metadata', marketplace.metadata.version, expectedVersion);
  }
  check(Array.isArray(marketplace.plugins), '.agents/plugins marketplace plugins must be an array');
  if (!Array.isArray(marketplace.plugins)) {
    return;
  }

  const expectedOrder = ['jobops', 'jobops-ic'];
  const actualOrder = marketplace.plugins.map((plugin) => plugin.name);
  check(
    actualOrder.join(',') === expectedOrder.join(','),
    `.agents/plugins marketplace plugin order must be ${expectedOrder.join(',')}, got ${actualOrder.join(',')}`
  );

  for (const plugin of PLUGINS) {
    const entry = marketplace.plugins.find((candidate) => candidate.name === plugin.name);
    check(entry, `.agents/plugins marketplace must include ${plugin.name}`);
    if (!entry) {
      continue;
    }
    check(entry.source && entry.source.source === 'local', `${plugin.name} marketplace source.source must be local`);
    assertRelativePath(`${plugin.name} marketplace source.path`, entry.source && entry.source.path, `./${plugin.root}`);
    check(entry.policy, `${plugin.name} marketplace entry must include policy`);
    if (entry.policy) {
      check(entry.policy.installation === 'AVAILABLE', `${plugin.name} policy.installation must be AVAILABLE`);
      check(entry.policy.authentication === 'ON_INSTALL', `${plugin.name} policy.authentication must be ON_INSTALL`);
    }
    check(entry.category === plugin.category, `${plugin.name} category must be ${plugin.category}`);
  }
}

function validateCodexManifestPathField(pluginName, field, value, pluginRoot) {
  if (value === undefined) {
    return;
  }
  check(typeof value === 'string', `${pluginName} ${field} must be a string path`);
  if (typeof value !== 'string') {
    return;
  }
  check(value.startsWith('./'), `${pluginName} ${field} path must start with ./`);
  const normalized = path.normalize(value);
  check(!normalized.startsWith('..'), `${pluginName} ${field} path must stay inside plugin root`);
  check(
    fs.existsSync(path.join(REPO_ROOT, pluginRoot, value)),
    `${pluginName} ${field} points to missing path ${value}`
  );
}

function validatePluginManifests(plugin, expectedVersion) {
  const claude = readJson(`${plugin.root}/.claude-plugin/plugin.json`);
  const codex = readJson(`${plugin.root}/.codex-plugin/plugin.json`);

  if (claude) {
    check(claude.name === plugin.name, `${plugin.name} Claude manifest name must be ${plugin.name}`);
    assertVersion(`${plugin.name} Claude manifest`, claude.version, expectedVersion);
  }

  if (!codex) {
    return;
  }

  check(codex.name === plugin.name, `${plugin.name} Codex manifest name must be ${plugin.name}`);
  assertVersion(`${plugin.name} Codex manifest`, codex.version, expectedVersion);
  assertRelativePath(`${plugin.name} Codex manifest skills`, codex.skills, './skills/');
  check(codex.license === 'ISC', `${plugin.name} Codex manifest license must be ISC`);
  check(
    codex.repository === 'https://github.com/reggiechan74/JobOps',
    `${plugin.name} Codex manifest repository must be https://github.com/reggiechan74/JobOps`
  );
  check(codex.interface, `${plugin.name} Codex manifest must include interface metadata`);
  if (codex.interface) {
    check(codex.interface.displayName === plugin.displayName, `${plugin.name} interface.displayName must be ${plugin.displayName}`);
    check(Boolean(codex.interface.shortDescription), `${plugin.name} interface.shortDescription must be present`);
    check(Boolean(codex.interface.longDescription), `${plugin.name} interface.longDescription must be present`);
    check(codex.interface.developerName === 'Reggie Chan', `${plugin.name} interface.developerName must be Reggie Chan`);
    check(codex.interface.category === plugin.category, `${plugin.name} interface.category must be ${plugin.category}`);
    check(Array.isArray(codex.interface.capabilities), `${plugin.name} interface.capabilities must be an array`);
    check(codex.interface.websiteURL === 'https://github.com/reggiechan74/JobOps', `${plugin.name} interface.websiteURL must be repository URL`);
  }

  check(!Object.prototype.hasOwnProperty.call(codex, 'dependencies'), `${plugin.name} Codex manifest must not declare dependencies`);
  for (const field of ['apps', 'mcpServers', 'hooks']) {
    validateCodexManifestPathField(plugin.name, field, codex[field], plugin.root);
  }
}

function validateSkillFrontmatter(plugin) {
  const skillsRoot = repoPath(`${plugin.root}/skills`);
  check(fs.existsSync(skillsRoot), `${plugin.name} skills directory must exist`);
  if (!fs.existsSync(skillsRoot)) {
    return;
  }

  const skillDirs = fs.readdirSync(skillsRoot, { withFileTypes: true })
    .filter((entry) => entry.isDirectory())
    .map((entry) => entry.name)
    .sort();

  check(skillDirs.length > 0, `${plugin.name} must contain at least one skill`);

  for (const skillName of skillDirs) {
    const relativeSkillPath = `${plugin.root}/skills/${skillName}/SKILL.md`;
    const frontmatter = extractFrontmatter(relativeSkillPath);
    if (!frontmatter) {
      continue;
    }
    check(frontmatter.name === skillName, `${relativeSkillPath} frontmatter name must be ${skillName}`);
    check(Boolean(frontmatter.description), `${relativeSkillPath} frontmatter description must be present`);
    check(
      frontmatter['disable-model-invocation'] === 'true',
      `${relativeSkillPath} must keep disable-model-invocation: true`
    );
  }
}

function main() {
  const packageJson = readJson('package.json');
  if (!packageJson) {
    process.exit(1);
  }

  check(packageJson.version === '2.6.1', `package.json version must be 2.6.1 for this release, got ${packageJson.version}`);
  validatePackageScripts(packageJson);
  validateClaudeMarketplace(packageJson.version);
  validateCodexMarketplace(packageJson.version);

  for (const plugin of PLUGINS) {
    validatePluginManifests(plugin, packageJson.version);
    validateSkillFrontmatter(plugin);
  }

  if (failures > 0) {
    console.error(`\nCodex compatibility validation failed with ${failures} issue(s).`);
    process.exit(1);
  }

  console.log('\nCodex compatibility validation passed.');
}

main();
```

- [ ] **Step 2: Make it executable**

Run:

```bash
chmod +x scripts/validate/validate-codex-plugin-compatibility.js
```

Expected: command exits `0`.

- [ ] **Step 3: Add package scripts**

Edit `package.json` so the scripts object becomes:

```json
  "scripts": {
    "install-browsers": "npx playwright install chrome",
    "install-all": "npm install && npm run install-browsers",
    "validate:codex-plugin": "node scripts/validate/validate-codex-plugin-compatibility.js",
    "test": "npm run validate:codex-plugin"
  },
```

- [ ] **Step 4: Document the validator**

In `scripts/validate/README.md`, insert this section after `validate-legacy-profile-block.js`:

````markdown
### validate-codex-plugin-compatibility.js

Validates that JobOps can be discovered and installed as Codex plugins without drifting from the existing Claude Code plugin manifests.

**Usage:**
```bash
node scripts/validate/validate-codex-plugin-compatibility.js
```

**Checks:**
- Codex repo marketplace exists at `.agents/plugins/marketplace.json`
- `jobops` and `jobops-ic` have `.codex-plugin/plugin.json`
- Codex marketplace entries use `./plugins/jobops` and `./plugins/jobops-ic`
- Codex manifests point to `./skills/`
- Versions match across `package.json`, Claude manifests, Codex manifests, and marketplace metadata
- Every bundled `SKILL.md` has `name` and `description` frontmatter
- Codex manifests do not point to missing `apps`, `mcpServers`, or `hooks` paths

**Exit codes:**
- 0: Codex compatibility contract passes
- 1: One or more compatibility requirements are missing
````

- [ ] **Step 5: Run the validator and verify RED**

Run:

```bash
node scripts/validate/validate-codex-plugin-compatibility.js
```

Expected: `FAIL` lines including:

```text
FAIL: Missing file: .agents/plugins/marketplace.json
FAIL: Missing file: plugins/jobops/.codex-plugin/plugin.json
FAIL: Missing file: plugins/jobops-ic/.codex-plugin/plugin.json
```

Expected final line:

```text
Codex compatibility validation failed with
```

- [ ] **Step 6: Commit the failing contract**

Run:

```bash
git add package.json scripts/validate/validate-codex-plugin-compatibility.js scripts/validate/README.md
git commit -m "test: define Codex plugin compatibility contract"
```

---

### Task 2: Add Codex Marketplace and Plugin Manifests

**Files:**
- Create: `.agents/plugins/marketplace.json`
- Create: `plugins/jobops/.codex-plugin/plugin.json`
- Create: `plugins/jobops-ic/.codex-plugin/plugin.json`

This task satisfies the packaging part of the contract. The validator should still fail afterward because skill frontmatter names are not present yet.

- [ ] **Step 1: Create the Codex marketplace directory**

Run:

```bash
mkdir -p .agents/plugins plugins/jobops/.codex-plugin plugins/jobops-ic/.codex-plugin
```

Expected: command exits `0`.

- [ ] **Step 2: Add `.agents/plugins/marketplace.json`**

Create `.agents/plugins/marketplace.json`:

```json
{
  "name": "jobops-marketplace",
  "interface": {
    "displayName": "JobOps Marketplace"
  },
  "metadata": {
    "description": "Intelligence-driven career management plugins for Claude Code and Codex",
    "version": "2.6.1"
  },
  "plugins": [
    {
      "name": "jobops",
      "source": {
        "source": "local",
        "path": "./plugins/jobops"
      },
      "policy": {
        "installation": "AVAILABLE",
        "authentication": "ON_INSTALL"
      },
      "category": "Productivity"
    },
    {
      "name": "jobops-ic",
      "source": {
        "source": "local",
        "path": "./plugins/jobops-ic"
      },
      "policy": {
        "installation": "AVAILABLE",
        "authentication": "ON_INSTALL"
      },
      "category": "Productivity"
    }
  ]
}
```

- [ ] **Step 3: Add the core Codex manifest**

Create `plugins/jobops/.codex-plugin/plugin.json`:

```json
{
  "name": "jobops",
  "version": "2.6.1",
  "description": "Intelligence-driven job application system - resume development, interview prep, OSINT intelligence, career strategy, and crisis management using HAM-Z methodology",
  "author": {
    "name": "Reggie Chan"
  },
  "homepage": "https://github.com/reggiechan74/JobOps",
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": [
    "resume",
    "job-search",
    "interview-prep",
    "osint",
    "career",
    "assessment",
    "crisis-management"
  ],
  "skills": "./skills/",
  "interface": {
    "displayName": "JobOps",
    "shortDescription": "Resume, interview, OSINT, career strategy, and crisis workflows",
    "longDescription": "Use JobOps for structured career inventory setup, resume development, cover letters, candidate assessment, interview preparation, company intelligence, career strategy, and employment-crisis documentation.",
    "developerName": "Reggie Chan",
    "category": "Productivity",
    "capabilities": [
      "Interactive",
      "Read",
      "Write",
      "Web"
    ],
    "websiteURL": "https://github.com/reggiechan74/JobOps"
  }
}
```

- [ ] **Step 4: Add the IC Codex manifest**

Create `plugins/jobops-ic/.codex-plugin/plugin.json`:

```json
{
  "name": "jobops-ic",
  "version": "2.6.1",
  "description": "Independent contractor toolkit - services, client prospecting, pitch decks, proposals, rate cards, landing pages. Requires jobops.",
  "author": {
    "name": "Reggie Chan"
  },
  "homepage": "https://github.com/reggiechan74/JobOps",
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": [
    "independent-contractor",
    "consulting",
    "proposals",
    "rate-card",
    "pitch-deck",
    "landing-page"
  ],
  "skills": "./skills/",
  "interface": {
    "displayName": "JobOps IC",
    "shortDescription": "Independent contractor services, prospecting, proposals, rate cards, and landing pages",
    "longDescription": "Use JobOps IC with the core JobOps workspace to define services, find clients, create pitch decks, draft consulting proposals, publish rate cards, and build landing-page copy and pages.",
    "developerName": "Reggie Chan",
    "category": "Productivity",
    "capabilities": [
      "Interactive",
      "Read",
      "Write",
      "Web"
    ],
    "websiteURL": "https://github.com/reggiechan74/JobOps"
  }
}
```

- [ ] **Step 5: Validate JSON syntax**

Run:

```bash
python3 -m json.tool .agents/plugins/marketplace.json >/dev/null
python3 -m json.tool plugins/jobops/.codex-plugin/plugin.json >/dev/null
python3 -m json.tool plugins/jobops-ic/.codex-plugin/plugin.json >/dev/null
```

Expected: all commands exit `0` and print nothing.

- [ ] **Step 6: Run the validator and verify the remaining RED**

Run:

```bash
node scripts/validate/validate-codex-plugin-compatibility.js
```

Expected: the missing-manifest failures are gone. Remaining failures are skill frontmatter failures such as:

```text
frontmatter name must be
```

- [ ] **Step 7: Commit Codex manifests**

Run:

```bash
git add .agents/plugins/marketplace.json plugins/jobops/.codex-plugin/plugin.json plugins/jobops-ic/.codex-plugin/plugin.json
git commit -m "feat: add Codex plugin manifests"
```

---

### Task 3: Add Codex-Required Skill Names

**Files:**
- Modify: `plugins/jobops/skills/*/SKILL.md`
- Modify: `plugins/jobops-ic/skills/*/SKILL.md`

This is a mechanical frontmatter rewrite. Do not change skill bodies.

- [ ] **Step 1: Add `name` to every skill frontmatter**

Run this one-time mechanical rewrite:

```bash
node <<'NODE'
const fs = require('fs');
const path = require('path');

const roots = ['plugins/jobops/skills', 'plugins/jobops-ic/skills'];

for (const root of roots) {
  const skillDirs = fs.readdirSync(root, { withFileTypes: true })
    .filter((entry) => entry.isDirectory())
    .map((entry) => entry.name)
    .sort();

  for (const skillName of skillDirs) {
    const filePath = path.join(root, skillName, 'SKILL.md');
    let text = fs.readFileSync(filePath, 'utf8');

    if (!text.startsWith('---\n')) {
      throw new Error(`${filePath} does not start with YAML frontmatter`);
    }

    const frontmatterEnd = text.indexOf('\n---', 4);
    if (frontmatterEnd === -1) {
      throw new Error(`${filePath} does not close YAML frontmatter`);
    }

    const frontmatter = text.slice(4, frontmatterEnd);
    if (/^name:/m.test(frontmatter)) {
      continue;
    }

    text = `---\nname: ${skillName}\n${text.slice(4)}`;
    fs.writeFileSync(filePath, text);
  }
}
NODE
```

Expected: command exits `0`.

- [ ] **Step 2: Inspect the mechanical diff**

Run:

```bash
git diff --stat plugins/jobops/skills plugins/jobops-ic/skills
git diff -- plugins/jobops/skills/setup/SKILL.md plugins/jobops-ic/skills/setup/SKILL.md | sed -n '1,80p'
```

Expected: every changed `SKILL.md` has one added `name` line near the top, and that value matches the skill directory name. No body text changes.

- [ ] **Step 3: Verify every skill has a matching name**

Run:

```bash
node scripts/validate/validate-codex-plugin-compatibility.js
```

Expected:

```text
Codex compatibility validation passed.
```

- [ ] **Step 4: Commit skill frontmatter names**

Run:

```bash
git add plugins/jobops/skills plugins/jobops-ic/skills
git commit -m "feat: add Codex skill names"
```

---

### Task 4: Update User and Contributor Documentation

**Files:**
- Modify: `README.md`
- Modify: `SETUP.md`
- Modify: `CLAUDE.md`
- Modify: `docs/ARCHITECTURE.md`
- Modify: `CHANGELOG.md`

This task documents dual compatibility without rewriting every skill example.

- [ ] **Step 1: Update the README title copy**

In `README.md`, replace:

```markdown
Two Claude Code plugins for systematic, intelligence-driven career management — from resume development to independent consulting.
```

with:

```markdown
Two Claude Code and Codex plugins for systematic, intelligence-driven career management — from resume development to independent consulting.
```

- [ ] **Step 2: Split README Quick Start by platform**

Replace the existing `## Quick Start` command block in `README.md` with:

````markdown
## Quick Start

### Claude Code

```bash
# Add the marketplace
/plugin marketplace add reggiechan74/JobOps

# Install the core plugin
/plugin install jobops@jobops-marketplace

# Initialize your workspace
/jobops:setup

# Optional: install the independent contractor add-on
/plugin install jobops-ic@jobops-marketplace
/jobops-ic:setup
```

### Codex

```bash
# Add the marketplace
codex plugin marketplace add reggiechan74/JobOps

# Install the core plugin
codex plugin add jobops@jobops-marketplace

# Optional: install the independent contractor add-on
codex plugin add jobops-ic@jobops-marketplace
```

After installation, start a new Codex session and invoke `jobops:setup` through Codex's explicit skill/plugin invocation surface. If you installed `jobops-ic`, run `jobops-ic:setup` after the core setup completes.
````

- [ ] **Step 3: Update README contributor structure**

In the `Repository Structure` block in `README.md`, add Codex entries so the relevant section reads:

```text
JobOps/
  .claude-plugin/marketplace.json    # Claude Code marketplace manifest
  .agents/plugins/marketplace.json   # Codex marketplace manifest
  plugins/
    jobops/                          # Core plugin
      .claude-plugin/plugin.json
      .codex-plugin/plugin.json
      skills/                        # 35 skills (flat layout)
      agents/                        # 16 Claude Code agents
      templates/                     # Bundled templates
    jobops-ic/                       # IC add-on plugin
      .claude-plugin/plugin.json
      .codex-plugin/plugin.json
      skills/                        # 10 skills
      agents/                        # 1 Claude Code agent
      templates/                     # Bundled templates
```

- [ ] **Step 4: Update README development commands**

Replace the contributor validation snippet in `README.md` with:

```bash
# Test a Claude Code plugin locally
claude --plugin-dir plugins/jobops

# Validate Claude Code plugin structure
claude plugin validate plugins/jobops
claude plugin validate plugins/jobops-ic
claude plugin validate .

# Validate Codex plugin structure
npm test
codex plugin marketplace add ./
codex plugin list
```

- [ ] **Step 5: Update SETUP with platform installation sections**

In `SETUP.md`, replace the opening sentence:

```markdown
JobOps v2.0 ships as two Claude Code plugins distributed through a self-hosted marketplace. Setup is a three-step flow: install Claude Code, install the plugin(s), then initialize the workspace.
```

with:

```markdown
JobOps v2.6.1 ships as two plugins for Claude Code and Codex through self-hosted marketplaces. Setup is a three-step flow: install a supported agent client, install the plugin(s), then initialize the workspace.
```

Rename `## 1. Install Claude Code` to:

```markdown
## 1. Install a Supported Client
```

Under the Claude Code install commands, add:

````markdown
For Codex, install and authenticate Codex using OpenAI's current Codex setup flow, then verify:

```bash
codex --version
```
````

Rename `## 2. Install the Plugins` to:

```markdown
## 2. Install the Plugins
```

Keep the Claude Code commands, then add:

````markdown
For Codex:

```bash
codex plugin marketplace add reggiechan74/JobOps
codex plugin add jobops@jobops-marketplace
```

Optional independent contractor toolkit:

```bash
codex plugin add jobops-ic@jobops-marketplace
```
````

Under `## 3. Initialize the Workspace`, add after the Claude `/jobops:setup` block:

```markdown
In Codex, start a new session and invoke `jobops:setup` through the explicit skill/plugin invocation surface. If `jobops-ic` is installed, invoke `jobops-ic:setup` after the core setup completes.
```

- [ ] **Step 6: Update SETUP checklist and troubleshooting**

Replace the setup checklist in `SETUP.md` with:

```markdown
## Quick Start Checklist

### Claude Code

- [ ] Install Claude Code: `npm install -g @anthropic-ai/claude-code`
- [ ] In Claude: `/plugin marketplace add reggiechan74/JobOps`
- [ ] In Claude: `/plugin install jobops@jobops-marketplace`
- [ ] In Claude: `/jobops:setup`
- [ ] (Optional) `/plugin install jobops-ic@jobops-marketplace` then `/jobops-ic:setup`
- [ ] (Optional) `/jobops:install-pandoc` for Word export
- [ ] (Optional) `npx playwright install chromium` for PDF export
- [ ] (v1.x users) `/jobops:migrate` to relocate existing outputs

### Codex

- [ ] Verify Codex: `codex --version`
- [ ] In shell: `codex plugin marketplace add reggiechan74/JobOps`
- [ ] In shell: `codex plugin add jobops@jobops-marketplace`
- [ ] In a new Codex session: invoke `jobops:setup`
- [ ] (Optional) `codex plugin add jobops-ic@jobops-marketplace` then invoke `jobops-ic:setup`
- [ ] (Optional) invoke `jobops:install-pandoc` for Word export
- [ ] (Optional) `npx playwright install chromium` for PDF export
- [ ] (v1.x users) invoke `jobops:migrate` to relocate existing outputs
```

Add this troubleshooting item:

```markdown
**Codex marketplace not listed**: run `codex plugin marketplace list` and confirm `jobops-marketplace` appears. If it does not, rerun `codex plugin marketplace add reggiechan74/JobOps`, then start a new Codex session.
```

- [ ] **Step 7: Update CLAUDE.md contributor notes**

In `CLAUDE.md`, change the version description line to:

```markdown
**JobOps v2.6.1** — Intelligence-driven career management plugins for Claude Code and Codex, distributed via self-hosted marketplaces.
```

In the repository structure block, add:

```text
  .agents/plugins/marketplace.json    # Codex marketplace manifest
```

and add `.codex-plugin/plugin.json` below each `.claude-plugin/plugin.json`.

In the Development Commands section, add:

```bash
# Validate Codex compatibility contract
npm test

# Test Codex marketplace discovery locally
codex plugin marketplace add ./
codex plugin list
```

- [ ] **Step 8: Add platform compatibility contract to architecture docs**

In `docs/ARCHITECTURE.md`, add this section after the plugin overview:

```markdown
## Platform compatibility

JobOps is distributed to both Claude Code and Codex from the same canonical plugin directories:

- Claude Code reads `.claude-plugin/marketplace.json` and each plugin's `.claude-plugin/plugin.json`.
- Codex reads `.agents/plugins/marketplace.json` and each plugin's `.codex-plugin/plugin.json`.
- Both platforms read the same `skills/` directories. Every skill frontmatter must include `name`, `description`, and `disable-model-invocation: true`.
- Keep skill bodies platform-neutral where practical. When a workflow needs platform-specific orchestration, describe the smallest mapping inline instead of creating a duplicate skill tree.
- Claude Code `Task tool` instructions map conceptually to Codex subagent spawning. Do not rewrite an entire skill solely to rename the orchestration primitive.
- `${CLAUDE_PLUGIN_ROOT}` remains a Claude-specific variable. Existing uses are tolerated for this compatibility pass; subsequent template/style changes should introduce a small cross-platform path-resolution pattern rather than duplicating skills.

Run `npm test` before committing any manifest, marketplace, or skill-frontmatter change.
```

- [ ] **Step 9: Add CHANGELOG entry**

At the top of `CHANGELOG.md`, add an `Unreleased` entry if one does not already exist:

```markdown
## [Unreleased]

### Added

- **Codex plugin compatibility** — added a Codex repo marketplace at `.agents/plugins/marketplace.json`, Codex manifests for `jobops` and `jobops-ic`, and explicit `name` frontmatter on every skill so Codex can discover and install the existing plugin trees without a separate Codex copy. Added `scripts/validate/validate-codex-plugin-compatibility.js` and `npm test` to guard manifest version drift, marketplace shape, and skill metadata.
```

If `## [Unreleased]` already exists, add only the bullet under its `### Added` section.

- [ ] **Step 10: Run documentation sanity checks**

Run:

```bash
rg -n "Claude Code and Codex|codex plugin marketplace add|\\.agents/plugins|\\.codex-plugin|validate-codex-plugin-compatibility" README.md SETUP.md CLAUDE.md docs/ARCHITECTURE.md CHANGELOG.md scripts/validate/README.md
npm test
```

Expected:

```text
Codex compatibility validation passed.
```

- [ ] **Step 11: Commit docs**

Run:

```bash
git add README.md SETUP.md CLAUDE.md docs/ARCHITECTURE.md CHANGELOG.md scripts/validate/README.md package.json
git commit -m "docs: document Claude and Codex plugin installation"
```

If `package.json` was already committed in Task 1, `git add package.json` is harmless and should not stage unrelated changes.

---

### Task 5: Verify Local Codex Marketplace Discovery

**Files:**
- No intended file changes.

This task verifies Codex can see the local repo marketplace and plugin entries.

- [ ] **Step 1: Confirm the CLI is present**

Run:

```bash
codex --version
```

Expected: output begins with:

```text
codex-cli
```

If `codex` is not available, stop this task and record that Codex CLI smoke testing could not run.

- [ ] **Step 2: Add the local marketplace**

Run:

```bash
codex plugin marketplace add ./
```

Expected: output confirms the marketplace was added. If it reports the marketplace is already configured, continue.

- [ ] **Step 3: Confirm marketplace listing**

Run:

```bash
codex plugin marketplace list
```

Expected: output includes:

```text
jobops-marketplace
```

- [ ] **Step 4: Confirm plugin listing**

Run:

```bash
codex plugin list | rg "jobops|jobops-ic"
```

Expected: output includes both plugin selectors:

```text
jobops@jobops-marketplace
jobops-ic@jobops-marketplace
```

- [ ] **Step 5: Install the plugins from the local marketplace**

Run:

```bash
codex plugin add jobops@jobops-marketplace
codex plugin add jobops-ic@jobops-marketplace
```

Expected: both commands exit `0`. If the plugins are already installed, the commands may report that state; continue if they exit `0`.

- [ ] **Step 6: Confirm installed status**

Run:

```bash
codex plugin list | rg "jobops|jobops-ic"
```

Expected: both rows show installed and enabled status for the `jobops-marketplace` entries.

- [ ] **Step 7: Commit no files**

Run:

```bash
git status --short
```

Expected: no new workspace files from the Codex smoke test. Do not commit changes from this task because Codex marketplace installation updates user-level Codex config/cache, not repository source.

---

### Task 6: Verify Claude Compatibility Still Holds

**Files:**
- No intended file changes.

Run this task only when the `claude` command is installed.

- [ ] **Step 1: Check for Claude Code**

Run:

```bash
command -v claude
```

Expected if installed: prints a path and exits `0`. If it exits nonzero, skip the remaining steps in this task and record that Claude validation could not run locally.

- [ ] **Step 2: Validate the core plugin**

Run:

```bash
claude plugin validate plugins/jobops
```

Expected: exits `0` with no validation errors.

- [ ] **Step 3: Validate the IC plugin**

Run:

```bash
claude plugin validate plugins/jobops-ic
```

Expected: exits `0` with no validation errors.

- [ ] **Step 4: Validate the Claude marketplace**

Run:

```bash
claude plugin validate .
```

Expected: exits `0` with no validation errors.

---

### Task 7: Final Package Verification

**Files:**
- No intended file changes unless a previous verification step exposes an issue.

- [ ] **Step 1: Run the static validator through npm**

Run:

```bash
npm test
```

Expected:

```text
Codex compatibility validation passed.
```

- [ ] **Step 2: Validate all JSON files touched by this plan**

Run:

```bash
python3 -m json.tool package.json >/dev/null
python3 -m json.tool .agents/plugins/marketplace.json >/dev/null
python3 -m json.tool .claude-plugin/marketplace.json >/dev/null
python3 -m json.tool plugins/jobops/.claude-plugin/plugin.json >/dev/null
python3 -m json.tool plugins/jobops/.codex-plugin/plugin.json >/dev/null
python3 -m json.tool plugins/jobops-ic/.claude-plugin/plugin.json >/dev/null
python3 -m json.tool plugins/jobops-ic/.codex-plugin/plugin.json >/dev/null
```

Expected: all commands exit `0` and print nothing.

- [ ] **Step 3: Check for accidental duplicate plugin trees**

Run:

```bash
find . -maxdepth 3 -type d \( -name '*codex*' -o -name '*Codex*' \) -print
```

Expected output includes only intentional Codex metadata paths:

```text
./plugins/jobops/.codex-plugin
./plugins/jobops-ic/.codex-plugin
```

If `.agents` does not appear because `-name '*codex*'` does not match it, that is expected.

- [ ] **Step 4: Review the diff**

Run:

```bash
git status --short
git log --oneline -5
```

Expected:

- `git status --short` is empty.
- Recent commits include:
  - `test: define Codex plugin compatibility contract`
  - `feat: add Codex plugin manifests`
  - `feat: add Codex skill names`
  - `docs: document Claude and Codex plugin installation`

- [ ] **Step 5: Prepare final implementation summary**

Report:

```text
Implemented thin dual-manifest Codex compatibility.
Verification run:
- npm test
- JSON syntax checks
- Codex marketplace listing and plugin add smoke test
- Claude plugin validation, if claude was installed
```

Include any command that could not run and the reason.
