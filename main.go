package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:= gin.Default()
	//加载静态文件
	r.Static("/static","static")

	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	//代办事项
	//添加
	//查看
	//修改
	//删除

	r.Run(":8080")

}
