package routes

import (
	"errors"

	"github.com/Jose-P-C/DevOps-interview/database"
	"github.com/Jose-P-C/DevOps-interview/models"
	"github.com/gofiber/fiber/v2"
)

type Course struct {
	// not model for Course, see as serializer
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func CreateResponseCourse(courseModel models.Course) Course {
	return Course{ID: courseModel.ID, Name: courseModel.Name, Status: courseModel.Status}
}

func CreateCourse(c *fiber.Ctx) error {
	var course models.Course

	if err := c.BodyParser(&course); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&course)
	responseCourse := CreateResponseCourse(course)

	return c.Status(201).JSON(responseCourse)
}

func GetCourses(c *fiber.Ctx) error {
	courses := []models.Course{}

	database.Database.Db.Find(&courses)
	responseCourses := []Course{}

	for _, course := range courses {
		responseCourse := CreateResponseCourse(course)
		responseCourses = append(responseCourses, responseCourse)
	}

	return c.Status(200).JSON(responseCourses)
}

func findUser(id int, course *models.Course) error {
	database.Database.Db.Find(&course, "id = ?", id)

	if course.ID == 0 {
		return errors.New("Course does not exist")
	}
	return nil
}

func GetCourse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var course models.Course

	if err != nil {
		return c.Status(400).JSON("Please verify the course id")
	}

	if err := findUser(id, &course); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseCourse := CreateResponseCourse(course)
	return c.Status(200).JSON(responseCourse)
}

func UpdateCourse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var course models.Course

	if err != nil {
		return c.Status(400).JSON("Please verify the course id")
	}

	if err := findUser(id, &course); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	type UpdateCourse struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}

	var updateData UpdateCourse

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	course.Name = updateData.Name
	course.Status = updateData.Status

	database.Database.Db.Save(&course)

	responseCourse := CreateResponseCourse(course)
	return c.Status(200).JSON(responseCourse)
}

func DeleteCourse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var course models.Course

	if err != nil {
		return c.Status(400).JSON("Please verify the course id")
	}

	if err := findUser(id, &course); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&course).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(204).SendString("Succesfully deleted course")
}
