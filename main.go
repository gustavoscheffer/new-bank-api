package main

import (
	"log"
	"new-bank-api/database"
	"new-bank-api/router"

	"github.com/gofiber/fiber/v2"
)

func init() {
	database.ConnectDB()
}

func main() {

	app := fiber.New()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
