package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanie12/dts-task-planner/controllers"
)


func RegisterRoute(e *echo.Echo, taskController *controllers.TaskController, workerController *controllers.WorkerController) {
	e.GET("/tasks", taskController.Index)
	e.POST("/tasks", taskController.Store)
	e.GET("/tasks/:id/edit", taskController.Edit)
	e.POST("/tasks/:id/update", taskController.Update)
	e.GET("/tasks/:id/delete", taskController.Delete)

	e.GET("/workers", workerController.Index)
}