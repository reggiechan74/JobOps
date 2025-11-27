---
description: Create or update GitHub portfolio documentation from repository analysis
argument-hint: <repository-url> [-create|-update]
---

I'll help you create or update your GitHub_Repositories.md portfolio documentation.

**Usage:**
- `/github-portfolio <repository-url>` (default: creates if none exists, updates if exists)
- `/github-portfolio <repository-url> -create`: Force create new portfolio (first repository)
- `/github-portfolio <repository-url> -update`: Add repository to existing portfolio
- `/github-portfolio -update`: Full portfolio refresh/reorganization (no URL needed)

**Arguments:**
- `$1`: Repository URL (required for adding repositories) OR `-update` flag for refresh-only mode
- `$2`: Mode flag (optional, defaults to auto-detect)
  - `-create`: Force create new portfolio file
  - `-update`: Force update existing portfolio

---

## Step 1: Parse Arguments and Determine Mode

**Repository URL**: $1
**Mode Flag**: ${2:-auto}

Checking arguments:
- If `$1` is `-update` and no `$2`: **Refresh Mode** (reorganize existing portfolio)
- If `$1` is a URL: **Repository Addition Mode** (create or update)
  - If portfolio exists: Update mode
  - If portfolio doesn't exist: Create mode
  - Mode can be forced with `$2` flag

---

## Step 2: Check Portfolio Existence

Checking if portfolio already exists:
@ResumeSourceFolder/Technology/GitHub_Repositories.md

**Result**: [Portfolio exists/does not exist]

---

## Step 3: Mode Selection

Based on arguments and portfolio existence:
- **Mode**: [CREATE NEW | ADD TO EXISTING | REFRESH EXISTING]

---

## Mode: CREATE NEW PORTFOLIO

**Purpose**: Create initial GitHub_Repositories.md with first repository.

**Repository URL Provided**: $1

### Step 3.1: Attempt to Access Repository

I'll try to access the repository directly to read its README and structure.

Attempting to fetch: $1

**Access Methods**:
1. Try WebFetch to read README.md from GitHub URL
2. If that fails (403, private repo, etc.), provide manual analysis script

### Step 3.2A: If Repository Accessible

[If WebFetch succeeds]:
- Read README.md content
- Parse repository name, description
- Attempt to infer technical details from README
- Ask user for additional details (commits, file counts, specific metrics)
- Proceed to create portfolio entry

### Step 3.2B: If Repository NOT Accessible (Most Likely)

[If WebFetch fails with 403/private/inaccessible]:

**The repository is not publicly accessible or I cannot reach it directly.**

No problem! Please run this analysis script in the repository's codespace/terminal and paste the complete results back to me:

