// vehicle-tracker polls all active MTD bus locations and prints a live table.
//
// Refreshes every 30 seconds. Press Ctrl-C to stop.
//
// Usage:
//
//	CUMTD_API_KEY=your_key go run .
//	CUMTD_API_KEY=your_key CUMTD_INTERVAL=10s go run .
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/cumtd/types"
)

func main() {
	apiKey := os.Getenv("CUMTD_API_KEY")
	if apiKey == "" {
		log.Fatal("CUMTD_API_KEY not set")
	}

	interval := 30 * time.Second
	if raw := os.Getenv("CUMTD_INTERVAL"); raw != "" {
		d, err := time.ParseDuration(raw)
		if err != nil {
			log.Fatalf("invalid CUMTD_INTERVAL: %v", err)
		}
		interval = d
	}

	client := cumtd.New(apiKey)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	fmt.Printf("Tracking MTD vehicles every %s. Ctrl-C to stop.\n\n", interval)
	poll(ctx, client)

	for {
		select {
		case <-ticker.C:
			poll(ctx, client)
		case <-ctx.Done():
			fmt.Println("\nstopped.")
			return
		}
	}
}

func poll(ctx context.Context, client *cumtd.Client) {
	locs, err := client.GetVehicleLocations(ctx)
	if err != nil {
		var rl *cumtd.RateLimitError
		var api *cumtd.APIError
		switch {
		case errors.Is(err, context.Canceled):
			return
		case errors.As(err, &rl):
			fmt.Printf("[%s] rate limited, retry after %s\n", now(), rl.RetryAfter)
		case errors.As(err, &api):
			fmt.Printf("[%s] API error %d: %s\n", now(), api.StatusCode, api.Body)
		default:
			fmt.Printf("[%s] error: %v\n", now(), err)
		}
		return
	}

	// Group by route short name for a compact display.
	byRoute := make(map[string][]types.VehicleLocation)
	for _, v := range locs {
		route := "unknown"
		if v.Route != nil {
			route = v.Route.ID
			if v.Route.ShortName != nil {
				route = *v.Route.ShortName
			}
		}
		byRoute[route] = append(byRoute[route], v)
	}

	routes := make([]string, 0, len(byRoute))
	for r := range byRoute {
		routes = append(routes, r)
	}
	sort.Strings(routes)

	fmt.Printf("[%s] %d active buses across %d routes\n", now(), len(locs), len(routes))
	for _, route := range routes {
		vehicles := byRoute[route]
		fmt.Printf("  %-14s %d bus(es)", route, len(vehicles))
		// Show coordinates for the first bus as a sample.
		if v := vehicles[0]; v.Location != nil {
			fmt.Printf("  — first: %.4f, %.4f", v.Location.Latitude, v.Location.Longitude)
		}
		fmt.Println()
	}
	fmt.Println()
}

func now() string {
	return time.Now().Format("15:04:05")
}
