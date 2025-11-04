package main

import "fmt"

func main() {

	fmt.Println("Maps")

	// map - chave valor
	// map[tipoChave]tipoValor
	usuario := map[string]string{
		"nome":      "Joao",
		"Sobrenome": "Sipauba",
	}

	fmt.Println(usuario)              // map[Sobrenome:Sipauba nome:Joao]
	fmt.Println(usuario["nome"])      // Joao
	fmt.Println(usuario["Sobrenome"]) // Sipauba

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Joao",
			"sobrenome": "Sipauba",
		},
		"curso": {
			"nome":         "Computação",
			"universidade": "unb",
		},
	}

	fmt.Println(usuario2)                      // map[curso:map[nome:Computação universidade:unb] nome:map[primeiro:Joao sobrenome:Sipauba]]
	fmt.Println(usuario2["nome"]["primeiro"])  // Joao
	fmt.Println(usuario2["nome"]["Sobrenome"]) // Sipauba

	delete(usuario2, "nome") // deletar chave
	fmt.Println(usuario2)

	usuario["idade"] = "21" // adicionar chave-valor

}
