package main

import "fmt"

// defer é bom para se quer lembrar de fechar um recurso ao final da execução. Muito usado para fechar conexão do Banco de Dados
func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
