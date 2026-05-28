#!/usr/bin/env node
/**
 * Validate that assessment skills block deprecated candidate profile JSON files.
 *
 * Usage: node scripts/validate/validate-legacy-profile-block.js
 */

const fs = require('fs');
const path = require('path');

const REPO_ROOT = path.resolve(__dirname, '..', '..');

const SKILL_FILES = [
  'plugins/jobops/skills/assessjob/SKILL.md',
  'plugins/jobops/skills/assesscandidate/SKILL.md'
];

const REQUIRED_PATTERNS = [
  {
    label: 'legacy artifact check heading',
    pattern: /### Legacy Candidate Profile Artifact Check/
  },
  {
    label: 'nested legacy profile path',
    pattern: /`?\{source_path\}\/\.profile\/candidate_profile\.json`?/
  },
  {
    label: 'root legacy profile path',
    pattern: /`?\{source_path\}\/candidate_profile\.json`?/
  },
  {
    label: 'single-file candidate_profile.json guard',
    pattern: /single-file source[\s\S]{0,240}candidate_profile\.json/i
  },
  {
    label: 'halt before assessment',
    pattern: /stop[\s\S]{0,160}(before|without)[\s\S]{0,160}assessment/i
  },
  {
    label: 'direct source markdown explanation',
    pattern: /read source markdown directly/i
  },
  {
    label: 'delete instruction',
    pattern: /delete[\s\S]{0,120}(legacy|deprecated)[\s\S]{0,120}(file|artifact)/i
  },
  {
    label: 'delete offer',
    pattern: /offer[\s\S]{0,120}delete/i
  }
];

const SECTION_BOUNDARY = '### 2.1 Read Candidate Source';

let failures = 0;

for (const relativePath of SKILL_FILES) {
  const filePath = path.join(REPO_ROOT, relativePath);
  const content = fs.readFileSync(filePath, 'utf8');
  const boundaryIndex = content.indexOf(SECTION_BOUNDARY);
  const headingIndex = content.indexOf('### Legacy Candidate Profile Artifact Check');

  console.log(`Checking ${relativePath}`);

  for (const requirement of REQUIRED_PATTERNS) {
    if (!requirement.pattern.test(content)) {
      failures += 1;
      console.error(`  FAIL: missing ${requirement.label}`);
    }
  }

  if (headingIndex === -1 || boundaryIndex === -1 || headingIndex > boundaryIndex) {
    failures += 1;
    console.error('  FAIL: legacy profile check must appear before candidate source reads');
  }
}

if (failures > 0) {
  console.error(`\nLegacy profile block validation failed with ${failures} issue(s).`);
  process.exit(1);
}

console.log('\nLegacy profile block validation passed.');
