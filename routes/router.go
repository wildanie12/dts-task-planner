package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanie12/dts-task-planner/controllers"
)


func RegisterRoute(e *echo.Echo, taskController *controllers.TaskController, workerController *controllers.WorkerController) {
	e.GET("", taskController.Index)
	e.POST("/tasks/store", taskController.Store)
	e.GET("/tasks/:id/edit", taskController.Edit)
	e.POST("/tasks/:id/update", taskController.Update)
	e.GET("/tasks/:id/delete", taskController.Delete)
	
	e.POST("/workers/store", workerController.Store)
	e.GET("/workers/:id/edit", workerController.Edit)
	e.POST("/workers/:id/update", workerController.Update)
	e.GET("/workers/:id/delete", workerController.Delete)
}