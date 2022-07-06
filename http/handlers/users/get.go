package users

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/middleware"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

type userResponse struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h *handler) Get(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	user, err := h.service.GetUser(ctx, *authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	response := userResponse{
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "user has been returned", response, nil))
}
