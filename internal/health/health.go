package health

import "github.com/gofiber/fiber"

type HealthCheck struct {
	Status string `json:"status"`
}

func GetHealthCheck(c *fiber.Ctx) {
	_ = c.JSON(HealthCheck{
		Status: "OK",
	})
}
