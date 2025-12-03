---
description: Identify potential B2B clients matching service definition and ideal client profiles
argument-hint: [ideal-job-file] [--industry=X] [--size=startup|mid|enterprise] [--location=X] [--limit=N]
---

I'll help you identify potential B2B clients for your consulting/contractor services based on your service definition and ideal client profiles. This command performs intelligent client discovery with B2B fit scoring and entry point mapping.

**Input Options:**
- `<ideal-job-file>` (optional): Path to ideal job description file from `/idealjob` output
- `--industry=X`: Filter by specific industry (e.g., --industry="Real Estate")
- `--size=startup|mid|enterprise`: Filter by company size
- `--location=X`: Geographic filter (e.g., --location="Toronto")
- `--limit=N`: Maximum number of prospects to research (default: 20, max: 50)

**Output File**: `Client_Prospects/Prospects_[Domain]_[Date].md`

---

## Phase 1: Load Service Definition and Validate Prerequisites

### 1.1 Check for Service Definition

**CRITICAL PREREQUISITE**: This command requires an existing service definition created by `/defineservices`.

Search for service definition files:
```
Look for: Client_Prospects/Service_Definition_*.md
```

**If NOT found:**
```
❌ SERVICE DEFINITION REQUIRED

Before identifying clients, you must first define your service offerings.

Run this command first:
  /defineservices

This will create a comprehensive service catalog with:
- Service offerings and deliverables
- Ideal client profiles for each service
- Pricing structure and engagement models
- Competitive differentiation and proof points

Once you have a service definition, run /findclient again.
```

Stop execution and guide user to run `/defineservices` first.

**If multiple service definitions found:**
- Use the most recent file (latest date in filename)
- Note the version being used in console output

### 1.2 Load and Parse Service Definition

Read the service definition file and extract:

**From YAML frontmatter:**
- `consultant` name
- `generated_on` timestamp
- `version` number

**From markdown body - Services section:**
For each service offering, extract:
- `name`: Service name (e.g., "AI Agent Development for Due Diligence")
- `category`: Service category
- `ideal_client.industries[]`: Target industries
- `ideal_client.company_size[]`: Preferred company sizes
- `ideal_client.pain_points[]`: Client problems solved
- `ideal_client.decision_makers[]`: Typical buyer roles

**From markdown body - Target Market section (if exists):**
- `primary_markets[]`: Primary target industries with segments
- `secondary_markets[]`: Secondary opportunities
- `avoid_markets[]`: Markets to exclude
- `geographic_focus[]`: Geographic regions

**From markdown body - Differentiation section:**
- `competitive_advantages[]`: Key differentiators
- `proof_points[]`: Quantified achievements

Store this data for use in search query generation and prospect scoring.

### 1.3 Load Optional Ideal Job File

If `{{ARG1}}` is provided and looks like a file path (ends with `.md`):
- Read the ideal job file
- Extract additional targeting criteria:
  - Target industries mentioned
  - Company size indicators
  - Geographic location
  - Technology stack requirements
  - Organizational characteristics

Use this to supplement service definition targeting.

### 1.4 Parse Command-Line Filters

Extract filters from arguments:

**Industry filter (`--industry=X`):**
- If present, narrow search to specified industry only
- Validate against service definition industries
- Warn if industry not in service definition but proceed

**Size filter (`--size=startup|mid|enterprise`):**
- Map to standard company size categories:
  - `startup`: 1-50 employees
  - `mid`: 50-500 employees, or "Mid-Market"
  - `enterprise`: 500+ employees, or "Enterprise"
- Filter prospect list to match size criteria

**Location filter (`--location=X`):**
- Extract geographic constraint
- Use for both search targeting and prospect filtering

**Limit filter (`--limit=N`):**
- Default: 20 prospects
- Maximum recommended: 50 (research time constraints)
- Parse integer value from argument

### 1.5 Consolidate Targeting Criteria

Merge all sources of targeting data:

**Priority order:**
1. Command-line filters (highest priority - explicit user intent)
2. Ideal job file criteria (medium priority - specific opportunity)
3. Service definition (baseline - proven ideal clients)

**Create targeting profile:**
```
Target Profile:
- Industries: [Consolidated list from all sources]
- Company Sizes: [Filtered or all from service definition]
- Geographic Focus: [Location filter or service definition regions]
- Pain Points: [From service definition]
- Decision Makers: [From service definition]
```

---

## Phase 2: Generate Research Queries and Search Strategy

### 2.1 Identify Service-Industry Intersections

For each service in the service definition:
- Cross-reference with target industries
- Identify specific pain points per industry
- Generate industry-specific value propositions

Example:
```
Service: "AI Agent Development for Due Diligence"
Industries: ["Commercial Real Estate", "Private Equity", "Infrastructure"]

Intersections:
- CRE × AI Due Diligence → "Property acquisition automation"
- PE × AI Due Diligence → "Investment analysis acceleration"
- Infrastructure × AI Due Diligence → "Asset evaluation efficiency"
```

### 2.2 Generate Web Search Queries (10-15 queries)

Create targeted search queries for company discovery:

**Query Type 1: Industry + Contractor Signals**
```
"[industry] companies hiring consultants [year]"
"[industry] firms contractor jobs [location]"
"[industry] consultant opportunities [location]"
```

**Query Type 2: Pain Point + Service Match**
```
"[industry] companies struggling with [pain point]"
"[industry] [pain point] consulting RFP"
"[industry] need for [service capability]"
```

