package goroutines

import (
	"fmt"
	"strings"
	"time"
)

// para llamarlo de forma asincrona go MiNombreLentooo("Esteban Navas")
func MiNombreLentooo(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1 * time.Microsecond)
		fmt.Println(letra)
	}
	//para asignar valor al canal
	canal1 <- true
}

func CallAsinc() {
	canal1 := make(chan bool)
	go MiNombreLentooo("funcion asincrona", canal1)
	fmt.Println("Inicio del proceso")
	//esto es un await es decir va esperar hasta que la funcion termine
	<-canal1
	/* para asegurarce que las funciones asincronar se ejecuten
	defer func() {
		<-canal1
	}()*/
}
