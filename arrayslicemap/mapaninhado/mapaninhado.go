package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1200.50,
			"Gustavo Lima":   100000.00,
		},
		"J": {
			"João Silva":         32000.85,
			"Jasmine Santos":     20000.90,
			"Julio Cesar Castro": 1550.23,
		},
		"P": {
			"Pedro Junior": 800.99,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Printf("%s: %s\n", letra, funcs)
		for nome, salario := range funcs {
			fmt.Printf("Nome: %s, Salário: %.2f\n", nome, salario)
		}
	}
}
