package main

import "fmt"

func main() {
	func() {
		fmt.Println("Teste")
	}()

	func(txt string) {
		fmt.Println(txt)
	}("Teste 2")

	resultado := func(txt string) string {
		return fmt.Sprintf("Recebido: %s", txt)
	}("Teste 2")

	fmt.Println(resultado)
}
