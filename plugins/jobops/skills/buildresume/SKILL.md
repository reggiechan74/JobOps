---
name: buildresume
description: Build a complete resume through all 3 steps (draft, provenance check, final)
disable-model-invocation: true
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

Templates referenced by this skill: evidence_verification_framework

## Application Path Resolution

This skill writes to a per-application folder. Before writing any output:

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the job-posting filename, or honor `--app=<slug>` if supplied. The slug MUST be canonical: a **leading** PascalCase `{Company}` token (matching the `Company_Intelligence/{Company}/` folder so OSINT links), a PascalCase `{Role}` (underscores between words allowed), and a **trailing** compact 8-digit date (`20260519` — no hyphens, no time). Reject leading date/time prefixes such as `2026-04-15_214414_...`; if the source filename carries one, recompose it into canonical form (`{Company}_{Role}_{YYYYMMDD}`) before composing the folder path.
2. Compose the app folder: `{config.directories.applications_root}/{app_slug}/`.
3. Resolve this skill's sub-folder by category:
   - resume-development (buildresume, provenance-check) → `resume/`
   - cover-letter (coverletter) → `cover-letter/`
   - rubric / assessment (createrubric, assessjob, assesscandidate, auditjobposting) → `assessment/`
   - briefing / interview prep (briefing, interviewprep) → `interview/`
4. If the app folder does not exist, `mkdir -p` it, then copy
   `{config.directories.job_postings}/{filename}` → `{app_slug}/job_posting.md`
   so the pinned JD cannot silently change under completed work. Ensure the pinned copy
   begins with YAML front matter carrying `output_type: job_posting`: if the source JD
   already has a front-matter block, add the key to it; otherwise wrap a new block
   (`---` / `output_type: job_posting` / `source_jd: {filename}` / `---`) above the JD body.
5. Exact-slug collisions (same Company+Role+Date) are not auto-suffixed. If the folder
   already contains the same output type, require the user to pass `--app=<distinct-slug>`.
6. **Output filenames** (fixed; only the sub-folder above is resolved dynamically). Pass the
   fully resolved absolute path to each step's agent in its Task instruction:
   - Step 1 draft → `resume/step1_draft.md`
   - Step 1 manifest (revise mode only) → `resume/step1_manifest.md`
   - Step 2 provenance analysis → `resume/step2_provenance.md`
   - Step 3 final resume → `resume/step3_final.md`
   Any PDF/TeX/DOCX derivative shares the **exact basename** of its source `.md` and lives in
   the **same** sub-folder (e.g. `resume/step3_final.pdf`) — never a separate `latex/` folder.

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

## Build Mode Selection

Determine the build mode BEFORE dispatching any step agent.

