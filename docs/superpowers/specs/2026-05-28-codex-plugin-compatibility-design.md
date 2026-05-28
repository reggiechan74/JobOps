# JobOps Codex Plugin Compatibility

**Date:** 2026-05-28
**Status:** Approved (brainstorm phase)
**Branch:** `main`
**Next step:** Implementation plan via `superpowers:writing-plans`.

---

## 1. Purpose

JobOps currently ships as Claude Code plugins only. The immediate goal is to make the same `jobops` and `jobops-ic` plugin trees installable and discoverable in Codex without creating a separate Codex copy of every skill.

The compatibility target for this pass is:

1. Codex can discover the repository marketplace.
2. Codex can install both plugins from that marketplace.
3. Codex can scan all bundled skills.
4. Documentation explains both Claude Code and Codex installation paths.
5. The existing Claude Code install path remains unchanged.

Deep runtime translation is intentionally out of scope except for small, low-churn notes that prevent obvious first-use confusion.

## 2. Platform Findings

Verified against OpenAI Codex plugin and skill documentation on 2026-05-28:

- Codex repo marketplaces live at `.agents/plugins/marketplace.json`.
- Codex can also read legacy-compatible `.claude-plugin/marketplace.json`, but `.agents/plugins/marketplace.json` is the current repo marketplace location.
- Every Codex plugin requires `.codex-plugin/plugin.json` at the plugin root.
- Plugin manifest paths, including `skills`, are relative to the plugin root and should start with `./`.
- Codex skill folders contain `SKILL.md`; each `SKILL.md` must include `name` and `description` frontmatter.
- Optional `agents/openai.yaml` can set Codex UI metadata and `policy.allow_implicit_invocation`.

Primary references:

- `https://developers.openai.com/codex/plugins/build`
- `https://developers.openai.com/codex/skills`
- `https://developers.openai.com/codex/subagents`

## 3. Design Decisions

| Decision | Choice |
|---|---|
| Distribution model | One shared plugin tree with parallel Claude and Codex manifests |
| Codex marketplace path | `.agents/plugins/marketplace.json` |
| Claude marketplace path | Keep existing `.claude-plugin/marketplace.json` |
| Codex plugin manifests | Add `.codex-plugin/plugin.json` under each existing plugin |
| Skill source | Existing `plugins/<plugin>/skills/` directories |
| Skill frontmatter | Add explicit `name: <folder-name>` to every `SKILL.md` |
| Runtime rewrite scope | Minimal; avoid duplicated skill bodies |
| Agent workflow compatibility | Document Claude-to-Codex mapping; defer broad rewrites |
| Setup/root-path issue | Do not bulk-rewrite `${CLAUDE_PLUGIN_ROOT}` in this pass; flag for targeted future cleanup |
| Validation | Add static package contract checks plus CLI smoke checks |

## 4. Repository Structure

The new Codex-facing files are additive:

```text
JobOps/
  .agents/
    plugins/
      marketplace.json
  .claude-plugin/
    marketplace.json
  plugins/
    jobops/
      .claude-plugin/plugin.json
      .codex-plugin/plugin.json
      skills/
      agents/
      templates/
    jobops-ic/
      .claude-plugin/plugin.json
      .codex-plugin/plugin.json
      skills/
      agents/
      templates/
```

No separate `plugins-codex/` or generated mirror tree is introduced.

## 5. Codex Marketplace

Add `.agents/plugins/marketplace.json` with the marketplace name `jobops-marketplace`. It lists both plugins in install order:

1. `jobops`
2. `jobops-ic`

Each entry uses the documented local source object:

```json
{
  "name": "jobops",
  "source": {
    "source": "local",
    "path": "./plugins/jobops"
  },
  "policy": {
    "installation": "AVAILABLE",
    "authentication": "ON_INSTALL"
  },
  "category": "Productivity"
}
```

`jobops-ic` uses the same shape with `./plugins/jobops-ic`.

## 6. Codex Plugin Manifests

Add `plugins/jobops/.codex-plugin/plugin.json` and `plugins/jobops-ic/.codex-plugin/plugin.json`.

Each Codex manifest mirrors the current Claude manifest identity and version, then adds Codex install-surface metadata:

- `skills: "./skills/"`
- `interface.displayName`
- `interface.shortDescription`
- `interface.longDescription`
- `interface.developerName`
- `interface.category`
- `interface.capabilities`
- `interface.websiteURL`

