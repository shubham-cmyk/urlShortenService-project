package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/database"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/models"
)

func URLShortner(app *fiber.Ctx) error {

	requestBody := new(models.Request)

	if err := app.BodyParser(requestBody); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot Parse the Request ",
		})
	}

	db1 := database.CreateClient(1)
	defer db1.Close()

	// Generate a random ID of 6 char if URL_short is not there
	var rndId string
	if requestBody.URL_SHORT == "" {
		rndId = uuid.New().String()[:6]
	} else {
		rndId = requestBody.URL_SHORT
	}

	// Check Whether the Short URL is already present in the database or not
	// Possible value Empty or String
	longURL, err := db1.Get(database.Ctx, rndId).Result()
	if err != nil {
		return err
	}

	if longURL != "" {
		return app.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": " URL is already in the use",
		})
	}

	// Set the Short value into the database
	err = db1.Set(database.Ctx, rndId, requestBody.URL_LONG, requestBody.Expiration).Err()
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unable to connect to server",
		})
	}

	resp := models.Response{
		URL_LONG:  requestBody.URL_LONG,
		URL_SHORT: "localhost:3000" + rndId,
	}

	return app.JSON(resp)

}

func createShort() {

}

func reponseShortner() {

}
