package classes

import "fmt"

func Variables() {
	fmt.Println("=== Ejemplos de Variables en Go ===")

	// Variables con declaración explícita
	var cadena = "Hola, Mundo!"
	fmt.Println(cadena)

	var entero1, entero2 int = 34, 56
	fmt.Println(entero1, entero2)

	var booleano = true
	fmt.Println(booleano)

	var enteroSinInicializar int
	if entero2 > 50 {
		enteroSinInicializar = 100
	}
	fmt.Println(enteroSinInicializar)

	fruta := []string{"manzana", "banana", "cereza"}
	for _, f := range fruta {
		fmt.Println(f)
	}
}
