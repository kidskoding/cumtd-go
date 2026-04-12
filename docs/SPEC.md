# cumtd-go — Go API Wrapper for CUMTD API v3

> Production-grade, fully modular Go SDK for the Champaign-Urbana Mass Transit
> District (MTD) Developer API v3. Built against the official OpenAPI spec at
> https://mtd.dev. First community Go SDK for MTD API v3.

---

## Repository Structure

```
cumtd-go/                           # repo root = SDK module
├── CLAUDE.md
├── go.work
├── go.work.sum
├── go.mod                          # module github.com/<handle>/cumtd-go
├── go.sum
│
├── cumtd/                          # package cumtd — the importable library
│   ├── CLAUDE.md
│   ├── client.go
│   ├── transport.go
│   ├── errors.go
│   ├── routes.go
│   ├── stops.go
│   ├── departures.go
│   ├── trips.go
│   ├── vehicles.go
│   ├── shapes.go
│   └── types/
│       ├── common.go
│       ├── routes.go
│       ├── stops.go
│       ├── departures.go
│       ├── trips.go
│       ├── vehicles.go
│       └── shapes.go
│
├── internal/
│   ├── CLAUDE.md
│   ├── coerce/
│   │   └── coerce.go
│   └── testutil/
│       ├── server.go
│       └── fixtures/
│           ├── route_groups.json
│           ├── route_group.json
│           ├── route.json
│           ├── stops.json
│           ├── stop.json
│           ├── stop_search.json
│           ├── stop_schedule.json
│           ├── stop_trips.json
│           ├── stop_route_groups.json
│           ├── departures.json
│           ├── trips.json
│           ├── trip.json
│           ├── vehicles.json
│           ├── vehicle.json
│           ├── vehicle_location.json
│           ├── vehicle_locations.json
│           ├── vehicle_configurations.json
│           ├── vehicle_configuration.json
│           ├── shape.json
│           └── shape_polyline.json
│
├── cumtd/tests/
│   ├── client_test.go
│   ├── transport_test.go
│   ├── errors_test.go
│   ├── routes_test.go
│   ├── stops_test.go
│   ├── departures_test.go
│   ├── trips_test.go
│   ├── vehicles_test.go
│   └── shapes_test.go
│
├── examples/
│   ├── departure-board/
│   └── agent-tool/
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── release.yml
│
├── README.md
├── CONTRIBUTING.md
├── LICENSE
└── SPEC.md
```

---

## go.work

```
go 1.22

use (
    .
)
```

---

## Module Paths

| Directory | Module path |
|-----------|-------------|
| `.` (root) | `github.com/<handle>/cumtd-go` |

---

## CLAUDE.md Files

---

### `/CLAUDE.md`

```markdown
# cumtd-go workspace

Go workspace for the CUMTD API v3 SDK and example projects.
First community Go SDK for the Champaign-Urbana Mass Transit District API v3.

## Modules
- `.` — SDK (github.com/<handle>/cumtd-go) — standard library only, no third-party deps
- `examples/departure-board/` — minimal stdout example
- `examples/agent-tool/` — LLM agent tool example

## Go version
1.22+ (generics used in transport envelope decoder)

## Commands
- `go test -race ./...` — run all SDK tests
- `go vet ./...` — vet SDK
- `go work sync` — sync workspace

## Key conventions
- Every public method takes context.Context as first arg
- All optional params use *XxxOptions structs, never naked args
- All path params go through url.PathEscape
- Use errors.As for error type checking, never string matching
- No global state — everything through *Client
- Standard library only inside cumtd/ package

## API
- Base URL: https://api.mtd.dev/api/v3
- Auth: X-Api-Key header (NOT a query param — changed from v2)
- Docs: https://mtd.dev
- OpenAPI spec: npm package @mtd.org/developer-api-spec
```

---

### `/cumtd/CLAUDE.md`

