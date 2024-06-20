package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type TransController struct {
	model *models.TransModel
	productModel *models.ProductModel
}

func NewTransController(m *models.TransModel, productModel *models.ProductModel) *TransController {
	return &TransController{
		model: m,
		productModel: productModel,
	}
}


func (tc *TransController) ManageTransaction(loginData models.Employees){
	var inputMenu int = -1

	if inputMenu != 0 {
		fmt.Println("==== KELOLA DATA TRANSAKSI ====");
		fmt.Println(loginData.ID)
        fmt.Println("1. Tambah | 2. Data Transaksi | 0. Kembali")
		fmt.Print("Input menu :");
		fmt.Scanln(&inputMenu);

		switch inputMenu {
		case 1:
			_, err := tc.AddTransHistori(int(loginData.ID));

			getOneHistory, err := tc.model.GetOneTransHistory();

			fmt.Println("Data Terakhir History :", )

			if err != nil {
				fmt.Println("Error to create transaction history:", err)
			}

			success, err := tc.AddTransaction(int(getOneHistory.ID));

			if err != nil {
			    fmt.Println("Error to create Transaction :");
			}

			if success {
				fmt.Println("Transaction Successfull !")
			}
		
		case 2:
			dataTransaksi := tc.model.GetAllTransaction();

			if len(dataTransaksi) == 0 {
				fmt.Println("===== Data Transaksi Tidak Tersedia =====")
				fmt.Println()
			} else {
				fmt.Println("Daftar Transaksi:")
				for _, val := range dataTransaksi {
					fmt.Printf("%v | %v | %v | %v | %v\n", val.ID, val.TransID, val.ProductID, val.Quantity, val.SubTotal)
				}
				fmt.Println()
			}
		case 0:
			return
		}
	}
}

func(tc *TransController) AddTransHistori(employeeID int)(models.TransHistories, error){
	var IdMember int;

	fmt.Println("==== create history ====");
	fmt.Print("Input Member ID :");
	fmt.Scanln(&IdMember);

	newHistory := models.TransHistories{
		EmployeeID: employeeID,
		MemberID: IdMember,
		GrandTotal: 0,
	};

	_, err := tc.model.AddTransHistori(newHistory);

	if err != nil {
		return newHistory, err;
	}
	return newHistory, nil;
}

func(tc *TransController) AddTransaction(transID int) (bool, error) {

	if tc.productModel == nil {
        return false, fmt.Errorf("product model is not initialized")
    }

	var productID int;
	var qnt int;

	
	fmt.Println("==== CREATE TRANSACTIN ====");
	fmt.Print("Input Product ID :");
	fmt.Scanln(&productID);
	fmt.Print("Input Quantity :")
	fmt.Scanln(&qnt);

	product, err := tc.productModel.GetOneProduct(productID)
    if err != nil {
        fmt.Printf("Error fetching product: %v\n", err)
        return false, err
    }
	 
	var newTrans = models.TransProducts{
		TransID: transID,
		ProductID: productID,
		Quantity: qnt,
		SubTotal: product.Price * qnt,
	}

	_, err = tc.model.AddTransaction(newTrans)

	if err != nil {
		return false, nil;
	}

	if product.Stock < qnt {
		panic("Trasaksi Gagal, Data Stok Kurang !!")
	} 

    err = tc.model.UpdateGrandTotal(transID, newTrans.SubTotal);

	if err != nil {
		fmt.Println("Error updating grand total:", err)
		return false, err
	}

	err = tc.productModel.UpdateStockProduct(productID, qnt);

	if err != nil {
		return false, err;
	}

	fmt.Println("Transaction successfully added.")

	return true, nil;
}