**Query Type 3: Procurement and Vendor Discovery**
```
"[industry] vendor registration portal [location]"
"[industry] companies procurement process"
"[industry] consulting spend [location]"
```

**Query Type 4: Specific Company Discovery**
```
"top [industry] companies [location]"
"[industry] companies [company size] [location]"
"[industry] digital transformation projects [year]"
```

**Query Type 5: Intermediary Channels**
```
"[industry] consulting marketplaces"
"[industry] staffing agencies [location]"
"[domain] consultant placement firms"
```

**Query Type 6: Industry Association and Event Signals**
```
"[industry] association member directory [location]"
"[industry] conference sponsors [year]"
"[industry] trade show exhibitors"
```

**Query Customization:**
- Replace `[year]` with current year (2025)
- Replace `[location]` with geographic filter if provided
- Replace `[industry]` with each target industry
- Replace `[pain point]` with pain points from service definition

Generate 10-15 high-value queries based on targeting profile.

### 2.3 Define Research Time Budget

**Time allocation:**
- Company discovery: ~3-4 minutes (broad search)
- B2B signals research: ~2-3 minutes (contractor evidence)
- Entry point identification: ~1-2 minutes (procurement/contacts)
- Total: ~6-8 minutes for standard 20-prospect run

**Search strategy:**
- Execute 3-5 searches simultaneously for efficiency
- Prioritize queries by expected signal strength
- Stop when sufficient prospects identified (limit reached)

---

## Phase 3: Company Discovery and B2B Signal Research

### 3.1 Execute Company Discovery Searches

Run web searches using generated queries to identify potential client companies.

**For each search result, extract:**
- Company name
- Industry/sector
- Location (headquarters)
- Company size (employees, revenue if available)
- Website URL
- Brief description

**Discovery sources to prioritize:**
- LinkedIn company pages (verified business profiles)
- Industry association directories (vetted members)
- Conference sponsor/exhibitor lists (budget signals)
- Government vendor registries (procurement-ready)
- News articles about digital transformation
- Consulting marketplace listings

**Quality filters:**
- Exclude staffing agencies (intermediaries, not clients)
- Exclude individual consultants (competitors, not clients)
- Exclude very small companies (<5 employees) unless startup filter used
- Prioritize companies with online presence

Target: Identify 30-50 candidate companies (will score and filter to limit)

### 3.2 Research B2B and Contractor Signals

For each candidate company discovered, conduct focused research on B2B indicators:

**Signal Type 1: Contractor History**
Search for:
```
"[company name] consultant job postings"
"[company name] contractor positions"
"[company name] freelance opportunities"
site:linkedin.com "[company name]" consultant
site:glassdoor.com "[company name]" reviews contractor
```

**Evidence:**
- Job postings for contractors/consultants (recent = strong signal)
- Glassdoor reviews mentioning contractors positively
- LinkedIn employees with "Contractor" or "Consultant" titles
- Career page sections for vendors/contractors

**Signal Type 2: Procurement Accessibility**
Search for:
```
"[company name]" vendor registration
"[company name]" supplier portal
"[company name]" procurement process
"[company name]" small business program
```

**Evidence:**
- Vendor registration portal exists and is accessible
- Small business or diverse supplier programs
- Streamlined procurement process (not overly bureaucratic)
- Recent RFPs or consulting engagements announced

**Signal Type 3: Domain Alignment**
Cross-reference company with service definition:
- Industry match with target industries
- Pain points evident from company news/challenges
- Technology stack alignment with service capabilities
- Organizational initiatives matching service offerings

**Evidence:**
- Press releases about relevant initiatives
- Job postings showing technology needs
- News articles about challenges service could address
- Company website describing relevant pain points

**Signal Type 4: Size/Budget Fit**
Research financial capacity:
```
"[company name]" revenue
"[company name]" funding
"[company name]" employee count
site:linkedin.com "[company name]" employees
```

**Evidence:**
- Revenue indicators (public filings, news, estimates)
- Funding rounds (for startups)
- Employee count (LinkedIn, company website)
- Consulting spend disclosures (annual reports)

**Budget assessment:**
- Can they afford target rates? (Compare to service definition rate card)
- Enterprise (500+ employees): Likely budget for premium consulting
- Mid-market (50-500): Moderate budgets, selective consulting
- Startup (<50): Limited budgets, must show clear ROI

**Signal Type 5: Geographic Match**
Validate location compatibility:
- Headquarters location vs. geographic focus
- Remote-friendly culture (if serving outside local area)
- Multiple office locations (indicates scale)
- International presence (if relevant)

**Time management:**
- Spend ~30-60 seconds per company on signal research
- Use parallel searches when possible
- Focus on high-signal sources (LinkedIn, company website, Glassdoor)

### 3.3 Identify Entry Points and Contacts

For promising companies (preliminary scoring), research access paths:

**Entry Point Type 1: Procurement Portals**
```
Direct vendor registration available?
URL: [vendor portal URL]
Process: [Open/Quarterly/By invitation]
Notes: [Registration requirements, fees, timelines]
```

**Entry Point Type 2: Warm Introductions**
```
Search LinkedIn for:
- 2nd-degree connections at target company
- Shared professional associations (e.g., RICS, PMI)
- Alumni networks from same schools
- Former colleagues now at target company

Document:
- Contact name and title
- Relationship path (e.g., "2nd degree via [mutual contact]")
- Shared affiliation
```

**Entry Point Type 3: Intermediary Channels**
```
Staffing agencies serving this client:
- [Agency name] - [Specialization]
- Relationship: [Existing vendor? New contact?]

Consulting marketplaces:
- [Platform name] - [Whether company sources there]
```

