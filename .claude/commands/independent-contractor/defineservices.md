---
description: Define independent contractor service offerings with pricing and positioning
argument-hint: [--guided|--from-profile|--update]
---

I'll help you create a comprehensive service definition for independent contractor positioning, including service catalog, pricing strategy, and competitive differentiation.

**Mode Selection:**
- `--guided` (default): Interactive guided mode with strategic questions
- `--from-profile`: Automatic generation from candidate profile and vision preferences
- `--update`: Update existing service definition file

**Output File**: `Client_Prospects/Service_Definition_[Date].md`

---

## Phase 1: Determine Operation Mode

### 1.1 Parse Command Arguments

Check for mode flags:
- If `$1` is `--guided` or not provided: Run guided interactive mode (default)
- If `$1` is `--from-profile`: Run automatic profile-based generation
- If `$1` is `--update`: Run update mode for existing service definition

### 1.2 Load Required Resources

**Load JSON Schema:**
Read `/workspaces/resumeoptimizer/.claude/templates/service_definition_schema.json` to understand the complete service definition structure.

**Load Candidate Profile (for --from-profile and --update modes):**
- Check for `/workspaces/resumeoptimizer/ResumeSourceFolder/.profile/candidate_profile.json`
- If exists and recent (≤7 days old), use directly
- If missing or stale, generate using resume-summarizer agent

**Load Vision Preferences (if available):**
- Read `/workspaces/resumeoptimizer/ResumeSourceFolder/Preferences/Vision.md` if it exists
- Extract employment preferences, compensation requirements, work arrangement preferences
- Use to inform pricing and engagement model recommendations

**For --update mode only:**
- Search `Client_Prospects/` for existing `Service_Definition_*.md` files
- If multiple found, ask user which to update
- Parse existing file to extract current services and pricing

---

## Phase 2: Guided Mode (--guided) - Interactive Service Definition

### 2.1 Consultant Identity & Positioning

Ask strategic questions to define consultant identity:

**Questions to ask:**
1. **Name/Business Name**: "What name will you use for consulting? (Personal name or business name)"
2. **Tagline**: "In one sentence, what unique value do you offer clients? (e.g., 'AI-powered real estate intelligence for institutional investors')"
3. **Key Credentials**: "What credentials establish your authority? (e.g., P.Eng, CFA, MBA, AWS Certified, etc.)"
4. **Years of Experience**: "How many years of professional experience do you have?"
5. **Online Presence**: "Website URL (if available)?" and "LinkedIn profile URL?"

### 2.2 Service Catalog Development

For each service offering (recommend 3-5 services):

**Service Identification Questions:**
1. **Service Name**: "What do you call this service? (e.g., 'AI Agent Development for Due Diligence')"
2. **Category**: "Which category? (Strategic Advisory / Technical Implementation / Training & Enablement / Research & Analysis / Project Management / Interim Leadership / Other)"
3. **Description**: "In 2-3 sentences, what is this service and why do clients need it?"
4. **Deliverables**: "What specific deliverables do clients receive? (List 3-5 items)"
5. **Success Metrics**: "How do you measure success for this service? (e.g., '50% time reduction', '95% accuracy')"

**Ideal Client Profiling:**
1. **Target Industries**: "Which industries need this service most? (e.g., Real Estate, Finance, Healthcare)"
2. **Company Size**: "What size companies are ideal? (Startup / SMB / Mid-Market / Enterprise)"
3. **Pain Points**: "What specific problems does this solve for clients?"
4. **Decision Makers**: "Who typically buys this service? (e.g., CTO, VP Engineering, Head of Operations)"

**Pricing Strategy:**
1. **Pricing Model**: "How do you price this? (Hourly / Daily / Fixed-Price / Retainer / Value-Based / Hybrid)"
2. **Price Range**: "What's your pricing range?"
   - "Minimum (entry-level clients/simple projects):"
   - "Target (standard engagement):"
   - "Premium (complex/high-value/rush):"
3. **Typical Duration**: "How long does a typical engagement last? (e.g., '4-6 weeks', '3-month retainer')"
4. **Value Justification**: "What ROI or business value justifies this pricing? (e.g., '10x ROI through automation')"

