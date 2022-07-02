package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/lib"
)

type updatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=8"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
	Username    string `json:"username" binding:"required"`
}

// TODO: unfinished handler update password
func (h *handler) UpdatePassword(ctx *gin.Context) {
	var req updatePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}
}
