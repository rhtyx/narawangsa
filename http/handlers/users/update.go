package users

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/middleware"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

type updateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (h *handler) Update(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	arg := postgres.UpdateUserParams{
		Name:      req.Name,
		Email:     req.Email,
		UpdatedAt: time.Now(),
		Username:  *authPayload.Username,
	}

	user, err := h.service.UpdateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, lib.Response("success", "user has been updated", user, nil))
}
