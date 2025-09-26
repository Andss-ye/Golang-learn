package classes

import (
	"fmt"
)

func For() {
	fmt.Println("=== Ejemplos de Bucles For en Go ===")

	i := 0
	for i > -5 {
		fmt.Println(i)
		i--
	}

	fmt.Println("-----")

	for rango := 0; rango > -5; rango-- {
		fmt.Println(rango)
	}
}