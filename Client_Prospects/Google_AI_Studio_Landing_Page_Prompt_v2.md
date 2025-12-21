# Tenebrus Capital Landing Page — Google AI Studio Prompt v2.0

**Purpose:** Production-ready prompt optimized for AI generation quality and conversion
**Version:** 2.0 (Expert-reviewed revision)
**Generated:** 2025-12-21

---

# PART 1: AI GENERATION INSTRUCTIONS

## Output Requirements

**Format:** Single-page HTML with inline CSS (no external dependencies)
**Framework:** Vanilla HTML/CSS with minimal JavaScript for interactions
**Responsive:** Mobile-first, breakpoints at 768px and 1200px

## Priority Hierarchy (If Constraints Force Cuts)

1. **CRITICAL (Never omit):** Hero, Credentials Bar, Services Overview, Final CTA, Footer
2. **HIGH (Omit only if necessary):** Track Record Metrics, Getting Started Options
3. **MEDIUM (Can simplify):** Problem Section, Process Timeline
4. **LOW (Can omit entirely):** Who We Serve grid, Optional animations

## Generation Rules

1. Use EXACT copy provided — do not paraphrase or add language
2. When a specific value is given (e.g., "48px"), use that exact value
3. Do not add stock photography, illustrations of people, or placeholder images
4. Do not add words like "revolutionary," "cutting-edge," "transform," "AI-powered," or exclamation marks
5. All CTAs link to: `https://calendly.com/tenebruscapital/consultation` (placeholder)
6. All external links open in new tab with `rel="noopener noreferrer"`

---

# PART 2: DESIGN SYSTEM (Use Exactly)

## Color Palette (7 Colors Only)

```css
:root {
  --bg-primary: #0f172a;      /* Deep slate - hero, CTA sections */
  --bg-secondary: #1e293b;    /* Slate - alternating sections */
  --bg-card: #334155;         /* Card backgrounds */
  --accent-gold: #f59e0b;     /* Amber-500 - CTAs, highlights (WCAG AA compliant) */
  --accent-gold-hover: #d97706; /* Amber-600 - hover states */
  --text-primary: #f8fafc;    /* Slate-50 - headings */
  --text-secondary: #94a3b8;  /* Slate-400 - body, supporting */
}
```

## Typography (Specific Values)

```css
/* Headings: System serif stack */
--font-heading: Georgia, Cambria, "Times New Roman", serif;

/* Body: System sans stack */
--font-body: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;

/* Sizes - Desktop */
--h1: 52px;      /* Hero headline only */
--h2: 36px;      /* Section headlines */
--h3: 22px;      /* Card titles, subsections */
--body: 17px;    /* Paragraphs */
--small: 14px;   /* Labels, metadata */
--eyebrow: 13px; /* Section eyebrows */

/* Sizes - Mobile (<768px): Reduce all by 20% */
```

## Spacing (8px Base Unit)

```css
--space-xs: 8px;
--space-sm: 16px;
--space-md: 24px;
--space-lg: 48px;
--space-xl: 80px;
--space-2xl: 120px;

/* Section padding: 80px top/bottom desktop, 48px mobile */
/* Max content width: 1140px, centered */
/* Card padding: 32px */
/* Card gap: 24px */
```

## Component Specifications

**Primary Button:**
```css
.btn-primary {
  background: var(--accent-gold);
  color: var(--bg-primary);
  padding: 16px 32px;
  border-radius: 6px;
  font-weight: 600;
  font-size: 16px;
  border: none;
  cursor: pointer;
  transition: background 0.2s ease;
}
.btn-primary:hover {
  background: var(--accent-gold-hover);
}
```

**Ghost Button:**
```css
.btn-ghost {
  background: transparent;
  color: var(--accent-gold);
  border: 1px solid var(--accent-gold);
  padding: 14px 28px;
  border-radius: 6px;
  font-weight: 500;
}
```

**Card:**
```css
.card {
  background: var(--bg-card);
  border-radius: 12px;
  padding: 32px;
  border: 1px solid rgba(148, 163, 184, 0.1);
}
```

