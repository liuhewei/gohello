package main

import (
	"fmt"
)

const delta float64 = 0.00000000001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

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
	return z, nil
}

func main() {
	x, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}
