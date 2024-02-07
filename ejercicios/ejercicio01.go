package ejercicios

import (
	"strconv"
)

func StringToNumber(texto string) (int, string) {
	//Si quiero ignorar el error que pueda traer la func utilizo "_"
	n, _ := strconv.Atoi(texto)
	if n > 100 {
		return n, "Es mayor a 100"
	} else {
		return n, "Es menor a 100"
	}
}
