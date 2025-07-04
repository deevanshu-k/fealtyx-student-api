package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deevanshu-k/fealtyx-student-api/src/config"
	"github.com/deevanshu-k/fealtyx-student-api/src/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	studentApi := app.Group("/student")
	{
		studentApi.Post("", handlers.CreateStudent)                // Create a new student
		studentApi.Get("", handlers.GetAllStudent)                 // Get all students
		studentApi.Get("/:id", handlers.GetStudentById)            // Get a student by ID
		studentApi.Put("/:id", handlers.UpdateStudent)             // Update a student by ID
		studentApi.Delete("/:id", handlers.DeleteStudent)          // Delete a student by ID
		studentApi.Delete("/:id", handlers.GenerateStudentSummary) // Generate a summary for a student
	}

	// Handle Incorrect route
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "404, Not found!",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.PORT)))
}
