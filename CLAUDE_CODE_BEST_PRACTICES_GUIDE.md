# Claude Code Best Practices Development Guide

**Version:** 1.0
**Generated:** 2025-11-08
**Source:** Complete analysis of JobOps repository (resumeoptimizer)
**Analysis Depth:** Full codebase exploration, git history review, architectural analysis

---

## Executive Summary

This guide distills lessons learned from building **JobOps**, a mature Claude Code application with 15 specialized agents, 15 slash commands, and complex multi-step workflows. After analyzing the complete codebase, git history evolution, and architectural decisions, this guide provides actionable patterns for building enterprise-grade Claude Code applications.

**Key Insights:**
- **Progressive disclosure** > Overwhelming users with complexity upfront
- **Parallel agent deployment** > Sequential processing (6x speed improvement)
- **YAML-driven automation** > Heuristic file parsing
- **Evidence-based rigor** > Unverified claims
- **User-agnostic patterns** > Hardcoded configurations
- **Modular reusability** > Monolithic commands
- **Validation loops** > Trust-and-forget

**Target Audience:** Developers building complex Claude Code applications requiring multi-agent orchestration, document generation, workflow automation, or intelligence-driven systems.

---

## Table of Contents

1. [Architectural Patterns](#1-architectural-patterns)
2. [Slash Command Design](#2-slash-command-design)
3. [Agent Architecture](#3-agent-architecture)
4. [Workflow Orchestration](#4-workflow-orchestration)
5. [Documentation Strategies](#5-documentation-strategies)
6. [File Organization & Naming](#6-file-organization--naming)
7. [Quality Assurance & Validation](#7-quality-assurance--validation)
8. [Version Control & Release Management](#8-version-control--release-management)
9. [Performance Optimization](#9-performance-optimization)
10. [Common Pitfalls & Anti-Patterns](#10-common-pitfalls--anti-patterns)
11. [Evolution & Refactoring Lessons](#11-evolution--refactoring-lessons)
12. [Tactical Implementation Checklist](#12-tactical-implementation-checklist)

---

## 1. ARCHITECTURAL PATTERNS

### 1.1 Separation of Concerns Pattern

**Pattern:** Commands orchestrate, agents execute, guides provide templates.

**Implementation:**
```
.claude/commands/       # Workflow orchestration
.claude/agents/         # Specialized task execution
Guides/                 # Templates and standards
CLAUDE.md               # System memory and instructions
```

**Rationale:**
- Commands focus on user interaction and workflow logic
- Agents focus on specialized processing and complex tasks
- Guides provide reusable templates and quality standards
- CLAUDE.md maintains consistent system behavior

**Example from JobOps:**
```markdown
# /buildresume command orchestrates:
Step 1: Launch step1-resume-draft agent
Step 2: Launch step2-provenance-check agent
Step 3: Launch step3-final-resume agent
```

**Anti-Pattern:**
❌ Mixing orchestration logic with execution logic in single files
❌ Commands that do both coordination and detailed processing

### 1.2 Modular Composability Pattern

**Pattern:** Build complex commands from simpler, reusable components.

**Implementation:**
```markdown
# Option 1: Composed modular workflow
/createrubric → Creates reusable rubric
/assesscandidate → Uses existing rubric for evaluation

# Option 2: All-in-one convenience
/assessjob → Creates rubric + performs assessment
```

**Benefits:**
- **Reusability:** Rubric created once, used for multiple candidates
- **Consistency:** Same evaluation criteria across all assessments
- **Flexibility:** Users choose modular or integrated approach
- **Audit Trail:** Separate rubric files preserve methodology

**Lesson from Git History:**
- Commit 5411a8b (Sept 27): "Add modular assessment system"
- Initially only had `/assessjob` (monolithic)
- Added `/createrubric` + `/assesscandidate` for reusability
- Kept both options for different use cases

### 1.3 Parallel Agent Deployment Pattern

**Pattern:** Deploy multiple independent agents simultaneously in a single message.

**Critical Implementation:**
```markdown
**CRITICAL**: You must deploy all agents in parallel using a single message
with multiple Task tool calls. This ensures simultaneous execution rather
than sequential processing.
```

**Code Example:**
```markdown
Deploy all 6 OSINT agents simultaneously to research [COMPANY]:
- Corporate intelligence (structure, finances, strategy)
- Legal intelligence (litigation, regulatory, compliance)
- Leadership intelligence (executives, personnel, governance)
- Compensation intelligence (salaries, benefits, equity)
- Culture intelligence (employee sentiment, workplace)
- Market intelligence (industry analysis, competitive positioning)
```

**Performance Impact:**
- **Sequential:** 6 agents × 2 minutes = 12 minutes total
- **Parallel:** max(2 minutes) = ~2 minutes total
- **Speedup:** 6x improvement

**Requirements:**
- Agents must be independent (no shared state dependencies)
- Each agent saves to separate file paths
- Standardized output format for synthesis
- Clear domain boundaries to prevent conflicts

### 1.4 Two-Phase Hybrid Workflow Pattern

**Pattern:** Combine fast agent-based reconnaissance with thorough main-session processing.

**Use Case:** Job search where API search is instant but Playwright scraping is slower.

**Architecture:**
```
Phase 1: Agent-Based API Search (hiringcafe-search agent)
   - Fast API reconnaissance (seconds)
   - Metadata extraction (titles, companies, URLs)
   - Filtering and prioritization
   ↓
Phase 2: Main Session Playwright Scraping
   - Navigate to each job URL
   - Extract complete verbatim descriptions
   - Save structured results
```

**Why This Works:**
- Agents cannot access MCP tools (like Playwright)
- API search provides quick results
- Playwright provides completeness
- Combined: speed + thoroughness

**Lesson Learned:**
- Initially tried all-in-agent approach → couldn't access Playwright
- Refactored to hybrid: agent for API, main session for MCP
- Result: Best of both worlds

### 1.5 YAML-Driven Automation Pattern

**Pattern:** Embed structured metadata in every generated document for downstream automation.

**Mandatory Front Matter Template:**
```yaml
---
job_file: Job_Postings/{{ARG1}}
role: <role title>
company: <company name>
candidate: <full candidate name>
generated_by: /commandname
generated_on: <ISO8601 timestamp>
output_type: resume_step1|assessment|rubric|briefing
status: draft|final|analysis
version: 1.0
---
```

**Placement:**
- **BEFORE** any markdown headings
- **FIRST** block in file
- Always use triple-dash delimiters

**Benefits:**
- Automated parsing without heuristics
- Provenance chain tracking
- Version control for regeneration
- Downstream tool integration
- Clear metadata for sorting/filtering

**Example Provenance Chain:**
```yaml
# Step 1 Draft
generated_by: /buildresume step1-resume-draft
output_type: resume_step1

# Step 2 Analysis (references Step 1)
draft_file: OutputResumes/Step1_Draft_Director_Alto_2025-09-23.md
generated_by: /buildresume step2-provenance-check
output_type: resume_provenance

# Step 3 Final (references both)
draft_file: OutputResumes/Step1_Draft_Director_Alto_2025-09-23.md
provenance_file: OutputResumes/Step2_Provenance_Analysis_Director_Alto_2025-09-23.md
generated_by: /buildresume step3-final-resume
output_type: resume_final
```

---

## 2. SLASH COMMAND DESIGN

### 2.1 Progressive Disclosure Principle

**Principle:** Reveal complexity gradually, not all at once.

**Implementation Structure:**
```markdown
## Step 1: Initial Setup
[Simple, clear instructions]

## Step 2: Core Processing
[Main workflow]

## Step 3: Advanced Options
[Complex configuration]
```

**Real Example from /buildresume:**
```markdown
## Step 1: Creating Initial Resume Draft
✓ Initiating strike package assembly

## Step 2: Provenance Analysis
✓ Executing credibility verification sweep

## Step 3: Final Hardened Resume
✓ Producing deployment-ready final resume
```

**Benefits:**
- Reduces cognitive load
- Clear step-by-step progression
- Professional tactical branding
- User knows where they are in the process

**Anti-Pattern:**
❌ Dumping all instructions in single block
❌ No visual progress indicators
❌ Unclear workflow sequence

### 2.2 Mandatory User Safety Mechanisms

**Safety Categories:**
1. **Confirmations** - Prevent accidental execution
2. **Warnings** - Highlight critical requirements
3. **Validations** - Verify output quality

**Example: Confirmation Pattern**
```markdown
Before any resume creation work, you MUST present this cultural profile menu to the user:

Please select the cultural profile that best matches your target market:
1. American Corporate / Tech Sector
2. Canadian Corporate
[...]

**WAIT for user selection before proceeding.**
```

**Example: Warning Pattern**
```markdown
🚨 MANDATORY DETAILED SCORING REQUIREMENT 🚨
The dynamic rubric you create MUST include granular point allocation
for every single criterion.

**IF ANY SECTION LACKS DETAILED BREAKDOWNS, THE RUBRIC IS INCOMPLETE**
```

**Example: Validation Pattern**
```markdown
**MANDATORY QUALITY CHECK**: Before saving, verify the rubric includes:
- ✅ Every required skill has Expert/Proficient/Basic/None scoring
- ✅ Every preferred skill has Strong/Basic/None scoring
- ✅ Point allocations sum to 100 total points
- ✅ Detailed scoring guidance provided for each criterion

**IF VALIDATION FAILS, REGENERATE THE RUBRIC**
```

**Lesson from v1.0.1 Bug Fix:**
- Issue: Credentials omitted without verification
- Fix: Added mandatory validation loops to Step 1 and Step 2
- Learning: **Trust but verify** → Add validation checkpoints

### 2.3 Parameter Passing Conventions

**Positional Arguments:**
```markdown
Arguments:
- {{ARG1}}: Job description file path (required)
- {{ARG2}}: Cultural profile (optional, defaults to "Canadian")
- {{ARG3}}: Additional flags (optional)
```

**Flag-Based Options:**
```bash
/searchjobs "software engineer" "Toronto" --company=Google --save --limit=50
```

**File Path Normalization:**
```markdown
Read the job posting from `Job_Postings/{{ARG1}}`
(add .md extension if not provided)

If job posting doesn't exist in Job_Postings/,
check root directory for legacy files
```

**Best Practices:**
- Use {{ARG1}}, {{ARG2}}, {{ARG3}} for positional
- Use --flag=value for optional parameters
- Auto-add file extensions when sensible
- Provide sensible defaults
- Document all parameters clearly

### 2.4 Error Handling Documentation

**Pattern:** Document anticipated errors and recovery strategies.

```markdown
### Error Handling
If issues are encountered:
- **Missing Rubric**: Report if rubric file cannot be found
- **Incomplete Rubric**: Note missing detailed scoring breakdowns
- **Misaligned Rubric**: Flag if rubric doesn't match job posting
- **Insufficient Candidate Data**: Document where evidence is limited

### Recovery Actions
For each error type, provide:
1. Clear error message to user
2. Suggested corrective action
3. Alternative workflow if applicable
```

**Graceful Degradation:**
```markdown
If comprehensive master resume not available:
1. Request user provide relevant experience details
2. Generate draft based on available information
3. Flag sections needing additional detail
4. Provide template for completing missing sections
```

### 2.5 Embedded Template Structures

**Pattern:** Include complete output templates within command definitions.

**Example:**
```markdown
Create assessment report in `OutputResumes/Assessment_[Company]_[Role]_[Date].md` with:

# Candidate Assessment Report: [Role] at [Company]
Date: [Assessment Date]
Candidate: [Candidate Name]

## Executive Summary
[2-3 sentences: Overall fit, key strengths, primary concerns, recommendation]

## Overall Assessment Score
**Total Score:** [X/100 points]

## Scoring Breakdown
### Technical Skills (25 points)
[Detailed breakdown with evidence citations]
...
```

**Benefits:**
- Agents know exactly what format to produce
- Consistency across all outputs
- Easy to validate structure
- Users know what to expect

---

## 3. AGENT ARCHITECTURE

### 3.1 Single Responsibility Principle

**Principle:** Each agent has ONE focused job.

**Example Specialization:**
```
osint-corporate     → Corporate intelligence ONLY
osint-legal         → Legal/litigation ONLY
osint-compensation  → Salary/benefits ONLY
step1-resume-draft  → Initial draft creation ONLY
step2-provenance    → Credibility analysis ONLY
```

**Anti-Pattern:**
❌ "kitchen-sink" agents that do everything
❌ Mixing unrelated responsibilities
❌ Agents that orchestrate other agents

**Benefits:**
- Clear boundaries
- Easy to debug
- Reusable in different contexts
- Parallel execution possible
- Focused expertise

### 3.2 Agent Definition Structure

**Minimal Agent File:**
```yaml
---
name: agent-name
description: Clear, concise capability statement (1-2 sentences)
model: opus|sonnet  # Choose based on task complexity
---

[Detailed agent instructions]
```

**Model Selection Guidelines:**
- **Opus:** Complex reasoning, creative tasks, strategic analysis
- **Sonnet:** Fast execution, straightforward tasks, cost optimization

**Examples from JobOps:**
- `candidate-assessment` → Opus (complex HR-level evaluation)
- `hiringcafe-search` → Sonnet (straightforward API calls)
- `osint-*` → Opus (deep intelligence analysis)

### 3.3 Standardized Output Requirements

**Every agent MUST:**
1. Save to designated folder with naming convention
2. Include YAML front matter with metadata
3. Follow structured template for consistency
4. Provide source attribution with citations
5. Include confidence indicators for findings

**Example Output Requirements:**
```markdown
**FILE STORAGE**: Save to `/Intelligence_Reports/[Company]_Corporate_Intelligence_[Date].md`

**YAML FRONT MATTER**:
---
company: {{ARG1}}
report_type: corporate
generated_by: /osint
generated_on: <ISO8601 timestamp>
output_type: intelligence_report
status: final
version: 1.0
---

**STRUCTURE**: Follow Corporate Intelligence Report template (section 5.2)

**CITATIONS**: All factual claims MUST include inline citations with URLs

**CONFIDENCE**: Label findings as HIGH/MEDIUM/LOW CONFIDENCE
```

### 3.4 Evidence-Based Scoring & Citations

**Critical Pattern:** Mandatory verification for all claims.

**Implementation:**
```markdown
## CRITICAL: EVIDENCE-BASED SCORING VERIFICATION

### 1. Evidence Citation Requirement
For every score above "Basic" (1 point), you MUST:
- Quote specific CV text supporting the score
- Identify exact CV section/line numbers
- Distinguish between direct experience vs transferable skills
- Verify claimed experience duration and recency

### 2. Domain Specificity Verification
NEVER assume domain equivalency without explicit evidence:
- Different industries require different expertise
- Different functions require different skills
- Different company sizes require different capabilities
```

**Quality Control Checklist:**
```markdown
□ Every score ≥2 has specific CV citation with line numbers
□ Domain/industry experience explicitly verified (not inferred)
□ Years of experience match claimed specialization level
□ Leadership scope verified vs assumed
□ Technical skills backed by specific project examples
□ Achievement metrics directly quoted from CV
```

**Lesson from Git History:**
- Commit 60ff4a7 (Sept 27): "Add evidence-based scoring verification"
- Initially allowed uncited claims
- Added mandatory citation requirements
- Result: Much higher credibility and defensibility

### 3.5 Tool Access Patterns

**Pattern:** Grant agents only the tools they need.

**Example: Assessment Agent**
```yaml
Tools Available:
- WebSearch: Domain knowledge acquisition
- WebFetch: Industry research
- Read: Load candidate work history files
- Write: Generate assessment report
```

**Example: Job Search Agent**
```yaml
Tools Available:
- Bash: API calls via curl/wget
- Write: Save search results
- NO MCP tools (handled in main session)
```

**Principle:** Least privilege - minimize attack surface and confusion.

---

## 4. WORKFLOW ORCHESTRATION

### 4.1 Sequential Multi-Step Workflows

**Pattern:** Each step reads previous outputs and builds upon them.

**Example: 3-Step Resume Building**
```
Step 1: Create Initial Draft
   - Input: Job description, Master resume
   - Output: Step1_Draft_*.md
   ↓
Step 2: Provenance Analysis
   - Input: Step 1 draft, Master resume
   - Output: Step2_Provenance_Analysis_*.md
   ↓
Step 3: Final Hardened Resume
   - Input: Step 1 draft, Step 2 analysis
   - Output: Step3_Final_Resume_*.md
```

**State Management:**
- YAML front matter tracks provenance chain
- File naming convention maintains linkage
- Each step explicitly references predecessors
- Version numbers increment on regeneration

**Command Orchestration:**
```markdown
## Step 1: Creating Initial Resume Draft
I'll start by reading the job description:
@$1

Now I'll run the step1-resume-draft agent:
**✓ Deploying Step 1 Agent - Initial Draft Creation**

## Step 2: Provenance Analysis
**✓ Deploying Step 2 Agent - Provenance Verification**

## Step 3: Final Hardened Resume
**✓ Deploying Step 3 Agent - Hardened Resume Production**
```

### 4.2 Parallel Multi-Agent Workflows

**Pattern:** Deploy multiple independent agents simultaneously.

**Critical Implementation Detail:**
```markdown
**CRITICAL**: You must deploy all agents in parallel using a SINGLE MESSAGE
with MULTIPLE Task tool calls. DO NOT send separate messages.

❌ WRONG:
  Send message 1: Deploy agent 1
  Send message 2: Deploy agent 2

✅ CORRECT:
  Send single message: Deploy agents 1, 2, 3, 4, 5, 6 simultaneously
```

**Example Deployment:**
```markdown
Deploy all 6 OSINT agents simultaneously:
1. Corporate intelligence
2. Legal intelligence
3. Leadership intelligence
4. Compensation intelligence
5. Culture intelligence
6. Market intelligence

Each agent operates independently and saves results to /Intelligence_Reports/
```

**Synthesis Phase:**
```markdown
After all 6 agents complete:
1. Read all 6 intelligence reports
2. Cross-reference findings
3. Resolve contradictions
4. Synthesize into Master Intelligence Report
5. Save to /Intelligence_Reports/[Company]_Master_Intelligence_[Date].md
```

### 4.3 Hybrid Agent + Main Session Workflows

**Pattern:** Use agents for what they do best, main session for MCP tools.

**Example: Job Search**
```
Phase 1: Agent-Based API Search
   - hiringcafe-search agent
   - Fast API calls using Bash
   - Returns: Job metadata, URLs

Phase 2: Main Session Playwright Scraping
   - mcp__playwright__browser_navigate
   - mcp__playwright__browser_snapshot
   - Extract complete job descriptions
   - Save structured results
```

**Why Hybrid:**
- Agents cannot access MCP tools
- API search is instant (agent)
- Playwright scraping is thorough (main session)
- Best of both worlds

### 4.4 File Input/Output Patterns

**Pattern 1: Glob-Based Discovery (User-Agnostic)**
```markdown
Load ALL candidate work history files from `ResumeSourceFolder/`:
- Use glob pattern to find all markdown files
- Read each file to build complete candidate profile
- Pattern-based discovery avoids hardcoded names
```

**Pattern 2: Standardized Output Naming**
```
Step1_Draft_[Role]_[Company]_YYYY-MM-DD.md
Step2_Provenance_Analysis_[Role]_[Company]_YYYY-MM-DD.md
Step3_Final_Resume_[Role]_[Company]_YYYY-MM-DD.md
Assessment_[Company]_[Role]_YYYY-MM-DD.md
Rubric_[Company]_[Role]_YYYYMMDD.md
```

**Pattern 3: File Path Normalization**
```markdown
1. Try Job_Postings/{{ARG1}}.md
2. Try Job_Postings/{{ARG1}}
3. Try root directory/{{ARG1}}.md (legacy support)
4. Report error if not found
```

**Lesson from Git History:**
- Commit 8eafc6b (Sept 25): "Fix hardcoded filenames"
- Commit 5d54824 (Oct 2): "Make agents generic and user-agnostic"
- **Learning:** Pattern-based discovery > hardcoded paths

---

## 5. DOCUMENTATION STRATEGIES

### 5.1 Multi-Tier Documentation Architecture

**Tier 1: User-Facing Entry Points**
```
README.md           - Primary entry, features, quick start
SETUP.md            - Installation and configuration
Why_I_Built_This.md - Philosophy and motivation
JOB_SEARCH_GUIDE.md - Specific workflow documentation
```

**Tier 2: System Documentation**
```
CLAUDE.md           - System instructions (Claude's memory)
AGENTS.md           - Agent development guidelines
CHANGELOG.md        - Version history and release notes
```

**Tier 3: Methodology Documentation**
```
SourceMaterial/CV_Master_Resume_Guide_v3.md          - HAM-Z methodology
SourceMaterial/CV_Provenance_Hardening_Pass_Check.md - Verification framework
SourceMaterial/System_Dynamics_Analysis_*.md         - Theoretical foundation
```

**Tier 4: Templates & Standards**
```
Guides/Experience_Template.md           - Structured templates
Guides/TechnologyMatrix_Template.md
Guides/Metadata_Standards.md            - YAML standards
Guides/Maintenance_Checklist.md
```

### 5.2 Inline Command Documentation Pattern

**Pattern:** Commands contain complete documentation within the file.

**Sections to Include:**
```markdown
---
description: Brief one-line summary
argument-hint: <arg1> [arg2]
---

**Arguments:**
[Detailed parameter documentation]

**Process:**
[Step-by-step workflow]

### YAML front matter requirement
[Complete metadata template]

### Usage Examples
[Multiple concrete examples]

### Error Handling
[Anticipated errors and recovery]

### Quality Standards
[Validation requirements]
```

**Benefits:**
- Self-documenting code
- Single source of truth
- No separate documentation to maintain
- Users can read command files for details

### 5.3 Cross-Reference System

**Pattern:** Create consistent cross-references across documentation.

**Examples:**
```markdown
# CLAUDE.md references:
- Command files (.claude/commands/)
- Agent files (.claude/agents/)
- Directory structure
- Workflow sequences

# README.md links to:
- CHANGELOG.md for version history
- SETUP.md for installation
- Why_I_Built_This.md for philosophy

# Commands reference:
- Agent files they delegate to
- Output directory locations
- Template structures
```

**Navigation Aids:**
```markdown
See also:
- [/assessjob command](.claude/commands/assessjob.md)
- [candidate-assessment agent](.claude/agents/candidate-assessment.md)
- [Assessment workflow](README.md#assessment-workflow)
```

### 5.4 Enforcement Language for Critical Requirements

**Pattern:** Use strong, unambiguous language for mandatory rules.

**Hierarchy:**
```markdown
🔥 CRITICAL ENFORCEMENT RULES 🔥
**ABSOLUTE REQUIREMENTS**
**MANDATORY**
**MUST** / **NEVER**
**Required** / **Forbidden**
```

**Example:**
```markdown
## 🔥 CRITICAL ENFORCEMENT RULES 🔥

**ABSOLUTE REQUIREMENTS FOR EVERY RUBRIC:**
1. **EVERY REQUIRED SKILL** must have Expert (3) / Proficient (2) / Basic (1) / None (0)
2. **EVERY PREFERRED SKILL** must have Strong (2) / Basic (1) / None (0)

**VIOLATION CONSEQUENCES:**
- Any rubric missing detailed breakdowns is INCOMPLETE
- You must regenerate the rubric if any section lacks granular scoring
```

**Why This Works:**
- No ambiguity about criticality
- Clear pass/fail criteria
- Explicit consequences
- Forces compliance

---

## 6. FILE ORGANIZATION & NAMING

### 6.1 Input/Output Separation

**Pattern:** Organize directories by lifecycle stage.

**Directory Structure:**
```
# INPUTS (stable, version controlled)
ResumeSourceFolder/    - Master career inventory (source of truth)
Job_Postings/          - Target job descriptions
Guides/                - Templates and standards
SourceMaterial/        - Methodology documentation

# OUTPUTS (generated, git-ignored)
OutputResumes/         - Generated drafts, analyses, finals
Briefing_Notes/        - Interview preparation materials
Scoring_Rubrics/       - Reusable assessment frameworks
Intelligence_Reports/  - OSINT company intelligence

# CONFIGURATION (.claude/)
agents/                - Agent definitions
commands/              - Slash command definitions
settings.local.json    - Claude Code configuration

# DOCUMENTATION (root)
README.md, CLAUDE.md, CHANGELOG.md, SETUP.md
```

**Benefits:**
- Clear separation of concerns
- Easy to .gitignore outputs
- Predictable file locations
- Scalable organization

### 6.2 Consistent Naming Conventions

**Commands:**
```
buildresume.md         (lowercase, no hyphens)
assessjob.md
createrubric.md
searchjobs.md
```

**Agents:**
```
step1-resume-draft.md           (descriptive, hyphenated)
osint-corporate.md
candidate-assessment.md
interview-briefing.md
```

**Outputs:**
```
Step1_Draft_[Role]_[Company]_YYYY-MM-DD.md
Step2_Provenance_Analysis_[Role]_[Company]_YYYY-MM-DD.md
Step3_Final_Resume_[Role]_[Company]_YYYY-MM-DD.md
Assessment_[Company]_[Role]_YYYY-MM-DD.md
Rubric_[Company]_[Role]_YYYYMMDD.md
Briefing_[Company]_[Role]_[Mode]_YYYY-MM-DD.md
```

**Job Postings:**
```
CompanyName_Role_YYYY-MM-DD.md
SearchResults_[Query]_YYYY-MM-DD.md
```

**Rationale:**
- Commands: Simple, URL-friendly
- Agents: Descriptive, self-documenting
- Outputs: Structured, sortable, parsable
- Dates: ISO-8601 format for sorting

### 6.3 .gitkeep Strategy

**Pattern:** Use .gitkeep to preserve empty directories in version control.

**Example:**
```
OutputResumes/.gitkeep
Briefing_Notes/.gitkeep
Intelligence_Reports/.gitkeep
Scoring_Rubrics/.gitkeep
```

**Why:**
- Git doesn't track empty directories
- Commands expect directories to exist
- .gitkeep preserves structure
- Users don't get file-not-found errors

---

## 7. QUALITY ASSURANCE & VALIDATION

### 7.1 Multi-Level Validation Pattern

**Level 1: Pre-Execution Validation**
```markdown
Before starting:
- ✅ Job description file exists
- ✅ Master resume folder is populated
- ✅ Required tools installed (pandoc if converting)
- ✅ User selections confirmed (cultural profile, etc.)
```

**Level 2: During-Execution Validation**
```markdown
During processing:
- ✅ All required fields populated in YAML front matter
- ✅ File naming convention followed
- ✅ Required sections present in output
- ✅ Word count within limits
```

**Level 3: Post-Execution Validation**
```markdown
After completion:
- ✅ All citations include sources with URLs
- ✅ All scores have supporting evidence
- ✅ No contradictions between sections
- ✅ Provenance chain intact (references correct source files)
```

**Lesson from v1.0.1:**
```markdown
# Bug: Credentials omitted without verification
# Fix: Added validation loops at Step 1 and Step 2

Step 1 agent now:
- Reviews ALL credentials from master resume
- Explicitly justifies any credential exclusions

Step 2 provenance check now:
- Verifies all credentials evaluated
- Flags any excluded credentials that map to job requirements
```

### 7.2 Quality Assurance Checklists

**Pattern:** Embed checklists in commands and agents.

**Example: Resume Quality Checklist**
```markdown
## Quality Assurance Checks
Before saving final resume:

□ All dates and positions match source documents
□ No contradictions between sections
□ Resume addresses top 5 requirements from job description
□ **Length check:** Maximum 3 pages when converted to Word
□ **Word count:** Maximum 1000 words (perform actual count)
□ **HAM-Z compliance:** All bullets follow Hard Skill + Action + Metrics formula
□ **Provenance:** All achievements traceable to master resume
□ **Cultural alignment:** Tone matches selected profile
□ **No superlatives:** Removed unbounded claims (market-leading, transformational)
```

**Example: Rubric Completeness Checklist**
```markdown
**MANDATORY QUALITY CHECK**: Before saving, verify the rubric includes:
- ✅ Every required skill has Expert/Proficient/Basic/None scoring (0-3 pts)
- ✅ Every preferred skill has Strong/Basic/None scoring (0-2 pts)
- ✅ Point allocations sum to exactly 100 total points
- ✅ Detailed scoring guidance provided for each criterion
- ✅ No ambiguous or subjective criteria without clear definitions

**IF ANY SECTION LACKS DETAILED BREAKDOWNS, THE RUBRIC IS INCOMPLETE**
```

### 7.3 Confidence Indicators & Source Quality Ratings

**Pattern:** Label all intelligence with confidence levels and source reliability.

**Confidence Levels:**
```markdown
[VERIFIED]     - Confirmed from primary sources with URLs
[INFERRED]     - Logical conclusions drawn from available data
[CALCULATED]   - Mathematical derivations from source data
[ASSUMPTION]   - Reasonable assumptions with stated reasoning

Confidence: HIGH / MEDIUM / LOW
Source Reliability: HIGH / MEDIUM / LOW
```

**Example Usage:**
```markdown
## Financial Health Assessment
- Revenue trend: $X billion¹ [VERIFIED - HIGH CONFIDENCE]
- Market position suggests competitive advantage² [INFERRED - MEDIUM CONFIDENCE]
- Projected growth: ~15% annually³ [CALCULATED - MEDIUM CONFIDENCE]

## Footnotes
¹ [Company Annual Report 2024], https://..., (HIGH RELIABILITY) - Accessed 2025-11-08
² Industry analysis from [Gartner], https://..., (MEDIUM RELIABILITY)
³ Calculated from 5-year CAGR based on source¹
```

### 7.4 Cross-Document Consistency Verification

**Pattern:** Verify consistency across related documents.

**Example: Provenance Chain Verification**
```markdown
## Provenance Verification Protocol
For each achievement in final resume:
1. Trace back to Step 1 draft
2. Verify Step 2 analysis addressed any concerns
3. Confirm source exists in master resume
4. Check metrics are unchanged (or justified if changed)
5. Ensure no new unbounded claims introduced
```

**Example: OSINT Cross-Reference**
```markdown
## Cross-Domain Intelligence Correlations
### Corroborating Evidence
- Corporate report shows revenue decline
- Compensation report shows salary freezes
- Culture report shows morale issues
- **CORRELATION**: Financial stress confirmed across 3 domains

### Contradictory Intelligence
- Legal report shows no major litigation
- Culture report mentions "class action rumors"
- **CONTRADICTION**: Requires additional investigation
```

---

## 8. VERSION CONTROL & RELEASE MANAGEMENT

### 8.1 Semantic Versioning Discipline

**Pattern:** Follow semantic versioning strictly.

**Version Format:** MAJOR.MINOR.PATCH

**When to Increment:**
```markdown
PATCH (1.0.X) - Increment for:
- Bug fixes and minor improvements
- Documentation updates
- Small enhancements to existing features
- Non-breaking changes

MINOR (1.X.0) - Increment for:
- New features and capabilities
- New slash commands or agents
- Backward-compatible functionality additions
- Significant enhancements

MAJOR (X.0.0) - Increment for:
- Breaking changes or incompatible API changes
- Major system rewrites or architecture changes
- Removal of deprecated features
- Changes requiring user intervention
```

### 8.2 Multi-File Version Synchronization

**Pattern:** Keep version numbers synchronized across files.

**Files to Update:**
```
1. package.json        - Primary version source
2. README.md           - Version badge
3. CHANGELOG.md        - Release notes
4. CLAUDE.md           - Current version reference
```

**Update Process:**
```bash
# 1. Update package.json
"version": "1.0.1"

# 2. Update README.md badge
![Version](https://img.shields.io/badge/version-1.0.1-blue)

# 3. Update CHANGELOG.md
## [1.0.1] - 2025-11-07

# 4. Update CLAUDE.md
**Current Version:** 1.0.1

# 5. Commit with version tag
git add package.json README.md CHANGELOG.md CLAUDE.md
git commit -m "Release version 1.0.1"
git tag -a v1.0.1 -m "Version 1.0.1"
git push && git push --tags
```

### 8.3 CHANGELOG Best Practices

**Pattern:** Follow "Keep a Changelog" format.

**Structure:**
```markdown
# Changelog

## [Unreleased]
### Added
- New features in development

### Changed
- Changes to existing functionality

### Deprecated
- Soon-to-be-removed features

### Removed
- Removed features

### Fixed
- Bug fixes

### Security
- Security improvements

## [1.0.1] - 2025-11-07
### Fixed
- Credential completeness verification in /buildresume
  - Step 1 agent reviews ALL credentials
  - Explicit justification for exclusions
  - Step 2 flags missing credentials
```

**Key Elements:**
- Keep unreleased section at top
- Use ISO dates (YYYY-MM-DD)
- Group changes by type
- Provide context for fixes
- Link to issues/PRs when applicable

### 8.4 Git Commit Message Standards

**Pattern:** Use imperative mood, clear context.

**Format:**
```
<type>: <subject>

<body>

<footer>
```

**Examples from JobOps:**
```
✅ GOOD:
Add comprehensive OSINT intelligence system with rigorous citation standards
Enhance /buildresume with credential completeness verification
Fix hardcoded filenames in candidate-assessment agent

❌ BAD:
Updated stuff
Fixed bug
Changes
```

**Types:**
- Add: New features
- Fix: Bug fixes
- Enhance: Improvements to existing features
- Update: Documentation or dependency updates
- Refactor: Code restructuring
- Remove: Feature removal

---

## 9. PERFORMANCE OPTIMIZATION

### 9.1 Parallel Execution Strategies

**Pattern:** Maximize parallel tool usage for independent operations.

**Example: Parallel File Reading**
```markdown
# ✅ GOOD: Parallel reads
Read these files simultaneously:
- Job_Postings/AltoJobPost.md
- ResumeSourceFolder/Experience/Director_Experience.md
- ResumeSourceFolder/Technology/TechnologyMatrix.md

# ❌ BAD: Sequential reads
Read Job_Postings/AltoJobPost.md
Then read ResumeSourceFolder/Experience/Director_Experience.md
Then read ResumeSourceFolder/Technology/TechnologyMatrix.md
```

**Example: Parallel Agent Deployment**
```markdown
# ✅ GOOD: Single message with 6 agent deployments
Deploy all 6 agents simultaneously in one message

# ❌ BAD: 6 separate messages
Deploy agent 1 (wait)
Deploy agent 2 (wait)
...
```

**Performance Impact:**
- Parallel file reads: 3x faster
- Parallel agent deployment: 6x faster
- Total workflow: 50-70% time reduction

### 9.2 Batch Processing for Large Operations

**Pattern:** Process items in manageable batches.

**Example: Job Search Scraping**
```markdown
**Batch Processing Strategy:**
For searches with many results:
- Process jobs in batches of 5-10 to manage browser resources
- For each batch:
  1. Navigate to each apply URL
  2. Capture page content
  3. Extract text
  4. Close browser tab
  5. Repeat for next job

**Performance Expectations:**
- 1-10 jobs: ~1-2 minutes
- 11-20 jobs: ~2-4 minutes
- 21-50 jobs: ~5-10 minutes (recommended max)
```

**Resource Management:**
```markdown
- Close browser tabs after each batch
- Provide progress indicators for large result sets
- Recommend smaller batches for better performance
- Allow user to stop/resume for very large searches
```

### 9.3 Model Selection for Cost Optimization

**Pattern:** Use appropriate model for task complexity.

**Guidelines:**
```markdown
Use Opus when:
- Complex reasoning required (assessment, analysis)
- Creative tasks (resume writing, cover letters)
- Strategic analysis (OSINT synthesis)
- Multi-domain integration

Use Sonnet when:
- Straightforward execution (API calls, file parsing)
- Fast processing needed
- Cost optimization important
- Simple transformations
```

**Example Distribution in JobOps:**
```
Opus Agents (7):
- candidate-assessment
- osint-agent
- osint-corporate
- osint-legal
- osint-leadership
- osint-compensation
- osint-culture
- osint-market
- interview-briefing
- interview-question-generator

Sonnet Agents (1):
- hiringcafe-search (simple API calls)
```

### 9.4 Caching and Reusability

**Pattern:** Design for reusability to avoid redundant work.

**Example: Reusable Rubrics**
```markdown
Instead of:
- Creating new rubric for every candidate assessment

Do this:
- Create rubric once with /createrubric
- Reuse for all candidates with /assesscandidate
- Result: 50% time savings for multiple candidates
```

**Example: Reusable Templates**
```markdown
Instead of:
- Recreating YAML templates in every command

Do this:
- Define templates in Guides/Metadata_Standards.md
- Reference standard templates
- Update once, propagate everywhere
```

---

## 10. COMMON PITFALLS & ANTI-PATTERNS

### 10.1 Hardcoded User-Specific Data

**❌ ANTI-PATTERN:**
```markdown
Read from: ResumeSourceFolder/ReggieChan_Experience.md
```

**✅ CORRECT PATTERN:**
```markdown
Use glob pattern to find all experience files:
ResumeSourceFolder/Experience/*.md
```

**Lesson from Git History:**
- Commit 8eafc6b: "Fix hardcoded filenames in candidate-assessment agent"
- Commit 5d54824: "Make resume agents generic and user-agnostic with pattern-based file discovery"

**Why It Matters:**
- Application unusable for other users
- Every user must modify agent files
- Breaks on first use for new users

### 10.2 Trust Without Verification

**❌ ANTI-PATTERN:**
```markdown
Create resume draft.
Done.
```

**✅ CORRECT PATTERN:**
```markdown
Create resume draft.
Verify all credentials included.
Check against master resume.
Validate quality standards.
Report findings to user.
```

**Lesson from v1.0.1 Bug:**
- Issue: Relevant credentials excluded without notice
- Fix: Added verification loops at Steps 1 & 2
- Learning: Always validate critical requirements

### 10.3 Uncited Claims

**❌ ANTI-PATTERN:**
```markdown
Company revenue is strong.
Leadership team is experienced.
Salaries are competitive.
```

**✅ CORRECT PATTERN:**
```markdown
Company revenue: $X billion in 2024¹ [VERIFIED - HIGH CONFIDENCE]
Leadership: CEO has 20 years in industry² [VERIFIED - HIGH CONFIDENCE]
Salaries: 15% above market average for role³ [CALCULATED - MEDIUM CONFIDENCE]

Footnotes:
¹ [Annual Report 2024], URL, (HIGH RELIABILITY)
² [LinkedIn Profile], URL, (MEDIUM RELIABILITY)
³ Calculated from Levels.fyi and Glassdoor data, (MEDIUM RELIABILITY)
```

**Lesson from Git History:**
- Commit 60ff4a7: "Add evidence-based scoring verification protocol"
- All claims now require citations
- Confidence indicators mandatory

### 10.4 Kitchen Sink Agents

**❌ ANTI-PATTERN:**
```markdown
Agent: do-everything-agent
- Create resume
- Perform assessment
- Generate cover letter
- Do OSINT research
- Search for jobs
```

**✅ CORRECT PATTERN:**
```markdown
Agent: step1-resume-draft (ONLY creates initial draft)
Agent: candidate-assessment (ONLY performs assessment)
Agent: step4-cover-letter (ONLY generates cover letter)
Agent: osint-corporate (ONLY corporate intelligence)
```

**Why Single Responsibility Matters:**
- Clear boundaries
- Easy debugging
- Reusable in different contexts
- Parallel execution possible
- Focused expertise

### 10.5 Sequential When Parallel Would Work

**❌ ANTI-PATTERN:**
```markdown
Deploy corporate intelligence agent.
Wait for completion.
Deploy legal intelligence agent.
Wait for completion.
Deploy leadership intelligence agent.
...
(Total: 12 minutes)
```

**✅ CORRECT PATTERN:**
```markdown
Deploy all 6 intelligence agents simultaneously in single message.
Wait for all to complete.
Synthesize results.
(Total: ~2 minutes)
```

**Performance Impact:** 6x speedup

### 10.6 Missing Error Handling

**❌ ANTI-PATTERN:**
```markdown
Read job description from Job_Postings/{{ARG1}}.md
Generate resume.
```

**✅ CORRECT PATTERN:**
```markdown
Try reading from:
1. Job_Postings/{{ARG1}}.md
2. Job_Postings/{{ARG1}} (auto-add extension)
3. Root directory/{{ARG1}}.md (legacy support)

If not found:
- Report clear error to user
- Suggest checking file path
- List available job postings
- Provide example of correct usage
```

### 10.7 Ambiguous Quality Standards

**❌ ANTI-PATTERN:**
```markdown
Create a good rubric.
Make sure it's detailed.
```

**✅ CORRECT PATTERN:**
```markdown
## 🔥 CRITICAL ENFORCEMENT RULES 🔥

**ABSOLUTE REQUIREMENTS FOR EVERY RUBRIC:**
1. EVERY required skill MUST have Expert (3) / Proficient (2) / Basic (1) / None (0)
2. EVERY preferred skill MUST have Strong (2) / Basic (1) / None (0)
3. Point allocations MUST sum to exactly 100 points
4. EVERY criterion MUST have detailed scoring guidance

**IF ANY SECTION LACKS DETAILED BREAKDOWNS, THE RUBRIC IS INCOMPLETE**
**YOU MUST REGENERATE IF VALIDATION FAILS**
```

### 10.8 Missing YAML Front Matter

**❌ ANTI-PATTERN:**
```markdown
# Resume for Director Position

## Summary
...
```

**✅ CORRECT PATTERN:**
```markdown
---
job_file: Job_Postings/Alto_Director.md
role: Director of Real Property
company: Alto Properties
candidate: John Doe
generated_by: /buildresume step1-resume-draft
generated_on: 2025-11-08T14:30:00Z
output_type: resume_step1
status: draft
version: 1.0
---

# Resume for Director Position

## Summary
...
```

**Why It Matters:**
- Enables automated parsing
- Tracks provenance chain
- Supports versioning
- Downstream tool integration

---

## 11. EVOLUTION & REFACTORING LESSONS

### 11.1 Development Timeline Insights

**Analysis of 7-Week Evolution (Sept 23 - Nov 8):**

```
Week 1 (Sept 23-24): Foundation
- Initial project setup
- Basic resume creation
- Provenance hardening system
✅ Learning: Start with core value proposition

Week 2 (Sept 24-25): Assessment System
- Candidate assessment with scoring rubrics
- Dynamic rubric generation
- Interview preparation system
✅ Learning: Assessment-first approach validated

Week 3 (Sept 26-27): Intelligence & Modularity
- OSINT distributed system
- Job comparison capability
- Modular assessment (create rubric + assess candidate)
- Evidence-based scoring verification
✅ Learning: Modularity enables reusability

Week 4-5 (Sept 28-Oct 2): Polish & Branding
- Job search functionality
- User-agnostic agents (removed hardcoded names)
- Tactical warfare branding
- Version tracking system
✅ Learning: Generalization comes after validation

Week 6-8 (Oct 3-Nov 8): Expansion
- GitHub portfolio management
- Resume parser (/create-career-history)
- Standardized templates (Guides/)
- Credential completeness verification (v1.0.1)
✅ Learning: Onboarding and validation loops critical
```

### 11.2 Major Refactorings

**Refactoring 1: Hardcoded → User-Agnostic (Oct 2)**
```
Before: ResumeSourceFolder/ReggieChan_Experience.md
After:  ResumeSourceFolder/Experience/*.md (glob pattern)

Impact: Application now usable by anyone
Lesson: Build generic from start, or plan early refactoring
```

**Refactoring 2: Monolithic → Modular Assessment (Sept 27)**
```
Before: /assessjob (all-in-one)
After:  /createrubric + /assesscandidate (modular)
        /assessjob (convenience wrapper)

Impact: Rubric reusability across multiple candidates
Lesson: Modularity enables new use cases
```

**Refactoring 3: Single-Agent → Distributed OSINT (Sept 26)**
```
Before: Single osint-agent doing everything sequentially
After:  6 specialized agents running in parallel

Impact: 6x performance improvement
Lesson: Parallelization worth the architectural complexity
```

**Refactoring 4: Trust → Verify (Nov 7, v1.0.1)**
```
Before: Draft creation without validation
After:  Mandatory credential verification loops in Steps 1 & 2

Impact: Prevents credential omission bugs
Lesson: Validation loops catch issues early
```

### 11.3 Feature Addition Patterns

**Pattern 1: Validate Before Generalizing**
```
1. Build hardcoded version for self
2. Validate it works and provides value
3. Generalize for all users
4. Add templates and documentation

Example: Resume system
- Built for specific user first
- Validated with real job applications
- Made user-agnostic
- Added templates and guides
```

**Pattern 2: Start Sequential, Optimize to Parallel**
```
1. Build sequential workflow (simpler)
2. Identify independent operations
3. Refactor to parallel execution
4. Document parallel deployment requirements

Example: OSINT system
- Initially sequential research
- Identified 6 independent domains
- Refactored to parallel agents
- Documented single-message requirement
```

**Pattern 3: Monolithic → Modular → Composite**
```
1. Build all-in-one command (fast to market)
2. Identify reusable components
3. Extract to separate commands
4. Keep composite as convenience option

Example: Assessment system
- /assessjob initially all-in-one
- Extracted /createrubric
- Extracted /assesscandidate
- Kept /assessjob as convenience wrapper
```

### 11.4 Lessons from Bug Fixes

**Bug: Missing Credentials (v1.0.1)**
```
Root Cause: No verification that all credentials evaluated
Impact: Relevant qualifications excluded from resume
Fix: Added validation loops at Step 1 and Step 2

Lesson: Critical requirements need verification checkpoints
```

**Bug: Hardcoded Filenames**
```
Root Cause: Agent files contained specific user names
Impact: Application unusable for other users
Fix: Pattern-based file discovery with glob

Lesson: Test with different users before considering complete
```

**Bug: Uncited Claims in Assessments**
```
Root Cause: No requirement for evidence citations
Impact: Scores lacked credibility and defensibility
Fix: Mandatory citation requirements with line numbers

Lesson: All claims need verifiable evidence
```

### 11.5 When to Refactor vs. Rebuild

**Refactor When:**
- Core architecture is sound
- Changes are localized
- Backward compatibility possible
- Risk is manageable

**Rebuild When:**
- Fundamental design flaw discovered
- Multiple interconnected issues
- Tech debt exceeds value
- New approach clearly superior

**Example from JobOps:**
- Hardcoded filenames → Refactored (localized change)
- Single OSINT agent → Refactored to 6 agents (architectural improvement)
- Assessment system → Refactored from monolithic to modular
- Never rebuilt from scratch (architecture was sound)

---

## 12. TACTICAL IMPLEMENTATION CHECKLIST

### 12.1 Starting a New Claude Code Project

**Phase 1: Foundation (Week 1)**
```markdown
□ Create .claude/agents/ directory
□ Create .claude/commands/ directory
□ Create input directories (source data)
□ Create output directories (generated artifacts)
□ Add .gitkeep to preserve empty directories
□ Create .gitignore for outputs
□ Write initial CLAUDE.md with system instructions
□ Set up package.json with dependencies
□ Create README.md with quick start
```

**Phase 2: Core Functionality (Week 2-3)**
```markdown
□ Build first slash command (single workflow)
□ Create first specialized agent
□ Validate with real use case
□ Add YAML front matter to all outputs
□ Implement file naming conventions
□ Add error handling for common failures
□ Document usage in README
```

**Phase 3: Expansion (Week 4-6)**
```markdown
□ Add additional commands/agents
□ Implement parallel agent deployment
□ Add validation checkpoints
□ Create reusable templates (Guides/)
□ Add quality assurance checklists
□ Implement evidence-based verification
□ Write SETUP.md for installation
```

**Phase 4: Polish (Week 7-8)**
```markdown
□ Make all agents user-agnostic (no hardcoded names)
□ Add comprehensive error handling
□ Create CHANGELOG.md
□ Set up version tracking
□ Add sample outputs
□ Write usage guides for complex workflows
□ Implement optional branding/theming
```

### 12.2 Command Development Checklist

**For Each New Slash Command:**
```markdown
□ Define clear single purpose
□ Document arguments and parameters
□ Include YAML front matter template
□ Add usage examples
□ Document error scenarios
□ Include validation requirements
□ Add quality assurance checklist
□ Specify agent delegation if applicable
□ Test with edge cases
□ Update README.md command list
□ Update CLAUDE.md if needed
```

### 12.3 Agent Development Checklist

**For Each New Agent:**
```markdown
□ Define single clear responsibility
□ Choose appropriate model (opus vs sonnet)
□ Document required tools
□ Specify output file location and naming
□ Include YAML front matter requirements
□ Define output structure template
□ Add evidence/citation requirements
□ Include confidence indicators
□ Specify quality validation criteria
□ Test independently
□ Document in AGENTS.md if needed
```

### 12.4 Pre-Commit Validation Checklist

**Before Every Commit:**
```markdown
□ All new files follow naming conventions
□ YAML front matter present in all outputs
□ No hardcoded user-specific data
□ Error handling added for new features
□ Documentation updated (README, CLAUDE.md)
□ Version numbers synchronized if releasing
□ .gitignore covers new output directories
□ Test with fresh user setup (user-agnostic check)
```

### 12.5 Release Checklist

**For Version Releases:**
```markdown
□ Update package.json version
□ Update README.md version badge
□ Update CLAUDE.md current version
□ Write CHANGELOG.md entry
  □ Added features
  □ Changed functionality
  □ Deprecated features
  □ Removed features
  □ Fixed bugs
  □ Security improvements
□ Test all slash commands
□ Verify agents still work
□ Check documentation links
□ Commit with version message
□ Create git tag (vX.Y.Z)
□ Push commits and tags
□ Create GitHub release (optional)
```

### 12.6 User Onboarding Checklist

**For New Users:**
```markdown
□ Clear installation instructions (SETUP.md)
□ Quick start guide in README
□ First-run command (/create-career-history or similar)
□ Sample outputs to show what's possible
□ Template files in Guides/
□ FAQ for common questions
□ Troubleshooting section
□ Example workflow walkthrough
□ Video or screenshots (optional)
```

### 12.7 Quality Gate Checklist

**Before Considering Feature Complete:**
```markdown
□ Progressive disclosure implemented
□ User safety mechanisms in place (confirmations, warnings)
□ Validation checkpoints added
□ Error handling comprehensive
□ Evidence-based verification required
□ YAML metadata automation working
□ Pattern-based file discovery (no hardcoded paths)
□ Parallel execution where applicable
□ Quality assurance checklists embedded
□ Documentation complete
□ Tested with different users
□ Sample outputs generated
```

---

## 13. ADVANCED PATTERNS

### 13.1 Tiered Information Architecture

**Pattern:** Prioritize information by importance and relevance.

**Example from Resume Building:**
```markdown
TIER 1: Must Include (Executive Roles)
- Crisis management and high-stakes decisions
- Strategic foresight and market prediction
- Industry firsts and innovation
- Thought leadership and publications

TIER 2: High Value (All Levels)
- Portfolio/financial performance
- Technology innovation
- Board/C-suite engagement

TIER 3: Supporting Evidence
- Operational excellence
- Team development
- Stakeholder management

TIER 4: Career Context
- Additional experience listings
- Foundational roles
```

**Role-Weighted Allocation:**
```markdown
PRIMARY ROLE (Most Relevant):     10-15 bullets
SECONDARY ROLES (Supporting):      3-5 bullets
TERTIARY ROLES (Career Context):   1-2 bullets
ADDITIONAL EXPERIENCE:             1 bullet per role
```

**Benefits:**
- Focuses on highest-value information
- Prevents information overload
- Guides users through prioritization
- Ensures critical content included

### 13.2 Cultural Adaptation System

**Pattern:** Parameterize style variations for different contexts.

**Example: Resume Cultural Profiles**
```markdown
1. American Corporate / Tech Sector
   - Confident, individualistic tone
   - Results-obsessed language
   - Quantified impact emphasized
   - Individual attribution

2. Canadian Corporate
   - Professional, modest tone
   - Balanced team/individual credit
   - Evidence-based claims
   - Conservative language

3. UK / Commonwealth
   - Formal, precise tone
   - Strong evidence backing
   - Understated language
   - Focus on qualifications

[... additional profiles ...]
```

**Implementation:**
```markdown
Before any resume creation, present cultural profile menu:
1. American Corporate / Tech Sector
2. Canadian Corporate
3. UK / Commonwealth
[...]

WAIT for user selection.
Apply selected profile throughout all content generation.
```

**Adaptation Covers:**
- Tone and voice
- Action verb selection
- Individual vs team attribution
- Quantification emphasis
- Format and length expectations

### 13.3 Provenance Chain Pattern

**Pattern:** Maintain complete traceability from source to final output.

**Implementation:**
```yaml
# Master Resume (Source of Truth)
file: ResumeSourceFolder/Experience/Director_Experience.md
content: "Led $2B portfolio redevelopment..."

# Step 1 Draft (First Transformation)
source_file: ResumeSourceFolder/Experience/Director_Experience.md
generated_by: /buildresume step1-resume-draft
content: "Led $2B portfolio redevelopment..." (copied verbatim)

# Step 2 Analysis (Verification)
draft_file: Step1_Draft_Director_Alto_2025-09-23.md
source_files: ResumeSourceFolder/**/*.md
generated_by: /buildresume step2-provenance-check
finding: "✓ Claim verified in Director_Experience.md line 47"

# Step 3 Final (Hardened)
draft_file: Step1_Draft_Director_Alto_2025-09-23.md
provenance_file: Step2_Provenance_Analysis_Director_Alto_2025-09-23.md
generated_by: /buildresume step3-final-resume
content: "Led $2B portfolio redevelopment..." (verified claim preserved)
```

**Benefits:**
- Complete audit trail
- Defensibility in interviews
- Credibility verification
- Quality assurance
- Version control

### 13.4 Dynamic Rubric Generation

**Pattern:** Create job-specific evaluation criteria rather than generic templates.

**Process:**
```markdown
1. Parse job description for requirements
2. Extract required skills (15 pts max)
3. Extract preferred skills (10 pts max)
4. Extract experience requirements (25 pts)
5. Extract key responsibilities (20 pts)
6. Extract expected achievements (15 pts)
7. Extract education/certifications (10 pts)
8. Extract cultural fit indicators (5 pts)
9. Generate granular scoring levels for each
10. Total: 100 points
```

**Granular Scoring Example:**
```markdown
Required Skill: Financial Modeling
- Expert (3 pts): 5+ years building complex DCF models, LBO models, merger models
- Proficient (2 pts): 3-5 years with standard financial models
- Basic (1 pt): 1-3 years, limited to simple models
- None (0 pts): No demonstrated experience
```

**Why Dynamic > Static:**
- Job-specific criteria
- Relevant scoring levels
- Accurate point allocation
- Defensible evaluation
- Consistent across candidates

### 13.5 Distributed Intelligence Synthesis

**Pattern:** Combine findings from multiple specialized sources into unified insight.

**Architecture:**
```markdown
Phase 1: Parallel Specialized Collection
├─ Corporate Agent    → Corporate_Intelligence.md
├─ Legal Agent        → Legal_Intelligence.md
├─ Leadership Agent   → Leadership_Intelligence.md
├─ Compensation Agent → Compensation_Intelligence.md
├─ Culture Agent      → Culture_Intelligence.md
└─ Market Agent       → Market_Intelligence.md

Phase 2: Cross-Domain Analysis
- Read all 6 intelligence reports
- Identify corroborating evidence
- Flag contradictory intelligence
- Assess confidence levels
- Weight by source reliability

Phase 3: Synthesis
- Integrate findings by theme
- Resolve contradictions
- Fill intelligence gaps
- Generate unified assessment
- Produce actionable recommendations
```

**Synthesis Techniques:**
```markdown
## Cross-Domain Verification
Financial decline (Corporate) + Salary freezes (Compensation) + Low morale (Culture)
= CORROBORATED: Company in financial stress

## Contradiction Resolution
No litigation (Legal) + "Class action rumors" (Culture)
= INVESTIGATE: Verify litigation status, may be early stage

## Confidence Weighting
High confidence + High reliability source = Strong finding
Medium confidence + Low reliability = Flag for verification
```

---

## 14. FUTURE-PROOFING STRATEGIES

### 14.1 Design for Evolution

**Principles:**
- Modular components > Monolithic systems
- Configuration > Hardcoding
- Templates > Duplication
- Patterns > Specific instances

**Example:**
```markdown
Instead of:
- step1-resume-draft.md (hardcoded Resume logic)

Do this:
- document-draft-generator.md (parameterized by document type)
- resume-template.md (Resume-specific config)
- coverletter-template.md (Cover letter config)
```

### 14.2 Backward Compatibility Considerations

**When Adding Features:**
```markdown
□ New parameters should be optional with sensible defaults
□ Existing file formats should still work
□ Legacy file locations checked as fallback
□ Deprecation warnings before removal
□ Migration guides for breaking changes
```

**Example:**
```markdown
# Support both new and old job posting locations
1. Try Job_Postings/{{ARG1}}.md (new location)
2. Try root/{{ARG1}}.md (legacy location)
3. Provide gentle nudge to use new location
```

### 14.3 Extensibility Points

**Design Extension Points:**
```markdown
1. Custom cultural profiles (add to menu)
2. Custom templates (add to Guides/)
3. Custom agents (add to .claude/agents/)
4. Custom commands (add to .claude/commands/)
5. Custom validation rules (configure in settings)
```

**Plugin Architecture (Future):**
```markdown
.claude/plugins/
  └─ industry-specific/
      ├─ tech-resume-variant.md
      ├─ healthcare-resume-variant.md
      └─ finance-resume-variant.md
```

---

## 15. MEASURING SUCCESS

### 15.1 Technical Metrics

**Code Quality:**
- Commands follow progressive disclosure pattern
- All agents are user-agnostic (no hardcoded names)
- YAML front matter present in all outputs
- Error handling comprehensive
- Documentation complete and current

**Architecture Quality:**
- Clear separation of concerns
- Single responsibility per agent
- Parallel execution where applicable
- Modular and composable design

**Performance:**
- Parallel operations reduce time by X%
- Batch processing prevents resource exhaustion
- Appropriate model selection (opus vs sonnet)

### 15.2 User Experience Metrics

**Onboarding:**
- Time to first successful command execution
- Number of errors in first session
- Clarity of error messages

**Usability:**
- Command discoverability
- Parameter obviousness
- Error recovery ease

**Value Delivery:**
- Time saved vs manual approach
- Quality of generated outputs
- User satisfaction scores

### 15.3 Quality Metrics

**Output Quality:**
- Provenance chain intact
- All citations present with URLs
- No contradictions in outputs
- Quality checklists passing

**Process Quality:**
- Validation checkpoints working
- Evidence requirements enforced
- Confidence indicators present
- Source quality rated

---

## 16. CONCLUSION

### Key Takeaways

**1. Architecture Patterns**
- Separation of concerns: Commands orchestrate, agents execute
- Modular composability: Build complex from simple reusable parts
- Parallel execution: Deploy independent agents simultaneously
- YAML-driven automation: Metadata enables downstream processing

**2. Design Principles**
- Progressive disclosure > Overwhelming complexity
- Evidence-based rigor > Unverified claims
- User-agnostic patterns > Hardcoded configurations
- Validation loops > Trust-and-forget
- Single responsibility > Kitchen sink agents

**3. Quality Assurance**
- Multi-level validation (pre, during, post)
- Quality assurance checklists
- Confidence indicators and source ratings
- Cross-document consistency verification
- Mandatory evidence citations

**4. Evolution Strategy**
- Validate before generalizing
- Start sequential, optimize to parallel
- Monolithic → Modular → Composite
- Refactor when localized, rebuild when fundamental
- Design for extension from start

**5. Performance Optimization**
- Parallel file operations
- Simultaneous agent deployment
- Batch processing for large operations
- Appropriate model selection
- Caching and reusability

### Application Beyond JobOps

These patterns are applicable to any Claude Code project requiring:
- Multi-agent orchestration
- Complex workflow automation
- Document generation pipelines
- Intelligence gathering systems
- Assessment and evaluation frameworks
- Evidence-based decision support

### Final Recommendations

**For New Projects:**
1. Start with clear architecture (commands/agents/guides)
2. Build user-agnostic from day one
3. Add YAML metadata to all outputs
4. Implement validation loops early
5. Document inline as you build
6. Test with different users before release

**For Existing Projects:**
1. Audit for hardcoded user data
2. Add evidence requirements
3. Implement validation checkpoints
4. Extract reusable components
5. Add quality assurance checklists
6. Document tribal knowledge

**For Scale:**
1. Identify parallelization opportunities
2. Implement modular architecture
3. Create reusable templates
4. Automate quality validation
5. Monitor performance metrics
6. Iterate based on user feedback

---

## Appendix A: Reference Implementation

**Repository:** github.com/reggiechan74/JobOps
**Version Analyzed:** 1.0.1
**Commit Range:** 41bbb50 (initial) to 962d467 (Nov 8, 2025)
**Total Commits:** 50+
**Development Duration:** 7 weeks (Sept 23 - Nov 8, 2025)

**Key Statistics:**
- 15 specialized agents
- 15 slash commands
- 8-step application workflow
- 6-agent parallel OSINT system
- 100-point dynamic scoring rubrics
- 3-step provenance hardening
- 6 cultural adaptation profiles

---

## Appendix B: Additional Resources

**Claude Code Documentation:**
- https://docs.claude.com/claude-code

**Related Patterns:**
- Model Context Protocol (MCP) integration
- Playwright automation patterns
- Progressive disclosure UI patterns
- Evidence-based decision making
- System dynamics analysis

**Further Reading:**
- Semantic Versioning: https://semver.org/
- Keep a Changelog: https://keepachangelog.com/
- Conventional Commits: https://www.conventionalcommits.org/

---

**Document Version:** 1.0
**Last Updated:** 2025-11-08
**Author:** Synthesized from JobOps repository analysis
**License:** MIT (same as source repository)

---

*This guide represents distilled wisdom from building a complex, production-grade Claude Code application. Use these patterns to accelerate your own development while avoiding common pitfalls.*
