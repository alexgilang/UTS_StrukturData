package model

import "UTS/node"

var DaftarKota node.ListKota

func CreateKota(kota node.Kota) {
	tempLL := node.ListKota{
		Data: kota,
		Link: nil,
	}
	if DaftarKota.Link == nil {
		DaftarKota.Link = &tempLL
	} else {
		temp := &DaftarKota
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
	}
}

func ReadKota() []node.Kota {
	daftarKota := []node.Kota{}
	temp := &DaftarKota
	for temp.Link != nil {
		daftarKota = append(daftarKota, temp.Link.Data)
		temp = temp.Link
	}
	return daftarKota
}

func UpdateKota(namaKota string, kota node.Kota) bool {
	temp := DaftarKota.Link
	for temp != nil {
		if temp.Data.NamaKota == namaKota {
			temp.Data = kota
			return true
		}
		temp = temp.Link
	}
	return false
}

func DeleteKota(namaKota string) bool {
	temp := &DaftarKota
	for temp.Link != nil {
		if temp.Link.Data.NamaKota == namaKota {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}
