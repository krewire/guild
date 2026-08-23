# Vision Compact — For Agent Fast-Load

> Load this (5 lines) instead of `internal/docs/project-vision.md` + `AGENTS.md` per subagent. Full docs: `internal/docs/project-vision.md` and `internal/docs/specs/index.md`.

**Workload Matrix (9 workloads, 8 kinds):** `cli` (`framework/tui`) · `worker` (`framework/worker`) · `infra` (`framework/infra` AWS/K8s) · `site` (`framework/web/ssg`) · `book` (`mdbind`) · `frontend` (`framework/runtime` WASM) · `backend`/`app` (`framework/web`+`framework/app`) · `monolith` (modular, `KWF-5ZHQV`) · `microservice` (`framework/service`)

**Config:** `krewire.yaml` only (no `ssg.yaml`); `libs/config`+`libs/validate` typed for all kinds.

**CLI:** `krewire new/init` → `krewire build` (binary/`site/`/plan) → `krewire run/dev/worker/deploy/dashboard/generate` → `krewire test/info/version`; `krewire build --plan` for infra, `krewire deploy --preview` for PR envs.

**Roadmap:** Phase 0 vision/specs ✅ → Phase 1 WASM runtime (`KWF-T4X9P`) → Phase 2 infra (`KWF-B7N3D`) → Phase 3 service/worker (`KWF-L5H2F`) → Phase 4 DX → Phase 5 polish

**Routing:** `runtime` for WASM/islands, `infra` for provider/state, `service` for registry/gateway/resilience, `worker` for queues, `build` orchestrates cross-cutting.

**Spec Index:** `internal/docs/specs/index.md` (49 specs) — `🔜 Planned` = spec-first before code; `✅ Shipped` = implement directly.
