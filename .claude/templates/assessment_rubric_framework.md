# Assessment Rubric Framework Template v2.0

**DO NOT USE THIS FILE DIRECTLY** - This is a meta-template for generating job-specific rubrics.

## Purpose

This template defines the canonical **200-point deduplicated scoring structure** with **role-based weight variants** for candidate assessments. Each assessment criterion appears in ONE location only to prevent double-counting. Skill scoring is PROFICIENCY-BASED, not years-based.

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
**Rubric Version:** 2.0

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

## CATEGORY 1: SKILLS INVENTORY

**Default: 50 points | Adjusted: [X] points per selected variant**

### 1A. Required Technical Skills (36 points base)

**MANDATORY: Score 6 required skills using PROFICIENCY scale (0-6 each)**

**PROFICIENCY SCALE - NOT YEARS-BASED**

| Level | Score | Definition | Behavioral Anchors |
|-------|-------|------------|-------------------|
| Thought Leader | 6 | Recognized authority who shapes industry practice | Publishes/speaks, influences standards, sought for expertise externally |
| Expert | 5 | Handles complex/novel scenarios, mentors others | Solves edge cases, trains others, trusted for critical decisions |
| Advanced | 4 | Complex work with minimal oversight | Independent on difficult tasks, identifies improvements |
| Proficient | 3 | Solid independent execution | Reliable delivery, handles standard complexity |
| Developing | 2 | Effective with guidance | Productive with support, growing capability |
| Basic | 1 | Foundational knowledge, needs support | Understands concepts, requires direction |
| None | 0 | No demonstrated capability | No evidence in materials |

#### Required Skills Assessment