```bash
#!/bin/bash
echo "=== REPOSITORY ANALYSIS FOR PORTFOLIO DOCUMENTATION ==="
echo ""
echo "## Basic Repository Information"
echo "Repository: $(basename $(git rev-parse --show-toplevel 2>/dev/null || pwd))"
echo "Current Branch: $(git branch --show-current 2>/dev/null || echo 'N/A')"
echo "Total Commits: $(git rev-list --count HEAD 2>/dev/null || echo 'N/A')"
echo "Last Commit: $(git log -1 --format='%h - %s (%ar)' 2>/dev/null || echo 'N/A')"
echo ""

echo "## File Structure"
echo "Total Files: $(find . -type f -not -path '*/\.*' | wc -l)"
echo ""
echo "Directory Tree (top 3 levels):"
tree -L 3 -I 'node_modules|.git|__pycache__|*.pyc' 2>/dev/null || find . -type d -not -path '*/\.*' -maxdepth 3 | sort
echo ""

echo "## README Content"
if [ -f README.md ]; then
    echo "--- START README.md ---"
    cat README.md
    echo "--- END README.md ---"
else
    echo "No README.md found"
fi
echo ""

echo "## Package/Dependency Information"
if [ -f package.json ]; then
    echo "--- package.json ---"
    cat package.json
fi
if [ -f requirements.txt ]; then
    echo "--- requirements.txt ---"
    cat requirements.txt
fi
if [ -f pyproject.toml ]; then
    echo "--- pyproject.toml ---"
    cat pyproject.toml
fi
echo ""

echo "## Code Statistics"
echo "Python files: $(find . -name '*.py' -not -path '*/\.*' | wc -l)"
echo "JavaScript files: $(find . -name '*.js' -not -path '*/\.*' | wc -l)"
echo "TypeScript files: $(find . -name '*.ts' -not -path '*/\.*' | wc -l)"
echo "Bash scripts: $(find . -name '*.sh' -not -path '*/\.*' | wc -l)"
echo "Markdown files: $(find . -name '*.md' -not -path '*/\.*' | wc -l)"
echo ""

echo "## Sample Code (largest Python/Bash/TypeScript files)"
for file in $(find . \( -name '*.py' -o -name '*.sh' -o -name '*.ts' \) | grep -v '.git' | grep -v 'node_modules' | head -5); do
    echo "--- $file ($(wc -l < "$file") lines) ---"
    head -50 "$file"
    echo ""
done

echo "## Git Activity"
echo "Commit History (last 10):"
git log --oneline -10 2>/dev/null || echo "Git history not available"
echo ""

echo "Recent Contributors:"
git shortlog -sn --no-merges | head -5 2>/dev/null || echo "N/A"
echo ""

echo "=== END ANALYSIS ==="
```

**Instructions**:
1. Open a terminal in the repository's codespace or local clone
2. Copy the entire script above
3. Paste it into the terminal and press Enter
4. Copy ALL the output (from "=== REPOSITORY ANALYSIS" to "=== END ANALYSIS ===")
5. Paste the complete output in your next message

Once you provide the analysis, I'll create your initial portfolio entry.

---

## Mode: ADD TO EXISTING PORTFOLIO

**Purpose**: Add new repository to existing portfolio.

**Repository URL Provided**: $1

### Step 3.1: Read Existing Portfolio

@ResumeSourceFolder/Technology/GitHub_Repositories.md

**Current Portfolio Status**:
- Total repositories: [count from file]
- Next repository number: [N+1]
- Portfolio version: [current version]
- Last updated: [date from file]

### Step 3.2: Attempt to Access Repository

Attempting to fetch: $1

**Access Methods**:
1. Try WebFetch to read README.md from GitHub URL
2. If that fails, provide manual analysis script

### Step 3.3A: If Repository Accessible

[If WebFetch succeeds]:
- Parse repository information
- Ask for additional metrics if needed
- Create comprehensive repository entry
- Update portfolio

### Step 3.3B: If Repository NOT Accessible

**The repository is not publicly accessible.**

Please run this analysis script in the repository's codespace and paste the results:

```bash
#!/bin/bash
echo "=== REPOSITORY ANALYSIS FOR PORTFOLIO DOCUMENTATION ==="
echo ""
echo "## Basic Repository Information"
echo "Repository: $(basename $(git rev-parse --show-toplevel 2>/dev/null || pwd))"
echo "Current Branch: $(git branch --show-current 2>/dev/null || echo 'N/A')"
echo "Total Commits: $(git rev-list --count HEAD 2>/dev/null || echo 'N/A')"
echo "Last Commit: $(git log -1 --format='%h - %s (%ar)' 2>/dev/null || echo 'N/A')"
echo ""

echo "## File Structure"
echo "Total Files: $(find . -type f -not -path '*/\.*' | wc -l)"
echo ""
echo "Directory Tree (top 3 levels):"
tree -L 3 -I 'node_modules|.git|__pycache__|*.pyc' 2>/dev/null || find . -type d -not -path '*/\.*' -maxdepth 3 | sort
echo ""

echo "## README Content"
if [ -f README.md ]; then
    echo "--- START README.md ---"
    cat README.md
    echo "--- END README.md ---"
else
    echo "No README.md found"
fi
echo ""

echo "## Package/Dependency Information"
if [ -f package.json ]; then
    echo "--- package.json ---"
    cat package.json
fi
if [ -f requirements.txt ]; then
    echo "--- requirements.txt ---"
    cat requirements.txt
fi
if [ -f pyproject.toml ]; then
    echo "--- pyproject.toml ---"
    cat pyproject.toml
fi
echo ""

echo "## Code Statistics"
echo "Python files: $(find . -name '*.py' -not -path '*/\.*' | wc -l)"
echo "JavaScript files: $(find . -name '*.js' -not -path '*/\.*' | wc -l)"
echo "TypeScript files: $(find . -name '*.ts' -not -path '*/\.*' | wc -l)"
echo "Bash scripts: $(find . -name '*.sh' -not -path '*/\.*' | wc -l)"
echo "Markdown files: $(find . -name '*.md' -not -path '*/\.*' | wc -l)"
echo ""

echo "## Sample Code (largest Python/Bash/TypeScript files)"
for file in $(find . \( -name '*.py' -o -name '*.sh' -o -name '*.ts' \) | grep -v '.git' | grep -v 'node_modules' | head -5); do
    echo "--- $file ($(wc -l < "$file") lines) ---"
    head -50 "$file"
    echo ""
done

echo "## Git Activity"
echo "Commit History (last 10):"
git log --oneline -10 2>/dev/null || echo "Git history not available"
echo ""

echo "Recent Contributors:"
git shortlog -sn --no-merges | head -5 2>/dev/null || echo "N/A"
echo ""

echo "=== END ANALYSIS ==="
```

**After receiving analysis**:
1. Parse repository analysis output
2. Create comprehensive repository entry following established format
3. Insert into Primary Repositories section (next available number)
4. Update capability matrices with new skills/technologies
5. Update conclusion to reference new repository count
6. Increment portfolio version number
7. Update "Last Updated" date

---

## Mode: REFRESH EXISTING PORTFOLIO

**Purpose**: Full portfolio reorganization and content refresh without adding new repositories.

**Triggered When**: User runs `/github-portfolio -update` with no URL

### Step 3.1: Read Existing Portfolio

@ResumeSourceFolder/Technology/GitHub_Repositories.md

**Current Portfolio Status**:
- Total repositories: [count]
- Portfolio version: [current version]
- Last updated: [date]

### Step 3.2: Comprehensive Review Tasks

1. **Repository Entries Consistency**:
   - Verify all entries follow current template format
   - Check for complete sections in each entry
   - Ensure consistent emoji usage and status indicators
   - Verify all repository links work

2. **Capability Matrices Accuracy**:
   - Review all 11 matrices for current proficiency levels
   - Add any new skills/technologies from recent work
   - Ensure honest positioning (AI-collaborative vs independent)
   - Check evidence references are current

3. **Positioning and Messaging**:
   - Refresh conclusion with current repository count
   - Update "What I Bring" / "What I Don't Claim" sections
   - Review application language examples for relevance
   - Ensure "AI-Collaborative Builder" philosophy is consistent

4. **Table of Contents**:
   - Verify all anchor links work correctly
   - Ensure all repositories are listed
   - Check capability matrix links
   - Confirm career positioning sections are linked

5. **Version and Date**:
   - Increment portfolio version (minor: X.Y+1)
   - Update lastUpdated to today's date in YAML frontmatter
   - Update "Last Updated" at bottom of document

6. **Quality Checks**:
   - Remove any outdated information
   - Fix broken internal references
   - Ensure development philosophy remains honest and accurate
   - Verify all repository counts match

### Step 3.3: Generate Update Summary

After refresh, provide summary of:
- Changes made
- New version number
- Items requiring user attention (if any)

---

## Repository Entry Template

When adding a new repository, use this comprehensive structure:

```markdown
### N. [repository-name](https://github.com/username/repo-name) 🎯
**Status**: Production (vX.Y.Z) | **License**: [License] | **Commits**: N

One-paragraph description highlighting core purpose and key technologies.

**Technical Architecture**:
- **Key Component 1**: Description with metrics
- **Key Component 2**: Description with metrics
- **Key Component 3**: Description with metrics
- **Key Component 4**: Description with metrics
- **Key Component 5**: Description with metrics

**Technical Sophistication**:
- **Scale Metric**: Total files, lines of code, complexity indicators
- **Advanced Pattern 1**: Specific implementation details
- **Advanced Pattern 2**: Specific implementation details
- **Advanced Pattern 3**: Specific implementation details
- **Advanced Pattern 4**: Specific implementation details

**Skills Demonstrated**:
- **Skill 1** (proficiency level: specific evidence)
- **Skill 2** (proficiency level: specific evidence)
- **Skill 3** (proficiency level: specific evidence)
- **Skill 4** (proficiency level: specific evidence)
- **Skill 5** (proficiency level: specific evidence)
- **Skill 6** (proficiency level: specific evidence)

**Business Impact**:
- Quantified improvement 1
- Quantified improvement 2
- Quantified improvement 3
- Use case or application context

**Development Approach**:
How this was built, key learning journey, relationship to other repositories, development philosophy demonstrated.

**Domain Specialization** (if applicable):
Specific industry domain, technical specialty, or problem space this addresses.

**Key Innovation**:
**Bold Innovation Name**: One-paragraph description of the unique approach, pattern, or solution this repository demonstrates. What makes this different or notable.

**[Optional] Component Breakdown**:
- **Category 1**: List of components/features
- **Category 2**: List of components/features
```

---

## Capability Matrices to Maintain

When updating portfolio, ensure these capability matrices remain comprehensive and accurate:

1. **Core Programming & Software Engineering** (7 capabilities)
2. **Data Science & Machine Learning** (5 capabilities)
3. **Financial Domain & Real Estate** (4 capabilities)
4. **AI Collaboration & Modern Development** (6 capabilities)
5. **AI Agent Systems & Orchestration** (6 capabilities)
6. **API & Data Integration** (5 capabilities)
7. **Bash Scripting & System Automation** (6 capabilities)
8. **Research Automation & Knowledge Management** (7 capabilities)
9. **Performance Engineering & Optimization** (5 capabilities)
10. **Workflow Automation & Process Design** (6 capabilities)
11. **Strategic & Analytical Capabilities** (5 capabilities)

**Matrix Format**:
```markdown
| Capability | Proficiency | Evidence | Development Method |
|---|---|---|---|
| **Capability Name** | Proficiency description | Specific evidence from repositories | How acquired/developed |
```

---

## Development Philosophy Principles

**Always maintain these core positioning principles**:

1. **Radical Honesty**: "Zero software engineering experience, ALL built through Claude Code collaboration"
2. **AI-Collaborative Builder**: Domain expert who learned to leverage AI as development partner
3. **What I Bring**: Domain expertise, architectural thinking, AI collaboration skills, proof of capability
4. **What I Don't Claim**: Traditional SWE credentials, ability to code independently, ML engineering qualification
5. **Value Proposition**: Living case study of domain expert → AI builder transformation

---

## Output Requirements

**File Location**: `ResumeSourceFolder/Technology/GitHub_Repositories.md`

**YAML Front Matter**:
```yaml
---
title: GitHub Repositories & Derived Technical Capabilities
person: Reggie Chan
lastUpdated: YYYY-MM-DD (today's date)
category: Technology Portfolio
keywords: AI-assisted-development, Python, machine-learning, financial-systems, Claude-Code
shareable: true
profileType: technical-portfolio
---
```

**Version Tracking**:
- Update `lastUpdated` field to today's date
- Increment `Portfolio Version` at end of document (format: X.Y)
  - Major version (X): Adding new repository or major restructuring
  - Minor version (Y): Content updates, refinements, corrections

**Quality Checks**:
- [ ] All repository entries follow consistent format
- [ ] Table of Contents links work correctly
- [ ] Capability matrices are complete and accurate
- [ ] Conclusion reflects current repository count
- [ ] Application language examples are current
- [ ] Repository Links section is complete
- [ ] Version and date are updated

---

Let me proceed with the selected mode.
