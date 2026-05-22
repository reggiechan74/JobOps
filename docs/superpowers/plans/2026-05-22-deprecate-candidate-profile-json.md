# Deprecate `candidate_profile.json` — Implementation Plan (Plan B)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Eliminate the `candidate_profile.json` intermediate layer entirely. Eleven downstream skills currently read this JSON (created by the `resume-summarizer` agent from `ResumeSourceFolder/` markdown); they will be migrated to read source markdown directly. The `resume-summarizer` agent and `candidate_profile_schema.json` template are then deleted. Source markdown becomes the only authoritative representation of candidate data.

**Architecture:** Plan A (shipped in v2.1.0) introduced `/jobops:audit-source` to ensure source markdown is structurally complete. Plan B removes the JSON intermediate that previously sat between source and downstream skills. Hallucination class (LLM-extracted enums baked into JSON, then trusted forever) is eliminated. Each downstream skill makes its judgments in the context of the specific task that needs them, reading the source files it requires.

**Tech Stack:** SKILL.md edits (no executable code), git history-preserving deletes (`git rm`), no schema changes — the schema goes away.

**Spec source:** Inventory dated 2026-05-22 (see this conversation's Explore agent output). Each affected skill's reference points and field consumption are documented per task below.

**Testing philosophy:** This plugin has no unit-test suite — verification uses (a) grep checks that no skill references `candidate_profile.json` after migration, (b) per-skill smoke tests via Claude Code invocation against `ResumeSourceFolder/` (manual, not gated), and (c) end-of-plan repo-wide grep for orphan references.

---

## Pre-flight check

Before starting, verify the prerequisite is met:
- [ ] `/jobops:audit-source` is shipped in v2.1.0 — confirmed by `git log --oneline | grep -i audit-source`
- [ ] No active work in flight that depends on `candidate_profile.json` schema — check `git diff main` for any open changes touching the schema

If either fails, stop and reconcile before proceeding.

---

## File Structure

### Files modified (11 skills + setup + CHANGELOG + version manifests)
- `plugins/jobops/skills/code-red/SKILL.md` (Task 1)
- `plugins/jobops/skills/severance-review/SKILL.md` (Task 2)
- `plugins/jobops/skills/reference-shield/SKILL.md` (Task 3)
- `plugins/jobops/skills/unemployment-prep/SKILL.md` (Task 4)
- `plugins/jobops/skills/idealjob/SKILL.md` (Task 5)
- `plugins/jobops-ic/skills/ratecard/SKILL.md` (Task 6)
- `plugins/jobops-ic/skills/proposaltemplate/SKILL.md` (Task 7)
- `plugins/jobops-ic/skills/pitchdeck/SKILL.md` (Task 8)
- `plugins/jobops-ic/skills/defineservices/SKILL.md` (Task 9)
- `plugins/jobops/skills/assessjob/SKILL.md` (Task 10)
- `plugins/jobops/skills/assesscandidate/SKILL.md` (Task 11)
- `plugins/jobops/skills/setup/SKILL.md` (Task 14 — remove schema from template list)
- `CHANGELOG.md` (Task 15)
- All 6 version files (Task 15) — same set as v2.1.0 bump

### Files deleted
- `plugins/jobops/agents/resume-summarizer.md` (Task 12)
- `plugins/jobops/templates/candidate_profile_schema.json` (Task 13)

### Files not touched
- `/jobops:audit-source` skill family (Plan A) — already in place, prerequisite for this work
- Source markdown under `ResumeSourceFolder/` — never touched by this plan

---

## Phased execution

| Phase | Risk | Tasks | Description |
|---|---|---|---|
| 1 | Low | 1–5 | Optional-load consumers (skill works without profile) |
| 2 | Medium | 6–9 | jobops-ic consumers (depend on profile for service/rate generation) |
| 3 | High | 10–11 | Critical consumers + creators (also trigger resume-summarizer subagent) |
| 4 | Cleanup | 12–14 | Delete agent, schema, update setup |
| 5 | Release | 15 | Version bump to 2.2.0, CHANGELOG, tag |

Phases must run in order. Within a phase, tasks can run sequentially or in parallel (each touches a different file). Phase 3 must complete before Phase 4 begins (cleanup is unsafe while consumers still reference the agent/schema).

---

## Reusable patterns (referenced by Tasks 1–11)

### Source-mapping table (the contract Phase 1–3 tasks implement)

| JSON field/section | Replacement source markdown |
|---|---|
| `candidate.name`, `candidate.current_title` | `ResumeSourceFolder/Identity/Name.md`, `Identity/CurrentRole.md` |
| `candidate.years_total_experience` | Compute from `WorkHistory/*.md` Start/End dates (deterministic) |
| `technical_skills[]` (names, categories) | `Technology/TechStack.md` |
| `technical_skills[].proficiency_level` | **Judge per task** from `WorkHistory/*.md` mentions — no global default |
| `technical_skills[].years_experience` | Compute from earliest WorkHistory mention to latest (date arithmetic) |
| `work_history[]` (company, title, dates, industry) | `WorkHistory/NN_*.md` files (one per role) |
| `work_history[].company_size`, `scope_indicators.*` | **Judge per task** from WorkHistory prose — no global default |
| `work_history[].achievements[].impact_category` | **Judge per task** from achievement context — no global default |
| `certifications[]` | `Technology/Certifications.md` |
| `education[]` | `ResumeSourceFolder/Education/*.md` if present (file-existence-checked) |
| `projects[]` | `ResumeSourceFolder/Projects/*.md` |
| `domain_expertise[]` | Synthesized per task from WorkHistory industries — no precomputed list |
| `thought_leadership.*` | `ResumeSourceFolder/Thought_Leadership/*.md` if present |
| `leadership_experience.*` | Synthesized per task from WorkHistory leadership prose |
| `job_preferences.*` | `Preferences/Vision.md`, `Preferences/AntiVision.md`, `Preferences/Compensation.md` (if present) |
| `evidence.file + evidence.lines` | Skill cites source file:line directly when making claims |

### Replacement preamble (drop into each migrated skill where the JSON read used to be)

```markdown
## Source Reading

Read the source markdown files this skill needs from `{config.directories.resume_source}`. Do not look for or require `candidate_profile.json` — that artifact is deprecated as of v2.2.0 and downstream skills now read source directly. The canonical source layout is documented in `{plugin_root}/skills/audit-source/source_layout.md`; if a required file is missing, prompt the user to run `/jobops:audit-source` and stop.

Specific files this skill needs:
- <file 1 and why>
- <file 2 and why>
- ...

For judgment-call fields (proficiency_level, company_size, impact_category, etc.) that the old JSON pre-baked: make those judgments now, in the context of the specific task you are doing, anchored to source lines you cite directly.
```

### Resume-summarizer subagent dispatch removal

Some skills (assesscandidate, assessjob, pitchdeck, defineservices, ratecard, proposaltemplate) currently dispatch the `resume-summarizer` agent via Task tool when the profile is missing or stale. Migration: **delete every such dispatch**. The skill reads source markdown directly; there is no subagent.

If the source layout is broken, the skill prompts the user to run `/jobops:audit-source` — it does NOT attempt to fix the source itself.

---

## Task 1: Migrate /jobops:code-red

**File:**
- Modify: `plugins/jobops/skills/code-red/SKILL.md` (current JSON reference at line ~68)

**Field mapping:** code-red is an emergency-mode skill; profile load was always optional. Replace with direct source reads when context is needed, but the skill must continue to function without any source folder at all (emergency mode).

- [ ] **Step 1: Read the current SKILL.md**

```bash
grep -n "candidate_profile.json\|resume-summarizer" /home/reggiechan/JobOps/plugins/jobops/skills/code-red/SKILL.md
```

Expected: 1+ matches near line 68. Note the surrounding paragraph for context.

- [ ] **Step 2: Replace the JSON reference**

Use Edit. The text varies slightly per file but follows this pattern:

- **Find**: a paragraph saying "If available, read `{config.directories.resume_source}/.profile/candidate_profile.json` for career context" (or similar)
- **Replace** with:

```
If career context would help, read `{config.directories.resume_source}/Identity/CurrentRole.md` and the most-recent `WorkHistory/NN_*.md` file. Do NOT attempt to load `candidate_profile.json` — that artifact is removed in v2.2.0. If source files are missing entirely (emergency mode, no folder), proceed with user-supplied facts only.
```

- [ ] **Step 3: Remove any `resume-summarizer` Task dispatch** (if present)

Look for any block dispatching the `resume-summarizer` agent. There likely isn't one in code-red (optional load), but search:

```bash
grep -n "resume-summarizer\|subagent_type.*resume" /home/reggiechan/JobOps/plugins/jobops/skills/code-red/SKILL.md
```

If matches found, delete the entire dispatch block. Code-red runs in emergency mode and must not block on subagent dispatches.

- [ ] **Step 4: Verify**

```bash
grep -c "candidate_profile.json" /home/reggiechan/JobOps/plugins/jobops/skills/code-red/SKILL.md
grep -c "resume-summarizer" /home/reggiechan/JobOps/plugins/jobops/skills/code-red/SKILL.md
```

Both must equal 0.

- [ ] **Step 5: Commit**

Stage by exact path. Never `git add -A`.

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops/skills/code-red/SKILL.md && git status --short && git commit -m "refactor(code-red): read source markdown instead of candidate_profile.json

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 2: Migrate /jobops:severance-review

**File:**
- Modify: `plugins/jobops/skills/severance-review/SKILL.md` (current JSON reference at line ~159)

**Field mapping:** severance-review uses profile for career context only (tenure, industry). Replace with WorkHistory direct reads when needed.

- [ ] **Step 1: Read the current SKILL.md and find the reference**

```bash
grep -n "candidate_profile.json\|resume-summarizer" /home/reggiechan/JobOps/plugins/jobops/skills/severance-review/SKILL.md
```

- [ ] **Step 2: Replace the JSON reference**

Find the "If available, read candidate profile" paragraph and replace with:

```
If career context is needed (tenure for severance formula, industry benchmarks), read `WorkHistory/*.md` files directly from `{config.directories.resume_source}`. Compute tenure from explicit Start/End dates. Do NOT load `candidate_profile.json` — removed in v2.2.0.
```

- [ ] **Step 3: Remove any resume-summarizer dispatch** (search, delete if found)

```bash
grep -n "resume-summarizer\|subagent_type.*resume" /home/reggiechan/JobOps/plugins/jobops/skills/severance-review/SKILL.md
```

- [ ] **Step 4: Verify**

`grep -c candidate_profile.json` and `grep -c resume-summarizer` both 0.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops/skills/severance-review/SKILL.md && git status --short && git commit -m "refactor(severance-review): read WorkHistory markdown instead of candidate_profile.json

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 3: Migrate /jobops:reference-shield

**File:**
- Modify: `plugins/jobops/skills/reference-shield/SKILL.md` (current JSON reference at line ~60)

**Field mapping:** uses work_history for supervisor names, tenure overlap, and 10-year career timeline.

- [ ] **Step 1: Find reference**

```bash
grep -n "candidate_profile.json\|resume-summarizer" /home/reggiechan/JobOps/plugins/jobops/skills/reference-shield/SKILL.md
```

- [ ] **Step 2: Replace**

The replacement text:

```
For career timeline and supervisor inventory, read `WorkHistory/*.md` directly. Each file should contain supervisor names, dates, and role context. Do NOT load `candidate_profile.json` — removed in v2.2.0.
```

- [ ] **Step 3: Remove any resume-summarizer dispatch**

- [ ] **Step 4: Verify**

`grep -c` both 0.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops/skills/reference-shield/SKILL.md && git status --short && git commit -m "refactor(reference-shield): read WorkHistory markdown instead of candidate_profile.json

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 4: Migrate /jobops:unemployment-prep

**File:**
- Modify: `plugins/jobops/skills/unemployment-prep/SKILL.md` (current JSON reference at line ~98)

**Field mapping:** uses work_history for tenure/insurable-hours calculation.

- [ ] **Step 1: Find reference**

```bash
grep -n "candidate_profile.json\|resume-summarizer" /home/reggiechan/JobOps/plugins/jobops/skills/unemployment-prep/SKILL.md
```

- [ ] **Step 2: Replace**

```
For employment timeline and insurable-hours calculation, read `WorkHistory/*.md` directly. Compute dates and tenure from explicit Start/End fields. Do NOT load `candidate_profile.json` — removed in v2.2.0.
```

- [ ] **Step 3: Remove resume-summarizer dispatch if present**

- [ ] **Step 4: Verify** — `grep -c` both 0

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops/skills/unemployment-prep/SKILL.md && git status --short && git commit -m "refactor(unemployment-prep): read WorkHistory markdown instead of candidate_profile.json

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 5: Migrate /jobops:idealjob

**File:**
- Modify: `plugins/jobops/skills/idealjob/SKILL.md` (current JSON reference at line ~64)

**Field mapping:** uses profile narrative for self-assessment. Replace with direct reads of Vision.md and CareerHighlights.

- [ ] **Step 1: Find reference**

```bash
grep -n "candidate_profile.json\|resume-summarizer" /home/reggiechan/JobOps/plugins/jobops/skills/idealjob/SKILL.md
```

- [ ] **Step 2: Replace**

```
For self-assessment narrative, read `Preferences/Vision.md` (preferred role/industry signals) and the file `CareerHighlights/CareerHighlights_Summary.md` if present (career narrative). For evidence anchors, read individual `WorkHistory/*.md` files as needed. Do NOT load `candidate_profile.json` — removed in v2.2.0.
```

- [ ] **Step 3: Remove resume-summarizer dispatch if present**

- [ ] **Step 4: Verify** — `grep -c` both 0

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops/skills/idealjob/SKILL.md && git status --short && git commit -m "refactor(idealjob): read Vision/CareerHighlights markdown instead of candidate_profile.json

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 6: Migrate /jobops-ic:ratecard

**File:**
- Modify: `plugins/jobops-ic/skills/ratecard/SKILL.md` (current JSON reference at line ~65)

**Field mapping:** uses years_total_experience, technical_skills (Expert), certifications (Active), work_history budget/scope, leadership_experience.

- [ ] **Step 1: Find reference**

```bash
grep -n "candidate_profile.json\|resume-summarizer\|--from-profile" /home/reggiechan/JobOps/plugins/jobops-ic/skills/ratecard/SKILL.md
```

Note: ratecard may also have a `--from-profile` flag that auto-extracts from JSON. That flag must either be removed or renamed/repurposed.

- [ ] **Step 2: Replace the JSON read with direct source reads**

```
For rate-card generation, read these source files:
- `Preferences/Vision.md` — explicit rate preferences if user has stated them (overrides market research)
- `Technology/TechStack.md` — identify Expert-level skills (judged in-context from WorkHistory mentions, not from a precomputed enum)
- `Technology/Certifications.md` — Active certifications only
- `WorkHistory/*.md` — for years of experience, leadership scope, budget magnitude

Compute `years_total_experience` from earliest WorkHistory Start date to today. Do NOT load `candidate_profile.json` — removed in v2.2.0.
```

- [ ] **Step 3: Handle `--from-profile` flag**

If the flag exists in argument-hint or argument-parsing, replace with `--from-source` (semantic: still auto-extract, just from markdown). Update any inline help text. If the flag is rarely used, deletion is also acceptable — document the choice in the commit message.

- [ ] **Step 4: Remove resume-summarizer dispatch**

```bash
grep -n "resume-summarizer\|subagent_type.*resume" /home/reggiechan/JobOps/plugins/jobops-ic/skills/ratecard/SKILL.md
```

Delete the entire dispatch block (typically a Phase 2 "regenerate profile if stale" section).

- [ ] **Step 5: Verify**

```bash
grep -c "candidate_profile.json" /home/reggiechan/JobOps/plugins/jobops-ic/skills/ratecard/SKILL.md
grep -c "resume-summarizer" /home/reggiechan/JobOps/plugins/jobops-ic/skills/ratecard/SKILL.md
grep -c "\.profile/" /home/reggiechan/JobOps/plugins/jobops-ic/skills/ratecard/SKILL.md
```

All three must equal 0.

- [ ] **Step 6: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops-ic/skills/ratecard/SKILL.md && git status --short && git commit -m "refactor(ratecard): read source markdown directly; remove --from-profile / resume-summarizer dispatch

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 7: Migrate /jobops-ic:proposaltemplate

**File:**
- Modify: `plugins/jobops-ic/skills/proposaltemplate/SKILL.md` (current JSON reference at line ~59)

**Field mapping:** likely similar to pitchdeck (uses work_history, projects, certifications for proposal credentials).

- [ ] **Step 1: Find reference and inventory exact field usage**

```bash
grep -n "candidate_profile.json\|resume-summarizer\|\.profile/" /home/reggiechan/JobOps/plugins/jobops-ic/skills/proposaltemplate/SKILL.md
```

Read 20 lines of context around each match. The skill's full data-extraction pattern may be more elaborate than the line-59 prereq check; inventory before editing.

- [ ] **Step 2: Replace JSON references with source reads**

Mapping pattern (adapt to actual skill content):

```
For proposal-relevant credentials, read:
- `WorkHistory/*.md` — relevant role history, achievements with metrics
- `Projects/*.md` — case studies with outcomes
- `Technology/Certifications.md` — Active certifications relevant to the proposal scope
- `CareerHighlights/*.md` if present — quantified accomplishments
- `Preferences/Vision.md` — engagement and pricing preferences

Do NOT load `candidate_profile.json` — removed in v2.2.0.
```

- [ ] **Step 3: Remove any resume-summarizer dispatch**

- [ ] **Step 4: Verify** — three greps all 0

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops-ic/skills/proposaltemplate/SKILL.md && git status --short && git commit -m "refactor(proposaltemplate): read source markdown directly instead of candidate_profile.json

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 8: Migrate /jobops-ic:pitchdeck

**File:**
- Modify: `plugins/jobops-ic/skills/pitchdeck/SKILL.md` (current JSON references at lines ~35 and ~50)

**Field mapping:** heaviest provenance discipline of any consumer. Currently builds an "evidence index" from JSON fields to validate claims with ≥90% provenance rate. Post-migration, the skill cites source `file:line` directly (which is actually more robust — no intermediate to drift).

- [ ] **Step 1: Read the whole file**

```bash
wc -l /home/reggiechan/JobOps/plugins/jobops-ic/skills/pitchdeck/SKILL.md
grep -n "candidate_profile.json\|resume-summarizer\|evidence\|\.profile/" /home/reggiechan/JobOps/plugins/jobops-ic/skills/pitchdeck/SKILL.md | head -30
```

Pitchdeck is a substantial file. Read it before editing.

- [ ] **Step 2: Replace the JSON load + evidence-index construction**

Replace:
- The current "Phase 2: Load candidate_profile.json and build evidence index" section
- All references to `evidence.file + evidence.lines` indexed FROM the JSON

With direct source-citation discipline:

```
## Source Reading for Pitch Deck

Read these source files directly:
- `CareerHighlights/CareerHighlights_Summary.md` and `CareerHighlights/*.md` — quantified achievements with metrics, timeframe, mechanism
- `Technology/TechStack.md` — skills (judge "Expert" level in-context from depth of WorkHistory mentions, not from a precomputed enum)
- `Technology/Certifications.md` — Active certs only (Expired/In-Progress excluded)
- `Projects/*.md` — case studies
- `Thought_Leadership/*.md` if present — publications, frameworks, awards
- `WorkHistory/*.md` — relevant roles for industry/scope context

Do NOT load `candidate_profile.json` — removed in v2.2.0.

## Provenance Discipline

The pitch deck's 90%+ provenance validation rate is now enforced by direct source citation: every quantified claim on a slide must cite `{filepath}:{line_number}` and that line must contain the verbatim metric. No claim survives without a direct source pointer. Validate by re-reading the cited line before finalizing each slide.

Failure path unchanged: validation rate <80% requires major revision.
```

- [ ] **Step 3: Remove all resume-summarizer dispatch blocks**

```bash
grep -n "resume-summarizer\|subagent_type.*resume" /home/reggiechan/JobOps/plugins/jobops-ic/skills/pitchdeck/SKILL.md
```

Delete every match's surrounding dispatch block.

- [ ] **Step 4: Verify**

```bash
grep -c "candidate_profile.json" /home/reggiechan/JobOps/plugins/jobops-ic/skills/pitchdeck/SKILL.md
grep -c "resume-summarizer" /home/reggiechan/JobOps/plugins/jobops-ic/skills/pitchdeck/SKILL.md
grep -c "\.profile/" /home/reggiechan/JobOps/plugins/jobops-ic/skills/pitchdeck/SKILL.md
grep -c "evidence.file" /home/reggiechan/JobOps/plugins/jobops-ic/skills/pitchdeck/SKILL.md
```

First three must equal 0. The `evidence.file` count may drop substantially but check that any remaining mentions are about source-file citation (not JSON-field references).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops-ic/skills/pitchdeck/SKILL.md && git status --short && git commit -m "refactor(pitchdeck): cite source files directly; remove JSON evidence-index

Provenance discipline now enforced via direct source file:line citations.
Removes resume-summarizer dispatch and candidate_profile.json dependency.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 9: Migrate /jobops-ic:defineservices

**File:**
- Modify: `plugins/jobops-ic/skills/defineservices/SKILL.md` (current JSON references at lines ~46 and ~104)

**Field mapping:** uses Expert technical_skills → service categories, leadership metrics → Strategic Advisory, domain_expertise (5+ years) → Domain Expertise, publications/frameworks → Thought Leadership. Also reads Vision.md for pricing.

- [ ] **Step 1: Find references**

```bash
grep -n "candidate_profile.json\|resume-summarizer\|--from-profile\|\.profile/" /home/reggiechan/JobOps/plugins/jobops-ic/skills/defineservices/SKILL.md
```

- [ ] **Step 2: Replace the JSON load with direct source reads**

```
For service-catalog generation, read these source files:
- `Technology/TechStack.md` — judge Expert-level skills in-context from depth/recency of WorkHistory mentions; each Expert skill maps to a Technical Implementation service
- `WorkHistory/*.md` — extract leadership scope (team size, budget, P&L mentions); leadership signals map to Strategic Advisory services
- `Industries/*.md` if present, otherwise infer from WorkHistory industries — domains with 5+ years of mentions map to Domain Expertise services
- `Thought_Leadership/*.md` if present (publications, frameworks, awards) — these map to Thought Leadership services
- `Preferences/Vision.md` — engagement preferences and pricing anchors; if absent, use market-rate formula (documented in the skill)

Service identification is a judgment task done in the context of this skill invocation. Do NOT pre-compute service candidates from a global enum; judge each potential service against the source evidence.

Do NOT load `candidate_profile.json` — removed in v2.2.0.
```

- [ ] **Step 3: Handle `--from-profile` flag**

If the flag exists: replace with `--from-source` or delete (same as Task 6 for ratecard). Update help text.

- [ ] **Step 4: Remove resume-summarizer dispatch**

Delete every dispatch block.

- [ ] **Step 5: Verify** — four greps (`candidate_profile.json`, `resume-summarizer`, `\.profile/`, `--from-profile`) all 0

- [ ] **Step 6: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops-ic/skills/defineservices/SKILL.md && git status --short && git commit -m "refactor(defineservices): read source markdown directly; remove --from-profile and resume-summarizer

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 10: Migrate /jobops:assessjob

**File:**
- Modify: `plugins/jobops/skills/assessjob/SKILL.md` (current JSON references at lines ~213, ~223, ~228, ~361)

**Field mapping:** the most heavily refactored skill. Currently has two code paths: (a) folder source → trigger resume-summarizer subagent → load JSON profile; (b) single-file source → skip profile, read file directly. Post-migration, both paths read source directly.

- [ ] **Step 1: Read the whole file and locate all JSON-related sections**

```bash
wc -l /home/reggiechan/JobOps/plugins/jobops/skills/assessjob/SKILL.md
grep -n "candidate_profile.json\|resume-summarizer\|\.profile/" /home/reggiechan/JobOps/plugins/jobops/skills/assessjob/SKILL.md
```

There are at least four references; inventory each and the surrounding logic before editing.

- [ ] **Step 2: Replace the folder-source profile-generation flow**

Find the "Phase 2: Profile generation" or equivalent section that dispatches `resume-summarizer`. Replace with:

```
## Phase 2: Source-Reading Strategy

Determine source structure:
- **Single-file source** (resume.md or similar passed by user): read that file directly. Skip everything below; you have the candidate data in-hand.
- **Folder source** (`{config.directories.resume_source}/`): read these files directly for the rubric scoring you are about to do:
  - `Identity/Name.md`, `Identity/CurrentRole.md` (candidate identity)
  - `Technology/TechStack.md` (skill inventory)
  - `Technology/Certifications.md` (active credentials)
  - `WorkHistory/*.md` (roles, achievements, scope)
  - `Projects/*.md` if present (case studies)
  - `Education/*.md` if present
  - `Preferences/Vision.md` (cultural-fit signals)

If a required file is missing, prompt the user to run `/jobops:audit-source` and stop. Do NOT attempt to generate or load `candidate_profile.json` — that artifact is removed in v2.2.0.

Token budget: most rubric scoring needs ~30K of source markdown loaded at once for a thorough assessment. Read what you need; do not pre-summarize.
```

- [ ] **Step 3: Replace the JSON-profile scoring references**

Find scoring-phase sections that say "read from candidate_profile.json" or cite `evidence.file + evidence.lines` indexed via the JSON. Replace with:

```
Score each rubric category against the source files you read in Phase 2. For each score, cite the specific source `{filepath}:{line_number}` you anchored on. Do not invent enums (proficiency_level, company_size, impact_category) — judge from the source prose with citation.
```

- [ ] **Step 4: Remove every resume-summarizer dispatch**

```bash
grep -n "resume-summarizer\|subagent_type.*resume\|Task.*subagent.*summari" /home/reggiechan/JobOps/plugins/jobops/skills/assessjob/SKILL.md
```

Delete each block. Folder-source path no longer needs subagent.

- [ ] **Step 5: Verify**

```bash
grep -c "candidate_profile.json" /home/reggiechan/JobOps/plugins/jobops/skills/assessjob/SKILL.md
grep -c "resume-summarizer" /home/reggiechan/JobOps/plugins/jobops/skills/assessjob/SKILL.md
grep -c "\.profile/" /home/reggiechan/JobOps/plugins/jobops/skills/assessjob/SKILL.md
```

All 0.

- [ ] **Step 6: Smoke-test (manual, not automated)**

If you have a real `Job_Postings/<file>.md`, manually invoke `/jobops:assessjob` against it in a Claude Code session and confirm:
- No error about missing `.profile/` or `candidate_profile.json`
- Rubric scoring produces output anchored to source `file:line` citations
- Token usage remains reasonable (i.e., the skill doesn't load every file in `ResumeSourceFolder/` blindly)

Document the smoke-test outcome in the commit message.

- [ ] **Step 7: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops/skills/assessjob/SKILL.md && git status --short && git commit -m "refactor(assessjob): read source markdown directly; remove resume-summarizer dispatch

Both folder-source and single-file-source paths now read source directly.
Rubric scoring cites source file:line. Smoke-tested against [job posting / fixture].

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 11: Migrate /jobops:assesscandidate

**File:**
- Modify: `plugins/jobops/skills/assesscandidate/SKILL.md` (current JSON references at lines ~152, ~162, ~198)

**Field mapping:** structurally identical to assessjob — also dispatches resume-summarizer for folder sources, then reads JSON profile. Replace with direct source reads using the same Phase 2 pattern as Task 10.

- [ ] **Step 1: Inventory references**

```bash
grep -n "candidate_profile.json\|resume-summarizer\|\.profile/" /home/reggiechan/JobOps/plugins/jobops/skills/assesscandidate/SKILL.md
```

- [ ] **Step 2: Replace with the same Source-Reading Strategy as Task 10**

Use the exact "Phase 2: Source-Reading Strategy" replacement block from Task 10. The file list is identical (assesscandidate reads the same sources to score a rubric that was created earlier).

- [ ] **Step 3: Update the rubric-scoring section**

Find any scoring instruction that cites `evidence.file + evidence.lines` as a JSON field path. Replace with direct source citation language (same as Task 10 Step 3).

- [ ] **Step 4: Remove resume-summarizer dispatch**

- [ ] **Step 5: Verify**

`grep -c candidate_profile.json`, `grep -c resume-summarizer`, `grep -c \.profile/` all 0.

- [ ] **Step 6: Smoke-test manually**

Invoke `/jobops:assesscandidate <rubric> <job-posting>` in Claude Code and confirm output anchored to source citations, no profile errors.

- [ ] **Step 7: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops/skills/assesscandidate/SKILL.md && git status --short && git commit -m "refactor(assesscandidate): read source markdown directly; remove resume-summarizer dispatch

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 12: Delete resume-summarizer agent

**Files:**
- Delete: `plugins/jobops/agents/resume-summarizer.md`

**Prerequisite:** Tasks 1–11 must be complete. Verify no skill references the agent before deleting:

- [ ] **Step 1: Repo-wide verification**

```bash
cd /home/reggiechan/JobOps && grep -rn "resume-summarizer" plugins/ --include="*.md"
```

Expected: empty output. If any match remains, STOP — return to the relevant task and complete the migration first.

- [ ] **Step 2: Delete the agent**

```bash
cd /home/reggiechan/JobOps && git rm plugins/jobops/agents/resume-summarizer.md
```

- [ ] **Step 3: Verify deletion**

```bash
ls plugins/jobops/agents/resume-summarizer.md 2>&1
```

Expected: "No such file or directory".

- [ ] **Step 4: Commit**

```bash
cd /home/reggiechan/JobOps && git status --short && git commit -m "remove: resume-summarizer agent (no longer referenced by any skill)

All 11 consumer skills migrated to read source markdown directly in
Tasks 1-11. Agent's only purpose was creating candidate_profile.json,
which is also removed in this plan.

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 13: Delete candidate_profile_schema.json

**Files:**
- Delete: `plugins/jobops/templates/candidate_profile_schema.json`

**Prerequisite:** Task 12 must be complete.

- [ ] **Step 1: Repo-wide verification**

```bash
cd /home/reggiechan/JobOps && grep -rn "candidate_profile_schema" plugins/ --include="*.md" --include="*.json"
```

Expected: only matches inside `plugins/jobops/skills/setup/SKILL.md` (the template list — to be removed in Task 14) or `.jobops/` workspace files (not in this repo's plugin tree).

If matches exist anywhere else in plugins/, stop and address.

- [ ] **Step 2: Delete the schema**

```bash
cd /home/reggiechan/JobOps && git rm plugins/jobops/templates/candidate_profile_schema.json
```

- [ ] **Step 3: Verify**

```bash
ls plugins/jobops/templates/candidate_profile_schema.json 2>&1
```

Expected: "No such file or directory".

- [ ] **Step 4: Commit**

```bash
cd /home/reggiechan/JobOps && git status --short && git commit -m "remove: candidate_profile_schema.json template (deprecated by source-direct migration)

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 14: Update setup skill — remove schema from bundled templates

**Files:**
- Modify: `plugins/jobops/skills/setup/SKILL.md` (removes candidate_profile_schema from template list)

- [ ] **Step 1: Find the reference**

```bash
grep -n "candidate_profile_schema" /home/reggiechan/JobOps/plugins/jobops/skills/setup/SKILL.md
```

The setup skill lists `candidate_profile_schema` in its bundled-template list (likely around line 91 or 129 per the inventory). Find every mention.

- [ ] **Step 2: Remove the references**

Use Edit to remove every line or list-entry that mentions `candidate_profile_schema`. Adjacent template names (assessment_rubric_framework, etc.) stay intact.

- [ ] **Step 3: Verify**

```bash
grep -c "candidate_profile_schema" /home/reggiechan/JobOps/plugins/jobops/skills/setup/SKILL.md
```

Expected: 0.

Also verify the bundled-templates count if the setup skill cites a number ("ships 5 templates" etc. — the count decreases by 1).

- [ ] **Step 4: Commit**

```bash
cd /home/reggiechan/JobOps && git add plugins/jobops/skills/setup/SKILL.md && git status --short && git commit -m "feat(setup): remove candidate_profile_schema from bundled templates

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

---

## Task 15: Release v2.2.0

**Files:**
- Modify: `plugins/jobops/.claude-plugin/plugin.json` (version 2.1.0 → 2.2.0)
- Modify: `plugins/jobops-ic/.claude-plugin/plugin.json` (version 2.1.0 → 2.2.0)
- Modify: `.claude-plugin/marketplace.json` (metadata.version 2.1.0 → 2.2.0)
- Modify: `package.json` (version 2.1.0 → 2.2.0)
- Modify: `package-lock.json` (both `version` fields 2.1.0 → 2.2.0)
- Modify: `README.md` (version line; skill count may change if any net additions — none expected)
- Modify: `CLAUDE.md` (version line)
- Modify: `SETUP.md` (skill count unchanged at 32 — no edit needed unless count changes)
- Modify: `CHANGELOG.md` (add `[2.2.0] - <date>` entry summarizing the migration)

- [ ] **Step 1: Bump versions in all 7 files**

Use Edit per file. Pattern is identical to the v2.1.0 release (commit `106c071` for reference). Skip SETUP.md if no skill-count change.

- [ ] **Step 2: Add CHANGELOG entry**

Add at the top of CHANGELOG.md (above the `[2.1.0]` entry):

```markdown
## [2.2.0] - <today's date in YYYY-MM-DD>

### Removed

- **`plugins/jobops/agents/resume-summarizer.md`** — agent that created `candidate_profile.json` from source markdown. No longer needed: downstream skills now read source markdown directly.
- **`plugins/jobops/templates/candidate_profile_schema.json`** — JSON Schema for the deprecated profile artifact.
- Generated artifact `<resume_source>/.profile/candidate_profile.json` is no longer created. Users may delete the `.profile/` directory at their convenience.

### Changed

- **Eleven skills migrated to source-direct reads** (Tasks 1–11 of plan `docs/superpowers/plans/2026-05-22-deprecate-candidate-profile-json.md`):
  - jobops: `code-red`, `severance-review`, `reference-shield`, `unemployment-prep`, `idealjob`, `assessjob`, `assesscandidate`
  - jobops-ic: `ratecard`, `proposaltemplate`, `pitchdeck`, `defineservices`
- Each migrated skill now reads the specific source files it needs and makes judgments (proficiency_level, company_size, impact_category, etc.) in the context of the specific task, anchored to source file:line citations.
- Provenance discipline strengthened: skills cite `{filepath}:{line_number}` directly instead of indexing JSON `evidence.file`/`evidence.lines` fields. No intermediate layer to drift.
- `ratecard` and `defineservices` `--from-profile` flag removed (or renamed to `--from-source`).

### Notes

- Companion skill `/jobops:audit-source` (shipped in v2.1.0) is the recommended way to keep source markdown structurally complete. Downstream skills will prompt users to run audit-source if required source files are missing.
- This is the second half of the candidate-profile-JSON deprecation effort. Plan A was v2.1.0 (audit-source); Plan B is this release.
```

- [ ] **Step 3: Repo-wide grep — final verification before release**

```bash
cd /home/reggiechan/JobOps && grep -rn "candidate_profile.json\|resume-summarizer\|candidate_profile_schema" plugins/ --include="*.md" --include="*.json"
```

Expected: empty output. If anything remains in plugins/, do NOT release — return to fix.

```bash
cd /home/reggiechan/JobOps && grep -rn "candidate_profile\|resume-summarizer" docs/ README.md CLAUDE.md CHANGELOG.md
```

Expected: only matches in historical docs (CHANGELOG entries for prior versions, plan files in docs/superpowers/plans/). No live references in CLAUDE.md or README.md.

- [ ] **Step 4: Stage and commit**

Stage by explicit paths (never `git add -A`):

```bash
cd /home/reggiechan/JobOps && git add \
  plugins/jobops/.claude-plugin/plugin.json \
  plugins/jobops-ic/.claude-plugin/plugin.json \
  .claude-plugin/marketplace.json \
  package.json \
  package-lock.json \
  README.md \
  CLAUDE.md \
  CHANGELOG.md \
  && git status --short && git commit -m "chore: release 2.2.0

- Migrate 11 skills off candidate_profile.json (read source markdown directly)
- Remove resume-summarizer agent and candidate_profile_schema.json
- Update version manifests, README, CLAUDE.md, CHANGELOG

Co-Authored-By: Claude Opus 4.7 (1M context) <noreply@anthropic.com>"
```

(Add SETUP.md to the stage list only if its skill count changed.)

- [ ] **Step 5: Tag and push**

```bash
cd /home/reggiechan/JobOps && git tag -a v2.2.0 -m "Release 2.2.0 — deprecate candidate_profile.json; eleven skills now read source markdown directly"
git push origin <feature-branch>  # or main if direct-committing
git push origin v2.2.0
```

---

## Self-review checklist (run after all tasks complete)

- [ ] Eleven skills migrated; per-skill `grep -c candidate_profile.json` returns 0 for each
- [ ] No `resume-summarizer` references anywhere under `plugins/`
- [ ] No `candidate_profile_schema` references anywhere under `plugins/`
- [ ] Agent file `plugins/jobops/agents/resume-summarizer.md` does not exist
- [ ] Schema file `plugins/jobops/templates/candidate_profile_schema.json` does not exist
- [ ] All 6 version files show 2.2.0
- [ ] CHANGELOG has dated `[2.2.0]` entry
- [ ] Tag `v2.2.0` exists locally and on origin
- [ ] At least one smoke-test of `/jobops:assessjob` or `/jobops:assesscandidate` against a real source folder succeeds without profile errors

---

## Out of scope (intentionally deferred)

- **Per-skill GREEN-phase verification**: unlike Plan A, this plan does not run pressure-scenario tests against each migrated skill. The skills' behavior change is mechanical (read source instead of JSON); any latent issues will surface in normal usage. If desired, add a post-release verification round.
- **Token-cost measurement**: the migration shifts token cost from "one ~10K JSON read per skill" to "N source-file reads per skill" where N varies. Measure if usage patterns suggest a problem, not preemptively.
- **`build_source_index.py` helper**: mentioned in earlier conversation as an optional deterministic indexer for the "scan everything" skills. Defer until a specific skill demonstrates a token-cost problem.
- **Renaming `.profile/` directory in `config.directories`**: the user can manually delete their workspace's `.profile/` after upgrading. No config-schema change required.
