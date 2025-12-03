---
description: Generate professional consulting proposal templates with pricing calculation
argument-hint: [--client=name] [--service=name] [--type=project|retainer|staff-aug|workshop] [--value=amount]
---

I'll generate a professional consulting proposal template tailored to your client, service offering, and engagement type with comprehensive pricing calculations.

**Arguments:**
- `--client=name`: Client/prospect company name (optional, defaults to "Generic")
- `--service=name`: Service name from your service definition (optional, uses primary service if not specified)
- `--type=TYPE`: Engagement type (optional, defaults to "project")
  - `project`: Fixed-scope project engagement
  - `retainer`: Ongoing advisory/support retainer
  - `staff-aug`: Staff augmentation/embedded consulting
  - `workshop`: Training/facilitation workshop
- `--value=amount`: Approximate engagement value for pricing (optional, calculated from rate card if not specified)

**Example Usage:**
```bash
# Generic project proposal
/proposaltemplate

# Client-specific AI Agent Development proposal
/proposaltemplate --client="Hatch" --service="AI Agent Development"

# Retainer proposal for portfolio advisory
/proposaltemplate --client="OxfordProperties" --service="Portfolio Advisory" --type=retainer

# Workshop proposal with target value
/proposaltemplate --service="AI Training" --type=workshop --value=15000
```

---

## Phase 1: Validate Prerequisites and Load Inputs

### 1.1 Check for Service Definition

**Search for service definition file:**
```bash
find Client_Prospects/ -name "Service_Definition_*.md" -type f | sort -r | head -1
```

**If no service definition found:**
```
❌ PREREQUISITE MISSING: Service Definition

No service definition file found in Client_Prospects/

**Required Action:**
Run /defineservices first to create your service catalog and pricing structure.

The service definition contains:
- Service offerings with deliverables
- Rate card (hourly, daily, retainer rates)
- Competitive differentiation and credentials
- Payment terms and engagement models

Usage:
  /defineservices              # Interactive guided mode
  /defineservices --from-profile  # Auto-generate from your profile

After creating your service definition, run /proposaltemplate again.
```

Stop execution and guide user.

**If service definition found:**
- Select the most recent by date (extract from filename `Service_Definition_YYYYMMDD.md`)
- Display confirmation:
  ```
  📄 Using service definition: Client_Prospects/Service_Definition_[Date].md
  Generated: [Date from file metadata]
  ```

### 1.2 Load Service Definition

Read the selected service definition file and parse:

**From markdown body - Consultant section:**
- `name`: Consultant/business name
- `tagline`: One-sentence value proposition
- `credentials[]`: Key credentials and certifications
- `years_experience`: Total professional experience
- `website`, `linkedin`, `email`: Contact information (if available)

**From markdown body - Services section:**
For each service offering, extract:
- `name`: Service name
- `category`: Service category
- `description`: Service description with value prop
- `deliverables[]`: Specific deliverables
- `ideal_client.industries[]`: Target industries
- `ideal_client.pain_points[]`: Problems solved
- `ideal_client.decision_makers[]`: Buyer roles
- `pricing.model`: Pricing model
- `pricing.range`: Minimum/target/premium pricing
- `pricing.typical_duration`: Typical engagement duration
- `pricing.value_justification`: ROI explanation
- `success_metrics[]`: How success is measured
- `case_studies[]`: Real examples (if available)

**From markdown body - Rate Card section:**
- `hourly`: Minimum/target/premium hourly rates
- `daily`: Minimum/target/premium daily rates
- `retainer`: Part-time/half-time/full-time options (if available)
- `payment_terms`: Deposit, invoice frequency, payment due, late fees
- `expenses`: Expense handling policy

**From markdown body - Engagement Models section:**
- `preferred[]`: Preferred engagement types
- `acceptable[]`: Acceptable engagement types
- `avoid[]`: Engagement types to decline
- `minimum_engagement`: Minimum engagement size
- `concurrent_clients`: Maximum concurrent clients (if available)

**From markdown body - Differentiation section:**
- `unique_value`: Core UVP (2-3 sentences)
- `competitive_advantages[]`: Advantages with impact
- `proof_points[]`: Quantified achievements
- `certifications[]`: Relevant certifications

**Validate required sections exist:**
- [ ] rate_card section with hourly and daily rates
- [ ] consultant name and credentials
- [ ] services catalog (at least 1 service)
- [ ] payment_terms with deposit and payment due

**If rate_card section missing or incomplete:**
```
❌ INVALID SERVICE DEFINITION: Missing Rate Card

Your service definition at Client_Prospects/Service_Definition_[Date].md
is missing the required rate_card section.

Please update your service definition with complete pricing information:
  /defineservices --update

Then run /proposaltemplate again.
```

Stop execution.

### 1.3 Parse Command-Line Arguments

**Extract client argument (`--client=name`):**
- If present: Extract client/company name
- If absent: Default to "Generic"
- Sanitize client name (remove special characters, standardize capitalization)

**Extract service argument (`--service=name`):**
- If present: Search for matching service in service definition
  - Use fuzzy matching (e.g., "AI Agent" matches "AI Agent Development for Due Diligence")
  - If no match found, show error with available services
- If absent: Use primary service (first service in catalog, or highest priority service)

**Service matching error:**
```
❌ SERVICE NOT FOUND: [service-name]

No service matching "[service-name]" in your service definition.

Available services from your service definition:
1. [Service 1 name]
2. [Service 2 name]
3. [Service 3 name]

Usage: /proposaltemplate --service="[exact or partial service name]"

Example: /proposaltemplate --service="AI Agent Development"
```

**Extract engagement type argument (`--type=TYPE`):**
- If present: Validate against allowed types: `project`, `retainer`, `staff-aug`, `workshop`
- If absent: Default to `project`
- If invalid type provided:
  ```
  ❌ INVALID ENGAGEMENT TYPE: [type]

  Valid engagement types:
  - project: Fixed-scope project engagement
  - retainer: Ongoing advisory/support retainer
  - staff-aug: Staff augmentation/embedded consulting
  - workshop: Training/facilitation workshop

  Usage: /proposaltemplate --type=project

  Defaulting to "project" type.
  ```

**Extract value argument (`--value=amount`):**
- If present: Parse numeric value (strip currency symbols, commas)
- If absent: Set to null (will be calculated from rate card)

**Argument summary confirmation:**
```
📋 PROPOSAL CONFIGURATION

Client: [Client name or "Generic"]
Service: [Service name from definition]
Type: [project|retainer|staff-aug|workshop]
Target Value: [Value if specified, or "TBD - will calculate from rate card"]

Generating proposal...
```

### 1.4 Load Pitch Deck (Optional - for prospect research integration)

**If client is not "Generic", search for existing pitch deck:**
```bash
find Client_Prospects/ -name "Pitch_*${client}*.md" -type f | sort -r | head -1
```

**If pitch deck found:**
- Load pitch deck file
- Extract `prospect_research` section from YAML frontmatter or markdown body
- Extract:
  - `company_name`: Verified company name
  - `pain_points_identified[]`: Client-specific pain points
  - `strategic_priorities[]`: Strategic priorities from research
  - `research_date`: When research was conducted
  - `sources_consulted[]`: Research sources
  - `decision_maker_status[]`: Current decision-makers

**Create prospect context:**
```yaml
prospect_context:
  has_prior_research: true
  research_date: "[Date]"
  pain_points: [pain points array]
  strategic_priorities: [priorities array]
  decision_makers: [decision maker array]
```

**If pitch deck not found:**
```yaml
prospect_context:
  has_prior_research: false
  research_date: null
  pain_points: []  # Will use generic service pain points
  strategic_priorities: []
  decision_makers: []
```

Display information message:
```
ℹ️ No pitch deck found for client "[Client]"

Proposal will use generic service positioning.

To create a more tailored proposal:
1. Run /pitchdeck --prospect="[Client]" to conduct prospect research
2. Then regenerate proposal: /proposaltemplate --client="[Client]"

Continuing with available information...
```

### 1.5 Load Candidate Profile (Optional - for credentials section)

**Check for candidate profile:**
`ResumeSourceFolder/.profile/candidate_profile.json`

**If profile exists:**
- Load certifications (Active status only)
- Load top 3-5 quantified achievements from work_history
- Extract thought leadership (publications, speaking)
- Use for "Team & Qualifications" section

**If profile missing:**
- Display warning but continue
- Use only service definition credentials
- Recommend running `/assessjob` to generate profile

```
⚠️ CANDIDATE PROFILE NOT FOUND

No candidate profile at ResumeSourceFolder/.profile/candidate_profile.json

"Team & Qualifications" section will have limited content.

To enhance proposal credentials:
  /assessjob <any-job-posting-file.md>

Continuing with service definition credentials only...
```

---

## Phase 2: Calculate Pricing

**THIS IS A CRITICAL SECTION - PRICING MUST BE TRANSPARENT AND DEFENSIBLE**

### 2.1 Determine Pricing Calculation Method

**Based on engagement type, select pricing algorithm:**

#### Algorithm 1: Project-Based Pricing

**When**: `--type=project`

**Calculation steps:**

**Step 1: Estimate project hours**
- If service has `typical_duration` (e.g., "4-6 weeks"), estimate hours:
  - Extract weeks from duration string
  - Assume 20-30 hours per week for typical project
  - Use midpoint: `estimated_hours = avg_weeks × 25 hours/week`
