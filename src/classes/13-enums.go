package classes

import "fmt"

// Enumeraciones en Go
// Go no tiene soporte directo para enums como otros lenguajes (por ejemplo, Java o C#).
// Sin embargo, podemos simular enums usando constantes y el tipo iota.

// Definición de un enum usando iota
const (
	Lunes = iota // iota comienza en 0 y se incrementa automáticamente
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo
)

// Otra forma de usar enums con valores personalizados
const (
	Rojo  = "rojo"
	Verde = "verde"
	Azul  = "azul"
)

func Enums() {
	fmt.Println("=== Enumeraciones en Go - Guía Completa ===")

	// 1. ¿Qué son las enumeraciones?
	fmt.Println("\n--- 1. ¿Qué son las Enumeraciones? ---")
	fmt.Println("Las enumeraciones son una forma de definir un conjunto de valores constantes con nombres.")
	fmt.Println("En Go, se simulan usando constantes y el tipo iota.")

	// 2. Uso básico de iota
	fmt.Println("\n--- 2. Uso Básico de iota ---")
	fmt.Printf("Lunes: %d, Martes: %d, Miercoles: %d\n", Lunes, Martes, Miercoles)
	fmt.Printf("Jueves: %d, Viernes: %d, Sabado: %d, Domingo: %d\n", Jueves, Viernes, Sabado, Domingo)

	// 3. Uso de enums con valores personalizados
	fmt.Println("\n--- 3. Uso de Enums con Valores Personalizados ---")
	fmt.Printf("Rojo: %s, Verde: %s, Azul: %s\n", Rojo, Verde, Azul)

	// 4. Ejemplo práctico: Días de la semana
	fmt.Println("\n--- 4. Ejemplo Práctico: Días de la Semana ---")
	imprimirDia := func(dia int) {
		switch dia {
		case Lunes:
			fmt.Println("Hoy es Lunes")
		case Martes:
			fmt.Println("Hoy es Martes")
		case Miercoles:
			fmt.Println("Hoy es Miércoles")
		case Jueves:
			fmt.Println("Hoy es Jueves")
		case Viernes:
			fmt.Println("Hoy es Viernes")
		case Sabado:
			fmt.Println("Hoy es Sábado")
		case Domingo:
			fmt.Println("Hoy es Domingo")
		default:
			fmt.Println("Día no válido")
		}
	}
	imprimirDia(Lunes)
	imprimirDia(Viernes)
	imprimirDia(10) // Día no válido

	// 5. Consideraciones importantes
	fmt.Println("\n--- 5. Consideraciones Importantes ---")
	fmt.Println("• iota es útil para generar valores secuenciales automáticamente.")
	fmt.Println("• Las constantes en Go son inmutables y tienen un alcance de paquete.")
	fmt.Println("• Puedes usar enums con valores personalizados para mayor flexibilidad.")

	fmt.Println("\n=== Fin de la Clase de Enumeraciones ===")
}
