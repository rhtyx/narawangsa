package booklists

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/lib"
)

func (h *handler) Delete(ctx *gin.Context) {
	bookId, err := strconv.Atoi(ctx.Param("book_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}
	err = h.service.DeleteBookList(ctx, int64(bookId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "booklist has been deleted", nil, nil))
}
