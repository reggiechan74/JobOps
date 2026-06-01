# Forward-Facing Cover Letter Mode — Design

**Date:** 2026-06-01
**Status:** Approved for planning
**Affects:** `plugins/jobops` — `step4-cover-letter` agent, `coverletter` skill, `setup` skill, config contract

## Problem

The current cover letter is **backward-looking**: it matches the major job-description
requirements against what the candidate has previously done. The `step4-cover-letter`
agent enforces this throughout — most explicitly with a hard ban on future-tense
self-promises ("I would bring," "I would contribute," "I would welcome") and a
requirement that every "On X:" paragraph carry a past-tense, named, quantified
accomplishment.

We want an **alternate forward-facing mode**, selectable via the user's config file, in
which the candidate states what he is **going to do** going forward, scoped to the
role's requirements. Because forward intent is not derivable from the resume alone, this
mode requires interviewing the user for specifics before drafting.

## Goals

- Add a forward-facing cover letter mode alongside the existing retrospective mode.
- Make the mode selectable via `.jobops/config.json`, with a per-invocation override.
- Keep forward statements **credible** — anchored to past proof, not aspirational.
- Gather the forward specifics through a structured interview at the skill level.
- Preserve everything that makes the retrospective letter strong (voice discipline,
  provenance verification, sub-agent review) in the new mode.

## Non-goals

- No persisted interview artifact in this iteration (the interview runs each forward
  invocation; persistence may be added later if re-runs prove tedious).
- No change to the retrospective mode's behavior or output. It remains the default.
- No new output path or filename — forward mode writes to the same
  `cover-letter/cover_letter.md`.

## Decisions (locked during brainstorming)

| Decision | Choice |
|----------|--------|
| Credibility model | **Evidence-anchored forward** — each forward commitment is grounded in a past proof point, then pivots to intent. |
| Interview flow | **Mandatory skill-level interview** — the `/coverletter` skill runs a structured interview in the main conversation before dispatch; forward mode never runs without it. |
| Mode selection | **Config default + flag override** — `preferences.cover_letter_mode` sets the default; `--mode=` overrides per-invocation. |
| Code structure | **Separate agent file** — a new `step4-cover-letter-forward.md`; the skill routes to the correct agent by mode. |

## Architecture

### Component 1 — Config schema

Add one key to `preferences` in `.jobops/config.json`:

```jsonc
"preferences": {
  "cultural_profile": "...",
  "default_jurisdiction": "...",
  "cover_letter_mode": "retrospective"   // "retrospective" (default) | "forward"
}
```

- **Backward compatibility:** configs written before this change have no
  `cover_letter_mode` key. Consumers treat a missing key as `retrospective`. No
  migration is required.
- `/jobops:setup` Step 4 (Preferences interview) gains a third question:
  - **Cover letter mode** — enum `retrospective` | `forward`. Default `retrospective`.
    One line of help: retrospective matches requirements to past work; forward states
    intended actions scoped to requirements and requires a short interview each time.
- `/jobops:setup` Step 6 schema gains the `cover_letter_mode` key.

### Component 2 — Mode resolution in `/coverletter`

- Add a `--mode=retrospective|forward` flag to the skill's argument handling (mirrors the
  existing `--jurisdiction=` override pattern used by crisis skills).
- **Effective mode** = `--mode` flag if present, else `config.preferences.cover_letter_mode`,
  else `retrospective`.
- After resolving the effective mode, the skill routes to the matching agent:
  - `retrospective` → dispatch `step4-cover-letter` (unchanged path).
  - `forward` → run the intake interview (Component 3), then dispatch
    `step4-cover-letter-forward` with the interview answers included in the prompt.

### Component 3 — Forward intake interview (skill-level, mandatory)

When the effective mode is `forward`, `/coverletter` conducts a structured interview in
the main conversation **before** dispatching the agent. (The agent runs as a
non-interactive sub-agent and cannot ask the user anything mid-run, so intake must happen
at the skill level.)

Fixed question set:

1. **Role thesis** — what is this role actually there to solve? The binding constraint,
   in the candidate's own words.
2. **First priorities** — the 2–3 things the candidate would tackle first, and roughly in
   what order or horizon (e.g., first 90 days / first 12 months).
3. **Per-requirement approach** — for each of the top requirements, how the candidate
   would approach it *and* the past work that backs that approach. This is the
   evidence-anchoring input; without it the letter would be aspirational.
4. **Gap & plan** — the real gap the candidate carries and how he would work around or
   with it, **without** trivializing it as quickly closeable.
5. **Company-specific notes** — anything about the firm's situation the candidate wants
   reflected in the context paragraph.

The skill collects the answers and passes them inline to the forward agent in its
dispatch prompt. Answers are not persisted to disk in this iteration.

### Component 4 — New agent `step4-cover-letter-forward.md`

A self-contained sibling to `step4-cover-letter.md`. Because agents do not reliably read
sibling agent files at runtime, all shared, mode-agnostic rules are **reproduced
verbatim** in the new file rather than referenced. The cost is duplication and a drift
risk between the two files; that risk is accepted in exchange for each agent being
fully self-contained. (The implementation plan should note that future edits to shared
rules must be applied to both files.)

**Reproduced verbatim from the retrospective agent:**

