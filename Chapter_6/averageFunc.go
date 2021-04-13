package main

import "fmt"

func average(xs ...float64) (av float64) {
	av = 0.0
	for _, value := range xs {
		av += value
	}

	return av / float64(len(xs))
}

func main() {
	x := []float64{34.5, 1.4, 23.4, 4.1, 100.1}

	fmt.Println(average(x...))
}
