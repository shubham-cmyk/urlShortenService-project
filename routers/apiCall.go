package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/database"
)

func ApiCallURL(app *fiber.Ctx) error {

	// Short URL from the Params
	url := app.Params("url")

	db1 := database.CreateClient(1)
	defer db1.Close()

	// Get the Long URL from the Short URL
	longURL, err := db1.Get(database.Ctx, url).Result()

	if err != nil {
		return err
	}

	// Redirect the request to the long URL
	return app.Redirect(longURL, fiber.StatusPermanentRedirect)
}