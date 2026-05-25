# JobOps latex-pdf LaTeX templates

This directory is plugin-bundled. `/jobops:setup` copies it to
`.jobops/templates/default/latex/`. Treat `default/` as read-only — edit
copies under `.jobops/templates/custom/latex/` and switch via
`config.templates.active.latex_config = "custom"`.

The final `.tex` is assembled as **base + doctype delta + body**.

Files:
- `preamble.base.tex.template` — shared base: documentclass, packages, palette,
  fonts, de-numbering, helper macros (`\tightlist`/`\real`/etc.), `keepaspectratio`,
  the navy section style, `\role`/`\subrole`, and table helpers (`\rowshade`/`\thcell`).
- `preamble.resume.tex.template` — resume delta (body spacing only).
- `preamble.coverletter.tex.template` — coverletter delta (parskip, requirements-table tuning).
- `preamble.document.tex.template` — generic document delta (parskip, sub/subsub headings, tables).
- `omers-filter.lua` — pandoc Lua filter that maps the JobOps markdown onto the
  OMERS LaTeX vocabulary (header block, `\section`, `\role`/`\subrole`, navy+zebra
  tables). Pass-through for the `document` doctype.
- `config.json` — themes + doctypes registry. See `../../skills/latex-pdf/SKILL.md` for schema.

The skill GENERATES the final `.tex` (base + delta + body). You do not edit
`.tex` files directly. See the SKILL.md "Markdown authoring contract" section
for the markdown shape the filter expects.
