package main

import (
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"time"
	"log"
	"encoding/gob"
)
// store for flash messages
var store = sessions.NewCookieStore([]byte("V!sual1D@ta"))

type ViewData struct {
	Flash[] FlashMessage
	Data    interface{}
}
// 404 response
func notFound(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("app/resources/index.html", "app/resources/includes/nav.html", "app/resources/includes/message.html", "app/resources/error/404.html")
	str := r.Method + " " + r.URL.Path;
	vd := ViewData{
		Flash: getFlashMessages(w, r),
		Data: str,
	}
	t.Execute(w, vd)
}
// main function, registering FlashMessage, and defining server variables
func main(){
	gob.Register(&FlashMessage{})
	srv := &http.Server{
		Addr: "0.0.0.0:6500",
		Handler: routes(),
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
