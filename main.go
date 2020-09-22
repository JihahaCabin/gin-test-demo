package main

import (
	"github.com/haha/gin-web/dao"
	"github.com/haha/gin-web/routers"
)

func main() {

	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	//程序退出，关闭连接
	defer dao.CloseMySQL()

	//绑定模型,初始化表
	//dao.InitModel()
	//注册路由
	r := routers.SetupRouter()
	//启动
	r.Run(":8080")

}
