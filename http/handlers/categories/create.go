package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rhtyx/narawangsa/lib"
)

type categoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *handler) Create(ctx *gin.Context) {
	var req categoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	err := h.service.CreateCategory(ctx, req.Name)
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

	ctx.JSON(http.StatusOK, lib.Response("success", "category has been created", nil, nil))
}