**Entry Point Type 4: Decision-Maker Identification**
```
Search LinkedIn for target decision-makers:
- [Title from ideal_client.decision_makers]
- Current employees matching role
- Accessibility (LinkedIn profile, contact info)

Document:
- Name: [Decision maker name]
- Title: [Exact title]
- LinkedIn: [Profile URL]
- Accessibility: [Open to connections? Premium member?]
```

**Entry Point Type 5: Content/Thought Leadership**
```
Engagement opportunities:
- Industry conferences company attends
- Company blog accepting guest posts
- LinkedIn activity (comment on posts)
- Webinars or events company hosts
```

Time budget: ~1-2 minutes per high-priority prospect (score 6+)

---

## Phase 4: B2B Fit Scoring Methodology

### 4.1 Apply 10-Point Scoring Framework

For each candidate company, calculate a B2B fit score using 5 weighted factors:

**Scoring Formula:**
```
Total Score = (Contractor History × 0.25) +
              (Procurement Accessibility × 0.20) +
              (Domain Alignment × 0.25) +
              (Size/Budget Fit × 0.15) +
              (Geographic Match × 0.15)
```

Each factor scored 0-10 points, weighted sum produces final 0-10 score.

### 4.2 Factor 1: Contractor History (25% weight, max 10 points)

**Scoring criteria:**

| Score | Evidence Level |
|-------|----------------|
| 10 | Active contractor program + recent job posts (< 3 months) + positive Glassdoor reviews mentioning consultants |
| 8-9 | Recent contractor job posts (< 6 months) OR consulting spend disclosed in annual reports |
| 6-7 | Some evidence of contractor use (LinkedIn profiles, older job posts 6-12 months) |
| 4-5 | Indirect signals (staffing agency mentions, industry norms suggest consultant use) |
| 2-3 | Minimal evidence, but company size suggests some consulting (enterprise) |
| 0-1 | No evidence of contractor/consultant engagement history |

**Evidence to cite:**
- URL of job posting with date
- Glassdoor review quote mentioning contractors
- LinkedIn profile of current/former contractor
- Annual report page number with consulting spend

**Confidence level:**
- HIGH: Direct evidence (job post, vendor portal, annual report)
- MEDIUM: Indirect evidence (employee reviews, LinkedIn profiles)
- LOW: Inferred from industry norms or company size

### 4.3 Factor 2: Procurement Accessibility (20% weight, max 10 points)

**Scoring criteria:**

| Score | Evidence Level |
|-------|----------------|
| 10 | Public vendor portal + small business program + streamlined process (documented) |
| 8-9 | Vendor portal exists + reasonable registration process |
| 6-7 | Procurement contact identified + process documented publicly |
| 4-5 | General procurement info available, suggests openness to new vendors |
| 2-3 | No public procurement info, but company size suggests process exists |
| 0-1 | Closed vendor network OR overly bureaucratic (e.g., strict pre-qualification) |

**Evidence to cite:**
- Vendor portal URL
- Procurement policy document
- Contact email/phone for vendor inquiries
- Small business program description

**Accessibility indicators:**
- Open registration (no pre-qualification)
- Small business friendly (lower thresholds)
- Quarterly or rolling vendor acceptance
- Clear SOW/contracting process

**Red flags:**
- "Approved vendors only" language
- Excessive insurance/bonding requirements
- Long qualification timelines (>6 months)
- Exclusively enterprise vendors

### 4.4 Factor 3: Domain Alignment (25% weight, max 10 points)

**Scoring criteria:**

| Score | Evidence Level |
|-------|----------------|
| 10 | Perfect industry match + active pain point evident + recent initiative matching service |
| 8-9 | Target industry + strong pain point signals OR adjacent industry with perfect pain point match |
| 6-7 | Industry match + some pain point evidence OR adjacent industry + good fit |
| 4-5 | Adjacent industry + transferable pain points |
| 2-3 | Weak industry alignment, speculative pain point fit |
| 0-1 | Industry mismatch OR no evidence of relevant pain points |

**Domain alignment checks:**

**Industry Match:**
- Exact match with service definition target industries: +4 points
- Adjacent industry (transferable expertise): +2 points
- Unrelated industry: 0 points

**Pain Point Evidence:**
- Press releases describing exact pain point service addresses: +3 points
- Job postings showing technology/skill gaps: +2 points
- Industry reports suggesting common challenges: +1 point
- No pain point evidence: 0 points

**Service Need Signals:**
- Recent initiative directly matching service offering: +3 points
- Strategic priority aligned with service value: +2 points
- General interest area: +1 point
- No relevant initiatives: 0 points

**Evidence to cite:**
- Press release URL with initiative description
- Job posting showing skill gap
- Company website strategic priorities section
- Industry report mentioning company challenges

### 4.5 Factor 4: Size/Budget Fit (15% weight, max 10 points)

**Scoring criteria:**

| Score | Evidence Level |
|-------|----------------|
| 10 | Enterprise (1000+ employees) + disclosed consulting spend > $1M/yr |
| 8-9 | Enterprise (500-1000 employees) OR mid-market with strong revenue ($50M+) |
| 6-7 | Mid-market (100-500 employees) with adequate revenue ($10M-$50M) |
| 4-5 | Small business (50-100 employees) with funding/revenue indicators |
| 2-3 | Startup (<50 employees) with Series A+ funding |
| 0-1 | Very small (<20 employees) with no budget indicators |

**Budget assessment formula:**

