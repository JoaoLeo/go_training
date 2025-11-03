package main

import (
	"fmt"
	"reflect"
)

func main() {

	// arrays
	var array1 [5]int
	array1[0] = 1
	array2 := [5]int{1, 2, 3, 4, 5}

	array3 := [...]int{2, 3, 3, 45, 52, 3}
	fmt.Println(array1, array2, array3)
	fmt.Println(reflect.TypeOf(array1))

	// slice
	slice1 := []int{1, 4, 53, 64, 64, 23445, 54, 543, 543}
	slice1 = append(slice1, 2)

	// indice 1 ao 2 [inclusivo, exclusivo]
	slice2 := array2[1:3]

	fmt.Println(slice1, slice2)
	fmt.Println(reflect.TypeOf(slice1))

	// arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade - maximo
}
