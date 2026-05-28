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

function validateCodexComponentPathField(pluginName, field, value, pluginRoot, expectedBasename) {
  validateCodexManifestPathField(pluginName, field, value, pluginRoot);
  if (value === undefined || typeof value !== 'string') {
    return;
  }
  check(
    path.basename(value) === expectedBasename,
    `${pluginName} ${field} path basename must be ${expectedBasename}`
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
  validateCodexComponentPathField(plugin.name, 'apps', codex.apps, plugin.root, '.app.json');
  validateCodexComponentPathField(plugin.name, 'mcpServers', codex.mcpServers, plugin.root, '.mcp.json');
  validateCodexManifestPathField(plugin.name, 'hooks', codex.hooks, plugin.root);
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

  check(
    typeof packageJson.version === 'string' && packageJson.version.trim() !== '',
    'package.json version must be present'
  );
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
