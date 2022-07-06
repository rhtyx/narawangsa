package notifications

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhtyx/narawangsa/lib"
)

type sendNotificationRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Quote string `json:"quote" binding:"required"`
}

func (h *handler) Send(ctx *gin.Context) {
	var req sendNotificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.ErrorResponse(err))
		return
	}

	ctx.Set("name", req.Name)
	ctx.Set("email", req.Email)
	ctx.Set("quote", req.Quote)

	err := h.service.SendNotifications(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lib.Response("success", "notification has been sent", nil, nil))
}
