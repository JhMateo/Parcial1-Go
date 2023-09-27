package puntos

import (
	"fmt"
)

/*	Punto 2:
Definir un programa que dada una matriz n*m (ambos ingresados por el usuario,
nótese que ambos pueden ser valores diferentes) me genere y muestre una matriz en
caracol.
	Por ejemplo:
	 1  2  3  4  5
	16 17 18 19  6
	15 24 25 20  7
	14 23 22 21  8
	13 12 11 10  9	*/

func GenerarMatrizCaracol() {
	var n, m int

	for {
		fmt.Print("Por favor, ingresa el tamaño de la matriz n*m\n")
		fmt.Print("Ingrese n: ")
		_, err := fmt.Scan(&n)
		fmt.Print("Ingrese m: ")
		_, err2 := fmt.Scan(&m)

		if err != nil || err2 != nil {
			fmt.Println("Error al leer la entrada:", err, err2)
		} else {
			break
		}
	}

	GenMatrizCaracol(n, m)
}

func GenMatrizCaracol(n int, m int) {
	matriz := make([][]string, n)

	for i := range matriz {
		matriz[i] = make([]string, m)
	}

	filaInicio := 0
	filaFin := n - 1
	columnaInicio := 0
	columnaFin := m - 1
	contador := 1

	for contador <= n*m {
		for j := columnaInicio; j <= columnaFin; j++ {
			if matriz[filaInicio][j] == "" {
				matriz[filaInicio][j] = FormatearNumero(contador)
				contador++
			}
		}
		for i := filaInicio + 1; i <= filaFin; i++ {
			if matriz[i][columnaFin] == "" {
				matriz[i][columnaFin] = FormatearNumero(contador)
				contador++
			}
		}
		for j := columnaFin - 1; j >= columnaInicio; j-- {
			if matriz[filaFin][j] == "" {
				matriz[filaFin][j] = FormatearNumero(contador)
				contador++
			}
		}
		for i := filaFin - 1; i > filaInicio; i-- {
			if matriz[i][columnaInicio] == "" {
				matriz[i][columnaInicio] = FormatearNumero(contador)
				contador++
			}
		}

		filaInicio++
		filaFin--
		columnaInicio++
		columnaFin--
	}
	fmt.Println("> Matriz Caracol")
	ImprimirMatriz(matriz, n, m)
}
