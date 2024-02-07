package funciones

import "fmt"

func Calculos() {
	var numero3 int = 32
	var numero4 int = 64
	//puede acceder a las variables de func padre
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}
	fmt.Println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		//puedo cambiar su logica, lo que no puedo cambiar es el tipo de parametros y su cantidad
		return numero1 * numero2 * numero3 * numero4
	}
	fmt.Println(calculo(10, 25))
}
