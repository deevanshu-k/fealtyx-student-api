package main

import (
	"fmt"
	"log"

	"github.com/deevanshu-k/fealtyx-student-api/src/config"
	"github.com/deevanshu-k/fealtyx-student-api/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

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
		return c.Status(500).JSON(fiber.Map{
			"error": "Not found!",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.PORT)))
}
