package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	//创建数据库
	//连接数据库
	DB, err = gorm.Open("mysql", "root:tangguo@tcp(localhost:3306)/blog_service?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open failed,", err)
		return
	}
	//测试连通性
	err = DB.DB().Ping()
	return
}

func CloseMySQL() {
	DB.Close()
}

// 会有循环依赖问题，注释掉
//func InitModel() {
//	DB.AutoMigrate(models.Todo{})
//}
