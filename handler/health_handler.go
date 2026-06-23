package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-org/your-app/contract"
)

type HealthHandler struct {
	service contract.HealthService
}

func (c *HealthHandler) InitService(svc contract.HealthService) {
	c.service = svc
}

func (h *HealthHandler) GetHealthStatus(ctx *gin.Context) {
	response := h.service.GetStatus()

	ctx.JSON(http.StatusOK, response)
}
