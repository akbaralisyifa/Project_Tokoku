package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type ProductController struct {
	model *models.ProductModel
};

func NewProductController(m *models.ProductModel) *ProductController{
	return &ProductController{
		model: m,
	}
};

func(pc *ProductController) AddTodo(id int)(bool, error){
	var newProduct models.Products;

	fmt.Println("==== CREATE PRODUCT ====");
	fmt.Print("Input Product Name : ");
	fmt.Scanln(&newProduct.Name);
	fmt.Print("Input Price :");
	fmt.Scanln(&newProduct.Price);
	fmt.Print("Input Stock Product :");
	fmt.Scanln(&newProduct.Stock);

	newProduct.EmployeeID = id;

	_, err := pc.model.AddProduct(newProduct);

	if err != nil {
		return false, err
	}

	return true, nil
}


func(pc *ProductController) EditProduct(EmployeeID int)(bool, error){
	var productID int;

	fmt.Println("==== UPDATE STOCK PRODUCT ====");
	fmt.Println("Input '0' if you want to cancel changes!")
	fmt.Print("Input Product ID :");
	fmt.Scanln(&productID)

	if productID != 0 {
		return true, nil
	}

	return false, nil;
}
