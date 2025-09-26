package classes

import ( 
	"fmt" 
)

func Condiciones() {
	fmt.Println("=== Ejemplos de Condiciones en Go ===")

	nombre := "Andrew"
	edad := 19

	if edad >= 18 {
		fmt.Println(nombre, "es mayor de edad.")
	} else {
		fmt.Println(nombre, "es menor de edad.")
	}

	if edad % 2 == 0 {
		fmt.Println("La edad es un número par.")
	} else {
		fmt.Println("La edad es un número impar.")
	}

	if edad > 0 && edad < 18 {
		fmt.Println("La edad es positiva y menor que 18.")
	}

	if numero := edad * 2; numero > 30 {
		fmt.Println("El doble de la edad es mayor que 30.")
	} else if numero == 30 {
		fmt.Println("El doble de la edad es igual a 30.")
	} else {
		fmt.Println("El doble de la edad es menor que 30.")
	}
}