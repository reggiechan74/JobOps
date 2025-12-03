---
description: Generate professional rate card document from service definition with pricing validation
argument-hint: [--format=md|pdf|html] [--currency=CAD|USD] [--include-justification]
---

I'll generate a professional rate card document from your service definition with comprehensive pricing validation.

**Arguments:**
- `$1`: Output format (optional): `md` | `pdf` | `html` (default: md)
- `$2`: Currency (optional): `CAD` | `USD` (default: CAD)
- `$3`: Additional flag (optional): `--include-justification` to add credentials/ROI section

**Example Usage:**
```bash
# Markdown rate card (default)
/ratecard

# PDF rate card with justification
/ratecard pdf CAD --include-justification

# HTML rate card in USD
/ratecard html USD

# Markdown with justification
/ratecard md CAD --include-justification
```

---

## Phase 1: Validate Prerequisites

### 1.1 Check for Service Definition

**Search for service definition file:**
```bash
find Client_Prospects/ -name "Service_Definition_*.md" -type f
```

**If no service definition found:**
```
❌ PREREQUISITE MISSING: Service Definition

No service definition file found in Client_Prospects/

**Required Action:**
Run /defineservices first to create your service catalog and pricing structure.

The service definition contains:
- Service offerings with pricing
- Rate card (hourly, daily, retainer rates)
- Competitive differentiation and credentials
- Engagement models and payment terms

Usage:
  /defineservices              # Interactive guided mode
  /defineservices --from-profile  # Auto-generate from your profile

After creating your service definition, run /ratecard again.
```

**If multiple service definitions found:**
- Select the most recent by date (extract from filename `Service_Definition_YYYYMMDD.md`)
- Display confirmation message:
  ```
  📄 Using service definition: Client_Prospects/Service_Definition_[Date].md
  Generated: [Date from file metadata]
  ```

### 1.2 Load Service Definition

Read the selected service definition file and parse:
- **Consultant identity**: name, tagline, credentials, years_experience
- **Services catalog**: All service offerings with names and descriptions
- **Rate card**: hourly, daily, retainer pricing tiers
- **Payment terms**: deposit, invoice frequency, payment due, late fees
- **Differentiation**: unique_value, competitive_advantages, proof_points
- **Engagement models**: preferred, acceptable, avoid

**Validate required sections exist:**
- [ ] rate_card section with hourly and daily rates
- [ ] consultant name and credentials
- [ ] services catalog (at least 1 service)
- [ ] payment_terms

**If rate_card section missing or incomplete:**
```
❌ INVALID SERVICE DEFINITION: Missing Rate Card

Your service definition at Client_Prospects/Service_Definition_[Date].md
is missing the required rate_card section.

**Required Structure:**

## Rate Card

### Hourly Rates
- **Minimum**: $[X]/hr - [Context]
- **Target**: $[Y]/hr - [Context]
- **Premium**: $[Z]/hr - [Context]

### Daily Rates
- **Minimum**: $[A]/day
- **Target**: $[B]/day
- **Premium**: $[C]/day

Please update your service definition with complete pricing information,
then run /ratecard again.

To regenerate your service definition:
  /defineservices --update
```

### 1.3 Load Candidate Profile (Optional - for justification section)

**If `--include-justification` flag present:**

Check for candidate profile at:
`ResumeSourceFolder/.profile/candidate_profile.json`

**If profile exists:**
- Load certifications (Active status only)
- Load top 3-5 quantified achievements from work_history
- Extract thought leadership (publications, speaking)
- Use for "Why These Rates?" section

**If profile missing:**
- Display warning but continue without justification section
- Recommend running `/assessjob` to generate profile

---

## Phase 2: Pricing Validation & Consistency Checks

### 2.1 Internal Pricing Validation

Run comprehensive validation checks on rate card pricing:

#### Check 1: Daily Rate Consistency

**Rule**: Daily rate should equal hourly rate × 8 (±10% tolerance)

For each tier (minimum, target, premium):
```python
def validate_daily_rate(hourly_rate, daily_rate, tier_name):
    expected_daily = hourly_rate * 8
    variance = abs(daily_rate - expected_daily) / expected_daily

    if variance > 0.10:
        return {
            "valid": False,
            "tier": tier_name,
            "hourly": hourly_rate,
            "daily_stated": daily_rate,
            "daily_expected": expected_daily,
            "variance_pct": variance * 100,
            "auto_corrected": expected_daily
        }
    return {"valid": True}
```

