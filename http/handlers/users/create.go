package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/lib"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *handler) Create(ctx *gin.Context, service users.IUsers) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	hashedPassword, err := lib.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	arg := postgres.CreateUserParams{
		Name:     req.Name,
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	err = service.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, lib.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "user has been created", nil))
}
