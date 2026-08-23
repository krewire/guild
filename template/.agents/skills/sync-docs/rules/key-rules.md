# Sync Docs — Key Rules

- Source of truth order: `project-vision.md` → `specs/index.md` → `core/workload.go` → `docs/architecture.md` → `README`/`manuscript` → `AGENTS.md` → `vision-compact.md`. Never update upstream to match downstream.
- `framework/tui` is the package; `cli` is the `project.kind` and `--cli` flag. Docs must use `framework/tui` for imports and `tui.NewApp`, not `framework/cli`/`cli.NewApp`.
- Original `<project>/docs/specs/` must stay as `MOVED.md` + redirect; real specs are `internal/docs/specs/<project>/`.
- After any workload matrix change, atomically update `project-vision.md` + `core/workload.go` + `vision-compact.md` + all `README.md` + `AGENTS.md`.
- Manuscript `01-introduction.md` must describe unified 9-workload vision, not old meta-framework phrasing.
