package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 34.5
	x[1] = 1.4
	x[2] = 432.3
	x[3] = 4.1
	x[4] = 100

	var y float64
	for _, i := range x {
		y += i
	}

	fmt.Println(y / float64(len(x)))
}
