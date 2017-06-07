package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"html/template"
)

func add_sensor_view(w http.ResponseWriter, r *http.Request) {
	response := &Response{0, nil, "", ""}
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/addsensor.html")
	t.Execute(w, response)
}

func add_sensor(w http.ResponseWriter, r *http.Request) {
	location := Position{r.FormValue("x_coordinate"), r.FormValue("y_coordinate"), r.FormValue("floor"), "CHIBB"}
	s := Sensor{r.FormValue("sensor_id"), r.FormValue("sensorType"), r.FormValue("nodeName"), r.FormValue("nodeType"), location, "active"}
	b, err := json.Marshal(s)
	if err != nil {
		print(err)
	}
	response := post_data(b, BaseUrl + "/sensor")
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/addsensor.html")
	t.Execute(w, response)
}

func edit_sensor_view(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := BaseUrl + "/sensor/" + vars["sensor_id"]
	result := getDataSingle(url)
	t, err := template.New("index.html").Funcs(template.FuncMap{"tostring": ToString}).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/editsensor.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: result,
	}
	t.Execute(w, vd)
}

func edit_sensor(w http.ResponseWriter, r *http.Request) {
	location := Position{r.FormValue("x_coordinate"), r.FormValue("y_coordinate"), r.FormValue("floor"), "CHIBB"}
	s := Sensor{r.FormValue("sensor_id"), r.FormValue("sensorType"), r.FormValue("nodeName"), r.FormValue("nodeType"), location, "active"}
	b, err := json.Marshal(s)
	if err != nil {
		print(err)
	}
	post_data(b, BaseUrl + "/sensor/update")
	addFlashMessage(w, r, FlashMessage{Message: "sensor updated", Type: "success"})
	url, err := mux.CurrentRoute(r).Subrouter().Get("sensorEdit").URL("sensor_id", r.FormValue("sensor_id"))
	http.Redirect(w, r, url.String(), 302)
}