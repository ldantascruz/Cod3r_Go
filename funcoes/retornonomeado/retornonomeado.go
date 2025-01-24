package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo... como já foi atribuido o valor que quero retornar podemos colocar apenas o "return" ao invés de "return segundo, primeiro"
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(4, 5)
	fmt.Println(segundo, primeiro)
}
