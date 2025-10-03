package exercises

import "fmt"

// Ejercicio 4: Arrays
// Practica con arrays de tamaño fijo

func Arrays() {
	fmt.Println("=== Ejercicio 4: Arrays ===")

	// TODO 1: Temperaturas de la semana
	// Crea un array [7]float64 con temperaturas de una semana
	// Calcula y muestra el promedio usando un for loop

	// TODO 2: Encontrar máximo y mínimo
	// Del array de temperaturas, encuentra:
	// - La temperatura más alta y en qué día fue
	// - La temperatura más baja y en qué día fue
	// Crea también un array con los nombres de los días

	// TODO 3: Matriz de multiplicación
	// Crea una matriz [3][3]int
	// Lléna con las tablas del 1, 2 y 3 usando for loops anidados
	// Muestra cada fila

	// TODO 4: Sistema de inventario
	// Crea 3 arrays paralelos:
	// - productos [5]string (nombres de productos)
	// - cantidades [5]int (stock disponible)
	// - precios [5]float64 (precio unitario)
	// Muestra el inventario completo con valor total por producto

	// TODO 5: Buscar en inventario
	// Busca un producto específico en el array de productos
	// Si lo encuentras, muestra su información completa
	// Si no existe, muestra mensaje de "no encontrado"

	fmt.Println("\n¡Los arrays tienen tamaño fijo pero son muy útiles!")
} // DESAFÍO EXTRA: Sudoku validator (3x3)
// Valida si una cuadrícula 3x3 contiene números del 1-9 sin repetir
func ValidarCuadricula() {
	cuadricula := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// TODO: Implementa la validación
	// Debe verificar que cada número del 1-9 aparezca exactamente una vez
	fmt.Printf("Cuadrícula: %v\n", cuadricula)
	fmt.Println("¿Es válida? (implementa la lógica)")
}
