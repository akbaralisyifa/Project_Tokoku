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

	fmt.Print("Username: ")
	fmt.Scanln(&employee.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&employee.Password)
	fmt.Println()

	if employee.Username == "admin" && employee.Password == "admin" {
		employee.AdminAccess = true
		return employee, true
	}

	loginData, isLogin := uc.model.Login(employee)
	return loginData, isLogin
}

func (uc *UsersController) Register() {
	var employee models.Employees

	fmt.Print("Username: ")
	fmt.Scanln(&employee.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&employee.Password)
	fmt.Print("Admin Access (0/1): ")
	fmt.Scanln(&employee.AdminAccess)

	uc.model.Register(employee)
}
