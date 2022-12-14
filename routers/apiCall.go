package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/database"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/models"
)

func ApiCallURL(app *fiber.Ctx) error {

	// Short URL from the Params
	url := app.Params("url")

	db1 := database.CreateClient(1)
	defer db1.Close()

	// Get the Long URL from the Short URL
	longURL, err := db1.Get(database.Ctx, url).Result()

	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "unable to find the key",
		})
	}

	reponsebody := models.Response{
		URL:   longURL,
		SHORT: url,
	}
	return app.JSON(reponsebody)
}
