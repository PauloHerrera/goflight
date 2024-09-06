package main

import (
	"encoding/json"
	"fmt"

	flightprovider "gihub.com/pauloherrera/goflight/flight_provider"
)

func main() {
	//Sample data for simulate search params
	arg := flightprovider.SearchParams{
		DepartureID:   "CGH",
		DepartureDate: "2024-09-07",
		ArrivalID:     "GIG",
		ArrivalDate:   "2024-09-12",
		FlightType:    1,
	}

	flightList := flightprovider.FlightsWithDiscount(arg)

	jsonData, err := json.Marshal(flightList)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
