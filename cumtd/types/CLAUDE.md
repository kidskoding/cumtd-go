# cumtd/types — domain types

Pure type definitions. No logic, no imports outside stdlib.

## File map

| File | Types |
|------|-------|
| `common.go` | `Coordinates`, `TripDirection`, `DayType` |
| `routes.go` | `RouteGroup`, `Route` |
| `stops.go` | `StopBase`, `BoardingPoint`, `StopGroup`, `StopSearchResult`, `StopTime` |
| `departures.go` | `Departure`, `DepartureTrip`, `DepartureRoute` |
| `trips.go` | `Trip` |
| `vehicles.go` | `Vehicle`, `VehicleLocation`, `VehicleLocationTrip`, `VehicleConfiguration`, `VehicleType`, `VehiclePowertrainType` |
| `shapes.go` | `Shape`, `ShapePoint`, `ShapePolyline` |

## Rules

- Nullable/optional fields → pointer types (`*string`, `*Coordinates`)
- Fields the API types as `int | string` → `any` (use `internal/coerce` at call site)
- No methods on types — types are data only
- JSON tags must match upstream API field names exactly (camelCase)
- Do not add helper constructors or validation — callers handle that
