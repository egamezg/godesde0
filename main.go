package main

import (
	"fmt"
	"github.com/egamezg/godesde0/ejercicios"
)

func main() {

	/*
		fmt.Println(variables.DevuelveEntera())

		variables.RestoVariables()

		entero := 25
		enteronuevo := variables.InttoInt32(entero)
		fmt.Printf("El entero %T convertido a int32 es %T \n", entero, enteronuevo)

		estado, texto := variables.ConviertoaTexto(5)

		fmt.Println("Estado:", estado)
		fmt.Printf("Texto: es de tipo %T\n", texto)

		if os := runtime.GOOS; os == "darwin" {
			println("This is a Mac")
		}

		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Printf("This is %s \n", os)
		default:
			fmt.Println("This is not a Mac")
		}
	*/

	conversion, mensaje := ejercicios.StringtoInt("125")
	fmt.Printf("%d ", conversion)
	fmt.Println(mensaje)
}
