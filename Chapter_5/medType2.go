package main

import "fmt"

func main() {
	x := [5]float64{34.5, 1.4, 400.1, 4.1, 100}

	var y float64
	for _, i := range x {
		y += i
	}

	fmt.Println(y / float64(len(x)))
}
