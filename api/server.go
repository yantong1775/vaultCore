package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccounts)

	server.router.POST("/transfers", server.createTransfer)
	server.router.GET("/transfers/:id", server.getTransfer)
	server.router.GET("/transfers", server.listTransfers)

	server.router.POST("/users", server.createUser)

	return server
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
