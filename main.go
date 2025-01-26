package main

import (
	"log"
	"net/http"
	"produk_go/config"
	"produk_go/controller/categoriescontroller"
	"produk_go/controller/homecontroller"
	"produk_go/controller/productcontroller"
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

	//products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("server running on port 5000")
	http.ListenAndServe(":5000", nil)
}