package controllers

import (
	"strconv"
	"strings"
	"time"

	"gihub.com/rivaldi-fsociety/golang-go-fiber/database"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShowAllEmployee(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var employee []models.Employee
	results := database.DB.Limit(intLimit).Offset(offset).Find(&employee)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(employee), "employee": employee})
}

func CreateEmployee(c *fiber.Ctx) error {
	payload := new(models.CreateEmployee)
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(payload); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"message": "Error Validation",
			"error":   err.Error(),
		})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newEmployee := models.Employee{
		Name:  payload.Name,
		Email: payload.Email,
	}

	result := database.DB.Create(&newEmployee)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}
	return c.JSON(fiber.Map{
		"message": "Employee Created Successfully",
		"data":    newEmployee,
	})
}

func UpdateEmployee(c *fiber.Ctx) error {
	employeeId := c.Params("id")

	payload := new(models.UpdateEmployee)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var employee models.Employee
	result := database.DB.First(&employee, "uuid = ?", employeeId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Employee with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Name != "" {
		updates["name"] = payload.Name
	}
	if payload.Email != "" {
		updates["email"] = payload.Email
	}
	updates["updated_at"] = time.Now()

	database.DB.Model(&employee).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"employee": employee}})
}

func GetByIdEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee

	result := database.DB.First(&employee, "uuid = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Employee with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"employee": employee}})
}

func DeleteEmployee(c *fiber.Ctx) error {
	employeeId := c.Params("id")

	result := database.DB.Delete(&models.Employee{}, "uuid = ?", employeeId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "failed", "message": "No Employee with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.JSON(fiber.Map{
		"message": "Employee Deleted Successfully",
	})
}
