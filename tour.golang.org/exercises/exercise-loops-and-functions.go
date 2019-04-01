package main

import (
	"fmt"
)

const delta float64 = 0.00000000001

func Sqrt(x float64) float64 {
	z := x / 2
	var t float64
	for {
		fmt.Println(z, "\n")

		t = z
		z -= (z*z - x) / (2 * z)
		if t-z >= -delta && t-z <= delta {
			break
		}

	}
	return z
}

func main() {
	x := 3.0
	fmt.Println(Sqrt(x))
}
