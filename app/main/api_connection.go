package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

type Result struct {
	Data string
}

func get_data(url string) *Result{
	fmt.Println(url)
	client :=  &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "dev_test")
	resp, err := client.Do(req)
	log.Print("before-before")
	if err != nil {
		log.Print("before-Err")
		log.Print(err)
		log.Print("after-Err")
		return &Result{Data: ""}
	}
	log.Print("before-close")
	defer resp.Body.Close()
	log.Print("before")
	resp_body, err := ioutil.ReadAll(resp.Body)
	log.Print("after")
	return &Result{Data: string(resp_body)}
	//return string(resp_body)
}