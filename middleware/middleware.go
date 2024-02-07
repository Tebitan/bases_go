package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

//definicion:Es un interceptor que se ejecuta antes de nuestra funcion
func MiMiddleware() {
	fmt.Println("Inicio")
	result := operacionMidd(sumar)(5, 6)
	fmt.Println(result)
	result = operacionMidd(restar)(10, 2)
	fmt.Println(result)
	result = operacionMidd(multiplicar)(7, 1)
	fmt.Println(result)
}

//Deben retornar el mismo tipo que ingresa es decir return int
func operacionMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio Operacion")
		return f(a, b)
	}
}
