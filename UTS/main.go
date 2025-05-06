package main

import (
	"UTS/controller"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/api/cities", controller.GetCities)
	http.HandleFunc("/api/cities/add", controller.AddCity)
	http.HandleFunc("/api/cities/update", controller.UpdateCity)
	http.HandleFunc("/api/cities/delete", controller.DeleteCity)

	log.Println("Server Berjalan di : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
