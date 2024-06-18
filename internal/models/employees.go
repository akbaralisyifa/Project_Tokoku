package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	Username       string
	Password       string
	Email          string
	Phone          string
	AdminAccess    bool
	Products       Products       `gorm:"foreignKey:EmployeeID"`
	TransHistories TransHistories `gorm:"foreignKey:EmployeeID"`
	StockRecepits  StockRecepits  `gorm:"foreignKey:EmployeeID"`
}

type EmployeesModel struct {
	db *gorm.DB
}

func NewEmployeesModel(connection *gorm.DB) *EmployeesModel {
	return &EmployeesModel{
		db: connection,
	}
}

func (em *EmployeesModel) Login(employee Employees) (Employees, bool) {
	var loginData Employees
	err := em.db.Where("username = ? AND password = ?", employee.Username, employee.Password).First(&loginData).Error

	if loginData.ID == 0 {
		fmt.Printf("\nGagal Login! Pastikan Username dan Password sudah benar!\n\n")
		return loginData, false
	} else if err != nil {
		fmt.Printf("%v\n\n", err)
		return loginData, false
	}

	return loginData, true
}

func (em *EmployeesModel) CreateEmployee(employee Employees) {
	err := em.db.Create(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nUser %v berhasil ditambahkan!\n\n", employee.Username)
	}
}

func (em *EmployeesModel) ReadAllEmployees() []Employees {
	var employees []Employees
	err := em.db.Find(&employees).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return employees
}

func (em *EmployeesModel) ReadEmployee(userID int) Employees {
	var employee Employees
	err := em.db.Where("id = ?", userID).First(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return employee
}

func (em *EmployeesModel) UpdateEmployee(employee Employees) {
	err := em.db.Save(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nUser [%v] %v berhasil diedit!\n\n", employee.ID, employee.Username)
	}
}

func (em *EmployeesModel) DeleteEmployee(employee Employees) {
	err := em.db.Delete(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nUser [%v] %v berhasil dihapus!\n\n", employee.ID, employee.Username)
	}
}
