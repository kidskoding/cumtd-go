# cumtd-go

[![CI](https://github.com/kidskoding/cumtd-go/actions/workflows/ci.yml/badge.svg)](https://github.com/kidskoding/cumtd-go/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/kidskoding/cumtd-go/cumtd.svg)](https://pkg.go.dev/github.com/kidskoding/cumtd-go/cumtd)
[![License](https://img.shields.io/github/license/kidskoding/cumtd-go)](LICENSE)

Go SDK for the [Champaign-Urbana Mass Transit District (MTD) API v3](https://mtd.dev).
First community Go wrapper for the MTD API v3.

## Install

```bash
go get github.com/kidskoding/cumtd-go/cumtd
```

## Quick start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/kidskoding/cumtd-go/cumtd"
)

func main() {
    client := cumtd.New("YOUR_API_KEY")

    deps, err := client.GetDepartures(context.Background(), "STOP1", nil)
    if err != nil {
        log.Fatal(err)
    }

    for _, d := range deps {
        fmt.Printf("%s — departs in %s\n", d.UniqueID, d.DepartsIn)
    }
}
```

Get an API key at [developer.mtd.org](https://developer.mtd.org).

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

## Error handling

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

## Options

```go
client := cumtd.New("YOUR_API_KEY",
    cumtd.WithBaseURL("https://api.mtd.dev/api/v3"), // override base URL
    cumtd.WithTimeout(30*time.Second),               // custom timeout
    cumtd.WithHTTPClient(myHTTPClient),              // bring your own client
    cumtd.WithUserAgent("my-app/1.0"),               // custom User-Agent
)
```

Optional parameters use `*Options` structs:

```go
// Filter departures by route and time
deps, err := client.GetDepartures(ctx, "STOP1", &cumtd.GetDeparturesOptions{
    Routes: "ILLINI,GOLDLINE",
    Time:   "08:00:00",
})

// Exclude boarding points from stop results
stops, err := client.GetStops(ctx, &cumtd.GetStopsOptions{
    ExcludeBoardingPoints: true,
})
```

## Examples

- [`examples/departure-board/`](examples/departure-board/) — print upcoming departures to stdout
- [`examples/agent-tool/`](examples/agent-tool/) — wrap endpoints as LLM tool definitions

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[Apache 2.0](LICENSE)
