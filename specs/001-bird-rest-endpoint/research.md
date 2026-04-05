# Research: Bird REST endpoint (001-bird-rest-endpoint)

## REST parameter binding

**Decision**: Expose **GET** `/getBird` with a **single query parameter** `bird` carrying the selection token (e.g. `bird=Eagle`).

**Rationale**: Matches existing style in `main.go` (`GET` + query on `/getMammal`). Safe, idempotent read; easy to exercise with `curl`; one logical parameter without a JSON body.

**Alternatives considered**:

- Path segment (`GET /getBird/Eagle`): Slightly nicer URLs but diverges from current handler patterns and encodes spaces/special chars awkwardly.
- **POST** + JSON body: Possible but heavier for a two-outcome lookup; not aligned with existing codebase.

## HTTP status and errors

**Decision**: Return **200 OK** for every request that reaches the handler successfully, including missing or empty `bird` query (body = `Chickadee` per spec).

**Rationale**: Spec treats missing/non-`Eagle` as **normal** resolution to `Chickadee`, not a client error.

**Alternatives considered**: **400** for missing parameter — rejected; contradicts edge-case requirements.

## Plain-text response

**Decision**: Set `Content-Type: text/plain; charset=utf-8` and write the label string; allow implementation to use `Fprintln`-style output (label + one newline) to stay consistent with other handlers.

**Rationale**: Aligns with clarification session and **FR-006**; matches existing `fmt.Fprintln` usage in `main.go`.

## Package layout

**Decision**: New package **`bird`** with `ResolveLabel(selection string) string` (name can be adjusted in implementation but responsibility is fixed).

**Rationale**: Satisfies **FR-005** (single source of truth); unit-testable without HTTP.

**Alternatives considered**: Logic only in `main.go` — rejected; duplicates risk if REST evolves.
