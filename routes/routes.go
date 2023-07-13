package routes

import (
	"crud-simple-api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, ac *controllers.AuthController) {

	e.POST("/login", ac.Login())

	apiGroup := e.Group("/api")

	
	UserRoutes(apiGroup.Group("/users"))
}	