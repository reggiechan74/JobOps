---
name: coverletter
description: Generate a strategic cover letter with requirements-matching table from Step 3 resume
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

## Application Path Resolution

This skill writes to a per-application folder. Before writing any output:

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the job-posting filename, or honor `--app=<slug>` if supplied.
2. Compose the app folder: `{config.directories.applications_root}/{app_slug}/`.
3. Resolve this skill's sub-folder by category:
   - resume-development (buildresume, provenance-check) → `resume/`
   - cover-letter (coverletter) → `cover-letter/`
   - rubric / assessment (createrubric, assessjob, assesscandidate, auditjobposting) → `assessment/`
   - briefing / interview prep (briefing, interviewprep) → `interview/`
4. If the app folder does not exist, `mkdir -p` it, then copy
   `{config.directories.job_postings}/{filename}` → `{app_slug}/job_posting.md`
   so the pinned JD cannot silently change under completed work.
5. Exact-slug collisions (same Company+Role+Date) are not auto-suffixed. If the folder
   already contains the same output type, require the user to pass `--app=<distinct-slug>`.

## Arguments

- `$1`: Step 3 final resume file path (required)
- `$2`: Job description file path (required)
- `$3`: Hiring manager name (optional, defaults to "Hiring Manager")

Produces a compelling cover letter based on the validated Step 3 resume, featuring a strategic requirements-matching table that directly demonstrates fit for the role. Process:

1. Load and analyze the validated Step 3 resume.
2. Extract critical requirements from the job description.
3. Create a requirements-matching table with evidence.
4. Generate a compelling narrative with cultural alignment.
5. Produce an interview-ready cover letter.

## Step 1: Loading Final Resume

First, let me read your Step 3 final resume:

@$1

## Step 2: Analyzing Job Requirements

Now loading the job description to extract critical requirements:

@$2

## Step 3: Generating Cover Letter

**Running Step 4 Agent...**

I'm launching the `step4-cover-letter` agent to create your cover letter. The agent enforces a refined methodology: every cover letter has the same seven structural elements in the same order, written in a declarative first-person voice that demonstrates insight rather than stating intent.

### The Seven-Element Structure

The agent writes the letter in this order. Each element is mandatory unless explicitly marked optional in the agent.

1. **Thesis cold-open.** A diagnostic claim about the organization's actual situation and what the role really is. No "I am writing to apply for…". The opening synthesizes 2–4 verified primary sources as their *consequences*, not their headlines. Demonstrates insight, never states intent.
2. **Role reframe.** Names the role's core function and the next 12–24 months in the candidate's own words, then stakes the claim that this intersection is where the candidate's career sits. Fit is inferred from positioning, not asserted.
3. **Requirements Alignment table (mandatory).** Two columns, capped at 5 rows. Each evidence cell carries a named system/project/counterparty and at least one quantity. Row construction varies — not every row is `verb + quantity + result`.
4. **"On X:" evidence paragraphs (2–3).** Each opens with a colon-led role-demand phrase ("On POC-to-production:", "On stakeholder negotiation:") and carries exactly one named, quantified accomplishment. One paragraph = one artifact. Each paragraph ties to the cold-open's diagnostic.
5. **Honest-limitation paragraph.** Uses the move "What I do not bring is X. What I do bring is rarer: Y." Real gap named honestly; pivot to differentiated strength; tie back to the cold-open. **Never trivialize the gap as quickly closeable** — phrasing like "X is learnable in a few weeks" or "I can pick this up fast" reads as a lack of humility toward genuine expertise and is banned. When the gap is an AI-built capability, the explicit AI-authorship split is required (candidate owns spec/schema/architecture; implementation is pair-programmed with Claude Code and Codex CLI by name).
6. **Forward-looking close (3–4 lines).** Ties to the role's near-term mandate and ends with a confident specific ask. No "thank you for your consideration," no "I look forward to hearing from you," no contact-info restatement.
7. **Signature.** `Sincerely,` / signature image / `{Candidate Name}, {post-nominals}`.

