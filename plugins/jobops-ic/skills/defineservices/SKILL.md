---
description: Define independent contractor service offerings with pricing and positioning
disable-model-invocation: true
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup (then /jobops-ic:setup) to initialize your workspace.

Use `config.directories.contractor_root` for output paths in this skill.
Use `config.preferences.default_currency` for pricing if applicable.

## Templates

For each template used by this skill, resolve the full path as:

  {config.templates.base_dir}/{config.templates.active.<template_name>}/<filename>

Templates referenced by this skill: service_definition_schema

## Modes

- `--guided` (default): Interactive questions
- `--from-source`: Auto-generate by reading source markdown files directly
- `--update`: Refresh existing service definition

## Output

`{config.directories.contractor_root}/services/service_definition_[YYYYMMDD].md`

---

## Phase 1: Setup

### Parse Mode
- Default or `--guided`: Interactive mode
- `--from-source`: Automatic generation from source markdown files
- `--update`: Update existing

### Load Resources
1. **JSON Schema**: Read `{config.templates.base_dir}/{config.templates.active[template_name]}/service_definition_schema.json`
2. **Source files** (--from-source, --update):

   For service-catalog generation, read these source files:
   - `{config.directories.resume_source}/Technology/TechStack.md` — judge Expert-level skills in-context from depth/recency of WorkHistory mentions; each Expert skill maps to a Technical Implementation service
   - `{config.directories.resume_source}/WorkHistory/*.md` — extract leadership scope (team size, budget, P&L mentions); leadership signals map to Strategic Advisory services
   - `{config.directories.resume_source}/Industries/*.md` if present, otherwise infer from WorkHistory industries — domains with 5+ years of mentions map to Domain Expertise services
   - `{config.directories.resume_source}/Thought_Leadership/*.md` if present (publications, frameworks, awards) — these map to Thought Leadership services
   - `{config.directories.resume_source}/Preferences/Vision.md` — engagement preferences and pricing anchors; if absent, use market-rate formula (documented in the skill)

   Service identification is a judgment task done in the context of this skill invocation. Do NOT pre-compute service candidates from a global enum; judge each potential service against the source evidence.

   Do NOT load `candidate_profile.json` — removed in v2.2.0.

3. **Vision.md**: Read `{config.directories.resume_source}/Preferences/Vision.md` for pricing/engagement preferences
4. **Existing Definition** (--update only):
   - Find `{config.directories.contractor_root}/services/service_definition_*.md` files
   - If multiple, ask user to select
   - Parse current services and pricing

---

## Phase 2: Guided Mode (--guided)

### 2.1 Identity Questions
1. Name/Business Name
2. Tagline (one sentence unique value)
3. Key Credentials (P.Eng, CFA, certifications)
4. Years of Experience
5. Website URL, LinkedIn URL

### 2.2 Service Catalog (3-5 services)
**Per Service:**
- **Identification**: Name, Category, Description (2-3 sentences), Deliverables (3-5), Success Metrics
- **Ideal Client**: Target Industries, Company Size, Pain Points, Decision Makers
- **Pricing**: Model (Hourly/Daily/Fixed/Retainer/Value/Hybrid), Range (Min/Target/Premium), Duration, Value Justification
- **Case Study** (if available): Client type, Challenge, Solution, Quantified results

Repeat until user declines additional services.

### 2.3 Engagement Models
1. **Preferred**: Project-Based, Retainer, Fractional Executive, Hourly Advisory, Value-Based Partnership, Equity + Cash (select multiple)
2. **Acceptable**: (not preferred but will consider)
3. **Avoid**: (decline these)
4. **Parameters**: Minimum Engagement, Typical Duration, Concurrent Clients

### 2.4 Rate Card
**Hourly Rates**: Minimum, Target, Premium, Rush Multiplier (optional)
**Daily Rates**: Explain "daily = hourly x 8", collect Min/Target/Premium
**Validation**: Flag if daily != hourly x 8 (>10% variance)
**Retainers** (optional): Part-time, Half-time, Full-time, Overage rate
**Payment Terms**: Deposit %, Invoice Frequency, Payment Due, Late Fee, Expense Handling

### 2.5 Differentiation
- **UVP**: 2-3 sentences on unique qualifications
- **Competitive Advantages** (3-5): Specific advantage + why it matters
- **Proof Points** (3-7): Quantified achievements with context
- **Authority Builders**: Certifications, Publications, Speaking

### 2.6 Target Markets (Optional)
**Primary Markets** (per market): Industry, Segment, Decision-maker roles, Market size
**Secondary Markets**: Adjacent opportunities
**Avoid Markets**: Client types to decline
**Geographic Focus**: Regional preferences

---

## Phase 3: Source-Based Mode (--from-source)

### 3.1 Extract Identity
Read `{config.directories.resume_source}/WorkHistory/*.md` and `{config.directories.resume_source}/Technology/TechStack.md`:
- Name: From the most recent WorkHistory file header or author attribution
- Years: Count span of dated WorkHistory entries
- Credentials: From any certifications block in WorkHistory or TechStack files (Active status only) + education references
- Tagline: Top 2-3 Expert-level skills from TechStack + dominant industry from WorkHistory → "[Capability] for [domain]"

