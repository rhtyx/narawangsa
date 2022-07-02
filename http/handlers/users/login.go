package users

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/lib"
)

type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=8"`
}

// TODO: unfinished Login User handler
func (h *handler) LoginUser(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	user, err := h.service.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	err = lib.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, lib.ErrorResponse(err))
		return
	}

	accessToken, err := h.token.CreateToken(user.Username)
}