**Eyebrow Label:**
```css
.eyebrow {
  color: var(--accent-gold);
  font-size: 13px;
  font-weight: 600;
  letter-spacing: 2px;
  text-transform: uppercase;
  margin-bottom: 16px;
}
```

**Metric Display:**
```css
.metric-number {
  font-size: 56px;
  font-weight: 700;
  color: var(--accent-gold);
  line-height: 1;
}
.metric-context {
  font-size: 15px;
  color: var(--text-secondary);
  margin-top: 8px;
}
```

---

# PART 3: PAGE STRUCTURE (8 Sections)

## Section 1: Navigation

**Layout:** Fixed top, height 72px, transparent initially, solid `--bg-primary` on scroll

**Left:** Logo text
```
TENEBRUS CAPITAL
```
Style: `--text-primary`, 18px, font-weight 600, letter-spacing 1px

**Right:** CTA button only (no nav links on single-page)
```
Schedule Consultation
```
Style: Small primary button (padding 12px 24px)

---

## Section 2: Hero

**Background:** `--bg-primary`
**Layout:** Centered, max-width 800px, padding 120px top, 80px bottom

### Credential Badge (First Element)
```
FRICS  ·  CFA  ·  COLUMBIA AI/ML  ·  24 YEARS
```
Style: `--accent-gold`, 13px, letter-spacing 3px, uppercase, flex row with dots as separators

### Headline
```
Your Portfolio Has Scaled.
Your Operations Haven't.
```
Style: `--font-heading`, 52px, `--text-primary`, line-height 1.15, text-align center
Note: Line break after first sentence

### Subheadline
```
Workflow engineering for institutional real estate—delivered by someone who has managed $750M+ in assets, not consultants learning your business on your budget.
```
Style: `--font-body`, 20px, `--text-secondary`, max-width 680px, text-align center, margin-top 24px

### Social Proof Line (Below Subheadline)
```
141% returns vs. 35% market  ·  50% faster processing  ·  $5M+ value captured
```
Style: `--text-secondary`, 15px, margin-top 32px, flex row with dot separators

### Primary CTA
```
Schedule a Consultation
```
Style: Primary button, margin-top 40px

### Risk Reversal
```
30 minutes  ·  No fee  ·  Actionable recommendations even if we don't work together
```
Style: `--text-secondary`, 14px, margin-top 16px

---

## Section 3: Credibility Bar

**Background:** `--bg-secondary`
**Layout:** Full-width, padding 40px vertical, horizontal flex with 4 items evenly spaced, max-width 900px centered

### Items (Left to Right)

**Item 1:**
```
FRICS
Fellow, Royal Institution
of Chartered Surveyors
```

**Item 2:**
```
CFA
Chartered Financial
Analyst
```

**Item 3:**
```
COLUMBIA
AI & Machine Learning
Certificate
```

**Item 4:**
```
$750M+
Institutional Assets
Managed Directly
```

Style per item:
- Top line: `--text-primary`, 20px, font-weight 700
- Bottom lines: `--text-secondary`, 13px
- Vertical dividers between items (1px, rgba white 10%)

---

## Section 4: The Problem (Condensed)

**Background:** `--bg-primary`
**Eyebrow:** `THE CONSTRAINT`

### Headline
```
Technology Consultants Don't Understand Lease Economics.
Real Estate Advisors Can't Implement Systems.
```
Style: `--font-heading`, 36px, `--text-primary`, max-width 900px

### Supporting Text
```
Your portfolio needs both. That combination is held by fewer than 100 professionals globally.
```
Style: 18px, `--text-secondary`, margin-top 24px

### Problem Cards (2x2 Grid)

**Card 1:**
```
60% of analyst time spent on administrative work instead of deal sourcing
```
Icon: Clock (simple line icon)

**Card 2:**
```
$200K+ wasted on technology implementations that don't fit real estate workflows
```
Icon: Chart trending down

**Card 3:**
```
3-4 week competitive disadvantage vs. peers with optimized operations
```
Icon: Users/competitors

