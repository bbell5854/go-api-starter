package main

import (
	"fmt"
	"go-api-starter/internal/database"
	"go-api-starter/internal/health"
	v1 "go-api-starter/routers/api/v1"
	"log"
	"os"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/joho/godotenv"
)

const port = "9000"

func init() {
	var dbType = "mysql"
	var dbConnString = fmt.Sprintf("%s:%s@tcp(%s:%s)", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	// If we're not in production, lets load an env file
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Println(err)
		}
	}

	if err := database.Connect(dbType, dbConnString); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer database.DB.Close()

	app := fiber.New()

	// Setup Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(cors.New())

	// Health Route
	app.Get("/health", health.GetHealthCheck)

	// Setup V1 Routes
	v1.SetupRoutes(app)

	log.Printf("Server Listening -  %s", port)
	_ = app.Listen(port)
}
