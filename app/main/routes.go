package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func routes() *mux.Router{
	r := mux.NewRouter()
	// exclude route matching of assets folder
	fs := http.FileServer(http.Dir("app/resources/assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	// graph page
	r.HandleFunc("/", multiLineGraph)
	r.HandleFunc("/visualisation", multiLineGraph)
	r.HandleFunc("/visualisation/{house}/{floor}", multiLineGraph)
	r.HandleFunc("/visualisation/{house}/{floor}/{sensors}/{types}", multiLineGraph)
	r.HandleFunc("/visualisation/{house}/{floor}/{sensors}/{types}/{from}/{to}", multiLineGraph)
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