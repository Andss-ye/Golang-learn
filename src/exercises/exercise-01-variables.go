package exercises

import "fmt"

// Ejercicio 1: Variables y Tipos Básicos
// Practica declaración y uso de variables

func VariablesBasicas() {
	fmt.Println("=== Ejercicio 1: Variables y Tipos Básicos ===")

	// TODO 1: Declara las siguientes variables usando var:
	// - nombre (string): tu nombre
	// - edad (int): tu edad
	// - altura (float64): tu altura en metros
	// - estudiante (bool): si eres estudiante o no

	var nombre string = "Andrew"
	var edad int = 19
	var altura float32 = 1.82
	var estudiante bool = true

	// TODO 2: Muestra toda la información usando Printf con formato

	fmt.Printf(" nombre: %s\n edad: %v\n altura: %v\n estudiante: %v\n", nombre, edad, altura, estudiante)

	// TODO 3: Declara 3 variables más usando inferencia de tipo (:=):
	// - país donde vives
	// - código postal (número)
	// - salario mensual (decimal)

	pais := "Colombia"
	postal_code := "111111"
	salary := 3000000

	fmt.Printf(" país: %s\n código postal: %s\n salario mensual: %d\n", pais, postal_code, salary)
	
	// TODO 4: Calcula y muestra tu año aproximado de nacimiento
	añoActual := 2025
	añoNacimiento := añoActual - edad
	fmt.Printf("Naciste aproximadamente en: %d\n", añoNacimiento)

	fmt.Println("\n¡Completa los TODOs y ejecuta para ver los resultados!")
}

// SOLUCIÓN (descomenta cuando hayas intentado el ejercicio):
/*
func VariablesBasicasSolucion() {
	fmt.Println("=== SOLUCIÓN: Variables y Tipos Básicos ===")

	var nombre string = "Juan"
	var edad int = 25
	var altura float64 = 1.75
	var estudiante bool = true

	fmt.Printf("Nombre: %s, Edad: %d, Altura: %.2fm, Estudiante: %t\n", nombre, edad, altura, estudiante)

	pais := "España"
	codigoPostal := 28001
	salario := 35000.50

	fmt.Printf("País: %s, CP: %d, Salario: %.2f€\n", pais, codigoPostal, salario)

	añoActual := 2025
	añoNacimiento := añoActual - edad
	fmt.Printf("Naciste aproximadamente en: %d\n", añoNacimiento)
}
*/
