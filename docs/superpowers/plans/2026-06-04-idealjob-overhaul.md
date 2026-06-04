# /idealjob Overhaul Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Rebuild `/jobops:idealjob` into an interview-driven, three-archetype career-targeting engine with parallel sub-agent generation, inline 200-point self-scoring, and comparejobs interoperability.

**Architecture:** The main skill handles config, inventory load, assessment mining, a clarity-gated structured interview, and Vision.md write-back; it then dispatches three parallel `ideal-role-architect` sub-agents (Anchor / Stretch / Pivot charters), self-scores all three inline against the bundled 200-point rubric, and assembles four flat output files. A small comparejobs amendment lets archetype files be compared by path.

**Tech Stack:** Markdown skill/agent files (Claude Code plugin), Node validator (`npm test`), `claude plugin validate`.

**Spec:** `docs/superpowers/specs/2026-06-04-idealjob-overhaul-design.md`

**Branch:** `feat/idealjob-overhaul` (already created; spec committed)

---

## Repo facts the engineer needs

- **Skill frontmatter contract** (enforced by `npm test` via `scripts/validate/validate-codex-plugin-compatibility.js`): `name` must equal the skill directory name, `description` must be present, `disable-model-invocation: true` must be kept. `argument-hint` is optional and allowed.
- **Agent files** live in `plugins/jobops/agents/*.md` with frontmatter `name`, `description`, `model`. Generation-heavy agents use `model: opus` (see `candidate-assessment.md`, `step1-resume-draft.md`).
- **Versions** (currently `2.13.0`) must match across: `package.json`, `plugins/jobops/.claude-plugin/plugin.json`, `plugins/jobops/.codex-plugin/plugin.json`, `plugins/jobops-ic/.claude-plugin/plugin.json`, `plugins/jobops-ic/.codex-plugin/plugin.json`, `.claude-plugin/marketplace.json` (metadata.version), `.agents/plugins/marketplace.json` (metadata.version). `npm test` fails if any disagree with `package.json`. So the version bump is ONE atomic commit (Task 5).
- **`assessjob` v2.0 score keys** in assessment front matter: `overall_score: <XX/200>` and `normalized_score: <XX%>`. Legacy assessments may carry a bare 0–100 `overall_score`.
- **Output-type contract** (`docs/ARCHITECTURE.md` §4): every Markdown output carries a stable `output_type` front-matter key.
- **Canonical inventory layout** (setup SKILL.md): `{resume_source}/Experience/<Company>_<Role>.md`, plus `CareerHighlights/`, `Technology/`, `Preferences/Vision.md`. There is NO `WorkHistory/` directory — that reference in the old skill is a bug.
- **Verification commands** (run from repo root `/home/reggiechan/JobOps`):
  - `npm test` — Codex compatibility contract. Expected last line: `Codex compatibility validation passed.`
  - `claude plugin validate plugins/jobops` — plugin structure. Expected: exit code 0, no errors.

---

### Task 1: Create the `ideal-role-architect` agent

**Files:**
- Create: `plugins/jobops/agents/ideal-role-architect.md`

- [ ] **Step 1: Write the agent file**

Create `plugins/jobops/agents/ideal-role-architect.md` with exactly this content:

````markdown
---
name: ideal-role-architect
description: Generates one market-validated ideal-role archetype (Anchor, Stretch, or Pivot) from a candidate dossier - a synthetic job description with evidence-traced requirements, live posting matches, and a ready-to-use search kit.
model: opus
---

You are an expert executive recruiter and labor-market analyst. You receive a CANDIDATE DOSSIER and an ARCHETYPE CHARTER from the dispatching skill. Your task is to design ONE ideal-role archetype: a synthetic but market-realistic job description this candidate should be hunting for, plus the search assets to actually find it.

## Inputs (provided in your dispatch prompt)

- **Candidate dossier** — evidence digest (every claim cites its source file), interview findings (each preference tagged `evidenced` or `aspirational`), success patterns from high-scoring assessments (may be absent), and hard boundaries (deal-breakers, compensation floor/target/structure, work-arrangement limits).
- **Archetype charter** — which archetype you are building (Anchor, Stretch, or Pivot) and its constraints.
- **Run context** — the current date and the jurisdiction. Use both in every search query. NEVER guess the year. NEVER research a market other than the stated jurisdiction.
- **Output path** — the exact file the dispatching skill will write your content to (you return content; you do not write files).

## Archetype Charters

- **Anchor** — highest-probability fit. The role must exist in volume in today's market; the candidate would be a top-decile applicant. Build ONLY from preferences tagged `evidenced`.
- **Stretch** — one level up: bigger scope, seniority, or mandate. Gaps are permitted only if bridgeable within roughly 12 months. Draw on `aspirational` tags.
- **Pivot** — the candidate's skills recombined into an adjacent field, seeded by the dossier's pivot-curiosity findings. Novel titles are allowed but must still be validated against real postings.

