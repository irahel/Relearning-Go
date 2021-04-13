package main

import "fmt"

func maxSeq(seq ...int) (max int) {
	for index, value := range seq {
		if index == 0 {
			max = value
		} else {
			if value > max {
				max = value
			}
		}
	}
	return
}

func main() {

	fmt.Println(maxSeq(1, 2, 30, 40, 6, 9, 50, 99, 69, 1, 3))
}
