# Forward-Facing Cover Letter Mode Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a config-selectable forward-facing cover letter mode in which the candidate proposes a dual-anchored first-90-days plan (real evidenced problem + past proof point), gathered through a mandatory skill-level interview.

**Architecture:** A new self-contained `step4-cover-letter-forward.md` agent sits beside the existing retrospective agent. The `/coverletter` skill resolves an effective mode (`--mode=` flag > `config.preferences.cover_letter_mode` > `retrospective`), runs a structured intake interview for forward mode, and routes to the matching agent. `/jobops:setup` gains a preference question and schema key.

**Tech Stack:** Markdown agent/skill prompt files; JSON config (`.jobops/config.json`); validation via `claude plugin validate plugins/jobops` and `npm test` (Codex compatibility contract). No runtime code or unit-test harness — verification is structural validation plus targeted content (grep) checks and manual functional review.

**Spec:** `docs/superpowers/specs/2026-06-01-forward-facing-cover-letter-mode-design.md`

**Note on "tests" in this plan:** these are prompt/config files, not executable code. Each task's verification uses the repo's real validators (`claude plugin validate`, `npm test`) plus `grep` presence/consistency checks that assert the required content exists. There is no unit-test framework to add; do not invent one.

**Cross-file rule-drift warning:** `step4-cover-letter-forward.md` reproduces shared, mode-agnostic rules from `step4-cover-letter.md` verbatim (contact header, Step 3a pipeline, voice/style, banned list, Step 6a review). Any future edit to a shared rule must be applied to BOTH files. Task 6 adds a contributor note recording this.

---

## File Structure

| File | Responsibility | Action |
|------|----------------|--------|
| `plugins/jobops/skills/setup/SKILL.md` | Preference interview + config schema | Modify (Step 4 question, Step 6 schema) |
| `plugins/jobops/agents/step4-cover-letter-forward.md` | Forward-mode drafting agent | Create (copy retrospective agent, apply diffs) |
| `plugins/jobops/skills/coverletter/SKILL.md` | Mode resolution, interview, agent routing | Modify |
| `docs/ARCHITECTURE.md` | Config contract documentation | Modify |
| `README.md` | User-facing feature mention | Modify |
| `CHANGELOG.md` | Release notes | Modify |
| Version files (×7) | Version bump 2.12.0 → 2.13.0 | Modify (via version-bump skill) |
| `CLAUDE.md` | Contributor note on the two-agent pairing | Modify |

Implementation order: config (Task 1) → forward agent (Tasks 2–3) → skill routing (Task 4) → validation (Task 5) → docs + contributor note (Task 6) → version + changelog (Task 7).

---

## Task 1: Add `cover_letter_mode` preference to setup

**Files:**
- Modify: `plugins/jobops/skills/setup/SKILL.md` (Step 4 interview, ~lines 72–84; Step 6 schema, ~lines 154–157)

- [ ] **Step 1: Add the third preference question to Step 4**

In `plugins/jobops/skills/setup/SKILL.md`, locate the Step 4 question list that currently ends with the "Default jurisdiction" item (around line 78–80). After that item and before the `Do **not** ask for default_currency` note, insert:

```markdown
3. **Cover letter mode** — enum `retrospective` | `forward`. Default `retrospective`.
   Controls how `/jobops:coverletter` writes the letter. `retrospective` matches the
   job's requirements to what the candidate has already done. `forward` proposes a
   dual-anchored first-90-days plan (each action sits on a real, research-evidenced
   problem and a past proof point) and runs a short intake interview every time it is
   used. The mode can be overridden per-invocation with `/jobops:coverletter --mode=`.
```

- [ ] **Step 2: Add the key to the Step 6 config schema**

In the same file, in the `"preferences"` block of the Step 6 JSON schema (currently lines 154–157), add the third key so the block reads:

```json
  "preferences": {
    "cultural_profile": "<step-4 value>",
    "default_jurisdiction": "<step-4 value>",
    "cover_letter_mode": "<step-4 value, default retrospective>"
  },
```

- [ ] **Step 3: Verify the content is present and the JSON block is well-formed**

Run:
```bash
cd /home/reggiechan/JobOps
grep -n 'Cover letter mode' plugins/jobops/skills/setup/SKILL.md
grep -n 'cover_letter_mode' plugins/jobops/skills/setup/SKILL.md
```
Expected: the question line matches once, `cover_letter_mode` matches twice (question prose + schema). Confirm the schema block still has matched braces and the new line ends with a comma only if a key follows it (it is the last key in `preferences`, so it must NOT have a trailing comma — verify the preceding `default_jurisdiction` line now ends with a comma and the new line does not).

- [ ] **Step 4: Validate plugin structure unaffected**

