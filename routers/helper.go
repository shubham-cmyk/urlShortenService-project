package routers

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shubham-cmyk/infraCloud-URL-Shortner/models"
)

func checkLongURL(ctx context.Context, url string, clientDB2 *redis.Client) (string, bool) {

	short, err := clientDB2.Get(ctx, url).Result()
	if err != nil {
		return "", false
	}

	return short, true
}

func generateshortURL(ctx context.Context, db1 *redis.Client, requestBody *models.Request) string {
	// Generate a random ID of 6 char if URL_short is not there
	var rndId string
	if requestBody.Short == "" {
		rndId = uuid.New().String()[:6]
	} else {
		rndId = requestBody.Short
	}

	return rndId

}

func checkShortInUse(ctx context.Context, db1 *redis.Client, app *fiber.Ctx, shortGenerated string) error {

	longURL, _ := db1.Get(ctx, shortGenerated).Result()

	if longURL != "" {
		return app.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": " URL is already in the use",
		})
	}

	return nil
}

func saveShort(ctx context.Context, db1 *redis.Client, requestBody *models.Request, app *fiber.Ctx, shortGenerated string) error {
	err := db1.Set(ctx, shortGenerated, requestBody.URL, requestBody.Expiration*time.Minute).Err()
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unable to connect to server",
		})
	}
	return nil
}

func saveKeyPair(ctx context.Context, db2 *redis.Client, short string, long string, request *models.Request) error {
	return db2.Set(ctx, long, short, request.Expiration).Err()

}
