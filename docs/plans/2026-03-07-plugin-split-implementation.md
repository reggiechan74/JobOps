# JobOps Plugin Split Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Split JobOps into two Claude Code plugins (`jobops` + `jobops-ic`) distributed via a self-hosted marketplace, with a configuration system and template management.

**Architecture:** Commands become skills (with `disable-model-invocation: true`) in flat plugin directories. Templates are bundled in plugins and copied to workspace via setup scripts. A `SessionStart` hook writes `${CLAUDE_PLUGIN_ROOT}` to `/tmp/` for path resolution. Configuration lives in `.jobops/config.json`.

**Tech Stack:** Claude Code plugin system, shell scripts, JSON config, markdown skills/agents

**Design Doc:** `docs/plans/2026-03-07-plugin-split-design.md`

---

### Task 1: Create feature branch

**Files:**
- None (git operation only)

**Step 1: Create and checkout feature branch**

```bash
git checkout -b feature/plugin-split
```

**Step 2: Verify branch**

```bash
git branch --show-current
```

Expected: `feature/plugin-split`

**Step 3: Push branch to remote**

```bash
git push -u origin feature/plugin-split
```

---

### Task 2: Scaffold plugin directory structure

**Files:**
- Create: `plugins/jobops/.claude-plugin/` (directory)
- Create: `plugins/jobops/hooks/` (directory)
- Create: `plugins/jobops/scripts/` (directory)
- Create: `plugins/jobops/templates/` (directory)
- Create: `plugins/jobops/skills/` (directory)
- Create: `plugins/jobops/agents/` (directory)
- Create: `plugins/jobops-ic/.claude-plugin/` (directory)
- Create: `plugins/jobops-ic/hooks/` (directory)
- Create: `plugins/jobops-ic/scripts/` (directory)
- Create: `plugins/jobops-ic/templates/` (directory)
- Create: `plugins/jobops-ic/skills/` (directory)
- Create: `plugins/jobops-ic/agents/` (directory)
- Create: `.claude-plugin/` (directory, for marketplace)

**Step 1: Create all directories**

```bash
mkdir -p plugins/jobops/{.claude-plugin,hooks,scripts,templates,skills,agents}
mkdir -p plugins/jobops-ic/{.claude-plugin,hooks,scripts,templates,skills,agents}
mkdir -p .claude-plugin
```

**Step 2: Create skill subdirectories for jobops**

One directory per skill. Each will contain a `SKILL.md`.

```bash
cd plugins/jobops/skills
mkdir setup buildresume provenance-check coverletter \
      assessjob assesscandidate createrubric comparejobs \
      briefing interviewprep osint auditjobposting \
      idealjob change-one-thing assess-job-offer \
      code-red severance-review workplace-documentation \
      non-compete-analysis reference-shield unemployment-prep \
      discrimination-assessment investigation-prep accommodation-request \
      layoff-intel constructive-dismissal \
      convert-to-pdf convert-to-word markdown-to-pdf \
      install-pandoc github-portfolio
cd ../../..
```

**Step 3: Create skill subdirectories for jobops-ic**

```bash
cd plugins/jobops-ic/skills
mkdir setup defineservices findclient pitchdeck \
      proposaltemplate ratecard copywrite copywriting-spec \
      create-landing-page css-template
cd ../../..
```

**Step 4: Verify structure**

```bash
find plugins/ -type d | head -60
```

Expected: All directories listed above exist.

**Step 5: Commit**

```bash
git add plugins/ .claude-plugin/
git commit -m "Scaffold plugin directory structure for jobops and jobops-ic"
```

---

### Task 3: Create plugin manifests and marketplace config

**Files:**
- Create: `plugins/jobops/.claude-plugin/plugin.json`
- Create: `plugins/jobops-ic/.claude-plugin/plugin.json`
- Create: `.claude-plugin/marketplace.json`

**Step 1: Create jobops plugin.json**

Write to `plugins/jobops/.claude-plugin/plugin.json`:

```json
{
  "name": "jobops",
  "description": "Intelligence-driven job application system - resume development, interview prep, OSINT intelligence, career strategy, and crisis management using HAM-Z methodology",
  "version": "2.0.0",
  "author": {
    "name": "Reggie Chan"
  },
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": ["resume", "job-search", "interview-prep", "osint", "career", "assessment"]
}
```

