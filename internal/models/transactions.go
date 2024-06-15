package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name          string
	Price         int
	Stock         int
	EmployeeID    int
	TransProducts TransProducts `gorm:"foreignKey:ProductID"`
}

type TransHistories struct {
	gorm.Model
	EmployeeID    int
	MemberID      int
	GrandTotal    int
	TransProducts TransProducts `gorm:"foreignKey:TransID"`
}

type TransProducts struct {
	gorm.Model
	TransID   int
	ProductID int
	Quantity  int
	SubTotal  int
}

type TransModel struct {
	db *gorm.DB
}

func NewTransModel(connection *gorm.DB) *TransModel {
	return &TransModel{
		db: connection,
	}
}
