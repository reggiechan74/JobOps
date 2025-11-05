# Metadata standards

All resume source files use YAML front matter for automation and Obsidian queries.

## Required fields
- `title`: Short name for the document (e.g., "Hydro One Networks Inc. – Senior Project Consultant").
- `organization`: Employer or institution name. Use `personal` for individual achievements.
- `tenure`: Date range in `YYYY-MM` format (`2025-02 – Present`).
- `roleLevel`: Career level or classification (`Executive`, `Director`, `Consultant`, `Certification`).
- `keywords`: Comma-separated tags supporting search and filtering.
- `sourceDoc`: Original document reference (e.g., `Comprehensive CV MASTER COPY v36`).
- `lastUpdated`: ISO date indicating the latest substantive change.
- `techRefs`: List of relevant anchors in the technology matrix (e.g., `[AI-ML, Documentation-Workflow]`).
- `shareable`: `true` or `false` to signal whether the content can be used in external resumes without editing.

## Optional fields
- `roleIds`: Stable identifiers for linking with tooling or workflows.
- `location`: City, province/state, and country.
- `mandate`: One-sentence summary of the role or highlight.
- `artefacts`: Inline list of supporting notes or vault paths.

## Front matter example
```yaml
---
title: Hydro One Networks Inc. – Senior Project Consultant
organization: Hydro One Networks Inc.
tenure: 2025-02 – Present
roleLevel: Consultant
keywords: infrastructure,land-acquisition,ai-workflows
sourceDoc: Comprehensive CV MASTER COPY v36
lastUpdated: 2025-10-02
techRefs:
  - AI-ML
  - Documentation-Workflow
shareable: false
roleIds:
  - honi-consultant-2025
location: Toronto, ON, Canada
mandate: Support Hydro One with land acquisition process modernization and cost optimization.
---
```

## Usage notes
- Place front matter at the top of every Markdown file before the first heading.
- Keep field names lowercase and hyphenated when necessary for Obsidian Dataview compatibility.
- Update `lastUpdated` whenever new material is added.
- Preserve `sourceDoc` and include references in `artefacts` when content traces back to external evidence.
