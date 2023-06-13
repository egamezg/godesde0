package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var max int = 10
var number int
var err error

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Introduce un numbero del 1 al 10: ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error, no es un número.")
				continue
			} else {
				break
			}
		}
	}
	for i := 0; i <= max; i++ {
		result := number * i
		fmt.Printf("%d x %d = %d \n", number, i, result)
	}

}
