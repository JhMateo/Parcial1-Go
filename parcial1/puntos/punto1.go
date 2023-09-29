package puntos

import (
	"fmt"
)

/*
GenerarFigurasZMas generates and displays two figures, "Z" and "+", formed by the values of an n*n matrix.
a. The "Z" figure is formed by numbers starting from 1 in the top and bottom rows, and a diagonal line connecting them.
b. The "+" figure is formed by numbers starting from 1 in the middle row and column, horizontally and vertically.

Parameters:
- n: The size of the n*n matrix entered by the user.

Functions:
- GenerarFigurasZMas(): Reads the size of the matrix from the user and generates both "Z" and "+" figures.
- GenerarFiguraZ(n int): Generates the "Z" figure for the given size n and prints it.
- GenerarFiguraMas(n int): Generates the "+" figure for the given size n and prints it.
- crearMatrizCuadrada(n int) [][]string: Creates an empty square matrix of size n*n.
- ImprimirMatriz(matriz [][]string, n int, m int): Prints the given matrix of size n*m.
- FormatearNumero(numero int) string: Formats the given number as a string with leading space if it is less than 10.
*/

func GenerarFigurasZMas() {
	var n int

	for {
		fmt.Print("Por favor, ingresa el tamaño de la matriz n*n: ")
		_, err := fmt.Scan(&n)

		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
		} else if n < 3 {
			fmt.Println("El tamaño de la matriz debe ser mayor o igual a tres.")
		} else {
			break
		}
	}

	GenerarFiguraZ(n)
	GenerarFiguraMas(n)
}

func GenerarFiguraZ(n int) {
	matrizZ := crearMatrizCuadrada(n)

	contadorZ := 1

	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			for j := 0; j < n; j++ {
				matrizZ[i][j] = FormatearNumero(contadorZ)
				contadorZ++
			}
		} else {
			matrizZ[i][n-i-1] = FormatearNumero(contadorZ)
			contadorZ++
		}
	}

	fmt.Println("> Figura Z")
	ImprimirMatriz(matrizZ, n, n)
}

func GenerarFiguraMas(n int) {
	matrizMas := crearMatrizCuadrada(n)

	contadorMas := 1
	filaColumnaMas := n / 2

	for i := 0; i < n; i++ {
		if i == filaColumnaMas {
			for j := 0; j < n; j++ {
				matrizMas[i][j] = FormatearNumero(contadorMas)
				contadorMas++
			}
		} else {
			matrizMas[i][filaColumnaMas] = FormatearNumero(contadorMas)
			contadorMas++
		}
	}

	fmt.Println("> Figura +")
	ImprimirMatriz(matrizMas, n, n)
}

func crearMatrizCuadrada(n int) [][]string {
	matriz := make([][]string, n)

	// Inicializar la matriz con campos vacíos
	for i := range matriz {
		matriz[i] = make([]string, n)
		for j := range matriz[i] {
			matriz[i][j] = "  "
		}
	}

	return matriz
}

func ImprimirMatriz(matriz [][]string, n int, m int) {
	fmt.Printf("Matriz de tamaño %v*%v: \n", n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%v ", matriz[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func FormatearNumero(numero int) string {
	if numero < 10 {
		return fmt.Sprintf(" %d", numero)
	}
	return fmt.Sprintf("%d", numero)
}
