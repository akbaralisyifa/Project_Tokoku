package main

import (
	"fmt"
	"tokoku/config"
	"tokoku/internal/controllers"
	"tokoku/internal/models"

	"gorm.io/gorm"
)

func main() {
	setup := config.InportSetting()
	connection, err := config.ConnectDB(setup)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mainMenu(connection)
}

func mainMenu(connection *gorm.DB) {
	fmt.Println("================== TOKOKU APP PROJECT ===================")
	fmt.Println("By: Muhammad Akbar Ali Syifa & Muhammad Farhan Adriansyah")
	fmt.Println()

	um := models.NewUsersModel(connection)
	uc := controllers.NewUsersController(um)
	// tm := models.NewTransModel(connection)
	// tc := controllers.NewTransController(tm)

	var input int = -1
	var isLogin bool

	for input != 0 {
		fmt.Println("===== Main Menu =====")
		fmt.Println("1. Login")
		fmt.Println("9. Migrate Database")
		fmt.Println("0. Keluar")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == 1 && !isLogin {
			var loginData models.Employees
			loginData, isLogin = uc.Login()
			isLogin = controllers.Dashboard(loginData, uc)

		} else if input == 9 {
			err := connection.AutoMigrate(&models.Employees{}, &models.Members{}, &models.Products{}, &models.TransHistories{}, &models.TransProducts{})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Database Successfully Migrated!")
			}

			fmt.Println()
		}
	}

	fmt.Println("Program Selesai. Terima Kasih!")
}
