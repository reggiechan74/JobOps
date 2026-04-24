# JobOps Setup

JobOps v2.0 ships as two Claude Code plugins distributed through a self-hosted marketplace. Setup is a three-step flow: install Claude Code, install the plugin(s), then initialize the workspace.

## 1. Install Claude Code

Claude Code is the only hard requirement.

```bash
npm install -g @anthropic-ai/claude-code
claude --version
```

More at [claude.ai/code](https://claude.ai/code).

## 2. Install the Plugins

From inside any project directory where you want to use JobOps, launch Claude Code (`claude`) and run:

```text
/plugin marketplace add reggiechan74/JobOps
/plugin install jobops@jobops-marketplace
```

The core `jobops` plugin provides 31 skills: resume development, interview prep, OSINT, career strategy, crisis management, and application finalization.

**Optional — independent contractor toolkit:**

```text
/plugin install jobops-ic@jobops-marketplace
```

`jobops-ic` adds 10 skills for service definitions, client prospecting, pitch decks, proposals, rate cards, and landing pages. It depends on `jobops`, so install the core plugin first.

## 3. Initialize the Workspace

```text
/jobops:setup
```

This creates `.jobops/` in the current workspace:

```
.jobops/
  config.json                # Directory paths + template settings
  templates/
    default/                 # Plugin-bundled templates (read-only by convention)
    custom/                  # Your overrides
```

It also scaffolds the default output directories (`ResumeSourceFolder/`, `Job_Postings/`, `Applications/`, `Company_Intelligence/`, `Career_Analysis/`, `Crisis_Management/`). Every path is configurable — edit `.jobops/config.json` to relocate anything.

If you installed `jobops-ic`, run its setup afterward:

```text
/jobops-ic:setup
```

This extends `.jobops/config.json` with `contractor_root` (default `Contractor/`) and copies IC-specific templates.

## Configurable Output Layout

All output paths are resolved from `.jobops/config.json` — skills never hardcode directories. Defaults:

| Config key | Default | Purpose |
|------------|---------|---------|
| `resume_source` | `ResumeSourceFolder/` | Master HAM-Z career inventory (input) |
| `job_postings` | `Job_Postings/` | Target job descriptions (input) |
| `applications_root` | `Applications/` | Per-application tree with fixed `resume/`, `cover-letter/`, `assessment/`, `interview/` subfolders |
| `company_intelligence` | `Company_Intelligence/` | OSINT output, one folder per company |
| `career_analysis` | `Career_Analysis/` | Flat timestamped career-level outputs |
| `crisis_management` | `Crisis_Management/` | Flat timestamped crisis outputs |
| `contractor_root` | `Contractor/` | `jobops-ic` outputs (added by `/jobops-ic:setup`) |

See `docs/ARCHITECTURE.md` for the full contract.

## Optional Dependencies

### Pandoc — `/jobops:convert-to-word`

Needed to convert markdown to DOCX. Install via the bundled skill:

```text
/jobops:install-pandoc
```

Or manually:
- **Ubuntu/Debian**: `sudo apt-get install pandoc`
- **macOS**: `brew install pandoc`
- **Windows**: [pandoc.org/installing.html](https://pandoc.org/installing.html)

### Playwright — `/jobops:convert-to-pdf`, `/jobops:markdown-to-pdf`

PDF conversion uses Playwright for pixel-perfect rendering. If the skill reports a missing browser:

```bash
npx playwright install chromium
```

## Migrating from v1.x

If you previously used JobOps v1.x (flat `OutputResumes/` and `Client_Prospects/` layouts):

```text
/jobops:setup       # Creates .jobops/config.json and new directory tree
/jobops:migrate     # Relocates v1 outputs into the v2 app-centric layout
```

`/jobops:migrate` moves per-application artifacts into `Applications/{Company}_{Role}_{YYYYMMDD}/`, consolidates OSINT files into `Company_Intelligence/{Company}/`, and routes career/crisis/contractor outputs to their dedicated roots.

## Quick Start Checklist

- [ ] Install Claude Code: `npm install -g @anthropic-ai/claude-code`
- [ ] In Claude: `/plugin marketplace add reggiechan74/JobOps`
- [ ] In Claude: `/plugin install jobops@jobops-marketplace`
- [ ] In Claude: `/jobops:setup`
- [ ] (Optional) `/plugin install jobops-ic@jobops-marketplace` then `/jobops-ic:setup`
- [ ] (Optional) `/jobops:install-pandoc` for Word export
- [ ] (Optional) `npx playwright install chromium` for PDF export
- [ ] (v1.x users) `/jobops:migrate` to relocate existing outputs

## Troubleshooting

**`/plugin` commands not recognized**: update Claude Code (`npm i -g @anthropic-ai/claude-code@latest`) — the plugin system ships in recent versions only.

**Skills not appearing after install**: restart Claude Code. Verify the plugin is listed via `/plugin list`.

**`/jobops-ic:*` skills fail with prerequisite error**: run `/jobops:setup` before `/jobops-ic:setup`. The IC plugin extends the shared `.jobops/config.json` and will not run standalone.

**Config drift after editing `.jobops/config.json`**: skills read config at invocation time, so changes take effect immediately — no restart needed. If a skill writes to an unexpected path, confirm the key it reads is present and spelled correctly (see `docs/ARCHITECTURE.md`).

## Getting Help

- **README.md** — Full skill catalog and workflow guide
- **CLAUDE.md** — Contributor notes on repo structure
- **docs/ARCHITECTURE.md** — Config contract and output-layout rules
- **CHANGELOG.md** — Version history and v2.0 migration notes
