---
description: Generate professional consulting proposal templates with pricing calculation
argument-hint: [--client=name] [--service=name] [--type=project|retainer|staff-aug|workshop] [--value=amount]
---

Generate McKinsey/BCG-style proposal with transparent pricing calculations.

**Arguments:**
- `--client=name`: Client name (default: "Generic")
- `--service=name`: Service from definition (default: primary service)
- `--type=TYPE`: `project|retainer|staff-aug|workshop` (default: project)
- `--value=amount`: Target pricing (default: calculated from rate card)

---

## Phase 1: Prerequisites & Input Loading

### 1.1 Load Service Definition

Find most recent: `find Client_Prospects/ -name "Service_Definition_*.md" -type f | sort -r | head -1`

**If missing:** Stop execution. Guide user to run `/defineservices` first.

**Parse required sections:**
- Consultant: name, credentials, years_experience, contact info
- Services: name, deliverables, pricing range, typical_duration, success_metrics, case_studies
- Rate Card: hourly/daily/retainer rates, payment_terms, expenses
- Engagement Models: preferred/acceptable types, minimum_engagement
- Differentiation: unique_value, competitive_advantages, proof_points

**Validate:** Require rate_card with hourly/daily rates, consultant credentials, ≥1 service, payment_terms.

### 1.2 Parse Arguments

- `--client`: Extract name, sanitize, default "Generic"
- `--service`: Fuzzy match against services, default to primary service, error if no match
- `--type`: Validate against [project|retainer|staff-aug|workshop], default "project"
- `--value`: Parse numeric, default null (calculate from rate card)

### 1.3 Load Optional Context

**Pitch Deck** (if client ≠ "Generic"):
- Find: `Pitch_*${client}*.md`
- Extract: prospect_research (pain_points, strategic_priorities, decision_makers, research_date)
- Set `prospect_context.has_prior_research = true/false`

**Candidate Profile**:
- Find: `ResumeSourceFolder/.profile/candidate_profile.json`
- Extract: certifications (Active), top achievements, thought_leadership
- Use for Team & Qualifications section, warn if missing

---

## Phase 2: Pricing Calculation

**CRITICAL: Transparent, defensible pricing with full traceability**

### 2.1 Calculate by Engagement Type

**PROJECT**: `estimated_hours × hourly.target × volume_discount`
- Hours from: typical_duration (weeks × 25 hrs/week) OR --value/hourly.target OR target_price/hourly.target
- Volume discounts: >$100K: 10%, >$50K: 5%, >$25K: 2.5%

**RETAINER**: `monthly_hours × effective_rate`
- Effective rates: Full-time (160h): minimum rate, Half (80h): target×0.85, Quarter (40h): target×0.90, Part (<40h): target×0.95

**STAFF-AUG**: `20_days × daily.target × duration_discount`
- Duration discounts: ≥6mo: 10%, ≥3mo: 5%

**WORKSHOP**: `(workshop_days × daily.premium) + (prep_days × daily.target) + materials + followup`
- Prep days: max(1, workshop_days × 0.5)
- Materials: $500/participant, Followup: 10 hours × hourly.target

### 2.2 Validation

**Range Check**: Compare to service pricing.range (allow 20% buffer), warn if outside
**Minimum Check**: Enforce minimums (project: $5K, retainer: $2K/mo, staff-aug: $8K/mo, workshop: $3K)

Generate pricing_calculation YAML metadata for frontmatter with method, rates, hours/days, discounts, final price.

---

## Phase 3: Generate 10-Section McKinsey/BCG Proposal

**Structure**: Cover → Executive Summary → Understanding Challenge → Proposed Approach → Deliverables/Timeline → Team/Qualifications → Investment → Terms → Next Steps → Appendix

**Generation Rules**:
- Use prospect_context for client-specific content (if has_prior_research=true), otherwise generic
- Include pricing_calculation YAML frontmatter
- Add evidence citations: `[Evidence: source_file → section]`
- Customize by engagement type (project/retainer/staff-aug/workshop)

### Section Templates (Adapt by Engagement Type)

**Cover Page**: YAML frontmatter with pricing_calculation metadata + title/client/consultant/date/validity (14 days)/reference (PROP-YYYYMMDD-ClientInit-ServiceInit)

