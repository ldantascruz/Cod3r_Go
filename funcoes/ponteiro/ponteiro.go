package main

import "fmt"

func inc1(n int) {
	n++
}

// revisando... um ponteiro é representado por um *

func inc2(n *int) {
	// revisando... o * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n) // POR VALOR...
	fmt.Println(n)

	// revisando... & é usado para obter o endereço da variável
	inc2(&n) // POR REFERÊNCIA...
	fmt.Println(n)
}
