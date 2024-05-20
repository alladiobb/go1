package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Saldo  int `json:"saldo"`
	Numero int `json:"N"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 200}

	resultado, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(resultado))

	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(conta)

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"N":2,"S":200}`)
	var conta2 Conta
	err = json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		println(err)
	}
	println(conta2.Saldo)

}
