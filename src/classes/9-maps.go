package classes

import "fmt"

func Maps() {
	fmt.Println("=== Maps en Go - Guía Completa ===")

	// 1. ¿QUÉ SON LOS MAPS?
	fmt.Println("\n--- 1. ¿Qué son los Maps? ---")
	fmt.Println("Los maps son estructuras de datos clave-valor:")
	fmt.Println("• Similar a diccionarios en Python o objetos en JavaScript")
	fmt.Println("• Cada valor se identifica por una clave única")
	fmt.Println("• Las claves deben ser comparables (strings, números, bools)")
	fmt.Println("• Los valores pueden ser de cualquier tipo")

	// 2. DECLARACIÓN DE MAPS
	fmt.Println("\n--- 2. Declaración de Maps ---")

	// Map vacío con make
	edades := make(map[string]int)
	fmt.Printf("Map vacío: %v\n", edades)

	// Map con valores iniciales
	capitales := map[string]string{
		"España":   "Madrid",
		"Francia":  "París",
		"Italia":   "Roma",
		"Portugal": "Lisboa",
	}
	fmt.Printf("Map de capitales: %v\n", capitales)

	// Map de diferentes tipos
	estudiante := map[string]interface{}{
		"nombre":       "Ana García",
		"edad":         22,
		"calificacion": 8.7,
		"activo":       true,
	}
	fmt.Printf("Estudiante: %v\n", estudiante)

	// 3. OPERACIONES BÁSICAS
	fmt.Println("\n--- 3. Operaciones Básicas ---")

	// Agregar elementos
	edades["Juan"] = 25
	edades["María"] = 30
	edades["Carlos"] = 28
	fmt.Printf("Después de agregar: %v\n", edades)

	// Acceder a elementos
	edadJuan := edades["Juan"]
	fmt.Printf("Edad de Juan: %d\n", edadJuan)

	// Verificar si una clave existe
	edad, existe := edades["Pedro"]
	if existe {
		fmt.Printf("Pedro tiene %d años\n", edad)
	} else {
		fmt.Println("Pedro no está en el map")
	}

	// Modificar elementos
	edades["Juan"] = 26
	fmt.Printf("Juan ahora tiene: %d años\n", edades["Juan"])

	// 4. ELIMINAR ELEMENTOS
	fmt.Println("\n--- 4. Eliminar Elementos ---")

	fmt.Printf("Antes de eliminar: %v\n", edades)
	delete(edades, "Carlos")
	fmt.Printf("Después de eliminar a Carlos: %v\n", edades)

	// Verificar si se eliminó
	_, existe = edades["Carlos"]
	fmt.Printf("¿Carlos existe? %t\n", existe)

	// 5. RECORRER MAPS
	fmt.Println("\n--- 5. Recorrer Maps ---")

	fmt.Println("Recorriendo con range:")
	for pais, capital := range capitales {
		fmt.Printf("  %s -> %s\n", pais, capital)
	}

	// Solo claves
	fmt.Println("\nSolo las claves:")
	for pais := range capitales {
		fmt.Printf("  • %s\n", pais)
	}

	// Solo valores
	fmt.Println("\nSolo los valores:")
	for _, capital := range capitales {
		fmt.Printf("  • %s\n", capital)
	}

	// 6. MAPS ANIDADOS
	fmt.Println("\n--- 6. Maps Anidados ---")

	// Map de maps para representar información de empleados
	empleados := map[string]map[string]interface{}{
		"E001": {
			"nombre":       "Luis Martín",
			"departamento": "IT",
			"salario":      45000,
			"activo":       true,
		},
		"E002": {
			"nombre":       "Ana López",
			"departamento": "Marketing",
			"salario":      42000,
			"activo":       true,
		},
		"E003": {
			"nombre":       "Carlos Ruiz",
			"departamento": "Ventas",
			"salario":      38000,
			"activo":       false,
		},
	}

	fmt.Println("Información de empleados:")
	for id, info := range empleados {
		fmt.Printf("ID: %s\n", id)
		fmt.Printf("  Nombre: %s\n", info["nombre"])
		fmt.Printf("  Departamento: %s\n", info["departamento"])
		fmt.Printf("  Salario: €%d\n", info["salario"])
		fmt.Printf("  Activo: %t\n", info["activo"])
		fmt.Println()
	}

	// 7. MAPS CON SLICES COMO VALORES
	fmt.Println("\n--- 7. Maps con Slices ---")

	// Map donde cada clave tiene múltiples valores
	materias := map[string][]string{
		"Matemáticas": {"Álgebra", "Cálculo", "Geometría"},
		"Ciencias":    {"Física", "Química", "Biología"},
		"Idiomas":     {"Español", "Inglés", "Francés"},
	}

	fmt.Println("Materias por área:")
	for area, lista := range materias {
		fmt.Printf("%s:\n", area)
		for _, materia := range lista {
			fmt.Printf("  - %s\n", materia)
		}
	}

	// Agregar elemento a un slice dentro del map
	materias["Ciencias"] = append(materias["Ciencias"], "Geología")
	fmt.Printf("\nCiencias actualizado: %v\n", materias["Ciencias"])

	// 8. CONTADORES CON MAPS
	fmt.Println("\n--- 8. Contadores con Maps ---")

	// Contar frecuencia de palabras
	texto := []string{"go", "es", "genial", "go", "es", "poderoso", "go", "rocks"}
	contador := make(map[string]int)

	for _, palabra := range texto {
		contador[palabra]++
	}

	fmt.Println("Frecuencia de palabras:")
	for palabra, frecuencia := range contador {
		fmt.Printf("  '%s': %d veces\n", palabra, frecuencia)
	}

	// Encontrar la palabra más frecuente
	maxPalabra := ""
	maxFrecuencia := 0

	for palabra, frecuencia := range contador {
		if frecuencia > maxFrecuencia {
			maxFrecuencia = frecuencia
			maxPalabra = palabra
		}
	}

	fmt.Printf("Palabra más frecuente: '%s' (%d veces)\n", maxPalabra, maxFrecuencia)

	// 9. MAPS COMO CACHE/LOOKUP
	fmt.Println("\n--- 9. Maps como Cache ---")

	// Cache de resultados costosos de calcular
	factorialCache := make(map[int]int)

	calcularFactorial := func(n int) int {
		if resultado, existe := factorialCache[n]; existe {
			fmt.Printf("  Factorial de %d encontrado en cache: %d\n", n, resultado)
			return resultado
		}

		resultado := 1
		for i := 1; i <= n; i++ {
			resultado *= i
		}

		factorialCache[n] = resultado
		fmt.Printf("  Factorial de %d calculado y guardado: %d\n", n, resultado)
		return resultado
	}

	// Probar el cache
	fmt.Println("Calculando factoriales:")
	calcularFactorial(5)
	calcularFactorial(3)
	calcularFactorial(5) // Este debería usar cache
	calcularFactorial(4)

	fmt.Printf("Cache actual: %v\n", factorialCache)

	// 10. MAPS DE FUNCIONES
	fmt.Println("\n--- 10. Maps de Funciones ---")

	// Map que almacena funciones como valores
	operaciones := map[string]func(float64, float64) float64{
		"suma": func(a, b float64) float64 {
			return a + b
		},
		"resta": func(a, b float64) float64 {
			return a - b
		},
		"multiplicacion": func(a, b float64) float64 {
			return a * b
		},
		"division": func(a, b float64) float64 {
			if b != 0 {
				return a / b
			}
			return 0
		},
	}

	// Usar las funciones del map
	a, b := 10.0, 3.0
	for nombre, funcion := range operaciones {
		resultado := funcion(a, b)
		fmt.Printf("%s(%.1f, %.1f) = %.2f\n", nombre, a, b, resultado)
	}

	// 11. CASOS PRÁCTICOS
	fmt.Println("\n--- 11. Casos Prácticos ---")

	// Sistema de inventario
	inventario := map[string]map[string]interface{}{
		"LAPTOP001": {
			"nombre":    "MacBook Pro",
			"precio":    1299.99,
			"stock":     15,
			"categoria": "Electrónicos",
		},
		"LIBRO001": {
			"nombre":    "Go Programming",
			"precio":    39.99,
			"stock":     50,
			"categoria": "Libros",
		},
	}

	// Función para mostrar producto
	mostrarProducto := func(codigo string) {
		if producto, existe := inventario[codigo]; existe {
			fmt.Printf("Producto: %s\n", codigo)
			fmt.Printf("  Nombre: %s\n", producto["nombre"])
			fmt.Printf("  Precio: €%.2f\n", producto["precio"])
			fmt.Printf("  Stock: %d unidades\n", producto["stock"])
			fmt.Printf("  Categoría: %s\n", producto["categoria"])
		} else {
			fmt.Printf("Producto %s no encontrado\n", codigo)
		}
	}

	fmt.Println("Consultando inventario:")
	mostrarProducto("LAPTOP001")
	mostrarProducto("LIBRO001")
	mostrarProducto("NOEXISTE")

	// 12. FUNCIONES ÚTILES CON MAPS
	fmt.Println("\n--- 12. Funciones Útiles ---")

	// Función para obtener todas las claves
	obtenerClaves := func(m map[string]int) []string {
		claves := make([]string, 0, len(m))
		for clave := range m {
			claves = append(claves, clave)
		}
		return claves
	}

	// Función para obtener todos los valores
	obtenerValores := func(m map[string]int) []int {
		valores := make([]int, 0, len(m))
		for _, valor := range m {
			valores = append(valores, valor)
		}
		return valores
	}

	// Función para verificar si un valor existe
	existeValor := func(m map[string]int, buscar int) bool {
		for _, valor := range m {
			if valor == buscar {
				return true
			}
		}
		return false
	}

	// Probar las funciones
	puntuaciones := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}

	fmt.Printf("Puntuaciones: %v\n", puntuaciones)
	fmt.Printf("Claves: %v\n", obtenerClaves(puntuaciones))
	fmt.Printf("Valores: %v\n", obtenerValores(puntuaciones))
	fmt.Printf("¿Existe puntuación 92? %t\n", existeValor(puntuaciones, 92))
	fmt.Printf("¿Existe puntuación 100? %t\n", existeValor(puntuaciones, 100))

	// 13. RENDIMIENTO Y CONSIDERACIONES
	fmt.Println("\n--- 13. Consideraciones Importantes ---")
	fmt.Printf("Longitud del map de edades: %d\n", len(edades))
	fmt.Printf("Longitud del map de capitales: %d\n", len(capitales))

	fmt.Println("\n=== Puntos Clave de Maps ===")
	fmt.Println("✅ Los maps son de referencia, no de valor")
	fmt.Println("✅ El orden de iteración no está garantizado")
	fmt.Println("✅ Las claves deben ser comparables")
	fmt.Println("✅ Usar make() para maps vacíos")
	fmt.Println("✅ Verificar existencia con value, ok := map[key]")
	fmt.Println("✅ Usar delete() para eliminar elementos")
	fmt.Println("✅ len() devuelve el número de pares clave-valor")
	fmt.Println("✅ Excelentes para lookups rápidos O(1) promedio")
}
