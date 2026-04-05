# Data model: Bird lookup (001-bird-rest-endpoint)

This feature is **stateless**; there is no database or persisted entity. Below are **logical** constructs for documentation and tests.

## Selection input

| Concept | Description |
|---------|-------------|
| **Selection parameter** | Single string provided by caller (REST query `bird` or service argument). |
| **Validation** | No trimming in v1. Comparison to `Eagle` is **case-sensitive** byte-for-byte equality. |
| **Missing / empty** | Absent `bird` query → empty string input → maps to **Chickadee**. |
| **Whitespace-only** | If the value is only spaces/tabs (e.g. `bird=%20`), it is **not** equal to `Eagle` → **Chickadee**. |

## Bird label (output)

| Value | When |
|--------|------|
| `American Bald Eagle` | Selection is exactly `Eagle`. |
| `Chickadee` | Any other string, including empty. |

## Relationships

- **One input** → **one output** (function); no collections, no lifecycle, no IDs.