1. **Flag shortcuts.** `--from-scratch` → mode = `scratch`, skip to dispatch.
   `--base=<path>` → verify the file exists and is `.md`; mode = `revise` with that
   base, skip to dispatch. (Stamping and fit assessment are skipped; if the forced base lacks `role_family`, the promotion offer will ask for one.)
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
     [revise from <file> / build from scratch / pick another base]. "Pick another base" re-presents the remaining candidates with their verdicts; the user's choice proceeds to dispatch as the selected base.
   - Only PARTIAL candidates → present both options, naming the specific gaps
     (e.g., "base covers the summary and primary role, but the JD's P&L emphasis is
     not covered"); offer [revise from <file> / build from scratch].
   - All POOR → mode = `scratch`; state why no base fits.
6. Record the chosen mode and (in revise mode) the absolute base path. Carry both
   through every step dispatch below.

## Output metadata

Every markdown file generated during this skill must begin with a YAML front-matter block so downstream tooling can parse metadata without heuristics. Populate the fields with real values before writing any body content. The output path for each step is determined by step 6 of the Application Path Resolution protocol above (step 1 → `resume/step1_draft.md`, step 2 → `resume/step2_provenance.md`, step 3 → `resume/step3_final.md`); `job_file` in the front matter is always the absolute path to the input JD.

- **Step 1 draft** — set `generated_by: /buildresume step1-resume-draft`, `output_type: resume_step1`, `status: draft`, `version: 1.0`.
- **Step 2 provenance analysis** — set `generated_by: /buildresume step2-provenance-check`, `output_type: resume_provenance`, `status: analysis`.
- **Step 3 final resume** — set `generated_by: /buildresume step3-final-resume`, `output_type: resume_final`, `status: final`. Increment `version` if regenerating after revisions.

Common fields for all three:

```yaml
---
job_file: <absolute path to JD>
role: <target role title>
company: <hiring company>
candidate: <full candidate name>
generated_on: <ISO8601 timestamp>
---
```

Always write the front matter before any markdown headings or narrative body.

All three files also carry `build_mode: revise` or `build_mode: scratch`. In revise
mode, step 1 and step 3 outputs additionally carry `base_resume: <absolute path to the
selected base>`. The step agents' own front-matter templates do not list these fields — I pass the `build_mode` value (and, in revise mode, the `base_resume` path) in every step agent's Task instruction with the instruction to include them in the output's front matter.

Revise mode writes one additional file — the change manifest:

- **Step 1 manifest** (revise mode only) — `resume/step1_manifest.md`, with
  `generated_by: /buildresume step1-resume-revise`, `output_type: resume_manifest`,
  `status: manifest`, and `base_resume` set. Written by the step1-resume-revise agent
  before it touches the draft.

## Step 1: Creating Initial Resume Draft
✓ Initiating strike package assembly

I'll start by reading the job description file and creating the initial draft using the $2 cultural profile.

First, let me read the job description:

@$1

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
of Application Path Resolution) to the agent in its Task instruction. I also pass `build_mode: scratch` for the agent to include in the draft's front matter.

## Step 2: Provenance Analysis
✓ Executing credibility verification sweep

**✓ Deploying Step 2 Agent - Provenance Verification**

Now I'll launch the step2-provenance-check agent to analyze the Step 1 draft for:
- Credibility issues and evidence gaps
- Risk factors for each major claim
- Defensibility of quantified achievements
- Consistency across all experience sections
- **Missing credentials: Verify ALL education/designations from master resume are evaluated**
- **Credential-to-job mapping: Flag any excluded credentials that map to job requirements**
- Recommendations for strengthening credibility

I pass the Step 1 draft path (`{app_slug}/resume/step1_draft.md`) to read and the resolved output path (`{app_slug}/resume/step2_provenance.md`) to write, both in the agent's Task instruction.

In revise mode I additionally pass `build_mode: revise`, the base resume path, and the
manifest path (`{app_slug}/resume/step1_manifest.md`) so the agent tags every finding
with its origin (`BASE` = carried from the base, including the user's manual edits;
`NEW` = introduced or modified by the manifest) per its Revise-Mode Delta Tagging
section. In both modes I pass the `build_mode` value for the agent to include in the provenance file's front matter.

## Step 3: Final Hardened Resume
✓ Producing deployment-ready final resume

**✓ Deploying Step 3 Agent - Hardened Resume Production**

Finally, I'll launch the step3-final-resume agent to:
- Incorporate all recommendations from Step 2 analysis
- Address credibility concerns while maintaining competitive positioning
- Produce a fully defensible final resume
- Ensure maximum interview potential

I pass the Step 1 draft (`{app_slug}/resume/step1_draft.md`) and Step 2 provenance analysis (`{app_slug}/resume/step2_provenance.md`) paths to read and the resolved output path (`{app_slug}/resume/step3_final.md`) to write, all in the agent's Task instruction.

In revise mode I additionally pass `build_mode: revise` and the base resume path so the
agent follows its Revise-Mode Edit Protocol: copy the draft to the final, apply
targeted edits only for Step 2 findings, and prove with a diff gate that nothing else
changed. In both modes I pass the `build_mode` value (and, in revise mode, the `base_resume` path) for the agent to include in the final's front matter.

## Mission Summary

All three steps of the resume assembly process will be completed:
- Step 1: Initial targeted draft created
- Step 2: Comprehensive provenance verification performed
- Step 3: Final hardened resume produced

Your deployment-ready resume will be ready for mission execution with full credibility and competitive positioning.

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
