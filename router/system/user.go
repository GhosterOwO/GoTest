package system

import (
	"go_server/api/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (e *UserRouter) UserInit(Router *gin.RouterGroup) {
	Router.GET("/user", controller.GetUser)
}
