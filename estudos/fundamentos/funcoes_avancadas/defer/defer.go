package main

import "fmt"

func funcao1() {
	fmt.Println("1")
}

func funcao2() {
	fmt.Println("2")
}

// defer adia a execução da função ate ultimo momento possivel
func main() {
	defer funcao1()
	funcao2()
}
