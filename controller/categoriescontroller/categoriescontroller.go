package categoriescontroller

import (
	"html/template"
	"net/http"
	"produk_go/entities"
	"produk_go/models/categorymodel"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request){
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}
	temp, err := template.ParseFiles("views/category/index.html")
	if err != err {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request){
	//menampilkan form add
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	//proses add
	if r.Method == "POST" {
		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		categorymodel.Create(category)
	}
}

func Edit(w http.ResponseWriter, r *http.Request){

}

func Delete(w http.ResponseWriter, r *http.Request){

}