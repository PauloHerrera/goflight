package flightprovider

import (
	"log"
	"strconv"

	g "github.com/serpapi/google-search-results-golang"

	"gihub.com/pauloherrera/goflight/util"
)

// Search for available flights through external API integration
func GetFlights(arg SearchParams) (resultFlights []interface{}, err error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
		return nil, err
	}

	args := map[string]string{
		"engine":        "google_flights",
		"hl":            "pt-br",
		"gl":            "br",
		"departure_id":  arg.DepartureID,
		"arrival_id":    arg.ReturnID,
		"outbound_date": arg.DepartureDate,
		"return_date":   arg.ReturnDate,
		"currency":      "BRL",
		"adults":        "1",
		"type":          strconv.Itoa(arg.FlightType),
		"travel_class":  "1",
	}

	search := g.NewGoogleSearch(args, config.FlightApiKey)
	results, err := search.GetJSON()

	if err != nil {
		log.Fatal("Integration error:", err)
		return nil, err
	}

	if results["best_flights"] != nil {
		resultFlights = append(resultFlights, results["best_flights"].([]interface{})...)
	}

	if results["other_flights"] != nil {
		resultFlights = append(resultFlights, results["other_flights"].([]interface{})...)
	}

	return
}
