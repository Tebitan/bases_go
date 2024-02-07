package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablasMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string
	for {
		fmt.Println("Ingrese número a Multiplicar : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		//Sprintf concatena el string segun el format
		texto += fmt.Sprintf("%d X %d = %d \n", numero, i, numero*i)
	}
	return texto
}
