# JobOps Plugin Split Design

**Date:** 2026-03-07
**Status:** Approved
**Branch:** Feature branch (TBD) - PR to main when ready

## Overview

Split JobOps v1.7.0 from a monolithic `.claude/` project into two Claude Code plugins distributed via a self-hosted marketplace in the same repository.

- **`jobops`** — Job search, resume development, interview prep, career strategy, career crisis, application finalization
- **`jobops-ic`** — Independent contractor: service definitions, client prospecting, pitch decks, proposals, rate cards, landing pages

## Design Decisions

| Decision | Choice |
|----------|--------|
| Goal | Open-source distribution + modularity |
| Plugin names | `jobops` + `jobops-ic` |
| Marketplace | `jobops-marketplace` in same repo |
| Dependency | Hard prerequisite — jobops-ic requires jobops |
| Repo experience | Plugins only — `.claude/commands/` and `.claude/agents/` removed |
| Component type | Skills with `disable-model-invocation: true` (not commands) |
| Grouping | Flat — no subdirectories in skills |
| Templates | Bundled in plugin, copied to `.jobops/templates/default/` via setup script |
| Template customization | `.jobops/templates/custom/` with config toggle |
| Configuration | `.jobops/config.json` — directories, active templates |
| Path resolution | `SessionStart` hook writes `${CLAUDE_PLUGIN_ROOT}` to `/tmp/`, setup script uses it |
| Setup | `/jobops:setup` (full init + career history), `/jobops-ic:setup` (extends config) |
| CLAUDE.md | Contributor-only, not operational |
| OSINT agents | In jobops, referenced by jobops-ic via dependency |
| Stealth-browser MCP | Dropped |
| searchjobs / hiringcafe-search | Dropped |
| setup-aliases | Dropped |

## Repository Structure

```
JobOps/
  .claude-plugin/
    marketplace.json
  plugins/
    jobops/
      .claude-plugin/
        plugin.json
      hooks/
        hooks.json
      scripts/
        copy-templates.sh
      templates/
        assessment_rubric_framework.md
        evidence_verification_framework.md
        assessment_report_structure.md
        candidate_profile_schema.json
      skills/
        setup/SKILL.md
        buildresume/SKILL.md
        provenance-check/SKILL.md
        coverletter/SKILL.md
        assessjob/SKILL.md
        assesscandidate/SKILL.md
        createrubric/SKILL.md
        comparejobs/SKILL.md
        briefing/SKILL.md
        interviewprep/SKILL.md
        osint/SKILL.md
        auditjobposting/SKILL.md
        idealjob/SKILL.md
        change-one-thing/SKILL.md
        assess-job-offer/SKILL.md
        code-red/SKILL.md
        severance-review/SKILL.md
        workplace-documentation/SKILL.md
        non-compete-analysis/SKILL.md
        reference-shield/SKILL.md
        unemployment-prep/SKILL.md
        discrimination-assessment/SKILL.md
        investigation-prep/SKILL.md
        accommodation-request/SKILL.md
        layoff-intel/SKILL.md
        constructive-dismissal/SKILL.md
        convert-to-pdf/SKILL.md
        convert-to-word/SKILL.md
        markdown-to-pdf/SKILL.md
        install-pandoc/SKILL.md
        github-portfolio/SKILL.md
      agents/
        step1-resume-draft.md
        step2-provenance-check.md
        step3-final-resume.md
        step4-cover-letter.md
        candidate-assessment.md
        interview-briefing.md
        interview-question-generator.md
        resume-summarizer.md
        osint-agent.md
        osint-compensation.md
        osint-corporate.md
        osint-culture.md
        osint-leadership.md
        osint-legal.md
        osint-market.md
        osint-person.md
    jobops-ic/
      .claude-plugin/
        plugin.json
      hooks/
        hooks.json
      scripts/
        copy-templates.sh
      templates/
        service_definition_schema.json
      skills/
        setup/SKILL.md
        defineservices/SKILL.md
        findclient/SKILL.md
        pitchdeck/SKILL.md
        proposaltemplate/SKILL.md
        ratecard/SKILL.md
        copywrite/SKILL.md
        copywriting-spec/SKILL.md
        create-landing-page/SKILL.md
        css-template/SKILL.md
      agents/
        landing-page-copywriter.md
  README.md
  CHANGELOG.md
  CLAUDE.md              (contributor-only)
  LICENSE
  package.json
  ResumeSourceFolder/    (user data, not part of plugins)
  Job_Postings/
  OutputResumes/
  ...
```

## Plugin Manifests

### jobops plugin.json

```json
{
  "name": "jobops",
  "description": "Intelligence-driven job application system - resume development, interview prep, OSINT intelligence, career strategy, and crisis management using HAM-Z methodology",
  "version": "1.7.0",
  "author": {
    "name": "Reggie Chan"
  },
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": ["resume", "job-search", "interview-prep", "osint", "career", "assessment"]
}
```

