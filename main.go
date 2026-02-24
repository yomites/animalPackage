package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x
	//iteration := 0
	//previous_z := 0.0
	/**
	for n := 0; n < 100; n++ {
		z -= (z*z - x) / (2*z)

		if z - previous_z == 0.0 {
			iteration++
			break
		}

		//fmt.Printf("z value %v \n", z)
		//fmt.Printf("Previous z value %v \n", previous_z)
		//fmt.Println(n+1)
		//fmt.Println("---------------------------------------")
		previous_z = z


	}
	**/

	for z*z > x {
		z -= (z*z - x) / (2 * z)

		fmt.Printf("z value %v \n", z)
		//fmt.Printf("Previous z value %v \n", previous_z)
		//fmt.Println(n+1)
		//fmt.Println("---------------------------------------")
		//previous_z = z
	}
	fmt.Printf("z value %v \n", z)
	//fmt.Printf("Previous z value %v \n", previous_z)
	//fmt.Println("---------------------------------------")
	//fmt.Printf("Number of iteration is %v \n", iteration)
	return z
}

func main() {
	number := 89998988.0
	fmt.Printf("The square root of %v is %v \n", number, Sqrt(number))
}
