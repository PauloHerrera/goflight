package flightprovider

import "log"

type SearchParams struct {
	DepartureID   string `json:"departure_id"`
	DepartureDate string `json:"departure_date"`
	ReturnID      string `json:"return_id"`
	ReturnDate    string `json:"return_date"`
	FlightType    int    `json:"flight_type"`
}

type Flight struct {
	RegularPrice  float64          `json:"regular_price"`
	BoardingFee   float64          `json:"boarding_fee"`
	AmparoFee     float64          `json:"amparo_fee"`
	FinalPrice    float64          `json:"final_price"`
	DiscountRate  int              `json:"discount_rate"`
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

func FlightsWithDiscount(args SearchParams) (flights []*Flight, err error) {
	results, err := GetFlights(args)
	if err != nil {
		log.Fatal("Flight search error:", err)
		return nil, err
	}

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

		calculatedPrice := PriceCalculator(flightMap["price"].(float64), segments[0].Airline)

		flights = append(flights, &Flight{
			RegularPrice:  flightMap["price"].(float64),
			BoardingFee:   calculatedPrice.BoardingFee,
			AmparoFee:     calculatedPrice.AmparoFee,
			FinalPrice:    calculatedPrice.FinalPrice,
			DiscountRate:  calculatedPrice.DiscountRate,
			TotalDuration: int(flightMap["total_duration"].(float64)),
			Airline:       segments[0].Airline, // For now
			Airlinelogo:   flightMap["airline_logo"].(string),
			Segments:      segments,
		})
	}

	return
}
