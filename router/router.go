package router

import (
	_ "new-bank-api/docs"
	"new-bank-api/handler"

	swagger "github.com/arsmn/fiber-swagger/v2"
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

	app.Get("/docs/*", swagger.Handler)
}
