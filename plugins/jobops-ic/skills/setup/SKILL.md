---
description: Extend JobOps workspace with independent-contractor directories, templates, and config - requires /jobops:setup to have run first
disable-model-invocation: true
---

# JobOps IC Setup

This skill extends an existing JobOps workspace with the directories,
templates, and config entries needed by the `jobops-ic` plugin. It does not
create `.jobops/config.json` — it adds to the one created by `/jobops:setup`.

---

## Step 1: Prerequisite check

Read `.jobops/config.json`.

If the file does not exist, stop with:

> JOBOPS NOT CONFIGURED
> Run /jobops:setup first, then retry /jobops-ic:setup.

(The plugin `dependencies` field ensures `jobops` is installed alongside
`jobops-ic`; this check catches the separate case of the user skipping the
base setup.)

---

## Step 2: Welcome

Print:

> **JobOps IC Setup**
>
> This extends your existing JobOps workspace with:
> - A contractor output tree (services, prospects, proposals, pitches, rate cards, landing pages)
> - The IC service-definition template
> - Currency preference for pricing outputs

Continue.

---

## Step 3: Interview

Ask in order:

1. **Contractor root directory** — default `./Contractor`. Validate the parent is writable.
2. **Default currency** — ISO 4217 code, enum: `CAD` | `USD` | `EUR` | `GBP` | `AUD`. Default `CAD`. Used by rate cards and proposals.

---

## Step 4: Create directory tree

Run:
```bash
mkdir -p <contractor_root>/services \
         <contractor_root>/prospects \
         <contractor_root>/proposals \
         <contractor_root>/pitches \
         <contractor_root>/rate-cards \
         <contractor_root>/landing-pages
```

Report created vs exists.

---

## Step 5: Template installation

Copy the IC service-definition schema into the workspace template tree:

```bash
cp ${CLAUDE_PLUGIN_ROOT}/templates/service_definition_schema.json .jobops/templates/default/
```

If the source file is missing, stop with a clear error. Do not silently skip.

---

## Step 6: Extend `.jobops/config.json`

Read the existing config, add the following keys, and write it back
atomically (write to `.jobops/config.json.tmp`, then `mv`):

- `directories.contractor_root = "<step-3 value>"`
- `preferences.default_currency = "<step-3 value>"`
- `templates.active.service_definition_schema = "default"`

Preserve all existing keys and values. Do not change `version`.

---

## Step 7: Gitignore update

If `.gitignore` has a `# JobOps workspace` block, append (if not already
present) the `contractor_root` path on a new line inside that block.
Default: `Contractor/`. If the block is missing, print a reminder to the
user that `/jobops:setup` controls the block and that they can re-run it
to refresh.

---

## Step 8: Summary

Print:

1. Contractor root path.
2. Default currency.
3. Template installed.
4. Recommended next steps: `/jobops-ic:defineservices`, `/jobops-ic:ratecard`.

Exit.
