---
description: Generate professional rate card document from service definition with pricing validation
argument-hint: [--format=md|pdf|html] [--currency=CAD|USD] [--include-justification]
---

Generate professional rate card from service definition with pricing validation.

**Arguments:**
- `$1`: Format (default: md): `md` | `pdf` | `html`
- `$2`: Currency (default: CAD): `CAD` | `USD`
- `$3`: `--include-justification` to add credentials/ROI section

---

## Phase 1: Validate Prerequisites

### 1.1 Find and Load Service Definition

Find most recent: `Client_Prospects/Service_Definition_*.md`

**If not found:**
```
❌ PREREQUISITE MISSING: Service Definition

Run /defineservices first to create service catalog and pricing.
```

**Load and validate:**
- rate_card (hourly/daily rates for minimum/target/premium tiers)
- consultant identity (name, tagline, credentials)
- services catalog (≥1 service)
- payment_terms

**If rate_card missing:**
```
❌ INVALID SERVICE DEFINITION: Missing Rate Card

Required structure:
## Rate Card
### Hourly Rates
- **Minimum**: $[X]/hr
- **Target**: $[Y]/hr
- **Premium**: $[Z]/hr
### Daily Rates
- **Minimum**: $[A]/day
- **Target**: $[B]/day
- **Premium**: $[C]/day

Fix with: /defineservices --update
```

### 1.2 Load Candidate Profile (if --include-justification)

Check: `ResumeSourceFolder/.profile/candidate_profile.json`
- Extract: certifications (Active), top 3-5 achievements, thought leadership
- If missing: Warn but continue without justification section

---

## Phase 2: Pricing Validation

### 2.1 Validation Checks

**Check 1: Daily Rate Consistency**
- Rule: `daily_rate ≈ hourly_rate × 8` (±10% tolerance)
- If fails: Offer auto-correction or keep with explanation

**Check 2: Rate Tier Progression**
- Rule: `minimum < target < premium` (both hourly and daily)
- If fails: Refuse generation, require manual fix

**Check 3: Premium Multiplier**
- Rule: `1.3x ≤ premium/target ≤ 2.0x`
- If outside: Advisory warning only, continue

**Check 4: Retainer Discounts (if applicable)**
- Rule: Discounts increase with commitment length
- Warn if discount <5% or >40%

### 2.2 Validation Summary

Display results:
```
✅ PRICING VALIDATION COMPLETE
✅ Daily rates align with hourly × 8
✅ Rate tiers ascending
✅ Premium multiplier reasonable (1.3x-2.0x)
⚠️ [Advisory warnings if any]
🔧 [Auto-corrections if any]
```

---

## Phase 3: Generate Rate Card

### 3.1 Parse Arguments

- Format: Default md, fallback to md if invalid
- Currency: Default CAD, assume CAD if unspecified in service definition
- Justification: Check for `--include-justification` flag

### 3.2 Currency Conversion (if needed)

If output currency ≠ service definition currency:
1. Web search for current CAD/USD exchange rate
2. Convert all monetary values
3. Add footnote with conversion rate and date

### 3.3 Markdown Rate Card Structure

**YAML Frontmatter:**
```yaml
consultant: "[Name]"
effective_date: [YYYY-MM-DD]
valid_until: [+6 months]
currency: [CAD|USD]
generated_on: [ISO8601]
pricing_validated: [true/false]
source_definition: "Client_Prospects/Service_Definition_[Date].md"
```

**Sections:**

1. **Header**: Consultant name, tagline, credentials
2. **Service Offerings**: Top 5 services with descriptions, deliverables
3. **Engagement Models & Rates**:
   - Hourly: Standard/Premium with use cases, minimum engagement
   - Daily: Standard/Premium (8-hour equivalents)
   - Retainer (if applicable): Part/half/full-time with effective rates and discounts
   - Project-Based (if applicable): Typical ranges by service type
4. **Volume & Term Discounts**: 3/6/12-month retainer, project thresholds, referrals
5. **Expense Policy**: Included (standard tools, local travel) vs. Billed separately (travel, specialized licenses)
6. **Terms & Conditions**:
   - Payment terms (retainers, projects, hourly, late fees)
   - Cancellation policy (30-day notice retainers, prorated projects, 48-hour session notice)
   - Confidentiality & IP (standard confidentiality, work product ownership, anonymized case studies)
   - Rate validity (6 months, annual review)
7. **Why These Rates?** (if --include-justification):
   - Credentials (certifications, years experience, publications, speaking)
   - Proven ROI (top 3-5 proof points with business impact)
   - Unique value proposition
8. **Contact & Next Steps**: Contact details, 4-step engagement process
9. **Footer**: Version, generation date, source file

