---
description: Identify potential B2B clients matching service definition and ideal client profiles
argument-hint: [ideal-job-file] [--industry=X] [--size=startup|mid|enterprise] [--location=X] [--limit=N]
---

Identifies potential B2B clients through intelligent discovery with 10-point B2B fit scoring and entry point mapping.

**Input Options:**
- `<ideal-job-file>`: Optional ideal job file from `/idealjob`
- `--industry=X`: Filter by industry
- `--size=startup|mid|enterprise`: Company size filter (startup: 1-50, mid: 50-500, enterprise: 500+)
- `--location=X`: Geographic filter
- `--limit=N`: Max prospects (default: 20, max: 50)

**Output**: `Client_Prospects/Prospects_[Domain]_[Date].md`

---

## Phase 1: Load Prerequisites and Build Targeting Profile

### 1.1 Check Service Definition Prerequisite

Search for `Client_Prospects/Service_Definition_*.md`. If NOT found, stop with:
```
❌ SERVICE DEFINITION REQUIRED
Run /defineservices first to create service catalog with ideal client profiles, pricing, and differentiation.
```

If multiple found, use most recent (latest date in filename).

### 1.2 Parse Service Definition

Extract from YAML: `consultant`, `generated_on`, `version`

Extract from markdown body:
- **Services**: name, category, ideal_client (industries, company_size, pain_points, decision_makers)
- **Target Market**: primary_markets, secondary_markets, avoid_markets, geographic_focus
- **Differentiation**: competitive_advantages, proof_points

### 1.3 Parse Command-Line Filters and Optional Ideal Job

**Filters:**
- `--industry=X`: Narrow to specified industry
- `--size=startup|mid|enterprise`: Filter by company size
- `--location=X`: Geographic constraint
- `--limit=N`: Default 20, max 50

**If ideal job file provided** ({{ARG1}} ends with `.md`):
Extract: target industries, company size, location, tech stack, organizational characteristics

### 1.4 Consolidate Targeting Profile

**Priority order:** Command-line filters > Ideal job criteria > Service definition

Create merged targeting profile with industries, company sizes, geographic focus, pain points, decision makers.

---

## Phase 2: Research Queries and Discovery Strategy

### 2.1 Generate 10-15 Web Search Queries

Create targeted queries using service-industry intersections:

**Query Types:**
1. **Industry + Contractor**: "[industry] companies hiring consultants 2025", "[industry] contractor jobs [location]"
2. **Pain Point + Service**: "[industry] [pain point] consulting RFP", "[industry] need for [service]"
3. **Procurement**: "[industry] vendor registration portal", "[industry] consulting spend"
4. **Company Discovery**: "top [industry] companies [location]", "[industry] digital transformation 2025"
5. **Intermediaries**: "[industry] consulting marketplaces", "[industry] staffing agencies"
6. **Associations/Events**: "[industry] association directory", "[industry] conference sponsors 2025"

Customize with current year (2025), location filter, target industries, and pain points.

**Time Budget:** ~6-8 minutes for 20 prospects (discovery 3-4min, B2B signals 2-3min, entry points 1-2min)

---

## Phase 3: Discovery and B2B Signal Research

### 3.1 Execute Company Discovery

Run web searches to identify 30-50 candidate companies (will filter to limit after scoring).

**Extract:** Company name, industry, HQ location, size (employees/revenue), website, description

**Prioritize sources:** LinkedIn pages, industry directories, conference sponsors, government registries, transformation news

**Quality filters:** Exclude staffing agencies, individual consultants, <5 employees (unless startup filter)

### 3.2 Research 5 B2B Signal Types (30-60 sec per company)

**1. Contractor History (25% weight)**
- Search: "[company] consultant/contractor job postings", site:linkedin.com "[company] consultant", site:glassdoor.com "[company] contractor"
- Evidence: Job posts (recent = strong), Glassdoor contractor mentions, LinkedIn contractor profiles, vendor pages

