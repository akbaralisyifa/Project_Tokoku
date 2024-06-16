package models

import "gorm.io/gorm"

type TransHistories struct { // Transaksi
	gorm.Model
	EmployeeID    int
	MemberID      int
	GrandTotal    int
	TransProducts TransProducts `gorm:"foreignKey:TransID"`
}

type TransProducts struct { // Detail transaksi
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
