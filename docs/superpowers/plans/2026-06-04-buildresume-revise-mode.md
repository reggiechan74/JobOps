# Revise-First `/buildresume` Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make `/buildresume` revise an existing base resume from a user-curated `Tailored_CV/` library by default (manifest + surgical edits + diff gates), with from-scratch as the fallback, and a promotion loop that grows the library.

**Architecture:** A new skill front-end ("Build Mode Selection") scans `config.directories.tailored_cv`, runs an inline fit assessment against user-defined `role_family` labels, and dispatches either a new `step1-resume-revise` agent (copy base → targeted edits → diff gate) or the untouched `step1-resume-draft` agent. Step 2 gains delta tagging (`BASE`/`NEW` origin per finding); Step 3 gains an edit-mode protocol with a second diff gate. After Step 3, either mode offers to promote the final into the library.

**Tech Stack:** Markdown skill/agent prompt files (Claude Code plugin), `.jobops/config.json` contract, `claude plugin validate`, `npm test` (Codex contract validator).

**Spec:** `docs/superpowers/specs/2026-06-04-buildresume-revise-mode-design.md`

**Context for the implementer (read first):**
- This repo ships *prompt files*, not executable code. There is no unit-test framework for skill content. Verification per task = (a) `grep` assertions that the exact required text landed, (b) `claude plugin validate plugins/jobops`, (c) `npm test` for the Codex contract. Run them exactly as written.
- Agents in `plugins/jobops/agents/` are auto-discovered — no manifest registration anywhere.
- Skill/agent conventions are defined in `docs/ARCHITECTURE.md` Section 5 (skill-authoring contract). The buildresume skill already complies; preserve all existing sections you are not told to change.
- Line numbers below are anchors from the current files; if drifted, locate by the quoted heading/text instead.

---

### Task 1: `tailored_cv` config key in setup skill

**Files:**
- Modify: `plugins/jobops/skills/setup/SKILL.md` (directory table ~line 43; config JSON ~line 152)

- [ ] **Step 1: Add the directory-interview table row**

In the `## Step 2: Directory interview` table, insert this row between the `crisis_management` row and the `application_tracker` row:

```markdown
| `tailored_cv` | `./Tailored_CV` | Base resume library for revise-first `/jobops:buildresume` — one user-curated `.md` base per role family; written only via the explicit promotion offer |
```

- [ ] **Step 2: Add the key to the config JSON schema**

In `## Step 6: Write .jobops/config.json`, inside the `"directories"` object, insert after the `"crisis_management"` line:

```json
    "tailored_cv": "<step-2 value>",
```

(Keep `"application_tracker"` as the last directories key.)

- [ ] **Step 3: Verify**

Run:
```bash
grep -c "tailored_cv" plugins/jobops/skills/setup/SKILL.md
```
Expected: `2`

Note: `## Step 3: Create directories` loops over every `config.directories` key except `application_tracker`, so `Tailored_CV/` gets `mkdir -p` automatically — no edit needed there. Confirm by reading that section; it must not special-case `tailored_cv`.

- [ ] **Step 4: Validate and commit**

```bash
claude plugin validate plugins/jobops
git add plugins/jobops/skills/setup/SKILL.md
git commit -m "feat(setup): add tailored_cv base-library directory key"
```

---

### Task 2: New agent `step1-resume-revise`

**Files:**
- Create: `plugins/jobops/agents/step1-resume-revise.md`

- [ ] **Step 1: Create the agent file with this exact content**

