package users

import (
	"github.com/rhtyx/narawangsa/internal/domain/authentications"
	"github.com/rhtyx/narawangsa/internal/domain/userlevels"
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

type handler struct {
	config                lib.Config
	token                 token.Maker
	service               users.IUsers
	userLevelService      userlevels.IUserLevels
	authenticationService authentications.IAuthentications
}

func NewHandler(service users.IUsers, userLevelService userlevels.IUserLevels, authenticationService authentications.IAuthentications, token token.Maker, config lib.Config) *handler {
	return &handler{
		service:               service,
		userLevelService:      userLevelService,
		authenticationService: authenticationService,
		token:                 token,
		config:                config,
	}
}
