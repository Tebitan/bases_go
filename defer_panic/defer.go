package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es el Mensaje numero 1")
	//funciona como se fuera finaly es decir esto se ejecuta el final de todo
	defer fmt.Println("Este es el Mensaje final")

	fmt.Println("Este es el Mensaje numero 2")
}

func EjemploPanic() {
	//ejemplo de funcion anonima con defer
	defer func() {
		//recover -> se estar acompañado por (defer, func anonima)
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		//aborta la ejecucion y muestra el mensaje
		panic("Se encontro el valor 1")
	}
}