**Case Study (if available):**
Ask: "Do you have a real example of delivering this service?"
If yes, collect:
- Client type (anonymized, e.g., "Global REIT with $2B AUM")
- Challenge they faced
- Your solution
- Quantified results

Repeat for each service offering. Ask after each service: "Do you want to define another service? (yes/no)"

### 2.3 Engagement Models & Working Preferences

**Engagement Structure Questions:**
1. **Preferred Models**: "Which engagement models do you prefer? (Select all: Project-Based / Retainer / Fractional Executive / Hourly Advisory / Value-Based Partnership / Equity + Cash)"
2. **Acceptable Models**: "Which are acceptable but not preferred?"
3. **Avoid Models**: "Which engagement types do you want to decline? (e.g., 'Long-term full-time employment', 'Fixed-price unlimited scope')"
4. **Minimum Engagement**: "What's your minimum engagement size? (e.g., '$10,000 project minimum', '20 hours minimum')"
5. **Typical Duration**: "What's a typical engagement length? (e.g., '3-6 months', '10-40 hours per project')"
6. **Concurrent Clients**: "How many concurrent client engagements can you handle? (e.g., 3-5)"

### 2.4 Rate Card Development

**Standard Rate Structure:**

**Hourly Rates:**
1. "Minimum hourly rate (entry-level clients): (e.g., '$150/hr')"
2. "Target hourly rate (standard clients): (e.g., '$200/hr')"
3. "Premium hourly rate (complex/specialized work): (e.g., '$350/hr')"
4. "Rush multiplier for urgent work? (e.g., 1.5x for <2 week turnaround, or 'None')"

**Daily Rates:**
Explain: "Daily rates should typically be hourly × 8 for consistency"
1. "Minimum day rate: (should be ~minimum hourly × 8)"
2. "Target day rate: (should be ~target hourly × 8)"
3. "Premium day rate: (should be ~premium hourly × 8)"

**Validation**: Check that daily rates ≈ hourly rates × 8 (±10% acceptable)
If mismatch >10%, flag: "⚠️ Daily rate doesn't align with hourly rate (daily should be ~8× hourly). This may confuse clients. Recommend: [calculated value]"

**Retainer Options (optional):**
Ask: "Do you offer monthly retainers? (yes/no)"
If yes:
1. "Part-time retainer (e.g., '$8,000/month for 40 hours'):"
2. "Half-time retainer (e.g., '$16,000/month for 80 hours'):"
3. "Full-time retainer (e.g., '$32,000/month for 160 hours'):"
4. "Overage rate for hours beyond retainer (e.g., '$250/hr'):"

**Payment Terms:**
1. "Upfront deposit percentage? (e.g., 50%, or 'None')"
2. "Invoice frequency? (Weekly / Bi-Weekly / Monthly / Milestone-Based / Net-30)"
3. "Payment due? (e.g., 'Net 15', 'Upon receipt')"
4. "Late payment fee? (e.g., '2% per month', 'None')"
5. "How are expenses handled? (e.g., 'Billed separately', 'Included in rates')"

### 2.5 Competitive Differentiation

**Unique Value Proposition:**
Ask: "In 2-3 sentences, what makes you uniquely qualified to deliver these services? What can you do that most competitors cannot?"

**Competitive Advantages (collect 3-5):**
For each advantage:
1. "What's a specific competitive advantage? (e.g., 'Domain expertise: 15+ years real estate + AI engineering')"
2. "Why does this matter to clients? (e.g., 'I speak both real estate and tech languages fluently')"

**Proof Points (quantified achievements):**
Ask: "What quantified achievements prove your capabilities? (e.g., '$500M+ assets analyzed', '12 AI agents in production')"
Collect 3-7 proof points with context.

**Authority Builders:**
1. "Relevant certifications that differentiate you:"
2. "Published works or thought leadership:"
3. "Speaking engagements or conference presentations:"

### 2.6 Target Market Definition (Optional)

Ask: "Do you want to define specific target markets? (yes/no - helps focus business development)"

