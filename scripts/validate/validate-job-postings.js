#!/usr/bin/env node
/**
 * Job Posting Validator
 *
 * Validates job postings in Job_Postings/ directory for:
 * - YAML front matter presence and completeness
 * - Filename convention compliance (Company_Role_YYYY-MM-DD.md)
 * - Required metadata fields
 *
 * Usage: node scripts/validate/validate-job-postings.js [file-or-directory]
 */

const fs = require('fs');
const path = require('path');

// Required YAML fields for job postings
const REQUIRED_FIELDS = [
  'company',
  'role',
  'location',
  'posting_date',
  'source_url',
  'date_saved',
  'status'
];

// Optional but recommended fields
const RECOMMENDED_FIELDS = [
  'salary_range',
  'employment_type',
  'remote_policy',
  'department',
  'seniority_level'
];

// Valid status values
const VALID_STATUSES = ['active', 'archived', 'applied', 'rejected', 'interview', 'offer'];

/**
 * Extract YAML front matter from markdown file
 */
function extractYAML(content) {
  const yamlRegex = /^---\s*\n([\s\S]*?)\n---/;
  const match = content.match(yamlRegex);

  if (!match) {
    return null;
  }

  // Simple YAML parser (for basic key: value pairs)
  const yaml = {};
  const lines = match[1].split('\n');

  for (const line of lines) {
    const colonIndex = line.indexOf(':');
    if (colonIndex > 0) {
      const key = line.substring(0, colonIndex).trim();
      let value = line.substring(colonIndex + 1).trim();

      // Remove comments
      const commentIndex = value.indexOf('#');
      if (commentIndex >= 0) {
        value = value.substring(0, commentIndex).trim();
      }

      yaml[key] = value;
    }
  }

  return yaml;
}

/**
 * Validate filename convention: Company_Role_YYYY-MM-DD.md
 */
function validateFilename(filename) {
  const errors = [];

  // Skip template and example files
  if (filename.startsWith('.') || filename.startsWith('Example_')) {
    return { valid: true, errors: [], skipped: true };
  }

  // Check .md extension
  if (!filename.endsWith('.md')) {
    errors.push(`File must have .md extension`);
    return { valid: false, errors };
  }

  // Extract parts
  const nameWithoutExt = filename.replace('.md', '');
  const parts = nameWithoutExt.split('_');

  // Should have at least 3 parts: Company, Role (may be multi-word), Date
  if (parts.length < 3) {
    errors.push(`Filename should follow pattern: Company_Role_YYYY-MM-DD.md`);
    return { valid: false, errors };
  }

  // Check if last part is a date (YYYY-MM-DD)
  const datePart = parts[parts.length - 1];
  const dateRegex = /^\d{4}-\d{2}-\d{2}$/;

  if (!dateRegex.test(datePart)) {
    errors.push(`Date should be in format YYYY-MM-DD, got: ${datePart}`);
  }

  return { valid: errors.length === 0, errors };
}

/**
 * Validate YAML metadata
 */
function validateYAML(yaml, filename) {
  const errors = [];
  const warnings = [];

  if (!yaml) {
    errors.push('Missing YAML front matter (should start with --- and end with ---)');
    return { valid: false, errors, warnings };
  }

  // Check required fields
  for (const field of REQUIRED_FIELDS) {
    if (!yaml[field] || yaml[field] === '') {
      errors.push(`Missing required field: ${field}`);
    }
  }

  // Check recommended fields
  for (const field of RECOMMENDED_FIELDS) {
    if (!yaml[field] || yaml[field] === '') {
      warnings.push(`Missing recommended field: ${field}`);
    }
  }

  // Validate status value
  if (yaml.status && !VALID_STATUSES.includes(yaml.status)) {
    errors.push(`Invalid status: ${yaml.status}. Must be one of: ${VALID_STATUSES.join(', ')}`);
  }

  // Validate date formats
  const dateFields = ['posting_date', 'date_saved'];
  const dateRegex = /^\d{4}-\d{2}-\d{2}$/;

  for (const field of dateFields) {
    if (yaml[field] && !dateRegex.test(yaml[field])) {
      errors.push(`${field} should be in format YYYY-MM-DD, got: ${yaml[field]}`);
    }
  }

  return { valid: errors.length === 0, errors, warnings };
}

