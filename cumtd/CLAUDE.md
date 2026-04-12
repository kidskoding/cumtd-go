# package cumtd — MTD API v3 SDK

Core SDK package. Library only — no `main.go`, no third-party deps.

## File responsibilities

| File | Responsibility |
|------|---------------|
| `client.go` | `Client` struct, `New()`, functional `Option` pattern |
| `transport.go` | `get()` method, envelope decoder, header injection |
| `errors.go` | `APIError`, `RateLimitError`, `ValidationError`, `require()` |
| `routes.go` | `GetRouteGroups`, `GetRouteGroup`, `GetRoute` |
| `stops.go` | `GetStops`, `GetStop`, `SearchStops`, `GetStopSchedule`, `GetStopTrips`, `GetStopRouteGroups` |
| `departures.go` | `GetDepartures` — primary real-time endpoint |
| `trips.go` | `GetTrips`, `GetTrip` |
| `vehicles.go` | `GetVehicles`, `GetVehicle`, `GetVehicleLocation`, `GetVehicleLocations`, `GetVehicleConfigurations`, `GetVehicleConfiguration` |
| `shapes.go` | `GetShape`, `GetShapePolyline` |

## All 20 endpoints

| Method | Path |
|--------|------|
| `GetRouteGroups` | `GET /routes/groups` |
| `GetRouteGroup` | `GET /routes/groups/{id}` |
| `GetRoute` | `GET /routes/{id}` |
| `GetStops` | `GET /stops` |
| `GetStop` | `GET /stops/{id}` |
| `SearchStops` | `GET /stops/search` |
| `GetStopSchedule` | `GET /stops/{id}/schedule` |
| `GetStopTrips` | `GET /stops/{id}/trips` |
| `GetStopRouteGroups` | `GET /stops/{id}/route-groups` |
| `GetDepartures` | `GET /stops/{id}/departures` |
| `GetTrips` | `GET /trips` |
| `GetTrip` | `GET /trips/{id}` |
| `GetVehicles` | `GET /vehicles` |
| `GetVehicle` | `GET /vehicles/{id}` |
| `GetVehicleLocation` | `GET /vehicles/{id}/location` |
| `GetVehicleLocations` | `GET /vehicles/locations` |
| `GetVehicleConfigurations` | `GET /vehicles/configurations` |
| `GetVehicleConfiguration` | `GET /vehicles/configurations/{id}` |
| `GetShape` | `GET /shapes/{id}` |
| `GetShapePolyline` | `GET /shape/{id}/polyline` |

## Critical implementation notes

- `GetShape` uses `/shapes/{id}` (plural)
- `GetShapePolyline` uses `/shape/{id}/polyline` (singular — upstream spec bug, match exactly)
- Auth header is `X-Api-Key`, not a query param
- Fields typed as `any` (`MinutesTillDeparture`, `SortNumber`, `StopSequence`) — use `internal/coerce`
- All nullable fields are pointer types (`*string`, `*Coordinates`, etc.)
- Call `require(field, s)` at top of every method with a required path param

## Adding a new endpoint

1. Add method to the appropriate domain file (or create new file if new domain)
2. Call `require()` for any path params
3. Build `url.Values` for query params
4. Call `c.get(ctx, path, params, &out)`
5. Add fixture JSON in `internal/testutil/fixtures/`
6. Add test cases in `cumtd/tests/`

## Error handling pattern

```go
deps, err := client.GetDepartures(ctx, stopID, nil)
if err != nil {
    var apiErr *cumtd.APIError
    var rateLimitErr *cumtd.RateLimitError
    var valErr *cumtd.ValidationError
    switch {
    case errors.As(err, &rateLimitErr):
        // retry after rateLimitErr.RetryAfter
    case errors.As(err, &apiErr):
        // apiErr.StatusCode, apiErr.Body
    case errors.As(err, &valErr):
        // valErr.Field, valErr.Message
    }
}
```
