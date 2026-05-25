# markdown-to-LaTeX-PDF Skill — Phase 1 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Ship a new JobOps skill (`latex-pdf`) that converts a markdown resume to a polished, multi-page PDF via a config-driven LaTeX pipeline, replacing `/pdf`'s pandoc+wkhtmltopdf path for the resume doctype with zero-config defaults that reproduce the proven `navy-serif` design.

**Architecture:** Pipeline is `markdown → pandoc(markdown→LaTeX body) → generate preamble from config → assemble → xelatex (2 passes) → PDF`. The user only edits markdown content and an optional JSON config; the skill GENERATES the `.tex`. Phase 1 ships the resume doctype with a single bundled preamble template and a default config that encodes the navy-serif theme exactly. Cover letter and generic-document doctypes (Phase 2/3) and `/pdf` repointing + old-skill deprecation (Phase 3) are explicitly out of scope here.

**Tech Stack:** Skill markdown (Claude-executed), `pandoc` (md→latex body), `xelatex` (texlive-xetex), `fontspec` + `titlesec` + `needspace` + `enumitem` + `microtype` + `xcolor` + `hyperref`, optional `pdftoppm` for preview PNG. Font checking via `fc-list`. Template substitution via `envsubst`-style `sed` pipeline (no Python dependency).

**Source spec:** `SourceMaterial/JobOps_MarkdownToLatexPdf_Skill_Kickstarter.md` — the embedded navy-serif preamble in that file is the load-bearing payload and is reproduced verbatim in the Phase 1 template.

**Environment note:** This plan was authored on a machine without `xelatex` or `pandoc` installed. Task 9 (smoke test) MUST run on a workstation with `texlive-xetex`, `texlive-fonts-extra` (for Lato; otherwise the skill falls back to TeX Gyre Heros), and `pandoc` installed. If unavailable, install via `sudo apt-get install -y texlive-xetex texlive-fonts-extra pandoc poppler-utils` (poppler for `pdftoppm` previews).

---

## File Structure

**New files:**
- `plugins/jobops/skills/latex-pdf/SKILL.md` — the skill itself (prose pipeline with embedded bash, the only file Claude reads at invocation).
- `plugins/jobops/templates/latex/preamble.resume.tex.template` — verbatim navy-serif preamble with `__PLACEHOLDER__` tokens for config-driven values.
- `plugins/jobops/templates/latex/config.json` — default config (single `navy-serif` theme + `resume` doctype).
- `plugins/jobops/templates/latex/README.md` — short note explaining the templates dir is plugin-bundled defaults; users override via `.jobops/templates/custom/latex/`.

**Modified files (Phase 1 scope):**
- `plugins/jobops/skills/setup/SKILL.md` — Step 5 (template installation) must additionally copy `templates/latex/config.json` to `.jobops/templates/default/latex/config.json` and register a new `templates.active.latex_config: "default"` key in the emitted `.jobops/config.json`.
- `CHANGELOG.md` — add an `[Unreleased]` entry under "Added: `latex-pdf` skill (Phase 1, resume doctype)".

**NOT touched in Phase 1 (explicitly deferred):**
- `markdown-to-pdf/SKILL.md` and `convert-to-pdf/SKILL.md` — no deprecation stubs yet. Phase 3.
- `/pdf` command repointing — Phase 3.
- `plugin.json` version bump and marketplace.json — done at Phase 1 release time via `/version-bump`, not in this plan.

---

## Conventions to Match (already verified against the repo)

1. Skill frontmatter: `description:` + `disable-model-invocation: true` + optional `argument-hint:`.
2. Skill body begins with `## Configuration` block reading `.jobops/config.json`; if missing, emits the canonical "JOBOPS NOT CONFIGURED / Run /jobops:setup" stop message.
3. Template path resolution: `{config.templates.base_dir}/{config.templates.active.<key>}/<filename>`. For Phase 1 we add the key `latex_config` whose value is the subdir name (default: `"default"`), and the file resolves to `{config.templates.base_dir}/{latex_config_value}/latex/config.json`.
4. Plugin-bundled assets are referenced via `${CLAUDE_PLUGIN_ROOT}/templates/latex/...` only when the user has NOT yet run setup (zero-config bootstrap path).
5. Output for resume goes to `{applications_root}/{app_slug}/resume/<basename>.pdf` per the Application Path Resolution protocol in `docs/ARCHITECTURE.md` §4.

---

### Task 1: Scaffold the templates dir and the default LaTeX preamble template

**Files:**
- Create: `plugins/jobops/templates/latex/preamble.resume.tex.template`
- Create: `plugins/jobops/templates/latex/README.md`

