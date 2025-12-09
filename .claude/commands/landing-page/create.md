---
description: Create a professional landing page using tactical CSS design system and copywriting best practices
argument-hint: <page-name> [--template=tactical|minimal|corporate] [--output-dir=docs]
---

I'll create a professional landing page using the tactical CSS design system and strategic copywriting principles.

**Arguments:**
- `$1`: Page name/identifier (required) - e.g., `services`, `about`, `product-launch`
- `--template`: CSS template style (optional): `tactical` | `minimal` | `corporate` (default: tactical)
- `--output-dir`: Output directory (optional, default: `docs`)

## Phase 1: Discovery & Content Planning

### 1.1: Gather Page Requirements

Before creating the landing page, I need to understand:

1. **Page Purpose**: What is this landing page for?
   - Service promotion
   - Product launch
   - Lead generation
   - Event registration
   - Portfolio showcase
   - Company/brand introduction

2. **Target Audience**: Who should this page resonate with?
   - Industry/sector
   - Job titles/roles
   - Pain points and motivations
   - Technical sophistication level

3. **Desired Action**: What should visitors do?
   - Book a call/consultation
   - Download a resource
   - Sign up for newsletter
   - Request a quote
   - Start a trial

### 1.2: Invoke Copywriter Agent

Use the **landing-page-copywriter** agent to develop strategic copy:

```
Task tool with subagent_type='landing-page-copywriter'
```

The copywriter will deliver:
- Compelling headline and hook
- Hero section messaging
- Problem/agitation content
- Solution positioning
- Benefits (not features)
- Social proof framework
- Call-to-action copy
- Supporting section copy

## Phase 2: Load Design System

### 2.1: Read CSS Template

Based on the `--template` argument, load the appropriate CSS:

| Template | CSS File | Best For |
|----------|----------|----------|
| tactical | `docs/subpage-styles.css` | Tech, consulting, B2B services |
| minimal | `docs/minimal-styles.css` | Creative, portfolio, personal brands |
| corporate | `docs/corporate-styles.css` | Enterprise, financial, healthcare |

**Default**: Use `docs/subpage-styles.css` (tactical theme) which includes:
- Military/tactical aesthetic with gradient backgrounds
- Feature grids and method cards
- Terminal-style code displays
- Pipeline/step indicators
- Responsive breakpoints (tablet 768px, mobile 480px)

### 2.2: Identify Available Components

From the loaded CSS, catalog available component classes:
- Navigation: `.nav`, `.nav-link`, `.nav-cta`
- Hero: `.hero`, `.hero-content`, `.hero-title`, `.hero-subtitle`
- Sections: `.content-section`, `.section-title`, `.highlight-text`
- Feature displays: `.feature-grid`, `.feature-item`, `.feature-icon`
- Method cards: `.method-cards`, `.method-card`, `.method-title`
- Pipeline steps: `.pipeline-steps`, `.pipeline-step`
- CTAs: `.cta-section`, `.cta-btn`, `.secondary-btn`
- Footer: `.footer`, `.footer-links`

## Phase 3: Generate Landing Page

### 3.1: Invoke Frontend Design Skill

Use the **frontend-design skill** to generate production-grade HTML:

```
Skill: frontend-design:frontend-design
```

Provide the skill with:
1. The strategic copy from the copywriter agent
2. The CSS design system components
3. The page structure requirements

### 3.2: Page Structure Template

Generate HTML following this proven landing page structure:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>[Page Title] | [Brand Name]</title>
    <meta name="description" content="[SEO meta description from copywriter]">
    <link rel="stylesheet" href="subpage-styles.css">
