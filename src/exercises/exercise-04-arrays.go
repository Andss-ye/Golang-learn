package exercises

import "fmt"

// Ejercicio 4: Arrays
// Practica con arrays de tamaño fijo

func Arrays() {
	fmt.Println("=== Ejercicio 4: Arrays ===")

	// TODO 1: Temperaturas de la semana
	// Crea un array [7]float64 con temperaturas de una semana
	// Calcula y muestra el promedio usando un for loop
	var sum float64
	temperaturas := [7]float64{40, 34.3, -4, 41, 20, -3.3, -10.3}

	for _, v := range temperaturas {
		fmt.Printf("v: %v\n", v)
		sum += v
	}

	average := sum / 7
	fmt.Printf("El promedio de las temperaturas de la ciudad es: %v\n", average)

	// TODO 2: Encontrar máximo y mínimo
	// Del array de temperaturas, encuentra:
	// - La temperatura más alta y en qué día fue
	// - La temperatura más baja y en qué día fue
	// Crea también un array con los nombres de los días
	dias := [7]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}
	menor := temperaturas[0]
	mayor := temperaturas[0]
	dia_menor := 0
	dia_mayor := 0

	for i, v := range temperaturas {
		if v < menor {
			menor = v
			dia_menor = i
		}
		if v > mayor {
			mayor = v
			dia_mayor = i
		}
	}

	fmt.Printf("menor temp: %v --- dia menor: %v\n", menor, dias[dia_menor])
	fmt.Printf("mayor temp: %v --- dia mayor: %v\n", mayor, dias[dia_mayor])

	// TODO 3: Matriz de multiplicación
	// Crea una matriz [3][3]int
	// Lléna con las tablas del 1, 2 y 3 usando for loops anidados
	// Muestra cada fila
	matriz := [3][3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = (i + 1) * (j + 1)
		}
	}

	fmt.Println("Matriz de multiplicación:")
	for i := 0; i < 3; i++ {
		fmt.Println(matriz[i])
	}

	// TODO 4: Sistema de inventario
	// Crea 3 arrays paralelos:
	// - productos [5]string (nombres de productos)
	// - cantidades [5]int (stock disponible)
	// - precios [5]float64 (precio unitario)
	// Muestra el inventario completo con valor total por producto
	productos := [5]string{"Manzanas", "Bananas", "Naranjas", "Peras", "Uvas"}
	cantidades := [5]int{10, 5, 8, 12, 7}
	precios := [5]float64{1.5, 0.8, 1.2, 1.0, 2.0}

	fmt.Printf("productos ---- cantidades ---- precios ---- total\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ---- %v ---- %v ---- %v\n", productos[i], cantidades[i], precios[i], precios[i]*float64(cantidades[i]))
	}

	// TODO 5: Buscar en inventario
	// Busca un producto específico en el array de productos
	// Si lo encuentras, muestra su información completa
	// Si no existe, muestra mensaje de "no encontrado"
	buscar := "Naranjas" // Cambia este valor para buscar otro producto
	encontrado := false

	for i := 0; i < len(productos); i++ {
		if productos[i] == buscar {
			fmt.Printf("Producto encontrado: %v\nCantidad: %v\nPrecio: %v\nTotal: %v\n",
				productos[i], cantidades[i], precios[i], precios[i]*float64(cantidades[i]))
			encontrado = true
			break
		}
	}
	if !encontrado {
		fmt.Printf("Producto '%v' no encontrado en el inventario.\n", buscar)
	}

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
