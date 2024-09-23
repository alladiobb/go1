package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	client := http.Client{Timeout: time.Second}

	respose, error := client.Get("http://google.com")
	if error != nil {
		panic(error)
	}

	defer respose.Body.Close()

	body, error := ioutil.ReadAll(respose.Body)
	if error != nil {
		panic(error)
	}
	println(string(body))

}
