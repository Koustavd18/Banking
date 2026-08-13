package api

import (
	"net/http"

	db "github.com/Koustavd18/Banking/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//endpoints
	router.GET("/health", server.health)
	router.GET("account/:id", server.getAccount)
	router.GET("/accounts", server.getAllAcoounts)
	router.POST("/account", server.createAccount)

	server.router = router
	return server
}

func (server *Server) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"health": "healthy"})
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
