package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/yantong1775/vaultCore/db/sqlc"
)

// Server serves the http requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new server
func NewServer(store *db.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}

	server.router.POST("/accounts", server.createAccount)
	return server
}
