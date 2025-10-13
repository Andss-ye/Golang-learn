package main

import (
	"bufio"
	"fmt"
	"go-learning/src/classes"
	"go-learning/src/exercises"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMainMenu()

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

		if !mainExecute(option) {
			break
		}

		fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	}

	fmt.Println("¡Hasta luego! 👋")
}

func showMainMenu() {
	fmt.Println("=== Menú Principal ===")
	fmt.Println("1. Ver Clases")
	fmt.Println("2. Ver Ejercicios")
	fmt.Println("0. Salir")
	fmt.Println("=====================")
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
	fmt.Println("8. Slices - Estructura de Slices en Go")
	fmt.Println("9. Maps - Estructura de Mapas en Go")
	fmt.Println("10. Funciones - Declaración y uso de funciones")
	fmt.Println("11. Punteros - Introducción a punteros en Go")
	fmt.Println("12. Structs - Introducción a structs en Go")
	fmt.Println("0. Volver al menú principal")
	fmt.Println("========================")
}

func mainExecute(option int) bool {
	switch option {
	case 1:
		return showClassesSubMenu()
	case 2:
		return showExercisesSubMenu()
	case 0:
		return false
	default:
		fmt.Println("❌ Opción no válida. Por favor, selecciona 1, 2 o 0 para salir.")
		return true
	}
}

func showClassesSubMenu() bool {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()

		fmt.Print("Selecciona una clase: ")
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

	return true
}

func showExercisesSubMenu() bool {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showExercisesMenu()

		fmt.Print("Selecciona un ejercicio: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		option, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("❌ Por favor, ingresa un número válido")
			continue
		}

		if !executeExercises(option) {
			break
		}

		fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	}

	return true
}

func showExercisesMenu() {
	fmt.Println("💪 Ejercicios de Práctica Go")
	fmt.Println("============================")
	fmt.Println("1. Variables Básicas - Declaración y tipos")
	fmt.Println("2. Condicionales - If/else y switch")
	fmt.Println("3. Ciclos - Loops")
	fmt.Println("4. Arrays")
	fmt.Println("5. Slices")
	fmt.Println("6. Proyecto Integrador")
	fmt.Println("0. Volver al menú principal")
	fmt.Println("============================")
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
	case 8:
		fmt.Println("\n📚 Ejecutando Clase 8: Slices")
		fmt.Println("-----------------------------")
		classes.Slices()
		return true
	case 9:
		fmt.Println("\n🚀 Ejecutando Clase 9: Mapas")
		fmt.Println("-----------------------------")
		classes.Maps()
		return true
	case 10:
		fmt.Println("\n🔧 Ejecutando Clase 10: Funciones")
		fmt.Println("-----------------------------")
		classes.Functions()
		return true
	case 11:
		fmt.Println("\n📍 Ejecutando Clase 11: Punteros")
		fmt.Println("-----------------------------")
		classes.Punteros()
		return true
	case 12:
		fmt.Println("\n🏗️ Ejecutando Clase 12: Structs")
		fmt.Println("-----------------------------")
		classes.Structs()
		return true
	case 0:
		return false
	default:
		fmt.Println("❌ Opción no válida. Por favor, selecciona 1, 2 o 0 para salir.")
		return true
	}
}

func executeExercises(option int) bool {
	switch option {
	case 1:
		fmt.Println("\n🔍 Ejercicio 1: Variables Básicas")
		fmt.Println("----------------------------------")
		exercises.VariablesBasicas()
		return true
	case 2:
		fmt.Println("\n🔀 Ejercicio 2: Condicionales")
		fmt.Println("------------------------------")
		exercises.Condicionales()
		return true
	case 3:
		fmt.Println("\n🔀 Ejercicio 3: loops")
		fmt.Println("------------------------------")
		exercises.BuclesFor()
		return true
	case 4:
		fmt.Println("\n🔀 Ejercicio 4: Arrays")
		fmt.Println("------------------------------")
		exercises.Arrays()
		return true
	case 5:
		fmt.Println("\n🔀 Ejercicio 5: Slices")
		fmt.Println("------------------------------")
		exercises.SlicesAvanzados()
		return true
	case 6:
		fmt.Println("\n🔀 Ejercicio 6: ProyectoIntegrador")
		fmt.Println("------------------------------")
		exercises.ProyectoIntegrador()
		return true
	case 0:
		return false
	default:
		fmt.Println("❌ Opción no válida. Por favor, selecciona del 1-6 o 0 para volver.")
		return true
	}
}
