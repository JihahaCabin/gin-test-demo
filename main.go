package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func (Todo) TableName() string {
	return "todo"
}

var (
	DB *gorm.DB
)

func initMySQL()(err error) {
	//创建数据库
	//连接数据库
	DB, err = gorm.Open("mysql", "root:tangguo@tcp(localhost:3306)/blog_service?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open failed,",err)
		return
	}
	//测试连通性
	err = DB.DB().Ping()
	return
}

func main() {

	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	//程序退出，关闭连接
	defer DB.Close()

	//绑定模型,初始化表
	DB.AutoMigrate(Todo{})

	r:= gin.Default()
	//加载静态文件
	r.Static("/static","static")

	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			err2 := DB.Create(&todo).Error
			if err2 != nil {
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK,todo)
		})
		//查看所有代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todos []Todo
			err2 := DB.Find(&todos).Error
			if err2 != nil {
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK,todos)
		})
		//查看某一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok{
				c.JSON(http.StatusOK,gin.H{
					"error":"get id failed",
				})
				return
			}
			var todo Todo
			if err := DB.Where("id = ? ", id).First(&todo).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK,todo)

		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id,ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK,gin.H{
					"error":"get id failed",
				})
				return
			}

			var todo Todo
			if err2 := DB.Where("id = ?", id).First(&todo).Error;err2!=nil{
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
				return
			}
			//将修改的属性，添加到todo 结构体中
			c.BindJSON(&todo)
			//保存数据，如果失败，则返回错误
			if err := DB.Save(&todo).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
				return
			}
			//成功，返回数据
			c.JSON(http.StatusOK,todo)

		})

		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok{
				c.JSON(http.StatusOK,gin.H{
					"error":"delete failed",
				})
				return
			}

			if err := DB.Where("id = ?", id).Delete(&Todo{}).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK,gin.H{
				id: "deleted",
			})

		})
	}



	r.Run(":8080")

}
