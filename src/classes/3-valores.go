package classes

import (
	"fmt"
)

func Valores() {
	fmt.Println("=== Ejemplos de Valores en Go ===")

	var numero int = 10
	numero2 := 20
	fmt.Println(numero, numero2)

	var numeroEntero = 10
	var numeroDoble = 20.5
	
	resultado := float64(numeroEntero) + numeroDoble
	fmt.Println("Resultado de la suma:", resultado)

	var nombre string = "Andrew"
	apellido := "Avila"
	nombreCompleto := nombre + " " + apellido

	edad := 19

	fmt.Println("Nombre completo:", nombreCompleto)
	fmt.Println("Edad:", edad)
}