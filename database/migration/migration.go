package migration

import (
	"fmt"

	"gihub.com/rivaldi-fsociety/golang-go-fiber/database"
	"gihub.com/rivaldi-fsociety/golang-go-fiber/models"
)

func RunMigrate() {
	err := database.DB.AutoMigrate(&models.Employee{})
	if err != nil {
		panic(err)
	}
	err = database.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to Migrate")
}
