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

	MainMenu(connection)
}

func MainMenu(connection *gorm.DB) {
	fmt.Println("================== TOKOKU APP PROJECT ===================")
	fmt.Println("By: Muhammad Akbar Ali Syifa & Muhammad Farhan Adriansyah")
	fmt.Println()

	em := models.NewEmployeesModel(connection)
	ec := controllers.NewEmployeesController(em)
	mm := models.NewMembersModel(connection)
	mc := controllers.NewMembersController(mm)
	// tm := models.NewTransModel(connection)
	// tc := controllers.NewTransController(tm)
	pm := models.NewProductModel(connection);
	pc := controllers.NewProductController(pm)


	var input int = -1

	for input != 0 {
		fmt.Println("===== Main Menu =====")
		fmt.Println("1. Login")
		fmt.Println("9. Migrate Database")
		fmt.Println("0. Keluar")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == 1 {
			loginData, isLogin := ec.Login()
			Dashboard(loginData, isLogin, ec, mc, pc)
		} else if input == 9 {
			err := connection.AutoMigrate(&models.Employees{}, &models.Members{}, &models.Products{}, &models.TransHistories{}, &models.TransProducts{}, &models.StockRecepits{})

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

func Dashboard(loginData models.Employees, isLogin bool, ec *controllers.EmployeesController, mc *controllers.MembersController, pc *controllers.ProductController) {
	var input int = -1

	for input != 0 && isLogin {
		input = -1

		fmt.Println("===== Dashboard =====")
		fmt.Printf("Username: [%v] %v | Admin Access: %v\n\n", loginData.ID, loginData.Username, loginData.AdminAccess)
		fmt.Println("1. Transaksi")
		fmt.Println("2. Kelola Data Produk")
		fmt.Println("3. Kelola Data Member")

		if loginData.AdminAccess {
			fmt.Println("4. Kelola Data Pegawai")
		}

		fmt.Println("0. Logout")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == 1 {
			fmt.Println("TRANSAKSI BERLANGSUNG")
			fmt.Println()
		} else if input == 2 {
			pc.ManageProduct(loginData)
		} else if input == 3 {
			mc.ManageMembers()
		} else if input == 4 {
			if loginData.AdminAccess {
				ec.ManageEmployees()
			} else {
				fmt.Printf("User %v tidak memiliki Admin Access!\n\n", loginData.Username)
			}
		} else if input == 0 {
			isLogin = false
		}
	}
}