- Contact header (§4.0), sourced from `config.candidate`.
- Step 3a primary-source acquisition and verification pipeline (read OSINT specialist
  files → role-targeted WebSearch → WebFetch verification → context-mode decision).
- Voice and style (§5a) **except** the future-tense rule (see below).
- Banned-construction list **except** the future-tense self-promise line (see below).
- Step 6a sub-agent sentence review (What / So What / Now What), plus one added forward
  check (see below).
- Quality checks (§6) and the regression self-check, adapted to the forward structure.

**Forward-specific structure** (the elements that differ from retrospective):

| Element | Retrospective | Forward |
|---------|---------------|---------|
| Opening (fit-led) | Names role, leads with candidate's record, honest pivot | **Unchanged** |
| Context & role reframe | Verified-source synthesis, binding constraint | **Unchanged** |
| Requirements element | **Requirements Alignment** table: requirement → past evidence | **Forward Priorities** table: requirement → *intended approach*, with a proof anchor in each cell (named entity + quantity required). Still capped at 5 rows. |
| Body paragraphs | "On X:" — one past-tense named, quantified artifact each | **"How I'd approach X:"** — each opens with a role demand, states the intended action scoped to that requirement, then grounds it in one concrete past artifact (named system + quantity). Intent + proof, never intent alone. |
| Honest-limitation | gap → rarer strength → tie to binding constraint | **Unchanged in structure**; gap-trivializing ban **stays** (no "learnable in N weeks"). |
| Forward close | near-term mandate + confident ask | **Unchanged** |
| Signature | Sincerely / image / name + post-nominals | **Unchanged** |

**The critical voice change — future-tense ban is narrowed, not lifted:**

- **Allowed in forward mode:** future-tense intent when it is (a) scoped to a named
  requirement AND (b) anchored to a concrete proof point.
- **Still banned:** ungrounded or generic future promises ("I would bring my passion,"
  "I would contribute my dedication," any value-proposition claim without a proof
  anchor), and trivializing a gap as quickly closeable.
- The Step 6a reviewer gets **one added check**: flag any forward claim that has no proof
  anchor (a promise with nothing behind it is a CUT).

**Agent input:** the forward agent receives, in addition to the Step 3 resume and JD, the
five interview answers from Component 3.

### Component 5 — Skill, docs, and versioning

- `coverletter/SKILL.md`: document the `--mode=` argument, the mode-resolution logic, the
  interview flow for forward mode, and the routing to the matching agent. Add a short
  forward-mode section parallel to the existing structure description.
- `setup/SKILL.md`: Step 4 question and Step 6 schema key.
- `docs/ARCHITECTURE.md`: add `cover_letter_mode` to the config contract; note the
  missing-key-defaults-to-retrospective rule.
- `CHANGELOG.md` and version bump across the four version files (`package.json`, both
  `plugin.json` files, `README.md`) per the repo's version-management convention.
- `README.md`: brief mention of forward mode where cover letters are described.

## Data flow

```
/coverletter $resume $jd [$manager] [--mode=forward]
        │
        ▼
  read .jobops/config.json
        │
   resolve effective mode  ── retrospective ──► dispatch step4-cover-letter ──► cover_letter.md
        │
      forward
        │
        ▼
  skill-level intake interview (5 questions, main conversation)
        │
        ▼
  dispatch step4-cover-letter-forward
     (resume + JD + interview answers)
        │
        ▼
  Step 3a primary-source verification ──► draft ──► Step 6 self-check
        │
        ▼
  Step 6a independent sub-agent review (+ forward proof-anchor check)
        │
        ▼
  cover_letter.md  (same path, with primary_sources YAML ledger)
```

## Error handling & edge cases

- **Missing config key:** treat as `retrospective`. No error.
- **Invalid `--mode` value:** reject with a clear message listing the two valid values;
  do not silently fall back.
- **Forward mode, user declines/abandons the interview:** do not draft. Forward mode is
  defined as never running without the interview. Surface that retrospective mode is
  available if the user wants a letter without the interview.
- **Insufficient verified primary sources (<2):** same fallback as retrospective — leaner
  context paragraph, `primary_sources: []`, skip-condition comment. The fit-led opening
  and forward body are unaffected.
- **Sub-agent review fails twice:** same as retrospective — surface the reviewer report
  rather than ship a weak letter.

## Testing

- `claude plugin validate plugins/jobops` passes after the new agent file is added.
- `npm test` (Codex compatibility contract) passes.
- Manual: run `/coverletter` with no flag and a config set to `forward` → interview runs,
  forward agent dispatched. Run with `--mode=retrospective` against the same config →
  retrospective agent dispatched, no interview. Run with a fresh config lacking the key →
  retrospective.
- Manual review of a generated forward letter: every forward claim has a proof anchor; no
  banned generic promises; gap not trivialized; em-dash count zero; structure complete.

## Open risks

- **Rule drift between the two agent files.** Mitigation: the implementation plan calls
  out that shared-rule edits must touch both files; consider a contributor note in
  CLAUDE.md or a comment header in each agent pointing to its sibling.
- **Interview friction.** If running the interview every forward invocation proves
  tedious, the persisted-artifact option (saving answers to the app folder) is the
  natural next iteration.