### 3.2 Identify Services (3-5)
**Technical Implementation**: Expert skills in `{config.directories.resume_source}/Technology/TechStack.md` cross-referenced with depth/recency of mentions in WorkHistory files; each Expert skill with recent project outcomes maps to a service
**Strategic Advisory**: Leadership scope extracted from WorkHistory — team size >5, budget ownership, P&L responsibility, or transformation programmes; each qualifying signal maps to a service
**Domain Expertise**: Industries with 5+ years of combined WorkHistory mentions (or from `{config.directories.resume_source}/Industries/*.md` if present) — each qualifying domain maps to a service
**Thought Leadership**: Publications, frameworks, and awards from `{config.directories.resume_source}/Thought_Leadership/*.md` if present — Peer-reviewed, Whitepaper, Conference presentations, and named frameworks each map to a service

**Per Service Extract**:
- Name: Capability + outcome (judged from source evidence)
- Category: Map to activity type (Technical Implementation / Strategic Advisory / Domain Expertise / Thought Leadership)
- Description: Synthesize from WorkHistory achievements + technologies — cite source file and line
- Deliverables: From project outcomes in WorkHistory
- Ideal Client: Industries from WorkHistory/Industries files, company size inferred from WorkHistory scope, pain points inferred from transformation narratives, decision-maker level inferred from stakeholder references

### 3.3 Generate Pricing
**From Vision.md** (if exists):
- Extract hourly rates (min/target/ideal) → rate_card.hourly
- Calculate daily = hourly x 8
- Map employment preferences → engagement models
- Extract "avoid" preferences

**If Vision.md missing**:
- Research market rates for service type + seniority
- Formula: Min = (salary / 2000) x 1.5, Target = Min x 1.3, Premium = Target x 1.5-2.0

### 3.4 Build Differentiation
**UVP**: Synthesize dominant domain + Expert technical skills + years span + track record from WorkHistory
**Competitive Advantages**: Cross-domain expertise (from Industries files or WorkHistory), proven scale (team size/budget from WorkHistory), active certifications, proprietary frameworks (from Thought_Leadership)
**Proof Points**: Quantified achievements with metrics from WorkHistory — cite file and line; publications + speaking count from Thought_Leadership
**Authority Builders**: Active certifications from WorkHistory/TechStack, Peer-reviewed/Whitepaper/Conference publications, speaking engagements from Thought_Leadership

### 3.5 Define Target Markets
**Primary**: Most frequent + recent industries from WorkHistory (or Industries files), inferred company segment from scope of roles, common decision-maker level from stakeholder references, research market size
**Avoid**: Vision.md avoid_industries + industries absent from work history
**Geographic**: Geographic scope patterns from WorkHistory locations; international if applicable

---

## Phase 4: Update Mode (--update)

### 4.1 Load Existing
Find `{config.directories.contractor_root}/services/service_definition_*.md` files. If multiple, ask user to select. Parse: services, pricing, engagement models, differentiation.

### 4.2 Identify Updates
Compare existing definition with current source files (WorkHistory, TechStack, Industries, Thought_Leadership):
- **New Services**: Recent WorkHistory entries (12 months), new certifications, new Expert-level skills in TechStack, new domain coverage
- **Pricing**: Compare with Vision.md, validate daily = hourly x 8, check market shifts
- **Proof Points**: Recent achievements from WorkHistory, new publications/speaking from Thought_Leadership, new case studies

### 4.3 Interactive Update
Present findings + ask:
```
NEW CAPABILITIES: [List]
PRICING: Current vs Vision.md recommendations
NEW PROOF POINTS: [List]

Update options:
1. Add new service
2. Update pricing
3. Add proof points/case studies
4. Update engagement models
5. Refresh all sections
6. Cancel
```
Proceed with targeted updates based on selection.

### 4.4 Version Control
- Keep original intact
- Create new file: `{config.directories.contractor_root}/services/service_definition_[NewDate].md`
- Reference previous version in metadata
- Increment version (1.0->1.1 minor, 1.0->2.0 major)

---

## Phase 5: Pricing Validation

### 5.1 Consistency Checks
**Hourly vs Daily**: For each tier, calculate expected_daily = hourly x 8. If variance >10%, flag:
```
PRICING INCONSISTENCY
Tier: [Min/Target/Premium]
Hourly: $X/hr | Daily stated: $Y/day | Expected: $Z/day
Variance: [%]% - Recommend alignment
```

**Ascending Tiers**: Verify min < target < premium (both hourly/daily). If not, refuse generation until corrected.

**Service vs Rate Card**: Verify service pricing within rate card ranges. Flag if exceeds premium (may be value-based) or below minimum (below-market).

### 5.2 Market Validation (--from-source only)
Search: "[Service type] consultant rates [year]", "[Domain] consulting hourly rate [region]", "[Seniority] [skill] contractor rates [year]"
- Flag if <50% market median (underpriced)
- Flag if >200% market 75th percentile (overpriced unless justified)
- Provide market context in output

---

## Phase 6: Generate Document

