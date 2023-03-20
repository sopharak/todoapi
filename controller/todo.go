package controller

import (
	"net/http"
	"todoapi/domain"
	"todoapi/model"

	"github.com/labstack/echo/v4"
)

func (s *Storage) CreateTodo(ctx echo.Context) error {
	req := domain.TodoReq{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	tx := s.db.Create(&model.Todo{
		Task: req.Task,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return ctx.JSON(http.StatusOK, &domain.Message{
		Message: "success",
	})
}

func (s *Storage) GetTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	var todo model.Todo
	tx := s.db.Find(&todo, id)

	if tx.Error != nil {
		return tx.Error
	}
	return ctx.JSON(http.StatusOK, model.ToTodoDomain(&todo))
}

func (s *Storage) GetTodos(ctx echo.Context) error {
	var todos []*model.Todo
	tx := s.db.Find(&todos)

	if tx.Error != nil {
		return tx.Error
	}
	return ctx.JSON(http.StatusOK, model.ToTodoDomains(todos))
}

func (s *Storage) UpdateTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	req := domain.TodoReq{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	todo := &model.Todo{}
	tx := s.db.First(&todo, id)
	if tx.Error != nil {
		return tx.Error
	}

	todo.Task = req.Task
	tx = s.db.Save(todo)
	if tx.Error != nil {
		return tx.Error
	}

	return ctx.JSON(http.StatusOK, &domain.Message{
		Message: "success",
	})
}

func (s *Storage) DeleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	tx := s.db.Delete(&model.Todo{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return ctx.JSON(http.StatusOK, &domain.Message{
		Message: "success",
	})
}
