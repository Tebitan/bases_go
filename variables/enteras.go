package variables

import "fmt"

//Si el Metodo empieza con Mayuscula es Publico
func MostrarEnteros() {
	//por declaracion
	var intcomun int
	// por asignacion
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun", intcomun)
	fmt.Println("intde32", intde32)
	fmt.Println("intde64", intde64)
}
