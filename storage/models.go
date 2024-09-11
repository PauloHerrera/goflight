package storage

import "time"

type UserFlight struct {
	UserID           string    `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
	DepartureFlights []Flight  `json:"departure_flights"`
	ReturnFlights    []Flight  `json:"return_flights"`
}

type Flight struct {
	Order         int             `json:"order"`
	TotalDuration int             `json:"duration"`
	Airline       string          `json:"airline"`
	Airlinelogo   string          `json:"airline_logo"`
	PriceInfo     PriceInfo       `json:"price_info"`
	FlightInfo    []FlightSegment `json:"flight_info"`
}

type PriceInfo struct {
	RegularPrice float64 `json:"regular_price"`
	BoardingFee  float64 `json:"boarding_fee"`
	AirlineFee   float64 `json:"airline_fee"`
	AmparoFee    float64 `json:"amparo_fee"`
	FinalPrice   float64 `json:"final_price"`
	DiscountRate int     `json:"discount_rate"`
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
