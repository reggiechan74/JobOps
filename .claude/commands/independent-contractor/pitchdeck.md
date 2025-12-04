---
description: Generate professional B2B service pitch deck tailored to prospects or industries
argument-hint: [--prospect=company-name] [--industry=X] [--service=service-name] [--format=md|pptx]
---

Generate 10-12 slide B2B pitch deck with provenance-hardened claims. Supports prospect-specific, industry-generic, service-focused, or general capabilities presentations.

**Modes:** `--prospect=name` | `--industry=X` | `--service=name` | default (general)
**Format:** `--format=md|pptx` (markdown default, pptx requires pandoc)
**Output:** `Client_Prospects/Pitch_[Target]_[Date].md`

---

## Phase 1: Validate Prerequisites and Load Data

### 1.1 Prerequisites Check

**Required files (stop if missing):**

1. **Service Definition**: `Client_Prospects/Service_Definition_*.md`
   - If missing: Show error, instruct to run `/defineservices`, stop execution

2. **Candidate Profile**: `ResumeSourceFolder/.profile/candidate_profile.json` (≤7 days old)
   - If missing/stale: Show error, instruct to run `/assessjob <job-file.md>`, stop execution

### 1.2 Load Service Definition

Use most recent `Service_Definition_*.md` if multiple exist. Extract:
- YAML: consultant, generated_on, version, source_profile
- Consultant: name, tagline, credentials, years_experience
- Services: name, category, description, deliverables, ideal_client (industries, pain_points, decision_makers), pricing (model, range, value_justification), success_metrics, case_studies
- Rate Card: hourly/daily/retainer rates, payment_terms
- Differentiation: unique_value, competitive_advantages, proof_points, certifications, publications, speaking
- Target Market: primary_markets, secondary_markets, geographic_focus

### 1.3 Load Candidate Profile

Read `candidate_profile.json` and extract: work_history, technical_skills, domain_expertise, certifications (Active only), projects, thought_leadership, leadership_experience, achievements. Create evidence index mapping all quantified metrics/credentials to source file + lines for provenance validation.

### 1.4 Parse Arguments

**Targeting modes (mutually exclusive):**
- `--prospect=name`: Load `Prospects_*.md` file (use most recent). If not found, show error with available prospects, stop.
- `--industry=X`: Validate against service definition industries. Warn if not listed, but proceed.
- `--service=name`: Fuzzy match against service definition. If no match, show error with available services, stop.
- Default: General capabilities (all services, industry-agnostic)

**Format:**
- `--format=pptx`: Check pandoc installed. If not, warn and use markdown only. If installed, convert after generation.
- Default: markdown

---

## Phase 2: Conduct Target Research (5-10 minutes)

### 2.1 Prospect Mode (5-7 min)
Load prospect data from file. Conduct fresh research:
1. **News & Initiatives (2-3 min)**: Recent press releases, strategic priorities, funding, leadership changes, technology adoption
2. **Pain Points (2 min)**: Glassdoor reviews, job postings (skill gaps), RFPs, industry reports
3. **Decision-Makers (1-2 min)**: LinkedIn verification, role changes, recent activity

### 2.2 Industry Mode (5-8 min)
1. **Trends & Challenges (3-4 min)**: Top trends, pain points, technology adoption, regulatory issues, market indicators
2. **Consulting Market (2-3 min)**: Market size, demand signals, engagement models, competitive landscape
3. **Decision-Makers (1-2 min)**: Common titles, priorities, KPIs, budget authority, procurement

### 2.3 Service Mode (4-6 min)
1. **Market Landscape (2-3 min)**: Market demand, ROI data, success stories, pricing benchmarks
2. **Pain Points (2-3 min)**: Cost of inaction, prevalence, solutions, ROI expectations

### 2.4 Quality Check
Validate: ≥5 credible sources, ≤6 months old, quantified data, current decision-makers, supported pain points. If LOW quality, warn but proceed.

---

## Phase 3: Generate Pitch Deck (10-12 Slides)

