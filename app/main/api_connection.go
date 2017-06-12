package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"fmt"
)

var BaseUrl = "http://localhost:4567";
var API_Key = "dev";

type Response struct {
	Code int `json:"statuscode"`
	Data[] interface{} `json:"data"`
	Message string `json:"message"`
	Status string `json:"status"`
}

type ResponseSingle struct {
	Code int `json:"statuscode"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Status string `json:"status"`
}

func getData(url string) *Response{
	client :=  &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", API_Key)
	req.Header.Add("From", "Arko")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
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

func getDataSingle(url string) *ResponseSingle{
	client :=  &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", API_Key)
	req.Header.Add("From", "Arko")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	var responseSingle = new(ResponseSingle)
	err = json.Unmarshal(resp_body, responseSingle)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	return responseSingle
}

func post_data(data []byte, url string) *Response {
	client :=  &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Authorization", API_Key)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("From", "Arko")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
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