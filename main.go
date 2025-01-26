package main

import (
	"log"
	"net/http"
	"produk_go/config"
	"produk_go/controller/categoriescontroller"
	"produk_go/controller/homecontroller"
)

func main() {
	config.ConnectDB()

	//1. sama dengan fungsi route (homepage)
	http.HandleFunc("/", homecontroller.Welcome)

	//categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	log.Println("server running on port 5000")
	http.ListenAndServe(":5000", nil)
}