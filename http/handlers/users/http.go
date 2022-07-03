package users

import (
	"github.com/rhtyx/narawangsa/internal/domain/userlevels"
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

type handler struct {
	config           lib.Config
	token            token.Maker
	service          users.IUsers
	userLevelService userlevels.IUserLevels
}

func NewHandler(service users.IUsers, userLevelService userlevels.IUserLevels, token token.Maker, config lib.Config) *handler {
	return &handler{
		service:          service,
		userLevelService: userLevelService,
		token:            token,
		config:           config,
	}
}
