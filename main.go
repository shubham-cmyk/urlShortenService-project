package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/routers"
)

func main() {

	app := fiber.New()

	setupRoutes(app)

	err := app.Listen(":3000")
	fmt.Printf("got an error while listen %s", err)
}

func setupRoutes(app *fiber.App) {

	app.Get("/:url", routers.ApiCallURL)
	app.Post("/api", routers.URLShortner)

}
