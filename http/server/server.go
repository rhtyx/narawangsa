package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/handlers/base"
	"github.com/rhtyx/narawangsa/internal/storage"
)

type server struct {
	store  storage.ExecTx
	router *gin.Engine
}

func New(store storage.ExecTx) *server {
	server := &server{store: store}
	router := gin.Default()

	base := base.NewHandler()

	router.GET("/ping", base.Ping)
	server.router = router
	return server
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}
