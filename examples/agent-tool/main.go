// agent-tool exposes CUMTD API endpoints as LLM tool definitions.
//
// This example shows how to wrap SDK methods as generic tool schemas that any
// LLM agent (OpenAI, Anthropic, Gemini, etc.) can call. The tool definitions
// and dispatcher are plain Go — no LLM SDK required.
//
// Usage:
//
//	CUMTD_API_KEY=your_key go run .
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kidskoding/cumtd-go/cumtd"
)

// Tool describes a single LLM-callable function using the standard JSON Schema
// shape understood by OpenAI, Anthropic, and Gemini tool-calling APIs.
type Tool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
}

// Tools returns all tool definitions. Pass this slice to your LLM's
// tools/functions field.
func Tools() []Tool {
	return []Tool{
		{
			Name:        "search_stops",
			Description: "Search for MTD bus stops by name or code. Returns stop IDs, names, and locations.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query": map[string]any{
						"type":        "string",
						"description": "Stop name or code to search for, e.g. \"Illinois Terminal\" or \"STOP1\"",
					},
				},
				"required": []string{"query"},
			},
		},
		{
			Name:        "get_departures",
			Description: "Get upcoming bus departures for a stop. Returns route, headsign, and minutes until departure. Use search_stops first if you only have a stop name.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"stop_id": map[string]any{
						"type":        "string",
						"description": "MTD stop ID (e.g. \"STOP1\"). Obtain from search_stops.",
					},
					"routes": map[string]any{
						"type":        "string",
						"description": "Optional comma-separated route IDs to filter (e.g. \"ILLINI,GOLDLINE\"). Omit for all routes.",
					},
				},
				"required": []string{"stop_id"},
			},
		},
		{
			Name:        "get_vehicle_locations",
			Description: "Get real-time GPS locations for all active MTD buses. Returns vehicle ID, route, direction, and coordinates.",
			Parameters: map[string]any{
				"type":       "object",
				"properties": map[string]any{},
			},
		},
	}
}

// Dispatch executes a tool call by name with the given JSON arguments.
// Returns the result as a JSON string suitable for feeding back to the LLM.
func Dispatch(ctx context.Context, client *cumtd.Client, name string, argsJSON []byte) (string, error) {
	switch name {
	case "search_stops":
		var args struct {
			Query string `json:"query"`
		}
		if err := json.Unmarshal(argsJSON, &args); err != nil {
			return "", err
		}
		results, err := client.SearchStops(ctx, args.Query)
		if err != nil {
			return "", err
		}
		b, _ := json.Marshal(results)
		return string(b), nil

	case "get_departures":
		var args struct {
			StopID string `json:"stop_id"`
			Routes string `json:"routes"`
		}
		if err := json.Unmarshal(argsJSON, &args); err != nil {
			return "", err
		}
		var opts *cumtd.GetDeparturesOptions
		if args.Routes != "" {
			opts = &cumtd.GetDeparturesOptions{Routes: args.Routes}
		}
		deps, err := client.GetDepartures(ctx, args.StopID, opts)
		if err != nil {
			return "", err
		}
		b, _ := json.Marshal(deps)
		return string(b), nil

	case "get_vehicle_locations":
		locs, err := client.GetVehicleLocations(ctx)
		if err != nil {
			return "", err
		}
		b, _ := json.Marshal(locs)
		return string(b), nil

	default:
		return "", fmt.Errorf("unknown tool: %s", name)
	}
}

func main() {
	apiKey := os.Getenv("CUMTD_API_KEY")
	if apiKey == "" {
		log.Fatal("CUMTD_API_KEY not set")
	}

	client := cumtd.New(apiKey)
	ctx := context.Background()

	// Print tool definitions — paste these into your LLM API call.
	defs, _ := json.MarshalIndent(Tools(), "", "  ")
	fmt.Println("=== Tool Definitions ===")
	fmt.Println(string(defs))

	// Demo: simulate an agent calling search_stops then get_departures.
	fmt.Println("\n=== Demo Dispatch ===")

	searchResult, err := Dispatch(ctx, client, "search_stops", []byte(`{"query":"Illinois Terminal"}`))
	if err != nil {
		var api *cumtd.APIError
		if errors.As(err, &api) {
			log.Fatalf("API error %d: %s", api.StatusCode, api.Body)
		}
		log.Fatal(err)
	}
	fmt.Printf("search_stops result (first 200 chars):\n%.200s...\n", searchResult)
}
