package classes

import ( 
	"fmt"
	"time"
)

func Switch() {
	fmt.Println("=== Ejemplos de Switch en Go ===")

	i := 2

	switch i {
	case 1:
		fmt.Println("El valor es 1")
	case 2:
		fmt.Println("El valor es 2")
	case 3:
		fmt.Println("El valor es 3")
	default:
		fmt.Println("El valor no es 1, 2 ni 3")
	}
	
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("Es fin de semana")
		default:
			fmt.Println("Es un día laborable")
	}
}