---
description: DEPRECATED — wraps /jobops:latex-pdf. Will be removed in a future release.
disable-model-invocation: true
argument-hint: "<resume.md> [theme] [pages]"
---

# DEPRECATED

> This skill has been replaced by `/jobops:latex-pdf`, which uses a
> config-driven LaTeX pipeline for higher-quality output. This stub remains
> for one release so existing workflows keep working; it will be removed in
> the following release.
>
> Please migrate any saved invocations / scripts to `/jobops:latex-pdf`.

## Behavior

Print the deprecation notice above to the user, then run the `/jobops:latex-pdf`
pipeline with the same arguments:

1. Read and follow `plugins/jobops/skills/latex-pdf/SKILL.md`.
2. Pass `$1` (resume markdown), `$2` (theme), `$3` (pages) straight through.

The new skill supports the same themes (`modern`, `classic`, `minimal`, plus the
default `navy-serif`) and the same `pages: 1|2|3|auto` shape via `$3`.