The manifests do not declare hooks, apps, MCP servers, or dependencies unless the corresponding Codex-supported component files exist. `jobops-ic` remains documented as requiring `jobops`; dependency enforcement is handled by docs and setup checks unless Codex adds a supported dependency field for this shape.

## 7. Skill Frontmatter

Every existing skill keeps its body and Claude metadata. Add only the Codex-required `name` key:

```yaml
---
name: buildresume
description: Build a complete resume through all 3 steps (draft, provenance check, final)
disable-model-invocation: true
---
```

The name must equal the containing folder name. This is behavior-neutral for Claude Code and required for Codex skill discovery.

## 8. Lean Dual-Compatibility Notes

Do not duplicate all skill instructions for Codex. Add a short contributor-facing compatibility section to `docs/ARCHITECTURE.md` and user docs:

- Claude invocation examples use `/jobops:skill`.
- Codex explicit skill invocation should use the Codex skill/plugin invocation surface.
- Claude `Task tool` subagent instructions map conceptually to Codex subagent spawning.
- Claude-specific `${CLAUDE_PLUGIN_ROOT}` commands remain known follow-up work; future edits should prefer a small platform-neutral helper pattern rather than copying whole skills.

This keeps runtime compatibility work targeted and prevents the plugin from growing two parallel instruction sets.

## 9. Documentation

Update:

- `README.md`: describe JobOps as Claude Code and Codex compatible, add Codex quick-start, and update contributor structure.
- `SETUP.md`: add Codex install path and troubleshooting.
- `CLAUDE.md`: contributor notes mention both manifest families.
- `docs/ARCHITECTURE.md`: add the platform compatibility contract.
- `CHANGELOG.md`: add a new entry for Codex install/discovery compatibility.

## 10. Validation

Add a static package contract test under `scripts/validate/` or another existing validation location. It checks:

- `.agents/plugins/marketplace.json` exists and has valid JSON.
- Both `.codex-plugin/plugin.json` files exist and have valid JSON.
- Codex marketplace paths are `./plugins/jobops` and `./plugins/jobops-ic`.
- Codex manifests point `skills` to `./skills/`.
- Package, Claude manifests, Codex manifests, and marketplace metadata use the same version.
- Every `plugins/*/skills/*/SKILL.md` has `name` and `description` frontmatter.
- No Codex manifest references nonexistent `apps`, `mcpServers`, or `hooks`.

Run:

```bash
npm test --if-present
node scripts/validate/validate-codex-plugin-compatibility.js
python3 -m json.tool .agents/plugins/marketplace.json >/dev/null
python3 -m json.tool plugins/jobops/.codex-plugin/plugin.json >/dev/null
python3 -m json.tool plugins/jobops-ic/.codex-plugin/plugin.json >/dev/null
codex plugin marketplace add ./
codex plugin list
```

If the local CLI supports direct add from the configured marketplace, also run:

```bash
codex plugin add jobops@jobops-marketplace
codex plugin add jobops-ic@jobops-marketplace
```

Claude validation remains:

```bash
claude plugin validate plugins/jobops
claude plugin validate plugins/jobops-ic
claude plugin validate .
```

Run Claude validation only when `claude` is available.

## 11. Out of Scope

- Duplicating the plugin tree for Codex.
- Rewriting all Claude `Task tool` workflows in this pass.
- Creating Codex custom agent TOML files for all existing Claude agents.
- Replacing every `/jobops:*` reference inside generated examples.
- Solving every `${CLAUDE_PLUGIN_ROOT}` occurrence before Codex discovery works.
- Changing JobOps workspace config semantics.

## 12. Risks and Mitigations

**Risk:** Codex installs the plugin but some agent-heavy workflows still contain Claude-specific operational language.

**Mitigation:** Document the platform mapping now and prioritize targeted rewrites only when a workflow is proven broken in Codex.

**Risk:** Manifest versions drift between Claude and Codex files.

**Mitigation:** Static validation fails on drift.

**Risk:** Codex implicit skill invocation causes a sensitive career/legal/crisis skill to trigger unexpectedly.

**Mitigation:** Preserve explicit-invocation intent. Add Codex `agents/openai.yaml` policy files later if empirical testing shows implicit invocation is too aggressive; do not add 45 metadata files prematurely.

**Risk:** `jobops-ic` dependency enforcement differs by platform.

**Mitigation:** Keep `jobops-ic:setup` runtime checks, document `jobops` as prerequisite, and install-order the marketplace entries.
