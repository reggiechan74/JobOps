---
description: Manage CSS design system templates for landing pages - view, analyze, or create component libraries
argument-hint: [--view|--analyze|--create] [template-name]
---

I'll help you manage the CSS design system templates for landing page development.

**Arguments:**
- `--view [template]`: Display available CSS templates and their components
- `--analyze [template]`: Deep-dive analysis of a specific template's design system
- `--create [name]`: Create a new CSS template based on existing patterns

**Default behavior (no arguments)**: List all available templates with summaries

## Available CSS Templates

| Template | File | Description |
|----------|------|-------------|
| **tactical** | `docs/subpage-styles.css` | Military/tech aesthetic with gradients, grid layouts, terminal displays |
| **minimal** | `docs/minimal-styles.css` | Clean, whitespace-focused, modern typography |
| **corporate** | `docs/corporate-styles.css` | Professional, conservative, enterprise-friendly |

---

## Mode: View Templates (`--view`)

### Tactical Template (`docs/subpage-styles.css`)

**Design Philosophy**: Mission-critical, data-driven, technically sophisticated

**Color Palette**:
```css
:root {
    --bg-dark: #0a0f1c;
    --bg-darker: #050810;
    --text-primary: #e0e6ed;
    --text-secondary: #8892a0;
    --accent: #4a9eff;
    --accent-secondary: #00d4aa;
    --border-color: #1a2234;
    --card-bg: #0d1424;
}
```

**Key Components**:

| Component | Class | Usage |
|-----------|-------|-------|
| Navigation | `.nav`, `.nav-link`, `.nav-cta` | Top navigation with CTA button |
| Hero | `.hero`, `.hero-title`, `.hero-subtitle` | Full-width gradient hero section |
| Content Sections | `.content-section`, `.section-title` | Standard content containers |
| Feature Grid | `.feature-grid`, `.feature-item` | 2-3 column benefit displays |
| Method Cards | `.method-cards`, `.method-card` | Step-by-step methodology |
| Pipeline Steps | `.pipeline-steps`, `.pipeline-step` | Linear process visualization |
| Terminal Display | `.terminal-mini`, `.terminal-header` | Code/command display |
| Scoring Grid | `.scoring-grid`, `.score-item` | Metrics and ratings |
| CTA Section | `.cta-section`, `.cta-btn`, `.secondary-btn` | Call-to-action blocks |
| Footer | `.footer`, `.footer-links` | Site footer |

**Responsive Breakpoints**:
- Tablet: `max-width: 768px`
- Mobile: `max-width: 480px`

**Animation Features**:
- Fade-in on scroll (`.fade-in-up`)
- Hover transitions on cards
- Gradient animations on backgrounds

---

## Mode: Analyze Template (`--analyze`)

When `--analyze tactical` (or other template) is specified:

### Deep Analysis Process

1. **Read the CSS file completely**
   ```
   Read: docs/subpage-styles.css
   ```

2. **Extract Design System Variables**
   - Color palette (primary, secondary, accent)
   - Typography scale (headings, body, small)
   - Spacing scale (margins, padding)
   - Border radius and shadows
   - Animation timings

3. **Catalog All Components**
   For each component, document:
   - Class name and variants
   - HTML structure required
   - Responsive behavior
   - Interaction states (hover, focus, active)
   - Dependencies on other components

4. **Generate Component Usage Examples**
   ```html
   <!-- Feature Grid Example -->
   <div class="feature-grid">
       <div class="feature-item">
           <div class="feature-icon">Icon</div>
           <h3>Feature Title</h3>
           <p>Feature description text</p>
       </div>
       <!-- Repeat for each feature -->
   </div>
   ```

5. **Identify Extension Points**
   - Which variables can be overridden
   - Where new components can be added
   - What patterns should be followed

6. **Accessibility Audit**
   - Color contrast ratios
   - Focus state visibility
   - Touch target sizes
   - Semantic structure requirements

### Analysis Output Format

