package booklists

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/middleware"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

func (h *handler) List(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	response, err := h.service.ListBookList(ctx, authPayload.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "booklists have been returned", response, nil))
}
