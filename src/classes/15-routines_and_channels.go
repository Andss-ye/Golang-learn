package classes

import (
	"fmt"
	"sync"
	"time"
)

func RoutinesAndChannels() {
	fmt.Println("=== Go Routines y Channels - Guía Básica ===")

	// 1. ¿Qué son las Go Routines?
	fmt.Println("\n--- 1. ¿Qué son las Go Routines? ---")
	fmt.Println("Las Go Routines son funciones que se ejecutan de manera concurrente.")
	fmt.Println("Se crean con la palabra clave 'go' seguida de una llamada a función.")
	fmt.Println("Ejemplo: go funcion()")

	// 2. Ejemplo básico: Crear una Go Routine
	fmt.Println("\n--- 2. Ejemplo Básico: Crear una Go Routine ---")
	fmt.Println("Ejecutando dos tareas de forma concurrente:")

	go func() {
		fmt.Println("Tarea 1: Ejecutándose en una Go Routine")
	}()

	fmt.Println("Tarea 2: Ejecutándose en el main")

	// Esperar para que la Go Routine termine
	time.Sleep(100 * time.Millisecond)

	// 3. Problema: Condiciones de carrera
	fmt.Println("\n--- 3. Problema: Condiciones de Carrera ---")
	fmt.Println("Sin sincronización, es difícil saber cuándo termina una Go Routine.")
	fmt.Println("Solución: Usar WaitGroup para sincronizar Go Routines.")

	// 4. Usar WaitGroup para sincronizar
	fmt.Println("\n--- 4. Usar WaitGroup para Sincronizar ---")
	var wg sync.WaitGroup

	// Agregar 2 Go Routines al contador de WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Go Routine 1 completada")
		wg.Done() // Marca como completada
	}()

	go func() {
		fmt.Println("Go Routine 2 completada")
		wg.Done() // Marca como completada
	}()

	wg.Wait() // Espera a que ambas Go Routines terminen
	fmt.Println("Todas las Go Routines han terminado")

	// 5. ¿Qué son los Channels?
	fmt.Println("\n--- 5. ¿Qué son los Channels? ---")
	fmt.Println("Los Channels son conductos para comunicación entre Go Routines.")
	fmt.Println("Permiten enviar y recibir datos de manera segura.")
	fmt.Println("Se crean con: make(chan TipoDato)")

	// 6. Ejemplo básico de un Channel
	fmt.Println("\n--- 6. Ejemplo Básico de un Channel ---")
	canal := make(chan string)

	// Go Routine que envía un mensaje
	go func() {
		canal <- "Hola desde la Go Routine"
	}()

	// Recibir el mensaje del canal
	mensaje := <- canal
	fmt.Println("Mensaje recibido:", mensaje)

	// 7. Channel con múltiples valores
	fmt.Println("\n--- 7. Channel con Múltiples Valores ---")
	numeros := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			numeros <- i
			fmt.Printf("Enviado: %d\n", i)
		}
		close(numeros) // Cerrar el canal cuando termines
	}()

	// Recibir los valores del canal
	for valor := range numeros {
		fmt.Printf("Recibido: %d\n", valor)
	}

	// 8. Channels con búfer
	fmt.Println("\n--- 8. Channels con Búfer ---")
	fmt.Println("Un channel sin búfer espera hasta que alguien reciba el valor.")
	fmt.Println("Un channel con búfer puede almacenar valores hasta su capacidad.")

	canalBuffer := make(chan string, 2) // Canal con capacidad para 2 elementos

	// Enviar dos valores sin que nadie reciba
	canalBuffer <- "Valor 1"
	canalBuffer <- "Valor 2"

	fmt.Println("Recibido:", <-canalBuffer)
	fmt.Println("Recibido:", <-canalBuffer)

	// 9. Select: Esperar múltiples Channels
	fmt.Println("\n--- 9. Select: Esperar Múltiples Channels ---")
	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		canal1 <- "Mensaje del canal 1"
	}()

	go func() {
		canal2 <- "Mensaje del canal 2"
	}()

	// Select espera el primer mensaje disponible
	select {
	case msg1 := <-canal1:
		fmt.Println("Recibido:", msg1)
	case msg2 := <-canal2:
		fmt.Println("Recibido:", msg2)
	}

	// 10. Resumen: Patrón común - Productor/Consumidor
	fmt.Println("\n--- 10. Patrón: Productor/Consumidor ---")
	datos := make(chan int)

	// Productor
	go func() {
		for i := 1; i <= 3; i++ {
			datos <- i * 10
			time.Sleep(100 * time.Millisecond)
		}
		close(datos)
	}()

	// Consumidor
	for valor := range datos {
		fmt.Printf("Consumido: %d\n", valor)
	}

	// 11. Consideraciones Importantes
	fmt.Println("\n--- 11. Consideraciones Importantes ---")
	fmt.Println("• Usa 'go' para crear Go Routines de manera simple.")
	fmt.Println("• Usa WaitGroup para sincronizar Go Routines cuando sea necesario.")
	fmt.Println("• Los Channels permiten comunicación segura entre Go Routines.")
	fmt.Println("• Siempre cierra los Channels cuando termines de enviar datos.")
	fmt.Println("• Los Channels sin búfer son de sincronización directa.")
	fmt.Println("• Los Channels con búfer pueden almacenar valores.")

	fmt.Println("\n=== Fin de la Clase Básica de Go Routines y Channels ===")
}
