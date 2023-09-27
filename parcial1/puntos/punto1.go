package puntos

import (
	"fmt"
)

/*	Punto 1:
Definir un programa que dada una matriz n*n (ambos ingresados por el
usuario) me genere y muestre los valores formando las siguientes figuras:
	a. “Z”
	b. “+”	*/

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
	columnaZ := n - 2

	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			for j := 0; j < n; j++ {
				if contadorZ < 10 {
					matrizZ[i][j] = FormatearNumero(contadorZ)
				} else {
					matrizZ[i][j] = FormatearNumero(contadorZ)
				}
				contadorZ++
			}
		} else {
			if contadorZ < 10 {
				matrizZ[i][columnaZ] = FormatearNumero(contadorZ)
			} else {
				matrizZ[i][columnaZ] = FormatearNumero(contadorZ)
			}
			contadorZ++
			columnaZ--
		}
	}

	fmt.Println("> Figura Z")
	ImprimirMatriz(matrizZ, n, n)
}

func GenerarFiguraMas(n int) {
	matrizMas := crearMatrizCuadrada(n)

	contadorMas := 1
	var filaColumnaMas int = n / 2

	for i := 0; i < n; i++ {
		if i == filaColumnaMas {
			for j := 0; j < n; j++ {
				if contadorMas < 10 {
					matrizMas[i][j] = FormatearNumero(contadorMas)
				} else {
					matrizMas[i][j] = FormatearNumero(contadorMas)
				}
				contadorMas++
			}
		} else {
			if contadorMas < 10 {
				matrizMas[i][filaColumnaMas] = FormatearNumero(contadorMas)
			} else {
				matrizMas[i][filaColumnaMas] = FormatearNumero(contadorMas)
			}
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
