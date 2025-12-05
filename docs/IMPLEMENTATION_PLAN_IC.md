# Implementation Plan: Independent Contractor Landing Pages

## Summary

This plan details the implementation of 7 new HTML files to add Independent Contractor capabilities to the JobOps landing page website. The implementation follows existing patterns established by the OSINT Intelligence feature (overview page with linked subpages).

## Key Decisions Made

1. **File Structure**: Following the OSINT pattern with `docs/arsenal/independent-contractor.html` as overview and `docs/arsenal/contractor/*.html` for subpages
2. **Navigation Flow**: Linear navigation through commands in logical workflow order (defineservices -> ratecard -> findclient -> pitchdeck -> proposaltemplate)
3. **Icons**: Selected thematic emoji that reflect each command's purpose
4. **Theme**: "Not everyone wants a boss. Some people want clients." - positioning independent contracting as liberation from traditional employment
5. **CSS Strategy**: 100% reuse of existing subpage-styles.css with minimal page-specific styles

---

## File 1: MODIFY docs/index.html

### Location
Insert new feature card after OSINT Intelligence card (approximately line 1500, within `.features-grid`)

### New Arsenal Card HTML
```html
<a href="arsenal/independent-contractor.html" class="feature-card">
    <div class="feature-icon">💼</div>
    <h3 class="feature-title">Independent Contractor</h3>
    <p class="feature-desc">Skip the job hunt entirely. Define your services, set your rates, and find clients who pay what you're worth.</p>
    <div class="feature-tags">
        <span class="feature-tag">Service Catalog</span>
        <span class="feature-tag">Rate Cards</span>
        <span class="feature-tag">Client Discovery</span>
    </div>
</a>
```

### Rationale
- Icon: `💼` (briefcase) - represents professional services/consulting
- Copy focuses on benefits (skip job hunt, set your rates) not features
- Tags preview the three main value propositions

---

## File 2: CREATE docs/arsenal/independent-contractor.html

### Purpose
Overview page introducing the Independent Contractor arsenal with links to 5 command subpages

### Page Title
"Build Your Own Practice"

### Page Subtitle
"Not everyone wants a boss. Some people want clients. Define your services, set your rates, find prospects, and close deals - without asking permission to work."

### Icon
`💼`

### Badge
"Weapons System"

### Command Display
```html
<div class="command-display multi">
    <code>/defineservices</code>
    <span class="or">→</span>
    <code>/ratecard</code>
    <span class="or">→</span>
    <code>/findclient</code>
    <span class="or">→</span>
    <code>/pitchdeck</code>
    <span class="or">→</span>
    <code>/proposaltemplate</code>
</div>
```

### Sections

#### Section 1: The Problem It Solves
**Copy:**
"The job market treats you like a commodity - one of hundreds applying for the same position, waiting for someone else to decide your worth. Independent contracting flips that dynamic: you define your services, you set your rates, you choose your clients. But most people don't know where to start."

#### Section 2: The 5-Step Workflow
**Card Grid (agent-grid pattern from OSINT):**

| Icon | Title | Agent ID | Bullet Points |
|------|-------|----------|---------------|
| `📋` | Define Your Services | /defineservices | Define service catalog, Set pricing strategy, Establish engagement models, Build competitive differentiation |
| `💵` | Generate Rate Card | /ratecard | Professional rate card, Pricing validation, Volume discounts, Payment terms |
| `🎯` | Find Your Clients | /findclient | B2B prospect discovery, 10-point fit scoring, Entry point mapping, Outreach strategy |
| `🎪` | Create Pitch Deck | /pitchdeck | Provenance-hardened claims, Prospect-specific, Industry-tailored, Service-focused |
| `📄` | Write Proposals | /proposaltemplate | McKinsey/BCG-style, Transparent pricing, Engagement-type templates, Professional terms |

#### Section 3: How It Works
**Process Flow (4 steps):**
1. **Define Your Offering** - What you do, who you serve, what you charge
2. **Find Your Market** - Discover companies that need what you offer
3. **Build Your Pitch** - Credible, evidence-backed presentations
4. **Close the Deal** - Professional proposals that convert

