package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问实例 /login?param1=a&param2=b&ids[1]=f&ids[2]=h
	r.GET("/login", func(c *gin.Context) {
		param1 := c.DefaultQuery("param1", "c") //设置初始值
		param2 := c.Query("param2")
		c.String(http.StatusOK, "Hello world "+param1+"_"+param2)
	})
	r.Run(":8282")
}