Compare service definition rate card to estimated project budgets:

**Rate card reference:**
- Minimum hourly: $[X]/hr → Minimum project ~$[X × 40] (40 hours)
- Target hourly: $[Y]/hr → Typical project ~$[Y × 80-160] (2-4 weeks)
- Premium hourly: $[Z]/hr → Large project ~$[Z × 200+]

**Company budget capacity (estimate):**
- Enterprise (500+ employees): $50K-$500K+ projects feasible
- Mid-market (100-500 employees): $10K-$100K projects typical
- Small business (50-100 employees): $5K-$25K projects
- Startup (<50 employees): <$10K unless well-funded

**Scoring logic:**
- Can afford premium rate + large projects (160+ hours): 10 points
- Can afford target rate + medium projects (80-160 hours): 8 points
- Can afford minimum rate + small projects (40-80 hours): 6 points
- Stretch for minimum rate + small projects: 4 points
- Likely below minimum rate tolerance: 2 points
- Almost certainly cannot afford services: 0 points

**Evidence to cite:**
- Revenue figures (public filings, Crunchbase, news)
- Employee count (LinkedIn, company website)
- Funding rounds and amounts (Crunchbase, TechCrunch)
- Consulting spend disclosures (annual reports)

**Size verification:**
- LinkedIn company page follower count (rough proxy)
- "About" page employee count claims
- Office locations and scale indicators

### 4.6 Factor 5: Geographic Match (15% weight, max 10 points)

**Scoring criteria:**

| Score | Evidence Level |
|-------|----------------|
| 10 | HQ in target geography + remote-friendly culture confirmed |
| 8-9 | HQ in target geography OR remote-friendly + adjacent geography |
| 6-7 | National presence including target geography + remote option |
| 4-5 | Different geography but strong remote culture |
| 2-3 | Different geography, limited remote signals |
| 0-1 | Different geography + no remote culture (on-site only) |

**Geographic considerations:**

**If location filter provided:**
- HQ in specified location: +5 points
- Office/presence in specified location: +3 points
- Adjacent region (e.g., Toronto when "Ontario" specified): +2 points
- National/international with remote work: +2 points
- No presence in location + no remote: 0 points

**If service definition has geographic_focus:**
- Match with primary region: +3 points
- Match with secondary regions: +2 points
- International presence (if global focus): +2 points

**Remote work culture assessment:**
- Job postings explicitly remote-friendly: +3 points
- Glassdoor reviews mentioning remote flexibility: +2 points
- Career page remote work policy: +2 points
- COVID-era remote adoption signals: +1 point
- No remote signals: 0 points

**Distance considerations:**
- Same city/metro: Premium for in-person engagements
- Same province/state: Reasonable for hybrid
- Different region but remote-friendly: No penalty
- International with no remote culture: Potential barrier

**Evidence to cite:**
- Company careers page remote work policy
- Job postings with "Remote" or "Hybrid" tags
- Glassdoor review quotes about flexibility
- LinkedIn employee locations (distributed = remote-friendly)

### 4.7 Calculate Final B2B Fit Score

**Weighted calculation:**
```
Score = (Contractor_History × 0.25) +
        (Procurement_Access × 0.20) +
        (Domain_Alignment × 0.25) +
        (Size_Budget × 0.15) +
        (Geographic × 0.15)

Result: 0.0 to 10.0 (round to 1 decimal place)
```

**Priority classification:**

| Score Range | Priority Level | Recommendation |
|-------------|----------------|----------------|
| 8.0 - 10.0 | HIGH | Active outreach strongly recommended |
| 5.0 - 7.9 | MEDIUM | Worth pursuing with right introduction or timing |
| 1.0 - 4.9 | LOW | Long-shot, deprioritize unless specific opportunity arises |

**Confidence level for score:**
- HIGH: 4+ factors scored with direct evidence
- MEDIUM: 2-3 factors with direct evidence, rest inferred
- LOW: Mostly inferred scoring, limited direct evidence

---

## Phase 5: Generate Prospect Report and Outreach Strategy

### 5.1 Sort and Filter Prospects

**Sorting:**
1. Sort all scored prospects by B2B fit score (descending)
2. Group by priority level (HIGH 8-10, MEDIUM 5-7.9, LOW 1-4.9)

**Filtering:**
- Limit total prospects to `--limit` value (default 20)
- Within limit, maximize HIGH priority prospects
- Include representative MEDIUM prospects for pipeline
- Include 1-2 LOW prospects for context (what to avoid)

**Distribution target (for limit=20):**
- HIGH priority: 8-12 prospects (top opportunities)
- MEDIUM priority: 6-10 prospects (qualified pipeline)
- LOW priority: 0-2 prospects (learning examples)

### 5.2 Generate Prospect Report Structure

Create comprehensive report at: `Client_Prospects/Prospects_[Domain]_[Date].md`

**Domain naming:**
- Use primary service category or industry focus
- Examples: "AI_Consulting", "RealEstate_Advisory", "DataAnalytics"
- Sanitize for filename (replace spaces with underscores)

**Date format:** YYYYMMDD (e.g., 20251203)

### 5.3 Prospect Report YAML Frontmatter

```yaml
---
consultant: [Name from service definition]
service_definition: [Path to service definition file used]
generated_by: /findclient
generated_on: [ISO8601 timestamp]
output_type: client_prospects
status: final
version: 1.0
search_criteria:
  industries: [List of target industries]
  company_sizes: [List of target sizes or "All"]
  location: [Location filter or "Global"]
  limit: [Number limit]
prospect_count:
  high_priority: [Count]
  medium_priority: [Count]
  low_priority: [Count]
  total: [Total count]
---
```

