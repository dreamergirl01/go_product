package homecontroller

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request){
	//panggil file index
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}