package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // Não passo nos colchetes o valor, mas o compilador conta quantos elementos temos entre as chaves

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	// ignorando o índice
	for _, num := range numeros {
		fmt.Println(num)
	}
}
