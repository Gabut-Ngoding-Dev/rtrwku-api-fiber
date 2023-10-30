package routes

import (
	"rtrwku-api-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func v1Route(app *fiber.App) {
	v1 := app.Group("/api/v1")

	// User
	user := v1.Group("/user")
	user.Post("/", controllers.CreateUser)
	user.Get("/", controllers.GetAllUser)
	user.Get("/:id", controllers.GetUserByID)
	user.Patch("/:id", controllers.UpdateUserByID)
	user.Delete("/:id", controllers.DeleteUserById)
}
