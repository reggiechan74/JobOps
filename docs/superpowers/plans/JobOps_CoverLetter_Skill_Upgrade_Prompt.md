---
title: JobOps Cover-Letter Skill Upgrade Prompt
date: 2026-05-28
keywords: [jobops, cover-letter, skill-upgrade, template, prompt, phone-slot, ats]
lastUpdated: 2026-05-28 21:21 EST
category: professional
documentType: prompt
status: ready-to-send
relatedDocs: [cover-letter/cover_letter.md, resume/resume_final.md]
---

# JobOps Cover-Letter Skill Upgrade Prompt

**Purpose:** Paste the fenced block below into the *other* Claude Code repo that owns
the JobOps skills/agents and templates. It folds the lessons from the Equinix xScale
cover-letter edit (May 28, 2026) back into the generator so future letters start at
this quality.

**Key correction (read first):** the other cover-letter template has **no telephone
slot at all** — phone is not a field in the template. The prompt instructs the repo to
*add* a `{phone}` field to both the markdown and LaTeX headers, not merely fix a
separator. The fused `(647) 892-8074reggie.chan@gmail.com` artifact happened because
the phone was hand-jammed into a template that never had a place for it.

---

````text
# Task: Upgrade the JobOps cover-letter agent/skill + templates

I refined a cover letter by hand in another repo and want to fold the lessons back
into the JobOps cover-letter generator so future letters start at this quality.

## Step 0 — Locate the relevant files
Find and read, then update as needed:
- The cover-letter agent (likely `step4-cover-letter`) and/or the `/coverletter` command/skill
- Any cover-letter TEMPLATE files (markdown header/body templates)
- Any LaTeX template used by `latex-pdf` for cover letters (header/contact block)
- The repo's CLAUDE.md / MEMORY / skill instructions where authoring rules live

Report what you found before editing.

## Required template fix (do this for sure) — ADD A PHONE FIELD
IMPORTANT: the current cover-letter template has NO telephone slot at all. Phone is
not a field. Because there was nowhere to put it, the number got jammed in next to the
email and rendered as `(647) 892-8074reggie.chan@gmail.com`.

ADD a `{phone}` field to the contact/header block (markdown AND LaTeX), so the header
renders distinct, separately-delimited fields:

    {name}, {credentials}
    {location} | {phone} | {email} | LinkedIn: {linkedin}

Rules:
- Introduce a real `{phone}` slot in the template variables/schema (it does not exist today).
- Join fields with ` | ` and OMIT any empty field cleanly (no orphan separators).
- Never adjacency-concatenate two fields. Apply the same change to the LaTeX header block.
- If the candidate profile/source data lacks a phone field, add one there too and wire it through.

## Authoring principles to bake into the agent/skill instructions
For each, encode the RULE and keep a short before/after so the model internalizes it.

1. OPENING = FIT-LED, not diagnostic-led, not "please find attached."
   - First line names the role; then state core-competency match on the candidate's
     own track record; close the paragraph with an honest pivot if there's a gap.
   - Do NOT open by diagnosing the employer's business (reads presumptuous to an
     insider hiring manager). Demote any company/market insight to paragraph 2 as
     CONTEXT that frames the role, not as the lead.
   - Banned opener: "Please find attached… I am interested in this position because…"
     (low-signal at senior level and now an AI tell).

2. NEVER recite the target's own facts/figures.
   Don't quote their earnings stats, AUM, headcount, or their own announced deals
   back to them ("promote me, don't remind them of their press release"). Convert
   any such figure into qualitative synthesis that shows operational insight.
   - Before: "roughly six of every ten of the largest deals now tied to AI"
   - After:  "demand has shifted from projected to contracted"
   - When you cut a cited fact, RECONCILE the `primary_sources` frontmatter note
     to record that the figure was removed and why.

3. NO em dashes anywhere. This candidate does not use them; em dashes are an AI
   artifact. Use colons, semicolons, or periods.

4. Avoid repeated antithetical cadence. The "X is new, Y is not" / "not A, but B"
   construction may appear AT MOST ONCE in a letter; vary all other closers. Repeated
   parallel symmetry is a primary AI-detection tell.

5. De-aging / no gratuitous tenure stamps. Do not state explicit total years of
   experience ("25 years") in the cover letter or résumé summary — it's an age proxy.
   Convey seniority through PROOF (named roles, $ figures, scope), which still clears
   "10+ years" requirements. On résumés, drop graduation years on degrees earned
   >15 years ago. Keep recent credential dates (they read as currency).

6. Keep idiosyncratic, un-fakeable specifics as the human-authenticity anchor.
   Named deals, specific counterparties, the candidate's own publications, and
   concrete numbers defeat AI-detection far better than generic polish. Prefer
   "a megawatt underwrites differently from a square foot" over abstract claims.

7. Requirements→Experience TABLE is a supported, encouraged device (ATS-friendly,
   skimmable). Map JD requirements to specific proof. Keep it.

8. Two-reader calibration. The letter must survive BOTH the recruiter/ATS gate
   (fast skim + keywords → served by the fit-led opener + table) AND the senior
   hiring manager (rewards demonstrated understanding → served by paragraphs 2+).
   Lead plain for the screener; deliver sophistication after.

