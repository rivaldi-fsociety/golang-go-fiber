package routers

import (
	"gihub.com/rivaldi-fsociety/golang-go-fiber/controllers"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App) {
	c.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello World")
	})

	api := c.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/employee", middlewares.Auth, controllers.ShowAllEmployee)
	v1.Get("/employee/:id", middlewares.Auth, controllers.GetByIdEmployee)
	v1.Post("/employee", middlewares.Auth, controllers.CreateEmployee)
	v1.Delete("/employee/:id", middlewares.Auth, controllers.DeleteEmployee)
	v1.Patch("/employee/:id", middlewares.Auth, controllers.UpdateEmployee)

	v1.Post("/login", controllers.Login)
	v1.Post("/user", controllers.CreateUser)
}