- If `--value` argument provided, back-calculate hours:
  - `estimated_hours = value / hourly_target_rate`
- If neither available, use service pricing range:
  - Use target price from service definition
  - Back-calculate: `estimated_hours = target_price / hourly_target_rate`

**Step 2: Calculate base price**
```python
base_price = estimated_hours × rate_card.hourly.target
```

**Step 3: Apply volume discounts (if applicable)**
```python
if base_price > 100000:
    discount_pct = 10  # 10% discount for projects >$100K
    discounted_price = base_price × 0.90
elif base_price > 50000:
    discount_pct = 5  # 5% discount for projects >$50K
    discounted_price = base_price × 0.95
elif base_price > 25000:
    discount_pct = 2.5  # 2.5% discount for projects >$25K
    discounted_price = base_price × 0.975
else:
    discount_pct = 0
    discounted_price = base_price
```

**Step 4: Round to nearest $100**
```python
final_price = round(discounted_price, -2)
```

**Pricing metadata (for YAML frontmatter):**
```yaml
pricing_calculation:
  method: "project pricing formula"
  estimated_hours: [hours]
  base_hourly_rate: "$[rate]/hour"
  subtotal: "$[base_price]"
  discount_applied: "[discount_pct]% (project >$[threshold])"
  total_investment: "$[final_price]"
```

#### Algorithm 2: Retainer Pricing

**When**: `--type=retainer`

**Calculation steps:**

**Step 1: Determine monthly hours**
- If `--value` argument provided, estimate hours:
  - Use rate card retainer options to find closest match
  - Or calculate: `monthly_hours = (value / rate_card.hourly.target) / 12 months`
- If service has retainer recommendations, use those
- Default assumptions:
  - Part-time: 20 hours/month
  - Half-time: 80 hours/month
  - Full-time: 160 hours/month

**Step 2: Determine effective hourly rate (with retainer discount)**
```python
if monthly_hours >= 160:  # Full-time
    effective_rate = rate_card.hourly.minimum  # ~30% discount
    retainer_tier = "full-time"
elif monthly_hours >= 80:  # Half-time
    effective_rate = rate_card.hourly.target × 0.85  # ~15% discount
    retainer_tier = "half-time"
elif monthly_hours >= 40:  # Quarter-time
    effective_rate = rate_card.hourly.target × 0.90  # ~10% discount
    retainer_tier = "quarter-time"
else:  # Part-time (<40 hours)
    effective_rate = rate_card.hourly.target × 0.95  # ~5% discount
    retainer_tier = "part-time"
```

**Step 3: Calculate monthly fee**
```python
monthly_fee = monthly_hours × effective_rate
```

**Step 4: Round to nearest $100**
```python
final_monthly_fee = round(monthly_fee, -2)
```

**Pricing metadata (for YAML frontmatter):**
```yaml
pricing_calculation:
  method: "retainer pricing formula"
  monthly_hours: [hours]
  retainer_tier: "[tier]"
  effective_hourly_rate: "$[effective_rate]/hour"
  discount_from_standard: "[discount_pct]%"
  monthly_investment: "$[final_monthly_fee]"
  annual_investment: "$[final_monthly_fee × 12]"
```

#### Algorithm 3: Staff Augmentation Pricing

**When**: `--type=staff-aug`

**Calculation steps:**

**Step 1: Determine daily rate**
- Use `rate_card.daily.target` as base rate
- If client is enterprise or long-term engagement (>3 months), consider minimum rate
- If rush or specialized expertise, consider premium rate

**Step 2: Calculate monthly fee**
```python
days_per_month = 20  # Standard business days
monthly_fee = days_per_month × rate_card.daily.target
```

**Step 3: Apply long-term engagement discount (if applicable)**
```python
if engagement_duration_months >= 6:
    discount_pct = 10
    monthly_fee = monthly_fee × 0.90
elif engagement_duration_months >= 3:
    discount_pct = 5
    monthly_fee = monthly_fee × 0.95
else:
    discount_pct = 0
```

**Step 4: Round to nearest $100**
```python
final_monthly_fee = round(monthly_fee, -2)
```

**Pricing metadata (for YAML frontmatter):**
```yaml
pricing_calculation:
  method: "staff augmentation pricing formula"
  daily_rate: "$[rate]/day"
  days_per_month: 20
  discount_applied: "[discount_pct]% (engagement >[X] months)"
  monthly_investment: "$[final_monthly_fee]"
  estimated_total: "$[final_monthly_fee × duration_months]"
```

#### Algorithm 4: Workshop Pricing

**When**: `--type=workshop`

**Calculation steps:**

**Step 1: Determine workshop days**
- If service has typical duration for workshops, use that
- If `--value` provided, back-calculate: `workshop_days = value / (daily.premium × 1.5)`
- Default: 1-day workshop

**Step 2: Calculate preparation days**
```python
# Prep is typically 50-100% of delivery time
prep_days = max(1, workshop_days × 0.5)  # Minimum 1 day prep
```

**Step 3: Calculate total fee**
```python
# Workshops use premium daily rate (includes materials, customization)
delivery_fee = workshop_days × rate_card.daily.premium
prep_fee = prep_days × rate_card.daily.target

total_fee = delivery_fee + prep_fee
```

**Step 4: Add materials and follow-up support (optional)**
```python
materials_fee = 500  # Per participant workbook/materials
followup_hours = 10  # Post-workshop support hours
followup_fee = followup_hours × rate_card.hourly.target

if include_materials_and_followup:
    total_fee += (materials_fee × estimated_participants) + followup_fee
```

**Step 5: Round to nearest $100**
```python
final_fee = round(total_fee, -2)
```

**Pricing metadata (for YAML frontmatter):**
```yaml
pricing_calculation:
  method: "workshop pricing formula"
  workshop_days: [days]
  preparation_days: [prep_days]
  delivery_rate: "$[daily.premium]/day"
  preparation_rate: "$[daily.target]/day"
  materials_and_followup: "$[materials_fee + followup_fee]"
  total_investment: "$[final_fee]"
```

### 2.2 Pricing Validation

**Validate calculated price against service definition pricing range:**

```python
matching_service = find_matching_service(services[], selected_service_name)

if matching_service and matching_service.pricing.range:
    # Parse pricing range (handles formats like "$5,000-$15,000", "$150-$200/hr")
    range_min, range_max = parse_price_range(matching_service.pricing.range)

    issues = []

    # Check if calculated price is within range (allow 20% buffer)
    if calculated_price < range_min × 0.8:
        issues.append({
            "severity": "WARNING",
            "message": f"Calculated price ${calculated_price} below service range minimum ${range_min}",
            "recommendation": "Consider increasing scope or using minimum pricing tier"
        })

    if calculated_price > range_max × 1.2:
        issues.append({
            "severity": "WARNING",
            "message": f"Calculated price ${calculated_price} above service range maximum ${range_max}",
            "recommendation": "Consider breaking into multiple phases or using phased pricing"
        })
```

**Check minimum viable engagement:**
```python
min_engagement = {
    "project": 5000,
    "retainer": 2000,  # Per month
    "staff-aug": 8000,  # Per month
    "workshop": 3000
}

if calculated_price < min_engagement.get(engagement_type, 0):
    issues.append({
        "severity": "ERROR",
        "message": f"Price ${calculated_price} below minimum viable {engagement_type} engagement (${min_engagement[engagement_type]})",
        "recommendation": "Increase scope or decline engagement"
    })
```

**Display validation results:**

```
✅ PRICING VALIDATION COMPLETE

Calculated Investment: $[calculated_price]
Service Range: $[range_min] - $[range_max]
Status: [WITHIN RANGE | BELOW MINIMUM | ABOVE MAXIMUM]

[If warnings exist:]
⚠️ PRICING ADVISORIES:
- [Warning 1]
  Recommendation: [Recommendation 1]
- [Warning 2]
  Recommendation: [Recommendation 2]

[If errors exist:]
❌ PRICING ERRORS:
- [Error 1]
  Recommendation: [Recommendation 1]

Proceeding with calculated pricing. Review manually before sending to client.
```

**If price is below minimum viable engagement:**
```
⚠️ BELOW MINIMUM ENGAGEMENT

Calculated price: $[calculated_price]
Minimum viable [type] engagement: $[minimum]

**Options:**
1. Increase scope to meet minimum ($[difference] more)
2. Combine with another service for bundled pricing
3. Decline engagement as not economically viable

Proposal will include calculated price, but flagged for your review.
```

---

## Phase 3: Generate Proposal Document (McKinsey/BCG Style)

### 3.1 Proposal Structure Definition

**Standard consulting proposal structure (10 sections):**

| Section | Content | Purpose |
|---------|---------|---------|
| **Cover Page** | Title, client, date, consultant | Professional first impression |
| **Executive Summary** | Problem, solution, value, investment | Decision-maker one-pager |
| **Understanding Your Challenge** | Client situation, pain points, stakes | Demonstrate understanding |
| **Proposed Approach** | Methodology, phases, activities | Show how work gets done |
| **Deliverables & Timeline** | What client receives, milestones | Set clear expectations |
| **Team & Qualifications** | Credentials, relevant experience | Build trust and credibility |
| **Investment** | Pricing, payment terms, options | Transparent cost structure |
| **Terms & Conditions** | Scope boundaries, assumptions, IP | Protect both parties |
| **Next Steps** | How to proceed, timeline to start | Clear call-to-action |
| **Appendix** | Detailed credentials, case studies | Supporting documentation |

