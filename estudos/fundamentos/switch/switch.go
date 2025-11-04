package main

import "fmt"

// so cai em um caso do switch caso nao haja fallthrough
func diaSemana(num int) string {
	switch num {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabado"
	default:
		return "opção invalida"
	}
}

func diaSemana2(num int) string {
	switch {
	case num == 1:
		return "domingo"
	case num == 2:
		return "segunda"
	case num == 3:
		return "terça"
	case num == 4:
		return "quarta"
	case num == 5:
		return "quinta"
	case num == 6:
		return "sexta"
	case num == 7:
		return "sabado"
	default:
		return "opção invalida"
	}
}

func diaSemana3(num int) string {
	var diaSemana string
	switch {
	case num == 1:
		diaSemana = "domingo"
		fallthrough // joga o codigo para a proxima condição
	case num == 2:
		diaSemana = "segunda"
	case num == 3:
		diaSemana = "terça"
	case num == 4:
		diaSemana = "quarta"
	case num == 5:
		diaSemana = "quinta"
	case num == 6:
		diaSemana = "sexta"
	case num == 7:
		diaSemana = "sabado"
	default:
		diaSemana = "opção invalida"
	}
	return diaSemana
}

func main() {
	dia := diaSemana(5)
	dia2 := diaSemana2(6)
	dia3 := diaSemana3(6)
	fmt.Println(dia, dia2, dia3)
}
