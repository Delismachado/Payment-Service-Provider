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
	//health.HealthCheck()
	HealthCheck := health.HealthCheck()

	return nil
}

func (m *manage) Routes(g *gin.RouterGroup) {
	m.health.RegisterRoutes(g)
}
