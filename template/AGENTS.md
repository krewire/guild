# Krewire Guild — Agent Constitution

This document is the primary operating ruleset for the AI agent in projects
that install this template. It is tuned for the **Krewire ecosystem** (Go
projects driven by the `krewire` CLI). If the current project is not a Krewire
project, follow the generic principles below and skip the Krewire-specific
sections marked *Krewire*.

## Language

- **Use English in all agent output** — responses, reports, summaries, commit messages, logs.
- The human may write to you in Indonesian (or any language); always reply in English.
- Code, identifiers, and documentation follow the project's own conventions.

## Mission

Help developers get work done correctly, quickly, and safely — without breaking
what already works. When facing an unfamiliar project, never assume: map it out
before acting.

## Identify the Project Kind (Krewire)

The Krewire ecosystem has exactly one config file, `krewire.yaml`, and eight
project shapes covering the full web-service spectrum. Detect which one you are
in before doing anything:

| Kind      | Detection                                                   | Driven by                              |
| --------- | ----------------------------------------------------------- | -------------------------------------- |
| `app`     | `project.kind: app` or a root `main.go` (fullstack monolith)| `krewire run`, `krewire dev`             |
| `cli`     | `project.kind: cli`                                         | `krewire run <args...>`                 |
| `site`    | `project.kind: site` or an `ssg:` key in `krewire.yaml`     | `krewire build`, `krewire serve`         |
| `book`    | `project.kind: book` or a `manuscript/` directory            | `krewire build`, `krewire serve`         |
| `worker`  | `project.kind: worker` (job queues, cron, retries)          | `krewire run`, `krewire worker`, `krewire dev` |
| `service` | `project.kind: service` (microservice)                      | `krewire run`, `krewire dev`             |
| `infra`   | `project.kind: infra` (cloud IaC)                           | `krewire deploy`, `krewire build --plan` |
| kernel    | only `go.mod`, `krewire.yaml`, `main.go`, `.gitignore` (pre-`init`) | `krewire init` |

Validate with `krewire info` (prints the detected kind). Never guess: read
`krewire.yaml` and check for `manuscript/` before deciding.
See the unified vision: [`KWF-M8K2Q`](framework/docs/specs/KWF-ARCH-M8K2Q-unified-framework-vision.md).

## The `krewire` Command Matrix (Krewire)

`krewire` is the single entry point for the whole ecosystem. Never invoke
project-specific `cmd/` binaries for build/serve/run.

| Command | Purpose | Works on |
| ------- | ------- | -------- |
| `krewire new <name>` | Scaffold a minimal kernel (go.mod, krewire.yaml, main.go, .gitignore) | any new project |
| `krewire init` | Equip a kernel in place (default: fullstack app) | kernel |
| `krewire init --static` | Equip a declarative static site (`ssg:` in krewire.yaml) | kernel |
| `krewire init --book` | Equip a manuscript book (mdbind) | kernel |
| `krewire init --cli` | Equip a command-line application (framework/tui) | kernel |
| `krewire init --template <git-url>` | Clone a starter repository | empty dir |
| `krewire build` | Build the project (binary for app/cli/worker/service, `site/` for site/book, plan for infra) | all |
| `krewire serve` | Serve the built site over HTTP | site, book |
| `krewire run [args...]` | Build and run the app/CLI/worker/service binary | app, cli, worker, service |
| `krewire dev` | Rebuild + auto-restart on change (incl. WASM for frontend) | app, cli, worker, service |
| `krewire worker` | Run background workers / job queues | worker |
| `krewire deploy` | Provision infra + deploy (`--plan`, `--preview`, `--destroy`) | app, site, book, worker, service, infra |
| `krewire dashboard` | Local dev dashboard (services, logs, traces, infra) | worker, service, infra |
| `krewire generate` | Generate code (OpenAPI, config, etc.) | all |
| `krewire test` | Run `go test ./...` of the current module | all |
| `krewire info` | Print environment and detected project kind | all |
| `krewire version` | Print CLI and framework versions | all |
| `krewire guild install` | Install this template | any project |

## Core Conventions (Krewire)

- **Config lives in `krewire.yaml` only.** There is no `ssg.yaml`; a declarative
  site uses the `ssg:` key inside `krewire.yaml`.
- **Entry point** — `app` and `cli` projects build from the root `main.go`
  (`go build .`). `site`/`book` projects have **no entry point**.
- **Canonical layout (app)** — `main.go` (thin entry), `internal/` (app
  assembly, config, http), `web/` (layouts, pages, theme), `assets/` or
  `public/` (embedded static assets).
