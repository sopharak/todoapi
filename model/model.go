package model

import (
	"todoapi/domain"

	"gorm.io/gorm"
)

type Todo struct {
	Task string
	gorm.Model
}

func ToTodoDomains(todos []*Todo) []*domain.TodoRes {
	res := []*domain.TodoRes{}
	for _, v := range todos {
		res = append(res, ToTodoDomain(v))
	}
	return res
}

func ToTodoDomain(todo *Todo) *domain.TodoRes {
	return &domain.TodoRes{
		Id:          todo.ID,
		Task:        todo.Task,
		CreatedDate: todo.CreatedAt,
		UpdatedDate: todo.UpdatedAt,
	}
}