This is the proven navy-serif preamble from the kickstarter, with config-driven values turned into `__TOKEN__` placeholders the SKILL.md will substitute. Spacing/needspace/font-metric values are calibrated — copy verbatim, do not re-tune.

- [ ] **Step 1: Write the template file**

Write `plugins/jobops/templates/latex/preamble.resume.tex.template` with this exact content:

```latex
% JobOps latex-pdf — resume preamble (navy-serif default).
% Generated by /jobops:latex-pdf — do NOT hand-edit; edit the source markdown
% and the JSON config instead. This file is regenerated on every run.

\documentclass[__FONT_SIZE_PT__pt]{article}

\usepackage[letterpaper,top=__MARGIN_TOP__in,bottom=__MARGIN_BOTTOM__in,left=__MARGIN_LEFT__in,right=__MARGIN_RIGHT__in]{geometry}
\usepackage{fontspec}
\usepackage[table]{xcolor}
\usepackage{titlesec}
\usepackage{enumitem}
\usepackage{array}
\usepackage{microtype}
\usepackage{needspace}
\usepackage[hidelinks]{hyperref}

% ---- Palette (from theme.accent_rgb / theme.muted_rgb) ----
\definecolor{accent}{RGB}{__ACCENT_R__,__ACCENT_G__,__ACCENT_B__}
\definecolor{muted}{RGB}{__MUTED_R__,__MUTED_G__,__MUTED_B__}

% ---- Fonts (from theme.main_font / theme.heading_font / theme.heading_font_fallback) ----
\setmainfont{__MAIN_FONT__}
\IfFontExistsTF{__HEADING_FONT__}{%
  \newfontfamily\headfont{__HEADING_FONT__}%
  \newfontfamily\lightfont{__HEADING_FONT__}%
}{%
  \newfontfamily\headfont{__HEADING_FONT_FALLBACK__}%
  \newfontfamily\lightfont{__HEADING_FONT_FALLBACK__}%
}

\hypersetup{colorlinks=true,urlcolor=accent,linkcolor=accent}

% ---- Section style (from doctype.section_letterspace) ----
\titleformat{\section}
  {\headfont\bfseries\color{accent}\fontsize{11}{13}\selectfont\addfontfeature{LetterSpace=__SECTION_LETTERSPACE__}}
  {}{0pt}{}
  [{\vspace{1pt}\color{accent}\titlerule[0.6pt]}]
\titlespacing*{\section}{0pt}{11pt}{5pt}

% ---- Lists (from doctype.list_itemsep_pt) ----
\setlist[itemize]{leftmargin=1.25em,itemsep=__LIST_ITEMSEP_PT__pt,topsep=2.6pt,parsep=0pt,
  label=\textcolor{accent}{\small\textbullet}}

% ---- Body spacing (from theme.line_spread) ----
\setlength{\parindent}{0pt}
\setlength{\parskip}{0.40em}
\linespread{__LINE_SPREAD__}

% ---- Role helpers (needspace prevents orphaned role headings) ----
\newcommand{\role}[2]{%
  \needspace{5\baselineskip}%
  \vspace{4pt}%
  \noindent{\headfont\bfseries\color{accent}\fontsize{11}{13}\selectfont #1}%
  \hfill{\headfont\small\color{muted}#2}\par\vspace{0.5pt}}
\newcommand{\subrole}[1]{\noindent\textit{#1}\par\vspace{1pt}}

\begin{document}
__BODY__
\end{document}
```

Token list (must match exactly what Task 4 substitutes): `__FONT_SIZE_PT__`, `__MARGIN_TOP__`, `__MARGIN_BOTTOM__`, `__MARGIN_LEFT__`, `__MARGIN_RIGHT__`, `__ACCENT_R__`, `__ACCENT_G__`, `__ACCENT_B__`, `__MUTED_R__`, `__MUTED_G__`, `__MUTED_B__`, `__MAIN_FONT__`, `__HEADING_FONT__`, `__HEADING_FONT_FALLBACK__`, `__SECTION_LETTERSPACE__`, `__LIST_ITEMSEP_PT__`, `__LINE_SPREAD__`, `__BODY__`.

- [ ] **Step 2: Write the README**

Write `plugins/jobops/templates/latex/README.md`:

```markdown
# JobOps latex-pdf LaTeX templates

This directory is plugin-bundled. `/jobops:setup` copies it to
`.jobops/templates/default/latex/`. Treat `default/` as read-only — edit
copies under `.jobops/templates/custom/latex/` and switch via
`config.templates.active.latex_config = "custom"`.

Files:
- `preamble.resume.tex.template` — navy-serif resume preamble with `__TOKEN__`
  substitutions. Tokens are listed at the top of the template.
- `config.json` — themes + doctypes registry. See SKILL.md for schema.

The skill GENERATES the final `.tex`. You do not edit `.tex` files directly.
```

