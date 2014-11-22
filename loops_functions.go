/* Exercise: Loops and Functions */
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	old_z := 0.0
	for math.Abs(old_z-z) > 1e-5 {
		old_z = z
		z = z - (z*z-x)/(2.0*z)
	}

	return z
}

func main() {
	sqrt := Sqrt(2)
	math_sqrt := math.Sqrt(2)
	fmt.Println(sqrt, math_sqrt, sqrt-math_sqrt)
}
