package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
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
	now_time := now.UTC().Format(time.RFC3339)
	then_time := then.UTC().Format(time.RFC3339)
	fmt.Println(now)
	vars := mux.Vars(r)
	fmap := template.FuncMap{
		"marshal": func(v interface {}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
		}}
	// should be dynamic
	url := BaseUrl + "/sensor/CHIBB-Test-01/" + then_time + "/" + now_time + "/Temperature"
	if vars["sensor_id"] != "" {
		url = BaseUrl + "/sensor/"+ vars["sensor_id"]+"/2017-05-27T00:00:00/2017-05-29T13:15:00/Temperature"
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
	url := BaseUrl + "/house/CHIBB"
	result := getData(url)

	t, err := template.New("index.html").Funcs(fmap).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/dashboard.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}

	t.Execute(w, result)
}

func house(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	url := BaseUrl + "/house/" + vars["house"]
	if vars["floor"] != "" {
		//fmt.Println(vars["floor"])
		url = url + "/" + vars["floor"]
	} else {
		fmt.Println("No floor")
	}
	result := getData(url)
	//result := "test"
	fmt.Println(result.Data)
	t, _ := template.ParseFiles("dist/index.html")
	t.Execute(w, r)
}

func add_sensor_view(w http.ResponseWriter, r *http.Request) {
	//response := &Response{200, nil, "Toegevoegd", "success"}
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
	//t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/editsensor.html")
	//t.Execute(w, result)
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

func check_login(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/pages/login.html")
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("password"))
	t.Execute(w, r)
}
func main(){
	r := mux.NewRouter()
	// exclude route matching of assets folder
	fs := http.FileServer(http.Dir("dist/assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.HandleFunc("/", home)
	r.HandleFunc("/home/{sensor_id}", home)
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/login", check_login).Methods("POST")

	r.HandleFunc("/dashboard", dashboard)
	r.HandleFunc("/sensor/add", add_sensor_view).Methods("GET")
	r.HandleFunc("/sensor/add", add_sensor).Methods("POST")
	r.HandleFunc("/sensor/edit/{sensor_id}", edit_sensor_view).Methods("GET")
	r.HandleFunc("/sensor/edit", edit_sensor).Methods("POST")
	//r.HandleFunc("/floorplan/{house}", house).Methods("GET")
	//r.HandleFunc("/floorplan/{house}/{floor}", house).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":6500", nil)
}