Note: Version bumped to 2.0.0 — this is a breaking change (plugin architecture replaces `.claude/` layout).

**Step 2: Create jobops-ic plugin.json**

Write to `plugins/jobops-ic/.claude-plugin/plugin.json`:

```json
{
  "name": "jobops-ic",
  "description": "Independent contractor toolkit - service definitions, client prospecting, pitch decks, proposals, rate cards, and landing pages. Requires jobops plugin.",
  "version": "2.0.0",
  "author": {
    "name": "Reggie Chan"
  },
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": ["independent-contractor", "consulting", "proposals", "rate-card", "pitch-deck", "landing-page"]
}
```

**Step 3: Create marketplace.json**

Write to `.claude-plugin/marketplace.json`:

```json
{
  "name": "jobops-marketplace",
  "owner": {
    "name": "Reggie Chan"
  },
  "metadata": {
    "description": "Intelligence-driven career management plugins for Claude Code",
    "version": "1.0.0"
  },
  "plugins": [
    {
      "name": "jobops",
      "source": "./plugins/jobops",
      "description": "Resume development, interview prep, OSINT, career strategy, and crisis management"
    },
    {
      "name": "jobops-ic",
      "source": "./plugins/jobops-ic",
      "description": "Independent contractor toolkit - services, clients, pitches, proposals, rate cards. Requires jobops."
    }
  ]
}
```

**Step 4: Commit**

```bash
git add plugins/jobops/.claude-plugin/plugin.json \
       plugins/jobops-ic/.claude-plugin/plugin.json \
       .claude-plugin/marketplace.json
git commit -m "Add plugin manifests and marketplace configuration"
```

---

### Task 4: Create hooks and scripts for both plugins

**Files:**
- Create: `plugins/jobops/hooks/hooks.json`
- Create: `plugins/jobops/scripts/copy-templates.sh`
- Create: `plugins/jobops-ic/hooks/hooks.json`
- Create: `plugins/jobops-ic/scripts/copy-templates.sh`

**Step 1: Create jobops hooks.json**

Write to `plugins/jobops/hooks/hooks.json`:

```json
{
  "hooks": {
    "SessionStart": [
      {
        "hooks": [
          {
            "type": "command",
            "command": "echo ${CLAUDE_PLUGIN_ROOT} > /tmp/.jobops-plugin-root"
          }
        ]
      }
    ]
  }
}
```

**Step 2: Create jobops copy-templates.sh**

Write to `plugins/jobops/scripts/copy-templates.sh`:

```bash
#!/bin/bash
set -euo pipefail

PLUGIN_ROOT_FILE="/tmp/.jobops-plugin-root"
if [ ! -f "$PLUGIN_ROOT_FILE" ]; then
  echo "ERROR: Plugin root not found. Restart Claude Code to initialize." >&2
  exit 1
fi

PLUGIN_ROOT="$(cat "$PLUGIN_ROOT_FILE" | tr -d '[:space:]')"
SOURCE_DIR="$PLUGIN_ROOT/templates"
TARGET_DIR="${1:-.jobops/templates/default}"

if [ ! -d "$SOURCE_DIR" ]; then
  echo "ERROR: Template source not found at $SOURCE_DIR" >&2
  exit 1
fi

mkdir -p "$TARGET_DIR"
cp "$SOURCE_DIR"/* "$TARGET_DIR/"
COUNT=$(ls -1 "$TARGET_DIR" | wc -l)
echo "Copied $COUNT templates to $TARGET_DIR"
```

**Step 3: Make script executable**

```bash
chmod +x plugins/jobops/scripts/copy-templates.sh
```

**Step 4: Create jobops-ic hooks.json**

Write to `plugins/jobops-ic/hooks/hooks.json`:

```json
{
  "hooks": {
    "SessionStart": [
      {
        "hooks": [
          {
            "type": "command",
            "command": "echo ${CLAUDE_PLUGIN_ROOT} > /tmp/.jobops-ic-plugin-root"
          }
        ]
      }
    ]
  }
}
```

**Step 5: Create jobops-ic copy-templates.sh**

Write to `plugins/jobops-ic/scripts/copy-templates.sh`:

