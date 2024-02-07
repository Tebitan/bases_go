package variables

import (
	"fmt"
	"strconv"
	"time"
)

// Si las variables Empiezan con Mayuscula son Globales
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Esteban Navas"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println("Nombre", Nombre)
	fmt.Println("Estado", Estado)
	fmt.Println("Sueldo", Sueldo)
	fmt.Println("Fecha", Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	var texto string = strconv.Itoa(numero)
	return true, texto
}
