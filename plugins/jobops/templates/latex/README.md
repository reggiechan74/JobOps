# JobOps latex-pdf LaTeX templates

This directory is plugin-bundled. `/jobops:setup` copies it to
`.jobops/templates/default/latex/`. Treat `default/` as read-only — edit
copies under `.jobops/templates/custom/latex/` and switch via
`config.templates.active.latex_config = "custom"`.

Files:
- `preamble.resume.tex.template` — resume preamble with `__TOKEN__` substitutions.
- `preamble.coverletter.tex.template` — coverletter preamble (wider margins, longtable styling, signature support).
- `preamble.document.tex.template` — generic document preamble (rubrics, briefings, analyses).
- `config.json` — themes + doctypes registry. See `../../skills/latex-pdf/SKILL.md` for schema.

The skill GENERATES the final `.tex`. You do not edit `.tex` files directly.
