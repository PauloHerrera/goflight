package flightprovider

import (
	"math"
)

type CalculatedPrice struct {
	BoardingFee  float64 `json:"boarding_fee"`
	AmparoFee    float64 `json:"amparo_fee"`
	AirlineFee   float64 `json:"airline_fee"`
	FinalPrice   float64 `json:"final_price"`
	DiscountRate int     `json:"discount_rate"`
}

const (
	amparoFeePercent = 10
	boardingFee      = 30
)

var airlineDiscountMap = map[string]float64{
	"Gol":   0.80,
	"Latam": 0.30,
	"Azul":  0.00,
}

func PriceCalculator(regularPrice float64, airline string) *CalculatedPrice {
	// Check the airline for discounts. Applies no discount for a unknown airline
	airlineDiscount, ok := airlineDiscountMap[airline]
	if !ok {
		airlineDiscount = 0.0
	}

	airlineFinalPrice := (1 - airlineDiscount) * math.Max(0, (regularPrice-boardingFee))
	amparoFee := float64(amparoFeePercent) / 100 * airlineFinalPrice
	finalPrice := airlineFinalPrice + amparoFee + boardingFee

	var discountRate int
	if regularPrice > 0 && airlineDiscount > 0 {
		discountRate = int(math.Ceil((1 - finalPrice/regularPrice) * 100))
	} else {
		discountRate = 0
	}

	calculatedPrice := CalculatedPrice{
		BoardingFee:  math.Round(boardingFee*100) / 100,
		AmparoFee:    math.Round(amparoFee*100) / 100,
		AirlineFee:   math.Round(airlineFinalPrice*100) / 100,
		FinalPrice:   math.Round(finalPrice*100) / 100,
		DiscountRate: discountRate,
	}

	return &calculatedPrice
}
