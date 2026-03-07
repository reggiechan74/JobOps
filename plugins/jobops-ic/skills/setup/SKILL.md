---
description: Initialize JobOps Independent Contractor workspace - extends jobops config with IC directories and templates
disable-model-invocation: true
---

## 1. Prerequisite Check

Read the file `.jobops/config.json`.

If the file does not exist, stop and display:

> Run `/jobops:setup` first to initialize the base JobOps workspace.

Verify the jobops plugin is installed by checking if the `resume-summarizer` agent is available. If it is not available, stop and display:

> PREREQUISITE MISSING: jobops plugin
> Install it with: `/plugin install jobops@jobops-marketplace`

## 2. Welcome

Display the following message:

> **JobOps IC Setup**
>
> This command extends your existing JobOps workspace with independent contractor directories and templates. It will:
> - Create a client prospects directory for tracking B2B opportunities
> - Install IC-specific templates (service definitions, rate cards, proposals)
> - Update your workspace configuration with IC settings

## 3. IC Directory Configuration

Ask the user:

> Where would you like to store client prospect files?
> (default: `./Client_Prospects`)

Accept the user's input. If no input is provided, use `./Client_Prospects`.

## 4. Create Directory

Create the directory chosen in step 3 if it does not already exist.

```
mkdir -p <client_prospects_directory>
```

## 5. Template Installation

Run the IC template copy script to install independent contractor templates:

```
bash "$(cat /tmp/.jobops-ic-plugin-root)/scripts/copy-templates.sh" ".jobops/templates/default"
```

## 6. Update config.json

Read the existing `.jobops/config.json` file. Add the following entries to the configuration:

- Set `config.directories.client_prospects` to the directory path chosen by the user in step 3.
- Set `config.templates.active.service_definition_schema` to `"default"`.

Write the updated configuration back to `.jobops/config.json`.

## 7. Summary

Display the following summary:

> **IC Workspace Initialized**
>
> Added to your workspace:
> - Client prospects directory: `<client_prospects_directory>`
> - IC templates installed to `.jobops/templates/default`
> - Configuration updated in `.jobops/config.json`
>
> **Next steps:**
> - Try `/jobops-ic:defineservices` to create your service catalog
> - Try `/jobops-ic:ratecard` to generate a professional rate card
> - Try `/jobops-ic:findclient` to discover potential clients
