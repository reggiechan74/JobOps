---
description: Perform provenance check on a draft resume against comprehensive work histories
disable-model-invocation: true
---

## Configuration

Read `.jobops/config.json`. If not found, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories` for all file paths in this skill.
Use `config.templates.active` to resolve template locations — for each template needed,
read from: `{config.templates.base_dir}/{active_value}/{filename}`

I'll perform a comprehensive provenance analysis on the selected draft resume against your master work histories in the {config.directories.resume_source}.

**Analyzing**: `$1`

## Step 1: Loading Source Materials

I'll first load all your comprehensive work histories from the {config.directories.resume_source} to use as the source of truth:

**Source Documents**: All files in `{config.directories.resume_source}/`

## Step 2: Reading Draft Resume

Now reading the draft resume to analyze:

@$1

## Step 3: Running Provenance Analysis

I'm launching the step2-provenance-check agent to perform a comprehensive provenance analysis. This will:

### Analysis Scope:
- **Evidence Verification**: Cross-reference every claim against master resume documentation
- **Quantitative Accuracy**: Verify all metrics, numbers, and percentages
- **Timeline Consistency**: Check date ranges and chronological accuracy
- **Technical Claims**: Validate technology and skill assertions against documented experience
- **Geographic Scope**: Verify location and jurisdiction claims
- **Team/Budget Claims**: Confirm management scope and financial figures

### Risk Assessment Categories:
- **High Risk**: Claims that could be easily challenged or verified as false
- **Medium Risk**: Assertions lacking sufficient evidence or specificity
- **Low Risk**: Minor issues or confidentiality concerns

### Deliverables:
1. **Detailed Risk Report**: Categorized list of all credibility issues
2. **Evidence Mapping**: Line-by-line verification against source materials
3. **Recommendations**: Specific fixes for each identified issue
4. **Defensibility Score**: Overall assessment of resume credibility

The analysis will use the CV Provenance Hardening methodology to ensure your resume can withstand scrutiny during:
- Background checks
- Reference verification
- Technical interviews
- Behavioral interviews

## Output

The provenance analysis will be saved with detailed findings including:
- Specific line references to issues in the draft
- Cross-references to supporting evidence in master documents
- Actionable recommendations for each credibility concern
- Priority ranking for addressing issues

Start the report with:

```yaml
---
draft_file: $1
generated_by: /provenance
generated_on: <ISO8601 timestamp>
output_type: resume_provenance
status: analysis
version: 1.0
---
```

Update timestamps and `version` when re-running.

This analysis ensures your resume maintains competitive positioning while being completely defensible under scrutiny.
