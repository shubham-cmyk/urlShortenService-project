package routers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/database"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/models"
)

func URLShortner(app *fiber.Ctx) error {

	requestBody := new(models.Request)

	if err := app.BodyParser(&requestBody); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot Parse the Request ",
		})
	}

	db1 := database.CreateClient(1)
	defer db1.Close()
	db2 := database.CreateClient(2)
	defer db2.Close()
	var ctx = database.Ctx
	var (
		shortGenerated string
	)

	short, exist := checkLongURL(ctx, requestBody.URL, db2)

	switch exist {
	case true:
		shortGenerated = short
		return urlResponse(requestBody, shortGenerated, app)
	case false:
		shortGenerated = generateshortURL(ctx, db1, requestBody)
		err := checkShortInUse(ctx, db1, app, shortGenerated)
		if err != nil {
			return err
		}
		err = saveShort(ctx, db1, requestBody, app, shortGenerated)
		if err != nil {
			return err
		}
		err = saveKeyPair(ctx, db2, shortGenerated, requestBody.URL, requestBody)
		if err != nil {
			return err
		}
		return urlResponse(requestBody, shortGenerated, app)

	}

	return nil
}

func urlResponse(requestBody *models.Request, shortGenerated string, app *fiber.Ctx) error {

	if requestBody.Expiration == 0 {
		requestBody.Expiration = 20
	}

	resp := models.Response{
		URL:        requestBody.URL,
		SHORT:      os.Getenv("IP") + "/" + shortGenerated,
		Expiration: requestBody.Expiration,
	}

	return app.JSON(resp)

}
