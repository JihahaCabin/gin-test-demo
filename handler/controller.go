package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haha/gin-web/models"
	"net/http"
)

/**
 url ————>handler ---> logic ---> model
请求来了---> 控制器 ————> 业务逻辑 ---> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, todo)
}

func ListTodo(c *gin.Context) {
	todos, err := models.QueryTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "get id failed",
		})
		return
	}

	todo, err := models.GetATodo(id)

	fmt.Printf("todo : %#v  %#v\n", todo, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	//将修改的属性，添加到todo 结构体中
	c.BindJSON(&todo)
	//保存数据，如果失败，则返回错误
	if err := models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//成功，返回数据
	c.JSON(http.StatusOK, todo)

}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "delete failed",
		})
		return
	}

	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		id: "deleted",
	})

}