**Card 4:**
```
Compliance burden growing while headcount remains flat
```
Icon: Scale/balance

Card style: `--bg-card`, 24px padding, icon in `--accent-gold` top-left, text 16px `--text-secondary`

---

## Section 5: Track Record (Move Up for Social Proof)

**Background:** `--bg-secondary`
**Eyebrow:** `TRACK RECORD`

### Headline
```
Results Across Institutional Portfolios
```

### Metrics Grid (3 columns x 2 rows)

**Row 1:**

| Metric | Label |
|--------|-------|
| 141% | portfolio returns vs. 35% benchmark (4× outperformance) |
| $750M+ | institutional assets managed at Artis REIT |
| 50% | processing time reduction with improved compliance |

**Row 2:**

| Metric | Label |
|--------|-------|
| $5M+ | operational value captured at Hydro One |
| 204,000 | acres of linear infrastructure managed |
| 90% | voluntary settlement rate on complex negotiations |

Metric style: Number in `--accent-gold`, 56px, bold. Label below in `--text-secondary`, 15px.

---

## Section 6: Services (Simplified to 3 Entry Points)

**Background:** `--bg-primary`
**Eyebrow:** `HOW TO START`

### Headline
```
Three Pathways to Better Operations
```

### Subheadline
```
Most clients begin with a workshop or pilot project. Enterprise engagements follow once ROI is proven.
```

### Service Cards (3-column grid, equal height)

---

**CARD 1: LOWEST BARRIER**

```
WORKSHOP

Get your team productive with workflow automation in one day.

What's Included:
• Pre-workshop needs assessment
• Role-specific curriculum (leasing, asset management, accounting, or appraisal)
• Hands-on exercises with your actual documents
• Reference guides and 30-day email support

Best For: Teams new to AI-augmented workflows

Duration: Half-day to full-day
Investment: $3,500 – $10,000
```

CTA Button (Ghost style): `Learn More`

---

**CARD 2: PROVE ROI**

```
PILOT PROJECT

Prove measurable ROI on one high-impact workflow before committing further.

What's Included:
• Operational assessment and bottleneck identification
• Implementation on single workflow (lease abstraction, credit scoring, or compliance)
• Training and documentation
• 30-day post-implementation support

Typical Results: 40-60% time reduction on target workflow

Duration: 4-8 weeks
Investment: $10,000 – $25,000
```

CTA Button (Primary style): `Start a Pilot`

---

**CARD 3: ENTERPRISE**

```
FULL IMPLEMENTATION

Complete operational transformation with platform deployment and system integration.

What's Included:
• 11 financial calculators, 25 automated workflows
• Integration with Yardi, MRI, VTS, Argus
• Multi-track training program
• 6-month support with monthly reviews

Typical Results: 85-95% time savings on routine analysis

Duration: 2-3 months
Investment: $50,000 – $150,000
```

CTA Button (Ghost style): `Discuss Requirements`

---

Card visual hierarchy:
- Card 2 (Pilot) should be visually emphasized: Slightly larger, gold border-top 3px, or subtle glow
- Badge on Card 2: "MOST POPULAR" in small gold pill above card title

---

## Section 7: Process (Simplified Timeline)

**Background:** `--bg-secondary`
**Eyebrow:** `THE PROCESS`

### Headline
```
From Assessment to Independence in 90 Days
```

### Timeline (Horizontal on desktop, vertical on mobile)

**Step 1**
```
ASSESS
1-2 weeks

We analyze workflows and identify the highest-ROI opportunities. You receive actionable recommendations regardless of next steps.
```

**Step 2**
```
IMPLEMENT
3-8 weeks

We deploy solutions with minimal disruption. All systems include documentation and training.
```

**Step 3**
```
TRANSFER
30-90 days

We ensure your team operates independently. Optional ongoing advisory available.
```

Visual: Numbered circles (1, 2, 3) in `--accent-gold`, connected by horizontal line. Each step in a subtle card below its number.

---

## Section 8: Final CTA + Footer Combined

