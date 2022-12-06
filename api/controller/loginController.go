package controller

import (
	"fmt"
	"go_server/api/dao"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadHtml(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") == nil {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	c.HTML(http.StatusOK, "index.html", nil)
}

func Login(c *gin.Context) {
	var loginCmd LoginCommand
	c.ShouldBind(&loginCmd)
	if loginCmd.Username == "" || loginCmd.Password == "" {
		c.JSON(200, gin.H{"error": 2, "msg": "參數遺失"})
		return
	}

	users, err := dao.LoginUser(loginCmd.Username, loginCmd.Password)
	if err != nil {
		fmt.Println("err :", err)
		c.JSON(200, gin.H{"error": 1, "msg": "登入失敗 帳號密碼有誤"})
		return
	}
	session := sessions.Default(c)
	session.Set("User", loginCmd.Username)
	session.Save()
	c.JSON(200, gin.H{"error": 0, "msg": "登入成功", "users": users})
	return
}

func Add(c *gin.Context) {
	var loginCmd LoginCommand
	c.ShouldBind(&loginCmd)

	if loginCmd.Username == "" || loginCmd.Password == "" {
		c.JSON(200, gin.H{"error": 2, "msg": "參數遺失"})
		return
	}
	users, err := dao.AddUser(loginCmd.Username, loginCmd.Password)
	if err != nil {
		fmt.Println("err :", err)
		c.JSON(200, gin.H{"error": 1, "msg": "創建失敗"})
		return
	}

	c.JSON(200, gin.H{"error": 0, "msg": "創建成功", "users": users})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("User")
	session.Save()
	c.JSON(200, gin.H{"error": 0, "msg": "登出成功"})
	return
}
