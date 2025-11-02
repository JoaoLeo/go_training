package main

import "fmt"

// funcoes sao tipos, enquanto pode ser usado em condicoes e passar funcoes como parametro de funcoes

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// mais de um retorno / declarando tipo dos parametros somente no fim caso sejam iguais
func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	resultado := soma(2, 2)
	fmt.Println(resultado)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Teste f")

	resultado1, resultado2 := calculos(5, 6)
	resultado3, _ := calculos(5, 6) // _ quando não é necessario usar a variavel
	fmt.Println(resultado1, resultado2, resultado3)

}
