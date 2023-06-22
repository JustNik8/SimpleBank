package api

import (
	db "SimpleBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer created a new HTTP server an setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router

	return server
}

// Start runs the HTTP server on a specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}