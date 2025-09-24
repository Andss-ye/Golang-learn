package main

import (
	"fmt"
	"os"
)

func main() {
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
}