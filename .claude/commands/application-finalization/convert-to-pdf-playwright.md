---
description: Convert markdown resume to professionally designed PDF using Playwright
argument-hint: <markdown-resume-file> [theme] [pages]
---

I'll convert your markdown resume into a professionally designed PDF using the frontend-design skill and Playwright rendering.

**Arguments:**
- `$1`: Path to markdown resume file (required)
- `$2`: Theme choice (optional): `modern` | `classic` | `minimal` (default: modern)
- `$3`: Target page count (optional): `1` | `2` | `3` (default: auto-detect based on content)

**Available Themes:**
- **modern** - Clean, contemporary design with blue accent colors and modern typography
- **classic** - Traditional serif fonts, centered header, formal styling for conservative industries
- **minimal** - Ultra-clean, whitespace-focused, lightweight typography

**Page Count Options:**
- **1** - Compact single-page resume (entry-level, early career, or targeted applications)
- **2** - Standard two-page resume (mid-career professionals, 5-15 years experience)
- **3** - Extended three-page resume (senior executives, academics, or extensive portfolios)
- **auto** - Automatically determine optimal page count based on content length (default)

## Step 1: Validate Input File

First, I'll verify the markdown resume file exists and read its contents:
- File path: `$1`
- Selected theme: `$2` (defaults to `modern` if not specified)
- Target pages: `$3` (defaults to `auto` if not specified)

Read the resume file and validate it contains resume-like content (name, experience, skills, etc.).

**Parse page count argument:**
- If `$3` is `1`, `2`, or `3`: Use as target page count
- If `$3` is `auto` or not specified: Auto-detect based on content length
- If `$3` is invalid: Default to `auto` with a warning

## Step 2: Parse Resume Structure

Analyze the markdown content to identify resume sections:
- **Header**: Name, title, contact information
- **Summary/Profile**: Professional summary or objective
- **Experience**: Work history with company, title, dates, achievements
- **Skills**: Technical and soft skills
- **Education**: Degrees, institutions, dates
- **Certifications**: Professional certifications
- **Projects**: Notable projects (if present)
- **Additional sections**: Publications, awards, volunteer work, etc.

## Step 3: Load CSS Theme Files

Load the base CSS and selected theme:
1. Read `.claude/styles/resume-base.css` (core print/PDF styles)
2. Read `.claude/styles/resume-themes/{theme}.css` (theme-specific styles)

Combine into a single stylesheet for the HTML document.

## Step 4: Generate Professional HTML

Using the **frontend-design skill**, generate a professionally structured HTML resume that:

1. **Uses semantic HTML5 structure:**
   ```html
   <div class="resume">
     <header class="resume-header">...</header>
     <section class="summary">...</section>
     <section class="experience">...</section>
     <section class="skills">...</section>
     <section class="education">...</section>
   </div>
   ```

2. **Applies the CSS classes defined in resume-base.css:**
   - `.resume-header`, `.resume-name`, `.resume-title`, `.contact-info`
   - `.section-title`, `.experience-item`, `.item-header`, `.item-title`
   - `.achievements`, `.skill-tags`, `.skill-tag`
   - `.education-item`, `.degree`, `.institution`

3. **Maintains content fidelity:**
   - Preserve all text exactly as written in the markdown
   - Keep bullet points and achievements intact
   - Maintain date formats and company names

4. **Optimizes layout for target page count:**

   **For 1-page resume:**
   | Setting | Value |
   |---------|-------|
   | Base font | 9pt |
   | @page margin | 0.4in 0.5in |
   | Section margin-bottom | 0.6rem |
   | Experience item margin-bottom | 0.5rem |
   | Achievement line-height | 1.3 |
   | Achievement margin-bottom | 0.15rem |
   | Summary | Condensed to 2-3 sentences max |
   | Skills | Inline comma-separated, no bullets |

   **For 2-page resume (default):**
   | Setting | Value |
   |---------|-------|
   | Base font | 10pt |
   | @page margin | 0.55in 0.6in |
   | Section margin-bottom | 1.15rem |
   | Experience item margin-bottom | 1rem |
   | Achievement line-height | 1.48 |
   | Achievement margin-bottom | 0.3rem |
   | Summary | Full professional summary |
   | Skills | Categorized with bullets |

   **For 3-page resume:**
   | Setting | Value |
   |---------|-------|
   | Base font | 10.5pt |
   | @page margin | 0.6in 0.65in |
   | Section margin-bottom | 1.4rem |
   | Experience item margin-bottom | 1.2rem |
   | Achievement line-height | 1.55 |
   | Achievement margin-bottom | 0.4rem |
   | Summary | Comprehensive with key highlights |
   | Skills | Expanded with proficiency levels |