- [ ] **Step 3: Commit**

```bash
git add plugins/jobops/templates/latex/preamble.resume.tex.template plugins/jobops/templates/latex/README.md
git commit -m "feat(latex-pdf): add bundled resume LaTeX preamble template"
```

---

### Task 2: Default LaTeX config.json

**Files:**
- Create: `plugins/jobops/templates/latex/config.json`

This is the zero-config default. Reproduces the navy-serif theme + resume doctype exactly. Schema is the same shape proposed in the kickstarter; the SKILL.md (Task 4) validates against it.

- [ ] **Step 1: Write the config**

```json
{
  "default_theme": "navy-serif",
  "themes": {
    "navy-serif": {
      "main_font": "TeX Gyre Pagella",
      "heading_font": "Lato",
      "heading_font_fallback": "TeX Gyre Heros",
      "accent_rgb": [27, 54, 93],
      "muted_rgb": [96, 103, 112],
      "font_size_pt": 11,
      "line_spread": 1.05
    }
  },
  "doctypes": {
    "resume": {
      "theme": "navy-serif",
      "margins_in": { "top": 0.7, "bottom": 0.7, "left": 0.85, "right": 0.85 },
      "page_fill_target": 0.92,
      "list_itemsep_pt": 2.3,
      "section_letterspace": 7.0
    }
  }
}
```

- [ ] **Step 2: Verify it parses**

Run: `python3 -c "import json; json.load(open('plugins/jobops/templates/latex/config.json'))" && echo OK`
Expected: `OK`

- [ ] **Step 3: Commit**

```bash
git add plugins/jobops/templates/latex/config.json
git commit -m "feat(latex-pdf): add default LaTeX config (navy-serif + resume doctype)"
```

---

### Task 3: Update /jobops:setup to scaffold the new templates and register the new active key

**Files:**
- Modify: `plugins/jobops/skills/setup/SKILL.md`

Two surgical edits in Step 5 (Template installation) and Step 6 (Write config). The setup skill must:
(a) copy the new `latex/` subdir into `.jobops/templates/default/latex/`, and
(b) include `"latex_config": "default"` in the emitted `templates.active` block so `latex-pdf` can resolve the path via the standard contract.

- [ ] **Step 1: Modify Step 5 (Template installation)**

Locate the `## Step 5: Template installation` section in `plugins/jobops/skills/setup/SKILL.md`. The current copy command uses a flat glob; change it so the `latex/` subdir is also copied (preserving structure).

Change FROM:
```bash
cp ${CLAUDE_PLUGIN_ROOT}/templates/* .jobops/templates/default/
```

Change TO:
```bash
cp ${CLAUDE_PLUGIN_ROOT}/templates/*.md .jobops/templates/default/ 2>/dev/null || true
cp -r ${CLAUDE_PLUGIN_ROOT}/templates/latex .jobops/templates/default/
```

Also update the "Report the count of files copied" line: expected count is now **3 markdown templates + 1 `latex/` subdir** (containing `config.json`, `preamble.resume.tex.template`, `README.md`).

Update the fallback `find` line equivalently:
```bash
find ${CLAUDE_PLUGIN_ROOT}/templates -maxdepth 1 -type f -exec cp {} .jobops/templates/default/ \;
cp -r ${CLAUDE_PLUGIN_ROOT}/templates/latex .jobops/templates/default/
```

- [ ] **Step 2: Modify Step 6 (Write config) to register latex_config**

In the JSON block under `## Step 6: Write .jobops/config.json`, find the `templates.active` object and add a new key. Change FROM:

```json
    "active": {
      "assessment_rubric_framework": "default",
      "evidence_verification_framework": "default",
      "assessment_report_structure": "default"
    }
```

Change TO:

```json
    "active": {
      "assessment_rubric_framework": "default",
      "evidence_verification_framework": "default",
      "assessment_report_structure": "default",
      "latex_config": "default"
    }
```

- [ ] **Step 3: Verify the setup skill still parses cleanly**

Run: `head -200 plugins/jobops/skills/setup/SKILL.md | grep -E "(latex|active|Step 5|Step 6)"`
Expected output contains both the new `cp -r ...latex` line and the new `"latex_config": "default"` line.

- [ ] **Step 4: Commit**

```bash
git add plugins/jobops/skills/setup/SKILL.md
git commit -m "feat(setup): scaffold latex/ templates and register latex_config key"
```

---

### Task 4: Write the latex-pdf SKILL.md (the main deliverable)

**Files:**
- Create: `plugins/jobops/skills/latex-pdf/SKILL.md`

This is the only file Claude actually executes when the user invokes the skill. Everything else is data. The skill is prose-with-bash, modeled on the existing `markdown-to-pdf` and `buildresume` skills.

