package flightprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPriceCalculator(t *testing.T) {
	t.Run("with known airline discount", func(t *testing.T) {
		regularPrice := 3000.0
		airline := "Gol"

		expected := &CalculatedPrice{
			BoardingFee:  30,
			AmparoFee:    59.40,
			AirlineFee:   594.00,
			FinalPrice:   683.40,
			DiscountRate: 78,
		}

		result := PriceCalculator(regularPrice, airline)

		require.Equal(t, expected, result)
	})

	t.Run("with unknown airline", func(t *testing.T) {
		regularPrice := 3000.0
		airline := "Gopher Airline"

		expected := &CalculatedPrice{
			BoardingFee: 30,
			AmparoFee:   297,
			AirlineFee:  2970,
			FinalPrice:  3297,
		}

		result := PriceCalculator(regularPrice, airline)

		//No discount should be applied
		//The boarding fee remains constant
		require.Equal(t, result.BoardingFee, expected.BoardingFee)
		//Amparo fee considers the usual 10% of regularPrice - boarding fee
		require.Equal(t, result.AmparoFee, expected.AmparoFee)
		//The airline fee remaing the same: regularPrice - boarding fee
		require.Equal(t, result.AirlineFee, expected.AirlineFee)
		//Since there is no discount, the final price is higher than regular price, considering Amparo's fee
		require.Equal(t, result.FinalPrice, expected.FinalPrice)
		//There is no discount, so the discount rate should be zero
		require.Zero(t, result.DiscountRate)

	})

	t.Run("with zero as price parameter", func(t *testing.T) {
		regularPrice := -5.0
		airline := "Latam"

		expected := &CalculatedPrice{
			BoardingFee:  30,
			AmparoFee:    0,
			AirlineFee:   0,
			FinalPrice:   30,
			DiscountRate: 0,
		}

		result := PriceCalculator(regularPrice, airline)

		assert.Equal(t, expected, result)
	})
}
