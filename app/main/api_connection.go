package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"bytes"
)

type Result struct {
	Data string
}

type Sensor struct {
	ID string `json:"sensor_id"`
	Type string `json:"type"`
	NodeName string `json:"nodeName"`
	NodeType string `json:"nodeType"`
	Location Position `json:"position"`
}

type Position struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
	Floor string `json:"floor"`
}

func get_data(url string) *Result{
	client :=  &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "dev_test")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return &Result{Data: ""}
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	return &Result{Data: string(resp_body)}
}

func post_data(data []byte, url string)  {
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
	//resp_body, err := ioutil.ReadAll(resp.Body)
	//return &Result{Data: string(resp_body)}
}