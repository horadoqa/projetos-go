package main

import (
	"fmt"
)

func main() {
	nome, idade := nomeIdade()
	fmt.Println("Meu nome é", nome, "e tenho", idade, "anos")
}

func nomeIdade() (string, int) {
	nome := "Ricardo Fahham"
	idade := 53
	return nome, idade
}