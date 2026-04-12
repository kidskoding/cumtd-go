# examples/

Usage examples for the cumtd-go SDK. Language-agnostic — Go preferred but not required.

## Examples

### `departure-board/`

Minimal example — fetch and print upcoming departures for a stop to stdout.

**Purpose:** Copy-paste starting point. Shows the full request/response cycle in the fewest lines.

**What it demonstrates:**
- Authenticating with `X-Api-Key` header
- Calling `GET /stops/{id}/departures`
- Handling errors
- Printing structured output

**Config:** `CUMTD_API_KEY` + `CUMTD_STOP_ID` env vars

**Rules:**
- Keep short and readable — under 80 lines
- No non-essential dependencies
- Comments explain each step for readers new to the API

### `agent-tool/`

Wraps CUMTD API endpoints as LLM tool definitions for use with an agent.

**Purpose:** Show how to expose transit data to an LLM. Demonstrates the tool-calling pattern without tying to any specific LLM SDK.

**Tools implemented:**
| Tool | Endpoint |
|------|----------|
| `search_stops` | `GET /stops/search` |
| `get_departures` | `GET /stops/{id}/departures` |
| `get_vehicle_locations` | `GET /vehicles/locations` |

**Config:** `CUMTD_API_KEY` env var

**Rules:**
- No LLM SDK dependency — tool definitions are plain data structures
- Generic dispatcher pattern, copy-paste friendly
- Comments explain the tool pattern for readers new to agent development

## Adding a new example

1. Create `examples/<name>/` directory
2. Implement in any language that can make HTTP requests
3. Add `CUMTD_API_KEY` env var support
4. Keep it focused on one concept
5. No build artifacts committed
