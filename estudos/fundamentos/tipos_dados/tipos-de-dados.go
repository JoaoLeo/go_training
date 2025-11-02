package main

import (
	"errors"
	"fmt"
)

func main() {
	// numeros inteiros - suporta 8, 16, 32 e 64 bits
	var teste int8 = 1

	// int - usa a arquitetura do computador como base
	var numero int = 1000
	numero2 := 1000

	// uint - unsygned int; numero sem sinal; apenas positivo
	var numero3 uint8 = 1

	// alias - rune = int32
	var numero4 rune = 1234

	// alias - byte = uint8
	var numero5 byte = 123

	fmt.Println(teste, numero, numero2, numero3, numero4, numero5)

	// numeros reais
	var numeroReal1 float32 = 1.5
	var numeroReal2 float64 = 1.5
	numeroReal3 := 1.6

	fmt.Println(numeroReal1, numeroReal2, numeroReal3)

	// strings
	var str string = "Teste"
	str2 := "Teste2"
	fmt.Println(str, str2)

	// char - valor tabela ASCI
	char := '%'
	fmt.Println(char)

	// bool
	var ligado bool = true
	fmt.Println(ligado, !ligado)

	// error
	var erro error = errors.New("texto de erro")
	var erro2 error
	fmt.Println(erro, erro2)

	// valor zero - valor atribuido a uma variavel sem valor; string = ""; int = 0; bool = false; error = <nil>
	var textoZero string
	fmt.Println(textoZero)

}
