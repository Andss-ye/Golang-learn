package classes

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
	Activo bool
}

func (p Persona) Saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años.\n", p.Nombre, p.Edad)
}

func (p *Persona) Activar() {
	p.Activo = true
}

type Direccion struct {
	Calle  string
	Ciudad string
}

type Empleado struct {
	Persona
	Puesto    string
	Direccion Direccion
}

func Structs() {
	fmt.Println("=== Structs en Go - Guía Completa ===")

	// 1. ¿QUÉ SON LAS STRUCTS?
	fmt.Println("\n--- 1. ¿Qué son las Structs? ---")
	fmt.Println("Las structs son tipos de datos personalizados que agrupan campos bajo un mismo nombre.")
	fmt.Println("• Son similares a las clases en otros lenguajes, pero sin herencia.")
	fmt.Println("• Se utilizan para modelar entidades con múltiples atributos.")

	// 2. DECLARACIÓN DE UNA STRUCT
	fmt.Println("\n--- 2. Declaración de una Struct ---")

	// Crear una instancia de una struct
	p1 := Persona{Nombre: "Ana", Edad: 30, Activo: true}
	fmt.Printf("Persona: %+v\n", p1)

	// Acceder a los campos
	fmt.Printf("Nombre: %s, Edad: %d, Activo: %t\n", p1.Nombre, p1.Edad, p1.Activo)

	// Modificar campos
	p1.Edad = 31
	fmt.Printf("Nueva edad: %d\n", p1.Edad)

	// 3. PUNTEROS A STRUCTS
	fmt.Println("\n--- 3. Punteros a Structs ---")
	p2 := &p1
	p2.Nombre = "Luis"
	fmt.Printf("Persona modificada: %+v\n", p1)

	// 4. MÉTODOS EN STRUCTS
	fmt.Println("\n--- 4. Métodos en Structs ---")
	p1.Saludar()
	p1.Activar()
	fmt.Printf("¿Activo?: %t\n", p1.Activo)

	// 5. ANIDACIÓN DE STRUCTS
	fmt.Println("\n--- 5. Anidación de Structs ---")
	e1 := Empleado{
		Persona:   Persona{Nombre: "Carlos", Edad: 25, Activo: true},
		Puesto:    "Desarrollador",
		Direccion: Direccion{Calle: "Av. Principal", Ciudad: "Madrid"},
	}
	fmt.Printf("Empleado: %+v\n", e1)
	fmt.Printf("Ciudad: %s\n", e1.Direccion.Ciudad)

	// 6. COMPARACIÓN DE STRUCTS
	fmt.Println("\n--- 6. Comparación de Structs ---")
	p3 := Persona{Nombre: "Ana", Edad: 30, Activo: true}
	p4 := Persona{Nombre: "Ana", Edad: 30, Activo: true}
	fmt.Printf("¿p3 y p4 son iguales?: %t\n", p3 == p4)

	// 7. USO DE STRUCTS EN SLICES Y MAPS
	fmt.Println("\n--- 7. Uso de Structs en Slices y Maps ---")
	personas := []Persona{
		{Nombre: "Juan", Edad: 20, Activo: true},
		{Nombre: "María", Edad: 25, Activo: false},
	}
	for _, persona := range personas {
		fmt.Printf("%+v\n", persona)
	}

	empleados := map[string]Empleado{
		"E001": e1,
		"E002": {
			Persona:   Persona{Nombre: "Laura", Edad: 28, Activo: true},
			Puesto:    "Diseñadora",
			Direccion: Direccion{Calle: "Calle 2", Ciudad: "Barcelona"},
		},
	}
	for id, empleado := range empleados {
		fmt.Printf("ID: %s, Nombre: %s, Puesto: %s\n", id, empleado.Nombre, empleado.Puesto)
	}

	// 8. CONSIDERACIONES IMPORTANTES
	fmt.Println("\n--- 8. Consideraciones Importantes ---")
	fmt.Println("• Las structs son valores por defecto, no referencias.")
	fmt.Println("• Usa punteros si necesitas modificar el valor original.")
	fmt.Println("• Las structs pueden anidarse para modelar estructuras complejas.")

	fmt.Println("\n=== Fin de la Clase de Structs ===")
}
