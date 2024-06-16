package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type UsersController struct {
	model *models.UsersModel
}

func NewUsersController(m *models.UsersModel) *UsersController {
	return &UsersController{
		model: m,
	}
}

func (uc *UsersController) Login() (models.Employees, bool) {
	var employee models.Employees

	fmt.Println("===== LOGIN PEGAWAI =====")
	fmt.Print("Username: ")
	fmt.Scanln(&employee.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&employee.Password)
	fmt.Println()

	if employee.Username == "admin" && employee.Password == "admin" {
		employee.ID = 0
		employee.AdminAccess = true
		return employee, true
	}

	loginData, isLogin := uc.model.Login(employee)
	return loginData, isLogin
}

func (uc *UsersController) ManageEmployees(loginData models.Employees) {
	var input int = -1

	for input != 0 {
		fmt.Println("===== KELOLA DATA PEGAWAI =====")
		fmt.Printf("Username: [%v] %v | Admin Access: %v\n\n", loginData.ID, loginData.Username, loginData.AdminAccess)

		employeeList := uc.model.ReadAllEmployees()

		if len(employeeList) == 0 {
			fmt.Println("===== DATA PEGAWAI TIDAK TERSEDIA =====")
			fmt.Println()
		} else {
			fmt.Println("===== DAFTAR PEGAWAI =====")
			for _, val := range employeeList {
				fmt.Printf("%v | %v | %v | %v | %v | %v\n", val.ID, val.Username, val.Password, val.Email, val.Phone, val.AdminAccess)
			}
			fmt.Println()
		}

		fmt.Println("1. Tambah | 2. Edit | 3. Hapus | 0. Kembali")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == 1 {
			uc.CreateEmployees()
		} else if input == 2 {
			uc.UpdateEmployees()
		} else if input == 3 {
			uc.DeleteEmployees()
		}
	}
}

func (uc *UsersController) CreateEmployees() {
	var employee models.Employees

	fmt.Println("===== TAMBAH PEGAWAI BARU =====")
	fmt.Print("Username: ")
	fmt.Scanln(&employee.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&employee.Password)
	fmt.Print("Email: ")
	fmt.Scanln(&employee.Email)
	fmt.Print("Phone: ")
	fmt.Scanln(&employee.Phone)
	fmt.Print("Admin Access (0/1): ")
	fmt.Scanln(&employee.AdminAccess)
	fmt.Println()

	uc.model.CreateEmployees(employee)
}

func (uc *UsersController) UpdateEmployees() {
	var employee models.Employees
	var userID int

	fmt.Println("===== EDIT DATA PEGAWAI =====")
	fmt.Print("Masukkan ID Pegawai: ")
	fmt.Scanln(&userID)

	employee = uc.model.ReadEmployees(userID)

	fmt.Printf("New Username [%v]: ", employee.Username)
	fmt.Scanln(&employee.Username)
	fmt.Printf("New Password [%v]: ", employee.Password)
	fmt.Scanln(&employee.Password)
	fmt.Printf("New Email [%v]: ", employee.Email)
	fmt.Scanln(&employee.Email)
	fmt.Printf("New Phone [%v]: ", employee.Phone)
	fmt.Scanln(&employee.Phone)
	fmt.Printf("New Admin Access (0/1) [%v]: ", employee.AdminAccess)
	fmt.Scanln(&employee.AdminAccess)
	fmt.Println()

	uc.model.UpdateEmployees(employee)
}

func (uc *UsersController) DeleteEmployees() {
	var employee models.Employees
	var userID, confirm int

	fmt.Println("===== HAPUS DATA PEGAWAI =====")
	fmt.Print("Masukkan ID Pegawai: ")
	fmt.Scanln(&userID)

	employee = uc.model.ReadEmployees(userID)

	fmt.Printf("\nData Pegawai [%v] %v akan DIHAPUS!\nMasukkan '1' untuk konfirmasi, '0' untuk membatalkan.\n", employee.ID, employee.Username)
	fmt.Print("Konfirmasi: ")
	fmt.Scanln(&confirm)

	if confirm == 1 {
		uc.model.DeleteEmployees(employee)
	}
}