Run:
```bash
cd /home/reggiechan/JobOps && claude plugin validate plugins/jobops
```
Expected: validation passes (no schema/frontmatter errors).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add plugins/jobops/skills/setup/SKILL.md
git commit -m "feat(setup): add cover_letter_mode preference (retrospective|forward)

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 2: Create the forward agent file from the retrospective base

This task creates the new agent file as an exact copy of the retrospective agent, so the shared rules are reproduced verbatim. Task 3 then applies the forward-specific diffs. Splitting it this way keeps each diff reviewable.

**Files:**
- Create: `plugins/jobops/agents/step4-cover-letter-forward.md` (from `plugins/jobops/agents/step4-cover-letter.md`)

- [ ] **Step 1: Copy the retrospective agent to the new filename**

```bash
cd /home/reggiechan/JobOps
cp plugins/jobops/agents/step4-cover-letter.md plugins/jobops/agents/step4-cover-letter-forward.md
```

- [ ] **Step 2: Update the frontmatter `name` and `description`**

In `plugins/jobops/agents/step4-cover-letter-forward.md`, change the frontmatter (lines 1–11). Replace:

```yaml
name: step4-cover-letter
description: Creates a compelling cover letter with requirements-matching table based on the final Step 3 resume
```

with:

```yaml
name: step4-cover-letter-forward
description: Creates a forward-facing cover letter proposing a dual-anchored first-90-days plan (evidenced problem + past proof) based on the final Step 3 resume and a skill-level intake interview
```

Leave the `tools:` list unchanged.

- [ ] **Step 3: Verify the copy and frontmatter**

Run:
```bash
cd /home/reggiechan/JobOps
grep -n 'name: step4-cover-letter-forward' plugins/jobops/agents/step4-cover-letter-forward.md
grep -c 'name: step4-cover-letter$' plugins/jobops/agents/step4-cover-letter-forward.md
```
Expected: the forward `name:` matches once; the bare retrospective `name: step4-cover-letter` (no `-forward`) matches 0 times (confirms the frontmatter name was changed, not left as the retrospective one).

- [ ] **Step 4: Validate plugin structure**

Run:
```bash
cd /home/reggiechan/JobOps && claude plugin validate plugins/jobops
```
Expected: passes (the new agent is discovered, frontmatter valid).

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add plugins/jobops/agents/step4-cover-letter-forward.md
git commit -m "feat(agent): scaffold step4-cover-letter-forward from retrospective base

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 3: Apply forward-mode diffs to the new agent

All edits are in `plugins/jobops/agents/step4-cover-letter-forward.md`. Apply each step's exact replacement.

**Files:**
- Modify: `plugins/jobops/agents/step4-cover-letter-forward.md`

- [ ] **Step 1: Rewrite the Overview to describe forward mode**

Replace the `## Overview` paragraph:

```markdown
## Overview
I create compelling, tailored cover letters based on the final Step 3 resume, featuring a strategic requirements-matching table that directly maps job requirements to your proven experience.
```

with:

```markdown
## Overview
I create forward-facing cover letters based on the final Step 3 resume and a structured intake interview supplied by the `/coverletter` skill. Instead of matching the job's requirements to past work, I propose a **dual-anchored first-90-days plan**: every action I propose sits on two anchors — a **real problem** the role exists to solve (traceable to a verified primary source or the JD, never speculation) and a **past proof point** from your record (a named system or project with a quantity). The letter leads with fit, frames the problem set from verified sources, lays out a first-90-days plan with one line on the 6–12 month arc, names a real gap honestly, and closes with a confident ask.
```

- [ ] **Step 2: Add the interview-input contract to Input Validation (§1)**

Replace the `### 1. Input Validation` block:

```markdown
### 1. Input Validation
First, I'll verify that I have:
- The final Step 3 resume (hardened and verified)
- The original job description
- Company and role details for personalization
```

with:

```markdown
### 1. Input Validation
First, I'll verify that I have:
- The final Step 3 resume (hardened and verified)
- The original job description
- Company and role details for personalization
- **The forward intake interview answers** passed by the `/coverletter` skill: (1) the role thesis / problem set, (2) the candidate's first-90-days actions plus a line on the 6–12 month arc, (3) the past proof backing each action, (4) the real gap and how the candidate would work with it, (5) any company-specific notes. **Forward mode does not run without these.** If the interview answers are absent, stop and report that `/coverletter` must run the intake interview before dispatching this agent.
```

- [ ] **Step 3: Add the problem-anchoring contract to the Step 3a pipeline (§3a, Step 4 — "Decide context mode")**

In the `**Step 4 — Decide context mode**` block of §3a, after the existing two bullets (`≥ 2 verified sources` and `< 2 verified sources`), append:

