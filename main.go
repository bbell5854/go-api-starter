package main

import (
	"fmt"
	"go-api-starter/internal/health"
	"go-api-starter/internal/mysql"
	v1 "go-api-starter/routers/api/v1"
	"log"
	"os"
	"time"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

const port = "9000"

func init() {
	var err error

	// If we're not in production, lets load an env file
	if os.Getenv("ENVIRONMENT") != "production" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Println(err)
		}
	}

	mysql.DB, err = dbConnect()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer mysql.DB.Close()

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

func dbConnect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
