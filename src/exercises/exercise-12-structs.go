package exercises

import "fmt"

type Persona_struct struct {
	nombre string
	edad int
	activo bool
}

func (p Persona_struct) Saludar(tiempo string) {
	fmt.Printf("hola %v, tienes %v y estas %s a las %v\n", p.nombre, p.edad, map[bool]string{true: "activo", false: "inactivo"}[p.activo], tiempo)
}

func (p *Persona_struct) cambio_estado() {
	p.activo = !p.activo
}

type Empleado struct {
	Persona_struct
	puesto string
}

func ExerciseStructs() {
	fmt.Println("=== Ejercicios de Structs ===")

	// TODO: Crear una struct llamada Persona con campos Nombre, Edad y Activo. Crear una instancia y mostrar sus valores.
	andrew := Persona_struct{
		nombre: "andrew",
		edad: 19,
		activo: true,
	}
	fmt.Printf("nombre: %v\nedad: %v\nactivo: %v\n", andrew.nombre, andrew.edad, andrew.activo)

	// TODO: Crear un método para la struct Persona que imprima un saludo con el nombre y la edad.
	andrew.Saludar("21:05")

	// TODO: Crear un puntero a una struct Persona y modificar uno de sus campos.
	andrew.cambio_estado()
	fmt.Printf("nuevo estado de andrew: %v\n", andrew.activo)

	// TODO: Crear una struct llamada Empleado que anide la struct Persona y agregue un campo Puesto. Crear una instancia y mostrar sus valores.
	e1 := Empleado{
		Persona_struct: andrew,
		puesto: "Backend developer",
	}
	fmt.Printf("Empleado: %+v\n", e1)
	// TODO: Comparar dos structs Persona y verificar si son iguales.
	p1 := Persona_struct{nombre: "andrew", edad: 19, activo: true}
	p2 := Persona_struct{nombre: "andrew", edad: 19, activo: true}
	sonIguales := p1 == p2
	fmt.Printf("¿p1 y p2 son iguales? %v\n", sonIguales)

	// TODO: Crear un slice de structs Persona y recorrerlo con un bucle para mostrar sus valores.
	personas := []Persona_struct{
		{nombre: "andrew", edad: 19, activo: true},
		{nombre: "maria", edad: 22, activo: false},
		{nombre: "juan", edad: 30, activo: true},
	}
	fmt.Println("Slice de personas:")
	for i, p := range personas {
		fmt.Printf("%d: nombre=%v, edad=%v, activo=%v\n", i, p.nombre, p.edad, p.activo)
	}

	// TODO: Crear un map donde las claves sean strings y los valores sean structs Persona. Agregar entradas y mostrarlas.
	personaMap := map[string]Persona_struct{
		"andrew": {nombre: "andrew", edad: 19, activo: true},
		"maria":  {nombre: "maria", edad: 22, activo: false},
	}
	fmt.Println("Map de personas:")
	for k, v := range personaMap {
		fmt.Printf("clave=%v, valor={nombre=%v, edad=%v, activo=%v}\n", k, v.nombre, v.edad, v.activo)
	}

	fmt.Println("=== Fin de los Ejercicios de Structs ===")
}
