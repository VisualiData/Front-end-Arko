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
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	return &Result{Data: string(resp_body)}
	//return string(resp_body)
}