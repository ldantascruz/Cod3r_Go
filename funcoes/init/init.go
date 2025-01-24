package main

import "fmt"

// PODEMOS TER UM init PARA CADA ARQUIVO
// ELES SEMPRE SERÃO EXECUTADOS ANTES DO main

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