### 3.1 Slide Structure
| # | Title | Content | Source |
|---|-------|---------|--------|
| 1 | Title | Identity, tagline, contact | Service definition |
| 2 | Problem | Pain points | Research + service definition |
| 3 | Cost of Inaction | Quantified impact | Research + industry data |
| 4 | Solution | Service overview | Service definition |
| 5 | How It Works | Methodology | Service definition + evidence |
| 6 | Results & Proof | Case studies, metrics | Candidate profile + service def |
| 7 | Why Me/Us | Credentials, differentiation | Both sources |
| 8 | Engagement Options | Pricing, packages | Rate card |
| 9 | Relevant Experience | Industry examples | Candidate profile (filtered) |
| 10 | Next Steps | CTA, timeline | Mode-specific |
| 11 | Q&A | Objections | Generated |
| 12 | Appendix | Credentials | Candidate profile |

Adjust: Service mode 8-10 slides, Prospect 12, Industry 10-11.

### 3.2-3.13 Slide Generation Instructions

**General Rules (all slides):**
- Include evidence citations for all quantified claims
- Customize "Presented to" field based on mode
- Use ONLY data from service definition and candidate profile
- Maximum 3-4 main points per slide (focus over volume)

**Slide 2 (Problem):** Mode-specific pain points with evidence. Prospect: company-specific from research. Industry: common challenges with statistics. Service: pain point service addresses with prevalence. Max 3 pain points, cite sources.

**Slide 3 (Cost of Inaction):** Impact table with quantified costs (4 areas), opportunity costs, compounding effects, urgency rationale. Prioritize prospect data, use industry benchmarks if unavailable. Cite sources.

**Slide 4 (Solution):** Service overview. Multi-service: 2-3 services with top deliverables. Single-service: full deliverables + success metrics. Prospect: challenge-solution-outcome table. Use exact deliverables from service definition.

**Slide 5 (How It Works):** 3-5 methodology phases with duration, activities, deliverables, client involvement. Extract methodology from candidate profile work_history. Timeline must align with service definition typical_duration.

**Slide 6 (Results & Proof):** CRITICAL PROVENANCE SLIDE. 3-5 top achievements with Result/Context/Relevance format. Include case study if available. Add credentials. ALL claims must have evidence citations from candidate profile. Rank by relevance (industry match 10pts, service match 10pts, recency 5pts, scale 5pts). Validate: quantified metric, evidence.file + evidence.lines exist, timeframe present, mechanism described.

**Slide 7 (Why Me/Us):** Unique value proposition (2-3 sentences). 3-4 competitive advantages with What/Why It Matters/Proof format. Comparison table (Typical vs. Us). Use competitive_advantages directly from service definition.

**Slide 8 (Engagement Options):** Service packages with pricing model, investment range (standard/entry/premium), duration, included deliverables. Engagement models table (project/retainer/hourly/value-based). Payment terms. Use ONLY pricing from service definition.

**Slide 9 (Relevant Experience):** Filter work_history by mode. Prospect: industry-matching roles (2-3). Industry: industry-specific roles (2-4). Service: cross-industry capability demonstration (3-5). Include technical depth (Expert-level skills). All must have evidence citations.

**Slide 10 (Next Steps):** Prospect: Timeline with discovery/proposal/launch phases, required stakeholders, 3 engagement options including entry point from prospect file. Generic: typical timeline, engagement options (call/proposal/download), contact info, testimonials if available.

**Slide 11 (Q&A):** 4-6 common objections based on mode. Prospect: company-type objections. Industry: industry-specific concerns. Service: technical feasibility, ROI validation. Include answers with evidence references.

**Slide 12 (Appendix):** Certifications table (Active only), education, thought leadership (publications/speaking/frameworks), professional associations, technical proficiencies (Expert + top 5-7 Proficient), awards, reference policy.

---

## Phase 4: Provenance Hardening (CRITICAL QUALITY GATE)

### 4.1-4.3 Claims Validation Process

**Extract claims:** Scan slides for quantified metrics, credentials, superlatives, client references, technologies, timeframes, ROI claims. Create inventory with slide #, type, source.

**Validate each claim:**
1. **Locate evidence:** Match METRIC/ACHIEVEMENT to work_history/projects/leadership_experience. CREDENTIAL to certifications (Active only). TECHNICAL SKILL to technical_skills (Expert/Proficient). THOUGHT LEADERSHIP to publications/speaking/frameworks.
2. **Verify quality:** evidence.file + evidence.lines exist, file verifiable, metric has timeframe + mechanism.
3. **Classify:** VALIDATED (HIGH: complete evidence, quantified, timeframe, mechanism). WEAK (MEDIUM: evidence incomplete, no timeframe, unclear mechanism). FAIL (none found, contradicts, superlative without benchmark, expired cert).

