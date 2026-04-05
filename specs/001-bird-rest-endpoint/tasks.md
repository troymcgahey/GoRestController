---
description: "Task list for feature 001-bird-rest-endpoint (bird service + GET /getBird)"
---

# Tasks: Bird lookup via service and REST

**Input**: Design documents from `/specs/001-bird-rest-endpoint/`  
**Prerequisites**: [plan.md](./plan.md), [spec.md](./spec.md), [data-model.md](./data-model.md), [contracts/bird-http.md](./contracts/bird-http.md), [research.md](./research.md), [quickstart.md](./quickstart.md)

**Tests**: Included because [plan.md](./plan.md) requires `bird/label_test.go` and `go test ./...` for verification.

**Constitution**: After each task, **commit** (message references task ID) and open a **PR** before starting the next task ([`.specify/memory/constitution.md`](../../.specify/memory/constitution.md)).

**Organization**: Phases follow dependencies: service logic first, then **User Story 1 (REST)** and **User Story 2 (service tests)** in parallel where safe, then polish.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files; shared prerequisite complete)
- **[Story]**: [US1] = spec User Story 1 (REST), [US2] = spec User Story 2 (service)
- Paths are relative to repository root `GoRestController/` unless noted

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Confirm the existing Go module is ready; no new services or databases.

- [x] T001 Confirm module readiness by running `go list ./...` from `GoRestController/` (validates `GoRestController/go.mod`)

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Bird resolution **service** required by **FR-001**, **FR-005**, and the HTTP contract before any handler or story-specific test work.

**⚠️ CRITICAL**: Complete **T002** before **T003** or **T004**.

- [x] T002 Implement `ResolveLabel(selection string) string` in `GoRestController/bird/label.go` per [contracts/bird-http.md](./contracts/bird-http.md) (service contract) and [data-model.md](./data-model.md)

**Checkpoint**: Service package compiles (`go build ./...` from `GoRestController/`).

---

## Phase 3: User Story 1 — Look up bird label through the API (Priority: P1) 🎯 MVP

**Goal**: Expose **GET /getBird?bird=...** with **plain-text** body and correct **Content-Type** (**FR-004**, **FR-006**).

**Independent Test**: `curl` against `/getBird` with `bird=Eagle` and a non-eagle value; compare bodies to [spec.md](./spec.md) acceptance scenarios (see [quickstart.md](./quickstart.md)).

### Implementation for User Story 1

- [x] T003 [US1] Add `getBird` HTTP handler and register `GET /getBird` in `GoRestController/main.go` reading query `bird`, calling `bird.ResolveLabel`, setting `Content-Type: text/plain; charset=utf-8`, writing only the label (optional single trailing newline per **FR-006**) per [contracts/bird-http.md](./contracts/bird-http.md)

**Checkpoint**: Manual or scripted HTTP checks pass for US1; handler delegates to `bird` only (**FR-005**).

---

## Phase 4: User Story 2 — Same mapping from the core service (Priority: P2)

**Goal**: Lock **FR-002**/**FR-003** and edge cases with **direct package tests** (independent of HTTP).

**Independent Test**: `go test ./bird -run Test...` from `GoRestController/` covers `Eagle`, non-eagle, empty, whitespace-only, and case variants per [spec.md](./spec.md) **Edge Cases**.

### Tests for User Story 2

- [x] T004 [P] [US2] Add table-driven tests in `GoRestController/bird/label_test.go` for **FR-002**, **FR-003**, and edge cases (missing/empty equivalent, whitespace-only, `eagle` / `EAGLE`) per [spec.md](./spec.md)

**Checkpoint**: `go test ./...` passes; service behavior provable without the HTTP server.

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Validate docs and end-to-end quickstart after code exists.

- [x] T005 Run `go test ./...` from `GoRestController/` and execute curl steps in `specs/001-bird-rest-endpoint/quickstart.md` against a local server

---

## Dependencies & Execution Order

### Phase Dependencies

- **Phase 1 (Setup)**: **T001** — no predecessors
- **Phase 2 (Foundational)**: **T002** — after **T001**
- **Phase 3 (US1)** / **Phase 4 (US2)**: **T003** and **T004** — both after **T002**; **T003** and **T004** may run in **parallel** (different files: `main.go` vs `bird/label_test.go`)
- **Phase 5 (Polish)**: **T005** — after **T003** and **T004**

### User Story Dependencies

- **US1 (P1)**: Depends on **T002** (service must exist for **FR-005**)
- **US2 (P2)**: Depends on **T002** only; does not require **T003**

### Parallel Opportunities

- After **T002**: run **T003** [US1] and **T004** [P] [US2] in parallel if two contributors (or sequential either order).

---

## Parallel Example: After T002

```bash
# Contributor A — REST (US1)
# Task T003: edit GoRestController/main.go only

# Contributor B — service tests (US2)
# Task T004: add GoRestController/bird/label_test.go only
```

---

## Implementation Strategy

### MVP First (User Story 1 minimal path)

1. **T001** → **T002** → **T003** → stop and validate with [quickstart.md](./quickstart.md) (HTTP only).  
2. Add **T004** then **T005** before merge to main if you require tests green in CI.

### Constitution-aligned delivery

1. Complete **T001** → commit + PR → merge.  
2. Complete **T002** → commit + PR → merge.  
3. Complete **T003** → commit + PR → merge (or **T004** first if you prefer tests before HTTP in your branch strategy—still one task per PR).  
4. Complete **T004** → commit + PR → merge.  
5. Complete **T005** → commit + PR → merge.

### Single-developer sequential order

**T001** → **T002** → **T003** → **T004** → **T005**

---

## Notes

- Every task lists concrete paths under `GoRestController/` for implementation.
- **[P]** on **T004** marks parallel eligibility with **T003** after **T002**.
- If CI is added later, wire `go test ./...` in a separate task/PR per constitution.

---

## Summary

| Metric | Value |
|--------|--------|
| **Total tasks** | 5 |
| **Per user story** | US1: 1 impl task (**T003**); US2: 1 test task (**T004**) |
| **Setup / foundation** | **T001**, **T002** |
| **Polish** | **T005** |
| **Parallel after T002** | **T003** and **T004** |
| **Suggested MVP** | **T001** + **T002** + **T003** (add **T004**/**T005** before production if policy requires tests) |
