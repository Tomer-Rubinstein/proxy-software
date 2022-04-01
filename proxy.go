package main

import (
	"fmt"
	"net/http"
	"html"
	"io"
	"os"
	"time"
	"encoding/json"
)

type Config struct {
	Logging	string
	Restrict []string
}
var config Config


func main(){
	readConfig("config.json")
	http.HandleFunc("/", sendRequest)

	http.ListenAndServe(":8000", nil)
}


func sendRequest(w http.ResponseWriter, req *http.Request){
	if config.Logging == "true" {
		log(req.Host, req.RemoteAddr)
	}

	for _, s := range config.Restrict {
		if s == req.Host {
			fmt.Fprintf(w, "Website cannot be accessed :(")
			return
		}
	}

	resp, err := http.Get(html.EscapeString("http://" + req.Host))
	check(err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)

	fmt.Println(string(body)) // DEBUG
	fmt.Fprintf(w, string(body))
}


/*
	func log(url, ip) logs the requested url from the proxy server
	it appends the string(URL + IP + DATE) to the log.txt file
	params:
		- url string, the url the user as requested from the proxy server
		- ip string, the IP address of the user
	returns: null
	outputs: string(URL + IP + DATE) >> log.txt
*/ 
func log(url string, ip string) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	date := time.Now()
	_, err = f.WriteString(url+"\t" + ip+"\t" + date.Format("01-02-2006 15:04:05")+"\n");
	check(err)
}


// simple func for error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func readConfig reads & parses config.json to the config(typeof Config) struct
func readConfig(filename string) {
	data, err := os.ReadFile(filename)
	check(err)

	err = json.Unmarshal(data, &config)
	check(err)
}