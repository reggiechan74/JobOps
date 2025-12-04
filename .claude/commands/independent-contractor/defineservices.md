---
description: Define independent contractor service offerings with pricing and positioning
argument-hint: [--guided|--from-profile|--update]
---

**Modes:**
- `--guided` (default): Interactive questions
- `--from-profile`: Auto-generate from candidate profile + Vision.md
- `--update`: Refresh existing service definition

**Output**: `Client_Prospects/Service_Definition_[Date].md` (+ JSON for --from-profile)

---

## Phase 1: Setup

### Parse Mode
- Default or `--guided`: Interactive mode
- `--from-profile`: Automatic generation
- `--update`: Update existing

### Load Resources
1. **JSON Schema**: Read `.claude/templates/service_definition_schema.json`
2. **Candidate Profile** (--from-profile, --update):
   - Check `ResumeSourceFolder/.profile/candidate_profile.json` (≤7 days old)
   - Generate via resume-summarizer agent if missing/stale
3. **Vision.md**: Read `ResumeSourceFolder/Preferences/Vision.md` for pricing/engagement preferences
4. **Existing Definition** (--update only):
   - Find `Client_Prospects/Service_Definition_*.md` files
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
**Daily Rates**: Explain "daily = hourly × 8", collect Min/Target/Premium
**Validation**: Flag if daily ≠ hourly × 8 (>10% variance)
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

## Phase 3: Profile-Based Mode (--from-profile)

### 3.1 Extract Identity
From `candidate_profile.json`:
- Name: `candidate.name`
- Years: `candidate.years_total_experience`
- Credentials: `certifications` (Active) + `education`
- Tagline: Top 2-3 Expert `technical_skills` + Expert `domain_expertise` → "[Capability] for [domain]"

### 3.2 Identify Services (3-5)
**Technical Implementation**: Expert programming/cloud/analytics skills + quantified project outcomes
**Strategic Advisory**: Leadership (team >5 or budget exists) + transformation projects
**Domain Expertise**: Expert/Proficient `domain_expertise` (5+ years) + domain achievements
**Thought Leadership**: Publications (Peer-reviewed/Whitepaper/Conference) + frameworks created

**Per Service Extract**:
- Name: Capability + outcome
- Category: Map to activity type
- Description: Synthesize from achievements + technologies
- Deliverables: From project outcomes
- Ideal Client: Industries (`domain_expertise`, `work_history.industry`), Size (`company_size`), Pain points (inferred), Decision makers (`stakeholder_level`)

### 3.3 Generate Pricing
**From Vision.md** (if exists):
- Extract hourly rates (min/target/ideal) → rate_card.hourly
- Calculate daily = hourly × 8
- Map employment preferences → engagement models
- Extract "avoid" preferences

**If Vision.md missing**:
- Research market rates for service type + seniority
- Formula: Min = (salary / 2000) × 1.5, Target = Min × 1.3, Premium = Target × 1.5-2.0

### 3.4 Build Differentiation
**UVP**: Synthesize domain + technical skill intersection + years + track record
**Competitive Advantages**: Cross-domain expertise, proven scale (team size/budget), certifications (Active), proprietary frameworks
**Proof Points**: Achievements with metrics, assets managed, project outcomes, publications + speaking count
**Authority Builders**: Active certifications, Peer-reviewed/Whitepaper/Conference publications, speaking engagements

### 3.5 Define Target Markets
**Primary**: Most frequent + recent industries, inferred segment from assets/company_size, common stakeholder_level, research market size
**Avoid**: Vision.md avoid_industries + industries not in work history
**Geographic**: `geographic_scope` patterns, international if applicable

---

## Phase 4: Update Mode (--update)

### 4.1 Load Existing
Find `Service_Definition_*.md` files. If multiple, ask user to select. Parse: services, pricing, engagement models, differentiation.

### 4.2 Identify Updates
Compare existing with current candidate profile:
- **New Services**: Recent work history (12 months), new certifications, new technical skills/domain expertise
- **Pricing**: Compare with Vision.md, validate daily = hourly × 8, check market shifts
- **Proof Points**: Recent achievements, new publications/speaking, new case studies

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
- Create new file: `Service_Definition_[NewDate].md`
- Reference previous version in metadata
- Increment version (1.0→1.1 minor, 1.0→2.0 major)

---

## Phase 5: Pricing Validation

### 5.1 Consistency Checks
**Hourly vs Daily**: For each tier, calculate expected_daily = hourly × 8. If variance >10%, flag:
```
⚠️ PRICING INCONSISTENCY
Tier: [Min/Target/Premium]
Hourly: $X/hr | Daily stated: $Y/day | Expected: $Z/day
Variance: [%]% - Recommend alignment
```

**Ascending Tiers**: Verify min < target < premium (both hourly/daily). If not, refuse generation until corrected.