Pipeline the skill must implement, in order:
1. **Configuration block** (canonical "JOBOPS NOT CONFIGURED" stop pattern).
2. **Argument parsing**: `$1` = source markdown file (required), `$2` = theme name override (optional), `$3` = doctype override (optional; defaults to inferred). Phase 1: theme arg is wired but only `navy-serif` is present; doctype is forced to `resume`. Theme arg validation is in place so Phase 2 can drop in `modern`/`classic`/`minimal` without re-touching the skill.
3. **Dependency check**: verify `pandoc`, `xelatex` exist; if not, print the install command for the user's platform and stop. (Auto-install only if `sudo -n` succeeds AND user previously consented; otherwise print and stop.)
4. **Config load**: read `.jobops/config.json`; resolve `latex_config_dir = {templates.base_dir}/{templates.active.latex_config}/latex`; read `{latex_config_dir}/config.json`. If the user-level config file is missing, fall back to `${CLAUDE_PLUGIN_ROOT}/templates/latex/config.json` and warn the user that setup hasn't been run.
5. **Config validation**: keys present, `accent_rgb` and `muted_rgb` are 3-int arrays in [0,255], margins numeric, fonts are strings. On failure, print the offending key path and stop.
6. **Resolve effective settings**: pick doctype, then theme name = arg-override or `doctype.theme` or `default_theme`. Merge theme + doctype settings.
7. **Font availability check**: run `fc-list | grep -Fi "<font>"` for the configured `main_font` and `heading_font`. If `heading_font` is missing, the LaTeX `\IfFontExistsTF` guard in the template will fall back at compile time, but warn the user up front with the substitution that will happen.
8. **Application Path Resolution**: parse `{Company}_{Role}_{YYYYMMDD}` from the markdown filename or its parent folder. Output dir = `{applications_root}/{slug}/resume/`. If parsing fails, write next to the source file and warn.
9. **Pandoc body generation**: `pandoc "$source.md" -t latex -o "$tmp/body.tex" --no-highlight --wrap=none`. (Phase 1 styling fidelity: bold/italic/sections/lists from standard pandoc output; the `\role`/`\subrole` macros stay defined in the preamble for future Lua-filter Phase 3 work but are not invoked here.)
10. **Preamble generation**: copy `{latex_config_dir}/preamble.resume.tex.template` to `$tmp/resume.tex`, then run a single `sed -i` pipeline replacing every `__TOKEN__` with the resolved value. The `__BODY__` token is replaced via a final `sed` that reads the body file. (Exact sed pipeline is given below.)
11. **Compile**: `cd $tmp && xelatex -interaction=nonstopmode resume.tex && xelatex -interaction=nonstopmode resume.tex` (two passes for refs). Capture exit code.
12. **Outputs**: copy `resume.pdf` and `resume.tex` into the resolved output dir as `<sourcebasename>.pdf` and `<sourcebasename>.tex`. Generate `<sourcebasename>_page1.png` via `pdftoppm` if available (silently skip if not).
13. **Page/fill report**: count pages with `pdfinfo` (or `pdftk dump_data`); compute approximate bottom-fill on the last page by rasterizing it and measuring the lowest non-white pixel row (`pdftoppm` + `python -c` one-liner). Print `pages=N, last_page_fill≈XX%`. If `pdfinfo`/`pdftoppm` unavailable, print `pages=N, fill_report=unavailable (install poppler-utils)`.
14. **Final report**: list each output path + the page/fill numbers + any font fallbacks that happened.

- [ ] **Step 1: Write the skill**

Write `plugins/jobops/skills/latex-pdf/SKILL.md`:

````markdown
---
description: Convert a markdown resume to a polished PDF via a config-driven LaTeX pipeline (replaces wkhtmltopdf-based /pdf for resumes)
disable-model-invocation: true
argument-hint: "<source.md> [theme] [doctype]"
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.

## Templates

This skill resolves the LaTeX config and preamble via:

  {config.templates.base_dir}/{config.templates.active.latex_config}/latex/<filename>

Templates referenced: `config.json`, `preamble.resume.tex.template`.

If `config.templates.active.latex_config` is unset (older config), fall back to
`${CLAUDE_PLUGIN_ROOT}/templates/latex/` and warn the user to re-run
`/jobops:setup` so their workspace owns a copy they can customize.

## Application Path Resolution

This skill writes a PDF into a per-application folder.

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the markdown filename OR from the
   parent folder name (the existing `buildresume` flow produces files inside
   `{applications_root}/{slug}/resume/`, so the parent's parent is usually the
   slug).
2. Output dir: `{config.directories.applications_root}/{slug}/resume/`.
3. If the slug cannot be parsed, write the PDF next to the source markdown and
   warn the user that application-folder placement was skipped.