**2. Procurement Accessibility (20% weight)**
- Search: "[company] vendor registration", "[company] supplier portal", "[company] small business program"
- Evidence: Vendor portals, small business programs, streamlined processes, RFPs

**3. Domain Alignment (25% weight)**
- Check: Industry match, pain point evidence (press releases, job postings, news), tech stack, initiatives
- Evidence: Press releases on relevant initiatives, job postings showing gaps, company website priorities

**4. Size/Budget Fit (15% weight)**
- Search: "[company] revenue/funding/employee count", site:linkedin.com "[company] employees"
- Evidence: Revenue (filings/news), funding rounds, employee count, consulting spend
- Budget capacity: Enterprise (500+): $50K-$500K+, Mid (100-500): $10K-$100K, Small (50-100): $5K-$25K, Startup (<50): <$10K

**5. Geographic Match (15% weight)**
- Check: HQ location, remote culture, multiple offices, international presence
- Evidence: Remote job postings, Glassdoor reviews on flexibility, distributed employee locations

### 3.3 Identify Entry Points (1-2 min for high-scoring prospects)

**5 Entry Point Types:**
1. **Procurement Portals**: URL, process (Open/Quarterly/Invite), requirements, ease
2. **Warm Introductions**: 2nd-degree LinkedIn connections, associations, alumni, former colleagues
3. **Intermediaries**: Staffing agencies serving client, consulting marketplaces
4. **Decision-Makers**: LinkedIn profiles matching ideal_client.decision_makers titles, accessibility
5. **Content/Events**: Conferences, guest posting, LinkedIn engagement, webinars

---

## Phase 4: 10-Point B2B Fit Scoring

### 4.1 Scoring Formula

```
Total Score = (Contractor History × 0.25) + (Procurement Accessibility × 0.20) +
              (Domain Alignment × 0.25) + (Size/Budget Fit × 0.15) + (Geographic Match × 0.15)
```

Each factor: 0-10 points. Final: 0.0-10.0 (1 decimal)

**Priority:** HIGH (8-10): Active outreach | MEDIUM (5-7.9): Worth pursuing | LOW (1-4.9): Deprioritize

**Confidence:** HIGH (4+ factors with direct evidence) | MEDIUM (2-3 direct) | LOW (mostly inferred)

### 4.2 Factor Scoring Tables

**Factor 1: Contractor History (25%)**

| Score | Evidence |
|-------|----------|
| 10 | Active program + job posts <3mo + positive Glassdoor |
| 8-9 | Job posts <6mo OR consulting spend in reports |
| 6-7 | LinkedIn contractors, job posts 6-12mo |
| 4-5 | Indirect (staffing mentions, industry norms) |
| 2-3 | Minimal, inferred from size (enterprise) |
| 0-1 | No evidence |

Evidence: Job post URLs/dates, Glassdoor quotes, LinkedIn profiles, report citations. Confidence: HIGH (direct), MED (indirect), LOW (inferred)

**Factor 2: Procurement Accessibility (20%)**

| Score | Evidence |
|-------|----------|
| 10 | Public portal + small biz program + streamlined |
| 8-9 | Portal exists + reasonable process |
| 6-7 | Contact identified + public process |
| 4-5 | General info, suggests openness |
| 2-3 | No public info, size suggests process |
| 0-1 | Closed network OR bureaucratic |

Evidence: Portal URLs, policy docs, contacts. Red flags: "Approved only", excessive insurance, >6mo qualification

**Factor 3: Domain Alignment (25%)**

| Score | Evidence |
|-------|----------|
| 10 | Perfect industry + active pain point + recent initiative |
| 8-9 | Target industry + strong pain point OR adjacent + perfect match |
| 6-7 | Industry + some pain point OR adjacent + good fit |
| 4-5 | Adjacent + transferable pain points |
| 2-3 | Weak alignment, speculative fit |
| 0-1 | Mismatch OR no pain point evidence |

