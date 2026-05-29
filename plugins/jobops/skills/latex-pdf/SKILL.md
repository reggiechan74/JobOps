---
name: latex-pdf
description: Convert a markdown resume to a polished PDF via a config-driven LaTeX pipeline (replaces wkhtmltopdf-based /pdf for resumes)
disable-model-invocation: true
argument-hint: "<source.md> [theme] [pages]"
---

## Configuration

Read `.jobops/config.json`. If missing, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup to initialize your workspace.

Use `config.directories.<key>` for all file paths in this skill.

## Templates

This skill resolves the LaTeX config and preamble via:

  {config.templates.base_dir}/{config.templates.active.latex_config}/latex/<filename>

Templates referenced: `config.json`, `preamble.base.tex.template`,
`preamble.resume.tex.template`, `preamble.coverletter.tex.template`,
`preamble.document.tex.template`, and `omers-filter.lua` (the Tier 2 pandoc
filter). The final `.tex` is assembled as **base + doctype delta + body**:
`preamble.base.tex.template` holds everything shared (fonts, palette,
de-numbering, helper macros, section style, `\role`/`\subrole`, table helpers);
each `preamble.<doctype>.tex.template` is a small delta (margins live in the
base via tokens; the delta sets paragraph spacing, lists, tables, and
`\begin{document}…__BODY__…\end{document}`).

If `config.templates.active.latex_config` is unset (older config), fall back to
`${CLAUDE_PLUGIN_ROOT}/templates/latex/` and warn the user to re-run
`/jobops:setup` so their workspace owns a copy they can customize.

## Markdown authoring contract

The OMERS Lua filter maps a specific markdown shape onto the gold-standard
LaTeX. The `buildresume` and `coverletter` skills MUST emit markdown in this
shape; the filter is tolerant but parity depends on it.

**Resume** (`output_type: resume_*`):

```markdown
# Name, PostNoms
**One-line tagline**
City, ST • (555) 555-5555 • [email](mailto:…) • [linkedin.com/in/…](https://…) • [github.com/…](https://…)

## SECTION NAME

### **COMPANY** (optional parenthetical)
*Title | Location | Dates*

- Achievement bullet with a metric.
```

- The first `#` (name) plus the tagline and contact lines become the centered
  header block — not a section.
- Each `## ` becomes a navy ruled `\section`.
- Each `### **COMPANY**` immediately followed by an italic `*…*` subline becomes
  `\role{COMPANY}{Dates}` + `\subrole{Title — Location}`. The subline is split on
  ` | `; the **last** segment is the right-aligned dates, the rest is the subrole.
  If there is no ` | `, the whole italic line becomes the subrole with no dates.
- Markdown thematic breaks (`---`) are **dropped** — they would render as a
  distracting standalone centered rule. Use them freely as source separators;
  the navy `\section` headings are the only visible dividers.

**Cover letter** (`output_type: cover_letter`):

```markdown
# Name, PostNoms
City, ST | (555) 555-5555 | [email](mailto:…) | LinkedIn: [linkedin.com/in/…](https://…)

Month DD, YYYY

Recipient Name\
Recipient Title\
Recipient Org

Dear FirstName:

Body paragraph…

## Requirements Alignment

| Requirement | Evidence |
|---|---|
| … | … |

On X: evidence paragraph…

Sincerely,

![](signature.png){width=2in}

Name, PostNoms
```

- The first `#` plus the contact line become the left letterhead (name + navy
  rule + contact). The date line that follows is NOT consumed into the header.
- **Contact line fields.** The contact line carries up to four distinct fields:
  `{location} | {phone} | {email} | LinkedIn: {linkedin}`. Phone is its own
  field — never concatenate it onto another field (the fused
  `(555) 555-0123name@example.com` artifact came from a header that had no
  phone slot). Join fields with ` | ` (space-pipe-space) and **omit any empty
  field cleanly** — no orphan separators, no doubled ` |  | `. If the candidate
  has no phone on record, drop the field and the adjacent separator so the line
  reads `City, ST | email | LinkedIn: …`. The filter renders the contact line
  verbatim, so the separation must be correct in the markdown.
