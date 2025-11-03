package main

import "fmt"

func main() {
	// atribuindo por copia
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	// ponteiro - referencia de memoria
	var variavel3 int = 100
	var variavel4 *int

	variavel4 = &variavel3 // o ponteiro variavel4 recebe endereco de memoria da variavel3

	variavel3 = 150

	fmt.Println(variavel3, *variavel4) // desreferenciacao: pega o valor dentro do endereco de memoria
}
