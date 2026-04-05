# Implementation Plan: Bird lookup via service and REST

**Branch**: `001-bird-rest-endpoint` | **Date**: 2026-04-04 | **Spec**: [spec.md](./spec.md)  
**Input**: Feature specification from `specs/001-bird-rest-endpoint/spec.md`

**Note**: This plan is produced by `/speckit.plan`. Execution follows Spec Kit phases below.

## Summary

Add a **bird label resolution** capability: a small **service** maps one selection string to either `American Bald Eagle` (only when the value is exactly `Eagle`) or `Chickadee` otherwise. A new **REST** handler exposes the same behavior using **GET**, a **single query parameter**, and a **plain-text** body per **FR-006**. Implementation stays in the existing Go HTTP server in `main.go`, with mapping logic in a dedicated package so the handler stays a thin delegate (**FR-005**). Downstream **`tasks.md`** MUST split work so each task can be committed and opened as its **own PR** per the project constitution.

## Technical Context

**Language/Version**: Go 1.23.5 (per `go.mod`)  
**Primary Dependencies**: `net/http` (stdlib), existing `github.com/joho/godotenv`, `github.com/prometheus/client_golang/prometheus/promhttp` (unchanged for this feature)  
**Storage**: N/A (stateless mapping)  
**Testing**: `go test ./...` — no `_test.go` files exist today; add tests with the bird package/handler  
**Target Platform**: Linux/macOS (local dev); container/K8s as already deployed for this repo  
**Project Type**: single-module web service (HTTP API)  
**Performance Goals**: No new targets; in-process string compare is trivial  
**Constraints**: Match **FR-006** (`text/plain; charset=utf-8`, body is label only, optional single trailing newline); case-sensitive match to `Eagle`; no trimming in v1  
**Scale/Scope**: One endpoint, two outcomes, no persistence

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Evidence |
|-----------|--------|----------|
| I. Single-task changes | **Pass** | Plan defers work breakdown to **`tasks.md`** (`/speckit.tasks`): implementers take **one** task per change set (e.g. bird package first, then HTTP wiring, then tests) — not one giant PR. |
| II. Commit when task done | **Pass** | Plan assumes each task ends with a **commit** referencing that task. |
| III. PR per completed task | **Pass** | Plan assumes **one PR per task** after each commit. |
| IV. Simplicity | **Pass** | Single pure function (or small API) + thin handler; no new infra. |

**Post-design re-check**: Design adds only `bird` package + route registration + tests — still one concern per task when tasks are written. **No violations.**

## Project Structure

### Documentation (this feature)

```text
specs/001-bird-rest-endpoint/
├── plan.md
├── research.md
├── data-model.md
├── quickstart.md
├── contracts/
│   └── bird-http.md
├── spec.md
└── tasks.md              # Phase 2 (/speckit.tasks) — not created by this command
```

### Source Code (repository root)

```text
GoRestController/
├── main.go                 # register GET /getBird, delegate to bird package
├── bird/
│   ├── label.go            # ResolveLabel(selection string) string
│   └── label_test.go       # table-driven unit tests for mapping + edges
├── go.mod
├── go.sum
└── Dockerfile / kubernetes_deployment_yaml  # unchanged unless task explicitly requires
```

**Structure Decision**: Introduce package **`bird`** for the resolution rule so **FR-005** is satisfied (handler does not duplicate mapping). Keep route registration in **`main.go`** to match existing handlers (`/getMammal`, `/getEnvironment`). Tests live next to logic (`bird/label_test.go`); optional later task for `httptest` handler integration if desired.

## Complexity Tracking

> No constitution violations required for this feature; table not used.
