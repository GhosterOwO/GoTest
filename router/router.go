package router

import (
	"fmt"
	"go_server/base"
	"go_server/global"
	"go_server/router/system"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	System system.RouterGroup
}

func Routers() {
	fmt.Println("開始執行router")
	store := cookie.NewStore([]byte("BNJ123456"))
	router := gin.Default()
	router.Use(sessions.Sessions("BNJ", store))
	global.Mysql = base.InitDB()
	PublicGroup := router.Group("")
	router.LoadHTMLGlob("views/*")
	router.Static("/public", "./public")
	system.RouterGroupApp.LoginInit(PublicGroup)
	system.RouterGroupApp.UserInit(PublicGroup)
	system.RouterGroupApp.ImgInit(PublicGroup)
	router.Run(":5158")
}
