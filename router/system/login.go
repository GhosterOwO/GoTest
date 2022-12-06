package system

import (
	"go_server/api/controller"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct{}

func (e *LoginRouter) LoginInit(Router *gin.RouterGroup) {
	Router.POST("/login", controller.Login)
	Router.POST("/logout", controller.Logout)
	Router.POST("/add", controller.Add)
	Router.GET("/", controller.LoadHtml)
}
