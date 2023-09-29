package puntos

import (
	"fmt"
)

func GenerarMatrizCaracol() {
	var n, m int

	for {
		fmt.Print("Por favor, ingresa el tamaño de la matriz n*m\n")
		fmt.Print("Ingrese n: ")
		_, err := fmt.Scan(&n)
		fmt.Print("Ingrese m: ")
		_, err2 := fmt.Scan(&m)

		if err != nil || err2 != nil || n <= 0 || m <= 0 {
			fmt.Println("Error: Ingresa valores válidos para n y m (n y m deben ser números positivos).")
		} else {
			break
		}
	}

	matriz := GenMatrizCaracol(n, m)

	fmt.Println("> Matriz Caracol")
	ImprimirMatrizCaracol(matriz)
}

func GenMatrizCaracol(n int, m int) [][]int {
	matriz := make([][]int, n)

	for i := range matriz {
		matriz[i] = make([]int, m)
	}

	contador := 1
	direccion := "derecha"
	fila, columna := 0, 0

	for contador <= n*m {
		matriz[fila][columna] = contador
		contador++

		switch direccion {
		case "derecha":
			if columna+1 < m && matriz[fila][columna+1] == 0 {
				columna++
			} else {
				direccion = "abajo"
				fila++
			}
		case "abajo":
			if fila+1 < n && matriz[fila+1][columna] == 0 {
				fila++
			} else {
				direccion = "izquierda"
				columna--
			}
		case "izquierda":
			if columna-1 >= 0 && matriz[fila][columna-1] == 0 {
				columna--
			} else {
				direccion = "arriba"
				fila--
			}
		case "arriba":
			if fila-1 >= 0 && matriz[fila-1][columna] == 0 {
				fila--
			} else {
				direccion = "derecha"
				columna++
			}
		}
	}

	return matriz
}

func ImprimirMatrizCaracol(matriz [][]int) {
	maximo := 0
	for _, fila := range matriz {
		for _, valor := range fila {
			if valor > maximo {
				maximo = valor
			}
		}
	}

	width := len(fmt.Sprint(maximo))

	for _, fila := range matriz {
		for _, valor := range fila {
			fmt.Printf("%*d ", width, valor)
		}
		fmt.Println()
	}
}