## Gold-standard exemplar (reproduce this quality/structure)
```markdown
# Reggie Chan, CFA, FRICS

Toronto, ON | (647) 892-8074 | reggie.chan@gmail.com | LinkedIn: /in/reggiechan

May 28, 2026

xScale Asset Management
Equinix, Inc.

Dear xScale Asset Management Team:

I'm applying for the Senior Manager, xScale Asset Management role. I have spent my
career doing the core of this job on institutional and REIT portfolios: investor and
JV reporting, performance against underwriting, and the commercial analysis behind
contract terms. The asset class is new to me; the discipline is not.

By design, xScale has front-loaded the hard parts: institutional partners hold the
majority of the equity while Equinix runs the platform as developer-manager on a
minority stake, a perpetual-vehicle model rather than the on-balance-sheet IBX
business, and demand has shifted from projected to contracted. What does not scale
automatically with capital and demand is governance. The scarce function is the
individual contributor who keeps performance reporting honest against underwriting
across a hyperscale portfolio being built faster than most reporting disciplines can
keep up. That is this role: relationship manager for equity partners, author of the
QBRs and annual plans, analyst on the contract terms that deviate from criteria and
need JV sign-off, and the person who surfaces a variance before a partner finds it.
That intersection of investor governance, financial reporting, and commercial
analysis is the center of my experience.

| **Your Requirements** | **My Proven Experience** |
|----------------------|--------------------------|
| JV / external-partner governance and investor reporting | At AIG Global Real Estate I coordinated a US$1.8B Asia-Pacific portfolio across three external JV partners, preparing the quarterly investment-committee decks and monthly NAV reporting that kept those partners aligned. |
| Performance, financial, and commercial reporting cadences | At Artis REIT I owned monthly executive-committee reporting and quarterly board reporting across a 40-property, ~$865M portfolio over a nine-year VP tenure. |
| Governance under demanding stakeholders | Emboss Capital: appointed by independent directors to a distressed HK$800M Macau asset after the prior managers were arrested, I reported to UK fund directors and a five-bank lender syndicate and sold the asset at HK$780M, a 56% premium, with full lender recovery. |
| Commercial analysis of complex deal and contract terms tied to underwriting | I analyzed 700+ lease transactions against breakeven and accretion hurdles using a self-built approval model that gated every deal against the owner's financing stack. |
| Performance against underwriting and value creation | The Artis industrial book grew 141% in fair market value (Q4 2017 to Q2 2021) against 35% market growth, with hold/sell calls made on 10-year DCF and IRR attribution. |

On investor governance: the work that matters is not the deck, it is the reporting
cadence that keeps a sophisticated partner from being surprised. At AIG I ran that
cadence across three external JV partners on a US$1.8B book, pairing quarterly IC
presentations with monthly cash-flow forecasts and variance analysis. The harder
version was Emboss Capital. Independent directors handed me a distressed HK$800M
asset after a fraud, and I reported the turnaround to UK fund directors and a
five-bank syndicate every month until the asset sold at a 56% premium and the lenders
were made whole. That is the part of this job I am most confident in.

On performance against underwriting: surfacing a variance before a partner finds it
is the entire job. At Artis I built annual budgets and multi-year asset business plans
with variance tracking, and ran hold/sell decisions on 10-year DCF with IRR
attribution that separated NOI growth from cap-rate movement. The reporting told us
when to lease and when to sell. The industrial book grew 141% in fair market value
over roughly four years against a market that moved 35%.

On commercial contract analysis: a rate structure or a ramp profile only matters once
it lands in the model and the underwriting. At New Star I traded away Japanese lease
termination rights for rental collars and held 90%-plus tenant retention through the
2008 crisis, then carried the same logic into a self-built approval model that tested
every deal against the owner's financing stack. I published the underlying framework,
an objective levered and unlevered breakeven that prices a lease against true
cash-flow-neutral recovery rather than budget, in *Real Estate Finance* in 2015,
writing specifically for REIT asset managers. The contract terms, the model, and the
JV sign-off are one conversation, not three.

What I do not bring is data-center operating knowledge: hyperscale lease economics,
power and PUE, the way a megawatt underwrites differently from a square foot. That
expertise is real and earned, and I will not pretend it transfers overnight. What I do
bring is rarer: institutional fund and REIT asset-management depth, governance tested
under sophisticated LPs and lender syndicates and the CFA and FRICS valuation rigor
this role is gated on. That discipline is what keeps LPs whole, and it does not reset
when the asset class changes.

The next phase of the U.S. xScale JV is about contributing assets into the vehicle and
reporting their performance against what the partners underwrote. I want to be the
person accountable for that reporting being right.

Sincerely,

Reggie Chan, CFA, FRICS
```

## Acceptance criteria
- A NEW `{phone}` field exists in the template variables/schema and renders in the
  header (markdown + LaTeX) with ` | ` separators and clean omission of empty fields.
- Agent/skill instructions encode principles 1–8 with at least one before/after each.
- A regression note/checklist exists so the generator self-checks: em dashes, recited
  target facts, repeated antithetical cadence, explicit tenure numbers, missing phone
  field, and a fit-led (not diagnostic-led, not "please find attached") opening.
- Show me a diff of every file you changed and a one-paragraph summary.
````
