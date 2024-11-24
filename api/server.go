package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/yantong1775/vaultCore/db/sqlc"
	"github.com/yantong1775/vaultCore/token"
	"github.com/yantong1775/vaultCore/util"
)

// Server serves the http requests
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new server
func NewServer(Config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(Config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{
		config:     Config,
		store:      store,
		tokenMaker: tokenMaker,
		router:     gin.Default(),
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

// Set up router
func (server *Server) setupRouter() {
	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccounts)

	server.router.POST("/transfers", server.createTransfer)
	server.router.GET("/transfers/:id", server.getTransfer)
	server.router.GET("/transfers", server.listTransfers)

	server.router.POST("/users", server.createUser)
	server.router.POST("/users/login", server.loginUser)
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
