package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_config "github.com/wildanie12/dts-task-planner/config"
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

func (controller WorkerController) Store(c echo.Context) error {
	workerRequest := _models.Worker{}
	c.Bind(&workerRequest)

	controller.db.Save(&workerRequest)

	return c.Redirect(http.StatusFound, "/")
}

func (controller WorkerController) Edit(c echo.Context) error {
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}
	
	tasks := []_models.Task{}
	workers := []_models.Worker{}
	controller.db.Find(&tasks)
	controller.db.Find(&workers)

	worker := _models.Worker{}
	controller.db.First(&worker, id)

	baseURL := "http://" + _config.Get().App.Host
	if _config.Get().App.Port != "80" {
		baseURL += ":" + _config.Get().App.Port
	}
	
	return c.Render(http.StatusOK, "worker-edit.html", map[string]interface{} {
		"worker": worker,
		"workers": workers,
		"tasks": tasks,
		"baseURL": baseURL,
		"isMale": map[bool]bool{ true: true, false: false }[worker.Gender == "L"],
		"isFemale": map[bool]bool{ true: true, false: false }[worker.Gender == "P"],
	})

} 


func (controller WorkerController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	workerRequest := _models.Worker{}
	c.Bind(&workerRequest)

	controller.db.Model(_models.Worker{}).Where("id = ?", id).Updates(&workerRequest)

	return c.Redirect(http.StatusFound, "/")
}


func (controller WorkerController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	workerRequest := _models.Worker{}
	c.Bind(&workerRequest)

	controller.db.Where("id = ?", id).Delete(&_models.Worker{})

	return c.Redirect(http.StatusFound, "/")

}