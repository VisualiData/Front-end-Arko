package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"fmt"
)

type Result struct {
	Data string
}

type Response struct {
	Code int `json:"statuscode"`
	Data[] Sensor `json:"data"`
	Message string `json:"message"`
	Status string `json:"status"`
}

type Sensor struct {
	ID string `json:"sensor_id"`
	Type string `json:"type"`
	NodeName string `json:"nodeName"`
	NodeType string `json:"nodeType"`
	Location Position `json:"position"`
}

type Position struct {
	X string `json:"x"`
	Y string `json:"y"`
	Floor string `json:"floor"`
	House string `json:"house"`
}

func get_data(url string) *Response{
	client :=  &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "dev_test")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		//return &Result{Data: ""}
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	var response = new(Response)
	err = json.Unmarshal(resp_body, response)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	//fmt.Print(response)
	//return &Result{Data: string(resp_body)}
	return response
}

func post_data(data []byte, url string) *Response {
	print(bytes.NewBuffer(data))
	client :=  &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Authorization", "dev_test")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		//return &Result{Data: ""}
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	var response = new(Response)
	err = json.Unmarshal(resp_body, response)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	return response
}