If yes, for each primary market:
1. "Target industry: (e.g., 'Commercial Real Estate')"
2. "Specific segment: (e.g., 'Institutional investors with $1B+ AUM')"
3. "Decision-maker roles to target: (e.g., 'CIO', 'Head of Acquisitions')"
4. "Estimated market size or opportunity:"

Then ask:
- "Secondary or adjacent markets to explore:"
- "Markets or client types to avoid:"
- "Geographic focus: (e.g., 'North America', 'EMEA')"

---

## Phase 3: Profile-Based Mode (--from-profile) - Automatic Generation

### 3.1 Extract Consultant Identity

From candidate profile JSON:
- **Name**: `candidate.name`
- **Years of experience**: `candidate.years_total_experience`
- **Key credentials**: Extract from `certifications` array (status: Active) and `education` array
- **Tagline**: Generate from combination of:
  - Top 2-3 technical skills from `technical_skills` (Expert level)
  - Primary domain expertise from `domain_expertise` (Expert level)
  - Format: "[Technical capability] for [domain/industry]"
  - Example: "AI-powered process automation for commercial real estate"

### 3.2 Analyze Capabilities for Service Catalog

**Identify Service Opportunities:**

Analyze work history, projects, and technical skills to identify 3-5 distinct service offerings:

**Technical Implementation Services:**
- If `technical_skills` contains Expert-level programming/cloud/analytics skills
- If `projects` shows technical delivery with quantified outcomes
- Extract technologies used, typical project scope, and measurable results

**Strategic Advisory Services:**
- If `leadership_experience.max_team_size` > 5 or `leadership_experience.max_budget` exists
- If `work_history` shows strategic decision-making or transformation projects
- Extract scope indicators: assets managed, stakeholder levels, geographic scope

**Domain Expertise Services:**
- If `domain_expertise` shows Expert or Proficient depth with 5+ years
- Cross-reference with achievements showing domain-specific impact
- Focus on intersection of technical skills + domain knowledge

**Thought Leadership Services:**
- If `thought_leadership.publications` exists (Peer-reviewed, Whitepaper, Conference)
- If `thought_leadership.frameworks_created` shows methodologies adopted by others
- Training, enablement, or advisory services based on published frameworks

**For each service identified:**
1. **Name**: Combine capability + target outcome (e.g., "AI Agent Development for Due Diligence")
2. **Category**: Map to primary activity (Strategic Advisory / Technical Implementation / etc.)
3. **Description**: Synthesize from achievements and technologies
4. **Deliverables**: Extract from project outcomes and typical deliverables
5. **Ideal Client**:
   - Industries: From `domain_expertise` and `work_history.industry`
   - Company size: From `work_history.company_size` patterns
   - Pain points: Infer from problems solved in achievements
   - Decision makers: From `work_history.scope_indicators.stakeholder_level`

### 3.3 Generate Pricing from Vision.md (if available)

If `ResumeSourceFolder/Preferences/Vision.md` exists:

**Extract from Vision:**
- Employment type preferences → map to engagement models
- Hourly rate preferences → use as rate card basis
- Compensation requirements → inform minimum pricing thresholds

**Hourly Rate Extraction:**
Look for hourly rate specifications in Vision.md:
- `minimum` hourly rate → use as rate_card.hourly.minimum
- `target` hourly rate → use as rate_card.hourly.target
- `ideal` hourly rate → use as rate_card.hourly.premium

**Calculate Daily Rates:**
- minimum daily = minimum hourly × 8
- target daily = target hourly × 8
- premium daily = premium hourly × 8

**If Vision.md missing or incomplete:**
Generate market-based pricing:
1. Research market rates for the service type and seniority level
2. Use formula:
   - Minimum = (salary target / 2000 hours) × 1.5 (covers overhead, taxes, benefits)
   - Target = Minimum × 1.3
   - Premium = Target × 1.5 to 2.0

**Engagement Models:**
- If Vision prefers "Consultant/Contractor" → prefer Project-Based, Retainer, Hourly Advisory
- If Vision includes "Entrepreneur" → include Value-Based Partnership, Equity + Cash
- Avoid models: Extract from Vision "avoid" preferences

