package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese número 1 : ")
	if scanner.Scan() {
		// se utiliza := cuando las variables NO se han creado previamente
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			//aborta el programa
			panic("El dato ingresado no es correcto " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2 : ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado no es correcto " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
