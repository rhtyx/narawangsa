package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/lib"
)

type CreateBookRequest struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Year     string `json:"year" binding:"required"`
	Pages    int32  `json:"pages" binding:"required"`
	Synopsis string `json:"synopsis" binding:"required"`
}

func (h *handler) Create(ctx *gin.Context) {
	var req CreateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	arg := postgres.CreateBookParams{
		Title:    req.Title,
		Author:   req.Author,
		Year:     req.Year,
		Pages:    req.Pages,
		Synopsis: req.Synopsis,
	}

	err := h.service.CreateBook(ctx, arg)
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

	ctx.JSON(http.StatusCreated, lib.Response("success", "book has been created", nil, nil, nil))
}
