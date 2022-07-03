package categorybooks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/lib"
)

type categoryBooksRequest struct {
	BookID     int64 `json:"book_id" binding:"required"`
	CategoryID int64 `json:"category_id" binding:"required"`
}

func (h *handler) Create(ctx *gin.Context) {
	var req categoryBooksRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	arg := postgres.CreateBookCategoryParams{
		BookID:     req.BookID,
		CategoryID: req.CategoryID,
	}

	err := h.service.CreateBookCategory(ctx, arg)
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

	ctx.JSON(http.StatusCreated, lib.Response("success", "books category has been created", nil, nil))
}
