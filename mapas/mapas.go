package mapas

import "fmt"

func MostrarMapas() {
	//usa la estructura KEY-VALUE
	paises := make(map[string]string)
	//add map
	paises["Colombia"] = "Bogota"
	paises["Venezuela"] = "Caracas"
	paises["Chile"] = "Santiago"
	paises["Ecuador"] = "Quito"
	paises["Peru"] = "Lima"

	fmt.Println(paises)
	//Me muestra el valor
	fmt.Println(paises["Colombia"])
	//creacion por asignacion directa
	liga := map[string]int{
		"Atl.Bucaramanga": 10,
		"Atl.Nacional":    9,
		"America":         8,
		"Cucuta Dep":      7,
	}
	fmt.Println(liga)
	//quitar elemento
	delete(liga, "Cucuta Dep")
	//recorrer map
	for equipo, puntos := range liga {
		fmt.Printf("Equipo: %s, Puntos: %d \n", equipo, puntos)
	}
	//buscar por KEY
	value, exist := liga["Juventus"]
	fmt.Printf("El puntaje capturado es: %d, y el equipo existe %t \n", value, exist)

}
