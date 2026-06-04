# /idealjob Overhaul — Design

**Date:** 2026-06-04
**Status:** Approved
**Target version:** v2.14.0

## Goal

Rebuild `/jobops:idealjob` from a batch single-JD synthesizer into an interview-driven, three-archetype career-targeting engine whose outputs are scored, market-validated, and consumable by `/comparejobs`.

## Problems with the current skill

1. **Score-scale bug.** Phase 2.1 greps for `overall_score: 9` (assumes 0–100), but `assessjob` v2.0 writes `overall_score: <XX/200>` + `normalized_score: <XX%>`. High-scorer detection never matches current assessments.
2. **Stale hardcoding.** Search queries say "2025"; "Canada" is hardcoded instead of `config.preferences.default_jurisdiction`; `Sample_Output/` is not part of the config directory contract.
3. **Internal inconsistency.** Phase 1.1 reads `Experience/`; Phase 1.2 says `WorkHistory/*.md`. Canonical structure is `Experience/<Company>_<Role>.md` (setup SKILL.md).
4. **Zero interactivity.** Pure batch synthesis; never asks the user anything.
5. **Generic output.** One averaged JD that mirrors the resume; predicted "95+ fit score" asserted, never tested.
6. **No downstream loop.** Output feeds nothing — not comparejobs, not a search workflow.

## Architecture (Approach A — hybrid)

Inline interview and synthesis in the main skill; three parallel sub-agents for archetype generation.

```
/idealjob [output-file]

Phase 1  Career Intelligence (inline)
Phase 2  Structured Interview (inline, always runs, clarity-gated)
Phase 3  Vision.md Write-back (inline, confirmation-gated)
Phase 4  Dossier Assembly + 3 parallel ideal-role-architect sub-agents
Phase 5  Self-Score & Revise (inline, 200-pt rubric)
Phase 6  Differentiation Check + Assembly (inline)
```

### Phase 1 — Career Intelligence

- Load `Experience/`, `CareerHighlights/`, `Technology/`, `Preferences/Vision.md` from `{resume_source}`.
- Mine assessments: glob `{applications_root}/*/assessment/assessment.md`; select `normalized_score ≥ 90%`. Legacy files without a `/200` score fall back to the old 0–100 `overall_score` interpretation (≥ 90). Read top 3–5; extract success patterns (role characteristics, skill alignment, cultural factors, gap patterns).
- `Sample_Output/` reference removed entirely.

### Phase 2 — Structured Interview

Always runs. Four themed rounds, each **clarity-gated, with no question budget**. Questions go out in batches of up to 4 via AskUserQuestion; a round loops with follow-up batches until every exit criterion is met. The skill states explicitly: *"Do not proceed to the next round until every exit criterion is satisfied. There is no maximum number of questions."*

Each question pre-seeds from Vision.md (current value shown as the first option) so the interview doubles as a Vision.md audit.

**Probing discipline (all rounds):**
- Vague answer → drill down (e.g., "autonomy" over what — methods, priorities, hours, headcount, P&L?).
- Contradiction with Vision.md, a prior answer, or career-inventory evidence → surface verbatim, resolve immediately.
- Concrete over abstract: every preference anchored to ≥1 real episode from the inventory; unanchored preferences tagged `aspirational`, anchored ones `evidenced`.
- Quantify when quantifiable: comp, team size, travel %, commute, on-call — numbers, not adjectives.

**Rounds and exit criteria (round repeats until ALL are true):**

| Round | Exit criteria |
|---|---|
| 1. Energy audit | ≥3 specific energizers + ≥3 specific drainers, each tied to a named role/project from `Experience/`; proudest achievement identified and the *why* articulated; ≥1 "never again" with underlying cause (not surface event) |
| 2. Forced trade-offs | comp-vs-mission, autonomy-vs-mentorship, IC-vs-leadership, stability-vs-stretch all resolved with stated strength of preference (hard requirement vs. lean); every "it depends" decomposed until conditions are explicit |
| 3. Context | Explicit industry in-list AND out-list (empty out-list confirmed, not assumed); company stage, size range, ownership structure (public/PE/VC/family/gov) each addressed; work arrangement with hard limits (max office days, travel %); geography confirmed against `default_jurisdiction` |
| 4. Boundaries & ambition | Every Anti-Vision item confirmed, amended, or retired; ≥1 new deal-breaker probed for; comp floor + target + structure (base/bonus/equity mix) as numbers; pivot curiosity explored with ≥2 adjacent fields rated; 3-year aspiration stated and checked for consistency against everything above |

**Round closure ritual:** the skill plays back a numbered summary ("Here's what I now believe about your [theme]…") and asks for confirmation. A round closes only on confirmed playback; corrections reopen probing.

**Final gate:** cross-round consistency check — the skill must state it found no unresolved contradictions between rounds (or resolve them) before Phase 3.

### Phase 3 — Vision.md Write-back

- Present a summarized diff (added / changed / confirmed) of interview findings vs. current Vision.md.
- Explicit confirmation required before writing. Only confirmed deltas merge; Vision.md's existing section structure is preserved; changed sections get an appended `> Last updated via /idealjob interview, {date}` line.
- Decline → skip write-back, continue (answers still feed this run).
- Vision.md missing → offer to create it.

### Phase 4 — Dossier & parallel dispatch

**Dossier** (assembled inline, passed verbatim to each agent):
- Evidence digest condensed from `Experience/`, `CareerHighlights/`, `Technology/` — every claim cites its source file.
- Interview findings: the four confirmed round playbacks, preferences tagged `evidenced` or `aspirational`.
- Success patterns from mined ≥90% assessments (or a "no success patterns available" note — see error handling).
- Hard boundaries: deal-breakers, comp floor/target/structure, work-arrangement limits.
- Run context: current date and jurisdiction passed explicitly — agents never guess the year or default market.