```markdown

**Step 5 — Anchor the forward plan to verified problems (forward mode only)**

The interview gave the candidate's read on the problem set, primed only from the JD and existing OSINT files. Now bind each proposed first-90-days action to evidence:

- An action whose target **problem** traces to a `verified` primary source (or is stated plainly in the JD) is anchored — it may appear in the table and body.
- An action whose problem traces to **neither** a verified source nor the JD is **speculation**. Do not ship it. Either reframe the action against a JD-stated requirement, or drop it.
- Record any dropped/reframed action in the closing notes to the user, naming which interview input was set aside and why (so the candidate can supply evidence and re-run).

With `< 2 verified sources`, forward actions may only anchor their problem to the JD itself; keep them conservative and JD-scoped rather than citing a market situation the research could not confirm.
```

- [ ] **Step 4: Convert the Requirements Alignment table (§4.3) to the First-90-Days Plan table**

Replace the entire `#### 4.3 Requirements Alignment table (mandatory)` section (heading, constraints, and the sample table) with:

```markdown
#### 4.3 First-90-Days Plan table (mandatory)

A three-column table that is the scannable core of the forward letter. Each row binds a **real problem** to the **action** the candidate would take in the first 90 days and the **proof** that the candidate can do it. The table is mandatory.

**Constraints:**
- **Cap at 5 rows.** Six or more becomes a wall.
- **Every problem cell must trace** to a verified primary source or the JD (no speculation; see §3a Step 5).
- **Every proof cell must contain a named entity** (project, deal, system, agency, counterparty) and at least one quantity.
- **Vary row construction.** Do not let every action read `verb + object`; lead some rows with the problem or the constraint.
- **No dead verbs** (see banned-construction list).

| **The problem (evidenced)** | **What I'd do in the first 90 days** | **Why I can (proof)** |
|------------------------------|--------------------------------------|------------------------|
| [Problem 1 — traceable to a verified source or the JD] | [Concrete first-90-days action scoped to that problem] | [Named system/project + quantity from the Step 3 resume] |
| [Lead with the constraint: the binding limit the role exists to resolve] | [Action that resolves it] | [Named proof + quantity] |
| [Problem 3] | [Action] | [Named proof + quantity] |
| [Problem 4] | [Action] | [Named proof + quantity] |
| [Problem 5] | [Action] | [Named proof + quantity] |
```

- [ ] **Step 5: Convert the "On X:" paragraphs (§4.4) to "How I'd approach X:" paragraphs**

Replace the entire `#### 4.4 "On X:" evidence paragraphs (2–3 paragraphs)` section with:

```markdown
#### 4.4 "How I'd approach X:" paragraphs (2–3 paragraphs)

Each body paragraph opens with one of the role's key demands as a colon-led header phrase, then carries the **dual anchor**: the real problem, the intended first-90-days action, and the proof. Examples of openers:
- "How I'd approach POC-to-production: …"
- "How I'd approach the workshop side: …"
- "How I'd approach lease restructuring under board oversight: …"

**Construction rules — each paragraph has three moves, in order:**
1. **Name the real problem** the demand represents, traceable to a verified source or the JD (not invented). One sentence.
2. **State the intended action** the candidate would take in the first 90 days, scoped to that problem.
3. **Ground it in one concrete past artifact**: a named system/project/program (with timeframe) and at least one quantity (users, dollars, headcount, percentage, term length, commits, adoption rate).

Rules:
- **One paragraph = one problem = one action = one proof.** Not a list of three.
- Future-tense intent is allowed here **only when both anchors are present** (real problem + concrete proof). A future-tense sentence with no proof behind it, or aimed at a problem the research cannot support, is banned (see §5a and the banned list).
- Do not paraphrase resume bullets. The proof adds what the bullet cannot: the constraint, the trade-off, the through-line to the action.
- Each paragraph must connect to a problem named in the context paragraph (4.2) or the JD. If it cannot connect to evidence, cut it.
```

- [ ] **Step 6: Convert the close (§4.6) to the layered horizon**

Replace the entire `#### 4.6 Forward-looking close (3–4 lines)` section with:

```markdown
#### 4.6 Layered-horizon close (3–4 lines)

Close on the plan's horizon, layered:
1. One sentence on the **first-90-days throughline** — the single thing the candidate's early actions add up to.
2. One sentence on the **6–12 month arc** — where that early work leads, tied to the role's near-term mandate.
3. A confident, specific **ask**. Not a hope, not a thank-you, not a contact-info restatement. Example: "I would like to be inside the room when those decisions are made."

**Banned in the close:**
- "Thank you for your consideration"
- "I look forward to hearing from you"
- "I would welcome the opportunity to discuss"
- "Please feel free to contact me"
- Any restatement of contact info that already appears in the header
```

- [ ] **Step 7: Narrow the future-tense rule in Voice (§5a)**

In `### 5a. Voice and Style`, under `**Voice:**`, replace the bullet:

```markdown
- **Declarative, first-person, confident.** No future-tense self-promises. No hedge framing.
```

with:

