package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type TransController struct {
	model *models.TransModel
}

func NewTransController(m *models.TransModel) *TransController {
	return &TransController{
		model: m,
	}
}

func Dashboard(loginData models.Employees) bool {
	var input int = -1

	for input != 0 {
		fmt.Println("===== DASHBOARD =====")
		fmt.Printf("Username: %v | Admin Access: %v\n\n", loginData.Username, loginData.AdminAccess)
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

		} else if input == 3 {

		} else if input == 4 {
			if loginData.AdminAccess {
				ManageEmployee()
			} else {
				fmt.Printf("\nUser %v tidak memiliki Admin Access!\n", loginData.Username)
			}
		}
	}

	return false
}

func ManageEmployee() {

}
