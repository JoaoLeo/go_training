package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	profissao string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	pessoa1 := pessoa{"teste", 21, "dev"}
	estudante := estudante{pessoa1, "computação", "unb"}
	fmt.Println(estudante)
	fmt.Println(estudante.nome)
}
