package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func Heatmap(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//url := BaseUrl + "/sensor/" + vars["sensor_id"]
	//result := getDataSingle(url)
	t, err := template.New("index.html").Funcs(template.FuncMap{"tostring": ToString}).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/heatmap.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: nil,
	}
	t.Execute(w, vd)
}
