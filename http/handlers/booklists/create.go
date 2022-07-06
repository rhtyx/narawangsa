package booklists

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rhtyx/narawangsa/http/middleware"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

type booklistRequest struct {
	BookID    int64     `json:"book_id" binding:"required"`
	IsRead    bool      `json:"is_read"`
	PagesRead int32     `json:"pages_read"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

func (h *handler) Create(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	var req booklistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	arg := postgres.CreateBookListParams{
		UserID:    authPayload.UserId,
		BookID:    req.BookID,
		IsRead:    req.IsRead,
		PagesRead: req.PagesRead,
		EndDate:   req.EndDate,
	}

	err := h.service.CreateBookList(ctx, arg)
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

	ctx.JSON(http.StatusCreated, lib.Response("success", "booklist has been created", nil, nil, nil))
}