```bash
#!/bin/bash
set -euo pipefail

PLUGIN_ROOT_FILE="/tmp/.jobops-ic-plugin-root"
if [ ! -f "$PLUGIN_ROOT_FILE" ]; then
  echo "ERROR: Plugin root not found. Restart Claude Code to initialize." >&2
  exit 1
fi

PLUGIN_ROOT="$(cat "$PLUGIN_ROOT_FILE" | tr -d '[:space:]')"
SOURCE_DIR="$PLUGIN_ROOT/templates"
TARGET_DIR="${1:-.jobops/templates/default}"

if [ ! -d "$SOURCE_DIR" ]; then
  echo "ERROR: Template source not found at $SOURCE_DIR" >&2
  exit 1
fi

mkdir -p "$TARGET_DIR"
cp "$SOURCE_DIR"/* "$TARGET_DIR/"
COUNT=$(ls -1 "$TARGET_DIR" | wc -l)
echo "Copied $COUNT templates to $TARGET_DIR"
```

**Step 6: Make script executable**

```bash
chmod +x plugins/jobops-ic/scripts/copy-templates.sh
```

**Step 7: Commit**

```bash
git add plugins/jobops/hooks/ plugins/jobops/scripts/ \
       plugins/jobops-ic/hooks/ plugins/jobops-ic/scripts/
git commit -m "Add SessionStart hooks and template copy scripts for both plugins"
```

---

### Task 5: Copy templates into plugins

**Files:**
- Copy: `.claude/templates/assessment_rubric_framework.md` -> `plugins/jobops/templates/`
- Copy: `.claude/templates/evidence_verification_framework.md` -> `plugins/jobops/templates/`
- Copy: `.claude/templates/assessment_report_structure.md` -> `plugins/jobops/templates/`
- Copy: `.claude/templates/candidate_profile_schema.json` -> `plugins/jobops/templates/`
- Copy: `.claude/templates/service_definition_schema.json` -> `plugins/jobops-ic/templates/`

**Step 1: Copy jobops templates**

```bash
cp .claude/templates/assessment_rubric_framework.md plugins/jobops/templates/
cp .claude/templates/evidence_verification_framework.md plugins/jobops/templates/
cp .claude/templates/assessment_report_structure.md plugins/jobops/templates/
cp .claude/templates/candidate_profile_schema.json plugins/jobops/templates/
```

**Step 2: Copy jobops-ic templates**

```bash
cp .claude/templates/service_definition_schema.json plugins/jobops-ic/templates/
```

**Step 3: Verify**

```bash
ls plugins/jobops/templates/
ls plugins/jobops-ic/templates/
```

Expected: 4 files in jobops, 1 file in jobops-ic.

**Step 4: Commit**

```bash
git add plugins/jobops/templates/ plugins/jobops-ic/templates/
git commit -m "Copy templates into plugin directories"
```

---

### Task 6: Move jobops agents (flatten from subdirectories)

**Files:**
- Move: `.claude/agents/resume-development/*.md` -> `plugins/jobops/agents/`
- Move: `.claude/agents/interview-prep/*.md` -> `plugins/jobops/agents/`
- Move: `.claude/agents/osint/*.md` -> `plugins/jobops/agents/`
- Move: `.claude/agents/performance/resume-summarizer.md` -> `plugins/jobops/agents/`
- Skip: `.claude/agents/job-search/hiringcafe-search.md` (dropped)

**Step 1: Copy agents to plugin (flatten)**

```bash
cp .claude/agents/resume-development/step1-resume-draft.md plugins/jobops/agents/
cp .claude/agents/resume-development/step2-provenance-check.md plugins/jobops/agents/
cp .claude/agents/resume-development/step3-final-resume.md plugins/jobops/agents/
cp .claude/agents/resume-development/step4-cover-letter.md plugins/jobops/agents/
cp .claude/agents/interview-prep/candidate-assessment.md plugins/jobops/agents/
cp .claude/agents/interview-prep/interview-briefing.md plugins/jobops/agents/
cp .claude/agents/interview-prep/interview-question-generator.md plugins/jobops/agents/
cp .claude/agents/performance/resume-summarizer.md plugins/jobops/agents/
cp .claude/agents/osint/osint-agent.md plugins/jobops/agents/
cp .claude/agents/osint/osint-compensation.md plugins/jobops/agents/
cp .claude/agents/osint/osint-corporate.md plugins/jobops/agents/
cp .claude/agents/osint/osint-culture.md plugins/jobops/agents/
cp .claude/agents/osint/osint-leadership.md plugins/jobops/agents/
cp .claude/agents/osint/osint-legal.md plugins/jobops/agents/
cp .claude/agents/osint/osint-market.md plugins/jobops/agents/
cp .claude/agents/osint/osint-person.md plugins/jobops/agents/
```

