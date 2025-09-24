---
description: Convert markdown documents to professional Word DOCX format using pandoc
argument-hint: <file-path-or-pattern> [output-directory]
---

I'll convert your specified markdown documents to professional Word DOCX format using pandoc.

**Arguments:**
- `$1`: File path, pattern, or type to convert (required)
  - Specific file: `OutputResumes/Step3_Final_Resume_Director_Alto_2025-09-23.md`
  - Pattern: `OutputResumes/Step3_*.md` (all Step 3 resumes)
  - Type: `resume` or `coverletter` or `all`
- `$2`: Output directory (optional, defaults to same directory as source files)

**Supported Conversion Types:**

### Single File Conversion
```
/convert OutputResumes/Step3_Final_Resume_Director_Alto_2025-09-23.md
```

### Pattern-Based Conversion
```
/convert "OutputResumes/Step3_*.md"
/convert "OutputResumes/Step4_*.md"
```

### Type-Based Conversion
```
/convert resume          # All final resumes (Step 3)
/convert coverletter     # All cover letters (Step 4)
/convert all            # All documents in OutputResumes/
```

## Step 1: System Verification

First, I'll check if pandoc is installed and ready:

## Step 2: File Discovery and Validation

I'll locate and validate the specified files based on your input pattern: @$1

## Step 3: Direct Conversion Process

I'll convert the files directly using pandoc in the main shell session:

### Conversion Features:
- **Professional Formatting**: Clean, business-appropriate styling
- **Table Preservation**: Requirements tables in cover letters render correctly
- **Font Optimization**: Uses standard professional fonts (Calibri, Arial)
- **Layout Optimization**: Proper spacing, margins, and page breaks
- **ATS Compatibility**: Ensures compatibility with applicant tracking systems

### Quality Checks:
- ✓ Verify pandoc conversion completed successfully
- ✓ Check DOCX file integrity and size
- ✓ Validate table formatting (especially requirements matrices)
- ✓ Ensure proper character encoding
- ✓ Confirm professional appearance

### File Organization:
- Original `.md` files remain unchanged
- New `.docx` files created with matching names
- Consistent directory structure maintained
- Clear success/error reporting

## Implementation

I'll execute the conversion using these steps:

1. **Verify pandoc installation**
2. **Find matching files** using glob patterns
3. **Convert each file** using pandoc with professional formatting options
4. **Verify successful conversion** and report results

**Enhanced Pandoc command used:**
```bash
pandoc "input.md" -o "input.docx" \
  --from markdown \
  --to docx \
  --standalone \
  --wrap=none \
  --metadata title="Professional Resume" \
  --highlight-style=monochrome \
  --metadata margin-top=1in \
  --metadata margin-bottom=1in \
  --metadata margin-left=1in \
  --metadata margin-right=1in
```

**Professional formatting targets:**
- **Length**: Target 3-page maximum (like reference PDF)
- **Bullet Points**: Clean rendering without line break issues
- **Typography**: Professional fonts and consistent spacing
- **Margins**: Standard 1-inch margins for business documents
- **Layout**: Compact but readable, matching industry standards

**Formatting options applied:**
- `--standalone`: Complete document structure
- `--wrap=none`: Prevents bullet point formatting issues
- `--highlight-style=monochrome`: Professional appearance
- `--metadata`: Sets margins and document properties
- Target output matches 3-page reference PDF format

## Expected Output

Your documents will be converted with these naming patterns:
- `Step3_Final_Resume_Director_Alto_2025-09-23.md` → `Step3_Final_Resume_Director_Alto_2025-09-23.docx`
- `Step4_CoverLetter_Director_Alto_2025-09-23.md` → `Step4_CoverLetter_Director_Alto_2025-09-23.docx`

## Professional Formatting Applied

The converted documents feature:
- **Resumes**: Single-page optimization, professional fonts, consistent bullet styling
- **Cover Letters**: Business format, properly formatted requirements tables, signature space
- **Both**: Clean margins, appropriate spacing, ATS-friendly formatting

## Success Verification

After conversion, I'll verify:
- ✓ DOCX files created successfully
- ✓ File sizes are reasonable
- ✓ Professional formatting preserved

Your markdown documents will be transformed into professional Word documents ready for employer submission while preserving all the strategic content and formatting from your carefully crafted application materials.