### 3.2 Generate Cover Page

```markdown
---
proposal_type: [project | retainer | staff-aug | workshop]
client: "[Client name or Generic]"
service: "[Service name]"
generated: [ISO8601 timestamp]
valid_until: [14 days from generation]
total_investment: "$[calculated_price]"
pricing_calculation:
  method: "[pricing method from calculation]"
  base_rate: "$[rate used]"
  [additional pricing metadata from Phase 2]
  discount_applied: "[discount description]"
consultant: "[Consultant name]"
service_definition_source: "[File path to service definition]"
---

# PROPOSAL

## [Service Name]

### Prepared for [Client Name]

---

**Prepared by:**
[Consultant Name]
[Credentials separated by |]

**Date:** [Current date formatted as "December 3, 2025"]
**Valid Until:** [Date 14 days from now]
**Proposal Reference:** PROP-[YYYYMMDD]-[ClientInitials]-[ServiceInitials]

---

**Contact Information:**
[Email if available]
[Phone if available]
[LinkedIn if available]
[Website if available]

---

*This proposal outlines our approach to delivering [Service Name] for [Client Name]. All information is confidential and proprietary.*

---
```

### 3.3 Generate Executive Summary

**Content depends on whether prospect research available:**

**If prospect_context.has_prior_research = true (client-specific):**
```markdown
## Executive Summary

[Client Name] is [brief context from research - e.g., "navigating significant portfolio expansion while managing operational complexity across distributed assets"].

Based on our understanding of your needs, we propose a [engagement type] engagement focused on [service name] to address:

1. **[Primary pain point from research]**: [Brief impact statement]
2. **[Secondary pain point from research]**: [Brief impact statement]
3. **[Tertiary pain point or strategic priority]**: [Brief impact statement]

### Our Solution

[Service description from service definition - 2-3 sentences explaining approach]

### Expected Outcomes

By engaging our services, [Client Name] can expect:

- **[Success metric 1 from service definition]**: [Explanation]
- **[Success metric 2 from service definition]**: [Explanation]
- **[Success metric 3 from service definition]**: [Explanation]

### Investment

**Total Investment:** $[calculated_price]

[If retainer:] **Monthly Investment:** $[monthly_fee]
[If staff-aug:] **Monthly Investment:** $[monthly_fee] ([X] days/month)
[If project:] **Duration:** [Typical duration from service definition]
[If workshop:] **Duration:** [Workshop days] days delivery + [Prep days] days preparation

**Payment Terms:** [Deposit percentage]% upfront, [payment structure from payment_terms]

### Next Steps

To proceed:
1. Review and approve this proposal by [Date 7 days from now]
2. Sign engagement agreement
3. Schedule kickoff meeting
4. Begin [Phase 1 name] immediately

**Timeline to Start:** [X] days from approval

---

[Evidence: Service_Definition_[Date].md → [Service name] section]
[If research used: Pitch_[Client]_[Date].md → prospect_research section]
```

**If prospect_context.has_prior_research = false (generic):**
```markdown
## Executive Summary

Organizations in [industries from ideal_client] frequently face challenges with [pain points from service definition].

This proposal outlines our [engagement type] approach to [service name], designed to [value proposition from service description].

### The Challenge

[Pain point 1 from service ideal_client.pain_points]

[Pain point 2 from service ideal_client.pain_points]

[Pain point 3 from service ideal_client.pain_points]

### Our Solution

[Service description from service definition - 2-3 sentences]

### Expected Outcomes

- **[Success metric 1]**: [Explanation]
- **[Success metric 2]**: [Explanation]
- **[Success metric 3]**: [Explanation]

### Investment

**Total Investment:** $[calculated_price]

[Rest of investment details as above]

### Next Steps

To discuss how this approach can be tailored to your specific needs:
1. Schedule discovery call: [Calendar link or email]
2. Provide additional context about your situation
3. Receive customized proposal with specific pricing

**Timeline to Start:** [X] days from approval

---

[Evidence: Service_Definition_[Date].md → [Service name] section]
```

### 3.4 Generate "Understanding Your Challenge" Section

**If prospect research available:**
```markdown
## Understanding Your Challenge

### Current Situation

Based on our research and understanding of [Client Name], you are currently:

[If strategic priorities exist:]
**Strategic Focus:**
- [Priority 1 from research]
- [Priority 2 from research]

[If pain points identified:]
**Key Challenges:**

#### [Pain point 1 title]
[Description of pain point from research]

**Business Impact:**
- [Specific impact if available, or infer from service definition]
- [Cost/risk/opportunity cost if quantified in research]

#### [Pain point 2 title]
[Description of pain point from research]

**Business Impact:**
- [Specific impact]

### Why This Matters Now

[Urgency driver from research - e.g., "With portfolio expansion planned for Q1 2026..."]

[Competitive pressure, regulatory deadline, technology shift, or market opportunity from research]

### Desired Future State

[Outcome goals that service can deliver - aligned with success metrics]

---

**Sources:** [If research conducted, cite URLs from prospect_research.sources_consulted]

[Evidence: Pitch_[Client]_[Date].md → prospect_research section]
```

**If no prospect research (generic):**
```markdown
## Understanding Your Challenge

### Common Challenges in [Industry/Domain]

Organizations like yours in [industry from ideal_client] typically face:

#### [Pain point 1 from service definition]
[Description of pain point]

**Typical Business Impact:**
- [Generic impact statement]
- [Cost/risk pattern common to this pain point]

#### [Pain point 2 from service definition]
[Description of pain point]

**Typical Business Impact:**
- [Generic impact statement]

### Why Organizations Seek This Service

[Value justification from service pricing section]

### What Success Looks Like

After engagement completion, clients typically achieve:
- [Success metric 1]
- [Success metric 2]
- [Success metric 3]

---

[Evidence: Service_Definition_[Date].md → [Service].ideal_client.pain_points]
```

### 3.5 Generate "Proposed Approach" Section

**Structure depends on engagement type:**

**For PROJECT engagements:**
```markdown
## Proposed Approach

Our [service name] methodology follows a proven [X]-phase approach designed to deliver results efficiently while minimizing disruption to your operations.

### Phase 1: [Phase name - e.g., "Discovery & Assessment"]

**Duration:** [Timeframe - e.g., "Week 1" or "Days 1-5"]

**Objective:** [What this phase achieves]

**Activities:**
- [Activity 1 - specific to service]
- [Activity 2]
- [Activity 3]
- [Activity 4]

**Deliverables:**
- [Deliverable 1 from service deliverables list]
- [Deliverable 2]

**Client Involvement Required:**
- [Expectation 1 - e.g., "2-hour kickoff meeting"]
- [Expectation 2 - e.g., "Access to [systems/people/data]"]

---

### Phase 2: [Phase name - e.g., "Analysis & Design"]

**Duration:** [Timeframe - e.g., "Weeks 2-3"]

**Objective:** [What this phase achieves]

**Activities:**
- [Activity 1]
- [Activity 2]
- [Activity 3]

**Deliverables:**
- [Deliverable 3 from service deliverables list]
- [Deliverable 4]

**Client Involvement Required:**
- [Expectation 1]
- [Expectation 2]

---

### Phase 3: [Phase name - e.g., "Implementation"]

**Duration:** [Timeframe - e.g., "Weeks 4-6"]

**Objective:** [What this phase achieves]

**Activities:**
- [Activity 1]
- [Activity 2]
- [Activity 3]

**Deliverables:**
- [Deliverable 5 from service deliverables list]
- [Deliverable 6]

**Client Involvement Required:**
- [Expectation 1]
- [Expectation 2]

---

### Phase 4: [Phase name - e.g., "Optimization & Handoff"]

**Duration:** [Timeframe - e.g., "Week 7"]

**Objective:** [What this phase achieves]

**Activities:**
- [Activity 1]
- [Activity 2]
- [Activity 3]

**Deliverables:**
- [Deliverable 7 from service deliverables list]
- [Final deliverable - e.g., "Final presentation and recommendations"]

**Client Involvement Required:**
- [Expectation 1 - e.g., "Final review meeting"]
- [Expectation 2 - e.g., "Sign-off on deliverables"]

---

### Methodology Notes

**Our Approach:**
[Brief explanation of methodology - e.g., "Agile project management with weekly sprints" or "Structured consulting framework based on [framework name]"]

**Quality Assurance:**
- [QA measure 1 - e.g., "Weekly status meetings and progress reports"]
- [QA measure 2 - e.g., "Peer review of all deliverables"]
- [QA measure 3 - e.g., "Client feedback loops at each phase gate"]

**Project Management:**
- [PM approach - e.g., "Dedicated project manager for coordination"]
- [Communication cadence - e.g., "Daily standups, weekly status reports"]
- [Tools used - e.g., "Asana for task tracking, Slack for communication"]

---

[Evidence: Service_Definition_[Date].md → [Service] deliverables + engagement_models]
```

