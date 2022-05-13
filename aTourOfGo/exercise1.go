// Execrise1
// https://go-tour-jp.appspot.com/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var z_old float64
	for math.Abs(z-z_old) > 0.00001 {
		z_old = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
}