**Step 2: Verify count**

```bash
ls -1 plugins/jobops/agents/ | wc -l
```

Expected: 16

**Step 3: Commit**

```bash
git add plugins/jobops/agents/
git commit -m "Move and flatten jobops agents into plugin directory"
```

---

### Task 7: Move jobops-ic agent

**Files:**
- Move: `.claude/agents/landing-page/landing-page-copywriter.md` -> `plugins/jobops-ic/agents/`

**Step 1: Copy agent to plugin**

```bash
cp .claude/agents/landing-page/landing-page-copywriter.md plugins/jobops-ic/agents/
```

**Step 2: Commit**

```bash
git add plugins/jobops-ic/agents/
git commit -m "Move landing-page-copywriter agent into jobops-ic plugin"
```

---

### Task 8: Convert jobops resume-development commands to skills

Each command `.md` file becomes a `SKILL.md` inside a skill directory. The transformation:
1. Replace frontmatter `description` + `argument-hint` with skill frontmatter adding `disable-model-invocation: true`
2. Prepend the standard configuration preamble after frontmatter
3. Replace hardcoded directory paths (`ResumeSourceFolder/`, `OutputResumes/`, etc.) with config references
4. Replace hardcoded template paths (`.claude/templates/`) with config-based template resolution

**Files:**
- Source: `.claude/commands/resume-development/buildresume.md`
- Source: `.claude/commands/resume-development/provenance-check.md`
- Source: `.claude/commands/resume-development/coverletter.md`
- Create: `plugins/jobops/skills/buildresume/SKILL.md`
- Create: `plugins/jobops/skills/provenance-check/SKILL.md`
- Create: `plugins/jobops/skills/coverletter/SKILL.md`

**Step 1: Read each source command file**

Read all three source files to understand their full content.

**Step 2: Transform each command to a skill**

For each file, apply these transformations:

a) **Frontmatter**: Change from:
```yaml
---
description: <text>
argument-hint: <text>
---
```
To:
```yaml
---
description: <text>
disable-model-invocation: true
---
```

b) **Prepend configuration preamble** immediately after frontmatter:
```markdown
## Configuration

Read `.jobops/config.json`. If not found, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories` for all file paths in this skill.
Use `config.templates.active` to resolve template locations — for each template needed,
read from: `{config.templates.base_dir}/{active_value}/{filename}`
```

c) **Replace hardcoded paths**: Search and replace throughout the file body:
- `ResumeSourceFolder/` -> `{config.directories.resume_source}/`
- `OutputResumes/` -> `{config.directories.output_resumes}/`
- `Job_Postings/` -> `{config.directories.job_postings}/`
- `Scoring_Rubrics/` -> `{config.directories.scoring_rubrics}/`
- `Briefing_Notes/` -> `{config.directories.briefing_notes}/`
- `Intelligence_Reports/` -> `{config.directories.intelligence_reports}/`
- `Client_Prospects/` -> `{config.directories.client_prospects}/`
- `.claude/templates/` -> reference via config template resolution

**Step 3: Write each transformed SKILL.md**

Write to:
- `plugins/jobops/skills/buildresume/SKILL.md`
- `plugins/jobops/skills/provenance-check/SKILL.md`
- `plugins/jobops/skills/coverletter/SKILL.md`

**Step 4: Verify skill count**

```bash
ls plugins/jobops/skills/buildresume/SKILL.md plugins/jobops/skills/provenance-check/SKILL.md plugins/jobops/skills/coverletter/SKILL.md
```

**Step 5: Commit**

```bash
git add plugins/jobops/skills/buildresume/ plugins/jobops/skills/provenance-check/ plugins/jobops/skills/coverletter/
git commit -m "Convert resume-development commands to jobops skills"
```

---

### Task 9: Convert jobops interview-prep commands to skills

Same transformation pattern as Task 8.

**Files:**
- Source: `.claude/commands/interview-prep/assessjob.md`
- Source: `.claude/commands/interview-prep/assesscandidate.md`
- Source: `.claude/commands/interview-prep/createrubric.md`
- Source: `.claude/commands/interview-prep/comparejobs.md`
- Source: `.claude/commands/interview-prep/briefing.md`
- Source: `.claude/commands/interview-prep/interviewprep.md`
- Create: `plugins/jobops/skills/{assessjob,assesscandidate,createrubric,comparejobs,briefing,interviewprep}/SKILL.md`

**Step 1: Read all 6 source command files**

**Step 2: Transform each using the pattern from Task 8**

Apply frontmatter change, configuration preamble, and path replacements.

**Step 3: Write each transformed SKILL.md**

**Step 4: Commit**

```bash
git add plugins/jobops/skills/assessjob/ plugins/jobops/skills/assesscandidate/ \
       plugins/jobops/skills/createrubric/ plugins/jobops/skills/comparejobs/ \
       plugins/jobops/skills/briefing/ plugins/jobops/skills/interviewprep/