**For RETAINER engagements:**
```markdown
## Proposed Approach

### Retainer Structure

**Monthly Hour Allocation:** [Monthly hours from calculation]
**Retainer Tier:** [Part-time/Half-time/Full-time]
**Effective Rate:** $[effective_rate]/hour ([discount_pct]% discount from standard rate)

### Services Covered Under Retainer

Your retainer includes access to:

**Strategic Advisory:**
- [Service activity 1 - e.g., "Monthly strategy sessions (2-4 hours)"]
- [Service activity 2 - e.g., "Ad-hoc consultation via email/calls"]
- [Service activity 3 - e.g., "Quarterly planning and roadmap reviews"]

**Tactical Support:**
- [Support activity 1 - e.g., "Review of documents and deliverables"]
- [Support activity 2 - e.g., "Attendance at key meetings as needed"]
- [Support activity 3 - e.g., "Troubleshooting and problem-solving"]

**Deliverables (Monthly):**
- [Deliverable 1 - e.g., "Monthly advisory report with recommendations"]
- [Deliverable 2 - e.g., "Ad-hoc analyses and memos as needed"]
- [Deliverable 3 - e.g., "Quarterly strategic planning document"]

### How the Retainer Works

**Hour Tracking:**
- Hours logged in [timeframe increments - e.g., "15-minute increments"]
- Monthly report provided showing hours used and activities
- Unused hours do not roll over month-to-month

**Overage Hours:**
- Hours beyond monthly allocation billed at [overage rate from rate_card or standard hourly rate]
- Overages invoiced separately at month-end
- You'll receive notification when approaching hour limit

**Scope:**
- Retainer covers [scope description - e.g., "advisory and strategic guidance related to [service area]"]
- Excluded: [exclusions - e.g., "Implementation work, hands-on execution, training delivery"]
- For projects beyond retainer scope, we'll provide separate proposals

**Availability:**
- Response time: [commitment - e.g., "Within 24 hours for email, same-day for urgent calls"]
- Scheduling: [approach - e.g., "Recurring monthly meetings + ad-hoc sessions as needed"]
- Communication: [channels - e.g., "Email, phone, video calls, Slack"]

### Typical Monthly Cadence

**Week 1:**
- Monthly strategy session (2 hours)
- Review of previous month's activities

**Weeks 2-3:**
- Ad-hoc support as needed
- Document reviews and feedback
- Short consultations (30-60 min each)

**Week 4:**
- Monthly report preparation and delivery
- Planning for next month

**Continuous:**
- Email support and quick questions
- Monitoring and proactive recommendations

---

[Evidence: Service_Definition_[Date].md → rate_card.retainer + engagement_models]
```

**For STAFF-AUGMENTATION engagements:**
```markdown
## Proposed Approach

### Staff Augmentation Model

**Role:** [Service name - e.g., "Senior AI Engineer"]
**Commitment:** [Days per month from calculation - typically 20 days/month]
**Daily Rate:** $[daily_rate]/day
**Monthly Investment:** $[monthly_fee]

### Role Definition

**Primary Responsibilities:**
- [Responsibility 1 aligned with service deliverables]
- [Responsibility 2]
- [Responsibility 3]
- [Responsibility 4]
- [Responsibility 5]

**Deliverables (Monthly):**
- [Deliverable 1 - e.g., "Code/implementations per sprint plan"]
- [Deliverable 2 - e.g., "Technical documentation"]
- [Deliverable 3 - e.g., "Weekly progress reports"]

**Skills Applied:**
[Top 5-7 technical skills from service definition or candidate profile relevant to service]

### Working Arrangement

**Reporting Structure:**
- Reports to: [Your decision-maker role - e.g., "CTO" or "VP Engineering"]
- Day-to-day direction: [Team lead role if different]
- Performance reviews: [Cadence - e.g., "Monthly with [Name]"]

**Work Schedule:**
- [Days per week - e.g., "Monday-Friday, full days"]
- [Hours per day - e.g., "8 hours/day, 9am-5pm [timezone]"]
- [Flexibility - e.g., "Flexible within core hours 10am-3pm [timezone]"]

**Location:**
- [Remote/Hybrid/On-site]
- [If hybrid:] [On-site days and location]
- [If on-site:] [Workspace requirements - e.g., "Dedicated desk with dual monitors"]

**Tools & Access:**
- [System access 1 - e.g., "GitHub repository access"]
- [System access 2 - e.g., "Production environment read access"]
- [Tools provided by you - e.g., "Laptop, development environment"]
- [Tools I provide - e.g., "My own laptop and software licenses"]

### Integration Approach

**Week 1: Onboarding**
- Introduction to team and stakeholders
- System access setup and verification
- Review of codebase, architecture, and documentation
- Establish communication channels and workflows

**Weeks 2-4: Ramping Up**
- Shadow team members on existing workflows
- Take ownership of well-defined tasks
- Begin contributing to sprints/projects
- Daily standups and integration into team rituals

**Month 2+: Full Contribution**
- Fully integrated into team workflows
- Leading technical workstreams
- Mentoring junior team members (if applicable)
- Contributing to architectural decisions

### Success Metrics

We'll measure success by:
- [Metric 1 - e.g., "Velocity: Story points completed per sprint"]
- [Metric 2 - e.g., "Quality: Code review approval rate >90%"]
- [Metric 3 - e.g., "Impact: [Business outcome metric from service definition]"]
- [Metric 4 - e.g., "Team integration: Feedback from team lead (monthly)"]

---

[Evidence: Service_Definition_[Date].md → [Service] + rate_card.daily]
```

**For WORKSHOP engagements:**
```markdown
## Proposed Approach

### Workshop Overview

**Title:** [Workshop name based on service]
**Format:** [Delivery days]-day [format - e.g., "interactive training workshop"]
**Audience:** [Target participants - e.g., "Technical team leads and senior engineers (10-20 participants)"]
**Learning Objectives:**
- [Learning objective 1 from service deliverables]
- [Learning objective 2]
- [Learning objective 3]

### Workshop Agenda

#### Day 1: [Module name]

**Morning (9:00 AM - 12:00 PM)**
- [Session 1 title] (60 min)
  - [Topics covered]
- [Session 2 title] (90 min)
  - [Topics covered]

**Afternoon (1:00 PM - 5:00 PM)**
- [Session 3 title] (90 min)
  - [Topics covered]
- [Hands-on exercise/workshop] (90 min)
  - [Activity description]

[If multi-day workshop, continue for each day]

#### Day 2: [Module name]

[Similar structure]

### Learning Materials

**Included in Workshop:**
- [Material 1 - e.g., "Printed workbook with exercises and templates"]
- [Material 2 - e.g., "Digital slide deck (PDF)"]
- [Material 3 - e.g., "Code samples and reference implementations"]
- [Material 4 - e.g., "Resource list and further reading"]
- [Material 5 - e.g., "Completion certificate"]

**Post-Workshop Support:**
- [Support 1 - e.g., "10 hours of follow-up consultation (via email/calls) over 90 days"]
- [Support 2 - e.g., "Access to private Slack channel for questions"]
- [Support 3 - e.g., "Monthly office hours (30 min) for 3 months"]

### Preparation Required

**From Your Team:**
- [Requirement 1 - e.g., "Identify participants (10-20 people)"]
- [Requirement 2 - e.g., "Pre-workshop survey completion (15 min per person)"]
- [Requirement 3 - e.g., "Prepare 2-3 real use cases to workshop during hands-on sessions"]

**Logistics:**
- [Logistics 1 - e.g., "Meeting room with projector and whiteboard"]
- [Logistics 2 - e.g., "WiFi access for all participants"]
- [Logistics 3 - e.g., "Participant laptops with [software] pre-installed"]

**Timeline:**
- [Timeline item 1 - e.g., "2 weeks before: Finalize participant list and pre-workshop survey"]
- [Timeline item 2 - e.g., "1 week before: Pre-workshop call with organizer (30 min)"]
- [Timeline item 3 - e.g., "Workshop delivery: [Dates]"]
- [Timeline item 4 - e.g., "Follow-up support: 90 days post-workshop"]

### Expected Outcomes

After completing this workshop, participants will:
- [Outcome 1 - e.g., "Understand core concepts of [topic] and be able to explain them to colleagues"]
- [Outcome 2 - e.g., "Apply [methodology/framework] to their own projects"]
- [Outcome 3 - e.g., "Have hands-on experience building [deliverable]"]
- [Outcome 4 - e.g., "Know where to find resources and continue learning"]

**Organizational Impact:**
- [Impact 1 from service success_metrics]
- [Impact 2 from service success_metrics]

---

[Evidence: Service_Definition_[Date].md → [Service] deliverables]
```

### 3.6 Generate "Deliverables & Timeline" Section

**Extract all deliverables from service definition and map to phases:**

