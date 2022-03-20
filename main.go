package main

import (
	"log"
	"new-bank-api/config"
	"new-bank-api/database"
	"new-bank-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func init() {
	database.ConnectDB()
}

func main() {

	app := fiber.New()

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} - ${pid} ${locals:requestid} ${status} - ${method} ${path}​\n​",
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":" + config.Config("APP_PORT")))
}
