package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/handlers/base"
	u "github.com/rhtyx/narawangsa/http/handlers/users"
	"github.com/rhtyx/narawangsa/http/middleware"
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

	v1 := router.Group("/v1")
	{
		v1.POST("/signup", user.Create)
		v1.POST("/login", user.LoginUser)
		v1.GET("/logout")

		users := v1.Group("/users").Use(middleware.AuthMiddleware(token))
		{
			users.GET("/", user.Get)
			users.PUT("/")
			users.PATCH("/update_password")
			users.DELETE("/", user.Delete)
		}

		books := v1.Group("/books")
		{
			books.GET("/")
			books.GET("/:book_id")
			books.POST("/")
			books.PUT("/:book_id")
			books.DELETE("/:book_id")
		}

		categories := v1.Group("/categories")
		{
			categories.GET("/")
			categories.POST("/")
			categories.PUT("/:category_id")
			categories.DELETE("/:category_id")
		}

		booklists := v1.Group("/booklists").Use(middleware.AuthMiddleware(token))
		{
			booklists.GET("/")
			booklists.POST("/")
			booklists.DELETE("/:book_id")
		}

		readConfirmations := v1.Group("/read_confirmations").Use(middleware.AuthMiddleware(token))
		{
			readConfirmations.GET("/")
			readConfirmations.POST("/")
		}

		userLevels := v1.Group("/user_levels").Use(middleware.AuthMiddleware(token))
		{
			userLevels.PUT("/")
		}
	}
	// TODO: change the uri to validation
	server.router = router
	return server
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}
