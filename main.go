package main

import (
	"log"

	"github.com/Jose-P-C/DevOps-interview/database"
	"github.com/Jose-P-C/DevOps-interview/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my Courses API")
}

func setupRoutes(app *fiber.App) {
	// endpoints
	app.Get("/api", welcome)
	app.Post("/api/courses", routes.CreateCourse)
	app.Get("/api/courses", routes.GetCourses)
	app.Get("/api/courses/:id", routes.GetCourse)
	app.Put("/api/courses/:id", routes.UpdateCourse)
	app.Delete("/api/courses/:id", routes.DeleteCourse)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