**If validation fails for any tier:**
```
⚠️ PRICING INCONSISTENCY DETECTED

Tier: [Minimum/Target/Premium]
Hourly rate: $[X]/hr
Daily rate stated: $[Y]/day
Expected daily rate: $[Z]/day (hourly × 8)
Variance: [XX.X]%

**Auto-Correction Available:**
Would you like to use the corrected daily rate ($[Z]/day) for the rate card?

1. Yes - Use corrected rate ($[Z]/day)
2. No - Keep stated rate ($[Y]/day) and add explanation
3. Cancel - Exit to manually fix service definition

Choice (1-3):
```

#### Check 2: Rate Tier Progression

**Rule**: Rates must be in ascending order: minimum < target < premium

For both hourly and daily rates:
```python
def validate_rate_progression(rates):
    errors = []

    # Hourly progression
    if not (rates.hourly.minimum < rates.hourly.target < rates.hourly.premium):
        errors.append({
            "type": "HOURLY_PROGRESSION",
            "message": "Hourly rates not in ascending order",
            "rates": {
                "minimum": rates.hourly.minimum,
                "target": rates.hourly.target,
                "premium": rates.hourly.premium
            }
        })

    # Daily progression
    if not (rates.daily.minimum < rates.daily.target < rates.daily.premium):
        errors.append({
            "type": "DAILY_PROGRESSION",
            "message": "Daily rates not in ascending order",
            "rates": {
                "minimum": rates.daily.minimum,
                "target": rates.daily.target,
                "premium": rates.daily.premium
            }
        })

    return errors
```

**If progression check fails:**
```
❌ INVALID RATE STRUCTURE

Your rate tiers are not in ascending order:
[Minimum: $X, Target: $Y, Premium: $Z]

Rate cards must follow: Minimum < Target < Premium

Please fix your service definition pricing structure:
  /defineservices --update

Then run /ratecard again.
```

#### Check 3: Premium Multiplier Reasonableness

**Rule**: Premium should be 1.3x to 2.0x of target rate

```python
def validate_premium_multiplier(target, premium, rate_type):
    multiplier = premium / target
    warnings = []

    if multiplier < 1.3:
        warnings.append({
            "type": "LOW_PREMIUM",
            "rate_type": rate_type,
            "multiplier": multiplier,
            "message": f"{rate_type} premium ({multiplier:.2f}x) may be too close to standard rate"
        })
    elif multiplier > 2.0:
        warnings.append({
            "type": "HIGH_PREMIUM",
            "rate_type": rate_type,
            "multiplier": multiplier,
            "message": f"{rate_type} premium ({multiplier:.2f}x) may be too high - could limit market"
        })

    return warnings
```

**If multiplier warnings exist:**
Display informational warnings but continue generation:
```
⚠️ PRICING ADVISORY

[Rate type] premium multiplier: [X.XX]x target rate

• Multiplier < 1.3x: Premium tier may not sufficiently differentiate for urgent/complex work
• Multiplier > 2.0x: Premium tier may price out potential clients

This is advisory only - rate card will be generated.
Consider adjusting if you want to optimize pricing strategy.
```

#### Check 4: Retainer Discount Validation (if retainers offered)

**Rule**: Retainer discounts should increase with commitment length

If `rate_card.retainer` exists:
```python
def validate_retainer_structure(retainer_rates, hourly_target):
    # Extract monthly fee and hours from strings like "$8,000/month for 40 hours"

    warnings = []
    for retainer_tier in ["part_time", "half_time", "full_time"]:
        if retainer_tier in retainer_rates:
            effective_rate = calculate_effective_hourly(retainer_rates[retainer_tier])
            discount_pct = ((hourly_target - effective_rate) / hourly_target) * 100

            if discount_pct < 5:
                warnings.append(f"{retainer_tier} offers minimal discount ({discount_pct:.1f}%) - may not incentivize retainers")
            elif discount_pct > 40:
                warnings.append(f"{retainer_tier} discount ({discount_pct:.1f}%) is very high - verify intentional")

    return warnings
```

### 2.2 Validation Summary Report

After all checks, generate validation summary:

