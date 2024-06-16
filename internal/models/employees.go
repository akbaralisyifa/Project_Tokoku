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
}

type UsersModel struct {
	db *gorm.DB
}

func NewUsersModel(connection *gorm.DB) *UsersModel {
	return &UsersModel{
		db: connection,
	}
}

func (um *UsersModel) Login(employee Employees) (Employees, bool) {
	var loginData Employees
	err := um.db.Where("username = ? AND password = ?", employee.Username, employee.Password).First(&loginData).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
		return loginData, false
	}

	return loginData, true
}

func (um *UsersModel) CreateEmployees(employee Employees) {
	err := um.db.Create(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nUser %v berhasil ditambahkan!\n\n", employee.Username)
	}
}

func (um *UsersModel) ReadAllEmployees() []Employees {
	var employee []Employees
	err := um.db.Find(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return employee
}

func (um *UsersModel) ReadEmployees(userID int) Employees {
	var employee Employees
	err := um.db.Where("id = ?", userID).First(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return employee
}

func (um *UsersModel) UpdateEmployees(employee Employees) {
	err := um.db.Save(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nUser [%v] %v berhasil diedit!\n\n", employee.ID, employee.Username)
	}
}

func (um *UsersModel) DeleteEmployees(employee Employees) {
	err := um.db.Delete(&employee).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nUser [%v] %v berhasil dihapus!\n\n", employee.ID, employee.Username)
	}
}
