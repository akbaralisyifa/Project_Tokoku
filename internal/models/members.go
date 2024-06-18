package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Members struct {
	gorm.Model
	Name           string
	Phone          string
	Address        string
	TransHistories TransHistories `gorm:"foreignKey:MemberID"`
}

type MembersModel struct {
	db *gorm.DB
}

func NewMembersModel(connection *gorm.DB) *MembersModel {
	return &MembersModel{
		db: connection,
	}
}

func (mm *MembersModel) CreateMember(members Members) {
	err := mm.db.Create(&members).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nMember %v berhasil ditambahkan!\n\n", members.Name)
	}
}

func (mm *MembersModel) ReadAllMembers() []Members {
	var members []Members
	err := mm.db.Find(&members).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return members
}

func (mm *MembersModel) ReadMember(memberID int) Members {
	var member Members
	err := mm.db.Where("id = ?", memberID).First(&member).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	}

	return member
}

func (mm *MembersModel) UpdateMember(member Members) {
	err := mm.db.Save(&member).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nMember %v berhasil diedit!\n\n", member.Name)
	}
}

func (mm *MembersModel) DeleteMember(member Members) {
	err := mm.db.Delete(&member).Error

	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("\nUser %v berhasil dihapus!\n\n", member.Name)
	}
}
