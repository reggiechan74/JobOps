---
name: steelman
description: Steelman a prepared job application package by OSINTing the hiring manager, building an evidence-based psychographic and behavioral profile, simulating the hiring manager's review, and recommending truthful improvements that increase interview odds
disable-model-invocation: true
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.
Use `config.preferences.cultural_profile` if this skill generates resume-style content.
Use `config.preferences.default_jurisdiction` if this skill has jurisdiction-sensitive logic.

## Purpose

This skill stress-tests a prepared application package before submission.

It has two jobs:

1. Build a public-source hiring-manager profile: professional history, decision style, communication preferences, likely risk sensitivities, domain priorities, and interview-selection heuristics.
2. Put yourself in that hiring manager's seat and decide whether the current package earns an interview. Then make the package stronger without lying, exaggerating, inventing facts, or creating claims that cannot be defended from source evidence.

The goal is adversarial improvement: knock the application down, identify what would cause a rational hiring manager to pass, then recommend the highest-leverage fixes.

## Arguments

- `$1`: Application slug, application folder, or job posting path (required)
- `$2`: Hiring manager name, profile URL, or person identifier (required unless already present in the job/application materials)
- `$3`: Optional package path or mode:
  - `--package=<path>` to review a specific package file or folder
  - `--refresh-osint` to rerun hiring-manager OSINT even if a prior report exists

If the application package or hiring manager identity cannot be resolved, ask the user for the missing item and stop. Do not guess the person.
If web access is unavailable, stop and tell the user that `/steelman` requires live web access because hiring-manager OSINT is the core input. Do not run a local-only substitute.

## Application Path Resolution

This skill writes to a per-application folder. Before writing any output:

1. Resolve the application:
   - If `$1` is an existing folder, use it as the application folder.
   - If `$1` is an application slug, use `{config.directories.applications_root}/$1/`.
   - If `$1` is a job-posting filename or path, parse `{Company}_{Role}_{YYYYMMDD}` from the job-posting filename, or honor `--app=<slug>` if supplied.
2. Compose the output folder: `{application_folder}/assessment/`.
3. If the app folder does not exist and a job posting was supplied, create it and copy the job posting to `{application_folder}/job_posting.md`.
4. Write final outputs:
   - `{application_folder}/assessment/steelman.md`
   - `{application_folder}/assessment/hiring_manager_profile.md`
   - `{application_folder}/assessment/package_improvement_plan.md`

## Source Discipline

Use only lawful, public, work-relevant OSINT and local application materials.

Allowed public sources:
- Company website, leadership bios, press releases, public job posting
- LinkedIn public profile snippets or user-provided profile exports
- Conference talks, podcasts, articles, publications, patents, GitHub, public slide decks
- Public social posts only when professionally relevant
- Previously generated JobOps intelligence reports

Label every profile claim as:
- **Observed**: directly supported by a cited source
- **Inferred**: reasoned from observed professional evidence
- **Speculative**: weak signal; use only as a hypothesis, never as a basis for claims in the application

Any recommendation that depends on a speculative claim must be marked "low confidence" or removed.

## Workflow

### Phase 1: Load the Application Package

Read the application materials that exist for the resolved app:

- `{application_folder}/job_posting.md`
- `{application_folder}/resume/*.md`
- `{application_folder}/cover-letter/*.md`
- `{application_folder}/assessment/rubric.md`
- `{application_folder}/assessment/assessment.md`
- Any file or folder passed via `--package=<path>`

Identify package components as:
- Final resume
- Cover letter
- Requirements table
- Rubric or candidate assessment
- Supporting notes or interview brief

If the final resume or cover letter is missing, continue with what exists but state the gap prominently.

### Phase 2: Hiring-Manager OSINT

Build a hiring-manager profile. Live web research is required. If web access is unavailable, stop with:

> STEELMAN REQUIRES WEB ACCESS
> Hiring-manager OSINT is the core input for this skill. Re-run when web access is available.

If an existing person report is available under `{config.directories.company_intelligence}/{Company}/people/`, read it first as context. Still verify or refresh key hiring-manager facts through web research, especially identity, current role, employer, and recent public signals. If the report is stale, thin, or `--refresh-osint` is passed, refresh it fully.

When sub-agents are available, dispatch `osint-person` for the target person and pass:

- Hiring manager name and employer
- Role being applied for
- Target output path: `{application_folder}/assessment/hiring_manager_profile.md`
- Instruction to restrict analysis to public, work-relevant signals and to avoid protected traits

If sub-agents are not available, perform the research directly.

The hiring-manager profile must include:

