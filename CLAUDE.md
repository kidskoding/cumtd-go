# cumtd-go workspace

Go workspace for the CUMTD API v3 SDK and example projects.
First community Go SDK for the Champaign-Urbana Mass Transit District API v3.

## graphify

This project has a graphify knowledge graph at graphify-out/.

Rules:
- Before answering architecture or codebase questions, read graphify-out/GRAPH_REPORT.md for god nodes and community structure
- If graphify-out/wiki/index.md exists, navigate it instead of reading raw files
- After modifying code files in this session, run `python3 -c "from graphify.watch import _rebuild_code; from pathlib import Path; _rebuild_code(Path('.'))"` to keep the graph current

## Modules

- `.` — SDK (`github.com/kidskoding/cumtd-go`) — standard library only, no third-party deps
- `examples/departure-board/` — minimal stdout departure board example
- `examples/agent-tool/` — LLM agent tool wrapping SDK methods

## Go version

1.26+ (workspace uses go.work; module requires go 1.26.2)

## Commands

```bash
go test -race ./...                          # run all SDK tests
go test -race -coverpkg=./... ./...          # run with cross-package coverage
go vet ./...                                 # vet SDK
go work sync                                 # sync workspace
```

## Key conventions

- Every public method takes `context.Context` as first arg
- All optional params use `*XxxOptions` structs, never naked args
- All path params go through `url.PathEscape`
- Use `errors.As` for error type checking, never string matching
- No global state — everything through `*Client`
- Standard library only inside `cumtd/` package

## API

- Base URL: `https://api.mtd.dev`
- Auth: `X-Api-Key` header (NOT a query param — changed from v2)
- Docs: https://mtd.dev

## Project structure

```
cumtd/          # importable SDK package
cumtd/types/    # domain types (one file per domain)
cumtd/tests/    # black-box tests (package cumtd_test)
internal/       # non-importable helpers
examples/       # usage examples (any language)
```

## Release checklist

- [ ] All 20 endpoints implemented
- [ ] All tests passing with -race
- [ ] go vet ./... clean
- [ ] staticcheck ./... clean
- [ ] Coverage ≥ 80%
- [ ] All exported symbols have godoc comments
- [ ] README complete with badges
- [ ] `git tag v0.1.0 && git push origin v0.1.0`