**Handle failures:** Try alternative evidence, soften claim, or remove. Priority: remove superlatives first, unbounded metrics second, keep partial evidence if critical (mark "Estimated").

### 4.4-4.5 Provenance Trail & Quality Gate

**Generate PROVENANCE TRAIL section** at end of markdown with: Generated timestamp, source files, research date, Claims Validation Summary (total/validated/weak/removed counts, validation rate %), Evidence Mapping table (slide, claim, evidence source, confidence, notes), Research Sources (URLs), Validation Notes (HIGH/MEDIUM/removed explanations), Pre-Use Checklist (6 items).

**Quality gate:** ≥90% = HIGH confidence (ready). 80-89% = MEDIUM (review recommended). <80% = LOW (major revision required, show issues/recommendations).

---

## Phase 5: Generate Output Files

**Filename:** `Pitch_[Target]_[Date].md` where Target = company name (prospect mode), industry name (industry mode), service name (service mode), or "GeneralCapabilities" (default). Sanitize: remove spaces/special chars, PascalCase, max 50 chars. Date: YYYYMMDD.

**File structure:** YAML frontmatter (pitch_type, target, generated, consultant, service_definition, candidate_profile, research_conducted, provenance_validated, validation_rate, confidence_level, version). 12 slides with `---` separators. PROVENANCE TRAIL section. Usage notes.

**Save:** `Client_Prospects/Pitch_[Target]_[Date].md` (UTF-8, Unix LF).

**PPTX conversion (if `--format=pptx`):** Run `pandoc ... -t pptx --slide-level=1`. Evidence citations → slide notes, Provenance Trail → hide in presentation.

---

## Phase 6: Summary Report

Show: Output files (markdown + pptx if applicable), Pitch details (target, type, slides, research time/sources), Content summary (services, proof points, advantages, pricing models), Provenance validation (rate %, confidence, status + warnings if MEDIUM/LOW), Research highlights (3 key findings), Next steps (6 items), Customization recommendations (4 items), Before presenting checklist (6 items).

---

## Error Handling

**Critical errors (stop execution):**
1. No service definition → instruct `/defineservices`
2. No/stale candidate profile → instruct `/assessjob`
3. Prospect not found → list available, suggest alternatives
4. Service not found → list available services

**Warnings (continue with warnings):**
5. Validation rate <80% → show issues, recommendations, save anyway
6. Research timeout/failure → show impact, manual research needed
7. Pandoc not found (pptx requested) → proceed markdown-only

---

## Usage Examples

**Prospect pitch:** `/pitchdeck --prospect="Hatch"` → Research company, tailor to pain points, include decision-makers
**Industry pitch:** `/pitchdeck --industry="Commercial Real Estate"` → Industry trends, common challenges, generic
**Service pitch:** `/pitchdeck --service="AI Agent Development"` → Deep-dive methodology, cross-industry examples
**General:** `/pitchdeck` → All services, top achievements, flexible
**PowerPoint:** Add `--format=pptx` to any command for PPTX conversion

---

## Important Notes

**Provenance:** 90%+ = HIGH (ready), 80-89% = MEDIUM (review), <80% = LOW (major review). All claims need evidence: metrics → achievements, credentials → certifications (Active), skills → technical_skills (Expert/Proficient), case studies → service def or projects, advantages → differentiation section.

**Research ethics:** Public info only (news, LinkedIn, Glassdoor, websites). No private data. 5-10 min time limit. Cite all sources.

**Before presenting:** Add logo, customize colors, add images, update Slide 10 dates/contacts, refresh if >30 days old, remove Provenance Trail for external use.

**Lifecycle:** Quarterly refresh (every 3 months: research, proof points, pricing, certifications). Each generation creates new dated file.

**Workflow:** 1) `/defineservices` 2) `/findclient` 3) `/pitchdeck --prospect=[name]` 4) `/assessjob` for follow-up. Supporting: `/ratecard`, `/idealjob`.

---

Now executing pitch deck generation...
