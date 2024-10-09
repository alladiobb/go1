package main

import (
	"context"
	"fmt"
)

func main() {
	ctxt := context.Background()
	ctxt = context.WithValue(ctxt, "token", "senha")
	bookHotel(ctxt, "Hotel")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
}
