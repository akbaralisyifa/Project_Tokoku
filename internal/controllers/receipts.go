package controllers

import (
	"bufio"
	"fmt"
	"os"
	"tokoku/internal/models"
)

type ReceiptsController struct {
	model *models.ReceiptsModel
}

func NewReceiptsController(m *models.ReceiptsModel) *ReceiptsController {
	return &ReceiptsController{
		model: m,
	}
}

func (rc *ReceiptsController) ManageReceipts(loginData models.Employees) {
	var input int = -1

	for input != 0 {
		input = -1
		fmt.Println("===== Kelola Penerimaan Stok =====")
		fmt.Println()

		receipts := rc.model.ReadAllReceipts()

		if len(receipts) == 0 {
			fmt.Println("===== Resi Penerimaan Tidak Tersedia =====")
			fmt.Println()
		} else {
			fmt.Println("Daftar Resi Penerimaan Barang:")
			for _, val := range receipts {
				fmt.Printf("%v | %v | [%v] %v | %v | %v | %v | %v | %v\n", val.ID, val.Date.Format("02-01-2006"), val.ProductID, val.ProductName, val.Supplier, val.OldStock, val.IncomingStock, val.TotalStock, val.EmployeeName)
			}
		}

		fmt.Println()
		fmt.Println("1. Terima Stok Baru | 2. Batalkan Penerimaan Barang | 0. Kembali")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == 1 {
			rc.CreateReceipt(loginData)
		} else if input == 2 {
			rc.DeleteReceipt()
		}
	}
}

func (rc *ReceiptsController) CreateReceipt(loginData models.Employees) {
	var receipt models.StockReceipts
	var product models.Products
	var productID int

	fmt.Println("===== Terima Stok Baru =====")
	fmt.Println("Masukkan '0' untuk kembali.")
	fmt.Println()
	fmt.Print("Masukkan ID Produk: ")
	fmt.Scanln(&productID)

	if productID != 0 {
		product = rc.model.ReadProduct(productID)

		if product.ID != 0 {
			fmt.Printf("\nProduk: %v | Stok Saat Ini: %v\n", product.Name, product.Stock)
			for len(receipt.Supplier) == 0 {
				fmt.Print("Supplier: ")
				supplier := bufio.NewScanner(os.Stdin)
				supplier.Scan()

				if len(supplier.Text()) == 0 {
					fmt.Printf("\nSupplier tidak boleh kosong!\n")
				} else {
					receipt.Supplier = supplier.Text()
				}
			}

			receipt.ProductID = int(product.ID)
			receipt.OldStock = int(product.Stock)

			var incomingStock int
			for incomingStock == 0 {
				fmt.Print("Stok Datang: ")
				fmt.Scanln(&incomingStock)

				if incomingStock == 0 {
					fmt.Printf("\nStok Datang tidak boleh kosong!\n")
				} else {
					receipt.IncomingStock = incomingStock
				}
			}

			receipt.TotalStock = receipt.OldStock + receipt.IncomingStock
			receipt.EmployeeID = int(loginData.ID)
			product.Stock = receipt.TotalStock

			rc.model.CreateReceipt(receipt, product)
		} else {
			fmt.Printf("Produk tidak ditemukan!\n\n")
		}
	}
}

func (rc *ReceiptsController) DeleteReceipt() {
	var receipt models.StockReceipts
	var product models.Products
	var receiptID int
	var input int

	fmt.Println("===== Batalkan Penerimaan Barang =====")
	fmt.Println("Masukkan '0' untuk kembali.")
	fmt.Println()
	fmt.Print("Masukkan ID Resi: ")
	fmt.Scanln(&receiptID)

	if receiptID != 0 {
		receipt = rc.model.ReadReceipt(receiptID)
		product = rc.model.ReadProduct(receipt.ProductID)

		if receipt.ID != 0 {
			fmt.Printf("\nPenerimaan %v (%v) sebanyak %v Pcs akan dibatalkan dan diretur ke %v\nMasukkan '1' untuk konfirmasi, '0' untuk membatalkan.\n", product.Name, receipt.CreatedAt.Format("02-01-2006"), receipt.IncomingStock, receipt.Supplier)
			fmt.Print("Konfirmasi: ")
			fmt.Scanln(&input)

			if input == 1 {
				product.Stock -= receipt.IncomingStock
				rc.model.DeleteReceipt(receipt, product)
			}
		} else {
			fmt.Printf("Resi Penerimaan Barang tidak ditemukan!\n\n")
		}
	}
}
