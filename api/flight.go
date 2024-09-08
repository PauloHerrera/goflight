package api

import (
	"net/http"

	f "gihub.com/pauloherrera/goflight/flight_provider"
	"github.com/gin-gonic/gin"
)

type getFlightRequest struct {
	DepartureAirport string `form:"departure_airport" binding:"required"`
	DepartureDate    string `form:"departure_date" binding:"required"`
	ReturnAirport    string `form:"return_airport"`
	ReturnDate       string `form:"return_date"`
	FlightType       int    `form:"flight_type" binding:"required,oneof=1 2"`
}

// Return the available flights with discount
func (server *Server) GetFlights(ctx *gin.Context) {
	var req getFlightRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := f.SearchParams{
		DepartureID:   req.DepartureAirport,
		DepartureDate: req.DepartureDate,
		ReturnID:      req.ReturnAirport,
		ReturnDate:    req.ReturnDate,
		FlightType:    req.FlightType,
	}

	flightList, err := f.FlightsWithDiscount(arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, flightList)
}

// TODO : POST FLIGHTS
type flightPreferencesRequest struct {
	FlightID string `json:"flight_id" binding:"required"`
}

// Stores user flight preferences
func (server *Server) PostFlights(ctx *gin.Context) {
	var req flightPreferencesRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req)
}
