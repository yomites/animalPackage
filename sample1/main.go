package main

import (
	"fmt"
	"time"
)

func Sqrt(x float64) float64 {
	z := x
	previous_z := 0.0

	for n := 0; n < 100; n++ {
		z -= (z*z - x) / (2 * z)

		fmt.Printf("z value %v \n", z)
		fmt.Printf("Previous z value %v \n", previous_z)
		fmt.Println(n + 1)
		fmt.Println("---------------------------------------")

		if z-previous_z == 0.0 {
			break
		}

		previous_z = z
	}

	fmt.Printf("z value %v \n", z)
	fmt.Printf("Previous z value %v \n", previous_z)
	fmt.Println("---------------------------------------")
	return z
}

func main() {
	start := time.Now()
	number := 4533.45
	fmt.Printf("The square root of %v is %v \n", number, Sqrt(number))

	fmt.Printf("The call took %v to run.\n", time.Since(start))
}
