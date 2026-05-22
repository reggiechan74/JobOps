# Canonical ResumeSourceFolder Layout

This file is the contract that `/jobops:audit-source` enforces and that downstream
skills (resume builder, cover letter, pitch deck, etc.) rely on. The path to the
source root is `config.directories.resume_source` — usually `./ResumeSourceFolder/`
but configurable via `/jobops:setup`.

## Required structure

```
<resume_source>/
  Identity/
    Name.md                  REQUIRED
    CurrentRole.md           REQUIRED
    Contact.md               optional
  WorkHistory/
    NN_company_role.md       REQUIRED (>=1 file, reverse chronological)
  Technology/
    TechStack.md             REQUIRED
    Certifications.md        REQUIRED (may say "None")
    DomainExpertise.md       optional
  Achievements/              optional (NN_topic.md grouping)
  Publications/              optional (NN_title.md grouping)
  Preferences/
    Vision.md                REQUIRED
    AntiVision.md            REQUIRED (may say "None")
    Compensation.md          optional
    WorkArrangement.md       optional
```

## Per-file required content

### Identity/Name.md
First non-blank line is the candidate's full name as a heading.

### Identity/CurrentRole.md
Must contain: a job title (level-1 or level-2 heading), a company name, and a
start_date in `YYYY-MM` format. If candidate is unemployed, file should say so
explicitly with a `Status: Unemployed since <YYYY-MM>` line.

### WorkHistory/NN_company_role.md
- Filename prefix `NN_` is a 2-digit reverse-chronological order key.
- Must contain (in any order, anywhere in the file):
  - A level-1 heading with `<Company> — <Title>` or equivalent
  - `Start: YYYY-MM`
  - `End: YYYY-MM` or `End: Present`
  - At least one `## Responsibilities` or `## Achievements` section header
- Each numeric claim (`%`, `$`, count of N, `Xx faster`) should have nearby
  context (within 5 lines) describing what it measures.

### Technology/TechStack.md
A flat enumerable structure (bullets, table rows, or sub-headings) listing
skills the candidate has used. Optionally grouped by category. Every skill
that appears in any WorkHistory file MUST also appear here. The reverse
(skills here not in any WorkHistory file) is allowed and means "skill exists,
no role-specific context."

### Technology/Certifications.md
Each cert as a bullet or sub-section with:
- Name
- Issuer
- Date obtained (YYYY-MM)
- Status (Active / Expired / In-Progress)

Example: `- AWS Solutions Architect Associate, AWS, 2022-06, Active`

If the candidate holds no certifications, file should contain a single line:
`None`

### Preferences/Vision.md and AntiVision.md
Free-form prose. No structural requirements beyond "non-empty or explicitly
states None". These files are NEVER parsed into enums by automation; downstream
skills read them as prose in context.

## Layout violations

Auditsource flags as **blocking** any of:
- A REQUIRED file is missing
- A WorkHistory file lacks Start or End date
- A WorkHistory file has neither Responsibilities nor Achievements section
- A Certifications.md entry has no issuer OR no date

Auditsource flags as **advisory** any of:
- A numeric claim lacks nearby context
- A skill appears in WorkHistory but not TechStack
- A WorkHistory file's End date precedes a more-recent role's Start date (timeline conflict)
- A WorkHistory file's start_date leaves a gap > 30 days from the previous role's end_date

## Migration from non-canonical layout

When a user runs `/jobops:audit-source --migrate-layout`, the skill walks the
current source folder file-by-file and proposes a target path in this canonical
layout, requiring confirmation per file. Never moves silently.