- Use a trailing `\` for hard line breaks in the recipient block.
- The 2-column requirements table renders with a navy header row, white bold
  text, zebra striping, and a 29/71 column split.
- Keep the signature image width in the markdown (`{width=2in}` ≈ 5 cm);
  `keepaspectratio` in the base preamble prevents distortion regardless.

## Application Path Resolution

This skill writes a PDF into a per-application folder.

1. Parse `{Company}_{Role}_{YYYYMMDD}` from the markdown filename OR from the
   parent folder name (the existing `buildresume` / `coverletter` flows produce
   files inside `{applications_root}/{slug}/{sub}/`, so the parent's parent is
   usually the slug).
2. Output dir is per-doctype: `{applications_root}/{slug}/resume/` for the
   `resume` doctype; `{applications_root}/{slug}/cover-letter/` for the
   `coverletter` doctype.
3. If the slug cannot be parsed, write the PDF next to the source markdown and
   warn the user that application-folder placement was skipped.

## Arguments

- `$1`: Source markdown file (required, absolute or relative to workspace).
- `$2`: Theme name override (optional). Available themes: `navy-serif`, `modern`, `classic`, `minimal`. Unknown name: stop with "Theme '<name>' is not defined in latex/config.json. Available: <list>".
- `$3`: Target page count (optional). One of `1`, `2`, `3`, or `auto` (default `auto`). When 1/2/3, the skill iteratively tunes font_size_pt and list_itemsep_pt (resume) or parskip_em (coverletter) toward the target, up to 3 iterations. When `auto`, the skill compiles once with the doctype's defaults.

Flags:
- `--doctype=resume|coverletter` — override doctype autodetection. Rarely needed.

## Step 1: Dependency check

```bash
missing=()
command -v pandoc  >/dev/null || missing+=("pandoc")
command -v xelatex >/dev/null || missing+=("xelatex (texlive-xetex)")
command -v fc-list >/dev/null || missing+=("fc-list (fontconfig)")
command -v jq      >/dev/null || missing+=("jq")
if [ ${#missing[@]} -gt 0 ]; then
  echo "MISSING DEPENDENCIES: ${missing[*]}"
  echo "Install on Debian/Ubuntu:"
  echo "  sudo apt-get install -y texlive-xetex texlive-fonts-extra pandoc poppler-utils fontconfig jq"
  echo "On macOS (Homebrew):"
  echo "  brew install --cask mactex-no-gui && brew install pandoc poppler jq"
  exit 1
fi

# Optional: pdfinfo (from poppler-utils) — used for page count + page-target iteration.
# If missing, the skill still produces the PDF but cannot enforce $3 page-count targeting.
if ! command -v pdfinfo >/dev/null; then
  echo "WARN: pdfinfo not found (from poppler-utils). Page-count targeting via \$3 will be skipped."
  echo "      Install with: sudo apt-get install -y poppler-utils  (or brew install poppler)"
fi
```

Stop on missing deps; do NOT proceed.

## Step 1.5: Expand $1 into a list of source files

The skill accepts `$1` as a file, a glob, or a directory (matching the old `/pdf` shape). Expand to a list of files; iterate Steps 2-11 once per file.

```bash
input="$1"
if [ -d "$input" ]; then
  # Directory: list *.md non-recursively
  shopt -s nullglob
  sources=( "$input"/*.md )
  shopt -u nullglob
elif [[ "$input" == *[\*\?]* ]]; then
  # Glob expression: expand
  shopt -s nullglob
  # shellcheck disable=SC2206
  sources=( $input )
  shopt -u nullglob
elif [ -f "$input" ]; then
  sources=( "$input" )
else
  echo "ERROR: '$input' is not a file, glob, or directory."
  exit 1
fi

if [ ${#sources[@]} -eq 0 ]; then
  echo "ERROR: no markdown files matched '$input'."
  exit 1
fi

echo "Will process ${#sources[@]} file(s):"
printf "  - %s\n" "${sources[@]}"
```

Steps 2 through 11 below execute for EACH `$src` in `${sources[@]}`. Wrap them in:

```bash
for src in "${sources[@]}"; do
  echo ""
  echo "=== Processing: $src ==="
  # ... Steps 2-11 below, using $src in place of any $1 reference ...
done
```

Inside the loop, `$src` is the current source file. References to `$1` in Step 2's front-matter check and Step 6's pandoc input use `$src` instead.

## Step 2: Detect doctype

Determine the doctype in this order (first match wins):

1. `--doctype=<value>` flag on the command line.
2. YAML front-matter `output_type` field in `$src` (the file's first 30 lines suffice):
   - `output_type ∈ {resume_final, resume_step1, resume_step2, resume_provenance}` → doctype `resume`
   - `output_type ∈ {cover_letter, coverletter}` → doctype `coverletter`
   - (any other value or absent → fall through to next check)
3. Parent-directory name: if the file's immediate parent dir is `cover-letter` → `coverletter`; if `resume` → `resume`.
4. Filename heuristic:
   - contains `cover_letter` → `coverletter`
   - contains `step1`, `step2`, `step3`, or `resume` → `resume`
   - otherwise → fall through
5. Default fallback: `document`.

```bash
# Pull --doctype= flag if present in any positional arg
doctype=""
for a in "$@"; do
  case "$a" in --doctype=*) doctype="${a#--doctype=}";; esac
done

if [ -z "$doctype" ]; then
  fm_type=$(awk '/^---$/{n++;next} n==1 && /^output_type:/{print $2; exit}' "$src" | tr -d '"' || true)
  case "$fm_type" in
    resume_final|resume_step1|resume_step2|resume_provenance) doctype="resume";;
    cover_letter|coverletter)                                  doctype="coverletter";;
  esac
fi

if [ -z "$doctype" ]; then
  case "$(basename "$(dirname "$src")")" in
    cover-letter) doctype="coverletter";;
    resume)       doctype="resume";;
  esac
fi

if [ -z "$doctype" ]; then
  case "$(basename "$src")" in
    *cover_letter*)                          doctype="coverletter";;
    *step1*|*step2*|*step3*|*resume*)        doctype="resume";;
  esac
fi

# Final fallback: document
doctype="${doctype:-document}"

case "$doctype" in
  resume|coverletter|document) ;;
  *) echo "Doctype '$doctype' not supported (Phase 3: resume, coverletter, document)"; exit 1;;
esac
echo "Detected doctype: $doctype"
```

## Step 3: Load and validate LaTeX config

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

## Step 4: Resolve effective settings

Doctype is already resolved (Step 2). Determine theme: `$2` override → `doctypes[doctype].theme` → `default_theme`. If the named theme is not in `themes`, stop with the "available themes" message.

```bash
# Theme resolution
theme_arg="${2:-}"
if [ -n "$theme_arg" ]; then
  THEME_NAME="$theme_arg"
elif jq -e --arg d "$doctype" '.doctypes[$d].theme' "$cfg" >/dev/null; then
  THEME_NAME=$(jq -r --arg d "$doctype" '.doctypes[$d].theme' "$cfg")
else
  THEME_NAME=$(jq -r '.default_theme' "$cfg")
fi
if ! jq -e --arg t "$THEME_NAME" '.themes[$t]' "$cfg" >/dev/null; then
  avail=$(jq -r '.themes | keys | join(", ")' "$cfg")
  echo "Theme '$THEME_NAME' is not defined in latex/config.json. Available: $avail"
  exit 1
fi

# Theme-level values (from .themes[$THEME_NAME])
FONT_SIZE_PT=$(jq -r --arg t "$THEME_NAME" '.themes[$t].font_size_pt'         "$cfg")
LINE_SPREAD=$( jq -r --arg t "$THEME_NAME" '.themes[$t].line_spread'          "$cfg")
MAIN_FONT=$(   jq -r --arg t "$THEME_NAME" '.themes[$t].main_font'            "$cfg")
HEADING_FONT=$(jq -r --arg t "$THEME_NAME" '.themes[$t].heading_font'         "$cfg")
HEADING_FONT_FALLBACK=$(jq -r --arg t "$THEME_NAME" '.themes[$t].heading_font_fallback' "$cfg")
ACCENT_R=$(jq -r --arg t "$THEME_NAME" '.themes[$t].accent_rgb[0]' "$cfg")
ACCENT_G=$(jq -r --arg t "$THEME_NAME" '.themes[$t].accent_rgb[1]' "$cfg")
ACCENT_B=$(jq -r --arg t "$THEME_NAME" '.themes[$t].accent_rgb[2]' "$cfg")
MUTED_R=$( jq -r --arg t "$THEME_NAME" '.themes[$t].muted_rgb[0]'  "$cfg")
MUTED_G=$( jq -r --arg t "$THEME_NAME" '.themes[$t].muted_rgb[1]'  "$cfg")
MUTED_B=$( jq -r --arg t "$THEME_NAME" '.themes[$t].muted_rgb[2]'  "$cfg")

# Doctype-level values (from .doctypes[$doctype])
MARGIN_TOP=$(   jq -r --arg d "$doctype" '.doctypes[$d].margins_in.top'    "$cfg")
MARGIN_BOTTOM=$(jq -r --arg d "$doctype" '.doctypes[$d].margins_in.bottom' "$cfg")
MARGIN_LEFT=$(  jq -r --arg d "$doctype" '.doctypes[$d].margins_in.left'   "$cfg")
MARGIN_RIGHT=$( jq -r --arg d "$doctype" '.doctypes[$d].margins_in.right'  "$cfg")
SECTION_LETTERSPACE=$(jq -r --arg d "$doctype" '.doctypes[$d].section_letterspace // "0"' "$cfg")

# Doctype-specific
if [ "$doctype" = "resume" ]; then
  LIST_ITEMSEP_PT=$(jq -r '.doctypes.resume.list_itemsep_pt' "$cfg")
elif [ "$doctype" = "coverletter" ]; then
  PARSKIP_EM=$(jq -r '.doctypes.coverletter.parskip_em' "$cfg")
elif [ "$doctype" = "document" ]; then
  PARSKIP_EM=$(jq -r '.doctypes.document.parskip_em' "$cfg")
fi
```

## Step 5: Font availability check

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

## Step 6: Pandoc body generation

```bash
# (src is set by the per-file loop wrapper from Step 1.5)
basename=$(basename "$src" .md)
work=$(mktemp -d)

# Copy sibling image assets (signature, logos) so relative \includegraphics paths
# resolve inside the temp compile dir. Without this, xelatex can't read the
# missing image's dimensions and (because the base preamble sets
# keepaspectratio) degrades into a Package-graphics "Division by 0" hard failure
# instead of a clean "file not found". One copy suffices — $work persists across
# the recompile/page-target iteration loop.
srcdir=$(dirname "$src")
find "$srcdir" -maxdepth 1 -type f \( -iname '*.png' -o -iname '*.jpg' -o -iname '*.jpeg' \) \
  -exec cp {} "$work/" \; 2>/dev/null || true

# The OMERS Lua filter maps the markdown onto the hand-built LaTeX vocabulary
# (header block, \section, \role/\subrole, navy+zebra tables). It needs to know
# the doctype; it is pass-through for `document`.
pandoc "$src" -t latex -o "${work}/body.tex" --no-highlight --wrap=none \
  --lua-filter="${latex_dir}/omers-filter.lua" -M doctype="$doctype"
```

## Step 7: Generate preamble + assemble

Copy the template and run a single sed pipeline. Use a delimiter unlikely to
appear in values (`|`):

```bash
base_file="preamble.base.tex.template"
delta_file="preamble.${doctype}.tex.template"
for f in "$base_file" "$delta_file"; do
  if [ ! -f "${latex_dir}/${f}" ]; then
    echo "ERROR: ${latex_dir}/${f} not found."
    exit 1
  fi
done
# Assemble base + doctype delta into the working .tex, then substitute tokens.
cat "${latex_dir}/${base_file}" "${latex_dir}/${delta_file}" > "${work}/${basename}.tex"
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
  -e "s|__LIST_ITEMSEP_PT__|${LIST_ITEMSEP_PT:-2.3}|g" \
  -e "s|__PARSKIP_EM__|${PARSKIP_EM:-0.4}|g" \
  -e "s|__LINE_SPREAD__|${LINE_SPREAD}|g" \
  "$tex"

# Substitute __BODY__ from the file (sed's r command + d on the marker):
sed -i -e "/__BODY__/{r ${work}/body.tex" -e "d;}" "$tex"
```

## Step 8: Compile with optional page-count targeting

````bash
pages_target="${3:-auto}"
case "$pages_target" in 1|2|3|auto) ;; *) echo "Bad \$3 '$pages_target'; expected 1|2|3|auto"; exit 1;; esac

compile_once() {
  ( cd "$work" && xelatex -interaction=nonstopmode "${basename}.tex" >/dev/null \
                && xelatex -interaction=nonstopmode "${basename}.tex" >/dev/null )
}

regen_preamble() {
  # Re-assemble base + delta from scratch so iterative tuning starts from the
  # templates, not from an already-substituted file.
  cat "${latex_dir}/${base_file}" "${latex_dir}/${delta_file}" > "${work}/${basename}.tex"
  sed -i \
    -e "s|__FONT_SIZE_PT__|${FONT_SIZE_PT}|g" \
    -e "s|__MARGIN_TOP__|${MARGIN_TOP}|g" \
    -e "s|__MARGIN_BOTTOM__|${MARGIN_BOTTOM}|g" \
    -e "s|__MARGIN_LEFT__|${MARGIN_LEFT}|g" \
    -e "s|__MARGIN_RIGHT__|${MARGIN_RIGHT}|g" \
    -e "s|__ACCENT_R__|${ACCENT_R}|g" -e "s|__ACCENT_G__|${ACCENT_G}|g" -e "s|__ACCENT_B__|${ACCENT_B}|g" \
    -e "s|__MUTED_R__|${MUTED_R}|g"   -e "s|__MUTED_G__|${MUTED_G}|g"   -e "s|__MUTED_B__|${MUTED_B}|g" \
    -e "s|__MAIN_FONT__|${MAIN_FONT}|g" \
    -e "s|__HEADING_FONT_FALLBACK__|${HEADING_FONT_FALLBACK}|g" \
    -e "s|__HEADING_FONT__|${HEADING_FONT}|g" \
    -e "s|__SECTION_LETTERSPACE__|${SECTION_LETTERSPACE}|g" \
    -e "s|__LIST_ITEMSEP_PT__|${LIST_ITEMSEP_PT:-2.3}|g" \
    -e "s|__PARSKIP_EM__|${PARSKIP_EM:-0.4}|g" \
    -e "s|__LINE_SPREAD__|${LINE_SPREAD}|g" \
    "${work}/${basename}.tex"
  sed -i -e "/__BODY__/{r ${work}/body.tex" -e "d;}" "${work}/${basename}.tex"
}

# Corrected last-page fill: % of page height down to the lowest inked row,
# EXCLUDING the bottom ~7% footer band where the centered page number lives
# (otherwise a nearly blank page reports ~97% full and hides an orphan).
# Echoes an integer 0-100, or empty if the tooling is unavailable.
measure_fill() {
  command -v pdftoppm >/dev/null && command -v python3 >/dev/null || { echo ""; return; }
  local pg="$1"
  rm -f "${work}"/fillchk-*.png 2>/dev/null
  pdftoppm -png -r 100 -f "$pg" -l "$pg" "${work}/${basename}.pdf" "${work}/fillchk" 2>/dev/null
  local img; img=$(find "${work}" -maxdepth 1 -name 'fillchk-*.png' | head -1)
  [ -n "$img" ] || { echo ""; return; }
  python3 - "$img" <<'PY'
import sys
from PIL import Image
im = Image.open(sys.argv[1]).convert("L")
w, h = im.size
px = im.load()
footer_cut = int(h * 0.93)   # ignore bottom 7% (page-number band)
last_ink_row = 0
for y in range(footer_cut):
    for x in range(w):
        if px[x, y] < 240:
            last_ink_row = y
            break
print(last_ink_row * 100 // h)
PY
}

compile_once || { echo "xelatex compile failed; see ${work}/${basename}.log"; exit 1; }
pages=$(pdfinfo "${work}/${basename}.pdf" 2>/dev/null | awk '/^Pages:/{print $2}')
echo "Initial compile: ${pages} pages (target=${pages_target})"

# Auto mode: detect a near-empty trailing (orphan) page and re-target one fewer
# page using the SAME tuning loop below, with a 10pt font floor so text never
# shrinks into illegibility. The explicit 1|2|3 path is unchanged (FONT_FLOOR
# stays empty there, so the floor clamp is a no-op).
auto_retarget=0
FONT_FLOOR=""
if [ "$pages_target" = "auto" ] && [ -n "$pages" ] && [ "$pages" -gt 1 ]; then
  orphan_fill=$(measure_fill "$pages")
  if [ -n "$orphan_fill" ] && [ "$orphan_fill" -lt 35 ]; then
    echo "Auto: trailing page only ${orphan_fill}% full (orphan) — re-targeting $((pages - 1)) pages (font floor 10pt)."
    pages_target=$((pages - 1))
    FONT_FLOOR=10.0
    auto_retarget=1
  fi
fi

if [ "$pages_target" != "auto" ] && [ -n "$pages" ] && [ "$pages" != "$pages_target" ]; then
  for iter in 1 2 3; do
    if [ "$pages" -gt "$pages_target" ]; then
      # too many pages: tighten. FONT_FLOOR (set only for auto orphan re-target)
      # clamps the font size so it never shrinks below legibility; empty/0 = no clamp.
      FONT_SIZE_PT=$(awk -v v="$FONT_SIZE_PT" -v f="${FONT_FLOOR:-0}" 'BEGIN{nv=v-0.25; if(f>0 && nv<f) nv=f; printf "%.2f", nv}')
      LIST_ITEMSEP_PT=$(awk -v v="${LIST_ITEMSEP_PT:-2.3}" 'BEGIN{printf "%.2f", v-0.3}')
      PARSKIP_EM=$(awk -v v="${PARSKIP_EM:-0.4}" 'BEGIN{printf "%.2f", v-0.05}')
    else
      # too few pages: loosen
      FONT_SIZE_PT=$(awk -v v="$FONT_SIZE_PT" 'BEGIN{printf "%.2f", v+0.25}')
      LIST_ITEMSEP_PT=$(awk -v v="${LIST_ITEMSEP_PT:-2.3}" 'BEGIN{printf "%.2f", v+0.3}')
      PARSKIP_EM=$(awk -v v="${PARSKIP_EM:-0.4}" 'BEGIN{printf "%.2f", v+0.05}')
    fi
    echo "Iter ${iter}: tuning font_size_pt→${FONT_SIZE_PT}, list_itemsep_pt→${LIST_ITEMSEP_PT}, parskip_em→${PARSKIP_EM}"
    regen_preamble
    compile_once || { echo "xelatex compile failed during iter ${iter}; see ${work}/${basename}.log"; exit 1; }
    pages=$(pdfinfo "${work}/${basename}.pdf" 2>/dev/null | awk '/^Pages:/{print $2}')
    echo "Iter ${iter} result: ${pages} pages"
    [ "$pages" = "$pages_target" ] && break
  done
  if [ "$pages" != "$pages_target" ]; then
    if [ "$auto_retarget" = "1" ]; then
      echo "WARN: could not clear the orphan page within the 10pt font floor. Final: ${pages} pages."
      echo "      Pass an explicit page target (1|2|3) and choose which content to cut."
    else
      echo "WARN: could not hit target after 3 iterations. Final: ${pages} pages."
    fi
  fi
fi
````

On any xelatex failure, leave `$work` in place and tell the user the log path.

## Step 9: Resolve output dir and copy artifacts

```bash
# Try to detect slug from parent folder structure (buildresume layout):
parent=$(dirname "$src")
grandparent=$(dirname "$parent")
slug=$(basename "$grandparent")
apps_root=$(jq -r '.directories.applications_root' .jobops/config.json)

case "$doctype" in
  resume)       sub="resume";;
  coverletter)  sub="cover-letter";;
  "document")   sub="";;
esac

if [ "$doctype" = "document" ]; then
  # Documents always write next to source — no app-slug subfolder.
  outdir="$parent"
elif [[ "$slug" =~ ^[A-Za-z0-9]+_[A-Za-z0-9]+_[0-9]{8}$ ]] \
     && [[ "$grandparent" == "${apps_root}/"* ]]; then
  outdir="${grandparent}/${sub}"
else
  outdir="$parent"
  echo "WARN: could not resolve application slug; writing PDF next to source ($outdir)."
fi
mkdir -p "$outdir"

cp "${work}/${basename}.pdf" "${outdir}/${basename}.pdf"
cp "${work}/${basename}.tex" "${outdir}/${basename}.tex"   # advanced override artifact
```

## Step 10: Page/fill report

```bash
pages="?"; fill="unavailable"
if command -v pdfinfo >/dev/null; then
  pages=$(pdfinfo "${outdir}/${basename}.pdf" | awk '/^Pages:/{print $2}')
fi
if command -v pdftoppm >/dev/null && command -v python3 >/dev/null; then
  png="${work}/lastpage.png"
  pdftoppm -png -r 100 -f "$pages" -l "$pages" "${outdir}/${basename}.pdf" "${work}/lastpage" 2>/dev/null
  # pdftoppm appends -N for page N; find the produced file:
  produced=$(find "${work}" -maxdepth 1 -name 'lastpage-*.png' | head -1)
  if [ -n "$produced" ]; then
    fill=$(python3 - "$produced" <<'PY'
import sys
from PIL import Image
im = Image.open(sys.argv[1]).convert("L")
w, h = im.size
px = im.load()
footer_cut = int(h * 0.93)   # ignore bottom 7% (page-number band) so an orphan
                             # page doesn't report ~97% full off the page number
last_ink_row = 0
for y in range(footer_cut):
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

## Step 11: Final report

Print:

```
latex-pdf complete.
  Source:   $src
  Theme:    $THEME_NAME
  Doctype:  ${doctype}
  Output:   ${outdir}/${basename}.pdf
  TeX:      ${outdir}/${basename}.tex   (advanced-override artifact)
  Pages:    ${pages}
  Fill:     ${fill}
  Fallbacks: ${fallback_msg:-none}
```

Done. Do not delete `$work` — it's a `mktemp -d` directory the OS reclaims,
and keeping it lets the user inspect the .log on failures.
