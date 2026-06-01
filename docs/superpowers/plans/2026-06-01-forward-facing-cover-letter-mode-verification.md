# Forward Cover-Letter Mode - End-to-End Routing Verification

**Purpose:** Confirm that the four routing scenarios in `plugins/jobops/skills/coverletter/SKILL.md` resolve unambiguously, the forward agent's stop condition guards against mis-wired dispatch, and the dual-anchor guardrail is consistent across the three places it appears in `plugins/jobops/agents/step4-cover-letter-forward.md`.

**Branch:** `feat/forward-cover-letter-mode`
**Verified on:** 2026-06-01

---

## 1. Routing Scenario Table

| # | Scenario | Expected outcome | Verdict | SKILL.md line(s) that determine it |
|---|----------|-----------------|---------|-------------------------------------|
| 1 | No `--mode` flag; `config.preferences.cover_letter_mode: forward` | Forward intake interview runs; `step4-cover-letter-forward` dispatched | PASS | Line 20 (flag absent - fall to config; config = `forward`); line 23 (`forward` branch) |
| 2 | `--mode=retrospective`; `config.preferences.cover_letter_mode: forward` | No interview; `step4-cover-letter` dispatched (flag wins over config) | PASS | Line 20 (flag present and valid - flag wins); line 22 (`retrospective` branch) |
| 3 | No `--mode` flag; `cover_letter_mode` key absent from config | No interview; `step4-cover-letter` dispatched (absence defaults to `retrospective`) | PASS | Line 20 ("else `retrospective` (a config written before this key existed has no `cover_letter_mode`; treat the absence as `retrospective`)"); line 22 (`retrospective` branch) |
| 4 | `--mode=sideways` (invalid value) | Rejected immediately with message: `Invalid --mode value. Use retrospective or forward.` | PASS | Line 20 ("Reject an invalid `--mode=` value with: `Invalid --mode value. Use retrospective or forward.`") |

All four scenarios resolve to exactly one outcome with no ambiguity. The resolution order is explicit and unambiguous: (1) flag if present and valid, (2) config key if present, (3) `retrospective` as the final default. An invalid flag value is rejected before the config is consulted, so the config can never rescue a bad flag value.

---

## 2. Forward Agent Stop Condition (section 1 - Input Validation)

**File:** `plugins/jobops/agents/step4-cover-letter-forward.md`

**Line 25:** "**Forward mode does not run without these.** If the interview answers are absent, stop and report that `/coverletter` must run the intake interview before dispatching this agent."

This stop condition is correctly placed in section 1 (Input Validation), before any drafting work begins. A direct or mis-wired dispatch that omits the intake-interview answers will cause the agent to halt and surface an actionable error - it cannot silently produce an unanchored letter. The stop message directs the user back to the skill entry point (`/coverletter`), which is the only correct way to trigger a forward-mode dispatch.

**Verdict: CONFIRMED.**

---

## 3. Dual-Anchor Guardrail - Three-Place Consistency Check

The "no speculative action / proof-anchor-required" rule must appear and agree in three places. Evidence:

### 3a. Section 3a Step 5 - Anchoring each action to a verified problem or the JD

**Lines 85-87:** "An action whose target **problem** traces to a `verified` primary source (or is stated plainly in the JD) is anchored - it may appear in the table and body. An action whose problem traces to **neither** a verified source nor the JD is **speculation**. Do not ship it."

Rule stated: every proposed action requires an evidenced-problem anchor (verified source or JD). Speculation is dropped or reframed. The second anchor (past proof) is addressed in the table constraint at lines 83-89 where proof cells require a named entity and quantity.

### 3b. Section 4.4 - Construction rules for "How I'd approach X:" paragraphs

**Lines 196-197:** "Future-tense intent is allowed here **only when both anchors are present** (real problem + concrete proof). A future-tense sentence with no proof behind it, or aimed at a problem the research cannot support, is banned (see §5a and the banned list)."

Rule stated: both anchors required - real evidenced problem AND concrete past proof. Either anchor absent = banned. This is the explicit dual-anchor statement that names both conditions.

### 3c. Section 6a - Reviewer dual-anchor CUT checks

**Lines 481-482:** "**Forward claims missing a proof anchor** - any 'what I'd do' sentence with no named, quantified past artifact behind it. CUT." and "**Forward claims aimed at an unverifiable problem** - any proposed action whose target problem is not traceable to a verified primary source (per the `primary_sources` ledger) or the JD. CUT."

Rule stated: reviewer independently enforces both anchor conditions at the sentence level. Each anchor failure is its own CUT category.

### Consistency assessment

All three placements are mutually consistent:
- 3a enforces the evidenced-problem anchor at the drafting stage (before the table and body paragraphs are written).
- 4.4 enforces both anchors simultaneously during paragraph construction.
- 6a enforces each anchor independently at review, catching any drafting-stage miss before output is written to disk.

No contradictions. The three checks form a layered defence: draft-time anchor binding (3a), construction rule (4.4), and independent sentence-level review (6a).

**Verdict: CONFIRMED - three placements present and mutually consistent.**

---

## 4. Overall Verdict

**PASS.** All four routing scenarios resolve unambiguously. The stop condition guards correctly. The dual-anchor guardrail is present and consistent across all three enforcement points. No ambiguities, contradictions, or silent failure paths found.