## Arguments

- `$1`: Source markdown file (required, absolute or relative to workspace).
- `$2`: Theme name override (optional). Phase 1 ships only `navy-serif`. Any
  other value: stop with "Theme '<name>' is not defined in latex/config.json.
  Available: <list>". Phase 2 will add `modern`/`classic`/`minimal`.
- `$3`: Doctype override (optional). Phase 1 supports only `resume`. Any other
  value: stop with "Doctype '<name>' not supported in Phase 1 (resume only).
  Phase 2 will add coverletter; Phase 3 will add document."

## Step 1: Dependency check

```bash
missing=()
command -v pandoc  >/dev/null || missing+=("pandoc")
command -v xelatex >/dev/null || missing+=("xelatex (texlive-xetex)")
command -v fc-list >/dev/null || missing+=("fc-list (fontconfig)")
if [ ${#missing[@]} -gt 0 ]; then
  echo "MISSING DEPENDENCIES: ${missing[*]}"
  echo "Install on Debian/Ubuntu:"
  echo "  sudo apt-get install -y texlive-xetex texlive-fonts-extra pandoc poppler-utils fontconfig"
  echo "On macOS (Homebrew):"
  echo "  brew install --cask mactex-no-gui && brew install pandoc poppler"
  exit 1
fi
```

Stop on missing deps; do NOT proceed.

## Step 2: Load and validate LaTeX config

Read the LaTeX config:

```bash
# Resolve config path from .jobops/config.json
base_dir=$(jq -r '.templates.base_dir' .jobops/config.json)
latex_dir_name=$(jq -r '.templates.active.latex_config // "default"' .jobops/config.json)
latex_dir="${base_dir}/${latex_dir_name}/latex"

cfg="${latex_dir}/config.json"
if [ ! -f "$cfg" ]; then
  echo "WARN: ${cfg} not found; falling back to plugin-bundled defaults."
  echo "      Run /jobops:setup to install workspace copies you can customize."
  latex_dir="${CLAUDE_PLUGIN_ROOT}/templates/latex"
  cfg="${latex_dir}/config.json"
fi
```

Validate:

```bash
jq -e '
  .default_theme           and
  (.themes | type == "object")    and
  (.doctypes | type == "object")  and
  (.themes[.default_theme] // empty) and
  (.themes | to_entries | all(.value.accent_rgb | length == 3 and all(. >= 0 and . <= 255))) and
  (.themes | to_entries | all(.value.muted_rgb  | length == 3 and all(. >= 0 and . <= 255)))
' "$cfg" >/dev/null || { echo "INVALID latex/config.json"; exit 1; }
```

On failure, print the offending key path and stop.

## Step 3: Resolve effective settings

Determine doctype (`$3` override, else `resume`). Phase 1 hard-rejects anything
other than `resume`.

Determine theme: `$2` override → `doctype.theme` → `default_theme`. If the
named theme is not in `themes`, stop with the "available themes" message.

Pull these resolved values into shell variables (use `jq -r`):

```
THEME_NAME, FONT_SIZE_PT, LINE_SPREAD,
MAIN_FONT, HEADING_FONT, HEADING_FONT_FALLBACK,
ACCENT_R, ACCENT_G, ACCENT_B,
MUTED_R, MUTED_G, MUTED_B,
MARGIN_TOP, MARGIN_BOTTOM, MARGIN_LEFT, MARGIN_RIGHT,
LIST_ITEMSEP_PT, SECTION_LETTERSPACE
```

## Step 4: Font availability check

```bash
fallback_msg=""
if ! fc-list | grep -Fiq "$MAIN_FONT"; then
  echo "ERROR: main_font '$MAIN_FONT' is not installed. Install texlive-fonts-extra or change the theme."
  exit 1
fi
if ! fc-list | grep -Fiq "$HEADING_FONT"; then
  fallback_msg="heading_font '$HEADING_FONT' not found; falling back to '$HEADING_FONT_FALLBACK' at compile time."
  echo "WARN: $fallback_msg"
fi
```

## Step 5: Pandoc body generation

```bash
src="$1"
basename=$(basename "$src" .md)
work=$(mktemp -d)
pandoc "$src" -t latex -o "${work}/body.tex" --no-highlight --wrap=none
```

## Step 6: Generate preamble + assemble

Copy the template and run a single sed pipeline. Use a delimiter unlikely to
appear in values (`|`):

