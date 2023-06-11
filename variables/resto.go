package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Eduardo"
	Estado = true
	Sueldo = 1000.50
	Fecha = time.Now()

	fmt.Println("Nombre:", Nombre)
	fmt.Println("Estado:", Estado)
	fmt.Println("Sueldo:", Sueldo)
	fmt.Println("Fecha:", Fecha.Day(), Fecha.Month(), Fecha.Year())
}

func ConviertoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto
}
