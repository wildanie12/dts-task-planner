package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	_config "github.com/wildanie12/dts-task-planner/config"
	_controllers "github.com/wildanie12/dts-task-planner/controllers"
	_routes "github.com/wildanie12/dts-task-planner/routes"
	_renderer "github.com/wildanie12/dts-task-planner/utils/renderer"
)

func main() {
	config := _config.Get()
	e := echo.New()
	e.Renderer = _renderer.New()
	e.Static("/", "public")
	
	taskController := _controllers.NewTaskControlller()

	_routes.RegisterRoute(e, taskController)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}