```markdown
---
name: step1-resume-revise
description: Revises an existing base resume toward a new job description using a change manifest and surgical edits. This is the revise-mode Step 1 of the three-step resume process. The agent performs gap analysis between the base resume and the job description, writes an explicit change manifest in which every change is justified by a named JD requirement, copies the base resume, applies only the manifest changes as targeted edits, and verifies with a diff gate that all unchanged content remains byte-identical to the base.
model: opus
---

You are a resume revision specialist. Your role is the revise-mode Step 1 of a three-step resume creation process. You NEVER regenerate a resume from scratch — you make the minimum set of surgical changes that adapt a proven base resume to a new job description.

## Your Mission
Adapt an existing, already-hardened base resume to a specific job description through an explicit change manifest and targeted edits, preserving every part of the base that does not need to change.

## Inputs (all passed by the dispatching skill in the Task instruction)
- **Base resume path** — the selected base from the user's Tailored_CV library
- **Job description path** — the target role requirements
- **Master inventory root** — the resume source folder (supplies swap-in material for JD demands the base does not cover)
- **Manifest output path** — convention: `{app_slug}/resume/step1_manifest.md`
- **Draft output path** — convention: `{app_slug}/resume/step1_draft.md`
- **Cultural profile** — only if the user passed one explicitly; otherwise inherit the base's voice and do NOT present the cultural profile or positioning menus

Do not invent paths. If any input path is missing from the Task instruction, stop and report.

## The Frozen-by-Default Rule
Everything in the base resume is FROZEN unless a manifest change explicitly targets it. You must not:
- Rephrase, "improve," or re-flow any bullet, heading, or sentence outside the manifest
- Change formatting conventions (bullet glyphs, separators, ALL-CAPS company names, date formats)
- Reorder sections or bullets except via an explicit manifest reorder change
- Re-wrap lines or normalize whitespace in untouched regions

The measure of success is that `diff <base> <draft>` shows ONLY the manifest changes plus the front-matter update.

## Phase 1 — Gap Analysis
1. Read the base resume completely.
2. Read the job description completely. Extract the top requirements, ATS keywords, and emphasis (e.g., P&L, team scale, domain).
3. Read master inventory files as needed to find swap-in candidates for JD demands the base does not cover (follow the same source-file discovery patterns as the from-scratch step1 agent: Experience/, Education, Publications, Skills files).
4. For each JD requirement, classify the base's coverage: COVERED (leave frozen) / WEAK (candidate for rewrite or re-emphasis) / MISSING (candidate for swap-in from inventory).

## Phase 2 — Change Manifest
Write the manifest to the manifest output path BEFORE touching the draft. Front matter:

    ---
    job_file: <absolute path to JD>
    role: <role title from JD>
    company: <hiring company>
    candidate: <full candidate name from base>
    base_resume: <absolute path to base>
    generated_by: /buildresume step1-resume-revise
    generated_on: <ISO8601 timestamp>
    output_type: resume_manifest
    status: manifest
    version: 1.0
    ---

Body — one entry per change, numbered:

    ## Change N: <short title>
    - **Target:** <section name + quoted anchor text identifying the exact location>
    - **Operation:** rewrite | swap-in | swap-out | reorder | keyword-injection | re-emphasis
    - **JD requirement:** "<quoted requirement text from the JD>"
    - **Before:** "<exact current text from the base, verbatim>" (or `N/A` for swap-in)
    - **After:** "<exact replacement text>" (or `N/A` for swap-out)
    - **Source evidence:** <master-inventory file + line reference for any NEW claim; or `carried from base`>

Manifest rules:
- Every change must cite a JD requirement. No "general improvements."
- Every NEW claim (swap-in or rewrite that adds facts) must cite master-inventory evidence — same zero-fabrication standard as the from-scratch pipeline.
- New content follows the HAM-Z formula ("Achieved [result] by leveraging [hard skill] to [process]") and the base's voice.
- An EMPTY manifest is legal: if the base already fits the JD, write the front matter plus `# Revision Manifest` and `No changes required — base resume covers all JD requirements.` Then copy the base through unchanged (front-matter update only).
- Typical change count is 3–10. If you find yourself writing more than ~15 changes, stop and report that the base is a poor fit and from-scratch may be better — do not rebuild the resume via the manifest.

