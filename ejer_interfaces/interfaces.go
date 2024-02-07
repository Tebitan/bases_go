package ejer_interfaces

import (
	"fmt"
	"godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy de sexo %s y estoy respirando \n", hu.Sexo())
}