#### Section 4: What You'll Build
**Feature List:**
- **Service Catalog**: 3-5 defined services with deliverables, ideal client profiles, and pricing
- **Professional Rate Card**: Validated hourly/daily/retainer rates with payment terms
- **Prospect Pipeline**: Scored and prioritized list of potential clients
- **Sales Materials**: Pitch decks and proposal templates ready to customize

#### Section 5: Who This Is For
**Feature List:**
- **Career Changers**: Transition from employment to self-employment
- **Consultants**: Formalize existing practice with professional materials
- **Freelancers**: Scale from gig work to strategic engagements
- **Executives**: Offer fractional/advisory services between roles

#### Section 6: Usage Example
**Terminal Mini:**
```
$ /defineservices --from-profile
Extracting services from career history...
✓ 4 services identified from work history
✓ Pricing aligned with market research
✓ Service definition generated

$ /findclient --industry="Commercial Real Estate" --limit=25
Deploying B2B discovery agents...
✓ 25 prospects scored and prioritized
✓ HIGH: 8 | MEDIUM: 12 | LOW: 5
→ Client_Prospects/Prospects_CommercialRealEstate_2025-12-04.md
```

### Navigation
- Previous: OSINT Intelligence (arsenal/osint-intelligence.html)
- Next: Return to Arsenal Overview (../index.html#features)

### Page-Specific Styles
Reuse `.agent-grid`, `.agent-card`, `.process-flow`, `.feature-list` from OSINT page - no new styles needed.

---

## File 3: CREATE docs/arsenal/contractor/defineservices.html

### Page Title
"Stop Selling Your Time. Start Selling Solutions."

### Page Subtitle
"Most consultants compete on hourly rate - a race to the bottom. Service definition transforms your expertise into productized offerings with clear value propositions and defensible pricing."

### Icon
`📋`

### Badge
"Contractor Command"

### Agent ID Display
`/defineservices`

### Sections

#### Mission Profile (2 paragraphs)
**Paragraph 1 (Stakes/Fear):**
"Without a defined service catalog, you're just another freelancer competing on price. Clients don't know what they're buying, you don't know what to charge, and every engagement starts from scratch. That's exhausting and unsustainable."

**Paragraph 2 (Benefit Questions):**
"What exactly do you offer? Who is your ideal client? What makes you different from every other consultant? How do you justify your rates? Service definition answers these questions with evidence from your actual career history."

#### Three Modes Section
**Method Cards Grid:**
1. **Guided Mode** (`--guided`): Interactive questions walk you through service definition
2. **Profile-Based Mode** (`--from-profile`): Auto-generates from your career history with evidence mapping
3. **Update Mode** (`--update`): Refresh existing definition with new achievements and pricing

#### What You'll Define
**Feature List:**
1. **Service Catalog (3-5 Services)**: Name, category, description, deliverables, success metrics
2. **Ideal Client Profiles**: Target industries, company sizes, pain points, decision makers
3. **Pricing Strategy**: Hourly/daily/retainer rates, payment terms, expense handling
4. **Engagement Models**: Preferred (project/retainer/fractional), acceptable, avoid
5. **Competitive Differentiation**: Unique value, advantages with proof points

#### Pricing Validation
**Assessment Box:**
"Your rates must pass consistency validation: daily rates align with hourly (within 10%), tiers ascend properly (min < target < premium), and premium multiplier falls within market norms (1.3x-2.0x). Invalid pricing is flagged before you embarrass yourself with clients."

#### Sample Output Preview
**Findings Preview (4 items):**
- **SERVICE**: AI Strategy Advisory - Expert strategic planning for AI adoption
- **PRICING**: $250-$400/hr | $2,000-$3,200/day | Retainers from $8K/mo
- **IDEAL CLIENT**: Mid-market enterprises ($50M-$500M revenue) facing digital transformation
- **DIFFERENTIATION**: 15+ years enterprise AI + P.Eng credential + 47 peer-reviewed publications

#### Output Structure
**Folder Tree:**
```
Client_Prospects/
├── Service_Definition_20251204.md
└── .cache/
    └── Service_Definition_20251204.json  (profile-based only)
```

### Navigation
- Previous: Back to Independent Contractor Overview (../independent-contractor.html)
- Next: Rate Card Generation (ratecard.html)

---

## File 4: CREATE docs/arsenal/contractor/ratecard.html

### Page Title
"Know Your Worth. Prove Your Worth."

### Page Subtitle
"A professional rate card isn't just a price list - it's a sales tool that establishes credibility before you ever meet a client. Stop negotiating from weakness."

### Icon
`💵`

### Badge
"Contractor Command"

### Agent ID Display
`/ratecard`

### Sections

#### Mission Profile
**Paragraph 1:**
"Consultants who can't articulate their pricing lose deals. They fumble when asked about rates, make up numbers on the spot, and undercut themselves out of insecurity. A rate card fixes this - it's your pricing authority in document form."

**Paragraph 2:**
"What's your hourly rate? What about daily? Do you offer retainers? What's included? What's extra? When's payment due? Professional rate cards answer every pricing question before clients ask, positioning you as established - not desperate."

#### Pricing Validation Section
**Requirements Box:**
- **Daily = Hourly x 8**: Rates must align (within 10% tolerance)
- **Ascending Tiers**: Min < Target < Premium (both hourly and daily)
- **Premium Multiplier**: Between 1.3x and 2.0x of target
- **Retainer Discounts**: Increase with commitment (5-40% off hourly)

#### Output Formats
**Method Cards (3):**
1. **Markdown** (default): Clean, portable, version-controlled
2. **HTML**: Professional web-ready with embedded styling
3. **PDF**: Pixel-perfect via Playwright, ready to send

#### What's Included
**Feature List:**
- Hourly rates (minimum/target/premium tiers)
- Daily rates (8-hour equivalents, calculated for consistency)
- Retainer packages (part-time/half-time/full-time)
- Volume and term discounts
- Expense policy (included vs. billable)
- Payment terms (deposit, invoice frequency, late fees)
- 6-month validity period

#### Credentials Justification (--include-justification)
**Assessment Box:**
"Add the `--include-justification` flag to include a 'Why These Rates?' section with your credentials, certifications, proof points, and ROI evidence. Especially powerful for premium pricing that needs defending."

#### Sample Rate Card Preview
**Report Section (simulated):**
```
═══════════════════════════════════════════════════════════════
[CONSULTANT NAME] | CAD | VALID: Dec 2024 to Jun 2025
───────────────────────────────────────────────────────────────
Hourly:    $250/hr   | $350/hr   | $450/hr
           Minimum     Target      Premium
Daily:     $2,000/day| $2,800/day| $3,600/day
───────────────────────────────────────────────────────────────
Retainer:  Part $4K/mo | Half $8K/mo | Full $14K/mo
PAYMENT: 50% upfront, Net 30 | MINIMUM: 8 hours
═══════════════════════════════════════════════════════════════
```

### Navigation
- Previous: Define Services (defineservices.html)
- Next: Find Clients (findclient.html)

---

## File 5: CREATE docs/arsenal/contractor/findclient.html

### Page Title
"Stop Waiting for Jobs. Start Finding Clients."

### Page Subtitle
"Job boards are where employees go to beg. B2B client discovery puts you in front of companies actively spending on consulting - with entry points mapped and outreach strategies ready."

### Icon
`🎯`

### Badge
"Contractor Command"

### Agent ID Display
`/findclient`

### Sections

#### Mission Profile
**Paragraph 1:**
"Most consultants rely on referrals and hope. They don't know how to systematically find new clients, so they wait for the phone to ring. When it doesn't, they panic and take bad engagements at bad rates."

**Paragraph 2:**
"Which companies hire consultants like you? How do you get in front of them? Who makes the buying decision? What pain points justify your rates? Client discovery answers these questions with researched, scored prospects and mapped entry points."

#### 10-Point B2B Fit Scoring
**Scoring Grid (5 categories):**
- **Contractor History (25%)**: Do they hire consultants? Recent job posts, vendor programs
- **Procurement Access (20%)**: How easy to get in? Vendor portals, small business programs
- **Domain Alignment (25%)**: Industry match, pain point evidence, service need
- **Size/Budget Fit (15%)**: Can they afford you? Revenue, funding, consulting spend
- **Geographic Match (15%)**: Location alignment, remote-friendly culture

#### Priority Classification
**Score Ranges:**
- **HIGH (8-10)**: Active outreach recommended - full detail report
- **MEDIUM (5-7.9)**: Worth pursuing with right introduction - condensed report
- **LOW (1-4.9)**: Deprioritize - table format only

#### Entry Point Mapping
**Feature List:**
- **Procurement Portals**: Vendor registration URLs and processes
- **Warm Introductions**: 2nd-degree LinkedIn connections, alumni networks
- **Intermediaries**: Staffing agencies, consulting marketplaces
- **Decision-Makers**: LinkedIn profiles, accessibility notes
- **Content/Events**: Conferences, guest posting, engagement opportunities

#### Filters
**Pattern List:**
- `--industry="X"`: Filter by industry
- `--size=startup|mid|enterprise`: Company size (startup: 1-50, mid: 50-500, enterprise: 500+)
- `--location="X"`: Geographic filter
- `--limit=N`: Max prospects (default: 20, max: 50)

#### Sample Output Preview
**Findings Preview:**
- **HIGH 8.7**: Hatch Ltd. - Engineering consulting, active contractor program, Toronto HQ
- **HIGH 8.2**: Oxford Properties - CRE, vendor portal, recent AI initiative press
- **MEDIUM 6.5**: Clio - Legal tech, startup culture, Vancouver (remote-friendly)
- **LOW 3.2**: SmallCo Inc. - No contractor signals, limited budget indicators

### Navigation
- Previous: Rate Card (ratecard.html)
- Next: Pitch Deck (pitchdeck.html)

---

## File 6: CREATE docs/arsenal/contractor/pitchdeck.html

### Page Title
"Credibility That Closes Deals"

### Page Subtitle
"Generic pitch decks get ignored. Prospect-specific, provenance-hardened presentations with evidence-backed claims show clients you're not just another consultant - you're the right one."

### Icon
`🎪`

### Badge
"Contractor Command"

### Agent ID Display
`/pitchdeck`

### Sections

#### Mission Profile
**Paragraph 1:**
"Consultants lose deals because they can't prove their claims. They say 'I increased revenue' without numbers, 'I led transformations' without specifics, 'I'm an expert' without evidence. Skeptical buyers tune out."

**Paragraph 2:**
"What results can you prove? What pain points does your prospect have? How does your solution address their specific situation? Provenance-hardened pitch decks answer these questions with traced, verified claims from your actual career history."

#### Four Targeting Modes
**Method Cards:**
1. **Prospect Mode** (`--prospect=name`): Company-specific research, tailored pain points
2. **Industry Mode** (`--industry=X`): Industry trends, common challenges
3. **Service Mode** (`--service=name`): Deep-dive on single offering
4. **General Mode** (default): All services, industry-agnostic capabilities

#### 12-Slide Structure
**Feature List:**
1. Title - Identity, tagline, contact
2. Problem - Pain points (mode-specific)
3. Cost of Inaction - Quantified impact
4. Solution - Service overview
5. How It Works - Methodology phases
6. Results & Proof - Case studies with evidence
7. Why Me/Us - Credentials, differentiation
8. Engagement Options - Pricing packages
9. Relevant Experience - Industry-filtered examples
10. Next Steps - CTA, timeline
11. Q&A - Anticipated objections with answers
12. Appendix - Full credentials

#### Provenance Hardening
**Assessment Box:**
"Every quantified claim is validated against your candidate profile with file + line citations. Validation rate targets: 90%+ = HIGH confidence (ready to present), 80-89% = MEDIUM (review recommended), <80% = LOW (major revision required). No more undefendable claims."

#### Research Depth
**Process Flow:**
1. **News & Initiatives (2-3 min)**: Press releases, strategic priorities, funding
2. **Pain Points (2 min)**: Glassdoor, job postings, RFPs
3. **Decision-Makers (1-2 min)**: LinkedIn verification, recent activity
4. **Validate Sources**: 5+ credible sources, <6 months old

#### Sample Provenance Trail
**Report Section (Evidence Mapping):**
| Slide | Claim | Evidence Source | Confidence |
|-------|-------|-----------------|------------|
| 6 | "Reduced costs by $2.4M" | Experience/Manulife.md:47-52 | HIGH |
| 7 | "P.Eng, PMP certifications" | certifications.json | HIGH |
| 9 | "Led 47-person team" | Experience/AECOM.md:23-28 | HIGH |

### Navigation
- Previous: Find Clients (findclient.html)
- Next: Proposal Template (proposaltemplate.html)

---

## File 7: CREATE docs/arsenal/contractor/proposaltemplate.html

### Page Title
"Proposals That Win"

### Page Subtitle
"Amateur proposals get rejected. McKinsey/BCG-style templates with transparent pricing calculations show clients you're a professional worth the investment - before negotiations even start."

### Icon
`📄`

### Badge
"Contractor Command"

### Agent ID Display
`/proposaltemplate`

### Sections

#### Mission Profile
**Paragraph 1:**
"Most consultant proposals are poorly structured, vague on scope, and confusing on pricing. Clients don't know what they're buying, doubt they'll get value, and negotiate harder out of uncertainty."

**Paragraph 2:**
"What exactly will you deliver? How long will it take? How did you calculate that price? What are the terms? Professional proposals answer every question before it's asked, positioning you as the low-risk choice."

#### Four Engagement Types
**Method Cards:**
1. **Project** (default): Fixed-scope engagements with phase-based delivery
2. **Retainer**: Ongoing advisory with monthly hour allocations
3. **Staff Augmentation**: Embedded consulting with daily rates
4. **Workshop**: Training/facilitation with prep and followup

#### 10-Section Structure
**Feature List:**
1. Cover Page - Metadata, reference number, validity
2. Executive Summary - Problem, solution, outcomes, investment
3. Understanding Your Challenge - Pain points, desired state
4. Proposed Approach - Methodology, phases, activities
5. Deliverables & Timeline - Gantt-style, milestones
6. Team & Qualifications - Credentials, achievements
7. Investment - Pricing table, payment terms, expenses
8. Terms & Conditions - Scope, IP, termination
9. Next Steps - Review process, kickoff timeline
10. Appendix - Credentials, case studies, references

#### Transparent Pricing Calculations
**Assessment Box:**
"Pricing is calculated from your rate card with full methodology in YAML frontmatter. Clients see: estimated hours x hourly rate x volume discount = final price. Transparent pricing builds trust and reduces negotiation friction."

#### Pricing Formulas by Type
**Pattern List:**
- **Project**: `hours x hourly.target x volume_discount`
- **Retainer**: `monthly_hours x effective_rate` (discounts for commitment)
- **Staff-Aug**: `20_days x daily.target x duration_discount`
- **Workshop**: `(workshop_days x premium) + (prep x target) + materials`

#### Client Context Integration
**Feature List:**
- **With Prior Research**: Uses pain points and decision-makers from pitch deck
- **Generic Mode**: Uses ideal client profiles from service definition
- **Candidate Profile**: Pulls credentials and achievements for Team section

#### Sample Proposal Preview
**Report Header:**
```
PROPOSAL: AI Strategy Advisory Engagement
CLIENT: Hatch Ltd.
REFERENCE: PROP-20251204-HA-AISA
VALID: 14 days | INVESTMENT: $47,500 CAD
```

### Navigation
- Previous: Pitch Deck (pitchdeck.html)
- Next: Return to Independent Contractor Overview (../independent-contractor.html)

---

## CSS Considerations

### Reused Components (from subpage-styles.css)
All existing styles are sufficient:
- `.arsenal-badge`, `.arsenal-icon` - page header styling
- `.command-display` - command display box
- `.content-section`, `.feature-list` - content sections
- `.method-cards`, `.method-card` - mode/type cards
- `.scoring-grid`, `.score-category` - scoring breakdown
- `.terminal-mini` - usage examples
- `.findings-preview`, `.finding` - sample output
- `.nav-steps`, `.nav-step` - page navigation
- `.folder-tree`, `.tree-item` - output structure
- `.report-header`, `.report-section` - sample reports

### Page-Specific Styles
Minimal additions needed. Add styles inline within `<style>` tag at bottom of each page (following OSINT subpage pattern).

### New Styles Needed (Add inline per page)

**IMPORTANT**: The `.assessment-box` pattern is NOT in subpage-styles.css. All pages using assessment boxes must define inline:
```css
.assessment-box { padding: 1rem; border-left: 3px solid; }
.assessment-box.positive { background: rgba(34, 197, 94, 0.1); border-color: var(--accent-green); }
.assessment-label { font-family: 'JetBrains Mono', monospace; font-size: 0.65rem; text-transform: uppercase; letter-spacing: 0.1em; color: var(--accent-green); display: block; margin-bottom: 0.5rem; }
.assessment-box p { font-size: 0.85rem; margin: 0; line-height: 1.5; }
```

**defineservices.html**:
- Needs `.assessment-box` for pricing validation section

**ratecard.html**:
- Rate card table/visualization - adapt `.salary-breakdown`, `.comp-table`, `.comp-row` patterns from compensation.html
- Needs `.assessment-box` for credentials justification section

**findclient.html**:
- B2B scoring breakdown - adapt `.scoring-grid`, `.score-category` from existing patterns
- Priority classification display - reuse `.priority-tag` classes (already in subpage-styles.css)

**pitchdeck.html**:
- Needs `.assessment-box` for provenance hardening section
- Provenance evidence table - reuse `.table-preview` with custom headers
- Research process flow - reuse `.process-flow`, `.flow-step` patterns

**proposaltemplate.html**:
- Needs `.assessment-box` for transparent pricing section
- Pricing formula display - adapt `.formula-display` patterns
- Engagement type comparison - reuse `.method-cards` grid

All other components use existing CSS classes without modification.

---

## Navigation Structure

### Linear Flow
```
index.html#features
    └── independent-contractor.html (overview)
            ├── contractor/defineservices.html
            ├── contractor/ratecard.html
            ├── contractor/findclient.html
            ├── contractor/pitchdeck.html
            └── contractor/proposaltemplate.html
```

### Back Links
- All subpages link back to independent-contractor.html
- Overview links back to index.html#features
- Previous/Next navigation follows workflow order

---

## Implementation Sequence

### Phase 1: Index Modification (GitHub Issue #32)
1. Add 5th feature card to docs/index.html

### Phase 2: Overview Page (GitHub Issue #33)
1. Create docs/arsenal/independent-contractor.html

### Phase 3: Subpages (GitHub Issues #34-38)
Create in order:
1. docs/arsenal/contractor/defineservices.html (#34)
2. docs/arsenal/contractor/ratecard.html (#35)
3. docs/arsenal/contractor/findclient.html (#36)
4. docs/arsenal/contractor/pitchdeck.html (#37)
5. docs/arsenal/contractor/proposaltemplate.html (#38)

### Phase 4: Testing
1. Verify all navigation links work
2. Test responsive behavior on mobile
3. Validate consistent styling with OSINT pages

---

## Open Questions

1. **Icon Selection**: Should we use emoji or consider SVG icons for consistency with a potential future design system?

2. **Order in Arsenal Section**: The new card is placed after OSINT. Should the order be: HAM-Z, Provenance, Cultural, OSINT, Contractor - or rearranged by workflow?

3. **Sample Data**: The sample outputs reference specific companies and dollar amounts. Should these be more generic or use obviously fictional companies?

4. **Currency**: Rate card examples use CAD. Should examples show USD for broader appeal, or keep CAD to match the system's Canadian cultural profile default?

---

## Critical Files for Implementation

- `/workspaces/resumeoptimizer/docs/index.html` - Add 5th Arsenal card in features-grid section (lines 1459-1500)
- `/workspaces/resumeoptimizer/docs/arsenal/osint-intelligence.html` - Primary pattern for overview page structure
- `/workspaces/resumeoptimizer/docs/arsenal/osint/corporate.html` - Primary pattern for subpage structure
- `/workspaces/resumeoptimizer/docs/subpage-styles.css` - Existing CSS to reuse (no modifications needed)
- `/workspaces/resumeoptimizer/.claude/commands/independent-contractor/defineservices.md` - Source for defineservices page content