- **CLI layout** — `main.go` (tui.App harness) + `internal/commands/`.
- **Site layout** — a `ssg:` key with `layouts`, `components`, `pages`; output
  to `site/`; links normalized with trailing slashes (dir-based routing).
- **Book layout** — `manuscript/` markdown chapters, assembled by mdbind.
- **Worker layout** — `worker:` key in `krewire.yaml`, jobs under `internal/worker/`, queues via `framework/worker`.
- **Service layout** — `service:` key, registry/config/gateway/resilience via `framework/service`; modular monolith default, opt-in extraction.
- **Infra layout** — `infra:` key, provider-agnostic declarations under `infra/` compiled by `framework/infra` to AWS/Kubernetes.
- **Config struct** — typed `krewire.yaml` structs loaded with `libs/config`,
  validated with `libs/validate` (`validate:"required"` tags) and business rules in `libs/core` (`Kind`/`Workload`/`SpecID`).
- **Exit codes** — `0` success, `1` runtime failure, `2` usage error
  (`libs/core.ExitCodeSuccess/Failure/Usage`).
- **Control plane** — `libs/core` (declarative: business rules, workload registry) + `libs/kern` (imperative: `Kernel`/`Module`/`Registry`/`Executor`/`Supervisor`) are the ecosystem center; `framework` and `krewire` compose via `kern`.
- **Modules** — `github.com/krewire/framework` (unified framework: `tui`, `web`+`ssg`, `ui`, `app`, `runtime`, `worker`, `service`, `infra`), `github.com/krewire/libs` (`core`+`kern`+`config`/`validate`/`term`), `github.com/krewire/mdbind`,
  `github.com/krewire/guild`. Cross-repo testing uses temporary `replace`
  directives in `go.mod` — never `go.work` with committed references.

## Spec-Driven Development (Krewire)

- **Spec first.** Before any code, write a specification in `docs/specs/` of
  the owning repo. New features get a requirement row with an ID
  (e.g. `RND-BLD-001`, `FRK-SSG-010`).
- **File name** `{ProjectId}-{Scope}-{SpecID}-{slug}.md`; SpecID is a unique
  random 5-character code (do not reuse or sequence).
- **Metadata table** — `SpecID`, `Title`, `Status`, `Date`, `Author`, `Domain`.
  No `Version`, `Last updated`, or `Changes` fields.
- **Revision history lives in git**, not in the file: verify changes with
  `git log -1 -- <path>`.
- Implementation software repos follow this order: framework → mdbind → krewire
  → docs/landing, then `gofmt`/`go vet`/`go test`, then `push`+`tag` and
  propagate the version to downstream `go.mod`.

## Workflow

1. **Understand** the request. Ask if ambiguous — better to ask than to go the wrong way.
2. **Plan** — for non-trivial tasks, make a short plan (this may use the `plan` agent or `/spec`).
3. **Implement** — make the smallest change following conventions.
4. **Verify** — run the relevant tests/lint/build. Never finish a task without verification.
5. **Summarize** — report what changed, why, and how it was verified.

## Mapping the Project

Run `/kickoff` (or `krewire info`) when the project is unmapped. Record: detected
kind, the `krewire` commands it needs, layout, and conventions. Don't make
another agent re-map from scratch.

## Edit Discipline

- Use precise edit tools; avoid replacing whole files unnecessarily.
- Do not add comments unless asked — clear code beats comments.
- Do not change code unrelated to the task.
- Be careful with destructive operations (`rm`, force-push, overwrite) — ask for confirmation when unsure.

## Quality Gates

- Krewire Go projects: `gofmt -l .`, `go vet ./...`, `go test ./...` in each repo.
- Per-kind build gate — `app`/`cli`: `go build .`; `site`/`book`:
  `krewire build` then spot-check the `site/` output.
- Run the real commands and record results; do not claim a gate passes without evidence.

## Testing

- `krewire test` (Go: `go test ./...`). Find repo-local test patterns first.
- After changing behavior, add/update tests when the project's patterns expect it.
- If tests fail, use the `debugger` agent for root-cause analysis instead of trying random fixes.

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
Before committing, check `git status`/`git diff` and never commit secrets.
**Do not commit large changes in a single commit** — split large changes into multiple small, atomic commits, each with a clear scope and message. Large commits make review difficult and increase risk of conflicts.

## Documentation

- Update README/docs when public behavior changes.
- For complex features/design, write or update a spec in `docs/specs/` before implementing.
- Follow the ecosystem spec conventions above.

## Project-Specific Customization

Add rules specific to this project below here: agreed architecture, concrete
naming conventions, specific test commands, compatibility targets, or the
`krewire` commands this project needs. Keep them minimal and verifiable.