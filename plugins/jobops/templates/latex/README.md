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
