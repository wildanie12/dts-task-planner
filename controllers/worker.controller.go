package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_models "github.com/wildanie12/dts-task-planner/models"
	"gorm.io/gorm"
)

type WorkerController struct {
	db *gorm.DB
}

func NewWorkerController(db *gorm.DB) *WorkerController {
	return &WorkerController{
		db: db,
	}
}

func (controller WorkerController) Index(c echo.Context) error {
	
	workers := []_models.Worker{}
	controller.db.Find(&workers)

	return c.Render(http.StatusOK, "worker-index.html", map[string]interface{} {
		"workers": workers,
	})
} 