| # | Skill | Job Requirement | Proficiency Score (0-6) | Evidence Citation | Confidence |
|---|-------|-----------------|------------------------|-------------------|------------|
| 1 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [H/M/L] |
| 2 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [H/M/L] |
| 3 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [H/M/L] |
| 4 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [H/M/L] |
| 5 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [H/M/L] |
| 6 | [Skill from posting] | [Specific requirement] | [ ] | [Source: Resume/Portfolio line#] | [H/M/L] |

**Required Skills Subtotal:** _____ / 36

### 1B. Preferred Skills (14 points base)

**MANDATORY: Score preferred/nice-to-have skills (variable count, allocate points proportionally)**

| # | Preferred Skill | Points Available | Score | Evidence Citation | Confidence |
|---|-----------------|------------------|-------|-------------------|------------|
| 1 | [Skill from posting] | [X pts] | [ ] | [Source] | [H/M/L] |
| 2 | [Skill from posting] | [X pts] | [ ] | [Source] | [H/M/L] |
| 3 | [Skill from posting] | [X pts] | [ ] | [Source] | [H/M/L] |
| [Continue as needed] | | | | | |

**Preferred Skills Subtotal:** _____ / 14

### Category 1 Total: _____ / 50 (or adjusted variant total)

**Normalized for variant:** _____ / [Variant Points]

---

## CATEGORY 2: EXPERIENCE RELEVANCE

**Default: 40 points | Adjusted: [X] points per selected variant**

### 2A. Industry/Domain Alignment (20 points)

**MANDATORY: Assess alignment with target industry and domain requirements**

| Factor | Weight | Scoring Criteria | Score | Evidence |
|--------|--------|------------------|-------|----------|
| Primary Industry Match | 10 pts | 10=Exact same industry, 7-9=Adjacent industry, 4-6=Transferable industry, 1-3=Unrelated with some overlap, 0=No relevance | [ ] | [Company/role from resume] |
| Domain Expertise | 6 pts | 6=Deep domain expert, 4-5=Solid domain knowledge, 2-3=Some domain exposure, 0-1=No domain experience | [ ] | [Specific domain evidence] |
| Market/Segment Knowledge | 4 pts | 4=Same market segment, 2-3=Related segment, 0-1=Different segment | [ ] | [Market experience cited] |

**Industry/Domain Subtotal:** _____ / 20

### 2B. Career Trajectory & Growth (12 points)

**MANDATORY: Assess progression pattern and growth trajectory**

| Factor | Weight | Scoring Criteria | Score | Evidence |
|--------|--------|------------------|-------|----------|
| Progression Pattern | 6 pts | 6=Clear upward trajectory with increasing scope, 4-5=Steady advancement, 2-3=Lateral moves with skill growth, 0-1=Stagnant or regression | [ ] | [Title progression] |
| Scope Expansion | 4 pts | 4=Consistently expanded responsibility, 2-3=Some scope growth, 0-1=Static scope | [ ] | [Budget/team/geo growth] |
| Role Stability | 2 pts | 2=Appropriate tenure (2-5 yrs avg), 1=Short tenure with explanation, 0=Job hopping pattern | [ ] | [Tenure analysis] |

**Career Trajectory Subtotal:** _____ / 12

### 2C. Recency of Relevant Work (8 points)

**MANDATORY: Weight recent experience more heavily**

| Recency Window | Weight | Scoring Criteria | Score | Evidence |
|----------------|--------|------------------|-------|----------|
| Current/Last Role (0-2 years) | 4 pts | 4=Highly relevant current work, 2-3=Moderately relevant, 0-1=Not relevant | [ ] | [Current role duties] |
| Recent History (2-5 years) | 3 pts | 3=Strong relevant history, 1-2=Some relevance, 0=No relevance | [ ] | [Recent roles] |
| Earlier Career (5+ years) | 1 pt | 1=Foundational relevance, 0=Not applicable | [ ] | [Earlier experience] |

**Recency Subtotal:** _____ / 8

### Category 2 Total: _____ / 40 (or adjusted variant total)

---

## CATEGORY 3: DEMONSTRATED IMPACT

**Default: 60 points | Adjusted: [X] points per selected variant**

**NOTE: This is the highest-weighted category (30% default). Evidence must be specific and quantified.**

### 3A. Quantified Achievements (30 points)

**MANDATORY: Score based on measurable results with verified metrics**

| Achievement Type | Max Points | Scoring Anchors | Score | Specific Achievement Cited | Confidence |
|------------------|------------|-----------------|-------|---------------------------|------------|
| Financial Impact | 10 pts | 10=>$10M impact, 7-9=$1M-$10M, 4-6=$100K-$1M, 1-3=<$100K documented, 0=No financial metrics | [ ] | [Exact claim from resume] | [H/M/L] |
| Efficiency/Quality Gains | 10 pts | 10=>50% improvement, 7-9=25-50%, 4-6=10-25%, 1-3=<10% documented, 0=No metrics | [ ] | [Exact claim from resume] | [H/M/L] |
| Growth Metrics | 10 pts | 10=Exceptional growth (>100%), 7-9=Strong growth (50-100%), 4-6=Solid growth (25-50%), 1-3=Modest growth (<25%), 0=No metrics | [ ] | [Exact claim from resume] | [H/M/L] |

**Quantified Achievements Subtotal:** _____ / 30

### 3B. Scale & Complexity Managed (18 points)

**MANDATORY: Assess scope indicators without duplicating Category 2**

| Scale Factor | Max Points | Scoring Anchors | Score | Evidence |
|--------------|------------|-----------------|-------|----------|
| Budget/P&L Responsibility | 6 pts | 6=>$100M, 4-5=$10M-$100M, 2-3=$1M-$10M, 1=<$1M, 0=No budget responsibility | [ ] | [Budget amount cited] |
| Team/Organization Size | 6 pts | 6=>100 direct/indirect, 4-5=25-100, 2-3=5-25, 1=<5, 0=Individual contributor only | [ ] | [Team size cited] |
| Geographic/Complexity Scope | 6 pts | 6=Global/multi-region, 4-5=National/multi-site, 2-3=Regional, 1=Single location, 0=Limited scope | [ ] | [Scope indicators] |

**Scale & Complexity Subtotal:** _____ / 18

### 3C. Innovation & Initiative (12 points)

**MANDATORY: Assess self-directed improvements and creative problem-solving**

| Factor | Max Points | Scoring Anchors | Score | Evidence |
|--------|------------|-----------------|-------|----------|
| Process/Product Innovation | 6 pts | 6=Created new offerings/methods adopted broadly, 4-5=Significant improvements with measurable impact, 2-3=Incremental improvements, 0-1=No innovation evidence | [ ] | [Innovation example] |
| Initiative & Ownership | 6 pts | 6=Identified and solved major problems proactively, 4-5=Led improvements beyond job scope, 2-3=Contributed to improvement efforts, 0-1=Executed assigned work only | [ ] | [Initiative example] |

**Innovation & Initiative Subtotal:** _____ / 12

### Category 3 Total: _____ / 60 (or adjusted variant total)

---

## CATEGORY 4: CREDENTIALS

**Default: 20 points | Adjusted: [X] points per selected variant**

### 4A. Education (12 points)

**MANDATORY: Score against stated requirements**

| Factor | Max Points | Scoring Criteria | Score | Evidence |
|--------|------------|------------------|-------|----------|
| Degree Level | 8 pts | 8=Exceeds requirement (e.g., PhD when Master's required), 6-7=Advanced degree at requirement level, 4-5=Meets exact requirement, 2-3=Below requirement with compensating factors, 0-1=Does not meet minimum | [ ] | [Degree, Institution] |
| Field Relevance | 4 pts | 4=Exact field match, 2-3=Related field, 1=Unrelated but compensated, 0=No relevance | [ ] | [Major/Concentration] |

**Education Subtotal:** _____ / 12

### 4B. Certifications (8 points)

**MANDATORY: Score against required and preferred certifications**

| Certification | Requirement Level | Status | Points | Evidence |
|---------------|------------------|--------|--------|----------|
| [Required Cert 1] | Required | [ ] Current / [ ] Expired / [ ] In Progress / [ ] None | /3 | [Cert details] |
| [Required Cert 2] | Required | [ ] Current / [ ] Expired / [ ] In Progress / [ ] None | /3 | [Cert details] |
| [Preferred Certs] | Preferred | [ ] Has / [ ] None | /2 | [Cert details] |

**Certification Scoring:**
- Required cert current: 3 pts each
- Required cert expired/in progress: 1-2 pts each
- Preferred certs: 0.5-1 pt each (max 2 pts total)

**Certifications Subtotal:** _____ / 8

### Category 4 Total: _____ / 20 (or adjusted variant total)

---

## CATEGORY 5: FIT & READINESS

**Default: 30 points | Adjusted: [X] points per selected variant**

### 5A. Communication Evidence (12 points)

**MANDATORY: Assess demonstrated communication capabilities**

| Factor | Max Points | Scoring Anchors | Score | Evidence |
|--------|------------|-----------------|-------|----------|
| Written Communication | 4 pts | 4=Executive-level writing samples/experience, 2-3=Professional writing demonstrated, 0-1=Basic or no evidence | [ ] | [Writing evidence] |
| Presentation/Speaking | 4 pts | 4=Conference speaker/board presentations, 2-3=Regular senior stakeholder presentations, 0-1=Limited or no evidence | [ ] | [Presentation evidence] |
| Cross-Functional Collaboration | 4 pts | 4=Led complex multi-stakeholder initiatives, 2-3=Regular cross-functional work, 0-1=Siloed experience | [ ] | [Collaboration evidence] |

**Communication Subtotal:** _____ / 12

### 5B. Cultural Alignment (10 points)

**MANDATORY: Assess alignment with stated company values and environment**

| Factor | Max Points | Scoring Anchors | Score | Evidence |
|--------|------------|-----------------|-------|----------|
| Values Alignment | 4 pts | 4=Clear evidence of matching values, 2-3=Some alignment indicators, 0-1=No evidence or misalignment | [ ] | [Values evidence] |
| Work Environment Fit | 3 pts | 3=Perfect match (remote/hybrid/onsite preference), 2=Acceptable fit, 0-1=Potential friction | [ ] | [Work style evidence] |
| Adaptability/Growth Mindset | 3 pts | 3=Strong evidence of adaptability and learning, 2=Some evidence, 0-1=Rigid or no evidence | [ ] | [Adaptability examples] |

**Cultural Alignment Subtotal:** _____ / 10

### 5C. Role-Specific Readiness (8 points)

**MANDATORY: Assess immediate readiness to perform in role**

| Factor | Max Points | Scoring Anchors | Score | Evidence |
|--------|------------|-----------------|-------|----------|
| Ramp-Up Time | 4 pts | 4=Immediately productive (done this exact work), 2-3=Quick ramp (similar work), 0-1=Extended ramp needed | [ ] | [Readiness indicators] |
| Role-Specific Requirements | 4 pts | 4=Meets all role-specific needs (travel, hours, etc.), 2-3=Meets most, 0-1=Potential conflicts | [ ] | [Specific req evidence] |

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
| HIGH | H | Direct evidence in materials | Score as stated |
| MEDIUM | M | Inferred from context/related evidence | Consider -1 point adjustment |
| LOW | L | Assumption or weak inference | Consider -2 point adjustment, flag for interview |

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
- [ ] All raters used the same rubric version
- [ ] Weight variant was agreed upon before scoring
- [ ] Critical barriers were reviewed jointly
- [ ] Proficiency scale definitions were calibrated

### Score Reconciliation

| Category | Rater 1 | Rater 2 | Variance | Reconciled Score |
|----------|---------|---------|----------|------------------|
| 1. Skills Inventory | | | | |
| 2. Experience Relevance | | | | |
| 3. Demonstrated Impact | | | | |
| 4. Credentials | | | | |
| 5. Fit & Readiness | | | | |
| **TOTAL** | | | | |

**Variance Threshold:** If any category varies by >20%, discussion required before finalizing.

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
| 2.0 | [Creation Date] | [Assessor] | Initial rubric creation from framework v2.0 |
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

### Scoring Best Practices

1. **Evidence-based scoring** - Every score needs a citation to source materials
2. **Proficiency over tenure** - Years don't equal skill; use behavioral anchors
3. **Confidence flagging** - Be honest about inference vs. direct evidence
4. **No double-counting** - Check that each achievement/skill appears in ONE category only
5. **Calibrate with team** - For important hires, use inter-rater protocol

### Common Mistakes to Avoid

| Mistake | Correct Approach |
|---------|------------------|
| Scoring years of experience as a skill | Use proficiency scale (0-6) based on capability demonstrated |
| Counting same achievement in multiple categories | Place each achievement in its primary category only |
| Inflating scores for prestigious employers | Score based on actual work done, not employer name |
| Skipping confidence flagging | Every score needs H/M/L confidence assignment |
| Ignoring bias checklist | Complete bias review before finalizing |
| Using default weights for all roles | Select appropriate variant for role type |

---

## Appendix: Example Anchor Patterns

### Skills Proficiency Anchors by Domain

**Software Engineering Example:**
- 6 (Thought Leader): Authored framework used by 1000+ developers, conference keynote speaker
- 5 (Expert): Architected systems handling 1M+ TPS, mentors senior engineers
- 4 (Advanced): Designs complex distributed systems, minimal code review needed
- 3 (Proficient): Ships features independently, solid code quality
- 2 (Developing): Contributes with guidance, growing technical judgment
- 1 (Basic): Understands fundamentals, needs pairing/oversight

**Finance/Accounting Example:**
- 6 (Thought Leader): Published author on GAAP interpretation, Big 4 partner-level expertise
- 5 (Expert): Handles complex M&A accounting, trains CPAs
- 4 (Advanced): Manages audit relationships, identifies technical accounting issues
- 3 (Proficient): Prepares financial statements independently
- 2 (Developing): Competent with standard transactions, needs review
- 1 (Basic): Understands double-entry, requires supervision

**Leadership/Management Example:**
- 6 (Thought Leader): Recognized for leadership methodology, executive coach to C-suite
- 5 (Expert): Turned around struggling teams, develops future leaders
- 4 (Advanced): Manages managers, handles complex people situations
- 3 (Proficient): Effective direct team leadership, good retention
- 2 (Developing): Growing leadership skills, inconsistent results
- 1 (Basic): Individual contributor with some project coordination

### Quantified Achievement Anchors

**Financial Impact:**
- >$10M: "Delivered $15M cost reduction through supply chain optimization"
- $1M-$10M: "Generated $3.2M incremental revenue from new product line"
- $100K-$1M: "Reduced operating costs by $450K annually"
- <$100K: "Saved $75K through process automation"

**Efficiency Gains:**
- >50%: "Reduced deployment time from 4 hours to 15 minutes (94%)"
- 25-50%: "Improved customer satisfaction scores by 35%"
- 10-25%: "Increased team velocity by 18% over 6 months"
- <10%: "Reduced defect rate by 8%"

**Scale Indicators:**
- Global: "Managed operations across 12 countries, 3 time zones"
- National: "Oversaw 15 locations across 8 states"
- Regional: "Led Western region with 5 branch offices"
- Local: "Managed single location with 25 staff"
