package users

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/lib"
)

type getUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (h *handler) Get(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	user, err := h.service.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "user has been returned", user))
}