git commit -m "Convert interview-prep commands to jobops skills"
```

---

### Task 10: Convert jobops job-search commands to skills

Same transformation pattern. Note: `searchjobs` is **dropped** — do not convert it.

**Files:**
- Source: `.claude/commands/job-search/osint.md`
- Source: `.claude/commands/job-search/auditjobposting.md`
- Create: `plugins/jobops/skills/{osint,auditjobposting}/SKILL.md`

**Step 1: Read both source command files**

**Step 2: Transform each using the pattern from Task 8**

**Step 3: Write each transformed SKILL.md**

**Step 4: Commit**

```bash
git add plugins/jobops/skills/osint/ plugins/jobops/skills/auditjobposting/
git commit -m "Convert job-search commands to jobops skills (searchjobs dropped)"
```

---

### Task 11: Convert jobops career-strategy commands to skills

**Files:**
- Source: `.claude/commands/career-strategy/idealjob.md`
- Source: `.claude/commands/career-strategy/change-one-thing.md`
- Source: `.claude/commands/career-strategy/assess-job-offer.md`
- Create: `plugins/jobops/skills/{idealjob,change-one-thing,assess-job-offer}/SKILL.md`

**Step 1-4: Same pattern as Task 8, then commit**

```bash
git add plugins/jobops/skills/idealjob/ plugins/jobops/skills/change-one-thing/ \
       plugins/jobops/skills/assess-job-offer/
