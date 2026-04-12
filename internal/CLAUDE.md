# internal/ — non-importable helpers

Packages here are not importable outside this module.

## Packages

### `coerce/`

Helpers for API spec fields that allow `int | string` (typed as `any` in Go structs).

```go
coerce.Int(v any) (int, error)
coerce.Float64(v any) (float64, error)
```

Handles: `float64` (JSON default), `int`, `string`, `nil`. Returns error for unsupported types.

**Never** use raw type assertions on `any`-typed fields outside of `coerce/`.

### `testutil/`

Mock HTTP server for tests. Never make real API calls in tests.

```go
srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
    "/routes/groups": {
        StatusCode: 200,
        Body:       testutil.MustReadFixture(t, "route_groups.json"),
    },
})
client := cumtd.New("test-key", cumtd.WithBaseURL(srv.URL))
```

`AssertQuery` — assert query params were sent correctly:
```go
"/stops/STOP1/departures": {
    StatusCode: 200,
    Body:       testutil.MustReadFixture(t, "departures.json"),
    AssertQuery: func(t *testing.T, q url.Values) {
        if q.Get("routes") != "ILLINI" { t.Errorf(...) }
    },
},
```

`MustReadFixture` locates files relative to `server.go` using `runtime.Caller(0)` — works from any test package in the module.

## testutil/fixtures/

One JSON file per endpoint. Always the API envelope format:
```json
{ "data": [...], "error": null }
```

Include ≥ 2 items for collections. Include all fields (including nullable ones) for single-item endpoints.

| Fixture | Endpoint |
|---------|----------|
| `route_groups.json` | `GET /routes/groups` |
| `route_group.json` | `GET /routes/groups/{id}` |
| `route.json` | `GET /routes/{id}` |
| `stops.json` | `GET /stops` |
| `stop.json` | `GET /stops/{id}` |
| `stop_search.json` | `GET /stops/search` |
| `stop_schedule.json` | `GET /stops/{id}/schedule` |
| `stop_trips.json` | `GET /stops/{id}/trips` |
| `stop_route_groups.json` | `GET /stops/{id}/route-groups` |
| `departures.json` | `GET /stops/{id}/departures` |
| `trips.json` | `GET /trips` |
| `trip.json` | `GET /trips/{id}` |
| `vehicles.json` | `GET /vehicles` |
| `vehicle.json` | `GET /vehicles/{id}` |
| `vehicle_location.json` | `GET /vehicles/{id}/location` |
| `vehicle_locations.json` | `GET /vehicles/locations` |
| `vehicle_configurations.json` | `GET /vehicles/configurations` |
| `vehicle_configuration.json` | `GET /vehicles/configurations/{id}` |
| `shape.json` | `GET /shapes/{id}` |
| `shape_polyline.json` | `GET /shape/{id}/polyline` |
