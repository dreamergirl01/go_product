package productcontroller

import (
	"html/template"
	"net/http"
	"produk_go/models/productmodel"
)

func Index(w http.ResponseWriter, r *http.Request){
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}
	temp, err := template.ParseFiles("views/product/index.html")
	if err != err {
		panic(err)
	}
	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request){

}
func Add(w http.ResponseWriter, r *http.Request){

}

func Edit(w http.ResponseWriter, r *http.Request){

}

func Delete(w http.ResponseWriter, r *http.Request){
	
}