package exercises

import "fmt"

func ExerciseFunctions() {
	fmt.Println("=== Ejercicios de Funciones ===")

	// TODO: Crear una función que reciba dos enteros y devuelva su suma.
	suma := func(a, b int) int {
		return a + b
	}
	fmt.Printf("la suma de 5 + 6; es: %v\n", suma(5, 6))

	// TODO: Crear una función que reciba un slice de enteros y devuelva el promedio.
	promedio := func(notas[]int) float64 {
		var suma int
		for nota := range notas {
			suma += nota
		}
		return float64(suma) / float64((len(notas)))
	}
	notas := []int{5, 4, 2, 5, 5, 0, 5}
	resultadoPromedio := promedio(notas)
	fmt.Printf("El promedio de las notas es: %.2f\n", resultadoPromedio)

	// TODO: Crear una función anónima que imprima un mensaje y ejecutarla.
	anonima := func() {
		fmt.Println("Mensaje anonimo...")
	}
	anonima()

	// TODO: Crear una función que reciba otra función como parámetro y la ejecute.
	aplicar_operacion := func(a, b int, operacion func(int, int) int) int {
		return operacion(a, b)
	}
	fmt.Printf("La suma de 5 y 12334 es: %v\n", aplicar_operacion(5, 12334, suma))

	// TODO: Crear una función closure que mantenga un contador y lo incremente cada vez que se llame.
	contador := func() func() int {
		contador := 0
		return func() int {
			contador++
			return contador
		}
	}()
	fmt.Println("Contador:", contador())
	fmt.Println("Contador:", contador())
	// TODO: Crear una función variádica que reciba un número arbitrario de strings y los imprima.
	imprimirNombres := func(nombres ...string) {
		for _, nombre := range nombres {
			fmt.Printf("	%v\n", nombre)
		}
	}
	imprimirNombres("Andrew", "Luis", "Michael")

	fmt.Println("=== Fin de los Ejercicios de Funciones ===")
}
