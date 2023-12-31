package puntos

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/*
DibujarTexto draws the given text using matrix representations of characters. It transforms the input text based on predefined rules,
scales the characters to fit the specified font size, and prints the scaled characters with a specified delay between characters.

Functions:
- DibujarTexto(): Reads the text, delay, and font size from the user and draws the text.
- transformarTexto(texto string) string: Transforms the input text using predefined rules.
- escalarRepresentacion(repOriginal []string, n int) ([]string, error): Scales the original character representation to fit the specified font size.
- imprimirTextoEnMatriz(text string,  retrasoSegundos time.Duration, n int): Prints the transformed and scaled text in a matrix format.
*/

var reglas = map[string]string{
	"A":        "4",
	"E":        "3",
	"I":        "1",
	"O":        "0",
	"U":        "#",
	"4":        "A",
	"3":        "E",
	"1":        "I",
	"0":        "O",
	"#":        "U",
	"AMOR":     "❤",
	"MÚSICA":   "♫",
	"FELIZ":    "☺",
	"BALANCE":  "☯",
	"PELIGRO":  "☠",
	"ATOMICO":  "☢",
	"ESTRELLA": "★",
	"INF":      "∞",
}

// Definir un mapa para almacenar las representaciones en matriz de los caracteres
var matrizCaracteres = map[string][]string{
	"A": {
		"  *  ",
		" * * ",
		"*   *",
		"*****",
		"*   *",
	},
	"B": {
		"**** ",
		"*   *",
		"**** ",
		"*   *",
		"**** ",
	},
	"C": {
		" ****",
		"*    ",
		"*    ",
		"*    ",
		" ****",
	},
	"D": {
		"**** ",
		"*   *",
		"*   *",
		"*   *",
		"**** ",
	},
	"E": {
		"*****",
		"*    ",
		"**** ",
		"*    ",
		"*****",
	},
	"F": {
		"*****",
		"*    ",
		"**** ",
		"*    ",
		"*    ",
	},
	"G": {
		" ****",
		"*    ",
		"*  **",
		"*   *",
		" *** ",
	},
	"H": {
		"*   *",
		"*   *",
		"*****",
		"*   *",
		"*   *",
	},
	"I": {
		"*****",
		"  *  ",
		"  *  ",
		"  *  ",
		"*****",
	},
	"J": {
		"  ***",
		"    *",
		"    *",
		"*   *",
		" *** ",
	},
	"K": {
		"*   *",
		"*  * ",
		"**   ",
		"*  * ",
		"*   *",
	},
	"L": {
		"*    ",
		"*    ",
		"*    ",
		"*    ",
		"*****",
	},
	"M": {
		"*   *",
		"** **",
		"* * *",
		"*   *",
		"*   *",
	},
	"N": {
		"*   *",
		"**  *",
		"* * *",
		"*  **",
		"*   *",
	},
	"O": {
		" *** ",
		"*   *",
		"*   *",
		"*   *",
		" *** ",
	},
	"P": {
		"**** ",
		"*   *",
		"**** ",
		"*    ",
		"*    ",
	},
	"Q": {
		" *** ",
		"*   *",
		"*   *",
		"*  * ",
		" ** *",
	},
	"R": {
		"**** ",
		"*   *",
		"**** ",
		"*  * ",
		"*   *",
	},
	"S": {
		" ****",
		"*    ",
		" *** ",
		"    *",
		"**** ",
	},
	"T": {
		"*****",
		"  *  ",
		"  *  ",
		"  *  ",
		"  *  ",
	},
	"U": {
		"*   *",
		"*   *",
		"*   *",
		"*   *",
		" *** ",
	},
	"V": {
		"*   *",
		"*   *",
		"*   *",
		" * * ",
		"  *  ",
	},
	"W": {
		"*   *",
		"*   *",
		"* * *",
		"** **",
		"*   *",
	},
	"X": {
		"*   *",
		" * * ",
		"  *  ",
		" * * ",
		"*   *",
	},
	"Y": {
		"*   *",
		" * * ",
		"  *  ",
		"  *  ",
		"  *  ",
	},
	"Z": {
		"*****",
		"   * ",
		"  *  ",
		" *   ",
		"*****",
	},
	"0": {
		" *** ",
		"*  **",
		"* * *",
		"**  *",
		" *** ",
	},
	"1": {
		"  *  ",
		" **  ",
		"  *  ",
		"  *  ",
		"*****",
	},
	"2": {
		" ****",
		"*   *",
		"   * ",
		"  *  ",
		"*****",
	},
	"3": {
		"*****",
		"    *",
		" ****",
		"    *",
		"*****",
	},
	"4": {
		"  ** ",
		" * * ",
		"*  * ",
		"*****",
		"   * ",
	},
	"5": {
		"*****",
		"*    ",
		"**** ",
		"    *",
		"**** ",
	},
	"6": {
		" ****",
		"*    ",
		"**** ",
		"*   *",
		" ****",
	},
	"7": {
		"*****",
		"   * ",
		"  *  ",
		" *   ",
		"*    ",
	},
	"8": {
		" ****",
		"*   *",
		" ****",
		"*   *",
		" ****",
	},
	"9": {
		" ****",
		"*   *",
		" ****",
		"    *",
		"**** ",
	},
	"#": {
		" * * ",
		"** **",
		" * *",
		"** **",
		" * * ",
	},
	"❤": {
		" * * ",
		"* * *",
		"*   *",
		" * * ",
		"  *  ",
	},
	"♫": {
		" ****",
		" *  *",
		" *  *",
		"** **",
		"** **",
	},
	"☺": {
		"     ",
		"*   *",
		"     ",
		"*   *",
		" *** ",
	},
	"☯": {
		" *** ",
		"* * *",
		"**  *",
		"*****",
		" *** ",
	},
	"☠": {
		"*   *",
		" * *",
		"*   *",
		" * * ",
		"*   *",
	},
	"☢": {
		" *** ",
		"*   *",
		"* * *",
		"*   *",
		" *** ",
	},
	"★": {
		"  *  ",
		"  *  ",
		"*****",
		" * * ",
		"*   *",
	},
	"∞": {
		"     ",
		" * * ",
		"* * *",
		" * * ",
		"     ",
	},
}

