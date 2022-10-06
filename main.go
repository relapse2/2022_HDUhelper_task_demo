package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var State = make(map[string]interface{})
var users = make(map[string]interface{})

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	_, bool1 := users[username]
	if bool1 {
		users[username] = password

		State["state"] = 1
		State["text"] = "注册成功！"
		c.String(200, "%v", State)
	} else {

		State["state"] = 2
		State["text"] = "此用户已存在！"
		c.String(307, "%v", State)
	}
	c.String(http.StatusOK, "%v", State)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	_, bool1 := users[username]
	if bool1 {
		State["state"] = 6
		State["text"] = "登录失败！此用户尚未注册！"
		c.String(307, "%v", State)
	} else {
		if password == users[username] {

			State["state"] = 4
			State["text"] = "登录成功！"
			c.String(200, "%v", State)
		} else {
			State["state"] = 5
			State["text"] = "密码错误！"
			c.String(401, "%v", State)
		}
	}

}

func GetUser(c *gin.Context) {
	if State["state"] == 4 {
		name := c.Param("name")
		c.String(http.StatusOK, "这是用户 %s 的页面", name)
	} else {
		state := "未登录"
		c.String(401, "%v", state)
	}
}
func Logout(c *gin.Context) {
	if State["state"] == 4 {
		State["state"] = 0
		State["text"] = "已登出"
		c.String(200, "%v", State)
	} else {
		state := "未登录"
		c.String(401, "%v", state)
	}
}
func main() {
	r := gin.Default()
	p1 := r.Group("/user")
	{
		p1.POST("/Register", Register)
		p1.GET("/Login", Login)
		p1.GET("/GetUser/:name", GetUser)
		p1.GET("/Logout", Logout)
	}
	r.Run(":8000")
}
