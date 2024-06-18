package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type EmployeesController struct {
	model *models.EmployeesModel
}

func NewEmployeesController(m *models.EmployeesModel) *EmployeesController {
	return &EmployeesController{
		model: m,
	}
}

func (ec *EmployeesController) Login() (models.Employees, bool) {
	var employee models.Employees

	fmt.Println("===== Login Pegawai =====")
	fmt.Print("Username: ")
	fmt.Scanln(&employee.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&employee.Password)
	fmt.Println()

	if employee.Username == "admin" && employee.Password == "admin" {
		employee.AdminAccess = true
		return employee, true
	}

	loginData, isLogin := ec.model.Login(employee)
	return loginData, isLogin
}

func (ec *EmployeesController) ManageEmployees(loginData models.Employees) {
	var input int = -1

	for input != 0 {
		fmt.Println("===== Kelola Data Pegawai =====")
		fmt.Printf("Username: [%v] %v | Admin Access: %v\n\n", loginData.ID, loginData.Username, loginData.AdminAccess)

		employeeList := ec.model.ReadAllEmployees()

		if len(employeeList) == 0 {
			fmt.Println("===== Data Pegawai Tidak Tersedia =====")
			fmt.Println()
		} else {
			fmt.Println("Daftar Pegawai")
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
			ec.CreateEmployee()
		} else if input == 2 {
			ec.UpdateEmployee()
		} else if input == 3 {
			ec.DeleteEmployee()
		}
	}
}

func (ec *EmployeesController) CreateEmployee() {
	var employee models.Employees

	fmt.Println("===== Tambah Pegawai Baru =====")
	fmt.Print("Username: ")
	fmt.Scanln(&employee.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&employee.Password)
	fmt.Print("Email: ")
	fmt.Scanln(&employee.Email)
	fmt.Print("No. HP: ")
	fmt.Scanln(&employee.Phone)
	fmt.Print("Admin Access (0/1): ")
	fmt.Scanln(&employee.AdminAccess)
	fmt.Println()

	ec.model.CreateEmployee(employee)
}

func (ec *EmployeesController) UpdateEmployee() {
	var employee models.Employees
	var employeeID int

	fmt.Println("===== Edit Data Pegawai =====")
	fmt.Print("Masukkan ID Pegawai: ")
	fmt.Scanln(&employeeID)

	employee = ec.model.ReadEmployee(employeeID)

	fmt.Println("Kosongkan input jika data tidak akan dirubah.")
	fmt.Printf("Username Baru [%v]: ", employee.Username)
	fmt.Scanln(&employee.Username)
	fmt.Printf("Password Baru [%v]: ", employee.Password)
	fmt.Scanln(&employee.Password)
	fmt.Printf("Email Baru [%v]: ", employee.Email)
	fmt.Scanln(&employee.Email)
	fmt.Printf("No. HP Baru [%v]: ", employee.Phone)
	fmt.Scanln(&employee.Phone)
	fmt.Printf("Admin Access Baru (0/1) [%v]: ", employee.AdminAccess)
	fmt.Scanln(&employee.AdminAccess)
	fmt.Println()

	ec.model.UpdateEmployee(employee)
}

func (ec *EmployeesController) DeleteEmployee() {
	var employee models.Employees
	var employeeID, input int

	fmt.Println("===== Hapus Data Pegawai =====")
	fmt.Print("Masukkan ID Pegawai: ")
	fmt.Scanln(&employeeID)

	employee = ec.model.ReadEmployee(employeeID)

	fmt.Printf("\nData Pegawai [%v] %v akan DIHAPUS!\nMasukkan '1' untuk konfirmasi, '0' untuk membatalkan.\n", employee.ID, employee.Username)
	fmt.Print("Konfirmasi: ")
	fmt.Scanln(&input)

	if input == 1 {
		ec.model.DeleteEmployee(employee)
	}
}
