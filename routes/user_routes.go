package routes

import (
	"crud-simple-api/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group) {

	g.GET("", controllers.GetAllUsers)
	g.POST("", controllers.CreateUser)
	g.GET("/:id", controllers.GetUserById)
	g.PUT("/:id", controllers.UpdateUser)
	g.DELETE("/:id", controllers.DeleteUser)
}