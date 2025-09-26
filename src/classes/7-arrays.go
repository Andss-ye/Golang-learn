package classes

import (
	"fmt"
)

func Arrays() {
	fmt.Println("=== Ejemplos de Arrays en Go ===")

	var array[5]float32

	fmt.Println(array)
	array[4] = 100.0
	array[3] = 1232.42
	fmt.Println(array)

	array = [5]float32{1, 2, 3, 4, 5}
	fmt.Println(array)

	var newArray = [5]float32{1, 2, 3, 4, 5.3434}
	fmt.Println(newArray)

	arrayDynamic := [...]int{}
	fmt.Println(arrayDynamic, len(arrayDynamic))
}