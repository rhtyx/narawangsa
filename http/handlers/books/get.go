package books

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rhtyx/narawangsa/lib"
)

func (h *handler) Get(ctx *gin.Context) {
	bookName, ok := ctx.GetQuery("book_name")

	if !ok {
		limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
			return
		}
		response, err := h.service.ListBooks(ctx, int32(limit))
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, lib.Response("success", "books have been returned", response, nil, nil))
		return
	}

	response, err := h.service.GetBook(ctx, bookName)
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

	ctx.JSON(http.StatusOK, lib.Response("success", "book has been returned", response, nil, nil))
}
