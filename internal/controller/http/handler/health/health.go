package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type health struct{}

func NewHealth() *health {
	return &health{}
}

func (h *health) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "ok!")
}

func (h *health) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/heath", h.HealthCheck)
}
