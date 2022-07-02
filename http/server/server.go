package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/handlers/base"
	u "github.com/rhtyx/narawangsa/http/handlers/users"
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/storage"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

type server struct {
	config lib.Config
	token  token.Maker
	store  storage.ExecTx
	router *gin.Engine
}

func New(store *postgres.Queries, storetx *postgres.TxInContext, config lib.Config, token token.Maker) *server {
	server := &server{store: storetx}
	router := gin.Default()

	userService := users.NewUserService(store, storetx)

	base := base.NewHandler()
	user := u.NewHandler(userService, token, config)

	router.GET("/ping", base.Ping)
	router.POST("/users/create", user.Create)
	// TODO: change the uri to validation
	router.GET("/users/:username", user.Get)
	router.DELETE("/users/delete/:username", user.Delete)
	router.POST("/users/update_password/:username")
	server.router = router
	return server
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}
