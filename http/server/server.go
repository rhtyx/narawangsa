package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/handlers/base"
	u "github.com/rhtyx/narawangsa/http/handlers/users"
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/storage"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type server struct {
	store  storage.ExecTx
	router *gin.Engine
}

func New(store *postgres.Queries, storetx *postgres.TxInContext) *server {
	server := &server{store: storetx}
	router := gin.Default()

	userService := users.NewUserService(store, storetx)

	base := base.NewHandler()
	user := u.NewHandler(userService)

	router.GET("/ping", base.Ping)
	router.POST("/users/create", user.Create)
	router.DELETE("/users/delete/:username", user.Delete)
	server.router = router
	return server
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}
