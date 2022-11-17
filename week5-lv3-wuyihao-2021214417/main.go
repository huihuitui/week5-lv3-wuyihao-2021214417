package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

type User struct {
	Name     string
	Password string
}

var datafile *os.File
var str string
var Usermap map[string]User

func init() {
	file, err := os.OpenFile("D:/user_info.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	datafile = file
}
func main() {
	r := gin.Default()
	r.GET("/register", func(c *gin.Context) {
		UserInfo() //输入账号密码
		ReadInfo()
		UserInfoWrite() //每次输入存入数据
		c.SetCookie("account", str, 60, "/", "localhost", false, true)
		c.String(200, "success!")
	})
	r.GET("/login", MiddleWar(), func(c *gin.Context) {
		c.String(200, "登陆成功")
	})
	r.Run(":8080")
}
