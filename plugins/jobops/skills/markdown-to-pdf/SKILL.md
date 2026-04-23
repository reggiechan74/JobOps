---
description: Convert markdown documents to professional PDF format with Obsidian-style preview
disable-model-invocation: true
---

I'll convert your specified markdown documents to professional PDF format using Obsidian-style preview styling.

**Arguments:**
- `$1`: File path, glob pattern, or directory to convert (required, absolute or relative to workspace)
  - Specific file: `Applications/AcmeCorp_Director_2026-04-23/assessment/rubric.md`
  - Pattern: `Applications/AcmeCorp_Director_2026-04-23/*.md` (all markdown in folder)
  - Multiple files: Space-separated list
- `$2`: Output directory (optional, defaults to the same directory as each source file)

**Supported Conversion Patterns:**

### Single File Conversion
```
/pdf Applications/AcmeCorp_Director_2026-04-23/assessment/rubric.md
/pdf Applications/AcmeCorp_Director_2026-04-23/interview/briefing.md
/pdf Career_Analysis/idealjob_20260423.md
/pdf Crisis_Management/severance_review_20260423.md
```

### Pattern-Based Conversion
```
/pdf "Applications/AcmeCorp_Director_2026-04-23/assessment/*.md"
/pdf "Applications/*/interview/briefing.md"
/pdf "Company_Intelligence/AcmeCorp/*.md"
```

### Directory-Based Conversion
```
/pdf Applications/AcmeCorp_Director_2026-04-23/
/pdf Career_Analysis/
/pdf Crisis_Management/
```

### With Custom Output Directory
```
/pdf Applications/AcmeCorp_Director_2026-04-23/assessment/rubric.md PDF_Exports/
/pdf "Applications/*/interview/briefing.md" PDF_Output/
```

## Step 1: System Verification

First, I'll verify that the required tools are installed:
- **pandoc**: Markdown to PDF conversion engine
- **wkhtmltopdf**: PDF rendering engine for high-quality output

If either tool is missing, I'll install it automatically.

## Step 2: File Discovery and Validation

I'll locate and validate the specified files based on your input pattern: @$1

Path resolution:
- Absolute paths are used verbatim.
- Relative paths are resolved against the current working directory.
- Glob patterns and directories are expanded to a list of `.md` files.

I'll check:
- Files exist and are accessible
- Files are markdown format (.md extension)
- Output directory exists (or can be created)
- No naming conflicts with existing PDF files

## Step 3: PDF Conversion with Obsidian Styling

I'll convert each file using pandoc with custom Obsidian-style CSS for beautiful, readable PDFs.

### Obsidian Styling Features:
- **Clean Typography**: Modern sans-serif fonts matching Obsidian's preview
- **Professional Spacing**: Proper line height and margins for readability
- **Code Block Styling**: GitHub-style code blocks with syntax highlighting
- **Table Formatting**: Clean, bordered tables with hover effects
- **Heading Hierarchy**: Clear visual hierarchy with bottom borders
- **Link Styling**: Purple accent color matching Obsidian theme
- **Blockquote Design**: Left border with subtle background
- **List Formatting**: Proper indentation and spacing
- **Print Optimization**: Clean page breaks and PDF-specific styling

### Conversion Command:
```bash
pandoc "input.md" -o "output.pdf" \
  --pdf-engine=wkhtmltopdf \
  --css="${CLAUDE_PLUGIN_ROOT}/styles/obsidian.css" \
  --standalone \
  --embed-resources \
  --metadata title="Document Title"
```

### Quality Features:
- **High Resolution**: Professional print-quality output
- **Readable Layout**: 800px max width, centered content
- **Consistent Formatting**: All markdown elements properly styled
- **Professional Colors**: Clean color scheme with good contrast
- **Page Break Control**: Smart page breaks for headings and tables
- **Link Preservation**: Clickable links in PDF output
- **Embedded Styling**: Self-contained PDFs with embedded CSS

## Step 4: Conversion Process

For each markdown file, I'll:

1. **Validate the file** exists and is readable
2. **Determine output path** — same directory as the source file by default, or the custom directory from `$2`
3. **Execute pandoc conversion** with Obsidian CSS stylesheet
4. **Verify PDF creation** and check file integrity
5. **Report success/errors** for each file

### File Naming:
- Original: `rubric.md`
- Output: `rubric.pdf`
- Location: Same directory as source (or custom output directory from `$2`)

## Step 5: Success Verification

After conversion, I'll verify:
- PDF files created successfully
- File sizes are reasonable (not empty)
- All requested files were processed
- Clear success summary with file locations

## Expected Output

Your documents will be converted with these patterns:
- `rubric.md` -> `rubric.pdf`
- `briefing.md` -> `briefing.pdf`
- `idealjob_20260423.md` -> `idealjob_20260423.pdf`

## Obsidian-Style Features Applied

The converted PDFs feature:
- **Modern Typography**: -apple-system, Segoe UI, Roboto font stack
- **Clean Headers**: H1 and H2 with bottom borders, proper hierarchy
- **Code Blocks**: GitHub-style with subtle background and border
- **Tables**: Professional borders, hover effects, header styling
- **Blockquotes**: Left border accent with light background
- **Lists**: Proper indentation and spacing
- **Links**: Purple accent color (#7c3aed)
- **Spacing**: 1.6 line height for optimal readability
- **Page Layout**: 40px margins, 800px max width, centered

## Advanced Options

### Batch Conversion with Progress
I'll process multiple files sequentially and provide progress updates for each file.

### Error Handling
If any file fails to convert:
- I'll report the specific error
- Continue processing remaining files
- Provide summary of successes and failures

### Output Organization
- Original markdown files remain unchanged
- PDF files created with matching names
- Output directory defaults to the source file's own directory
- Optional custom output directory via `$2`

## Implementation

I'll execute the conversion following these steps:

1. **Check system dependencies** (pandoc, wkhtmltopdf)
2. **Install missing tools** if needed (automatic)
3. **Resolve file patterns** to get list of files to convert
4. **Create output directory** if custom location specified
5. **Convert each file** with Obsidian CSS styling
6. **Verify results** and report success/failure for each file
7. **Provide summary** with locations of created PDFs

Your markdown documents will be transformed into beautiful, Obsidian-styled PDFs with professional formatting, clean typography, and excellent readability - perfect for sharing, printing, or archival purposes.
