package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctxt := context.Background()
	ctxt, cancel := context.WithTimeout(ctxt, time.Second*2)
	defer cancel()

	bookHotel(ctxt)
}

func bookHotel(ctxt context.Context) {
	select {
	case <-ctxt.Done():
		fmt.Println("Compra não realizada. Limite de tempo excedido!")
		return
	case <-time.After(3 * time.Second):
		fmt.Println("Compra realizada!")
	}
}
