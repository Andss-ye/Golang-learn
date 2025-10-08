package exercises

import "fmt"

// Ejercicio 5: Slices Avanzados
// Practica manipulación dinámica de datos

func SlicesAvanzados() {
	fmt.Println("=== Ejercicio 5: Slices Avanzados ===")

	// TODO 1: Sistema de calificaciones dinámico
	// Crea un slice vacío []float64
	// Agrega 5 calificaciones usando append()
	// Calcula y muestra el promedio
	calificaciones := []float64{}
	i := 0
	for i < 5 {
		calificaciones = append(calificaciones, float64(i))
		i++
	}

	// Calcular el promedio
	var suma float64
	for _, cal := range calificaciones {
		suma += cal
	}
	promedio := suma / float64(len(calificaciones))
	fmt.Printf("Calificaciones: %v\n", calificaciones)
	fmt.Printf("Promedio: %.2f\n", promedio)
	

	// TODO 2: Filtrar números pares
	// Crea un slice con números del 1 al 20
	// Crea un nuevo slice solo con los números pares
	// Usa append() dentro de un for loop
	pares := []int{}
	i = 0
	for i <= 20 {
		if i % 2 == 0 {
			pares = append(pares, i)
		}
		i++
	}
	fmt.Printf("%v\n", pares)

	// TODO 3: Lista de tareas interactiva
	tareas := []string{"Comprar leche", "Estudiar Go", "Hacer ejercicio"}
	tareas = append(tareas, "Leer un libro", "Llamar a un amigo")
	// Eliminar la segunda tarea ("Estudiar Go")
	tareas = append(tareas[:1], tareas[2:]...)
	fmt.Printf("Lista de tareas final: %v\n", tareas)

	// TODO 4: Encontrar palabras largas
	palabras := []string{"gato", "elefante", "sol", "programacion", "auto", "computadora"}
	largas := []string{}
	for _, palabra := range palabras {
		if len(palabra) > 5 {
			largas = append(largas, palabra)
		}
	}
	fmt.Printf("Palabras largas: %v\n", largas)

	// TODO 5: Matriz dinámica
	matriz := make([][]int, 3)
	num := 1
	for i := range matriz {
		matriz[i] = make([]int, 3)
		for j := range matriz[i] {
			matriz[i][j] = num
			num++
		}
	}
	fmt.Println("Matriz dinámica:")
	for _, fila := range matriz {
		fmt.Println(fila)
	}

	fmt.Println("\n¡Los slices son dinámicos y muy poderosos!")
}
