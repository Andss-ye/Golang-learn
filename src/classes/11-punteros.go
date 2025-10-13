package classes

import "fmt"

func Punteros() {
	fmt.Println("=== Punteros en Go - Guía Completa ===")

	// 1. ¿QUÉ SON LOS PUNTEROS?
	fmt.Println("\n--- 1. ¿Qué son los Punteros? ---")
	fmt.Println("Un puntero es una variable que almacena la dirección de memoria de otra variable.")
	fmt.Println("• Los punteros permiten modificar el valor original de una variable desde otra función.")
	fmt.Println("• Se declaran usando el operador * para el tipo y el operador & para obtener la dirección.")

	// 2. DECLARACIÓN Y USO DE PUNTEROS
	fmt.Println("\n--- 2. Declaración y Uso de Punteros ---")
	x := 10
	p := &x // p es un puntero a x
	fmt.Printf("Valor de x: %d\n", x)
	fmt.Printf("Dirección de x: %p\n", &x)
	fmt.Printf("Valor de p (dirección de x): %p\n", p)
	fmt.Printf("Valor al que apunta p: %d\n", *p)

	// Modificar el valor de x a través del puntero
	*p = 20
	fmt.Printf("Nuevo valor de x: %d\n", x)

	// 3. PUNTEROS Y FUNCIONES
	fmt.Println("\n--- 3. Punteros y Funciones ---")
	modificarValor := func(ptr *int) {
		*ptr = 50
	}
	fmt.Printf("Antes de modificar: %d\n", x)
	modificarValor(&x)
	fmt.Printf("Después de modificar: %d\n", x)

	// 4. PUNTEROS A STRUCTS
	fmt.Println("\n--- 4. Punteros a Structs ---")
	type Persona struct {
		nombre string
		edad   int
	}
	p1 := Persona{nombre: "Ana", edad: 30}
	p2 := &p1 // p2 es un puntero a p1
	fmt.Printf("Persona: %v\n", p1)
	fmt.Printf("Puntero a Persona: %p\n", p2)
	fmt.Printf("Nombre desde el puntero: %s\n", p2.nombre)

	// Modificar valores a través del puntero
	p2.edad = 35
	fmt.Printf("Nueva edad de la persona: %d\n", p1.edad)

	// 5. PUNTEROS A PUNTEROS
	fmt.Println("\n--- 5. Punteros a Punteros ---")
	ptr := &x
	pptr := &ptr
	fmt.Printf("Valor de x: %d\n", x)
	fmt.Printf("Dirección de x: %p\n", ptr)
	fmt.Printf("Dirección de ptr: %p\n", pptr)
	fmt.Printf("Valor al que apunta pptr: %d\n", **pptr)

	// 6. NIL Y PUNTEROS
	fmt.Println("\n--- 6. Nil y Punteros ---")
	var punteroNulo *int
	if punteroNulo == nil {
		fmt.Println("El puntero está inicializado a nil (no apunta a ninguna dirección válida).")
	}

	// 7. PUNTEROS EN SLICES Y MAPS
	fmt.Println("\n--- 7. Punteros en Slices y Maps ---")
	slice := []int{1, 2, 3}
	ptrSlice := &slice[0]
	fmt.Printf("Primer elemento del slice: %d\n", *ptrSlice)
	*ptrSlice = 10
	fmt.Printf("Slice modificado: %v\n", slice)

	// 8. CONSIDERACIONES IMPORTANTES
	fmt.Println("\n--- 8. Consideraciones Importantes ---")
	fmt.Println("• Los punteros son útiles para evitar copias innecesarias de datos grandes.")
	fmt.Println("• Usar punteros incorrectamente puede causar errores como accesos a memoria no válida.")
	fmt.Println("• En Go, el recolector de basura maneja la memoria, por lo que no necesitas liberar memoria manualmente.")

	fmt.Println("\n=== Fin de la Clase de Punteros ===")
}
