package categories

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/lib"
)

func (h *handler) Get(ctx *gin.Context) {
	categoryName, ok := ctx.GetQuery("category_name")
	if !ok {
		limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "1"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
			return
		}
		offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
			return
		}

		var req = postgres.ListCategoriesParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		}

		response, err := h.service.ListCategories(ctx, req)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, lib.Response("success", "categories have been returned", response, nil, nil))
		return
	}

	response, err := h.service.GetCategory(ctx, categoryName)
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

	ctx.JSON(http.StatusOK, lib.Response("success", "category has been returned", response, nil, nil))
}
