package controllers

import (
	"log"
	"time"

	"gihub.com/rivaldi-fsociety/golang-go-fiber/database"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/models"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	payload := new(models.Auth)

	if err := ctx.BodyParser(payload); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(payload)
	if errValidate != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user models.User
	err := database.DB.First(&user, "email = ?", payload.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	isValid := utils.VerifyPassword(payload.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Email or Password Wrong!",
		})
	}

	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
