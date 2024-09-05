package main

import (
	"fmt"
	"log"

	"gihub.com/pauloherrera/goflight/util"
	g "github.com/serpapi/google-search-results-golang"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
	}

	args := map[string]string{
		"engine":        "google_flights",
		"hl":            "pt-br",
		"gl":            "br",
		"departure_id":  "CGH",
		"arrival_id":    "GIG",
		"outbound_date": "2024-09-07",
		"return_date":   "2024-09-10",
		"currency":      "BRL",
		"adults":        "1",
		"type":          "1",
		"travel_class":  "1",
	}

	fmt.Println(config.FlightApiKey)
	search := g.NewGoogleSearch(args, config.FlightApiKey)
	results, err := search.GetJSON()

	if err != nil {
		fmt.Println("Integration error %w", err)
	}

	// Just printing the result to check the integration result
	fmt.Println(results)
}