### 5.4 Executive Summary Section

```markdown
# Client Prospect Report: [Domain Focus]

**Generated:** [Date]
**Consultant:** [Name]
**Search Focus:** [Primary industries and services]

## Executive Summary

### Prospect Pipeline Overview
- **HIGH Priority (8-10 score)**: [X] companies - Active outreach recommended
- **MEDIUM Priority (5-7.9 score)**: [Y] companies - Worth pursuing with right introduction
- **LOW Priority (1-4.9 score)**: [Z] companies - Deprioritize unless specific opportunity

### Top 3 Opportunities
1. **[Company 1]** (Score: [X.X]/10) - [One-line value proposition]
2. **[Company 2]** (Score: [X.X]/10) - [One-line value proposition]
3. **[Company 3]** (Score: [X.X]/10) - [One-line value proposition]

### Market Intelligence Summary
- [Key finding about target market receptiveness]
- [Common pain points identified across prospects]
- [Entry point patterns observed]

### Recommended Next Steps
1. [Immediate action for top prospect]
2. [Outreach strategy for high-priority group]
3. [Pipeline development approach]
```

### 5.5 HIGH Priority Prospects (Full Detail)

For each HIGH priority prospect (score 8-10), provide comprehensive profile:

```markdown
---

## [Company Name]

**B2B Fit Score: [X.X]/10** ⭐ HIGH PRIORITY

### Company Overview
- **Industry:** [Primary industry]
- **Size:** [Employee count] employees | [Revenue if known]
- **Headquarters:** [City, State/Province, Country]
- **Website:** [URL]

### B2B Fit Scoring Breakdown

| Factor | Score | Weight | Contribution | Evidence Quality |
|--------|-------|--------|--------------|------------------|
| Contractor History | [X]/10 | 25% | [X.XX] | [HIGH/MED/LOW] |
| Procurement Access | [X]/10 | 20% | [X.XX] | [HIGH/MED/LOW] |
| Domain Alignment | [X]/10 | 25% | [X.XX] | [HIGH/MED/LOW] |
| Size/Budget Fit | [X]/10 | 15% | [X.XX] | [HIGH/MED/LOW] |
| Geographic Match | [X]/10 | 15% | [X.XX] | [HIGH/MED/LOW] |
| **TOTAL** | **[X.X]/10** | 100% | **[X.XX]** | **[HIGH/MED/LOW]** |

**Overall Confidence:** [HIGH/MEDIUM/LOW] - [Brief rationale]

### Score Rationale
[2-3 sentence explanation of why this company scored well, highlighting strongest factors]

### Service Alignment

| Service Offering | Fit Level | Evidence |
|------------------|-----------|----------|
| [Service 1 from definition] | HIGH/MED/LOW | [Specific evidence of need] |
| [Service 2 from definition] | HIGH/MED/LOW | [Specific evidence of need] |

**Primary Value Proposition:**
[Service-specific value prop for this client based on pain points and service alignment]

### Entry Points & Outreach Strategy

#### Entry Point 1: [Type - e.g., Procurement Portal]
- **Access:** [URL or contact method]
- **Process:** [Registration steps, timeline, requirements]
- **Notes:** [Any specific requirements or timing considerations]
- **Ease:** [EASY/MODERATE/COMPLEX]

#### Entry Point 2: [Type - e.g., Warm Introduction]
- **Contact:** [Name, Title]
- **Relationship:** [Connection path, e.g., "2nd degree via Jane Smith (RICS)"]
- **Outreach Strategy:** [How to leverage connection]
- **Ease:** [EASY/MODERATE/COMPLEX]

#### Entry Point 3: [Type - e.g., Direct Decision-Maker]
- **Decision Maker:** [Name, Title]
- **LinkedIn:** [Profile URL if found]
- **Accessibility:** [Open to connections? Active on LinkedIn?]
- **Approach:** [Recommended outreach method]
- **Ease:** [EASY/MODERATE/COMPLEX]

### Recommended Outreach Approach

**Priority:** [IMMEDIATE/THIS WEEK/THIS MONTH]

**Step-by-Step Strategy:**
1. [First action - e.g., "Register on vendor portal (Q1 deadline approaching)"]
2. [Second action - e.g., "Request LinkedIn introduction via [mutual contact]"]
3. [Third action - e.g., "Prepare pitch deck focusing on [specific pain point]"]
4. [Follow-up - e.g., "Follow up with [decision maker] within 2 weeks"]

**Pitch Focus:**
- **Lead with:** [Primary pain point or value proposition]
- **Emphasize:** [Competitive advantage from service definition]
- **Proof points to highlight:** [Specific proof points relevant to this client]

**Estimated Engagement Value:**
- **Likely engagement type:** [Project/Retainer/Hourly from service definition]
- **Estimated project size:** [Hours or dollar range]
- **Timeline:** [Typical duration from service definition]

### Evidence & Sources

**Contractor History Evidence:**
- [Source 1 with URL and key quote or finding]
- [Source 2 with URL and key quote or finding]

**Procurement Accessibility Evidence:**
- [Source 1 with URL]

**Domain Alignment Evidence:**
- [Source 1 with URL - pain point or initiative]
- [Source 2 with URL - technology need]

**Size/Budget Evidence:**
- [Source 1 - employee count, revenue, funding]

**Geographic Evidence:**
- [Source 1 - location, remote culture]

### Red Flags & Risk Factors
[If any concerns identified, list them here. Otherwise state "None identified."]

---
```

