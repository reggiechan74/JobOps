---
description: Generate strategic landing page copy using proven copywriting frameworks and persuasion psychology
argument-hint: <page-purpose> [--tone=professional|casual|bold] [--framework=PAS|AIDA|StoryBrand]
---

I'll generate strategic landing page copy using proven copywriting frameworks and persuasion psychology.

**Arguments:**
- `$1`: Page purpose/topic (required) - e.g., `consulting-services`, `saas-product`, `event-registration`
- `--tone`: Writing tone (optional): `professional` | `casual` | `bold` (default: professional)
- `--framework`: Copywriting framework (optional): `PAS` | `AIDA` | `StoryBrand` (default: PAS)

## Copywriting Frameworks

### PAS (Problem-Agitate-Solution) - Default
Best for: Services, consulting, B2B offerings

| Stage | Purpose | Landing Page Section |
|-------|---------|---------------------|
| **Problem** | Identify the pain | Hero subtitle, Problem section |
| **Agitate** | Intensify the urgency | Statistics, consequences |
| **Solution** | Present your offering | Benefits, How It Works |

### AIDA (Attention-Interest-Desire-Action)
Best for: Products, e-commerce, direct response

| Stage | Purpose | Landing Page Section |
|-------|---------|---------------------|
| **Attention** | Hook with headline | Hero title |
| **Interest** | Build curiosity | Problem/opportunity |
| **Desire** | Create want | Benefits, social proof |
| **Action** | Drive conversion | CTA sections |

### StoryBrand
Best for: Brand building, services, complex offerings

| Element | Purpose | Landing Page Section |
|---------|---------|---------------------|
| **Hero (Customer)** | Position visitor as hero | Hero messaging |
| **Problem** | Define villain/challenge | Problem section |
| **Guide (You)** | Position as trusted guide | Credentials, about |
| **Plan** | Show clear path | How It Works |
| **Call to Action** | Direct next step | CTA sections |
| **Success/Failure** | Paint outcomes | Benefits/consequences |

---

## Phase 1: Context Gathering

### 1.1: Invoke Copywriter Agent

Launch the **landing-page-copywriter** agent for comprehensive copy development:

```
Task tool with subagent_type='landing-page-copywriter'

Prompt: Generate landing page copy for [page-purpose] using [framework] framework with [tone] tone.

Context to provide:
- Page purpose: $1
- Target audience (if known)
- Key differentiators
- Desired action/conversion goal
- Any existing brand voice guidelines
```

### 1.2: Information Required

The copywriter agent will need (will ask if not provided):

| Information | Why It Matters | Example |
|-------------|----------------|---------|
| **Who** is the target audience? | Voice, pain points, motivations | "CTOs at mid-market companies" |
| **What** problem do you solve? | Core messaging foundation | "Manual processes waste 20hrs/week" |
| **Why** should they choose you? | Differentiator, positioning | "Only solution with X feature" |
| **What** action should they take? | CTA clarity | "Book a demo" |

---

## Phase 2: Copy Development

### 2.1: Headline Formulas

The copywriter will generate headlines using proven formulas:

**Result-Based Headlines**:
- "Get [Desired Result] in [Timeframe]"
- "[Number] [Audience] Now [Achieve Result]"
- "The [Adjective] Way to [Desired Outcome]"

**Problem-Based Headlines**:
- "Stop [Pain Point] Forever"
- "Why [Common Approach] Fails (And What Works Instead)"
- "Tired of [Frustration]?"

**Curiosity Headlines**:
- "The Secret to [Desired Result]"
- "What [Successful People] Know About [Topic]"
- "How [Surprising Method] [Achieves Result]"

### 2.2: Section-by-Section Copy

**Hero Section** (3-second hook):
```
Headline: [Clear value proposition - what they get]
Subheadline: [How they get it + why it matters]
CTA: [Specific action verb + outcome]
```

**Problem Section** (Agitation):
```
Opening: [Acknowledge their current situation]
Pain Points: [3-5 specific frustrations they experience]
Stakes: [What happens if they don't solve this]
Statistics: [Data that validates the problem]
```

**Solution Section** (Your Offering):
```
Transition: [Bridge from problem to solution]
Core Promise: [What you deliver]
Key Benefits: [3-5 benefits, not features]
Proof: [Why this works]
```