```markdown
## Deliverables & Timeline

### Complete Deliverables List

[For each deliverable from service definition:]

#### [Deliverable name]

**Description:** [What this deliverable includes]

**Format:** [Format - e.g., "PDF report", "Excel model", "Code repository", "Slide deck"]

**Delivery Date:** [Phase and timeline - e.g., "End of Phase 1 (Week 2)", "Monthly recurring", "Final delivery"]

**Acceptance Criteria:**
- [Criterion 1 - what defines completion/success]
- [Criterion 2]

---

[Continue for all deliverables]

### Project Timeline (Gantt-Style Table)

[For PROJECT type:]

| Phase | Week(s) | Key Activities | Deliverables | Status Gates |
|-------|---------|----------------|--------------|--------------|
| Phase 1: [Name] | 1 | [Activities] | [Deliverables] | Kickoff meeting |
| Phase 2: [Name] | 2-3 | [Activities] | [Deliverables] | Phase 1 approval |
| Phase 3: [Name] | 4-6 | [Activities] | [Deliverables] | Phase 2 approval |
| Phase 4: [Name] | 7 | [Activities] | [Deliverables] | Final sign-off |

**Total Duration:** [X] weeks from kickoff

[For RETAINER type:]

| Month | Recurring Deliverables | Ad-Hoc Support |
|-------|------------------------|----------------|
| Month 1 | [Monthly deliverables] | [Hour allocation] hours |
| Month 2 | [Monthly deliverables] | [Hour allocation] hours |
| Month 3+ | [Monthly deliverables] | [Hour allocation] hours |

**Ongoing:** Monthly reports, quarterly planning sessions

[For STAFF-AUG type:]

| Timeframe | Deliverables | Integration Milestones |
|-----------|--------------|------------------------|
| Week 1 | Onboarding complete, access verified | Team introduction complete |
| Weeks 2-4 | [Sprint deliverables per sprint plan] | Ramping up, contributing to sprints |
| Month 2+ | [Ongoing sprint deliverables] | Fully integrated, leading workstreams |

**Continuous:** Code contributions, documentation, weekly progress reports

[For WORKSHOP type:]

| Milestone | Date | Deliverable |
|-----------|------|-------------|
| 2 weeks prior | [Date] | Pre-workshop survey sent, participant list confirmed |
| 1 week prior | [Date] | Pre-workshop call completed, materials finalized |
| Workshop Day 1 | [Date] | [Module 1 completion], workbook Part 1 |
| Workshop Day 2 | [Date] | [Module 2 completion], final presentations, certificates |
| Post-workshop | [90 days] | Follow-up support available, office hours |

### Milestones & Check-ins

**Key Milestones:**
1. **[Milestone 1 - e.g., "Project Kickoff"]**: [Date relative to start - e.g., "Day 1"]
   - Deliverable: [Kickoff meeting notes, project plan confirmed]

2. **[Milestone 2 - e.g., "Phase 1 Completion"]**: [Date - e.g., "End of Week 2"]
   - Deliverable: [Deliverable name]
   - Gate: Client approval to proceed to Phase 2

3. **[Milestone 3]**: [Date]
   - Deliverable: [Deliverable name]
   - Gate: [Approval/review required]

4. **[Final Milestone - e.g., "Final Delivery"]**: [Date - e.g., "End of Week 7"]
   - Deliverable: [All final deliverables]
   - Gate: Client sign-off and project closure

**Regular Check-ins:**
- [Frequency - e.g., "Weekly"] status meetings ([duration - e.g., "30 minutes"])
- [Frequency - e.g., "Bi-weekly"] steering committee updates (if applicable)
- [Frequency - e.g., "Monthly"] executive summaries

---

[Evidence: Service_Definition_[Date].md → [Service].deliverables + typical_duration]
```

### 3.7 Generate "Team & Qualifications" Section

**Content depends on availability of candidate profile:**

**If candidate profile available:**
```markdown
## Team & Qualifications

### Lead Consultant: [Consultant name]

**Credentials:**
[Credentials from service definition separated by " | "]

**Experience:**
[Years experience] years of professional experience across [industries/domains from candidate profile]

### Relevant Expertise

**Domain Knowledge:**
[If candidate profile has domain_expertise matching service ideal_client industries:]
- **[Domain 1]**: [Years] years, [depth level - e.g., "Expert" or "Deep domain knowledge"]
  - [Context from work history if available]
- **[Domain 2]**: [Years] years, [depth level]

**Technical Capabilities:**
[Extract from technical_skills[] where relevant to service:]
- **[Skill 1]**: [Proficiency level - Expert/Proficient] ([Years] years)
- **[Skill 2]**: [Proficiency level] ([Years] years)
- **[Skill 3]**: [Proficiency level] ([Years] years)

### Relevant Achievements

[Extract top 3-5 achievements from candidate profile relevant to service:]

#### [Achievement 1 - most relevant to service]
[Quantified metric from work_history.achievements]

**Context:** [Brief explanation of what was achieved]

**Relevance:** [Why this matters for this engagement]

[Evidence: candidate_profile.json → work_history[X].achievements[Y] → [source file], lines [A-B]]

---

[Continue for 2-4 more achievements]

### Certifications & Credentials

[Extract from certifications[] where status = "Active" and relevant to service:]

| Certification | Issuing Organization | Year | Status |
|---------------|---------------------|------|--------|
| [Cert 1] | [Organization] | [Year] | Active |
| [Cert 2] | [Organization] | [Year] | Active |
| [Cert 3] | [Organization] | [Year] | Active |

[Evidence: candidate_profile.json → certifications[]]

### Thought Leadership (if available)

[If thought_leadership exists in profile:]

**Publications:**
- [Publication 1 - if relevant to service]
- [Publication 2]

**Speaking Engagements:**
- [Engagement 1 - if relevant to service]
- [Engagement 2]

**Frameworks/Methodologies Created:**
- [Framework 1 - if relevant to service]

[Evidence: candidate_profile.json → thought_leadership]

### Why This Experience Matters

[2-3 sentences explaining how consultant's background specifically qualifies them for this service:]

[Unique value from service definition differentiation.unique_value]

---

[Evidence: Service_Definition_[Date].md + candidate_profile.json]
```

**If candidate profile NOT available (use service definition only):**
```markdown
## Team & Qualifications

### Lead Consultant: [Consultant name]

**Credentials:**
[Credentials from service definition separated by " | "]

**Experience:**
[Years experience] years of professional experience

### Competitive Advantages

[Extract from service definition differentiation.competitive_advantages:]

#### [Advantage 1]
**What:** [Advantage description]

**Impact:** [Why this matters to clients]

[Evidence: Service_Definition_[Date].md → differentiation.competitive_advantages[0]]

---

[Continue for 2-3 advantages]

### Proven Results

[Extract from service definition differentiation.proof_points:]

- **[Proof point 1]**: [Metric with context]
- **[Proof point 2]**: [Metric with context]
- **[Proof point 3]**: [Metric with context]

[Evidence: Service_Definition_[Date].md → differentiation.proof_points]

### Unique Value Proposition

[Service definition differentiation.unique_value - 2-3 sentences]

---

**Detailed credentials and case studies available in Appendix.**

[Evidence: Service_Definition_[Date].md → differentiation section]
```

### 3.8 Generate "Investment" Section

```markdown
## Investment

### Pricing Structure

[Content varies by engagement type:]

[For PROJECT:]

| Component | Investment |
|-----------|------------|
| Phase 1: [Name] | $[calculated from hours/rate] |
| Phase 2: [Name] | $[calculated] |
| Phase 3: [Name] | $[calculated] |
| Phase 4: [Name] | $[calculated] |
| **Total Project Investment** | **$[total_calculated_price]** |

[If volume discount applied:]
*Pricing includes [discount_pct]% discount for project value >[threshold].*

---

[For RETAINER:]

| Retainer Tier | Monthly Investment | Hours Included | Effective Rate |
|---------------|-------------------|----------------|----------------|
| **[Selected tier]** | **$[monthly_fee]** | [hours] hours | $[effective_rate]/hr |

**Annual Investment:** $[monthly_fee × 12] ([X]% savings vs. hourly)

**Overage Rate:** $[overage_rate or standard_hourly]/hour for hours beyond monthly allocation

---

[For STAFF-AUG:]

| Component | Rate | Investment |
|-----------|------|------------|
| Daily rate | $[daily_rate]/day | - |
| Days per month | 20 days | $[monthly_fee] |
| **Monthly Investment** | | **$[monthly_fee]** |

[If long-term discount:]
*Pricing includes [discount_pct]% discount for [X]+ month commitment.*

**Estimated [Duration] Month Total:** $[monthly_fee × duration]

---

[For WORKSHOP:]

| Component | Investment |
|-----------|------------|
| Workshop delivery ([X] days) | $[delivery_fee] |
| Preparation and customization ([Y] days) | $[prep_fee] |
| Materials and workbooks ([Z] participants) | $[materials_fee] |
| Post-workshop support (90 days) | $[followup_fee] |
| **Total Workshop Investment** | **$[total_fee]** |

### Payment Terms

**Deposit:** [Deposit percentage from payment_terms]% ($[deposit_amount]) due upon signing

[For PROJECT:]
**Progress Payments:**
- [Payment 1]: [Percentage]% ($[amount]) upon [milestone - e.g., "Phase 1 completion"]
- [Payment 2]: [Percentage]% ($[amount]) upon [milestone - e.g., "Phase 2 completion"]
- [Final payment]: [Percentage]% ($[amount]) upon [milestone - e.g., "final delivery and acceptance"]

[For RETAINER:]
**Monthly Billing:**
- Retainer fee invoiced [invoice_frequency from payment_terms - e.g., "monthly in advance"]
- Overage hours invoiced separately at month-end
- Payment due [payment_due from payment_terms - e.g., "Net 15 days from invoice"]

[For STAFF-AUG:]
**Monthly Billing:**
- Invoiced [invoice_frequency - e.g., "bi-weekly"] based on days worked
- Payment due [payment_due - e.g., "Net 15 days from invoice"]

[For WORKSHOP:]
**Payment Schedule:**
- 50% ($[amount]) upon signing
- 50% ($[amount]) [timeframe - e.g., "7 days prior to workshop delivery"]

**Invoice Terms:** [Payment due from payment_terms - e.g., "Net 15 days from invoice date"]

**Late Payments:** [Late fee from payment_terms - e.g., "2% per month" or "None"]

**Accepted Payment Methods:**
- Wire transfer (preferred)
- Check
- ACH/EFT
- Credit card ([note surcharge if applicable - e.g., "3% processing fee applies"])

### Expenses

[Extract from rate_card.expenses:]

**Included in Pricing:**
- [Included 1 - e.g., "Standard software and collaboration tools"]
- [Included 2 - e.g., "Virtual meeting platforms (Zoom, Teams, etc.)"]
- [Included 3 - e.g., "Document preparation and delivery"]
- [Included 4 - e.g., "Reasonable local travel (within [region])"]

**Billed Separately (at cost + 10% admin fee):**
- [Billable 1 - e.g., "Air travel and accommodation for on-site work"]
- [Billable 2 - e.g., "Specialized software licenses (if required for project)"]
- [Billable 3 - e.g., "Third-party data subscriptions or research services"]
- [Billable 4 - e.g., "Printing and materials for large volumes (>100 pages)"]

*Client will be notified and approval obtained before incurring any billable expenses.*

### Pricing Transparency

**How We Calculated This Investment:**

[Insert pricing calculation metadata from Phase 2:]

- **Method:** [pricing_method - e.g., "Project pricing formula"]
- **Base Rate:** $[base_rate]/[unit - hour/day]
- [For project:] **Estimated Hours:** [hours] hours
- [For retainer:] **Monthly Hours:** [hours] hours ([tier] retainer)
- [If discount:] **Discount Applied:** [discount_pct]% ([reason - e.g., "project value >$50K"])
- **Calculation:** [Show calculation - e.g., "200 hours × $250/hr × 0.95 (5% discount) = $47,500"]

**Price Validation:**
- Service range for [service name]: $[range_min] - $[range_max]
- This proposal: $[calculated_price] ✓ [WITHIN RANGE | BELOW MINIMUM | ABOVE MAXIMUM]

[If warnings exist from validation:]
⚠️ **Pricing Notes:**
- [Warning or note from validation]

### Value Justification

[Extract from service pricing.value_justification:]

[Value justification text explaining ROI or business value]

**Expected ROI:**
- [ROI metric 1 from service success_metrics]
- [ROI metric 2 from service success_metrics]

---

[Evidence: Service_Definition_[Date].md → rate_card + [Service].pricing]
```

