package main

import (
	"github.com/KimmyLps/test/config"
	r "github.com/KimmyLps/test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	godotenv.Load()

	config.InitialDb()

	r.RouteTest(app)
	r.RouteEmployee(app)

	app.Listen(":3000")

	// test
}
