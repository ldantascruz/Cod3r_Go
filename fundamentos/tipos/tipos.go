package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numéricos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("\nLiteral inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255 // apelido para uint8
	fmt.Println("\nO byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("\nO valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("\nO rune i2 é", reflect.TypeOf(i2))
	fmt.Println("O valor do i2 é", i2)

	// números reais... float32 float64
	var x float32 = 49.99
	fmt.Println("\nO tipo de x é", reflect.TypeOf(x))
	fmt.Println("O valor de x é", x)
	fmt.Println("O valor liteal de 49.99 é", reflect.TypeOf(49.99)) // float64

	var x2 = float32(39.99)
	fmt.Println("\nO tipo de x é", reflect.TypeOf(x2))
	fmt.Println("O valor de x é", x2)
	fmt.Println("O valor liteal de 49.99 é", reflect.TypeOf(39.99)) // float64

	// boolean
	bo := true
	fmt.Println("\nO tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Lucas"
	fmt.Println("\nO tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println("O valor de s1 é", s1)
	fmt.Println("O tamanho da string s1 é:", len(s1))

	// string com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Lucas`
	fmt.Println("\nO tamanho da string s2 é:", len(s2))

	// char???
	char := 'a'
	fmt.Println("\nO tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
