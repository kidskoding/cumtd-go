// departure-board prints upcoming departures for a stop to stdout.
//
// Usage:
//
//	CUMTD_API_KEY=your_key CUMTD_STOP_ID=STOP1 go run .
//
// Get a key at https://developer.mtd.org.
// Find stop IDs via the MTD system map or SearchStops.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kidskoding/cumtd-go/cumtd"
)

func main() {
	apiKey := os.Getenv("CUMTD_API_KEY")
	stopID := os.Getenv("CUMTD_STOP_ID")
	if apiKey == "" {
		log.Fatal("CUMTD_API_KEY not set")
	}
	if stopID == "" {
		log.Fatal("CUMTD_STOP_ID not set")
	}

	client := cumtd.New(apiKey)

	deps, err := client.GetDepartures(context.Background(), stopID, nil)
	if err != nil {
		var rl *cumtd.RateLimitError
		var api *cumtd.APIError
		switch {
		case errors.As(err, &rl):
			log.Fatalf("rate limited, retry after %s", rl.RetryAfter)
		case errors.As(err, &api):
			log.Fatalf("API error %d: %s", api.StatusCode, api.Body)
		default:
			log.Fatal(err)
		}
	}

	if len(deps) == 0 {
		fmt.Println("no upcoming departures")
		return
	}

	fmt.Printf("Departures — stop %s\n\n", stopID)
	for _, d := range deps {
		// Route short name falls back to route ID when absent.
		route := ""
		if d.Route != nil {
			route = d.Route.ID
			if d.Route.ShortName != nil {
				route = *d.Route.ShortName
			}
		}

		headsign := ""
		if d.Headsign != nil {
			headsign = " → " + *d.Headsign
		}

		realtime := ""
		if d.IsRealTime {
			realtime = " [live]"
		}

		fmt.Printf("  %-10s  %-12s%s%s\n", d.DepartsIn, route, headsign, realtime)
	}
}
