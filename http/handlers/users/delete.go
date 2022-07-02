package users

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/lib"
)

type deleteAccountRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (h *handler) Delete(ctx *gin.Context) {
	var req deleteAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	err := h.service.DeleteUser(ctx, req.Username)
	if err != nil {
		// TODO: Error if the row that has been deleted still return no error
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, lib.Response("success", "user has been deleted", nil))
}
