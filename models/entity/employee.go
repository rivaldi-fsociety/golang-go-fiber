package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id           uint   `json:"id" gorm:"primarykey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Birthday     time.Time
	Age          uint8
	MemberNumber string `json:"member_number"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
