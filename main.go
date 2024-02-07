package main

import (
	"godesde0/webserver"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1991)
	fmt.Println("Estado", estado)
	fmt.Println("Texto", texto)
	condicionIf()
	condicionSwitch()

	numero, texto := ejercicios.StringToNumber("50")
	fmt.Println("numero", numero)
	fmt.Println("texto", texto)
	teclado.IngresoNumeros()
	iteraciones.Iterar()
	fmt.Println(ejercicios.TablasMultiplicar())*/
	webserver.MiWebServer()
}

/*
func condicionIf() {
	//Permite realizar una asignacion, pero solo vive en el IF
	if os := runtime.GOOS; os == "linux" || os == "OS X" {
		fmt.Println("SO no es Windows", os)
	} else {
		fmt.Println("SO Es Windows", os)
	}
}


func condicionSwitch() {
	//Permite realizar una asignacion, pero solo vive en el switch
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "mac":
		fmt.Println("Esto es mac")
	default:
		fmt.Printf("El OS %s \n", os)
	}
}*/
