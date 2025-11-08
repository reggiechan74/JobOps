# Repository Guidelines

## Project Structure & Module Organization
- `README.md` and `SETUP.md` capture the HAM-Z™ workflow and install steps; keep them current alongside feature work.
- `.claude/agents/` hosts agent briefs and `.claude/commands/` defines slash commands; update both whenever behaviour changes.
- Inputs live in `ResumeSourceFolder/` and `Job_Postings/`; generated artefacts belong in `OutputResumes/` and `Briefing_Notes/`. Retain `.gitkeep` sentinels.
- Reference collections (`Sample_Output/`, `Scoring_Rubrics/`, `Intelligence_Reports/`, `SourceMaterial/`) show accepted formatting—mirror their naming when sharing exemplars.

## Build, Test, and Development Commands
- `npm run install-all` installs npm dependencies plus Playwright Chrome; use `npm install` or `npm run install-browsers` if you only need packages or browsers.
- `npx playwright test` runs the suite (place specs in `tests/`); add `--project=chromium` for quick smoke checks.
- `claude --workspace .` spins up the Claude Code CLI so you can verify command registration before review.

## Coding Style & Naming Conventions
- Prefer TypeScript modules under `src/`, using 2-space indentation and descriptive `verbNoun` command handlers or `nounAgent` orchestrators.
- Name job inputs `Job_Postings/Company_Role_YYYY-MM-DD.md` and outputs `OutputResumes/Step{N}_Final_Resume_Role_Company_Date.md` to keep automation parsable.
- Markdown briefs begin with an H1, use sentence-case headings, and wrap shell snippets in fenced code blocks.

## Testing Guidelines
- Cover new flows with `@playwright/test`; organise files as `tests/<feature>.spec.ts` and assert on Markdown content plus filesystem side effects.
- Reuse fixtures in `Sample_Output/` when checking formatting, updating those exemplars if expectations shift.
- Record manual verification in `Briefing_Notes/` only when it must persist; otherwise discard scratch artefacts locally.

## Commit & Pull Request Guidelines
- Follow the existing imperative commit tone (`Add job search workflow support`) and keep subjects under ~72 characters, noting key directories in the body when needed.
- Pull requests should include a succinct change summary, affected assets, and validation evidence (`npx playwright test`, screenshots, or sample diffs).
- Call out edits to `.claude/agents` or `.claude/commands` so reviewers rerun command registration.

## Agent & Data Handling Tips
- Add new slash commands alongside their agent brief and refresh the Core Commands table in `README.md`.
- Treat customer materials as sensitive: scrub PII before committing and keep proprietary data in ignored folders when possible.
- Confirm every published output traces back to sources in `ResumeSourceFolder/` to honour the provenance standard.
