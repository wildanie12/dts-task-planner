package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

	fmt.Println(taskRequest)

	controller.db.Save(&taskRequest)

	return c.Redirect(http.StatusMovedPermanently, "/tasks")
}