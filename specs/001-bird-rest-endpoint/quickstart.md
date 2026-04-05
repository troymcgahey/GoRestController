# Quickstart: Bird endpoint

Prerequisites: Go 1.23+, repo root `GoRestController`.

## Run the server

```bash
cd /Users/troymcgahey/Projects/Go/GoRestController
go run .
```

Default listen address: `:8080` (see `main.go`).

## Try the bird endpoint

After implementation (see [plan.md](./plan.md)):

```bash
# Eagle → American Bald Eagle
curl -s "http://localhost:8080/getBird?bird=Eagle"

# Anything else → Chickadee
curl -s "http://localhost:8080/getBird?bird=sparrow"

# Missing parameter → Chickadee
curl -s "http://localhost:8080/getBird"
```

Expect **plain text** responses and `Content-Type: text/plain; charset=utf-8`.

## Verify tests

```bash
go test ./...
```

## Contract

See [contracts/bird-http.md](./contracts/bird-http.md).
