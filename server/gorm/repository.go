package gorm

import (
	model "github.com/jackyuan2010/todoapp/server/model"
)

type Repository interface {
	QueryById(id string) *model.ToDoTask
}