</head>
<body>
    <!-- Navigation -->
    <nav class="nav">...</nav>

    <!-- Hero Section -->
    <section class="hero">
        <div class="hero-content">
            <h1 class="hero-title">[Headline]</h1>
            <p class="hero-subtitle">[Subheadline]</p>
            <a href="#cta" class="cta-btn">[Primary CTA]</a>
        </div>
    </section>

    <!-- Problem/Agitation -->
    <section class="content-section">
        <h2 class="section-title">[Problem Hook]</h2>
        <!-- Pain points, statistics, urgency -->
    </section>

    <!-- Solution/Benefits -->
    <section class="content-section">
        <h2 class="section-title">[Solution Hook]</h2>
        <div class="feature-grid">
            <!-- Benefits as feature items -->
        </div>
    </section>

    <!-- How It Works / Method -->
    <section class="content-section">
        <h2 class="section-title">[Process Hook]</h2>
        <div class="method-cards">
            <!-- Steps or methodology -->
        </div>
    </section>

    <!-- Social Proof -->
    <section class="content-section">
        <h2 class="section-title">[Credibility Hook]</h2>
        <!-- Testimonials, logos, metrics -->
    </section>

    <!-- Call to Action -->
    <section class="cta-section" id="cta">
        <h2>[CTA Headline]</h2>
        <p>[CTA Supporting Copy]</p>
        <a href="[action-url]" class="cta-btn">[Primary CTA Button]</a>
        <a href="[secondary-url]" class="secondary-btn">[Secondary CTA]</a>
    </section>

    <!-- Footer -->
    <footer class="footer">...</footer>
</body>
</html>
```

### 3.3: Design Requirements

The frontend-design skill MUST ensure:

1. **Visual Hierarchy**
   - Single primary CTA color (never compete with hero)
   - Headline is the largest text element
   - Subheadline supports, doesn't distract
   - Clear visual flow down the page

2. **Mobile-First Responsive**
   - Touch-friendly tap targets (min 44px)
   - Readable text without zooming
   - Stacked layouts on mobile
   - No horizontal scrolling

3. **Performance**
   - No external font dependencies (use font stacks)
   - Optimized images (lazy loading if images included)
   - Minimal JavaScript (CSS-only animations preferred)

4. **Accessibility**
   - Semantic HTML structure
   - Proper heading hierarchy (h1 > h2 > h3)
   - Sufficient color contrast (WCAG AA)
   - Focus states for interactive elements

## Phase 4: Quality Review

### 4.1: Landing Page Checklist

Before finalizing, verify:

| Category | Requirement | Status |
|----------|-------------|--------|
| **Copy** | Headline hooks in <3 seconds | [ ] |
| **Copy** | Benefits before features | [ ] |
| **Copy** | Single, clear CTA | [ ] |
| **Copy** | Social proof present | [ ] |
| **Design** | Consistent visual theme | [ ] |
| **Design** | Mobile-responsive | [ ] |
| **Design** | Above-fold CTA visible | [ ] |
| **Technical** | Valid HTML5 | [ ] |
| **Technical** | CSS linked correctly | [ ] |
| **Technical** | Meta tags populated | [ ] |

### 4.2: Preview with Playwright

Use Playwright MCP to preview the landing page:

1. Navigate to the generated HTML file
2. Take screenshots at desktop (1280px) and mobile (375px) widths
3. Verify visual appearance matches design intent
4. Check for layout issues or broken elements

## Phase 5: Output & Delivery

### 5.1: Save Files

Save the generated landing page:
- Location: `{output-dir}/{page-name}.html`
- Default: `docs/{page-name}.html`

### 5.2: Report Results

Provide:
- File location
- Screenshot previews (if taken)
- Any copywriting notes or suggestions
- Recommendations for A/B testing variations

## Example Usage

```bash
# Create a services landing page with tactical theme
/landing-page:create services

# Create a product launch page with minimal theme
/landing-page:create product-launch --template=minimal

# Create an about page in a custom directory
/landing-page:create about --output-dir=public/pages

# Create with corporate styling for enterprise clients
/landing-page:create enterprise-solutions --template=corporate
```

---

Now executing the landing page creation pipeline for `$1`...
