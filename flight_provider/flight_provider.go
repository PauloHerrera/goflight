package flightprovider

type SearchParams struct {
	DepartureID   string `json:"departure_id"`
	DepartureDate string `json:"departure_date"`
	ArrivalID     string `json:"arrival_id"`
	ArrivalDate   string `json:"arrival_date"`
	FlightType    int    `json:"flight_type"`
}

type Flight struct {
	RegularPrice  float64          `json:"regular_price"`
	BoardingFee   float64          `json:"boarding_fee"`
	AmparoFee     float64          `json:"amparo_fee"`
	FinalPrice    float64          `json:"final_price"`
	TotalDuration int              `json:"duration"`
	Airline       string           `json:"airline"`
	Airlinelogo   string           `json:"airline_logo"`
	Segments      []*FlightSegment `json:"segments"`
}

type FlightSegment struct {
	ID               string `json:"id"`
	DepartureDate    string `json:"departure_date"`
	DepartureAirport string `json:"departure_airport"`
	ArrivalDate      string `json:"arrival_date"`
	ArrivalAirport   string `json:"arrival_airport"`
	Duration         int    `json:"duration"`
	Airline          string `json:"airline"`
}

func FlightsWithDiscount(args SearchParams) []*Flight {
	var flights []*Flight

	results := getFlights(args)

	for _, flight := range results {
		flightMap := flight.(map[string]interface{})

		var segments []*FlightSegment

		for _, flight := range flightMap["flights"].([]interface{}) {
			flightItem := flight.(map[string]interface{})

			segments = append(segments, &FlightSegment{
				ID:               flightItem["flight_number"].(string),
				Airline:          flightItem["airline"].(string),
				DepartureDate:    flightItem["departure_airport"].(map[string]interface{})["time"].(string),
				DepartureAirport: flightItem["departure_airport"].(map[string]interface{})["name"].(string),
				ArrivalDate:      flightItem["arrival_airport"].(map[string]interface{})["time"].(string),
				ArrivalAirport:   flightItem["arrival_airport"].(map[string]interface{})["name"].(string),
				Duration:         int(flightItem["duration"].(float64)),
			})
		}

		//TODO: Flight discount calculator
		flights = append(flights, &Flight{
			RegularPrice:  flightMap["price"].(float64),
			BoardingFee:   flightMap["price"].(float64),
			AmparoFee:     flightMap["price"].(float64),
			FinalPrice:    flightMap["price"].(float64),
			TotalDuration: int(flightMap["total_duration"].(float64)),
			Airline:       segments[0].Airline, // For now
			Airlinelogo:   flightMap["airline_logo"].(string),
			Segments:      segments,
		})
	}

	return flights
}
