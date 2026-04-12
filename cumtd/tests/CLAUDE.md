# cumtd/tests — SDK tests

Black-box tests for the `cumtd` package. All files use `package cumtd_test`.

## File map

| File | Tests |
|------|-------|
| `client_test.go` | `New()` defaults, each `With*` option |
| `transport_test.go` | Headers on every request, context cancel propagation |
| `errors_test.go` | `errors.As` for all three error types |
| `routes_test.go` | Routes happy path, 404, 429, empty ID → ValidationError |
| `stops_test.go` | Stops happy path, `ExcludeBoardingPoints`, 404, empty ID |
| `departures_test.go` | Departures happy path, Routes param, Time param, 429, empty ID |
| `trips_test.go` | Trips happy path, 404, empty ID |
| `vehicles_test.go` | Vehicles happy path, locations (no ID), 404, empty ID |
| `shapes_test.go` | GetShape (plural path), GetShapePolyline (singular path), empty ID |

## Conventions

- All tests table-driven with `t.Run` where multiple cases exist
- All tests `-race` compatible (no shared mutable state)
- No real HTTP calls — always use `testutil.NewMockServer`
- Fixtures loaded via `testutil.MustReadFixture(t, "filename.json")`
- Test query params with `AssertQuery` on `MockRoute`

## Adding a test

```go
func TestGetFoo(t *testing.T) {
    srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
        "/foo/BAR": {StatusCode: 200, Body: testutil.MustReadFixture(t, "foo.json")},
    })
    c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
    result, err := c.GetFoo(context.Background(), "BAR")
    if err != nil { t.Fatalf("unexpected error: %v", err) }
    // assert result fields
}
```

## Coverage target

≥ 80% across the module (`go test -race -coverpkg=./... ./...`).
Current: 84.3%.
