package exercises

import "fmt"

// Ejercicio 2: Condicionales
// Practica estructuras de control if/else y switch

func Condicionales() {
	fmt.Println("=== Ejercicio 2: Condicionales ===")

	// TODO 1: Clasificador de edad
	// Declara una variable edad con un valor
	// Usa if/else if/else para clasificar:
	// 0-12: Niño, 13-17: Adolescente, 18-64: Adulto, 65+: Adulto Mayor
	edad := 1

	if edad > 65 {
		println("Adulto Mayor")
	} else if edad >= 18 && edad <= 64 {
		println("Adulto")
	} else if edad >= 13 && edad <= 17 {
		println("Adolescente")
	} else {
		println("Niño")
	}

	// TODO 2: Calculadora de notas
	// Declara una variable nota (float64)
	// Convierte a letra usando if o switch:
	// 90-100: A, 80-89: B, 70-79: C, 60-69: D, <60: F
	nota := 95
	var letra string
	if nota >= 90 && nota <= 100 {
		letra = "A"
	} else if nota >= 80 && nota < 90 {
		letra = "B"
	} else if nota >= 70 && nota < 80 {
		letra = "C"
	} else if nota >= 60 && nota < 70 {
		letra = "D"
	} else {
		letra = "F"
	}
	fmt.Println("Nota en letra:", letra)

	// TODO 3: Verificador de contraseña
	// Declara una variable password
	// Verifica si es válida:
	// - Debe tener al menos 8 caracteres (len(password) >= 8)
	// - No debe ser "password" ni "123456"

	password := "123456"
	if len(password) < 8 {
		println("No es valida su contrasena")
	} else if password == "password" || password == "123456" {
		println("No es valida su contrasena")
	} else {
		println("Es valida su contrasena")
	}

	// TODO 4: Calculadora simple
	// Declara dos números y una operación (string)
	// Usa switch para hacer la operación: +, -, *, /
	// Muestra el resultado

	n1 := 2
	n2 := 5
	op := "/"

	switch op {
	case "+":
		fmt.Println("Resultado:", n1+n2)
	case "-":
		fmt.Println("Resultado:", n1-n2)
	case "*":
		fmt.Println("Resultado:", n1*n2)
	case "/":
		fmt.Println("Resultado:", n1/n2)
	default:
		fmt.Println("Operacion no valida")
	}

	fmt.Println("\n¡Experimenta cambiando los valores!")

	EsBisiesto(2024)
} // DESAFÍO EXTRA: Función para determinar si un año es bisiesto
// Un año es bisiesto si:
// - Es divisible por 4 Y (no es divisible por 100 O es divisible por 400)
func EsBisiesto(año int) bool {
	if año%4 == 0 && !(año%100 == 0 && año%400 != 0) {
		fmt.Printf("%d es bisiesto\n", año)
	}
	return false
}
