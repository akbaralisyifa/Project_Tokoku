package models

import (
	"fmt"

	"gorm.io/gorm"
)

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

func (tm *TransModel) AddTransHistori(newTransHistory TransHistories)(TransHistories, error){

	err := tm.db.Create(&newTransHistory).Error

	if err != nil {
		return TransHistories{}, err
	}

	return newTransHistory, nil;
}

// get history get one
func (tm *TransModel) GetOneTransHistory()(TransHistories, error){
	var dataHistory TransHistories;

	err := tm.db.Last(&dataHistory).Error;

	if err != nil {
		return TransHistories{}, err;
	}

	return dataHistory, nil;
}

func(tm *TransModel) UpdateGrandTotal(transID int, amount int)error{
	var transHistory TransHistories;

	result := tm.db.First(&transHistory, transID);

	if result.Error != nil {
		return result.Error
	}

	transHistory.GrandTotal += amount;

	return tm.db.Save(&transHistory).Error
}


func (tm *TransModel) AddTransaction(newTrans TransProducts) (TransProducts, error) {
    err := tm.db.Create(&newTrans).Error
	
    if err != nil {
        return TransProducts{}, err
    }

    return newTrans, nil
}

// Get All Transaksi
func(tm *TransModel) GetAllTransaction()[]TransProducts{
	var dataTransaksi []TransProducts;

	err := tm.db.Find(&dataTransaksi).Error;

	if err != nil {
		fmt.Println(err)
	}

	return dataTransaksi;
}
