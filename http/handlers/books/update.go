package books

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/lib"
)

type updateBookRequest struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Year     string `json:"year" binding:"required"`
	Pages    int32  `json:"pages" binding:"required"`
	Synopsis string `json:"synopsis" binding:"required"`
}

func (h *handler) Update(ctx *gin.Context) {
	var req updateBookRequest

	bookId, err := strconv.Atoi(ctx.Param("book_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	arg := postgres.UpdateBookParams{
		Title:     req.Title,
		Author:    req.Author,
		Year:      req.Year,
		Pages:     req.Pages,
		Synopsis:  req.Synopsis,
		UpdatedAt: time.Now(),
		ID:        int64(bookId),
	}

	err = h.service.UpdateBook(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, lib.Response("success", "book has been updated", nil, nil, nil))
}
