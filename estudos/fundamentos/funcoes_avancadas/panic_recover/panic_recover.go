package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado")
	}
}

// panic - interrompe o fluxo normal do programa
func aprovado(n1, n2 float64) bool {
	defer recuperar()
	media := (n1 + n1) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("Média 6")
}

func main() {

	fmt.Println(aprovado(6, 6))
}
