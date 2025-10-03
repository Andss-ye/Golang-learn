package exercises

import "fmt"

// Ejercicio 3: Bucles For
// Practica diferentes tipos de loops

func BuclesFor() {
	fmt.Println("=== Ejercicio 3: Bucles For ===")

	// TODO 1: Tabla de multiplicar
	// Crea la tabla del 7 del 1 al 10
	// Usa un for loop y muestra: "7 x 1 = 7", "7 x 2 = 14", etc.

	// TODO 2: Sumar números del 1 al 100
	// Usa un for loop para sumar todos los números del 1 al 100
	// Muestra el resultado final

	// TODO 3: Contar números pares e impares
	// Del 1 al 20, cuenta cuántos números son pares y cuántos impares
	// Usa el operador % (módulo) para verificar si es par: numero%2 == 0

	// TODO 4: Pirámide de asteriscos
	// Crea una pirámide de 5 niveles usando for loops anidados:
	// *
	// **
	// ***
	// ****
	// *****

	// TODO 5: Factorial
	// Calcula el factorial de 5 (5! = 5 * 4 * 3 * 2 * 1 = 120)

	// TODO 6: Secuencia Fibonacci
	// Genera los primeros 10 números de Fibonacci: 0, 1, 1, 2, 3, 5, 8, 13...
	// Cada número es la suma de los dos anteriores

	fmt.Println("\n¡Practica con diferentes valores y patrones!")
} // DESAFÍO EXTRA: Números primos
// Encuentra todos los números primos entre 2 y 50
func EncontrarPrimos() {
	fmt.Println("\n=== DESAFÍO: Números Primos ===")
	// TODO: Implementa la lógica para encontrar números primos
	// Un número primo solo es divisible por 1 y por sí mismo

	fmt.Println("Números primos entre 2 y 50:")
	// for num := 2; num <= 50; num++ {
	//     esPrimo := true
	//     for i := 2; i < num; i++ {
	//         if num%i == 0 {
	//             esPrimo = false
	//             break
	//         }
	//     }
	//     if esPrimo {
	//         fmt.Printf("%d ", num)
	//     }
	// }
	// fmt.Println()
}
