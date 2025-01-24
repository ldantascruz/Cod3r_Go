package main

import "fmt"

// Estrutura de uma função em Go:
// palavra "func", nome da função, (parâmetros), (tipos de retorno), chaves {}

// função sem parâmetros e sem retornos
func f1() {
	fmt.Println("F1")
}

// função com parâmetros, mas sem retornos
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// função sem parâmetros, mas com 1 retorno
func f3() string {
	return "F3"
}

// função com parâmetros (forma simplificada de passar tipos iguais), e com 1 retorno
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s\n", p1, p2) // Sprintf apenas organiza, não exibe no console
}

// função sem parâmetros, mas com 2 retorno
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}
