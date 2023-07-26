package main

import (
	"apiProject/api"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	api.Initialize()
	app := fiber.New()
	api.RegisterEndpoints(app)
	log.Fatal(app.Listen(":3000"))
}
