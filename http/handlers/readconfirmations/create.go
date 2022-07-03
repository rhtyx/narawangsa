package readconfirmations

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/http/middleware"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

type readConfirmationsRequest struct {
	BookId    int64 `json:"book_id" binding:"required"`
	Pagesread int32 `json:"pages_read" binding:"required"`
}

func (h *handler) Create(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	var req readConfirmationsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	booklists, err := h.booklistService.ListBookList(ctx, authPayload.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	arg := postgres.CreateReadConfirmationParams{
		PagesRead: req.Pagesread,
	}

	for _, booklist := range booklists {
		if booklist.BookID == req.BookId {
			arg.BookListID = booklist.ID
		}
	}

	err = h.service.CreateReadConfirmation(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, lib.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "read confirmation has been created", nil, nil))
}
