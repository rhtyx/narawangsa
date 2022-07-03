package readconfirmations

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/lib"
)

type listReadConfirmationRequest struct {
	BookListID int64 `json:"book_list_id" binding:"required"`
	Limit      int32 `json:"limit" binding:"required"`
}

func (h *handler) List(ctx *gin.Context) {
	var req listReadConfirmationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}
	arg := postgres.ListReadConfirmationsParams{
		BookListID: req.BookListID,
		Limit:      req.Limit,
	}

	response, err := h.service.ListReadConfirmations(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "read confirmations has been returned", response, nil))
}
