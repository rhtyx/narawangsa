package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/internal/storage"
)

type server struct {
	store  storage.ExecTx
	router *gin.Engine
}

func NewServer(store storage.ExecTx) *server {
	server := &server{store: store}
	router := gin.Default()

	server.router = router
	return server
}
