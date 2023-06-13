package database

import (
	"fmt"
	"hallo-corona/models"
	"hallo-corona/pkg/postgres"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Consultation{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
