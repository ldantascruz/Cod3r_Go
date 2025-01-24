package main

import "fmt"

func main() {

	// APPEND PODE FAZER O SLICE CRESCER, O COPY NÃO PODE FAZER ISSO
	array1 := [3]int{1, 2, 3}

	var slice1 []int

	//array1 = append(array1, 4, 5, 6) // esse código não é possível, pois ele espera que passe um slice, não um array
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
