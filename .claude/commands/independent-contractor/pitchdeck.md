---
description: Generate professional B2B service pitch deck tailored to prospects or industries
argument-hint: [--prospect=company-name] [--industry=X] [--service=service-name] [--format=md|pptx]
---

I'll help you create a professional B2B service pitch deck for your consulting/contractor business. This command generates 10-12 slide presentations tailored to specific prospects or industries with provenance-hardened claims.

**Input Options:**
- `--prospect=company-name`: Generate prospect-specific pitch with company research (requires prospect file from `/findclient`)
- `--industry=X`: Generate industry-generic pitch for a specific industry vertical
- `--service=service-name`: Deep-dive pitch focused on single service offering
- `--format=md|pptx`: Output format (markdown default, pptx requires pandoc conversion)

**Output File**: `Client_Prospects/Pitch_[Target]_[Date].md`

---

## Phase 1: Validate Prerequisites and Load Service Definition

### 1.1 Check Required Files

**CRITICAL PREREQUISITES**: This command requires two files to exist:

1. **Service Definition** (REQUIRED):
   - Search for: `Client_Prospects/Service_Definition_*.md`
   - If NOT found, guide user:
   ```
   ❌ SERVICE DEFINITION REQUIRED

   Before creating pitch decks, you must define your service offerings.

   Run this command first:
     /defineservices

   This creates a comprehensive service catalog with:
   - Service offerings and deliverables
   - Pricing structure and engagement models
   - Competitive differentiation and proof points
   - Target market definitions

   Once you have a service definition, run /pitchdeck again.
   ```
   Stop execution and guide user.

2. **Candidate Profile** (REQUIRED):
   - Check for: `ResumeSourceFolder/.profile/candidate_profile.json`
   - If NOT found or stale (>7 days old), guide user:
   ```
   ❌ CANDIDATE PROFILE REQUIRED

   Pitch deck generation requires a current candidate profile for provenance validation.

   Run this command first:
     /assessjob <any-job-posting-file.md>

   This automatically generates an optimized candidate profile with:
   - Technical skills and domain expertise
   - Work history and achievements with evidence
   - Certifications and thought leadership
   - Complete provenance trail for all claims

   Once you have a candidate profile, run /pitchdeck again.
   ```
   Stop execution and guide user.

### 1.2 Load Service Definition

**If multiple service definitions found:**
- Use the most recent file (latest date in filename)
- Log which version is being used

**Extract from service definition:**

**From YAML frontmatter:**
- `consultant` name
- `generated_on` timestamp
- `version` number
- `source_profile` path (if available)

**From markdown body - Consultant section:**
- `name`: Consultant/business name
- `tagline`: One-sentence value proposition
- `credentials[]`: Key credentials and certifications
- `years_experience`: Total professional experience

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
- `pricing.value_justification`: ROI explanation
- `success_metrics[]`: How success is measured
- `case_studies[]`: Real examples (if available)

**From markdown body - Rate Card section:**
- `hourly`: Minimum/target/premium hourly rates
- `daily`: Minimum/target/premium daily rates
- `retainer`: Part-time/half-time/full-time options (if available)
- `payment_terms`: Deposit, invoice frequency, payment due

**From markdown body - Differentiation section:**
- `unique_value`: Core UVP (2-3 sentences)
- `competitive_advantages[]`: Advantages with impact
- `proof_points[]`: Quantified achievements
- `certifications[]`: Relevant certifications
- `publications[]`: Thought leadership (if available)
- `speaking[]`: Speaking engagements (if available)

**From markdown body - Target Market section (if exists):**
- `primary_markets[]`: Primary industries with segments
- `secondary_markets[]`: Secondary opportunities
- `geographic_focus[]`: Geographic regions

Store all extracted data for use in pitch deck generation.

### 1.3 Load Candidate Profile

Read `ResumeSourceFolder/.profile/candidate_profile.json` and extract:

**For provenance validation:**
- `work_history[]`: All roles with achievements, metrics, and evidence
- `technical_skills[]`: Skills with proficiency levels and evidence
- `domain_expertise[]`: Domain knowledge with depth and years
- `certifications[]`: Certifications with status (Active only)
- `projects[]`: Project outcomes with evidence
- `thought_leadership`: Publications, speaking, frameworks created
- `leadership_experience`: Team size, budget, scope indicators
- `achievements`: All quantified achievements with evidence

**Create evidence index:**
Map every quantified metric, credential, and claim to source file + line numbers for instant provenance validation during pitch deck generation.

### 1.4 Parse Command-Line Arguments

**Pitch targeting mode (mutually exclusive):**

**Option 1: Prospect-specific pitch (`--prospect=company-name`)**
- Search for prospect file: `Client_Prospects/Prospects_*.md` containing company-name
- If multiple files found, use most recent
- If NOT found:
  ```
  ❌ PROSPECT NOT FOUND: [company-name]

  No prospect file found matching "[company-name]".

  Available prospects from your most recent /findclient run:
  [List companies from most recent Prospects_*.md file]

  Options:
  1. Use one of the available prospects: /pitchdeck --prospect=[available-name]
  2. Run new prospect discovery: /findclient [filters]
  3. Generate industry-generic pitch: /pitchdeck --industry=[industry-name]
  ```
  Stop execution and list available prospects.

**Option 2: Industry-generic pitch (`--industry=X`)**
- Extract industry name from argument
- Validate against service definition target industries
- If industry not in service definition, warn but proceed:
  ```
  ⚠️ INDUSTRY NOT IN SERVICE DEFINITION

  "[Industry]" is not listed in your service definition target industries.

  Target industries from your service definition:
  [List industries from service definition]

  Proceeding with industry-generic pitch, but service alignment may be weaker.
  Consider updating service definition if "[Industry]" is a primary target market.
  ```

**Option 3: Service-focused pitch (`--service=service-name`)**
- Search for matching service in service definition
- Use fuzzy matching (e.g., "AI Agent" matches "AI Agent Development for Due Diligence")
- If no match found:
  ```
  ❌ SERVICE NOT FOUND: [service-name]

  Available services from your service definition:
  [List all service names]

  Usage: /pitchdeck --service="[exact or partial service name]"
  ```
  Stop execution and list available services.

**Option 4: Default (no targeting specified)**
- Use all services from service definition
- Generate comprehensive pitch covering full service catalog
- Industry-agnostic, suitable for general networking

**Output format (`--format=md|pptx`):**
- Default: `md` (markdown)
- If `pptx` specified:
  - Check if pandoc is installed: `pandoc --version`
  - If NOT found, warn and proceed with markdown:
    ```
    ⚠️ PANDOC NOT FOUND - MARKDOWN OUTPUT ONLY

    PowerPoint (pptx) output requires pandoc for conversion.

    Install pandoc:
      /install-pandoc

    Proceeding with markdown output. You can convert manually later.
    ```
  - If found, markdown will be converted to pptx after generation

---

## Phase 2: Conduct Target Research (5-10 minutes max)

### 2.1 Prospect-Specific Research (if `--prospect` mode)

**Load prospect data from file:**
- Company overview (industry, size, location, website)
- B2B fit score breakdown
- Service alignment analysis
- Pain point evidence
- Entry points and decision-makers
- Sources and evidence from discovery research

**Conduct fresh research (time budget: 5-7 minutes):**

**Research Type 1: Recent News & Initiatives (2-3 minutes)**
Search queries:
```
"[company name]" news [current year]
"[company name]" strategic initiatives
"[company name]" digital transformation
"[company name]" [pain point from prospect file]
```

Extract:
- Recent press releases (last 90 days)
- Strategic priorities or transformation projects
- Funding announcements or major contracts
- Leadership changes or organizational shifts
- Technology adoption or modernization efforts

**Research Type 2: Current Pain Points (2 minutes)**
Search queries:
```
"[company name]" challenges [current year]
site:glassdoor.com "[company name]" reviews [pain point]
"[company name]" job postings [technology relevant to service]
"[company name]" RFP consultant
```

Extract:
- Glassdoor reviews mentioning relevant pain points
- Job postings showing skill gaps or technology needs
- RFPs or consulting engagements announced
- Industry reports mentioning company challenges

**Research Type 3: Decision-Maker Update (1-2 minutes)**
Search queries:
```
site:linkedin.com "[company name]" [decision maker title]
"[decision maker name from prospect file]" "[company name]"
```

Extract:
- Verify decision-maker still in role (LinkedIn profile update date)
- New decision-makers if role has changed
- Recent LinkedIn activity or posts (engagement opportunities)
- Professional background and priorities

**Compile research findings:**
```yaml
prospect_research:
  company: "[Company name]"
  research_date: "[ISO8601 timestamp]"
  sources_consulted:
    - url: "[URL]"
      extracted: "[Key finding]"
  recent_initiatives:
    - "[Initiative with date]"
  current_pain_points:
    - "[Pain point with evidence]"
  decision_maker_status:
    - name: "[Name]"
      title: "[Title]"
      status: "Current" | "Changed"
      updated: "[Date]"
  strategic_priorities:
    - "[Priority from news/job postings]"
```

