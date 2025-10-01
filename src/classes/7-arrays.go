package classes

import (
	"fmt"
)

func Arrays() {
	fmt.Println("=== Arrays en Go - Guía Completa ===")

	// 1. DECLARACIÓN BÁSICA DE ARRAYS
	fmt.Println("\n--- 1. Declaración básica ---")
	var array [5]float32 // Array de 5 elementos, inicializado con valores cero
	fmt.Printf("Array vacío: %v\n", array)
	fmt.Printf("Valor por defecto de float32: %.1f\n", array[0])

	// Asignar valores individuales
	array[0] = 10.5
	array[4] = 100.0
	array[3] = 1232.42
	fmt.Printf("Array con algunos valores: %v\n", array)

	// 2. INICIALIZACIÓN CON VALORES
	fmt.Println("\n--- 2. Inicialización con valores ---")
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array inicializado: %v\n", numbers)

	// Inicialización parcial
	partial := [5]int{1, 2} // Los demás son 0
	fmt.Printf("Inicialización parcial: %v\n", partial)

	// Inicialización con índices específicos
	indexed := [5]int{0: 10, 2: 20, 4: 40}
	fmt.Printf("Inicialización con índices: %v\n", indexed)

	// 3. ARRAYS CON TAMAÑO AUTOMÁTICO
	fmt.Println("\n--- 3. Arrays con tamaño automático ---")
	autoSize := [...]string{"Go", "Python", "JavaScript", "Rust"}
	fmt.Printf("Array auto-size: %v (longitud: %d)\n", autoSize, len(autoSize))

	// Array vacío con tamaño automático (¡Cuidado!)
	empty := [...]int{}
	fmt.Printf("Array vacío: %v (longitud: %d)\n", empty, len(empty))

	// 4. OPERACIONES CON ARRAYS
	fmt.Println("\n--- 4. Operaciones con arrays ---")
	scores := [5]int{85, 92, 78, 96, 88}

	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Primer score: %d\n", scores[0])
	fmt.Printf("Último score: %d\n", scores[len(scores)-1])
	fmt.Printf("Longitud del array: %d\n", len(scores))

	// Recorrer con for tradicional
	fmt.Println("\nRecorriendo con for tradicional:")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("  scores[%d] = %d\n", i, scores[i])
	}

	// Recorrer con range
	fmt.Println("\nRecorriendo con range:")
	for index, value := range scores {
		fmt.Printf("  índice %d: valor %d\n", index, value)
	}

	// Solo valores (ignorar índice)
	fmt.Println("\nSolo valores:")
	for _, value := range scores {
		fmt.Printf("  %d ", value)
	}
	fmt.Println()

	// 5. ARRAYS MULTIDIMENSIONALES
	fmt.Println("\n--- 5. Arrays multidimensionales ---")
	var matrix [3][3]int

	// Llenar la matriz
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}

	fmt.Println("Matriz 3x3:")
	for i := 0; i < 3; i++ {
		fmt.Printf("  %v\n", matrix[i])
	}

	// Matriz inicializada directamente
	board := [3][3]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}

	fmt.Println("\nTablero de tic-tac-toe:")
	for _, row := range board {
		fmt.Printf("  %v\n", row)
	}

	// 6. COMPARACIÓN DE ARRAYS
	fmt.Println("\n--- 6. Comparación de arrays ---")
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v\n", arr3)
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3)

	// 7. ARRAYS vs SLICES (DIFERENCIAS IMPORTANTES)
	fmt.Println("\n--- 7. Arrays vs Slices ---")
	fmt.Println("ARRAYS:")
	fmt.Println("  - Tamaño fijo definido en tiempo de compilación")
	fmt.Println("  - Se pasan por valor (se copian)")
	fmt.Println("  - Tipo incluye el tamaño: [5]int != [3]int")

	// Demostrar que se copian
	original := [3]int{1, 2, 3}
	copy := original
	copy[0] = 999

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copia modificada: %v\n", copy)
	fmt.Println("¡El original no cambió porque se copió!")

	// 8. CASOS DE USO PRÁCTICOS
	fmt.Println("\n--- 8. Casos de uso prácticos ---")

	// Días de la semana
	weekdays := [7]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
	fmt.Println("Días de la semana:")
	for i, day := range weekdays {
		fmt.Printf("  %d: %s\n", i+1, day)
	}

	// Temperaturas de la semana
	temperatures := [7]float64{22.5, 25.0, 23.8, 21.2, 24.5, 26.8, 25.5}
	var sum float64
	for _, temp := range temperatures {
		sum += temp
	}
	average := sum / float64(len(temperatures))
	fmt.Printf("\nTemperatura promedio de la semana: %.1f°C\n", average)

	// Encontrar el máximo
	max := temperatures[0]
	maxDay := 0
	for i, temp := range temperatures {
		if temp > max {
			max = temp
			maxDay = i
		}
	}
	fmt.Printf("Día más caluroso: %s con %.1f°C\n", weekdays[maxDay], max)

	fmt.Println("\n=== Consejos importantes ===")
	fmt.Println("• Los arrays tienen tamaño fijo y se definen en tiempo de compilación")
	fmt.Println("• Se pasan por valor (se copian), no por referencia")
	fmt.Println("• Para tamaños dinámicos, usa slices ([]int en lugar de [5]int)")
	fmt.Println("• Los arrays son útiles cuando conoces el tamaño exacto de antemano")
}
