# Assessment Rubric Framework Template v2.2

**DO NOT USE THIS FILE DIRECTLY** - This is a meta-template for generating job-specific rubrics.

## Purpose

This template defines the canonical **200-point deduplicated scoring structure** with **role-based weight variants** for candidate assessments. Each assessment criterion appears in ONE location only to prevent double-counting. Skill scoring is PROFICIENCY-BASED, not years-based.

**v2.2 CRITICAL UPDATE:** All scoring criteria now use **discrete point-level anchors** - no range-bound scoring. Every point value has its own unique, distinguishable criteria to ensure inter-rater reliability.

## Changelog

| Version | Date | Changes |
|---------|------|---------|
| 2.2 | 2026-02-04 | Eliminated all range-bound scoring; every point level now has discrete anchor criteria |
| 2.1 | 2026-02-04 | Added Evidence Verification Protocol (domain specificity, experience classification, citation requirements) |
| 2.0 | 2026-02-04 | Initial 200-point framework with role-based weight variants |

## Commands That Reference This Template

- `/createrubric` - Creates standalone job-specific rubrics
- `/assessjob` - Creates rubric as part of full assessment workflow
- `/assesscandidate` - Applies existing rubric to candidate evaluation

---

# TEMPLATE BEGINS HERE

---

# Scoring Rubric: [Role Title] at [Company]

**Generated:** [Date]
**Job Posting:** [Job Posting File Path]
**Rubric Version:** 2.2

---

## ALIGNMENT STATEMENT

**MANDATORY: Complete this section before scoring**

| Dimension | Job Requirement | Rubric Coverage |
|-----------|-----------------|-----------------|
| Role Type | [Technical IC / People Manager / Executive] | Weight variant applied: [variant] |
| Primary Function | [Core job function] | Category 1-3 alignment |
| Industry Context | [Industry/sector] | Category 2 alignment |
| Success Metrics | [How success is measured] | Category 3 alignment |
| Critical Skills | [Top 3 must-haves] | Category 1 alignment |

**Alignment Confirmation:** This rubric has been customized to assess [Role Title] candidates against [Company]'s specific requirements as stated in the job posting dated [Date].

---

## WEIGHT VARIANT SELECTION

**MANDATORY: Select ONE weight variant based on role type**

### Variant Selection Guide

| Role Indicators | Select Variant |
|-----------------|----------------|
| Deep technical expertise required, individual contributor, architect/engineer titles | **Technical IC** |
| Direct reports, team leadership, coaching responsibilities, manager/director titles | **People Manager** |
| Strategic decision-making, P&L ownership, C-suite/VP titles, board interaction | **Executive** |

### Weight Allocation Table

| Category | Default | Technical IC | People Manager | Executive |
|----------|---------|--------------|----------------|-----------|
| 1. Skills Inventory | 25% (50 pts) | **30% (60 pts)** | 20% (40 pts) | 15% (30 pts) |
| 2. Experience Relevance | 20% (40 pts) | 20% (40 pts) | 20% (40 pts) | **25% (50 pts)** |
| 3. Demonstrated Impact | 30% (60 pts) | 25% (50 pts) | 25% (50 pts) | **35% (70 pts)** |
| 4. Credentials | 10% (20 pts) | 10% (20 pts) | 10% (20 pts) | 5% (10 pts) |
| 5. Fit & Readiness | 15% (30 pts) | 15% (30 pts) | **25% (50 pts)** | 20% (40 pts) |
| **TOTAL** | **200 pts** | **200 pts** | **200 pts** | **200 pts** |

**Selected Variant:** [Technical IC / People Manager / Executive / Default]

### Weight Justification Table

**MANDATORY: Document why this variant was selected**

| Justification Factor | Evidence from Job Posting |
|---------------------|---------------------------|
| Primary Role Function | [Quote or paraphrase from posting] |
| Reporting Structure | [Reports to X, manages Y] |
| Success Metrics Type | [Technical deliverables / Team outcomes / Business results] |
| Decision Authority Level | [Individual / Team / Strategic] |

---

## CRITICAL BARRIERS TABLE

**MANDATORY: Identify disqualifying factors BEFORE detailed scoring**

