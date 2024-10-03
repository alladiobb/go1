package main

import (
	"io/ioutil"
	"net/http"
)

func main() {

	client := http.Client{}
	request, error := http.NewRequest("GET", "http://google.com", nil)

	if error != nil {
		panic(error)
	}

	request.Header.Set("Accept", "application/json")
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	println(string(body))

}
