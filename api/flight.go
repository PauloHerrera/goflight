package api

import (
	"errors"
	"net/http"
	"time"

	f "gihub.com/pauloherrera/goflight/flight_provider"
	"gihub.com/pauloherrera/goflight/storage"
	"github.com/gin-gonic/gin"
)

type getFlightRequest struct {
	DepartureAirport string `form:"departure_airport" binding:"required,validAirport"`
	DepartureDate    string `form:"departure_date" binding:"required,validDate"`
	ReturnAirport    string `form:"return_airport" binding:"validAirport"`
	ReturnDate       string `form:"return_date" binding:"validDate"`
	FlightType       int    `form:"flight_type" binding:"required,oneof=1 2"`
}

// Return the available flights with discount
func (server *Server) GetFlights(ctx *gin.Context) {
	var req getFlightRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponseList(err))
		return
	}

	arg := f.SearchParams{
		DepartureAirport: req.DepartureAirport,
		DepartureDate:    req.DepartureDate,
		ReturnAirport:    req.ReturnAirport,
		ReturnDate:       req.ReturnDate,
		FlightType:       f.FlightType(req.FlightType),
	}

	flightList, err := f.FlightsWithDiscount(arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, flightList)
}

type flightPreferencesParams struct {
	Order         int                     `json:"order" binding:"required"`
	Direction     string                  `json:"flight_direction" binding:"required,oneof=departure return"`
	RegularPrice  float64                 `json:"regular_price"`
	BoardingFee   float64                 `json:"boarding_fee"`
	AirlineFee    float64                 `json:"airline_fee"`
	AmparoFee     float64                 `json:"amparo_fee"`
	FinalPrice    float64                 `json:"final_price"`
	DiscountRate  int                     `json:"discount_rate"`
	TotalDuration int                     `json:"duration"`
	Airline       string                  `json:"airline"`
	Airlinelogo   string                  `json:"airline_logo"`
	Segments      []*flightSegmentsParams `json:"segments"`
}

type flightSegmentsParams struct {
	FlightNumber     string `json:"flight_number"`
	DepartureDate    string `json:"departure_date"`
	DepartureAirport string `json:"departure_airport"`
	ArrivalDate      string `json:"arrival_date"`
	ArrivalAirport   string `json:"arrival_airport"`
	Duration         int    `json:"duration"`
	Airline          string `json:"airline"`
}

// Stores user flight preferences
func (server *Server) PostFlights(ctx *gin.Context) {
	var req []flightPreferencesParams

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//Gets the user_id through query parameter, since we still don't have auth
	userID := ctx.Query("user_id")
	if userID == "" {
		err := errors.New("user id not found")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args, err := buildFlight(userID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	result, err := server.db.PostUserFlights(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func buildFlight(userID string, flightPreferences []flightPreferencesParams) (flight *storage.UserFlight, err error) {
	var departureFlights []storage.Flight
	var returnFlights []storage.Flight

	for _, f := range flightPreferences {
		var flightSegments []storage.FlightSegment

		for _, s := range f.Segments {
			flightSegments = append(flightSegments, storage.FlightSegment{
				FlightNumber:     s.FlightNumber,
				DepartureDate:    s.DepartureDate,
				DepartureAirport: s.DepartureAirport,
				ArrivalDate:      s.ArrivalDate,
				ArrivalAirport:   s.ArrivalAirport,
				Duration:         s.Duration,
				Airline:          s.Airline,
			})
		}

		flightItem := storage.Flight{
			Order:         f.Order,
			TotalDuration: f.TotalDuration,
			Airline:       f.Airline,
			Airlinelogo:   f.Airlinelogo,
			PriceInfo: storage.PriceInfo{
				RegularPrice: f.RegularPrice,
				BoardingFee:  f.BoardingFee,
				AirlineFee:   f.AirlineFee,
				AmparoFee:    f.AmparoFee,
				FinalPrice:   f.FinalPrice,
				DiscountRate: f.DiscountRate,
			},
			FlightInfo: flightSegments,
		}

		if f.Direction == "departure" {
			departureFlights = append(departureFlights, flightItem)
		}

		if f.Direction == "return" {
			returnFlights = append(returnFlights, flightItem)
		}
	}

	flight = &storage.UserFlight{
		UserID:           userID,
		CreatedAt:        time.Now(),
		DepartureFlights: departureFlights,
		ReturnFlights:    returnFlights,
	}

	return
}