**IMPORTANT Design Requirements:**
- The HTML must be self-contained with embedded CSS (no external dependencies)
- All fonts must use web-safe font stacks (no Google Fonts imports)
- Design must be ATS-friendly (parseable text, no images for text)
- Target Letter size paper (8.5" x 11")
- Match target page count through spacing adjustments, not content truncation

## Step 5: Create Complete HTML Document

Generate a complete HTML document structure:

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{Candidate Name} - Resume</title>
  <style>
    /* Base CSS content here */
    /* Theme CSS content here */
    /* Any additional dynamic styles */
  </style>
</head>
<body>
  <div class="resume">
    <!-- Generated resume HTML -->
  </div>
</body>
</html>
```

## Step 6: Save Temporary HTML File

Save the generated HTML to a temporary file for Playwright to render:
- Location: `/tmp/resume_temp_{timestamp}.html`
- This file will be used by Playwright and cleaned up after PDF generation

## Step 7: Render PDF with Playwright

Use the Playwright MCP to render the HTML to PDF:

1. **Navigate to the HTML file:**
   ```
   mcp__playwright__browser_navigate to file:///tmp/resume_temp_{timestamp}.html
   ```

2. **Wait for content to load:**
   ```
   mcp__playwright__browser_wait_for with appropriate timeout
   ```

3. **Generate PDF with professional settings:**
   - Use Letter size paper format
   - Set appropriate margins (0.5 inch)
   - Enable print background for colors
   - Generate high-quality output

4. **Save PDF to output location:**
   - Default: Same directory as input file
   - Filename: `{original_name}_formatted.pdf`

## Step 8: Review PDF Output and Iterative Refinement

**CRITICAL: The formatting process is NOT complete until the PDF has been visually reviewed and confirmed to be properly formatted.**

### 8.1: Visual Review of Generated PDF

After generating the PDF, you MUST:

1. **Navigate to the PDF file** using Playwright to view it
2. **Take screenshots** of each page to assess the layout
3. **Evaluate the following quality criteria:**

| Criteria | Target | Red Flags |
|----------|--------|-----------|
| Page count | **Match target** (1, 2, or 3 pages as specified) | Off by ±1 page from target |
| Page balance | Even content distribution across pages | >25% empty space on final page |
| Header spacing | Clear hierarchy, breathing room | Cramped or excessive gaps |
| Section spacing | Consistent, professional margins | Sections touching or too spread |
| Font readability | Clear at normal viewing distance | Too small (<8.5pt) or too large (>12pt) |
| Content overflow | All text visible and complete | Cut-off text, orphaned bullets |
| Page breaks | Break between logical sections | Mid-bullet or mid-sentence breaks |

**Page count validation:**
- If target is `1` and result is `2+`: Tighten spacing aggressively
- If target is `2` and result is `1`: Expand spacing to fill second page
- If target is `2` and result is `3`: Reduce spacing to fit on two pages
- If target is `3` and result is `2`: Expand spacing for comfortable three-page layout

### 8.2: Automatic Formatting Corrections

If ANY quality criteria fail, automatically adjust and regenerate based on target page count:

**Condensing (need fewer pages):**
| Adjustment | Small (-0.25 page) | Medium (-0.5 page) | Large (-1 page) |
|------------|-------------------|-------------------|-----------------|
| Base font | -0.25pt | -0.5pt | -1pt |
| Line-height | -0.02 | -0.05 | -0.1 |
| Section margin | -10% | -20% | -30% |
| Bullet spacing | -0.05rem | -0.1rem | -0.15rem |
| @page margin | -0.05in | -0.1in | -0.15in |

**Expanding (need more pages):**
| Adjustment | Small (+0.25 page) | Medium (+0.5 page) | Large (+1 page) |
|------------|-------------------|-------------------|-----------------|
| Base font | +0.25pt | +0.5pt | +0.75pt |
| Line-height | +0.02 | +0.05 | +0.08 |
| Section margin | +10% | +20% | +30% |
| Bullet spacing | +0.05rem | +0.1rem | +0.15rem |
| @page margin | +0.05in | +0.08in | +0.1in |

**For page 1 mostly empty (content pushed to later pages):**
- Check for overly aggressive `page-break-inside: avoid` on sections
- Ensure only `.experience-item` and similar small elements have page-break protection
- Remove page-break-inside from entire `section` elements

**For uneven page distribution:**
- Calculate ideal content split for target page count
- Adjust spacing incrementally (0.05rem steps)
- Target 85-95% content fill on each page

**1-page specific constraints:**
- Minimum font size: 8.5pt (readability floor)
- If content cannot fit at 8.5pt, warn user that content may need manual reduction
- Consider suggesting removal of older/less relevant experience

**3-page specific constraints:**
- Maximum font size: 11pt (avoid looking sparse)
- Ensure page 3 has at least 40% content (not just footer sections)
- If content doesn't justify 3 pages, suggest 2-page format instead

### 8.3: User Feedback Loop

After automatic adjustments, if the result still needs refinement:

1. **Show the user the current PDF screenshot**
2. **Ask specific questions:**
   - "Does this spacing look appropriate?"
   - "Would you prefer more/less whitespace between sections?"
   - "Any specific elements that need adjustment?"

3. **Accept user tweaks:**
   - Font size adjustments
   - Margin preferences
   - Section ordering suggestions
   - Color/accent preferences

4. **Regenerate with user feedback** and repeat until approved

### 8.4: Maximum Iterations

- **Automatic adjustments:** Up to 3 iterations without user input
- **User-guided refinements:** Continue until user approves
- **Timeout:** After 5 total iterations, present best result and ask for explicit approval

## Step 9: Cleanup and Final Report

1. **Delete temporary HTML file(s)** - remove all `/tmp/resume_temp_*.html` files created during iteration
2. **Close Playwright browser**
3. **Verify final PDF was created successfully**
4. **Report:**
   - Output location and file size
   - Number of pages
   - Number of iterations performed
   - Any adjustments made from original settings

## Output Format

The generated PDF will include:
- Professional typography appropriate to the selected theme
- Clean visual hierarchy with proper spacing
- ATS-compatible text layer
- Print-optimized formatting
- Clickable links for email, LinkedIn, GitHub, portfolio

## Example Usage

```bash
# Modern theme, auto page count (default)
/formatresume OutputResumes/Step3_Final_Resume_Director_CompanyName_2025-01-15.md