/**
 * Validate a single job posting file
 */
function validateFile(filePath) {
  const filename = path.basename(filePath);
  const result = {
    file: filename,
    path: filePath,
    valid: true,
    errors: [],
    warnings: [],
    skipped: false
  };

  // Validate filename
  const filenameValidation = validateFilename(filename);
  if (filenameValidation.skipped) {
    result.skipped = true;
    return result;
  }

  if (!filenameValidation.valid) {
    result.valid = false;
    result.errors.push(...filenameValidation.errors.map(e => `[Filename] ${e}`));
  }

  // Read file content
  let content;
  try {
    content = fs.readFileSync(filePath, 'utf8');
  } catch (err) {
    result.valid = false;
    result.errors.push(`Failed to read file: ${err.message}`);
    return result;
  }

  // Extract and validate YAML
  const yaml = extractYAML(content);
  const yamlValidation = validateYAML(yaml, filename);

  if (!yamlValidation.valid) {
    result.valid = false;
  }

  result.errors.push(...yamlValidation.errors.map(e => `[YAML] ${e}`));
  result.warnings.push(...yamlValidation.warnings.map(w => `[YAML] ${w}`));

  return result;
}

/**
 * Validate all job postings in a directory
 */
function validateDirectory(dirPath) {
  const files = fs.readdirSync(dirPath);
  const results = [];

  for (const file of files) {
    const filePath = path.join(dirPath, file);
    const stat = fs.statSync(filePath);

    if (stat.isFile() && file.endsWith('.md')) {
      results.push(validateFile(filePath));
    }
  }

  return results;
}

/**
 * Print validation results
 */
function printResults(results) {
  let totalFiles = 0;
  let validFiles = 0;
  let skippedFiles = 0;
  let filesWithErrors = 0;
  let filesWithWarnings = 0;

  console.log('\n=== Job Posting Validation Report ===\n');

  for (const result of results) {
    totalFiles++;

    if (result.skipped) {
      skippedFiles++;
      console.log(`⊘ SKIPPED: ${result.file} (template or example)`);
      continue;
    }

    if (result.valid && result.warnings.length === 0) {
      validFiles++;
      console.log(`✓ VALID: ${result.file}`);
    } else if (result.valid && result.warnings.length > 0) {
      validFiles++;
      filesWithWarnings++;
      console.log(`⚠ WARNING: ${result.file}`);
      result.warnings.forEach(w => console.log(`    ${w}`));
    } else {
      filesWithErrors++;
      console.log(`✗ ERROR: ${result.file}`);
      result.errors.forEach(e => console.log(`    ${e}`));
      if (result.warnings.length > 0) {
        result.warnings.forEach(w => console.log(`    ${w}`));
      }
    }
  }

  console.log('\n=== Summary ===');
  console.log(`Total files: ${totalFiles}`);
  console.log(`Valid: ${validFiles}`);
  console.log(`Errors: ${filesWithErrors}`);
  console.log(`Warnings: ${filesWithWarnings}`);
  console.log(`Skipped: ${skippedFiles}`);

  return filesWithErrors === 0;
}

/**
 * Main execution
 */
function main() {
  const args = process.argv.slice(2);
  const target = args[0] || 'Job_Postings';

  let results = [];

  if (fs.existsSync(target)) {
    const stat = fs.statSync(target);

    if (stat.isDirectory()) {
      results = validateDirectory(target);
    } else if (stat.isFile()) {
      results = [validateFile(target)];
    }
  } else {
    console.error(`Error: Path not found: ${target}`);
    process.exit(1);
  }

  const success = printResults(results);
  process.exit(success ? 0 : 1);
}

if (require.main === module) {
  main();
}

module.exports = { validateFile, validateDirectory, extractYAML, validateYAML, validateFilename };
