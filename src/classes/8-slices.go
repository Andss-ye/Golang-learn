package classes

import "fmt"

func Slices() {
	fmt.Println("=== Slices en Go - Guía Completa ===")

	// 1. DIFERENCIAS CLAVE: ARRAYS vs SLICES
	fmt.Println("\n--- 1. Arrays vs Slices ---")
	fmt.Println("ARRAYS:")
	fmt.Println("  [5]int     - Tamaño fijo, parte del tipo")
	fmt.Println("  Se copian por valor")
	fmt.Println("\nSLICES:")
	fmt.Println("  []int      - Tamaño dinámico")
	fmt.Println("  Se pasan por referencia")
	fmt.Println("  Son como 'ventanas' a un array subyacente")

	// 2. DECLARACIÓN DE SLICES
	fmt.Println("\n--- 2. Formas de crear slices ---")

	// Slice literal
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice literal: %v (len=%d, cap=%d)\n", numeros, len(numeros), cap(numeros))

	// Slice vacío
	var vacio []int
	fmt.Printf("Slice vacío: %v (len=%d, cap=%d, nil=%t)\n", vacio, len(vacio), cap(vacio), vacio == nil)

	// Con make
	conMake := make([]int, 3) // longitud 3, capacidad 3
	fmt.Printf("Con make([]int, 3): %v (len=%d, cap=%d)\n", conMake, len(conMake), cap(conMake))

	conMakeYCap := make([]int, 3, 5) // longitud 3, capacidad 5
	fmt.Printf("Con make([]int, 3, 5): %v (len=%d, cap=%d)\n", conMakeYCap, len(conMakeYCap), cap(conMakeYCap))

	// 3. LONGITUD vs CAPACIDAD
	fmt.Println("\n--- 3. Longitud vs Capacidad ---")
	frutas := make([]string, 2, 5)
	frutas[0] = "🍎"
	frutas[1] = "🍌"
	fmt.Printf("Frutas: %v\n", frutas)
	fmt.Printf("len(frutas) = %d (elementos actuales)\n", len(frutas))
	fmt.Printf("cap(frutas) = %d (capacidad máxima sin reasignar memoria)\n", cap(frutas))

	// 4. APPEND - AGREGANDO ELEMENTOS
	fmt.Println("\n--- 4. Agregando elementos con append ---")

	// Agregar un elemento
	frutas = append(frutas, "🍊")
	fmt.Printf("Después de agregar 🍊: %v (len=%d, cap=%d)\n", frutas, len(frutas), cap(frutas))

	// Agregar múltiples elementos
	frutas = append(frutas, "🍇", "🥝")
	fmt.Printf("Después de agregar 🍇,🥝: %v (len=%d, cap=%d)\n", frutas, len(frutas), cap(frutas))

	// Agregar otro slice
	masFrutas := []string{"🍓", "🥭"}
	frutas = append(frutas, masFrutas...) // Los ... expanden el slice
	fmt.Printf("Después de agregar otro slice: %v (len=%d, cap=%d)\n", frutas, len(frutas), cap(frutas))

	// 5. SLICING - CREANDO SUB-SLICES
	fmt.Println("\n--- 5. Slicing - Creando sub-slices ---")
	numeros = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", numeros)

	// Diferentes formas de slicing
	fmt.Printf("numeros[2:5]: %v\n", numeros[2:5]) // índices 2, 3, 4
	fmt.Printf("numeros[:4]: %v\n", numeros[:4])   // desde el inicio hasta índice 4 (exclusive)
	fmt.Printf("numeros[6:]: %v\n", numeros[6:])   // desde índice 6 hasta el final
	fmt.Printf("numeros[:]: %v\n", numeros[:])     // copia completa

	// 6. SLICES COMPARTEN MEMORIA
	fmt.Println("\n--- 6. Los slices comparten memoria ---")
	original := []int{1, 2, 3, 4, 5}
	subSlice := original[1:4] // [2, 3, 4]

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Sub-slice: %v\n", subSlice)

	// Modificar el sub-slice afecta al original
	subSlice[0] = 999
	fmt.Printf("Después de modificar sub-slice[0] = 999:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Sub-slice: %v\n", subSlice)
	fmt.Println("¡El original cambió porque comparten memoria!")

	// 7. COPY - COPIANDO SLICES
	fmt.Println("\n--- 7. Copiando slices independientes ---")
	source := []int{10, 20, 30, 40, 50}
	destination := make([]int, len(source))

	copied := copy(destination, source)
	fmt.Printf("Elementos copiados: %d\n", copied)
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Destination: %v\n", destination)

	// Modificar la copia no afecta al original
	destination[0] = 999
	fmt.Printf("Después de modificar destination[0] = 999:\n")
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Destination: %v\n", destination)
	fmt.Println("¡El original NO cambió porque son independientes!")

	// 8. ELIMINANDO ELEMENTOS
	fmt.Println("\n--- 8. Eliminando elementos ---")
	colores := []string{"rojo", "verde", "azul", "amarillo", "violeta"}
	fmt.Printf("Colores originales: %v\n", colores)

	// Eliminar elemento en índice 2 (azul)
	indiceAEliminar := 2
	colores = append(colores[:indiceAEliminar], colores[indiceAEliminar+1:]...)
	fmt.Printf("Después de eliminar índice 2: %v\n", colores)

	// Eliminar el primer elemento
	colores = colores[1:]
	fmt.Printf("Después de eliminar el primero: %v\n", colores)

	// Eliminar el último elemento
	colores = colores[:len(colores)-1]
	fmt.Printf("Después de eliminar el último: %v\n", colores)

	// 9. RECORRIENDO SLICES
	fmt.Println("\n--- 9. Recorriendo slices ---")
	ciudades := []string{"Madrid", "Barcelona", "Valencia", "Sevilla"}

	// For tradicional
	fmt.Println("Con for tradicional:")
	for i := 0; i < len(ciudades); i++ {
		fmt.Printf("  [%d]: %s\n", i, ciudades[i])
	}

	// Con range (índice y valor)
	fmt.Println("Con range (índice y valor):")
	for i, ciudad := range ciudades {
		fmt.Printf("  [%d]: %s\n", i, ciudad)
	}

	// Con range (solo valor)
	fmt.Println("Con range (solo valor):")
	for _, ciudad := range ciudades {
		fmt.Printf("  • %s\n", ciudad)
	}

	// Con range (solo índice)
	fmt.Println("Con range (solo índice):")
	for i := range ciudades {
		fmt.Printf("  Índice: %d\n", i)
	}

	// 10. SLICES MULTIDIMENSIONALES
	fmt.Println("\n--- 10. Slices multidimensionales ---")

	// Matriz dinámica
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matriz:")
	for i, fila := range matrix {
		fmt.Printf("  Fila %d: %v\n", i, fila)
	}

	// Agregar una nueva fila
	matrix = append(matrix, []int{10, 11, 12})
	fmt.Printf("Después de agregar fila: %v\n", matrix)

	// 11. CASOS PRÁCTICOS
	fmt.Println("\n--- 11. Casos prácticos ---")

	// Lista de tareas
	tareas := []string{}
	tareas = append(tareas, "Estudiar Go")
	tareas = append(tareas, "Hacer ejercicio")
	tareas = append(tareas, "Leer un libro")

	fmt.Println("Lista de tareas:")
	for i, tarea := range tareas {
		fmt.Printf("  %d. %s\n", i+1, tarea)
	}

	// Filtrar números pares
	todosNumeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var pares []int

	for _, num := range todosNumeros {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}

	fmt.Printf("Números originales: %v\n", todosNumeros)
	fmt.Printf("Solo pares: %v\n", pares)

	// Encontrar máximo
	calificaciones := []float64{8.5, 9.2, 7.8, 9.6, 8.8, 9.9, 7.5}
	max := calificaciones[0]

	for _, cal := range calificaciones {
		if cal > max {
			max = cal
		}
	}

	fmt.Printf("Calificaciones: %v\n", calificaciones)
	fmt.Printf("Calificación máxima: %.1f\n", max)

	// 12. FUNCIONES ÚTILES CON SLICES
	fmt.Println("\n--- 12. Funciones útiles ---")
	demostrateSliceFunctions()

	fmt.Println("\n=== Consejos importantes ===")
	fmt.Println("• Los slices son dinámicos y se pasan por referencia")
	fmt.Println("• Usa append() para agregar elementos")
	fmt.Println("• Los sub-slices comparten memoria con el original")
	fmt.Println("• Usa copy() para crear copias independientes")
	fmt.Println("• len() = elementos actuales, cap() = capacidad máxima")
	fmt.Println("• Go reasigna automáticamente memoria cuando se supera la capacidad")
}