Sub-scoring: Industry match (+4 exact, +2 adjacent), Pain point (+3 exact, +2 job gaps, +1 reports), Service need (+3 direct, +2 strategic, +1 general)

**Factor 4: Size/Budget Fit (15%)**

| Score | Evidence |
|-------|----------|
| 10 | 1000+ employees + consulting spend >$1M/yr |
| 8-9 | 500-1000 employees OR mid + $50M+ revenue |
| 6-7 | 100-500 employees + $10M-$50M revenue |
| 4-5 | 50-100 employees + funding/revenue |
| 2-3 | <50 employees + Series A+ funding |
| 0-1 | <20 employees, no budget indicators |

Budget capacity: Enterprise $50K-$500K+, Mid $10K-$100K, Small $5K-$25K, Startup <$10K. Evidence: Revenue, employee count, funding, consulting spend

**Factor 5: Geographic Match (15%)**

| Score | Evidence |
|-------|----------|
| 10 | HQ in target + remote-friendly confirmed |
| 8-9 | HQ in target OR remote + adjacent |
| 6-7 | National presence + remote option |
| 4-5 | Different geo + strong remote |
| 2-3 | Different geo, limited remote |
| 0-1 | Different geo + no remote (on-site) |

Location filter scoring: HQ match +5, office +3, adjacent +2, remote national +2. Remote signals: Job posts +3, Glassdoor +2, policy +2

---

## Phase 5: Prospect Report Generation

### 5.1 Sort and Filter

Sort by score (descending), group by priority. Limit to `--limit` (default 20). Target distribution (limit=20): HIGH 8-12, MEDIUM 6-10, LOW 0-2.

### 5.2 Report Structure

**File:** `Client_Prospects/Prospects_[Domain]_[YYYYMMDD].md` (domain = service category/industry, sanitized)

**YAML Frontmatter:**
```yaml
---
consultant: [Name]
service_definition: [Path]
generated_by: /findclient
generated_on: [ISO8601]
output_type: client_prospects
status: final
version: 1.0
search_criteria:
  industries: [List]
  company_sizes: [List or "All"]
  location: [Filter or "Global"]
  limit: [Number]
prospect_count:
  high_priority: [Count]
  medium_priority: [Count]
  low_priority: [Count]
  total: [Total]
---
```

### 5.3 Executive Summary

```markdown
# Client Prospect Report: [Domain Focus]

**Generated:** [Date] | **Consultant:** [Name] | **Search Focus:** [Industries/services]

## Executive Summary

### Prospect Pipeline Overview
- **HIGH Priority (8-10)**: [X] companies - Active outreach recommended
- **MEDIUM Priority (5-7.9)**: [Y] companies - Worth pursuing with right introduction
- **LOW Priority (1-4.9)**: [Z] companies - Deprioritize

### Top 3 Opportunities
1. **[Company 1]** ([X.X]/10) - [One-line value prop]
2. **[Company 2]** ([X.X]/10) - [One-line value prop]
3. **[Company 3]** ([X.X]/10) - [One-line value prop]

### Market Intelligence
- [Market receptiveness finding]
- [Common pain points]
- [Entry point patterns]

### Recommended Next Steps
1. [Immediate action for top prospect]
2. [High-priority group strategy]
3. [Pipeline development]
```

### 5.4 HIGH Priority Format (Full Detail for ALL)

