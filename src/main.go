package main

import (
	"fmt"
	"log"

	"github.com/deevanshu-k/fealtyx-student-api/src/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.PORT)))
}
