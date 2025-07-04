package handlers

import (
	"net/http"

	"github.com/deevanshu-k/fealtyx-student-api/src/db"
	"github.com/deevanshu-k/fealtyx-student-api/src/structs"
	"github.com/deevanshu-k/fealtyx-student-api/src/summarizer"
	"github.com/gofiber/fiber/v2"
)

func CreateStudent(c *fiber.Ctx) error {
	var student struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&student); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if student.Name == "" || student.Age == 0 || student.Email == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	createdStudent, err := db.CreateStudent(structs.Student{
		ID:    "",
		Name:  student.Name,
		Age:   student.Age,
		Email: student.Email,
	})
	if err != nil {
		if err.Error() == "STUDENT_ALREADY_EXISTS" {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"message": "Student already exists",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create student",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Student created successfully",
		"data":    createdStudent,
	})
}

func GetAllStudent(c *fiber.Ctx) error {
	students, err := db.GetAllStudents()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get students",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Students retrieved successfully",
		"data":    students,
	})
}

func GetStudentById(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := db.GetStudent(id)
	if err != nil {
		if err.Error() == "STUDENT_NOT_FOUND" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Student not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get student",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student retrieved successfully",
		"data":    student,
	})
}

func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	var student struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&student); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if student.Name == "" && student.Age == 0 && student.Email == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "No fields to update",
		})
	}

	fetchedStudent, err := db.GetStudent(id)
	if err != nil {
		if err.Error() == "STUDENT_NOT_FOUND" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Student not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update student",
		})
	}

	if student.Age != 0 {
		fetchedStudent.Age = student.Age
	}

	if student.Name != "" {
		fetchedStudent.Name = student.Name
	}

	if student.Email != "" {
		fetchedStudent.Email = student.Email
	}

	updatedStudent, err := db.UpdateStudent(fetchedStudent)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update student",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student updated successfully",
		"data":    updatedStudent,
	})
}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := db.DeleteStudent(id); err != nil {
		if err.Error() == "STUDENT_NOT_FOUND" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Student not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete student",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student deleted successfully",
	})
}

func GenerateStudentSummary(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := db.GetStudent(id)
	if err != nil {
		if err.Error() == "STUDENT_NOT_FOUND" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Student not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate student summary",
		})
	}

	summary, err := summarizer.SummarizeStudent(student)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate student summary",
		})
	}

	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"message": "Student summary generated successfully",
		"data":    summary,
	})
}