```markdown
# package cumtd — MTD API v3 SDK

This is the core SDK package. Library only — no main.go, no third-party deps.

## File responsibilities
- client.go — Client struct, New(), functional Option pattern
- transport.go — single get() method, envelope decoder, header injection
- errors.go — APIError, RateLimitError, ValidationError, require()
- routes.go — GetRouteGroups, GetRouteGroup, GetRoute
- stops.go — GetStops, GetStop, SearchStops, GetStopSchedule, GetStopTrips, GetStopRouteGroups
- departures.go — GetDepartures (separate from stops — primary real-time endpoint)
- trips.go — GetTrips, GetTrip
- vehicles.go — GetVehicles, GetVehicle, GetVehicleLocation, GetVehicleLocations, GetVehicleConfigurations, GetVehicleConfiguration
- shapes.go — GetShape, GetShapePolyline
- types/ — all domain types, one file per domain group

## All 20 endpoints
| Method | Path |
|--------|------|
| GetRouteGroups | GET /routes/groups |
| GetRouteGroup | GET /routes/groups/{id} |
| GetRoute | GET /routes/{id} |
| GetStops | GET /stops |
| GetStop | GET /stops/{id} |
| SearchStops | GET /stops/search |
| GetStopSchedule | GET /stops/{id}/schedule |
| GetStopTrips | GET /stops/{id}/trips |
| GetStopRouteGroups | GET /stops/{id}/route-groups |
| GetDepartures | GET /stops/{id}/departures |
| GetTrips | GET /trips |
| GetTrip | GET /trips/{id} |
| GetVehicles | GET /vehicles |
| GetVehicle | GET /vehicles/{id} |
| GetVehicleLocation | GET /vehicles/{id}/location |
| GetVehicleLocations | GET /vehicles/locations |
| GetVehicleConfigurations | GET /vehicles/configurations |
| GetVehicleConfiguration | GET /vehicles/configurations/{id} |
| GetShape | GET /shapes/{id} |
| GetShapePolyline | GET /shape/{id}/polyline |

## Critical notes
- GetShape uses /shapes/{id} (plural)
- GetShapePolyline uses /shape/{id}/polyline (singular — upstream spec bug, match exactly)
- Auth header is X-Api-Key, not a query param
- Fields typed as `any` (MinutesTillDeparture, SortNumber, StopSequence) — use internal/coerce
- All nullable fields are pointer types (*string, *Coordinates, etc)
- Call require(field, s) at top of every method with a required path param

## Error handling pattern
```go
deps, err := client.GetDepartures(ctx, stopID, nil)
if err != nil {
    var apiErr *cumtd.APIError
    var rateLimitErr *cumtd.RateLimitError
    var validErr *cumtd.ValidationError
    switch {
    case errors.As(err, &rateLimitErr):
    case errors.As(err, &apiErr):
    case errors.As(err, &validErr):
    }
}
```

## Testing
- All tests in cumtd/tests/
- Use internal/testutil.NewMockServer and MustReadFixture
- Table-driven with t.Run
- Run with -race flag
- No real HTTP calls in any test
- Target ≥ 80% coverage
```

---

### `/internal/CLAUDE.md`

```markdown
# internal/

Internal packages — not importable outside this module.

## coerce/coerce.go
Helpers for spec fields that allow int | string:
- coerce.Int(v any) (int, error)
- coerce.Float64(v any) (float64, error)
Never use raw type assertions on any-typed fields outside of coerce/.

## testutil/server.go
Mock HTTP server for tests. Never make real API calls in tests.

### Usage
```go
srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
    "/routes/groups": {
        StatusCode: 200,
        Body:       testutil.MustReadFixture(t, "route_groups.json"),
    },
})
client := cumtd.New("test-key", cumtd.WithBaseURL(srv.URL))
```

### AssertQuery
```go
"/stops/STOP1/departures": {
    StatusCode: 200,
    Body:       testutil.MustReadFixture(t, "departures.json"),
    AssertQuery: func(t *testing.T, q url.Values) {
        assert.Equal(t, "route1,route2", q.Get("routes"))
    },
},
```

## testutil/fixtures/
One JSON file per endpoint. Always the API envelope format:
```json
{ "data": [...], "error": null }
```
Include at least 2 items for collections. Include all fields including
nullable ones for single-item endpoints.
```

---

### `/examples/departure-board/`

Minimal example — fetch and print upcoming departures for a stop to stdout.

**Purpose:** Copy-paste starting point for new users. Any language that can make HTTP requests works.

**What it demonstrates**
- Authenticating with `X-Api-Key` header
- Calling `GET /stops/{id}/departures`
- Printing departure info to stdout

**Config**
- `CUMTD_API_KEY` env var — required
- `CUMTD_STOP_ID` env var — required

**Rules**
- Keep it short and readable
- No non-essential dependencies
- Comments explain each step for readers unfamiliar with the API

---

### `/examples/agent-tool/`

Wraps CUMTD API endpoints as LLM tool definitions.

**Purpose:** Show how to expose transit data to an LLM agent. Any language works.

