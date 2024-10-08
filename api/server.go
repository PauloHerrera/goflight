package api

import (
	"net/http"

	"gihub.com/pauloherrera/goflight/storage"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	db     *storage.Worker
	router *gin.Engine
}

func NewServer(mongodb *storage.Worker) *Server {
	server := &Server{
		db: mongodb,
	}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validAirport", validAirport)
		v.RegisterValidation("validDate", validDate)
	}

	router.GET("/health", server.healthCheck)
	router.GET("/flights", server.GetFlights)
	router.PUT("/flights", server.PutFlights)

	server.router = router

	return server
}

// Starts and run the server
func (server *Server) Start(port string) error {
	return server.router.Run(":" + port)
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

func (server *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is ready"})
}
