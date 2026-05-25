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
```

Stop on missing deps; do NOT proceed.

## Step 2: Detect doctype

Determine the doctype in this order (first match wins):

1. `--doctype=<value>` flag on the command line.
2. YAML front-matter `output_type` field in `$1` (read the front matter; the file's first 30 lines suffice):
   - `output_type ∈ {resume_final, resume_step1, resume_step2, resume_provenance}` → doctype `resume`
   - `output_type ∈ {cover_letter, coverletter}` → doctype `coverletter`
3. Filename heuristic: if `basename "$1"` contains the substring `cover_letter` → `coverletter`, else → `resume`.

```bash
# Pull --doctype= flag if present in any of the positional args
doctype=""
for a in "$@"; do
  case "$a" in --doctype=*) doctype="${a#--doctype=}";; esac
done

if [ -z "$doctype" ]; then
  # Try YAML front-matter
  fm_type=$(awk '/^---$/{n++;next} n==1 && /^output_type:/{print $2; exit}' "$1" | tr -d '"' || true)
  case "$fm_type" in
    resume_final|resume_step1|resume_step2|resume_provenance) doctype="resume";;
    cover_letter|coverletter)                                  doctype="coverletter";;
  esac
fi

if [ -z "$doctype" ]; then
  case "$(basename "$1")" in
    *cover_letter*) doctype="coverletter";;
    *)              doctype="resume";;
  esac
fi

case "$doctype" in
  resume|coverletter) ;;
  *) echo "Doctype '$doctype' not supported (Phase 2: resume, coverletter)"; exit 1;;
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

Doctype is already resolved (Step 2). Determine theme: `$2` override → `doctype.theme` → `default_theme`. If the named theme is not in `themes`, stop with the "available themes" message.

Pull these resolved values into shell variables (use `jq -r`):

```
THEME_NAME, FONT_SIZE_PT, LINE_SPREAD,
MAIN_FONT, HEADING_FONT, HEADING_FONT_FALLBACK,
ACCENT_R, ACCENT_G, ACCENT_B,
MUTED_R, MUTED_G, MUTED_B,
MARGIN_TOP, MARGIN_BOTTOM, MARGIN_LEFT, MARGIN_RIGHT,
SECTION_LETTERSPACE
```

Then, doctype-specific:
- If `doctype == "resume"`: also pull `LIST_ITEMSEP_PT` from `.doctypes.resume.list_itemsep_pt`.
- If `doctype == "coverletter"`: also pull `PARSKIP_EM` from `.doctypes.coverletter.parskip_em` and `SIGNATURE_IMAGE` from `.doctypes.coverletter.signature_image` (may be `null`).

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
src="$1"
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

## Step 8: Compile (2 passes)

```bash
( cd "$work" && xelatex -interaction=nonstopmode "${basename}.tex" >/dev/null \
              && xelatex -interaction=nonstopmode "${basename}.tex" >/dev/null ) \
  || { echo "xelatex compile failed; see ${work}/${basename}.log"; exit 1; }
```

On failure, leave `$work` in place and tell the user the log path.

## Step 9: Resolve output dir and copy artifacts

```bash
# Try to detect slug from parent folder structure (buildresume layout):
parent=$(dirname "$src")
grandparent=$(dirname "$parent")
slug=$(basename "$grandparent")
apps_root=$(jq -r '.directories.applications_root' .jobops/config.json)

case "$doctype" in
  resume)      sub="resume";;
  coverletter) sub="cover-letter";;
esac

if [[ "$slug" =~ ^[A-Za-z0-9]+_[A-Za-z0-9]+_[0-9]{8}$ ]] \
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
