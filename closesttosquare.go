package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var before = z
	for {
		z -= (z*z - x) / (2 * z)
		if before == z {
			return z
		}
		before = z
	}
}

func main() {
	fmt.Println(Sqrt(10))
}
