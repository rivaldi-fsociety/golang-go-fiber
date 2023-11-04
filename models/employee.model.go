package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	// "github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Employee struct {
	UUID         uuid.UUID `gorm:"type:char(36);primary_key" json:"uuid,omitempty"`
	Name         string    `gorm:"varchar(30);not null" json:"name"`
	Email        string    `gorm:"varchar(20);uniqueIndex;not null" json:"email"`
	Birthday     time.Time
	Age          uint8
	MemberNumber string `json:"member_number"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (employee *Employee) BeforeCreate(scope *gorm.DB) (err error) {
	employee.UUID = uuid.NewV4()
	return
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type CreateEmployee struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateEmployee struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
