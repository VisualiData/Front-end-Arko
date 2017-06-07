package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
	"log"
	"encoding/gob"
)

var store = sessions.NewCookieStore([]byte("V!sual1D@ta"))

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

func Marshal(value interface{}) template.JS {
	a, _ := json.Marshal(value)
	return template.JS(a);
}

type ViewData struct {
	Flash[] FlashMessage
	Data    interface{}
}

func home(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	current_time := time.Now().Format(time.RFC3339)
	two_days_ago := time.Now().AddDate(0, 0, -2).Format(time.RFC3339)

	url := BaseUrl + "/sensor/CHIBB-Test-01/" + two_days_ago + "/" + current_time + "/Temperature"
	if vars["sensor_id"] != "" {
		url = BaseUrl + "/sensor/"+ vars["sensor_id"]+"/" + two_days_ago + "/" + current_time + "/Temperature"
	}
	result := getData(url)
	t, err := template.New("index.html").Funcs(template.FuncMap{"marshal": Marshal}).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/home.html")
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
	if len(sensors.Data) > 0 {
		t, err := template.New("index.html").Funcs(template.FuncMap{"marshal": Marshal}).
				ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/dashboard.html")
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
	}else {
		notFound(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/login.html")
	t.Execute(w, r)
}

func notFound(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/error/404.html")
	str := r.Method + " " + r.URL.Path;
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: str,
	}
	t.Execute(w, vd)
}

func check_login(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/login.html")
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("password"))
	t.Execute(w, r)
}
func main(){
	gob.Register(&FlashMessage{})
	srv := &http.Server{
		Addr: "192.168.0.18:6500",
		Handler: routes(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