**Service vs Rate Card**: Verify service pricing within rate card ranges. Flag if exceeds premium (may be value-based) or below minimum (below-market).

### 5.2 Market Validation (--from-profile only)
Search: "[Service type] consultant rates [year]", "[Domain] consulting hourly rate [region]", "[Seniority] [skill] contractor rates [year]"
- Flag if <50% market median (underpriced)
- Flag if >200% market 75th percentile (overpriced unless justified)
- Provide market context in output

---

## Phase 6: Generate Document

### 6.1 Output Structure
See `.claude/templates/service_definition_schema.json` for complete structure.

**Markdown sections**:
1. **Header**: Metadata (consultant, generated_by, version, timestamps, source paths)
2. **Identity**: Name, Tagline, Credentials, Experience
3. **Service Offerings** (per service): Category, Description, Deliverables, Ideal Client (Industries/Size/Pain Points/Decision Makers), Pricing (Model/Range/Duration/Value), Success Metrics, Case Study (if available), Evidence
4. **Engagement Models**: Preferred/Acceptable/Avoid, Parameters (Min Engagement/Duration/Concurrent Clients)
5. **Rate Card**: Hourly (Min/Target/Premium/Rush), Daily (calculated ≈ hourly × 8), Retainers (optional: Part/Half/Full-time, Overage), Payment Terms (Deposit/Frequency/Due/Late Fee/Expenses)
6. **Competitive Differentiation**: UVP, Advantages (with evidence), Proof Points (with evidence), Authority Builders (Certifications/Publications/Speaking)
7. **Target Markets** (if defined): Primary (Industry/Segment/Decision Makers/Market Size), Secondary, Avoid, Geographic Focus
8. **Next Steps**: Immediate Actions, Discovery Questions, Proposal Framework
9. **Appendix**: Evidence Mapping table (Service/Claim | Evidence File | Lines | Key Quote)
10. **Version History**: Version log

### 6.2 Save Files
**Primary**: `Client_Prospects/Service_Definition_[YYYYMMDD].md`
**Secondary** (--from-profile only): `Client_Prospects/.cache/Service_Definition_[YYYYMMDD].json` (follow schema exactly)

### 6.3 Summary Report
```
✅ SERVICE DEFINITION GENERATED

Output: Service_Definition_[Date].md (+ JSON if --from-profile)
Services: [X] | Pricing: $[min]-$[max]/hr, $[min]-$[max]/day
Competitive Advantages: [X] with evidence

Pricing Validation:
✅ Daily = hourly × 8 (within 10%)
✅ Ascending tiers (min < target < premium)
[⚠️ Warnings if any]

Next Steps: Review accuracy, update LinkedIn, create portfolio, identify 10-20 prospects, use for SOW templates

Market Validation (--from-profile): Rates align with [percentile] - [insights]
```

---

## Phase 7: Validation Checklists

### 7.1 Pricing
- [ ] Daily ≈ hourly × 8 (±10% all tiers)
- [ ] Min < Target < Premium (both hourly/daily)
- [ ] Service pricing within rate card (or justified exceptions)
- [ ] Payment terms clear and enforceable
- [ ] Currency specified

### 7.2 Evidence (--from-profile)
- [ ] Services → work history/project evidence
- [ ] Advantages → file + line citations
- [ ] Proof points → source achievements
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
- [ ] Evidence mapping (profile-based)
- [ ] Next steps

---

## Error Handling

**Missing Data**:
- Profile missing (--from-profile): Generate via resume-summarizer
- Vision.md missing: Use market-based pricing + warning
- Insufficient work history: Fall back to guided mode

**Pricing Errors**:
- Daily ≠ hourly × 8 (>10%): Report error, request correction
- Non-ascending rates: Refuse until fixed
- Unrealistic rates (<$50 or >$500/hr): Flag for confirmation

**Evidence Gaps**:
- Service without evidence: Mark [EVIDENCE NEEDED]
- Unverifiable proof: Exclude or mark "Self-reported - verification recommended"

**Update Conflicts**:
- Custom edits in existing definition: Warn --from-profile may overwrite, recommend targeted updates

---

## Usage Examples

```bash
/defineservices              # Default guided mode
/defineservices --guided     # Explicit guided mode
/defineservices --from-profile  # Auto-generate from profile
/defineservices --update     # Update existing definition
```

---

## Key Requirements

- **Pricing Consistency**: Clients notice hourly/daily misalignment. Always validate.
- **Evidence-Based**: --from-profile only includes verifiable claims.
- **Market Research**: Validate pricing competitiveness (--from-profile).
- **Confidentiality**: Anonymize case study client details.
- **Version Control**: Create new timestamped files, never overwrite.
- **Schema Compliance**: JSON must validate against service_definition_schema.json.

Now executing...