# Classic theme, auto page count
/formatresume OutputResumes/Step3_Final_Resume_VP_FinancialServices_2025-01-15.md classic

# Minimal theme, auto page count
/formatresume OutputResumes/Step3_Final_Resume_Engineer_StartupName_2025-01-15.md minimal

# Modern theme, force 1-page (entry-level/targeted)
/formatresume OutputResumes/Step3_Final_Resume_Analyst_Startup_2025-01-15.md modern 1

# Classic theme, force 2-page (standard professional)
/formatresume OutputResumes/Step3_Final_Resume_Manager_Enterprise_2025-01-15.md classic 2

# Modern theme, force 3-page (executive/academic)
/formatresume OutputResumes/Step3_Final_Resume_VP_GlobalCorp_2025-01-15.md modern 3
```

**When to use each page count:**

| Pages | Best For | Typical Experience |
|-------|----------|-------------------|
| 1 | Entry-level, career changers, targeted applications | 0-5 years |
| 2 | Mid-career professionals, most job applications | 5-15 years |
| 3 | Senior executives, academics, extensive portfolios | 15+ years |

## Implementation Notes

### Using the Frontend Design Skill

Invoke the frontend-design skill to generate the HTML structure. The skill will:
- Create distinctive, production-grade HTML
- Apply creative design within the theme constraints
- Ensure the output avoids generic "AI-generated" aesthetics
- Optimize the layout for the specific content length

### Handling Long Resumes

For resumes with extensive experience:
- Reduce font sizes slightly while maintaining readability
- Tighten line-height and margins
- Consider two-column layout for skills section
- Allow graceful page breaks between experience items

### Error Handling

If any step fails:
- Report the specific error with context
- Suggest corrective actions
- Clean up any temporary files created

---

Now executing the resume formatting pipeline with file `$1`, theme `$2`, and target pages `$3`...
