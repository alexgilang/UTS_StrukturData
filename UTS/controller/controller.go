package controller

import (
	"UTS/model"
	"UTS/node"
	"encoding/json"
	"net/http"
)

// Fungsi untuk mendapatkan daftar kota
func GetCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cities := model.ReadKota() // Mengambil data kota
	json.NewEncoder(w).Encode(cities)
}

// Fungsi untuk menambahkan kota
func AddCity(w http.ResponseWriter, r *http.Request) {
	var kota node.Kota
	if err := json.NewDecoder(r.Body).Decode(&kota); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.CreateKota(kota)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kota)
}

// Fungsi untuk mengupdate kota
func UpdateCity(w http.ResponseWriter, r *http.Request) {
	var updateData struct {
		NamaKota string    `json:"nama_kota"`
		Kota     node.Kota `json:"kota"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	success := model.UpdateKota(updateData.NamaKota, updateData.Kota)
	if success {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updateData.Kota)
	} else {
		http.Error(w, "Kota tidak ditemukan", http.StatusNotFound)
	}
}

// Fungsi untuk menghapus kota
func DeleteCity(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		NamaKota string `json:"nama_kota"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	success := model.DeleteKota(requestData.NamaKota)
	if success {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Kota tidak ditemukan", http.StatusNotFound)
	}
}
