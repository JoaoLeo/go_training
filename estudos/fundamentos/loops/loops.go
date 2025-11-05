package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		time.Sleep(time.Second)
		fmt.Println(j)
	}

	nomes := [3]string{"joao", "sipauba", "debian"}

	for index, nome := range nomes {
		fmt.Println(index, nome)
	}

	for index, letra := range "PALAVRA" {
		fmt.Println(index, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Joao",
		"sobrenome": "Sipauba",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
