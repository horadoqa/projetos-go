package main

import (
	"fmt"
	"reflect"
)

// var primeiro_numero int = 15
// var segundo_numero int = 15
// var total_plus int = primeiro_numero + segundo_numero
// var total_less int = primeiro_numero - segundo_numero
// var total_times int = primeiro_numero * segundo_numero
// var total_divide int = primeiro_numero / segundo_numero
// var idade, ano int = 51, 1971
// var versao float32 = 1.1
// var nome string = "Fahham"


// Pode remover a tipagem
// var primeiro_numero = 15
// var segundo_numero = 15
// var total_plus = primeiro_numero + segundo_numero
// var total_less = primeiro_numero - segundo_numero
// var total_times = primeiro_numero * segundo_numero
// var total_divide = primeiro_numero / segundo_numero
// var idade, ano = 51, 1971
// var versao = 1.1
// var nome = "Fahham"




func main() {
	
	primeiro_numero := 15
	segundo_numero := 15
	total_plus := primeiro_numero + segundo_numero
	total_less := primeiro_numero - segundo_numero
	total_times := primeiro_numero * segundo_numero
	total_divide := primeiro_numero / segundo_numero
	idade, ano := 53, 1971
	versao := 1.22
	nome := "Fahham"


	fmt.Println("")
	fmt.Println("A soma de", primeiro_numero, "+", segundo_numero, "=", total_plus)
	fmt.Println("A subtração de", primeiro_numero, "-", segundo_numero, "=", total_less)
	fmt.Println("A Multiplicação de", primeiro_numero, "X", segundo_numero, "=", total_times)
	fmt.Println("A divisão de", primeiro_numero, "/", segundo_numero, "=", total_divide)
	fmt.Println("Tenho", idade, "anos, nasci em", ano)
	fmt.Println("Este programa está na versão:", versao)
	fmt.Println("Meu nome é", nome)
	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))

	fmt.Println("")
}
