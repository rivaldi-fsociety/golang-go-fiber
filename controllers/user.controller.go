package controllers

import (
	"log"
	"strings"

	"gihub.com/rivaldi-fsociety/golang-go-fiber/database"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/models"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	payload := new(models.CreateUser)
	if err := ctx.BodyParser(payload); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(payload); err != nil {
		return ctx.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"message": "Error Validation",
			"error":   err.Error(),
		})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
		})
	}

	newUser := models.User{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	result := database.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Email already exist, please use another email"})
	} else if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}
	return ctx.JSON(fiber.Map{
		"message": "User Created Successfully",
		"data":    newUser,
	})
}
