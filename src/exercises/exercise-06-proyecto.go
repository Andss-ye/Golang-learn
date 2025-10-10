package exercises

import (
	"fmt"
	"math"
)

// Ejercicio 6: Proyecto Integrador
// Integra todo: variables, condicionales, bucles, arrays y slices

func ProyectoIntegrador() {
	fmt.Println("=== Proyecto Integrador: Sistema de Estudiantes ===")

	// TODO 1: Crear datos básicos
	// Crea slices para almacenar información de 5 estudiantes:
	// - nombres []string
	// - edades []int
	// - calificaciones []float64
	nombre := []string{"Andrew", "Leon", "Daniel", "Nicolas", "Angel"}
	edades := []int{19, 10, 18, 19, 11}
	calificaciones := []float64{3.5, 3.1, 4.3, 5.0, 8.3}

	// TODO 2: Mostrar información completa
	// Recorre los datos con un for loop
	// Muestra: "Nombre (edad años) - Calificación: X.X"
	for i := 0; i < len(nombre); i++ {
		fmt.Printf("%v (edad %v) - Calificacion: %v\n", nombre[i], edades[i], calificaciones[i])
	}

	// TODO 3: Calcular estadísticas
	// Calcula y muestra:
	// - Edad promedio de todos los estudiantes
	// - Calificación promedio general
	// - Mejor estudiante (mayor calificación)
	// - Peor estudiante (menor calificación)
	var edad_suma int
	var calificacion_suma float64
	var menor float64 = math.MaxFloat64
	var mayor float64 = 0.0
	for i := 0; i < 5; i++ {
		edad_suma += edades[i]
		calificacion_suma += calificaciones[i]

		if calificaciones[i] < calificaciones[0] {
			menor = calificaciones[i]
		}
		if calificaciones[i] > calificaciones[0] {
			mayor = calificaciones[i]
		}
	}
	
	promedio_edad := edad_suma / 5
	promedio_calificaciones := calificacion_suma / 5
	fmt.Printf("Promedio de edades: %v\n", promedio_edad)
	fmt.Printf("Promedio de calificaciones: %v\n", promedio_calificaciones)
	fmt.Printf("Nota menor: %v ---- Nota mayor: %v\n", menor, mayor)

	// TODO 4: Filtrar estudiantes
	// Crea nuevos slices con:
	// - Estudiantes menores de 20 años
	// - Estudiantes con calificación >= 8.0 (excelentes)
	// - Estudiantes con calificación < 6.0 (necesitan ayuda)
	menores_20 := []string{}
	excelentes := []string{}
	regulares := []string{}
	for i := 0; i < len(nombre); i++ {
		if edades[i] < 20 {
			menores_20 = append(menores_20, nombre[i])
		}
	}

	for i := 0; i < len(menores_20); i++ {
		if calificaciones[i] >= 8.0 {
			excelentes = append(excelentes, menores_20[i])
		}
		if calificaciones[i] < 6.0 {
			regulares = append(regulares, menores_20[i])
		}
	}

	if len(excelentes) == 0 {
		fmt.Println("No hubieron estudiantes excelentes")
	} else {
		fmt.Printf("Los estudiantes excelentes fueron: \n")
		for i := 0; i < len(excelentes); i++ {
			fmt.Printf("%v\n", excelentes[i])
		}
	}

	if len(regulares) == 0 {
		fmt.Println("No hubieron estudiantes regulares :D")
	} else {
		fmt.Printf("Los estudiantes regulares fueron: \n")
		for i := 0; i < len(regulares); i++ {
			fmt.Printf("%v\n", regulares[i])
		}
	}


	// TODO 5: Sistema de mejora
	// Simula que algunos estudiantes mejoran:
	// - Agrega 1.0 punto a estudiantes con calificación < 6.0
	// - Recalcula y muestra las nuevas estadísticas
	// - Compara antes vs después
	var suma float64
	fmt.Println("---------------------------------------------")
	for i := 0; i < len(nombre); i++ {
		fmt.Printf("calificacion actual: %v\n", calificaciones[i])
		if calificaciones[i] < 6.0 {
			calificaciones[i] = calificaciones[i] + 1.0
		}
		fmt.Printf("nueva calificacion: %v\n", calificaciones[i])
		suma += calificaciones[i]
	}
	nuevo_promedio := suma / float64(len(nombre))
	fmt.Printf("Nuevo promedio de calificaciones: %v\n", nuevo_promedio)
	fmt.Printf("Promedio anterior de calificaciones: %v\n", promedio_calificaciones)

	fmt.Println("\n¡Este proyecto usa TODO lo que aprendiste!")
}
