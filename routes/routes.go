package routes

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) {


	apiGroup := e.Group("/api")

	
	UserRoutes(apiGroup.Group("/users"))
}	