```markdown
- **Declarative, first-person, confident.** No hedge framing. Future-tense intent is the point of this mode, but it is allowed **only when dual-anchored**: the action must address a real problem (traceable to a verified source or the JD) AND rest on a concrete past proof point. Generic or ungrounded future promises ("I would bring my passion," "I would contribute my dedication," any value-proposition claim with nothing behind it) remain banned, as does any action aimed at a problem the research cannot support.
```

- [ ] **Step 8: Update the banned-constructions list for forward mode**

In the `### Banned Constructions` section, under `**Banned phrases:**`, replace the line:

```markdown
- "I would bring," "I would contribute," "I would welcome" (any future-tense self-promise)
```

with:

```markdown
- "I would bring," "I would contribute," "I would welcome" — **only when ungrounded.** In this forward mode a future-tense statement is permitted when it is dual-anchored (real evidenced problem + concrete past proof). A future-tense statement with no proof anchor, or aimed at a problem the research cannot support, is banned.
- Speculative actions: any proposed action whose target problem is not traceable to a verified primary source or the JD.
```

Leave the `"is learnable in [timeframe]"` / gap-trivializing ban exactly as is — it stays in force.

- [ ] **Step 9: Replace the §4.4 reference in the honest-limitation note about future tense**

In `#### 4.4` we removed the "No future-tense promises" line by rewriting the section, but the honest-limitation section (§4.5) is unchanged and still correct. No edit needed in §4.5. Confirm §4.5 still reads "What I do not bring is X. What I do bring is rarer: Y." and still bans trivializing the gap.

Run:
```bash
cd /home/reggiechan/JobOps
grep -n 'What I do not bring is X' plugins/jobops/agents/step4-cover-letter-forward.md
grep -n 'is learnable in' plugins/jobops/agents/step4-cover-letter-forward.md
```
Expected: both match (honest-limitation move intact; gap-trivializing ban intact).

- [ ] **Step 10: Add the dual-anchor check to the Step 6a sub-agent review**

In `### 6a. Sub-Agent Sentence Review`, under `**What the reviewer is specifically hunting for:**`, append two bullets:

```markdown
- **Forward claims missing a proof anchor** — any "what I'd do" sentence with no named, quantified past artifact behind it. CUT.
- **Forward claims aimed at an unverifiable problem** — any proposed action whose target problem is not traceable to a verified primary source (per the `primary_sources` ledger) or the JD. CUT.
```

- [ ] **Step 11: Replace the gold-standard exemplar (§5b) with a forward exemplar**

Replace the prose letter inside `### 5b. Gold-Standard Exemplar` (the block-quoted John Smith letter, from `> John Smith, CFA, FRICS` through `> John Smith, CFA, FRICS` at the signature) with the forward exemplar below. Keep the surrounding framing sentences but update them to describe forward structure:

```markdown
> John Smith, CFA, FRICS
> Toronto, ON | (555) 555-0123 | john.smith@example.com | LinkedIn: /in/johnsmith | GitHub: /johnsmith
>
> May 28, 2026
>
> ABC Inc.
>
> Dear John:
>
> I'm applying for the Associate Director, Customer Success and Innovation role. I have spent my career taking technical requirements all the way to running software that non-technical colleagues actually use: internal-facing systems, plain-language query tools, and the workshops that turn leaders into hands-on operators. Enterprise AI governance at this scale is new to me; the translation work between business demand and delivery is not.
>
> By design, ABC Inc. has front-loaded the hard parts. The demand, the governance, and the delivery capacity already exist, and oversight of AI now sits at the Board-committee level against a policy that draws a hard line between Custom and Public AI. What does not arrive with capacity is translation. The binding constraint is no longer capability but the customer-facing function inside the technology group that turns business-unit demand into governed, delivered tools. That is this role.
>
> [First-90-Days Plan table: each row is problem → first-90-days action → proof. E.g. "Internal AI requests stall at proof-of-concept (no owned path to production)" → "Stand up a triage-to-production lane and move two POCs to live tools in the first quarter" → "relationship-intelligence system, 3,000+ commits, plain-language queries colleagues run daily"; "Build-vs-buy decisions lack a Custom-versus-Public test" → "Publish a one-page decision rule mapped to the AI Governance Policy" → "enterprise SaaS rollout, 90% adoption month 1".]
>
> How I'd approach POC-to-production: the problem the policy creates is that internal AI requests stall between a working demo and a governed, production tool, because no one owns the path between them. In the first 90 days I would stand up a single triage-to-production lane and move two live requests through it end to end. I can do this because I have already built the kind of system this role describes: my relationship-intelligence system (3,000+ commits) lets a non-technical user query a database in plain language and get structured answers back, the same pattern ABC Inc. would use to seat AI beside a property asset manager. I take ideas to working software people use, not slideware, which is exactly where most internal AI programs stall.
>
> How I'd approach the workshop side: the gap once tools exist is adoption, and the JD names enablement as core to the function. In the first quarter I would run a builder workshop for one business unit and convert its analysts from prompt-readers into agent-builders. At JKL Corp. (2022 to 2024) I designed and delivered an AI-powered onboarding program (600+ pages of content and twenty fifteen-minute audio episodes, built in two months) that cut new-hire ramp-up by 50%. The methodology is portable. The same architecture runs against ABC Inc. teams inside their own domains.
>
> What I do not bring is enterprise-scale data-warehouse experience; I own the specification, schema, and query architecture, and pair-program the build with Claude Code and Codex. What I do bring is rarer: institutional CRE fluency at platform depth, plus a track record of taking requirements all the way to running systems. That is the exact intersection where the AI Governance Policy now needs translating into customer-facing work.
>
> The first 90 days are about proving the triage-to-production lane on two real requests; the first year is about making it the default way the technology group turns business demand into governed tools. I would like to be inside the room when those decisions are made.
>
> Sincerely,
> [signature image]
> John Smith, CFA, FRICS
```

