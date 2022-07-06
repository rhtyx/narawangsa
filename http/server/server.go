package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/app"
	"github.com/rhtyx/narawangsa/http/handlers/base"
	"github.com/rhtyx/narawangsa/http/handlers/booklists"
	"github.com/rhtyx/narawangsa/http/handlers/books"
	"github.com/rhtyx/narawangsa/http/handlers/categories"
	"github.com/rhtyx/narawangsa/http/handlers/categorybooks"
	"github.com/rhtyx/narawangsa/http/handlers/notifications"
	"github.com/rhtyx/narawangsa/http/handlers/readconfirmations"
	"github.com/rhtyx/narawangsa/http/handlers/userlevels"
	"github.com/rhtyx/narawangsa/http/handlers/users"
	"github.com/rhtyx/narawangsa/http/middleware"
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

	app := app.NewContainer(store, storetx)

	base := base.NewHandler()
	user := users.NewHandler(app.UsersService, app.UserLevelsService, token, config)
	userlevel := userlevels.NewHandler(app.UserLevelsService)
	category := categories.NewHandler(app.CategoriesService)
	book := books.Newhandler(app.BooksService)
	categorybook := categorybooks.NewHandler(app.CategoryBooksService)
	readconfirmation := readconfirmations.NewHandler(app.ReadConfirmationsService, app.BooklistsService)
	booklist := booklists.NewHandler(app.BooklistsService)
	notification := notifications.NewHandler(app.NotificationsService)

	router.GET("/ping", base.Ping)

	v1 := router.Group("/v1")
	{
		v1.POST("/signup", user.Create)
		v1.POST("/login", user.LoginUser)
		v1.GET("/logout")
		v1.POST("/sendnotifications", notification.Send)

		users := v1.Group("/users").Use(middleware.AuthMiddleware(token))
		{
			users.GET("/", user.Get)
			users.PUT("/", user.Update)
			users.PATCH("/updatepassword", user.UpdatePassword)
			users.DELETE("/", user.Delete)
		}

		books := v1.Group("/books")
		{
			books.GET("/", book.Get)
			books.POST("/", book.Create)
			books.PUT("/:book_id", book.Update)
			books.DELETE("/:book_id", book.Delete)
		}

		categories := v1.Group("/categories")
		{
			categories.GET("/", category.Get)
			categories.POST("/", category.Create)
			categories.PUT("/:category_id", category.Update)
			categories.DELETE("/:category_id", category.Delete)
		}

		categorybooks := v1.Group("/categorybooks")
		{
			categorybooks.POST("/", categorybook.Create)
			categorybooks.DELETE("/", categorybook.Delete)
		}

		booklists := v1.Group("/booklists").Use(middleware.AuthMiddleware(token))
		{
			booklists.GET("/", booklist.List)
			booklists.POST("/", booklist.Create)
			booklists.PUT("/", booklist.Update)
			booklists.DELETE("/:book_id", booklist.Delete)
		}

		readConfirmations := v1.Group("/readconfirmations").Use(middleware.AuthMiddleware(token))
		{
			readConfirmations.GET("/", readconfirmation.List)
			readConfirmations.POST("/", readconfirmation.Create)
		}

		userLevels := v1.Group("/userlevels").Use(middleware.AuthMiddleware(token))
		{
			userLevels.GET("/", userlevel.Get)
			userLevels.PUT("/", userlevel.Update)
		}
	}
	// TODO: change the uri to validation
	server.router = router
	return server
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}