**Background:** `--bg-primary` with subtle radial gradient (lighter center)
**Layout:** Two parts stacked

### CTA Section (Top)

**Headline:**
```
Explore What's Possible for Your Portfolio
```
Style: `--font-heading`, 40px, `--text-primary`, centered

**Supporting Text:**
```
30-minute consultation. No fee. No commitment required.
Actionable recommendations whether we work together or not.
```
Style: 18px, `--text-secondary`, max-width 600px, centered, margin-top 16px

**Primary CTA:**
```
Schedule a Consultation
```
Style: Large primary button (padding 20px 48px, font-size 18px), margin-top 32px

**Alternative:**
```
Or email directly: reggie@tenebruscapital.com
```
Style: 14px, `--text-secondary`, margin-top 16px, email as gold link

### Footer (Bottom, border-top 1px)

**Layout:** 3 columns on desktop, stacked on mobile, padding 48px vertical

**Column 1:**
```
TENEBRUS CAPITAL

Workflow engineering for institutional
real estate and infrastructure.

© 2025 Tenebrus Capital
```

**Column 2:**
```
SERVICES

Workshops
Pilot Projects
Enterprise Implementation
```

**Column 3:**
```
CONTACT

Reggie Chan, FRICS, CFA
Principal

Toronto, Ontario
LinkedIn (link icon)
```

LinkedIn URL: `https://linkedin.com/in/reggiechan`

Style: Column headers in `--accent-gold`, 13px uppercase. Content in `--text-secondary`, 14px.

---

# PART 4: RESPONSIVE BEHAVIOR

## Desktop (≥1200px)
- Full layouts as described
- Navigation shows CTA button
- Service cards in 3-column grid
- Metrics in 3x2 grid
- Timeline horizontal

## Tablet (768px - 1199px)
- Service cards in 2-column (third card full-width below)
- Metrics in 2x3 grid
- Timeline horizontal but condensed
- Reduce section padding to 64px

## Mobile (<768px)
- Navigation: Logo left, hamburger right (opens slide-out with CTA)
- Hero headline: 36px
- All grids become single column
- Timeline becomes vertical
- Service cards stack
- Footer columns stack
- Section padding: 48px
- Card padding: 24px
- Metric numbers: 44px

---

# PART 5: INTERACTIONS (Optional, Apply If Supported)

## Required (Implement These)
1. **Sticky nav darkening:** Transparent initially, solid `--bg-primary` after 100px scroll
2. **Button hover states:** Defined in component specs above
3. **Smooth scroll:** Anchor links scroll smoothly (CSS `scroll-behavior: smooth`)

## Optional (Nice to Have)
1. **Fade-in on scroll:** Sections fade in as they enter viewport (use `IntersectionObserver`)
2. **Metric counter:** Numbers count up from 0 when section enters viewport
3. **Card hover lift:** Cards translate up 4px and increase shadow on hover

---

# PART 6: ANTI-PATTERNS (Do Not Generate)

The following will cause rejection of the output:

1. **Language violations:**
   - "AI-powered," "AI-driven," "machine learning" in marketing copy
   - "Revolutionary," "cutting-edge," "game-changing," "transform"
   - "Best-in-class," "world-class," "industry-leading"
   - Exclamation marks anywhere
   - "Book a call," "Let's chat," "Get started" (use "Schedule a Consultation")

2. **Visual violations:**
   - Stock photography of any kind
   - Illustrations of people
   - Generic tech/abstract backgrounds with blue gradients
   - Light/white color scheme
   - More than 3 font sizes per element type
   - Rounded pill buttons (use 6px radius)

3. **UX violations:**
   - Popup modals on load
   - Auto-playing anything
   - Parallax scrolling effects
   - More than one CTA per viewport
   - Navigation dropdowns

4. **Content violations:**
   - Adding copy not in this prompt
   - Removing or paraphrasing provided copy
   - Changing numbers or metrics
   - Adding testimonial quotes (none provided)

---

# PART 7: VALIDATION CHECKLIST

Before finalizing, verify:

