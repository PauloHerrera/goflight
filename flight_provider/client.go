package flightprovider

import (
	"log"
	"strconv"

	g "github.com/serpapi/google-search-results-golang"

	"gihub.com/pauloherrera/goflight/util"
)

// Search for the available flights through external API integration
func getFlights(arg SearchParams) []interface{} {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
	}

	args := map[string]string{
		"engine":        "google_flights",
		"hl":            "pt-br",
		"gl":            "br",
		"departure_id":  arg.DepartureID,
		"arrival_id":    arg.ArrivalID,
		"outbound_date": arg.DepartureDate,
		"return_date":   arg.DepartureDate,
		"currency":      "BRL",
		"adults":        "1",
		"type":          strconv.Itoa(arg.FlightType),
		"travel_class":  "1",
	}

	search := g.NewGoogleSearch(args, config.FlightApiKey)
	results, err := search.GetJSON()

	if err != nil {
		log.Fatal("Integration error:", err)
	}

	combinedFlights := append(results["best_flights"].([]interface{}), results["best_flights"].([]interface{})...)

	return combinedFlights
}
