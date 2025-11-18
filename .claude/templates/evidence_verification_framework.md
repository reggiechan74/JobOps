# Evidence-Based Scoring Verification Framework

**MANDATORY APPLICATION** - All assessment commands must apply this framework when scoring candidates.

## Purpose

This framework defines the rigorous evidence requirements and verification protocols that must be applied when assessing candidates against rubrics. These rules ensure objectivity, defensibility, and consistency across all assessments.

## Commands That Reference This Framework

- `/createrubric` - Embeds this framework in generated rubrics
- `/assessjob` - Applies this framework during assessment scoring
- `/assesscandidate` - Applies this framework during assessment scoring

---

## CRITICAL: EVIDENCE-BASED SCORING VERIFICATION

Before finalizing any assessment score:

### 1. Evidence Citation Requirement

For every score above "Basic" (1 point), you MUST:
- Quote specific CV text supporting the score
- Identify exact CV section/line numbers
- Distinguish between direct experience vs transferable skills
- Verify claimed experience duration and recency

### 2. Domain Specificity Verification

NEVER assume domain equivalency without explicit evidence:
- Different industries require different expertise (healthcare ≠ finance ≠ retail ≠ manufacturing)
- Different functions require different skills (sales ≠ operations ≠ engineering ≠ marketing)
- Different company sizes require different capabilities (startup ≠ enterprise ≠ mid-market)
- Different asset/product types require specialized knowledge

### 3. Experience Type Classification

Before scoring, verify and classify experience as:
- **Direct**: Exact same role/industry/function as job requirement
- **Adjacent**: Related but different (specify differences)
- **Transferable**: Different but applicable skills (justify transfer logic)
- **Assumed**: No explicit evidence (score as "None" or "Basic" only)

### 4. Cross-Reference Protocol

Before scoring each section:
- Ask: "Where exactly in the CV does this evidence appear?"
- Verify: Does the evidence match the specific skill/experience requirement?
- Challenge: What evidence contradicts a high score?
- Validate: Are years of experience consistent with claimed expertise level?

### 5. Quantitative Verification

For any metrics-based scoring:
- Verify portfolio sizes, team sizes, budget amounts are explicitly stated
- Distinguish between managed vs contributed to vs exposed to
- Check timeframes and ensure experience is recent/relevant
- Validate geographic scope and market coverage claims

### 6. Quality Control Checklist

□ Every score ≥2 has specific CV citation with line numbers
□ Domain/industry experience explicitly verified (not inferred)
□ Years of experience match claimed specialization level
□ Leadership scope verified vs assumed
□ Technical skills backed by specific project examples
□ Achievement metrics directly quoted from CV
□ No conflation of different industries/functions/company types

### 7. Scoring Revision Protocol

If evidence doesn't support initial score:
- Immediately revise score downward to match actual evidence
- Document the correction reasoning in assessment notes
- Recalculate total score and recommendation level
- Update executive summary to reflect corrected analysis

### 8. Red Flag Verification

Automatically verify when scoring shows:
- High scores (≥3) across multiple categories for any candidate
- Perfect or near-perfect technical skills alignment
- Claims of expertise without supporting project details
- Leadership experience without team size/scope specifics
- Industry experience without explicit company/role evidence

## Application Guidelines

### When Creating Rubrics

Include this framework in the "Usage Guidelines" section with clear instructions that rubric users must apply these verification protocols.

### When Performing Assessments

Apply this framework systematically:
1. **Before scoring each category** - Review the verification requirements
2. **During scoring** - Document specific evidence for each score ≥2
3. **After initial scoring** - Run through the quality control checklist
4. **Before finalizing** - Execute red flag verification for high scores

### Documentation Requirements

For each score in your assessment, document:
- **Evidence Location**: Specific CV section and line numbers
- **Evidence Type**: Direct, Adjacent, Transferable, or Assumed
- **Verification Notes**: Why this evidence supports the assigned score
- **Gap Analysis**: What evidence is missing or weak

## Common Violations to Avoid

### ❌ Inference Without Evidence
- **Wrong**: "Candidate likely has project management experience given their senior role"
- **Right**: "CV line 45-47 documents leading 5-person project team for 18 months"

### ❌ Domain Conflation
- **Wrong**: "Real estate finance experience transfers perfectly to healthcare finance"
- **Right**: "Real estate finance is adjacent but different: RE focuses on asset valuation and cap rates, healthcare on reimbursement models and regulatory compliance. Score as Adjacent (2/3) not Direct."

### ❌ Assumed Scope
- **Wrong**: "VP title implies managing large teams"
- **Right**: "No team size mentioned in CV. Cannot verify leadership scope. Score as Basic (1/3) for leadership."

### ❌ Unbounded Metrics
- **Wrong**: "Managed budgets - Expert level"
- **Right**: "CV line 89: 'Managed $2.5M operating budget' - Good impact for mid-level role (5/10)"

### ❌ Transferable Skills Without Justification
- **Wrong**: "Marketing skills transfer to product management"
- **Right**: "Marketing analytics and customer research (CV lines 34-38) provide data analysis foundation for product role, but lacks product roadmap and engineering collaboration experience. Score as Transferable (2/3)."

## Enforcement

Assessment outputs that violate this framework should be rejected and regenerated with proper evidence documentation.