**Content:**
- [ ] Hero credential badge shows: FRICS · CFA · COLUMBIA AI/ML · 24 YEARS
- [ ] Hero headline has line break after first sentence
- [ ] All 6 metrics appear with exact numbers (141%, $750M+, 50%, $5M+, 204,000, 90%)
- [ ] Three service cards present (Workshop, Pilot, Enterprise)
- [ ] Pilot card has "MOST POPULAR" indicator
- [ ] Footer includes LinkedIn link
- [ ] Email shows: reggie@tenebruscapital.com
- [ ] No prohibited language from Part 6

**Design:**
- [ ] Dark background throughout (no white/light sections)
- [ ] Gold accent (#f59e0b) used for CTAs and highlights only
- [ ] Primary CTA button appears exactly 3 times (nav, hero, final)
- [ ] No stock photography or people illustrations
- [ ] Proper heading hierarchy (one H1, H2s for sections)

**Functionality:**
- [ ] All CTA buttons link to Calendly placeholder
- [ ] LinkedIn opens in new tab
- [ ] Mobile hamburger menu functional
- [ ] Smooth scroll on anchor links

---

# PART 8: SAMPLE HTML STRUCTURE (Reference)

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Tenebrus Capital | Workflow Engineering for Institutional Real Estate</title>
  <meta name="description" content="Operational optimization for commercial real estate portfolios. 24 years institutional experience. FRICS, CFA credentials.">
  <style>
    /* CSS variables and styles from Part 2 */
  </style>
</head>
<body>
  <nav class="navbar"><!-- Section 1 --></nav>
  <header class="hero"><!-- Section 2 --></header>
  <section class="credentials"><!-- Section 3 --></section>
  <section class="problem"><!-- Section 4 --></section>
  <section class="track-record"><!-- Section 5 --></section>
  <section class="services"><!-- Section 6 --></section>
  <section class="process"><!-- Section 7 --></section>
  <footer class="footer"><!-- Section 8 --></footer>
  <script>
    // Sticky nav, smooth scroll, optional animations
  </script>
</body>
</html>
```

---

# APPENDIX A: QUICK REFERENCE CARD

**Brand:** Tenebrus Capital
**Principal:** Reggie Chan, FRICS, CFA
**Tagline:** (None - avoid taglines, lead with credentials)
**Colors:** Dark slate (#0f172a), Amber accent (#f59e0b)
**Font:** Georgia (headings), System sans (body)
**CTA:** "Schedule a Consultation"
**Link:** calendly.com/tenebruscapital/consultation
**Email:** reggie@tenebruscapital.com
**LinkedIn:** linkedin.com/in/reggiechan
**Location:** Toronto, Ontario

---

# APPENDIX B: COPY QUICK-PASTE (All Sections)

**Hero Headline:**
Your Portfolio Has Scaled.
Your Operations Haven't.

**Hero Subheadline:**
Workflow engineering for institutional real estate—delivered by someone who has managed $750M+ in assets, not consultants learning your business on your budget.

**Hero Social Proof:**
141% returns vs. 35% market · 50% faster processing · $5M+ value captured

**Hero CTA:**
Schedule a Consultation

**Hero Risk Reversal:**
30 minutes · No fee · Actionable recommendations even if we don't work together

**Problem Headline:**
Technology Consultants Don't Understand Lease Economics.
Real Estate Advisors Can't Implement Systems.

**Problem Support:**
Your portfolio needs both. That combination is held by fewer than 100 professionals globally.

**Track Record Headline:**
Results Across Institutional Portfolios

**Services Headline:**
Three Pathways to Better Operations

**Services Subheadline:**
Most clients begin with a workshop or pilot project. Enterprise engagements follow once ROI is proven.

**Process Headline:**
From Assessment to Independence in 90 Days

**Final CTA Headline:**
Explore What's Possible for Your Portfolio

**Final CTA Support:**
30-minute consultation. No fee. No commitment required.
Actionable recommendations whether we work together or not.

---

*End of Prompt — Version 2.0*