### 3.9 Generate "Terms & Conditions" Section

```markdown
## Terms & Conditions

### Scope of Work

**Included:**
- [Deliverable 1 from service deliverables]
- [Deliverable 2]
- [Deliverable 3]
- [All deliverables listed]

**Explicitly Not Included:**
- [Exclusion 1 - e.g., "Ongoing maintenance beyond [timeframe]"]
- [Exclusion 2 - e.g., "Implementation work (advisory only)"]
- [Exclusion 3 - e.g., "Training for end-users (management training only)"]
- [Exclusion 4 - e.g., "Third-party software licensing costs"]

*Any work beyond the defined scope will require a separate proposal or change order.*

### Assumptions

This proposal assumes:

1. **Client Access & Cooperation:**
   - [Assumption 1 - e.g., "Timely access to key stakeholders (within 48 hours)"]
   - [Assumption 2 - e.g., "Access to relevant systems and data as needed"]
   - [Assumption 3 - e.g., "Client SMEs available for [X] hours per week"]

2. **Timeline:**
   - [Assumption 4 - e.g., "Project kickoff within [X] days of contract signing"]
   - [Assumption 5 - e.g., "No delays due to client resource availability"]
   - [Assumption 6 - e.g., "Client feedback provided within [X] business days"]

3. **Environment & Tools:**
   - [Assumption 7 - e.g., "Client provides necessary system access/credentials"]
   - [Assumption 8 - e.g., "Existing infrastructure is documented and accessible"]

*If assumptions prove incorrect, timeline and pricing may be adjusted via change order process.*

### Change Management

**Scope Changes:**
- Changes to scope must be requested in writing
- We'll provide impact assessment (timeline, cost) within 3 business days
- Changes require mutual written approval via change order
- Approved changes will update timeline and investment accordingly

**Timeline Changes:**
- Client-caused delays do not extend our commitment or reduce fees
- We'll work to accommodate reasonable client-requested timeline shifts
- Major delays (>2 weeks) may require resource reallocation and rescheduling

### Intellectual Property

[Choose appropriate IP clause based on engagement model:]

[For PROJECT/WORKSHOP:]
**Work Product Ownership:**
- All deliverables become client property upon full payment
- Client receives perpetual, irrevocable license to use all work product
- We retain right to use general methodologies and approaches for other clients
- We may reference this project (anonymized) in case studies with client approval

**Pre-Existing IP:**
- Any pre-existing frameworks, tools, or methodologies remain our property
- Client receives non-exclusive license to use pre-existing IP for their internal purposes
- No resale or distribution of pre-existing IP without separate agreement

[For RETAINER/STAFF-AUG:]
**Work Product Ownership:**
- All work product created during engagement belongs to client
- Client owns all code, documentation, and deliverables upon payment
- We retain right to general knowledge and skills developed
- Non-compete or IP assignment beyond work product requires separate negotiation

**Confidentiality:**
- We agree to keep all client information confidential
- Client agrees to keep our methodologies and pricing confidential
- Mutual NDA available upon request (recommended for sensitive engagements)

### Confidentiality

**Information Protection:**
- All client data, systems, and business information treated as confidential
- We implement industry-standard security practices for data handling
- No disclosure to third parties without client written consent
- Confidentiality survives engagement termination

**Exceptions:**
- Information already publicly available
- Information independently developed without use of client confidential information
- Information required to be disclosed by law (with notice to client)

### Warranties & Limitations

**Our Warranties:**
- Services performed in professional, workmanlike manner
- Deliverables will meet specifications outlined in this proposal
- We have authority and ability to enter this engagement

**Limitation of Liability:**
- Our liability limited to fees paid for this engagement
- No liability for indirect, consequential, or punitive damages
- No liability for client's use or implementation of our recommendations

**Client Responsibilities:**
- Client responsible for implementation decisions and consequences
- Client responsible for ensuring recommendations comply with applicable laws/regulations
- We provide advisory services; client retains final decision authority

### Termination

**Termination for Convenience:**

[For PROJECT:]
- Either party may terminate with [X days notice - e.g., "14 days"] written notice
- Client pays for work completed plus non-cancellable costs
- We deliver all completed work product upon payment
- Prorated refund for unused prepaid fees

[For RETAINER:]
- Either party may terminate with [X days notice - e.g., "30 days"] written notice
- No refund for current month retainer (prepaid in advance)
- Unused future month retainers refunded in full
- Overage hours billed at end of termination notice period

[For STAFF-AUG:]
- Either party may terminate with [X days notice - e.g., "14 days"] written notice
- Payment for days worked through termination date
- Knowledge transfer period ([Y days - e.g., "5 days"]) included in notice period

**Termination for Cause:**
- Material breach of agreement with [X days - e.g., "7 days"] to cure
- Non-payment of fees (after [X days past due - e.g., "30 days"])
- Immediate termination for illegal activity or gross misconduct

### Engagement Period

**Proposal Validity:** This proposal is valid until [14 days from generation]

**Engagement Start:** Work begins [X days - e.g., "3-5 business days"] after contract signing and deposit payment

**Engagement Duration:**
[For PROJECT:] [Duration from typical_duration - e.g., "6-8 weeks from kickoff"]
[For RETAINER:] Ongoing, month-to-month with [X days] notice for termination
[For STAFF-AUG:] [Duration - e.g., "3-6 months initial term, renewable"]
[For WORKSHOP:] [Workshop dates + 90 days post-workshop support]

**Extension/Renewal:**
- [For project:] Extensions require new proposal or change order
- [For retainer:] Auto-renews monthly unless terminated
- [For staff-aug:] Renewable in [X-month increments] by mutual agreement

### Governing Law & Disputes

**Jurisdiction:** This agreement governed by laws of [Location - e.g., "Ontario, Canada"]

**Dispute Resolution:**
1. Good faith negotiation (30 days)
2. Mediation (if negotiation fails)
3. Binding arbitration or litigation (as last resort)

**Attorney Fees:** Prevailing party in dispute entitled to reasonable attorney fees

---

[Evidence: Service_Definition_[Date].md → payment_terms + engagement_models]
```

### 3.10 Generate "Next Steps" Section

