package exercises

import "fmt"

func ExerciseMaps() {
	fmt.Println("=== Ejercicios de Maps ===")

	// TODO: Crear un map vacío para almacenar edades de personas y agregar al menos 3 entradas.
	edades := map[string]int{
		"Andrew": 20,
		"Luis": 25,
		"Jeronimo": 19,
	}

	// TODO: Acceder a un valor en el map y verificar si una clave existe.
	if edad, existe := edades["Andrew"]; existe {
		fmt.Printf("La edad de Andrew existe: %v\n", edad)
	}

	// TODO: Eliminar una entrada del map y verificar si se eliminó correctamente.
	delete(edades, "Jeronimo")
	if _, existe := edades["Jeronimo"]; !existe {
		fmt.Printf("Se elimino la entrada de 'Jeronimo'\n")
	}
	// TODO: Recorrer un map con un bucle y mostrar todas las claves y valores.
	for clave, valor := range edades {
		fmt.Printf("Clave: %v Valor: %v\n", clave, valor)
	}

	// TODO: Crear un map anidado para representar información de empleados (nombre, departamento, salario).
	empleados := map[string]map[string]interface{}{
		"E001": {
			"nombre": "Andrew",
			"departamento": "Ingenieria",
			"salario": 2000,
		},
	}

	for id, empleados := range empleados {
		fmt.Printf("ID: %v\n", id)
		for atributo, propiedad := range empleados {
			fmt.Printf("	%v: %v\n", atributo, propiedad)
		}
	}

	// TODO: Crear un map donde las claves sean strings y los valores sean slices, y agregar elementos a los slices.
	notas := map[string][]int{
		"Andrew": {4, 5, 3, 4, 2},
		"Luis": {5, 5, 0, 2, 3},
	}
	fmt.Printf("Map antes de agregar cosas: %v\n", notas)
	notas["Andrew"] = append(notas["Andrew"], 5)
	notas["Luis"] = append(notas["Luis"], 2)
	fmt.Printf("Map despues de agregar cosas: %v\n", notas)

	// TODO: Contar la frecuencia de palabras en un slice de strings usando un map.
	mensaje := []string{"hola", "como", "estas", "hola", "hola", "como"}
	conteo := make(map[string]int)

	for _, palabra := range mensaje {
		conteo[palabra] += 1
	}

	fmt.Printf("  %v\n", conteo)

	// TODO: Implementar un sistema de cache simple usando un map para almacenar resultados de cálculos.

	cache := make(map[int]int)

	calcular := func(n int) int {
		if resultado, existe := cache[n]; existe {
			fmt.Printf("Resultado de %d obtenido del cache\n", n)
			return resultado
		}
		resultado := n * n
		cache[n] = resultado
		return resultado
	}

	fmt.Println(calcular(4))
	fmt.Println(calcular(5))
	fmt.Println(calcular(4))

	fmt.Println("=== Fin de los Ejercicios de Maps ===")
}