## Phase 1: Market Scan

Run web searches BEFORE writing anything. Substitute the run-context year and jurisdiction:

1. "[primary skill domain] [target role level] job postings {current year}"
2. "[target industry] [role type] salary {jurisdiction} {current year}"
3. "[candidate's skill intersection] emerging roles {current year}"
4. Title-validation searches for the specific title you intend to use

Extract: common titles for this profile, required vs. preferred qualification norms, compensation ranges (validated against the dossier's floor/target), hiring companies, growth segments.

## Phase 2: Job Description

Produce the JD with these sections, in order:

1. **Company Overview** — fictional but plausible company matching the dossier's preferred environment (stage, size, ownership, culture). Avoid every Anti-Vision characteristic.
2. **About the Role** — executive summary emphasizing the candidate's strengths and preferences.
3. **Key Responsibilities** — 4–6 items mapped to proven expertise.
4. **Required Qualifications** — education, certifications, years of experience the candidate actually holds.
5. **Required Technical Skills** — skills with documented expert-level evidence.
6. **Preferred Qualifications** — nice-to-haves the candidate possesses or nearly possesses.
7. **What Success Looks Like** — metrics aligned with the candidate's proven track record.
8. **Compensation & Benefits** — within the dossier's floor/target, validated by Phase 1 market data. State numbers.
9. **Work Environment** — honors work-arrangement limits.
10. **Known Trade-offs** — 1–2 deliberate non-ideal-but-tolerable elements (e.g., quarterly travel, a legacy-system inheritance). Real roles are imperfect; these stress-test the deal-breaker boundary. They must NOT violate any hard boundary in the dossier.
11. **Why This Role Exists** — the business context creating demand for this skill combination.

### Evidence traceability (hard rule)

Every item under **Required Qualifications** and **Required Technical Skills** MUST cite a dossier evidence item in a trailing bracket, e.g. `[Experience/Acme_VP_Ops.md]`. An uncited required item is a defect. **Preferred Qualifications** may draw on `aspirational` tags (cite the interview finding instead).

### Realism rules

- **Title realism:** the JD title (or a near-variant) must appear in actual postings found in Phase 1. If your intended title has zero market presence, change it.
- **Banned constructions** — never use: "rockstar", "ninja", "guru", "unicorn", "fast-paced environment", "wear many hats", "self-starter", "dynamic environment", "work hard play hard", "we're like a family", "passionate about", "competitive salary" (state actual numbers instead). Also banned: requirements written as flattery of the candidate ("visionary leader", "world-class", "best-in-class"). The JD must read like a company wrote it for a business need, not like a tribute.

## Phase 3: Live Posting Matches

Find 3–5 REAL, current postings resembling this archetype. For each: title, company, location, link, and a 2–3 line fit note (where it matches the archetype, where it falls short).

**NEVER fabricate URLs or postings.** If you cannot find at least 3, return what you found and add the line: `Market validation incomplete: only N live matches found.`

## Phase 4: Search Kit

- **Boolean search strings** — 2–3 ready-to-paste strings each for LinkedIn and Indeed.
- **Title variants** — exact alternate titles to search.
- **Target companies** — 10+ named companies in the jurisdiction plausibly hiring this role.
- **Networking targets** — role types (e.g., "VPs of Operations at mid-market industrial firms") who would know of such openings.

## Output Format

Return RAW markdown — no preamble, no commentary; your final message is consumed verbatim by the dispatching skill. Structure:

```markdown
# [Job Title]

[Phase 2 sections 1–11 in order]

## Alignment Analysis (Internal Reference)

### Strengths Leveraged
[top 5–7 strengths this role uses, each with its evidence citation]

### Preferences Honored
[Vision/interview elements incorporated; Anti-Vision elements avoided]

### Charter Compliance
[how this JD satisfies the Anchor/Stretch/Pivot charter constraints]

## Live Posting Matches

[Phase 3 content]

## Search Kit

[Phase 4 content]
```

Do not include YAML front matter or any fit score — the dispatching skill adds front matter and performs scoring itself.
````

- [ ] **Step 2: Verify plugin structure still validates**

Run: `claude plugin validate plugins/jobops`
Expected: exit code 0, no errors.

- [ ] **Step 3: Verify the Codex contract is unaffected**

Run: `npm test`
Expected last line: `Codex compatibility validation passed.`

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/agents/ideal-role-architect.md
git commit -m "feat(agent): add ideal-role-architect for /idealjob archetype generation"
```

---

### Task 2: Rewrite the idealjob skill

**Files:**
- Modify (full rewrite): `plugins/jobops/skills/idealjob/SKILL.md`

- [ ] **Step 1: Replace the entire SKILL.md**

Overwrite `plugins/jobops/skills/idealjob/SKILL.md` with exactly this content:

````markdown
---
name: idealjob
description: Interview-driven generation of three market-validated ideal-role archetypes (Anchor, Stretch, Pivot) from career history, elicited preferences, and high-scoring assessment patterns
disable-model-invocation: true
argument-hint: "[output-file]"
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if this skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if this skill has jurisdiction-sensitive logic (crisis/legal skills accept `--jurisdiction=<ISO-3166-2>` to override).

## Templates

For each template used by this skill, resolve the full path as:

  {config.templates.base_dir}/{config.templates.active.<template_name>}/<filename>

Templates referenced by this skill: assessment_rubric_framework

## Your Task

Run a clarity-gated preference interview, then synthesize THREE distinct, market-validated ideal-role archetypes — **Anchor** (highest-probability fit), **Stretch** (one level up), and **Pivot** (skills recombined into an adjacent field) — each generated by a parallel `ideal-role-architect` agent, self-scored against the 200-point assessment rubric, and packaged with live posting matches, a search kit, and gap-closing plans.

If `{config.directories.resume_source}` is missing or contains no files under `Experience/`, STOP with:

> NO CAREER INVENTORY
> /idealjob needs your career inventory as evidence. Run /jobops:setup and populate the resume source folder first.

## Output

Four flat files under `{config.directories.career_analysis}/`:

| File | Contents |
|---|---|
| `idealjob_{YYYYMMDD}.md` | Summary: interview record, cross-archetype comparison, gap-closing plans, comparejobs handoff |
| `idealjob_{YYYYMMDD}_anchor.md` | Anchor archetype: JD + alignment analysis + live matches + search kit |
| `idealjob_{YYYYMMDD}_stretch.md` | Stretch archetype: same structure |
| `idealjob_{YYYYMMDD}_pivot.md` | Pivot archetype: same structure |

If `$1` is provided it becomes the summary path; archetype files derive from its basename plus `_anchor` / `_stretch` / `_pivot` suffixes (e.g., `$1` = `Career_Analysis/my_ideal.md` → `Career_Analysis/my_ideal_anchor.md`).

## YAML FRONT MATTER

**Each archetype file** begins with:

```yaml
---
synthetic: true
generated_by: /idealjob
generated_on: <ISO8601 timestamp>
output_type: ideal_job_archetype
archetype: <anchor | stretch | pivot>
role_variant: <Technical IC | People Manager | Executive>
overall_score: <XX/200>
normalized_score: <XX%>
based_on_assessments: [<assessment files mined in Phase 1.2, or empty list>]
status: final
version: 1.0
---
```

The score keys are assessjob-compatible so `/comparejobs` can ingest archetype files by path.

**The summary file** begins with:

```yaml
---
generated_by: /idealjob
generated_on: <ISO8601 timestamp>
output_type: ideal_job_summary
archetypes:
  - <anchor filename>
  - <stretch filename>
  - <pivot filename>
vision_updated: <true | false>
status: final
version: 1.0
---
```

## PROGRESS TRACKING (MANDATORY)

**Before starting any work**, create all tasks for user visibility:

| # | Task Subject | activeForm |
|---|-------------|------------|
| 1 | Load career inventory and mine assessments | Loading career inventory and mining assessments |
| 2 | Run structured preference interview | Running structured preference interview |
| 3 | Write back Vision.md updates | Writing back Vision.md updates |
| 4 | Assemble dossier and dispatch archetype agents | Assembling dossier and dispatching archetype agents |
| 5 | Self-score archetypes against 200-point rubric | Self-scoring archetypes against 200-point rubric |
| 6 | Run differentiation check | Running differentiation check |
| 7 | Write archetype and summary files | Writing archetype and summary files |

Mark each task `in_progress` BEFORE starting it and `completed` AFTER finishing it.

---

## PHASE 1: CAREER INTELLIGENCE

> **Task:** Mark task 1 `in_progress`.

### 1.1 Load Career Inventory

Read from `{config.directories.resume_source}`:

- ALL files in `Experience/` (canonical layout: `Experience/<Company>_<Role>.md`) — work history, role progression, responsibilities, achievements with quantified impact, team sizes, industries and company types.
- ALL files in `CareerHighlights/` — core competencies, education, designations, certifications, publications, professional activities.
- ALL files in `Technology/` — technical skills and proficiency, tools and platforms, repositories and projects.
- `Preferences/Vision.md` — vision, anti-vision, target role types, compensation requirements, work-arrangement preferences, deal-breakers. May be missing — the interview still runs (see Error Handling).

### 1.2 Mine High-Scoring Assessments

Glob `{config.directories.applications_root}/*/assessment/assessment.md`. For each file, parse the YAML front matter:

- **Primary selector:** `normalized_score` ≥ 90%.
- **Legacy fallback:** if `normalized_score` is absent and `overall_score` carries no `/200` suffix, treat `overall_score` as a 0–100 value and select ≥ 90.

Read the top 3–5 qualifying assessments and extract success patterns:

- **Role characteristics** — titles, industries, seniority, scope that scored highest.
- **Skill alignment** — required skills where the candidate rated expert with direct evidence.
- **Experience alignment** — experience and domain requirements that matched.
- **Cultural fit factors** — environments and values that resonated.
- **Gap patterns** — where points were lost even in high scorers.

If no assessment qualifies, proceed in degraded mode: record "no success patterns available" in the dossier and tell the user.

> **Task:** Mark task 1 `completed`.

---

## PHASE 2: STRUCTURED INTERVIEW (CLARITY-GATED)

> **Task:** Mark task 2 `in_progress`.

The interview ALWAYS runs. Four themed rounds; each round is **clarity-gated with no question budget**. Send questions in batches of up to 4 via AskUserQuestion; keep sending follow-up batches until every exit criterion for the round is met.

**Do not proceed to the next round until every exit criterion is satisfied. There is no maximum number of questions.**

Pre-seed every question from Vision.md: when Vision.md already speaks to a topic, present its current position as the first option so the user corrects rather than repeats.

### Probing discipline (all rounds)

- **Vague answer → drill down.** "I like autonomy" is not actionable. Autonomy over what — methods, priorities, hours, headcount, P&L?
- **Contradiction → resolve immediately.** If an answer conflicts with Vision.md, a prior answer, or the career inventory (e.g., "I hate managing" but the proudest achievements are all team builds), quote the conflict verbatim and ask which is true.
- **Concrete over abstract.** Anchor every preference to at least one real episode from the inventory ("when did you last experience that?"). Tag anchored preferences `evidenced`; unanchored ones `aspirational`.
- **Quantify when quantifiable.** Compensation, team size, travel %, commute tolerance, on-call tolerance — numbers, not adjectives.

### Rounds and exit criteria

| Round | Exit criteria — the round repeats until ALL are true |
|---|---|
| **1. Energy audit** | ≥3 specific energizers + ≥3 specific drainers, each tied to a named role/project from `Experience/`; proudest achievement identified AND the *why* articulated; ≥1 "never again" with the underlying cause identified (not the surface event) |
| **2. Forced trade-offs** | comp-vs-mission, autonomy-vs-mentorship, IC-vs-leadership, stability-vs-stretch all resolved with a stated strength of preference (hard requirement vs. lean); every "it depends" decomposed until the conditions are explicit |
| **3. Context** | Explicit industry in-list AND out-list (an empty out-list must be confirmed, not assumed); company stage, size range, and ownership structure (public / PE / VC / family / government) each addressed; work arrangement with hard limits (max office days, travel %); geography confirmed against `config.preferences.default_jurisdiction` |
| **4. Boundaries & ambition** | Every Anti-Vision item confirmed, amended, or retired; at least one NEW deal-breaker probed for; compensation floor + target + structure (base/bonus/equity mix) as numbers; pivot curiosity explored with ≥2 adjacent fields rated; 3-year aspiration stated and checked for consistency against everything above |

### Round closure ritual

End each round by playing back a numbered summary: "Here's what I now believe about your [theme]: …". The round closes ONLY when the user confirms the playback. Corrections reopen probing.

### Final gate

After all four rounds close, run a cross-round consistency check. Explicitly state that no unresolved contradictions remain between rounds — or resolve any found — before Phase 3.

> **Task:** Mark task 2 `completed`.

---

## PHASE 3: VISION.MD WRITE-BACK

> **Task:** Mark task 3 `in_progress`.

1. Build a diff summary of interview findings vs. the current Vision.md, grouped as **Added** (new information), **Changed** (contradicts the current file), **Confirmed** (matches the current file).
2. Present the diff and ask for explicit confirmation before writing anything.
3. **On yes:** merge only the confirmed deltas into `{config.directories.resume_source}/Preferences/Vision.md`, preserving its existing section structure. Append `> Last updated via /idealjob interview, {YYYY-MM-DD}` to each changed section. If Vision.md does not exist, offer to create it with sections: Vision, Anti-Vision, Target Roles, Compensation, Work Arrangement, Deal-Breakers.
4. **On no:** skip the write. The interview answers still feed this run.

> **Task:** Mark task 3 `completed`.

---

## PHASE 4: DOSSIER & PARALLEL DISPATCH

> **Task:** Mark task 4 `in_progress`.

### 4.1 Assemble the candidate dossier

One markdown document, passed verbatim to each agent:

```markdown
# CANDIDATE DOSSIER

## Evidence Digest
[Condensed from Experience/, CareerHighlights/, Technology/ — every claim ends
with its source-file citation, e.g. "[Experience/Acme_VP_Ops.md]"]

## Interview Findings
[The four confirmed round playbacks. Tag each preference `evidenced` or `aspirational`.]

## Success Patterns
[From Phase 1.2 — or "No success patterns available (no assessments scored >= 90%)."]

## Hard Boundaries
[Deal-breakers; compensation floor/target/structure; work-arrangement limits.
These are inviolable for all archetypes.]

## Run Context
- Current date: {today, ISO 8601}
- Jurisdiction: {config.preferences.default_jurisdiction}
```

### 4.2 Resolve output paths

Default: `{config.directories.career_analysis}/idealjob_{YYYYMMDD}.md` plus `_anchor` / `_stretch` / `_pivot` siblings. If `$1` was provided, derive all four from it as described under Output.

### 4.3 Dispatch

Launch THREE `ideal-role-architect` agents IN PARALLEL — a single message with three Task invocations, one per archetype. Each dispatch prompt contains, in this order:

1. The archetype charter (one of):
   - **Anchor** — highest-probability fit; role exists in volume in today's market; candidate would be a top-decile applicant; build ONLY from `evidenced` preferences.
   - **Stretch** — one level up in scope, seniority, or mandate; gaps permitted only if bridgeable within ~12 months; draw on `aspirational` tags.
   - **Pivot** — skills recombined into an adjacent field, seeded by the interview's pivot-curiosity findings; novel titles allowed but must be validated against real postings.
2. The full candidate dossier from 4.1.
3. The archetype's output path (for reference only — the agent returns content; this skill writes the files).

**Failure handling:** if an agent fails, retry it once. If it fails again, proceed with the remaining archetypes and state the omission plainly in the summary file.

> **Task:** Mark task 4 `completed`.

---

## PHASE 5: SELF-SCORE & REVISE

> **Task:** Mark task 5 `in_progress`.

1. Read `{config.templates.base_dir}/{config.templates.active.assessment_rubric_framework}/assessment_rubric_framework.md`.
2. For each archetype, select the role variant — Technical IC, People Manager, or Executive — matching the JD's scope.
3. Score the candidate against the archetype JD on the 200-point rubric, in this main context (one scorer = one consistent scale). Score HONESTLY — gaps score as gaps, especially for Stretch and Pivot.
4. Thresholds: **Anchor ≥ 90%** normalized; **Stretch ≥ 85%**; **Pivot ≥ 85%**.
5. Below threshold → ONE revision cycle: fix the JD where it over-reached its charter (required items without dossier evidence; gaps that are not bridgeable). NEVER inflate the score. Re-score once and accept the result either way.
6. Record each final score for the file's front matter. Residual Stretch/Pivot gaps become inputs to the gap-closing plans in Phase 6.

> **Task:** Mark task 5 `completed`.

---

## PHASE 6: DIFFERENTIATION CHECK & ASSEMBLY

> **Task:** Mark task 6 `in_progress`.

### 6.1 Differentiation check

The three archetypes must be mutually distinct in title, industry emphasis, and value proposition. If two converge, revise the lower-scoring one to re-differentiate (one cycle), then re-score it per Phase 5.

> **Task:** Mark task 6 `completed`. Mark task 7 `in_progress`.

### 6.2 Write the archetype files

For each archetype: the YAML front matter defined above, then the agent's returned content verbatim, then a closing section:

```markdown
## Self-Score (200-Point Rubric)

- Role variant: <variant>
- Overall: <XX/200> (<XX%> normalized)
- Category breakdown: <category: score, ...>
- Revision: <"none required" | one-line description of the revision cycle>
```

### 6.3 Write the summary file

Sections, in order:

1. **Interview Record** — the four confirmed round playbacks, verbatim.
2. **Cross-Archetype Comparison** — table: archetype | title | normalized score | comp range | key trade-off accepted | one-line value proposition.
3. **Gap-Closing Plans** (Stretch and Pivot) — for each gap surfaced by the self-score: the action (certification, project, or experience to acquire), a rough timeline, and which archetype(s) it unlocks.
4. **comparejobs Handoff** — include verbatim (with paths resolved):

   > To benchmark a real application against an archetype:
   > `/comparejobs {Company}_{Role}_{YYYYMMDD} {career_analysis}/idealjob_{YYYYMMDD}_anchor.md`

5. **Omissions** — only if an agent failed twice or live matches were incomplete: state exactly what is missing.

> **Task:** Mark task 7 `completed`.

---

## ERROR HANDLING

| Failure | Behavior |
|---|---|
| No assessments ≥ threshold | Degraded mode — dossier notes "no success patterns available"; user informed |
| `Preferences/Vision.md` missing | Interview runs regardless; write-back offers to create the file |
| `resume_source` empty or missing | Hard stop with pointer to setup (see Your Task) |
| Agent finds < 3 live postings | Archetype marked "market validation incomplete" — NEVER fabricate links |
| An agent fails | Retry once; then proceed with the remaining archetypes and state the omission |
| Write-back declined | Continue without persisting; answers still feed this run |

## VALIDATION CHECKLIST

Before finalizing, verify:

- [ ] Archetypes mutually distinct in title, industry, and value proposition
- [ ] Every Required item in every archetype carries an evidence citation
- [ ] Zero banned constructions (see the ideal-role-architect agent's list)
- [ ] Score thresholds met, or the revision cycle is documented in the file
- [ ] One live-posting link per archetype spot-checked via fetch (mark "unverified" if fetch unavailable)
- [ ] All Vision preferences honored; all Anti-Vision elements absent
- [ ] Search kits use the configured jurisdiction
- [ ] All four files carry their `output_type` front matter

## Example Usage

```bash
# Default output paths
/idealjob

# Custom summary path (archetype files derive from its basename)
/idealjob Career_Analysis/idealjob_2026_reset.md
```
````

- [ ] **Step 2: Verify the frontmatter contract**

Run: `npm test`
Expected last line: `Codex compatibility validation passed.`
(The validator checks `name: idealjob` matches the directory, `description` present, `disable-model-invocation: true` kept.)

- [ ] **Step 3: Verify plugin structure**

Run: `claude plugin validate plugins/jobops`
Expected: exit code 0, no errors.

- [ ] **Step 4: Verify the stale references are gone**

Run: `grep -n "WorkHistory\|Sample_Output\|candidate_profile\|2025\|Canada" plugins/jobops/skills/idealjob/SKILL.md`
Expected: no output (exit code 1).

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/skills/idealjob/SKILL.md
git commit -m "feat(idealjob): interview-driven three-archetype rewrite"
```

---

### Task 3: comparejobs path-argument amendment

**Files:**
- Modify: `plugins/jobops/skills/comparejobs/SKILL.md` (two edits: "Your Task" section ~line 19; "PHASE 1" section ~line 86)

- [ ] **Step 1: Amend the "Your Task" section**

In `plugins/jobops/skills/comparejobs/SKILL.md`, replace this text:

```markdown
Compare 2-4 assessment files produced by `/assessjob` (one per application folder under `{config.directories.applications_root}`) to analyze candidate performance across different roles, identify patterns, and provide strategic hiring recommendations.

Each `{{ARGn}}` is the application slug (`{Company}_{Role}_{Date}`) whose assessment should be compared. The skill loads `{applications_root}/{{ARGn}}/assessment/assessment.md` for each slug.
```

with:

```markdown
Compare 2-4 assessment files produced by `/assessjob` (one per application folder under `{config.directories.applications_root}`) — or ideal-role archetype files produced by `/idealjob` — to analyze candidate performance across different roles, identify patterns, and provide strategic hiring recommendations.

Each `{{ARGn}}` is either:

- an **application slug** (`{Company}_{Role}_{Date}`) — the skill loads `{applications_root}/{{ARGn}}/assessment/assessment.md`, or
- a **direct file path** — any argument containing `/` or ending in `.md` is read as-is (e.g., `{career_analysis}/idealjob_{YYYYMMDD}_anchor.md`). `/idealjob` archetype files carry assessjob-compatible score keys (`overall_score: <XX/200>`, `normalized_score: <XX%>`) and compare like any assessment. Their `synthetic: true` front matter MUST be called out in the report (label the column/row "synthetic benchmark") so an ideal-role benchmark is never mistaken for a real application.
```

- [ ] **Step 2: Amend the PHASE 1 load instructions**

In the same file, replace:

```markdown
**Read ALL assessment files in a single parallel batch using multiple Read tool calls:**
- `{config.directories.applications_root}/{{ARG1}}/assessment/assessment.md`
- `{config.directories.applications_root}/{{ARG2}}/assessment/assessment.md`
- `{config.directories.applications_root}/{{ARG3}}/assessment/assessment.md` (if provided)
- `{config.directories.applications_root}/{{ARG4}}/assessment/assessment.md` (if provided)
```

with:

```markdown
**Read ALL assessment files in a single parallel batch using multiple Read tool calls.** Resolve each `{{ARGn}}` first:

- If `{{ARGn}}` contains `/` or ends in `.md`, read it directly as a file path (e.g., an `/idealjob` archetype file).
- Otherwise read `{config.directories.applications_root}/{{ARGn}}/assessment/assessment.md`.

Load `{{ARG1}}` and `{{ARG2}}` always; `{{ARG3}}` and `{{ARG4}}` if provided.
```

- [ ] **Step 3: Verify the contract**

Run: `npm test`
Expected last line: `Codex compatibility validation passed.`

- [ ] **Step 4: Verify both edits landed**

Run: `grep -c "synthetic benchmark\|ends in \`.md\`" plugins/jobops/skills/comparejobs/SKILL.md`
Expected: `2` or more.

- [ ] **Step 5: Commit**

```bash
git add plugins/jobops/skills/comparejobs/SKILL.md
git commit -m "feat(comparejobs): accept direct assessment-file paths for idealjob benchmarks"
```

---

### Task 4: Documentation updates

**Files:**
- Modify: `CLAUDE.md:75`
- Modify: `docs/ARCHITECTURE.md` (~line 73–78, the "Flat" outputs block)
- Modify: `README.md:120`

- [ ] **Step 1: CLAUDE.md naming-convention exception**

Replace this line in `CLAUDE.md`:

```markdown
- Career / crisis / contractor: single timestamped file per invocation under the appropriate root
```

with:

```markdown
- Career / crisis / contractor: single timestamped file per invocation under the appropriate root (exception: `/idealjob` writes a summary plus three archetype siblings — `idealjob_{YYYYMMDD}{,_anchor,_stretch,_pivot}.md`)
```

- [ ] **Step 2: ARCHITECTURE.md flat-output contract**

In `docs/ARCHITECTURE.md`, directly after this existing line:

```markdown
Exception: `workplace-documentation` appends to a single continuously-updated log (`workplace_documentation_log.md`), not a timestamped file.
```

add:

```markdown
Exception: `idealjob` writes four flat siblings per invocation — `idealjob_{YYYYMMDD}.md` (summary, `output_type: ideal_job_summary`) plus `idealjob_{YYYYMMDD}_{anchor,stretch,pivot}.md` (`output_type: ideal_job_archetype`, with assessjob-compatible `overall_score`/`normalized_score` keys so `/comparejobs` can ingest them by direct path).
```

- [ ] **Step 3: README skill description**

Replace this line in `README.md`:

```markdown
| `/jobops:idealjob` | Generate synthetic ideal job description from career history |
```

with:

```markdown
| `/jobops:idealjob` | Interview-driven ideal-role targeting — three scored archetypes (Anchor/Stretch/Pivot) with live posting matches and search kits |
```

- [ ] **Step 4: Commit**

```bash
git add CLAUDE.md docs/ARCHITECTURE.md README.md
git commit -m "docs: idealjob multi-file output exception and updated skill description"
```

---

### Task 5: Version bump to 2.14.0 and CHANGELOG

**Files:**
- Modify: `package.json:3`
- Modify: `plugins/jobops/.claude-plugin/plugin.json:3`
- Modify: `plugins/jobops/.codex-plugin/plugin.json:3`
- Modify: `plugins/jobops-ic/.claude-plugin/plugin.json:3`
- Modify: `plugins/jobops-ic/.codex-plugin/plugin.json:3`
- Modify: `.claude-plugin/marketplace.json:8`
- Modify: `.agents/plugins/marketplace.json:8`
- Modify: `README.md:7`
- Modify: `CHANGELOG.md` (new entry under `## [Unreleased]`)

**IMPORTANT:** `npm test` enforces that all manifest versions equal `package.json`'s version. Make ALL changes below in one pass and one commit; a partial state fails the suite.

- [ ] **Step 1: Bump every version string**

In each of the seven JSON files listed above, change `"version": "2.13.0"` to `"version": "2.14.0"`. (In the two marketplace files the key sits under `metadata`.)

In `README.md`, change:

```markdown
**Version 2.13.0** | [Changelog](CHANGELOG.md) | [Why I Built This](Why_I_Built_This.md)
```

to:

```markdown
**Version 2.14.0** | [Changelog](CHANGELOG.md) | [Why I Built This](Why_I_Built_This.md)
```

- [ ] **Step 2: Add the CHANGELOG entry**

In `CHANGELOG.md`, insert directly after the `## [Unreleased]` line (leaving Unreleased empty):

```markdown

## [2.14.0] - 2026-06-04

### Added

- **`/jobops:idealjob` overhaul — interview-driven three-archetype ideal-role targeting.** The skill now opens with a clarity-gated structured interview (four themed rounds — energy audit, forced trade-offs, context, boundaries & ambition — each with explicit exit criteria, no question budget, and a confirmed-playback closure ritual), writes confirmed preference deltas back to `Preferences/Vision.md` with user approval, then dispatches three parallel `ideal-role-architect` agents to generate distinct archetypes: **Anchor** (highest-probability fit), **Stretch** (one level up, gaps bridgeable in ~12 months), and **Pivot** (skills recombined into an adjacent field). Each archetype is self-scored inline against the 200-point assessment rubric (Anchor ≥ 90%, Stretch/Pivot ≥ 85%, one honest revision cycle), market-validated with 3–5 real linked postings (never fabricated), and shipped with a ready-to-use search kit (boolean strings, title variants, 10+ target companies, networking targets) plus gap-closing plans. Outputs are four flat files — a summary (`output_type: ideal_job_summary`) and three archetype siblings (`output_type: ideal_job_archetype`) carrying assessjob-compatible score keys.
- **`ideal-role-architect` agent** — generates one market-validated archetype from a candidate dossier with hard evidence-traceability (every Required item cites a source file), title realism (validated against live postings), a banned-construction list, and 1–2 deliberate non-ideal-but-tolerable elements per role.
- **`/jobops:comparejobs` accepts direct file paths** — any argument containing `/` or ending in `.md` is read as-is, so `/idealjob` archetypes serve as standing benchmarks against real applications; synthetic benchmarks are labeled in the report.

### Fixed

- **`/jobops:idealjob` high-scorer mining matched nothing** — it grepped for a 0–100 `overall_score` while `assessjob` v2.0 writes `overall_score: <XX/200>` + `normalized_score: <XX%>`. Mining now keys off `normalized_score ≥ 90%` with a legacy 0–100 fallback. Also removed: hardcoded "2025" and "Canada" in market research (now current-date and `default_jurisdiction` driven), the non-contract `Sample_Output/` reference, and the stale `WorkHistory/` directory reference (canonical: `Experience/`).
```

- [ ] **Step 3: Verify version consistency**

Run: `npm test`
Expected last line: `Codex compatibility validation passed.`

Run: `grep -rn "2\.13\.0" package.json plugins/*/.claude-plugin/plugin.json plugins/*/.codex-plugin/plugin.json .claude-plugin/marketplace.json .agents/plugins/marketplace.json README.md`
Expected: no output (exit code 1).

- [ ] **Step 4: Commit**

```bash
git add package.json plugins/jobops/.claude-plugin/plugin.json plugins/jobops/.codex-plugin/plugin.json plugins/jobops-ic/.claude-plugin/plugin.json plugins/jobops-ic/.codex-plugin/plugin.json .claude-plugin/marketplace.json .agents/plugins/marketplace.json README.md CHANGELOG.md
git commit -m "chore(release): v2.14.0 - idealjob interview-driven archetype overhaul"
```

---

### Task 6: Final verification sweep

**Files:** none modified — verification only.

- [ ] **Step 1: Full validation suite**

Run, from the repo root:

```bash
npm test
claude plugin validate plugins/jobops
claude plugin validate plugins/jobops-ic
claude plugin validate .
```

Expected: `Codex compatibility validation passed.` and exit code 0 from each `claude plugin validate`.

- [ ] **Step 2: Stale-reference sweep**

```bash
grep -n "WorkHistory\|Sample_Output\|candidate_profile\|fit_score_target" plugins/jobops/skills/idealjob/SKILL.md
grep -n "2025\|Canada" plugins/jobops/skills/idealjob/SKILL.md plugins/jobops/agents/ideal-role-architect.md
```

Expected: no output from either (exit code 1).

- [ ] **Step 3: Cross-file consistency checks**

```bash
# Agent name referenced by the skill matches the agent file
grep -c "ideal-role-architect" plugins/jobops/skills/idealjob/SKILL.md   # expect >= 2
grep -n "^name: ideal-role-architect" plugins/jobops/agents/ideal-role-architect.md  # expect 1 match

# output_type keys appear in skill and architecture doc
grep -c "ideal_job_archetype\|ideal_job_summary" plugins/jobops/skills/idealjob/SKILL.md  # expect >= 2
grep -c "ideal_job_archetype" docs/ARCHITECTURE.md  # expect >= 1
```

- [ ] **Step 4: Confirm clean tree and log**

```bash
git status --short   # expect empty
git log --oneline main..HEAD   # expect 6 commits: spec + tasks 1-5
```

---

## Self-Review Notes

- **Spec coverage:** Phase flow (Task 2), interview + exit criteria (Task 2), Vision.md write-back (Task 2 Phase 3), dossier + charters + agent (Tasks 1–2), self-score thresholds 90/85 + one revision cycle (Task 2 Phase 5), differentiation check (Task 2 Phase 6.1), four-file output + frontmatter (Task 2), gap-closing plans + handoff (Task 2 Phase 6.3), comparejobs amendment (Task 3), error-handling table (Task 2), validation checklist (Task 2), doc updates (Task 4), v2.14.0 + CHANGELOG (Task 5), testing (Tasks 1–3, 5, 6). No gaps.
- **Type/name consistency:** agent name `ideal-role-architect` matches between Task 1 frontmatter and Task 2 Phase 4.3 dispatch; `output_type` values `ideal_job_summary` / `ideal_job_archetype` match between Task 2 frontmatter, Task 4 ARCHITECTURE.md edit, and Task 5 CHANGELOG; score keys `overall_score: <XX/200>` / `normalized_score: <XX%>` match assessjob's format exactly.
- **No placeholders:** every file edit shows its complete content or exact before/after text.