### 2.2 Industry-Generic Research (if `--industry` mode)

**Time budget: 5-8 minutes**

**Research Type 1: Industry Trends & Challenges (3-4 minutes)**
Search queries:
```
"[industry]" trends [current year]
"[industry]" digital transformation [current year]
"[industry]" challenges consulting
"[industry]" market outlook [current year]
"[industry]" technology adoption
```

Extract:
- Top 3-5 industry trends
- Common pain points across industry
- Technology adoption patterns
- Regulatory or compliance challenges
- Market growth or contraction indicators

**Research Type 2: Industry Consulting Market (2-3 minutes)**
Search queries:
```
"[industry]" consulting services demand
"[industry]" [service type] market size
"[industry]" consultant hiring [current year]
"[industry]" consulting spend
```

Extract:
- Market size for relevant consulting services
- Demand signals (job postings, RFPs)
- Typical engagement models in industry
- Competitive landscape (who's winning business)

**Research Type 3: Decision-Maker Titles & Priorities (1-2 minutes)**
Search queries:
```
"[industry]" "[decision maker title]" priorities
"[industry]" CTO challenges
"[industry]" digital transformation leaders
```

Extract:
- Common decision-maker titles for service type
- Typical priorities and KPIs
- Budget authority levels
- Procurement processes

**Compile research findings:**
```yaml
industry_research:
  industry: "[Industry name]"
  research_date: "[ISO8601 timestamp]"
  sources_consulted:
    - url: "[URL]"
      extracted: "[Key finding]"
  industry_trends:
    - "[Trend with context]"
  common_pain_points:
    - "[Pain point with prevalence data]"
  market_size:
    estimate: "[Market size or 'Unknown']"
    growth_rate: "[Percentage or 'Unknown']"
  decision_makers:
    - title: "[Title]"
      typical_priorities: "[Priorities]"
  competitive_landscape:
    - "[Competitor or pattern]"
```

### 2.3 Service-Focused Research (if `--service` mode)

**Time budget: 4-6 minutes**

**Research Type 1: Service Market Landscape (2-3 minutes)**
Search queries:
```
"[service type]" consulting market [current year]
"[service type]" ROI case studies
"[service type]" implementation success stories
"[service type]" pricing benchmarks
```

Extract:
- Market demand for specific service
- Typical ROI or business impact
- Success stories and case studies
- Pricing benchmarks and engagement models

**Research Type 2: Pain Points Solved (2-3 minutes)**
Search queries:
```
"[pain point from service definition]" cost impact
"[pain point]" business consequences
companies struggling with "[pain point]"
"[pain point]" consultant solutions
```

Extract:
- Quantified cost of inaction (business impact data)
- Prevalence of pain point across industries
- Typical solutions and approaches
- ROI expectations from buyers

**Compile research findings:**
```yaml
service_research:
  service: "[Service name]"
  research_date: "[ISO8601 timestamp]"
  sources_consulted:
    - url: "[URL]"
      extracted: "[Key finding]"
  market_demand:
    level: "HIGH" | "MEDIUM" | "LOW"
    evidence: "[Evidence source]"
  typical_roi:
    metric: "[ROI metric]"
    range: "[Percentage or dollar value]"
  pain_point_impact:
    - pain_point: "[Pain point]"
      annual_cost: "[Cost estimate]"
      prevalence: "[Percentage or description]"
  competitive_differentiation:
    - "[Opportunity to differentiate]"
```

### 2.4 Research Quality Check

**Before proceeding to pitch deck generation:**

Validate research quality:
- [ ] At least 5 credible sources consulted (company websites, news, Glassdoor, LinkedIn, industry reports)
- [ ] Recent information (≤6 months old prioritized, note if using older data)
- [ ] Quantified data where possible (costs, percentages, market sizes)
- [ ] Decision-maker information current (if prospect mode)
- [ ] Pain points have supporting evidence (not just assumed)

**If research quality is LOW:**
```
⚠️ LIMITED RESEARCH DATA AVAILABLE

Research completed but data quality concerns:
- [Issue 1: e.g., "Limited recent news for [company]"]
- [Issue 2: e.g., "Market size data unavailable for [industry]"]

Proceeding with pitch deck generation using available data.
Pitch deck will note areas requiring manual validation before use.
```

Log research quality level: HIGH / MEDIUM / LOW

---

## Phase 3: Generate Pitch Deck Content (10-12 Slides)

### 3.1 Slide Structure Definition

**Standard 10-12 slide pitch deck structure:**

| Slide # | Title | Content Type | Primary Data Source |
|---------|-------|--------------|---------------------|
| 1 | Title | Identity, tagline, contact | Service definition |
| 2 | The Problem | Industry pain points | Research + service definition |
| 3 | The Cost of Inaction | Quantified business impact | Research + industry data |
| 4 | The Solution | Service offering overview | Service definition |
| 5 | How It Works | Methodology/approach | Service definition + evidence |
| 6 | Results & Proof | Case studies, metrics | Candidate profile + service definition |
| 7 | Why Me/Us | Credentials, differentiation | Service definition + candidate profile |
| 8 | Engagement Options | Pricing models, packages | Service definition rate card |
| 9 | Relevant Experience | Industry-specific examples | Candidate profile (filtered) |
| 10 | Next Steps | Clear CTA, timeline | Customizable based on mode |
| 11 | Q&A | Common objections | Generated based on service/industry |
| 12 | Appendix | Detailed credentials, references | Candidate profile |

**Slide count adjustment:**
- Service-focused mode: May reduce to 8-10 slides (single service deep-dive)
- Prospect-specific mode: Full 12 slides (comprehensive pitch)
- Industry-generic mode: 10-11 slides (Q&A may be optional)

### 3.2 Slide 1: Title Slide

```markdown
# [Consultant Name]
## [Tagline from service definition]

**Presented to:** [Prospect name] OR [Industry name] OR "General Capabilities"
**Date:** [Current date]

---

**Contact Information:**
- **Email:** [Extract from service definition or use placeholder]
- **LinkedIn:** [URL from service definition if available]
- **Website:** [URL from service definition if available]

**Credentials:** [Comma-separated credentials from service definition]

---

[Evidence: Service_Definition_[Date].md, lines X-Y]
```

**Customization by mode:**
- Prospect mode: "Presented to: [Company Name]"
- Industry mode: "Presented to: [Industry] Organizations"
- Service mode: "Presented to: Organizations Seeking [Service Type]"
- Default: "General Capabilities Overview"

### 3.3 Slide 2: The Problem

**Prospect-specific mode:**
```markdown
# The Challenge Facing [Company Name]

Based on our research, [Company Name] is navigating:

## [Primary Pain Point from research]
[Evidence from recent news, job postings, or industry reports]

## [Secondary Pain Point from prospect file]
[Evidence from prospect discovery research]

## [Tertiary Pain Point if available]
[Evidence from service definition alignment]

---

**Industry Context:**
[Industry] companies are experiencing similar challenges:
- [Industry trend 1 with statistic]
- [Industry trend 2 with statistic]

**Why This Matters Now:**
[Urgency factor from research - e.g., regulatory deadline, competitive pressure, technology shift]

---

**Sources:**
[List research sources with URLs]
```

**Industry-generic mode:**
```markdown
# The Challenge Facing [Industry]

[Industry] organizations are struggling with:

## [Industry Pain Point 1]
[Statistic or prevalence data from research]

**Business Impact:**
- [Impact metric 1]
- [Impact metric 2]

## [Industry Pain Point 2]
[Statistic or prevalence data]

**Business Impact:**
- [Impact metric 1]
- [Impact metric 2]

## [Industry Pain Point 3]
[Statistic or prevalence data]

**Business Impact:**
- [Impact metric 1]

---

**Industry Trends Driving These Challenges:**
- [Trend 1 from research]
- [Trend 2 from research]

---

**Sources:**
[List research sources]
```

**Service-focused mode:**
```markdown
# The Problem: [Pain Point Service Addresses]

Organizations across industries struggle with [pain point]:

## The Challenge
[Description of pain point with business context]

## Who This Affects
- **Industries:** [List industries from service ideal_client]
- **Company Types:** [Company sizes/types from service ideal_client]
- **Decision Makers:** [Roles from service ideal_client]

## Why Traditional Approaches Fail
- [Failure mode 1 from research or expertise]
- [Failure mode 2]
- [Failure mode 3]

---

**Prevalence:**
[Statistic about how common this problem is from research]

---

**Sources:**
[List research sources]
```

**Content rules:**
- Use ONLY pain points with supporting evidence (from research or prospect file)
- Quantify impact where possible (percentages, dollar costs, time waste)
- Avoid generic/unsupported claims (e.g., "Many companies struggle with...")
- Cite sources at bottom of slide
- Maximum 3 pain points per slide (focus over volume)

### 3.4 Slide 3: The Cost of Inaction

**Generate quantified business impact table:**

```markdown
# The Cost of Inaction

## What Happens If This Problem Persists?

| Impact Area | Annual Cost | Evidence |
|-------------|-------------|----------|
| [Impact Area 1 - e.g., "Inefficiency"] | [Quantified cost - e.g., "$500K-$1M in wasted hours"] | [Source] |
| [Impact Area 2 - e.g., "Missed Opportunities"] | [Quantified cost - e.g., "15-20% slower time-to-market"] | [Source] |
| [Impact Area 3 - e.g., "Competitive Disadvantage"] | [Quantified cost - e.g., "Market share erosion"] | [Source] |
| [Impact Area 4 - e.g., "Compliance Risk"] | [Quantified cost - e.g., "Regulatory penalties up to $X"] | [Source] |

**Total Estimated Impact:** [Range estimate if calculable]

---

## Beyond Direct Costs

**Opportunity Cost:**
- [Missed opportunity 1 - e.g., "Delayed product launches"]
- [Missed opportunity 2 - e.g., "Inability to scale operations"]

**Compounding Effects:**
- [How problem worsens over time]
- [Downstream consequences]

---

**The Window of Opportunity:**
[Why acting now is critical - urgency driver from research]

---

**Sources:**
[List sources for cost data - industry reports, case studies, research]
```

**Data sourcing rules:**
- Prioritize prospect-specific cost data (from research) if available
- Use industry benchmarks when specific data unavailable
- Clearly label estimates vs. verified data
- Cite all quantified claims
- If cost data unavailable, use qualitative impact description

**Quality check:**
- [ ] At least 2 impact areas have quantified costs
- [ ] Sources cited for all quantified claims
- [ ] Opportunity cost section addresses strategic impact
- [ ] Urgency/timing rationale provided

### 3.5 Slide 4: The Solution

**Multi-service pitch (default mode):**
```markdown
# The Solution: [Consultant Name] Service Portfolio

## [Service 1 Name]
[One-sentence description from service definition]

**Deliverables:** [Top 2-3 deliverables]

---

## [Service 2 Name]
[One-sentence description]

**Deliverables:** [Top 2-3 deliverables]

---

## [Service 3 Name]
[One-sentence description]

**Deliverables:** [Top 2-3 deliverables]

---

## How These Services Work Together
[Explain synergies or typical engagement progression]

**Example Client Journey:**
[Brief example of how services combine to solve comprehensive problem]

---

[Evidence: Service_Definition_[Date].md, Services section]
```

**Single-service pitch (service-focused mode):**
```markdown
# The Solution: [Service Name]

## What It Is
[Full description from service definition - 2-3 sentences]

## What You Get

### Deliverables
- [Deliverable 1 with brief explanation]
- [Deliverable 2 with brief explanation]
- [Deliverable 3 with brief explanation]
- [Deliverable 4 with brief explanation]
- [Deliverable 5 with brief explanation]

### Success Metrics
We measure success by:
- [Success metric 1 from service definition]
- [Success metric 2 from service definition]
- [Success metric 3 from service definition]

---

## Why This Approach Works
[Value justification from service definition pricing]

**Expected Outcomes:**
- [Outcome 1 with timeframe]
- [Outcome 2 with timeframe]

---

[Evidence: Service_Definition_[Date].md, [Service Name] section]
```

**Prospect-specific pitch:**
```markdown
# The Solution: Tailored [Service Name] for [Company Name]

## Addressing Your Specific Needs

| Your Challenge | Our Solution | Expected Outcome |
|----------------|--------------|------------------|
| [Pain point 1 from research] | [Service deliverable addressing it] | [Success metric] |
| [Pain point 2 from research] | [Service deliverable addressing it] | [Success metric] |
| [Pain point 3 from research] | [Service deliverable addressing it] | [Success metric] |

---

## The [Service Name] Package

**Core Deliverables:**
- [Deliverable 1 customized to prospect context]
- [Deliverable 2 customized to prospect context]
- [Deliverable 3 customized to prospect context]

**Timeline:** [Typical duration from service definition]

**Success Criteria:**
- [Metric 1 relevant to prospect]
- [Metric 2 relevant to prospect]

---

[Evidence: Service_Definition_[Date].md, [Service Name] section]
```

**Content rules:**
- Use exact deliverables from service definition (provenance critical)
- Match deliverables to pain points from Slide 2 where possible
- Avoid adding deliverables not in service definition (provenance violation)
- Include success metrics to set expectations
- Cite service definition as evidence source

### 3.6 Slide 5: How It Works

**Generate methodology/process flow:**

```markdown
# How It Works: Our Methodology

## [Phase 1 Name - e.g., "Discovery & Assessment"]
**Duration:** [Timeframe - e.g., "Week 1"]

**Activities:**
- [Activity 1]
- [Activity 2]
- [Activity 3]

**Deliverable:** [What client receives]

---

## [Phase 2 Name - e.g., "Design & Planning"]
**Duration:** [Timeframe - e.g., "Weeks 2-3"]

**Activities:**
- [Activity 1]
- [Activity 2]
- [Activity 3]

**Deliverable:** [What client receives]

---

## [Phase 3 Name - e.g., "Implementation"]
**Duration:** [Timeframe - e.g., "Weeks 4-6"]

**Activities:**
- [Activity 1]
- [Activity 2]
- [Activity 3]

**Deliverable:** [What client receives]

---

## [Phase 4 Name - e.g., "Optimization & Handoff"]
**Duration:** [Timeframe - e.g., "Weeks 7-8"]

**Activities:**
- [Activity 1]
- [Activity 2]

**Deliverable:** [What client receives]

---

## Client Involvement

**Your Role:**
- [Expectation 1 - e.g., "Weekly 1-hour status meetings"]
- [Expectation 2 - e.g., "SME access for 2-3 hours per week"]
- [Expectation 3 - e.g., "Final approval on deliverables"]

**Our Role:**
- [Commitment 1]
- [Commitment 2]

---

[Evidence: Service_Definition_[Date].md + candidate_profile.json work_history methodologies]
```

**Content generation approach:**

1. **Extract methodology from candidate profile:**
   - Search `work_history[].description` and `projects[].description` for methodology keywords
   - Look for process frameworks (Agile, SDLC, Design Thinking, etc.)
   - Extract typical project phases and timelines

2. **Match to service definition:**
   - Align extracted methodology with service deliverables
   - Sequence deliverables into logical phases
   - Estimate phase durations based on "typical_duration" from service pricing

3. **Validate provenance:**
   - Methodology must be supportable from work history (not fabricated)
   - Timelines must align with service definition typical duration
   - If no clear methodology in profile, use generic phased approach (Discovery → Design → Implementation → Handoff)

**Quality check:**
- [ ] 3-5 phases defined with clear start/end
- [ ] Each phase has activities and deliverables
- [ ] Total timeline aligns with service definition typical_duration
- [ ] Client expectations set (involvement level)
- [ ] Methodology is defensible from work history

### 3.7 Slide 6: Results & Proof

**THIS IS A CRITICAL PROVENANCE SLIDE - ALL CLAIMS MUST BE VALIDATED**

```markdown
# Results & Proof: Track Record of Success

## Quantified Achievements

### [Achievement 1 - Most Impressive]
**Result:** [Quantified metric from candidate profile]

**Context:** [Brief explanation of what was achieved]

**Relevance:** [Why this matters to prospect/industry]

[Evidence: candidate_profile.json → work_history[X].achievements[Y] → source file, lines A-B]

---

### [Achievement 2 - Industry-Relevant]
**Result:** [Quantified metric from candidate profile]

**Context:** [Brief explanation]

**Relevance:** [Connection to prospect's needs]

[Evidence: candidate_profile.json → work_history[X].achievements[Y] → source file, lines A-B]

---

### [Achievement 3 - Service-Specific]
**Result:** [Quantified metric from candidate profile]

**Context:** [Brief explanation]

**Relevance:** [How this proves capability for offered service]

[Evidence: candidate_profile.json → work_history[X].achievements[Y] → source file, lines A-B]

---

## Case Study: [Client Type from service definition case_studies if available]

**Challenge:** [What client faced]

**Solution:** [What was delivered]

**Results:**
- [Quantified result 1]
- [Quantified result 2]
- [Quantified result 3]

[Evidence: Service_Definition_[Date].md → [Service].case_studies[X] OR candidate_profile.json projects[X]]

---

## Credentials & Authority

- [Credential 1 from service definition] - [Brief context why it matters]
- [Credential 2] - [Context]
- [Thought leadership: Publication/speaking count] - [Context]

[Evidence: Service_Definition_[Date].md differentiation.certifications + candidate_profile.json certifications]
```

**PROVENANCE VALIDATION PROCESS FOR THIS SLIDE:**

**Step 1: Extract achievements from candidate profile**
```
Search candidate_profile.json:
  - work_history[].achievements[] where metrics field exists
  - projects[].outcomes where quantified results present
  - leadership_experience with max_budget, max_team_size, scope_indicators

Filter criteria:
  - Must have quantified metric (numbers, percentages, dollar values)
  - Must have evidence.file and evidence.lines fields
  - Prefer achievements relevant to target (prospect industry, service type)
```

**Step 2: Rank achievements by relevance**
```
Scoring criteria:
  - Industry match (10 points): Achievement from same industry as prospect/target
  - Service match (10 points): Achievement demonstrates service capability
  - Recency (5 points): Achievement within last 5 years
  - Scale/Impact (5 points): Large dollar value, team size, or percentage improvement

Select top 3-5 achievements with highest combined score.
```

**Step 3: Validate each achievement**
```
For each selected achievement:
  - Verify evidence.file exists in ResumeSourceFolder
  - Verify evidence.lines point to actual content
  - Confirm metric is quantified (not vague "improved performance")
  - Check metric has timeframe (avoid unbounded claims)
  - Validate mechanism described (not just outcome without "how")

If validation fails: EXCLUDE achievement, select next highest-ranked
```

**Step 4: Case study extraction (if available)**
```
Priority 1: Service definition case_studies[] for exact service being pitched
Priority 2: candidate_profile.json projects[] matching industry/service
Priority 3: Synthesize from multiple work_history achievements (mark as "Representative Example")

All case study elements must have provenance:
  - Challenge: From work history or project description
  - Solution: From deliverables in work history
  - Results: From achievements with evidence
```

**Step 5: Generate evidence citations**
```
For each achievement/claim, append evidence line:
  [Evidence: candidate_profile.json → work_history[2].achievements[0] → Experience/Company_Role.md, lines 45-47]

Format ensures traceability back to source documents.
```

**Quality check before finalizing slide:**
- [ ] All quantified metrics have evidence citations
- [ ] No unbounded metrics (all have timeframes)
- [ ] No superlatives without proof ("market-leading" needs benchmark)
- [ ] Credentials match Active certifications from profile
- [ ] Case study (if included) has complete provenance
- [ ] Relevance explained for each proof point (why it matters here)

**If provenance validation fails:**
```
⚠️ PROVENANCE VALIDATION CONCERNS - SLIDE 6

Some claims could not be fully validated:
- [Claim 1]: [Issue - e.g., "No source file found for metric"]
- [Claim 2]: [Issue - e.g., "Unbounded metric without timeframe"]

EXCLUDED from pitch deck:
- [Excluded claim 1]
- [Excluded claim 2]

Proceeding with validated proof points only.
Validation rate: [X]% ([Y] of [Z] claims validated)
```

**Target validation rate: >90% for HIGH confidence pitch**

### 3.8 Slide 7: Why Me/Us

```markdown
# Why [Consultant Name]?

## Unique Value Proposition

[2-3 sentence unique_value from service definition differentiation section]

[Evidence: Service_Definition_[Date].md differentiation.unique_value]

---

## Competitive Advantages

### [Competitive Advantage 1]
**What:** [Advantage description from service definition]

**Why It Matters:** [Impact description from service definition]

**Proof:** [Related proof point or credential from candidate profile]

[Evidence: Service_Definition_[Date].md → differentiation.competitive_advantages[0]]

---

### [Competitive Advantage 2]
**What:** [Advantage description]

**Why It Matters:** [Impact description]

**Proof:** [Related proof point or credential]

[Evidence: Service_Definition_[Date].md → differentiation.competitive_advantages[1]]

---

### [Competitive Advantage 3]
**What:** [Advantage description]

**Why It Matters:** [Impact description]

**Proof:** [Related proof point or credential]

[Evidence: Service_Definition_[Date].md → differentiation.competitive_advantages[2]]

---

## What Sets Us Apart

| Typical Consultant | [Consultant Name] |
|-------------------|-------------------|
| [Generic capability 1] | [Specific differentiation from service def] |
| [Generic capability 2] | [Specific differentiation from service def] |
| [Generic capability 3] | [Specific differentiation from service def] |

---

[Evidence: Service_Definition_[Date].md differentiation section + candidate_profile.json]
```

**Content generation rules:**
- Use competitive_advantages[] directly from service definition (pre-validated)
- Match advantages to proof points from candidate profile for credibility
- Comparison table: Left column is "table stakes", right column is differentiation
- Avoid creating new advantages not in service definition (provenance violation)
- Maximum 3-4 advantages (focus over volume)

### 3.9 Slide 8: Engagement Options

```markdown
# Engagement Options & Pricing

## Service Packages

### [Service 1 Name]
**Pricing Model:** [Pricing model from service definition]

**Investment Range:**
- **Standard Engagement:** [Target pricing from service definition]
- **Entry-Level/Pilot:** [Minimum pricing]
- **Complex/Premium:** [Premium pricing]

**Typical Duration:** [Duration from service definition]

**What's Included:** [Top 3-4 deliverables]

---

### [Service 2 Name]
[Same structure]

---

### [Service 3 Name]
[Same structure]

---

## Flexible Engagement Models

| Model | Best For | Pricing |
|-------|----------|---------|
| **Project-Based** | Defined scope, clear deliverables | [Fixed-price range from service] |
| **Retainer** | Ongoing advisory, recurring needs | [Monthly retainer from rate card] |
| **Hourly Advisory** | Ad-hoc support, small engagements | [Hourly rate from rate card] |
| **Value-Based Partnership** | Strategic initiatives with measurable ROI | Custom (negotiated based on value) |

[Evidence: Service_Definition_[Date].md rate_card + services[].pricing]

---

## Payment Terms

- **Deposit:** [Deposit percentage from payment_terms]
- **Invoicing:** [Invoice frequency from payment_terms]
- **Payment Due:** [Payment due terms]
- **Expenses:** [Expense policy from rate_card]

---

## Volume Discounts & Packages

[If multiple services purchased together or long-term retainer, note discount policy]

---

**All pricing in [Currency - CAD/USD]**
**Rates valid through [Date 6 months out]**

[Evidence: Service_Definition_[Date].md rate_card.payment_terms]
```

**Pricing presentation strategy:**

**Prospect-specific mode:**
- Lead with service most relevant to prospect's pain points
- Emphasize pricing model that matches prospect's likely preference (based on company size and engagement history from prospect file)
- If prospect file shows budget indicators, subtly position within their likely range

**Industry-generic mode:**
- Present all services with equal weight
- Show model flexibility (project, retainer, hourly)
- Emphasize typical engagement patterns for industry

**Service-focused mode:**
- Deep dive on single service pricing
- Show package options (pilot vs. full implementation)
- Provide ROI justification from value_justification field

**Content rules:**
- Use ONLY pricing from service definition (do not invent packages)
- If pricing ranges seem wide, explain factors (scope, complexity, timeline)
- Include payment terms for transparency
- Note currency and validity period
- Cite service definition as evidence source

### 3.10 Slide 9: Relevant Experience

**THIS SLIDE MUST BE FILTERED FOR RELEVANCE**

```markdown
# Relevant Experience

## [Industry/Domain] Expertise

[If prospect-specific or industry-generic mode, filter work history to relevant industry]

### [Role 1 - Most Relevant]
**Company:** [Company type - anonymized if needed]

**Industry:** [Industry from work_history]

**Scope:** [Scope indicators - budget, team size, geographic scope]

**Key Achievements:**
- [Achievement 1 relevant to prospect/industry]
- [Achievement 2 relevant to prospect/industry]

[Evidence: candidate_profile.json → work_history[X] → source file, lines A-B]

---

### [Role 2 - Industry-Relevant]
**Company:** [Company type]

**Industry:** [Industry]

**Scope:** [Scope]

**Key Achievements:**
- [Achievement 1]
- [Achievement 2]

[Evidence: candidate_profile.json → work_history[Y] → source file, lines A-B]

---

## Cross-Industry Impact

[If service-focused mode, show service capability across industries]

**Service Capability Demonstrated Across:**

| Industry | Project Type | Outcome |
|----------|--------------|---------|
| [Industry 1] | [Project from candidate profile] | [Quantified result] |
| [Industry 2] | [Project from candidate profile] | [Quantified result] |
| [Industry 3] | [Project from candidate profile] | [Quantified result] |

[Evidence: candidate_profile.json projects[]]

---

## Technical Depth

**Expert-Level Skills Relevant to This Engagement:**

[Filter technical_skills[] where proficiency_level = "Expert" AND relevant to service]

- [Skill 1] - [Years experience] - [Context from work history]
- [Skill 2] - [Years experience] - [Context]
- [Skill 3] - [Years experience] - [Context]

[Evidence: candidate_profile.json technical_skills[] + work_history evidence]
```

**Filtering logic:**

**Prospect-specific mode:**
```
Filter criteria:
  1. work_history[] WHERE industry matches prospect industry (exact or adjacent)
  2. If <2 roles match, include work_history with transferable achievements
  3. Limit to 2-3 most relevant roles
  4. Extract only achievements relevant to prospect's pain points
```

**Industry-generic mode:**
```
Filter criteria:
  1. work_history[] WHERE industry matches target industry
  2. Show 2-4 roles demonstrating industry depth
  3. Emphasize industry-specific outcomes and metrics
```

**Service-focused mode:**
```
Filter criteria:
  1. work_history[] and projects[] WHERE demonstrates service capability
  2. Show cross-industry application (proves versatility)
  3. Emphasize technical skills and methodologies used
  4. Include 3-5 examples from different contexts
```

**Default mode (no targeting):**
```
Filter criteria:
  1. Most recent 3-4 roles
  2. Highest-impact achievements (largest metrics, most impressive outcomes)
  3. Diverse industry representation
```

**Quality check:**
- [ ] All work history cited has evidence (file + lines)
- [ ] Achievements selected are relevant to pitch target
- [ ] No more than 3-4 roles (focus over volume)
- [ ] Technical skills shown are Expert or Proficient level
- [ ] Clear connection to service offerings

### 3.11 Slide 10: Next Steps

**Prospect-specific mode:**
```markdown
# Next Steps: Let's Move Forward

## Proposed Timeline

**Week 1: Initial Discovery**
- Kickoff meeting with [decision makers from prospect file]
- Stakeholder interviews (2-3 hours)
- Current state assessment

**Week 2: Proposal Development**
- Detailed scope and timeline
- Customized engagement plan
- Budget finalization

**Week 3-X: Engagement Launch**
- [Phase 1 of methodology from Slide 5]
- [Expected start date based on typical duration]

---

## What We Need From You

1. **Decision to Proceed:** [Date 2 weeks out]
2. **Stakeholder Availability:** [Roles needed, hours per week]
3. **Data/System Access:** [Any prerequisites for starting]

---

## How to Get Started

**Option 1: Schedule Discovery Call**
- **Duration:** 60 minutes
- **Participants:** [Your decision makers] + [Consultant name]
- **Agenda:** Validate fit, discuss scope, answer questions
- **Book:** [Calendar link or email]

**Option 2: Request Detailed Proposal**
- Email: [Email address]
- Include: [Scope requirements, timeline constraints, budget range]
- Turnaround: [Response time - e.g., "3-5 business days"]

**Option 3: [Entry point from prospect file - e.g., "Vendor Portal Registration"]**
- [Instructions specific to prospect's procurement process]

---

## Questions?

**Contact:**
- **Email:** [Email]
- **Phone:** [Phone if available]
- **LinkedIn:** [Profile URL]

**Response Time:** [Commitment - e.g., "Within 24 hours"]

---

[Evidence: Entry points from Client_Prospects/Prospects_*.md for this prospect]
```

**Industry-generic or service-focused mode:**
```markdown
# Next Steps: Start the Conversation

## Typical Engagement Timeline

**Phase 1: Discovery (1-2 weeks)**
- Initial consultation and needs assessment
- Stakeholder interviews
- Current state evaluation

**Phase 2: Proposal (1 week)**
- Detailed scope, timeline, and budget
- Customized engagement plan
- Contract and terms

**Phase 3: Kickoff ([Timeline from service definition])**
- Project launch
- [Methodology phase 1 from Slide 5]

---

## How to Engage

### 1. Schedule a Discovery Call
**Duration:** 30-60 minutes

**We'll Discuss:**
- Your specific challenges and goals
- Potential solutions and approaches
- Fit assessment and next steps

**Book:** [Calendar link or email]

---

### 2. Request a Custom Proposal
**Provide:**
- Brief description of challenge/need
- Desired timeline and outcomes
- Budget parameters (if known)

**Email:** [Email address]

**Turnaround:** 3-5 business days

---

### 3. Download Our Service Catalog
[If service definition has been published online or as PDF]

**Includes:**
- Detailed service descriptions
- Pricing and engagement models
- Case studies and references

**Download:** [URL or "Email to request"]

---

## Questions?

**Contact:**
- **Email:** [Email]
- **LinkedIn:** [LinkedIn URL]
- **Website:** [Website URL if available]

**We typically respond within 24 hours.**

---

## What Our Clients Say

[If thought_leadership.speaking_engagements or testimonials available from candidate profile, include 1-2 brief quotes]

[Evidence: candidate_profile.json thought_leadership OR "Testimonials available upon request"]
```

**Content customization rules:**
- Use entry points from prospect file if prospect-specific (highest accessibility entry point first)
- Timeline should align with service definition typical_duration
- Decision-maker names from prospect file if available (show you've done homework)
- Generic mode: Multiple engagement options (call, proposal, download)
- Include clear contact information
- Set expectations (response time, next steps, prerequisites)

### 3.12 Slide 11: Q&A / Common Objections (Optional)

**Generate based on common objections for industry/service:**

```markdown
# Common Questions

## "How do you handle [common concern]?"

**Answer:**
[Address concern with evidence from methodology or work history]

[Evidence: candidate_profile.json or service_definition]

---

## "What makes you different from [competitor type]?"

**Answer:**
[Reference Slide 7 competitive advantages, provide specific differentiation]

[Evidence: Service_Definition differentiation.competitive_advantages]

---

## "How do you ensure ROI?"

**Answer:**
[Reference success metrics from service definition + proof points from Slide 6]

[Evidence: Service_Definition services[].success_metrics + candidate_profile achievements]

---

## "What if the project scope changes?"

**Answer:**
[Describe change management process from payment_terms or methodology]

[Evidence: Service_Definition rate_card.payment_terms]

---

## "Do you offer references?"

**Answer:**
[Yes/No with explanation. If yes, note how to request. If no, point to case studies and public proof points]

[Evidence: candidate_profile.json thought_leadership.publications OR "Available upon request for qualified opportunities"]

---

## "What's your availability?"

**Answer:**
[Based on engagement_models.concurrent_clients and current capacity]

[Evidence: Service_Definition engagement_models.concurrent_clients]
```

**Objection generation logic:**

**Prospect-specific mode:**
- Address objections specific to company type (e.g., enterprise: "procurement process", startup: "budget constraints")
- Reference red flags or concerns from prospect research
- Anticipate questions based on entry point (e.g., vendor portal: "insurance requirements")

**Industry-generic mode:**
- Address common industry objections (research-based)
- Regulatory/compliance concerns if relevant
- Integration with existing systems/processes

**Service-focused mode:**
- Technical feasibility questions
- Implementation complexity concerns
- ROI validation requests

**Quality check:**
- [ ] 4-6 questions maximum (focused on top objections)
- [ ] Answers reference earlier slides or service definition (consistency)
- [ ] Evidence cited for all factual claims
- [ ] Tone is confident but not defensive

### 3.13 Slide 12: Appendix - Detailed Credentials

```markdown
# Appendix: Credentials & References

## Professional Certifications

[Extract from candidate_profile.json certifications[] WHERE status = "Active"]

| Certification | Issuing Organization | Year Obtained | Status |
|---------------|---------------------|---------------|--------|
| [Cert 1] | [Organization] | [Year] | Active |
| [Cert 2] | [Organization] | [Year] | Active |
| [Cert 3] | [Organization] | [Year] | Active |

[Evidence: candidate_profile.json certifications[]]

---

## Education

[Extract from candidate_profile.json education[]]

- **[Degree 1]** - [Institution] ([Year])
- **[Degree 2]** - [Institution] ([Year])

[Evidence: candidate_profile.json education[]]

---

## Thought Leadership

### Publications
[Extract from candidate_profile.json thought_leadership.publications[]]

- [Publication 1] - [Type: Peer-reviewed/Whitepaper/etc.] ([Year])
- [Publication 2] - [Type] ([Year])

### Speaking Engagements
[Extract from thought_leadership.speaking_engagements count or list]

- [Speaking engagement 1] - [Venue] ([Year])
- [Speaking engagement 2] - [Venue] ([Year])

### Frameworks & Methodologies Created
[Extract from thought_leadership.frameworks_created[]]

- [Framework 1] - [Brief description]

[Evidence: candidate_profile.json thought_leadership]

---

## Professional Associations

[If available from candidate profile or service definition]

- [Association 1] - [Membership level]
- [Association 2] - [Membership level]

---

## Technical Proficiencies

**Expert Level:**
[List all technical_skills[] WHERE proficiency_level = "Expert"]

**Proficient Level:**
[List all technical_skills[] WHERE proficiency_level = "Proficient", limit to top 5-7]

[Evidence: candidate_profile.json technical_skills[]]

---

## Awards & Recognition

[If available from candidate profile achievements or awards field]

- [Award 1] - [Organization] ([Year])
- [Award 2] - [Organization] ([Year])

[Evidence: candidate_profile.json OR "Available upon request"]

---

## References

**Professional References Available Upon Request**

[For qualified opportunities, we provide:]
- Client references from similar engagements
- Peer references from industry colleagues
- Academic/institutional references

**To request references:** Email [email address] with:
- Your organization and role
- Specific service/engagement under consideration
- Timeline for decision
```

**Content rules:**
- Include ONLY Active certifications (filter out Expired/Lapsed)
- Prioritize credentials relevant to service offerings
- If credential list is very long, categorize (Technical, Business, Industry-Specific)
- All data must come from candidate profile (no fabrication)
- Evidence citations for all sections

---

## Phase 4: Provenance Hardening - Validate All Claims

**THIS IS A CRITICAL QUALITY GATE - DO NOT SKIP**

### 4.1 Extract All Quantified Claims from Generated Pitch Deck

**Scan all slides for:**
- Quantified metrics (numbers, percentages, dollar values, timeframes)
- Credentials and certifications
- Superlatives ("leading", "top", "best", "transformational")
- Company names and client references
- Technologies and methodologies claimed
- Timeframes and durations
- ROI claims and value justifications

**Create claims inventory:**
```
Claim #1: "[Exact text from slide]"
  - Slide: [Number]
  - Type: [Metric/Credential/Superlative/ROI/etc.]
  - Source claimed: [Evidence citation in slide]

Claim #2: "[Exact text]"
  - Slide: [Number]
  - Type: [Type]
  - Source claimed: [Evidence citation]

[Continue for all claims...]
```

### 4.2 Validate Each Claim Against Candidate Profile

**For each claim in inventory:**

**Step 1: Locate evidence in candidate profile**
```
If claim type is METRIC or ACHIEVEMENT:
  - Search work_history[].achievements[] for matching metric
  - Search projects[].outcomes for matching result
  - Search leadership_experience for scope indicators

If claim type is CREDENTIAL:
  - Search certifications[] for exact certification name
  - Verify status = "Active" (not Expired)

If claim type is TECHNICAL SKILL:
  - Search technical_skills[] for skill name
  - Verify proficiency_level = "Expert" or "Proficient"

If claim type is THOUGHT LEADERSHIP:
  - Search thought_leadership.publications[]
  - Search thought_leadership.speaking_engagements
  - Search thought_leadership.frameworks_created[]
```

**Step 2: Verify evidence quality**
```
For located evidence, check:
  - evidence.file field exists and points to valid file
  - evidence.lines field exists
  - If file is in ResumeSourceFolder, verify it exists
  - If metric has timeframe (not unbounded)
  - If metric has mechanism (how it was achieved, not just outcome)
```

**Step 3: Classify validation result**
```
VALIDATED (HIGH confidence):
  - Evidence found with file + lines citation
  - Metric is quantified with timeframe
  - Mechanism described or clear from context
  - Source file verifiable

WEAK EVIDENCE (MEDIUM confidence):
  - Evidence found but missing file/lines
  - Metric without clear timeframe
  - Mechanism unclear
  - Inferred from multiple sources

UNSUBSTANTIATED (FAIL):
  - No evidence found in candidate profile
  - Evidence contradicts claim (e.g., claimed "Expert" but profile shows "Basic")
  - Superlative without benchmark
  - Certification shown as Expired
```

**Step 4: Create validation report**
```
Claim: "[Text]"
Validation: VALIDATED | WEAK | FAIL
Evidence: [File path + lines OR "None found"]
Confidence: HIGH | MEDIUM | LOW
Notes: [Any concerns or caveats]
```

### 4.3 Handle Validation Failures

**For each FAIL (unsubstantiated claim):**

**Option 1: Find alternative evidence**
- Search candidate profile more broadly for supporting data
- Check if claim can be reworded to match available evidence
- Example: "Leading provider" → "Delivered for [X] clients in [industry]" (if evidence exists)

**Option 2: Soften claim**
- Add qualifier: "Demonstrated expertise in..." instead of "Expert in..."
- Add evidence caveat: "Based on representative projects..." instead of absolute claim
- Reduce specificity: "Significant improvement" instead of "50% improvement" (if exact metric unavailable)

**Option 3: Remove claim**
- If no alternative evidence and claim is not essential, delete
- Mark slide for manual review before use
- Log removal in provenance report

**Prioritization for removal:**
1. Remove superlatives first (lowest value, highest risk)
2. Remove unbounded metrics second ("improved efficiency" without %)
3. Keep claims with partial evidence if critical to pitch (mark as "Estimated - Validation Required")

### 4.4 Generate Provenance Trail Section

**Add to end of pitch deck markdown:**

```markdown
---

# PROVENANCE TRAIL

**Generated:** [ISO8601 timestamp]
**Service Definition Source:** [File path + version]
**Candidate Profile Source:** [File path + extraction date]
**Research Conducted:** [Date + time budget]

---

## Claims Validation Summary

**Total Claims:** [Count of all quantified claims, credentials, achievements]
**Validated (HIGH confidence):** [Count] ([Percentage]%)
**Weak Evidence (MEDIUM confidence):** [Count] ([Percentage]%)
**Unsubstantiated (removed):** [Count] ([Percentage]%)

**Overall Validation Rate:** [Percentage]%

---

## Evidence Mapping

| Slide | Claim | Evidence Source | Confidence | Notes |
|-------|-------|-----------------|------------|-------|
| 6 | "[Achievement metric]" | candidate_profile.json → work_history[2].achievements[0] → Experience/Company_Role.md, lines 45-47 | HIGH | Verified |
| 6 | "[Another metric]" | candidate_profile.json → projects[1].outcomes → Projects/Project_Name.md, lines 12-15 | HIGH | Verified |
| 7 | "[Certification]" | candidate_profile.json → certifications[0] | HIGH | Active status verified |
| 9 | "[Technical skill]" | candidate_profile.json → technical_skills[5] | MEDIUM | Proficient level (not Expert) |
| [X] | "[Removed claim]" | None found | FAIL | Removed from final pitch |

[Continue for all significant claims...]

---

## Research Sources

### Prospect Research (if applicable)
[List all URLs consulted for prospect-specific research with dates]

- [Source 1 URL] - [What was extracted]
- [Source 2 URL] - [What was extracted]

### Industry Research (if applicable)
[List all industry research sources]

- [Source 1 URL] - [What was extracted]
- [Source 2 URL] - [What was extracted]

---

## Validation Notes

**HIGH Confidence Claims ([X]):**
All claims have complete provenance with file + line citations, quantified metrics with timeframes, and verifiable mechanisms.

**MEDIUM Confidence Claims ([Y]):**
[List any claims with weak evidence and explain limitation]

**Removed Claims ([Z]):**
[List removed claims and reason for removal]

---

## Pre-Use Checklist

Before presenting this pitch deck:

- [ ] Verify decision-maker information is current (if prospect-specific)
- [ ] Update any time-sensitive research (news, initiatives) if >30 days old
- [ ] Confirm pricing and availability with [consultant name]
- [ ] Review MEDIUM confidence claims and determine if acceptable risk
- [ ] Prepare detailed provenance documentation if client requests verification
- [ ] Check for any industry/prospect news in last 7 days (refresh research)

---

**Validation Completed:** [ISO8601 timestamp]
**Validated By:** /pitchdeck command v[version]
**Validation Standard:** >90% for HIGH confidence pitch, >80% acceptable, <80% requires review
```

### 4.5 Validation Quality Gate

**Determine if pitch deck meets quality standards:**

```
if validation_rate >= 90%:
    confidence = "HIGH"
    status = "READY FOR USE"

elif validation_rate >= 80%:
    confidence = "MEDIUM"
    status = "REVIEW RECOMMENDED"
    warning = "Some claims have weak evidence. Review MEDIUM confidence items before presenting."

else:  # validation_rate < 80%
    confidence = "LOW"
    status = "MAJOR REVISION REQUIRED"
    error = "Too many unsubstantiated claims. Pitch deck requires significant rework."
```

**Report validation results to user:**

```
✅ PROVENANCE VALIDATION COMPLETE

**Validation Rate:** [X]% ([Y] of [Z] claims validated)
**Confidence Level:** [HIGH/MEDIUM/LOW]
**Status:** [READY/REVIEW RECOMMENDED/REVISION REQUIRED]

**Claims Breakdown:**
- HIGH confidence: [Count] claims with complete evidence
- MEDIUM confidence: [Count] claims with partial evidence
- Removed: [Count] unsubstantiated claims

[If MEDIUM or LOW confidence, show warning/error message]

**Provenance Trail:** See end of pitch deck markdown for complete evidence mapping.
```

**If validation rate < 80%, stop and report issues:**
```
❌ PROVENANCE VALIDATION FAILED

Validation rate: [X]% (target: >80%)

**Issues Identified:**
1. [Issue 1 - e.g., "5 achievement metrics lack evidence in candidate profile"]
2. [Issue 2 - e.g., "2 certifications shown as Expired"]
3. [Issue 3 - e.g., "Unbounded metrics without timeframes"]

**Removed Claims:**
- [Claim 1 removed]
- [Claim 2 removed]

**Recommendations:**
1. Review candidate profile completeness (may need to run /assessjob with comprehensive resume)
2. Update service definition if missing recent credentials
3. Remove or soften claims without evidence
4. Consider generating pitch deck with fewer quantified claims (focus on qualitative differentiation)

Pitch deck saved with current validation status.
**WARNING:** Review manually before presenting to prospects.
```

---

## Phase 5: Generate Final Output File

### 5.1 Determine Output Filename

**Filename format:** `Pitch_[Target]_[Date].md`

**Target naming logic:**

**Prospect-specific mode (`--prospect`):**
```
Target = [Company name from prospect file, sanitized]
Example: Pitch_Hatch_20251203.md
```

**Industry-generic mode (`--industry`):**
```
Target = [Industry name, sanitized]
Example: Pitch_CommercialRealEstate_20251203.md
```

**Service-focused mode (`--service`):**
```
Target = [Service name abbreviated, sanitized]
Example: Pitch_AIAgentDevelopment_20251203.md
```

**Default mode (no targeting):**
```
Target = "GeneralCapabilities"
Example: Pitch_GeneralCapabilities_20251203.md
```

**Sanitization rules:**
- Remove spaces: "Commercial Real Estate" → "CommercialRealEstate"
- Remove special characters: "AI/ML Services" → "AIMLServices"
- Use PascalCase for readability
- Maximum 50 characters for target portion

**Date format:** YYYYMMDD (e.g., 20251203)

### 5.2 Create Output Directory

```
if not exists("Client_Prospects/"):
    create_directory("Client_Prospects/")
```

### 5.3 Construct Final Markdown File

**File structure:**

```markdown
---
pitch_type: [prospect-specific | industry-generic | service-focused | general]
target: "[Target name - company, industry, or service]"
generated: [ISO8601 timestamp]
consultant: "[Name from service definition]"
service_definition: "[Path to service definition file]"
candidate_profile: "[Path to candidate profile JSON]"
research_conducted: [true | false]
provenance_validated: true
validation_rate: [Percentage]
confidence_level: [HIGH | MEDIUM | LOW]
version: "1.0"
---

[Slide 1: Title]

---

[Slide 2: The Problem]

---

[Slide 3: The Cost of Inaction]

---

[Slide 4: The Solution]

---

[Slide 5: How It Works]

---

[Slide 6: Results & Proof]

---

[Slide 7: Why Me/Us]

---

[Slide 8: Engagement Options]

---

[Slide 9: Relevant Experience]

---

[Slide 10: Next Steps]

---

[Slide 11: Q&A - Optional]

---

[Slide 12: Appendix]

---

[PROVENANCE TRAIL section]

---

**End of Pitch Deck**

---

## Usage Notes

**Presentation Format:**
This markdown file is optimized for conversion to slide decks using:
- Pandoc: `pandoc Pitch_[Target]_[Date].md -o Pitch_[Target]_[Date].pptx -t pptx`
- Marp: Use `---` as slide separators for Marp presentation
- Manual: Copy slides to PowerPoint/Google Slides/Keynote

**Before Presenting:**
- Review Provenance Trail section at end of file
- Verify all MEDIUM confidence claims are acceptable
- Update time-sensitive research if deck is >30 days old
- Customize Slide 10 (Next Steps) with specific dates/contacts
- Remove Provenance Trail section if presenting externally (keep internal copy for reference)

**Customization:**
- Add prospect/company logo to title slide
- Adjust color scheme to match service brand
- Include relevant images or charts where helpful (references maintained in notes)
- Reorder slides if needed for specific audience (keep provenance trail updated)
```

### 5.4 Save Primary Output File

**Save markdown to:**
`Client_Prospects/Pitch_[Target]_[Date].md`

**File encoding:** UTF-8

**Line endings:** Unix (LF)

### 5.5 Optional: Convert to PowerPoint (if `--format=pptx`)

**If --format=pptx specified AND pandoc installed:**

```bash
# Create PPTX using pandoc
pandoc "Client_Prospects/Pitch_[Target]_[Date].md" \
  -o "Client_Prospects/Pitch_[Target]_[Date].pptx" \
  -t pptx \
  --slide-level=1

# Check if conversion succeeded
if [ $? -eq 0 ]; then
    echo "✅ PowerPoint file generated: Pitch_[Target]_[Date].pptx"
else
    echo "⚠️ PowerPoint conversion failed. Markdown file available."
fi
```

**Conversion notes:**
- Slide breaks on `---` separators
- Markdown formatting preserved (bullets, tables, bold, italic)
- Evidence citations appear as slide notes (not visible in presentation mode)
- Provenance Trail becomes final slide(s) - recommend hiding in presentation

**Output files:**
- Markdown: `Pitch_[Target]_[Date].md` (always generated)
- PowerPoint: `Pitch_[Target]_[Date].pptx` (if conversion successful)

---

## Phase 6: Generate Summary Report

### 6.1 Console Output Summary

```
✅ PITCH DECK GENERATION COMPLETE

**Output Files:**
- Markdown: Client_Prospects/Pitch_[Target]_[Date].md
[- PowerPoint: Client_Prospects/Pitch_[Target]_[Date].pptx]

**Pitch Deck Details:**
- **Target:** [Prospect/Industry/Service name]
- **Type:** [Prospect-specific/Industry-generic/Service-focused/General]
- **Slides:** [Count] slides
- **Research Conducted:** [Yes/No] ([X] minutes, [Y] sources)

**Content Summary:**
- **Services Featured:** [Service names from deck]
- **Key Proof Points:** [Count] quantified achievements
- **Competitive Advantages:** [Count] differentiators highlighted
- **Pricing Models:** [Models shown - e.g., "Project-based, Retainer, Hourly"]

**Provenance Validation:**
- **Validation Rate:** [X]% ([Y] of [Z] claims validated)
- **Confidence Level:** [HIGH/MEDIUM/LOW]
- **Status:** [READY FOR USE/REVIEW RECOMMENDED/REVISION REQUIRED]

[If MEDIUM confidence:]
⚠️ **Review Recommended:** [X] claims have weak evidence. See provenance trail for details.

[If LOW confidence:]
❌ **Major Revision Required:** Validation rate below 80%. Manual review critical before use.

**Research Highlights** (if research conducted):
- [Key finding 1 from research]
- [Key finding 2 from research]
- [Key finding 3 from research]

**Next Steps:**
1. Review pitch deck markdown in [file path]
2. Check Provenance Trail section for evidence citations
3. [If PPTX generated:] Open PowerPoint file and customize design/branding
4. [If prospect-specific:] Verify decision-maker info is current before outreach
5. [If research conducted:] Refresh research if presenting >30 days from now
6. Remove Provenance Trail section before external presentation (keep internal copy)

**Customization Recommendations:**
- Add [Company/Industry] logo to title slide
- Include relevant charts/images for visual impact
- Adjust slide order based on audience priorities
- Tailor Slide 10 (Next Steps) with specific dates and contacts

**Before Presenting:**
- [ ] Verify decision-maker information (if prospect-specific)
- [ ] Update time-sensitive research (if >30 days old)
- [ ] Review MEDIUM confidence claims (if any)
- [ ] Customize branding and design
- [ ] Prepare detailed provenance if client asks for verification
- [ ] Check for recent news about prospect/industry (last 7 days)

---

**Pitch Deck Ready for Review**
```

### 6.2 Error and Warning Summary

**If warnings or errors occurred during generation:**

```
⚠️ WARNINGS DURING GENERATION:

1. [Warning 1 - e.g., "No case studies found in service definition for [Service X]"]
   Impact: Slide 6 uses general achievements instead of service-specific case study
   Recommendation: Add case study to service definition and regenerate

2. [Warning 2 - e.g., "Decision-maker from prospect file may have changed roles"]
   Impact: Slide 10 references [Name] as [Title] - verification recommended
   Recommendation: Check LinkedIn before outreach

3. [Warning 3 - e.g., "Limited research data for [Industry] trends"]
   Impact: Slide 2 uses general industry knowledge instead of current data
   Recommendation: Conduct manual research for fresher insights
```

**If critical errors prevented full generation:**

```
❌ ERRORS DURING GENERATION:

1. [Error 1 - e.g., "Service [Name] not found in service definition"]
   Impact: Cannot generate service-focused pitch
   Resolution: Check service name spelling or use --service with available service

2. [Error 2 - e.g., "Candidate profile missing critical data (work history empty)"]
   Impact: Slides 6 and 9 have minimal content
   Resolution: Run /assessjob with comprehensive resume to generate complete profile

[Pitch deck generated with available data, but manual review required before use.]
```

---

## Error Handling

### Error 1: No Service Definition Found

```
❌ SERVICE DEFINITION REQUIRED

No service definition file found in Client_Prospects/

Before creating pitch decks, define your service offerings:
  /defineservices

This creates:
- Service catalog with offerings and deliverables
- Pricing structure and engagement models
- Competitive differentiation with proof points
- Target market definitions

Once you have a service definition, run /pitchdeck again.
```

**Stop execution.**

### Error 2: No Candidate Profile Found

```
❌ CANDIDATE PROFILE REQUIRED

No candidate profile found at:
  ResumeSourceFolder/.profile/candidate_profile.json

Pitch deck generation requires a current candidate profile for provenance validation.

Generate a profile by running:
  /assessjob <any-job-posting-file.md>

This automatically creates an optimized candidate profile with:
- Technical skills and domain expertise
- Work history and achievements with evidence
- Certifications and thought leadership
- Complete provenance trail for all claims

Once you have a candidate profile, run /pitchdeck again.
```

**Stop execution.**

### Error 3: Prospect Not Found

```
❌ PROSPECT NOT FOUND: [company-name]

No prospect file found matching "[company-name]".

Available prospects from your most recent /findclient run:
[List HIGH and MEDIUM priority prospects from most recent Prospects_*.md file]

Options:
1. Use available prospect: /pitchdeck --prospect=[available-name]
2. Run new prospect discovery: /findclient [filters]
3. Generate industry-generic pitch: /pitchdeck --industry=[industry]
4. Generate service-focused pitch: /pitchdeck --service=[service-name]
```

**Stop execution.**

### Error 4: Service Not Found

```
❌ SERVICE NOT FOUND: [service-name]

No service matching "[service-name]" in your service definition.

Available services:
1. [Service 1 name]
2. [Service 2 name]
3. [Service 3 name]

Usage: /pitchdeck --service="[exact or partial service name]"

Example: /pitchdeck --service="AI Agent Development"
```

**Stop execution.**

### Error 5: Validation Rate Too Low

```
❌ PROVENANCE VALIDATION FAILED

Validation rate: [X]% (target: >80%)

**Critical Issues:**
1. [Issue 1 - e.g., "8 achievement metrics lack evidence in candidate profile"]
2. [Issue 2 - e.g., "3 certifications shown as Expired, removed from deck"]
3. [Issue 3 - e.g., "5 unbounded metrics without timeframes, softened or removed"]

**Removed Claims:**
- [Claim 1]
- [Claim 2]
- [Claim 3]

**Recommendations:**
1. Review candidate profile completeness:
   - Run /assessjob with comprehensive resume to populate profile
   - Ensure all certifications are current (update profile if needed)

2. Update service definition:
   - Add case studies with complete evidence
   - Verify all proof points are current and defensible

3. Reduce quantified claims:
   - Focus on qualitative differentiation instead of metrics
   - Use comparative language ("improved", "enhanced") instead of specific percentages

Pitch deck saved with current validation status.

⚠️ WARNING: Manual review required before presenting to prospects.
          Validation rate below acceptable threshold.
```

**Continue execution but warn user.**

### Error 6: Research Timeout or Failure

```
⚠️ RESEARCH LIMITATIONS

Research phase encountered issues:
- [X] web searches failed or timed out
- [Y] sources returned limited/no data
- Total research time: [Z] minutes (target: 5-10)

**Impact:**
- Prospect-specific customization is limited
- Slide 2 (The Problem) uses generic pain points
- Slide 3 (Cost of Inaction) uses industry estimates, not prospect data
- Slide 10 (Next Steps) may have outdated decision-maker info

**Recommendations:**
1. Conduct manual research before presenting:
   - Recent news about prospect/industry
   - Current decision-maker verification
   - Fresh pain point evidence

2. Consider regenerating deck when research services are responsive

Pitch deck generated with available data.
Marked slides requiring manual validation in provenance trail.
```

**Continue execution but warn user and mark affected slides.**

### Error 7: Pandoc Not Found (if --format=pptx)

```
⚠️ PANDOC NOT FOUND - MARKDOWN OUTPUT ONLY

PowerPoint (pptx) output requires pandoc for conversion.

Install pandoc:
  /install-pandoc

Or install manually:
  - macOS: brew install pandoc
  - Ubuntu/Debian: apt-get install pandoc
  - Windows: Download from https://pandoc.org/installing.html

Proceeding with markdown output.
You can convert manually later:
  pandoc Pitch_[Target]_[Date].md -o Pitch_[Target]_[Date].pptx -t pptx
```

**Continue with markdown-only output.**

---

## Usage Examples

### Example 1: Prospect-Specific Pitch

```bash
/pitchdeck --prospect="Hatch"
```

**Behavior:**
- Load prospect file: `Client_Prospects/Prospects_*Hatch*.md`
- Conduct fresh research on Hatch (5-7 minutes)
- Tailor all slides to Hatch's specific pain points and context
- Include decision-maker names from prospect file in Slide 10
- Validate all claims against candidate profile
- Output: `Client_Prospects/Pitch_Hatch_20251203.md`

**Use Case:** High-priority prospect from /findclient, ready for direct outreach

---

### Example 2: Industry-Generic Pitch

```bash
/pitchdeck --industry="Commercial Real Estate"
```

**Behavior:**
- Research CRE industry trends and pain points (5-8 minutes)
- Generate pitch addressing common CRE challenges
- Show service portfolio relevant to CRE
- Include industry-specific proof points from candidate profile
- Generic Next Steps (no specific company targeting)
- Output: `Client_Prospects/Pitch_CommercialRealEstate_20251203.md`

**Use Case:** Networking events, conference presentations, general CRE prospect outreach

---

### Example 3: Service-Focused Deep Dive

```bash
/pitchdeck --service="AI Agent Development"
```

**Behavior:**
- Focus entire pitch on single service
- Research service market demand and ROI (4-6 minutes)
- Deep-dive on methodology for this service (Slide 5)
- Show cross-industry applications (Slide 9)
- Detailed pricing and package options (Slide 8)
- Output: `Client_Prospects/Pitch_AIAgentDevelopment_20251203.md`

**Use Case:** RFP response, service-specific inquiry, thought leadership presentation

---

### Example 4: General Capabilities Deck

```bash
/pitchdeck
```

**Behavior:**
- No specific targeting (prospect/industry/service)
- Comprehensive overview of all services
- Highlights top achievements regardless of industry
- Generic problem/solution framing
- Flexible Next Steps for various engagement types
- Output: `Client_Prospects/Pitch_GeneralCapabilities_20251203.md`

**Use Case:** Networking, initial conversations, website download, capability statement

---

### Example 5: PowerPoint Output

```bash
/pitchdeck --prospect="Hatch" --format=pptx
```

**Behavior:**
- Generate prospect-specific pitch as above
- Convert markdown to PowerPoint using pandoc
- Output files:
  - `Client_Prospects/Pitch_Hatch_20251203.md` (source)
  - `Client_Prospects/Pitch_Hatch_20251203.pptx` (presentation)

**Use Case:** Ready-to-present deck for in-person or virtual meetings

---

### Example 6: Multiple Formats Workflow

```bash
# Generate markdown first
/pitchdeck --industry="Infrastructure" --format=md

# Review and customize markdown
# Edit Client_Prospects/Pitch_Infrastructure_20251203.md

# Convert to PowerPoint manually when ready
# pandoc Client_Prospects/Pitch_Infrastructure_20251203.md -o Pitch_Infrastructure_20251203.pptx -t pptx
```

**Behavior:**
- Markdown-first workflow allows customization before conversion
- Preserve evidence trail in markdown, clean version in PPTX

**Use Case:** When significant customization is needed before presentation

---

## Important Notes

### Provenance Standards

- **90%+ validation rate:** HIGH confidence, ready for use without additional review
- **80-89% validation rate:** MEDIUM confidence, review weak evidence claims before presenting
- **<80% validation rate:** LOW confidence, major manual review required

**All claims must have evidence trail:**
- Quantified metrics → candidate profile achievements with file + lines
- Credentials → candidate profile certifications (Active status)
- Technical skills → candidate profile technical_skills (Expert/Proficient level)
- Case studies → service definition case_studies OR candidate profile projects
- Competitive advantages → service definition differentiation section

### Research Ethics

- **Public information only:** Web searches use publicly available sources (news, LinkedIn, Glassdoor, company websites)
- **No private data:** Never access confidential information or internal documents
- **Time-bounded:** Research limited to 5-10 minutes to ensure freshness while respecting rate limits
- **Source citation:** All research findings cited with URLs for transparency

### Customization After Generation

**Before presenting:**
1. Add company/prospect logo to title slide
2. Customize color scheme to match brand
3. Include relevant images or charts (stock photos, industry graphics)
4. Tailor Slide 10 with specific dates and contact details
5. Review and update time-sensitive information (if deck >30 days old)
6. Remove Provenance Trail section for external presentation (keep internal copy)

**Maintain provenance:**
- If adding new claims, manually add evidence citations
- If modifying metrics, verify against source documents
- If changing services/pricing, ensure alignment with service definition
- Keep markdown file as source of truth, export to PPTX as needed

### Deck Lifecycle

**Quarterly refresh recommended:**
```
Every 3 months:
- Refresh research (news, pain points, decision-makers)
- Update proof points with latest achievements
- Verify pricing and availability
- Check certifications are still Active
- Regenerate deck: /pitchdeck [same arguments]
```

**Version control:**
- Each generation creates new dated file (no overwrites)
- Compare versions to see how pitch has evolved
- Maintain provenance trail in all versions

### Integration with Other Commands

**Recommended workflow:**

1. **Define services:** `/defineservices` → Create service catalog
2. **Discover prospects:** `/findclient` → Identify target clients
3. **Generate pitch:** `/pitchdeck --prospect=[name]` → Create tailored pitch
4. **Follow up:** After initial pitch, use `/assessjob` for role-specific proposals

**Supporting commands:**
- `/ratecard` → Generate professional rate card to include as appendix or leave-behind
- `/idealjob` → Understand ideal roles to inform service positioning
- `/assessjob` → Keep candidate profile current for provenance validation

---

Now executing pitch deck generation...
