package main

import "fmt"

// valor zero: valor zero de todos campos { ,0}
type pessoa struct {
	nome  string
	idade uint8
}

func main() {
	var joao pessoa = pessoa{nome: "João"}
	joao.nome = "João Sipauba"
	joao.idade = 21
	fmt.Println(joao)
	fmt.Println(joao.nome)
	fmt.Println(joao.idade)

	teste := pessoa{"João", 21}
	fmt.Println(teste)
}
