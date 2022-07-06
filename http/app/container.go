package app

import (
	"github.com/rhtyx/narawangsa/internal/domain/booklists"
	"github.com/rhtyx/narawangsa/internal/domain/books"
	"github.com/rhtyx/narawangsa/internal/domain/categories"
	"github.com/rhtyx/narawangsa/internal/domain/categorybooks"
	"github.com/rhtyx/narawangsa/internal/domain/notifications"
	"github.com/rhtyx/narawangsa/internal/domain/readconfirmations"
	"github.com/rhtyx/narawangsa/internal/domain/userlevels"
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

type container struct {
	UsersService             users.IUsers
	UserLevelsService        userlevels.IUserLevels
	CategoriesService        categories.ICategories
	BooksService             books.IBooks
	CategoryBooksService     categorybooks.ICategoryBooks
	ReadConfirmationsService readconfirmations.IReadConfirmations
	BooklistsService         booklists.IBooklists
	NotificationsService     notifications.INotifications
}

func NewContainer(store *postgres.Queries, storetx *postgres.TxInContext) *container {
	return &container{
		UsersService:             users.NewUserService(store, storetx),
		UserLevelsService:        userlevels.NewUserLevelsService(store, storetx),
		CategoriesService:        categories.NewCategoriesService(store, storetx),
		BooksService:             books.NewBooksService(store, storetx),
		CategoryBooksService:     categorybooks.NewCategoryBooksService(store, storetx),
		ReadConfirmationsService: readconfirmations.NewReadConfirmationsService(store, storetx),
		BooklistsService:         booklists.NewBookListsService(store, storetx),
		NotificationsService:     notifications.NewNotificationsService(),
	}
}
