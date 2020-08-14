package v1

import (
	"go-api-starter/internal/users"
	"go-api-starter/middleware/auth"

	"github.com/gofiber/fiber"
)

// SetupRoutes Initializes routes for the API
func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1", auth.Protected())
	apiV1.Get("/user/settings", users.GetSettings)
}
