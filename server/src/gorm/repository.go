package gorm

import (
	model "github.com/jackyuan2010/todoapp/server/src/model"
)

type Repository interface {
	QueryById(todoTask model.ToDoTask, id string) map[string]interface{}
}