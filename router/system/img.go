package system

import (
	"go_server/api/controller"

	"github.com/gin-gonic/gin"
)

type ImgRouter struct{}

func (e *ImgRouter) ImgInit(Router *gin.RouterGroup) {
	Router.GET("/getimg", controller.GetImg)
	Router.POST("/addimg", controller.AddImg)
	Router.POST("/delimg", controller.DelImg)
	Router.POST("/uploadimg", controller.UploadImg)
}
