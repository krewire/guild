# .agents/

The canonical agent directory for this template — works with any AI coding agent that follows
the `.agents/` convention (OpenCode, Claude Code, Cursor, etc.).

```
.agents/
├── agents/                # Agents & subagents (Markdown + YAML frontmatter)
│   ├── build.md           # Primary orchestrator — delegates slices
│   ├── plan.md            # Vision-aware planner (internal/docs/project-vision.md)
│   ├── orchestrator.md    # Vision-aware router & parallelizer
│   ├── scout.md           # 8-kind detector (app/cli/site/book/worker/service/infra/kernel)
│   ├── reviewer.md        # Vision-aware reviewer (SSR parity, secrets, opt-in cost)
│   ├── tester.md          # Per-kind test strategy (WASM hydration, plan idempotence, DLQ)
│   ├── debugger.md
│   ├── refactor.md
│   ├── docs.md            # Centralized specs (krewire/internal)
│   ├── security.md        # Secrets/ WASM CSP / mTLS aware
│   ├── deploy.md          # Unified krewire deploy (preview envs)
│   ├── runtime.md         # WASM frontend (Go→WASM, VDOM, islands)
│   ├── infra.md           # Cloud infra (AWS/K8s, state/locking, plan/apply)
│   ├── service.md         # Microservice patterns (registry/gateway/resilience)
│   └── worker.md          # Background jobs (queues/cron/DLQ)
├── commands/              # Slash commands (Markdown + YAML frontmatter)
│   ├── kickoff.md
│   ├── new-project.md     # 8 variants (app/static/book/cli/worker/service/infra)
│   ├── vision.md          # Load unified vision (workload matrix)
│   ├── runtime.md         # /runtime — frontend/WASM
│   ├── infra.md           # /infra — plan/deploy/preview/destroy
│   ├── service.md         # /service — registry/gateway/resilience
│   ├── worker.md          # /worker — queues/cron/DLQ
│   ├── review.md
│   ├── test.md
│   ├── fix.md
│   ├── commit.md
│   ├── spec.md
│   ├── refactor.md
│   ├── deploy.md          # Unified deploy flow
│   ├── release.md
│   └── triage.md
├── skills/                # Skills, one folder each with SKILL.md + rules/
│   ├── vision/            # Shared vision loader (compact context)
│   ├── context-awareness/ # 4-layer situational intelligence (project/vision/history/env)
│   ├── agent-workflow/    # Handoffs, state, parallel vs sequential, lifecycle
│   ├── project-init/      # 8-kind scaffolding
│   ├── kickoff/
│   ├── requirement-gathering/
│   ├── spec-writing/
│   ├── test-driving/
│   ├── quality-gate/      # Vision-aware gates + WASM/infra gates
│   ├── ci-cd/             # 8-kind CI (fakes vs live backends)
│   ├── release/
│   ├── bug-triage/
│   ├── conventional-commit/
│   ├── security-review/
│   ├── dependency-audit/
│   ├── wasm-runtime/      # Go→WASM, VDOM, islands, theming (KWF-T4X9P)
│   ├── infra-provision/   # Provider/state/plan, AWS/K8s (KWF-B7N3D)
│   ├── service-mesh/      # Registry/gateway/resilience/tracing (KWF-L5H2F)
│   ├── worker-queue/      # Job queues/cron/DLQ (KWF-L5H2F)
│   └── agent-optimization/ # Context compaction & parallel orchestration
└── context/
    └── vision-compact.md  # 5-line workload matrix — fast-load for subagents
```

## File Format

All files are plain **Markdown with YAML frontmatter**:

```yaml
---
description: Human-readable description
agent: build  # for commands: which agent runs this command
mode: primary # for agents: primary | subagent
---
```

This format is portable — adapt to any tool that accepts agent definitions, commands, or skills.

## How Different Tools Load This Directory

| Tool | Agents | Commands | Skills |
|------|--------|----------|--------|
| **OpenCode** | Via `@capybearista/opencode-agents-loader` plugin | Via same plugin | Native (reads `.agents/skills/*/SKILL.md`) |
| **Claude Code** | Copy to `.claude/agents/` | Copy to `.claude/commands/` | Reference in `CLAUDE.md` or copy to `.claude/` |
| **Cursor** | Convert to `.cursor/rules/*.mdc` | Convert to `.cursor/rules/*.mdc` | Convert each skill to `.cursor/rules/<name>.mdc` |
| **Copilot** | Reference in `.github/copilot-instructions.md` | N/A | Reference in instructions |
| **Codex** | Reference in `AGENTS.md` or `CODEX.md` | N/A | Reference in instructions |
| **Windsurf** | Reference in `.windsurf/` | Reference in `.windsurf/` | Convert to `.windsurf/rules/` |
| **Generic** | Read directly as Markdown | Read directly as Markdown | Read directly as Markdown |

## Skill Rules Structure

Every skill separates its non-negotiable rules from its workflow:

```
skills/<name>/
├── SKILL.md                 # workflow / how-to (points to key-rules.md)
└── rules/
    ├── key-rules.md         # the primary, non-negotiable rules
    └── <rule-name>.md       # additional/extended rule sets
```

`SKILL.md` instructs the agent to read `rules/key-rules.md` before acting; supporting rule
files are read on demand when the agent needs them.