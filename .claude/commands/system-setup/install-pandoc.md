---
description: Install pandoc for document conversion functionality
argument-hint: [force]
---

I'll install pandoc, the universal document converter required for Step 5 document conversion functionality.

**Arguments:**
- `$1`: `force` (optional) - Forces reinstallation even if pandoc is already installed

## Step 1: System Detection and Current Status

First, let me check your current system and pandoc installation status:

```bash
echo "System Information:"
uname -a
echo ""
echo "Checking for existing pandoc installation:"
which pandoc && pandoc --version || echo "Pandoc not found"
```

## Step 2: Platform-Specific Installation

Based on your system, I'll use the appropriate installation method:

### Linux (Ubuntu/Debian)
```bash
sudo apt-get update
sudo apt-get install -y pandoc
```

### Linux (CentOS/RHEL/Fedora)
```bash
# For DNF-based systems (Fedora, newer RHEL/CentOS)
sudo dnf install -y pandoc

# For YUM-based systems (older RHEL/CentOS)
sudo yum install -y pandoc
```

### macOS
```bash
# Using Homebrew (recommended)
brew install pandoc

# Alternative: MacPorts
# sudo port install pandoc
```

### Windows (WSL/Git Bash)
```bash
# Using Chocolatey
choco install pandoc

# Using Scoop
# scoop install pandoc
```

## Step 3: Installation Verification

After installation, I'll verify pandoc is working correctly:

```bash
echo "Verifying pandoc installation:"
pandoc --version
echo ""
echo "Testing basic conversion functionality:"
echo "# Test Document" | pandoc -f markdown -t html
```

## Step 4: Feature Testing

I'll test pandoc's ability to handle the specific conversions we need:

```bash
echo "Testing Word document conversion capability:"
echo -e "# Test Resume\n\n**Experience**\n- Achievement 1\n- Achievement 2" | pandoc -f markdown -t docx -o test_conversion.docx
ls -la test_conversion.docx 2>/dev/null && echo "✓ DOCX conversion successful" || echo "✗ DOCX conversion failed"
rm -f test_conversion.docx
```

## Step 5: Configuration Recommendations

After successful installation, I'll provide configuration recommendations:

### Optional Enhancements
- **LaTeX Support**: For advanced PDF generation
- **Additional Formats**: Support for more output formats
- **Custom Templates**: Professional Word templates for better formatting

## Installation Troubleshooting

### Common Issues and Solutions:

**Permission Errors:**
- Use `sudo` for system-wide installation on Linux/macOS
- Run as administrator on Windows

**Package Manager Not Found:**
- Ubuntu/Debian: Install `apt` packages
- macOS: Install Homebrew first: `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
- Windows: Install Chocolatey or Scoop first

**Outdated Package Repositories:**
- Run `sudo apt-get update` on Ubuntu/Debian
- Run `brew update` on macOS

**Network Issues:**
- Check internet connectivity
- Try alternative package managers
- Download from pandoc.org directly if needed

## Success Confirmation

Upon successful installation, you'll see:
```
✓ Pandoc installed successfully
✓ Version: [version number]
✓ DOCX conversion capability confirmed
✓ Ready for Step 5 document conversion
```

## Next Steps

After pandoc installation:
1. **Test the conversion system**: `/convert` command will now work
2. **Convert existing documents**: Try `/convert all` to convert your current resume/cover letter files
3. **Professional formatting**: Your markdown documents can now be converted to submission-ready Word format

## Integration with Resume Optimizer

With pandoc installed, you can now:
- Convert Step 3 final resumes to professional Word format
- Convert Step 4 cover letters with properly formatted requirements tables
- Batch convert all documents for job applications
- Maintain consistent professional formatting across all application materials

The `/convert` command will automatically detect pandoc and provide high-quality document conversion for your resume optimization workflow.

## Verification Command

To verify everything is working correctly after installation:
```bash
pandoc --version && echo "✓ Pandoc ready for Resume Optimizer Step 5"
```