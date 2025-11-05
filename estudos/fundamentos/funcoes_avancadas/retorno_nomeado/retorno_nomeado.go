package main

import "fmt"

// retorna as variaveis descritas no retorno
func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	n1, n2 := calculos(5, 6)
	fmt.Println(n1, n2)
}
