package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	ctxt := context.Background()
	ctxt, cancel := context.WithTimeout(ctxt, time.Microsecond)
	// ctxt, cancel := context.WithCancel(ctxt)
	defer cancel()

	request, error := http.NewRequestWithContext(ctxt, "GET", "http://google.com", nil)
	if error != nil {
		panic(error)
	}

	response, error := http.DefaultClient.Do(request)
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