### 3.4 Build Differentiation from Profile Evidence

**Unique Value Proposition:**
Synthesize from:
- Intersection of domain expertise + technical skills (rare combinations)
- Years of experience in specific domains
- Proven track record with quantified achievements

**Competitive Advantages:**
Extract from:
- `domain_expertise` with Expert depth + `technical_skills` Expert level = cross-domain advantage
- `leadership_experience.max_team_size` and `max_budget` = proven scale
- `certifications` (Active status) = credential authority
- `thought_leadership.frameworks_created` = proprietary methodologies

**Proof Points:**
Extract from:
- `work_history.achievements` with `metrics` field → quantified results
- `work_history.scope_indicators.assets_managed` + description → scale proof
- `projects.outcomes` → measurable project results
- `thought_leadership.publications` count + speaking_engagements count

**Authority Builders:**
- Certifications: From `certifications` array (Active only)
- Publications: From `thought_leadership.publications` (prioritize Peer-reviewed, Whitepaper, Conference)
- Speaking: From `thought_leadership.speaking_engagements` count

### 3.5 Define Target Market from History

**Primary Markets:**
- Industries: Extract from `work_history.industry` (most frequent + recent)
- Segment: Infer from `work_history.scope_indicators.assets_managed` and company_size patterns
- Decision makers: From `work_history.scope_indicators.stakeholder_level` (most common)
- Market size: Research based on identified industries

**Avoid Markets:**
- Extract from Vision.md "avoid_industries" if available
- Identify industries NOT in work history (lack of evidence)

**Geographic Focus:**
- Extract from `work_history.scope_indicators.geographic_scope` patterns
- If International experience exists, include broader geographic focus

---

## Phase 4: Update Mode (--update) - Refresh Existing Service Definition

### 4.1 Load Existing Service Definition

1. Search `Client_Prospects/` for `Service_Definition_*.md` files
2. If multiple found, list them with dates and ask user to select:
   ```
   Found multiple service definitions:
   1. Service_Definition_20251115.md (2024-11-15)
   2. Service_Definition_20251201.md (2024-12-01)

   Which service definition do you want to update? (1-2)
   ```
3. Parse the selected file to extract:
   - Current services catalog
   - Current pricing structure
   - Current engagement models
   - Current differentiation points

### 4.2 Identify Update Opportunities

Compare existing service definition with current candidate profile:

**Check for New Services:**
- Analyze recent work history (last 12 months) for new capabilities
- Check for new certifications added since last update
- Identify new technical skills or domain expertise gained

**Check for Pricing Updates:**
- Compare existing rate card with current Vision.md preferences
- Validate daily rate = hourly rate × 8 consistency
- Flag if market rates have shifted significantly

**Check for New Proof Points:**
- Scan recent achievements for quantified outcomes
- Check for new publications or speaking engagements
- Identify new case study opportunities from recent projects

### 4.3 Interactive Update Process

Present findings and ask:
```
I've analyzed your current service definition and identified potential updates:

NEW CAPABILITIES:
- [New service opportunity based on recent work]

PRICING:
- Current rates: $X-Y/hr, $A-B/day
- Vision preferences: $Z/hr target
- Recommendation: [Update/Keep current]

NEW PROOF POINTS:
- [New quantified achievement]
- [New certification or publication]

What would you like to update?
1. Add new service
2. Update pricing
3. Add new proof points / case studies
4. Update engagement models
5. Refresh all sections
6. Cancel update
```

Based on selection, proceed with targeted updates using guided or profile-based approach for modified sections only.

### 4.4 Version Control

When updating:
- Keep original file intact
- Create new file with updated timestamp: `Service_Definition_[NewDate].md`
- In metadata section, reference previous version
- Increment version number (e.g., 1.0 → 1.1 for minor updates, 1.0 → 2.0 for major changes)

---

## Phase 5: Pricing Validation & Consistency Checks

### 5.1 Validate Pricing Consistency

