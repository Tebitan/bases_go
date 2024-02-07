package arreglos_slices

import "fmt"

//No se define su tamaño
var sliceSimple []int = []int{2, 5, 6}
var arregloInt2 [10]int = [10]int{6, 5, 4, 8, 9, 11, 23, 25, 15}

func MuestraSlice() {
	//Slice creado desde un vector , desde la posicion 3
	sPorcion1 := arregloInt2[3:]
	//Slice creado desde un vector , desde la posicion 0 hasta la 5
	sPorcion2 := arregloInt2[:5]
	//Slice creado desde un vector , desde la posicion 6 hasta la 7
	sPorcion3 := arregloInt2[6:7]

	fmt.Println(sliceSimple)
	fmt.Println(sPorcion1)
	fmt.Println(sPorcion2)
	fmt.Println(sPorcion3)
}

func Capacidad() {
	//capacidad es el espacio en memoria que en este caso es 20
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo: %d, Capacidad: %d \n", len(elementos), cap(elementos))
	//declaramos el slice sin dimencion y sin capacidad
	sliceVacio := make([]int, 0)
	for i := 0; i < 10; i++ {
		//adicionamos item al slice
		sliceVacio = append(sliceVacio, i)
	}
	fmt.Printf("Largo: %d, Capacidad: %d \n", len(sliceVacio), cap(sliceVacio))

}
