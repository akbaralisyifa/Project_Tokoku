package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	Username       string
	Password       string
	AdminAccess    bool
	Products       Products       `gorm:"foreignKey:EmployeeID"`
	TransHistories TransHistories `gorm:"foreignKey:EmployeeID"`
}

type Members struct {
	gorm.Model
	Name           string
	TransHistories TransHistories `gorm:"foreignKey:MemberID"`
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
		fmt.Printf("%v\n", err)
		return loginData, false
	}

	return loginData, true
}

func (um *UsersModel) Register(employee Employees) {
	err := um.db.Create(&employee).Error

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
