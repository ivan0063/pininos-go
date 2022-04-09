package main

import (
	"fmt"
	"math"
)

func cuadratica() (float64, float64) {
	var a, b, c int
	fmt.Println("Dame los valores de tu ecuacion en la forma:")
	fmt.Println("ax2 +/- by +/- c")

	fmt.Println("a: ")
	fmt.Scan(&a)
	fmt.Println("b: ")
	fmt.Scan(&b)
	fmt.Println("c: ")
	fmt.Scan(&c)

	coeficienteRaiz := (b * b) - (4 * a * c)
	negativeB := float64(b * -1)

	prevSqrt := float64(coeficienteRaiz)
	sqrt := math.Sqrt(prevSqrt)
	abajo := float64(2 * a)

	/*fmt.Printf("coeficienteRaiz=%f\n", coeficienteRaiz)
	fmt.Printf("prevSqrt=%f\n", prevSqrt)
	fmt.Printf("negativeB=%f\n", negativeB)
	fmt.Printf("cuadrado=%f\n", sqrt)
	fmt.Printf("abajo=%f\n", abajo)*/

	x := (negativeB + sqrt) / abajo
	y := (negativeB - sqrt) / abajo

	return x, y
}

func main() {
	// assabling reader
	//reader := bufio.NewReader(os.Stdin)
	var option int

	fmt.Println("Que quires hacer?")
	fmt.Println("[1] Ecuaacion Cuadratica")
	fmt.Scan(&option)

	switch option {
	case 1:
		resultA, resultB := cuadratica()
		fmt.Printf("Reusltados son x=%f , y=%f\n", resultA, resultB)
		break
	}
}