```bash
cp "${latex_dir}/preamble.resume.tex.template" "${work}/${basename}.tex"
tex="${work}/${basename}.tex"

sed -i \
  -e "s|__FONT_SIZE_PT__|${FONT_SIZE_PT}|g" \
  -e "s|__MARGIN_TOP__|${MARGIN_TOP}|g" \
  -e "s|__MARGIN_BOTTOM__|${MARGIN_BOTTOM}|g" \
  -e "s|__MARGIN_LEFT__|${MARGIN_LEFT}|g" \
  -e "s|__MARGIN_RIGHT__|${MARGIN_RIGHT}|g" \
  -e "s|__ACCENT_R__|${ACCENT_R}|g" \
  -e "s|__ACCENT_G__|${ACCENT_G}|g" \
  -e "s|__ACCENT_B__|${ACCENT_B}|g" \
  -e "s|__MUTED_R__|${MUTED_R}|g" \
  -e "s|__MUTED_G__|${MUTED_G}|g" \
  -e "s|__MUTED_B__|${MUTED_B}|g" \
  -e "s|__MAIN_FONT__|${MAIN_FONT}|g" \
  -e "s|__HEADING_FONT_FALLBACK__|${HEADING_FONT_FALLBACK}|g" \
  -e "s|__HEADING_FONT__|${HEADING_FONT}|g" \
  -e "s|__SECTION_LETTERSPACE__|${SECTION_LETTERSPACE}|g" \
  -e "s|__LIST_ITEMSEP_PT__|${LIST_ITEMSEP_PT}|g" \
  -e "s|__LINE_SPREAD__|${LINE_SPREAD}|g" \
  "$tex"

# Substitute __BODY__ from the file (sed's r command + d on the marker):
sed -i -e "/__BODY__/{r ${work}/body.tex" -e "d;}" "$tex"
```

## Step 7: Compile (2 passes)

```bash
( cd "$work" && xelatex -interaction=nonstopmode "${basename}.tex" >/dev/null \
              && xelatex -interaction=nonstopmode "${basename}.tex" >/dev/null ) \
  || { echo "xelatex compile failed; see ${work}/${basename}.log"; exit 1; }
```

On failure, leave `$work` in place and tell the user the log path.

## Step 8: Resolve output dir and copy artifacts

```bash
# Try to detect slug from parent folder structure (buildresume layout):
parent=$(dirname "$src")
grandparent=$(dirname "$parent")
slug=$(basename "$grandparent")
apps_root=$(jq -r '.directories.applications_root' .jobops/config.json)

if [[ "$slug" =~ ^[A-Za-z0-9]+_[A-Za-z0-9]+_[0-9]{8}$ ]] \
   && [[ "$grandparent" == "${apps_root}/"* ]]; then
  outdir="${grandparent}/resume"
else
  outdir="$parent"
  echo "WARN: could not resolve application slug; writing PDF next to source ($outdir)."
fi
mkdir -p "$outdir"

cp "${work}/${basename}.pdf" "${outdir}/${basename}.pdf"
cp "${work}/${basename}.tex" "${outdir}/${basename}.tex"   # advanced override artifact
```

## Step 9: Page/fill report

```bash
pages="?"; fill="unavailable"
if command -v pdfinfo >/dev/null; then
  pages=$(pdfinfo "${outdir}/${basename}.pdf" | awk '/^Pages:/{print $2}')
fi
if command -v pdftoppm >/dev/null && command -v python3 >/dev/null; then
  png="${work}/lastpage.png"
  pdftoppm -png -r 100 -f "$pages" -l "$pages" "${outdir}/${basename}.pdf" "${work}/lastpage" 2>/dev/null
  # pdftoppm appends -N for page N; find the produced file:
  produced=$(ls "${work}"/lastpage-*.png 2>/dev/null | head -1)
  if [ -n "$produced" ]; then
    fill=$(python3 - "$produced" <<'PY'
import sys
from PIL import Image
im = Image.open(sys.argv[1]).convert("L")
w, h = im.size
px = im.load()
last_ink_row = 0
for y in range(h):
    for x in range(w):
        if px[x, y] < 240:
            last_ink_row = y
            break
print(f"{last_ink_row*100//h}%")
PY
)
  fi
fi
echo "REPORT: pages=${pages}, last_page_fill≈${fill}"
```

(If Pillow is not installed, the python step will fail silently; that's
acceptable for Phase 1 — the report just stays "unavailable" and the user
sees the install hint.)

## Step 10: Final report

Print:

```
latex-pdf complete.
  Source:   $src
  Theme:    $THEME_NAME
  Doctype:  resume
  Output:   ${outdir}/${basename}.pdf
  TeX:      ${outdir}/${basename}.tex   (advanced-override artifact)
  Pages:    ${pages}
  Fill:     ${fill}
  Fallbacks: ${fallback_msg:-none}
```

Done. Do not delete `$work` — it's a `mktemp -d` directory the OS reclaims,
and keeping it lets the user inspect the .log on failures.
````

- [ ] **Step 2: Sanity-check the skill markdown shape**

