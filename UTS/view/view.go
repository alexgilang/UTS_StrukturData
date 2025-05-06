package view

import (
	"UTS/model"
	"UTS/node"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InsertKota() {
	var nama string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Kota: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	kota := node.Kota{
		NamaKota: nama,
		Stasiun:  nil,
	}

	model.CreateKota(kota)
	fmt.Println("== Kota berhasil ditambahkan ==")
	fmt.Println()
}

func ViewsKota() {
	fmt.Println("Daftar Kota")
	for i, kota := range model.ReadKota() {
		fmt.Println("--- Kota ke -", i+1, " ---")
		fmt.Println("Nama Kota\t: ", kota.NamaKota)
		fmt.Println("Daftar Stasiun:")
		if kota.Stasiun == nil {
			fmt.Println("  - Tidak ada stasiun.")
		} else {
			temp := kota.Stasiun
			for j := 0; temp != nil; j++ {
				fmt.Printf("  - Stasiun ke - %d : %s, Alamat Stasiun: %s, Notelp: %s \n", j+1, temp.Data.Nama, temp.Data.Alamat, temp.Data.Notelp)
				temp = temp.Link
			}
		}
		fmt.Println()
	}
}

func UpdateKota() {
	var namaKota string
	var nama string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Kota yang akan diupdate: ")
	namaKota, _ = reader.ReadString('\n')
	namaKota = strings.TrimSpace(namaKota)

	fmt.Print("Masukkan Nama Kota baru: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	kota := node.Kota{
		NamaKota: nama,
		Stasiun:  nil,
	}

	cek := model.UpdateKota(namaKota, kota)
	if cek {
		fmt.Println("== Kota berhasil diupdate ==")
	} else {
		fmt.Println("Kota gagal diupdate")
	}
	fmt.Println()
}

func DeleteKota() {
	var namaKota string
	fmt.Print("Masukkan Nama Kota yang akan dihapus: ")
	reader := bufio.NewReader(os.Stdin)
	namaKota, _ = reader.ReadString('\n')
	namaKota = strings.TrimSpace(namaKota)

	cek := model.DeleteKota(namaKota)
	if cek {
		fmt.Println("== Kota berhasil dihapus ==")
	} else {
		fmt.Println("Kota gagal dihapus")
	}
	fmt.Println()
}

func InsertStasiun() {
	var kotaName string
	var stasiunNama string
	var alamat string
	var notelp string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Kota untuk menambah Stasiun: ")
	kotaName, _ = reader.ReadString('\n')
	kotaName = strings.TrimSpace(kotaName)

	fmt.Print("Masukkan Nama Stasiun: ")
	stasiunNama, _ = reader.ReadString('\n')
	stasiunNama = strings.TrimSpace(stasiunNama)

	fmt.Print("Masukkan Alamat Stasiun: ")
	alamat, _ = reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)

	fmt.Print("Masukkan No Telp Stasiun: ")
	notelp, _ = reader.ReadString('\n')
	notelp = strings.TrimSpace(notelp)

	stasiun := node.Stasiun{
		Nama:   stasiunNama,
		Alamat: alamat,
		Notelp: notelp,
	}

	cek := model.AddStasiunToKota(kotaName, stasiun)
	if cek {
		fmt.Println("== Stasiun berhasil ditambahkan ke Kota ==")
	} else {
		fmt.Println("Kota tidak ditemukan. Pastikan nama kota benar.")
	}
	fmt.Println()
}

func DeleteStasiun() {
	var kotaName string
	var stasiunName string

	fmt.Print("Masukkan Nama Kota untuk menghapus Stasiun: ")
	reader := bufio.NewReader(os.Stdin)
	kotaName, _ = reader.ReadString('\n')
	kotaName = strings.TrimSpace(kotaName)

	fmt.Print("Masukkan Nama Stasiun yang akan dihapus: ")
	stasiunName, _ = reader.ReadString('\n')
	stasiunName = strings.TrimSpace(stasiunName)

	cek := model.DeleteStasiunFromKota(kotaName, stasiunName)
	if cek {
		fmt.Println("== Stasiun berhasil dihapus dari Kota ==")
	} else {
		fmt.Println("Stasiun gagal dihapus. Pastikan nama stasiun dan kota benar.")
	}
	fmt.Println()
}
