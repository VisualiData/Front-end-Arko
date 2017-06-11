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
		fmt.Println(sensorUrl)
		results = append(results, getData(sensorUrl).Data)
	}
	//fmt.Println(strings.Split(vars["sensors"], ","))
	//fmt.Println(url)
	//fmt.Println(getData(url))
	//current_time := time.Now().Format(time.RFC3339)
	//two_days_ago := time.Now().AddDate(0, 0, -2).Format(time.RFC3339)
	//
	//sensorurl := BaseUrl + "/sensor/CHIBB-Test-01/" + vars["from"] + "/" + vars["to"] + "/" + vars["types"]
	//fmt.Println(sensorurl)
	//results = append(results, getData(sensorurl).Data)
	//results = append(results, getData(sensorurl).Data)
	//fmt.Println(results)
	//if vars["sensor_id"] != "" {
	//	url = BaseUrl + "/sensor/"+ vars["sensor_id"]+"/" + two_days_ago + "/" + current_time + "/Temperature"
	//}
	//fmt.Println(getData(sensorurl))
	//result := getData(url)
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
