package main

import "fmt"

func ternario(condicao bool, txt1 any, txt2 any) any {
	if condicao {
		return txt1
	} else {
		return txt2
	}
}

func main() {
	texto := ternario(8 > 6, 1, "nao")
	fmt.Println(texto)
}
