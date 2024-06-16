package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name          string
	Price         int
	Stock         int
	EmployeeID    int
	TransProducts TransProducts `gorm:"foreignKey:ProductID"`
};

type ProductModel struct{
	db *gorm.DB
};

func NewProductModel(connection *gorm.DB) *ProductModel {
	return &ProductModel{
		db : connection,
	}
}

// add product
func (pm *ProductModel) AddProduct(newProduct Products)(Products, error){

	query := pm.db.Create(&newProduct);

	if query.Error != nil {
		return Products{}, query.Error
	}

	return newProduct, nil
}

// edit product
func(pm *ProductModel) EditProduct(newStock int, employeeID, productID uint)(error){

	err := pm.db.Model(&Products{}).Where("id = ? AND EmployeeID = ?", productID, employeeID).Update("Stock", newStock).Error;

	return err;
}




