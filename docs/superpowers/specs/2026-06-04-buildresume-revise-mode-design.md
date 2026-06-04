# Revise-First `/buildresume` — Design

**Date:** 2026-06-04
**Status:** Approved for planning
**Affects:** `plugins/jobops` — `buildresume` skill, `step2-provenance-check` and
`step3-final-resume` agents, new `step1-resume-revise` agent, `setup` skill, config
contract, `docs/ARCHITECTURE.md`, `CLAUDE.md`

## Problem

`/buildresume` always builds from scratch: Step 1 reads the entire master inventory
plus the JD and regenerates a complete resume; Step 2 provenance-audits every claim;
Step 3 rewrites the whole document. Full regeneration re-rolls the dice on content
that was already settled, producing four observed artifact classes:

1. **Lost manual edits** — hand-polished wording from previous versions does not
   survive a rebuild.
2. **Content drift / invention** — bullets get rephrased in ways that weaken or
   subtly distort claims; metrics are framed differently each run.
3. **Formatting/structure drift** — section ordering, header style, and bullet
   conventions vary between runs, breaking PDF/Word conversion consistency.
4. **Selection instability** — different runs feature different achievements, so
   output quality depends on the luck of the draw rather than converging.

The user maintains a `Tailored_CV/` folder of base resumes: pipeline finals that were
subsequently hand-polished, organized by role family (e.g., executive, technical,
consulting/advisory — the actual categories are the user's own). Across applications
within a role family, roughly 70–90% of resume content is stable; only the executive
summary, competency emphasis, bullet selection/ordering, and keywords genuinely change
per JD. A previously hardened base has already survived provenance review; rebuilding
from scratch discards that work and the human polish layered on top of it.

## Goals

- Make **revise-from-base the default** `/buildresume` behavior; from-scratch becomes
  the fallback when no suitable base exists.
- Enforce minimal-delta edits **structurally** (copy + surgical edits + diff gate),
  not as a prompt-level hope. Unchanged content must be byte-identical to the base.
- Keep the role-family taxonomy **entirely user-defined** — no hardcoded category
  values anywhere in the plugin. Taxonomy is emergent from base-file front matter.
- Keep Step 2 provenance **full-document** (the user's manual edits in bases were
  never audited) but **delta-aware**, so Step 3 only touches flagged items.
- Close the loop: offer to **promote** each new final back into `Tailored_CV/`, so
  the base library converges and from-scratch builds seed new role families.
- Preserve all downstream contracts: output filenames, application path resolution,
  front-matter conventions, `--app` semantics.

## Non-goals

- No change to the from-scratch pipeline itself (`step1-resume-draft` agent is
  untouched).
- No automatic writes to `Tailored_CV/` — promotion is always offered, never silent.
- No support for non-markdown bases (DOCX/PDF dropped into the folder are ignored
  with a note).
- No numeric fit scores — fit verdicts are qualitative tiers with stated rationale.

## Decisions (from brainstorming)

| Question | Decision |
|---|---|
| What is the base material? | Mix: pipeline finals subsequently hand-polished, kept in `Tailored_CV/` |
| Invocation shape | Revise-first becomes the `/buildresume` default; from-scratch is the fallback |
| Base discovery | New config key + fit assessment + user confirmation |
| Multiple bases | Yes — one per role family; categories user-defined, never hardcoded |
| Provenance scope in revise mode | Full document, delta-tagged (`BASE` vs `NEW`) |
| Lifecycle | Offer to promote finals back into `Tailored_CV/` |
| Enforcement mechanism | Approach A: change manifest + surgical edits + mechanical diff gate |

## Design

### 1. Mode selection (new front-end to the skill)

```
/buildresume <JD> [profile] [--base=<path>] [--from-scratch] [--app=<slug>]
```

1. Resolve config. New key `config.directories.tailored_cv` (default `Tailored_CV/`),
   added by `/jobops:setup`. Key missing or folder empty → from-scratch (current
   behavior), with a one-line note that a base library would enable revise mode.
2. Scan `tailored_cv/` for `.md` candidates.
3. **Fit assessment** — runs inline in the skill (main conversation, no sub-agent,
   since it ends in a user-facing menu). Compare the JD against each candidate on
   four built-in axes:
   - declared category match (user's own labels — see §2),
   - positioning level (1–5 IC→C-suite, the scale already used by Step 1),
   - domain/industry overlap,
   - requirement-keyword coverage.
   Verdict per candidate: **STRONG / PARTIAL / POOR**, each with a one-line rationale
   expressed in the user's own category labels. No numeric scores.
4. Decision gate (user always confirms):
   - STRONG → recommend revising from that base.
   - PARTIAL only → present both options, naming the specific gaps
     ("base covers the summary and primary role, but the JD's P&L emphasis is not
     covered").
   - All POOR → from-scratch, stating why.
5. `--base=<path>` forces a specific base file (skips assessment);
   `--from-scratch` skips base discovery entirely.

### 2. User-defined categories (emergent taxonomy)

The category list lives in the library, not the plugin. Each base file declares its
own category in front matter:

```yaml
---
output_type: resume_base
role_family: <free-text label chosen by the user>
---
```

- The skill never matches against a built-in category list; it reads whatever labels
  exist in the library at run time.
- **Untagged files** (pre-existing hand-managed copies): on first encounter, the
  skill proposes a label inferred from content, the user confirms or replaces it, and
  the skill stamps the front matter. One-time, per file.
- Fit assessment compares the JD against each base's declared category **and its
  actual content**, so the verdict does not hinge on label quality.
- Positioning level is a secondary signal informing the rationale; it never overrides
  the user's categories.
- Promotion (§5) asks which category a new final belongs to, offering existing labels
  plus "new category" — which is how the taxonomy grows.

*Alternative considered:* a central category list in `.jobops/config.json`. Rejected:
it creates a second source of truth to keep in sync with the files.

### 3. Revise pipeline (manifest + surgical edits)

**Step R1 — gap analysis → manifest → surgical apply.** New agent
`step1-resume-revise`; the existing `step1-resume-draft` agent is untouched.

- Reads the base, the JD, and the master inventory (the inventory supplies swap-in
  bullets for JD demands the base does not cover).
- Writes `resume/step1_manifest.md`: an explicit list of changes (summary rewrite,
  bullet swap/reorder, keyword injection, competency re-emphasis), each justified by
  a named JD requirement. **Everything not in the manifest is frozen.**
- Copies the base to `resume/step1_draft.md` (front matter updated:
  `build_mode: revise`, `base_resume: <path>`, version per existing rules), then
  applies the manifest via targeted Edit operations — never a wholesale rewrite.
- **Diff gate 1** (performed by the R1 agent before reporting completion):
  `diff <base> step1_draft.md`; every changed hunk must map to a manifest item.
  Out-of-manifest changes are reverted. The diff summary is included in the agent's
  report and surfaced to the user.
- Cultural profile and positioning menus are **inherited from the base** — not
  re-asked unless `$2` (profile) is passed explicitly.

**Step 2 — provenance, delta-aware.** Full-document audit (manual edits in bases
were never audited), with each finding tagged `BASE` or `NEW`. Existing agent gains a
revise-mode section; output file and format otherwise unchanged.

**Step 3 — hardening in edit mode.** Copy `step1_draft.md` → `step3_final.md`, apply
fixes **only** for flagged findings via targeted edits.
**Diff gate 2** (performed by the Step 3 agent): changes between draft and final
must map to Step 2 findings. No whole-document rewrite.

### 4. File & metadata conventions (downstream-safe)

Same fixed filenames as today — `resume/step1_draft.md`, `resume/step2_provenance.md`,
`resume/step3_final.md` — so coverletter, PDF/Word conversion, and interview prep are
unaffected. One new sibling: `resume/step1_manifest.md` (front matter:
`output_type: resume_manifest`). Resume front matter gains `build_mode: revise|scratch`
and, in revise mode, `base_resume: <path>`. These fields are dictated by the
buildresume SKILL.md "Output metadata" section (passed to agents in Task
instructions), so the from-scratch `step1-resume-draft` agent file itself stays
untouched.

### 5. Promotion lifecycle (library convergence)

After Step 3 completes in **either** mode, offer once:

> Promote `step3_final.md` to Tailored_CV as the {role_family} base?
> [update existing base / save as new variant / skip]

Promotion stamps front matter (`output_type: resume_base`, `role_family`,
`promoted_from: <app_slug>`, version bump) and copies the file into
`tailored_cv/`. From-scratch builds use the same offer — that is how a new role
family is seeded. Nothing in `Tailored_CV/` is ever modified without this explicit
confirmation; the user's hand-edit workflow is preserved.

### 6. Repo changes

| File | Change |
|---|---|
| `plugins/jobops/skills/buildresume/SKILL.md` | Mode selection, fit assessment, manifest/diff-gate protocol, promotion offer |
| `plugins/jobops/agents/step1-resume-revise.md` | **New** — gap-analysis + manifest + surgical-edit specialist |
| `plugins/jobops/agents/step2-provenance-check.md` | Revise-mode delta-tagging section |
| `plugins/jobops/agents/step3-final-resume.md` | Edit-mode hardening section + diff gate 2 |
| `plugins/jobops/skills/setup/SKILL.md` | `tailored_cv` config key |
| `docs/ARCHITECTURE.md` | Config-contract table entry; revise-mode flow notes |
| `CLAUDE.md` | User-data-directories table entry |
| Version artifacts | Minor version bump via `version-bump` skill at implementation end |

### 7. Edge cases

- **Bases must be markdown.** Non-`.md` files in `tailored_cv/` are ignored with a
  note.
- **Stale base vs. updated inventory:** fine — the manifest phase reads the master
  inventory and can swap newer achievements in as deltas.
- **Slug collisions, `--app` semantics:** unchanged from the current skill.
- **Empty manifest** (base already fits the JD perfectly): legal — Step R1 reports
  "no changes required," the base is copied through, and Steps 2–3 still run
  (keyword coverage is re-verified during fit assessment).
- **Conflicting diff:** if an edit cannot be applied cleanly (base drifted since
  scan), the skill stops and reports rather than falling back to regeneration.

## Testing / validation

- Dry-run the revise flow on a real base + JD pair and verify with `diff` that
  unchanged regions are byte-identical end-to-end (base → draft → final).
- Verify from-scratch fallback triggers when: no config key, empty folder, all-POOR
  verdicts, and `--from-scratch`.
- Verify untagged-base stamping happens exactly once per file.
- Verify downstream consumers (coverletter, convert-to-pdf) operate unchanged on a
  revise-mode `step3_final.md`.
- `claude plugin validate plugins/jobops` and `npm test` (Codex contract) pass.
