package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"-"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	response, error := json.Marshal(conta)
	if error != nil {
		println(error)
	}
	println(string(response))

	error = json.NewEncoder(os.Stdout).Encode(conta)
	if error != nil {
		println(error)
	}

	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	error = json.Unmarshal(jsonPuro, &contaX)
	if error != nil {
		println(error)
	}

	println(contaX.Saldo)

}
