package classes

import (
	"fmt"
	"os"
)

func Packages() {
	fmt.Println("Hello, World! from Go!!!!")

	envVar := os.Getenv("HOME")

	if envVar == "" {
		fmt.Println("La variable de HOME no existe :)")
	} else {
		fmt.Printf("El valor de HOME es: %s\n", envVar)
	}

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("Error creando el archivo: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("✅ Archivo test.txt creado exitosamente")

	// Escribir contenido al archivo
	_, err = file.WriteString("¡Hola desde Go!\nEste archivo fue creado por la clase Packages.")
	if err != nil {
		fmt.Printf("Error escribiendo al archivo: %v\n", err)
		return
	}

	fmt.Println("✅ Contenido escrito al archivo")
}
