package users

import (
	//"fmt"
	"fmt"
	"godesde0/modelos"
	"time"
)

func CrearUsuario() {
	//Instancio el Model
	user := new(modelos.User)
	//LLamo la func que actua como constructor
	user.UserConstructor(1, "Esteban", time.Now(), true)
	fmt.Println(user)
}
