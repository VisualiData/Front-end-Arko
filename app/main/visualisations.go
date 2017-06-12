package main

import (
	"net/http"
	"fmt"
	"html/template"
	"github.com/gorilla/mux"
	"strings"
)

func multiLineGraph(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var results []interface{}
	url := BaseUrl + "/house/CHIBB/0"
	if (vars["floor"] != "" && vars["house"] != ""){
		url = BaseUrl + "/house/" + vars["house"] + "/" + vars["floor"]
	}
	for _, element := range strings.Split(vars["sensors"], ","){
		sensorUrl := BaseUrl + "/sensor/"+ element + "/" + vars["from"] + "/" + vars["to"] + "/" + vars["types"]
		results = append(results, getData(sensorUrl).Data)
	}
	t, err := template.New("index.html").Funcs(template.FuncMap{"marshal": Marshal}).ParseFiles("app/resources/index.html", "app/resources/includes/nav.html", "app/resources/includes/message.html", "app/resources/pages/multilinegraph.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	data := map[string]interface{}{
		"sensors": getData(url).Data,
		"house": getDataSingle(BaseUrl + "/house-info/CHIBB").Data,
		"data": results,
		"floor": vars["floor"],
		"selectedSensors": vars["sensors"],
	}
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: data,
	}
	t.Execute(w, vd)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	get_sensors_url := BaseUrl + "/house/CHIBB"
	get_house_url := BaseUrl + "/house-info/CHIBB"
	if vars["house"] != "" {
		get_sensors_url = BaseUrl + "/house/" + vars["house"]
		get_house_url = BaseUrl + "/house-info/" + vars["house"]
	}

	house := getDataSingle(get_house_url)
	sensors := getData(get_sensors_url)
	t, err := template.New("index.html").Funcs(template.FuncMap{"marshal": Marshal}).
		ParseFiles("app/resources/index.html", "app/resources/includes/nav.html", "app/resources/includes/message.html", "app/resources/pages/dashboard.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	values := []interface{}{0: sensors.Data, 1: house.Data}
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: values,
	}
	t.Execute(w, vd)
}