**Tools implemented**
| Tool name | Endpoint |
|-----------|----------|
| search_stops | GET /stops/search |
| get_departures | GET /stops/{id}/departures |
| get_vehicle_locations | GET /vehicles/locations |

**Config**
- `CUMTD_API_KEY` env var — required

**Rules**
- No LLM SDK dependency — tool definitions are plain data structures
- Generic dispatcher pattern, copy-paste friendly
- Comments explain the tool pattern for readers new to agent development

---

## Architecture Principles

1. **One responsibility per file.**
2. **No global state.** Everything through `*Client`.
3. **Standard library only for the SDK.**
4. **Context everywhere.**
5. **Options structs** for all optional params.
6. **Typed errors** — `errors.As`, never string matching.
7. **Fixture-based tests** — no real HTTP calls.
8. **Table-driven tests** with `t.Run`.

---

## `cumtd/client.go`

```go
const (
    DefaultBaseURL = "https://api.mtd.dev/api/v3"
    DefaultTimeout = 10 * time.Second
    Version        = "0.1.0"
)

type Client struct {
    apiKey     string
    baseURL    string
    httpClient *http.Client
    userAgent  string
}

type Option func(*Client)

func New(apiKey string, opts ...Option) *Client
func WithHTTPClient(hc *http.Client) Option
func WithBaseURL(u string) Option
func WithUserAgent(ua string) Option
func WithTimeout(d time.Duration) Option
```

---

## `cumtd/transport.go`

```go
func (c *Client) get(ctx context.Context, path string, params url.Values, dst any) error

type envelope[T any] struct {
    Data  T           `json:"data"`
    Error *apiErrBody `json:"error,omitempty"`
}

func decode[T any](body []byte) (T, error)
```

---

## `cumtd/errors.go`

```go
type APIError struct {
    StatusCode int
    Body       string
}
type RateLimitError struct {
    RetryAfter string
}
type ValidationError struct {
    Field   string
    Message string
}
func require(field, s string) error
```

---

## `internal/coerce/coerce.go`

```go
func Int(v any) (int, error)
func Float64(v any) (float64, error)
```

---

## Types (`cumtd/types/`)

