package modelos

//Esto es Herencias
type Mujer struct {
	Hombre
}

//hereda todos sus atributos NO sus propiedades
func (mujer *Mujer) Respirar()    { mujer.respirando = true }
func (mujer *Mujer) Pensar()      { mujer.pensando = true }
func (mujer *Mujer) Comer()       { mujer.comiendo = true }
func (mujer *Mujer) Sexo() string { return "Femenino" }
