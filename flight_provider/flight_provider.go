package flightprovider

import (
	"log"
	"math"
	"sort"
)

// Departure: from -> to
// Return: to -> from
type FlightDirection string

const (
	Departure FlightDirection = "departure"
	Return    FlightDirection = "return"
)

type FlightType int

const (
	OneWay FlightType = iota
	RoundTrip
)

type SearchParams struct {
	DepartureAirport string     `json:"departure_airport"`
	DepartureDate    string     `json:"departure_date"`
	ReturnAirport    string     `json:"return_airport"`
	ReturnDate       string     `json:"return_date"`
	FlightType       FlightType `json:"type"`
}

type Flight struct {
	Direction     FlightDirection  `json:"flight_direction"`
	RegularPrice  float64          `json:"regular_price"`
	BoardingFee   float64          `json:"boarding_fee"`
	AirlineFee    float64          `json:"airline_fee"`
	AmparoFee     float64          `json:"amparo_fee"`
	FinalPrice    float64          `json:"final_price"`
	DiscountRate  int              `json:"discount_rate"`
	TotalDuration int              `json:"duration"`
	Airline       string           `json:"airline"`
	Airlinelogo   string           `json:"airline_logo"`
	Segments      []*FlightSegment `json:"segments"`
}

type FlightSegment struct {
	FlightNumber     string `json:"flight_number"`
	DepartureDate    string `json:"departure_date"`
	DepartureAirport string `json:"departure_airport"`
	ArrivalDate      string `json:"arrival_date"`
	ArrivalAirport   string `json:"arrival_airport"`
	Duration         int    `json:"duration"`
	Airline          string `json:"airline"`
}

func FlightsWithDiscount(searchArgs SearchParams) (flights []*Flight, err error) {
	results := make(chan []*Flight, 2)

	go func() {
		// Departure flight provider: from -> to
		args := getFlightParams{
			DepartureAirport:   searchArgs.DepartureAirport,
			DepartureDate:      searchArgs.DepartureDate,
			DestinationAirport: searchArgs.ReturnAirport,
		}

		flightResult, e := provideFlights(args, Departure)
		results <- flightResult
		if err != nil {
			err = e
		}
	}()

	go func() {
		// Return flight provider: to -> from (if necessary)
		if searchArgs.FlightType == RoundTrip {
			args := getFlightParams{
				DepartureAirport:   searchArgs.ReturnAirport,
				DepartureDate:      searchArgs.ReturnDate,
				DestinationAirport: searchArgs.DepartureAirport,
			}

			flightResult, e := provideFlights(args, Return)
			results <- flightResult
			if err != nil {
				err = e
			}
		} else {
			results <- nil
		}
	}()

	for i := 0; i < 2; i++ {
		flightResults := <-results
		if flightResults != nil {
			flights = append(flights, flightResults...)
		}
	}

	return flights, err
}

func provideFlights(args getFlightParams, fd FlightDirection) (flights []*Flight, err error) {
	results, err := GetFlights(args)
	if err != nil {
		log.Fatal("Flight api search error:", err)
		return nil, err
	}

	flights, err = proccessResult(results, fd)
	if err != nil {
		log.Fatal("Flight discount calculation error:", err)
		return nil, err
	}

	return
}

func proccessResult(resultFlights []interface{}, fd FlightDirection) (flights []*Flight, err error) {
	for _, flight := range resultFlights {
		flightMap := flight.(map[string]interface{})

		var segments []*FlightSegment

		for _, flight := range flightMap["flights"].([]interface{}) {
			flightItem := flight.(map[string]interface{})

			segments = append(segments, &FlightSegment{
				FlightNumber:     flightItem["flight_number"].(string),
				Airline:          flightItem["airline"].(string),
				Duration:         int(flightItem["duration"].(float64)),
				DepartureDate:    flightItem["departure_airport"].(map[string]interface{})["time"].(string),
				DepartureAirport: flightItem["departure_airport"].(map[string]interface{})["name"].(string),
				ArrivalDate:      flightItem["arrival_airport"].(map[string]interface{})["time"].(string),
				ArrivalAirport:   flightItem["arrival_airport"].(map[string]interface{})["name"].(string),
			})
		}

		calculatedPrice := PriceCalculator(flightMap["price"].(float64), segments[0].Airline)

		flights = append(flights, &Flight{
			Direction:     fd,
			RegularPrice:  flightMap["price"].(float64),
			BoardingFee:   calculatedPrice.BoardingFee,
			AmparoFee:     calculatedPrice.AmparoFee,
			AirlineFee:    calculatedPrice.AirlineFee,
			FinalPrice:    calculatedPrice.FinalPrice,
			DiscountRate:  calculatedPrice.DiscountRate,
			TotalDuration: int(flightMap["total_duration"].(float64)),
			Airline:       segments[0].Airline, // For now
			Airlinelogo:   flightMap["airline_logo"].(string),
			Segments:      segments,
		})
	}

	// Sort the list from the cheapest flight and return just the first five options
	sort.Slice(flights, func(i, j int) bool {
		return flights[i].FinalPrice < flights[j].FinalPrice
	})

	limitFlights := int(math.Min(float64(len(flights)), 5))
	return flights[:limitFlights], err
}