Then, immediately after the exemplar, replace the `**What this exemplar does that the model must imitate:**` bullet list so the forward-specific bullets are accurate. Update at minimum these three bullets:

```markdown
- **First-90-Days Plan table (4.3):** three columns — problem (evidenced) → first-90-days action → proof. Every problem traces to a verified source or the JD; every proof cell has a named system and a quantity.
- **"How I'd approach X:" paragraphs (4.4):** each carries the dual anchor in order — name the real problem, state the first-90-days action, ground it in one named, quantified past artifact. Future tense is earned by the proof, never floated alone.
- **Layered-horizon close (4.6):** first-90-days throughline, then the 6–12 month arc ("the first year is about…"), then the confident specific ask. No thank-you, no "look forward to," no contact restatement.
```

- [ ] **Step 12: Update the Step 6 quality checks and regression self-check for the forward structure**

In `### 6. Quality Checks` and the `#### Regression self-check`, update the structure and table checklist items to match the forward elements. Specifically replace the structure checklist item:

```markdown
- ✓ **Structure (4.0–4.7):** contact header → fit-led opening → context and role reframe → Requirements Alignment table → "On X:" evidence paragraphs (2–3) → honest-limitation paragraph → forward-looking close → signature with post-nominals. All elements present in order.
```

with:

```markdown
- ✓ **Structure (4.0–4.7):** contact header → fit-led opening → context and role reframe → First-90-Days Plan table → "How I'd approach X:" paragraphs (2–3) → honest-limitation paragraph → layered-horizon close → signature with post-nominals. All elements present in order.
```

And replace the Requirements-table checklist item:

```markdown
- ✓ **Requirements Alignment table:** no more than 5 rows; row construction varies; every evidence cell has a named entity and at least one quantity.
```

with:

```markdown
- ✓ **First-90-Days Plan table:** no more than 5 rows; row construction varies; every problem cell traces to a verified source or the JD; every proof cell has a named entity and at least one quantity.
```

And replace the "On X:" checklist item:

```markdown
- ✓ **Each "On X:" paragraph** opens with a colon-led role-demand phrase, carries exactly one artifact with a named system/project and at least one quantity, ties to the context paragraph's binding constraint, and does not paraphrase a resume bullet.
```

with:

```markdown
- ✓ **Each "How I'd approach X:" paragraph** carries the dual anchor in order: real problem (evidenced) → first-90-days action → one named, quantified past artifact. No forward claim lacks a proof anchor; no action targets a problem the research cannot support.
```

And replace the forward-looking-close checklist item:

```markdown
- ✓ **Forward-looking close** names a specific near-term phase of the firm's work and ends with a confident specific ask. No "thank you for your consideration," no "I look forward to hearing from you," no contact-info restatement.
```

with:

```markdown
- ✓ **Layered-horizon close** shows the first-90-days throughline, then one line on the 6–12 month arc, then a confident specific ask. No "thank you for your consideration," no "I look forward to hearing from you," no contact-info restatement.
```

- [ ] **Step 13: Verify all forward content is present and retrospective artifacts are gone**

Run:
```bash
cd /home/reggiechan/JobOps
F=plugins/jobops/agents/step4-cover-letter-forward.md
echo "--- present (each should be >=1) ---"
grep -c 'First-90-Days Plan' $F
grep -c "How I'd approach" $F
grep -c 'Layered-horizon close' $F
grep -c 'dual-anchor' $F
grep -c 'first 90 days' $F
echo "--- removed (each should be 0) ---"
grep -c 'Requirements Alignment table (mandatory)' $F
grep -c '"On X:" evidence paragraphs' $F
grep -c 'No future-tense self-promises' $F
```
Expected: every "present" count ≥ 1; every "removed" count = 0.

- [ ] **Step 14: Validate plugin structure**

