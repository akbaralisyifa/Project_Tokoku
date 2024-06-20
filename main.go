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

type AllControl struct {
	ec *controllers.EmployeesController
	mc *controllers.MembersController
	pc *controllers.ProductController
	rc *controllers.ReceiptsController
	tc *controllers.TransController
}

func MainMenu(connection *gorm.DB) {
	fmt.Println("================== TOKOKU APP PROJECT ===================")
	fmt.Println("By: Muhammad Akbar Ali Syifa & Muhammad Farhan Adriansyah")
	fmt.Println()

	em := models.NewEmployeesModel(connection)
	mm := models.NewMembersModel(connection)
	pm := models.NewProductModel(connection)
	rm := models.NewReceiptsModel(connection)
	tm := models.NewTransModel(connection)

	var con AllControl

	con.ec = controllers.NewEmployeesController(em)
	con.mc = controllers.NewMembersController(mm)
	con.pc = controllers.NewProductController(pm)
	con.rc = controllers.NewReceiptsController(rm)
	con.tc = controllers.NewTransController(tm)

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
			loginData, isLogin := con.ec.Login()
			Dashboard(loginData, isLogin, con)
		} else if input == 9 {
			err := connection.AutoMigrate(&models.Employees{}, &models.Members{}, &models.Products{}, &models.TransHistories{}, &models.TransProducts{}, &models.StockReceipts{})

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

func Dashboard(loginData models.Employees, isLogin bool, con AllControl) {
	var input int = -1

	for input != 0 && isLogin {
		input = -1

		fmt.Println("===== Dashboard =====")
		fmt.Printf("Username: [%v] %v | Admin Access: %v\n\n", loginData.ID, loginData.Username, loginData.AdminAccess)
		fmt.Println("1. Transaksi")
		fmt.Println("2. Kelola Data Produk")
		fmt.Println("3. Kelola Penerimaan Stok")
		fmt.Println("4. Kelola Data Member")

		if loginData.AdminAccess {
			fmt.Println("5. Kelola Data Pegawai")
		}

		fmt.Println("0. Logout")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == 1 {
			fmt.Println("TRANSAKSI BERLANGSUNG")
			fmt.Println()
		} else if input == 2 {
			con.pc.ManageProduct(loginData)
		} else if input == 3 {
			con.rc.ManageReceipts(loginData)
		} else if input == 4 {
			con.mc.ManageMembers()
		} else if input == 5 {
			if loginData.AdminAccess {
				con.ec.ManageEmployees(loginData)
			} else {
				fmt.Printf("User %v tidak memiliki Admin Access!\n\n", loginData.Username)
			}
		} else if input == 0 {
			isLogin = false
		}
	}
}
