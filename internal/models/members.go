package models

import "gorm.io/gorm"

type Members struct {
	gorm.Model
	Name           string
	Address        string
	Phone          string
	TransHistories TransHistories `gorm:"foreignKey:MemberID"`
}