```
✅ PRICING VALIDATION COMPLETE

**Consistency Checks:**
✅ Daily rates align with hourly × 8 (within 10% tolerance)
✅ Rate tiers in ascending order (minimum < target < premium)
✅ Premium multiplier within reasonable range (1.3x - 2.0x)
[✅/⚠️] Retainer discounts properly structured

**Advisory Warnings:** [count]
[List any non-blocking warnings]

**Auto-Corrections Applied:** [count]
[List any auto-corrected values used in rate card]

Proceeding with rate card generation...
```

---

## Phase 3: Generate Rate Card Document

### 3.1 Parse Command Arguments

**Extract format argument:**
- `$1`: Output format
  - If empty or "md": Use markdown
  - If "pdf": Generate PDF using Playwright
  - If "html": Generate styled HTML
  - If invalid: Default to markdown with warning

**Extract currency argument:**
- `$2`: Currency selection
  - If empty or "CAD": Use Canadian dollars (CAD)
  - If "USD": Use US dollars (USD)
  - If invalid: Default to CAD with warning

**Extract justification flag:**
- Check if `--include-justification` appears in any argument position
- If present: Include "Why These Rates?" section with credentials and ROI
- If absent: Omit justification section

### 3.2 Currency Conversion (if needed)

**If output currency differs from service definition currency:**

Detect service definition currency from rate values ($ = CAD/USD by default)

If conversion needed:
1. Use current CAD/USD exchange rate (web search for latest rate)
2. Convert all monetary values
3. Add footnote: "Rates converted from [SOURCE] to [TARGET] at [rate] on [date]"

**Note:** If service definition doesn't specify currency, assume CAD and display warning.

### 3.3 Generate Markdown Rate Card

Create comprehensive markdown document following this structure:

