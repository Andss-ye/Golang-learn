package classes

import (
	"errors"
	"fmt"
)

// Función para dividir dos números
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir entre cero")
	}
	return a / b, nil
}

// Tipo de error personalizado
type MiError struct {
	Mensaje string
	Codigo  int
}

func (e MiError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Codigo, e.Mensaje)
}

func manejarPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado de un panic:", r)
	}
}

func causarPanic() {
	defer manejarPanic()
	panic("esto es un panic")
}

func Errors() {
	fmt.Println("=== Manejo de Errores en Go - Guía Completa ===")

	// 1. ¿Qué son los errores en Go?
	fmt.Println("\n--- 1. ¿Qué son los Errores en Go? ---")
	fmt.Println("En Go, los errores son valores. Esto significa que puedes devolver errores como cualquier otro valor.")
	fmt.Println("El paquete 'errors' proporciona funciones para crear y manejar errores.")

	// 2. Crear un error simple
	fmt.Println("\n--- 2. Crear un Error Simple ---")
	err := errors.New("este es un error simple")
	fmt.Println("Error creado:", err)

	// 3. Devolver errores desde una función
	fmt.Println("\n--- 3. Devolver Errores desde una Función ---")
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error al dividir:", err)
	} else {
		fmt.Println("Resultado de la división:", resultado)
	}

	// 4. Usar fmt.Errorf para errores formateados
	fmt.Println("\n--- 4. Usar fmt.Errorf para Errores Formateados ---")
	errFormateado := fmt.Errorf("el archivo %s no se pudo abrir", "data.txt")
	fmt.Println("Error formateado:", errFormateado)

	// 5. Comparar errores
	fmt.Println("\n--- 5. Comparar Errores ---")
	err1 := errors.New("error único")
	err2 := errors.New("error único")
	if err1 == err2 {
		fmt.Println("Los errores son iguales")
	} else {
		fmt.Println("Los errores son diferentes")
	}

	// 6. Errores personalizados
	fmt.Println("\n--- 6. Errores Personalizados ---")
	errPersonalizado := MiError{Mensaje: "Algo salió mal", Codigo: 404}
	fmt.Println("Error personalizado:", errPersonalizado)

	// 7. Manejo de errores con defer, panic y recover
	fmt.Println("\n--- 7. Manejo de Errores con defer, panic y recover ---")
	fmt.Println("Causando un panic...")
	causarPanic()

	// 8. Consideraciones importantes
	fmt.Println("\n--- 8. Consideraciones Importantes ---")
	fmt.Println("• Siempre verifica los errores devueltos por las funciones.")
	fmt.Println("• Usa errores personalizados para proporcionar más contexto.")
	fmt.Println("• Evita usar panic/recover para manejar errores comunes.")

	fmt.Println("\n=== Fin de la Clase de Manejo de Errores ===")
}
