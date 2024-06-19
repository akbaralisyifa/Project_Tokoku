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


func(pc *ProductController) ManageProduct(loginData models.Employees) {
	var input int = -1

	for input != 0 {
		fmt.Println("===== Kelola Data Product =====")
		fmt.Printf("Username: [%v] %v | Admin Access: %v\n\n", loginData.ID, loginData.Username, loginData.AdminAccess)

		productList := pc.model.GetAllProduct()

		if len(productList) == 0 {
			fmt.Println("===== Data Product Tidak Tersedia =====")
			fmt.Println()
		} else {
			fmt.Println("Daftar Product:")
			for _, val := range productList {
				fmt.Printf("%v | %v | %v | %v\n", val.ID, val.Name, val.Price, val.Stock)
			}
			fmt.Println()
		}

		fmt.Println("1. Tambah | 2. Edit | 0. Kembali")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == 1 {
			pc.AddProduct()
		} else if input == 2 {
			pc.EditProduct()
		}
	}
}

func(pc *ProductController) AddProduct()(bool, error){
	var newProduct models.Products;
	var employeeID int;

	fmt.Println("==== CREATE PRODUCT ====");
	fmt.Print("Input ID Pegawai :");
	fmt.Scanln(&employeeID)
	fmt.Print("Input Product Name : ");
	fmt.Scanln(&newProduct.Name);
	fmt.Print("Input Price :");
	fmt.Scanln(&newProduct.Price);
	fmt.Print("Input Stock Product :");
	fmt.Scanln(&newProduct.Stock);

	newProduct.EmployeeID = employeeID;

	_, err := pc.model.AddProduct(newProduct);

	if err != nil {
		return false, err
	}

	return true, nil
}


func(pc *ProductController) EditProduct()(string, error){
	var productID int;
	var inputStock int;
	var succ = "Data Product Berhasil di Ubah."
	var failed = "Data Product Gagal Di Ubah!"

	productList := pc.model.GetAllProduct();

	if len(productList) == 0 {
		fmt.Println("===== Data Product Tidak Tersedia =====")
		fmt.Println()
	} else {
		fmt.Println("Daftar Product:")
		for _, val := range productList {
			fmt.Printf("%v | %v | %v | %v\n", val.ID, val.Name, val.Price, val.Stock)
		}
		fmt.Println()
	}


	fmt.Println("==== UPDATE STOCK PRODUCT ====");
	fmt.Print("Input Product ID :");
	fmt.Scanln(&productID);
	fmt.Print("Input Stock Product :");
	fmt.Scanln(&inputStock)

	err := pc.model.EditProduct(inputStock, productID)

	if err !=nil {
		fmt.Printf("Error updating product: %v\n", err)
		return failed, err;
	}

	return succ, nil;
}