```markdown
---
consultant: "[Consultant Name]"
effective_date: [Today's date in YYYY-MM-DD]
valid_until: [6 months from today in YYYY-MM-DD]
currency: [CAD or USD]
generated_on: [ISO8601 timestamp]
generated_by: "/ratecard"
pricing_validated: [true/false]
validation_notes: "[Summary of validation results]"
source_definition: "Client_Prospects/Service_Definition_[Date].md"
---

# Professional Services Rate Card

## [Consultant Name]
**[Tagline from service definition]**

[Credentials list - e.g., "FRICS | CFA | MBA | P.Eng"]

---

## Service Offerings

[For each service in services catalog:]

### [Service Name]
[Service description from service definition]

**Typical Engagement:** [Pricing.typical_duration if available, else "Varies by project scope"]
**Deliverables:** [List key deliverables from service definition]

[Repeat for all services - limit to top 5 if more exist]

---

## Engagement Models & Rates

### Hourly Consulting

| Tier | Rate ([Currency]) | When to Use |
|------|-------------------|-------------|
| **Standard** | $[target hourly]/hour | Advisory calls, reviews, ad-hoc support, standard timeline projects |
| **Premium** | $[premium hourly]/hour | Urgent requests (<2 week turnaround), specialized expertise, C-suite engagement |

*Minimum engagement: [From engagement_models.minimum_engagement if available, else "4 hours"]*

[If rush_multiplier exists:]
**Rush Fee:** [X]× standard rate for turnaround < [timeframe]

---

### Daily Rate

| Tier | Rate ([Currency]) | When to Use |
|------|------------------|-------------|
| **Standard** | $[target daily]/day | On-site work, workshops, intensive sessions, full-day engagements |
| **Premium** | $[premium daily]/day | Urgent timeline, travel required, executive-level facilitation |

*Full day = 8 hours*

---

[If retainer rates exist in service definition:]

### Monthly Retainer

| Hours/Month | Monthly Fee ([Currency]) | Effective Rate | Discount |
|-------------|--------------------------|----------------|----------|
| [Hours from part_time] | $[Amount] | $[Calculated]/hr | [X]% off standard |
| [Hours from half_time] | $[Amount] | $[Calculated]/hr | [X]% off standard |
| [Hours from full_time] | $[Amount] | $[Calculated]/hr | [X]% off standard |

**Retainer Terms:**
- Unused hours do not roll over month-to-month
- Additional hours beyond retainer billed at standard hourly rate
- [Overage rate if specified: "Overage rate: $[amount]/hr"]
- [Payment terms: "Paid monthly in advance" or from payment_terms.invoice_frequency]

---

[If project-based pricing examples exist in service catalog:]

### Project-Based Pricing

| Project Type | Typical Range ([Currency]) | Duration |
|--------------|----------------------------|----------|
| [Service 1] | $[min range] - $[max range] | [Duration] |
| [Service 2] | $[min range] - $[max range] | [Duration] |
| [Service 3] | $[min range] - $[max range] | [Duration] |

*Project pricing based on defined scope and deliverables. Custom quotes available for unique requirements.*

---

## Volume & Term Discounts

[If engagement_models includes retainer commitments or service definition specifies discounts:]

| Commitment | Discount |
|------------|----------|
| 3-month retainer | 5% off monthly fee |
| 6-month retainer | 10% off monthly fee |
| 12-month retainer | 15% off monthly fee |
| Project > $50,000 | 5% discount |
| Project > $100,000 | 10% discount |
| Referral from existing client | 5% discount on first project |

*Discounts cannot be combined unless explicitly stated.*

---

## Expense Policy

### Included in Rates
- Standard software and collaboration tools
- Virtual meeting platforms (Zoom, Teams, etc.)
- Document preparation and delivery
- Reasonable local travel ([define region, e.g., "within Greater Toronto Area"])

### Billed Separately (at cost + 10% admin fee)
- Air travel and accommodation for on-site engagements
- Specialized software licenses (if required for project)
- Third-party data subscriptions or research services
- Printing and materials for large volumes (>100 pages)

[From rate_card.expenses if more specific policy exists]

---

## Terms & Conditions

### Payment Terms
[Extract from rate_card.payment_terms:]
- **Retainers:** [Invoice_frequency] in advance (e.g., "Monthly in advance")
- **Projects:** [Deposit_percentage]% upfront, remainder [payment schedule, e.g., "on completion" or "milestone-based"]
  - Projects > $50,000: Milestone-based payments (typically 50% upfront, 25% mid-project, 25% completion)
- **Hourly:** [Payment_due] from invoice date (e.g., "Net 15 days")
- **Late Payment Fee:** [Late_fee] (e.g., "2% per month" or "None")

### Cancellation Policy
- **Retainers:** 30 days written notice required
- **Projects:** Prorated refund minus work completed and non-recoverable costs
- **Scheduled Sessions:** 48 hours notice for rescheduling (within retainer) or cancellation (billed engagements)
- **Same-day cancellation:** Full session fee applies

### Confidentiality & IP
- All client information treated as confidential
- Work product delivered belongs to client upon full payment
- Consultant retains right to reference project (anonymized) in case studies with client approval
- [If IP rights are important from Vision.md preferences, add specific terms]

### Rate Validity
This rate card is valid from **[Effective_date]** to **[Valid_until]** ([6 months from today]).

Rates are subject to annual review and may be adjusted with 60 days notice to existing clients.

---

[ONLY IF --include-justification flag present:]

## Why These Rates?

### Credentials & Expertise
[List certifications from differentiation.certifications:]
- **[Certification 1]** - [Brief description of significance]
- **[Certification 2]** - [Brief description of significance]

**[Years_experience] years** of cross-industry professional experience

[If differentiation.publications exists:]
**Published Work:**
- [Publication 1]
- [Publication 2]
[Limit to top 3]

[If differentiation.speaking exists:]
**Industry Recognition:**
- [Speaking engagement 1]
- [Speaking engagement 2]
[Limit to top 3]

### Proven ROI
[Extract from differentiation.proof_points - limit to top 3-5 most impressive metrics:]
- **[Proof point 1]** - [Context explaining business impact]
- **[Proof point 2]** - [Context explaining business impact]
- **[Proof point 3]** - [Context explaining business impact]

### Unique Value Proposition
[From differentiation.unique_value - 2-3 sentences explaining competitive differentiation]

---

## Contact & Next Steps

**[Consultant Name]**
[Extract from consultant metadata if available: email, phone, LinkedIn]

**To discuss your specific needs:**
1. Review the service offerings above to identify your requirements
2. Email with a brief overview of your challenge or project
3. Schedule a complimentary 30-minute discovery call
4. Receive a custom proposal aligned with your budget and timeline

*For custom proposals, enterprise agreements, or questions about pricing, please reach out directly.*

---

**Rate Card Version:** [metadata.version, e.g., "1.0"]
**Generated:** [ISO8601 date] using /ratecard
**Source:** [Path to service definition file]
```

