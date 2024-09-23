package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	client := http.Client{}
	request, error := http.NewRequest("GET", "http://google.com")

	if error != nil {
		panic(error)
	}

	request.Header.Set("Accept","application/json")
	response , error := client.Do(request)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	
	error != nil {
		panic(errerror)
	}

	println(string(body))
	
}
