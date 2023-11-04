package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID      uuid.UUID `gorm:"type:char(36);primary_key" json:"uuid,omitempty"`
	Email     string    `gorm:"varchar(20);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"varchar(255);not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(scope *gorm.DB) (err error) {
	user.UUID = uuid.NewV4()
	return
}

type CreateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
