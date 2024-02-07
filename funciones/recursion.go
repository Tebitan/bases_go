package funciones

import "fmt"

//funcion que se llama asi misma N veces
func Exponencia(valor int) {
	//se define cuando se termina
	if valor > 1000 {
		return
	}
	fmt.Println(valor)
	Exponencia(valor * 2)
}
