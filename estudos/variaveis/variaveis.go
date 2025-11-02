package main

import "fmt"

func main() {
	// tipo explicito
	var nome string = "variavel nome"

	// tipo implicito - inferencia de tipo
	nome2 := "Teste"

	// criar varias variaveis de tipo explicito
	var (
		nome3 string = "teste3"
		nome4 string = "teste4"
		// nome5 := "teste" - erro
	)

	// criando varias variaveis com inferencia de tipo
	nome5, nome6 := "teste5", "teste6 "

	// constante
	const nome7 string = "teste7"

	// inverter valores de variaveis
	nome5, nome6 = nome6, nome5

	fmt.Println(nome, nome2, nome3, nome4, nome5, nome6, nome7)
}
