/* Integrantes:
Kevin Santiago Martinez: 160004321
Pablo Santiago Bobadilla: 160004331
Joshep Mateo Granada: 160004314	*/

package main

import (
	"fmt"
	"github.com/JhMateo/parcial1/puntos"
)

func main() {
	var opcion int

	for {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Generar Figuras Z y +")
		fmt.Println("2. Generar Matriz en Forma de Caracol")
		fmt.Println("3. Dibujar Texto")
		fmt.Println("4. Salir")

		fmt.Print("Ingrese su opción: ")
		_, err := fmt.Scan(&opcion)

		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			continue
		}

		switch opcion {
		case 1:
			fmt.Println("\n# Punto 1:")
			puntos.GenerarFigurasZMas()
		case 2:
			fmt.Println("\n# Punto 2:")
			puntos.GenerarMatrizCaracol()
		case 3:
			fmt.Println("\n# Punto 3:")
			puntos.DibujarTexto()
		case 4:
			fmt.Println("Saliendo del programa.")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.\n")
		}
	}
}
