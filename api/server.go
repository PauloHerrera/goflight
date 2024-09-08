package api

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}

	router := gin.Default()

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
