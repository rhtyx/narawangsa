package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/handlers/base"
	bl "github.com/rhtyx/narawangsa/http/handlers/booklists"
	b "github.com/rhtyx/narawangsa/http/handlers/books"
	c "github.com/rhtyx/narawangsa/http/handlers/categories"
	cb "github.com/rhtyx/narawangsa/http/handlers/categorybooks"
	rc "github.com/rhtyx/narawangsa/http/handlers/readconfirmations"
	ul "github.com/rhtyx/narawangsa/http/handlers/userlevels"
	u "github.com/rhtyx/narawangsa/http/handlers/users"
	"github.com/rhtyx/narawangsa/http/middleware"
	"github.com/rhtyx/narawangsa/internal/domain/booklists"
	"github.com/rhtyx/narawangsa/internal/domain/books"
	"github.com/rhtyx/narawangsa/internal/domain/categories"
	"github.com/rhtyx/narawangsa/internal/domain/categorybooks"
	"github.com/rhtyx/narawangsa/internal/domain/readconfirmations"
	"github.com/rhtyx/narawangsa/internal/domain/userlevels"
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

	usersService := users.NewUserService(store, storetx)
	userLevelsService := userlevels.NewUserLevelsService(store, storetx)
	categoriesService := categories.NewCategoriesService(store, storetx)
	booksService := books.NewBooksService(store, storetx)
	categoryBooksService := categorybooks.NewCategoryBooksService(store, storetx)
	readConfirmationsService := readconfirmations.NewReadConfirmationsService(store, storetx)
	booklistsService := booklists.NewBookListsService(store, storetx)

	base := base.NewHandler()
	user := u.NewHandler(usersService, userLevelsService, token, config)
	userlevel := ul.NewHandler(userLevelsService)
	category := c.NewHandler(categoriesService)
	book := b.Newhandler(booksService)
	categorybook := cb.NewHandler(categoryBooksService)
	readconfirmation := rc.NewHandler(readConfirmationsService, booklistsService)
	booklist := bl.NewHandler(booklistsService)

	router.GET("/ping", base.Ping)

	v1 := router.Group("/v1")
	{
		v1.POST("/signup", user.Create)
		v1.POST("/login", user.LoginUser)
		v1.GET("/logout")

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
