package exercises

import "fmt"

type PersonaIntegrador struct {
	Nombre string
	Edad int
	Activo bool
}

func agregar(persona PersonaIntegrador, personas map[string]PersonaIntegrador, id string) {
	personas[id] = persona
}

func modificar(id string, personas map[string]PersonaIntegrador, personaModificada PersonaIntegrador) {
	personas[id] = personaModificada
}

func eliminar(id string, personas map[string]PersonaIntegrador) {
	delete(personas, id)
}

func ExerciseIntegrador() {
	fmt.Println("=== Ejercicio Integrador ===")

	// TODO: Crear un sistema completo que utilice Maps, Funciones, Punteros y Structs.
	// 1. Crear una struct Persona con campos Nombre, Edad y Activo.
	// 2. Crear un map donde las claves sean strings (IDs) y los valores sean structs Persona.
	personas := make(map[string]PersonaIntegrador)
	// 3. Implementar funciones para agregar, actualizar y eliminar personas del map.
	andrew := PersonaIntegrador{
		Nombre: "andrew",
		Edad: 19,
		Activo: true,
	}
	erick := PersonaIntegrador{
		Nombre: "erick",
		Edad: 25,
		Activo: false,
	}
	agregar(andrew, personas, "P001")
	agregar(erick, personas, "P002")
	fmt.Printf("Personas despues de agregar:\n%v\n", personas)
	agregar(andrew, personas, "P001")
	modificar("P002", personas, PersonaIntegrador{
		Nombre: "erick",
		Edad: 26,
		Activo: true,
	})
	fmt.Printf("Personas despues de modificar:\n%v\n", personas)
	eliminar("P001", personas)
	fmt.Printf("Personas despues de eliminar:\n%v\n", personas)
	// 4. Usar punteros para modificar directamente los valores de las personas en el map.
	if persona, existe := personas["P002"]; existe {
		personaPtr := &persona
		personaPtr.Edad = 27
		personas["P002"] = *personaPtr
	}
	// 5. Crear una función que recorra el map y muestre todas las personas activas.
	mostrarPersonasActivas := func(personas map[string]PersonaIntegrador) {
		for id, persona := range personas {
			if persona.Activo {
				fmt.Printf("ID: %s, Persona: %+v\n", id, persona)
			}
		}
	}
	mostrarPersonasActivas(personas)
	// 6. Crear una función que reciba un slice de IDs y devuelva un slice con las personas correspondientes.
	obtenerPersonasPorID := func(ids []string, personas map[string]PersonaIntegrador) []PersonaIntegrador {
		var resultado []PersonaIntegrador
		for _, id := range ids {
			if persona, existe := personas[id]; existe {
				resultado = append(resultado, persona)
			}
		}
		return resultado
	}
	personasSeleccionadas := obtenerPersonasPorID([]string{"P002"}, personas)
	fmt.Printf("Personas seleccionadas: %v\n", personasSeleccionadas)
	// 7. Usar closures para mantener un contador de personas activas y actualizarlo dinámicamente.
	contadorPersonasActivas := func() func() int {
		contador := 0
		return func() int {
			contador++
			return contador
		}
	}()
	fmt.Printf("Contador1: %d\n", contadorPersonasActivas())
	fmt.Printf("Contador2: %d\n", contadorPersonasActivas())
	// 8. Implementar un sistema de búsqueda que permita encontrar personas por nombre usando un map auxiliar.

	fmt.Println("=== Fin del Ejercicio Integrador ===")
}