Repeat full detail template for ALL high-priority prospects.

### 5.6 MEDIUM Priority Prospects (Condensed Format)

For each MEDIUM priority prospect (score 5-7.9), provide condensed profile:

```markdown
---

## [Company Name]

**B2B Fit Score: [X.X]/10** ⚠️ MEDIUM PRIORITY

**Industry:** [Industry] | **Size:** [Employee count] | **Location:** [City, State/Province]

### Why Medium Priority
[1-2 sentences explaining scoring - what's strong, what's weaker]

**Scoring Summary:** Contractor History: [X]/10, Procurement: [X]/10, Domain: [X]/10, Budget: [X]/10, Geographic: [X]/10

### Service Fit
- **Best Match:** [Service name] - [Brief evidence]
- **Value Prop:** [One-sentence pitch focus]

### Entry Points
1. **[Type]**: [Brief description] - Ease: [EASY/MODERATE/COMPLEX]
2. **[Type]**: [Brief description] - Ease: [EASY/MODERATE/COMPLEX]

### Outreach Strategy
**When to pursue:** [Condition - e.g., "When you have warm introduction" or "After establishing case studies in industry"]
**Approach:** [1-2 sentence recommendation]

### Key Sources
- [Most important source with URL]

---
```

### 5.7 LOW Priority Prospects (Brief Listing)

For LOW priority prospects (score 1-4.9), provide minimal listing:

```markdown
## Low Priority Prospects (Score 1-4.9)

*These prospects have limited B2B fit based on available evidence. Include for completeness but deprioritize unless specific opportunities arise.*

| Company | Score | Industry | Location | Key Gap |
|---------|-------|----------|----------|---------|
| [Company 1] | [X.X]/10 | [Industry] | [Location] | [Primary reason for low score] |
| [Company 2] | [X.X]/10 | [Industry] | [Location] | [Primary reason for low score] |

**Common issues:** [Pattern in low scores - e.g., "Most lack evidence of contractor use" or "Budget constraints"]
```

### 5.8 Market Intelligence & Patterns Section

```markdown
## Market Intelligence Summary

### Industry Insights

**[Industry 1] Market:**
- **Contractor Receptiveness:** [HIGH/MEDIUM/LOW based on signals across prospects]
- **Common Pain Points:** [Top 2-3 pain points identified across companies]
- **Typical Engagement Model:** [Pattern observed - project vs retainer]
- **Procurement Characteristics:** [Open/Moderate/Restrictive]

**[Industry 2] Market:**
[Same structure]

### Entry Point Analysis

**Most Accessible Entry Points:**
1. **[Entry type]** - [X]% of prospects have this ([X] companies)
   - Example: [Company with best example of this entry point]
2. **[Entry type]** - [X]% of prospects ([X] companies)
   - Example: [Company example]

**Most Challenging Entry Points:**
- [Challenge type] - [X] companies require this

### Competitive Landscape Observations
- [Insight about competitors serving these clients]
- [Gaps in market that consultant could fill]

### Geographic Patterns
- **Remote-Friendly:** [X]% of prospects accept remote consulting
- **Location Clusters:** [Cities/regions with multiple prospects]

### Budget & Pricing Insights
- **Enterprise clients ([X] prospects):** Can support $[range] projects
- **Mid-market clients ([X] prospects):** Typical budget $[range]
- **Pricing strategy recommendation:** [Insight based on market findings]
```

### 5.9 Next Steps & Action Plan

```markdown
## Recommended Action Plan

### Immediate Actions (This Week)

#### Top 3 Priority Outreach:
1. **[Company 1]** - [Specific action, e.g., "Register on vendor portal before Dec 15 deadline"]
2. **[Company 2]** - [Specific action, e.g., "Request intro via [mutual contact name]"]
3. **[Company 3]** - [Specific action, e.g., "Prepare tailored pitch deck for [pain point]"]

### Short-Term Development (This Month)

#### Warm Introduction Cultivation:
- Reach out to [X] mutual connections at HIGH priority companies
- Join [relevant association] to expand network in [industry]

#### Content/Positioning:
- Create case study highlighting [specific service] success
- Publish thought leadership on [pain point] common to prospects
- Update LinkedIn with [competitive advantage] positioning

#### Procurement Preparation:
- Register on [X] vendor portals before next cycle
- Prepare standard proposal templates for [service type]
- Gather required documentation (insurance, references, etc.)

### Medium-Term Pipeline Development (Next Quarter)

#### MEDIUM Priority Cultivation:
- Monitor [X] MEDIUM prospects for trigger events (funding, leadership changes, initiatives)
- Build credibility in [industry] through [speaking/writing/projects]
- Develop case studies to address domain skepticism

#### Market Positioning:
- Establish presence at [industry conference/association]
- Partner with [intermediary type] serving these clients
- Build referral relationships with [complementary service providers]

### Long-Term Strategy

#### Market Expansion:
- [If certain industries over-represented]: Diversify into [secondary industry]
- [If geographic gaps]: Expand remote service delivery capabilities
- [If size gaps]: Develop service packages for [underserved size segment]

#### Competitive Differentiation:
- Leverage [strongest competitive advantage] in all outreach
- Develop proprietary [methodology/tool/framework] aligned with market needs
- Build reference portfolio in [highest-value industry]
```

### 5.10 Research Methodology Documentation

