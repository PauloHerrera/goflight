package api

import (
	"net/http"

	f "gihub.com/pauloherrera/goflight/flight_provider"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	if !validRequest(ctx, req) {
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

func validRequest(ctx *gin.Context, req getFlightRequest) bool {
	var fieldErrors = map[string]string{}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range validationErrors {
				fieldErrors[err.Field()] = err.Tag()
			}
		} else {
			fieldErrors["error"] = err.Error()
		}

		ctx.JSON(http.StatusBadRequest, errorResponseList(fieldErrors))

		return false
	}

	return true
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