### Voice and Style (enforced by the agent)

- **Declarative, first-person, confident.** Sentence fragments allowed for emphasis ("That is this role." / "The methodology is portable.").
- **Em-dashes for elaboration: zero per letter.** Use commas, periods, parentheses, semicolons, or colons. Em-dashes are the most reliable AI-prose fingerprint.
- **Concrete metrics and named systems over adjectives.** Banned clichés include "excited to apply," "passionate," "team player," "fast-paced," "synergy," "leverage" used as filler. Banned dead verbs include spearhead, orchestrate, drive, deliver, champion, unlock, empower, facilitate.
- **Explicit AI-authorship split** whenever AI-built work is cited. Never vague "AI-assisted" phrasing.
- **Synthesize, do not recite.** Convert institutional facts into operational insight. Recitation reads as flattery and wastes lines.
- **Address the hiring manager by first name** when known.

### Provenance and the Requirements-Matching Table

- The table maps the top 5 requirements to specific achievements from the Step 3 resume with quantified evidence.
- Every institutional fact used in the cold-open or body must trace to a verified primary source. The `primary_sources` YAML block (see Output below) is the audit ledger; `status: verified` sources are the only ones that may appear in the letter prose.
- If a fact is cut during editing, the `primary_sources` block is reconciled so it stays accurate.

### Quality Assurance

- All claims traced to Step 3 validated content. No new claims that haven't been validated.
- All institutional facts traced to verified primary sources (Step 3a acquisition pipeline in the agent).
- An independent sub-agent reviews the draft sentence-by-sentence using **What / So What / Now What** framing before the letter is written to disk (Step 6a in the agent). Every sentence must address a real hiring-manager concern and move the reader toward a decision; filler is flagged CUT.
- Company and role names verified for accuracy.
- Cultural tone matches resume style; optimized for ATS and human review.

### Gold-Standard Exemplar (illustrative; names anonymized)

Imitate this letter for structure and voice, not for content. All names, firms, and product references below are placeholders (John Smith / ABC Inc. / XYZ Corp etc.) — the real letter substitutes the actual hiring manager, target firm, prior employers, and named systems from the candidate's record. Notice how each of the seven elements appears in order, how the cold-open does its work without recitation, and how the honest-limitation paragraph uses the explicit AI-authorship split.

> Dear John:
>
> ABC Inc. has the demand, the governance, and the delivery capacity for enterprise AI. What it does not yet have is a customer-facing function inside the technology group to translate between the three. That is this role. Every business unit (the property platform, the investment teams, HR, Finance) now runs its own AI ambitions against an AI Governance Policy that draws a hard line between Custom and Public AI, and oversight of AI now sits at the Board-committee level. AI is already in production; the binding constraint is no longer capability but translation between business-unit demand, governance constraint, and delivery capacity.
>
> The next 12 months are about extending that translation work into the units that will use it first: the property platform's asset teams, then the investment platforms and the operating functions, under explicit governance. The Associate Director hired into this role is that translator: triaging internal-customer AI requests, framing build-vs-buy decisions against Custom-versus-Public obligations, leading POCs to production, and running the workshops that turn leaders into hands-on users of the tools. Most enterprise AI programs miss that last move. That intersection is where I have built my career.
>
> [Requirements Alignment table maps each requirement to one evidence cell with metrics: enterprise SaaS rollout 90% adoption month 1; 25 production systems / 3,000+ commits; production SQLite schema design with FTS5 + sqlite-vec; XYZ Corp to DEF Inc. to GHI Holdings VP plus CFA/FRICS; CEO/executive quarterly reporting over a 9-year VP tenure.]
>
> On POC-to-production: I have already built similar things that this role describes, an internal-facing AI system that sits next to a non-technical user and answers their questions on demand. My relationship-intelligence system lets me query a database in plain language and get structured answers back, the same pattern ABC Inc. would use to seat AI beside a property asset manager or an investment analyst. A second system took a piece of my own consulting methodology and turned it into a tool a colleague can actually run. Across both, the through-line is that I take ideas to working software people use, not slideware, which is exactly where most internal AI programs stall.
>
> On the workshop side: at JKL Corp. (2022 to 2024) I designed and delivered an AI-powered property-acquisitions onboarding program (600+ pages of training content and twenty fifteen-minute audio episodes, built in two months). The program cut new-hire ramp-up time by 50%, validated with new hires. I have since formalized the underlying methodology into a CRE AI training curriculum for commercial real estate professionals. The methodology is portable. The same architecture would convert ABC Inc. analysts and managers from prompt-readers into agent-builders inside their respective customer domains.
>
> What I do not bring is enterprise-scale data-warehouse experience; I own the specification, schema, and query architecture, and pair-program the build with Claude Code and Codex. What I do bring is rarer: institutional CRE fluency at platform depth, plus a track record of taking requirements all the way to running systems. That is the exact intersection where the AI Governance Policy now needs translating into customer-facing work.
>
> The first 12 months of the Customer Success and Innovation function will define how that translation actually happens. I would like to be inside the room when those decisions are made.
>
> Sincerely,
> [signature image]
> John Smith, CFA, FRICS

