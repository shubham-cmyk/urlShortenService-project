package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/routers"
)

func main() {

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {

	app.Get("/:url", routers.ApiCallURL)
	app.Post("/api", routers.URLShortner)

}
