package categories

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/lib"
)

func (h *handler) Delete(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	err = h.service.DeleteCategory(ctx, int64(categoryId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, lib.Response("success", "category has been deleted", nil, nil, nil))
}