git commit -m "Convert career-strategy commands to jobops skills"
```

---

### Task 12: Convert jobops career-crisis commands to skills

**Files:**
- Source: `.claude/commands/career-crisis/*.md` (11 files)
- Create: `plugins/jobops/skills/{code-red,severance-review,workplace-documentation,non-compete-analysis,reference-shield,unemployment-prep,discrimination-assessment,investigation-prep,accommodation-request,layoff-intel,constructive-dismissal}/SKILL.md`

**Step 1: Read all 11 source command files**

**Step 2: Transform each using the pattern from Task 8**

Note: These commands are largely self-contained (no template dependencies), so the path replacements are primarily for `ResumeSourceFolder/` and `OutputResumes/`.

**Step 3: Write each transformed SKILL.md**

**Step 4: Commit**

```bash
git add plugins/jobops/skills/code-red/ plugins/jobops/skills/severance-review/ \
       plugins/jobops/skills/workplace-documentation/ plugins/jobops/skills/non-compete-analysis/ \
       plugins/jobops/skills/reference-shield/ plugins/jobops/skills/unemployment-prep/ \
       plugins/jobops/skills/discrimination-assessment/ plugins/jobops/skills/investigation-prep/ \
       plugins/jobops/skills/accommodation-request/ plugins/jobops/skills/layoff-intel/ \
       plugins/jobops/skills/constructive-dismissal/
git commit -m "Convert career-crisis commands to jobops skills"
```

---

### Task 13: Convert jobops application-finalization and system-setup commands to skills

**Files:**
- Source: `.claude/commands/application-finalization/convert-to-pdf-playwright.md`
- Source: `.claude/commands/application-finalization/convert-to-word.md`
- Source: `.claude/commands/application-finalization/markdown-to-pdf.md`
- Source: `.claude/commands/system-setup/install-pandoc.md`
- Source: `.claude/commands/system-setup/github-portfolio.md`
- Create: `plugins/jobops/skills/{convert-to-pdf,convert-to-word,markdown-to-pdf,install-pandoc,github-portfolio}/SKILL.md`
- Skip: `setup-aliases.md` (dropped)
- Skip: `git-delete.md` (not part of either plugin)
- Note: `create-career-history.md` is absorbed into the setup skill (Task 15)

**Step 1-4: Same pattern as Task 8, then commit**

```bash
git add plugins/jobops/skills/convert-to-pdf/ plugins/jobops/skills/convert-to-word/ \
       plugins/jobops/skills/markdown-to-pdf/ plugins/jobops/skills/install-pandoc/ \
       plugins/jobops/skills/github-portfolio/
git commit -m "Convert finalization and system-setup commands to jobops skills"
```

---

### Task 14: Convert jobops-ic commands to skills

Same transformation pattern, but with the **jobops-ic preamble** which includes the prerequisite check.

**Files:**
- Source: `.claude/commands/independent-contractor/defineservices.md`
- Source: `.claude/commands/independent-contractor/findclient.md`
- Source: `.claude/commands/independent-contractor/pitchdeck.md`
- Source: `.claude/commands/independent-contractor/proposaltemplate.md`
- Source: `.claude/commands/independent-contractor/ratecard.md`
- Source: `.claude/commands/landing-page/copywrite.md`
- Source: `.claude/commands/landing-page/copywriting-spec.md`
- Source: `.claude/commands/landing-page/create.md`
- Source: `.claude/commands/landing-page/css-template.md`
- Create: `plugins/jobops-ic/skills/{defineservices,findclient,pitchdeck,proposaltemplate,ratecard,copywrite,copywriting-spec,create-landing-page,css-template}/SKILL.md`

**Step 1: Read all 9 source command files**

**Step 2: Transform each command to a skill**

Apply same frontmatter and path transformations as Task 8, but **replace the configuration preamble** with the jobops-ic version:

```markdown
## Prerequisite Check

Before executing, verify the `jobops` plugin is installed by checking if the
`resume-summarizer` agent is available. If not found, stop with:

> PREREQUISITE MISSING: jobops plugin
>
> The jobops-ic plugin requires the jobops plugin to be installed first.
> Install it with: `/plugin install jobops@jobops-marketplace`
>
> Then retry this command.

## Configuration

Read `.jobops/config.json`. If not found, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories` for all file paths in this skill.
Use `config.templates.active` to resolve template locations — for each template needed,
read from: `{config.templates.base_dir}/{active_value}/{filename}`
```

**Step 3: Write each transformed SKILL.md**

Note: `create.md` maps to `create-landing-page/SKILL.md` (renamed to avoid ambiguity).

**Step 4: Commit**

```bash
git add plugins/jobops-ic/skills/
git commit -m "Convert independent-contractor and landing-page commands to jobops-ic skills"
```

---

### Task 15: Create /jobops:setup skill

This is a **new skill** that absorbs `create-career-history` functionality and adds directory configuration and template installation.

**Files:**
- Read: `.claude/commands/system-setup/create-career-history.md` (for career history import logic)
- Create: `plugins/jobops/skills/setup/SKILL.md`

**Step 1: Read create-career-history.md for reference**

Read the full file to understand the career history parsing logic to incorporate.

**Step 2: Write the setup SKILL.md**

Write to `plugins/jobops/skills/setup/SKILL.md`:

```yaml
---
description: Initialize JobOps workspace - configure directories, install templates, and optionally import career history
disable-model-invocation: true
---
```

The skill body should implement this flow:

1. **Welcome message** - Explain what setup does
2. **Check for existing config** - If `.jobops/config.json` exists, ask if user wants to reconfigure or skip
3. **Directory Configuration** - For each directory, ask user with sensible defaults:
   - `resume_source` (default: `./ResumeSourceFolder`)
   - `job_postings` (default: `./Job_Postings`)
   - `output_resumes` (default: `./OutputResumes`)
   - `scoring_rubrics` (default: `./Scoring_Rubrics`)
   - `briefing_notes` (default: `./Briefing_Notes`)
   - `intelligence_reports` (default: `./Intelligence_Reports`)
   - `sample_output` (default: `./Sample_Output`)
4. **Create directories** - Create any that don't exist
5. **Template Installation** - Run: `bash "$(cat /tmp/.jobops-plugin-root)/scripts/copy-templates.sh" ".jobops/templates/default"`
6. **Create custom template directory** - `mkdir -p .jobops/templates/custom`
7. **Write config.json** - Write `.jobops/config.json` with user's choices and default template active values
8. **Career History Import** (optional) - Ask if user has existing resume files. If yes, incorporate the parsing logic from `create-career-history.md` (read files, parse into ResumeSourceFolder structure with Experience/, CareerHighlights/, Technology/, Preferences/ subdirectories, apply HAM-Z methodology)
9. **Profile Selection** (optional) - Ask for cultural profile preference (Canadian/American)
10. **Summary** - Display configuration and next steps

**Step 3: Commit**

```bash
git add plugins/jobops/skills/setup/
git commit -m "Create /jobops:setup skill with directory config, template install, and career history import"
```

---

### Task 16: Create /jobops-ic:setup skill

**Files:**
- Create: `plugins/jobops-ic/skills/setup/SKILL.md`

**Step 1: Write the setup SKILL.md**

```yaml
---
description: Initialize JobOps Independent Contractor workspace - extends jobops config with IC directories and templates
disable-model-invocation: true
---
```

The skill body should implement:

1. **Prerequisite check** - Read `.jobops/config.json`. If not found, stop with: "Run `/jobops:setup` first"
2. **IC directory configuration** - Ask for `client_prospects` directory path (default: `./Client_Prospects`)
3. **Create directory** if it doesn't exist
4. **Template installation** - Run: `bash "$(cat /tmp/.jobops-ic-plugin-root)/scripts/copy-templates.sh" ".jobops/templates/default"`
5. **Update config.json** - Read existing config, add:
   - `config.directories.client_prospects` with user's choice
   - `config.templates.active.service_definition_schema` set to `"default"`
6. **Write updated config** back to `.jobops/config.json`
7. **Summary** - Display what was added and next steps

**Step 2: Commit**

```bash
git add plugins/jobops-ic/skills/setup/
git commit -m "Create /jobops-ic:setup skill with IC directory and template configuration"
```

---

### Task 17: Remove old .claude/ directory and dropped files

**Files:**
- Delete: `.claude/commands/` (entire directory)
- Delete: `.claude/agents/` (entire directory)
- Delete: `.claude/templates/` (entire directory)
- Delete: `.mcp.json`
- Delete: `.playwright-mcp/` (entire directory)
- Delete: `scripts/mcp/` (if exists)
- Delete: `AGENTS.md`

**Step 1: Verify nothing else lives in .claude/ that should be kept**

```bash
ls -la .claude/
```

Check for `settings.json`, `settings.local.json`, or other config files that should remain.

**Step 2: Remove old command, agent, and template directories**

```bash
rm -rf .claude/commands/ .claude/agents/ .claude/templates/
```

**Step 3: Remove dropped files**

```bash
rm -f .mcp.json AGENTS.md
rm -rf .playwright-mcp/
rm -rf scripts/mcp/ 2>/dev/null
```

**Step 4: Verify .claude/ only contains settings and plugin-dev files (if any)**

```bash
ls -la .claude/ 2>/dev/null
```

**Step 5: Commit**

```bash
git add -A
git commit -m "Remove old .claude/ commands, agents, templates, and dropped files

Removed:
- .claude/commands/ (moved to plugin skills)
- .claude/agents/ (moved to plugin agents)
- .claude/templates/ (bundled in plugins)
- .mcp.json (stealth-browser dropped)
- .playwright-mcp/ (dropped with stealth-browser)
- AGENTS.md (agents self-documenting in plugins)"
```

---

### Task 18: Update package.json

**Files:**
- Modify: `package.json`

**Step 1: Read current package.json**

Read `package.json` to see current content.

**Step 2: Update package.json**

Changes:
- Update `version` to `"2.0.0"`
- Update `description` to mention plugin marketplace
- Remove `@modelcontextprotocol/sdk`, `playwright-extra`, and `puppeteer-extra-plugin-stealth` from dependencies (stealth-browser dropped)
- Remove `stealth-browser` and `mcp` from keywords
- Add `plugin` and `marketplace` to keywords
- Keep `@playwright/test` in devDependencies (still used for PDF generation)

**Step 3: Commit**

```bash
git add package.json
git commit -m "Update package.json for v2.0.0 plugin architecture"
```

---

### Task 19: Rewrite CLAUDE.md as contributor-only

**Files:**
- Modify: `CLAUDE.md`

**Step 1: Read current CLAUDE.md**

**Step 2: Rewrite CLAUDE.md**

Strip all operational instructions (HAM-Z methodology, slash command reference, agent table, provenance rules, etc.). Keep only:
- Repository purpose (updated to mention plugin marketplace)
- Key directories (updated for plugin layout)
- File naming conventions
- Version management
- Contributor guidelines

The operational content previously in CLAUDE.md is now embedded directly in the skill and agent files within each plugin.

**Step 3: Commit**

```bash
git add CLAUDE.md
git commit -m "Rewrite CLAUDE.md as contributor-only guide for plugin architecture"
```

---

### Task 20: Rewrite README.md for marketplace

**Files:**
- Modify: `README.md`

**Step 1: Read current README.md**

**Step 2: Rewrite README.md**

New structure:
1. **Header** - JobOps v2.0.0 badge, description as plugin marketplace
2. **Quick Start** - Installation commands:
   ```
   /plugin marketplace add reggiechan74/JobOps
   /plugin install jobops@jobops-marketplace
   /jobops:setup
   ```
3. **Plugins** - Two-column overview of jobops and jobops-ic
4. **jobops Plugin** - Full skill list with descriptions, grouped by domain (resume, interview, job-search, career-strategy, career-crisis, finalization, setup)
5. **jobops-ic Plugin** - Full skill list with descriptions, prerequisite noted
6. **Configuration** - Explain `.jobops/config.json`, template customization
7. **For Contributors** - How to develop, test with `--plugin-dir`, validate with `/plugin validate`
8. **License**

**Step 3: Commit**

```bash
git add README.md
git commit -m "Rewrite README for plugin marketplace distribution"
```

---

### Task 21: Update .gitignore

**Files:**
- Modify: `.gitignore`

**Step 1: Read current .gitignore**

**Step 2: Add plugin-relevant ignores**

Add:
```
# Plugin runtime files
/tmp/.jobops-plugin-root
/tmp/.jobops-ic-plugin-root

# User workspace config (not part of the plugin source)
.jobops/
```

**Step 3: Commit**

```bash
git add .gitignore
git commit -m "Update .gitignore for plugin architecture"
```

---

### Task 22: Validate plugin structure

**Files:** None (validation only)

**Step 1: Validate jobops plugin**

```bash
claude plugin validate plugins/jobops
```

Expected: No errors. Fix any reported issues.

**Step 2: Validate jobops-ic plugin**

```bash
claude plugin validate plugins/jobops-ic
```

Expected: No errors. Fix any reported issues.

**Step 3: Validate marketplace**

```bash
claude plugin validate .
```

Expected: Marketplace validated successfully.

---

### Task 23: Test jobops plugin with --plugin-dir

**Files:** None (testing only)

**Step 1: Start Claude Code with jobops plugin**

```bash
claude --plugin-dir ./plugins/jobops
```

**Step 2: Verify skills are registered**

Run `/help` inside Claude Code. Verify all 27 jobops skills appear with the `jobops:` prefix.

**Step 3: Test a simple skill**

Run `/jobops:install-pandoc` to verify a basic skill loads and executes.

**Step 4: Verify agents are registered**

Run `/agents` and confirm all 16 agents appear.

**Step 5: Test hook**

Check that `/tmp/.jobops-plugin-root` exists and contains a valid path.

---

### Task 24: Test jobops-ic plugin with --plugin-dir

**Files:** None (testing only)

**Step 1: Start Claude Code with both plugins**

```bash
claude --plugin-dir ./plugins/jobops --plugin-dir ./plugins/jobops-ic
```

**Step 2: Verify skills are registered**

Run `/help` and verify both `jobops:` (27) and `jobops-ic:` (10) skill sets appear.

**Step 3: Verify both hooks fired**

```bash
cat /tmp/.jobops-plugin-root
cat /tmp/.jobops-ic-plugin-root
```

Both should contain valid paths.

**Step 4: Test setup flow**

Run `/jobops:setup` and walk through the guided configuration.
Then run `/jobops-ic:setup` to verify it extends the config.

---

### Task 25: Final commit and summary

**Step 1: Verify clean state**

```bash
git status
git log --oneline feature/plugin-split ^main
```

**Step 2: Push feature branch**

```bash
git push origin feature/plugin-split
```

**Step 3: Summary**

Report to user:
- Feature branch `feature/plugin-split` is ready for review
- All skills, agents, hooks, scripts, templates, and manifests are in place
- Old `.claude/` structure has been removed
- README, CLAUDE.md, and package.json updated
- Plugin validation passed
- Manual testing completed with `--plugin-dir`
- When ready, user can request PR to merge to main
