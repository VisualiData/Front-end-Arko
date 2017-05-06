package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("dist/index.html")
	t.Execute(w, r)
}

func test(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("dist/index.html")
	t.Execute(w, r)
}

func house(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	url := "http://localhost:4567/house/" + vars["house"]
	if vars["floor"] != "" {
		//fmt.Println(vars["floor"])
		url = url + "/" + vars["floor"]
	} else {
		fmt.Println("No floor")
	}
	result := get_data(url)
	//result := "test"
	//fmt.Println(result.Data)
	t, _ := template.ParseFiles("dist/index.html")
	t.Execute(w, result)
}
func main(){
	r := mux.NewRouter()
	// exclude matching of assets folder
	fs := http.FileServer(http.Dir("dist/assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.HandleFunc("/", home)
	r.HandleFunc("/test", test)
	r.HandleFunc("/floorplan/{house}", house).Methods("GET")
	r.HandleFunc("/floorplan/{house}/{floor}", house).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":6500", nil)
}