### 3.4 Save Markdown Rate Card

**Output location:**
`Client_Prospects/Rate_Card_[YYYYMMDD].[format]`

Use today's date for filename timestamp.

**If markdown format (default):**
Save markdown file directly.

**If HTML or PDF format:**
- Save markdown first as intermediate file
- Proceed to Phase 4 for HTML/PDF generation

---

## Phase 4: HTML/PDF Generation (if requested)

### 4.1 Convert Markdown to Styled HTML

If `$1` is `pdf` or `html`:

**Generate professional HTML document:**

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>[Consultant Name] - Professional Services Rate Card</title>
  <style>
    /* Reset and base styles */
    * { margin: 0; padding: 0; box-sizing: border-box; }

    body {
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
      line-height: 1.6;
      color: #2d3748;
      background: #f7fafc;
      padding: 2rem;
    }

    .rate-card-container {
      max-width: 900px;
      margin: 0 auto;
      background: white;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
      border-radius: 8px;
      overflow: hidden;
    }

    /* Header section */
    .rate-card-header {
      background: linear-gradient(135deg, #1a365d 0%, #2c5282 100%);
      color: white;
      padding: 3rem 2rem;
      text-align: center;
    }

    .consultant-name {
      font-size: 2.5rem;
      font-weight: 700;
      margin-bottom: 0.5rem;
      letter-spacing: -0.025em;
    }

    .tagline {
      font-size: 1.25rem;
      font-weight: 300;
      margin-bottom: 1rem;
      opacity: 0.95;
    }

    .credentials {
      font-size: 1rem;
      font-weight: 500;
      opacity: 0.9;
      letter-spacing: 0.05em;
    }

    /* Content section */
    .rate-card-content {
      padding: 2.5rem;
    }

    .section {
      margin-bottom: 2.5rem;
    }

    h2 {
      font-size: 1.75rem;
      color: #2c5282;
      border-bottom: 3px solid #4299e1;
      padding-bottom: 0.5rem;
      margin-bottom: 1.5rem;
      font-weight: 600;
    }

    h3 {
      font-size: 1.35rem;
      color: #2d3748;
      margin: 1.5rem 0 0.75rem;
      font-weight: 600;
    }

    /* Service offerings */
    .service-item {
      margin-bottom: 1.5rem;
      padding: 1rem;
      background: #f7fafc;
      border-left: 4px solid #4299e1;
      border-radius: 4px;
    }

    .service-name {
      font-size: 1.15rem;
      font-weight: 600;
      color: #2c5282;
      margin-bottom: 0.5rem;
    }

    .service-description {
      color: #4a5568;
      margin-bottom: 0.5rem;
    }

    .service-meta {
      font-size: 0.9rem;
      color: #718096;
      font-style: italic;
    }

    /* Tables */
    table {
      width: 100%;
      border-collapse: collapse;
      margin: 1rem 0 1.5rem;
      font-size: 0.95rem;
    }

    th {
      background-color: #2c5282;
      color: white;
      padding: 0.75rem;
      text-align: left;
      font-weight: 600;
      border: 1px solid #2c5282;
    }

    td {
      padding: 0.75rem;
      border: 1px solid #e2e8f0;
      background-color: white;
    }

    tr:nth-child(even) td {
      background-color: #f7fafc;
    }

    /* Tier-specific styling */
    .tier-standard {
      background-color: #f0f4f8 !important;
    }

    .tier-premium {
      background-color: #fefcbf !important;
      font-weight: 500;
    }

    /* Callout boxes */
    .note-box {
      background: #ebf8ff;
      border-left: 4px solid #4299e1;
      padding: 1rem;
      margin: 1rem 0;
      border-radius: 4px;
      font-size: 0.9rem;
    }

    .warning-box {
      background: #fffaf0;
      border-left: 4px solid #f6ad55;
      padding: 1rem;
      margin: 1rem 0;
      border-radius: 4px;
      font-size: 0.9rem;
    }

    /* Lists */
    ul {
      margin: 0.5rem 0 1rem 1.5rem;
    }

    li {
      margin-bottom: 0.35rem;
    }

    /* Contact section */
    .contact-section {
      background: #edf2f7;
      padding: 2rem;
      border-radius: 8px;
      margin-top: 2rem;
      text-align: center;
    }

    .contact-name {
      font-size: 1.5rem;
      font-weight: 600;
      color: #2c5282;
      margin-bottom: 1rem;
    }

    .contact-details {
      font-size: 1rem;
      color: #4a5568;
      margin-bottom: 1rem;
    }

    /* Footer metadata */
    .metadata-footer {
      margin-top: 2rem;
      padding-top: 1.5rem;
      border-top: 2px solid #e2e8f0;
      font-size: 0.85rem;
      color: #718096;
      text-align: center;
    }

    /* Print-specific styles for PDF generation */
    @media print {
      body {
        background: white;
        padding: 0;
      }

      .rate-card-container {
        box-shadow: none;
        border-radius: 0;
      }

      @page {
        margin: 0.75in 0.75in;
        size: letter;
      }

      h2, h3 {
        page-break-after: avoid;
      }

      table {
        page-break-inside: avoid;
      }

      .service-item {
        page-break-inside: avoid;
      }
    }
  </style>
</head>
<body>
  <div class="rate-card-container">
    <!-- Header -->
    <div class="rate-card-header">
      <div class="consultant-name">[Consultant Name]</div>
      <div class="tagline">[Tagline]</div>
      <div class="credentials">[Credentials separated by |]</div>
    </div>

    <!-- Content -->
    <div class="rate-card-content">
      <!-- Convert markdown sections to HTML here -->
      [Parsed markdown content with HTML structure]
    </div>
  </div>
</body>
</html>
```

**Markdown to HTML conversion:**
- Use standard markdown parser
- Apply CSS classes from stylesheet above
- Ensure tables render with proper styling
- Convert bullet lists to `<ul>` with styled `<li>`

### 4.2 Save HTML File (if html format requested)

**If `$1` is `html`:**
Save the generated HTML to:
`Client_Prospects/Rate_Card_[YYYYMMDD].html`

Display success message and skip to Phase 5.

### 4.3 Generate PDF with Playwright (if pdf format requested)

**If `$1` is `pdf`:**

Use Playwright MCP browser automation (same pattern as `/formatresume`):

**Step 1: Save temporary HTML file**
```bash
tmp_html="/tmp/ratecard_temp_$(date +%s).html"
# Write HTML content to tmp file
```

**Step 2: Navigate to HTML file with Playwright**
```
mcp__playwright__browser_navigate
  url: file://[tmp_html path]
```

**Step 3: Wait for content to load**
```
mcp__playwright__browser_wait_for
  time: 2  # seconds
```

**Step 4: Take screenshot for visual review (optional)**
```
mcp__playwright__browser_take_screenshot
  filename: Client_Prospects/.temp/ratecard_preview.png
  fullPage: true
```

**Step 5: Generate PDF**

Use Playwright's PDF generation with professional settings:
```
mcp__playwright__browser_take_screenshot equivalent for PDF
OR use browser_run_code with Playwright page.pdf() method:

await page.pdf({
  path: 'Client_Prospects/Rate_Card_[YYYYMMDD].pdf',
  format: 'Letter',
  printBackground: true,
  margin: {
    top: '0.75in',
    bottom: '0.75in',
    left: '0.75in',
    right: '0.75in'
  },
  displayHeaderFooter: false
});
```

**Step 6: Close browser and cleanup**
```
mcp__playwright__browser_close
```

Remove temporary HTML file:
```bash
rm -f /tmp/ratecard_temp_*.html
```

**PDF Generation Error Handling:**

If PDF generation fails:
```
⚠️ PDF GENERATION FAILED

Unable to generate PDF using Playwright browser automation.
Error: [Error message]

**Fallback Action:**
Saving as HTML instead: Client_Prospects/Rate_Card_[Date].html

You can manually convert to PDF by:
1. Opening the HTML file in Chrome/Edge
2. Print → Save as PDF
3. Use print settings: Letter size, 0.75in margins

Or save as Markdown: Client_Prospects/Rate_Card_[Date].md
```

Fallback to HTML format and save that instead.

---

## Phase 5: Generate Final Report & Summary

### 5.1 Verify Output File

Check that output file was created successfully:
- For markdown: `Client_Prospects/Rate_Card_[Date].md`
- For HTML: `Client_Prospects/Rate_Card_[Date].html`
- For PDF: `Client_Prospects/Rate_Card_[Date].pdf`

Get file size for reporting.

### 5.2 Generate Success Report

Display comprehensive summary:

```
✅ PROFESSIONAL RATE CARD GENERATED

**Output File:**
Client_Prospects/Rate_Card_[YYYYMMDD].[format]
Size: [XX] KB | Pages: [X] (for PDF)

**Source:**
Service Definition: Client_Prospects/Service_Definition_[Date].md

**Pricing Structure:**
Hourly Rates: $[min]/hr - $[premium]/hr ([Currency])
Daily Rates: $[min]/day - $[premium]/day ([Currency])
[If retainers exist:] Retainer: $[min]/mo - $[max]/mo

**Services Included:** [X] service offerings
[List service names]

**Validation Results:**
✅ Daily rates consistent with hourly × 8
✅ Rate tiers in ascending order
✅ Premium multiplier: [X.X]x ([reasonable/above range/below range])
[If retainers:] ✅ Retainer discounts structured progressively

[If auto-corrections applied:]
**Auto-Corrections Applied:** [count]
- [List corrected values]

[If warnings exist:]
**Advisory Warnings:** [count]
- [List warnings]

**Rate Card Validity:**
Effective: [Start date]
Valid Until: [End date] (6 months)

**Currency:** [CAD/USD]
[If converted:] *Converted from [SOURCE] at [rate] on [date]*

[If --include-justification used:]
**Credentials Section Included:**
- [X] certifications listed
- [X] proof points included
- Unique value proposition emphasized

---

**Next Steps:**

1. **Review the rate card** for accuracy and completeness
2. **Customize if needed:**
   - Add company-specific terms or policies
   - Adjust volume discount thresholds
   - Include special client tiers (nonprofit, startup, etc.)

3. **Distribution options:**
   - Email to prospects as PDF attachment
   - Host HTML version on your website
   - Include in proposal packages
   - Share via LinkedIn for transparency

4. **Keep it current:**
   - Review quarterly for market alignment
   - Update annually or when service offerings change
   - Version control: Increment metadata.version for changes

5. **Related commands:**
   - Update service definition: /defineservices --update
   - Create custom proposal: /proposaltemplate (coming soon)
   - Find target clients: /findclient (coming soon)

---

**Professional Tips:**

💡 **Rate Card Best Practices:**
- Be transparent: Clear pricing builds trust with professional buyers
- Show value tiers: Help clients self-select appropriate level
- Update regularly: Keep rates aligned with market and experience
- Version control: Date all rate cards clearly

📊 **Using Your Rate Card:**
- Qualify leads: Share early to filter budget-mismatched prospects
- Anchor pricing: Use in sales conversations to set expectations
- Simplify proposals: Reference rate card for standard engagements
- Demonstrate professionalism: Polished rate cards signal established practice

🎯 **When to customize:**
- Enterprise clients: Offer volume discounts beyond standard tiers
- Retainer clients: Consider loyalty discounts for renewals
- Startups/nonprofits: Create special tier with verification requirements
- Strategic partners: Equity + cash or revenue-share arrangements
```

### 5.3 Optional: Display Rate Card Summary Table

For quick reference, display condensed pricing table in console:

```
═══════════════════════════════════════════════════════════════
                    RATE CARD SUMMARY
═══════════════════════════════════════════════════════════════

CONSULTANT: [Name]
CURRENCY: [CAD/USD]
VALID: [Start date] to [End date]

───────────────────────────────────────────────────────────────
ENGAGEMENT TYPE     MINIMUM        TARGET         PREMIUM
───────────────────────────────────────────────────────────────
Hourly              $[XXX]/hr      $[XXX]/hr      $[XXX]/hr
Daily (8 hours)     $[X,XXX]/day   $[X,XXX]/day   $[X,XXX]/day
[If retainers:]
Part-Time Retainer  $[X,XXX]/mo ([XX] hrs) - [X]% discount
Half-Time Retainer  $[X,XXX]/mo ([XX] hrs) - [X]% discount
Full-Time Retainer  $[XX,XXX]/mo ([XXX] hrs) - [X]% discount
───────────────────────────────────────────────────────────────

PAYMENT TERMS: [Deposit]% upfront, [Payment due] net
LATE FEE: [Fee structure or "None"]
MINIMUM ENGAGEMENT: [Minimum]

═══════════════════════════════════════════════════════════════
```

---

## Error Handling & Edge Cases

### Missing Service Definition
```
❌ No service definition found
→ Action: Run /defineservices first
→ Recovery: Display usage instructions
```

### Incomplete Rate Card
```
❌ Service definition missing rate_card section
→ Action: Show required structure
→ Recovery: Suggest /defineservices --update
```

### Invalid Rate Progression
```
❌ Rates not in ascending order
→ Action: Refuse generation, display error
→ Recovery: User must fix service definition manually
```

### Pricing Inconsistency (Daily ≠ Hourly × 8)
```
⚠️ Daily rate doesn't match hourly × 8
→ Action: Offer auto-correction
→ Recovery: User can accept correction or keep original with explanation
```

### PDF Generation Failure
```
⚠️ Playwright PDF generation failed
→ Action: Fallback to HTML format
→ Recovery: Provide manual conversion instructions
```

### Currency Conversion Not Possible
```
⚠️ Unable to determine exchange rate
→ Action: Keep original currency, warn user
→ Recovery: User can manually update currency values
```

### Missing Candidate Profile (for justification)
```
⚠️ Candidate profile not found, --include-justification requested
→ Action: Generate rate card without justification section
→ Recovery: Suggest running /assessjob to create profile
```

### Invalid Format Argument
```
⚠️ Unknown format: [format]
→ Action: Default to markdown
→ Recovery: Display valid format options (md, pdf, html)
```

---

## Implementation Notes

### Playwright MCP Integration

For PDF generation, use the same pattern as `/formatresume`:
1. Generate self-contained HTML with embedded CSS
2. Save to temporary file
3. Use `mcp__playwright__browser_navigate` to load HTML
4. Use Playwright's `page.pdf()` method via `browser_run_code`
5. Close browser and cleanup temporary files

**Playwright code example:**
```javascript
await page.pdf({
  path: 'Client_Prospects/Rate_Card_20251202.pdf',
  format: 'Letter',
  printBackground: true,
  margin: { top: '0.75in', bottom: '0.75in', left: '0.75in', right: '0.75in' }
});
```

### Pricing Validation Philosophy

**Auto-correct when safe:**
- Daily rate alignment (hourly × 8 calculation)
- Minor rounding discrepancies

**Refuse generation when unsafe:**
- Rate tier progression violations (minimum > target > premium)
- Missing required fields

**Warn but continue:**
- Premium multiplier outside 1.3x-2.0x range
- Retainer discounts that seem unusual
- Large daily/hourly variance (>10% but user-confirmed)

### Rate Card Versioning

Rate cards should be versioned and dated:
- Include `effective_date` and `valid_until` in YAML frontmatter
- Recommend 6-month validity period
- Suggest annual rate reviews
- Keep old rate cards for reference (don't overwrite)

### HTML/PDF Styling Principles

**For professional B2B rate cards:**
- Clean, corporate aesthetic (blue/gray color scheme)
- Clear visual hierarchy with tables
- Generous whitespace
- Print-friendly (for PDF generation)
- No stock photos or decorative elements
- Focus on clarity and scanability

### Currency Handling

**Default behavior:**
- Assume CAD if not specified in service definition
- Support CAD and USD explicitly
- For other currencies: Keep as-is, add note

**Conversion:**
- Only convert if explicitly requested and exchange rate available
- Use real-time rate from web search
- Document conversion rate and date in footnote

---

## Testing Checklist

Before marking implementation complete, verify:

- [ ] Prerequisite validation works (missing service definition)
- [ ] Service definition parsing extracts all required fields
- [ ] Daily rate consistency check calculates correctly
- [ ] Rate tier progression validation catches errors
- [ ] Premium multiplier warnings trigger appropriately
- [ ] Markdown rate card generates with all sections
- [ ] HTML generation includes professional styling
- [ ] PDF generation via Playwright succeeds
- [ ] Currency argument is respected (CAD/USD)
- [ ] --include-justification flag adds credentials section
- [ ] Auto-correction for pricing inconsistencies offers choice
- [ ] Validation summary report displays correctly
- [ ] Error messages are clear and actionable
- [ ] Fallback to HTML works if PDF fails
- [ ] Output file is created in Client_Prospects/ directory
- [ ] File naming follows convention: Rate_Card_[YYYYMMDD].[ext]
- [ ] Final success report includes all key information

---

Now executing rate card generation with format `$1`, currency `$2`, and justification flag `$3`...