Save to: `Client_Prospects/Rate_Card_[YYYYMMDD].md`

---

## Phase 4: HTML/PDF Generation (if requested)

### 4.1 HTML Template

**CSS Styling (embedded):**
- Professional B2B aesthetic: Blue (#2c5282) gradient header, clean tables
- Layout: 900px container, 2.5rem content padding
- Typography: System fonts, 1.6 line-height
- Tables: Header (#2c5282), alternating row backgrounds
- Print styles: Letter size, 0.75in margins, page-break-avoid for headings/tables

**HTML Structure:**
- Header: Gradient background with name/tagline/credentials
- Content sections: Parse markdown to HTML with styled classes
- Contact: Gray background callout box
- Footer: Metadata (version, date, source)

### 4.2 PDF Generation with Playwright

If format is `pdf`:

1. Save HTML to `/tmp/ratecard_temp_[timestamp].html`
2. Navigate: `mcp__playwright__browser_navigate` to `file://[tmp_html]`
3. Wait: 2 seconds for content load
4. Generate PDF via `browser_run_code`:
   ```javascript
   await page.pdf({
     path: 'Client_Prospects/Rate_Card_[YYYYMMDD].pdf',
     format: 'Letter',
     printBackground: true,
     margin: { top: '0.75in', bottom: '0.75in', left: '0.75in', right: '0.75in' }
   });
   ```
5. Close browser and cleanup temp file

**If PDF fails:** Fallback to HTML with manual conversion instructions

If format is `html`:
- Save to `Client_Prospects/Rate_Card_[YYYYMMDD].html` and skip Phase 5

---

## Phase 5: Final Report

### 5.1 Success Report

```
✅ PROFESSIONAL RATE CARD GENERATED

**Output:** Client_Prospects/Rate_Card_[YYYYMMDD].[format] ([XX] KB)
**Source:** Service_Definition_[Date].md

**Pricing:**
Hourly: $[min]-$[premium]/hr ([Currency])
Daily: $[min]-$[premium]/day
[Retainer: $[min]-$[max]/mo]

**Services:** [X] offerings ([List names])

**Validation:**
✅ Daily = hourly × 8 (±10%)
✅ Tiers ascending
✅ Premium multiplier: [X.X]x
[Auto-corrections: [count] | Warnings: [count]]

**Validity:** [Start] to [End] (6 months)
[Currency conversion note if applicable]
[Credentials section if --include-justification]

**Next Steps:**
1. Review for accuracy
2. Customize: Company-specific terms, volume discounts, special tiers
3. Distribute: Email PDF, host HTML, include in proposals
4. Maintain: Quarterly reviews, annual updates
5. Related: /defineservices --update, /proposaltemplate, /findclient

**Rate Card Summary:**
═══════════════════════════════════════════════════════════════
[CONSULTANT] | [CURRENCY] | VALID: [Start] to [End]
───────────────────────────────────────────────────────────────
Hourly:    $[min]/hr | $[target]/hr | $[premium]/hr
Daily:     $[min]/day | $[target]/day | $[premium]/day
[Retainer: Part-time $[X]/mo, Half-time $[Y]/mo, Full-time $[Z]/mo]
───────────────────────────────────────────────────────────────
PAYMENT: [Deposit]% upfront, Net [days] | MINIMUM: [hours]
═══════════════════════════════════════════════════════════════
```

---

## Error Handling

**Missing Prerequisites:**
- No service definition → Run /defineservices
- Missing rate_card section → Show required structure, suggest /defineservices --update
- Missing profile (with --include-justification) → Warn, continue without justification

**Validation Failures:**
- Invalid rate progression (min > target > premium) → Refuse generation, require manual fix
- Daily ≠ hourly × 8 (>10% variance) → Offer auto-correction (yes/no/cancel)

**Generation Failures:**
- PDF generation fails → Fallback to HTML with manual conversion instructions
- Currency conversion unavailable → Keep original, warn user
- Invalid format → Default to markdown

---

## Implementation Notes

**Pricing Validation Philosophy:**
- Auto-correct: Daily rate alignment, minor rounding
- Refuse: Rate progression violations, missing required fields
- Warn but continue: Premium multiplier outside 1.3x-2.0x, unusual retainer discounts

**Rate Card Best Practices:**
- 6-month validity, annual reviews
- Version control in YAML frontmatter
- Keep old versions for reference
- Professional B2B aesthetic: blue/gray, clean tables, generous whitespace

**Playwright PDF Pattern:** Same as /formatresume - HTML with embedded CSS → temp file → navigate → page.pdf() → cleanup

---

Now executing rate card generation with format `$1`, currency `$2`, and justification flag `$3`...
