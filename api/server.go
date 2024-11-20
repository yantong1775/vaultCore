package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/yantong1775/vaultCore/db/sqlc"
)

// Server serves the http requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new server
func NewServer(store db.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}

	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccounts)

	server.router.POST("/transfers", server.createTransfer)
	server.router.GET("/transfers/:id", server.getTransfer)
	server.router.GET("/transfers", server.listTransfers)

	server.router.POST("/entries", server.createEntry)
	server.router.GET("/entries/:id", server.getEntry)
	server.router.GET("/entries", server.listEntry)

	return server
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