```markdown
## Appendix: Research Methodology

### Service Definition Source
- **File:** [Path to service definition]
- **Version:** [Version number]
- **Generated:** [Date]
- **Services Analyzed:** [Count and list]

### Search Strategy
- **Queries Executed:** [Number of web searches]
- **Primary Sources:** [Types of sources used - LinkedIn, Glassdoor, company websites, etc.]
- **Time Investment:** ~[X] minutes total research

### Scoring Framework
- **Method:** 5-factor weighted scoring (10-point scale)
- **Weights:** Contractor History 25%, Procurement 20%, Domain 25%, Budget 15%, Geographic 15%
- **Evidence Standards:** Direct evidence preferred, inferred scoring noted

### Limitations & Caveats
- **Public Information Only:** Research based on publicly available sources
- **Point-in-Time:** Company circumstances may change rapidly
- **Scoring Subjectivity:** Some factors require judgment (especially domain alignment)
- **Evidence Gaps:** [Note any systematic gaps - e.g., "Private companies lack financial data"]

### Recommended Follow-Up Research
For HIGH priority prospects before outreach:
- [ ] Verify decision-maker contact information
- [ ] Research recent company news (last 30 days)
- [ ] Check for mutual connections on LinkedIn
- [ ] Review latest job postings for signals
- [ ] Identify upcoming events/opportunities for engagement

---

**Report Generated:** [ISO8601 timestamp]
**Command:** `/findclient [arguments used]`
**Consultant:** [Name]
**Next Update Recommended:** [Date 3 months out - quarterly refresh]
```

---

## Phase 6: Save Output and Generate Summary

### 6.1 Save Prospect Report

**Primary Output:**
Save complete markdown report to: `Client_Prospects/Prospects_[Domain]_[YYYYMMDD].md`

**Domain sanitization:**
- Extract primary domain from service definition or ideal job
- Examples: "AI_Consulting", "RealEstate_Advisory", "DataScience"
- Remove spaces, special characters
- Use underscores for readability

**Date format:** YYYYMMDD using current date

### 6.2 Create Client Prospects Directory (if needed)

If `Client_Prospects/` directory doesn't exist:
```
mkdir -p Client_Prospects
```

### 6.3 Console Summary Output

After saving report, provide executive summary to console:

```
✅ CLIENT PROSPECT RESEARCH COMPLETE

**Output File:**
Client_Prospects/Prospects_[Domain]_[Date].md

**Prospect Pipeline:**
- HIGH Priority (8-10): [X] companies - Active outreach recommended
- MEDIUM Priority (5-7.9): [Y] companies - Worth pursuing with right introduction
- LOW Priority (1-4.9): [Z] companies - Deprioritize

**Top 3 Opportunities:**
1. [Company 1] ([Score]/10) - [Industry] - [Primary entry point]
2. [Company 2] ([Score]/10) - [Industry] - [Primary entry point]
3. [Company 3] ([Score]/10) - [Industry] - [Primary entry point]

**Market Intelligence Highlights:**
- [Key finding 1]
- [Key finding 2]
- [Key finding 3]

**Recommended Immediate Actions:**
1. [Action for prospect 1]
2. [Action for prospect 2]
3. [Action for prospect 3]

**Research Summary:**
- [X] companies researched
- [Y] web searches conducted
- ~[Z] minutes research time
- Evidence confidence: [HIGH/MEDIUM/LOW overall]

**Next Steps:**
1. Review full prospect profiles in saved report
2. Prioritize top 3 HIGH priority prospects for outreach
3. Prepare tailored pitch materials for each prospect
4. Execute recommended entry point strategies
5. Update prospect list quarterly as market evolves

**Note:** This research is point-in-time. Validate key details before outreach.
```

---

## Error Handling & Edge Cases

### Error 1: No Service Definition Found

```
❌ SERVICE DEFINITION REQUIRED

No service definition file found in Client_Prospects/

Before identifying clients, you must define your service offerings:
  /defineservices

This command creates:
- Service catalog with ideal client profiles
- Pricing structure and engagement models
- Competitive differentiation and proof points
- Target market definitions

Run /defineservices first, then return to /findclient.
```

### Error 2: No Prospects Found

```
⚠️ NO PROSPECTS IDENTIFIED

Search criteria:
- Industries: [list]
- Size: [filter]
- Location: [filter]

No companies matching criteria and minimum B2B fit threshold (score ≥ 1.0).

**Suggestions:**
1. Broaden industry criteria - try adjacent industries
2. Remove size or location filters
3. Review service definition ideal_client profiles - may be too narrow
4. Try different geographic region
5. Research may need manual supplementing (industry directories, associations)

**Alternative approaches:**
- Search industry association member directories manually
- Attend industry conferences to identify prospects in person
- Use LinkedIn Sales Navigator with broader filters
- Partner with industry-specific staffing agencies
```

### Error 3: All Low Priority Prospects

```
⚠️ LIMITED HIGH-QUALITY PROSPECTS

Research complete, but prospect quality concerns:
- HIGH Priority (8-10): 0 companies
- MEDIUM Priority (5-7.9): [X] companies
- LOW Priority (1-4.9): [Y] companies

**Likely issues:**
- [Most common gap - e.g., "Lack of contractor history evidence in target industries"]
- [Secondary issue - e.g., "Budget constraints in target company sizes"]

**Recommendations:**
1. Review MEDIUM priority prospects - some may warrant deeper research
2. Consider expanding target industries in service definition
3. Adjust company size targets (enterprise vs mid-market)
4. Build credibility through case studies before targeting these clients
5. Consider intermediary routes (staffing agencies, consulting marketplaces)

Report saved with available prospects for review.
```

