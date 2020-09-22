package routers

import (
	"github.com/gin-gonic/gin"
	controller "github.com/haha/gin-web/handler"
)

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()
	//加载静态文件
	r.Static("/static", "static")

	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.AddATodo)
		//查看所有代办事项
		v1Group.GET("/todo", controller.ListTodo)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
