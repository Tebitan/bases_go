package funciones

import "fmt"

//func return func
func tabla(valor int) func() int {
	//zona de variables
	numero := valor
	secuencia := 0
	//lo que importa es lo que se esta retornando
	return func() int {
		//logica
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	tablaNumber := 2
	//variable asignada a la funcion
	miTabla := tabla(tablaNumber)
	for i := 1; i <= 10; i++ {
		fmt.Println(miTabla())
	}
}
