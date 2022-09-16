package database

import (
	"fmt"
	"golangfnl/models"
	"golangfnl/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Product{}, &models.Cart{}, &models.Profile{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Berhasil Luur")
}
