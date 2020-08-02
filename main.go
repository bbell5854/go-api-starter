package main

import (
	v1 "go-api-starter/routers/api/v1"
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

const port = "9000"

type healthResponse struct {
	Status string `json:"status"`
}

func main() {
	app := fiber.New()

	// Setup Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(cors.New())

	// Health Route
	app.Get("/health", getHealth)

	// Setup V1 Routes
	v1.SetupRoutes(app)

	log.Printf("Server Listening -  %s", port)
	_ = app.Listen(port)
}

func getHealth(c *fiber.Ctx) {
	_ = c.JSON(healthResponse{
		Status: "OK",
	})
}
