package main

import "fmt"

func fibb(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	}

	return fibb(i-1) + fibb(i-2)
}

func main() {

	fmt.Println(fibb(0))
	fmt.Println(fibb(1))
	fmt.Println(fibb(2))
	fmt.Println(fibb(9))
}
