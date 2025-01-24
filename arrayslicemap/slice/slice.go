package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // ARRAY
	s1 := []int{1, 2, 3}  // SLICE
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// SLICE NÃO É UM ARRAY!!! Ele define um PEDAÇO de um array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slie, mas apontando para o MESMO array
	fmt.Println(a2, s3)

	// Você pode imaginar um slice como: tem um tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1] // criando o slice de um slice
	fmt.Println(s2, s4)
}
