# HTTP contract: GET /getBird

**Version**: 1.0 (feature 001-bird-rest-endpoint)  
**Spec reference**: [spec.md](../spec.md)

## Endpoint

| Item | Value |
|------|--------|
| Method | `GET` |
| Path | `/getBird` |
| Purpose | Resolve a bird **display label** from one selection token. |

## Request

### Query parameters

| Name | Required | Description |
|------|----------|-------------|
| `bird` | No | Selection token. If omitted or any value other than exactly `Eagle`, the response body is `Chickadee`. If exactly `Eagle`, body is `American Bald Eagle`. Matching is case-sensitive; no trimming. |

### Examples

```http
GET /getBird?bird=Eagle HTTP/1.1
Host: localhost:8080
```

```http
GET /getBird?bird=sparrow HTTP/1.1
Host: localhost:8080
```

```http
GET /getBird HTTP/1.1
Host: localhost:8080
```

## Response (success)

| Item | Value |
|------|--------|
| Status | `200 OK` |
| `Content-Type` | `text/plain; charset=utf-8` |
| Body | **Only** the label: `American Bald Eagle` or `Chickadee`. Implementation MAY suffix a **single** newline after the label (**FR-006** / spec). |

### Example bodies

- Input `bird=Eagle` → body (logical): `American Bald Eagle`
- Any other case → body (logical): `Chickadee`

## Service contract (in-process)

Same mapping as HTTP:

- **Input**: `selection string`
- **Output**: `American Bald Eagle` if `selection == "Eagle"`, else `Chickadee`

REST handler MUST call this logic without reimplementing the rules (**FR-005**).
