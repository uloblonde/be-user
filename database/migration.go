package database

import (
	"fmt"
	"livecode/models"
	mysql "livecode/pkg"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")

}