### jobops-ic plugin.json

```json
{
  "name": "jobops-ic",
  "description": "Independent contractor toolkit - service definitions, client prospecting, pitch decks, proposals, rate cards, and landing pages. Requires jobops plugin.",
  "version": "1.7.0",
  "author": {
    "name": "Reggie Chan"
  },
  "repository": "https://github.com/reggiechan74/JobOps",
  "license": "ISC",
  "keywords": ["independent-contractor", "consulting", "proposals", "rate-card", "pitch-deck", "landing-page"]
}
```

### marketplace.json

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

## Configuration System

### .jobops/config.json

Created by `/jobops:setup`, extended by `/jobops-ic:setup`:

```json
{
  "version": "1.0",
  "directories": {
    "resume_source": "./ResumeSourceFolder",
    "job_postings": "./Job_Postings",
    "output_resumes": "./OutputResumes",
    "scoring_rubrics": "./Scoring_Rubrics",
    "briefing_notes": "./Briefing_Notes",
    "intelligence_reports": "./Intelligence_Reports",
    "client_prospects": "./Client_Prospects",
    "sample_output": "./Sample_Output"
  },
  "templates": {
    "base_dir": "./.jobops/templates",
    "active": {
      "assessment_rubric_framework": "default",
      "evidence_verification_framework": "default",
      "assessment_report_structure": "default",
      "candidate_profile_schema": "default",
      "service_definition_schema": "default"
    }
  }
}
```

### Template Directory Structure

```
.jobops/
  config.json
  templates/
    default/              (copied from plugin on setup, read-only by convention)
      assessment_rubric_framework.md
      evidence_verification_framework.md
      assessment_report_structure.md
      candidate_profile_schema.json
      service_definition_schema.json
    custom/               (user creates modified versions here)
```

### Template Resolution in Skills

Every skill (except setup) starts with:

```
1. Read .jobops/config.json (fail fast if missing -> "Run /jobops:setup")
2. For each template needed:
   - Read config.templates.active[template_name] -> "default" or "custom"
   - Read from {config.templates.base_dir}/{active_value}/{filename}
3. For each directory reference:
   - Use config.directories[key] instead of hardcoded paths
```

## Plugin Root Path Resolution

### Problem

`${CLAUDE_PLUGIN_ROOT}` is only expanded in hooks and MCP configs, not in skill markdown. Setup skills need to copy templates from the plugin directory to the workspace.

### Solution

Each plugin uses a `SessionStart` hook to write its plugin root path to a known temp file. Setup scripts read from there.

**hooks/hooks.json (jobops):**

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

**hooks/hooks.json (jobops-ic):**

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

**scripts/copy-templates.sh (both plugins):**

```bash
#!/bin/bash
PLUGIN_ROOT="$(cat /tmp/.jobops-plugin-root)"  # or .jobops-ic-plugin-root
TARGET_DIR="$1"
mkdir -p "$TARGET_DIR"
cp -r "$PLUGIN_ROOT/templates/"* "$TARGET_DIR/"
echo "Copied $(ls "$TARGET_DIR" | wc -l) templates to $TARGET_DIR"
```

**Setup skill invokes:**

```
bash "$(cat /tmp/.jobops-plugin-root)/scripts/copy-templates.sh" ".jobops/templates/default"
```

## Dependency Enforcement

Every jobops-ic skill includes this prerequisite check:

```
## Prerequisite Check

Before executing, verify the jobops plugin is installed by checking if the
resume-summarizer agent is available. If not found, stop with:

  PREREQUISITE MISSING: jobops plugin

  The jobops-ic plugin requires the jobops plugin to be installed first.
  Install it with:
    /plugin install jobops@jobops-marketplace

  Then retry this command.
```

Additionally, `/jobops-ic:setup` checks for `.jobops/config.json` (created by `/jobops:setup`).

## Setup Flow

### /jobops:setup

Interactive guided flow:

1. **Directory Configuration** — Ask user for each directory path (with sensible defaults)
2. **Template Installation** — Run copy-templates.sh to copy defaults to `.jobops/templates/default/`
3. **Career History Import** (optional) — Parse existing resume files into ResumeSourceFolder
4. **Profile Selection** (optional) — Choose cultural profile (Canadian/American)
5. **Save** — Write `.jobops/config.json`

### /jobops-ic:setup

1. **Prerequisite Check** — Verify `.jobops/config.json` exists
2. **IC Directory Configuration** — Ask for `client_prospects` directory path
3. **Template Installation** — Copy `service_definition_schema.json` to `.jobops/templates/default/`
4. **Update Config** — Add IC template entry and directory to existing config.json

## Skill Inventory

### jobops (27 skills, 17 agents)

