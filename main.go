package main

import (
	"log"

	"gihub.com/rivaldi-fsociety/golang-go-fiber/database"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/database/migration"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	database.ConnectDB()
	migration.RunMigrate()
}

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	routers.RouterApp(app)

	log.Fatal(app.Listen(":8080"))
}
