package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


type TaskController struct {

}


func NewTaskControlller() *TaskController {
	return &TaskController{}	
}

func (controller TaskController) Index(c echo.Context) error  {
	return c.Render(http.StatusOK, "worker-index.html", map[string]interface{} {
		"hello": "Ngok",
	})
}