Run:
```bash
cd /home/reggiechan/JobOps && claude plugin validate plugins/jobops && npm test
```
Expected: both pass.

- [ ] **Step 15: Commit**

```bash
cd /home/reggiechan/JobOps
git add plugins/jobops/agents/step4-cover-letter-forward.md
git commit -m "feat(agent): forward-mode structure, dual-anchor rules, and exemplar

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 4: Mode resolution, interview, and routing in `/coverletter`

**Files:**
- Modify: `plugins/jobops/skills/coverletter/SKILL.md` (Arguments ~lines 46–49, Configuration ~lines 7–17, Step 3 dispatch ~lines 71–96)

- [ ] **Step 1: Document the `--mode` argument**

In `plugins/jobops/skills/coverletter/SKILL.md`, in the `## Arguments` section, after the `$3` line, add:

```markdown
- `--mode=retrospective|forward` (optional): overrides `config.preferences.cover_letter_mode` for this invocation. Invalid values are rejected with a message listing the two valid values; the skill does not silently fall back.
```

- [ ] **Step 2: Add a mode-resolution + interview section before the dispatch step**

In the `## Configuration` section, after the existing `Use config.preferences.cultural_profile …` lines, add:

```markdown

### Cover letter mode

Resolve the **effective mode**: the `--mode=` flag if present and valid, else `config.preferences.cover_letter_mode`, else `retrospective` (a config written before this key existed has no `cover_letter_mode`; treat the absence as `retrospective`). Reject an invalid `--mode=` value with: `Invalid --mode value. Use retrospective or forward.`

- **retrospective** → dispatch the `step4-cover-letter` agent (Step 3 below), unchanged.
- **forward** → run the forward intake interview, then dispatch the `step4-cover-letter-forward` agent.

#### Forward intake interview (mandatory for forward mode)

Forward mode never drafts without this interview, because the agent runs non-interactively and cannot ask the candidate anything mid-run. Before dispatch:

1. **Prime the problem set cheaply.** Read the JD ($2) and, if it exists, the specialist files under `{config.directories.company_intelligence}/{Company}/` (`corporate.md`, `legal.md`, `leadership.md`, `market.md`). Do not run web searches here — the agent's Step 3a pipeline does the deep verification. Summarize the candidate-facing problem set in 2–4 bullets so the candidate reacts to evidence rather than inventing problems.
2. **Ask the five questions, one at a time**, in the main conversation:
   1. **Role thesis / problem set** — given the primed bullets, what is this role actually there to solve?
   2. **First-90-days actions** — the 2–3 things you'd do first to address that problem set, plus one line on the 6–12 month arc.
   3. **Per-action proof** — for each action, the past work that backs your ability to do it.
   4. **Gap & plan** — the real gap you carry and how you'd work with it (without trivializing it as quickly closeable).
   5. **Company-specific notes** — anything about the firm's situation to reflect in the context paragraph or problem set.
3. **If the candidate declines or abandons the interview, do not draft.** Tell them forward mode requires the interview, and that retrospective mode (`--mode=retrospective`) produces a letter without one.
4. **Pass the answers and the primed problem set inline** to the `step4-cover-letter-forward` agent in its dispatch prompt.
```

- [ ] **Step 3: Make the Step 3 dispatch conditional on mode**

In `## Step 3: Generating Cover Letter`, replace the paragraph that begins `I'm launching the \`step4-cover-letter\` agent …` with:

```markdown
I dispatch the agent that matches the effective mode resolved in Configuration:

- **retrospective** → the `step4-cover-letter` agent: a contact header followed by seven body elements (fit-led opening, context/role reframe, Requirements Alignment table, "On X:" evidence paragraphs, honest-limitation, forward-looking close, signature), every claim traced to the Step 3 resume and verified primary sources.
- **forward** → the `step4-cover-letter-forward` agent, with the intake-interview answers and primed problem set passed in. It proposes a **dual-anchored first-90-days plan**: a First-90-Days Plan table (problem → action → proof) and "How I'd approach X:" paragraphs, where every proposed action sits on a real evidenced problem and a concrete past proof point. The fit-led opening, context paragraph, honest-limitation, voice discipline, primary-source verification, and Step 6a sub-agent review are shared with retrospective mode.
```

- [ ] **Step 4: Verify content present and consistent**

Run:
```bash
cd /home/reggiechan/JobOps
F=plugins/jobops/skills/coverletter/SKILL.md
grep -c -- '--mode=' $F
grep -c 'effective mode' $F
grep -c 'Forward intake interview' $F
grep -c 'step4-cover-letter-forward' $F
grep -c 'step4-cover-letter\b' $F
```
Expected: `--mode=` ≥ 2, `effective mode` ≥ 1, `Forward intake interview` ≥ 1, `step4-cover-letter-forward` ≥ 2, retrospective agent name still referenced ≥ 1.

