package api

import "github.com/gin-gonic/gin"

// createAccount request
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required, oneof=USD EUR"`
}

func (s *Server) createAccount(ctx *gin.Context) {
}
