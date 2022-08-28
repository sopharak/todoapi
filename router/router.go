package router

import (
	"net/http"
	"todoapi/controller"
	"todoapi/domain"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, storage *controller.Storage) {

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, domain.Health{
			Message: "OK",
		})
	})

	e.POST("/todo", storage.CreateTodo)
	e.GET("/todo/:id", storage.GetTodo)
	e.GET("/todos", storage.GetTodos)
	e.PUT("/todo/:id", storage.UpdateTodo)
	e.DELETE("/todo/:id", storage.DeleteTodo)

}