// Función auxiliar para demostrar operaciones con slices
func demostrateSliceFunctions() {
	// Función para verificar si un slice contiene un elemento
	contains := func(slice []string, item string) bool {
		for _, s := range slice {
			if s == item {
				return true
			}
		}
		return false
	}

	// Función para remover duplicados
	removeDuplicates := func(slice []int) []int {
		keys := make(map[int]bool)
		var result []int

		for _, item := range slice {
			if !keys[item] {
				keys[item] = true
				result = append(result, item)
			}
		}
		return result
	}

	// Función para reversar slice
	reverse := func(slice []string) []string {
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
		return slice
	}

	// Probar las funciones
	animales := []string{"perro", "gato", "pájaro", "pez"}
	fmt.Printf("¿Contiene 'gato'?: %t\n", contains(animales, "gato"))
	fmt.Printf("¿Contiene 'león'?: %t\n", contains(animales, "león"))

	conDuplicados := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	fmt.Printf("Con duplicados: %v\n", conDuplicados)
	fmt.Printf("Sin duplicados: %v\n", removeDuplicates(conDuplicados))

	palabras := []string{"hola", "mundo", "go", "es", "genial"}
	fmt.Printf("Original: %v\n", palabras)
	fmt.Printf("Reverso: %v\n", reverse(append([]string(nil), palabras...))) // Copiamos antes de reversar
}
