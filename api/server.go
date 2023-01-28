package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/libterty/g-micro/db/sqlc"
)

// Server servers HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts", server.ListAccounts)
	router.GET("/accounts/:id", server.GetAccountById)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
