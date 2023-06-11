package ejercicios

import (
	"strconv"
)

func StringtoInt(numero string) (int, string) {

	switch numeroint, _ := strconv.Atoi(numero); {
	case numeroint > 100:
		return numeroint, "Es mayor a 100"
	default:
		return numeroint, "Es menor a 100"
	}

}
