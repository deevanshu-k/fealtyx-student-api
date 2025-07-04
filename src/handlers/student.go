package handlers

import "github.com/gofiber/fiber/v2"

func CreateStudent(c *fiber.Ctx) error {
	return c.SendString("Create Student")
}

func GetAllStudent(c *fiber.Ctx) error {
	return c.SendString("Get All Students")
}

func GetStudentById(c *fiber.Ctx) error {
	return c.SendString("Get Student By Id")
}

func UpdateStudent(c *fiber.Ctx) error {
	return c.SendString("Update Student")
}

func DeleteStudent(c *fiber.Ctx) error {
	return c.SendString("Delete Student")
}

func GenerateStudentSummary(c *fiber.Ctx) error {
	return c.SendString("Generate Student Summary")
}
