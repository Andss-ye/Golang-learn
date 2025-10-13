package exercises

import "fmt"

type Persona struct {
	nombre string
	edad int
}

func ExercisePointers() {
	fmt.Println("=== Ejercicios de Punteros ===")

	// TODO: Crear una variable entera y un puntero que apunte a ella. Modificar el valor usando el puntero.
	number1 := 4
	ptr := &number1
	*ptr = 5

	// TODO: Crear una función que reciba un puntero a un entero y modifique su valor.
	prueba := func(ptr *int) {
		*ptr = 565
	}

	prueba(&number1)
	// TODO: Crear una struct y un puntero a esa struct. Modificar los campos de la struct usando el puntero.
	persona1 := Persona{nombre: "andrew", edad: 19}
	personaPtr := &persona1
	personaPtr.nombre = "juan"
	personaPtr.edad = 25
	fmt.Println("Persona modificada:", persona1)


	// TODO: Crear un puntero a un puntero y acceder al valor original.
	num := 10
	ptrNum := &num
	ptrPtrNum := &ptrNum
	fmt.Println("Valor original usando puntero a puntero:", **ptrPtrNum)

	// TODO: Crear un slice y un puntero al primer elemento. Modificar el slice usando el puntero.
	slice := []int{1, 2, 3}
	ptrSlice := &slice[0]
	*ptrSlice = 100
	fmt.Println("Slice modificado:", slice)

	// TODO: Verificar si un puntero está inicializado a nil.
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("El puntero está inicializado a nil")
	} else {
		fmt.Println("El puntero no está inicializado a nil")
	}

	fmt.Println("=== Fin de los Ejercicios de Punteros ===")
}