### 6.1 Output Structure
See `{config.templates.base_dir}/{config.templates.active[template_name]}/service_definition_schema.json` for complete structure.

**Markdown sections**:
1. **Header**: YAML frontmatter with the following fields:
   ```yaml
   consultant: <name>
   generated_on: <YYYY-MM-DD>
   version: <plugin-version>
   source_files_read:
     - {config.directories.resume_source}/Technology/TechStack.md
     - {config.directories.resume_source}/WorkHistory/*.md
     - ... (whichever files Phase 1 actually read)
   ```
2. **Identity**: Name, Tagline, Credentials, Experience
3. **Service Offerings** (per service): Category, Description, Deliverables, Ideal Client (Industries/Size/Pain Points/Decision Makers), Pricing (Model/Range/Duration/Value), Success Metrics, Case Study (if available), Evidence
4. **Engagement Models**: Preferred/Acceptable/Avoid, Parameters (Min Engagement/Duration/Concurrent Clients)
5. **Rate Card**: Hourly (Min/Target/Premium/Rush), Daily (calculated ~ hourly x 8), Retainers (optional: Part/Half/Full-time, Overage), Payment Terms (Deposit/Frequency/Due/Late Fee/Expenses)
6. **Competitive Differentiation**: UVP, Advantages (with evidence), Proof Points (with evidence), Authority Builders (Certifications/Publications/Speaking)
7. **Target Markets** (if defined): Primary (Industry/Segment/Decision Makers/Market Size), Secondary, Avoid, Geographic Focus
8. **Next Steps**: Immediate Actions, Discovery Questions, Proposal Framework
9. **Appendix**: Evidence Mapping table (Service/Claim | Evidence File | Lines | Key Quote)
10. **Version History**: Version log

### 6.2 Save Files
**Primary**: `{config.directories.contractor_root}/services/service_definition_[YYYYMMDD].md`

### 6.3 Summary Report
```
SERVICE DEFINITION GENERATED

Output: service_definition_[Date].md
Services: [X] | Pricing: $[min]-$[max]/hr, $[min]-$[max]/day
Competitive Advantages: [X] with evidence

Pricing Validation:
- Daily = hourly x 8 (within 10%)
- Ascending tiers (min < target < premium)
[Warnings if any]

Next Steps: Review accuracy, update LinkedIn, create portfolio, identify 10-20 prospects, use for SOW templates

Market Validation (--from-source): Rates align with [percentile] - [insights]
```

---

## Phase 7: Validation Checklists

### 7.1 Pricing
- [ ] Daily ~ hourly x 8 (+/-10% all tiers)
- [ ] Min < Target < Premium (both hourly/daily)
- [ ] Service pricing within rate card (or justified exceptions)
- [ ] Payment terms clear and enforceable
- [ ] Currency specified

### 7.2 Evidence (--from-source)
- [ ] Services → WorkHistory/TechStack/Thought_Leadership evidence cited
- [ ] Advantages → file + line citations
- [ ] Proof points → source achievements with file + line references
- [ ] Certifications Active only
- [ ] Publications/speaking verified (not inferred)

### 7.3 Market Realism
- [ ] Services differentiated (non-overlapping)
- [ ] Descriptions clear, jargon-free, client-value focused
- [ ] Ideal client profiles actionable
- [ ] Pricing defensible
- [ ] Target market reachable and sufficiently large

### 7.4 Completeness
- [ ] Identity + credentials
- [ ] 3-5 services (deliverables, pricing)
- [ ] Engagement models (preferred/acceptable/avoid)
- [ ] Complete rate card (hourly/daily/payment terms)
- [ ] Differentiation with proof points
- [ ] Evidence mapping (--from-source)
- [ ] Next steps

---

## Error Handling

**Missing Data**:
- Source files missing (--from-source): List which expected files are absent; fall back to guided mode if WorkHistory is unavailable
- Vision.md missing: Use market-based pricing + warning
- Insufficient work history: Fall back to guided mode

**Pricing Errors**:
- Daily != hourly x 8 (>10%): Report error, request correction
- Non-ascending rates: Refuse until fixed
- Unrealistic rates (<$50 or >$500/hr): Flag for confirmation

**Evidence Gaps**:
- Service without evidence: Mark [EVIDENCE NEEDED]
- Unverifiable proof: Exclude or mark "Self-reported - verification recommended"

**Update Conflicts**:
- Custom edits in existing definition: Warn --from-source may overwrite, recommend targeted updates

---

## Usage Examples

```bash
/defineservices              # Default guided mode
/defineservices --guided     # Explicit guided mode
/defineservices --from-source  # Auto-generate from source markdown files
/defineservices --update     # Update existing definition
```

---

## Key Requirements

- **Pricing Consistency**: Clients notice hourly/daily misalignment. Always validate.
- **Evidence-Based**: --from-source only includes verifiable claims cited to source files.
- **Market Research**: Validate pricing competitiveness (--from-source).
- **Confidentiality**: Anonymize case study client details.
- **Version Control**: Create new timestamped files, never overwrite.
- **Schema Compliance**: JSON must validate against service_definition_schema.json.

Now executing...