| Barrier Type | Threshold | Candidate Status | Action |
|--------------|-----------|------------------|--------|
| Required Certification/License | [Specific cert from posting] | [ ] Has / [ ] Lacks | If lacks: STOP - Disqualified |
| Minimum Education | [Degree requirement] | [ ] Meets / [ ] Below | If below: Document waiver rationale |
| Work Authorization | [Location/visa requirements] | [ ] Eligible / [ ] Ineligible | If ineligible: STOP - Disqualified |
| Security Clearance | [If applicable] | [ ] Has / [ ] Cannot Obtain | If cannot obtain: STOP - Disqualified |
| Non-Negotiable Skill | [Critical skill #1] | [ ] Demonstrated / [ ] Absent | If absent: Score Category 1 first |
| Non-Negotiable Skill | [Critical skill #2] | [ ] Demonstrated / [ ] Absent | If absent: Score Category 1 first |

**Barrier Review Result:** [ ] Proceed with full assessment / [ ] Candidate disqualified - Reason: ___________

---

## EVIDENCE VERIFICATION PROTOCOL

**MANDATORY: Apply this protocol throughout scoring**

### Evidence Citation Requirements

**CRITICAL:** Before finalizing any score above "Basic" (1 point):
- Quote specific CV/resume text supporting the score
- Identify exact source section and line numbers where possible
- Distinguish between **direct experience** vs **transferable skills**
- Verify claimed experience duration and recency
- Flag any claims that cannot be verified from provided materials

### Domain Specificity Verification

**NEVER assume domain equivalency without explicit evidence.** Common false equivalencies:

| Claimed Domain | ≠ | Related But Different Domain |
|----------------|---|------------------------------|
| Real estate investment | ≠ | Real estate asset management |
| Commercial real estate | ≠ | Residential real estate |
| Institutional investment | ≠ | Brokerage/advisory |
| Due diligence (M&A/investment) | ≠ | General analysis |
| Software engineering | ≠ | IT support/administration |
| Product management | ≠ | Project management |
| Data science/ML | ≠ | Data analytics/BI |
| Strategic planning | ≠ | Operational planning |
| P&L ownership | ≠ | Budget management |
| People leadership | ≠ | Project coordination |

**When in doubt:** Score as "Adjacent" or "Transferable" rather than "Direct Match"

### Experience Type Classification

**MANDATORY: Classify ALL experience claims using these categories:**

| Classification | Definition | Scoring Impact |
|----------------|------------|----------------|
| **Direct** | Exact same function, industry, and scope | Full points if proficiency demonstrated |
| **Adjacent** | Same function, different industry OR same industry, related function | Maximum 80% of points; note transfer gap |
| **Transferable** | Different function and industry but applicable skills | Maximum 60% of points; flag for interview verification |
| **Assumed** | No explicit evidence; assessor inference | Maximum "Basic" (1 pt) score; flag as LOW confidence |

### Verification Checklist

Before finalizing Category 1-3 scores:

- [ ] Each skill score has quoted evidence from source materials
- [ ] Domain-specific claims verified against classification table
- [ ] Experience type (Direct/Adjacent/Transferable/Assumed) documented
- [ ] Recency of relevant experience confirmed
- [ ] No scores above Basic without direct textual support
- [ ] Ambiguous claims flagged for interview verification

---

## CATEGORY 1: SKILLS INVENTORY

**Default: 50 points | Adjusted: [X] points per selected variant**

### 1A. Required Technical Skills (36 points base)

**MANDATORY: Score 6 required skills using PROFICIENCY scale (0-6 each)**

#### Proficiency Scale with Discrete Anchors

| Score | Level | Definition | Behavioral Anchors |
|-------|-------|------------|-------------------|
| 6 | Thought Leader | Recognized authority who shapes industry practice | Published books/papers on topic, keynote speaker at major conferences, created frameworks adopted industry-wide, serves on standards bodies, quoted in trade publications |
| 5 | Expert | Handles complex/novel scenarios, mentors others | Solves problems others cannot, trains/certifies practitioners, trusted for highest-stakes decisions, designs systems adopted by multiple teams, recognized internally as go-to authority |
| 4 | Advanced | Complex work with minimal oversight | Independently handles difficult/unusual cases, identifies process improvements proactively, reviews others' work, requires no supervision for complex tasks |
| 3 | Proficient | Solid independent execution | Reliably delivers standard work independently, handles typical complexity without assistance, meets quality standards consistently |
| 2 | Developing | Effective with guidance | Productive with periodic check-ins, handles routine tasks independently, asks appropriate questions, growing capability evident |
| 1 | Basic | Foundational knowledge, needs support | Understands concepts and terminology, requires supervision for most tasks, can assist but not lead |
| 0 | None | No demonstrated capability | No evidence in any materials, no mention of skill or related activities |

#### Required Skills Assessment

| # | Skill | Job Requirement | Proficiency Score (0-6) | Evidence Citation | Experience Type | Confidence |
|---|-------|-----------------|------------------------|-------------------|-----------------|------------|
| 1 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [D/A/T/As] | [H/M/L] |
| 2 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [D/A/T/As] | [H/M/L] |
| 3 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [D/A/T/As] | [H/M/L] |
| 4 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [D/A/T/As] | [H/M/L] |
| 5 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [D/A/T/As] | [H/M/L] |
| 6 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [D/A/T/As] | [H/M/L] |

**Required Skills Subtotal:** _____ / 36

### 1B. Preferred Skills (14 points base)

**MANDATORY: Score preferred/nice-to-have skills using same proficiency scale, points allocated proportionally**

**Allocation Method:** Divide 14 points across preferred skills based on job posting emphasis. Each skill uses the 0-6 proficiency scale, normalized to allocated points.

| # | Preferred Skill | Points Allocated | Proficiency (0-6) | Normalized Score | Evidence Citation | Confidence |
|---|-----------------|------------------|-------------------|------------------|-------------------|------------|
| 1 | [Skill from posting] | [X pts] | [ ] | [ ] | [Source] | [H/M/L] |
| 2 | [Skill from posting] | [X pts] | [ ] | [ ] | [Source] | [H/M/L] |
| 3 | [Skill from posting] | [X pts] | [ ] | [ ] | [Source] | [H/M/L] |

**Normalization Formula:** (Proficiency Score / 6) × Points Allocated = Normalized Score

**Preferred Skills Subtotal:** _____ / 14

### Category 1 Total: _____ / 50 (or adjusted variant total)

**Normalized for variant:** _____ / [Variant Points]

---

## CATEGORY 2: EXPERIENCE RELEVANCE

**Default: 40 points | Adjusted: [X] points per selected variant**

### 2A. Industry/Domain Alignment (20 points)

#### 2A-1. Primary Industry Match (10 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Example |
|-------|-------------------|---------|
| 10 | Exact industry AND exact sub-sector match | Target: SaaS B2B HR Tech → Candidate: SaaS B2B HR Tech |
| 9 | Exact industry, adjacent sub-sector | Target: SaaS B2B HR Tech → Candidate: SaaS B2B Recruiting Tech |
| 8 | Exact industry, different sub-sector | Target: SaaS B2B HR Tech → Candidate: SaaS B2B Finance Tech |
| 7 | Adjacent industry, same customer segment | Target: SaaS B2B → Candidate: Enterprise software (on-prem) B2B |
| 6 | Adjacent industry, overlapping business model | Target: SaaS B2B → Candidate: PaaS B2B |
| 5 | Same broad sector, different business model | Target: SaaS B2B → Candidate: IT Services consulting |
| 4 | Related sector with operational overlap | Target: SaaS B2B → Candidate: Hardware tech company |
| 3 | Tangentially related, some transferable context | Target: SaaS B2B → Candidate: Digital media B2B |
| 2 | Different industry, shared customer type only | Target: SaaS B2B → Candidate: Manufacturing B2B |
| 1 | Different industry, transferable business functions only | Target: SaaS B2B → Candidate: Retail (shared: sales, operations) |
| 0 | No industry relevance or transferability | Target: SaaS B2B → Candidate: Non-profit social services |

**Score:** [ ] / 10
**Evidence:** [Company/role from resume demonstrating match level]
**Experience Type:** [Direct/Adjacent/Transferable/Assumed]

#### 2A-2. Domain Expertise (6 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 6 | Deep domain expert - recognized authority in specific domain | Published domain expertise, domain certifications, domain-specific leadership roles, invited domain speaker |
| 5 | Strong domain practitioner - extensive hands-on domain experience | 5+ years domain-specific work, led domain initiatives, domain mentioned in multiple roles |
| 4 | Solid domain knowledge - competent domain performer | 3-5 years domain work, domain responsibilities clearly stated, domain outcomes documented |
| 3 | Working domain knowledge - functional in domain | 1-3 years domain exposure, domain tasks performed (not led), domain mentioned in job duties |
| 2 | Domain familiarity - exposure without ownership | Domain mentioned peripherally, worked alongside domain experts, domain training completed |
| 1 | Domain awareness - conceptual understanding only | Domain referenced in education or coursework, no professional domain experience |
| 0 | No domain experience or knowledge evident | No mention of domain in any materials |

**Score:** [ ] / 6
**Evidence:** [Specific domain experience cited]
**Experience Type:** [Direct/Adjacent/Transferable/Assumed]

#### 2A-3. Market/Segment Knowledge (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Exact market segment match with depth | Same customer size, geography, buying behavior; market expertise demonstrated |
| 3 | Same market segment, different geography or adjacent customer size | Enterprise → Mid-market, or US → Canada in same segment |
| 2 | Related market segment with transferable dynamics | B2B enterprise → B2B SMB, or adjacent verticals within segment |
| 1 | Different segment, some market mechanics transfer | B2B → B2C with complex sales, or different vertical entirely |
| 0 | No market segment relevance | Completely different market dynamics, no transferable knowledge |

**Score:** [ ] / 4
**Evidence:** [Market experience cited]
**Experience Type:** [Direct/Adjacent/Transferable/Assumed]

**Industry/Domain Subtotal:** _____ / 20

### 2B. Career Trajectory & Growth (12 points)

#### 2B-1. Progression Pattern (6 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Pattern |
|-------|-------------------|------------------|
| 6 | Exceptional upward trajectory - consistent promotions with expanding scope and title advancement | Promoted every 2-3 years, each role larger than previous, clear title progression (Analyst→Manager→Director) |
| 5 | Strong upward trajectory - regular advancement with increasing responsibility | Promoted every 3-4 years, scope increased (team size, budget, geography), titles reflect growth |
| 4 | Steady advancement - progression evident with some scope increase | Promotions present but not rapid, responsibility grew over time, logical career progression |
| 3 | Moderate progression - lateral moves with skill expansion | Moved across functions/companies at similar level, but acquired new capabilities each move |
| 2 | Limited progression - some growth but inconsistent | 1-2 promotions over long career, or recent progression after static period |
| 1 | Minimal progression - largely static career level | Same level roles throughout, no evidence of increased responsibility |
| 0 | Regression or stagnation - downward or flat trajectory | Demotions, reduced scope, or unexplained gaps with lower re-entry |

**Score:** [ ] / 6
**Evidence:** [Title progression from resume]

#### 2B-2. Scope Expansion (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Consistent scope expansion across multiple dimensions | Budget, team, geography, and complexity all increased over career; evidence of scaling |
| 3 | Scope expansion in 2-3 dimensions | Clear growth in team size AND budget, OR geography AND complexity |
| 2 | Scope expansion in 1 dimension | Growth in one area (e.g., team grew from 5 to 15, but budget/geo static) |
| 1 | Minor scope changes | Small increases that don't represent meaningful expansion |
| 0 | No scope expansion or scope contraction | Same or smaller scope throughout career |

**Score:** [ ] / 4
**Evidence:** [Budget/team/geo growth documented]

#### 2B-3. Role Stability (2 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Pattern |
|-------|-------------------|------------------|
| 2 | Appropriate tenure demonstrating commitment | Average 2-5 years per role, departures explainable (acquisition, relocation, growth opportunity) |
| 1 | Borderline tenure with reasonable explanation | Average 1-2 years, or one short stint among longer tenures, explanations plausible |
| 0 | Job-hopping pattern raising concerns | Multiple roles under 1 year, no long tenures, pattern suggests performance or fit issues |

**Score:** [ ] / 2
**Evidence:** [Tenure analysis]

**Career Trajectory Subtotal:** _____ / 12

### 2C. Recency of Relevant Work (8 points)

#### 2C-1. Current/Last Role Relevance (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Current role is highly relevant - doing this exact work now | Job duties match target role 80%+, same function, same industry context |
| 3 | Current role is substantially relevant - significant overlap | Job duties match target role 50-80%, same function OR same industry |
| 2 | Current role is partially relevant - some transferable work | Job duties match target role 25-50%, related function or industry |
| 1 | Current role has minimal relevance - tangential connection | Job duties match target role <25%, but some transferable skills evident |
| 0 | Current role not relevant - no meaningful connection | Current work has no relationship to target role |

**Score:** [ ] / 4
**Evidence:** [Current role duties cited]

#### 2C-2. Recent History Relevance (2-5 years ago) (3 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 3 | Strong relevant experience in recent history | 1+ roles in 2-5 year window directly relevant to target |
| 2 | Moderate relevant experience in recent history | Partial relevance in 2-5 year window, or short relevant stint |
| 1 | Limited relevant experience in recent history | Tangential relevance only in 2-5 year window |
| 0 | No relevant experience in recent history | Nothing in 2-5 year window relates to target role |

**Score:** [ ] / 3
**Evidence:** [Recent roles cited]

#### 2C-3. Earlier Career Relevance (5+ years ago) (1 point)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 1 | Foundational relevant experience exists | Relevant experience 5+ years ago provides foundation, even if not recent |
| 0 | No earlier relevant experience | Nothing in earlier career relates to target role |

**Score:** [ ] / 1
**Evidence:** [Earlier experience cited]

**Recency Subtotal:** _____ / 8

### Category 2 Total: _____ / 40 (or adjusted variant total)

---

## CATEGORY 3: DEMONSTRATED IMPACT

**Default: 60 points | Adjusted: [X] points per selected variant**

**NOTE: This is the highest-weighted category (30% default). Evidence must be specific and quantified.**

### 3A. Quantified Achievements (30 points)

#### 3A-1. Financial Impact (10 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 10 | Transformational financial impact >$50M | Documented $50M+ revenue, savings, or value creation with clear attribution |
| 9 | Major financial impact $25M-$50M | Documented $25M-$50M impact with clear attribution |
| 8 | Significant financial impact $10M-$25M | Documented $10M-$25M impact with clear attribution |
| 7 | Substantial financial impact $5M-$10M | Documented $5M-$10M impact with clear attribution |
| 6 | Strong financial impact $2M-$5M | Documented $2M-$5M impact with clear attribution |
| 5 | Solid financial impact $1M-$2M | Documented $1M-$2M impact with clear attribution |
| 4 | Meaningful financial impact $500K-$1M | Documented $500K-$1M impact with clear attribution |
| 3 | Moderate financial impact $250K-$500K | Documented $250K-$500K impact with clear attribution |
| 2 | Notable financial impact $100K-$250K | Documented $100K-$250K impact with clear attribution |
| 1 | Some financial impact <$100K documented | Documented impact under $100K, or vague financial claims |
| 0 | No financial metrics provided | No quantified financial impact in materials |

**Score:** [ ] / 10
**Evidence:** [Exact claim from resume]
**Confidence:** [H/M/L]

#### 3A-2. Efficiency/Quality Gains (10 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 10 | Transformational efficiency gain >75% improvement | Documented 75%+ improvement in speed, quality, or efficiency metric |
| 9 | Exceptional efficiency gain 60-75% improvement | Documented 60-75% improvement with clear baseline and outcome |
| 8 | Major efficiency gain 50-60% improvement | Documented 50-60% improvement with clear baseline and outcome |
| 7 | Significant efficiency gain 40-50% improvement | Documented 40-50% improvement with clear baseline and outcome |
| 6 | Strong efficiency gain 30-40% improvement | Documented 30-40% improvement with clear baseline and outcome |
| 5 | Solid efficiency gain 20-30% improvement | Documented 20-30% improvement with clear baseline and outcome |
| 4 | Meaningful efficiency gain 15-20% improvement | Documented 15-20% improvement with clear baseline and outcome |
| 3 | Moderate efficiency gain 10-15% improvement | Documented 10-15% improvement with clear baseline and outcome |
| 2 | Notable efficiency gain 5-10% improvement | Documented 5-10% improvement with clear baseline and outcome |
| 1 | Some efficiency gain <5% documented | Documented improvement under 5%, or vague efficiency claims |
| 0 | No efficiency/quality metrics provided | No quantified efficiency or quality impact in materials |

**Score:** [ ] / 10
**Evidence:** [Exact claim from resume]
**Confidence:** [H/M/L]

#### 3A-3. Growth Metrics (10 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 10 | Exceptional growth >200% | Documented 200%+ growth (revenue, users, market share, etc.) with attribution |
| 9 | Outstanding growth 150-200% | Documented 150-200% growth with clear attribution |
| 8 | Excellent growth 100-150% | Documented 100-150% growth with clear attribution |
| 7 | Strong growth 75-100% | Documented 75-100% growth with clear attribution |
| 6 | Solid growth 50-75% | Documented 50-75% growth with clear attribution |
| 5 | Good growth 35-50% | Documented 35-50% growth with clear attribution |
| 4 | Moderate growth 25-35% | Documented 25-35% growth with clear attribution |
| 3 | Steady growth 15-25% | Documented 15-25% growth with clear attribution |
| 2 | Modest growth 10-15% | Documented 10-15% growth with clear attribution |
| 1 | Minimal growth <10% documented | Documented growth under 10%, or vague growth claims |
| 0 | No growth metrics provided | No quantified growth impact in materials |

**Score:** [ ] / 10
**Evidence:** [Exact claim from resume]
**Confidence:** [H/M/L]

**Quantified Achievements Subtotal:** _____ / 30

### 3B. Scale & Complexity Managed (18 points)

#### 3B-1. Budget/P&L Responsibility (6 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 6 | Enterprise-scale budget >$100M | Full P&L or budget ownership exceeding $100M documented |
| 5 | Large-scale budget $50M-$100M | Full P&L or budget ownership $50M-$100M documented |
| 4 | Significant budget $25M-$50M | Budget responsibility $25M-$50M documented |
| 3 | Substantial budget $10M-$25M | Budget responsibility $10M-$25M documented |
| 2 | Moderate budget $5M-$10M | Budget responsibility $5M-$10M documented |
| 1 | Small budget $1M-$5M | Budget responsibility $1M-$5M documented |
| 0 | No budget responsibility or <$1M | No documented budget ownership, or individual contributor only |

**Score:** [ ] / 6
**Evidence:** [Budget amount cited]

#### 3B-2. Team/Organization Size (6 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 6 | Large organization >100 direct + indirect reports | Managed organization of 100+ people (direct + matrix/indirect) |
| 5 | Substantial organization 50-100 reports | Managed 50-100 people (direct + matrix/indirect) |
| 4 | Mid-size team 25-50 reports | Managed 25-50 people (direct + matrix/indirect) |
| 3 | Growing team 10-25 reports | Managed 10-25 people (direct + matrix/indirect) |
| 2 | Small team 5-10 reports | Managed 5-10 direct reports |
| 1 | Minimal team 2-4 reports | Managed 2-4 direct reports or led project teams without direct reports |
| 0 | Individual contributor | No direct reports, no team leadership responsibility |

**Score:** [ ] / 6
**Evidence:** [Team size cited]

#### 3B-3. Geographic/Complexity Scope (6 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 6 | Global scope - multiple continents | Responsibilities spanning 3+ continents, global strategy ownership |
| 5 | International scope - multiple countries | Responsibilities spanning multiple countries across 2 continents |
| 4 | Multi-national scope - 2+ countries | Responsibilities in 2+ countries, cross-border coordination |
| 3 | National scope - multiple regions | Country-wide responsibility with multiple regional offices/operations |
| 2 | Regional scope - multiple locations | Multi-site responsibility within a region (e.g., Northeast US, Western Canada) |
| 1 | Multi-site local scope | 2-3 locations in same metro area or adjacent areas |
| 0 | Single location | Single site responsibility, no geographic complexity |

**Score:** [ ] / 6
**Evidence:** [Scope indicators cited]

**Scale & Complexity Subtotal:** _____ / 18

### 3C. Innovation & Initiative (12 points)

#### 3C-1. Process/Product Innovation (6 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 6 | Industry-level innovation - created methods/products adopted beyond organization | Innovation adopted by competitors, became industry standard, or commercialized externally |
| 5 | Organization-wide innovation - created approaches adopted across company | Innovation scaled to multiple departments/business units with documented adoption |
| 4 | Department-wide innovation - new methods/products with measurable impact | Innovation improved department metrics significantly, documented outcomes |
| 3 | Team-level innovation - improvements beyond incremental | Created new approaches that changed how team operates, documented benefits |
| 2 | Incremental innovation - meaningful improvements to existing processes | Optimized existing processes with documented efficiency gains |
| 1 | Minor improvements - small optimizations | Made small improvements, participated in improvement efforts led by others |
| 0 | No innovation evidence | No evidence of process improvement or innovation in materials |

**Score:** [ ] / 6
**Evidence:** [Innovation example cited]

#### 3C-2. Initiative & Ownership (6 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 6 | Transformational initiative - identified and solved major organizational problem proactively | Self-initiated project that fundamentally changed business outcomes, executive visibility |
| 5 | Strategic initiative - led significant improvement beyond job scope | Took ownership of cross-functional problem without being asked, delivered major results |
| 4 | Substantial initiative - drove meaningful change proactively | Identified opportunity and led effort to address it, documented impact |
| 3 | Moderate initiative - contributed leadership to improvement efforts | Actively participated in improvement initiatives, took ownership of components |
| 2 | Some initiative - participated in improvements when asked | Contributed to improvement efforts initiated by others |
| 1 | Limited initiative - primarily executed assigned work | Met job requirements but little evidence of going beyond scope |
| 0 | No initiative evidence | Only executed specifically assigned tasks, no evidence of proactive contribution |

**Score:** [ ] / 6
**Evidence:** [Initiative example cited]

**Innovation & Initiative Subtotal:** _____ / 12

### Category 3 Total: _____ / 60 (or adjusted variant total)

---

## CATEGORY 4: CREDENTIALS

**Default: 20 points | Adjusted: [X] points per selected variant**

### 4A. Education (12 points)

#### 4A-1. Degree Level (8 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 8 | Exceeds requirement by 2+ levels | PhD when Bachelor's required, or Doctorate + additional graduate degrees |
| 7 | Exceeds requirement by 1 level with distinction | Master's when Bachelor's required, from top-tier institution or with honors |
| 6 | Exceeds requirement by 1 level | Master's when Bachelor's required, from accredited institution |
| 5 | Meets requirement with distinction | Required degree from top-tier institution, honors, or additional relevant credentials |
| 4 | Meets exact requirement | Has the degree specified in job posting from accredited institution |
| 3 | Below requirement with strong compensating experience | One level below requirement, but 10+ years relevant experience documented |
| 2 | Below requirement with some compensating factors | One level below requirement, with some compensating experience or credentials |
| 1 | Below requirement with minimal compensation | Does not meet requirement, minimal compensating factors |
| 0 | Does not meet minimum and no compensation | Significantly below requirement with no compensating credentials |

**Score:** [ ] / 8
**Evidence:** [Degree, Institution]

#### 4A-2. Field Relevance (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Exact field match | Degree field matches job requirement exactly (e.g., CS degree for software role) |
| 3 | Closely related field | Degree in adjacent field (e.g., Math degree for data science role) |
| 2 | Somewhat related field | Degree has partial relevance (e.g., Physics for engineering role) |
| 1 | Tangentially related field | Degree provides some foundational knowledge (e.g., Business for technical PM) |
| 0 | Unrelated field | Degree field has no relevance to role requirements |

**Score:** [ ] / 4
**Evidence:** [Major/Concentration]

**Education Subtotal:** _____ / 12

### 4B. Certifications (8 points)

#### Certification Scoring Structure

**Required Certifications (up to 6 points):**

| Certification | Requirement | Status | Points | Evidence |
|---------------|-------------|--------|--------|----------|
| [Required Cert 1] | Required | [ ] | /3 | [Details] |
| [Required Cert 2] | Required | [ ] | /3 | [Details] |

**Required Certification Scoring (per cert):**

| Score | Anchor Definition |
|-------|-------------------|
| 3 | Current and valid certification |
| 2 | Expired within 2 years OR actively in progress (exam scheduled) |
| 1 | Expired 2+ years OR stated intent to pursue |
| 0 | Not held and no plans to obtain |

**Preferred Certifications (up to 2 points):**

| Score | Anchor Definition |
|-------|-------------------|
| 2 | 2+ preferred certifications current |
| 1 | 1 preferred certification current, OR 2+ expired/in progress |
| 0 | No preferred certifications |

**Certifications Subtotal:** _____ / 8

### Category 4 Total: _____ / 20 (or adjusted variant total)

---

## CATEGORY 5: FIT & READINESS

**Default: 30 points | Adjusted: [X] points per selected variant**

### 5A. Communication Evidence (12 points)

#### 5A-1. Written Communication (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Executive-level written communication demonstrated | Published author, board/investor materials created, executive communications role |
| 3 | Strong professional writing with senior audience | Documented experience writing for leadership, strategy documents, external communications |
| 2 | Solid professional writing skills | Resume demonstrates clear writing, documented report/proposal writing |
| 1 | Basic professional writing | Resume readable but not polished, limited evidence of writing responsibilities |
| 0 | Writing quality concerns | Resume has errors, no evidence of professional writing, or ESL concerns for English-required role |

**Score:** [ ] / 4
**Evidence:** [Writing evidence cited]

#### 5A-2. Presentation/Speaking (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | High-profile speaking demonstrated | Conference keynotes, board presentations, media appearances documented |
| 3 | Regular senior stakeholder presentations | Documented experience presenting to executives, client presentations, all-hands |
| 2 | Professional presentation experience | Training delivery, team presentations, customer meetings documented |
| 1 | Limited presentation evidence | Some indication of presentation responsibilities, not documented |
| 0 | No presentation evidence | No evidence of presentation or public speaking experience |

**Score:** [ ] / 4
**Evidence:** [Presentation evidence cited]

#### 5A-3. Cross-Functional Collaboration (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Led complex multi-stakeholder initiatives | Documented leadership of initiatives spanning 3+ departments/functions |
| 3 | Regular cross-functional project leadership | Led initiatives involving 2-3 departments, stakeholder management documented |
| 2 | Participated in cross-functional work | Member of cross-functional teams, collaboration with other departments documented |
| 1 | Limited cross-functional exposure | Some interaction with other departments mentioned |
| 0 | Siloed experience | No evidence of working outside own department/function |

**Score:** [ ] / 4
**Evidence:** [Collaboration evidence cited]

**Communication Subtotal:** _____ / 12

### 5B. Cultural Alignment (10 points)

#### 5B-1. Values Alignment (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Strong values match with evidence | Candidate materials explicitly demonstrate company values (e.g., volunteer work for "community" value) |
| 3 | Likely values alignment | Experience and career choices suggest compatible values |
| 2 | Neutral/unclear values fit | Insufficient information to assess, no red flags |
| 1 | Potential values friction | Some indicators of different priorities or values |
| 0 | Values misalignment evident | Clear evidence of conflicting values or priorities |

**Score:** [ ] / 4
**Evidence:** [Values evidence cited]

#### 5B-2. Work Environment Fit (3 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 3 | Perfect environment match | Candidate's stated/demonstrated preference matches role (remote/hybrid/onsite) |
| 2 | Acceptable environment fit | Candidate can adapt to environment, no concerns expressed |
| 1 | Minor environment friction possible | Some adjustment needed (e.g., has been remote, role is hybrid) |
| 0 | Environment mismatch | Candidate preference conflicts with role requirements |

**Score:** [ ] / 3
**Evidence:** [Work style evidence cited]

#### 5B-3. Adaptability/Growth Mindset (3 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 3 | Strong adaptability evidence | Career pivots, new skill acquisition, thrived in changing environments documented |
| 2 | Some adaptability indicators | Career shows some adaptation, learning new technologies/methods evident |
| 1 | Limited adaptability evidence | Stable career with limited change, unclear how candidate handles ambiguity |
| 0 | Rigidity concerns | Very narrow experience, resistance to change indicated |

**Score:** [ ] / 3
**Evidence:** [Adaptability examples cited]

**Cultural Alignment Subtotal:** _____ / 10

### 5C. Role-Specific Readiness (8 points)

#### 5C-1. Ramp-Up Time (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Immediately productive - has done this exact work | Current role is nearly identical, same industry, same tools, same function |
| 3 | Quick ramp - highly similar work | Very similar role, minor adjustment needed (different company, same function) |
| 2 | Moderate ramp - transferable work | Related experience, will need to learn some new aspects |
| 1 | Extended ramp - significant learning needed | Relevant foundation but substantial new skills required |
| 0 | Long ramp - starting from adjacent base | Career change or significant stretch role |

**Score:** [ ] / 4
**Evidence:** [Readiness indicators cited]

#### 5C-2. Role-Specific Requirements (4 points)

**Discrete Scoring Anchors:**

| Score | Anchor Definition | Evidence Required |
|-------|-------------------|-------------------|
| 4 | Meets all role-specific requirements | Travel, hours, location, physical requirements all confirmed or clearly compatible |
| 3 | Meets most requirements, minor clarifications needed | Most requirements met, one small area to confirm |
| 2 | Meets core requirements, some flexibility needed | Core requirements met, but some accommodation or adjustment needed |
| 1 | Meets some requirements, significant gaps | Several requirements unconfirmed or require negotiation |
| 0 | Requirement conflicts evident | Clear conflicts with stated requirements (location, travel, hours) |

**Score:** [ ] / 4
**Evidence:** [Specific requirement evidence cited]

**Role-Specific Readiness Subtotal:** _____ / 8

### Category 5 Total: _____ / 30 (or adjusted variant total)

---

## SCORING SUMMARY

### Raw Score Compilation

| Category | Possible Points | Raw Score | Weighted Points (if variant applied) |
|----------|----------------|-----------|-------------------------------------|
| 1. Skills Inventory | [Base/Variant] | _____ | _____ |
| 2. Experience Relevance | [Base/Variant] | _____ | _____ |
| 3. Demonstrated Impact | [Base/Variant] | _____ | _____ |
| 4. Credentials | [Base/Variant] | _____ | _____ |
| 5. Fit & Readiness | [Base/Variant] | _____ | _____ |
| **TOTAL** | **200** | **_____** | **_____** |

### Normalized Score

**Normalized Score = (Raw Score / 200) x 100% = _____% **

### Score Interpretation Bands

| Band | Percentage | Interpretation | Typical Action |
|------|------------|----------------|----------------|
| Exceptional | 90-100% (180-200 pts) | Exceeds requirements significantly | Priority hire, expedite process |
| Excellent | 80-89% (160-179 pts) | Strong match, minor gaps if any | Proceed to final interviews |
| Good | 70-79% (140-159 pts) | Solid match, some development areas | Proceed with targeted questions |
| Potential | 60-69% (120-139 pts) | Meets core requirements, notable gaps | Proceed if high-potential indicators |
| Borderline | 50-59% (100-119 pts) | Significant gaps, consider if exceptional fit | Proceed only with strong justification |
| Not Recommended | Below 50% (<100 pts) | Major gaps in critical areas | Do not proceed |

---

## CONFIDENCE FLAGGING PROTOCOL

**MANDATORY: Flag confidence level for each scored item**

### Confidence Levels

| Level | Code | Definition | Scoring Guidance |
|-------|------|------------|------------------|
| HIGH | H | Direct evidence in materials - specific quote available | Score as stated |
| MEDIUM | M | Inferred from context - logical conclusion from related evidence | Consider -1 point adjustment |
| LOW | L | Assumption or weak inference - no direct support | Consider -2 point adjustment, flag for interview |

### Confidence Summary

| Category | High Confidence Items | Medium Confidence Items | Low Confidence Items |
|----------|----------------------|------------------------|---------------------|
| 1. Skills Inventory | [Count] | [Count] | [Count] |
| 2. Experience Relevance | [Count] | [Count] | [Count] |
| 3. Demonstrated Impact | [Count] | [Count] | [Count] |
| 4. Credentials | [Count] | [Count] | [Count] |
| 5. Fit & Readiness | [Count] | [Count] | [Count] |

**Overall Confidence Assessment:** [ ] High / [ ] Medium / [ ] Low

**Low Confidence Items Requiring Interview Verification:**
1. [Item and category]
2. [Item and category]
3. [Continue as needed]

---

## INTER-RATER RELIABILITY PROTOCOL

**MANDATORY for hiring decisions involving multiple assessors**

### Calibration Checklist

- [ ] All raters reviewed the same source materials
- [ ] All raters used the same rubric version (2.2)
- [ ] Weight variant was agreed upon before scoring
- [ ] Critical barriers were reviewed jointly
- [ ] Proficiency scale definitions were calibrated
- [ ] Discrete anchor definitions were reviewed together

### Score Reconciliation

| Category | Rater 1 | Rater 2 | Variance | Reconciled Score |
|----------|---------|---------|----------|------------------|
| 1. Skills Inventory | | | | |
| 2. Experience Relevance | | | | |
| 3. Demonstrated Impact | | | | |
| 4. Credentials | | | | |
| 5. Fit & Readiness | | | | |
| **TOTAL** | | | | |

**Variance Threshold:** If any category varies by >10%, discussion required before finalizing.

**v2.2 Note:** With discrete anchors, variance should be minimal. If variance exceeds threshold, review specific anchors where disagreement occurred.

---

## BIAS REVIEW SECTION

**MANDATORY: Complete before finalizing assessment**

### Bias Checklist

| Potential Bias | Self-Check Question | Response |
|----------------|---------------------|----------|
| Halo Effect | Am I letting one strong area inflate other scores? | [ ] Yes - Adjusted / [ ] No |
| Horn Effect | Am I letting one weak area deflate other scores? | [ ] Yes - Adjusted / [ ] No |
| Similarity Bias | Am I favoring candidates similar to myself/team? | [ ] Yes - Adjusted / [ ] No |
| Recency Bias | Am I overweighting recent experience vs. career arc? | [ ] Yes - Adjusted / [ ] No |
| Prestige Bias | Am I overweighting employer/school names vs. actual skills? | [ ] Yes - Adjusted / [ ] No |
| Confirmation Bias | Am I seeking evidence to confirm initial impression? | [ ] Yes - Adjusted / [ ] No |

### Objectivity Verification

- [ ] Scores are based on documented evidence, not impressions
- [ ] Each criterion scored independently (not influenced by other scores)
- [ ] Negative and positive evidence both considered fairly
- [ ] Equivalent standards applied as to other candidates for this role
- [ ] Discrete anchors applied consistently throughout

---

## ASSESSOR NOTES

### Key Strengths (Top 3)
1. [Strength with supporting evidence]
2. [Strength with supporting evidence]
3. [Strength with supporting evidence]

### Key Concerns (Top 3)
1. [Concern with specific gap identified]
2. [Concern with specific gap identified]
3. [Concern with specific gap identified]

### Interview Focus Areas
1. [Topic to probe based on Medium/Low confidence items]
2. [Topic to probe based on identified concerns]
3. [Topic to verify based on impressive claims]

### Special Considerations
[Any unique context: relocation, compensation expectations, competing offers, etc.]

---

## VERSION HISTORY

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 2.2 | [Creation Date] | [Assessor] | Initial rubric creation from framework v2.2 |
| | | | |

---

## QUALITY CONTROL CHECKLIST

**MANDATORY: Complete before submitting assessment**

### Rubric Completeness

- [ ] All 5 categories scored
- [ ] All subcategories within each category scored
- [ ] Evidence citations provided for each score
- [ ] Confidence levels assigned to all items
- [ ] Critical barriers table completed

### Scoring Integrity

- [ ] No double-counting (each factor scored in ONE place only)
- [ ] Proficiency scale used for skills (not years)
- [ ] Weight variant correctly applied
- [ ] Normalized score calculated correctly
- [ ] Score interpretation band identified
- [ ] Evidence verification protocol applied (all scores have citations)
- [ ] Experience type classified (Direct/Adjacent/Transferable/Assumed)
- [ ] Domain specificity verified (no false equivalencies)
- [ ] **Discrete anchors applied - no range-based scoring**

### Documentation Quality

- [ ] Alignment statement completed
- [ ] Weight justification documented
- [ ] Top 3 strengths identified with evidence
- [ ] Top 3 concerns identified with specifics
- [ ] Interview focus areas specified
- [ ] Bias review completed

### Final Verification

- [ ] Assessor name and date recorded
- [ ] Version history updated
- [ ] Ready for inter-rater reconciliation (if applicable)

---

# TEMPLATE ENDS HERE

---

## Framework Usage Guidelines

### When Generating a Job-Specific Rubric

1. **Read job posting thoroughly** - Extract all requirements, preferences, and success criteria
2. **Select weight variant** - Use Selection Guide to choose Technical IC, People Manager, or Executive
3. **Complete Critical Barriers** - Identify any disqualifying factors before detailed scoring
4. **Complete Alignment Statement** - Document how rubric maps to job requirements
5. **Fill Skills Inventory** - Extract 6 required skills and all preferred skills from posting
6. **Customize all categories** - Replace [bracketed placeholders] with job-specific content
7. **Remove unused sections** - Delete any certification rows, skill rows, etc. that don't apply
8. **Review discrete anchors** - Customize anchor examples for industry context where helpful

### Scoring Best Practices

1. **Evidence-based scoring** - Every score needs a citation to source materials
2. **Proficiency over tenure** - Years don't equal skill; use behavioral anchors
3. **Confidence flagging** - Be honest about inference vs. direct evidence
4. **No double-counting** - Check that each achievement/skill appears in ONE category only
5. **Calibrate with team** - For important hires, use inter-rater protocol
6. **Domain specificity** - Never assume equivalency; classify as Direct/Adjacent/Transferable/Assumed
7. **Quote evidence** - For scores above Basic, quote specific text supporting the score
8. **Use discrete anchors** - Select the single anchor that best matches; do not interpolate between anchors

### Common Mistakes to Avoid

| Mistake | Correct Approach |
|---------|------------------|
| Scoring years of experience as a skill | Use proficiency scale (0-6) based on capability demonstrated |
| Counting same achievement in multiple categories | Place each achievement in its primary category only |
| Inflating scores for prestigious employers | Score based on actual work done, not employer name |
| Skipping confidence flagging | Every score needs H/M/L confidence assignment |
| Ignoring bias checklist | Complete bias review before finalizing |
| Using default weights for all roles | Select appropriate variant for role type |
| Assuming domain equivalency | Real estate investment ≠ asset management; classify experience type explicitly |
| Scoring without quoted evidence | Every score above Basic needs specific text citation from source materials |
| Treating "adjacent" as "direct" experience | Adjacent experience caps at 80% points; document the transfer gap |
| **Interpolating between anchors** | **Select the single best-fit anchor; if between two, choose the lower score** |
| **Using "about" or "approximately"** | **Find the exact anchor match; when uncertain, round down** |

---

## Appendix: Anchor Customization Examples

### Industry Match Anchors by Sector

**Technology Sector:**
| Score | Example Anchor |
|-------|---------------|
| 10 | Enterprise SaaS → Enterprise SaaS |
| 9 | Enterprise SaaS → Mid-market SaaS |
| 8 | Enterprise SaaS → Enterprise on-prem software |
| 7 | Enterprise SaaS → IT Services/Consulting |
| 6 | Enterprise SaaS → Hardware tech company |
| 5 | Enterprise SaaS → Digital agency |
| 4 | Enterprise SaaS → E-commerce tech |
| 3 | Enterprise SaaS → FinTech (different domain) |
| 2 | Enterprise SaaS → Manufacturing (tech-enabled) |
| 1 | Enterprise SaaS → Traditional services |
| 0 | Enterprise SaaS → Non-tech industry |

**Financial Services:**
| Score | Example Anchor |
|-------|---------------|
| 10 | Investment banking M&A → Investment banking M&A |
| 9 | Investment banking M&A → Investment banking ECM/DCM |
| 8 | Investment banking M&A → Corporate development |
| 7 | Investment banking M&A → Private equity |
| 6 | Investment banking M&A → Management consulting (M&A practice) |
| 5 | Investment banking M&A → Venture capital |
| 4 | Investment banking M&A → Corporate finance (F500) |
| 3 | Investment banking M&A → Commercial banking |
| 2 | Investment banking M&A → Insurance |
| 1 | Investment banking M&A → Fintech startup |
| 0 | Investment banking M&A → Unrelated industry |

**Real Estate:**
| Score | Example Anchor |
|-------|---------------|
| 10 | CRE Investment (office) → CRE Investment (office) |
| 9 | CRE Investment (office) → CRE Investment (industrial) |
| 8 | CRE Investment (office) → CRE Development |
| 7 | CRE Investment (office) → REIT asset management |
| 6 | CRE Investment (office) → CRE brokerage (investment sales) |
| 5 | CRE Investment (office) → Real estate PE/debt |
| 4 | CRE Investment (office) → Property management |
| 3 | CRE Investment (office) → Residential investment |
| 2 | CRE Investment (office) → Construction/development |
| 1 | CRE Investment (office) → Facilities management |
| 0 | CRE Investment (office) → Unrelated industry |

### Skills Proficiency Anchors by Domain

**Software Engineering:**
| Score | Level | Anchor |
|-------|-------|--------|
| 6 | Thought Leader | Authored framework used by 1000+ developers, conference keynote speaker, core contributor to major OSS |
| 5 | Expert | Architected systems handling 1M+ TPS, mentors senior engineers, designs systems adopted org-wide |
| 4 | Advanced | Designs complex distributed systems independently, minimal code review needed, identifies improvements |
| 3 | Proficient | Ships features independently, solid code quality, handles standard complexity |
| 2 | Developing | Contributes with guidance, handles routine tasks, growing technical judgment |
| 1 | Basic | Understands fundamentals, needs pairing/oversight, can assist but not lead |
| 0 | None | No software engineering evidence |

**Finance/Accounting:**
| Score | Level | Anchor |
|-------|-------|--------|
| 6 | Thought Leader | Published author on GAAP/IFRS interpretation, Big 4 partner-level, shapes accounting standards |
| 5 | Expert | Handles complex M&A accounting, trains CPAs, trusted for technical accounting decisions |
| 4 | Advanced | Manages audit relationships, identifies technical accounting issues, reviews others' work |
| 3 | Proficient | Prepares financial statements independently, handles standard transactions |
| 2 | Developing | Competent with standard transactions, needs review for complex items |
| 1 | Basic | Understands double-entry, requires supervision for most tasks |
| 0 | None | No accounting/finance evidence |

**Leadership/Management:**
| Score | Level | Anchor |
|-------|-------|--------|
| 6 | Thought Leader | Recognized leadership methodology creator, executive coach to C-suite, published on leadership |
| 5 | Expert | Turned around struggling teams, develops future leaders, handles complex people situations |
| 4 | Advanced | Manages managers effectively, strong team performance record, handles ambiguous situations |
| 3 | Proficient | Effective direct team leadership, good retention, delivers through team |
| 2 | Developing | Growing leadership skills, manages small team with support |
| 1 | Basic | Individual contributor with project coordination exposure |
| 0 | None | No leadership evidence |

### Quantified Achievement Anchors

**Financial Impact Examples:**
| Score | Range | Example Statements |
|-------|-------|-------------------|
| 10 | >$50M | "Delivered $75M cost reduction through supply chain transformation" |
| 9 | $25M-$50M | "Generated $35M incremental revenue from new product line" |
| 8 | $10M-$25M | "Reduced operating costs by $18M through process automation" |
| 7 | $5M-$10M | "Drove $7.5M in new business through strategic partnerships" |
| 6 | $2M-$5M | "Achieved $3.2M savings through vendor consolidation" |
| 5 | $1M-$2M | "Generated $1.5M efficiency gains through workflow optimization" |
| 4 | $500K-$1M | "Reduced costs by $750K annually through process improvements" |
| 3 | $250K-$500K | "Delivered $400K in savings through automation project" |
| 2 | $100K-$250K | "Achieved $180K cost reduction through inventory optimization" |
| 1 | <$100K | "Saved $75K through process improvement initiative" |
| 0 | None | No financial metrics documented |

**Efficiency Gain Examples:**
| Score | Range | Example Statements |
|-------|-------|-------------------|
| 10 | >75% | "Reduced deployment time from 4 hours to 15 minutes (94%)" |
| 9 | 60-75% | "Improved processing speed by 68% through system redesign" |
| 8 | 50-60% | "Reduced defect rate by 55% through quality program" |
| 7 | 40-50% | "Improved customer satisfaction scores by 45%" |
| 6 | 30-40% | "Increased team velocity by 35% over 6 months" |
| 5 | 20-30% | "Reduced cycle time by 25% through lean implementation" |
| 4 | 15-20% | "Improved first-call resolution by 18%" |
| 3 | 10-15% | "Increased on-time delivery by 12%" |
| 2 | 5-10% | "Reduced error rate by 8%" |
| 1 | <5% | "Improved response time by 4%" |
| 0 | None | No efficiency metrics documented |

---

## Appendix: Scoring Decision Rules

### When Evidence Falls Between Anchors

**Rule 1: Default to Lower Score**
If evidence could reasonably support two adjacent anchors, select the lower score. This ensures conservative, defensible assessments.

**Rule 2: Document the Decision**
When a borderline decision is made, note in the evidence field: "Borderline X/Y, scored Y due to [specific reason]"

**Rule 3: Flag for Interview**
Borderline scores should be flagged as MEDIUM confidence and added to interview verification list.

### When Evidence is Absent

**Rule 4: Zero Means Zero Evidence**
Score 0 means no evidence exists in any provided materials. Do not infer capability from job titles alone.

**Rule 5: Low Confidence Maximum**
If scoring based on inference rather than evidence, maximum score is 1 (Basic) with LOW confidence flag.

### When Evidence Conflicts

**Rule 6: Most Recent Takes Precedence**
When resume shows conflicting information, weight the most recent evidence more heavily.

**Rule 7: Specific Beats General**
Specific quantified claims take precedence over general statements.

**Rule 8: Document Conflicts**
Note any conflicting evidence in the comments: "Resume states X but also Y - flagged for clarification"
