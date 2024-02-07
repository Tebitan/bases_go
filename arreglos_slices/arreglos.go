package arreglos_slices

import "fmt"

//Tener en cuenta las posiciones que no son asignadas el las iniciara EJ: INT -> 0 , STRING ""
var arregloInt [10]int = [10]int{1, 2, 3}
var matrizInt [20][30]int

func MuestraArreglos() {
	//asignar por posicion
	arregloInt[3] = 5
	arregloInt[4] = 6
	//creacion por asignacion
	arregloString := [10]string{"pablo", "pedro", "juan"}

	matrizInt[1][1] = 2
	matrizInt[1][2] = 3

	fmt.Println(arregloInt)
	fmt.Println(arregloString)
	fmt.Println(matrizInt)

	//recorrer un arreglo
	for i := 0; i < len(arregloInt); i++ {
		fmt.Println(arregloInt[i])
	}

}
