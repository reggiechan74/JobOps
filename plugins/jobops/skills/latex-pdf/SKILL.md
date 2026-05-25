---
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

Templates referenced: `config.json`, `preamble.resume.tex.template`, `preamble.coverletter.tex.template`.

If `config.templates.active.latex_config` is unset (older config), fall back to
`${CLAUDE_PLUGIN_ROOT}/templates/latex/` and warn the user to re-run
`/jobops:setup` so their workspace owns a copy they can customize.

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
pandoc "$src" -t latex -o "${work}/body.tex" --no-highlight --wrap=none
```

## Step 7: Generate preamble + assemble

Copy the template and run a single sed pipeline. Use a delimiter unlikely to
appear in values (`|`):

```bash
preamble_file="preamble.${doctype}.tex.template"
if [ ! -f "${latex_dir}/${preamble_file}" ]; then
  echo "ERROR: ${latex_dir}/${preamble_file} not found."
  exit 1
fi
cp "${latex_dir}/${preamble_file}" "${work}/${basename}.tex"
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
  # Re-run the sed pipeline against a fresh copy of the template so iterative
  # tuning starts from the template, not from an already-substituted file.
  cp "${latex_dir}/${preamble_file}" "${work}/${basename}.tex"
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

compile_once || { echo "xelatex compile failed; see ${work}/${basename}.log"; exit 1; }
pages=$(pdfinfo "${work}/${basename}.pdf" 2>/dev/null | awk '/^Pages:/{print $2}')
echo "Initial compile: ${pages} pages (target=${pages_target})"

if [ "$pages_target" != "auto" ] && [ -n "$pages" ] && [ "$pages" != "$pages_target" ]; then
  for iter in 1 2 3; do
    if [ "$pages" -gt "$pages_target" ]; then
      # too many pages: tighten
      FONT_SIZE_PT=$(awk -v v="$FONT_SIZE_PT" 'BEGIN{printf "%.2f", v-0.25}')
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
    echo "WARN: could not hit target after 3 iterations. Final: ${pages} pages."
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
