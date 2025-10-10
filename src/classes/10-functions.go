package classes

import "fmt"

func Functions() {
	fmt.Println("=== Funciones en Go - Guía Completa ===")

	// 1. ¿QUÉ SON LAS FUNCIONES?
	fmt.Println("\n--- 1. ¿Qué son las Funciones? ---")
	fmt.Println("Las funciones son bloques de código reutilizables que realizan una tarea específica.")
	fmt.Println("• Pueden recibir parámetros y devolver valores.")
	fmt.Println("• Ayudan a organizar y modularizar el código.")

	// 2. DECLARACIÓN DE FUNCIONES
	fmt.Println("\n--- 2. Declaración de Funciones ---")

	// Función básica
	saludo := func() {
		fmt.Println("Hola, mundo!")
	}
	saludo()

	// Función con parámetros
	sumar := func(a, b int) int {
		return a + b
	}
	resultado := sumar(3, 5)
	fmt.Printf("La suma de 3 y 5 es: %d\n", resultado)

	// Función con múltiples valores de retorno
	dividir := func(dividendo, divisor float64) (float64, error) {
		if divisor == 0 {
			return 0, fmt.Errorf("no se puede dividir entre cero")
		}
		return dividendo / divisor, nil
	}
	cociente, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("El cociente es: %.2f\n", cociente)
	}

	// 3. FUNCIONES ANÓNIMAS
	fmt.Println("\n--- 3. Funciones Anónimas ---")
	anonima := func(nombre string) {
		fmt.Printf("Hola, %s!\n", nombre)
	}
	anonima("Andrés")

	// 4. FUNCIONES COMO VALORES
	fmt.Println("\n--- 4. Funciones como Valores ---")
	operacion := func(a, b int) int {
		return a * b
	}
	fmt.Printf("El producto de 4 y 5 es: %d\n", operacion(4, 5))

	// 5. FUNCIONES COMO PARÁMETROS
	fmt.Println("\n--- 5. Funciones como Parámetros ---")
	aplicarOperacion := func(a, b int, operacion func(int, int) int) int {
		return operacion(a, b)
	}
	suma := func(a, b int) int {
		return a + b
	}
	fmt.Printf("La suma de 7 y 8 es: %d\n", aplicarOperacion(7, 8, suma))

	// 6. FUNCIONES CLOSURE
	fmt.Println("\n--- 6. Funciones Closure ---")
	contador := func() func() int {
		contador := 0
		return func() int {
			contador++
			return contador
		}
	}()
	fmt.Println("Contador:", contador())
	fmt.Println("Contador:", contador())

	// 7. FUNCIONES VARIÁDICAS
	fmt.Println("\n--- 7. Funciones Variádicas ---")
	imprimirNombres := func(nombres ...string) {
		for _, nombre := range nombres {
			fmt.Println(nombre)
		}
	}
	imprimirNombres("Ana", "Luis", "Carlos")

	// 8. DIFERENCIA ENTRE FUNCIONES Y MÉTODOS
	fmt.Println("\n--- 8. Diferencia entre Funciones y Métodos ---")
	type Persona struct {
		nombre string
	}
	saludarPersona := func(p Persona) {
		fmt.Printf("Hola, soy %s\n", p.nombre)
	}
	persona := Persona{nombre: "María"}
	saludarPersona(persona)

	// 9. FUNCIONES RECURSIVAS
	fmt.Println("\n--- 9. Funciones Recursivas ---")
	var factorial func(int) int
	factorial = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}
	fmt.Printf("El factorial de 5 es: %d\n", factorial(5))

	// 10. FUNCIONES DEFER
	fmt.Println("\n--- 10. Funciones Defer ---")
	ejemploDefer := func() {
		defer fmt.Println("Esto se ejecuta al final")
		fmt.Println("Esto se ejecuta primero")
	}
	ejemploDefer()

	fmt.Println("\n=== Fin de la Clase de Funciones ===")
}
