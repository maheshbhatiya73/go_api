package main

import (
	"goapi/config"
	"goapi/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	routes.AuthRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