Run: `head -20 plugins/jobops/skills/latex-pdf/SKILL.md`
Expected: frontmatter with `description:`, `disable-model-invocation: true`, `argument-hint:`.

Run: `grep -c "JOBOPS NOT CONFIGURED" plugins/jobops/skills/latex-pdf/SKILL.md`
Expected: `1`.

Run: `grep -E "^## Step" plugins/jobops/skills/latex-pdf/SKILL.md`
Expected: 10 step headings (Step 1 through Step 10).

- [ ] **Step 3: Commit**

```bash
git add plugins/jobops/skills/latex-pdf/SKILL.md
git commit -m "feat(latex-pdf): add Phase 1 skill (resume doctype, navy-serif default)"
```

---

### Task 5: Update CHANGELOG

**Files:**
- Modify: `CHANGELOG.md`

- [ ] **Step 1: Add an Unreleased entry**

Open `CHANGELOG.md`. If there is no `## [Unreleased]` section at the top, create one. Add this bullet under `### Added`:

```markdown
- `latex-pdf` skill (Phase 1): config-driven markdown→LaTeX→PDF pipeline for resumes.
  Ships the proven `navy-serif` theme as zero-config default. Reads
  `.jobops/templates/{active.latex_config}/latex/config.json`; users edit JSON,
  never LaTeX. Replaces `/pdf` for resumes once Phase 3 retires the old
  converters. See `docs/superpowers/plans/2026-05-25-markdown-to-latex-pdf-phase1.md`.
```

And under `### Changed`:

```markdown
- `/jobops:setup` now copies `templates/latex/` into the workspace and registers
  `templates.active.latex_config = "default"` so latex-pdf resolves via the
  standard template contract.
```

- [ ] **Step 2: Commit**

```bash
git add CHANGELOG.md
git commit -m "docs: changelog for latex-pdf Phase 1 + setup changes"
```

---

### Task 6: Smoke test — dependency probe (no-op-friendly)

**Files:**
- None new; this task exercises the skill on the current machine.

If `xelatex` and `pandoc` are absent (the typical Codespaces case), this task verifies the dependency check fires correctly and stops without producing garbage. If they ARE present, proceed to Task 7.

- [ ] **Step 1: Probe deps**

Run: `command -v xelatex >/dev/null && command -v pandoc >/dev/null && echo PRESENT || echo MISSING`

- [ ] **Step 2a: If MISSING — simulate the dep-check failure**

Run the dep-check snippet from `Step 1` of the skill verbatim against the current machine. Verify:
- It prints `MISSING DEPENDENCIES: ...`
- It prints the apt-get hint with `texlive-xetex texlive-fonts-extra pandoc poppler-utils fontconfig`
- Exit status is non-zero.

If the user is in a Codespaces / dev container, install before continuing:
```bash
sudo apt-get update && sudo apt-get install -y texlive-xetex texlive-fonts-extra pandoc poppler-utils fontconfig
```

Then re-run the probe; expected `PRESENT`.

- [ ] **Step 2b: If PRESENT — skip to Task 7**

---

### Task 7: Smoke test — zero-config resume

**Files:**
- Test input: a real resume markdown.

- [ ] **Step 1: Locate or scaffold a sample resume**

Run:
```bash
sample=$(find . -maxdepth 6 -name "step3*final*.md" -type f 2>/dev/null | head -1)
echo "$sample"
```

If empty, ask the user for a path to a step3 final resume. If they have none yet (fresh workspace), generate a minimal one inline at `/tmp/sample_resume.md`:

```markdown
# Jane Doe
**Director of Engineering**

jane@example.com · linkedin.com/in/janedoe · github.com/janedoe · 555-0100

## Summary

Engineering leader with 12+ years of experience scaling backend platforms and
high-performing teams across fintech and SaaS.

## Experience

### Acme Corp
*Director of Engineering — Toronto, ON — Jan 2022 to Present*

- Scaled platform team from 6 to 24 engineers across three squads.
- Delivered the migration from monolith to event-driven services; reduced
  median checkout latency 38%.
- Owned a $4.2M annual cloud budget; brought unit economics to break-even.

### Globex Inc.
*Engineering Manager — Toronto, ON — Mar 2018 to Dec 2021*

- Led the payments-platform team (8 engineers) through PCI-DSS certification.
- Shipped the dispute-management API used by 200+ merchants.

## Skills

Backend platforms, distributed systems, team scaling, FinOps, PCI-DSS,
incident response, hiring loop design.

## Education

**BSc Computer Science**, University of Toronto, 2011.
```

- [ ] **Step 2: Invoke the skill with zero config**

Make sure the workspace has been set up (`.jobops/config.json` exists). If not, run `/jobops:setup` first (accepting all defaults).