**Hourly vs Daily Rate Check:**
For each pricing tier (minimum/target/premium):
- Calculate expected_daily = hourly_rate × 8
- Compare to stated daily_rate
- If difference > 10%, flag inconsistency:
  ```
  ⚠️ PRICING INCONSISTENCY DETECTED

  Tier: [Minimum/Target/Premium]
  Hourly rate: $X/hr
  Daily rate stated: $Y/day
  Expected daily rate: $Z/day (hourly × 8)
  Variance: [percentage]%

  Recommendation: Align daily rate to $Z/day OR explain premium/discount for day-rate engagements
  ```

**Rate Card Ascending Check:**
Verify pricing tiers are ascending:
- minimum < target < premium (for both hourly and daily)
- If not ascending, flag error and refuse to generate until corrected

**Service Pricing vs Rate Card Check:**
For each service with hourly/daily pricing:
- Verify service pricing range falls within rate card ranges
- If service pricing exceeds premium rate, flag for review (may be value-based)
- If service pricing below minimum rate, flag as below-market concern

### 5.2 Market Validation (--from-profile mode only)

Conduct web research to validate pricing:

**Search queries:**
1. "[Primary service type] consultant rates [current year]"
2. "[Domain expertise] consulting hourly rate [geographic region]"
3. "[Seniority level] [technical skill] contractor rates [current year]"

**Validation checks:**
- If generated rates are <50% of market median → flag as potentially underpriced
- If generated rates are >200% of market 75th percentile → flag as potentially overpriced (unless justified by unique differentiation)
- Provide market context in final output

---

## Phase 6: Generate Service Definition Document

### 6.1 Create Markdown Output

Generate comprehensive markdown document at `Client_Prospects/Service_Definition_[Date].md`:

**Structure:**

