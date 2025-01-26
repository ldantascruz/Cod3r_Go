package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para um canal (escrita)
	<-ch    // recebendo dados do canal (leitura)
	//fmt.Println(<-ch)

	ch <- 2
	fmt.Println(<-ch)
}
