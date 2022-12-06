package controller

import (
	"fmt"
	"go_server/api/dao"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	data, err := dao.GetUser()
	if err != nil {
		fmt.Println("err :", err)
		c.JSON(200, gin.H{"error": 1, "msg": "取得會員訊息失敗"})
		return
	}
	fmt.Println("data :", data)
	c.JSON(200, gin.H{"error": 0, "msg": "取得會員訊息成功", "users": data})
	return
}
