package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	_config "github.com/wildanie12/dts-task-planner/config"
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
	controller.db.Preload("Worker").Find(&tasks)
	
	workers := []_models.Worker{}
	controller.db.Find(&workers)

	return c.Render(http.StatusOK, "index.html", map[string]interface{} {
		"tasks": tasks,
		"workers": workers,
	})
}

func (controller TaskController) Store(c echo.Context) error  {

	taskRequest := _models.Task{}
	c.Bind(&taskRequest)

	// Populate data 
	startDate, _ := time.Parse("2006-01-02", c.FormValue("start_date"))
	taskRequest.StartDate = startDate

	endDate, _ := time.Parse("2006-01-02", c.FormValue("end_date"))
	taskRequest.EndDate = endDate

	workerID, _ := strconv.Atoi(c.FormValue("worker_id"))
	taskRequest.WorkerID = uint(workerID)
	
	done, _ := strconv.ParseBool(c.FormValue("done"))
	taskRequest.Done = done


	controller.db.Save(&taskRequest)

	return c.Redirect(http.StatusFound, "/")
}

func (controller TaskController) Edit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	tasks := []_models.Task{}
	workers := []_models.Worker{}
	controller.db.Find(&tasks)
	controller.db.Find(&workers)
	
	task := _models.Task{}
	controller.db.First(&task, id)

	baseURL := "http://" + _config.Get().App.Host
	if _config.Get().App.Port != "80" {
		baseURL += ":" + _config.Get().App.Port
	}


	return c.Render(http.StatusOK, "task-edit.html", map[string]interface{} {
		"task": task,
		"tasks": tasks,
		"workers": workers,
		"baseURL": baseURL,
	})
}

func (controller TaskController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	taskRequest := _models.Task{}
	c.Bind(&taskRequest)

	// Populate data 
	startDate, _ := time.Parse("2006-01-02", c.FormValue("start_date"))
	taskRequest.StartDate = startDate

	endDate, _ := time.Parse("2006-01-02", c.FormValue("end_date"))
	taskRequest.EndDate = endDate

	workerID, _ := strconv.Atoi(c.FormValue("worker_id"))
	taskRequest.WorkerID = uint(workerID)
	
	done, _ := strconv.ParseBool(c.FormValue("done"))
	taskRequest.Done = done

	taskRequest.ID = uint(id)

	controller.db.Model(&_models.Task{}).Where("id = ?", id).Updates(taskRequest)


	return c.Redirect(http.StatusFound, "/")
}

func (controller TaskController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	controller.db.Where("id = ?", id).Delete(&_models.Task{})
	return c.Redirect(http.StatusFound, "/")
}