package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"html/template"
	"strings"
)

type Sensor struct {
	ID string `json:"sensor_id"`
	Type[] string `json:"types"`
	NodeName string `json:"nodeName"`
	NodeType string `json:"nodeType"`
	Location Position `json:"position"`
	Status string `json:"status"`
}
type Position struct {
	X string `json:"x"`
	Y string `json:"y"`
	Floor string `json:"floor"`
	House string `json:"house"`
}

func AddSensorView(w http.ResponseWriter, r *http.Request) {
	response := &Response{0, nil, "", ""}
	t, err := template.New("index.html").Funcs(template.FuncMap{"tostring": ToString, "join": Join}).ParseFiles("app/resources/index.html", "app/resources/includes/nav.html", "app/resources/includes/message.html", "app/resources/pages/addsensor.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: response,
	}
	t.Execute(w, vd)
}

func AddSensor(w http.ResponseWriter, r *http.Request) {
	location := Position{r.FormValue("x_coordinate"), r.FormValue("y_coordinate"), r.FormValue("floor"), "CHIBB"}
	s := Sensor{r.FormValue("sensor_id"), strings.Split(r.FormValue("sensorType"), ","), r.FormValue("nodeName"), r.FormValue("nodeType"), location, "active"}
	b, err := json.Marshal(s)
	if err != nil {
		print(err)
	}
	response := post_data(b, BaseUrl + "/sensor")
	addFlashMessage(w, r, FlashMessage{Message: response.Message, Type: response.Status})
	http.Redirect(w, r, "/sensor/add", 302)
}

func EditSensorView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := BaseUrl + "/sensor/" + vars["sensor_id"]
	result := getDataSingle(url)
	t, err := template.New("index.html").Funcs(template.FuncMap{"tostring": ToString, "join": Join}).ParseFiles("app/resources/index.html", "app/resources/includes/nav.html", "app/resources/includes/message.html", "app/resources/pages/editsensor.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: result.Data,
	}
	t.Execute(w, vd)
}

func EditSensor(w http.ResponseWriter, r *http.Request) {
	location := Position{r.FormValue("x_coordinate"), r.FormValue("y_coordinate"), r.FormValue("floor"), "CHIBB"}
	s := Sensor{r.FormValue("sensor_id"), strings.Split(r.FormValue("sensorType"), ","), r.FormValue("nodeName"), r.FormValue("nodeType"), location, "active"}
	fmt.Println(s)
	b, err := json.Marshal(s)
	if err != nil {
		print(err)
	}
	post_data(b, BaseUrl + "/sensor/update")
	addFlashMessage(w, r, FlashMessage{Message: "sensor updated", Type: "success"})
	http.Redirect(w, r, "/sensor/edit/"+r.FormValue("sensor_id"), 302)
}