## Phase 3 — Surgical Apply
1. Copy the base file to the draft output path with `cp`.
2. Update ONLY the YAML front matter of the draft:

    ---
    job_file: <absolute path to JD>
    role: <role title>
    company: <hiring company>
    candidate: <full candidate name>
    generated_by: /buildresume step1-resume-revise
    generated_on: <ISO8601 timestamp>
    output_type: resume_step1
    status: draft
    build_mode: revise
    base_resume: <absolute path to base>
    version: 1.0
    ---

   (Remove base-library keys such as `role_family` and `promoted_from` from the draft's front matter; they belong to the library copy.)
3. Apply each manifest change with a targeted Edit operation using the exact Before/After text. Never rewrite the whole file. Never apply changes not in the manifest.
4. If a change's Before text cannot be found in the draft (the base drifted since the manifest was written), STOP and report the failed change — do not regenerate the section or fall back to rewriting.

## Phase 4 — Diff Gate (MANDATORY before reporting completion)
1. Run `diff <base path> <draft path>`.
2. Verify every changed hunk maps to either (a) a numbered manifest change or (b) the front-matter update. 
3. If any out-of-manifest change appears, REVERT it (restore the base text via Edit) and re-run the diff.
4. Run `wc -w <draft path>`. If manifest changes grew the document materially past the base's length, flag it in your report — do not silently rebalance other sections to compensate.

## Output Expectations
After completing all phases, report:
1. Manifest path and change count (or "empty manifest — base passed through")
2. Draft path
3. The diff summary (hunk count, mapped to change numbers) proving the gate passed
4. Word count of draft vs base
5. Any JD requirements left uncovered (candidates for the user to address manually)
6. Recommendation to proceed to Step 2 (provenance check)

## Critical Operating Principles
- **Never fabricate** — every new claim cites master-inventory evidence
- **Frozen-by-default** — untouched content is byte-identical to the base
- **Minimum effective change** — fewer, better-justified changes beat broad rewrites
- **Inherit the base's voice** — no cultural-profile or positioning re-selection unless the dispatching skill passes one
```

- [ ] **Step 2: Verify frontmatter and key sections landed**

```bash
grep -c "^name: step1-resume-revise" plugins/jobops/agents/step1-resume-revise.md
grep -ci "frozen-by-default" plugins/jobops/agents/step1-resume-revise.md
grep -c "output_type: resume_manifest" plugins/jobops/agents/step1-resume-revise.md
grep -c "Diff Gate" plugins/jobops/agents/step1-resume-revise.md
```
Expected: exactly `1`, then ≥ `2`, ≥ `1`, ≥ `1`.

- [ ] **Step 3: Validate and commit**

```bash
claude plugin validate plugins/jobops
git add plugins/jobops/agents/step1-resume-revise.md
git commit -m "feat(agents): add step1-resume-revise manifest/surgical-edit agent"
```

---

### Task 3: buildresume SKILL.md — mode selection, dispatch, metadata, promotion

**Files:**
- Modify: `plugins/jobops/skills/buildresume/SKILL.md` (Arguments ~line 53; Output metadata ~line 63; Step 1 ~line 85; Step 2 ~line 109; Step 3 ~line 124; Mission Summary ~line 137)

- [ ] **Step 1: Replace the `## Arguments` section**

Replace:

```markdown
## Arguments

- `$1`: Job description file path (required)
- `$2`: Cultural profile (optional, defaults to "Canadian")

Runs the three-step resume build sequentially:
1. **Step 1**: Create initial tailored resume draft using HAM-Z methodology
2. **Step 2**: Perform comprehensive provenance analysis for credibility
3. **Step 3**: Create final hardened resume addressing all issues
```

with:

```markdown
## Arguments

- `$1`: Job description file path (required)
- `$2`: Cultural profile (optional). In scratch mode, defaults to "Canadian". In revise
  mode, the base resume's voice is inherited and the profile menus are skipped; passing
  `$2` explicitly overrides the inherited voice.
- `--base=<path>`: Force a specific base resume file. Skips fit assessment; implies revise mode.
- `--from-scratch`: Skip base discovery entirely; run the from-scratch pipeline.

First selects a build mode (see Build Mode Selection below), then runs the three-step
build sequentially:
1. **Step 1**: Revise the selected base toward the JD via change manifest and surgical
   edits (revise mode) or create an initial tailored draft using HAM-Z methodology
   (scratch mode)
2. **Step 2**: Perform comprehensive provenance analysis for credibility (delta-tagged
   `BASE`/`NEW` in revise mode)
3. **Step 3**: Create final hardened resume addressing all issues (edit-mode with diff
   gate in revise mode)
```

- [ ] **Step 2: Insert the `## Build Mode Selection` section**

Insert immediately after the (new) Arguments section, before `## Output metadata`:

```markdown
## Build Mode Selection

Determine the build mode BEFORE dispatching any step agent.

1. **Flag shortcuts.** `--from-scratch` → mode = `scratch`, skip to dispatch.
   `--base=<path>` → verify the file exists and is `.md`; mode = `revise` with that
   base, skip to dispatch.
2. **Library check.** Read `config.directories.tailored_cv`. If the key is absent, or
   the directory is missing or contains no `.md` files → mode = `scratch`, and tell
   the user once:

   > No base resume library found. Building from scratch. To enable revise mode,
   > re-run /jobops:setup and curate base resumes in your Tailored_CV directory —
   > a great first base is the final this run produces (see promotion offer at the end).

   Non-`.md` files in the directory are ignored with a one-line note.
3. **Stamp untagged bases (one-time, per file).** For each candidate `.md` lacking
   front matter with `output_type: resume_base` and a `role_family` value: infer a
   short category label from the file's content, propose it to the user, let them
   confirm or type their own label (free text — the taxonomy is entirely
   user-defined; never assume a built-in category list), then stamp the file's front
   matter with `output_type: resume_base` and the confirmed `role_family` (create the
   front-matter block if the file has none). Do not alter the resume body.
4. **Fit assessment (inline — do not dispatch a sub-agent).** Read the JD. For each
   candidate base, assess fit on four axes:
   - declared `role_family` match to the JD's role family,
   - positioning level (1–5 IC→C-suite scale, as used by the step1 agents),
   - domain/industry overlap,
   - requirement-keyword coverage.
   Verdict per candidate: **STRONG / PARTIAL / POOR**, each with a one-line rationale
   phrased in the user's own category labels. No numeric scores.
5. **Decision gate (the user always confirms).**
   - Any STRONG candidate → recommend revising from the strongest one; offer
     [revise from <file> / build from scratch / pick another base].
   - Only PARTIAL candidates → present both options, naming the specific gaps
     (e.g., "base covers the summary and primary role, but the JD's P&L emphasis is
     not covered"); offer [revise from <file> / build from scratch].
   - All POOR → mode = `scratch`; state why no base fits.
6. Record the chosen mode and (in revise mode) the absolute base path. Carry both
   through every step dispatch below.
```

- [ ] **Step 3: Extend the `## Output metadata` section**

After the existing common-fields YAML block (ends `generated_on: <ISO8601 timestamp>` / `---`), append:

```markdown
All three files also carry `build_mode: revise` or `build_mode: scratch`. In revise
mode, step 1 and step 3 outputs additionally carry `base_resume: <absolute path to the
selected base>`.

Revise mode writes one additional file — the change manifest:

- **Step 1 manifest** (revise mode only) — `resume/step1_manifest.md`, with
  `generated_by: /buildresume step1-resume-revise`, `output_type: resume_manifest`,
  `status: manifest`, and `base_resume` set. Written by the step1-resume-revise agent
  before it touches the draft.
```

- [ ] **Step 4: Make Step 1 mode-conditional**

Replace the body of `## Step 1: Creating Initial Resume Draft` (keep the heading and the `@$1` JD read) so the agent dispatch reads:

```markdown
**If mode = revise — Deploying Step 1 Agent - Base Revision**

I'm launching the `step1-resume-revise` agent. In its Task instruction I pass, as
absolute paths: the selected base resume, the JD ($1), the master inventory root
(`config.directories.resume_source`), the manifest output path
(`{app_slug}/resume/step1_manifest.md`), and the draft output path
(`{app_slug}/resume/step1_draft.md`). I pass `$2` only if the user supplied it —
otherwise the agent inherits the base's voice and skips the profile menus. This agent
will:
- Gap-analyze the base against the JD
- Write an explicit change manifest (every change justified by a named JD requirement)
- Copy the base and apply only the manifest changes as surgical edits
- Run the diff gate proving unchanged content is byte-identical to the base

**If mode = scratch — Deploying Step 1 Agent - Initial Draft Creation**

I'm launching the step1-resume-draft agent to create an initial tailored resume draft
based on the job requirements. This agent will:
- Analyze the job requirements from $1
- Apply the $2 cultural profile preferences
- Use the HAM-Z methodology for strategic positioning
- **Review ALL education and credentials from master resume against job requirements**
- **Explicitly justify any credential exclusions in agent output**
- Create a targeted first draft optimized for the role

I pass the resolved absolute output path (`{app_slug}/resume/step1_draft.md`, per step 6
of Application Path Resolution) to the agent in its Task instruction.
```

- [ ] **Step 5: Extend Step 2 dispatch**

In `## Step 2: Provenance Analysis`, after the existing sentence "I pass the Step 1 draft path ... in the agent's Task instruction.", append:

```markdown
In revise mode I additionally pass `build_mode: revise`, the base resume path, and the
manifest path (`{app_slug}/resume/step1_manifest.md`) so the agent tags every finding
with its origin (`BASE` = carried from the base, including the user's manual edits;
`NEW` = introduced or modified by the manifest) per its Revise-Mode Delta Tagging
section.
```

- [ ] **Step 6: Extend Step 3 dispatch**

In `## Step 3: Final Hardened Resume`, after the existing sentence ending "...all in the agent's Task instruction.", append:

```markdown
In revise mode I additionally pass `build_mode: revise` and the base resume path so the
agent follows its Revise-Mode Edit Protocol: copy the draft to the final, apply
targeted edits only for Step 2 findings, and prove with a diff gate that nothing else
changed.
```

- [ ] **Step 7: Add the promotion offer**

After the `## Mission Summary` section (end of file), append:

```markdown
## Promotion Offer (both modes)

After Step 3 delivers `resume/step3_final.md`, offer exactly once:

> Promote this final to your Tailored_CV library as a base for future applications?
> 1. Update an existing base (pick which)
> 2. Save as a new variant (pick a filename and a role_family label)
> 3. Skip

Skip this offer entirely if `config.directories.tailored_cv` is not configured.

On **update**: copy `step3_final.md` over the chosen base file, then restamp the
library copy's front matter: `output_type: resume_base`, the base's existing
`role_family` (confirm with the user), `promoted_from: <app_slug>`, and bump
`version`. Remove application-specific keys (`job_file`, `build_mode`, `status`).

On **new variant**: same restamp, but ask the user for the target filename and the
`role_family` label — offer the labels already present in the library plus
"new category" (free text). This is how the user's taxonomy grows.

Never write to the library without this explicit confirmation.
```

- [ ] **Step 8: Verify**

```bash
grep -c "Build Mode Selection" plugins/jobops/skills/buildresume/SKILL.md
grep -c "step1-resume-revise" plugins/jobops/skills/buildresume/SKILL.md
grep -c "Promotion Offer" plugins/jobops/skills/buildresume/SKILL.md
grep -c "resume_manifest" plugins/jobops/skills/buildresume/SKILL.md
grep -c "from-scratch" plugins/jobops/skills/buildresume/SKILL.md
```
Expected: all non-zero; `step1-resume-revise` ≥ 3 (mode selection reference, Step 1 dispatch, metadata section).

- [ ] **Step 9: Validate and commit**

```bash
claude plugin validate plugins/jobops
git add plugins/jobops/skills/buildresume/SKILL.md
git commit -m "feat(buildresume): revise-first mode selection, dispatch, and promotion offer"
```

---

### Task 4: step2-provenance-check — revise-mode delta tagging

**Files:**
- Modify: `plugins/jobops/agents/step2-provenance-check.md` (insert new section after `## Provenance Analysis Methodology`'s Risk Scoring block, ~line 39; extend output format ~lines 240–285)

- [ ] **Step 1: Insert the delta-tagging section**

Insert after the `### Risk Scoring` block (before `## Your Analysis Process`):

```markdown
## Revise-Mode Delta Tagging

When the dispatching skill's Task instruction includes `build_mode: revise`, it also
passes the base resume path and the change-manifest path. In that case:

1. Read the base resume and the manifest in addition to the draft.
2. Classify every line of the draft by origin:
   - **NEW** — text introduced or modified by a manifest change (match against each
     change's Before/After text)
   - **BASE** — text carried unchanged from the base. This includes the user's manual
     edits to the base, which have never been provenance-checked — audit them at full
     rigor; do not assume the base is clean.
3. The audit remains FULL-DOCUMENT. Delta tagging changes the reporting, not the
   scope.
4. Add an `Origin` field to every PROBLEM_STATEMENTS entry (see output format).
5. Split the RISK_ASSESSMENT_SUMMARY counts by origin and list NEW findings first in
   RECOMMENDATIONS_FOR_STEP3 — in revise mode, Step 3 may only touch flagged lines,
   so your findings define its entire edit surface.

In scratch mode (`build_mode: scratch` or absent), skip this section entirely and omit
the Origin field.
```

- [ ] **Step 2: Add the Origin field to the PROBLEM_STATEMENTS format**

In `### === PROBLEM_STATEMENTS ===`, after the line `- **Category:** {A|B|C|D|E|F|G|H|I|J|K}`, insert:

```markdown
- **Origin:** {BASE|NEW} (revise mode only — omit in scratch mode)
```

- [ ] **Step 3: Extend the summary and recommendations formats**

In `### === RISK_ASSESSMENT_SUMMARY ===`, after the `- **Low Risk Issues:**` line, insert:

```markdown
- **Origin Split (revise mode only):** {{N findings in NEW content / M findings in BASE content}}
```

In `### === RECOMMENDATIONS_FOR_STEP3 ===`, after the `- **CRITICAL - Cannot Proceed Until Fixed:**` line, insert:

```markdown
- **Revise-Mode Edit Surface (revise mode only):** {{NEW findings listed first; this list plus BASE findings defines the ONLY lines Step 3 may modify}}
```

- [ ] **Step 4: Verify**

```bash
grep -c "Revise-Mode Delta Tagging" plugins/jobops/agents/step2-provenance-check.md
grep -c "Origin" plugins/jobops/agents/step2-provenance-check.md
```
Expected: `1`; ≥ `4`.

- [ ] **Step 5: Validate and commit**

```bash
claude plugin validate plugins/jobops
git add plugins/jobops/agents/step2-provenance-check.md
git commit -m "feat(agents): delta-tag provenance findings BASE/NEW in revise mode"
```

---

### Task 5: step3-final-resume — edit-mode hardening + diff gate 2

**Files:**
- Modify: `plugins/jobops/agents/step3-final-resume.md` (insert new section after `## Input Requirements`, ~line 17)

- [ ] **Step 1: Insert the edit-protocol section**

Insert after the `## Input Requirements` block (before `## Your Step 3 Process`):

```markdown
## Revise-Mode Edit Protocol

When the dispatching skill's Task instruction includes `build_mode: revise` (it will
also pass the base resume path), you operate in EDIT MODE — you never rewrite the
document:

1. Copy `step1_draft.md` to the final output path with `cp`.
2. Update ONLY the YAML front matter (`generated_by: /buildresume step3-final-resume`,
   `generated_on`, `output_type: resume_final`, `status: final`, keep
   `build_mode: revise` and `base_resume`, bump `version` per the regeneration rule).
3. Apply targeted Edit operations ONLY for lines flagged in the Step 2 provenance
   analysis (all CRITICAL and High findings; Medium per its recommendations; Low
   optional). Text not flagged by Step 2 is FROZEN — no rephrasing, reformatting,
   reordering, or re-flowing, even if you could "improve" it.
4. **Diff gate (MANDATORY):** run `diff <step1_draft path> <step3_final path>` and
   verify every changed hunk maps to a Step 2 finding or the front-matter update.
   Revert any out-of-scope change and re-run the diff. Include the diff summary
   (hunks → finding references) in your completion report.
5. Skip the whole-document enhancement passes described below — in revise mode the
   sections "STEP 3C - Enhancement and Optimization" and global re-verification apply
   only to the lines you edited. Word-count and length checks still run and are
   reported.

In scratch mode (`build_mode: scratch` or absent), ignore this section and follow the
standard process below.
```

- [ ] **Step 2: Verify**

```bash
grep -c "Revise-Mode Edit Protocol" plugins/jobops/agents/step3-final-resume.md
grep -c "Diff gate" plugins/jobops/agents/step3-final-resume.md
```
Expected: `1`; ≥ `1`.

- [ ] **Step 3: Validate and commit**

```bash
claude plugin validate plugins/jobops
git add plugins/jobops/agents/step3-final-resume.md
git commit -m "feat(agents): edit-mode hardening with diff gate for revise builds"
```

---

### Task 6: Documentation — ARCHITECTURE.md and CLAUDE.md

**Files:**
- Modify: `docs/ARCHITECTURE.md` (Section 2 ~line 32; Section 4 output_type table ~line 94)
- Modify: `CLAUDE.md` (User Data Directories table)

- [ ] **Step 1: ARCHITECTURE.md Section 2 — config key**

After the `config.preferences.cover_letter_mode` paragraph, insert:

```markdown
`config.directories.tailored_cv` (default `./Tailored_CV`) is the base resume library for revise-first `/jobops:buildresume`: one user-curated `.md` base per role family, each carrying `output_type: resume_base` and a free-text `role_family` label in front matter (the taxonomy is entirely user-defined — no built-in category list exists). A config written before this key existed has no `tailored_cv`; `buildresume` treats the absent key (or an empty directory) as "no library" and falls back to the from-scratch pipeline with a setup hint — it never self-heals the config. The library is written only via buildresume's explicit promotion offer; the one-time `role_family` stamping of untagged files edits front matter only.
```

- [ ] **Step 2: ARCHITECTURE.md Section 4 — output types and library pattern**

Add two rows to the `output_type` table after the `resume_final` row:

```markdown
| `resume_manifest` | buildresume (revise-mode step 1) | `resume/step1_manifest.md` |
| `resume_base` | buildresume promotion offer | `{tailored_cv}/<base>.md` |
```

After the **Flat** destination block (after the `idealjob` exception paragraph), insert:

```markdown
**Base library** — user-curated revise-mode inputs, written only via buildresume's explicit promotion offer:

    {tailored_cv}/<base>.md            (output_type: resume_base, role_family: <user label>)
```

- [ ] **Step 3: CLAUDE.md — User Data Directories table**

Add a row to the config-key table after `crisis_management`:

```markdown
| `tailored_cv` | `Tailored_CV/` | Base resume library for revise-first `/buildresume` (one `.md` per user-defined role family; written only via promotion offer) |
```

- [ ] **Step 4: Verify**

```bash
grep -c "tailored_cv" docs/ARCHITECTURE.md
grep -c "resume_manifest\|resume_base" docs/ARCHITECTURE.md
grep -c "tailored_cv" CLAUDE.md
```
Expected: all ≥ 1 (ARCHITECTURE ≥ 3).

- [ ] **Step 5: Commit**

```bash
git add docs/ARCHITECTURE.md CLAUDE.md
git commit -m "docs: tailored_cv config contract, resume_base/resume_manifest output types"
```

---

### Task 7: Full validation + version bump

**Files:**
- Modify (via version-bump skill): `package.json`, `plugins/*/.claude-plugin/plugin.json`, `plugins/*/.codex-plugin/plugin.json`, marketplace manifests, `README.md`, `CHANGELOG.md`

- [ ] **Step 1: Run the full validation suite**

```bash
claude plugin validate plugins/jobops
claude plugin validate plugins/jobops-ic
claude plugin validate .
npm test
```
Expected: all pass. If `npm test` fails, fix the reported Codex-contract violation before proceeding (likely cause: malformed skill frontmatter).

- [ ] **Step 2: Version bump**

Invoke the repo's `version-bump` skill for a **minor** bump (new feature: revise-first buildresume). CHANGELOG entry:

```markdown
### Added
- Revise-first `/buildresume`: selects a base resume from the user-curated `Tailored_CV/` library (new `tailored_cv` config key) via fit assessment against user-defined `role_family` labels; revises it through a change manifest and surgical edits with diff gates so unchanged content is byte-identical; from-scratch build is the fallback. Step 2 provenance findings are delta-tagged `BASE`/`NEW`; Step 3 hardens in edit mode. Both modes end with an offer to promote the final into the library.
- New agent `step1-resume-revise`; new output types `resume_manifest` and `resume_base`.
```

- [ ] **Step 3: Final commit and tag (per version-bump skill flow)**

Follow the version-bump skill's commit/tag instructions. Do not push unless the user asks.

---

## Manual acceptance test (run after all tasks, in a real workspace)

Not automatable in this repo; run with the user's workspace:

1. `claude --plugin-dir plugins/jobops`, run `/jobops:setup` reconfigure → confirm the `tailored_cv` prompt appears with default `./Tailored_CV`.
2. Place one hand-edited base `.md` (no front matter) in `Tailored_CV/`, run `/jobops:buildresume <JD>` → confirm: one-time `role_family` stamping prompt; fit verdict with rationale in the user's label; revise recommendation menu.
3. Accept revise → confirm `step1_manifest.md` exists, and `diff <base> Applications/<slug>/resume/step1_draft.md` shows only manifest changes + front matter.
4. After Step 3 → confirm `diff .../step1_draft.md .../step3_final.md` maps to Step 2 findings only, and the promotion offer appears with [update/new variant/skip].
5. Run with `--from-scratch` → confirm legacy pipeline runs unchanged and the promotion offer still appears.
6. Empty `Tailored_CV/` → confirm scratch fallback with the setup hint, no errors.
