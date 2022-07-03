package booklists

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

func (h *handler) Update(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	var req booklistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	arg := postgres.UpdateBookListParams{
		UserID:    authPayload.UserId,
		BookID:    req.BookID,
		IsRead:    req.IsRead,
		PagesRead: req.PagesRead,
		EndDate:   req.EndDate,
		UpdatedAt: time.Now(),
	}

	err := h.service.UpdateBookList(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "booklist has been updated", nil, nil))
}
