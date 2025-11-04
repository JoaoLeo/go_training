package main

import "fmt"

func main() {

	// if else comum
	numero := -5

	if numero%2 == 0 {
		fmt.Println("par")
	} else {
		fmt.Println("impar")
	}

	// if init
	// variavel numero2 fica somente no escopo do if/else
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Maior que zero")
	} else {
		fmt.Println("Menor")
	}

}