```markdown
---
consultant: [Name]
generated_by: /defineservices [mode]
generated_on: [ISO8601 timestamp]
version: 1.0
source_profile: [path to candidate_profile.json if used]
source_vision: [path to Vision.md if used]
last_updated: [ISO8601 timestamp]
---

# [Consultant Name] - Independent Contractor Service Catalog

**Tagline**: [One-sentence value proposition]

**Credentials**: [Comma-separated credentials]

**Experience**: [X] years

---

## Service Offerings

### [Service 1 Name]

**Category**: [Category]

**Description**:
[2-3 sentence description with value proposition]

**Deliverables**:
- [Deliverable 1]
- [Deliverable 2]
- [Deliverable 3]

**Ideal Client**:
- **Industries**: [Industry 1], [Industry 2]
- **Company Size**: [Size range]
- **Pain Points**: [Problems solved]
- **Decision Makers**: [Typical buyer roles]

**Pricing**:
- **Model**: [Pricing model]
- **Range**: [Minimum] - [Premium]
- **Typical Duration**: [Duration]
- **Value Justification**: [ROI explanation]

**Success Metrics**:
- [Metric 1]
- [Metric 2]

**Case Study** (if available):
- **Client**: [Anonymized client type]
- **Challenge**: [Problem faced]
- **Solution**: [What was delivered]
- **Results**: [Quantified outcomes]

[Evidence: Source file, lines XX-YY]

---

[Repeat for each service]

---

## Engagement Models

**Preferred Models**:
- [Model 1]
- [Model 2]

**Acceptable Models**:
- [Model 3]
- [Model 4]

**Avoid**:
- [Declined engagement type 1]
- [Declined engagement type 2]

**Engagement Parameters**:
- **Minimum Engagement**: [Minimum size]
- **Typical Duration**: [Typical length]
- **Concurrent Clients**: [Maximum number]

---

## Rate Card

### Hourly Rates
- **Minimum**: $[X]/hr - [Context: entry-level clients, straightforward work]
- **Target**: $[Y]/hr - [Context: standard engagements]
- **Premium**: $[Z]/hr - [Context: complex/specialized/rush work]
- **Rush Multiplier**: [X]× for <[timeframe] turnaround (if applicable)

### Daily Rates
- **Minimum**: $[A]/day (≈ [X]hr × 8)
- **Target**: $[B]/day (≈ [Y]hr × 8)
- **Premium**: $[C]/day (≈ [Z]hr × 8)

### Monthly Retainers (if offered)
- **Part-time**: $[Amount]/month for [Hours] hours
- **Half-time**: $[Amount]/month for [Hours] hours
- **Full-time**: $[Amount]/month for [Hours] hours
- **Overage Rate**: $[Amount]/hr beyond retainer hours

### Payment Terms
- **Deposit**: [X]% upfront (if applicable)
- **Invoice Frequency**: [Frequency]
- **Payment Due**: [Terms]
- **Late Fee**: [Fee structure or 'None']
- **Expenses**: [How handled]

---

## Competitive Differentiation

### Unique Value Proposition
[2-3 sentences explaining what makes this consultant uniquely qualified and different from competitors]

### Competitive Advantages
1. **[Advantage 1]**: [Why it matters]
   [Evidence: Source, lines XX-YY]

2. **[Advantage 2]**: [Why it matters]
   [Evidence: Source, lines XX-YY]

3. **[Advantage 3]**: [Why it matters]
   [Evidence: Source, lines XX-YY]

### Proof Points
- [Quantified achievement 1] - [Context]
- [Quantified achievement 2] - [Context]
- [Quantified achievement 3] - [Context]

[Evidence citations]

### Authority Builders
**Certifications**:
- [Certification 1]
- [Certification 2]

**Publications**:
- [Publication 1]
- [Publication 2]

**Speaking**:
- [Speaking engagement 1]
- [Speaking engagement 2]

---

## Target Markets (if defined)

### Primary Markets

#### [Industry 1]
- **Segment**: [Specific segment description]
- **Decision Makers**: [Role 1], [Role 2], [Role 3]
- **Market Size**: [Estimate or "Research needed"]

#### [Industry 2]
- **Segment**: [Specific segment description]
- **Decision Makers**: [Role 1], [Role 2]
- **Market Size**: [Estimate or "Research needed"]

### Secondary Markets
- [Market 1]
- [Market 2]

### Avoid
- [Market to avoid 1] - [Reason]
- [Market to avoid 2] - [Reason]

### Geographic Focus
- [Region 1]
- [Region 2]

---

## Next Steps: Client Discovery & Business Development

### Immediate Actions
1. **Update LinkedIn profile** with consultant positioning and service offerings
2. **Create portfolio/case study deck** showcasing proof points and client results
3. **Identify 10-20 target prospects** in primary markets
4. **Draft outreach templates** aligned with service offerings

### Discovery Questions for Prospects
For each target market, prepare strategic discovery questions:
- [Question 1 aligned with pain point]
- [Question 2 to uncover budget/urgency]
- [Question 3 to identify decision process]

### Proposal Framework
Use this service definition as the foundation for:
- Statement of Work (SOW) templates
- Proposal deck structure
- Engagement agreement terms

---

## Appendix: Evidence Mapping

[If generated from profile, include evidence table showing which source files supported each service/proof point]

| Service/Claim | Evidence File | Lines | Key Quote |
|---------------|---------------|-------|-----------|
| [Service 1]   | [File path]   | XX-YY | [Quote]   |
| [Proof point 1] | [File path] | XX-YY | [Quote]   |

---

**Version History**:
- v1.0 ([Date]): Initial service definition created via [mode]
```

### 6.2 Save Output Files

**Primary Output:**
Save markdown document to: `Client_Prospects/Service_Definition_[YYYYMMDD].md`

Use current date for timestamp in filename.

**Secondary Output (for --from-profile mode only):**
Create `Client_Prospects/.cache/` directory if it doesn't exist.
Save structured JSON to: `Client_Prospects/.cache/Service_Definition_[YYYYMMDD].json`

Follow the JSON schema from `.claude/templates/service_definition_schema.json` exactly.

This enables:
- Programmatic access to service catalog
- Future automation of proposal generation
- Easy updates and version comparison

### 6.3 Generate Summary Report

After saving files, provide console summary:

