package exercises

import "fmt"

const (
	_ = iota
)

const (
	Admin = iota
	Usuario
	Invitado
)

func ExerciseEnums() {
	fmt.Println("=== Ejercicios de Enumeraciones ===")

	// TODO: Crear una función que reciba un día de la semana (como entero) y devuelva su nombre usando un switch.
	imprimirDia := func(dia int) {
		switch dia {
		case 1: 
			fmt.Println("Hoy es Lunes")
		case 2:
			fmt.Println("Hoy es Martes")
		case 3:
			fmt.Println("Hoy es Miércoles")
		case 4:
			fmt.Println("Hoy es Jueves")
		case 5:
			fmt.Println("Hoy es Viernes")
		case 6:
			fmt.Println("Hoy es Sábado")
		case 7:
			fmt.Println("Hoy es Domingo")
		default:
			fmt.Println("Día no válido")
		}
	}

	imprimirDia(1)
	imprimirDia(4)
	imprimirDia(7)

	// TODO: Crear un enum con iota para representar niveles de acceso (Admin, Usuario, Invitado) y mostrar el nivel de un usuario.

	mostrarNivel := func(nivel int) {
		switch nivel {
		case 1: 
			fmt.Println("Es Admin")
		case 2:
			fmt.Println("Es Usuario")
		case 3:
			fmt.Println("Es Invitado")
		}
	}

	mostrarNivel(1)
	mostrarNivel(2)
	mostrarNivel(3)

	// TODO: Usar un enum con valores personalizados para representar colores y mostrar un mensaje según el color seleccionado.
	const (
		Rojo = iota + 1
		Verde
		Azul
	)

	mostrarColor := func(color int) {
		switch color {
		case Rojo:
			fmt.Println("El color seleccionado es Rojo")
		case Verde:
			fmt.Println("El color seleccionado es Verde")
		case Azul:
			fmt.Println("El color seleccionado es Azul")
		default:
			fmt.Println("Color no válido")
		}
	}

	mostrarColor(Rojo)
	mostrarColor(Verde)
	mostrarColor(Azul)
	mostrarColor(0)

	fmt.Println("=== Fin de los Ejercicios de Enumeraciones ===")
}