**Archetype charters:**

| Archetype | Charter |
|---|---|
| Anchor | Highest-probability fit. Role exists in volume in today's market; candidate would be a top-decile applicant. Optimizes `evidenced` preferences only. |
| Stretch | One level up — bigger scope, seniority, or mandate. Gaps permitted only if bridgeable within ~12 months. Draws on `aspirational` tags. |
| Pivot | Skills recombined into an adjacent field, seeded by the interview's pivot-curiosity answers. Novel titles allowed but must be validated against real postings. |

**New agent: `plugins/jobops/agents/ideal-role-architect.md`** (used 3× in parallel). Tasks:
1. Market scan — web searches using the passed-in date and jurisdiction; comp validated against the dossier's floor/target.
2. JD generation — current output skeleton retained; every *Required* qualification/skill cites a dossier evidence item; *Preferred* items may be aspirational; banned-construction list (generic JD-speak: "rockstar," "fast-paced environment," "wear many hats," "self-starter," plus generic-flattery constructions); 1–2 deliberate non-ideal-but-tolerable elements in the role/company (stress-tests deal-breaker boundaries; real jobs are imperfect).
3. Title realism — JD title (or near-variant) must appear in actual postings found during the scan.
4. Live matches — 3–5 real, current postings with links and a 2–3 line fit note each. Never fabricate URLs.
5. Search kit — boolean strings for LinkedIn/Indeed, exact title variants, 10+ target companies, networking-target role types.
6. Returns structured markdown to the main context (raw data, no preamble).

### Phase 5 — Self-score & revision

- Main context scores all three archetypes (one scorer = one consistent scale) against the bundled 200-point `assessment_rubric_framework.md`, selecting the correct role variant (Technical IC / People Manager / Executive).
- Honest scoring: gaps score as gaps.
- Thresholds: **Anchor ≥ 90%** normalized; **Stretch / Pivot ≥ 85%**. Below threshold → one revision cycle (fix the JD where it over-reached its charter; never inflate the score), re-score, accept the result.
- Residual Stretch/Pivot gaps feed the gap-closing plan.
- Final scores land in each archetype file's frontmatter in assessjob-compatible form.

### Phase 6 — Differentiation check & assembly

- Verify archetypes are mutually distinct in title, industry, and value proposition; revise if two converge.
- Write the four output files; build gap-closing plans; emit comparejobs handoff instructions.

## Outputs

Four flat files under `{career_analysis}/` (flatness preserved; documented exception to the single-file convention):

```
idealjob_{YYYYMMDD}.md            summary: interview record (round playbacks),
                                  cross-archetype comparison table (titles, scores,
                                  comp, key trade-offs), gap-closing plans,
                                  comparejobs handoff instructions
idealjob_{YYYYMMDD}_anchor.md     per archetype: assessjob-compatible YAML
idealjob_{YYYYMMDD}_stretch.md    frontmatter (synthetic: true, generated_by,
idealjob_{YYYYMMDD}_pivot.md      archetype, overall_score: XX/200,
                                  normalized_score: XX%), full JD, alignment
                                  analysis, live posting matches, search kit
```

`$1` override: becomes the summary path; archetype files derive from its basename + `_anchor`/`_stretch`/`_pivot` suffixes.

**Gap-closing plan** (summary file, for Stretch + Pivot): each self-score gap gets an action (cert / project / experience), a rough timeline, and which archetype(s) it unlocks.

## comparejobs amendment

`{{ARGn}}` args containing a `/` or ending in `.md` are treated as direct assessment-file paths instead of app slugs, e.g.:

```
/comparejobs Acme_PM_20260601 Career_Analysis/idealjob_20260604_anchor.md
```

The archetype files' assessjob-compatible frontmatter makes them valid comparison inputs.

## Error handling

| Failure | Behavior |
|---|---|
| No assessments ≥ threshold | Degraded mode — dossier notes "no success patterns available," user informed |
| `Preferences/Vision.md` missing | Interview runs regardless; write-back offers to create the file |
| `resume_source` empty/missing | Hard stop with pointer to setup |
| No live postings found for an archetype | Mark "market validation incomplete" — never fabricate links |
| A sub-agent fails | Retry once; then proceed with remaining archetypes and state the omission |
| Write-back declined | Continue without persisting |

## Validation checklist (replaces current)

- Archetypes mutually distinct in title/industry/value-prop
- Every Required item evidence-cited
- Zero banned constructions
- Score thresholds met or revision documented
- One live-posting link per archetype spot-checked via fetch
- All Vision preferences honored; all Anti-Vision elements absent
- Search kits jurisdiction-correct

## Documentation & versioning

- CLAUDE.md: file-naming convention line gains the idealjob multi-file exception
- `docs/ARCHITECTURE.md`: career_analysis output contract updated
- README: skill description updated
- Version bump to **v2.14.0** across `package.json`, plugin.json files, README, CHANGELOG

## Testing

- `claude plugin validate plugins/jobops`
- `npm test` (Codex compatibility contract)
- Manual dry-run against the maintainer's own workspace

## Out of scope

- Full `/assessjob` pipeline runs per archetype (lightweight inline self-score chosen instead)
- Dashboard integration with career_analysis outputs
- Changes to the `candidate-assessment` agent (its own staleness — 100-pt rubric, hardcoded paths — is a separate fix)
