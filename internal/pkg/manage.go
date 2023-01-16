package manage

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/x/mongo/driver/auth"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup, auth auth.Authenticator)
}
