package v1

import (
	"github.com/gofiber/fiber"
)

type userResponse struct {
	Message string `json:"message"`
}

// GetUser Handler Function to Get User
func GetUser(c *fiber.Ctx) {
	_ = c.JSON(userResponse{
		Message: "Test",
	})
}
