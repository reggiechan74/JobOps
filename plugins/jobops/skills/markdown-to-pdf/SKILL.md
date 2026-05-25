---
description: DEPRECATED — wraps /jobops:latex-pdf. Will be removed in a future release.
disable-model-invocation: true
argument-hint: "<file|glob|dir> [output_dir_ignored]"
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
pipeline for `$1`:

1. Read and follow `plugins/jobops/skills/latex-pdf/SKILL.md`.
2. Pass `$1` through to that skill as its `$1` (file, glob, or directory).
3. `$2` of this old skill was a custom output directory; the new skill's
   output dir is determined by doctype + Application Path Resolution and does
   not accept an override. If `$2` was set, print:
   `WARN: /jobops:markdown-to-pdf $2 (output_dir) is no longer honored. latex-pdf writes per-doctype next to source or into the application folder.`

That's it. The new skill autodetects doctype per file (defaults to `document`
for non-resume / non-coverletter content, exactly matching the documents this
old skill handled).
