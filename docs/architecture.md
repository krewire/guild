# Architecture — Krewire Guild

## Module Structure

```
guild/
├── guild.go              # `//go:embed all:template` + Install library (force/dry-run, conflict detection)
├── template/
│   ├── AGENTS.md         # Unified constitution (8 kinds, full command matrix, spec conventions)
│   ├── opencode.json     # Base opencode config
│   └── .agents/
│       ├── agents/       # 15 agents: build, plan, orchestrator, scout, tester, runtime, infra, service, worker, ...
│       ├── commands/     # 16 commands: kickoff, vision, runtime, infra, service, worker, new-project, ...
│       ├── skills/       # 20 skills: vision, wasm-runtime, infra-provision, service-mesh, worker-queue, ...
│       └── context/      # vision-compact.md (5-line fast-load)
└── docs/
```

**Design decisions:**

- **Embed.FS + Install.** Template is shipped as `guild.Template` embed.FS; `guild.Install(target)` copies managed paths, refusing to overwrite unless `WithForce`.
- **Stdlib-only module.** No third-party imports; preserves `KWG-*` spec requirements.
- **Managed paths.** `AGENTS.md`, `opencode.json`, `.agents/agents|commands|skills|context`, `.agents/README.md` — installed via `krewire guild install`.


## Conventions

- Documentation in English, Markdown, spec-driven (`internal/docs/specs/guild/` in `krewire/internal`).
- Quality gates: `gofmt -l .`, `go vet ./...`, `go test ./...` in each Go repo; per-kind `krewire build` / `krewire build --plan` spot-checks.
- Cross-repo testing via temporary `replace` in `go.mod`; never `go.work`.
