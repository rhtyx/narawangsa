package categories

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/lib"
)

type updateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *handler) Update(ctx *gin.Context) {
	var req updateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
			return
		}
	}

	var categoryId, err = strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	arg := postgres.UpdateCategoryParams{
		Name:      req.Name,
		ID:        int64(categoryId),
		UpdatedAt: time.Now(),
	}

	err = h.service.UpdateCategory(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "category has been updated", nil, nil))
}
