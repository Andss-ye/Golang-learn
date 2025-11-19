package exercises

import (
	"errors"
	"fmt"
)

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir entre cero")
	}
	return a / b, nil
}

type MiError struct {
	Codigo  string
	Mensaje string
}

func (e MiError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Codigo, e.Mensaje)
}

func validarEdad(edad int) error {
	if edad < 18 {
		return MiError{
			Codigo:  "EDAD_INVALIDA",
			Mensaje: "La edad debe ser mayor o igual a 18",
		}
	}
	return nil
}

func manejarErrorCritico() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se recuperó de un error crítico:", r)
		}
	}()
	panic("¡Error crítico detectado!")
}

func ExerciseErrors() {
	fmt.Println("=== Ejercicios de Manejo de Errores ===")

	// TODO: Crear una función que divida dos números y maneje el error de división por cero.
	resultado, err := division(23.3, 0)
	if err != nil {
		fmt.Println("error al dividir", err)
	} else {
		fmt.Println("Resultado de la division", resultado)
	}

	// TODO: Crear un tipo de error personalizado y usarlo en una función que valide la edad de un usuario.
	err = validarEdad(16)
	if err != nil {
		fmt.Println("Error al validar edad:", err)
	} else {
		fmt.Println("Edad válida")
	}

	// TODO: Usar defer, panic y recover para manejar un error crítico en una función.
	manejarErrorCritico()

	fmt.Println("=== Fin de los Ejercicios de Manejo de Errores ===")
}
