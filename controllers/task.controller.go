package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_models "github.com/wildanie12/dts-task-planner/models"
	"gorm.io/gorm"
)


type TaskController struct {
	db *gorm.DB
}


func NewTaskControlller(db *gorm.DB) *TaskController {
	return &TaskController{
		db: db,
	}
}

func (controller TaskController) Index(c echo.Context) error  {
	tasks := []_models.Task{}
	controller.db.Find(&tasks)
	fmt.Println(tasks)

	return c.Render(http.StatusOK, "worker-index.html", map[string]interface{} {
		"tasks": tasks,
	})
}