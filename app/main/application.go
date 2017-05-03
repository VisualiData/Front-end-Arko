package main

import (
	"html/template"
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func home(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("dist/index.html")
	t.Execute(w, r)
}

func test(w http.ResponseWriter, r *http.Request){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:4567/house/CHIBB/1", nil)
	req.Header.Add("Authorization", "dev_test")
	resp, err := client.Do(req)
	if err != nil{
		log.Print(err)
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	t, _ := template.ParseFiles("dist/index.html")
	t.Execute(w, r)
}

func main() {
	fs := http.FileServer(http.Dir("dist/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/test", test)
	http.ListenAndServe(":6500", nil)
}
