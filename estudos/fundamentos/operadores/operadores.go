package main

import "fmt"

func main() {
	// aritmeticos : + - / * %

	// atribuição
	var variavel1 string = "Ttesteee"
	variavel2 := "testeee"

	fmt.Println(variavel1, variavel2)

	// relacionais
	fmt.Println(3 > 2, 32 <= 64, 6*2 == 12, 5 != 4)

	// logicos
	fmt.Println(true && false, true || false, !false)

	// unarios - interagem com 1 variavel por vez
	num := 10
	num++
	num += 121

	num--
	num -= 32

	num *= 4
	num /= 2
	fmt.Println(num)

	// ternario - nao existe, forma de fazer:
	var texto string
	if num > 5 {
		texto = "maior"
	} else {
		texto = "menor"
	}

	fmt.Println(texto)

}