func DibujarTexto() {
	var texto string
	var t int
	var n int

	// Validación del texto
	for {
		fmt.Print("Ingrese un texto: ")
		_, err := fmt.Scan(&texto)
		if err != nil {
			fmt.Println("Error al leer el texto. Inténtelo de nuevo.")
			continue
		}
		break
	}

	// Validación del tiempo de retraso
	for {
		fmt.Print("Ingrese el tiempo de retraso en segundos: ")
		_, err := fmt.Scan(&t)
		if err != nil || t <= 0 {
			fmt.Println("Tiempo de retraso inválido. Debe ser un número positivo mayor que cero.")
			continue
		}
		break
	}

	// Validación del tamaño de fuente
	for {
		fmt.Print("Ingrese el tamaño de fuente (número impar y mayor o igual a 5): ")
		_, err := fmt.Scan(&n)
		if err != nil || n < 5 || n%2 == 0 {
			fmt.Println("Tamaño de fuente inválido. Debe ser un número impar y mayor o igual a 5.")
			continue
		}
		break
	}

	// Convierte la cadena a mayúsculas
	textoMayusculas := strings.ToUpper(texto)

	textoTransformado := transformarTexto(textoMayusculas)

	imprimirTextoEnMatriz(textoTransformado, time.Duration(t), n)
}

func transformarTexto(texto string) string {
	if _, encontrado := reglas[texto]; encontrado {
		// La entrada es una palabra completa, devolver la transformación
		return reglas[texto]
	}
	// La entrada no es una palabra completa, aplicar transformación a vocales y números
	textoTransformado := ""
	for _, char := range texto {
		if nuevoChar, encontrado := reglas[string(char)]; encontrado {
			textoTransformado += nuevoChar
		} else {
			textoTransformado += string(char)
		}
	}
	return textoTransformado
}

func escalarRepresentacion(repOriginal []string, n int) ([]string, error) {
	if n%2 == 0 {
		return nil, errors.New("n debe ser un número impar")
	}

	m := len(repOriginal)
	if m%2 == 0 {
		return nil, errors.New("la representación original debe tener un tamaño impar")
	}

	escala := n / m
	var nuevaRepresentacion []string

	for _, fila := range repOriginal {
		nuevaFila := ""
		for _, caracter := range fila {
			nuevaFila += strings.Repeat(string(caracter), escala)
		}
		for i := 0; i < n/m; i++ {
			nuevaRepresentacion = append(nuevaRepresentacion, nuevaFila)
		}
	}

	return nuevaRepresentacion, nil
}

func imprimirTextoEnMatriz(texto string, retrasoSegundos time.Duration, n int) {
	for _, char := range texto {
		if representacion, encontrado := matrizCaracteres[string(char)]; encontrado {
			textoEscalado, err := escalarRepresentacion(representacion, n)
			if err != nil {
				fmt.Println("Error al escalar la representación:", err)
				return
			}
			for _, linea := range textoEscalado {
				fmt.Println(linea)
			}
		} else {
			fmt.Println("Carácter no reconocido:", string(char))
		}
		fmt.Println("")
		time.Sleep(retrasoSegundos * time.Second)
	}
}
