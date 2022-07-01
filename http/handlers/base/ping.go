package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ready to serve")
}