### `common.go`
```go
type Coordinates struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}
type TripDirection struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
type DayType struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

### `routes.go`
```go
type RouteGroup struct {
    ID             string  `json:"id"`
    SortNumber     any     `json:"sortNumber"`
    RouteGroupName string  `json:"routeGroupName"`
    Color          string  `json:"color"`
    TextColor      string  `json:"textColor"`
    Routes         []Route `json:"routes"`
}
type Route struct {
    ID                    string  `json:"id"`
    Number                *string `json:"number"`
    FirstTrip             string  `json:"firstTrip"`
    LastTrip              string  `json:"lastTrip"`
    LastTripAfterMidnight bool    `json:"lastTripAfterMidnight"`
    DayType               any     `json:"dayType"`
    GtfsRoutes            []any   `json:"gtfsRoutes"`
    RouteGroupID          string  `json:"routeGroupId"`
}
```

### `stops.go`
```go
type StopBase struct {
    ID             string          `json:"id"`
    Name           string          `json:"name"`
    Code           *string         `json:"code"`
    Location       *Coordinates    `json:"location"`
    BoardingPoints []BoardingPoint `json:"boardingPoints"`
    StopGroups     []StopGroup     `json:"stopGroups"`
}
type BoardingPoint struct {
    ID       string       `json:"id"`
    Name     *string      `json:"name"`
    Location *Coordinates `json:"location"`
}
type StopGroup struct {
    ID   string  `json:"id"`
    Name *string `json:"name"`
}
type StopSearchResult struct {
    ID       string       `json:"id"`
    Name     string       `json:"name"`
    Code     *string      `json:"code"`
    Location *Coordinates `json:"location"`
}
type StopTime struct {
    StopID                string         `json:"stopId"`
    TripID                string         `json:"tripId"`
    RouteID               *string        `json:"routeId"`
    GtfsRouteID           *string        `json:"gtfsRouteId"`
    Direction             *TripDirection `json:"direction"`
    StopSequence          any            `json:"stopSequence"`
    ArrivalTime           string         `json:"arrivalTime"`
    ArrivalPastMidnight   bool           `json:"arrivalPastMidnight"`
    DepartureTime         string         `json:"departureTime"`
    DeparturePastMidnight bool           `json:"departurePastMidnight"`
    StopHeadsign          *string        `json:"stopHeadsign"`
}
```

### `departures.go`
```go
type Departure struct {
    StopID               string          `json:"stopId"`
    Headsign             *string         `json:"headsign"`
    Trip                 *DepartureTrip  `json:"trip"`
    BlockID              *string         `json:"blockId"`
    RecordedTime         string          `json:"recordedTime"`
    ScheduledDeparture   *string         `json:"scheduledDeparture"`
    EstimatedDeparture   *string         `json:"estimatedDeparture"`
    VehicleID            *string         `json:"vehicleId"`
    OriginStopID         *string         `json:"originStopId"`
    DestinationStopID    *string         `json:"destinationStopId"`
    Location             *Coordinates    `json:"location"`
    ShapeID              *string         `json:"shapeId"`
    MinutesTillDeparture any             `json:"minutesTillDeparture"`
    IsRealTime           bool            `json:"isRealTime"`
    IsHopper             bool            `json:"isHopper"`
    Destination          *string         `json:"destination"`
    DepartsIn            string          `json:"departsIn"`
    IsIStop              bool            `json:"isIStop"`
    UniqueID             string          `json:"uniqueId"`
    Route                *DepartureRoute `json:"route"`
}
type DepartureTrip struct {
    TripID    *string        `json:"tripId"`
    Direction *TripDirection `json:"direction"`
}
type DepartureRoute struct {
    ID           string  `json:"id"`
    RouteGroupID *string `json:"routeGroupId"`
    GtfsRouteID  string  `json:"gtfsRouteId"`
    LongName     *string `json:"longName"`
    ShortName    *string `json:"shortName"`
    Color        *string `json:"color"`
    TextColor    *string `json:"textColor"`
}
```

### `trips.go`
```go
type Trip struct {
    ID        string         `json:"id"`
    BlockID   string         `json:"blockId"`
    ShapeID   string         `json:"shapeId"`
    Headsign  string         `json:"headsign"`
    Direction *TripDirection `json:"direction"`
    Route     *Route         `json:"route"`
}
```

### `vehicles.go`
```go
type Vehicle struct {
    ID                     string  `json:"id"`
    VehicleConfigurationID string  `json:"vehicleConfigurationId"`
    IsActive               bool    `json:"isActive"`
    DateInService          *string `json:"dateInService"`
}
type VehicleLocation struct {
    ID          string               `json:"id"`
    Location    *Coordinates         `json:"location"`
    LastUpdated *string              `json:"lastUpdated"`
    Trip        *VehicleLocationTrip `json:"trip"`
    Route       *DepartureRoute      `json:"route"`
}
type VehicleLocationTrip struct {
    TripID    *string        `json:"tripId"`
    Direction *TripDirection `json:"direction"`
}
type VehicleConfiguration struct {
    ID             string                `json:"id"`
    Type           VehicleType           `json:"type"`
    PowertrainType VehiclePowertrainType `json:"powertrainType"`
    Capacity       int                   `json:"capacity"`
}
type VehicleType           string
type VehiclePowertrainType string
```

### `shapes.go`
```go
type Shape struct {
    ID          string       `json:"id"`
    ShapePoints []ShapePoint `json:"shapePoints"`
}
type ShapePoint struct {
    Sequence int     `json:"sequence"`
    Lat      float64 `json:"lat"`
    Lon      float64 `json:"lon"`
}
type ShapePolyline struct {
    ID       string `json:"id"`
    Polyline string `json:"polyline"`
}
```

---

## All 20 Endpoint Methods

### `cumtd/routes.go`
```go
func (c *Client) GetRouteGroups(ctx context.Context) ([]types.RouteGroup, error)
func (c *Client) GetRouteGroup(ctx context.Context, routeGroupID string) (*types.RouteGroup, error)
func (c *Client) GetRoute(ctx context.Context, routeID string) (*types.Route, error)
```

### `cumtd/stops.go`
```go
type GetStopsOptions struct {
    ExcludeBoardingPoints bool
}
type GetStopScheduleOptions struct {
    RouteID string
    Date    string // YYYY-MM-DD
}
func (c *Client) GetStops(ctx context.Context, opts *GetStopsOptions) ([]types.StopBase, error)
func (c *Client) GetStop(ctx context.Context, stopID string) (*types.StopBase, error)
func (c *Client) SearchStops(ctx context.Context, query string) ([]types.StopSearchResult, error)
func (c *Client) GetStopSchedule(ctx context.Context, stopID string, opts *GetStopScheduleOptions) ([]types.StopTime, error)
func (c *Client) GetStopTrips(ctx context.Context, stopID string) ([]types.Trip, error)
func (c *Client) GetStopRouteGroups(ctx context.Context, stopID string) ([]types.RouteGroup, error)
```

### `cumtd/departures.go`
```go
type GetDeparturesOptions struct {
    Routes string
    Time   string
}
func (c *Client) GetDepartures(ctx context.Context, stopID string, opts *GetDeparturesOptions) ([]types.Departure, error)
```

### `cumtd/trips.go`
```go
func (c *Client) GetTrips(ctx context.Context) ([]types.Trip, error)
func (c *Client) GetTrip(ctx context.Context, tripID string) (*types.Trip, error)
```

### `cumtd/vehicles.go`
```go
func (c *Client) GetVehicles(ctx context.Context) ([]types.Vehicle, error)
func (c *Client) GetVehicle(ctx context.Context, vehicleID string) (*types.Vehicle, error)
func (c *Client) GetVehicleLocation(ctx context.Context, vehicleID string) (*types.VehicleLocation, error)
func (c *Client) GetVehicleLocations(ctx context.Context) ([]types.VehicleLocation, error)
func (c *Client) GetVehicleConfigurations(ctx context.Context) ([]types.VehicleConfiguration, error)
func (c *Client) GetVehicleConfiguration(ctx context.Context, configID string) (*types.VehicleConfiguration, error)
```

### `cumtd/shapes.go`
```go
// NOTE: upstream spec inconsistency — /shapes/{id} vs /shape/{id}/polyline (no 's').
// Match exactly. Document with a comment in shapes.go.
func (c *Client) GetShape(ctx context.Context, shapeID string) (*types.Shape, error)
func (c *Client) GetShapePolyline(ctx context.Context, shapeID string) (*types.ShapePolyline, error)
```

---

## Testing

### Mock server (`internal/testutil/server.go`)
```go
type MockRoute struct {
    StatusCode  int
    Body        []byte
    AssertQuery func(t *testing.T, q url.Values)
}
func NewMockServer(t *testing.T, routes map[string]MockRoute) *httptest.Server
func MustReadFixture(t *testing.T, name string) []byte
```

### Coverage requirements

| Test file | Cases required |
|-----------|---------------|
| `client_test.go` | Default options; each With* option applied correctly |
| `transport_test.go` | Headers on every request; context cancel propagates; body closed |
| `errors_test.go` | errors.As works for all three types |
| `routes_test.go` | Happy path; 404; 429; empty ID → ValidationError |
| `stops_test.go` | Happy path; ExcludeBoardingPoints encoded; 404; empty stopID → ValidationError |
| `departures_test.go` | Happy path; Routes param encoded; Time param encoded; 429; empty stopID → ValidationError |
| `trips_test.go` | Happy path; 404; empty tripID → ValidationError |
| `vehicles_test.go` | Happy path; GetVehicleLocations (no ID); 404; empty vehicleID → ValidationError |
| `shapes_test.go` | GetShape happy path; GetShapePolyline happy path; correct paths used |

All tests: table-driven with `t.Run`, `-race` compatible. Target ≥ 80% coverage.

---

## CI (`.github/workflows/ci.yml`)

```yaml
steps:
  - go vet ./...
  - go test -race -coverprofile=coverage.out ./...
  - go tool cover -func=coverage.out
  - staticcheck ./...
```

---

## Implementation Notes

- `X-Api-Key` header — NOT a query param. Changed from v2.
- `/shapes/{id}` vs `/shape/{id}/polyline` — match upstream exactly, comment it.
- All optional/nullable fields use pointer types. Never zero value for "absent".
- All path params go through `url.PathEscape`.
- `require(field, s)` at top of every method with a required path param.
- `any`-typed fields: use `internal/coerce`, never raw type assertions.

---

## README Outline

1. Badges: Go | pkg.go.dev | CI | Coverage | License
2. One-liner
3. `go get github.com/<handle>/cumtd-go`
4. Quick start (~15 lines, GetDepartures)
5. Full method reference table
6. Error handling with errors.As
7. Contributing
8. License (Apache-2.0)

---

## Release Checklist

- [ ] All 20 endpoints implemented
- [ ] All test files passing with -race
- [ ] go vet ./... clean
- [ ] staticcheck ./... clean
- [ ] Coverage ≥ 80%
- [ ] All exported symbols have godoc comments
- [ ] All CLAUDE.md files written
- [ ] README complete with badges
- [ ] `git tag v0.1.0 && git push origin v0.1.0`
- [ ] Visit pkg.go.dev/github.com/<handle>/cumtd-go to trigger indexing