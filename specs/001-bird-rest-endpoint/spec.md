# Feature Specification: Bird lookup via service and REST

**Feature Branch**: `001-bird-rest-endpoint`  
**Created**: 2026-04-04  
**Status**: Draft  
**Input**: User description: "Create a new service that returns a Bird. This service should be accessible via rest end point. Both the service and the rest end point should accept a single parameter. If 'Eagle' is passed then the service and the rest handler should return 'American Bald Eagle', otherwise the service and rest handler should return 'Chikadee'." *(Normative spelling for the default label in this document: **Chickadee**.)*

## Clarifications

### Session 2026-04-04

- Q: How should the REST response carry the bird label (plain text vs JSON)? → A: **Plain text** body containing **only** the label; use `Content-Type: text/plain; charset=utf-8` (or equivalent plain-text typing).

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Look up bird label through the API (Priority: P1)

An API consumer sends **one** input that selects which bird label they want resolved. When that input is exactly the eagle selection token, they receive the full eagle label; for any other input they receive the default small-bird label.

**Why this priority**: This is the primary way external clients integrate and delivers the visible behavior end-to-end.

**Independent Test**: Call the REST interface with the eagle token and with a non-eagle token; assert the two distinct labels are returned without involving unrelated features.

**Acceptance Scenarios**:

1. **Given** a REST client, **When** it submits a request whose single selection parameter equals `Eagle`, **Then** the response is **plain text** whose payload is the label `American Bald Eagle` (an optional single trailing newline is permitted).
2. **Given** a REST client, **When** it submits a request whose single selection parameter equals any value other than `Eagle` (for example `sparrow`, empty, or another token), **Then** the response is **plain text** whose payload is the label `Chickadee` (an optional single trailing newline is permitted).

---

### User Story 2 - Same mapping from the core service (Priority: P2)

An internal or in-process caller uses the bird resolution **service** with the same single selection parameter. Results match the REST behavior for the same parameter values.

**Why this priority**: Keeps one source of truth for labels and avoids REST-only logic.

**Independent Test**: Invoke the service directly with `Eagle` and with a non-eagle value; compare outcomes to the REST scenarios above.

**Acceptance Scenarios**:

1. **Given** the bird resolution service, **When** it is called with the parameter value `Eagle`, **Then** it returns `American Bald Eagle`.
2. **Given** the bird resolution service, **When** it is called with any parameter value other than `Eagle`, **Then** it returns `Chickadee`.

---

### Edge Cases

- **Missing parameter**: Treated as “not `Eagle`” and resolves to `Chickadee`.
- **Whitespace-only parameter**: Treated as “not `Eagle`” and resolves to `Chickadee` (no trimming in v1 per Assumptions).
- **Case variants** (`eagle`, `EAGLE`): Treated as “not `Eagle`” unless explicitly changed in a future amendment; matching is **case-sensitive** to the literal `Eagle`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST provide a **bird resolution service** that accepts **exactly one** selection parameter and returns a bird **display label** (the resolved “Bird” outcome for this feature).
- **FR-002**: When the selection parameter is **exactly** the string `Eagle`, the service MUST return the label `American Bald Eagle`.
- **FR-003**: When the selection parameter is **any** value other than exactly `Eagle`, the service MUST return the label `Chickadee`.
- **FR-004**: The system MUST expose a **REST** entry point that accepts the **same single selection parameter** (same meaning as the service input) and returns the **same** label outcomes as **FR-002** and **FR-003** for equivalent inputs.
- **FR-005**: The REST handler MUST delegate to (or otherwise reuse without duplicating rules) the bird resolution service so mapping rules exist in one place.
- **FR-006**: For successful REST responses, the response body MUST be **plain text** consisting of **only** the bird label (no JSON or other structured wrapper). The response MUST use a `Content-Type` appropriate for plain text (e.g. `text/plain; charset=utf-8`). A **single** trailing newline after the label is permitted.

### Key Entities

- **Selection parameter**: A single token provided by the caller; only the value `Eagle` selects the eagle label.
- **Bird label**: The human-readable string returned to the caller—either `American Bald Eagle` or `Chickadee`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: For a representative set of at least five non-`Eagle` inputs (including empty and at least one arbitrary word), **100%** of outcomes are the label `Chickadee` for both the service and the REST entry point.
- **SC-002**: For the input exactly `Eagle`, **100%** of outcomes are the label `American Bald Eagle` for both the service and the REST entry point.
- **SC-003**: For any tested input, the service and REST entry point return **identical** labels (no drift between layers).
- **SC-004**: For the REST entry point, **100%** of successful responses use a plain-text body and `Content-Type` consistent with **FR-006**, and the body (ignoring at most one trailing newline) equals exactly one of the two allowed labels.

## Assumptions

- REST shape is **`GET /getBird`** with query parameter **`bird`** (see [plan.md](./plan.md) and [contracts/bird-http.md](./contracts/bird-http.md)); semantics remain **one** logical selection input aligned with the service parameter.
- **String matching** for `Eagle` is **case-sensitive** and **no trimming** is applied unless the implementation plan explicitly adds trimming; if trimming is added, only `Eagle` after trimming counts as the eagle case.
- **Authentication, authorization, and rate limiting** are out of scope unless an existing platform policy already applies to all endpoints.
- The default label uses standard English spelling **`Chickadee`**.
