package models

import (
	"github.com/haha/gin-web/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func (Todo) TableName() string {
	return "todo"
}

// 增删改查
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func QueryTodo() (todos []*Todo, err error) {
	err = dao.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err := dao.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id= ?", id).Delete(Todo{}).Error
	return
}
