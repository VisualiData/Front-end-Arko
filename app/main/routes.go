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

	r.HandleFunc("/heatmap", Heatmap)
	// login routes
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/login", check_login).Methods("POST")
	// dashboard routes
	r.HandleFunc("/dashboard", dashboard).Name("dashboard")
	r.HandleFunc("/dashboard/{house}", dashboard)
	// sensor routes
	r.HandleFunc("/sensor/add", AddSensorView).Methods("GET").Name("sensorAdd")
	r.HandleFunc("/sensor/add", AddSensor).Methods("POST")
	r.HandleFunc("/sensor/edit", EditSensor).Methods("POST")
	r.HandleFunc("/sensor/edit/{sensor_id}", EditSensorView).Methods("GET").Name("sensorEdit")
	r.NotFoundHandler = http.HandlerFunc(notFound)
	return r
}
