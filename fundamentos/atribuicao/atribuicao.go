package main

import "fmt"

func main() {
	var z byte = 3
	fmt.Println(z)

	a := 3 // inferencia de tipo
	a += 3 // a = b + 3
	a -= 3 // a = a- 3
	a /= 3 // a = a / 3
	a *= 3 // a = a * 3
	a %= 3 // a = a % 3

	fmt.Println(a)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}
