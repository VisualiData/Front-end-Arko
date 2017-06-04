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

type ViewData struct {
	Flash[] FlashMessage
	Data    interface{}
}

func home(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	now := time.Now().Format(time.RFC3339)
	before := time.Now().AddDate(0, 0, -2).Format(time.RFC3339)

	fmap := template.FuncMap{
		"marshal": func(v interface {}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
		}}
	session, err := store.Get(r, "front-end")

	url := BaseUrl + "/sensor/CHIBB-Test-01/" + before + "/" + now + "/Temperature"
	if vars["sensor_id"] != "" {
		url = BaseUrl + "/sensor/"+ vars["sensor_id"]+"/" + before + "/" + now + "/Temperature"
	}
	result := getSensorData(url)
	t, err := template.New("index.html").Funcs(fmap).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/home.html")
	if err != nil {
		fmt.Fprint(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(session.Flashes())
	session.Save(r, w)
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: result,
	}
	t.Execute(w, vd)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "front-end")
	session.AddFlash("test")
	session.Save(r, w)
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
		t, err := template.New("index.html").Funcs(fmap).ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/dashboard.html")
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



func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/login.html")
	t.Execute(w, r)
}

func notFound(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/error/404.html")
	t.Execute(w, r)
}

func check_login(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("dist/index.html", "dist/includes/nav.html", "dist/includes/message.html", "dist/pages/login.html")
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("password"))
	t.Execute(w, r)
}
func main(){
	gob.Register(&FlashMessage{})
	//http.Handle("/", routes())
	//http.ListenAndServe(":6500", nil)
	srv := &http.Server{
		Addr: "192.168.0.18:6500",
		Handler: routes(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