```markdown
# CSS Template Analysis: [Template Name]

## Design System Overview
- **Primary Color**: [hex] - Used for [components]
- **Accent Color**: [hex] - Used for [components]
- **Typography**: [font-stack] at [scale]

## Component Inventory
| Component | Classes | Variants | Responsive |
|-----------|---------|----------|------------|
| ...       | ...     | ...      | ...        |

## Usage Patterns
[Code examples for each major component]

## Extension Guidelines
[How to add new components while maintaining consistency]

## Accessibility Notes
[WCAG compliance status and recommendations]
```

---

## Mode: Create Template (`--create`)

When `--create [name]` is specified:

### Template Creation Process

1. **Define Design Direction**
   Gather requirements for the new template:
   - Visual style (modern, classic, bold, subtle)
   - Industry/use case focus
   - Color scheme preferences
   - Typography preferences

2. **Generate Base CSS Structure**

   Using the **frontend-design skill**, create a new CSS file with:

   ```css
   /* ========================================
    * [Template Name] Design System
    * Generated: [Date]
    * Purpose: [Description]
    * ======================================== */

   /* Variables */
   :root {
       /* Colors */
       --bg-primary: [value];
       --bg-secondary: [value];
       --text-primary: [value];
       --text-secondary: [value];
       --accent: [value];
       --accent-hover: [value];

       /* Typography */
       --font-heading: [stack];
       --font-body: [stack];
       --font-mono: [stack];

       /* Spacing */
       --space-xs: 0.25rem;
       --space-sm: 0.5rem;
       --space-md: 1rem;
       --space-lg: 2rem;
       --space-xl: 4rem;

       /* Borders */
       --radius-sm: 4px;
       --radius-md: 8px;
       --radius-lg: 16px;

       /* Shadows */
       --shadow-sm: [value];
       --shadow-md: [value];
       --shadow-lg: [value];
   }

   /* Base Styles */
   * { box-sizing: border-box; margin: 0; padding: 0; }

   body {
       font-family: var(--font-body);
       background: var(--bg-primary);
       color: var(--text-primary);
       line-height: 1.6;
   }

   /* Navigation */
   .nav { ... }

   /* Hero Section */
   .hero { ... }

   /* Content Sections */
   .content-section { ... }

   /* Feature Grid */
   .feature-grid { ... }

   /* Cards */
   .card { ... }

   /* CTA */
   .cta-section { ... }
   .cta-btn { ... }

   /* Footer */
   .footer { ... }

   /* Responsive */
   @media (max-width: 768px) { ... }
   @media (max-width: 480px) { ... }
   ```

3. **Validate Template**
   - Test with sample landing page content
   - Verify responsive behavior
   - Check accessibility compliance
   - Preview with Playwright

4. **Save Template**
   - Location: `docs/[name]-styles.css`
   - Update template registry (if maintained)

---

## Quick Reference: Common Patterns

### Hero Section Pattern
```html
<section class="hero">
    <div class="hero-content">
        <h1 class="hero-title">Main Headline</h1>
        <p class="hero-subtitle">Supporting subheadline with value prop</p>
        <div class="hero-cta">
            <a href="#" class="cta-btn">Primary Action</a>
            <a href="#" class="secondary-btn">Secondary Action</a>
        </div>
    </div>
</section>
```

### Feature Grid Pattern
```html
<section class="content-section">
    <h2 class="section-title">Benefits Section</h2>
    <div class="feature-grid">
        <div class="feature-item">
            <div class="feature-icon">[Icon/Emoji]</div>
            <h3>Benefit One</h3>
            <p>Description of benefit</p>
        </div>
        <!-- Repeat 2-3 more items -->
    </div>
</section>
```

### CTA Section Pattern
```html
<section class="cta-section">
    <h2>Ready to Get Started?</h2>
    <p>Supporting copy that reinforces value</p>
    <a href="#" class="cta-btn">Take Action Now</a>
</section>
```

---

## Example Usage

```bash
# List all available templates
/landing-page:css-template

# View tactical template components
/landing-page:css-template --view tactical

# Deep analysis of the tactical template
/landing-page:css-template --analyze tactical

# Create a new minimal template
/landing-page:css-template --create saas-modern

# Create corporate template with specific requirements
/landing-page:css-template --create enterprise-clean
```

---

Now processing template operation: `$ARGUMENTS`...