```markdown
## [Company Name]

**B2B Fit Score: [X.X]/10** ⭐ HIGH PRIORITY

### Company Overview
**Industry:** [Industry] | **Size:** [Employees] employees | [Revenue] | **HQ:** [City, Province, Country] | **Website:** [URL]

### B2B Fit Scoring Breakdown

| Factor | Score | Weight | Contribution | Evidence Quality |
|--------|-------|--------|--------------|------------------|
| Contractor History | [X]/10 | 25% | [X.XX] | [HIGH/MED/LOW] |
| Procurement Access | [X]/10 | 20% | [X.XX] | [HIGH/MED/LOW] |
| Domain Alignment | [X]/10 | 25% | [X.XX] | [HIGH/MED/LOW] |
| Size/Budget Fit | [X]/10 | 15% | [X.XX] | [HIGH/MED/LOW] |
| Geographic Match | [X]/10 | 15% | [X.XX] | [HIGH/MED/LOW] |
| **TOTAL** | **[X.X]/10** | 100% | **[X.XX]** | **[HIGH/MED/LOW]** |

**Overall Confidence:** [HIGH/MED/LOW] - [Rationale]

### Score Rationale
[2-3 sentences explaining score, highlighting strongest factors]

### Service Alignment
| Service | Fit | Evidence |
|---------|-----|----------|
| [Service 1] | HIGH/MED/LOW | [Specific need evidence] |

**Primary Value Prop:** [Service-specific value prop]

### Entry Points & Outreach
#### Entry Point 1: [Type]
**Access:** [URL/method] | **Process:** [Steps/timeline] | **Ease:** [EASY/MOD/COMPLEX] | **Notes:** [Details]

#### Entry Point 2-3: [Repeat format]

### Recommended Outreach
**Priority:** [IMMEDIATE/THIS WEEK/THIS MONTH]

**Steps:**
1. [Action 1]
2. [Action 2]
3. [Action 3]

**Pitch:** Lead with [pain point], emphasize [advantage], highlight [proof points]

**Estimated Value:** [Engagement type], [project size], [timeline]

### Evidence & Sources
**Contractor History:** [Source URLs/quotes]
**Procurement:** [Source URLs]
**Domain:** [Source URLs]
**Budget:** [Source data]
**Geographic:** [Source data]

### Red Flags
[Concerns or "None identified"]
```

### 5.5 MEDIUM Priority Format (Condensed)

```markdown
## [Company Name]

**B2B Fit Score: [X.X]/10** ⚠️ MEDIUM PRIORITY

**Industry:** [Industry] | **Size:** [Employees] | **Location:** [City, Province]

**Why Medium:** [1-2 sentences on strengths/weaknesses]

**Scoring:** Contractor [X]/10, Procurement [X]/10, Domain [X]/10, Budget [X]/10, Geographic [X]/10

**Service Fit:** [Service] - [Brief evidence] | **Value Prop:** [One sentence]

**Entry Points:** 1. [Type]: [Brief] - Ease: [EASY/MOD/COMPLEX] | 2. [Type]: [Brief] - Ease: [EASY/MOD/COMPLEX]

**Outreach:** When: [Condition] | Approach: [1-2 sentences]

**Key Source:** [URL]
```

### 5.6 LOW Priority Format (Table)

```markdown
## Low Priority Prospects (1-4.9)

| Company | Score | Industry | Location | Key Gap |
|---------|-------|----------|----------|---------|
| [Co 1] | [X.X]/10 | [Ind] | [Loc] | [Reason] |

**Common issues:** [Pattern]
```

### 5.7 Market Intelligence & Action Plan

```markdown
## Market Intelligence

**[Industry 1]:** Contractor receptiveness: [H/M/L], Pain points: [List], Engagement: [Type], Procurement: [Open/Mod/Restrictive]

**Entry Points:** [Type] - [X]% ([X] companies) | [Most challenging]: [Type] - [X] companies

**Geographic:** [X]% remote-friendly, Clusters: [Regions]

**Budget:** Enterprise ([X]): $[range], Mid ([X]): $[range]

## Recommended Action Plan

**Immediate (This Week):**
1. [Company 1] - [Action]
2. [Company 2] - [Action]
3. [Company 3] - [Action]

**Short-Term (This Month):** Warm intros: [Actions], Content: [Actions], Procurement: [Actions]

**Medium-Term (Next Quarter):** MEDIUM cultivation: [Actions], Positioning: [Actions]

**Long-Term:** Market expansion: [Recommendations], Differentiation: [Recommendations]

## Appendix: Methodology

**Service Definition:** [Path] v[X] ([Date]), [X] services
**Search:** [X] queries, [Sources], ~[X] minutes
**Scoring:** 5-factor weighted (Contractor 25%, Procurement 20%, Domain 25%, Budget 15%, Geographic 15%)
**Limitations:** Public info only, point-in-time, domain scoring subjective, [evidence gaps]
**Follow-Up Checklist:** [ ] Verify contacts, [ ] Recent news, [ ] Mutual connections, [ ] Job postings, [ ] Events

**Generated:** [ISO8601] | **Command:** /findclient [args] | **Next Update:** [Date +3mo]
```