**How It Works** (Remove Friction):
```
Step 1: [Simple first action]
Step 2: [What happens next]
Step 3: [Outcome they achieve]
```

**Social Proof** (Build Trust):
```
Testimonials: [Specific results + attribution]
Logos: [Recognizable clients/partners]
Metrics: [Quantified success]
Credentials: [Authority indicators]
```

**CTA Section** (Drive Action):
```
Headline: [Restate value + urgency]
Supporting: [Address final objection]
Primary CTA: [Main action button]
Secondary CTA: [Lower-commitment alternative]
Risk Reversal: [Guarantee or assurance]
```

---

## Phase 3: Copy Best Practices

### 3.1: Writing Principles

| Principle | Application |
|-----------|-------------|
| **Benefits over features** | "Save 10 hours/week" not "Automated workflows" |
| **Specificity sells** | "47% increase" not "significant improvement" |
| **You > We** | "You'll achieve" not "We provide" |
| **Active voice** | "Transform your process" not "Your process is transformed" |
| **One idea per sentence** | Short, punchy, scannable |
| **Power words** | Free, New, Proven, Guaranteed, Instant, Secret |

### 3.2: Tone Calibration

**Professional** (default):
- Authoritative but approachable
- Data-driven claims
- Industry terminology appropriate
- Formal but not stiff

**Casual**:
- Conversational, friendly
- Contractions allowed
- Personality and humor
- Relatable examples

**Bold**:
- Strong claims with proof
- Direct, challenging
- Urgency and scarcity
- Polarizing (attracts right audience)

### 3.3: Conversion Psychology

| Principle | Implementation |
|-----------|----------------|
| **Social Proof** | "Join 5,000+ companies" |
| **Authority** | Credentials, logos, certifications |
| **Scarcity** | "Limited spots available" |
| **Urgency** | "Offer expires Friday" |
| **Reciprocity** | Free resource, valuable content |
| **Loss Aversion** | "Don't miss out on..." |

---

## Phase 4: Deliverables

### 4.1: Copy Document Structure

The copywriter agent will deliver a structured copy document:

```markdown
# Landing Page Copy: [Page Name]

## Meta Information
- **Title Tag**: [60 characters max]
- **Meta Description**: [155 characters max]

## Hero Section
- **Headline**: [Primary headline]
- **Subheadline**: [Supporting copy]
- **Primary CTA**: [Button text]
- **Secondary CTA**: [Alternative action]

## Problem Section
- **Section Title**: [Problem hook]
- **Opening Copy**: [Pain acknowledgment]
- **Pain Points**:
  1. [Pain point 1]
  2. [Pain point 2]
  3. [Pain point 3]
- **Stakes Copy**: [What's at risk]

## Solution Section
- **Section Title**: [Solution hook]
- **Core Promise**: [What you deliver]
- **Benefits**:
  1. **[Benefit 1 Title]**: [Description]
  2. **[Benefit 2 Title]**: [Description]
  3. **[Benefit 3 Title]**: [Description]

## How It Works
- **Section Title**: [Process hook]
- **Step 1**: [Title] - [Description]
- **Step 2**: [Title] - [Description]
- **Step 3**: [Title] - [Description]

## Social Proof
- **Section Title**: [Credibility hook]
- **Testimonials**:
  - "[Quote]" - [Name, Title, Company]
- **Metrics**: [Key statistics]

## Final CTA
- **Headline**: [Action-oriented]
- **Supporting Copy**: [Final value prop]
- **Primary CTA**: [Main button]
- **Risk Reversal**: [Guarantee copy]

## Alternative Copy Versions
[A/B test variations for key elements]
```

### 4.2: Output Location

Copy documents are saved to:
- Default: `OutputResumes/LandingPageCopy_[PageName]_[Date].md`
- Or provide to `/landing-page:create` command directly

---

## Example Usage

```bash
# Professional consulting services copy using PAS
/landing-page:copywrite consulting-services

# Casual SaaS product copy using AIDA
/landing-page:copywrite saas-onboarding --tone=casual --framework=AIDA

# Bold workshop registration using StoryBrand
/landing-page:copywrite ai-workshop --tone=bold --framework=StoryBrand

# Generate copy for existing page improvement
/landing-page:copywrite services-page-refresh --tone=professional
```

---

Now generating landing page copy for `$1`...
