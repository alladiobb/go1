package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	client := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Alladio"}`))
	respose, error := client.Post("http://google.com", "application/json", jsonVar)
	if error != nil {
		panic(error)
	}

	defer respose.Body.Close()
	io.CopyBuffer(os.Stdout, respose.Body, nil)
}