The exemplar is a model for *form*, not for *content*. The agent will not copy its institutional facts into a letter for a different firm, and will not reuse this exemplar's `primary_sources:` block.

## Output

Your strategic cover letter will be saved as:
`{applications_root}/{app_slug}/cover-letter/cover_letter.md`

Before the letter content, write this YAML metadata block with actual values. The `primary_sources` ledger is mandatory whenever the thesis cold-open uses verified primary sources; it is the audit trail that lets the hiring panel confirm the candidate did the work.

```yaml
---
job_file: $2
role: <role title>
company: <company name>
candidate: <full candidate name>
generated_by: /coverletter
generated_on: <ISO8601 timestamp>
output_type: cover_letter
status: final
version: 1.0
primary_sources:
  - title: <document title>
    date: <YYYY-MM-DD>
    url: <public URL fetched during verification>
    status: verified | unverified
    verified_on: <ISO8601 timestamp when WebFetch confirmed the document; null if unverified>
    used_in: <opening | body-paragraph-2 | table-row-3 | not-used>
    notes: <optional context, e.g. "paywalled", "retrieved from web archive", "OSINT corporate.md said 2025; primary source confirms December 2024">
---
```

Rules for the `primary_sources` block (enforced by the agent):
- Only `status: verified` sources may appear in the letter prose. Unverified candidates may be listed for reference but never cited.
- If a fact is cut during editing, the `primary_sources` block is updated so it stays accurate.
- If the same `primary_sources` block appears verbatim across letters to different firms, the letter is not firm-specific and must be rewritten from a fresh acquisition.
- When fewer than 2 verified primary sources are available, set `primary_sources: []` and have the agent use the fallback cold-open. Valid skip conditions: `private_firm_no_public_record`, `generic_role_any_firm`, `ats_screened_pipeline`, `insufficient_verified_primary_sources`.

Update timestamps and increment `version` on subsequent iterations.

The cover letter will:
- Demonstrate insight into the firm's actual situation, not state intent
- Provide visual proof of requirements match via the alignment table
- Name real gaps honestly and pivot to the differentiated strength
- Maintain a verified provenance chain from primary sources through master documents
- End with a confident, specific ask tied to the role's near-term mandate

## Next Steps

After generation, you'll have:
1. A compelling, evidence-based cover letter
2. A requirements-matching table for quick scanning
3. Consistent messaging across all application materials
4. Strong positioning for interview success

The cover letter and resume work together as a cohesive application package, with every claim defensible and every achievement traceable to your comprehensive work history.
