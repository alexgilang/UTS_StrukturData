package model

import "UTS/node"

func AddStasiunToKota(namaKota string, stasiun node.Stasiun) bool {
	temp := &DaftarKota
	for temp.Link != nil {
		if temp.Link.Data.NamaKota == namaKota {
			if temp.Link.Data.Stasiun == nil {
				temp.Link.Data.Stasiun = &node.ListStasiun{Data: stasiun, Link: nil}
			} else {
				stasiunList := temp.Link.Data.Stasiun
				for stasiunList.Link != nil {
					stasiunList = stasiunList.Link
				}
				stasiunList.Link = &node.ListStasiun{Data: stasiun, Link: nil}
			}
			return true
		}
		temp = temp.Link
	}
	return false
}

func DeleteStasiunFromKota(kotaName string, stasiunName string) bool {
	temp := &DaftarKota
	for temp.Link != nil {
		if temp.Link.Data.NamaKota == kotaName {
			if temp.Link.Data.Stasiun == nil {
				return false
			}

			if temp.Link.Data.Stasiun.Data.Nama == stasiunName {
				temp.Link.Data.Stasiun = temp.Link.Data.Stasiun.Link
				return true
			}

			stasiunList := temp.Link.Data.Stasiun
			for stasiunList.Link != nil {
				if stasiunList.Link.Data.Nama == stasiunName {
					stasiunList.Link = stasiunList.Link.Link
					return true
				}
				stasiunList = stasiunList.Link
			}
			return false
		}
		temp = temp.Link
	}
	return false
}
