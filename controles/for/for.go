package main

import (
	"fmt"
	"time"
)

func main() {
	// PARECIDO COM WHILE, INICIALIZA UMA VARIÁVEL FORA E VAI SEGUINDO ATÉ A CONDIÇÃO
	i := 1
	for i <= 10 { // Não tem While em Go
		fmt.Println(i)
		i++
	}

	// INICIALIZA A VARIÁVEL NO PRÓPRIO FOR, COLOCA A CONDIÇÃO E COLOCA A EXPRESSÃO (FOR NORMAL)
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("\n%d é Par", i)
		} else {
			fmt.Printf("\n%d é Impar", i)
		}
	}

	// LOOP INFINITO
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}

	// Veemos o foreach no capítulo de array
}