**Executive Summary**:
- Client-specific: Use pain_points from prospect_context + strategic_priorities
- Generic: Use ideal_client.pain_points from service definition
- Include: problem statement, solution approach, expected outcomes (success_metrics), investment summary, payment terms, next steps timeline

**Understanding Your Challenge**:
- Client-specific: Current situation, pain points with business impact, urgency drivers from research, desired future state
- Generic: Common challenges in industry, typical business impact, why organizations seek service, success outcomes

**Proposed Approach** (varies by type):
- PROJECT: 3-4 phases with duration/objective/activities/deliverables/client involvement per phase + methodology notes
- RETAINER: Monthly hour allocation, services covered (strategic advisory + tactical support), hour tracking, overage policy, typical monthly cadence
- STAFF-AUG: Role definition, responsibilities, skills, working arrangement (reporting/schedule/location/tools), integration timeline, success metrics
- WORKSHOP: Overview, agenda (day-by-day sessions), materials, post-workshop support, preparation requirements, expected outcomes

**Deliverables & Timeline**:
- Complete deliverables list with format/delivery date/acceptance criteria
- Timeline table (Gantt-style for project, recurring for retainer, continuous for staff-aug, milestone-based for workshop)
- Key milestones with gates, regular check-in cadence

**Team & Qualifications**:
- If candidate_profile exists: Credentials, domain expertise, technical capabilities, relevant achievements (top 3-5 with evidence), certifications, thought leadership
- If no profile: Use differentiation section (competitive_advantages, proof_points, unique_value)

**Investment**:
- Pricing table by type (project phases, retainer tiers, staff-aug breakdown, workshop components)
- Payment terms (deposit %, progress/recurring payments, invoice terms, late fees, accepted methods)
- Expenses (included vs. billable)
- Pricing transparency (calculation method, validation, value justification with ROI)

**Terms & Conditions**:
- Scope (included/excluded deliverables)
- Assumptions (client access, timeline, environment)
- Change management process
- IP ownership (work product → client, pre-existing → consultant with license, confidentiality)
- Warranties & limitations
- Termination (for convenience vs. cause, notice periods by type)
- Engagement period & renewal
- Governing law & disputes

**Next Steps**:
- Review process, discussion scheduling, approval deadline (7-14 days), contract signing, kickoff timeline
- Client preparation checklist (administrative, access/resources, information)
- Contact info and customization options

**Appendix**:
- Detailed credentials (certifications, education, thought leadership, technical proficiencies from profile)
- Case studies (from service_definition or candidate profile projects)
- References (available upon request)
- Terms/definitions, proposal metadata (generated date, reference, sources)

---

## Phase 4: Save & Report

### 4.1 File Output

**Filename**: `Client_Prospects/Proposal_[Client]_[Service]_[YYYYMMDD].md`
- Sanitize names: PascalCase, no special chars, service max 30 chars

### 4.2 Summary Report

Display:
- Output file path
- Proposal details (client, service, type, investment, duration)
- Pricing calculation (method, rates, discounts, validation status)
- Proposal contents (sections, deliverables count, customization level)
- Prospect research integration status
- Sources used
- Next steps: review, customize, convert (formatresume/convert), send, handle changes
- Best practices: respond quickly (24-48h), personalize, be specific, show expertise, price transparently

**Warnings**:
- Below minimum viable engagement: Flag for review with options
- Price outside service range: Explain and recommend
- Stale prospect research (>30 days): Suggest refresh

---

## Error Handling

**Stop Execution Errors**:
1. No service definition found → Guide to `/defineservices`
2. Service not found → List available services
3. Incomplete service definition (missing required sections) → Guide to `/defineservices --update`
4. Pricing calculation failed → Provide --value or update service definition

**Continue with Default**:
- Invalid engagement type → Default to "project", warn user

---

## Usage Examples

1. **Generic project** (default): `/proposaltemplate`
2. **Client-specific**: `/proposaltemplate --client="Hatch" --service="AI Agent Development"`
3. **Retainer**: `/proposaltemplate --client="OxfordProperties" --service="Portfolio Advisory" --type=retainer`
4. **Workshop with budget**: `/proposaltemplate --service="AI Training" --type=workshop --value=15000`
5. **Staff augmentation**: `/proposaltemplate --client="TechStartup" --service="Senior AI Engineer" --type=staff-aug`

---

**Now execute proposal generation following all phases above...**

