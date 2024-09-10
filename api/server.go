package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validAirport", validAirport)
		v.RegisterValidation("validDate", validDate)
	}

	router.GET("/flights", server.GetFlights)
	router.POST("/flights", server.PostFlights)

	server.router = router

	return server
}

// Starts and run the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func errorResponseList(err error) map[string]interface{} {
	var fieldErrors = map[string]string{}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			fieldErrors[err.Field()] = err.Tag()
		}
	} else {
		fieldErrors["error"] = err.Error()
	}

	return map[string]interface{}{
		"error": fieldErrors,
	}
}
