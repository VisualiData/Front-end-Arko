package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
	"log"
)

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	// Add whatever other types you need
	default:
		return ""
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	then := now.AddDate(0, 0, -2)
	now_time := now.Format(time.RFC3339)
	then_time := then.Format(time.RFC3339)
	vars := mux.Vars(r)
	fmap := template.FuncMap{
		"marshal": func(v interface {}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
		}}
	// should be dynamic
	url := BaseUrl + "/sensor/CHIBB-Test-01/" + then_time + "/" + now_time + "/Temperature"
	if vars["sensor_id"] != "" {
		url = BaseUrl + "/sensor/"+ vars["sensor_id"]+"/" + then_time + "/" + now_time + "/Temperature"
	}
	result := getSensorData(url)
	t, err := template.New("index.html").Funcs(fmap).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/home.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}

	t.Execute(w, result)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	fmap := template.FuncMap{
		"marshal": func(v interface {}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
	}}
	vars := mux.Vars(r)
	url := BaseUrl + "/house/CHIBB"
	if vars["house"] != "" {
		url = BaseUrl + "/house/" + vars["house"]
	}

	result := getData(url)
	if len(result.Data) > 0 {
		t, err := template.New("index.html").Funcs(fmap).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/dashboard.html")
		if err != nil {
			fmt.Fprint(w, "Error:", err)
			fmt.Println("Error:", err)
			return
		}

		t.Execute(w, result)
	}else {
		notFound(w, r)
	}
}

func house(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	url := BaseUrl + "/house/" + vars["house"]
	if vars["floor"] != "" {
		url = url + "/" + vars["floor"]
	} else {
		fmt.Println("No floor")
	}
	result := getData(url)
	fmt.Println(result.Data)
	t, _ := template.ParseFiles("dist/index.html")
	t.Execute(w, r)
}

func add_sensor_view(w http.ResponseWriter, r *http.Request) {
	response := &Response{0, nil, "", ""}
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/addsensor.html")
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
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/addsensor.html")
	t.Execute(w, response)
}

func edit_sensor_view(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := BaseUrl + "/sensor/" + vars["sensor_id"]
	result := getDataSingle(url)
	t, err := template.New("index.html").Funcs(template.FuncMap{"tostring": ToString}).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/editsensor.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}

	t.Execute(w, result)
}

func edit_sensor(w http.ResponseWriter, r *http.Request) {
	location := Position{r.FormValue("x_coordinate"), r.FormValue("y_coordinate"), r.FormValue("floor"), "CHIBB"}
	fmt.Println(r.FormValue("x_coordinate"));
	fmt.Println(r.FormValue("y_coordinate"));
	s := Sensor{r.FormValue("sensor_id"), r.FormValue("sensorType"), r.FormValue("nodeName"), r.FormValue("nodeType"), location, "active"}
	b, err := json.Marshal(s)
	if err != nil {
		print(err)
	}
	response := post_data(b, BaseUrl + "/sensor/update")
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/editsensor.html")
	t.Execute(w, response)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/login.html")
	t.Execute(w, r)
}

func notFound(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/error/404.html")
	t.Execute(w, r)
}

func check_login(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/login.html")
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("password"))
	t.Execute(w, r)
}
func main(){
	//http.Handle("/", routes())
	//http.ListenAndServe(":6500", nil)
	srv := &http.Server{
		Handler: routes(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
