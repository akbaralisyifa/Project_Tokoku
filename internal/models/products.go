package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name          string
	Price         int
	Stock         int
	EmployeeID    int
	TransProducts TransProducts `gorm:"foreignKey:ProductID"`
	StockRecepits StockRecepits `gorm:"foreignKey:ProductID"`
}

type StockRecepits struct {
	gorm.Model
	ProductID     int
	EmployeeID    int
	OldStock      int
	IncomingStock int
	TotalStock    int
}

type ProductModel struct {
	db *gorm.DB
}

func NewProductModel(connection *gorm.DB) *ProductModel {
	return &ProductModel{
		db: connection,
	}
}

// add product
func (pm *ProductModel) AddProduct(newProduct Products) (Products, error) {

	query := pm.db.Create(&newProduct)

	if query.Error != nil {
		return Products{}, query.Error
	}

	return newProduct, nil
}

// edit product
func (pm *ProductModel) EditProduct(newStock int, productID int) error {

	var product Products

	err := pm.db.Model(&product).Where("id = ? ", productID).Update("Stock", newStock).Error

	return err
}


// getAllProduct
func (pm *ProductModel) GetAllProduct() []Products{
	var getProduct []Products;

	err := pm.db.Find(&getProduct).Error;

	if err != nil{
		fmt.Println(err)
	}

	return getProduct;
}