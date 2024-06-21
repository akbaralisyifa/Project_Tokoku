package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tokoku/internal/models"
)

type MembersController struct {
	model *models.MembersModel
}

func NewMembersController(m *models.MembersModel) *MembersController {
	return &MembersController{
		model: m,
	}
}

func (mc *MembersController) ManageMembers() {
	var input int = -1

	for input != 0 {
		input = -1
		fmt.Println("===== Kelola Data Member =====")
		fmt.Println()
		memberList := mc.model.ReadAllMembers()

		if len(memberList) == 0 {
			fmt.Println("===== Data Member Tidak Tersedia =====")
			fmt.Println()
		} else {
			fmt.Println("Daftar Member:")
			for _, val := range memberList {
				fmt.Printf("%v | %v | %v | %v\n", val.ID, val.Name, val.Phone, val.Address)
			}
			fmt.Println()
		}

		fmt.Println("1. Tambah | 2. Edit | 3. Hapus | 0. Kembali")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == 1 {
			mc.CreateMember()
		} else if input == 2 {
			mc.UpdateMember()
		} else if input == 3 {
			mc.DeleteMember()
		}
	}
}

func (mc *MembersController) CreateMember() {
	var member models.Members

	fmt.Println("===== Tambah Member Baru =====")

	for len(member.Name) == 0 {
		fmt.Print("Nama: ")
		name := bufio.NewScanner(os.Stdin)
		name.Scan()

		if len(name.Text()) == 0 {
			fmt.Printf("\nNama tidak boleh kosong!\n")
		} else {
			member.Name = name.Text()
		}
	}

	for len(member.Phone) == 0 {
		fmt.Print("No. HP: ")
		fmt.Scanln(&member.Phone)

		if len(member.Phone) == 0 {
			fmt.Printf("\nNo. HP tidak boleh kosong!\n")
		}
	}

	for len(member.Address) == 0 {
		fmt.Print("Alamat: ")
		address := bufio.NewScanner(os.Stdin)
		address.Scan()

		if len(address.Text()) == 0 {
			fmt.Printf("\nAlamat tidak boleh kosong!\n")
		} else {
			member.Address = address.Text()
		}
	}

	fmt.Println()
	mc.model.CreateMember(member)
}

func (mc *MembersController) UpdateMember() {
	var member models.Members
	var memberID int

	fmt.Println("===== Edit Data Member =====")
	fmt.Println("Masukkan '0' untuk kembali.")
	fmt.Println()
	fmt.Print("Masukkan ID Member: ")
	fmt.Scanln(&memberID)

	if memberID != 0 {
		member = mc.model.ReadMember(memberID)

		if member.ID != 0 {
			fmt.Println("Length memberID:", len(strconv.Itoa(memberID)))
			fmt.Println("Kosongkan input jika data tidak akan dirubah.")
			fmt.Printf("\nNama Baru [%v]: ", member.Name)
			name := bufio.NewScanner(os.Stdin)
			name.Scan()

			if len(name.Text()) != 0 {
				member.Name = name.Text()
			}

			fmt.Printf("No. HP Baru [%v]: ", member.Phone)
			fmt.Scanln(&member.Phone)

			fmt.Printf("Alamat Baru [%v]: ", member.Address)
			address := bufio.NewScanner(os.Stdin)
			address.Scan()

			if len(address.Text()) != 0 {
				member.Address = address.Text()
			}

			mc.model.UpdateMember(member)
		} else {
			fmt.Printf("Member tidak ditemukan!\n\n")
		}
	}

	fmt.Println()
}

func (mc *MembersController) DeleteMember() {
	var member models.Members
	var memberID, input int

	fmt.Println("===== Hapus Data Member =====")
	fmt.Println("Masukkan '0' untuk kembali.")
	fmt.Println()
	fmt.Print("Masukkan ID Member: ")
	fmt.Scanln(&memberID)
	fmt.Println()

	if memberID != 0 {
		member = mc.model.ReadMember(memberID)

		if member.ID != 0 {
			fmt.Printf("\nData Member [%v] %v akan DIHAPUS!\nMasukkan '1' untuk konfirmasi, '0' untuk membatalkan.\n", member.ID, member.Name)
			fmt.Print("Konfirmasi: ")
			fmt.Scanln(&input)

			if input == 1 {
				mc.model.DeleteMember(member)
			}
		} else {
			fmt.Printf("Member tidak ditemukan!\n\n")
		}
	}
}
