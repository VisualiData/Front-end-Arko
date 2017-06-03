package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func run(){
	r := mux.NewRouter()
	// exclude route matching of assets folder
	fs := http.FileServer(http.Dir("dist/assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	// start page
	r.HandleFunc("/", home)

	r.HandleFunc("/home/{sensor_id}", home)
	// login routes
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/login", check_login).Methods("POST")
	r.HandleFunc("/dashboard", dashboard)
	r.HandleFunc("/dashboard/{house}", dashboard)
	r.HandleFunc("/sensor/add", add_sensor_view).Methods("GET")
	r.HandleFunc("/sensor/add", add_sensor).Methods("POST")
	r.HandleFunc("/sensor/edit/{sensor_id}", edit_sensor_view).Methods("GET").Name("sensorEdit")
	r.HandleFunc("/sensor/edit", edit_sensor).Methods("POST")
	//r.HandleFunc("/floorplan/{house}", house).Methods("GET")
	//r.HandleFunc("/floorplan/{house}/{floor}", house).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":6500", nil)
}
