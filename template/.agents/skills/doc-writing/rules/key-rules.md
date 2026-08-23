# Doc Writing — Key Rules

- Source of truth order: `project-vision.md` → `specs/index.md` → `core/workload.go` → `docs/architecture.md` → `README` → `manuscript` → `AGENTS.md` → `vision-compact.md`. Never reverse.
- `framework/tui` is the package; `cli` is the `project.kind` and `--cli` flag. Docs must use `framework/tui` for imports and `tui.NewApp`.
- After any workload matrix change, atomically update `project-vision.md` + `core/workload.go` + `vision-compact.md` + all `README.md` + `AGENTS.md`.
- Use `rules/doc-template.md` for the doc type; keep workload matrix to 9 rows and architecture tree to one page.
- `sync-docs` must report `In-sync` before docs PR is done.