| Skill | Original Location | Agents Used |
|-------|-------------------|-------------|
| setup | new (absorbs create-career-history) | resume-summarizer |
| buildresume | resume-development/buildresume | step1-resume-draft, step2-provenance-check, step3-final-resume |
| provenance-check | resume-development/provenance-check | step2-provenance-check |
| coverletter | resume-development/coverletter | step4-cover-letter |
| assessjob | interview-prep/assessjob | candidate-assessment, resume-summarizer |
| assesscandidate | interview-prep/assesscandidate | candidate-assessment |
| createrubric | interview-prep/createrubric | candidate-assessment |
| comparejobs | interview-prep/comparejobs | none |
| briefing | interview-prep/briefing | interview-briefing |
| interviewprep | interview-prep/interviewprep | interview-question-generator |
| osint | job-search/osint | osint-agent + 7 sub-agents |
| auditjobposting | job-search/auditjobposting | none |
| idealjob | career-strategy/idealjob | none |
| change-one-thing | career-strategy/change-one-thing | none |
| assess-job-offer | career-strategy/assess-job-offer | none |
| code-red | career-crisis/code-red | none |
| severance-review | career-crisis/severance-review | none |
| workplace-documentation | career-crisis/workplace-documentation | none |
| non-compete-analysis | career-crisis/non-compete-analysis | none |
| reference-shield | career-crisis/reference-shield | none |
| unemployment-prep | career-crisis/unemployment-prep | none |
| discrimination-assessment | career-crisis/discrimination-assessment | none |
| investigation-prep | career-crisis/investigation-prep | none |
| accommodation-request | career-crisis/accommodation-request | none |
| layoff-intel | career-crisis/layoff-intel | none |
| constructive-dismissal | career-crisis/constructive-dismissal | none |
| convert-to-pdf | application-finalization/convert-to-pdf-playwright | none |
| convert-to-word | application-finalization/convert-to-word | none |
| markdown-to-pdf | application-finalization/markdown-to-pdf | none |
| install-pandoc | system-setup/install-pandoc | none |
| github-portfolio | system-setup/github-portfolio | none |

### jobops-ic (10 skills, 1 agent)

| Skill | Original Location | Agents Used |
|-------|-------------------|-------------|
| setup | new | none (reads jobops config) |
| defineservices | independent-contractor/defineservices | resume-summarizer (from jobops) |
| findclient | independent-contractor/findclient | osint agents (from jobops) |
| pitchdeck | independent-contractor/pitchdeck | resume-summarizer (from jobops) |
| proposaltemplate | independent-contractor/proposaltemplate | none |
| ratecard | independent-contractor/ratecard | none |
| copywrite | landing-page/copywrite | landing-page-copywriter |
| copywriting-spec | landing-page/copywriting-spec | none |
| create-landing-page | landing-page/create | landing-page-copywriter |
| css-template | landing-page/css-template | none |

## Removed from Repository

| Item | Reason |
|------|--------|
| `.claude/commands/` | Moved into plugin skills |
| `.claude/agents/` | Moved into plugin agents |
| `.claude/templates/` | Bundled in plugins, copied on setup |
| `.mcp.json` | Stealth-browser dropped |
| `.playwright-mcp/` | Dropped with stealth-browser |
| `scripts/mcp/` | Dropped with stealth-browser |
| `AGENTS.md` | Agents self-documenting in plugins |
| `searchjobs` skill | Dropped - to be rebuilt separately |
| `hiringcafe-search` agent | Deprecated, dropped |
| `setup-aliases` command | Dropped |

## Modified Files

| File | Change |
|------|--------|
| `CLAUDE.md` | Strip operational instructions, keep contributor guidelines only |
| `README.md` | Rewrite for marketplace installation and plugin documentation |
| `package.json` | Remove stealth-browser deps, update description |

## User Installation Flow

```bash
# Add marketplace
/plugin marketplace add reggiechan74/JobOps

# Install base plugin
/plugin install jobops@jobops-marketplace

# Initialize workspace
/jobops:setup

# Optional: install contractor add-on
/plugin install jobops-ic@jobops-marketplace
/jobops-ic:setup
```

## Skill Frontmatter Pattern

All skills use:

```yaml
---
description: <skill description>
disable-model-invocation: true
---
```

## Standard Skill Preamble

Every skill (except setup) begins with:

```markdown
## Configuration

Read `.jobops/config.json`. If not found:

  JOBOPS NOT CONFIGURED
  Run /jobops:setup to initialize your workspace.

Use config.directories for all file paths.
Use config.templates.active to resolve template locations.
For each template needed, read from:
  {config.templates.base_dir}/{active_value}/{filename}
```

## Standard jobops-ic Preamble

Every jobops-ic skill (except setup) begins with:

```markdown
## Prerequisite Check

Verify the jobops plugin is installed by checking if the resume-summarizer
agent is available. If not found, stop with install instructions.

## Configuration

Read `.jobops/config.json`. If not found:

  JOBOPS NOT CONFIGURED
  Run /jobops:setup to initialize your workspace.
```