### Error 4: Web Research Failures

```
⚠️ RESEARCH LIMITATIONS

Some web searches failed or returned limited results:
- [X] queries blocked or rate-limited
- [Y] companies with insufficient public information

**Impact:**
- Scoring confidence reduced for [Z] prospects
- May be missing contractor signals or entry points

**Mitigation:**
- Marked low-confidence scores with [LOW] evidence quality
- Recommend manual validation before outreach
- Consider LinkedIn Sales Navigator or paid databases for better company intel

Report generated with available data. Review evidence quality ratings.
```

### Error 5: Invalid Filter Values

```
❌ INVALID FILTER: --size=[value]

Valid company size filters:
- startup (1-50 employees)
- mid (50-500 employees)
- enterprise (500+ employees)

Received: [invalid value]

Usage: /findclient [ideal-job-file] [--size=startup|mid|enterprise] [other filters]
```

### Error 6: Limit Too High

```
⚠️ LIMIT ADJUSTED

Requested: --limit=[X]
Maximum recommended: 50 (research time constraints)

Adjusted to 50 prospects for this run.

**Note:** Researching 50 companies with B2B signal analysis takes ~15-20 minutes.
For larger prospect lists, consider:
- Multiple runs with different industry filters
- Broader initial discovery, then manual filtering
- Focus on highest-value industries first
```

---

## Usage Examples

### Example 1: Basic Discovery (Using Service Definition Only)

```bash
/findclient
```

**Behavior:**
- Load service definition from `Client_Prospects/Service_Definition_*.md`
- Use ideal_client profiles from all services
- No additional filters (global search)
- Default limit of 20 prospects
- Output: `Client_Prospects/Prospects_[Domain]_20251203.md`

### Example 2: Industry-Specific Discovery

```bash
/findclient --industry="Commercial Real Estate"
```

**Behavior:**
- Filter to only CRE companies
- Ignore other industries from service definition
- Focus queries and scoring on CRE market
- Output: `Client_Prospects/Prospects_CommercialRealEstate_20251203.md`

### Example 3: Enterprise Clients in Specific Location

```bash
/findclient --size=enterprise --location="Toronto" --limit=15
```

**Behavior:**
- Only enterprise companies (500+ employees)
- Toronto headquarters or presence
- Limit to 15 prospects (focused research)
- Output: `Client_Prospects/Prospects_[Domain]_Toronto_20251203.md`

### Example 4: Ideal Job-Driven Discovery

```bash
/findclient Job_Postings/IdealJob_Synthetic_20251201.md --limit=25
```

**Behavior:**
- Load ideal job file for additional targeting criteria
- Merge with service definition ideal client profiles
- Industries/tech stack from ideal job guide research
- 25 prospect limit
- Output: `Client_Prospects/Prospects_[Domain]_20251203.md`

### Example 5: Comprehensive Discovery Across Multiple Industries

```bash
/findclient --limit=50
```

**Behavior:**
- All industries from service definition
- Maximum recommended limit (50 prospects)
- ~15-20 minutes research time
- Comprehensive market view
- Output: `Client_Prospects/Prospects_[Domain]_20251203.md`

### Example 6: Startup-Focused Discovery

```bash
/findclient --size=startup --industry="AI/ML" --limit=30
```

**Behavior:**
- Startups only (1-50 employees)
- AI/ML industry focus
- 30 prospect limit
- Emphasis on funding/budget validation
- Output: `Client_Prospects/Prospects_AI_Startups_20251203.md`

---

## Important Notes

### Research Ethics & Limitations

- **Public Information Only:** All research uses publicly available sources (LinkedIn, Glassdoor, company websites, news)
- **No Private Data:** Never access private databases, internal documents, or confidential information
- **Respect Privacy:** When identifying decision-makers, use only public LinkedIn profiles and published contact info
- **Point-in-Time:** Company circumstances change rapidly; validate before outreach
- **No Guarantees:** B2B fit score is probabilistic, not predictive

### Evidence Standards

- **Direct Evidence (HIGH confidence):** Job postings, vendor portals, annual reports, press releases
- **Indirect Evidence (MEDIUM confidence):** Glassdoor reviews, LinkedIn profiles, industry reports
- **Inferred (LOW confidence):** Based on company size/industry norms, calculated estimates

### Scoring Objectivity

The B2B fit scoring methodology is designed to be:
- **Systematic:** Same criteria applied to all prospects
- **Weighted:** Factors weighted by importance (contractor history + domain alignment = 50%)
- **Evidence-Based:** Scores tied to specific evidence (cited in report)
- **Transparent:** Scoring breakdown shown for each prospect

However, some subjectivity exists in:
- Domain alignment interpretation (pain point matching)
- Budget estimation (when financial data unavailable)
- Entry point ease assessment

### Follow-Up Research Before Outreach

Before contacting HIGH priority prospects, conduct:
1. **Fresh news check:** Google "[company] news" for last 30 days
2. **Decision-maker validation:** Verify contact still in role (LinkedIn)
3. **Warm intro confirmation:** Confirm mutual connection willing to introduce
4. **Procurement update:** Check vendor portal for any process changes

### Quarterly Refresh Recommended

Market conditions and company needs evolve. Update prospect list every 3 months:
```bash
/findclient [same filters] --limit=[same limit]
```

Compare new vs. old reports to identify:
- Companies that have moved to HIGH priority (trigger events)
- Procurement portal opening/closing cycles
- New decision-makers (outreach opportunities)
- Market trend shifts

---

Now executing client discovery research...
