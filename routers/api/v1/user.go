package v1

import (
	"github.com/gofiber/fiber"
)

type userResponse struct {
	Message string `json:"message"`
}

func GetUser(c *fiber.Ctx) {
	_ = c.JSON(userResponse{
		Message: "Test",
	})
}