```markdown
## Next Steps

### How to Proceed

To move forward with this engagement:

**Step 1: Review & Questions**
- Review this proposal in detail
- [If client-specific:] Verify our understanding of your situation matches your needs
- Prepare any questions or clarifications

**Step 2: Schedule Discussion (Optional)**
- [Contact method - e.g., "Reply to this proposal email or call [phone]"]
- 30-60 minute call to discuss proposal, answer questions, and refine if needed
- Available [availability - e.g., "within 48 hours"] for questions

**Step 3: Approval & Contract**
- Approve proposal by [deadline - 7-14 days from generation]
- Sign engagement agreement (provided upon approval)
- Submit deposit payment ([deposit_percentage]%: $[deposit_amount])

**Step 4: Kickoff**
- Schedule kickoff meeting within [X days - e.g., "3-5 business days"] of contract signing
- [For staff-aug:] Begin onboarding process
- [For workshop:] Finalize participant list and pre-workshop logistics

### Timeline

**If approved by [Date 7 days from now]:**

[For PROJECT:]
- **[Date 7 days + 3-5 days]**: Kickoff meeting
- **[Date + typical_duration]**: Project completion (estimated)

[For RETAINER:]
- **[Date 7 days + 3 days]**: Retainer begins, first advisory session scheduled
- **Ongoing**: Monthly retainer renews on [day of month]

[For STAFF-AUG:]
- **[Date 7 days + 3 days]**: Onboarding begins
- **[Date 7 days + 1 week]**: Full integration, productive contributions begin

[For WORKSHOP:]
- **[Date 2 weeks prior]**: Pre-workshop survey and participant confirmation
- **[Date 1 week prior]**: Pre-workshop call and final logistics
- **[Proposed workshop dates]**: Workshop delivery
- **[Proposed dates + 90 days]**: Post-workshop support period

**Total time from approval to start:** [X days - e.g., "3-5 business days"]

### What You Need to Prepare

Before kickoff, please:

**Administrative:**
- [Item 1 - e.g., "Designate primary point of contact and project sponsor"]
- [Item 2 - e.g., "Provide billing contact and PO (if required)"]
- [Item 3 - e.g., "Complete vendor onboarding (if applicable)"]

**Access & Resources:**
- [Item 1 - e.g., "Identify stakeholders for interviews/workshops"]
- [Item 2 - e.g., "Grant system access (list will be provided)"]
- [Item 3 - e.g., "Prepare workspace (if on-site work)"]

**Information:**
- [Item 1 - e.g., "Background documents (strategies, org charts, prior analyses)"]
- [Item 2 - e.g., "Context briefing for consultant (30-60 min call)"]

*We'll provide detailed onboarding checklist upon contract signing.*

### Questions or Customization Needed?

**Contact:**
- **Email:** [Email if available]
- **Phone:** [Phone if available]
- **LinkedIn:** [LinkedIn if available]
- **Response time:** [Commitment - e.g., "Within 24 hours"]

**Common Customizations:**
- Adjusting timeline or phases to match your constraints
- Adding or removing scope elements
- Modifying payment terms for your procurement process
- Creating phased approach (start smaller, expand later)

We're happy to refine this proposal to better fit your needs.

---

**Looking forward to working together.**

[Consultant Name]
[Credentials]

---

[Evidence: Service_Definition_[Date].md → payment_terms + engagement_models]
```

### 3.11 Generate "Appendix" Section

```markdown
## Appendix

### A. Detailed Credentials

[If candidate profile available:]

#### Professional Certifications

[Full list from certifications[] where status = "Active"]

| Certification | Issuing Organization | Year Obtained | Status |
|---------------|---------------------|---------------|--------|
| [Cert 1] | [Organization] | [Year] | Active |
| [Cert 2] | [Organization] | [Year] | Active |
| [Cert 3] | [Organization] | [Year] | Active |

[Evidence: candidate_profile.json → certifications[]]

#### Education

[Extract from education[]]

- **[Degree 1]** - [Institution] ([Year])
- **[Degree 2]** - [Institution] ([Year])

[Evidence: candidate_profile.json → education[]]

#### Thought Leadership

**Publications:**
[Extract from thought_leadership.publications[]]
- [Publication 1] - [Type: Peer-reviewed/Whitepaper/etc.] ([Year])
- [Publication 2] - [Type] ([Year])

**Speaking Engagements:**
[Extract from thought_leadership.speaking_engagements]
- [Speaking engagement 1] - [Venue] ([Year])
- [Speaking engagement 2] - [Venue] ([Year])

[Evidence: candidate_profile.json → thought_leadership]

#### Technical Proficiencies

**Expert Level:**
[List all technical_skills[] where proficiency_level = "Expert" and relevant to service]

**Proficient Level:**
[List all technical_skills[] where proficiency_level = "Proficient" and relevant to service, limit to top 5-7]

[Evidence: candidate_profile.json → technical_skills[]]

[If candidate profile NOT available, use service definition:]

#### Credentials & Differentiators

[Extract from differentiation section:]

**Certifications:**
- [Certification 1 from differentiation.certifications]
- [Certification 2]

**Unique Capabilities:**
- [Competitive advantage 1 from differentiation.competitive_advantages]
- [Competitive advantage 2]

**Proven Results:**
- [Proof point 1 from differentiation.proof_points]
- [Proof point 2]

[Evidence: Service_Definition_[Date].md → differentiation]

---

### B. Case Studies (If Available)

[If service definition has case_studies[] for this service:]

#### Case Study 1: [Client type from case_study]

**Challenge:**
[Challenge description from case_study]

**Solution:**
[Solution description from case_study]

**Results:**
[Results from case_study]

[Evidence: Service_Definition_[Date].md → [Service].case_studies[0]]

---

[Continue for additional case studies]

[If no case studies in service definition:]

**Representative Project Examples:**

[Extract from candidate profile projects[] or work_history[] relevant to service]

#### Project: [Project name or description]

**Context:** [Brief project context]

**Deliverables:** [What was delivered]

**Outcome:** [Quantified result if available]

[Evidence: candidate_profile.json → projects[X] OR work_history[Y]]

---

### C. References

**Professional References Available Upon Request**

For qualified opportunities, we provide:
- Client references from similar engagements
- Peer references from industry colleagues
- [If applicable:] Academic/institutional references

**To request references:**
Email [email] with:
- Your organization and role
- Specific service/engagement under consideration
- Timeline for decision

---

### D. Terms & Definitions

**Key Terms Used in This Proposal:**

- **Deliverable**: Specific work product or output provided to client
- **Phase Gate**: Approval point between project phases requiring client sign-off
- **Change Order**: Formal process for modifying scope, timeline, or budget
- **Overage Hours**: [For retainer] Hours beyond monthly allocation, billed separately
- **Work Product**: All materials, documents, code, analyses created during engagement
- **Confidential Information**: Any non-public client data, business information, or systems

**Abbreviations:**

- **SME**: Subject Matter Expert
- **SOW**: Statement of Work (this proposal serves as SOW)
- **PO**: Purchase Order
- **NDA**: Non-Disclosure Agreement

---

### E. Proposal Metadata

**Generated:** [ISO8601 timestamp]
**Proposal Reference:** PROP-[YYYYMMDD]-[ClientInitials]-[ServiceInitials]
**Valid Until:** [14 days from generation]
**Version:** 1.0

**Source Documents:**
- Service Definition: [File path]
- [If pitch deck used:] Prospect Research: [File path]
- [If profile used:] Candidate Profile: [File path]

**Generated By:** /proposaltemplate command

---

**End of Proposal**

```

---

## Phase 4: Save Output File

### 4.1 Determine Output Filename

**Filename format:** `Proposal_[Client]_[Service]_[Date].md`

**Client portion:**
- Use `--client` argument value if provided
- If client is "Generic", use "Generic"
- Sanitize: Remove spaces, special characters, use PascalCase
- Example: "Oxford Properties" → "OxfordProperties"

**Service portion:**
- Use abbreviated service name (first 2-3 significant words)
- Sanitize: Remove spaces, special characters, use PascalCase
- Maximum 30 characters
- Example: "AI Agent Development for Due Diligence" → "AIAgentDevelopment"

**Date portion:**
- Format: YYYYMMDD
- Example: "20251203"

**Complete examples:**
- `Proposal_Hatch_AIAgentDevelopment_20251203.md`
- `Proposal_Generic_PortfolioAdvisory_20251203.md`
- `Proposal_OxfordProperties_RealEstateAI_20251203.md`

### 4.2 Create Output Directory

```bash
# Ensure Client_Prospects directory exists
mkdir -p Client_Prospects/
```

### 4.3 Save Proposal File

Write complete proposal markdown to:
`Client_Prospects/Proposal_[Client]_[Service]_[Date].md`

**File encoding:** UTF-8
**Line endings:** Unix (LF)

Include complete proposal with all sections generated in Phase 3.

---

## Phase 5: Generate Summary Report

### 5.1 Console Output Summary