```
✅ SERVICE DEFINITION GENERATED

**Output Files**:
- Markdown: Client_Prospects/Service_Definition_[Date].md
- JSON: Client_Prospects/.cache/Service_Definition_[Date].json (if --from-profile)

**Service Catalog**: [X] services defined
**Pricing**: $[min]-$[max]/hr, $[min]-$[max]/day
**Competitive Advantages**: [X] advantages with evidence

**Pricing Validation**:
✅ Daily rates consistent with hourly (within 10%)
✅ Rate tiers ascending (minimum < target < premium)
[⚠️ Any warnings or recommendations]

**Next Steps**:
1. Review service definitions and pricing for accuracy
2. Update LinkedIn profile with consultant positioning
3. Create portfolio deck with case studies and proof points
4. Identify 10-20 target prospects in primary markets
5. Use this definition as foundation for SOW templates

**Market Validation** (if --from-profile):
- Your rates align with [percentile] of market (based on [source])
- [Any market insights or recommendations]
```

---

## Phase 7: Validation & Quality Checks

### 7.1 Pricing Validation Checklist

Before finalizing, verify:
- [ ] Daily rate ≈ Hourly rate × 8 (within 10% for all tiers)
- [ ] Minimum < Target < Premium (ascending for both hourly and daily)
- [ ] Service-specific pricing falls within rate card ranges (or justified exceptions documented)
- [ ] Payment terms are clear and enforceable
- [ ] Currency specified for all rates

### 7.2 Evidence Verification (for --from-profile mode)

Verify all claims have evidence:
- [ ] Each service maps to specific work history or project evidence
- [ ] Each competitive advantage has file + line citation
- [ ] Each proof point references source achievement
- [ ] Certifications are Active status only
- [ ] Publications and speaking engagements are real (not inferred)

### 7.3 Market Realism Check

Validate service catalog is market-realistic:
- [ ] Services are clearly differentiated (not overlapping)
- [ ] Service descriptions avoid jargon and explain client value
- [ ] Ideal client profiles are specific enough to be actionable
- [ ] Pricing is defensible based on value delivered
- [ ] Target market is reachable and sufficiently large

### 7.4 Completeness Check

Ensure all required sections are present:
- [ ] Consultant identity with credentials
- [ ] 3-5 well-defined services with deliverables and pricing
- [ ] Engagement models (preferred/acceptable/avoid)
- [ ] Complete rate card (hourly, daily, payment terms)
- [ ] Competitive differentiation with proof points
- [ ] Evidence mapping (for profile-based generation)
- [ ] Next steps for client discovery

---

## Error Handling

**Missing Required Data:**
- If candidate profile missing and --from-profile used: Generate profile first using resume-summarizer agent
- If Vision.md missing and pricing needed: Use market-based pricing with warning
- If insufficient work history for services: Fall back to guided mode and ask user

**Pricing Inconsistencies:**
- If daily ≠ hourly × 8 beyond 10% variance: Report error, ask user to correct
- If rates not ascending: Refuse to generate until fixed
- If rates seem unrealistic (<$50/hr or >$500/hr): Flag for user confirmation

**Evidence Gaps:**
- If service claimed but no supporting evidence in profile: Mark as [EVIDENCE NEEDED] and recommend user validation
- If proof point can't be verified: Exclude or mark as "Self-reported - verification recommended"

**Update Conflicts:**
- If existing service definition has custom edits not from profile: Warn user that --from-profile update may overwrite customizations
- Recommend manual merge or targeted updates only

---

## Example Usage

```bash
# Guided interactive mode (default)
/defineservices

# Guided mode (explicit)
/defineservices --guided

# Automatic generation from profile
/defineservices --from-profile

# Update existing service definition
/defineservices --update
```

---

## Important Notes

- **Pricing Consistency is Critical**: Clients will notice if hourly vs daily rates don't align. Always validate.
- **Evidence-Based Claims**: In --from-profile mode, only include services and proof points with verifiable evidence from source materials.
- **Market Research**: Use web search in --from-profile mode to validate pricing is market-competitive.
- **Confidentiality**: When creating case studies from real projects, anonymize client details appropriately.
- **Version Control**: Always create new timestamped file rather than overwriting existing service definitions.
- **JSON Schema Compliance**: Generated JSON must validate against service_definition_schema.json.

Now executing service definition generation...
