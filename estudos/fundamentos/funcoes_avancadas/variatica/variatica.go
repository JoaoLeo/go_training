package main

import "fmt"

// slice de int - numeros
// parametro variatico deve ser o ultimo da função e só pode haver um
func soma(texto string, numeros ...int) (soma int) {
	for _, n := range numeros {
		soma += n
	}
	fmt.Println(texto)
	return
}

func main() {
	n1 := soma("teste", 5, 6, 2, 42, 553, 64, 213)
	fmt.Println(n1)
}