1. **Identity and role confidence**: person, current title, employer, evidence, confidence.
2. **Professional trajectory**: career pattern, technical/domain background, management scope.
3. **Likely mandate**: what business problem this person is probably trying to solve with the hire.
4. **Decision style**: detail-oriented vs. narrative, risk-averse vs. opportunity-seeking, operator vs. strategist, consensus-driven vs. directive. Cite evidence or mark as inferred.
5. **Communication preferences**: signs they value brevity, metrics, technical depth, customer outcomes, governance, culture, speed, craft, or other signals.
6. **Credibility triggers**: what would make them trust the candidate quickly.
7. **Rejection triggers**: what would make them discard or down-rank the package.
8. **Interview hooks**: 3-5 conversation angles the package should make easy for them to ask about.
9. **Information gaps**: what is unknown and how that affects confidence.

### Phase 3: Hiring-Manager Seat Simulation

Review the application as the hiring manager. Be direct and skeptical.

Answer:

- Would this package get an interview from this hiring manager?
- What is the likely first impression after 30 seconds?
- What is the likely impression after a full read?
- What concerns would survive even if the reader likes the candidate?
- Which claims feel most credible, differentiated, or interview-worthy?
- Which claims feel generic, overbuilt, under-proven, misaligned, or risky?
- Where does the package fail to meet the hiring manager's likely mandate?

Use this decision scale:

- **Strong interview**: clear short-list package; only minor refinements needed.
- **Interview**: likely to advance, with some addressable weaknesses.
- **Borderline**: plausible but not safe; package needs targeted changes.
- **Unlikely**: reader has too many unresolved doubts.
- **No interview**: mismatch, credibility gap, or weak signal.

Do not soften the verdict to be encouraging. Do not be harsher than the evidence supports.

### Phase 4: Truthful Improvement Plan

Recommend fixes that improve interview odds without fabrication.

For each recommendation, include:

- **Target file/section**
- **Problem**
- **Hiring-manager reason**
- **Change to make**
- **Evidence needed**
- **Risk if unchanged**
- **Confidence**

Use this hierarchy:

1. Fix factual or provenance risk first.
2. Strengthen alignment to the hiring manager's likely mandate.
3. Replace generic positioning with specific evidence.
4. Cut content that creates doubt, distraction, or false precision.
5. Add honest limitation handling where the package invites a concern.
6. Improve ordering, scannability, and decision flow.

Never recommend:
- Inventing experience, credentials, relationships, metrics, or internal knowledge.
- Implying direct familiarity with the hiring manager beyond public sources.
- Flattery, parasocial familiarity, or language that would feel invasive.
- Claims based on protected traits or sensitive personal information.

### Phase 5: Write Outputs

Write all three markdown outputs with YAML front matter.

`hiring_manager_profile.md`:

```yaml
---
application: <app_slug>
company: <company>
role: <role>
hiring_manager: <name>
generated_by: /steelman
generated_on: <ISO8601 timestamp>
output_type: hiring_manager_profile
status: final
version: 1.0
confidence: high|medium|low
---
```

`steelman.md`:

```yaml
---
application: <app_slug>
company: <company>
role: <role>
hiring_manager: <name>
generated_by: /steelman
generated_on: <ISO8601 timestamp>
output_type: steelman_review
status: final
version: 1.0
verdict: strong_interview|interview|borderline|unlikely|no_interview
confidence: high|medium|low
---
```

`package_improvement_plan.md`:

```yaml
---
application: <app_slug>
company: <company>
role: <role>
hiring_manager: <name>
generated_by: /steelman
generated_on: <ISO8601 timestamp>
output_type: package_improvement_plan
status: final
version: 1.0
---
```

## Steelman Report Format

```markdown
# Steelman Review: <Company> - <Role>

## Verdict

**Decision:** <Strong interview | Interview | Borderline | Unlikely | No interview>
**Confidence:** <High | Medium | Low>

<Two to four blunt sentences explaining the decision.>

## Hiring-Manager Lens

- **Likely mandate:** ...
- **Likely selection heuristic:** ...
- **Credibility triggers:** ...
- **Rejection triggers:** ...

## 30-Second Read

What the hiring manager likely notices first.

## Full-Read Assessment

What gets stronger or weaker after scrutiny.

## Interview Case

The strongest argument for advancing the candidate.

## Pass Case

The strongest argument for rejecting or down-ranking the candidate.

## Failure Points

Ranked issues that could cost the interview.

## Recommendations

Ranked changes, each tied to the hiring-manager profile and truthful source evidence.

## Questions for the Candidate

Only ask questions that would materially improve the package or resolve a live risk.
```

## Completion Standard

The work is complete only when:

- The hiring-manager profile distinguishes observed facts from inferences.
- The steelman verdict is explicit.
- The improvement plan names concrete edits, not vague advice.
- Every recommendation preserves truthfulness and provenance.
- Open questions are limited to issues that cannot be resolved from available sources.
