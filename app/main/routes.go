package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func routes() *mux.Router{
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
	// dashboard routes
	r.HandleFunc("/dashboard", dashboard)
	r.HandleFunc("/dashboard/{house}", dashboard)
	// sensor routes
	s := r.PathPrefix("/sensor").Subrouter()
	s.HandleFunc("/add", add_sensor_view).Methods("GET")
	s.HandleFunc("/add", add_sensor).Methods("POST")

	s.HandleFunc("/edit/{sensor_id}", edit_sensor_view).Methods("GET").Name("sensorEdit")
	s.HandleFunc("/edit", edit_sensor).Methods("POST")
	//r.NotFoundHandler = http.HandleFunc(notFound)
	//r.HandleFunc("/floorplan/{house}", house).Methods("GET")
	//r.HandleFunc("/floorplan/{house}/{floor}", house).Methods("GET")
	return r
}