---

## Phase 6: Save and Summarize

### 6.1 Save Report

Save to: `Client_Prospects/Prospects_[Domain]_[YYYYMMDD].md` (domain = service category/industry, sanitized). Create directory if needed: `mkdir -p Client_Prospects`

### 6.2 Console Summary

```
✅ CLIENT PROSPECT RESEARCH COMPLETE

**Output:** Client_Prospects/Prospects_[Domain]_[Date].md

**Pipeline:** HIGH (8-10): [X] | MEDIUM (5-7.9): [Y] | LOW (1-4.9): [Z]

**Top 3:** 1. [Co1] ([Score]/10) - [Industry] - [Entry] | 2. [Co2] ... | 3. [Co3] ...

**Market Highlights:** [3 key findings]

**Immediate Actions:** [3 specific actions for top prospects]

**Research:** [X] companies, [Y] searches, ~[Z] min, Evidence: [H/M/L]

**Next:** Review report → Prioritize top 3 → Prepare pitches → Execute strategies → Quarterly refresh
**Note:** Point-in-time research. Validate before outreach.
```

---

## Error Handling

**No Service Definition:** Stop with message to run `/defineservices` first

**No Prospects:** Suggest broaden criteria, remove filters, review ideal_client profiles, try adjacent industries/regions, use association directories/conferences/LinkedIn Navigator

**All Low Priority:** Note quality concern, suggest review MEDIUM prospects, expand industries, adjust size targets, build credibility first, consider intermediaries

**Web Failures:** Note confidence reduced, mark LOW quality, recommend manual validation, suggest paid databases

**Invalid Filters:** Display valid options (size: startup|mid|enterprise)

**Limit Too High:** Adjust to 50 max, note 15-20min time, suggest multiple runs or focused industry

---

## Usage Examples

1. **Basic:** `/findclient` → Uses service def, all services, global, limit 20
2. **Industry:** `/findclient --industry="Commercial Real Estate"` → CRE only
3. **Enterprise + Location:** `/findclient --size=enterprise --location="Toronto" --limit=15` → 500+ employees, Toronto
4. **Ideal Job:** `/findclient Job_Postings/IdealJob_Synthetic_20251201.md --limit=25` → Merges criteria
5. **Comprehensive:** `/findclient --limit=50` → All industries, max limit, 15-20min
6. **Startups:** `/findclient --size=startup --industry="AI/ML" --limit=30` → 1-50 employees, AI/ML

---

## Important Notes

**Ethics:** Public info only, no private data, respect privacy, point-in-time, probabilistic scoring

**Evidence:** HIGH (job posts/portals/reports/press), MED (Glassdoor/LinkedIn/industry reports), LOW (size/industry norms/estimates)

**Scoring:** Systematic, weighted (contractor+domain=50%), evidence-based, transparent. Subjectivity: domain matching, budget estimates, entry ease

**Pre-Outreach:** Fresh news (30 days), verify contact role, confirm warm intro, check portal updates

**Quarterly Refresh:** Run same command every 3mo to track trigger events, procurement cycles, decision-maker changes, market shifts

---

Now executing client discovery research...
