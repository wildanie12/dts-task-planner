package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	_config "github.com/wildanie12/dts-task-planner/config"
	_controllers "github.com/wildanie12/dts-task-planner/controllers"
	_models "github.com/wildanie12/dts-task-planner/models"
	_routes "github.com/wildanie12/dts-task-planner/routes"
	_mysql "github.com/wildanie12/dts-task-planner/utils/mysql"
	_renderer "github.com/wildanie12/dts-task-planner/utils/renderer"
)

func main() {
	config := _config.Get()

	db := _mysql.Connect(config)
	db.AutoMigrate(&_models.Worker{})

	e := echo.New()
	e.Renderer = _renderer.New()
	e.Static("/", "public")
	
	taskController := _controllers.NewTaskControlller()
	workerController := _controllers.NewWorkerController(db)

	_routes.RegisterRoute(e, taskController, workerController)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}