package flightprovider

import (
	"log"

	g "github.com/serpapi/google-search-results-golang"

	"gihub.com/pauloherrera/goflight/util"
)

type getFlightParams struct {
	DepartureAirport   string `json:"departure_airport"`
	DepartureDate      string `json:"departure_date"`
	DestinationAirport string `json:"destination_airport"`
}

// Default Search values:
// Country and Language: business req only works on Brazil and with real based currency
// Adults: business req always flights for one person.
// SearchType: Round trip=1,one way=2,multi city=3.
// Business req: search always for one way flights. A round trip should be two one way flight search.
// TravelClass: Economy=1,Premium=2,Business=3,First=4. Business req: always economy flights
const (
	Engine      = "google_flights"
	Language    = "pt-br"
	Country     = "br"
	Currency    = "BRL"
	Adults      = "1"
	SearchType  = "2"
	TravelClass = "1"
)

// Search for available flights through external API integration
func GetFlights(arg getFlightParams) (resultFlights []interface{}, err error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
		return nil, err
	}

	args := map[string]string{
		"engine":        Engine,
		"hl":            Language,
		"gl":            Country,
		"departure_id":  arg.DepartureAirport,
		"arrival_id":    arg.DestinationAirport,
		"outbound_date": arg.DepartureDate,
		"currency":      Currency,
		"adults":        Adults,
		"type":          SearchType,
		"travel_class":  TravelClass,
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
