package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanie12/dts-task-planner/controllers"
)


func RegisterRoute(e *echo.Echo, taskController *controllers.TaskController, workerController *controllers.WorkerController) {
	e.GET("/", taskController.Index)

	e.GET("/workers", workerController.Index)
}