- [ ] **Step 5: Validate**

Run:
```bash
cd /home/reggiechan/JobOps && claude plugin validate plugins/jobops && npm test
```
Expected: both pass.

- [ ] **Step 6: Commit**

```bash
cd /home/reggiechan/JobOps
git add plugins/jobops/skills/coverletter/SKILL.md
git commit -m "feat(coverletter): resolve mode, run forward interview, route to agent

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 5: End-to-end functional verification (manual)

No automated harness exercises the prompt logic, so verify the three routing paths by reading the resolved instructions and confirming they are unambiguous. This task produces a short verification note, not code.

**Files:**
- (none modified; verification only)

- [ ] **Step 1: Trace the three mode paths against the skill text**

Read `plugins/jobops/skills/coverletter/SKILL.md` and confirm, for each scenario, exactly one path is determined with no ambiguity:
1. **No `--mode`, config `cover_letter_mode: forward`** → interview runs, `step4-cover-letter-forward` dispatched.
2. **`--mode=retrospective`, config `cover_letter_mode: forward`** → no interview, `step4-cover-letter` dispatched (flag wins).
3. **No `--mode`, config has no `cover_letter_mode` key** → no interview, `step4-cover-letter` dispatched (absence → retrospective).
4. **`--mode=sideways`** → rejected with the invalid-value message.

- [ ] **Step 2: Confirm the forward agent's stop condition for a missing interview**

Read `plugins/jobops/agents/step4-cover-letter-forward.md` §1 and confirm it halts and reports when interview answers are absent, so a direct/mis-wired dispatch cannot silently produce an unanchored letter.

- [ ] **Step 3: Confirm the dual-anchor guardrail is enforced in three places**

Confirm the "no speculative action / proof-anchor-required" rule appears in: §3a Step 5 (anchoring), §4.4 (construction rules), and §6a (reviewer check). All three must agree.

- [ ] **Step 4: Record the verification note and commit**

Append a dated note to the spec's "Testing" expectations or create `docs/superpowers/plans/2026-06-01-forward-facing-cover-letter-mode-verification.md` with the four traced paths and their outcomes. Then:

```bash
cd /home/reggiechan/JobOps
git add docs/superpowers/plans/2026-06-01-forward-facing-cover-letter-mode-verification.md
git commit -m "docs: forward cover-letter mode routing verification note

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 6: Documentation and contributor note

**Files:**
- Modify: `docs/ARCHITECTURE.md` (§2 Config file, ~line 30)
- Modify: `README.md` (cover letter row, ~line 89)
- Modify: `CLAUDE.md` (Coding Style or a new note)

- [ ] **Step 1: Document the config key in ARCHITECTURE.md**

In `docs/ARCHITECTURE.md` §2 "Config file", after the `config.candidate` paragraph (line 30), add:

```markdown
`config.preferences.cover_letter_mode` (`retrospective` | `forward`, default `retrospective`) selects how `/jobops:coverletter` writes the letter. A config written before this key existed has no `cover_letter_mode`; consumers treat the absence as `retrospective`. `/jobops:coverletter --mode=` overrides it per-invocation. `forward` mode runs a mandatory skill-level intake interview and dispatches `step4-cover-letter-forward` instead of `step4-cover-letter`; both agents share voice, provenance, and sub-agent-review rules.
```

- [ ] **Step 2: Mention forward mode in README**

In `README.md`, update the cover letter command row (line 89). Replace:

```markdown
| `/jobops:coverletter` | Strategic cover letter with requirements-matching table |
```

with:

```markdown
| `/jobops:coverletter` | Strategic cover letter — retrospective (requirements→proof) or forward mode (`--mode=forward`: dual-anchored first-90-days plan via a short interview), set by `preferences.cover_letter_mode` |
```

- [ ] **Step 3: Add the rule-drift contributor note to CLAUDE.md**

In `CLAUDE.md`, under `## Coding Style`, add a bullet:

```markdown
- `step4-cover-letter.md` (retrospective) and `step4-cover-letter-forward.md` (forward) are a paired set. Shared, mode-agnostic rules (contact header, Step 3a primary-source pipeline, voice/style §5a, banned-construction list, Step 6a sub-agent review) are reproduced verbatim in both. Any edit to a shared rule must be applied to BOTH files.
```

- [ ] **Step 4: Verify and validate**

Run:
```bash
cd /home/reggiechan/JobOps
grep -c 'cover_letter_mode' docs/ARCHITECTURE.md
grep -c 'mode=forward' README.md
grep -c 'paired set' CLAUDE.md
claude plugin validate plugins/jobops && npm test
```
Expected: each grep ≥ 1; validators pass.

- [ ] **Step 5: Commit**

