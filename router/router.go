package router

import (
	"new-bank-api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	v1 := app.Group("/v1")

	// User
	user := v1.Group("/users")

	user.Get("/", handler.GetAllUser)
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)
}