```
✅ PROFESSIONAL PROPOSAL GENERATED

**Output File:**
Client_Prospects/Proposal_[Client]_[Service]_[YYYYMMDD].md

**Proposal Details:**
- **Client:** [Client name or "Generic"]
- **Service:** [Service name]
- **Engagement Type:** [project|retainer|staff-aug|workshop]
- **Total Investment:** $[calculated_price]
- [If retainer:] **Monthly Investment:** $[monthly_fee]
- **Duration:** [Duration from typical_duration or calculation]

**Pricing Calculation:**
- **Method:** [pricing_method from calculation]
- **Base Rate:** $[base_rate]/[hour|day]
- [Calculation-specific details]
- **Discount Applied:** [discount_pct]% ([reason]) OR "None"
- **Final Price:** $[calculated_price]

**Price Validation:**
- Service range: $[range_min] - $[range_max]
- Calculated price: $[calculated_price] ✓ [WITHIN RANGE | BELOW | ABOVE]
- Status: [VALIDATED | WARNING | ERROR]

[If validation warnings:]
⚠️ **Pricing Advisories:**
- [Warning 1]
- [Warning 2]

**Proposal Contents:**
- **Sections:** 10 (Cover through Appendix)
- **Deliverables Listed:** [Count] deliverables
- **Timeline:** [Summary - e.g., "6-week project with 4 phases" or "Ongoing monthly retainer"]
- [If prospect research used:] **Customization:** Client-specific (used pitch deck research)
- [If generic:] **Customization:** Generic (can be tailored with prospect research)

**Prospect Research Integration:**
[If pitch deck found and used:]
✓ Proposal customized with prospect-specific pain points from pitch deck
  - Research date: [date]
  - Pain points: [count]
  - Decision-makers: [count]

[If pitch deck NOT found:]
ℹ️ Generic service positioning used (no pitch deck found for "[Client]")
  - To create more tailored proposal: /pitchdeck --prospect="[Client]"

**Sources:**
- Service Definition: Client_Prospects/Service_Definition_[Date].md
[If pitch deck:] - Prospect Research: Client_Prospects/Pitch_[Client]_[Date].md
[If profile:] - Candidate Profile: ResumeSourceFolder/.profile/candidate_profile.json

---

**Next Steps:**

1. **Review Proposal**
   - Open: Client_Prospects/Proposal_[Client]_[Service]_[Date].md
   - Review pricing calculation in YAML frontmatter
   - Check "Understanding Your Challenge" section for accuracy
   - Verify deliverables match client expectations
   - Review timeline and phases for feasibility

2. **Customize (Optional)**
   - Add client logo to cover page (if converting to PDF)
   - Personalize "Understanding Your Challenge" with specific client context
   - Adjust phases/timeline if needed
   - Add specific names to "Next Steps" section
   - Include any client-specific T&C requirements

3. **Convert to Professional Format (Optional)**
   - Markdown to PDF: /formatresume Proposal_[Client]_[Service]_[Date].md
   - Markdown to DOCX: /convert Proposal_[Client]_[Service]_[Date].md
   - Manual: Copy to Word/Google Docs and apply formatting

4. **Send to Client**
   - Remove YAML frontmatter and evidence citations if desired
   - Keep internal version with full provenance trail
   - Email as PDF attachment or share via DocuSign/PandaDoc
   - Follow up within 3-5 business days

5. **If Client Requests Changes**
   - Note requested changes
   - Update pricing if scope changes
   - Regenerate: /proposaltemplate --client="[Client]" --service="[Service]" [adjusted args]
   - Version control: Keep both versions for reference

---

**Professional Tips:**

💡 **Proposal Best Practices:**
- Respond quickly: Send within 24-48 hours of request
- Personalize: Use client name, not "your organization"
- Be specific: Concrete deliverables > vague promises
- Show expertise: Include relevant case studies and credentials
- Price transparently: Explain how pricing was calculated

📊 **Using Your Proposal:**
- Lead with Executive Summary: Busy decision-makers read this first
- Emphasize ROI: Connect price to business value
- Set clear expectations: Timeline, deliverables, client responsibilities
- Reduce risk perception: Case studies, credentials, terms & conditions
- Make it easy to say yes: Clear next steps and tight turnaround

🎯 **When to Customize Further:**
- Enterprise clients: Add compliance sections (SOC2, GDPR, etc.)
- Government/public sector: Include procurement-specific language
- Startups: Emphasize agility, consider equity options
- Strategic partners: Revenue-share or outcome-based pricing models
- Multi-phase opportunities: Create phased pricing (pilot → full implementation)

📄 **Document Management:**
- Internal version: Keep with YAML frontmatter and evidence citations
- Client version: Clean markdown or PDF without metadata
- Versioning: If updated, increment version and track changes
- Archive: Keep all versions for reference and future proposals

---

**Proposal Ready for Review and Delivery**
```

### 5.2 Warning Messages (If Applicable)

**If price is below minimum viable engagement:**
```
⚠️ BELOW MINIMUM VIABLE ENGAGEMENT

Calculated price: $[calculated_price]
Minimum viable [engagement_type] engagement: $[minimum]

**Recommendation:**
Consider one of the following before sending to client:
1. Increase scope to meet minimum viable engagement
2. Bundle with additional services
3. Use as loss-leader for strategic client (if justified)
4. Decline engagement as not economically viable

Proposal generated with calculated price, flagged for your review.
```

**If price is outside service range:**
```
⚠️ PRICE OUTSIDE SERVICE RANGE

Service range: $[range_min] - $[range_max]
Calculated price: $[calculated_price]

**Analysis:**
[Explanation of why price is outside range - e.g., "Smaller scope than typical engagement" or "Additional complexity requiring premium pricing"]

**Recommendation:**
[Specific recommendation - e.g., "Consider adding Phase X to reach minimum range" or "Justify premium pricing with additional value"]

Proposal generated with calculated price, flagged for your review.
```

**If prospect research is stale (>30 days):**
```
⚠️ STALE PROSPECT RESEARCH

Pitch deck research date: [date] ([X] days old)

**Recommendation:**
Before sending proposal, verify:
- Decision-makers are still in role (check LinkedIn)
- Strategic priorities haven't changed (check recent news)
- Pain points are still relevant (check job postings, reviews)

Or regenerate pitch deck for fresh research:
  /pitchdeck --prospect="[Client]"

Proposal uses existing research, but manual validation recommended.
```

---

## Error Handling

### Error 1: No Service Definition Found

```
❌ PREREQUISITE MISSING: Service Definition

No service definition file found in Client_Prospects/

**Required Action:**
Run /defineservices first to create your service catalog and pricing structure.

Usage:
  /defineservices              # Interactive guided mode
  /defineservices --from-profile  # Auto-generate from your profile

After creating your service definition, run /proposaltemplate again.
```

**Stop execution.**

### Error 2: Service Not Found

```
❌ SERVICE NOT FOUND: [service-name]

No service matching "[service-name]" in your service definition.

Available services from your service definition:
1. [Service 1 name]
2. [Service 2 name]
3. [Service 3 name]

Usage: /proposaltemplate --service="[exact or partial service name]"

Example: /proposaltemplate --service="AI Agent Development"
```

**Stop execution.**

### Error 3: Invalid Engagement Type

```
❌ INVALID ENGAGEMENT TYPE: [type]

Valid engagement types:
- project: Fixed-scope project engagement
- retainer: Ongoing advisory/support retainer
- staff-aug: Staff augmentation/embedded consulting
- workshop: Training/facilitation workshop

Usage: /proposaltemplate --type=project

Defaulting to "project" type.
```

**Continue execution with default.**

### Error 4: Incomplete Service Definition

```
❌ INCOMPLETE SERVICE DEFINITION

Your service definition at Client_Prospects/Service_Definition_[Date].md
is missing required sections:

Missing sections:
- [Missing section 1 - e.g., "rate_card"]
- [Missing section 2 - e.g., "payment_terms"]

Please update your service definition:
  /defineservices --update

Then run /proposaltemplate again.
```

**Stop execution.**

### Error 5: Pricing Calculation Failed

```
❌ PRICING CALCULATION ERROR

Unable to calculate pricing for [engagement_type] engagement.

**Issue:** [Specific issue - e.g., "No typical_duration specified in service definition and no --value argument provided"]

**Resolution:**
1. Provide --value argument: /proposaltemplate --value=25000
2. Or update service definition with typical_duration: /defineservices --update

Cannot generate proposal without pricing information.
```

**Stop execution.**

---

## Usage Examples

### Example 1: Generic Project Proposal (Default)

```bash
/proposaltemplate
```

**Behavior:**
- Uses first/primary service from service definition
- Client: "Generic"
- Type: "project"
- Calculates pricing from service definition and rate card
- No prospect-specific customization
- Output: `Client_Prospects/Proposal_Generic_[ServiceName]_20251203.md`

**Use Case:** Template for general inquiries, website download

---

### Example 2: Client-Specific AI Development Proposal

```bash
/proposaltemplate --client="Hatch" --service="AI Agent Development"
```

**Behavior:**
- Searches for pitch deck: `Client_Prospects/Pitch_*Hatch*.md`
- If found: Customizes "Understanding Your Challenge" with prospect research
- Uses "AI Agent Development" service
- Type: "project" (default)
- Calculates pricing from service definition
- Output: `Client_Prospects/Proposal_Hatch_AIAgentDevelopment_20251203.md`

**Use Case:** Following up on pitch deck with formal proposal

---

### Example 3: Retainer Proposal

```bash
/proposaltemplate --client="OxfordProperties" --service="Portfolio Advisory" --type=retainer
```

**Behavior:**
- Client: "OxfordProperties"
- Service: "Portfolio Advisory"
- Type: "retainer"
- Calculates monthly retainer fee with discount
- Shows ongoing deliverables, monthly cadence
- Output: `Client_Prospects/Proposal_OxfordProperties_PortfolioAdvisory_20251203.md`

**Use Case:** Ongoing advisory relationship proposal

---

### Example 4: Workshop with Target Value

```bash
/proposaltemplate --service="AI Training" --type=workshop --value=15000
```

**Behavior:**
- Client: "Generic"
- Service: "AI Training"
- Type: "workshop"
- Uses --value=15000 to back-calculate workshop days
- Calculates prep days, materials, follow-up support
- Output: `Client_Prospects/Proposal_Generic_AITraining_20251203.md`

**Use Case:** RFP response with specified budget

---

### Example 5: Staff Augmentation Proposal

```bash
/proposaltemplate --client="TechStartup" --service="Senior AI Engineer" --type=staff-aug
```

**Behavior:**
- Client: "TechStartup"
- Service: "Senior AI Engineer"
- Type: "staff-aug"
- Calculates monthly fee based on daily rate
- Shows onboarding, responsibilities, working arrangement
- Output: `Client_Prospects/Proposal_TechStartup_SeniorAIEngineer_20251203.md`

**Use Case:** Embedded consulting opportunity

---

Now executing proposal generation...