```bash
cd /home/reggiechan/JobOps
git add docs/ARCHITECTURE.md README.md CLAUDE.md
git commit -m "docs: document cover_letter_mode config key and agent pairing

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 7: Version bump and changelog

**Files:**
- Modify (via version-bump skill): `package.json`, `plugins/jobops/.claude-plugin/plugin.json`, `plugins/jobops/.codex-plugin/plugin.json`, `plugins/jobops-ic/.claude-plugin/plugin.json`, `plugins/jobops-ic/.codex-plugin/plugin.json`, `.claude-plugin/marketplace.json`, `.agents/plugins/marketplace.json`, `README.md`, `CHANGELOG.md`

- [ ] **Step 1: Run the version-bump skill**

This repo has a dedicated `version-bump` skill that updates every version file and the changelog scaffold consistently. Invoke it for a **minor** bump (new backward-compatible feature): `2.12.0 → 2.13.0`. Provide the changelog entry text below. If the skill is unavailable, edit the nine files listed above by hand to `2.13.0`.

- [ ] **Step 2: Write the CHANGELOG entry**

In `CHANGELOG.md`, under `## [Unreleased]` (or the new `## [2.13.0]` section the skill creates), add:

```markdown
### Added

- **Forward-facing cover letter mode (`/jobops:coverletter --mode=forward`, or `preferences.cover_letter_mode: forward`)** — an alternate to the default retrospective letter. Instead of matching the job's requirements to past work, forward mode proposes a **dual-anchored first-90-days plan**: every proposed action sits on a real problem (traceable to a verified primary source or the JD, never speculation) and a concrete past proof point. The `/coverletter` skill runs a mandatory five-question intake interview (primed cheaply from the JD and existing OSINT) before dispatching a new `step4-cover-letter-forward` agent. The agent reuses the retrospective letter's voice discipline, primary-source verification, and independent sub-agent review, and adds a dual-anchor reviewer check that cuts any forward claim lacking a proof anchor or aimed at an unverifiable problem. The table becomes a First-90-Days Plan (problem → action → proof), body paragraphs become "How I'd approach X:", and the close layers a first-90-days throughline with the 6–12 month arc. Retrospective remains the default; configs lacking the new key behave exactly as before.
```

- [ ] **Step 3: Verify versions are consistent**

Run:
```bash
cd /home/reggiechan/JobOps
grep -RH '"version": "2.13.0"' package.json plugins/*/.claude-plugin/plugin.json plugins/*/.codex-plugin/plugin.json .claude-plugin/marketplace.json .agents/plugins/marketplace.json | wc -l
grep -c '2.13.0' CHANGELOG.md
claude plugin validate plugins/jobops && claude plugin validate plugins/jobops-ic && claude plugin validate . && npm test
```
Expected: the version count is 7 (all version files at 2.13.0); CHANGELOG has the new version; all validators and `npm test` pass.

- [ ] **Step 4: Commit**

```bash
cd /home/reggiechan/JobOps
git add -A
git commit -m "chore(release): v2.13.0 — forward-facing cover letter mode

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Self-Review

**Spec coverage:**
- Config schema + setup question → Task 1. ✓
- Mode resolution (flag > config > default) → Task 4 Step 2. ✓
- Mandatory skill-level interview (5 questions, priming) → Task 4 Step 2. ✓
- Interview/agent sequencing contract → Task 3 Step 3 (§3a Step 5) + Task 4 Step 2. ✓
- New forward agent, shared rules reproduced → Tasks 2–3. ✓
- Dual-anchor model (problem + proof) → Task 3 Steps 4, 5, 7, 8, 10. ✓
- Layered horizon → Task 3 Steps 6, 11, 12. ✓
- Forward exemplar → Task 3 Step 11. ✓
- Narrowed future-tense ban → Task 3 Steps 7, 8. ✓
- Gap-trivializing ban retained → Task 3 Step 9. ✓
- Dual-anchor reviewer check → Task 3 Step 10. ✓
- Quality-check/regression updates → Task 3 Step 12. ✓
- Skill routing → Task 4 Step 3. ✓
- Edge cases (invalid flag, declined interview, unverifiable problem, <2 sources) → Task 4 Steps 1–2 + Task 3 Step 3 + Task 5. ✓
- Docs (ARCHITECTURE, README, CLAUDE.md drift note) → Task 6. ✓
- Version + changelog → Task 7. ✓

**Placeholder scan:** No "TBD"/"implement later". Every content step shows the exact markdown/JSON to insert. The forward exemplar is fully written. ✓

**Type/name consistency:** Agent name `step4-cover-letter-forward` is identical in frontmatter (Task 2), skill routing (Task 4), docs (Task 6), and changelog (Task 7). Config key `cover_letter_mode` and flag `--mode=` are spelled identically across Tasks 1, 4, 6, 7. Section names (First-90-Days Plan table, "How I'd approach X:", Layered-horizon close) match between the agent edits (Task 3) and the skill/docs descriptions. ✓
