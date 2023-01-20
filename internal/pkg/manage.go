package manage

import (
	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

type manage struct {
	health apps
}

func NewManage() *manage {
	HealthCheck := HealthCheck()
}
