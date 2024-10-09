package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// ctxt := context.Background()
	ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	request, error := http.NewRequestWithContext(ctxt, "GET", "http://localhost:8082", nil)
	if error != nil {
		panic(error)
	}
	response, error := http.DefaultClient.Do(request)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}
