package v1

import (
	"go-api-starter/internal/auth"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1", auth.ValidateUserToken())
	apiV1.Get("/user", GetUser)
}
