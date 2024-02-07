package files

import (
	"bufio"
	"fmt"
	"godesde0/ejercicios"
	"os"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablasMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creando el archivo" + err.Error())
		//sale del metodo
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func ConcatenaTablaArchivo() {
	var texto string = ejercicios.TablasMultiplicar()
	if !AppendTextFile(fileName, texto) {
		fmt.Println("Error al concatenar el archivo")
	}
}

func AppendTextFile(filen string, texto string) bool {
	//uso | para enviar varias opciones
	arch, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante en el AppendText" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante en el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	//os.Open solo es de lectura
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error leyendo el archivo" + err.Error())
		return
	}
	//lectura del archivo
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println("> " + linea)
	}
	archivo.Close()
}
