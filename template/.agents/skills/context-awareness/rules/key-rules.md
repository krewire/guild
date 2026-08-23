# Context Awareness — Key Rules

- Detect kind among 8 (`app`/`cli`/`site`/`book`/`worker`/`service`/`infra`/`kernel`) via `krewire.yaml` + `krewire info`; never guess.
- Load `.agents/context/vision-compact.md` before full vision; 60-80% token saving is mandatory for subagents.
- Progressive disclosure: scout → vision compact → spec index → single KWF-* spec → code.
- Centralized specs are `internal/docs/specs/<project>/` (repo `krewire/internal`); per-repo `docs/specs/` are redirects — never add specs there.
- Validate freshness: `git log -1 -- <path>` before editing docs; re-scout after branch changes.
- Environment: `bin/krewire` (workspace-built) for testing local changes, not installed `krewire`; never `go.work`.
