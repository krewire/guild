# AGENTS.md — Krewire Guild

Agent guide for AI agents working in the `github.com/krewire/guild` repository.

## Repository-Specific Rules

- The installable template lives under `template/` and is embedded via
  `//go:embed all:template` in `guild.go`. It is the source of truth that
  `kiw guild install` copies into target projects.
- Do not mistake `template/AGENTS.md` for this file: that copy governs projects
  that install the template. This root `AGENTS.md` governs this repository.
- Keep the module stdlib-only (no third-party imports) and preserve the GLD-IN
  requirements from `docs/specs/KWG-INSTALL-P9ZT4-guild-module-install.md`.
- Quality gates: `cd <repo> && gofmt -l . && go vet ./... && go test ./...`.

## Sorting (applies to EVERYTHING)

Sort ANY decision, backlog, task list, spec list, migration plan, release
plan, or todo by **impact-to-effort** (high impact, low effort first) first,
then by **dependency chain** (foundations before dependents). State the chosen
order up front so the user can correct it before work begins.

## Specifications (convention)

- Location: `docs/specs/`, with `index.md` as the index.
- ProjectId: `KWG-*` (this repo / guild). File name:
  `{ProjectId}-{Scope}-{SpecID}-{slug}.md`, e.g.
  `KWG-INSTALL-P9ZT4-guild-module-install.md`. Scope is a short category
  written in uppercase alphanumeric characters (`INSTALL`, `MODULE`,
  `SCAFFOLD`, ...) chosen in this repo.
- **Spec-per-initiative**: each initiative (feature, milestone, or cohesive
  change) gets its own spec file — never one spec for several unrelated
  initiatives. Scope creep inside a spec is a sign to split it.
- **SpecID is a unique random 5-character alphanumeric code** (e.g. `P9ZT4`).
  Pick it randomly; do not reuse, do not sequence. It lives in the metadata
  table and in the file name.
- Mandatory metadata table: `SpecID`, `Title`, `Status`, `Date`, `Author`,
  `Domain`. **No `Version`, `Last updated`, or `Changes` fields.** Revision
  history is tracked by `git log -- <path>`, not in-file metadata.
- **Write the spec before implementation.** Every new feature has a
  requirement row with an ID (e.g. `GLD-IN-001`).

## Language

- **Use English in all agent output** and repository docs — responses, reports,
  summaries, commit messages, logs.
- The human may write to you in Indonesian (or any language); always reply in English.

## Mission

Help developers get work done correctly, quickly, and safely — without breaking
what already works. When facing an unfamiliar project, never assume: map it out
before acting.

## Core Principles

1. **Read before writing.** Understand the structure, conventions, and surrounding code before changing anything.
2. **Smallest change.** Change as little as possible to fulfill the request. Do not "refactor along the way".
3. **Respect existing conventions.** Follow the patterns, style, and tooling the project already uses — not personal preferences.
4. **Verify, do not guess.** If unsure about the environment/tool/library, check first (manifest, docs, or query tool).
5. **Never break things.** No feature is worth more than working code.

## Workflow

1. **Understand** the request. Ask if ambiguous — better to ask than to go the wrong way.
2. **Plan** — for non-trivial tasks, make a short plan.
3. **Implement** — make the smallest change following conventions.
4. **Verify** — run the relevant tests/lint/build. Never finish a task without verification.
5. **Summarize** — report what changed, why, and how it was verified.

## Edit Discipline

- Use precise edit tools; avoid replacing whole files unnecessarily.
- Do not add comments unless asked — clear code beats comments.
- Do not change code unrelated to the task.
- Be careful with destructive operations (`rm`, force-push, overwrite) — ask for confirmation when unsure.

## Testing

- Run the module tests (`go test ./...`) after changing `guild.go` or its tests.
- When the template content changes, existing tests already assert the key
  installed paths (`template/AGENTS.md`, `opencode.json`, `.agents/...`).

## Commit

Use **Conventional Commits**:

```
<type>(<scope>): <short description>

feat:     new feature
fix:      bug fix
refactor: change without behavior change
docs:     documentation
test:     tests
chore:    other supporting work
```

Commit focuses on one logical change. Do not commit unrelated files or secrets.

## Documentation

- Update README/docs when public behavior changes.
- For complex features/design, write or update a spec in `docs/specs/` before implementing.
- Follow the spec template at `docs/specs/spec-template.md` and the
  conventions above (file naming, SpecID, metadata table).