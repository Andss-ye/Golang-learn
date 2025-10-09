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
	edades := []int{19, 20, 18, 19, 21}
	calificaciones := []float64{3.5, 3.1, 4.3, 5.0, 2.3}

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

	// TODO 5: Sistema de mejora
	// Simula que algunos estudiantes mejoran:
	// - Agrega 1.0 punto a estudiantes con calificación < 6.0
	// - Recalcula y muestra las nuevas estadísticas
	// - Compara antes vs después

	fmt.Println("\n¡Este proyecto usa TODO lo que aprendiste!")
}
