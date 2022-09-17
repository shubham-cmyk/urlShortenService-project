package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/routers"
)

func main() {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	// Create a Fiber named instance
	app := fiber.New()
	// Set up the routing path for the GET and POST method
	setupRoutes(app)

	err := app.Listen(":3000")
	// Raise the Panic if found error
	panic(err)
}

func setupRoutes(app *fiber.App) {

	app.Get("/:url", routers.ApiCallURL)
	app.Post("/api", routers.URLShortner)

}
