package main

import (
	"bufio"
	"fmt"
	"go-classes/src/classes"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()

		fmt.Print("Selecciona una opción: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		option, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("❌ Por favor, ingresa un número válido")
			continue
		}

		if !executeClass(option) {
			break
		}

		fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	}

	fmt.Println("¡Hasta luego! 👋")
}

func showMenu() {
	fmt.Println("🚀 Sistema de Clases Go")
	fmt.Println("========================")
	fmt.Println("1. Packages - Trabajo con paquetes y archivos")
	fmt.Println("2. Variables - Declaración y uso de variables")
	fmt.Println("3. Valores - Tipos de datos y valores en Go")
	fmt.Println("4. Condiciones - Estructuras condicionales en Go")
	fmt.Println("5. For - Bucles For en Go")
	fmt.Println("6. Switch - Estructura Switch en Go")
	fmt.Println("7. Arrays - Estructura de Arrays en Go")
	fmt.Println("0. Salir")
	fmt.Println("========================")
}

func executeClass(option int) bool {
	switch option {
	case 1:
		fmt.Println("\n📦 Ejecutando Clase 1: Packages")
		fmt.Println("-------------------------------")
		classes.Packages()
		return true
	case 2:
		fmt.Println("\n🔤 Ejecutando Clase 2: Variables")
		fmt.Println("--------------------------------")
		classes.Variables()
		return true
	case 3:
		fmt.Println("\n🔢 Ejecutando Clase 3: Valores")
		fmt.Println("--------------------------------")
		classes.Valores()
		return true
	case 4:
		fmt.Println("\n🔀 Ejecutando Clase 4: Condiciones")
		fmt.Println("----------------------------------")
		classes.Condiciones()
		return true
	case 5:
		fmt.Println("\n🔁 Ejecutando Clase 5: For")
		fmt.Println("---------------------------")
		classes.For()
		return true
	case 6:
		fmt.Println("\n🔄 Ejecutando Clase 6: Switch")
		fmt.Println("-----------------------------")
		classes.Switch()
		return true
	case 7:
		fmt.Println("\n🔢 Ejecutando Clase 7: Arrays")
		fmt.Println("-----------------------------")
		classes.Arrays()
		return true
	case 0:
		return false
	default:
		fmt.Println("❌ Opción no válida. Por favor, selecciona 1, 2 o 0 para salir.")
		return true
	}
}