Then trigger the skill. Either invoke from a Claude session or replicate the pipeline in a single shell pass. Recommended: invoke through Claude so the SKILL.md prose is actually exercised.

Expected outputs:
- `<basename>.pdf` in `<source_parent>/` (since the synthetic sample is not in an application slug folder, the "writing PDF next to source" warning fires).
- `<basename>.tex` next to it.
- Report line: `REPORT: pages=N, last_page_fill≈XX%`.
- No font-fallback warning if `texlive-fonts-extra` includes Lato; otherwise a single `WARN: heading_font 'Lato' not found; falling back to 'TeX Gyre Heros'`.

- [ ] **Step 3: Visually inspect**

Open the PDF. Confirm:
- Two-column header (name + role at top).
- Navy section rules under each `## SECTION`.
- Bulleted achievements with the small navy bullet glyph.
- No orphaned role headings at page bottoms.
- Reasonable page fill (>80% on the last page for a 1-page resume of this length).

- [ ] **Step 4: Record results in the plan**

Append to the bottom of THIS plan file (`docs/superpowers/plans/2026-05-25-markdown-to-latex-pdf-phase1.md`) a short "Smoke test results" section noting page count, fill %, and any anomalies.

- [ ] **Step 5: Commit any plan/changelog adjustments**

```bash
git add docs/superpowers/plans/2026-05-25-markdown-to-latex-pdf-phase1.md
git commit -m "test(latex-pdf): record zero-config smoke test results"
```

---

### Task 8: Smoke test — non-default override

**Files:**
- Test input: same sample resume.
- Modify: `.jobops/templates/default/latex/config.json` (workspace copy, not the plugin-bundled one).

Verify the config-driven path works. Override one easily-visible value, re-run, confirm the change appears.

- [ ] **Step 1: Tweak a theme value in the workspace config**

Edit `.jobops/templates/default/latex/config.json` and change the accent RGB to a bright burgundy:

```json
"accent_rgb": [128, 0, 32]
```

- [ ] **Step 2: Re-invoke latex-pdf**

Same command as Task 7 Step 2. Expected: the new PDF has burgundy section rules + headings instead of navy.

- [ ] **Step 3: Verify the .tex captures the change**

Run: `grep "definecolor{accent}" <outdir>/<basename>.tex`
Expected: `\definecolor{accent}{RGB}{128,0,32}`

- [ ] **Step 4: Restore the navy default and commit**

Restore `accent_rgb` to `[27, 54, 93]` in the workspace config (the plugin-bundled `plugins/jobops/templates/latex/config.json` should never have been touched). Do not commit the workspace config — it's gitignored.

If the plugin-bundled file was accidentally touched, restore it:
```bash
git checkout plugins/jobops/templates/latex/config.json
```

No commit needed for this task unless the plan file got further edits.

---

### Task 9: Self-review checklist

This is run by the implementer before declaring Phase 1 done.

- [ ] Confirm the skill stops cleanly when `.jobops/config.json` is absent.
- [ ] Confirm the skill stops cleanly when `xelatex` is absent.
- [ ] Confirm the skill stops cleanly on a malformed `latex/config.json` (delete a required key, re-invoke, observe a clear error pointing at the missing key).
- [ ] Confirm `latex-pdf` produces a PDF for the synthetic sample with zero workspace customization.
- [ ] Confirm an edit to `accent_rgb` in the workspace config changes the PDF on the next run without any change to LaTeX.
- [ ] Confirm the generated `.tex` is human-readable and could be hand-edited as the "advanced override" the kickstarter promises.
- [ ] Confirm `plugins/jobops/skills/markdown-to-pdf/SKILL.md` and `plugins/jobops/skills/convert-to-pdf/SKILL.md` are UNTOUCHED (deferred to Phase 3).
- [ ] Confirm `marketplace.json` / `plugin.json` version was NOT bumped here (release handled separately by `/version-bump`).

---

## Phase 2 and 3 (intentionally out of scope, for reference)

- **Phase 2:** add `coverletter` doctype + `preamble.coverletter.tex.template` (letterhead, longtable requirements table with navy header + alternating shade, signature image). Add named theme presets (`modern`/`classic`/`minimal`). Add page-count targeting (`$3` arg interpretation) by iterating `font_size_pt` / `list_itemsep_pt` toward `page_fill_target`.
- **Phase 3:** add `document` doctype (single-file / glob / directory like the old `/pdf`). Repoint the `/pdf` command to `latex-pdf`. Convert `markdown-to-pdf` and `convert-to-pdf` SKILL.md files into thin deprecation stubs that delegate to `latex-pdf` and print a notice for one release; remove in the following release. Optional Lua filter for exact `\role`/`\subrole` macro fidelity from vanilla markdown.
