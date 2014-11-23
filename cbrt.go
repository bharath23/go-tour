/* Advanced Exercise: Complex cube roots */
package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := x
	old_z := complex128(0)
	for cmplx.Abs(z-old_z) > 1e-5 {
		old_z = z
		z = z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))
	}

	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
