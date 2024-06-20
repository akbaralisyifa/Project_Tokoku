package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type StockReceipts struct {
	gorm.Model
	ProductID     int
	Supplier      string
	OldStock      int
	IncomingStock int
	TotalStock    int
	EmployeeID    int
}

type AllReceipts struct {
	ID            int
	Date          time.Time
	ProductID     int
	ProductName   string
	Supplier      string
	OldStock      int
	IncomingStock int
	TotalStock    int
	EmployeeName  string
}

type ReceiptsModel struct {
	db *gorm.DB
}

func NewReceiptsModel(connection *gorm.DB) *ReceiptsModel {
	return &ReceiptsModel{
		db: connection,
	}
}

func (rm *ReceiptsModel) ReadProduct(productID int) Products {
	var product Products
	err := rm.db.Where("id = ?", productID).First(&product).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return product
}

func (rm *ReceiptsModel) CreateReceipt(receipt StockReceipts, product Products) {
	err := rm.db.Create(&receipt).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		err = rm.db.Save(&product).Error

		if err != nil {
			fmt.Printf("%v\n\n", err)
		} else {
			fmt.Printf("\nResi penerimaan barang berhasil ditambahkan!\n\n")
		}
	}
}

func (rm *ReceiptsModel) ReadAllReceipts() []AllReceipts {
	var receipts []AllReceipts

	query := rm.db.Raw("select s.id, s.created_at as date, p.id as product_id, p.name as product_name, s.supplier, s.old_stock, s.incoming_stock, s.total_stock, e.username as employee_name from stock_receipts s join products p on p.id = s.product_id join employees e on e.id = s.employee_id where s.deleted_at is null")
	err := query.Scan(&receipts).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return receipts
}

func (rm *ReceiptsModel) ReadReceipt(receiptID int) StockReceipts {
	var receipt StockReceipts
	err := rm.db.Where("id = ?", receiptID).First(&receipt).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return receipt
}

func (rm *ReceiptsModel) DeleteReceipt(receipt StockReceipts, product Products) {
	err := rm.db.Delete(&receipt).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		err = rm.db.Save(&product).Error

		if err != nil {
			fmt.Printf("%v\n\n", err)
		} else {
			fmt.Printf("\n%v (%v) sebanyak %v Pcs telah diretur ke %v\n\n", product.Name, receipt.CreatedAt.Format("02-01-2006"), receipt.IncomingStock, receipt.Supplier)
		}
	}
}
