package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var max int = 10

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Introduce un numbero del 1 al 10: ")
		if scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error, no es un número.")
				continue
			}
			if number > 1 && number <= 10 {
				for i := 0; i <= max; i++ {
					result := number * i
					fmt.Printf("%d x %d = %d \n", number, i, result)
				}

			} else {
				fmt.Println("Solo imprimo las tablas del 1 al 10")
				continue
			}
			break
		}

	}
}
